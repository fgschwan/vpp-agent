// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/core/gre.api.json

/*
Package gre is a generated VPP binary API for 'gre' module.

It consists of:
	 13 enums
	  6 aliases
	  9 types
	  1 union
	  4 messages
	  2 services
*/
package gre

import (
	"bytes"
	"context"
	"io"
	"strconv"

	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"

	interface_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2009/interface_types"
	ip_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2009/ip_types"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "gre"
	// APIVersion is the API version of this module.
	APIVersion = "2.1.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x123bda70
)

type AddressFamily = ip_types.AddressFamily

// GreTunnelType represents VPP binary API enum 'gre_tunnel_type'.
type GreTunnelType uint8

const (
	GRE_API_TUNNEL_TYPE_L3     GreTunnelType = 0
	GRE_API_TUNNEL_TYPE_TEB    GreTunnelType = 1
	GRE_API_TUNNEL_TYPE_ERSPAN GreTunnelType = 2
)

var GreTunnelType_name = map[uint8]string{
	0: "GRE_API_TUNNEL_TYPE_L3",
	1: "GRE_API_TUNNEL_TYPE_TEB",
	2: "GRE_API_TUNNEL_TYPE_ERSPAN",
}

var GreTunnelType_value = map[string]uint8{
	"GRE_API_TUNNEL_TYPE_L3":     0,
	"GRE_API_TUNNEL_TYPE_TEB":    1,
	"GRE_API_TUNNEL_TYPE_ERSPAN": 2,
}

func (x GreTunnelType) String() string {
	s, ok := GreTunnelType_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

type IfStatusFlags = interface_types.IfStatusFlags

type IfType = interface_types.IfType

type IPDscp = ip_types.IPDscp

type IPEcn = ip_types.IPEcn

type IPProto = ip_types.IPProto

type LinkDuplex = interface_types.LinkDuplex

type MtuProto = interface_types.MtuProto

type RxMode = interface_types.RxMode

type SubIfFlags = interface_types.SubIfFlags

// TunnelEncapDecapFlags represents VPP binary API enum 'tunnel_encap_decap_flags'.
type TunnelEncapDecapFlags uint8

const (
	TUNNEL_API_ENCAP_DECAP_FLAG_NONE            TunnelEncapDecapFlags = 0
	TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_DF   TunnelEncapDecapFlags = 1
	TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_SET_DF    TunnelEncapDecapFlags = 2
	TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_DSCP TunnelEncapDecapFlags = 4
	TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_ECN  TunnelEncapDecapFlags = 8
	TUNNEL_API_ENCAP_DECAP_FLAG_DECAP_COPY_ECN  TunnelEncapDecapFlags = 16
)

var TunnelEncapDecapFlags_name = map[uint8]string{
	0:  "TUNNEL_API_ENCAP_DECAP_FLAG_NONE",
	1:  "TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_DF",
	2:  "TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_SET_DF",
	4:  "TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_DSCP",
	8:  "TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_ECN",
	16: "TUNNEL_API_ENCAP_DECAP_FLAG_DECAP_COPY_ECN",
}

var TunnelEncapDecapFlags_value = map[string]uint8{
	"TUNNEL_API_ENCAP_DECAP_FLAG_NONE":            0,
	"TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_DF":   1,
	"TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_SET_DF":    2,
	"TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_DSCP": 4,
	"TUNNEL_API_ENCAP_DECAP_FLAG_ENCAP_COPY_ECN":  8,
	"TUNNEL_API_ENCAP_DECAP_FLAG_DECAP_COPY_ECN":  16,
}

func (x TunnelEncapDecapFlags) String() string {
	s, ok := TunnelEncapDecapFlags_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// TunnelMode represents VPP binary API enum 'tunnel_mode'.
type TunnelMode uint8

const (
	TUNNEL_API_MODE_P2P TunnelMode = 0
	TUNNEL_API_MODE_MP  TunnelMode = 1
)

var TunnelMode_name = map[uint8]string{
	0: "TUNNEL_API_MODE_P2P",
	1: "TUNNEL_API_MODE_MP",
}

var TunnelMode_value = map[string]uint8{
	"TUNNEL_API_MODE_P2P": 0,
	"TUNNEL_API_MODE_MP":  1,
}

func (x TunnelMode) String() string {
	s, ok := TunnelMode_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

type AddressWithPrefix = ip_types.AddressWithPrefix

type InterfaceIndex = interface_types.InterfaceIndex

type IP4Address = ip_types.IP4Address

type IP4AddressWithPrefix = ip_types.IP4AddressWithPrefix

type IP6Address = ip_types.IP6Address

type IP6AddressWithPrefix = ip_types.IP6AddressWithPrefix

type Address = ip_types.Address

// GreTunnel represents VPP binary API type 'gre_tunnel'.
type GreTunnel struct {
	Type         GreTunnelType
	Mode         TunnelMode
	Flags        TunnelEncapDecapFlags
	SessionID    uint16
	Instance     uint32
	OuterTableID uint32
	SwIfIndex    InterfaceIndex
	Src          Address
	Dst          Address
}

func (*GreTunnel) GetTypeName() string { return "gre_tunnel" }

type IP4AddressAndMask = ip_types.IP4AddressAndMask

type IP4Prefix = ip_types.IP4Prefix

type IP6AddressAndMask = ip_types.IP6AddressAndMask

type IP6Prefix = ip_types.IP6Prefix

type Mprefix = ip_types.Mprefix

type Prefix = ip_types.Prefix

type PrefixMatcher = ip_types.PrefixMatcher

type AddressUnion = ip_types.AddressUnion

// GreTunnelAddDel represents VPP binary API message 'gre_tunnel_add_del'.
type GreTunnelAddDel struct {
	IsAdd  bool
	Tunnel GreTunnel
}

func (m *GreTunnelAddDel) Reset()                        { *m = GreTunnelAddDel{} }
func (*GreTunnelAddDel) GetMessageName() string          { return "gre_tunnel_add_del" }
func (*GreTunnelAddDel) GetCrcString() string            { return "6efc9c22" }
func (*GreTunnelAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// GreTunnelAddDelReply represents VPP binary API message 'gre_tunnel_add_del_reply'.
type GreTunnelAddDelReply struct {
	Retval    int32
	SwIfIndex InterfaceIndex
}

func (m *GreTunnelAddDelReply) Reset()                        { *m = GreTunnelAddDelReply{} }
func (*GreTunnelAddDelReply) GetMessageName() string          { return "gre_tunnel_add_del_reply" }
func (*GreTunnelAddDelReply) GetCrcString() string            { return "5383d31f" }
func (*GreTunnelAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// GreTunnelDetails represents VPP binary API message 'gre_tunnel_details'.
type GreTunnelDetails struct {
	Tunnel GreTunnel
}

func (m *GreTunnelDetails) Reset()                        { *m = GreTunnelDetails{} }
func (*GreTunnelDetails) GetMessageName() string          { return "gre_tunnel_details" }
func (*GreTunnelDetails) GetCrcString() string            { return "003bfbf1" }
func (*GreTunnelDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// GreTunnelDump represents VPP binary API message 'gre_tunnel_dump'.
type GreTunnelDump struct {
	SwIfIndex InterfaceIndex
}

func (m *GreTunnelDump) Reset()                        { *m = GreTunnelDump{} }
func (*GreTunnelDump) GetMessageName() string          { return "gre_tunnel_dump" }
func (*GreTunnelDump) GetCrcString() string            { return "f9e6675e" }
func (*GreTunnelDump) GetMessageType() api.MessageType { return api.RequestMessage }

func init() {
	api.RegisterMessage((*GreTunnelAddDel)(nil), "gre.GreTunnelAddDel")
	api.RegisterMessage((*GreTunnelAddDelReply)(nil), "gre.GreTunnelAddDelReply")
	api.RegisterMessage((*GreTunnelDetails)(nil), "gre.GreTunnelDetails")
	api.RegisterMessage((*GreTunnelDump)(nil), "gre.GreTunnelDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*GreTunnelAddDel)(nil),
		(*GreTunnelAddDelReply)(nil),
		(*GreTunnelDetails)(nil),
		(*GreTunnelDump)(nil),
	}
}

// RPCService represents RPC service API for gre module.
type RPCService interface {
	DumpGreTunnel(ctx context.Context, in *GreTunnelDump) (RPCService_DumpGreTunnelClient, error)
	GreTunnelAddDel(ctx context.Context, in *GreTunnelAddDel) (*GreTunnelAddDelReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpGreTunnel(ctx context.Context, in *GreTunnelDump) (RPCService_DumpGreTunnelClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpGreTunnelClient{stream}
	return x, nil
}

type RPCService_DumpGreTunnelClient interface {
	Recv() (*GreTunnelDetails, error)
}

type serviceClient_DumpGreTunnelClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpGreTunnelClient) Recv() (*GreTunnelDetails, error) {
	m := new(GreTunnelDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) GreTunnelAddDel(ctx context.Context, in *GreTunnelAddDel) (*GreTunnelAddDelReply, error) {
	out := new(GreTunnelAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack
