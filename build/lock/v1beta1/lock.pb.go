// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: lock/v1beta1/lock.proto

package lockV1Beta1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Resource      string                 `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Id            string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Ttl           *int64                 `protobuf:"varint,3,opt,name=ttl,proto3,oneof" json:"ttl,omitempty"`
	Wait          *int64                 `protobuf:"varint,4,opt,name=wait,proto3,oneof" json:"wait,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_lock_v1beta1_lock_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_lock_v1beta1_lock_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_lock_v1beta1_lock_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Request) GetTtl() int64 {
	if x != nil && x.Ttl != nil {
		return *x.Ttl
	}
	return 0
}

func (x *Request) GetWait() int64 {
	if x != nil && x.Wait != nil {
		return *x.Wait
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ok            bool                   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_lock_v1beta1_lock_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_lock_v1beta1_lock_proto_msgTypes[1]
	if x != nil {
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
	return file_lock_v1beta1_lock_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_lock_v1beta1_lock_proto protoreflect.FileDescriptor

const file_lock_v1beta1_lock_proto_rawDesc = "" +
	"\n" +
	"\x17lock/v1beta1/lock.proto\x12\flock.v1beta1\"v\n" +
	"\aRequest\x12\x1a\n" +
	"\bresource\x18\x01 \x01(\tR\bresource\x12\x0e\n" +
	"\x02id\x18\x02 \x01(\tR\x02id\x12\x15\n" +
	"\x03ttl\x18\x03 \x01(\x03H\x00R\x03ttl\x88\x01\x01\x12\x17\n" +
	"\x04wait\x18\x04 \x01(\x03H\x01R\x04wait\x88\x01\x01B\x06\n" +
	"\x04_ttlB\a\n" +
	"\x05_wait\"\x1a\n" +
	"\bResponse\x12\x0e\n" +
	"\x02ok\x18\x01 \x01(\bR\x02okB\x8c\x01ZBgithub.com/roadrunner-server/api/v4/build/lock/v1beta1;lockV1Beta1\xca\x02\x1bRoadRunner\\Lock\\DTO\\V1BETA1\xe2\x02'RoadRunner\\Lock\\DTO\\V1BETA1\\GPBMetadatab\x06proto3"

var (
	file_lock_v1beta1_lock_proto_rawDescOnce sync.Once
	file_lock_v1beta1_lock_proto_rawDescData []byte
)

func file_lock_v1beta1_lock_proto_rawDescGZIP() []byte {
	file_lock_v1beta1_lock_proto_rawDescOnce.Do(func() {
		file_lock_v1beta1_lock_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_lock_v1beta1_lock_proto_rawDesc), len(file_lock_v1beta1_lock_proto_rawDesc)))
	})
	return file_lock_v1beta1_lock_proto_rawDescData
}

var file_lock_v1beta1_lock_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_lock_v1beta1_lock_proto_goTypes = []any{
	(*Request)(nil),  // 0: lock.v1beta1.Request
	(*Response)(nil), // 1: lock.v1beta1.Response
}
var file_lock_v1beta1_lock_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lock_v1beta1_lock_proto_init() }
func file_lock_v1beta1_lock_proto_init() {
	if File_lock_v1beta1_lock_proto != nil {
		return
	}
	file_lock_v1beta1_lock_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_lock_v1beta1_lock_proto_rawDesc), len(file_lock_v1beta1_lock_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lock_v1beta1_lock_proto_goTypes,
		DependencyIndexes: file_lock_v1beta1_lock_proto_depIdxs,
		MessageInfos:      file_lock_v1beta1_lock_proto_msgTypes,
	}.Build()
	File_lock_v1beta1_lock_proto = out.File
	file_lock_v1beta1_lock_proto_goTypes = nil
	file_lock_v1beta1_lock_proto_depIdxs = nil
}
