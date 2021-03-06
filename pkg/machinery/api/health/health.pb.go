// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: health/health.proto

package health

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type HealthCheck_ServingStatus int32

const (
	HealthCheck_UNKNOWN     HealthCheck_ServingStatus = 0
	HealthCheck_SERVING     HealthCheck_ServingStatus = 1
	HealthCheck_NOT_SERVING HealthCheck_ServingStatus = 2
)

// Enum value maps for HealthCheck_ServingStatus.
var (
	HealthCheck_ServingStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "SERVING",
		2: "NOT_SERVING",
	}
	HealthCheck_ServingStatus_value = map[string]int32{
		"UNKNOWN":     0,
		"SERVING":     1,
		"NOT_SERVING": 2,
	}
)

func (x HealthCheck_ServingStatus) Enum() *HealthCheck_ServingStatus {
	p := new(HealthCheck_ServingStatus)
	*p = x
	return p
}

func (x HealthCheck_ServingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthCheck_ServingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_health_health_proto_enumTypes[0].Descriptor()
}

func (HealthCheck_ServingStatus) Type() protoreflect.EnumType {
	return &file_health_health_proto_enumTypes[0]
}

func (x HealthCheck_ServingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthCheck_ServingStatus.Descriptor instead.
func (HealthCheck_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return file_health_health_proto_rawDescGZIP(), []int{1, 0}
}

type ReadyCheck_ReadyStatus int32

const (
	ReadyCheck_UNKNOWN   ReadyCheck_ReadyStatus = 0
	ReadyCheck_READY     ReadyCheck_ReadyStatus = 1
	ReadyCheck_NOT_READY ReadyCheck_ReadyStatus = 2
)

// Enum value maps for ReadyCheck_ReadyStatus.
var (
	ReadyCheck_ReadyStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "READY",
		2: "NOT_READY",
	}
	ReadyCheck_ReadyStatus_value = map[string]int32{
		"UNKNOWN":   0,
		"READY":     1,
		"NOT_READY": 2,
	}
)

func (x ReadyCheck_ReadyStatus) Enum() *ReadyCheck_ReadyStatus {
	p := new(ReadyCheck_ReadyStatus)
	*p = x
	return p
}

func (x ReadyCheck_ReadyStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReadyCheck_ReadyStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_health_health_proto_enumTypes[1].Descriptor()
}

func (ReadyCheck_ReadyStatus) Type() protoreflect.EnumType {
	return &file_health_health_proto_enumTypes[1]
}

func (x ReadyCheck_ReadyStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReadyCheck_ReadyStatus.Descriptor instead.
func (ReadyCheck_ReadyStatus) EnumDescriptor() ([]byte, []int) {
	return file_health_health_proto_rawDescGZIP(), []int{3, 0}
}

type HealthWatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntervalSeconds int64 `protobuf:"varint,1,opt,name=interval_seconds,json=intervalSeconds,proto3" json:"interval_seconds,omitempty"`
}

func (x *HealthWatchRequest) Reset() {
	*x = HealthWatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthWatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthWatchRequest) ProtoMessage() {}

func (x *HealthWatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_health_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthWatchRequest.ProtoReflect.Descriptor instead.
func (*HealthWatchRequest) Descriptor() ([]byte, []int) {
	return file_health_health_proto_rawDescGZIP(), []int{0}
}

func (x *HealthWatchRequest) GetIntervalSeconds() int64 {
	if x != nil {
		return x.IntervalSeconds
	}
	return 0
}

type HealthCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status HealthCheck_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=health.HealthCheck_ServingStatus" json:"status,omitempty"`
}

func (x *HealthCheck) Reset() {
	*x = HealthCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheck) ProtoMessage() {}

func (x *HealthCheck) ProtoReflect() protoreflect.Message {
	mi := &file_health_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheck.ProtoReflect.Descriptor instead.
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return file_health_health_proto_rawDescGZIP(), []int{1}
}

func (x *HealthCheck) GetStatus() HealthCheck_ServingStatus {
	if x != nil {
		return x.Status
	}
	return HealthCheck_UNKNOWN
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*HealthCheck `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_health_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_health_health_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_health_health_proto_rawDescGZIP(), []int{2}
}

func (x *HealthCheckResponse) GetMessages() []*HealthCheck {
	if x != nil {
		return x.Messages
	}
	return nil
}

type ReadyCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ReadyCheck_ReadyStatus `protobuf:"varint,1,opt,name=status,proto3,enum=health.ReadyCheck_ReadyStatus" json:"status,omitempty"`
}

func (x *ReadyCheck) Reset() {
	*x = ReadyCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_health_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyCheck) ProtoMessage() {}

func (x *ReadyCheck) ProtoReflect() protoreflect.Message {
	mi := &file_health_health_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyCheck.ProtoReflect.Descriptor instead.
func (*ReadyCheck) Descriptor() ([]byte, []int) {
	return file_health_health_proto_rawDescGZIP(), []int{3}
}

func (x *ReadyCheck) GetStatus() ReadyCheck_ReadyStatus {
	if x != nil {
		return x.Status
	}
	return ReadyCheck_UNKNOWN
}

type ReadyCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*ReadyCheck `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *ReadyCheckResponse) Reset() {
	*x = ReadyCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_health_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyCheckResponse) ProtoMessage() {}

func (x *ReadyCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_health_health_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyCheckResponse.ProtoReflect.Descriptor instead.
func (*ReadyCheckResponse) Descriptor() ([]byte, []int) {
	return file_health_health_proto_rawDescGZIP(), []int{4}
}

func (x *ReadyCheckResponse) GetMessages() []*ReadyCheck {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_health_health_proto protoreflect.FileDescriptor

var file_health_health_proto_rawDesc = []byte{
	0x0a, 0x13, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x12, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x0b,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x39, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x22, 0x46, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x7a, 0x0a, 0x0a, 0x52, 0x65,
	0x61, 0x64, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x34, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x52,
	0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x22, 0x44, 0x0a, 0x12, 0x52, 0x65, 0x61, 0x64, 0x79, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x32, 0xc7, 0x01, 0x0a,
	0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x3c, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1a,
	0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x3b, 0x0a, 0x05, 0x52, 0x65, 0x61,
	0x64, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x73, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_health_health_proto_rawDescOnce sync.Once
	file_health_health_proto_rawDescData = file_health_health_proto_rawDesc
)

func file_health_health_proto_rawDescGZIP() []byte {
	file_health_health_proto_rawDescOnce.Do(func() {
		file_health_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_health_health_proto_rawDescData)
	})
	return file_health_health_proto_rawDescData
}

var (
	file_health_health_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
	file_health_health_proto_msgTypes  = make([]protoimpl.MessageInfo, 5)
	file_health_health_proto_goTypes   = []interface{}{
		(HealthCheck_ServingStatus)(0), // 0: health.HealthCheck.ServingStatus
		(ReadyCheck_ReadyStatus)(0),    // 1: health.ReadyCheck.ReadyStatus
		(*HealthWatchRequest)(nil),     // 2: health.HealthWatchRequest
		(*HealthCheck)(nil),            // 3: health.HealthCheck
		(*HealthCheckResponse)(nil),    // 4: health.HealthCheckResponse
		(*ReadyCheck)(nil),             // 5: health.ReadyCheck
		(*ReadyCheckResponse)(nil),     // 6: health.ReadyCheckResponse
		(*empty.Empty)(nil),            // 7: google.protobuf.Empty
	}
)

var file_health_health_proto_depIdxs = []int32{
	0, // 0: health.HealthCheck.status:type_name -> health.HealthCheck.ServingStatus
	3, // 1: health.HealthCheckResponse.messages:type_name -> health.HealthCheck
	1, // 2: health.ReadyCheck.status:type_name -> health.ReadyCheck.ReadyStatus
	5, // 3: health.ReadyCheckResponse.messages:type_name -> health.ReadyCheck
	7, // 4: health.Health.Check:input_type -> google.protobuf.Empty
	2, // 5: health.Health.Watch:input_type -> health.HealthWatchRequest
	7, // 6: health.Health.Ready:input_type -> google.protobuf.Empty
	4, // 7: health.Health.Check:output_type -> health.HealthCheckResponse
	4, // 8: health.Health.Watch:output_type -> health.HealthCheckResponse
	6, // 9: health.Health.Ready:output_type -> health.ReadyCheckResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_health_health_proto_init() }
func file_health_health_proto_init() {
	if File_health_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_health_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthWatchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_health_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_health_health_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_health_health_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyCheck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_health_health_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyCheckResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_health_health_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_health_health_proto_goTypes,
		DependencyIndexes: file_health_health_proto_depIdxs,
		EnumInfos:         file_health_health_proto_enumTypes,
		MessageInfos:      file_health_health_proto_msgTypes,
	}.Build()
	File_health_health_proto = out.File
	file_health_health_proto_rawDesc = nil
	file_health_health_proto_goTypes = nil
	file_health_health_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConnInterface
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthClient interface {
	Check(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	Watch(ctx context.Context, in *HealthWatchRequest, opts ...grpc.CallOption) (Health_WatchClient, error)
	Ready(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ReadyCheckResponse, error)
}

type healthClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthClient(cc grpc.ClientConnInterface) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/health.Health/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthClient) Watch(ctx context.Context, in *HealthWatchRequest, opts ...grpc.CallOption) (Health_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Health_serviceDesc.Streams[0], "/health.Health/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &healthWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Health_WatchClient interface {
	Recv() (*HealthCheckResponse, error)
	grpc.ClientStream
}

type healthWatchClient struct {
	grpc.ClientStream
}

func (x *healthWatchClient) Recv() (*HealthCheckResponse, error) {
	m := new(HealthCheckResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *healthClient) Ready(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ReadyCheckResponse, error) {
	out := new(ReadyCheckResponse)
	err := c.cc.Invoke(ctx, "/health.Health/Ready", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
type HealthServer interface {
	Check(context.Context, *empty.Empty) (*HealthCheckResponse, error)
	Watch(*HealthWatchRequest, Health_WatchServer) error
	Ready(context.Context, *empty.Empty) (*ReadyCheckResponse, error)
}

// UnimplementedHealthServer can be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (*UnimplementedHealthServer) Check(context.Context, *empty.Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func (*UnimplementedHealthServer) Watch(*HealthWatchRequest, Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

func (*UnimplementedHealthServer) Ready(context.Context, *empty.Empty) (*ReadyCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ready not implemented")
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Health_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HealthWatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HealthServer).Watch(m, &healthWatchServer{stream})
}

type Health_WatchServer interface {
	Send(*HealthCheckResponse) error
	grpc.ServerStream
}

type healthWatchServer struct {
	grpc.ServerStream
}

func (x *healthWatchServer) Send(m *HealthCheckResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Health_Ready_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Ready(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health.Health/Ready",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Ready(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "health.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
		{
			MethodName: "Ready",
			Handler:    _Health_Ready_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Health_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "health/health.proto",
}
