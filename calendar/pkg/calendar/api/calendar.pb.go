// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calendar.proto

package api

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Period int32

const (
	Period_DAY   Period = 0
	Period_WEEK  Period = 1
	Period_MONTH Period = 2
)

var Period_name = map[int32]string{
	0: "DAY",
	1: "WEEK",
	2: "MONTH",
}

var Period_value = map[string]int32{
	"DAY":   0,
	"WEEK":  1,
	"MONTH": 2,
}

func (x Period) String() string {
	return proto.EnumName(Period_name, int32(x))
}

func (Period) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{0}
}

type Event struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Header               string               `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Text                 string               `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	User                 string               `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	ReminderBefore       *duration.Duration   `protobuf:"bytes,7,opt,name=reminder_before,json=reminderBefore,proto3" json:"reminder_before,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{0}
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

func (m *Event) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *Event) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Event) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Event) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Event) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Event) GetReminderBefore() *duration.Duration {
	if m != nil {
		return m.ReminderBefore
	}
	return nil
}

type AddResponseResult struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponseResult) Reset()         { *m = AddResponseResult{} }
func (m *AddResponseResult) String() string { return proto.CompactTextString(m) }
func (*AddResponseResult) ProtoMessage()    {}
func (*AddResponseResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{1}
}

func (m *AddResponseResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponseResult.Unmarshal(m, b)
}
func (m *AddResponseResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponseResult.Marshal(b, m, deterministic)
}
func (m *AddResponseResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponseResult.Merge(m, src)
}
func (m *AddResponseResult) XXX_Size() int {
	return xxx_messageInfo_AddResponseResult.Size(m)
}
func (m *AddResponseResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponseResult.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponseResult proto.InternalMessageInfo

func (m *AddResponseResult) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *AddResponseResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AddResponseResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type UpdateResponseResult struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponseResult) Reset()         { *m = UpdateResponseResult{} }
func (m *UpdateResponseResult) String() string { return proto.CompactTextString(m) }
func (*UpdateResponseResult) ProtoMessage()    {}
func (*UpdateResponseResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{2}
}

func (m *UpdateResponseResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponseResult.Unmarshal(m, b)
}
func (m *UpdateResponseResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponseResult.Marshal(b, m, deterministic)
}
func (m *UpdateResponseResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponseResult.Merge(m, src)
}
func (m *UpdateResponseResult) XXX_Size() int {
	return xxx_messageInfo_UpdateResponseResult.Size(m)
}
func (m *UpdateResponseResult) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponseResult.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponseResult proto.InternalMessageInfo

func (m *UpdateResponseResult) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *UpdateResponseResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateResponseResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type DelResponseResult struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelResponseResult) Reset()         { *m = DelResponseResult{} }
func (m *DelResponseResult) String() string { return proto.CompactTextString(m) }
func (*DelResponseResult) ProtoMessage()    {}
func (*DelResponseResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{3}
}

func (m *DelResponseResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelResponseResult.Unmarshal(m, b)
}
func (m *DelResponseResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelResponseResult.Marshal(b, m, deterministic)
}
func (m *DelResponseResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelResponseResult.Merge(m, src)
}
func (m *DelResponseResult) XXX_Size() int {
	return xxx_messageInfo_DelResponseResult.Size(m)
}
func (m *DelResponseResult) XXX_DiscardUnknown() {
	xxx_messageInfo_DelResponseResult.DiscardUnknown(m)
}

var xxx_messageInfo_DelResponseResult proto.InternalMessageInfo

func (m *DelResponseResult) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *DelResponseResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DelResponseResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type EventsResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Events               []*Event `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventsResponse) Reset()         { *m = EventsResponse{} }
func (m *EventsResponse) String() string { return proto.CompactTextString(m) }
func (*EventsResponse) ProtoMessage()    {}
func (*EventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{4}
}

func (m *EventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventsResponse.Unmarshal(m, b)
}
func (m *EventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventsResponse.Marshal(b, m, deterministic)
}
func (m *EventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventsResponse.Merge(m, src)
}
func (m *EventsResponse) XXX_Size() int {
	return xxx_messageInfo_EventsResponse.Size(m)
}
func (m *EventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventsResponse proto.InternalMessageInfo

func (m *EventsResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *EventsResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *EventsResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type EventResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Event                *Event   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventResponse) Reset()         { *m = EventResponse{} }
func (m *EventResponse) String() string { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()    {}
func (*EventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{5}
}

func (m *EventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventResponse.Unmarshal(m, b)
}
func (m *EventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventResponse.Marshal(b, m, deterministic)
}
func (m *EventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventResponse.Merge(m, src)
}
func (m *EventResponse) XXX_Size() int {
	return xxx_messageInfo_EventResponse.Size(m)
}
func (m *EventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventResponse proto.InternalMessageInfo

func (m *EventResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *EventResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *EventResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type RequestUser struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestUser) Reset()         { *m = RequestUser{} }
func (m *RequestUser) String() string { return proto.CompactTextString(m) }
func (*RequestUser) ProtoMessage()    {}
func (*RequestUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{6}
}

func (m *RequestUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestUser.Unmarshal(m, b)
}
func (m *RequestUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestUser.Marshal(b, m, deterministic)
}
func (m *RequestUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestUser.Merge(m, src)
}
func (m *RequestUser) XXX_Size() int {
	return xxx_messageInfo_RequestUser.Size(m)
}
func (m *RequestUser) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestUser.DiscardUnknown(m)
}

var xxx_messageInfo_RequestUser proto.InternalMessageInfo

func (m *RequestUser) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type EventID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventID) Reset()         { *m = EventID{} }
func (m *EventID) String() string { return proto.CompactTextString(m) }
func (*EventID) ProtoMessage()    {}
func (*EventID) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{7}
}

func (m *EventID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventID.Unmarshal(m, b)
}
func (m *EventID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventID.Marshal(b, m, deterministic)
}
func (m *EventID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventID.Merge(m, src)
}
func (m *EventID) XXX_Size() int {
	return xxx_messageInfo_EventID.Size(m)
}
func (m *EventID) XXX_DiscardUnknown() {
	xxx_messageInfo_EventID.DiscardUnknown(m)
}

var xxx_messageInfo_EventID proto.InternalMessageInfo

func (m *EventID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type PeriodRequest struct {
	User                 string               `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Period               Period               `protobuf:"varint,2,opt,name=period,proto3,enum=api.Period" json:"period,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PeriodRequest) Reset()         { *m = PeriodRequest{} }
func (m *PeriodRequest) String() string { return proto.CompactTextString(m) }
func (*PeriodRequest) ProtoMessage()    {}
func (*PeriodRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{8}
}

func (m *PeriodRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeriodRequest.Unmarshal(m, b)
}
func (m *PeriodRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeriodRequest.Marshal(b, m, deterministic)
}
func (m *PeriodRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeriodRequest.Merge(m, src)
}
func (m *PeriodRequest) XXX_Size() int {
	return xxx_messageInfo_PeriodRequest.Size(m)
}
func (m *PeriodRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PeriodRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PeriodRequest proto.InternalMessageInfo

func (m *PeriodRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *PeriodRequest) GetPeriod() Period {
	if m != nil {
		return m.Period
	}
	return Period_DAY
}

func (m *PeriodRequest) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func init() {
	proto.RegisterEnum("api.Period", Period_name, Period_value)
	proto.RegisterType((*Event)(nil), "api.Event")
	proto.RegisterType((*AddResponseResult)(nil), "api.AddResponseResult")
	proto.RegisterType((*UpdateResponseResult)(nil), "api.UpdateResponseResult")
	proto.RegisterType((*DelResponseResult)(nil), "api.DelResponseResult")
	proto.RegisterType((*EventsResponse)(nil), "api.EventsResponse")
	proto.RegisterType((*EventResponse)(nil), "api.EventResponse")
	proto.RegisterType((*RequestUser)(nil), "api.RequestUser")
	proto.RegisterType((*EventID)(nil), "api.EventID")
	proto.RegisterType((*PeriodRequest)(nil), "api.PeriodRequest")
}

func init() { proto.RegisterFile("calendar.proto", fileDescriptor_e3d25d49f056cdb2) }

var fileDescriptor_e3d25d49f056cdb2 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4b, 0x6f, 0xd3, 0x4c,
	0x14, 0x4d, 0x9c, 0xd8, 0x71, 0x6f, 0xbe, 0xe4, 0x4b, 0x87, 0xaa, 0x72, 0xb3, 0x80, 0x60, 0x24,
	0x54, 0xb1, 0x70, 0x51, 0x2a, 0x10, 0x2c, 0x13, 0x1c, 0x20, 0x42, 0x3c, 0x6a, 0xa5, 0x42, 0xb0,
	0x89, 0x1c, 0xe6, 0xb6, 0x58, 0x4a, 0x6c, 0x33, 0x33, 0x46, 0xe5, 0x47, 0xb0, 0xe6, 0xef, 0x22,
	0xdf, 0x99, 0x90, 0x90, 0x34, 0x74, 0xd3, 0xdd, 0x3c, 0xce, 0xe3, 0xea, 0xcc, 0xb1, 0xa1, 0xfd,
	0x25, 0x9e, 0x63, 0xca, 0x63, 0x11, 0xe4, 0x22, 0x53, 0x19, 0xab, 0xc5, 0x79, 0xd2, 0xbd, 0x77,
	0x99, 0x65, 0x97, 0x73, 0x3c, 0xa1, 0xa3, 0x59, 0x71, 0x71, 0xa2, 0x92, 0x05, 0x4a, 0x15, 0x2f,
	0x72, 0x8d, 0xea, 0xde, 0xdd, 0x04, 0xf0, 0x42, 0xc4, 0x2a, 0xc9, 0x52, 0x7d, 0xef, 0xff, 0xb4,
	0xc0, 0x1e, 0x7d, 0xc7, 0x54, 0xb1, 0x36, 0x58, 0x09, 0xf7, 0xaa, 0xbd, 0xea, 0xf1, 0x5e, 0x64,
	0x25, 0x9c, 0x1d, 0x82, 0xf3, 0x15, 0x63, 0x8e, 0xc2, 0xb3, 0xe8, 0xcc, 0xec, 0x18, 0x83, 0xba,
	0xc2, 0x2b, 0xe5, 0xd5, 0xe8, 0x94, 0xd6, 0xec, 0x39, 0x80, 0x54, 0xb1, 0x50, 0xd3, 0xd2, 0xde,
	0xab, 0xf7, 0xaa, 0xc7, 0xcd, 0x7e, 0x37, 0xd0, 0xd6, 0xc1, 0xd2, 0x3a, 0x98, 0x2c, 0x67, 0x8b,
	0xf6, 0x08, 0x5d, 0xee, 0xd9, 0x13, 0x70, 0x31, 0xe5, 0x9a, 0x68, 0xdf, 0x48, 0x6c, 0x60, 0xca,
	0x89, 0xc6, 0xa0, 0x5e, 0x48, 0x14, 0x9e, 0xa3, 0xa7, 0x28, 0xd7, 0x6c, 0x08, 0xff, 0x0b, 0x5c,
	0x24, 0x29, 0x47, 0x31, 0x9d, 0xe1, 0x45, 0x26, 0xd0, 0x6b, 0x90, 0xe2, 0xd1, 0x96, 0x62, 0x68,
	0x52, 0x88, 0xda, 0x4b, 0xc6, 0x90, 0x08, 0xfe, 0x19, 0xec, 0x0f, 0x38, 0x8f, 0x50, 0xe6, 0x59,
	0x2a, 0x31, 0x42, 0x59, 0xcc, 0x55, 0x19, 0x85, 0x54, 0xb1, 0x2a, 0x24, 0xc5, 0xe3, 0x46, 0x66,
	0x67, 0x22, 0xb3, 0xfe, 0x44, 0x76, 0x00, 0x36, 0x0a, 0x91, 0x09, 0x93, 0x8d, 0xde, 0xf8, 0x13,
	0x38, 0x38, 0xcf, 0x79, 0xac, 0xf0, 0x56, 0x55, 0xcf, 0x60, 0x3f, 0xc4, 0xf9, 0xad, 0x4a, 0xce,
	0xa0, 0x4d, 0x55, 0x90, 0x4b, 0xd5, 0x9d, 0x7a, 0x3e, 0x38, 0x48, 0x48, 0xcf, 0xea, 0xd5, 0x8e,
	0x9b, 0x7d, 0x08, 0xe2, 0x3c, 0x09, 0x88, 0x1c, 0x99, 0x9b, 0x1d, 0x1e, 0x53, 0x68, 0x69, 0xd8,
	0x4d, 0x16, 0x3d, 0xb0, 0x49, 0x88, 0xa6, 0xfe, 0xdb, 0x41, 0x5f, 0xec, 0x30, 0xb8, 0x0f, 0xcd,
	0x08, 0xbf, 0x15, 0x28, 0xd5, 0xb9, 0xd4, 0x6d, 0xa5, 0x9e, 0x54, 0x57, 0x3d, 0xf1, 0x8f, 0xa0,
	0x41, 0x42, 0xe3, 0x70, 0xb3, 0xf4, 0xfe, 0x15, 0xb4, 0x3e, 0xa0, 0x48, 0x32, 0x6e, 0x34, 0xae,
	0xe3, 0xb3, 0x07, 0xe0, 0xe4, 0x04, 0xa2, 0xd9, 0xda, 0xfd, 0x26, 0xcd, 0x66, 0x78, 0xe6, 0x8a,
	0x05, 0x50, 0x2f, 0xdf, 0x9c, 0x86, 0xfb, 0x77, 0xa7, 0x09, 0xf7, 0xe8, 0x21, 0x38, 0x5a, 0x81,
	0x35, 0xa0, 0x16, 0x0e, 0x3e, 0x75, 0x2a, 0xcc, 0x85, 0xfa, 0xc7, 0xd1, 0xe8, 0x4d, 0xa7, 0xca,
	0xf6, 0xc0, 0x7e, 0xfb, 0xfe, 0xdd, 0xe4, 0x75, 0xc7, 0xea, 0xff, 0xb2, 0xc0, 0x7d, 0x61, 0xfe,
	0x04, 0xec, 0x31, 0xb8, 0x03, 0xce, 0xf5, 0xf7, 0xbb, 0x96, 0x50, 0xf7, 0x90, 0xd6, 0x5b, 0x45,
	0xf6, 0x2b, 0xec, 0x29, 0x34, 0x75, 0x19, 0xb7, 0x49, 0x47, 0xb4, 0xbe, 0xae, 0xaa, 0x7e, 0x85,
	0xf5, 0xc1, 0x0d, 0x71, 0xae, 0x49, 0xff, 0xad, 0x48, 0xe3, 0xd0, 0x78, 0x6d, 0x75, 0xd1, 0xaf,
	0xb0, 0x53, 0x68, 0xbd, 0x4c, 0x52, 0x3d, 0xde, 0xf0, 0xc7, 0x38, 0xdc, 0x20, 0xb2, 0xb5, 0x27,
	0x35, 0x54, 0xbf, 0xc2, 0x9e, 0x41, 0xeb, 0x15, 0xd2, 0xdb, 0xe9, 0x2e, 0xb2, 0x0e, 0xc1, 0xd6,
	0xde, 0xb4, 0x7b, 0x67, 0x45, 0x94, 0x2b, 0xe6, 0xd0, 0xfe, 0x5c, 0xfe, 0x12, 0x67, 0x0e, 0x45,
	0x7c, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x4f, 0xfc, 0xb5, 0x30, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalendarClient is the client API for Calendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarClient interface {
	AddEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*AddResponseResult, error)
	UpdateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*UpdateResponseResult, error)
	DelEvent(ctx context.Context, in *EventID, opts ...grpc.CallOption) (*DelResponseResult, error)
	FindEventByID(ctx context.Context, in *EventID, opts ...grpc.CallOption) (*EventResponse, error)
	GetUserEvents(ctx context.Context, in *RequestUser, opts ...grpc.CallOption) (*EventsResponse, error)
}

type calendarClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarClient(cc grpc.ClientConnInterface) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) AddEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*AddResponseResult, error) {
	out := new(AddResponseResult)
	err := c.cc.Invoke(ctx, "/api.Calendar/AddEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) UpdateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*UpdateResponseResult, error) {
	out := new(UpdateResponseResult)
	err := c.cc.Invoke(ctx, "/api.Calendar/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) DelEvent(ctx context.Context, in *EventID, opts ...grpc.CallOption) (*DelResponseResult, error) {
	out := new(DelResponseResult)
	err := c.cc.Invoke(ctx, "/api.Calendar/DelEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) FindEventByID(ctx context.Context, in *EventID, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/api.Calendar/FindEventByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) GetUserEvents(ctx context.Context, in *RequestUser, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/api.Calendar/GetUserEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServer is the server API for Calendar service.
type CalendarServer interface {
	AddEvent(context.Context, *Event) (*AddResponseResult, error)
	UpdateEvent(context.Context, *Event) (*UpdateResponseResult, error)
	DelEvent(context.Context, *EventID) (*DelResponseResult, error)
	FindEventByID(context.Context, *EventID) (*EventResponse, error)
	GetUserEvents(context.Context, *RequestUser) (*EventsResponse, error)
}

// UnimplementedCalendarServer can be embedded to have forward compatible implementations.
type UnimplementedCalendarServer struct {
}

func (*UnimplementedCalendarServer) AddEvent(ctx context.Context, req *Event) (*AddResponseResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEvent not implemented")
}
func (*UnimplementedCalendarServer) UpdateEvent(ctx context.Context, req *Event) (*UpdateResponseResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (*UnimplementedCalendarServer) DelEvent(ctx context.Context, req *EventID) (*DelResponseResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelEvent not implemented")
}
func (*UnimplementedCalendarServer) FindEventByID(ctx context.Context, req *EventID) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEventByID not implemented")
}
func (*UnimplementedCalendarServer) GetUserEvents(ctx context.Context, req *RequestUser) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserEvents not implemented")
}

func RegisterCalendarServer(s *grpc.Server, srv CalendarServer) {
	s.RegisterService(&_Calendar_serviceDesc, srv)
}

func _Calendar_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Calendar/AddEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).AddEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Calendar/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).UpdateEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_DelEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).DelEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Calendar/DelEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).DelEvent(ctx, req.(*EventID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_FindEventByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).FindEventByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Calendar/FindEventByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).FindEventByID(ctx, req.(*EventID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_GetUserEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetUserEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Calendar/GetUserEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetUserEvents(ctx, req.(*RequestUser))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calendar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEvent",
			Handler:    _Calendar_AddEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _Calendar_UpdateEvent_Handler,
		},
		{
			MethodName: "DelEvent",
			Handler:    _Calendar_DelEvent_Handler,
		},
		{
			MethodName: "FindEventByID",
			Handler:    _Calendar_FindEventByID_Handler,
		},
		{
			MethodName: "GetUserEvents",
			Handler:    _Calendar_GetUserEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calendar.proto",
}