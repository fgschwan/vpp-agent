// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: models/vpp/l3/vrf.proto

package vpp_l3

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type VrfTable_Protocol int32

const (
	VrfTable_IPV4 VrfTable_Protocol = 0
	VrfTable_IPV6 VrfTable_Protocol = 1
)

var VrfTable_Protocol_name = map[int32]string{
	0: "IPV4",
	1: "IPV6",
}

var VrfTable_Protocol_value = map[string]int32{
	"IPV4": 0,
	"IPV6": 1,
}

func (x VrfTable_Protocol) String() string {
	return proto.EnumName(VrfTable_Protocol_name, int32(x))
}

func (VrfTable_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1f737b4f5eb6705c, []int{0, 0}
}

type VrfTable struct {
	Id                   uint32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Protocol             VrfTable_Protocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=vpp.l3.VrfTable_Protocol" json:"protocol,omitempty"`
	Label                string            `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *VrfTable) Reset()         { *m = VrfTable{} }
func (m *VrfTable) String() string { return proto.CompactTextString(m) }
func (*VrfTable) ProtoMessage()    {}
func (*VrfTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f737b4f5eb6705c, []int{0}
}
func (m *VrfTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VrfTable.Unmarshal(m, b)
}
func (m *VrfTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VrfTable.Marshal(b, m, deterministic)
}
func (m *VrfTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VrfTable.Merge(m, src)
}
func (m *VrfTable) XXX_Size() int {
	return xxx_messageInfo_VrfTable.Size(m)
}
func (m *VrfTable) XXX_DiscardUnknown() {
	xxx_messageInfo_VrfTable.DiscardUnknown(m)
}

var xxx_messageInfo_VrfTable proto.InternalMessageInfo

func (m *VrfTable) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VrfTable) GetProtocol() VrfTable_Protocol {
	if m != nil {
		return m.Protocol
	}
	return VrfTable_IPV4
}

func (m *VrfTable) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (*VrfTable) XXX_MessageName() string {
	return "vpp.l3.VrfTable"
}
func init() {
	proto.RegisterEnum("vpp.l3.VrfTable_Protocol", VrfTable_Protocol_name, VrfTable_Protocol_value)
	proto.RegisterType((*VrfTable)(nil), "vpp.l3.VrfTable")
}

func init() { proto.RegisterFile("models/vpp/l3/vrf.proto", fileDescriptor_1f737b4f5eb6705c) }

var fileDescriptor_1f737b4f5eb6705c = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0xcd, 0x4f, 0x49,
	0xcd, 0x29, 0xd6, 0x2f, 0x2b, 0x28, 0xd0, 0xcf, 0x31, 0xd6, 0x2f, 0x2b, 0x4a, 0xd3, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b, 0x2b, 0x28, 0xd0, 0xcb, 0x31, 0x96, 0xd2, 0x4d, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x4b, 0x27,
	0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa6, 0xd4, 0xce, 0xc8, 0xc5, 0x11, 0x56,
	0x94, 0x16, 0x92, 0x98, 0x94, 0x93, 0x2a, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x1b, 0xc4, 0x94, 0x99, 0x22, 0x64, 0xca, 0xc5, 0x01, 0x56, 0x95, 0x9c, 0x9f, 0x23,
	0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x67, 0x24, 0xa9, 0x07, 0xb1, 0x46, 0x0f, 0xa6, 0x47, 0x2f, 0x00,
	0xaa, 0x20, 0x08, 0xae, 0x54, 0x48, 0x84, 0x8b, 0x35, 0x27, 0x31, 0x29, 0x35, 0x47, 0x82, 0x59,
	0x81, 0x51, 0x83, 0x33, 0x08, 0xc2, 0x51, 0x92, 0xe3, 0xe2, 0x80, 0xa9, 0x15, 0xe2, 0xe0, 0x62,
	0xf1, 0x0c, 0x08, 0x33, 0x11, 0x60, 0x80, 0xb2, 0xcc, 0x04, 0x18, 0x9d, 0xac, 0x4e, 0x3c, 0x96,
	0x63, 0x8c, 0x32, 0x41, 0x72, 0x7e, 0x4e, 0x66, 0x7a, 0x62, 0x49, 0x3e, 0xc8, 0xab, 0xba, 0x89,
	0xe9, 0xa9, 0x79, 0x25, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x28, 0xfe, 0xb7, 0x2e, 0x2b, 0x28, 0x88,
	0xcf, 0x31, 0x4e, 0x62, 0x03, 0xdb, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x87, 0x3a, 0xf8,
	0xe4, 0x1e, 0x01, 0x00, 0x00,
}