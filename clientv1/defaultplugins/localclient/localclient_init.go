// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package localclient

import (
	"github.com/ligato/cn-infra/datasync"
	"github.com/ligato/cn-infra/datasync/kvdbsync/local"
	"github.com/ligato/cn-infra/datasync/syncbase"
)

// Plugin implements the core.Plugin interface.
type Plugin struct {
}

// Init tries to sets the default transport (can be set only if it is nil)
func (plugin *Plugin) Init() error {
	return datasync.RegisterTransport(&syncbase.Adapter{Watcher: local.Get()})
}

// Close does nothing
func (plugin *Plugin) Close() error {
	return nil
}
