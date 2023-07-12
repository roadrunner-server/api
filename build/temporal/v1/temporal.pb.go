// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: temporal/v1/temporal.proto

package v1

import (
	v12 "github.com/roadrunner-server/api/v4/build/common/v1"
	v11 "go.temporal.io/api/common/v1"
	v1 "go.temporal.io/api/failure/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Frame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *Frame) Reset() {
	*x = Frame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_v1_temporal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Frame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Frame) ProtoMessage() {}

func (x *Frame) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_v1_temporal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Frame.ProtoReflect.Descriptor instead.
func (*Frame) Descriptor() ([]byte, []int) {
	return file_temporal_v1_temporal_proto_rawDescGZIP(), []int{0}
}

func (x *Frame) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

// Single communication message.
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique ID of the message (counter)
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// command name (if any)
	Command string `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	// command options in json format.
	Options []byte `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	// error response.
	Failure *v1.Failure `protobuf:"bytes,4,opt,name=failure,proto3" json:"failure,omitempty"`
	// invocation or result payloads.
	Payloads *v11.Payloads `protobuf:"bytes,5,opt,name=payloads,proto3" json:"payloads,omitempty"`
	// invocation or result payloads.
	Header *v11.Header `protobuf:"bytes,6,opt,name=header,proto3" json:"header,omitempty"`
	// workflow history length
	HistoryLength int64 `protobuf:"varint,7,opt,name=history_length,json=historyLength,proto3" json:"history_length,omitempty"`
	// rr_run id
	RunId string `protobuf:"bytes,8,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_v1_temporal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_v1_temporal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_temporal_v1_temporal_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Message) GetOptions() []byte {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *Message) GetFailure() *v1.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

func (x *Message) GetPayloads() *v11.Payloads {
	if x != nil {
		return x.Payloads
	}
	return nil
}

func (x *Message) GetHeader() *v11.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Message) GetHistoryLength() int64 {
	if x != nil {
		return x.HistoryLength
	}
	return 0
}

func (x *Message) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

type ReplayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowExecution *v11.WorkflowExecution `protobuf:"bytes,1,opt,name=workflow_execution,json=workflowExecution,proto3" json:"workflow_execution,omitempty"`
	WorkflowType      *v11.WorkflowType      `protobuf:"bytes,2,opt,name=workflow_type,json=workflowType,proto3" json:"workflow_type,omitempty"`
	SavePath          string                 `protobuf:"bytes,3,opt,name=save_path,json=savePath,proto3" json:"save_path,omitempty"`
}

func (x *ReplayRequest) Reset() {
	*x = ReplayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_v1_temporal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplayRequest) ProtoMessage() {}

func (x *ReplayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_v1_temporal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplayRequest.ProtoReflect.Descriptor instead.
func (*ReplayRequest) Descriptor() ([]byte, []int) {
	return file_temporal_v1_temporal_proto_rawDescGZIP(), []int{2}
}

func (x *ReplayRequest) GetWorkflowExecution() *v11.WorkflowExecution {
	if x != nil {
		return x.WorkflowExecution
	}
	return nil
}

func (x *ReplayRequest) GetWorkflowType() *v11.WorkflowType {
	if x != nil {
		return x.WorkflowType
	}
	return nil
}

func (x *ReplayRequest) GetSavePath() string {
	if x != nil {
		return x.SavePath
	}
	return ""
}

// https://cloud.google.com/apis/design/errors
type ReplayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *v12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ReplayResponse) Reset() {
	*x = ReplayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_v1_temporal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplayResponse) ProtoMessage() {}

func (x *ReplayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_v1_temporal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplayResponse.ProtoReflect.Descriptor instead.
func (*ReplayResponse) Descriptor() ([]byte, []int) {
	return file_temporal_v1_temporal_proto_rawDescGZIP(), []int{3}
}

func (x *ReplayResponse) GetStatus() *v12.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_temporal_v1_temporal_proto protoreflect.FileDescriptor

var file_temporal_v1_temporal_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x05, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0xbd,
	0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3a,
	0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x08,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x25, 0x0a, 0x0e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x22, 0xd1,
	0x01, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x58, 0x0a, 0x12, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x0d, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x61, 0x76, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x22, 0x3b, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x53, 0x5a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0xca, 0x02,
	0x1a, 0x52, 0x6f, 0x61, 0x64, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x5c, 0x54, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x5c, 0x44, 0x54, 0x4f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x26, 0x52, 0x6f,
	0x61, 0x64, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x5c, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x5c, 0x44, 0x54, 0x4f, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_v1_temporal_proto_rawDescOnce sync.Once
	file_temporal_v1_temporal_proto_rawDescData = file_temporal_v1_temporal_proto_rawDesc
)

func file_temporal_v1_temporal_proto_rawDescGZIP() []byte {
	file_temporal_v1_temporal_proto_rawDescOnce.Do(func() {
		file_temporal_v1_temporal_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_v1_temporal_proto_rawDescData)
	})
	return file_temporal_v1_temporal_proto_rawDescData
}

var file_temporal_v1_temporal_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_temporal_v1_temporal_proto_goTypes = []interface{}{
	(*Frame)(nil),                 // 0: temporal.v1.Frame
	(*Message)(nil),               // 1: temporal.v1.Message
	(*ReplayRequest)(nil),         // 2: temporal.v1.ReplayRequest
	(*ReplayResponse)(nil),        // 3: temporal.v1.ReplayResponse
	(*v1.Failure)(nil),            // 4: temporal.api.failure.v1.Failure
	(*v11.Payloads)(nil),          // 5: temporal.api.common.v1.Payloads
	(*v11.Header)(nil),            // 6: temporal.api.common.v1.Header
	(*v11.WorkflowExecution)(nil), // 7: temporal.api.common.v1.WorkflowExecution
	(*v11.WorkflowType)(nil),      // 8: temporal.api.common.v1.WorkflowType
	(*v12.Status)(nil),            // 9: common.v1.Status
}
var file_temporal_v1_temporal_proto_depIdxs = []int32{
	1, // 0: temporal.v1.Frame.messages:type_name -> temporal.v1.Message
	4, // 1: temporal.v1.Message.failure:type_name -> temporal.api.failure.v1.Failure
	5, // 2: temporal.v1.Message.payloads:type_name -> temporal.api.common.v1.Payloads
	6, // 3: temporal.v1.Message.header:type_name -> temporal.api.common.v1.Header
	7, // 4: temporal.v1.ReplayRequest.workflow_execution:type_name -> temporal.api.common.v1.WorkflowExecution
	8, // 5: temporal.v1.ReplayRequest.workflow_type:type_name -> temporal.api.common.v1.WorkflowType
	9, // 6: temporal.v1.ReplayResponse.status:type_name -> common.v1.Status
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_temporal_v1_temporal_proto_init() }
func file_temporal_v1_temporal_proto_init() {
	if File_temporal_v1_temporal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_v1_temporal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Frame); i {
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
		file_temporal_v1_temporal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_temporal_v1_temporal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplayRequest); i {
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
		file_temporal_v1_temporal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplayResponse); i {
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
			RawDescriptor: file_temporal_v1_temporal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_v1_temporal_proto_goTypes,
		DependencyIndexes: file_temporal_v1_temporal_proto_depIdxs,
		MessageInfos:      file_temporal_v1_temporal_proto_msgTypes,
	}.Build()
	File_temporal_v1_temporal_proto = out.File
	file_temporal_v1_temporal_proto_rawDesc = nil
	file_temporal_v1_temporal_proto_goTypes = nil
	file_temporal_v1_temporal_proto_depIdxs = nil
}
