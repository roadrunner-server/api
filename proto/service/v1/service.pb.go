// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: service.proto

package serviceV1

import (
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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Ok      bool   `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type Create struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Command         string            `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	ProcessNum      int64             `protobuf:"varint,3,opt,name=process_num,json=processNum,proto3" json:"process_num,omitempty"`
	ExecTimeout     int64             `protobuf:"varint,4,opt,name=exec_timeout,json=execTimeout,proto3" json:"exec_timeout,omitempty"`
	RemainAfterExit bool              `protobuf:"varint,5,opt,name=remain_after_exit,json=remainAfterExit,proto3" json:"remain_after_exit,omitempty"`
	Env             map[string]string `protobuf:"bytes,6,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RestartSec      uint64            `protobuf:"varint,7,opt,name=restart_sec,json=restartSec,proto3" json:"restart_sec,omitempty"`
	WorkingDir      string            `protobuf:"bytes,8,opt,name=working_dir,json=workingDir,proto3" json:"working_dir,omitempty"`
}

func (x *Create) Reset() {
	*x = Create{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Create) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Create) ProtoMessage() {}

func (x *Create) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Create.ProtoReflect.Descriptor instead.
func (*Create) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *Create) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Create) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Create) GetProcessNum() int64 {
	if x != nil {
		return x.ProcessNum
	}
	return 0
}

func (x *Create) GetExecTimeout() int64 {
	if x != nil {
		return x.ExecTimeout
	}
	return 0
}

func (x *Create) GetRemainAfterExit() bool {
	if x != nil {
		return x.RemainAfterExit
	}
	return false
}

func (x *Create) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *Create) GetRestartSec() uint64 {
	if x != nil {
		return x.RestartSec
	}
	return 0
}

func (x *Create) GetWorkingDir() string {
	if x != nil {
		return x.WorkingDir
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CPUPercent  float32 `protobuf:"fixed32,1,opt,name=CPUPercent,proto3" json:"CPUPercent,omitempty"`
	Pid         int32   `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	MemoryUsage uint64  `protobuf:"varint,3,opt,name=memoryUsage,proto3" json:"memoryUsage,omitempty"`
	Command     string  `protobuf:"bytes,4,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetCPUPercent() float32 {
	if x != nil {
		return x.CPUPercent
	}
	return 0
}

func (x *Status) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Status) GetMemoryUsage() uint64 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

func (x *Status) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []string `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *List) Reset() {
	*x = List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*List) ProtoMessage() {}

func (x *List) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use List.ProtoReflect.Descriptor instead.
func (*List) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *List) GetServices() []string {
	if x != nil {
		return x.Services
	}
	return nil
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0x22, 0x34, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f,
	0x6b, 0x22, 0xcf, 0x02, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x65,
	0x78, 0x65, 0x63, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x2a,
	0x0a, 0x11, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x65,
	0x78, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x72, 0x65, 0x6d, 0x61, 0x69,
	0x6e, 0x41, 0x66, 0x74, 0x65, 0x72, 0x45, 0x78, 0x69, 0x74, 0x12, 0x2d, 0x0a, 0x03, 0x65, 0x6e,
	0x76, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x6e, 0x76, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x63, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x1a, 0x36, 0x0a, 0x08, 0x45,
	0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x76, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x43, 0x50, 0x55, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0a, 0x43, 0x50, 0x55, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x1d, 0x0a, 0x07, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x0e,
	0x5a, 0x0c, 0x2e, 0x2f, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_proto_goTypes = []interface{}{
	(*Response)(nil), // 0: service.V1.Response
	(*Create)(nil),   // 1: service.V1.Create
	(*Status)(nil),   // 2: service.V1.Status
	(*Service)(nil),  // 3: service.V1.Service
	(*List)(nil),     // 4: service.V1.List
	nil,              // 5: service.V1.Create.EnvEntry
}
var file_service_proto_depIdxs = []int32{
	5, // 0: service.V1.Create.env:type_name -> service.V1.Create.EnvEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Create); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*List); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
