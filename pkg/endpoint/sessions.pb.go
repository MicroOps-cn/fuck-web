// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: endpoints/sessions.proto

package endpoint

import (
	fmt "fmt"
	github_com_MicroOps_cn_fuck_web_pkg_service_models "github.com/MicroOps-cn/fuck-web/pkg/service/models"
	models "github.com/MicroOps-cn/fuck-web/pkg/service/models"
	github_com_MicroOps_cn_fuck_web_pkg_utils_sign "github.com/MicroOps-cn/fuck-web/pkg/utils/sign"
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

// @sync-to-public:public/src/services/fuck-web/enums.ts:LoginType
type LoginType int32

const (
	LoginType_normal           LoginType = 0
	LoginType_mfa_totp         LoginType = 1
	LoginType_mfa_email        LoginType = 2
	LoginType_mfa_sms          LoginType = 3
	LoginType_email            LoginType = 4
	LoginType_sms              LoginType = 5
	LoginType_oauth2           LoginType = 6
	LoginType_enable_mfa_totp  LoginType = 10
	LoginType_enable_mfa_email LoginType = 11
	LoginType_enable_mfa_sms   LoginType = 12
)

var LoginType_name = map[int32]string{
	0:  "normal",
	1:  "mfa_totp",
	2:  "mfa_email",
	3:  "mfa_sms",
	4:  "email",
	5:  "sms",
	6:  "oauth2",
	10: "enable_mfa_totp",
	11: "enable_mfa_email",
	12: "enable_mfa_sms",
}

var LoginType_value = map[string]int32{
	"normal":           0,
	"mfa_totp":         1,
	"mfa_email":        2,
	"mfa_sms":          3,
	"email":            4,
	"sms":              5,
	"oauth2":           6,
	"enable_mfa_totp":  10,
	"enable_mfa_email": 11,
	"enable_mfa_sms":   12,
}

func (x LoginType) String() string {
	return proto.EnumName(LoginType_name, int32(x))
}

func (LoginType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_56667f8afe262146, []int{0}
}

type GetSessionsRequest struct {
	BaseListRequest      `protobuf:"bytes,1,opt,name=base_list_request,json=baseListRequest,proto3,embedded=base_list_request" json:"base_list_request"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"userId,omitempty" valid:"required,uuid"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSessionsRequest) Reset()         { *m = GetSessionsRequest{} }
func (m *GetSessionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetSessionsRequest) ProtoMessage()    {}
func (*GetSessionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56667f8afe262146, []int{0}
}
func (m *GetSessionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionsRequest.Unmarshal(m, b)
}
func (m *GetSessionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionsRequest.Marshal(b, m, deterministic)
}
func (m *GetSessionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionsRequest.Merge(m, src)
}
func (m *GetSessionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetSessionsRequest.Size(m)
}
func (m *GetSessionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionsRequest proto.InternalMessageInfo

func (m *GetSessionsRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type SessionInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" valid:"required,uuid"`
	CreateTime           string   `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"createTime" valid:"required"`
	Expiry               string   `protobuf:"bytes,3,opt,name=expiry,proto3" json:"expiry" valid:"required"`
	LastSeen             string   `protobuf:"bytes,4,opt,name=last_seen,json=lastSeen,proto3" json:"lastSeen,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionInfo) Reset()         { *m = SessionInfo{} }
func (m *SessionInfo) String() string { return proto.CompactTextString(m) }
func (*SessionInfo) ProtoMessage()    {}
func (*SessionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_56667f8afe262146, []int{1}
}
func (m *SessionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionInfo.Unmarshal(m, b)
}
func (m *SessionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionInfo.Marshal(b, m, deterministic)
}
func (m *SessionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionInfo.Merge(m, src)
}
func (m *SessionInfo) XXX_Size() int {
	return xxx_messageInfo_SessionInfo.Size(m)
}
func (m *SessionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SessionInfo proto.InternalMessageInfo

func (m *SessionInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SessionInfo) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *SessionInfo) GetExpiry() string {
	if m != nil {
		return m.Expiry
	}
	return ""
}

func (m *SessionInfo) GetLastSeen() string {
	if m != nil {
		return m.LastSeen
	}
	return ""
}

type GetSessionsResponse struct {
	BaseListResponse     `protobuf:"bytes,1,opt,name=base_list_response,json=baseListResponse,proto3,embedded=base_list_response" json:",omitempty"`
	Data                 []*SessionInfo `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetSessionsResponse) Reset()         { *m = GetSessionsResponse{} }
func (m *GetSessionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetSessionsResponse) ProtoMessage()    {}
func (*GetSessionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56667f8afe262146, []int{2}
}
func (m *GetSessionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionsResponse.Unmarshal(m, b)
}
func (m *GetSessionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionsResponse.Marshal(b, m, deterministic)
}
func (m *GetSessionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionsResponse.Merge(m, src)
}
func (m *GetSessionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetSessionsResponse.Size(m)
}
func (m *GetSessionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionsResponse proto.InternalMessageInfo

func (m *GetSessionsResponse) GetData() []*SessionInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeleteSessionRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" valid:"required,uuid"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSessionRequest) Reset()         { *m = DeleteSessionRequest{} }
func (m *DeleteSessionRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSessionRequest) ProtoMessage()    {}
func (*DeleteSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56667f8afe262146, []int{3}
}
func (m *DeleteSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSessionRequest.Unmarshal(m, b)
}
func (m *DeleteSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSessionRequest.Marshal(b, m, deterministic)
}
func (m *DeleteSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSessionRequest.Merge(m, src)
}
func (m *DeleteSessionRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSessionRequest.Size(m)
}
func (m *DeleteSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSessionRequest proto.InternalMessageInfo

func (m *DeleteSessionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AuthenticationRequest struct {
	Username             string                                                       `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string                                                       `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	AuthMethod           models.AuthMeta_Method                                       `protobuf:"varint,3,opt,name=auth_method,json=authMethod,proto3,enum=fuck_web.service.models.AuthMeta_Method" json:"authMethod,omitempty"`
	AuthAlgorithm        github_com_MicroOps_cn_fuck_web_pkg_utils_sign.AuthAlgorithm `protobuf:"bytes,4,opt,name=auth_algorithm,json=authAlgorithm,proto3,customtype=github.com/MicroOps-cn/fuck-web/pkg/utils/sign.AuthAlgorithm" json:"authAlgorithm,omitempty"`
	AuthKey              string                                                       `protobuf:"bytes,5,opt,name=auth_key,json=authKey,proto3" json:"authKey,omitempty"`
	AuthSecret           string                                                       `protobuf:"bytes,6,opt,name=auth_secret,json=authSecret,proto3" json:"authSecret,omitempty"`
	AuthSign             string                                                       `protobuf:"bytes,7,opt,name=auth_sign,json=authSign,proto3" json:"authSign,omitempty"`
	Payload              string                                                       `protobuf:"bytes,8,opt,name=payload,proto3" json:"-"`
	XXX_NoUnkeyedLiteral struct{}                                                     `json:"-"`
	XXX_unrecognized     []byte                                                       `json:"-"`
	XXX_sizecache        int32                                                        `json:"-"`
}

func (m *AuthenticationRequest) Reset()         { *m = AuthenticationRequest{} }
func (m *AuthenticationRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticationRequest) ProtoMessage()    {}
func (*AuthenticationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56667f8afe262146, []int{4}
}
func (m *AuthenticationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationRequest.Unmarshal(m, b)
}
func (m *AuthenticationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationRequest.Marshal(b, m, deterministic)
}
func (m *AuthenticationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationRequest.Merge(m, src)
}
func (m *AuthenticationRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticationRequest.Size(m)
}
func (m *AuthenticationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationRequest proto.InternalMessageInfo

func (m *AuthenticationRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthenticationRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AuthenticationRequest) GetAuthMethod() models.AuthMeta_Method {
	if m != nil {
		return m.AuthMethod
	}
	return models.AuthMeta_basic
}

func (m *AuthenticationRequest) GetAuthKey() string {
	if m != nil {
		return m.AuthKey
	}
	return ""
}

func (m *AuthenticationRequest) GetAuthSecret() string {
	if m != nil {
		return m.AuthSecret
	}
	return ""
}

func (m *AuthenticationRequest) GetAuthSign() string {
	if m != nil {
		return m.AuthSign
	}
	return ""
}

func (m *AuthenticationRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type UserLoginRequest struct {
	Username             string                                                    `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string                                                    `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty" valid:"email"`
	Phone                string                                                    `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Password             github_com_MicroOps_cn_fuck_web_pkg_service_models.Secret `protobuf:"bytes,4,opt,name=password,proto3,customtype=github.com/MicroOps-cn/fuck-web/pkg/service/models.Secret" json:"password,omitempty"`
	AutoLogin            bool                                                      `protobuf:"varint,5,opt,name=auto_login,json=autoLogin,proto3" json:"autoLogin,omitempty"`
	Type                 LoginType                                                 `protobuf:"varint,6,opt,name=type,proto3,enum=fuck_web.endpoint.LoginType" json:"type,omitempty"`
	Code                 string                                                    `protobuf:"bytes,7,opt,name=code,proto3" json:"code,omitempty"`
	Token                string                                                    `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
	FirstCode            string                                                    `protobuf:"bytes,9,opt,name=first_code,json=firstCode,proto3" json:"firstCode,omitempty"`
	SecondCode           string                                                    `protobuf:"bytes,10,opt,name=second_code,json=secondCode,proto3" json:"secondCode,omitempty"`
	BindingToken         string                                                    `protobuf:"bytes,11,opt,name=binding_token,json=bindingToken,proto3" json:"bindingToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                  `json:"-"`
	XXX_unrecognized     []byte                                                    `json:"-"`
	XXX_sizecache        int32                                                     `json:"-"`
}

func (m *UserLoginRequest) Reset()         { *m = UserLoginRequest{} }
func (m *UserLoginRequest) String() string { return proto.CompactTextString(m) }
func (*UserLoginRequest) ProtoMessage()    {}
func (*UserLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56667f8afe262146, []int{5}
}
func (m *UserLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginRequest.Unmarshal(m, b)
}
func (m *UserLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginRequest.Marshal(b, m, deterministic)
}
func (m *UserLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginRequest.Merge(m, src)
}
func (m *UserLoginRequest) XXX_Size() int {
	return xxx_messageInfo_UserLoginRequest.Size(m)
}
func (m *UserLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginRequest proto.InternalMessageInfo

func (m *UserLoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserLoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserLoginRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *UserLoginRequest) GetAutoLogin() bool {
	if m != nil {
		return m.AutoLogin
	}
	return false
}

func (m *UserLoginRequest) GetType() LoginType {
	if m != nil {
		return m.Type
	}
	return LoginType_normal
}

func (m *UserLoginRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *UserLoginRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UserLoginRequest) GetFirstCode() string {
	if m != nil {
		return m.FirstCode
	}
	return ""
}

func (m *UserLoginRequest) GetSecondCode() string {
	if m != nil {
		return m.SecondCode
	}
	return ""
}

func (m *UserLoginRequest) GetBindingToken() string {
	if m != nil {
		return m.BindingToken
	}
	return ""
}

type UserLoginResponseData struct {
	Token                string      `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	NextMethod           []LoginType `protobuf:"varint,2,rep,packed,name=next_method,json=nextMethod,proto3,enum=fuck_web.endpoint.LoginType" json:"nextMethod"`
	Email                string      `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber          string      `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserLoginResponseData) Reset()         { *m = UserLoginResponseData{} }
func (m *UserLoginResponseData) String() string { return proto.CompactTextString(m) }
func (*UserLoginResponseData) ProtoMessage()    {}
func (*UserLoginResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_56667f8afe262146, []int{6}
}
func (m *UserLoginResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginResponseData.Unmarshal(m, b)
}
func (m *UserLoginResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginResponseData.Marshal(b, m, deterministic)
}
func (m *UserLoginResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginResponseData.Merge(m, src)
}
func (m *UserLoginResponseData) XXX_Size() int {
	return xxx_messageInfo_UserLoginResponseData.Size(m)
}
func (m *UserLoginResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginResponseData proto.InternalMessageInfo

func (m *UserLoginResponseData) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UserLoginResponseData) GetNextMethod() []LoginType {
	if m != nil {
		return m.NextMethod
	}
	return nil
}

func (m *UserLoginResponseData) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserLoginResponseData) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

type UserLoginResponse struct {
	BaseResponse         `protobuf:"bytes,1,opt,name=base_response,json=baseResponse,proto3,embedded=base_response" json:",omitempty"`
	Data                 *UserLoginResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *UserLoginResponse) Reset()         { *m = UserLoginResponse{} }
func (m *UserLoginResponse) String() string { return proto.CompactTextString(m) }
func (*UserLoginResponse) ProtoMessage()    {}
func (*UserLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56667f8afe262146, []int{7}
}
func (m *UserLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginResponse.Unmarshal(m, b)
}
func (m *UserLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginResponse.Marshal(b, m, deterministic)
}
func (m *UserLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginResponse.Merge(m, src)
}
func (m *UserLoginResponse) XXX_Size() int {
	return xxx_messageInfo_UserLoginResponse.Size(m)
}
func (m *UserLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginResponse proto.InternalMessageInfo

func (m *UserLoginResponse) GetData() *UserLoginResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type SendLoginCaptchaRequest struct {
	Username             string    `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Type                 LoginType `protobuf:"varint,2,opt,name=type,proto3,enum=fuck_web.endpoint.LoginType" json:"type" valid:"required"`
	Email                string    `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty" valid:"email"`
	Phone                string    `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SendLoginCaptchaRequest) Reset()         { *m = SendLoginCaptchaRequest{} }
func (m *SendLoginCaptchaRequest) String() string { return proto.CompactTextString(m) }
func (*SendLoginCaptchaRequest) ProtoMessage()    {}
func (*SendLoginCaptchaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56667f8afe262146, []int{8}
}
func (m *SendLoginCaptchaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendLoginCaptchaRequest.Unmarshal(m, b)
}
func (m *SendLoginCaptchaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendLoginCaptchaRequest.Marshal(b, m, deterministic)
}
func (m *SendLoginCaptchaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendLoginCaptchaRequest.Merge(m, src)
}
func (m *SendLoginCaptchaRequest) XXX_Size() int {
	return xxx_messageInfo_SendLoginCaptchaRequest.Size(m)
}
func (m *SendLoginCaptchaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendLoginCaptchaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendLoginCaptchaRequest proto.InternalMessageInfo

func (m *SendLoginCaptchaRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SendLoginCaptchaRequest) GetType() LoginType {
	if m != nil {
		return m.Type
	}
	return LoginType_normal
}

func (m *SendLoginCaptchaRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SendLoginCaptchaRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type SendLoginCaptchaResponseData struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendLoginCaptchaResponseData) Reset()         { *m = SendLoginCaptchaResponseData{} }
func (m *SendLoginCaptchaResponseData) String() string { return proto.CompactTextString(m) }
func (*SendLoginCaptchaResponseData) ProtoMessage()    {}
func (*SendLoginCaptchaResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_56667f8afe262146, []int{9}
}
func (m *SendLoginCaptchaResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendLoginCaptchaResponseData.Unmarshal(m, b)
}
func (m *SendLoginCaptchaResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendLoginCaptchaResponseData.Marshal(b, m, deterministic)
}
func (m *SendLoginCaptchaResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendLoginCaptchaResponseData.Merge(m, src)
}
func (m *SendLoginCaptchaResponseData) XXX_Size() int {
	return xxx_messageInfo_SendLoginCaptchaResponseData.Size(m)
}
func (m *SendLoginCaptchaResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_SendLoginCaptchaResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_SendLoginCaptchaResponseData proto.InternalMessageInfo

func (m *SendLoginCaptchaResponseData) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SendLoginCaptchaResponse struct {
	BaseResponse         `protobuf:"bytes,1,opt,name=base_response,json=baseResponse,proto3,embedded=base_response" json:",omitempty"`
	Data                 *SendLoginCaptchaResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SendLoginCaptchaResponse) Reset()         { *m = SendLoginCaptchaResponse{} }
func (m *SendLoginCaptchaResponse) String() string { return proto.CompactTextString(m) }
func (*SendLoginCaptchaResponse) ProtoMessage()    {}
func (*SendLoginCaptchaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56667f8afe262146, []int{10}
}
func (m *SendLoginCaptchaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendLoginCaptchaResponse.Unmarshal(m, b)
}
func (m *SendLoginCaptchaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendLoginCaptchaResponse.Marshal(b, m, deterministic)
}
func (m *SendLoginCaptchaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendLoginCaptchaResponse.Merge(m, src)
}
func (m *SendLoginCaptchaResponse) XXX_Size() int {
	return xxx_messageInfo_SendLoginCaptchaResponse.Size(m)
}
func (m *SendLoginCaptchaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendLoginCaptchaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendLoginCaptchaResponse proto.InternalMessageInfo

func (m *SendLoginCaptchaResponse) GetData() *SendLoginCaptchaResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type OAuthLoginRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" valid:"required"`
	State                string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuthLoginRequest) Reset()         { *m = OAuthLoginRequest{} }
func (m *OAuthLoginRequest) String() string { return proto.CompactTextString(m) }
func (*OAuthLoginRequest) ProtoMessage()    {}
func (*OAuthLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56667f8afe262146, []int{11}
}
func (m *OAuthLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthLoginRequest.Unmarshal(m, b)
}
func (m *OAuthLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthLoginRequest.Marshal(b, m, deterministic)
}
func (m *OAuthLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthLoginRequest.Merge(m, src)
}
func (m *OAuthLoginRequest) XXX_Size() int {
	return xxx_messageInfo_OAuthLoginRequest.Size(m)
}
func (m *OAuthLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthLoginRequest proto.InternalMessageInfo

func (m *OAuthLoginRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OAuthLoginRequest) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *OAuthLoginRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func init() {
	proto.RegisterEnum("fuck_web.endpoint.LoginType", LoginType_name, LoginType_value)
	proto.RegisterType((*GetSessionsRequest)(nil), "fuck_web.endpoint.GetSessionsRequest")
	proto.RegisterType((*SessionInfo)(nil), "fuck_web.endpoint.SessionInfo")
	proto.RegisterType((*GetSessionsResponse)(nil), "fuck_web.endpoint.GetSessionsResponse")
	proto.RegisterType((*DeleteSessionRequest)(nil), "fuck_web.endpoint.DeleteSessionRequest")
	proto.RegisterType((*AuthenticationRequest)(nil), "fuck_web.endpoint.AuthenticationRequest")
	proto.RegisterType((*UserLoginRequest)(nil), "fuck_web.endpoint.UserLoginRequest")
	proto.RegisterType((*UserLoginResponseData)(nil), "fuck_web.endpoint.UserLoginResponseData")
	proto.RegisterType((*UserLoginResponse)(nil), "fuck_web.endpoint.UserLoginResponse")
	proto.RegisterType((*SendLoginCaptchaRequest)(nil), "fuck_web.endpoint.SendLoginCaptchaRequest")
	proto.RegisterType((*SendLoginCaptchaResponseData)(nil), "fuck_web.endpoint.SendLoginCaptchaResponseData")
	proto.RegisterType((*SendLoginCaptchaResponse)(nil), "fuck_web.endpoint.SendLoginCaptchaResponse")
	proto.RegisterType((*OAuthLoginRequest)(nil), "fuck_web.endpoint.OAuthLoginRequest")
}

func init() { proto.RegisterFile("endpoints/sessions.proto", fileDescriptor_56667f8afe262146) }

var fileDescriptor_56667f8afe262146 = []byte{
	// 1214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xcf, 0x6f, 0x1b, 0xc5,
	0x17, 0xef, 0x3a, 0x4e, 0x62, 0x3f, 0x27, 0xa9, 0x33, 0x71, 0x5a, 0x2b, 0xaa, 0xba, 0xfd, 0xee,
	0xf7, 0xd0, 0xa8, 0x6a, 0xec, 0x2a, 0x85, 0xa2, 0x42, 0x25, 0x54, 0xa7, 0xa2, 0xaa, 0x68, 0x5a,
	0xb4, 0x29, 0x02, 0x21, 0xa4, 0xd5, 0xd8, 0xfb, 0x62, 0x8f, 0xe2, 0xdd, 0x59, 0x76, 0xc6, 0x6d,
	0x7d, 0xe3, 0xc4, 0xdf, 0xc1, 0x99, 0x0b, 0x17, 0xe0, 0xc6, 0x8d, 0x03, 0x47, 0xce, 0x1c, 0xf6,
	0x0f, 0xf0, 0x09, 0xf5, 0x2f, 0x40, 0xf3, 0x63, 0xed, 0x75, 0xe2, 0x90, 0x70, 0xe0, 0xe4, 0x9d,
	0xcf, 0xbc, 0xf7, 0xe6, 0xbd, 0xf7, 0xf9, 0xcc, 0x1b, 0x43, 0x13, 0xe3, 0x30, 0xe1, 0x2c, 0x96,
	0xa2, 0x2d, 0x50, 0x08, 0xc6, 0x63, 0xd1, 0x4a, 0x52, 0x2e, 0x39, 0xd9, 0x3c, 0x1e, 0xf5, 0x4e,
	0x82, 0x37, 0xd8, 0x6d, 0xe5, 0x26, 0x3b, 0x8d, 0x3e, 0xef, 0x73, 0xbd, 0xdb, 0x56, 0x5f, 0xc6,
	0x70, 0xa7, 0x31, 0x0b, 0xd1, 0xa5, 0x02, 0x2d, 0xba, 0x15, 0xf1, 0x10, 0x87, 0xa2, 0x6d, 0x7e,
	0x0c, 0xe8, 0xfd, 0xe2, 0x00, 0x79, 0x8a, 0xf2, 0xc8, 0x9e, 0xe4, 0xe3, 0x37, 0x23, 0x14, 0x92,
	0x7c, 0x09, 0x9b, 0xca, 0x33, 0x18, 0x32, 0x21, 0x83, 0xd4, 0x80, 0x4d, 0xe7, 0x96, 0xb3, 0x5b,
	0xdb, 0xf7, 0x5a, 0x67, 0xd2, 0x68, 0x75, 0xa8, 0xc0, 0xe7, 0x4c, 0x48, 0xeb, 0xde, 0xa9, 0xfc,
	0x9e, 0xb9, 0x57, 0xfe, 0xc8, 0x5c, 0xc7, 0xbf, 0xda, 0x9d, 0xdf, 0x22, 0x9f, 0xc0, 0xea, 0x48,
	0x60, 0x1a, 0xb0, 0xb0, 0x59, 0xba, 0xe5, 0xec, 0x56, 0x3b, 0x7b, 0x93, 0xcc, 0xad, 0x2b, 0xe8,
	0x59, 0x78, 0x97, 0x47, 0x4c, 0x62, 0x94, 0xc8, 0xf1, 0xbb, 0xcc, 0xdd, 0x7e, 0x4d, 0x87, 0x2c,
	0xfc, 0xd0, 0x53, 0xa7, 0xb3, 0x14, 0xc3, 0xbb, 0xa3, 0x11, 0x0b, 0x3d, 0x7f, 0xc5, 0x98, 0x7a,
	0x7f, 0x39, 0x50, 0xb3, 0x59, 0x3f, 0x8b, 0x8f, 0x39, 0x69, 0x43, 0x89, 0x85, 0x3a, 0xc5, 0x6a,
	0xc7, 0x9d, 0x64, 0x6e, 0x89, 0x85, 0xe7, 0x07, 0x29, 0xb1, 0x90, 0x1c, 0x40, 0xad, 0x97, 0x22,
	0x95, 0x18, 0x48, 0x16, 0xa1, 0x4d, 0xc6, 0x9b, 0x64, 0x2e, 0x18, 0xf8, 0x15, 0x8b, 0xf0, 0x5d,
	0xe6, 0xd6, 0x4f, 0x45, 0xf0, 0xfc, 0xc2, 0x3e, 0x79, 0x00, 0x2b, 0xf8, 0x36, 0x61, 0xe9, 0xb8,
	0xb9, 0xa4, 0xfd, 0x6f, 0x4e, 0x32, 0xd7, 0x22, 0x0b, 0x7d, 0xed, 0x1e, 0xb9, 0x0f, 0xd5, 0x21,
	0x15, 0x32, 0x10, 0x88, 0x71, 0xb3, 0xac, 0x5d, 0xaf, 0x4d, 0x32, 0x97, 0x28, 0xf0, 0x08, 0x31,
	0x9e, 0x75, 0xc2, 0xaf, 0xe4, 0x98, 0xf7, 0x93, 0x03, 0x5b, 0x73, 0x5c, 0x89, 0x84, 0xc7, 0x02,
	0x09, 0x03, 0x52, 0x24, 0xcb, 0xa0, 0x96, 0xad, 0xff, 0xff, 0x23, 0x5b, 0xc6, 0xb4, 0x73, 0x2d,
	0xa7, 0x4b, 0x55, 0x5f, 0x38, 0xba, 0xde, 0x3d, 0x65, 0x49, 0xf6, 0xa1, 0x1c, 0x52, 0x49, 0x9b,
	0xa5, 0x5b, 0x4b, 0xbb, 0xb5, 0xfd, 0x9b, 0x0b, 0x82, 0x17, 0x38, 0xf1, 0xb5, 0xad, 0xf7, 0x14,
	0x1a, 0x4f, 0x70, 0x88, 0x12, 0xed, 0x56, 0xae, 0x84, 0x7f, 0xcb, 0x98, 0xf7, 0x6d, 0x19, 0xb6,
	0x1f, 0x8f, 0xe4, 0x00, 0x63, 0xc9, 0x7a, 0x54, 0x16, 0x42, 0xed, 0x40, 0x45, 0xc9, 0x22, 0xa6,
	0x91, 0xa9, 0xbb, 0xea, 0x4f, 0xd7, 0x6a, 0x2f, 0xa1, 0x42, 0xbc, 0xe1, 0xa9, 0x55, 0x9c, 0x3f,
	0x5d, 0x93, 0x1e, 0xd4, 0xe8, 0x48, 0x0e, 0x82, 0x08, 0xe5, 0x80, 0x87, 0x9a, 0xc3, 0x8d, 0xfd,
	0xdd, 0x59, 0x55, 0x02, 0xd3, 0xd7, 0xac, 0x87, 0x2d, 0x7b, 0x65, 0xd4, 0xe1, 0x87, 0x28, 0x69,
	0xeb, 0x50, 0xdb, 0x77, 0x9a, 0x93, 0xcc, 0x6d, 0x50, 0x03, 0x0e, 0x78, 0x41, 0xbe, 0x3e, 0xcc,
	0x50, 0xf2, 0x9d, 0x03, 0x1b, 0xfa, 0x14, 0x3a, 0xec, 0xf3, 0x94, 0xc9, 0x41, 0x64, 0x19, 0x0f,
	0x54, 0xdb, 0xff, 0xcc, 0xdc, 0x47, 0x7d, 0x26, 0x07, 0xa3, 0x6e, 0xab, 0xc7, 0xa3, 0xf6, 0x21,
	0xeb, 0xa5, 0xfc, 0x65, 0x22, 0xf6, 0x7a, 0x71, 0x5b, 0xa5, 0xb1, 0xf7, 0x06, 0xbb, 0xed, 0xe4,
	0xa4, 0xdf, 0x1e, 0x49, 0x36, 0x14, 0x6d, 0xc1, 0xfa, 0xb1, 0x4e, 0xe3, 0x71, 0x1e, 0x6b, 0x92,
	0xb9, 0xd7, 0x69, 0x11, 0x28, 0x64, 0xb1, 0x3e, 0xb7, 0x41, 0xee, 0x41, 0x45, 0xe7, 0x71, 0x82,
	0xe3, 0xe6, 0xb2, 0xce, 0x60, 0x7b, 0x92, 0xb9, 0x9b, 0x0a, 0xfb, 0x14, 0xc7, 0x05, 0xbf, 0x55,
	0x0b, 0x91, 0x87, 0xb6, 0x3f, 0x02, 0x7b, 0x29, 0xca, 0xe6, 0x8a, 0x76, 0x9a, 0x56, 0x7d, 0xa4,
	0xd1, 0xd3, 0x55, 0x1b, 0x54, 0x29, 0xdc, 0xb8, 0xb2, 0x7e, 0xdc, 0x5c, 0x9d, 0x29, 0x5c, 0x9b,
	0xb0, 0xfe, 0x9c, 0xc2, 0x73, 0x8c, 0xb8, 0xb0, 0x9a, 0xd0, 0xf1, 0x90, 0xd3, 0xb0, 0x59, 0xd1,
	0x2e, 0xcb, 0x93, 0xcc, 0x75, 0xf6, 0xfc, 0x1c, 0xf5, 0x7e, 0x28, 0x43, 0xfd, 0x73, 0x81, 0xe9,
	0x73, 0xde, 0x67, 0x97, 0x62, 0xff, 0x36, 0x2c, 0x63, 0x44, 0xd9, 0xd0, 0xde, 0xef, 0xcd, 0x77,
	0x99, 0xbb, 0x6e, 0x15, 0xa6, 0x71, 0xcf, 0x37, 0xfb, 0xa4, 0x01, 0xcb, 0xc9, 0x80, 0xc7, 0x68,
	0x2e, 0xb2, 0x6f, 0x16, 0x44, 0x14, 0xc4, 0x63, 0x48, 0xfb, 0xc2, 0x92, 0xf6, 0xf0, 0x32, 0xa4,
	0x59, 0xfd, 0xe4, 0x23, 0xd7, 0xb4, 0x46, 0x75, 0x21, 0x0f, 0x5a, 0xec, 0xc2, 0x54, 0x95, 0x0f,
	0x40, 0x35, 0x92, 0x07, 0x43, 0x55, 0xa4, 0x66, 0xaa, 0xd2, 0xb9, 0x3e, 0xc9, 0xdc, 0x2d, 0x85,
	0xea, 0xca, 0x0b, 0x6e, 0xd5, 0x29, 0x48, 0xee, 0x41, 0x59, 0x8e, 0x13, 0xd4, 0x34, 0x6d, 0xec,
	0xdf, 0x58, 0x70, 0x39, 0xb5, 0xdd, 0xab, 0x71, 0x82, 0xbe, 0xb6, 0x24, 0x04, 0xca, 0x3d, 0x1e,
	0xa2, 0xe1, 0xc7, 0xd7, 0xdf, 0xaa, 0x11, 0x92, 0x9f, 0x60, 0x6c, 0x18, 0xf0, 0xcd, 0x42, 0xe5,
	0x74, 0xcc, 0x52, 0x21, 0x03, 0x6d, 0x5f, 0xd5, 0xad, 0xd0, 0x39, 0x69, 0xf4, 0x80, 0x87, 0x58,
	0xcc, 0x69, 0x0a, 0x2a, 0x05, 0x09, 0xec, 0xf1, 0x38, 0x34, 0x8e, 0x30, 0x53, 0x90, 0x81, 0x4f,
	0x79, 0xc2, 0x0c, 0x25, 0x1f, 0xc3, 0x7a, 0x97, 0xc5, 0x21, 0x8b, 0xfb, 0x81, 0x49, 0xa8, 0xa6,
	0x9d, 0x77, 0x26, 0x99, 0x7b, 0xcd, 0x6e, 0xbc, 0x52, 0x78, 0xc1, 0x7d, 0xad, 0x88, 0x7b, 0x3f,
	0x3b, 0xb0, 0x5d, 0x10, 0x8b, 0x19, 0x61, 0x4f, 0xa8, 0xa4, 0xb3, 0x1a, 0x9d, 0x62, 0x8d, 0x87,
	0x50, 0x8b, 0xf1, 0xad, 0xcc, 0xa7, 0x81, 0x9a, 0x71, 0x17, 0xb4, 0xb1, 0xb3, 0xa1, 0x26, 0xa6,
	0x72, 0x32, 0x77, 0xdd, 0x2f, 0x7c, 0xab, 0x43, 0x8c, 0xf4, 0xac, 0xa2, 0x8c, 0xce, 0xfe, 0x07,
	0x6b, 0x5a, 0x5a, 0x41, 0x3c, 0x8a, 0xba, 0x98, 0x1a, 0x55, 0xf9, 0x35, 0x8d, 0xbd, 0xd0, 0x90,
	0xf7, 0xa3, 0x03, 0x9b, 0x67, 0xf2, 0x26, 0x5f, 0xc3, 0xba, 0x9e, 0xf2, 0xa7, 0x06, 0xbc, 0x7b,
	0xce, 0x80, 0xbf, 0x70, 0xb8, 0xaf, 0x75, 0x0b, 0x56, 0xe4, 0xd1, 0x74, 0xb0, 0xab, 0xa0, 0xbb,
	0x0b, 0x82, 0x2e, 0xec, 0xa4, 0x1d, 0xf1, 0xbf, 0x39, 0x70, 0xfd, 0x08, 0xe3, 0x50, 0xef, 0x1f,
	0xd0, 0x44, 0xf6, 0x06, 0xf4, 0x32, 0xb7, 0xf3, 0x85, 0x55, 0x6c, 0xe9, 0x62, 0xc5, 0x76, 0x6e,
	0x4c, 0x32, 0x57, 0x5b, 0x2f, 0x7c, 0x58, 0x8d, 0x9e, 0x6f, 0xcf, 0xb5, 0xfc, 0x32, 0xb7, 0xbd,
	0x5c, 0xb8, 0xed, 0xde, 0x7b, 0x70, 0xe3, 0x6c, 0x15, 0x17, 0xc9, 0xc6, 0xfb, 0xd5, 0x81, 0xe6,
	0x79, 0x6e, 0xff, 0x31, 0x6b, 0x07, 0x73, 0xac, 0xb5, 0x17, 0x3e, 0xc7, 0xe7, 0xd7, 0x63, 0xc9,
	0x63, 0xb0, 0xf9, 0x52, 0x3d, 0x29, 0x73, 0x33, 0xf5, 0x4e, 0xe1, 0x71, 0xde, 0x99, 0x3e, 0xce,
	0x67, 0xfb, 0xae, 0xfe, 0x49, 0x35, 0x60, 0x59, 0x48, 0x2a, 0xed, 0x7f, 0x28, 0xdf, 0x2c, 0xa6,
	0xb3, 0x65, 0x69, 0x36, 0x5b, 0xee, 0x7c, 0xef, 0x40, 0x75, 0xca, 0x28, 0x01, 0x58, 0x89, 0x79,
	0x1a, 0xd1, 0x61, 0xfd, 0x0a, 0x59, 0x83, 0x4a, 0x74, 0x4c, 0x03, 0xc9, 0x65, 0x52, 0x77, 0xc8,
	0x3a, 0x54, 0xd5, 0x4a, 0x73, 0x55, 0x2f, 0x91, 0x1a, 0xac, 0xaa, 0xa5, 0x88, 0x44, 0x7d, 0x89,
	0x54, 0x2d, 0xc7, 0xf5, 0x32, 0x59, 0x85, 0x25, 0x85, 0x2d, 0xab, 0x48, 0x5c, 0x3d, 0x22, 0xfb,
	0xf5, 0x15, 0xb2, 0x05, 0x57, 0x31, 0xa6, 0xdd, 0x21, 0x06, 0xd3, 0x80, 0x40, 0x1a, 0x50, 0x2f,
	0x80, 0xc6, 0xbf, 0x46, 0x08, 0x6c, 0x14, 0x50, 0x15, 0x6a, 0xad, 0xf3, 0xc1, 0x57, 0xef, 0x5f,
	0x66, 0xb2, 0xe7, 0x0d, 0xfe, 0x28, 0xff, 0xf8, 0xec, 0x4a, 0x77, 0x45, 0xff, 0xa5, 0xbe, 0xff,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x50, 0x4c, 0x9f, 0xc2, 0x0b, 0x00, 0x00,
}
