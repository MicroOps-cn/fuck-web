// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/tls.proto

package tls

import (
	fmt "fmt"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TLSOptions struct {
	CAFile               string   `protobuf:"bytes,1,opt,name=ca_file,json=caFile,proto3" json:"ca_file,omitempty"`
	CertFile             string   `protobuf:"bytes,2,opt,name=cert_file,json=certFile,proto3" json:"cert_file,omitempty"`
	KeyFile              string   `protobuf:"bytes,3,opt,name=key_file,json=keyFile,proto3" json:"key_file,omitempty"`
	ServerName           string   `protobuf:"bytes,4,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
	InsecureSkipVerify   bool     `protobuf:"varint,5,opt,name=insecure_skip_verify,json=insecureSkipVerify,proto3" json:"insecure_skip_verify,omitempty"`
	MinVersion           string   `protobuf:"bytes,6,opt,name=min_version,json=minVersion,proto3" json:"min_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TLSOptions) Reset()         { *m = TLSOptions{} }
func (m *TLSOptions) String() string { return proto.CompactTextString(m) }
func (*TLSOptions) ProtoMessage()    {}
func (*TLSOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c06a1c6ca94056c, []int{0}
}
func (m *TLSOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TLSOptions.Unmarshal(m, b)
}
func (m *TLSOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TLSOptions.Marshal(b, m, deterministic)
}
func (m *TLSOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLSOptions.Merge(m, src)
}
func (m *TLSOptions) XXX_Size() int {
	return xxx_messageInfo_TLSOptions.Size(m)
}
func (m *TLSOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_TLSOptions.DiscardUnknown(m)
}

var xxx_messageInfo_TLSOptions proto.InternalMessageInfo

func (m *TLSOptions) GetCAFile() string {
	if m != nil {
		return m.CAFile
	}
	return ""
}

func (m *TLSOptions) GetCertFile() string {
	if m != nil {
		return m.CertFile
	}
	return ""
}

func (m *TLSOptions) GetKeyFile() string {
	if m != nil {
		return m.KeyFile
	}
	return ""
}

func (m *TLSOptions) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *TLSOptions) GetInsecureSkipVerify() bool {
	if m != nil {
		return m.InsecureSkipVerify
	}
	return false
}

func (m *TLSOptions) GetMinVersion() string {
	if m != nil {
		return m.MinVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*TLSOptions)(nil), "fuck_web.client.tls.TLSOptions")
}

func init() { proto.RegisterFile("types/tls.proto", fileDescriptor_9c06a1c6ca94056c) }

var fileDescriptor_9c06a1c6ca94056c = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xd0, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x05, 0xe0, 0x3f, 0x3f, 0x90, 0xb6, 0x66, 0x40, 0x0a, 0x1d, 0x02, 0x0c, 0xad, 0x60, 0xe9,
	0xd2, 0x18, 0x89, 0xb1, 0x62, 0xa0, 0x48, 0x4c, 0x40, 0x51, 0x8b, 0x3a, 0xb0, 0x44, 0x89, 0x75,
	0x1b, 0xae, 0xec, 0xd8, 0x96, 0xed, 0xb4, 0xca, 0xc3, 0xc2, 0xc0, 0x93, 0x20, 0xdb, 0x62, 0xb2,
	0x75, 0xbe, 0xa3, 0x3b, 0x1c, 0x72, 0xe6, 0x7a, 0x0d, 0x96, 0x3a, 0x61, 0x0b, 0x6d, 0x94, 0x53,
	0xd9, 0xf9, 0xae, 0x63, 0xbc, 0x3c, 0x40, 0x5d, 0x30, 0x81, 0x20, 0x5d, 0xe1, 0x84, 0xbd, 0x1c,
	0x37, 0xaa, 0x51, 0xc1, 0xa9, 0xff, 0xc5, 0xea, 0xf5, 0x57, 0x42, 0xc8, 0xfb, 0xf3, 0x66, 0xa5,
	0x1d, 0x2a, 0x69, 0xb3, 0x1b, 0x32, 0x60, 0x55, 0xb9, 0x43, 0x01, 0x79, 0x32, 0x4d, 0x66, 0xa3,
	0x25, 0xf9, 0xf9, 0x9e, 0xa4, 0x8f, 0x0f, 0x4f, 0x28, 0x60, 0x9d, 0xb2, 0xca, 0xbf, 0xd9, 0x15,
	0x19, 0x31, 0x30, 0x2e, 0xd6, 0xfe, 0xfb, 0xda, 0x7a, 0xe8, 0x83, 0x80, 0x17, 0x64, 0xc8, 0xa1,
	0x8f, 0x76, 0x14, 0x6c, 0xc0, 0xa1, 0x0f, 0x34, 0x21, 0xa7, 0x16, 0xcc, 0x1e, 0x4c, 0x29, 0xab,
	0x16, 0xf2, 0xe3, 0xa0, 0x24, 0x46, 0xaf, 0x55, 0x0b, 0xd9, 0x2d, 0x19, 0xa3, 0xb4, 0xc0, 0x3a,
	0x03, 0xa5, 0xe5, 0xa8, 0xcb, 0x3d, 0x18, 0xdc, 0xf5, 0xf9, 0xc9, 0x34, 0x99, 0x0d, 0xd7, 0xd9,
	0x9f, 0x6d, 0x38, 0xea, 0x6d, 0x10, 0x7f, 0xb2, 0x45, 0xe9, 0x7b, 0x16, 0x95, 0xcc, 0xd3, 0x78,
	0xb2, 0x45, 0xb9, 0x8d, 0xc9, 0xf2, 0xfe, 0x63, 0xd1, 0xa0, 0xfb, 0xec, 0xea, 0x82, 0xa9, 0x96,
	0xbe, 0x20, 0x33, 0x6a, 0xa5, 0xed, 0x9c, 0x49, 0xea, 0x37, 0x9a, 0x1f, 0xa0, 0xa6, 0x9a, 0x37,
	0x34, 0xee, 0x44, 0x51, 0x3a, 0x30, 0xb2, 0x12, 0x7e, 0xcb, 0x85, 0x13, 0xf6, 0xed, 0x5f, 0x9d,
	0x86, 0x9d, 0xee, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xce, 0xe2, 0xa4, 0x6e, 0x65, 0x01, 0x00,
	0x00,
}
