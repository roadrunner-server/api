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
// source: temporal/api/enums/v1/common.proto

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

type EncodingType int32

const (
	EncodingType_ENCODING_TYPE_UNSPECIFIED EncodingType = 0
	EncodingType_ENCODING_TYPE_PROTO3      EncodingType = 1
	EncodingType_ENCODING_TYPE_JSON        EncodingType = 2
)

// Enum value maps for EncodingType.
var (
	EncodingType_name = map[int32]string{
		0: "ENCODING_TYPE_UNSPECIFIED",
		1: "ENCODING_TYPE_PROTO3",
		2: "ENCODING_TYPE_JSON",
	}
	EncodingType_value = map[string]int32{
		"ENCODING_TYPE_UNSPECIFIED": 0,
		"ENCODING_TYPE_PROTO3":      1,
		"ENCODING_TYPE_JSON":        2,
	}
)

func (x EncodingType) Enum() *EncodingType {
	p := new(EncodingType)
	*p = x
	return p
}

func (x EncodingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EncodingType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_common_proto_enumTypes[0].Descriptor()
}

func (EncodingType) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_common_proto_enumTypes[0]
}

func (x EncodingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EncodingType.Descriptor instead.
func (EncodingType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_common_proto_rawDescGZIP(), []int{0}
}

type IndexedValueType int32

const (
	IndexedValueType_INDEXED_VALUE_TYPE_UNSPECIFIED  IndexedValueType = 0
	IndexedValueType_INDEXED_VALUE_TYPE_TEXT         IndexedValueType = 1
	IndexedValueType_INDEXED_VALUE_TYPE_KEYWORD      IndexedValueType = 2
	IndexedValueType_INDEXED_VALUE_TYPE_INT          IndexedValueType = 3
	IndexedValueType_INDEXED_VALUE_TYPE_DOUBLE       IndexedValueType = 4
	IndexedValueType_INDEXED_VALUE_TYPE_BOOL         IndexedValueType = 5
	IndexedValueType_INDEXED_VALUE_TYPE_DATETIME     IndexedValueType = 6
	IndexedValueType_INDEXED_VALUE_TYPE_KEYWORD_LIST IndexedValueType = 7
)

// Enum value maps for IndexedValueType.
var (
	IndexedValueType_name = map[int32]string{
		0: "INDEXED_VALUE_TYPE_UNSPECIFIED",
		1: "INDEXED_VALUE_TYPE_TEXT",
		2: "INDEXED_VALUE_TYPE_KEYWORD",
		3: "INDEXED_VALUE_TYPE_INT",
		4: "INDEXED_VALUE_TYPE_DOUBLE",
		5: "INDEXED_VALUE_TYPE_BOOL",
		6: "INDEXED_VALUE_TYPE_DATETIME",
		7: "INDEXED_VALUE_TYPE_KEYWORD_LIST",
	}
	IndexedValueType_value = map[string]int32{
		"INDEXED_VALUE_TYPE_UNSPECIFIED":  0,
		"INDEXED_VALUE_TYPE_TEXT":         1,
		"INDEXED_VALUE_TYPE_KEYWORD":      2,
		"INDEXED_VALUE_TYPE_INT":          3,
		"INDEXED_VALUE_TYPE_DOUBLE":       4,
		"INDEXED_VALUE_TYPE_BOOL":         5,
		"INDEXED_VALUE_TYPE_DATETIME":     6,
		"INDEXED_VALUE_TYPE_KEYWORD_LIST": 7,
	}
)

func (x IndexedValueType) Enum() *IndexedValueType {
	p := new(IndexedValueType)
	*p = x
	return p
}

func (x IndexedValueType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IndexedValueType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_common_proto_enumTypes[1].Descriptor()
}

func (IndexedValueType) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_common_proto_enumTypes[1]
}

func (x IndexedValueType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IndexedValueType.Descriptor instead.
func (IndexedValueType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_common_proto_rawDescGZIP(), []int{1}
}

type Severity int32

const (
	Severity_SEVERITY_UNSPECIFIED Severity = 0
	Severity_SEVERITY_HIGH        Severity = 1
	Severity_SEVERITY_MEDIUM      Severity = 2
	Severity_SEVERITY_LOW         Severity = 3
)

// Enum value maps for Severity.
var (
	Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "SEVERITY_HIGH",
		2: "SEVERITY_MEDIUM",
		3: "SEVERITY_LOW",
	}
	Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"SEVERITY_HIGH":        1,
		"SEVERITY_MEDIUM":      2,
		"SEVERITY_LOW":         3,
	}
)

func (x Severity) Enum() *Severity {
	p := new(Severity)
	*p = x
	return p
}

func (x Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_common_proto_enumTypes[2].Descriptor()
}

func (Severity) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_common_proto_enumTypes[2]
}

func (x Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Severity.Descriptor instead.
func (Severity) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_common_proto_rawDescGZIP(), []int{2}
}

// State of a callback.
type CallbackState int32

const (
	// Default value, unspecified state.
	CallbackState_CALLBACK_STATE_UNSPECIFIED CallbackState = 0
	// Callback is standing by, waiting to be triggered.
	CallbackState_CALLBACK_STATE_STANDBY CallbackState = 1
	// Callback is in the queue waiting to be executed or is currently executing.
	CallbackState_CALLBACK_STATE_SCHEDULED CallbackState = 2
	// Callback has failed with a retryable error and is backing off before the next attempt.
	CallbackState_CALLBACK_STATE_BACKING_OFF CallbackState = 3
	// Callback has failed.
	CallbackState_CALLBACK_STATE_FAILED CallbackState = 4
	// Callback has succeeded.
	CallbackState_CALLBACK_STATE_SUCCEEDED CallbackState = 5
	// Callback is blocked (eg: by circuit breaker).
	CallbackState_CALLBACK_STATE_BLOCKED CallbackState = 6
)

// Enum value maps for CallbackState.
var (
	CallbackState_name = map[int32]string{
		0: "CALLBACK_STATE_UNSPECIFIED",
		1: "CALLBACK_STATE_STANDBY",
		2: "CALLBACK_STATE_SCHEDULED",
		3: "CALLBACK_STATE_BACKING_OFF",
		4: "CALLBACK_STATE_FAILED",
		5: "CALLBACK_STATE_SUCCEEDED",
		6: "CALLBACK_STATE_BLOCKED",
	}
	CallbackState_value = map[string]int32{
		"CALLBACK_STATE_UNSPECIFIED": 0,
		"CALLBACK_STATE_STANDBY":     1,
		"CALLBACK_STATE_SCHEDULED":   2,
		"CALLBACK_STATE_BACKING_OFF": 3,
		"CALLBACK_STATE_FAILED":      4,
		"CALLBACK_STATE_SUCCEEDED":   5,
		"CALLBACK_STATE_BLOCKED":     6,
	}
)

func (x CallbackState) Enum() *CallbackState {
	p := new(CallbackState)
	*p = x
	return p
}

func (x CallbackState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CallbackState) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_common_proto_enumTypes[3].Descriptor()
}

func (CallbackState) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_common_proto_enumTypes[3]
}

func (x CallbackState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CallbackState.Descriptor instead.
func (CallbackState) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_common_proto_rawDescGZIP(), []int{3}
}

// State of a pending Nexus operation.
type PendingNexusOperationState int32

const (
	// Default value, unspecified state.
	PendingNexusOperationState_PENDING_NEXUS_OPERATION_STATE_UNSPECIFIED PendingNexusOperationState = 0
	// Operation is in the queue waiting to be executed or is currently executing.
	PendingNexusOperationState_PENDING_NEXUS_OPERATION_STATE_SCHEDULED PendingNexusOperationState = 1
	// Operation has failed with a retryable error and is backing off before the next attempt.
	PendingNexusOperationState_PENDING_NEXUS_OPERATION_STATE_BACKING_OFF PendingNexusOperationState = 2
	// Operation was started and will complete asynchronously.
	PendingNexusOperationState_PENDING_NEXUS_OPERATION_STATE_STARTED PendingNexusOperationState = 3
	// Operation is blocked (eg: by circuit breaker).
	PendingNexusOperationState_PENDING_NEXUS_OPERATION_STATE_BLOCKED PendingNexusOperationState = 4
)

// Enum value maps for PendingNexusOperationState.
var (
	PendingNexusOperationState_name = map[int32]string{
		0: "PENDING_NEXUS_OPERATION_STATE_UNSPECIFIED",
		1: "PENDING_NEXUS_OPERATION_STATE_SCHEDULED",
		2: "PENDING_NEXUS_OPERATION_STATE_BACKING_OFF",
		3: "PENDING_NEXUS_OPERATION_STATE_STARTED",
		4: "PENDING_NEXUS_OPERATION_STATE_BLOCKED",
	}
	PendingNexusOperationState_value = map[string]int32{
		"PENDING_NEXUS_OPERATION_STATE_UNSPECIFIED": 0,
		"PENDING_NEXUS_OPERATION_STATE_SCHEDULED":   1,
		"PENDING_NEXUS_OPERATION_STATE_BACKING_OFF": 2,
		"PENDING_NEXUS_OPERATION_STATE_STARTED":     3,
		"PENDING_NEXUS_OPERATION_STATE_BLOCKED":     4,
	}
)

func (x PendingNexusOperationState) Enum() *PendingNexusOperationState {
	p := new(PendingNexusOperationState)
	*p = x
	return p
}

func (x PendingNexusOperationState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PendingNexusOperationState) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_common_proto_enumTypes[4].Descriptor()
}

func (PendingNexusOperationState) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_common_proto_enumTypes[4]
}

func (x PendingNexusOperationState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PendingNexusOperationState.Descriptor instead.
func (PendingNexusOperationState) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_common_proto_rawDescGZIP(), []int{4}
}

// State of a Nexus operation cancellation.
type NexusOperationCancellationState int32

const (
	// Default value, unspecified state.
	NexusOperationCancellationState_NEXUS_OPERATION_CANCELLATION_STATE_UNSPECIFIED NexusOperationCancellationState = 0
	// Cancellation request is in the queue waiting to be executed or is currently executing.
	NexusOperationCancellationState_NEXUS_OPERATION_CANCELLATION_STATE_SCHEDULED NexusOperationCancellationState = 1
	// Cancellation request has failed with a retryable error and is backing off before the next attempt.
	NexusOperationCancellationState_NEXUS_OPERATION_CANCELLATION_STATE_BACKING_OFF NexusOperationCancellationState = 2
	// Cancellation request succeeded.
	NexusOperationCancellationState_NEXUS_OPERATION_CANCELLATION_STATE_SUCCEEDED NexusOperationCancellationState = 3
	// Cancellation request failed with a non-retryable error.
	NexusOperationCancellationState_NEXUS_OPERATION_CANCELLATION_STATE_FAILED NexusOperationCancellationState = 4
	// The associated operation timed out - exceeded the user supplied schedule-to-close timeout.
	NexusOperationCancellationState_NEXUS_OPERATION_CANCELLATION_STATE_TIMED_OUT NexusOperationCancellationState = 5
	// Cancellation request is blocked (eg: by circuit breaker).
	NexusOperationCancellationState_NEXUS_OPERATION_CANCELLATION_STATE_BLOCKED NexusOperationCancellationState = 6
)

// Enum value maps for NexusOperationCancellationState.
var (
	NexusOperationCancellationState_name = map[int32]string{
		0: "NEXUS_OPERATION_CANCELLATION_STATE_UNSPECIFIED",
		1: "NEXUS_OPERATION_CANCELLATION_STATE_SCHEDULED",
		2: "NEXUS_OPERATION_CANCELLATION_STATE_BACKING_OFF",
		3: "NEXUS_OPERATION_CANCELLATION_STATE_SUCCEEDED",
		4: "NEXUS_OPERATION_CANCELLATION_STATE_FAILED",
		5: "NEXUS_OPERATION_CANCELLATION_STATE_TIMED_OUT",
		6: "NEXUS_OPERATION_CANCELLATION_STATE_BLOCKED",
	}
	NexusOperationCancellationState_value = map[string]int32{
		"NEXUS_OPERATION_CANCELLATION_STATE_UNSPECIFIED": 0,
		"NEXUS_OPERATION_CANCELLATION_STATE_SCHEDULED":   1,
		"NEXUS_OPERATION_CANCELLATION_STATE_BACKING_OFF": 2,
		"NEXUS_OPERATION_CANCELLATION_STATE_SUCCEEDED":   3,
		"NEXUS_OPERATION_CANCELLATION_STATE_FAILED":      4,
		"NEXUS_OPERATION_CANCELLATION_STATE_TIMED_OUT":   5,
		"NEXUS_OPERATION_CANCELLATION_STATE_BLOCKED":     6,
	}
)

func (x NexusOperationCancellationState) Enum() *NexusOperationCancellationState {
	p := new(NexusOperationCancellationState)
	*p = x
	return p
}

func (x NexusOperationCancellationState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NexusOperationCancellationState) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_common_proto_enumTypes[5].Descriptor()
}

func (NexusOperationCancellationState) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_common_proto_enumTypes[5]
}

func (x NexusOperationCancellationState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NexusOperationCancellationState.Descriptor instead.
func (NexusOperationCancellationState) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_common_proto_rawDescGZIP(), []int{5}
}

var File_temporal_api_enums_v1_common_proto protoreflect.FileDescriptor

var file_temporal_api_enums_v1_common_proto_rawDesc = []byte{
	0x0a, 0x22, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2a, 0x5f, 0x0a, 0x0c, 0x45,
	0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x45,
	0x4e, 0x43, 0x4f, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x4e,
	0x43, 0x4f, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x54,
	0x4f, 0x33, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x4e, 0x43, 0x4f, 0x44, 0x49, 0x4e, 0x47,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x02, 0x2a, 0x91, 0x02, 0x0a,
	0x10, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x22, 0x0a, 0x1e, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x4c,
	0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x45, 0x44,
	0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x45, 0x58, 0x54,
	0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44,
	0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x1d,
	0x0a, 0x19, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x1b, 0x0a,
	0x17, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x05, 0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4e,
	0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x44, 0x41, 0x54, 0x45, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x06, 0x12, 0x23, 0x0a, 0x1f, 0x49,
	0x4e, 0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x07,
	0x2a, 0x5e, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x14,
	0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x56,
	0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x02, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x03,
	0x2a, 0xde, 0x01, 0x0a, 0x0d, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x42, 0x59, 0x10, 0x01, 0x12, 0x1c,
	0x0a, 0x18, 0x43, 0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a,
	0x43, 0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x42,
	0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15,
	0x43, 0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x41, 0x4c, 0x4c, 0x42,
	0x41, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45,
	0x44, 0x45, 0x44, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43,
	0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10,
	0x06, 0x2a, 0xfd, 0x01, 0x0a, 0x1a, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4e, 0x65, 0x78,
	0x75, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x2d, 0x0a, 0x29, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x4e, 0x45, 0x58, 0x55,
	0x53, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x2b, 0x0a, 0x27, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x4e, 0x45, 0x58, 0x55, 0x53,
	0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x2d, 0x0a, 0x29,
	0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x4e, 0x45, 0x58, 0x55, 0x53, 0x5f, 0x4f, 0x50,
	0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x42, 0x41,
	0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x02, 0x12, 0x29, 0x0a, 0x25, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x4e, 0x45, 0x58, 0x55, 0x53, 0x5f, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x52, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x29, 0x0a, 0x25, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x5f, 0x4e, 0x45, 0x58, 0x55, 0x53, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10,
	0x04, 0x2a, 0xfe, 0x02, 0x0a, 0x1f, 0x4e, 0x65, 0x78, 0x75, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x2e, 0x4e, 0x45, 0x58, 0x55, 0x53, 0x5f, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x30, 0x0a, 0x2c, 0x4e, 0x45, 0x58,
	0x55, 0x53, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e,
	0x43, 0x45, 0x4c, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x32, 0x0a, 0x2e, 0x4e,
	0x45, 0x58, 0x55, 0x53, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x02, 0x12,
	0x30, 0x0a, 0x2c, 0x4e, 0x45, 0x58, 0x55, 0x53, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x2d, 0x0a, 0x29, 0x4e, 0x45, 0x58, 0x55, 0x53, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04,
	0x12, 0x30, 0x0a, 0x2c, 0x4e, 0x45, 0x58, 0x55, 0x53, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x44, 0x5f, 0x4f, 0x55, 0x54,
	0x10, 0x05, 0x12, 0x2e, 0x0a, 0x2a, 0x4e, 0x45, 0x58, 0x55, 0x53, 0x5f, 0x4f, 0x50, 0x45, 0x52,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44,
	0x10, 0x06, 0x42, 0x83, 0x01, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42,
	0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21,
	0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0xaa, 0x02, 0x17, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41,
	0x70, 0x69, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x1a, 0x54, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x45,
	0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_enums_v1_common_proto_rawDescOnce sync.Once
	file_temporal_api_enums_v1_common_proto_rawDescData = file_temporal_api_enums_v1_common_proto_rawDesc
)

func file_temporal_api_enums_v1_common_proto_rawDescGZIP() []byte {
	file_temporal_api_enums_v1_common_proto_rawDescOnce.Do(func() {
		file_temporal_api_enums_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_enums_v1_common_proto_rawDescData)
	})
	return file_temporal_api_enums_v1_common_proto_rawDescData
}

var file_temporal_api_enums_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_temporal_api_enums_v1_common_proto_goTypes = []any{
	(EncodingType)(0),                    // 0: temporal.api.enums.v1.EncodingType
	(IndexedValueType)(0),                // 1: temporal.api.enums.v1.IndexedValueType
	(Severity)(0),                        // 2: temporal.api.enums.v1.Severity
	(CallbackState)(0),                   // 3: temporal.api.enums.v1.CallbackState
	(PendingNexusOperationState)(0),      // 4: temporal.api.enums.v1.PendingNexusOperationState
	(NexusOperationCancellationState)(0), // 5: temporal.api.enums.v1.NexusOperationCancellationState
}
var file_temporal_api_enums_v1_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_enums_v1_common_proto_init() }
func file_temporal_api_enums_v1_common_proto_init() {
	if File_temporal_api_enums_v1_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_enums_v1_common_proto_rawDesc,
			NumEnums:      6,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_enums_v1_common_proto_goTypes,
		DependencyIndexes: file_temporal_api_enums_v1_common_proto_depIdxs,
		EnumInfos:         file_temporal_api_enums_v1_common_proto_enumTypes,
	}.Build()
	File_temporal_api_enums_v1_common_proto = out.File
	file_temporal_api_enums_v1_common_proto_rawDesc = nil
	file_temporal_api_enums_v1_common_proto_goTypes = nil
	file_temporal_api_enums_v1_common_proto_depIdxs = nil
}
