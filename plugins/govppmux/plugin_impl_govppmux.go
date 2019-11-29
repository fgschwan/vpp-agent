//  Copyright (c) 2019 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package govppmux

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"git.fd.io/govpp.git/adapter"
	govppapi "git.fd.io/govpp.git/api"
	govpp "git.fd.io/govpp.git/core"
	"github.com/ligato/cn-infra/datasync/resync"
	"github.com/ligato/cn-infra/health/statuscheck"
	"github.com/ligato/cn-infra/infra"
	"github.com/ligato/cn-infra/logging"
	"github.com/ligato/cn-infra/rpc/rest"
	"github.com/pkg/errors"

	"go.ligato.io/vpp-agent/v2/plugins/govppmux/vppcalls"

	_ "go.ligato.io/vpp-agent/v2/plugins/govppmux/vppcalls/vpp1904"
	_ "go.ligato.io/vpp-agent/v2/plugins/govppmux/vppcalls/vpp1908"
	_ "go.ligato.io/vpp-agent/v2/plugins/govppmux/vppcalls/vpp2001"
	_ "go.ligato.io/vpp-agent/v2/plugins/govppmux/vppcalls/vpp2001_324"
)

var (
	disabledSocketClient = os.Getenv("GOVPPMUX_NOSOCK") != ""
)

// Plugin is the govppmux plugin implementation.
type Plugin struct {
	Deps

	config *Config

	vppAdapter  adapter.VppAPI
	vppConn     *govpp.Connection
	vppConChan  chan govpp.ConnectionEvent
	lastConnErr error
	vppapiChan  govppapi.Channel

	statsAdapter adapter.StatsAPI
	statsConn    govppapi.StatsProvider

	// infoMu synchonizes access to fields
	// vppInfo and lastEvent
	infoMu    sync.Mutex
	vppInfo   VPPInfo
	lastEvent govpp.ConnectionEvent

	cancel context.CancelFunc
	wg     sync.WaitGroup
}

// Deps defines dependencies for the govppmux plugin.
type Deps struct {
	infra.PluginDeps
	HTTPHandlers rest.HTTPHandlers
	StatusCheck  statuscheck.PluginStatusWriter
	Resync       *resync.Plugin
}

// Init is the entry point called by Agent Core. A single binary-API connection to VPP is established.
func (p *Plugin) Init() (err error) {
	if p.config, err = p.loadConfig(); err != nil {
		return err
	}

	p.Log.Debugf("config: %+v", p.config)

	govpp.HealthCheckProbeInterval = p.config.HealthCheckProbeInterval
	govpp.HealthCheckReplyTimeout = p.config.HealthCheckReplyTimeout
	govpp.HealthCheckThreshold = p.config.HealthCheckThreshold
	govpp.DefaultReplyTimeout = p.config.ReplyTimeout

	// register REST API handlers
	p.registerHandlers(p.HTTPHandlers)

	if p.vppAdapter == nil {
		var address string
		useShm := disabledSocketClient || p.config.ConnectViaShm || p.config.ShmPrefix != ""
		if useShm {
			address = p.config.ShmPrefix
		} else {
			address = p.config.BinAPISocketPath
		}
		p.vppAdapter = NewVppAdapter(address, useShm)
	} else {
		// this is used for testing purposes
		p.Log.Info("Reusing existing vppAdapter")
	}

	// TODO: Async connect & automatic reconnect support is not yet implemented in the agent,
	// so synchronously wait until connected to VPP.
	startTime := time.Now()
	p.Log.Debugf("connecting to VPP..")

	p.vppConn, p.vppConChan, err = govpp.AsyncConnect(p.vppAdapter, p.config.RetryConnectCount, p.config.RetryConnectTimeout)
	if err != nil {
		return err
	}

	// wait for connection event
	for {
		event := <-p.vppConChan
		if event.State == govpp.Connected {
			break
		} else if event.State == govpp.Failed || event.State == govpp.Disconnected {
			return errors.Errorf("unable to establish connection to VPP (%v)", event.Error)
		} else {
			p.Log.Debugf("VPP connection state: %+v", event)
		}
	}

	connectDur := time.Since(startTime)
	p.Log.Debugf("connection to VPP established (took %s)", connectDur.Round(time.Millisecond))

	if err := p.updateVPPInfo(); err != nil {
		return errors.WithMessage(err, "retrieving VPP info failed")
	}

	// Connect to VPP status socket
	var statsSocket string
	if p.config.StatsSocketPath != "" {
		statsSocket = p.config.StatsSocketPath
	} else {
		statsSocket = adapter.DefaultStatsSocket
	}
	statsAdapter := NewStatsAdapter(statsSocket)
	if statsAdapter == nil {
		p.Log.Warnf("Unable to connect to the VPP statistics socket, nil stats adapter", err)
	} else if p.statsConn, err = govpp.ConnectStats(statsAdapter); err != nil {
		p.Log.Warnf("Unable to connect to the VPP statistics socket, %v", err)
		p.statsAdapter = nil
	}

	return nil
}

// AfterInit reports status check.
func (p *Plugin) AfterInit() error {
	// Register providing status reports (push mode)
	p.StatusCheck.Register(p.PluginName, nil)
	p.StatusCheck.ReportStateChange(p.PluginName, statuscheck.OK, nil)

	var ctx context.Context
	ctx, p.cancel = context.WithCancel(context.Background())

	p.wg.Add(1)
	go p.handleVPPConnectionEvents(ctx)

	return nil
}

// Close cleans up the resources allocated by the govppmux plugin.
func (p *Plugin) Close() error {
	p.cancel()
	p.wg.Wait()

	defer func() {
		if p.vppConn != nil {
			p.vppConn.Disconnect()
		}
		if p.statsAdapter != nil {
			if err := p.statsAdapter.Disconnect(); err != nil {
				p.Log.Errorf("VPP statistics socket adapter disconnect error: %v", err)
			}
		}
	}()

	return nil
}
func (p *Plugin) CheckCompatiblity(msgs ...govppapi.Message) error {
	p.infoMu.Lock()
	defer p.infoMu.Unlock()
	if p.vppapiChan == nil {
		apiChan, err := p.vppConn.NewAPIChannel()
		if err != nil {
			return err
		}
		p.vppapiChan = apiChan
	}
	return p.vppapiChan.CheckCompatiblity(msgs...)
}

func (p *Plugin) Stats() govppapi.StatsProvider {
	return p
}

func (p *Plugin) StatsConnected() bool {
	return p.statsConn != nil
}

// VPPInfo returns information about VPP session.
func (p *Plugin) VPPInfo() VPPInfo {
	p.infoMu.Lock()
	defer p.infoMu.Unlock()
	return p.vppInfo
}

// IsPluginLoaded returns true if plugin is loaded.
func (p *Plugin) IsPluginLoaded(plugin string) bool {
	p.infoMu.Lock()
	defer p.infoMu.Unlock()
	for _, p := range p.vppInfo.Plugins {
		if p.Name == plugin {
			return true
		}
	}
	return false
}

func (p *Plugin) updateVPPInfo() (err error) {
	if p.vppConn == nil {
		return fmt.Errorf("VPP connection is nil")
	}

	ctx := context.Background()

	p.vppapiChan, err = p.vppConn.NewAPIChannel()
	if err != nil {
		return err
	}

	vpeHandler := vppcalls.CompatibleHandler(p)
	if vpeHandler == nil {
		return errors.New("no compatible VPP handler found")
	}

	version, err := vpeHandler.RunCli(ctx, "show version verbose")
	if err != nil {
		p.Log.Warnf("RunCli error: %v", err)
	} else {
		p.Log.Debugf("vpp# show version verbose\n%s", version)
	}

	cmdline, err := vpeHandler.RunCli(ctx, "show version cmdline")
	if err != nil {
		p.Log.Warnf("RunCli error: %v", err)
	} else {
		out := strings.Replace(cmdline, "\n", "", -1)
		p.Log.Debugf("vpp# show version cmdline:\n%s", out)
	}

	ver, err := vpeHandler.GetVersion(ctx)
	if err != nil {
		return err
	}
	session, err := vpeHandler.GetSession(ctx)
	if err != nil {
		return err
	}
	p.Log.WithFields(logging.Fields{
		"PID":      session.PID,
		"ClientID": session.ClientIdx,
	}).Infof("VPP version: %v", ver.Version)

	modules, err := vpeHandler.GetModules(ctx)
	if err != nil {
		return err
	}
	p.Log.Debugf("VPP has %d core modules: %v", len(modules), modules)

	plugins, err := vpeHandler.GetPlugins(ctx)
	if err != nil {
		return err
	}
	p.Log.Debugf("VPP loaded %d plugins", len(plugins))
	for _, plugin := range plugins {
		p.Log.Debugf(" - plugin: %v", plugin)
	}

	p.infoMu.Lock()
	p.vppInfo = VPPInfo{
		Connected:   true,
		VersionInfo: *ver,
		SessionInfo: *session,
		Plugins:     plugins,
	}
	p.infoMu.Unlock()

	return nil
}

// handleVPPConnectionEvents handles VPP connection events.
func (p *Plugin) handleVPPConnectionEvents(ctx context.Context) {
	defer p.wg.Done()

	for {
		select {
		case event := <-p.vppConChan:
			if event.State == govpp.Connected {
				if err := p.updateVPPInfo(); err != nil {
					p.Log.Errorf("updating VPP info failed: %v", err)
				}

				if p.config.ReconnectResync && p.lastConnErr != nil {
					p.Log.Info("Starting resync after VPP reconnect")
					if p.Resync != nil {
						p.Resync.DoResync()
						p.lastConnErr = nil
					} else {
						p.Log.Warn("Expected resync after VPP reconnect could not start because of missing Resync plugin")
					}
				}
				p.StatusCheck.ReportStateChange(p.PluginName, statuscheck.OK, nil)
			} else if event.State == govpp.Failed || event.State == govpp.Disconnected {
				p.infoMu.Lock()
				p.vppInfo.Connected = false
				p.infoMu.Unlock()

				p.lastConnErr = errors.Errorf("VPP connection lost (event: %+v)", event)
				p.StatusCheck.ReportStateChange(p.PluginName, statuscheck.Error, p.lastConnErr)
			} else {
				p.Log.Debugf("VPP connection state: %+v", event)
			}

			p.infoMu.Lock()
			p.lastEvent = event
			p.infoMu.Unlock()

		case <-ctx.Done():
			return
		}
	}
}
