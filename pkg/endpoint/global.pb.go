// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: endpoints/global.proto

package endpoint

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

type GlobalLoginType struct {
	Type                 LoginType `protobuf:"varint,1,opt,name=type,proto3,enum=fuck_web.endpoint.LoginType" json:"type"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon                 string    `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	AutoLogin            bool      `protobuf:"varint,4,opt,name=auto_login,json=autoLogin,proto3" json:"autoLogin,omitempty"`
	AutoRedirect         bool      `protobuf:"varint,5,opt,name=auto_redirect,json=autoRedirect,proto3" json:"autoRedirect,omitempty"`
	Id                   string    `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GlobalLoginType) Reset()         { *m = GlobalLoginType{} }
func (m *GlobalLoginType) String() string { return proto.CompactTextString(m) }
func (*GlobalLoginType) ProtoMessage()    {}
func (*GlobalLoginType) Descriptor() ([]byte, []int) {
	return fileDescriptor_997cafa9b4dd7474, []int{0}
}
func (m *GlobalLoginType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalLoginType.Unmarshal(m, b)
}
func (m *GlobalLoginType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalLoginType.Marshal(b, m, deterministic)
}
func (m *GlobalLoginType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalLoginType.Merge(m, src)
}
func (m *GlobalLoginType) XXX_Size() int {
	return xxx_messageInfo_GlobalLoginType.Size(m)
}
func (m *GlobalLoginType) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalLoginType.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalLoginType proto.InternalMessageInfo

func (m *GlobalLoginType) GetType() LoginType {
	if m != nil {
		return m.Type
	}
	return LoginType_normal
}

func (m *GlobalLoginType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GlobalLoginType) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *GlobalLoginType) GetAutoLogin() bool {
	if m != nil {
		return m.AutoLogin
	}
	return false
}

func (m *GlobalLoginType) GetAutoRedirect() bool {
	if m != nil {
		return m.AutoRedirect
	}
	return false
}

func (m *GlobalLoginType) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GlobalConfig struct {
	LoginType            []*GlobalLoginType `protobuf:"bytes,1,rep,name=login_type,json=loginType,proto3" json:"loginType"`
	Title                string             `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	SubTitle             string             `protobuf:"bytes,3,opt,name=sub_title,json=subTitle,proto3" json:"subTitle,omitempty"`
	Logo                 string             `protobuf:"bytes,4,opt,name=logo,proto3" json:"logo,omitempty"`
	Copyright            string             `protobuf:"bytes,5,opt,name=copyright,proto3" json:"copyright,omitempty"`
	DefaultLoginType     LoginType          `protobuf:"varint,6,opt,name=DefaultLoginType,proto3,enum=fuck_web.endpoint.LoginType" json:"defaultLoginType"`
	Version              string             `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	ExternalUrl          string             `protobuf:"bytes,8,opt,name=external_url,json=externalUrl,proto3" json:"external_url,omitempty"`
	AdminUrl             string             `protobuf:"bytes,9,opt,name=admin_url,json=adminUrl,proto3" json:"admin_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GlobalConfig) Reset()         { *m = GlobalConfig{} }
func (m *GlobalConfig) String() string { return proto.CompactTextString(m) }
func (*GlobalConfig) ProtoMessage()    {}
func (*GlobalConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_997cafa9b4dd7474, []int{1}
}
func (m *GlobalConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalConfig.Unmarshal(m, b)
}
func (m *GlobalConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalConfig.Marshal(b, m, deterministic)
}
func (m *GlobalConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalConfig.Merge(m, src)
}
func (m *GlobalConfig) XXX_Size() int {
	return xxx_messageInfo_GlobalConfig.Size(m)
}
func (m *GlobalConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalConfig proto.InternalMessageInfo

func (m *GlobalConfig) GetLoginType() []*GlobalLoginType {
	if m != nil {
		return m.LoginType
	}
	return nil
}

func (m *GlobalConfig) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GlobalConfig) GetSubTitle() string {
	if m != nil {
		return m.SubTitle
	}
	return ""
}

func (m *GlobalConfig) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *GlobalConfig) GetCopyright() string {
	if m != nil {
		return m.Copyright
	}
	return ""
}

func (m *GlobalConfig) GetDefaultLoginType() LoginType {
	if m != nil {
		return m.DefaultLoginType
	}
	return LoginType_normal
}

func (m *GlobalConfig) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *GlobalConfig) GetExternalUrl() string {
	if m != nil {
		return m.ExternalUrl
	}
	return ""
}

func (m *GlobalConfig) GetAdminUrl() string {
	if m != nil {
		return m.AdminUrl
	}
	return ""
}

type GlobalConfigResponse struct {
	BaseResponse         `protobuf:"bytes,1,opt,name=base_response,json=baseResponse,proto3,embedded=base_response" json:",omitempty"`
	Data                 *GlobalConfig `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GlobalConfigResponse) Reset()         { *m = GlobalConfigResponse{} }
func (m *GlobalConfigResponse) String() string { return proto.CompactTextString(m) }
func (*GlobalConfigResponse) ProtoMessage()    {}
func (*GlobalConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_997cafa9b4dd7474, []int{2}
}
func (m *GlobalConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalConfigResponse.Unmarshal(m, b)
}
func (m *GlobalConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalConfigResponse.Marshal(b, m, deterministic)
}
func (m *GlobalConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalConfigResponse.Merge(m, src)
}
func (m *GlobalConfigResponse) XXX_Size() int {
	return xxx_messageInfo_GlobalConfigResponse.Size(m)
}
func (m *GlobalConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalConfigResponse proto.InternalMessageInfo

func (m *GlobalConfigResponse) GetData() *GlobalConfig {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GlobalLoginType)(nil), "fuck_web.endpoint.GlobalLoginType")
	proto.RegisterType((*GlobalConfig)(nil), "fuck_web.endpoint.GlobalConfig")
	proto.RegisterType((*GlobalConfigResponse)(nil), "fuck_web.endpoint.GlobalConfigResponse")
}

func init() { proto.RegisterFile("endpoints/global.proto", fileDescriptor_997cafa9b4dd7474) }

var fileDescriptor_997cafa9b4dd7474 = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xeb, 0x36, 0x6d, 0xe3, 0x49, 0x52, 0xca, 0x12, 0x05, 0x2b, 0x54, 0x72, 0xc8, 0x29,
	0x07, 0x1a, 0x4b, 0x89, 0x00, 0x09, 0x0e, 0x48, 0x06, 0x89, 0x0b, 0x88, 0xca, 0x6a, 0x2f, 0x08,
	0x61, 0xf9, 0xcf, 0xc6, 0x5d, 0xd5, 0xf1, 0x5a, 0xde, 0x35, 0x25, 0x4f, 0xc2, 0x33, 0xf0, 0x26,
	0x1c, 0x79, 0x02, 0x3f, 0x80, 0x1f, 0x02, 0xa1, 0x1d, 0xc7, 0x89, 0x69, 0x23, 0x71, 0xb1, 0x66,
	0x7f, 0xdf, 0xcc, 0xec, 0xee, 0xb7, 0x63, 0x18, 0xd0, 0x24, 0x4c, 0x39, 0x4b, 0xa4, 0xb0, 0xa2,
	0x98, 0xfb, 0x5e, 0x3c, 0x4d, 0x33, 0x2e, 0x39, 0x79, 0xb8, 0xc8, 0x83, 0x1b, 0xf7, 0x96, 0xfa,
	0xd3, 0x3a, 0x61, 0xd8, 0x8f, 0x78, 0xc4, 0x51, 0xb5, 0x54, 0x54, 0x25, 0x0e, 0x8d, 0x6d, 0x03,
	0x41, 0x85, 0x60, 0x3c, 0x11, 0x6b, 0xa5, 0xbf, 0x55, 0x7c, 0x4f, 0xd0, 0x8a, 0x8e, 0xff, 0x68,
	0xf0, 0xe0, 0x3d, 0xee, 0xf4, 0x81, 0x47, 0x2c, 0xb9, 0x5c, 0xa5, 0x94, 0xbc, 0x82, 0x96, 0x5c,
	0xa5, 0xd4, 0xd0, 0x46, 0xda, 0xe4, 0x64, 0x76, 0x36, 0xbd, 0xb7, 0xf7, 0x74, 0x93, 0x6b, 0xb7,
	0xcb, 0xc2, 0xc4, 0x6c, 0x07, 0xbf, 0x84, 0x40, 0x2b, 0xf1, 0x96, 0xd4, 0xd8, 0x1f, 0x69, 0x13,
	0xdd, 0xc1, 0x58, 0x31, 0x16, 0xf0, 0xc4, 0x38, 0xa8, 0x98, 0x8a, 0xc9, 0x0b, 0x00, 0x2f, 0x97,
	0xdc, 0x8d, 0x55, 0x27, 0xa3, 0x35, 0xd2, 0x26, 0x6d, 0xfb, 0x71, 0x59, 0x98, 0x8f, 0x14, 0xc5,
	0xf6, 0xcf, 0xf8, 0x92, 0x49, 0xba, 0x4c, 0xe5, 0xca, 0xd1, 0x37, 0x90, 0xbc, 0x81, 0x1e, 0xd6,
	0x65, 0x34, 0x64, 0x19, 0x0d, 0xa4, 0x71, 0x88, 0xa5, 0xc3, 0xb2, 0x30, 0x07, 0x4a, 0x70, 0xd6,
	0xbc, 0x51, 0xdd, 0x6d, 0x72, 0x72, 0x02, 0xfb, 0x2c, 0x34, 0x8e, 0xf0, 0x28, 0xfb, 0x2c, 0x1c,
	0xff, 0x38, 0x80, 0x6e, 0x65, 0xc0, 0x5b, 0x9e, 0x2c, 0x58, 0x44, 0x2e, 0x00, 0xf0, 0x50, 0xee,
	0xda, 0x83, 0x83, 0x49, 0x67, 0x36, 0xde, 0xe1, 0xc1, 0x1d, 0xd7, 0xec, 0x5e, 0x59, 0x98, 0x7a,
	0x5c, 0x2f, 0x9d, 0x6d, 0x48, 0xfa, 0x70, 0x28, 0x99, 0x8c, 0x6b, 0x53, 0xaa, 0x05, 0x99, 0x83,
	0x2e, 0x72, 0xdf, 0xad, 0x14, 0xb4, 0xc6, 0x1e, 0x94, 0x85, 0x49, 0x44, 0xee, 0x5f, 0x2a, 0xd6,
	0xb8, 0x41, 0xbb, 0x66, 0xca, 0xca, 0x98, 0x47, 0x1c, 0x0d, 0xd3, 0x1d, 0x8c, 0xc9, 0x19, 0xe8,
	0x01, 0x4f, 0x57, 0x19, 0x8b, 0xae, 0x2b, 0x3b, 0x74, 0x67, 0x0b, 0xc8, 0x57, 0x38, 0x7d, 0x47,
	0x17, 0x5e, 0x1e, 0xcb, 0xcd, 0x51, 0xf1, 0xf6, 0xff, 0x7b, 0xd8, 0x7e, 0x59, 0x98, 0xa7, 0xe1,
	0x9d, 0x4a, 0xe7, 0x5e, 0x2f, 0x62, 0xc0, 0xf1, 0x37, 0x9a, 0xa9, 0x41, 0x33, 0x8e, 0x71, 0xef,
	0x7a, 0x49, 0x9e, 0x42, 0x97, 0x7e, 0x97, 0x34, 0x4b, 0xbc, 0xd8, 0xcd, 0xb3, 0xd8, 0x68, 0xa3,
	0xdc, 0xa9, 0xd9, 0x55, 0x16, 0x93, 0x27, 0xa0, 0x7b, 0xe1, 0x92, 0x25, 0xa8, 0xeb, 0xa8, 0xb7,
	0x11, 0x5c, 0x65, 0xf1, 0xf8, 0xa7, 0x06, 0xfd, 0xe6, 0xcb, 0x38, 0x54, 0xa4, 0x3c, 0x11, 0x94,
	0x7c, 0x81, 0x9e, 0x9a, 0x60, 0x37, 0x5b, 0x03, 0x1c, 0xd4, 0xce, 0xcc, 0xdc, 0x71, 0x1f, 0xdb,
	0x13, 0xb4, 0xae, 0xb3, 0x07, 0xbf, 0x0a, 0x73, 0xef, 0x77, 0x61, 0x6a, 0x65, 0x61, 0x42, 0x73,
	0x40, 0xfc, 0x46, 0x16, 0x99, 0x43, 0x2b, 0xf4, 0xa4, 0x87, 0x8f, 0xb5, 0xbb, 0xe9, 0x3f, 0x87,
	0xc2, 0x64, 0xfb, 0xe5, 0xe7, 0xe7, 0x11, 0x93, 0xd7, 0xb9, 0x3f, 0x0d, 0xf8, 0xd2, 0xfa, 0xc8,
	0x82, 0x8c, 0x7f, 0x4a, 0xc5, 0x79, 0x90, 0x58, 0xaa, 0xfc, 0xfc, 0x96, 0xfa, 0x56, 0x7a, 0x13,
	0x59, 0x75, 0x8b, 0xd7, 0x75, 0x70, 0xb1, 0xe7, 0x1f, 0xe1, 0x8f, 0x38, 0xff, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0xea, 0x43, 0x09, 0x01, 0xfb, 0x03, 0x00, 0x00,
}
