// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: endpoints/logging.proto

package endpoint

import (
	fmt "fmt"
	_ "github.com/MicroOps-cn/fuck-web/config"
	_ "github.com/MicroOps-cn/fuck-web/pkg/service/models"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// Timestamp from public import google/protobuf/timestamp.proto
type Timestamp = types.Timestamp

type Event struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" valid:"required,uuid"`
	UpdateTime           string   `protobuf:"bytes,2,opt,name=update_time,json=updateTime,proto3" json:"updateTime" valid:"required"`
	CreateTime           string   `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"createTime" valid:"required"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"userId"`
	Username             string   `protobuf:"bytes,5,opt,name=username,proto3" json:"username"`
	Action               string   `protobuf:"bytes,6,opt,name=action,proto3" json:"action"`
	ClientIp             string   `protobuf:"bytes,7,opt,name=client_ip,json=clientIp,proto3" json:"client_ip"`
	Status               string   `protobuf:"bytes,8,opt,name=status,proto3" json:"status"`
	Took                 int64    `protobuf:"varint,9,opt,name=took,proto3" json:"took"`
	Message              string   `protobuf:"bytes,10,opt,name=message,proto3" json:"message"`
	Location             string   `protobuf:"bytes,11,opt,name=location,proto3" json:"location"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_9225461cd3a21153, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *Event) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *Event) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Event) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Event) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Event) GetClientIp() string {
	if m != nil {
		return m.ClientIp
	}
	return ""
}

func (m *Event) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Event) GetTook() int64 {
	if m != nil {
		return m.Took
	}
	return 0
}

func (m *Event) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Event) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type EventLog struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" valid:"required,uuid"`
	UpdateTime           string   `protobuf:"bytes,2,opt,name=update_time,json=updateTime,proto3" json:"updateTime" valid:"required"`
	CreateTime           string   `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"createTime" valid:"required"`
	EventId              string   `protobuf:"bytes,4,opt,name=event_id,json=eventId,proto3" json:"userId"`
	Log                  string   `protobuf:"bytes,5,opt,name=log,proto3" json:"log"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventLog) Reset()         { *m = EventLog{} }
func (m *EventLog) String() string { return proto.CompactTextString(m) }
func (*EventLog) ProtoMessage()    {}
func (*EventLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_9225461cd3a21153, []int{1}
}
func (m *EventLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventLog.Unmarshal(m, b)
}
func (m *EventLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventLog.Marshal(b, m, deterministic)
}
func (m *EventLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventLog.Merge(m, src)
}
func (m *EventLog) XXX_Size() int {
	return xxx_messageInfo_EventLog.Size(m)
}
func (m *EventLog) XXX_DiscardUnknown() {
	xxx_messageInfo_EventLog.DiscardUnknown(m)
}

var xxx_messageInfo_EventLog proto.InternalMessageInfo

func (m *EventLog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EventLog) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *EventLog) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *EventLog) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *EventLog) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

type GetEventsRequest struct {
	BaseListRequest      `protobuf:"bytes,1,opt,name=base_list_request,json=baseListRequest,proto3,embedded=base_list_request" json:",omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Action               string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	StartTime            string   `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"startTime" valid:"required,rfc3339"`
	EndTime              string   `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"endTime" valid:"required,rfc3339"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventsRequest) Reset()         { *m = GetEventsRequest{} }
func (m *GetEventsRequest) String() string { return proto.CompactTextString(m) }
func (*GetEventsRequest) ProtoMessage()    {}
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9225461cd3a21153, []int{2}
}
func (m *GetEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventsRequest.Unmarshal(m, b)
}
func (m *GetEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventsRequest.Marshal(b, m, deterministic)
}
func (m *GetEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventsRequest.Merge(m, src)
}
func (m *GetEventsRequest) XXX_Size() int {
	return xxx_messageInfo_GetEventsRequest.Size(m)
}
func (m *GetEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventsRequest proto.InternalMessageInfo

func (m *GetEventsRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetEventsRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *GetEventsRequest) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *GetEventsRequest) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

type GetEventsResponse struct {
	BaseListResponse     `protobuf:"bytes,1,opt,name=base_list_response,json=baseListResponse,proto3,embedded=base_list_response" json:",omitempty"`
	Data                 []*Event `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventsResponse) Reset()         { *m = GetEventsResponse{} }
func (m *GetEventsResponse) String() string { return proto.CompactTextString(m) }
func (*GetEventsResponse) ProtoMessage()    {}
func (*GetEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9225461cd3a21153, []int{3}
}
func (m *GetEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventsResponse.Unmarshal(m, b)
}
func (m *GetEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventsResponse.Marshal(b, m, deterministic)
}
func (m *GetEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventsResponse.Merge(m, src)
}
func (m *GetEventsResponse) XXX_Size() int {
	return xxx_messageInfo_GetEventsResponse.Size(m)
}
func (m *GetEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventsResponse proto.InternalMessageInfo

func (m *GetEventsResponse) GetData() []*Event {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetCurrentUserEventsRequest struct {
	BaseListRequest      `protobuf:"bytes,1,opt,name=base_list_request,json=baseListRequest,proto3,embedded=base_list_request" json:",omitempty"`
	Action               string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	StartTime            string   `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"startTime" valid:"required,rfc3339"`
	EndTime              string   `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"endTime" valid:"required,rfc3339"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrentUserEventsRequest) Reset()         { *m = GetCurrentUserEventsRequest{} }
func (m *GetCurrentUserEventsRequest) String() string { return proto.CompactTextString(m) }
func (*GetCurrentUserEventsRequest) ProtoMessage()    {}
func (*GetCurrentUserEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9225461cd3a21153, []int{4}
}
func (m *GetCurrentUserEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentUserEventsRequest.Unmarshal(m, b)
}
func (m *GetCurrentUserEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentUserEventsRequest.Marshal(b, m, deterministic)
}
func (m *GetCurrentUserEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentUserEventsRequest.Merge(m, src)
}
func (m *GetCurrentUserEventsRequest) XXX_Size() int {
	return xxx_messageInfo_GetCurrentUserEventsRequest.Size(m)
}
func (m *GetCurrentUserEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentUserEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentUserEventsRequest proto.InternalMessageInfo

func (m *GetCurrentUserEventsRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *GetCurrentUserEventsRequest) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *GetCurrentUserEventsRequest) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

type GetCurrentUserEventsResponse struct {
	BaseListResponse     `protobuf:"bytes,1,opt,name=base_list_response,json=baseListResponse,proto3,embedded=base_list_response" json:",omitempty"`
	Data                 []*Event `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrentUserEventsResponse) Reset()         { *m = GetCurrentUserEventsResponse{} }
func (m *GetCurrentUserEventsResponse) String() string { return proto.CompactTextString(m) }
func (*GetCurrentUserEventsResponse) ProtoMessage()    {}
func (*GetCurrentUserEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9225461cd3a21153, []int{5}
}
func (m *GetCurrentUserEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentUserEventsResponse.Unmarshal(m, b)
}
func (m *GetCurrentUserEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentUserEventsResponse.Marshal(b, m, deterministic)
}
func (m *GetCurrentUserEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentUserEventsResponse.Merge(m, src)
}
func (m *GetCurrentUserEventsResponse) XXX_Size() int {
	return xxx_messageInfo_GetCurrentUserEventsResponse.Size(m)
}
func (m *GetCurrentUserEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentUserEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentUserEventsResponse proto.InternalMessageInfo

func (m *GetCurrentUserEventsResponse) GetData() []*Event {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetEventLogsRequest struct {
	BaseListRequest      `protobuf:"bytes,1,opt,name=base_list_request,json=baseListRequest,proto3,embedded=base_list_request" json:",omitempty"`
	EventId              string   `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3" json:"eventId" valid:"required,uuid"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventLogsRequest) Reset()         { *m = GetEventLogsRequest{} }
func (m *GetEventLogsRequest) String() string { return proto.CompactTextString(m) }
func (*GetEventLogsRequest) ProtoMessage()    {}
func (*GetEventLogsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9225461cd3a21153, []int{6}
}
func (m *GetEventLogsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventLogsRequest.Unmarshal(m, b)
}
func (m *GetEventLogsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventLogsRequest.Marshal(b, m, deterministic)
}
func (m *GetEventLogsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventLogsRequest.Merge(m, src)
}
func (m *GetEventLogsRequest) XXX_Size() int {
	return xxx_messageInfo_GetEventLogsRequest.Size(m)
}
func (m *GetEventLogsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventLogsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventLogsRequest proto.InternalMessageInfo

func (m *GetEventLogsRequest) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

type GetEventLogsResponse struct {
	BaseListResponse     `protobuf:"bytes,1,opt,name=base_list_response,json=baseListResponse,proto3,embedded=base_list_response" json:",omitempty"`
	Data                 []*EventLog `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetEventLogsResponse) Reset()         { *m = GetEventLogsResponse{} }
func (m *GetEventLogsResponse) String() string { return proto.CompactTextString(m) }
func (*GetEventLogsResponse) ProtoMessage()    {}
func (*GetEventLogsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9225461cd3a21153, []int{7}
}
func (m *GetEventLogsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventLogsResponse.Unmarshal(m, b)
}
func (m *GetEventLogsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventLogsResponse.Marshal(b, m, deterministic)
}
func (m *GetEventLogsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventLogsResponse.Merge(m, src)
}
func (m *GetEventLogsResponse) XXX_Size() int {
	return xxx_messageInfo_GetEventLogsResponse.Size(m)
}
func (m *GetEventLogsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventLogsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventLogsResponse proto.InternalMessageInfo

func (m *GetEventLogsResponse) GetData() []*EventLog {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetCurrentUserEventLogsRequest struct {
	BaseListRequest      `protobuf:"bytes,1,opt,name=base_list_request,json=baseListRequest,proto3,embedded=base_list_request" json:",omitempty"`
	EventId              string   `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3" json:"eventId" valid:"required,uuid"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrentUserEventLogsRequest) Reset()         { *m = GetCurrentUserEventLogsRequest{} }
func (m *GetCurrentUserEventLogsRequest) String() string { return proto.CompactTextString(m) }
func (*GetCurrentUserEventLogsRequest) ProtoMessage()    {}
func (*GetCurrentUserEventLogsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9225461cd3a21153, []int{8}
}
func (m *GetCurrentUserEventLogsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentUserEventLogsRequest.Unmarshal(m, b)
}
func (m *GetCurrentUserEventLogsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentUserEventLogsRequest.Marshal(b, m, deterministic)
}
func (m *GetCurrentUserEventLogsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentUserEventLogsRequest.Merge(m, src)
}
func (m *GetCurrentUserEventLogsRequest) XXX_Size() int {
	return xxx_messageInfo_GetCurrentUserEventLogsRequest.Size(m)
}
func (m *GetCurrentUserEventLogsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentUserEventLogsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentUserEventLogsRequest proto.InternalMessageInfo

func (m *GetCurrentUserEventLogsRequest) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

type GetCurrentUserEventLogsResponse struct {
	BaseListResponse     `protobuf:"bytes,1,opt,name=base_list_response,json=baseListResponse,proto3,embedded=base_list_response" json:",omitempty"`
	Data                 []*EventLog `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetCurrentUserEventLogsResponse) Reset()         { *m = GetCurrentUserEventLogsResponse{} }
func (m *GetCurrentUserEventLogsResponse) String() string { return proto.CompactTextString(m) }
func (*GetCurrentUserEventLogsResponse) ProtoMessage()    {}
func (*GetCurrentUserEventLogsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9225461cd3a21153, []int{9}
}
func (m *GetCurrentUserEventLogsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentUserEventLogsResponse.Unmarshal(m, b)
}
func (m *GetCurrentUserEventLogsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentUserEventLogsResponse.Marshal(b, m, deterministic)
}
func (m *GetCurrentUserEventLogsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentUserEventLogsResponse.Merge(m, src)
}
func (m *GetCurrentUserEventLogsResponse) XXX_Size() int {
	return xxx_messageInfo_GetCurrentUserEventLogsResponse.Size(m)
}
func (m *GetCurrentUserEventLogsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentUserEventLogsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentUserEventLogsResponse proto.InternalMessageInfo

func (m *GetCurrentUserEventLogsResponse) GetData() []*EventLog {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "fuck_web.endpoint.Event")
	proto.RegisterType((*EventLog)(nil), "fuck_web.endpoint.EventLog")
	proto.RegisterType((*GetEventsRequest)(nil), "fuck_web.endpoint.GetEventsRequest")
	proto.RegisterType((*GetEventsResponse)(nil), "fuck_web.endpoint.GetEventsResponse")
	proto.RegisterType((*GetCurrentUserEventsRequest)(nil), "fuck_web.endpoint.GetCurrentUserEventsRequest")
	proto.RegisterType((*GetCurrentUserEventsResponse)(nil), "fuck_web.endpoint.GetCurrentUserEventsResponse")
	proto.RegisterType((*GetEventLogsRequest)(nil), "fuck_web.endpoint.GetEventLogsRequest")
	proto.RegisterType((*GetEventLogsResponse)(nil), "fuck_web.endpoint.GetEventLogsResponse")
	proto.RegisterType((*GetCurrentUserEventLogsRequest)(nil), "fuck_web.endpoint.GetCurrentUserEventLogsRequest")
	proto.RegisterType((*GetCurrentUserEventLogsResponse)(nil), "fuck_web.endpoint.GetCurrentUserEventLogsResponse")
}

func init() { proto.RegisterFile("endpoints/logging.proto", fileDescriptor_9225461cd3a21153) }

var fileDescriptor_9225461cd3a21153 = []byte{
	// 779 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xc1, 0x4e, 0xdb, 0x48,
	0x18, 0xc6, 0x0e, 0x24, 0x61, 0xc2, 0x6a, 0x61, 0x60, 0x59, 0x6f, 0x40, 0x6b, 0x64, 0x16, 0x6d,
	0x76, 0x05, 0xb1, 0x44, 0xb4, 0x5a, 0xb5, 0x3d, 0x54, 0x0a, 0xaa, 0x28, 0x52, 0xaa, 0x22, 0xab,
	0xbd, 0xf4, 0x12, 0xd9, 0x9e, 0x89, 0x3b, 0xc2, 0xf6, 0x18, 0xcf, 0x18, 0xd4, 0x27, 0xe8, 0xc3,
	0xb4, 0xa7, 0x5e, 0x7a, 0x69, 0x7b, 0x6e, 0x6f, 0x7d, 0x02, 0x3f, 0x80, 0x8f, 0xa8, 0x0f, 0x50,
	0x79, 0xc6, 0x8e, 0x03, 0x4d, 0xda, 0x1b, 0x85, 0x8b, 0xe7, 0xf7, 0x37, 0xdf, 0x7c, 0x99, 0xf9,
	0xbe, 0xdf, 0xa3, 0x80, 0xdf, 0x71, 0x88, 0x22, 0x4a, 0x42, 0xce, 0x4c, 0x9f, 0x7a, 0x1e, 0x09,
	0xbd, 0x6e, 0x14, 0x53, 0x4e, 0xe1, 0xca, 0x28, 0x71, 0x4f, 0x86, 0xe7, 0xd8, 0xe9, 0x96, 0x8c,
	0xf6, 0x9a, 0x47, 0x3d, 0x2a, 0x66, 0xcd, 0xbc, 0x92, 0xc4, 0xb6, 0xee, 0x51, 0xea, 0xf9, 0xd8,
	0x14, 0x6f, 0x4e, 0x32, 0x32, 0x39, 0x09, 0x30, 0xe3, 0x76, 0x10, 0x15, 0x84, 0x25, 0x97, 0x86,
	0x23, 0x52, 0xe8, 0xb6, 0xd7, 0xaa, 0x1f, 0x74, 0x6c, 0x86, 0x0b, 0x74, 0x35, 0xa0, 0x08, 0xfb,
	0xcc, 0x94, 0x83, 0x04, 0x8d, 0x2f, 0x35, 0xb0, 0xf0, 0xe0, 0x0c, 0x87, 0x1c, 0x9a, 0x40, 0x25,
	0x48, 0x53, 0xb6, 0x94, 0xce, 0x62, 0x5f, 0xcf, 0x52, 0x5d, 0x25, 0xe8, 0x22, 0xd5, 0x7f, 0x3b,
	0xb3, 0x7d, 0x82, 0xee, 0x1a, 0x31, 0x3e, 0x4d, 0x48, 0x8c, 0xd1, 0x6e, 0x92, 0x10, 0x64, 0x58,
	0x2a, 0x41, 0xf0, 0x00, 0xb4, 0x92, 0x08, 0xd9, 0x1c, 0x0f, 0xf3, 0xdd, 0x68, 0xaa, 0x58, 0x69,
	0x64, 0xa9, 0x0e, 0x24, 0xfc, 0x84, 0x04, 0xf8, 0x22, 0xd5, 0x97, 0xaf, 0x28, 0x18, 0xd6, 0xc4,
	0x7c, 0x2e, 0xe2, 0xc6, 0x78, 0x2c, 0x52, 0xab, 0x44, 0x24, 0x3c, 0x5b, 0xa4, 0x9a, 0x87, 0xdb,
	0xa0, 0x91, 0x30, 0x1c, 0x0f, 0x09, 0xd2, 0xe6, 0x85, 0x00, 0xc8, 0x52, 0xbd, 0x9e, 0x43, 0x47,
	0xc8, 0x2a, 0x46, 0xd8, 0x01, 0xcd, 0xbc, 0x0a, 0xed, 0x00, 0x6b, 0x0b, 0x82, 0xb5, 0x94, 0xa5,
	0xfa, 0x18, 0xb3, 0xc6, 0x15, 0x34, 0x40, 0xdd, 0x76, 0x39, 0xa1, 0xa1, 0x56, 0xaf, 0xd4, 0x24,
	0x62, 0x15, 0x23, 0xfc, 0x17, 0x2c, 0xba, 0x3e, 0xc1, 0x21, 0x1f, 0x92, 0x48, 0x6b, 0x08, 0xda,
	0x2f, 0x59, 0xaa, 0x57, 0xa0, 0xd5, 0x94, 0xe5, 0x51, 0x94, 0xeb, 0x31, 0x6e, 0xf3, 0x84, 0x69,
	0xcd, 0x4a, 0x4f, 0x22, 0x56, 0x31, 0xc2, 0x4d, 0x30, 0xcf, 0x29, 0x3d, 0xd1, 0x16, 0xb7, 0x94,
	0x4e, 0xad, 0xdf, 0xcc, 0x52, 0x5d, 0xbc, 0x5b, 0xe2, 0x09, 0x77, 0x40, 0x23, 0xc0, 0x8c, 0xd9,
	0x1e, 0xd6, 0x80, 0x90, 0x68, 0x65, 0xa9, 0x5e, 0x42, 0x56, 0x59, 0xe4, 0x47, 0xf4, 0xa9, 0x6b,
	0x8b, 0xad, 0xb7, 0xaa, 0x23, 0x96, 0x98, 0x35, 0xae, 0x8c, 0x97, 0x2a, 0x68, 0x8a, 0xd8, 0x07,
	0xd4, 0xbb, 0xcd, 0xc9, 0xef, 0x80, 0x26, 0x3e, 0x13, 0x86, 0x4f, 0x8b, 0xbe, 0x21, 0xe6, 0x8e,
	0x10, 0xfc, 0x03, 0xd4, 0x7c, 0xea, 0x15, 0xb1, 0x37, 0xb2, 0x54, 0xcf, 0x5f, 0xad, 0xfc, 0x61,
	0xbc, 0x53, 0xc1, 0xf2, 0x21, 0xe6, 0xc2, 0x0c, 0x66, 0xe1, 0xd3, 0x04, 0x33, 0x0e, 0x47, 0x60,
	0x25, 0xff, 0x70, 0x86, 0x3e, 0x61, 0x7c, 0x18, 0x4b, 0x50, 0x18, 0xd4, 0xda, 0x37, 0xba, 0xdf,
	0x7c, 0xb4, 0xdd, 0xbe, 0xcd, 0xf0, 0x80, 0x30, 0x5e, 0x2c, 0xef, 0xaf, 0x7f, 0x4c, 0xf5, 0xb9,
	0xcf, 0xa9, 0xae, 0xe4, 0xa7, 0xd9, 0xa5, 0x01, 0xe1, 0x38, 0x88, 0xf8, 0x0b, 0xeb, 0x57, 0xe7,
	0x32, 0x11, 0xb6, 0x27, 0x7a, 0x52, 0xb8, 0x38, 0xd1, 0x85, 0xeb, 0xe3, 0x2e, 0x14, 0xd6, 0x8c,
	0x3b, 0xef, 0x21, 0x00, 0x8c, 0xdb, 0x31, 0x97, 0xb6, 0xc9, 0x43, 0xff, 0x93, 0xb7, 0x9e, 0x40,
	0x0b, 0xd7, 0xb4, 0xab, 0xe1, 0xc5, 0x23, 0xb7, 0xd7, 0xeb, 0xdd, 0x31, 0xac, 0x8a, 0x06, 0xfb,
	0xa0, 0x89, 0x43, 0x24, 0x75, 0xa4, 0x35, 0x7f, 0xe7, 0x6d, 0x85, 0x43, 0xf4, 0x43, 0x95, 0x92,
	0x64, 0xbc, 0x56, 0xc0, 0xca, 0x84, 0x7d, 0x2c, 0xa2, 0x21, 0xc3, 0x90, 0x00, 0x38, 0xe9, 0x9f,
	0x44, 0x0b, 0x03, 0xb7, 0xbf, 0x6b, 0xa0, 0xa4, 0xce, 0x74, 0x70, 0xd9, 0xb9, 0xc2, 0x84, 0xbb,
	0x60, 0x1e, 0xd9, 0xdc, 0xd6, 0xd4, 0xad, 0x5a, 0xa7, 0xb5, 0xaf, 0x4d, 0x11, 0x17, 0x7b, 0xb3,
	0x04, 0xcb, 0x78, 0xa5, 0x82, 0x8d, 0x43, 0xcc, 0x0f, 0x92, 0x38, 0xc6, 0x21, 0x7f, 0xca, 0x70,
	0xfc, 0x73, 0x82, 0xbf, 0x1d, 0xe1, 0xbe, 0x55, 0xc0, 0xe6, 0x74, 0xb7, 0x6e, 0x7a, 0xce, 0x1f,
	0x14, 0xb0, 0x5a, 0xb6, 0xe5, 0x80, 0x7a, 0xd7, 0x9e, 0xef, 0xfd, 0x89, 0x7b, 0x49, 0x5e, 0x8f,
	0x7f, 0x09, 0xf7, 0xe5, 0x7d, 0x34, 0xfb, 0x76, 0x2d, 0x19, 0xc6, 0x1b, 0x05, 0xac, 0x5d, 0x3e,
	0xc0, 0xf5, 0x5b, 0x6e, 0x5e, 0xb2, 0x7c, 0x63, 0x96, 0xe5, 0x03, 0xea, 0x15, 0xae, 0x7f, 0x52,
	0xc0, 0x9f, 0x53, 0xfa, 0xe5, 0x56, 0x06, 0xf0, 0x5e, 0x01, 0xfa, 0xcc, 0xb3, 0xdc, 0xfc, 0x2c,
	0xfa, 0xff, 0x3f, 0xfb, 0xcf, 0x23, 0xfc, 0x79, 0xe2, 0x74, 0x5d, 0x1a, 0x98, 0x8f, 0x88, 0x1b,
	0xd3, 0xc7, 0x11, 0xdb, 0x73, 0x43, 0x33, 0x5f, 0xba, 0x77, 0x8e, 0x1d, 0x33, 0x3a, 0xf1, 0xcc,
	0x72, 0xf9, 0xbd, 0xb2, 0x38, 0x9e, 0x3b, 0x56, 0x9c, 0xba, 0xf8, 0x73, 0xd8, 0xfb, 0x1a, 0x00,
	0x00, 0xff, 0xff, 0x7e, 0x0e, 0x11, 0xf9, 0xba, 0x0a, 0x00, 0x00,
}
