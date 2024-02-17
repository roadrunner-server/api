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
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: temporal/api/enums/v1/query.proto

package enums

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

type QueryResultType int32

const (
	QueryResultType_QUERY_RESULT_TYPE_UNSPECIFIED QueryResultType = 0
	QueryResultType_QUERY_RESULT_TYPE_ANSWERED    QueryResultType = 1
	QueryResultType_QUERY_RESULT_TYPE_FAILED      QueryResultType = 2
)

// Enum value maps for QueryResultType.
var (
	QueryResultType_name = map[int32]string{
		0: "QUERY_RESULT_TYPE_UNSPECIFIED",
		1: "QUERY_RESULT_TYPE_ANSWERED",
		2: "QUERY_RESULT_TYPE_FAILED",
	}
	QueryResultType_value = map[string]int32{
		"QUERY_RESULT_TYPE_UNSPECIFIED": 0,
		"QUERY_RESULT_TYPE_ANSWERED":    1,
		"QUERY_RESULT_TYPE_FAILED":      2,
	}
)

func (x QueryResultType) Enum() *QueryResultType {
	p := new(QueryResultType)
	*p = x
	return p
}

func (x QueryResultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_query_proto_enumTypes[0].Descriptor()
}

func (QueryResultType) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_query_proto_enumTypes[0]
}

func (x QueryResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryResultType.Descriptor instead.
func (QueryResultType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_query_proto_rawDescGZIP(), []int{0}
}

type QueryRejectCondition int32

const (
	QueryRejectCondition_QUERY_REJECT_CONDITION_UNSPECIFIED QueryRejectCondition = 0
	// None indicates that query should not be rejected.
	QueryRejectCondition_QUERY_REJECT_CONDITION_NONE QueryRejectCondition = 1
	// NotOpen indicates that query should be rejected if workflow is not open.
	QueryRejectCondition_QUERY_REJECT_CONDITION_NOT_OPEN QueryRejectCondition = 2
	// NotCompletedCleanly indicates that query should be rejected if workflow did not complete cleanly.
	QueryRejectCondition_QUERY_REJECT_CONDITION_NOT_COMPLETED_CLEANLY QueryRejectCondition = 3
)

// Enum value maps for QueryRejectCondition.
var (
	QueryRejectCondition_name = map[int32]string{
		0: "QUERY_REJECT_CONDITION_UNSPECIFIED",
		1: "QUERY_REJECT_CONDITION_NONE",
		2: "QUERY_REJECT_CONDITION_NOT_OPEN",
		3: "QUERY_REJECT_CONDITION_NOT_COMPLETED_CLEANLY",
	}
	QueryRejectCondition_value = map[string]int32{
		"QUERY_REJECT_CONDITION_UNSPECIFIED":           0,
		"QUERY_REJECT_CONDITION_NONE":                  1,
		"QUERY_REJECT_CONDITION_NOT_OPEN":              2,
		"QUERY_REJECT_CONDITION_NOT_COMPLETED_CLEANLY": 3,
	}
)

func (x QueryRejectCondition) Enum() *QueryRejectCondition {
	p := new(QueryRejectCondition)
	*p = x
	return p
}

func (x QueryRejectCondition) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryRejectCondition) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_query_proto_enumTypes[1].Descriptor()
}

func (QueryRejectCondition) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_query_proto_enumTypes[1]
}

func (x QueryRejectCondition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryRejectCondition.Descriptor instead.
func (QueryRejectCondition) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_query_proto_rawDescGZIP(), []int{1}
}

var File_temporal_api_enums_v1_query_proto protoreflect.FileDescriptor

var file_temporal_api_enums_v1_query_proto_rawDesc = []byte{
	0x0a, 0x21, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2a, 0x72, 0x0a, 0x0f, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a,
	0x1d, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x1e, 0x0a, 0x1a, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x1c, 0x0a, 0x18, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x2a, 0xb6,
	0x01, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x22, 0x51, 0x55, 0x45, 0x52, 0x59,
	0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x1f, 0x0a, 0x1b, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x5f,
	0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01,
	0x12, 0x23, 0x0a, 0x1f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54,
	0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4f,
	0x50, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x30, 0x0a, 0x2c, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x52,
	0x45, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x5f, 0x43, 0x4c,
	0x45, 0x41, 0x4e, 0x4c, 0x59, 0x10, 0x03, 0x42, 0x82, 0x01, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x21, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0xaa, 0x02, 0x17, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x56, 0x31, 0xea,
	0x02, 0x1a, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70,
	0x69, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_enums_v1_query_proto_rawDescOnce sync.Once
	file_temporal_api_enums_v1_query_proto_rawDescData = file_temporal_api_enums_v1_query_proto_rawDesc
)

func file_temporal_api_enums_v1_query_proto_rawDescGZIP() []byte {
	file_temporal_api_enums_v1_query_proto_rawDescOnce.Do(func() {
		file_temporal_api_enums_v1_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_enums_v1_query_proto_rawDescData)
	})
	return file_temporal_api_enums_v1_query_proto_rawDescData
}

var file_temporal_api_enums_v1_query_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_temporal_api_enums_v1_query_proto_goTypes = []interface{}{
	(QueryResultType)(0),      // 0: temporal.api.enums.v1.QueryResultType
	(QueryRejectCondition)(0), // 1: temporal.api.enums.v1.QueryRejectCondition
}
var file_temporal_api_enums_v1_query_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_enums_v1_query_proto_init() }
func file_temporal_api_enums_v1_query_proto_init() {
	if File_temporal_api_enums_v1_query_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_enums_v1_query_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_enums_v1_query_proto_goTypes,
		DependencyIndexes: file_temporal_api_enums_v1_query_proto_depIdxs,
		EnumInfos:         file_temporal_api_enums_v1_query_proto_enumTypes,
	}.Build()
	File_temporal_api_enums_v1_query_proto = out.File
	file_temporal_api_enums_v1_query_proto_rawDesc = nil
	file_temporal_api_enums_v1_query_proto_goTypes = nil
	file_temporal_api_enums_v1_query_proto_depIdxs = nil
}
