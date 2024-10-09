// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/oauth2.proto

package oauth2

import (
	fmt "fmt"
	github_com_MicroOps_cn_fuck_safe "github.com/MicroOps-cn/fuck/safe"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	math_bits "math/bits"
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

type CustomType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomType) Reset()         { *m = CustomType{} }
func (m *CustomType) String() string { return proto.CompactTextString(m) }
func (*CustomType) ProtoMessage()    {}
func (*CustomType) Descriptor() ([]byte, []int) {
	return fileDescriptor_02bed1b4386e2ff5, []int{0}
}
func (m *CustomType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomType.Unmarshal(m, b)
}
func (m *CustomType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomType.Marshal(b, m, deterministic)
}
func (m *CustomType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomType.Merge(m, src)
}
func (m *CustomType) XXX_Size() int {
	return xxx_messageInfo_CustomType.Size(m)
}
func (m *CustomType) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomType.DiscardUnknown(m)
}

var xxx_messageInfo_CustomType proto.InternalMessageInfo

type Options struct {
	Name                     string                                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" valid:"required"`
	Icon                     string                                   `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon,omitempty"`
	TokenUrl                 string                                   `protobuf:"bytes,3,opt,name=token_url,json=tokenUrl,proto3" json:"token_url,omitempty" valid:"required,url"`
	AuthUrl                  string                                   `protobuf:"bytes,4,opt,name=auth_url,json=authUrl,proto3" json:"auth_url,omitempty" valid:"required,url"`
	ApiUrl                   string                                   `protobuf:"bytes,5,opt,name=api_url,json=apiUrl,proto3" json:"api_url,omitempty" valid:"required,url"`
	ClientId                 string                                   `protobuf:"bytes,6,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" valid:"required"`
	ClientSecret             *github_com_MicroOps_cn_fuck_safe.String `protobuf:"bytes,7,opt,name=client_secret,json=clientSecret,proto3,customtype=github.com/MicroOps-cn/fuck/safe.String" json:"client_secret,omitempty" valid:"required"`
	AutoLogin                bool                                     `protobuf:"varint,8,opt,name=auto_login,json=autoLogin,proto3" json:"auto_login,omitempty"`
	Id                       string                                   `protobuf:"bytes,9,opt,name=id,proto3" json:"id,omitempty" valid:"required"`
	LoginId                  string                                   `protobuf:"bytes,10,opt,name=login_id,json=loginId,proto3" json:"login_id,omitempty"`
	EmailAttributePath       string                                   `protobuf:"bytes,11,opt,name=email_attribute_path,json=emailAttributePath,proto3" json:"email_attribute_path,omitempty"`
	UsernameAttributePath    string                                   `protobuf:"bytes,12,opt,name=username_attribute_path,json=usernameAttributePath,proto3" json:"username_attribute_path,omitempty"`
	PhoneNumberAttributePath string                                   `protobuf:"bytes,13,opt,name=phone_number_attribute_path,json=phoneNumberAttributePath,proto3" json:"phone_number_attribute_path,omitempty"`
	FullNameAttributePath    string                                   `protobuf:"bytes,14,opt,name=full_name_attribute_path,json=fullNameAttributePath,proto3" json:"full_name_attribute_path,omitempty"`
	RoleAttributePath        string                                   `protobuf:"bytes,15,opt,name=role_attribute_path,json=roleAttributePath,proto3" json:"role_attribute_path,omitempty"`
	AvatarAttributePath      string                                   `protobuf:"bytes,16,opt,name=avatar_attribute_path,json=avatarAttributePath,proto3" json:"avatar_attribute_path,omitempty"`
	Scope                    string                                   `protobuf:"bytes,17,opt,name=scope,proto3" json:"scope,omitempty"`
	AutoRedirect             bool                                     `protobuf:"varint,18,opt,name=auto_redirect,json=autoRedirect,proto3" json:"auto_redirect,omitempty"`
	AllowRegister            bool                                     `protobuf:"varint,19,opt,name=allow_register,json=allowRegister,proto3" json:"allow_register,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                                 `json:"-"`
	XXX_unrecognized         []byte                                   `json:"-"`
	XXX_sizecache            int32                                    `json:"-"`
}

func (m *Options) Reset()         { *m = Options{} }
func (m *Options) String() string { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()    {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_02bed1b4386e2ff5, []int{1}
}
func (m *Options) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Options.Unmarshal(m, b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Options.Marshal(b, m, deterministic)
}
func (m *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(m, src)
}
func (m *Options) XXX_Size() int {
	return xxx_messageInfo_Options.Size(m)
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

func (m *Options) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Options) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Options) GetTokenUrl() string {
	if m != nil {
		return m.TokenUrl
	}
	return ""
}

func (m *Options) GetAuthUrl() string {
	if m != nil {
		return m.AuthUrl
	}
	return ""
}

func (m *Options) GetApiUrl() string {
	if m != nil {
		return m.ApiUrl
	}
	return ""
}

func (m *Options) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Options) GetAutoLogin() bool {
	if m != nil {
		return m.AutoLogin
	}
	return false
}

func (m *Options) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Options) GetLoginId() string {
	if m != nil {
		return m.LoginId
	}
	return ""
}

func (m *Options) GetEmailAttributePath() string {
	if m != nil {
		return m.EmailAttributePath
	}
	return ""
}

func (m *Options) GetUsernameAttributePath() string {
	if m != nil {
		return m.UsernameAttributePath
	}
	return ""
}

func (m *Options) GetPhoneNumberAttributePath() string {
	if m != nil {
		return m.PhoneNumberAttributePath
	}
	return ""
}

func (m *Options) GetFullNameAttributePath() string {
	if m != nil {
		return m.FullNameAttributePath
	}
	return ""
}

func (m *Options) GetRoleAttributePath() string {
	if m != nil {
		return m.RoleAttributePath
	}
	return ""
}

func (m *Options) GetAvatarAttributePath() string {
	if m != nil {
		return m.AvatarAttributePath
	}
	return ""
}

func (m *Options) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *Options) GetAutoRedirect() bool {
	if m != nil {
		return m.AutoRedirect
	}
	return false
}

func (m *Options) GetAllowRegister() bool {
	if m != nil {
		return m.AllowRegister
	}
	return false
}

func init() {
	proto.RegisterType((*CustomType)(nil), "fuck_web.client.oauth2.custom_type")
	proto.RegisterType((*Options)(nil), "fuck_web.client.oauth2.Options")
}

func init() { proto.RegisterFile("types/oauth2.proto", fileDescriptor_02bed1b4386e2ff5) }

var fileDescriptor_02bed1b4386e2ff5 = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xc7, 0x7f, 0xd9, 0x6f, 0x5b, 0x52, 0x6f, 0x1d, 0x9b, 0xd7, 0x81, 0x01, 0x41, 0xab, 0x0e,
	0x44, 0x2f, 0x58, 0x02, 0x9d, 0xf8, 0x23, 0x10, 0x20, 0x7a, 0x37, 0x09, 0xb6, 0x29, 0x13, 0x37,
	0xdc, 0x44, 0x4e, 0xe2, 0xa6, 0x56, 0xdd, 0x38, 0x38, 0xf6, 0x2a, 0x1e, 0x81, 0x87, 0xe1, 0x85,
	0xb8, 0xa8, 0x10, 0x8f, 0xb0, 0x27, 0x40, 0x3e, 0xe9, 0x24, 0xfa, 0x47, 0xdd, 0x55, 0xe2, 0xf3,
	0xfd, 0x7e, 0xce, 0xb1, 0x8f, 0x8f, 0x8c, 0xb0, 0xfe, 0x5e, 0xb0, 0x32, 0x90, 0xd4, 0xe8, 0x41,
	0xd7, 0x2f, 0x94, 0xd4, 0x12, 0xdf, 0xee, 0x9b, 0x64, 0x18, 0x8d, 0x59, 0xec, 0x27, 0x82, 0xb3,
	0x5c, 0xfb, 0x95, 0x7a, 0xaf, 0x91, 0xc9, 0x4c, 0x82, 0x25, 0xb0, 0x7f, 0x95, 0xbb, 0x5d, 0x47,
	0x5b, 0x89, 0x29, 0xb5, 0x1c, 0x45, 0x36, 0x55, 0xfb, 0xa7, 0x8b, 0xdc, 0xb3, 0x42, 0x73, 0x99,
	0x97, 0xb8, 0x83, 0xd6, 0x73, 0x3a, 0x62, 0xc4, 0x69, 0x39, 0x9d, 0x5a, 0xaf, 0x71, 0x35, 0x69,
	0xee, 0x5e, 0x52, 0xc1, 0xd3, 0x37, 0x6d, 0xc5, 0xbe, 0x19, 0xae, 0x58, 0xda, 0x0e, 0xc1, 0x81,
	0x31, 0x5a, 0xe7, 0x89, 0xcc, 0xc9, 0x9a, 0x75, 0x86, 0xf0, 0x8f, 0x5f, 0xa0, 0x9a, 0x96, 0x43,
	0x96, 0x47, 0x46, 0x09, 0xf2, 0x3f, 0xa4, 0x20, 0x57, 0x93, 0x66, 0x63, 0x2e, 0xc5, 0x53, 0xa3,
	0x44, 0x3b, 0xf4, 0xc0, 0xfa, 0x45, 0x09, 0x7c, 0x8c, 0x3c, 0xbb, 0x5d, 0xa0, 0xd6, 0x6f, 0xa0,
	0x5c, 0xeb, 0xb4, 0xd0, 0x73, 0xe4, 0xd2, 0x82, 0x03, 0xb3, 0x71, 0x03, 0xb3, 0x49, 0x0b, 0x5e,
	0x21, 0xb5, 0xaa, 0x3d, 0x11, 0x4f, 0xc9, 0xe6, 0x8a, 0x13, 0x7a, 0x95, 0xed, 0x24, 0xc5, 0x3f,
	0x1c, 0x54, 0x9f, 0x32, 0x25, 0x4b, 0x14, 0xd3, 0xc4, 0x6d, 0x39, 0x9d, 0xad, 0xee, 0xa1, 0xbf,
	0xbc, 0xe3, 0xfe, 0x3f, 0x8d, 0xed, 0x7d, 0xf8, 0x35, 0x69, 0x3e, 0xc9, 0xb8, 0x1e, 0x98, 0xd8,
	0x4f, 0xe4, 0x28, 0xf8, 0xcc, 0x13, 0x25, 0xcf, 0x8a, 0xf2, 0x28, 0xc9, 0x03, 0x8b, 0x07, 0x25,
	0xed, 0x33, 0xff, 0x42, 0x2b, 0x9e, 0x67, 0x4b, 0xf7, 0xb1, 0x5d, 0xe5, 0xbe, 0x80, 0xca, 0xf8,
	0x01, 0x42, 0xd4, 0x68, 0x19, 0x09, 0x99, 0xf1, 0x9c, 0x78, 0x2d, 0xa7, 0xe3, 0x85, 0x35, 0x1b,
	0xf9, 0x64, 0x03, 0xf8, 0x11, 0x5a, 0xe3, 0x29, 0xa9, 0xad, 0x38, 0xd6, 0x1a, 0x4f, 0xf1, 0x5d,
	0xe4, 0x01, 0x6f, 0x5b, 0x80, 0xe0, 0xea, 0x5c, 0x58, 0x9f, 0xa4, 0xf8, 0x19, 0x6a, 0xb0, 0x11,
	0xe5, 0x22, 0xa2, 0x5a, 0x2b, 0x1e, 0x1b, 0xcd, 0xa2, 0x82, 0xea, 0x01, 0xd9, 0x02, 0x1b, 0x06,
	0xed, 0xe3, 0xb5, 0x74, 0x4e, 0xf5, 0x00, 0xbf, 0x44, 0x77, 0x4c, 0xc9, 0x94, 0x9d, 0x87, 0x79,
	0x68, 0x1b, 0xa0, 0x83, 0x6b, 0x79, 0x96, 0x7b, 0x87, 0xee, 0x17, 0x03, 0x99, 0xb3, 0x28, 0x37,
	0xa3, 0x98, 0xa9, 0x79, 0xb6, 0x0e, 0x2c, 0x01, 0xcb, 0x29, 0x38, 0x66, 0xf1, 0x57, 0x88, 0xf4,
	0x8d, 0x10, 0xd1, 0xb2, 0xba, 0x3b, 0x55, 0x5d, 0xab, 0x9f, 0x2e, 0xd4, 0xf5, 0xd1, 0xbe, 0x92,
	0x62, 0x81, 0xb9, 0x05, 0xcc, 0x9e, 0x95, 0x66, 0xfd, 0x5d, 0x74, 0x40, 0x2f, 0xa9, 0xa6, 0x0b,
	0x3b, 0xdc, 0x05, 0x62, 0xbf, 0x12, 0x67, 0x99, 0x06, 0xda, 0x28, 0x13, 0x59, 0x30, 0xb2, 0x07,
	0x9e, 0x6a, 0x81, 0x0f, 0x51, 0x1d, 0xee, 0x4e, 0xb1, 0x94, 0x2b, 0x96, 0x68, 0x82, 0xe1, 0xfa,
	0xb6, 0x6d, 0x30, 0x9c, 0xc6, 0xf0, 0x63, 0xb4, 0x43, 0x85, 0x90, 0xe3, 0x48, 0xb1, 0x8c, 0x97,
	0x9a, 0x29, 0xb2, 0x0f, 0xae, 0x3a, 0x44, 0xc3, 0x69, 0xb0, 0xf7, 0xfe, 0xf7, 0x9f, 0x87, 0xce,
	0xd7, 0xd7, 0x2b, 0x06, 0xeb, 0x68, 0xcc, 0xe2, 0xa0, 0x18, 0x66, 0x41, 0x35, 0x3f, 0xd3, 0xb7,
	0xe2, 0x6d, 0xf5, 0x39, 0xff, 0x2f, 0xde, 0x84, 0x77, 0xe0, 0xf8, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xfe, 0xa9, 0x71, 0xb6, 0x4b, 0x04, 0x00, 0x00,
}

func (m *CustomType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Options) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovOauth2(uint64(l))
	}
	l = len(m.Icon)
	if l > 0 {
		n += 1 + l + sovOauth2(uint64(l))
	}
	l = len(m.TokenUrl)
	if l > 0 {
		n += 1 + l + sovOauth2(uint64(l))
	}
	l = len(m.AuthUrl)
	if l > 0 {
		n += 1 + l + sovOauth2(uint64(l))
	}
	l = len(m.ApiUrl)
	if l > 0 {
		n += 1 + l + sovOauth2(uint64(l))
	}
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovOauth2(uint64(l))
	}
	if m.ClientSecret != nil {
		l = m.ClientSecret.Size()
		n += 1 + l + sovOauth2(uint64(l))
	}
	if m.AutoLogin {
		n += 2
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovOauth2(uint64(l))
	}
	l = len(m.LoginId)
	if l > 0 {
		n += 1 + l + sovOauth2(uint64(l))
	}
	l = len(m.EmailAttributePath)
	if l > 0 {
		n += 1 + l + sovOauth2(uint64(l))
	}
	l = len(m.UsernameAttributePath)
	if l > 0 {
		n += 1 + l + sovOauth2(uint64(l))
	}
	l = len(m.PhoneNumberAttributePath)
	if l > 0 {
		n += 1 + l + sovOauth2(uint64(l))
	}
	l = len(m.FullNameAttributePath)
	if l > 0 {
		n += 1 + l + sovOauth2(uint64(l))
	}
	l = len(m.RoleAttributePath)
	if l > 0 {
		n += 1 + l + sovOauth2(uint64(l))
	}
	l = len(m.AvatarAttributePath)
	if l > 0 {
		n += 2 + l + sovOauth2(uint64(l))
	}
	l = len(m.Scope)
	if l > 0 {
		n += 2 + l + sovOauth2(uint64(l))
	}
	if m.AutoRedirect {
		n += 3
	}
	if m.AllowRegister {
		n += 3
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovOauth2(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOauth2(x uint64) (n int) {
	return sovOauth2(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
