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
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: temporal/api/enums/v1/reset.proto

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

// Reset reapplay(replay) options
// * RESET_REAPPLY_TYPE_SIGNAL (default) - Signals are reapplied when workflow is reset
// * RESET_REAPPLY_TYPE_NONE - nothing is reapplied
type ResetReapplyType int32

const (
	ResetReapplyType_RESET_REAPPLY_TYPE_UNSPECIFIED ResetReapplyType = 0
	ResetReapplyType_RESET_REAPPLY_TYPE_SIGNAL      ResetReapplyType = 1
	ResetReapplyType_RESET_REAPPLY_TYPE_NONE        ResetReapplyType = 2
)

// Enum value maps for ResetReapplyType.
var (
	ResetReapplyType_name = map[int32]string{
		0: "RESET_REAPPLY_TYPE_UNSPECIFIED",
		1: "RESET_REAPPLY_TYPE_SIGNAL",
		2: "RESET_REAPPLY_TYPE_NONE",
	}
	ResetReapplyType_value = map[string]int32{
		"RESET_REAPPLY_TYPE_UNSPECIFIED": 0,
		"RESET_REAPPLY_TYPE_SIGNAL":      1,
		"RESET_REAPPLY_TYPE_NONE":        2,
	}
)

func (x ResetReapplyType) Enum() *ResetReapplyType {
	p := new(ResetReapplyType)
	*p = x
	return p
}

func (x ResetReapplyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResetReapplyType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_reset_proto_enumTypes[0].Descriptor()
}

func (ResetReapplyType) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_reset_proto_enumTypes[0]
}

func (x ResetReapplyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResetReapplyType.Descriptor instead.
func (ResetReapplyType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_reset_proto_rawDescGZIP(), []int{0}
}

// Reset type options
type ResetType int32

const (
	ResetType_RESET_TYPE_UNSPECIFIED ResetType = 0
	// Resets to event of the first workflow task completed, or if it does not exist, the event after task scheduled.
	ResetType_RESET_TYPE_FIRST_WORKFLOW_TASK ResetType = 1
	// Resets to event of the last workflow task completed, or if it does not exist, the event after task scheduled.
	ResetType_RESET_TYPE_LAST_WORKFLOW_TASK ResetType = 2
)

// Enum value maps for ResetType.
var (
	ResetType_name = map[int32]string{
		0: "RESET_TYPE_UNSPECIFIED",
		1: "RESET_TYPE_FIRST_WORKFLOW_TASK",
		2: "RESET_TYPE_LAST_WORKFLOW_TASK",
	}
	ResetType_value = map[string]int32{
		"RESET_TYPE_UNSPECIFIED":         0,
		"RESET_TYPE_FIRST_WORKFLOW_TASK": 1,
		"RESET_TYPE_LAST_WORKFLOW_TASK":  2,
	}
)

func (x ResetType) Enum() *ResetType {
	p := new(ResetType)
	*p = x
	return p
}

func (x ResetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResetType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_reset_proto_enumTypes[1].Descriptor()
}

func (ResetType) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_reset_proto_enumTypes[1]
}

func (x ResetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResetType.Descriptor instead.
func (ResetType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_reset_proto_rawDescGZIP(), []int{1}
}

var File_temporal_api_enums_v1_reset_proto protoreflect.FileDescriptor

var file_temporal_api_enums_v1_reset_proto_rawDesc = []byte{
	0x0a, 0x21, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2a, 0x72, 0x0a, 0x10, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22,
	0x0a, 0x1e, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x50,
	0x50, 0x4c, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x4c, 0x10,
	0x01, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x50, 0x50,
	0x4c, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x2a, 0x6e,
	0x0a, 0x09, 0x52, 0x65, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x52,
	0x45, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x52, 0x45, 0x53, 0x45, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x57, 0x4f, 0x52, 0x4b,
	0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x52,
	0x45, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x57,
	0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x10, 0x02, 0x42, 0x82,
	0x01, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x6f, 0x2e, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xaa, 0x02, 0x17, 0x54,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x1a, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_enums_v1_reset_proto_rawDescOnce sync.Once
	file_temporal_api_enums_v1_reset_proto_rawDescData = file_temporal_api_enums_v1_reset_proto_rawDesc
)

func file_temporal_api_enums_v1_reset_proto_rawDescGZIP() []byte {
	file_temporal_api_enums_v1_reset_proto_rawDescOnce.Do(func() {
		file_temporal_api_enums_v1_reset_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_enums_v1_reset_proto_rawDescData)
	})
	return file_temporal_api_enums_v1_reset_proto_rawDescData
}

var file_temporal_api_enums_v1_reset_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_temporal_api_enums_v1_reset_proto_goTypes = []interface{}{
	(ResetReapplyType)(0), // 0: temporal.api.enums.v1.ResetReapplyType
	(ResetType)(0),        // 1: temporal.api.enums.v1.ResetType
}
var file_temporal_api_enums_v1_reset_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_enums_v1_reset_proto_init() }
func file_temporal_api_enums_v1_reset_proto_init() {
	if File_temporal_api_enums_v1_reset_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_enums_v1_reset_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_enums_v1_reset_proto_goTypes,
		DependencyIndexes: file_temporal_api_enums_v1_reset_proto_depIdxs,
		EnumInfos:         file_temporal_api_enums_v1_reset_proto_enumTypes,
	}.Build()
	File_temporal_api_enums_v1_reset_proto = out.File
	file_temporal_api_enums_v1_reset_proto_rawDesc = nil
	file_temporal_api_enums_v1_reset_proto_goTypes = nil
	file_temporal_api_enums_v1_reset_proto_depIdxs = nil
}
