// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

package config

import (
	fmt "fmt"
	_ "github.com/MicroOps-cn/fuck-web/pkg/client/email"
	github_com_MicroOps_cn_fuck_web_pkg_client_email "github.com/MicroOps-cn/fuck-web/pkg/client/email"
	_ "github.com/MicroOps-cn/fuck-web/pkg/client/geoip"
	github_com_MicroOps_cn_fuck_web_pkg_client_geoip "github.com/MicroOps-cn/fuck-web/pkg/client/geoip"
	_ "github.com/MicroOps-cn/fuck-web/pkg/client/ldap"
	github_com_MicroOps_cn_fuck_web_pkg_client_ldap "github.com/MicroOps-cn/fuck-web/pkg/client/ldap"
	oauth2 "github.com/MicroOps-cn/fuck-web/pkg/client/oauth2"
	github_com_MicroOps_cn_fuck_capacity "github.com/MicroOps-cn/fuck/capacity"
	github_com_MicroOps_cn_fuck_clients_gorm "github.com/MicroOps-cn/fuck/clients/gorm"
	github_com_MicroOps_cn_fuck_clients_redis "github.com/MicroOps-cn/fuck/clients/redis"
	github_com_MicroOps_cn_fuck_clients_tracing "github.com/MicroOps-cn/fuck/clients/tracing"
	github_com_MicroOps_cn_fuck_jwt "github.com/MicroOps-cn/fuck/jwt"
	github_com_MicroOps_cn_fuck_sets "github.com/MicroOps-cn/fuck/sets"
	github_com_MicroOps_cn_fuck_wrapper "github.com/MicroOps-cn/fuck/wrapper"
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

type LoginType int32

const (
	LoginType_normal LoginType = 0
	LoginType_email  LoginType = 4
	LoginType_sms    LoginType = 5
	LoginType_oauth2 LoginType = 6
)

var LoginType_name = map[int32]string{
	0: "normal",
	4: "email",
	5: "sms",
	6: "oauth2",
}

var LoginType_value = map[string]int32{
	"normal": 0,
	"email":  4,
	"sms":    5,
	"oauth2": 6,
}

func (x LoginType) String() string {
	return proto.EnumName(LoginType_name, int32(x))
}

func (LoginType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
}

// @sync-to-public:public/src/services/fuck-web/enums.ts:PasswordComplexity
type PasswordComplexity int32

const (
	PasswordComplexity_unsafe    PasswordComplexity = 0
	PasswordComplexity_general   PasswordComplexity = 1
	PasswordComplexity_safe      PasswordComplexity = 2
	PasswordComplexity_very_safe PasswordComplexity = 3
)

var PasswordComplexity_name = map[int32]string{
	0: "unsafe",
	1: "general",
	2: "safe",
	3: "very_safe",
}

var PasswordComplexity_value = map[string]int32{
	"unsafe":    0,
	"general":   1,
	"safe":      2,
	"very_safe": 3,
}

func (x PasswordComplexity) String() string {
	return proto.EnumName(PasswordComplexity_name, int32(x))
}

func (PasswordComplexity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{1}
}

type CustomType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomType) Reset()         { *m = CustomType{} }
func (m *CustomType) String() string { return proto.CompactTextString(m) }
func (*CustomType) ProtoMessage()    {}
func (*CustomType) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
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

type StorageRef struct {
	Storage              *Storage `protobuf:"bytes,1,opt,name=storage,proto3" json:"storage,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StorageRef) Reset()         { *m = StorageRef{} }
func (m *StorageRef) String() string { return proto.CompactTextString(m) }
func (*StorageRef) ProtoMessage()    {}
func (*StorageRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{1}
}
func (m *StorageRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageRef.Unmarshal(m, b)
}
func (m *StorageRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageRef.Marshal(b, m, deterministic)
}
func (m *StorageRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageRef.Merge(m, src)
}
func (m *StorageRef) XXX_Size() int {
	return xxx_messageInfo_StorageRef.Size(m)
}
func (m *StorageRef) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageRef.DiscardUnknown(m)
}

var xxx_messageInfo_StorageRef proto.InternalMessageInfo

func (m *StorageRef) GetStorage() *Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *StorageRef) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *StorageRef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Storage struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to Source:
	//	*Storage_Ref
	//	*Storage_Mysql
	//	*Storage_Redis
	//	*Storage_Ldap
	//	*Storage_Sqlite
	Source               isStorage_Source `protobuf_oneof:"source"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Storage) Reset()         { *m = Storage{} }
func (m *Storage) String() string { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()    {}
func (*Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{2}
}
func (m *Storage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storage.Unmarshal(m, b)
}
func (m *Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storage.Marshal(b, m, deterministic)
}
func (m *Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage.Merge(m, src)
}
func (m *Storage) XXX_Size() int {
	return xxx_messageInfo_Storage.Size(m)
}
func (m *Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_Storage proto.InternalMessageInfo

type isStorage_Source interface {
	isStorage_Source()
}

type Storage_Ref struct {
	Ref *StorageRef `protobuf:"bytes,10,opt,name=ref,proto3,oneof" json:"ref,omitempty"`
}
type Storage_Mysql struct {
	Mysql *github_com_MicroOps_cn_fuck_clients_gorm.MySQLClient `protobuf:"bytes,11,opt,name=mysql,proto3,oneof,customtype=github.com/MicroOps-cn/fuck/clients/gorm.MySQLClient" json:"mysql,omitempty"`
}
type Storage_Redis struct {
	Redis *github_com_MicroOps_cn_fuck_clients_redis.Client `protobuf:"bytes,12,opt,name=redis,proto3,oneof,customtype=github.com/MicroOps-cn/fuck/clients/redis.Client" json:"redis,omitempty"`
}
type Storage_Ldap struct {
	Ldap *github_com_MicroOps_cn_fuck_web_pkg_client_ldap.Client `protobuf:"bytes,13,opt,name=ldap,proto3,oneof,customtype=github.com/MicroOps-cn/fuck-web/pkg/client/ldap.Client" json:"ldap,omitempty"`
}
type Storage_Sqlite struct {
	Sqlite *github_com_MicroOps_cn_fuck_clients_gorm.SQLiteClient `protobuf:"bytes,14,opt,name=sqlite,proto3,oneof,customtype=github.com/MicroOps-cn/fuck/clients/gorm.SQLiteClient" json:"sqlite,omitempty"`
}

func (*Storage_Ref) isStorage_Source()    {}
func (*Storage_Mysql) isStorage_Source()  {}
func (*Storage_Redis) isStorage_Source()  {}
func (*Storage_Ldap) isStorage_Source()   {}
func (*Storage_Sqlite) isStorage_Source() {}

func (m *Storage) GetSource() isStorage_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Storage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Storage) GetRef() *StorageRef {
	if x, ok := m.GetSource().(*Storage_Ref); ok {
		return x.Ref
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Storage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Storage_Ref)(nil),
		(*Storage_Mysql)(nil),
		(*Storage_Redis)(nil),
		(*Storage_Ldap)(nil),
		(*Storage_Sqlite)(nil),
	}
}

type Storages struct {
	Default              *Storage                                                 `protobuf:"bytes,1,opt,name=default,proto3" json:"default,omitempty"`
	Session              *Storage                                                 `protobuf:"bytes,2,opt,name=session,proto3" json:"session,omitempty"`
	User                 *Storage                                                 `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Logging              *Storage                                                 `protobuf:"bytes,4,opt,name=logging,proto3" json:"logging,omitempty"`
	Geoip                *github_com_MicroOps_cn_fuck_web_pkg_client_geoip.Client `protobuf:"bytes,15,opt,name=geoip,proto3,customtype=github.com/MicroOps-cn/fuck-web/pkg/client/geoip.Client" json:"geoip,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                 `json:"-"`
	XXX_unrecognized     []byte                                                   `json:"-"`
	XXX_sizecache        int32                                                    `json:"-"`
}

func (m *Storages) Reset()         { *m = Storages{} }
func (m *Storages) String() string { return proto.CompactTextString(m) }
func (*Storages) ProtoMessage()    {}
func (*Storages) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{3}
}
func (m *Storages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storages.Unmarshal(m, b)
}
func (m *Storages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storages.Marshal(b, m, deterministic)
}
func (m *Storages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storages.Merge(m, src)
}
func (m *Storages) XXX_Size() int {
	return xxx_messageInfo_Storages.Size(m)
}
func (m *Storages) XXX_DiscardUnknown() {
	xxx_messageInfo_Storages.DiscardUnknown(m)
}

var xxx_messageInfo_Storages proto.InternalMessageInfo

func (m *Storages) GetDefault() *Storage {
	if m != nil {
		return m.Default
	}
	return nil
}

func (m *Storages) GetSession() *Storage {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *Storages) GetUser() *Storage {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Storages) GetLogging() *Storage {
	if m != nil {
		return m.Logging
	}
	return nil
}

type GlobalOptions struct {
	MaxUploadSize        github_com_MicroOps_cn_fuck_capacity.Capacities `protobuf:"bytes,1,opt,name=max_upload_size,json=maxUploadSize,proto3,customtype=github.com/MicroOps-cn/fuck/capacity.Capacities" json:"max_upload_size"`
	MaxBodySize          github_com_MicroOps_cn_fuck_capacity.Capacities `protobuf:"bytes,2,opt,name=max_body_size,json=maxBodySize,proto3,customtype=github.com/MicroOps-cn/fuck/capacity.Capacities" json:"max_body_size"`
	UploadPath           string                                          `protobuf:"bytes,3,opt,name=upload_path,json=uploadPath,proto3" json:"upload_path,omitempty"`
	Workspace            string                                          `protobuf:"bytes,4,opt,name=workspace,proto3" json:"workspace,omitempty"`
	AppName              string                                          `protobuf:"bytes,7,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	Title                string                                          `protobuf:"bytes,10,opt,name=title,proto3" json:"title,omitempty"`
	SubTitle             string                                          `protobuf:"bytes,11,opt,name=sub_title,json=subTitle,proto3" json:"sub_title,omitempty"`
	Logo                 string                                          `protobuf:"bytes,12,opt,name=logo,proto3" json:"logo,omitempty"`
	Copyright            string                                          `protobuf:"bytes,13,opt,name=copyright,proto3" json:"copyright,omitempty"`
	AdminEmail           string                                          `protobuf:"bytes,14,opt,name=admin_email,json=adminEmail,proto3" json:"admin_email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *GlobalOptions) Reset()         { *m = GlobalOptions{} }
func (m *GlobalOptions) String() string { return proto.CompactTextString(m) }
func (*GlobalOptions) ProtoMessage()    {}
func (*GlobalOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{4}
}
func (m *GlobalOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalOptions.Unmarshal(m, b)
}
func (m *GlobalOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalOptions.Marshal(b, m, deterministic)
}
func (m *GlobalOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalOptions.Merge(m, src)
}
func (m *GlobalOptions) XXX_Size() int {
	return xxx_messageInfo_GlobalOptions.Size(m)
}
func (m *GlobalOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalOptions proto.InternalMessageInfo

func (m *GlobalOptions) GetUploadPath() string {
	if m != nil {
		return m.UploadPath
	}
	return ""
}

func (m *GlobalOptions) GetWorkspace() string {
	if m != nil {
		return m.Workspace
	}
	return ""
}

func (m *GlobalOptions) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *GlobalOptions) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GlobalOptions) GetSubTitle() string {
	if m != nil {
		return m.SubTitle
	}
	return ""
}

func (m *GlobalOptions) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *GlobalOptions) GetCopyright() string {
	if m != nil {
		return m.Copyright
	}
	return ""
}

func (m *GlobalOptions) GetAdminEmail() string {
	if m != nil {
		return m.AdminEmail
	}
	return ""
}

type RateLimit struct {
	Name                 github_com_MicroOps_cn_fuck_wrapper.OneOrMore[string] `protobuf:"bytes,1,opt,name=name,proto3,customtype=github.com/MicroOps-cn/fuck/wrapper.OneOrMore[string]" json:"name"`
	Allower              Limiter                                               `protobuf:"bytes,2,opt,name=allower,proto3,customtype=Limiter" json:"-"`
	Limit                string                                                `protobuf:"bytes,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Burst                int32                                                 `protobuf:"varint,4,opt,name=burst,proto3" json:"burst,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                              `json:"-"`
	XXX_unrecognized     []byte                                                `json:"-"`
	XXX_sizecache        int32                                                 `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{5}
}
func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimit.Unmarshal(m, b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(m, src)
}
func (m *RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimit.Size(m)
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetLimit() string {
	if m != nil {
		return m.Limit
	}
	return ""
}

func (m *RateLimit) GetBurst() int32 {
	if m != nil {
		return m.Burst
	}
	return 0
}

type SecurityOptions struct {
	TrustIp              []github_com_MicroOps_cn_fuck_sets.IPNet   `protobuf:"bytes,1,rep,name=trust_ip,json=trustIp,proto3,customtype=github.com/MicroOps-cn/fuck/sets.IPNet" json:"trust_ip,omitempty"`
	DefaultLoginType     LoginType                                  `protobuf:"varint,2,opt,name=default_login_type,json=defaultLoginType,proto3,enum=fuck_web.config.LoginType" json:"default_login_type,omitempty"`
	AllowLoginType       []LoginType                                `protobuf:"varint,3,rep,packed,name=allow_login_type,json=allowLoginType,proto3,enum=fuck_web.config.LoginType" json:"allow_login_type,omitempty"`
	Oauth2               []*oauth2.Options                          `protobuf:"bytes,8,rep,name=oauth2,proto3" json:"oauth2,omitempty"`
	DisableLoginForm     bool                                       `protobuf:"varint,9,opt,name=disable_login_form,json=disableLoginForm,proto3" json:"disable_login_form,omitempty"`
	Secret               string                                     `protobuf:"bytes,5,opt,name=secret,proto3" json:"secret,omitempty"`
	Jwt                  *github_com_MicroOps_cn_fuck_jwt.JWTConfig `protobuf:"bytes,6,opt,name=jwt,proto3,customtype=github.com/MicroOps-cn/fuck/jwt.JWTConfig" json:"jwt,omitempty"`
	RateLimit            []*RateLimit                               `protobuf:"bytes,7,rep,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *SecurityOptions) Reset()         { *m = SecurityOptions{} }
func (m *SecurityOptions) String() string { return proto.CompactTextString(m) }
func (*SecurityOptions) ProtoMessage()    {}
func (*SecurityOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{6}
}
func (m *SecurityOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityOptions.Unmarshal(m, b)
}
func (m *SecurityOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityOptions.Marshal(b, m, deterministic)
}
func (m *SecurityOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityOptions.Merge(m, src)
}
func (m *SecurityOptions) XXX_Size() int {
	return xxx_messageInfo_SecurityOptions.Size(m)
}
func (m *SecurityOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityOptions.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityOptions proto.InternalMessageInfo

func (m *SecurityOptions) GetDefaultLoginType() LoginType {
	if m != nil {
		return m.DefaultLoginType
	}
	return LoginType_normal
}

func (m *SecurityOptions) GetAllowLoginType() []LoginType {
	if m != nil {
		return m.AllowLoginType
	}
	return nil
}

func (m *SecurityOptions) GetOauth2() []*oauth2.Options {
	if m != nil {
		return m.Oauth2
	}
	return nil
}

func (m *SecurityOptions) GetDisableLoginForm() bool {
	if m != nil {
		return m.DisableLoginForm
	}
	return false
}

func (m *SecurityOptions) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *SecurityOptions) GetRateLimit() []*RateLimit {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

type Config struct {
	Storage              *Storages                                                     `protobuf:"bytes,1,opt,name=storage,proto3" json:"storage,omitempty"`
	Global               *GlobalOptions                                                `protobuf:"bytes,2,opt,name=global,proto3" json:"global,omitempty"`
	Smtp                 *github_com_MicroOps_cn_fuck_web_pkg_client_email.SmtpOptions `protobuf:"bytes,3,opt,name=smtp,proto3,customtype=github.com/MicroOps-cn/fuck-web/pkg/client/email.SmtpOptions" json:"smtp,omitempty"`
	Security             *SecurityOptions                                              `protobuf:"bytes,4,opt,name=security,proto3" json:"security,omitempty"`
	Trace                *github_com_MicroOps_cn_fuck_clients_tracing.TraceOptions     `protobuf:"bytes,5,opt,name=trace,proto3,customtype=github.com/MicroOps-cn/fuck/clients/tracing.TraceOptions" json:"trace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                      `json:"-"`
	XXX_unrecognized     []byte                                                        `json:"-"`
	XXX_sizecache        int32                                                         `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{7}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetStorage() *Storages {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *Config) GetGlobal() *GlobalOptions {
	if m != nil {
		return m.Global
	}
	return nil
}

func (m *Config) GetSecurity() *SecurityOptions {
	if m != nil {
		return m.Security
	}
	return nil
}

type RuntimeSecurityConfig struct {
	ForceEnableMfa              bool               `protobuf:"varint,1,opt,name=force_enable_mfa,json=forceEnableMfa,proto3" json:"forceEnableMfa"`
	PasswordComplexity          PasswordComplexity `protobuf:"varint,2,opt,name=password_complexity,json=passwordComplexity,proto3,enum=fuck_web.config.PasswordComplexity" json:"passwordComplexity"`
	PasswordMinLength           uint32             `protobuf:"varint,3,opt,name=password_min_length,json=passwordMinLength,proto3" json:"passwordMinLength"`
	PasswordExpireTime          uint32             `protobuf:"varint,4,opt,name=password_expire_time,json=passwordExpireTime,proto3" json:"passwordExpireTime"`
	PasswordFailedLockThreshold uint32             `protobuf:"varint,5,opt,name=password_failed_lock_threshold,json=passwordFailedLockThreshold,proto3" json:"passwordFailedLockThreshold"`
	PasswordFailedLockDuration  uint32             `protobuf:"varint,6,opt,name=password_failed_lock_duration,json=passwordFailedLockDuration,proto3" json:"passwordFailedLockDuration"`
	PasswordHistory             uint32             `protobuf:"varint,7,opt,name=password_history,json=passwordHistory,proto3" json:"passwordHistory"`
	AccountInactiveLock         uint32             `protobuf:"varint,8,opt,name=account_inactive_lock,json=accountInactiveLock,proto3" json:"accountInactiveLock"`
	LoginSessionInactivityTime  uint32             `protobuf:"varint,9,opt,name=login_session_inactivity_time,json=loginSessionInactivityTime,proto3" json:"loginSessionInactivityTime"`
	LoginSessionMaxTime         uint32             `protobuf:"varint,10,opt,name=login_session_max_time,json=loginSessionMaxTime,proto3" json:"loginSessionMaxTime"`
	XXX_NoUnkeyedLiteral        struct{}           `json:"-"`
	XXX_unrecognized            []byte             `json:"-"`
	XXX_sizecache               int32              `json:"-"`
}

func (m *RuntimeSecurityConfig) Reset()         { *m = RuntimeSecurityConfig{} }
func (m *RuntimeSecurityConfig) String() string { return proto.CompactTextString(m) }
func (*RuntimeSecurityConfig) ProtoMessage()    {}
func (*RuntimeSecurityConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{8}
}
func (m *RuntimeSecurityConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuntimeSecurityConfig.Unmarshal(m, b)
}
func (m *RuntimeSecurityConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuntimeSecurityConfig.Marshal(b, m, deterministic)
}
func (m *RuntimeSecurityConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeSecurityConfig.Merge(m, src)
}
func (m *RuntimeSecurityConfig) XXX_Size() int {
	return xxx_messageInfo_RuntimeSecurityConfig.Size(m)
}
func (m *RuntimeSecurityConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeSecurityConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeSecurityConfig proto.InternalMessageInfo

func (m *RuntimeSecurityConfig) GetForceEnableMfa() bool {
	if m != nil {
		return m.ForceEnableMfa
	}
	return false
}

func (m *RuntimeSecurityConfig) GetPasswordComplexity() PasswordComplexity {
	if m != nil {
		return m.PasswordComplexity
	}
	return PasswordComplexity_unsafe
}

func (m *RuntimeSecurityConfig) GetPasswordMinLength() uint32 {
	if m != nil {
		return m.PasswordMinLength
	}
	return 0
}

func (m *RuntimeSecurityConfig) GetPasswordExpireTime() uint32 {
	if m != nil {
		return m.PasswordExpireTime
	}
	return 0
}

func (m *RuntimeSecurityConfig) GetPasswordFailedLockThreshold() uint32 {
	if m != nil {
		return m.PasswordFailedLockThreshold
	}
	return 0
}

func (m *RuntimeSecurityConfig) GetPasswordFailedLockDuration() uint32 {
	if m != nil {
		return m.PasswordFailedLockDuration
	}
	return 0
}

func (m *RuntimeSecurityConfig) GetPasswordHistory() uint32 {
	if m != nil {
		return m.PasswordHistory
	}
	return 0
}

func (m *RuntimeSecurityConfig) GetAccountInactiveLock() uint32 {
	if m != nil {
		return m.AccountInactiveLock
	}
	return 0
}

func (m *RuntimeSecurityConfig) GetLoginSessionInactivityTime() uint32 {
	if m != nil {
		return m.LoginSessionInactivityTime
	}
	return 0
}

func (m *RuntimeSecurityConfig) GetLoginSessionMaxTime() uint32 {
	if m != nil {
		return m.LoginSessionMaxTime
	}
	return 0
}

type RuntimeConfig struct {
	Security             *RuntimeSecurityConfig `protobuf:"bytes,1,opt,name=security,proto3" json:"security,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RuntimeConfig) Reset()         { *m = RuntimeConfig{} }
func (m *RuntimeConfig) String() string { return proto.CompactTextString(m) }
func (*RuntimeConfig) ProtoMessage()    {}
func (*RuntimeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{9}
}
func (m *RuntimeConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuntimeConfig.Unmarshal(m, b)
}
func (m *RuntimeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuntimeConfig.Marshal(b, m, deterministic)
}
func (m *RuntimeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeConfig.Merge(m, src)
}
func (m *RuntimeConfig) XXX_Size() int {
	return xxx_messageInfo_RuntimeConfig.Size(m)
}
func (m *RuntimeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeConfig proto.InternalMessageInfo

func (m *RuntimeConfig) GetSecurity() *RuntimeSecurityConfig {
	if m != nil {
		return m.Security
	}
	return nil
}

func init() {
	proto.RegisterEnum("fuck_web.config.LoginType", LoginType_name, LoginType_value)
	proto.RegisterEnum("fuck_web.config.PasswordComplexity", PasswordComplexity_name, PasswordComplexity_value)
	proto.RegisterType((*CustomType)(nil), "fuck_web.config.custom_type")
	proto.RegisterType((*StorageRef)(nil), "fuck_web.config.StorageRef")
	proto.RegisterType((*Storage)(nil), "fuck_web.config.Storage")
	proto.RegisterType((*Storages)(nil), "fuck_web.config.Storages")
	proto.RegisterType((*GlobalOptions)(nil), "fuck_web.config.GlobalOptions")
	proto.RegisterType((*RateLimit)(nil), "fuck_web.config.RateLimit")
	proto.RegisterType((*SecurityOptions)(nil), "fuck_web.config.SecurityOptions")
	proto.RegisterType((*Config)(nil), "fuck_web.config.Config")
	proto.RegisterType((*RuntimeSecurityConfig)(nil), "fuck_web.config.RuntimeSecurityConfig")
	proto.RegisterType((*RuntimeConfig)(nil), "fuck_web.config.RuntimeConfig")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_3eaf2c85e69e9ea4) }

var fileDescriptor_3eaf2c85e69e9ea4 = []byte{
	// 1543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0xdd, 0x72, 0x1b, 0xb7,
	0x15, 0x16, 0xcd, 0x7f, 0x28, 0xb4, 0x18, 0xf8, 0xa7, 0x1b, 0x3b, 0xf5, 0xaa, 0xcc, 0x4c, 0x46,
	0xf5, 0x44, 0x64, 0x47, 0x49, 0xe3, 0xa4, 0x71, 0x3b, 0x1d, 0x3a, 0x72, 0xa4, 0x56, 0x8a, 0x65,
	0x50, 0x9d, 0xce, 0xb4, 0x93, 0xd9, 0x01, 0x97, 0xe0, 0x0a, 0xd6, 0xee, 0x62, 0x0d, 0x60, 0x2d,
	0x31, 0x33, 0xbd, 0xc8, 0x7b, 0xf4, 0x49, 0xfa, 0x04, 0xbd, 0xe8, 0x13, 0xb0, 0x33, 0x7b, 0x5f,
	0x3d, 0x45, 0x07, 0x07, 0x20, 0x29, 0x99, 0x0a, 0xad, 0x5c, 0xe4, 0x8a, 0xc0, 0xf9, 0xfb, 0x70,
	0x70, 0xce, 0xf9, 0xb0, 0x44, 0xef, 0x85, 0x22, 0x1d, 0xf3, 0xa8, 0x9b, 0x49, 0xa1, 0x05, 0xde,
	0x18, 0xe7, 0xe1, 0x69, 0x70, 0xc6, 0x86, 0x5d, 0x2b, 0x7e, 0xd0, 0xd6, 0x93, 0x8c, 0xa9, 0x5e,
	0x3c, 0xa2, 0x99, 0x35, 0x79, 0xf0, 0xbe, 0x95, 0xb0, 0x84, 0xf2, 0xf8, 0xaa, 0x28, 0x62, 0x82,
	0xcf, 0xac, 0xb0, 0x15, 0x09, 0x9a, 0xeb, 0x93, 0x1d, 0x27, 0xbb, 0x1b, 0x89, 0x48, 0xc0, 0xb2,
	0x67, 0x56, 0x56, 0xda, 0x69, 0xa1, 0xf5, 0x30, 0x57, 0x5a, 0x24, 0x81, 0x71, 0xe9, 0x9c, 0x20,
	0x34, 0xd0, 0x42, 0xd2, 0x88, 0x11, 0x36, 0xc6, 0x3b, 0xa8, 0xae, 0xec, 0xce, 0x2b, 0x6d, 0x96,
	0xb6, 0xd6, 0x77, 0xbc, 0xee, 0x5b, 0x27, 0xec, 0xce, 0xac, 0x67, 0x86, 0x18, 0xa3, 0x4a, 0x46,
	0xf5, 0x89, 0x77, 0x6b, 0xb3, 0xb4, 0xd5, 0x24, 0xb0, 0x36, 0xb2, 0x94, 0x26, 0xcc, 0x2b, 0x5b,
	0x99, 0x59, 0x77, 0xfe, 0x59, 0x41, 0xf5, 0xc1, 0xc2, 0x07, 0xf4, 0xa5, 0x85, 0x1e, 0xf7, 0x50,
	0x59, 0xb2, 0xb1, 0x87, 0x00, 0xf7, 0xe1, 0x8f, 0xe2, 0xb2, 0xf1, 0xde, 0x1a, 0x31, 0x96, 0x38,
	0x45, 0xd5, 0x64, 0xa2, 0x5e, 0xc7, 0xde, 0x3a, 0xb8, 0x7c, 0xb8, 0xe4, 0x72, 0x29, 0xcf, 0xfe,
	0x17, 0xd3, 0xc2, 0xff, 0x2c, 0xe2, 0xfa, 0x24, 0x37, 0xea, 0xa4, 0x77, 0xc8, 0x43, 0x29, 0x5e,
	0x64, 0x6a, 0x3b, 0x4c, 0x7b, 0xc6, 0xaf, 0x17, 0xc6, 0x9c, 0xa5, 0x5a, 0xf5, 0x22, 0x21, 0x93,
	0xee, 0xe1, 0x64, 0xf0, 0xf2, 0xe0, 0x19, 0x48, 0xf6, 0xd6, 0x88, 0x85, 0xc1, 0xaf, 0x50, 0x55,
	0xb2, 0x11, 0x57, 0xde, 0x7b, 0x37, 0xc0, 0xfb, 0x6c, 0x5a, 0xf8, 0xbf, 0xb9, 0x09, 0x1e, 0x84,
	0xec, 0x2e, 0xb0, 0x60, 0x8f, 0x15, 0xaa, 0x98, 0x1e, 0xf0, 0x5a, 0x00, 0xf5, 0xab, 0x4b, 0x50,
	0x60, 0xd8, 0x85, 0x06, 0x39, 0x18, 0xd1, 0xec, 0x45, 0xa6, 0xb9, 0x48, 0x55, 0xff, 0x77, 0xd3,
	0xc2, 0xff, 0x7c, 0x05, 0xde, 0xf6, 0x19, 0x1b, 0xf6, 0xb2, 0xd3, 0xc8, 0xe1, 0xda, 0x0e, 0x9b,
	0xa3, 0x02, 0x18, 0x7e, 0x8d, 0x6a, 0xea, 0x75, 0xcc, 0x35, 0xf3, 0x6e, 0xdf, 0x20, 0xc3, 0x2f,
	0xa7, 0x85, 0xff, 0xdb, 0x1b, 0xdf, 0xe8, 0xe0, 0xe5, 0x01, 0xd7, 0x6c, 0x0e, 0xe8, 0x80, 0xfa,
	0x0d, 0x54, 0x53, 0x22, 0x97, 0x21, 0xeb, 0xfc, 0xe7, 0x16, 0x6a, 0xb8, 0x1a, 0x2b, 0xd3, 0x87,
	0x23, 0x36, 0xa6, 0x79, 0xac, 0xdf, 0xdd, 0x87, 0xce, 0x10, 0x7a, 0x97, 0x29, 0xc5, 0x45, 0x0a,
	0xad, 0xb8, 0xba, 0x77, 0xad, 0x21, 0xfe, 0x04, 0x55, 0x72, 0xc5, 0x24, 0xf4, 0xe9, 0x2a, 0x07,
	0xb0, 0x32, 0x08, 0xb1, 0x88, 0x22, 0x9e, 0x46, 0x5e, 0xe5, 0x5d, 0x08, 0xce, 0x10, 0xbf, 0x41,
	0x55, 0x98, 0x53, 0x6f, 0x03, 0x3c, 0x3e, 0x5a, 0xaa, 0xa4, 0x9d, 0xe2, 0x6f, 0x98, 0xd8, 0x3f,
	0x9a, 0xd5, 0xf2, 0xab, 0x69, 0xe1, 0x3f, 0xf9, 0x09, 0xb5, 0xb4, 0x21, 0xec, 0xdd, 0x12, 0x0b,
	0xd7, 0xf9, 0x5f, 0x19, 0xb5, 0xbe, 0x89, 0xc5, 0x90, 0xc6, 0x2e, 0x2a, 0xfe, 0x07, 0xda, 0x48,
	0xe8, 0x79, 0x90, 0x67, 0xb1, 0xa0, 0xa3, 0x40, 0xf1, 0xef, 0x67, 0x33, 0xbe, 0xba, 0xcc, 0x4f,
	0xfe, 0x5d, 0xf8, 0x6b, 0xd3, 0xc2, 0xef, 0xad, 0x2c, 0x35, 0xcd, 0x68, 0xc8, 0xf5, 0xa4, 0xfb,
	0xcc, 0x2e, 0x38, 0x53, 0xa4, 0x95, 0xd0, 0xf3, 0xbf, 0x00, 0xd8, 0x80, 0x7f, 0xcf, 0xf0, 0x04,
	0x19, 0x41, 0x30, 0x14, 0xa3, 0x89, 0x05, 0xbf, 0xf5, 0x73, 0x82, 0xaf, 0x27, 0xf4, 0xbc, 0x2f,
	0x46, 0x13, 0x80, 0xf6, 0xd1, 0xba, 0xcb, 0x1a, 0x88, 0xca, 0x92, 0x12, 0xb2, 0xa2, 0x23, 0x43,
	0x57, 0x1f, 0xa2, 0xe6, 0x99, 0x90, 0xa7, 0x2a, 0xa3, 0x21, 0x83, 0xd2, 0x36, 0xc9, 0x42, 0x80,
	0x3f, 0x40, 0x0d, 0x9a, 0x65, 0x01, 0x10, 0x56, 0x1d, 0x94, 0x75, 0x9a, 0x65, 0xdf, 0x1a, 0xce,
	0xba, 0x8b, 0xaa, 0x9a, 0xeb, 0x98, 0x01, 0x6b, 0x35, 0x89, 0xdd, 0xe0, 0x87, 0xa8, 0xa9, 0xf2,
	0x61, 0x60, 0x35, 0xeb, 0xa0, 0x69, 0xa8, 0x7c, 0x78, 0x0c, 0x4a, 0x8c, 0x2a, 0xb1, 0x88, 0x04,
	0x90, 0x48, 0x93, 0xc0, 0xda, 0xe0, 0x87, 0x22, 0x9b, 0x48, 0x1e, 0x9d, 0x68, 0x18, 0xf9, 0x26,
	0x59, 0x08, 0xcc, 0xf1, 0xe9, 0x28, 0xe1, 0x69, 0x00, 0x6f, 0x00, 0xcc, 0x66, 0x93, 0x20, 0x10,
	0xed, 0x1a, 0x49, 0xe7, 0x5f, 0x25, 0xd4, 0x24, 0x54, 0xb3, 0x03, 0x9e, 0x70, 0x8d, 0x5f, 0x5e,
	0xe6, 0xd6, 0xfe, 0xef, 0xdd, 0x0d, 0xae, 0x9c, 0xd4, 0x33, 0x49, 0xb3, 0x8c, 0xc9, 0xee, 0x8b,
	0x94, 0xbd, 0x90, 0x87, 0x42, 0xb2, 0xbf, 0x2b, 0x2d, 0x79, 0x1a, 0x7d, 0xe7, 0xa8, 0xb9, 0x8b,
	0xea, 0x34, 0x8e, 0xc5, 0x19, 0x93, 0x96, 0xe5, 0xfb, 0x77, 0x5d, 0xd4, 0x3a, 0x40, 0x32, 0x79,
	0x51, 0xf8, 0xa5, 0x6d, 0x32, 0x33, 0x32, 0xd7, 0x12, 0x1b, 0x85, 0xbb, 0x6a, 0xbb, 0x31, 0xd2,
	0x61, 0x2e, 0x95, 0x86, 0x1b, 0xae, 0x12, 0xbb, 0xe9, 0xfc, 0x50, 0x41, 0x1b, 0x03, 0x16, 0xe6,
	0x92, 0xeb, 0xc9, 0xac, 0x55, 0x77, 0x51, 0x43, 0xcb, 0x5c, 0xe9, 0x80, 0x67, 0x5e, 0x69, 0xb3,
	0xbc, 0xd5, 0xec, 0x3f, 0x9e, 0x16, 0xfe, 0xc7, 0xab, 0x52, 0x50, 0x4c, 0xab, 0xee, 0xfe, 0xd1,
	0xb7, 0x4c, 0x93, 0x3a, 0xf8, 0xee, 0x67, 0x78, 0x0f, 0x61, 0x47, 0x0e, 0x41, 0x2c, 0x22, 0x9e,
	0x42, 0x4f, 0x41, 0x06, 0xb7, 0x77, 0x1e, 0x2c, 0xf5, 0xdd, 0x81, 0x31, 0x39, 0x9e, 0x64, 0x8c,
	0xb4, 0x9d, 0xd7, 0x5c, 0x82, 0xbf, 0x46, 0x6d, 0xc8, 0xed, 0x72, 0x9c, 0xf2, 0x66, 0xf9, 0x1d,
	0x71, 0x6e, 0x83, 0xcf, 0x22, 0xca, 0x13, 0x54, 0xb3, 0x0f, 0xb4, 0xd7, 0xd8, 0x2c, 0x6f, 0xad,
	0xef, 0xf8, 0x4b, 0x64, 0xe0, 0xde, 0x6f, 0x77, 0x0f, 0xc4, 0x99, 0xe3, 0x4f, 0x10, 0x1e, 0x71,
	0x45, 0x87, 0x31, 0x73, 0x07, 0x18, 0x0b, 0x99, 0x78, 0xcd, 0xcd, 0xd2, 0x56, 0x83, 0xb4, 0x9d,
	0x06, 0x60, 0x9e, 0x0b, 0x99, 0xe0, 0xfb, 0xa8, 0xa6, 0x58, 0x28, 0x99, 0xf6, 0xaa, 0x70, 0xfd,
	0x6e, 0x87, 0xbf, 0x43, 0xe5, 0x57, 0x67, 0xda, 0xab, 0xdd, 0x60, 0xee, 0xb6, 0xa7, 0x85, 0xff,
	0xeb, 0x55, 0xd7, 0xfd, 0xea, 0x4c, 0x77, 0xff, 0xf4, 0xd7, 0xe3, 0x67, 0xe0, 0x4a, 0x4c, 0x5c,
	0xfc, 0x25, 0x42, 0x92, 0x6a, 0x16, 0xd8, 0xca, 0xd7, 0x21, 0xc3, 0xe5, 0xdb, 0x99, 0xf7, 0x29,
	0x69, 0xca, 0xd9, 0xb2, 0xf3, 0x43, 0x19, 0xd5, 0x6c, 0x28, 0xfc, 0xe9, 0xdb, 0x5f, 0x20, 0x1f,
	0xfc, 0x18, 0xc7, 0xaa, 0xc5, 0x27, 0xc8, 0xe7, 0xa8, 0x16, 0x01, 0xd7, 0x39, 0x52, 0x79, 0xb4,
	0xe4, 0x73, 0x85, 0x0a, 0x89, 0xb3, 0xc6, 0xe7, 0xa8, 0xa2, 0x12, 0x9d, 0x39, 0xfa, 0xef, 0x2c,
	0x95, 0xc3, 0x7e, 0x74, 0x0d, 0x12, 0x3d, 0x7f, 0x66, 0xff, 0x38, 0x2d, 0xfc, 0xa7, 0x3f, 0x81,
	0x9a, 0x97, 0x22, 0x10, 0x40, 0xc4, 0x4f, 0x51, 0x43, 0xb9, 0xa6, 0x77, 0x6f, 0xc9, 0xe6, 0x72,
	0x9e, 0x57, 0xa7, 0x82, 0xcc, 0x3d, 0x30, 0x41, 0x55, 0x2d, 0x0d, 0x57, 0x41, 0x81, 0xfb, 0x4f,
	0xa7, 0x85, 0xff, 0xc5, 0x4d, 0x5e, 0x62, 0xe3, 0xc4, 0xd3, 0xa8, 0x7b, 0x6c, 0x9c, 0x67, 0x61,
	0x6d, 0xa8, 0xce, 0x7f, 0x6b, 0xe8, 0x1e, 0xc9, 0x53, 0xcd, 0x13, 0x36, 0x03, 0x76, 0x25, 0x79,
	0x8a, 0xda, 0x63, 0x21, 0x43, 0x16, 0xb0, 0x14, 0x5a, 0x30, 0x19, 0x53, 0xa8, 0x4d, 0xa3, 0x8f,
	0x2f, 0x0a, 0xff, 0x36, 0xe8, 0x76, 0x41, 0x75, 0x38, 0xa6, 0xe4, 0xad, 0x3d, 0x8e, 0xd1, 0x9d,
	0x8c, 0x2a, 0x75, 0x26, 0xe4, 0x28, 0x08, 0x45, 0x92, 0xc5, 0xec, 0xdc, 0x24, 0x6d, 0xa7, 0xf0,
	0xa3, 0xa5, 0xa4, 0x8f, 0x9c, 0xed, 0xb3, 0xb9, 0x69, 0xff, 0xfe, 0x45, 0xe1, 0xe3, 0x6c, 0x49,
	0x4e, 0xae, 0x91, 0xe1, 0xdd, 0x4b, 0x68, 0x86, 0x32, 0x63, 0x96, 0x46, 0x8e, 0xf2, 0x5b, 0xfd,
	0x7b, 0x17, 0x85, 0xff, 0xfe, 0x4c, 0x7d, 0xc8, 0xd3, 0x03, 0x50, 0x92, 0x65, 0x11, 0xde, 0x43,
	0x77, 0xe7, 0x61, 0xd8, 0x79, 0xc6, 0x25, 0x0b, 0xcc, 0xc5, 0x40, 0xa9, 0x5a, 0x57, 0x0f, 0xb4,
	0x0b, 0xea, 0x63, 0x9e, 0x30, 0x72, 0x8d, 0x0c, 0x8f, 0xd0, 0xa3, 0x79, 0xa4, 0x31, 0xe5, 0x31,
	0x1b, 0x05, 0xb1, 0x08, 0x4f, 0x03, 0x7d, 0x22, 0x99, 0x3a, 0x11, 0xf1, 0x08, 0x6a, 0xd8, 0xea,
	0xfb, 0x17, 0x85, 0xff, 0x70, 0x66, 0xf9, 0x1c, 0x0c, 0x0f, 0x44, 0x78, 0x7a, 0x3c, 0x33, 0x23,
	0xab, 0x94, 0x98, 0xa2, 0x5f, 0x5e, 0x8b, 0x32, 0xca, 0x25, 0x35, 0x55, 0x86, 0xa1, 0x6f, 0xf5,
	0x1f, 0x5d, 0x14, 0xfe, 0x83, 0xe5, 0x38, 0x5f, 0x3b, 0x2b, 0xb2, 0x42, 0x87, 0xff, 0x80, 0xda,
	0x73, 0x88, 0x13, 0x6e, 0x26, 0x6f, 0x02, 0xaf, 0x61, 0xab, 0x7f, 0xe7, 0xa2, 0xf0, 0x37, 0x66,
	0xba, 0x3d, 0xab, 0x22, 0x6f, 0x0b, 0xf0, 0x9f, 0xd1, 0x3d, 0x1a, 0x86, 0x22, 0x4f, 0x75, 0xc0,
	0x53, 0x1a, 0x6a, 0xfe, 0x86, 0xc1, 0x19, 0xbd, 0x06, 0x04, 0xf9, 0xc5, 0x45, 0xe1, 0xdf, 0x71,
	0x06, 0xfb, 0x4e, 0x6f, 0xf0, 0xc9, 0x75, 0x42, 0x93, 0xaf, 0x25, 0x42, 0xf7, 0x21, 0x37, 0x0b,
	0xc9, 0xf5, 0xc4, 0x16, 0xaa, 0xb9, 0xc8, 0x17, 0x0c, 0x07, 0xd6, 0x6e, 0x7f, 0x6e, 0x06, 0x05,
	0x5b, 0xa1, 0xc3, 0x07, 0xe8, 0xfe, 0x55, 0x08, 0xf3, 0xf5, 0x02, 0xb1, 0xd1, 0xe2, 0xc0, 0x97,
	0xfd, 0x0f, 0xe9, 0x39, 0x04, 0xbd, 0x4e, 0xd8, 0x19, 0xa0, 0x96, 0x1b, 0x2e, 0x37, 0x54, 0xfd,
	0x4b, 0x04, 0x60, 0x89, 0xee, 0xe3, 0x65, 0xae, 0xbc, 0x6e, 0x1c, 0x17, 0x34, 0xf0, 0xf8, 0x09,
	0x6a, 0x2e, 0x1e, 0x17, 0x84, 0x6a, 0xa9, 0x90, 0x09, 0x8d, 0xdb, 0x6b, 0xb8, 0x89, 0xaa, 0x40,
	0x3c, 0xed, 0x0a, 0xae, 0xa3, 0xb2, 0x4a, 0x54, 0xbb, 0x6a, 0xf4, 0xf6, 0x35, 0x69, 0xd7, 0x1e,
	0x3f, 0x47, 0x78, 0x79, 0xce, 0x8c, 0x45, 0x9e, 0x2a, 0x3a, 0x66, 0xed, 0x35, 0xbc, 0x8e, 0xea,
	0x11, 0x4b, 0x99, 0xa4, 0x71, 0xbb, 0x84, 0x1b, 0xa8, 0x02, 0xe2, 0x5b, 0xb8, 0x85, 0x9a, 0x6f,
	0x98, 0x9c, 0x04, 0xb0, 0x2d, 0xf7, 0x7b, 0x7f, 0xdb, 0x7e, 0x17, 0x17, 0xda, 0x54, 0xbe, 0xb2,
	0x3f, 0x47, 0x95, 0x61, 0x0d, 0xfe, 0x85, 0x7e, 0xfa, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95,
	0x30, 0xa1, 0x9e, 0x08, 0x0f, 0x00, 0x00,
}
