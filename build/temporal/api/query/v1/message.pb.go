// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: temporal/api/query/v1/message.proto

package query

import (
	v1 "go.temporal.io/api/common/v1"
	v11 "go.temporal.io/api/enums/v1"
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

// See https://docs.temporal.io/docs/concepts/queries/
type WorkflowQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The workflow-author-defined identifier of the query. Typically a function name.
	QueryType string `protobuf:"bytes,1,opt,name=query_type,json=queryType,proto3" json:"query_type,omitempty"`
	// Serialized arguments that will be provided to the query handler.
	QueryArgs *v1.Payloads `protobuf:"bytes,2,opt,name=query_args,json=queryArgs,proto3" json:"query_args,omitempty"`
	// Headers that were passed by the caller of the query and copied by temporal
	// server into the workflow task.
	Header *v1.Header `protobuf:"bytes,3,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *WorkflowQuery) Reset() {
	*x = WorkflowQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_query_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowQuery) ProtoMessage() {}

func (x *WorkflowQuery) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_query_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowQuery.ProtoReflect.Descriptor instead.
func (*WorkflowQuery) Descriptor() ([]byte, []int) {
	return file_temporal_api_query_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowQuery) GetQueryType() string {
	if x != nil {
		return x.QueryType
	}
	return ""
}

func (x *WorkflowQuery) GetQueryArgs() *v1.Payloads {
	if x != nil {
		return x.QueryArgs
	}
	return nil
}

func (x *WorkflowQuery) GetHeader() *v1.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

// Answer to a `WorkflowQuery`
type WorkflowQueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Did the query succeed or fail?
	ResultType v11.QueryResultType `protobuf:"varint,1,opt,name=result_type,json=resultType,proto3,enum=temporal.api.enums.v1.QueryResultType" json:"result_type,omitempty"`
	// Set when the query succeeds with the results
	Answer *v1.Payloads `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	// Mutually exclusive with `answer`. Set when the query fails.
	ErrorMessage string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *WorkflowQueryResult) Reset() {
	*x = WorkflowQueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_query_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowQueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowQueryResult) ProtoMessage() {}

func (x *WorkflowQueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_query_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowQueryResult.ProtoReflect.Descriptor instead.
func (*WorkflowQueryResult) Descriptor() ([]byte, []int) {
	return file_temporal_api_query_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowQueryResult) GetResultType() v11.QueryResultType {
	if x != nil {
		return x.ResultType
	}
	return v11.QueryResultType(0)
}

func (x *WorkflowQueryResult) GetAnswer() *v1.Payloads {
	if x != nil {
		return x.Answer
	}
	return nil
}

func (x *WorkflowQueryResult) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type QueryRejected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status v11.WorkflowExecutionStatus `protobuf:"varint,1,opt,name=status,proto3,enum=temporal.api.enums.v1.WorkflowExecutionStatus" json:"status,omitempty"`
}

func (x *QueryRejected) Reset() {
	*x = QueryRejected{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_query_v1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRejected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRejected) ProtoMessage() {}

func (x *QueryRejected) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_query_v1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRejected.ProtoReflect.Descriptor instead.
func (*QueryRejected) Descriptor() ([]byte, []int) {
	return file_temporal_api_query_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *QueryRejected) GetStatus() v11.WorkflowExecutionStatus {
	if x != nil {
		return x.Status
	}
	return v11.WorkflowExecutionStatus(0)
}

var File_temporal_api_query_v1_message_proto protoreflect.FileDescriptor

var file_temporal_api_query_v1_message_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x24, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x0d,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3f, 0x0a, 0x0a,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x73, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x41, 0x72, 0x67, 0x73, 0x12, 0x36, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0xbd, 0x01, 0x0a, 0x13, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x47, 0x0a,
	0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x57, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x46, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x84,
	0x01, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x6f, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x71, 0x75, 0x65, 0x72, 0x79, 0xaa, 0x02,
	0x17, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x1a, 0x54, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_query_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_query_v1_message_proto_rawDescData = file_temporal_api_query_v1_message_proto_rawDesc
)

func file_temporal_api_query_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_query_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_query_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_query_v1_message_proto_rawDescData)
	})
	return file_temporal_api_query_v1_message_proto_rawDescData
}

var file_temporal_api_query_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_temporal_api_query_v1_message_proto_goTypes = []any{
	(*WorkflowQuery)(nil),            // 0: temporal.api.query.v1.WorkflowQuery
	(*WorkflowQueryResult)(nil),      // 1: temporal.api.query.v1.WorkflowQueryResult
	(*QueryRejected)(nil),            // 2: temporal.api.query.v1.QueryRejected
	(*v1.Payloads)(nil),              // 3: temporal.api.common.v1.Payloads
	(*v1.Header)(nil),                // 4: temporal.api.common.v1.Header
	(v11.QueryResultType)(0),         // 5: temporal.api.enums.v1.QueryResultType
	(v11.WorkflowExecutionStatus)(0), // 6: temporal.api.enums.v1.WorkflowExecutionStatus
}
var file_temporal_api_query_v1_message_proto_depIdxs = []int32{
	3, // 0: temporal.api.query.v1.WorkflowQuery.query_args:type_name -> temporal.api.common.v1.Payloads
	4, // 1: temporal.api.query.v1.WorkflowQuery.header:type_name -> temporal.api.common.v1.Header
	5, // 2: temporal.api.query.v1.WorkflowQueryResult.result_type:type_name -> temporal.api.enums.v1.QueryResultType
	3, // 3: temporal.api.query.v1.WorkflowQueryResult.answer:type_name -> temporal.api.common.v1.Payloads
	6, // 4: temporal.api.query.v1.QueryRejected.status:type_name -> temporal.api.enums.v1.WorkflowExecutionStatus
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_temporal_api_query_v1_message_proto_init() }
func file_temporal_api_query_v1_message_proto_init() {
	if File_temporal_api_query_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_api_query_v1_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*WorkflowQuery); i {
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
		file_temporal_api_query_v1_message_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*WorkflowQueryResult); i {
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
		file_temporal_api_query_v1_message_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*QueryRejected); i {
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
			RawDescriptor: file_temporal_api_query_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_query_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_query_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_api_query_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_query_v1_message_proto = out.File
	file_temporal_api_query_v1_message_proto_rawDesc = nil
	file_temporal_api_query_v1_message_proto_goTypes = nil
	file_temporal_api_query_v1_message_proto_depIdxs = nil
}
