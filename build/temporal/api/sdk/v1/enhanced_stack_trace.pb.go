// The MIT License
//
// Copyright (c) 2024 Temporal Technologies Inc.  All rights reserved.
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
// source: temporal/api/sdk/v1/enhanced_stack_trace.proto

package sdk

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

// Internal structure used to create worker stack traces with references to code.
type EnhancedStackTrace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Information pertaining to the SDK that the trace has been captured from.
	Sdk *StackTraceSDKInfo `protobuf:"bytes,1,opt,name=sdk,proto3" json:"sdk,omitempty"`
	// Mapping of file path to file contents.
	Sources map[string]*StackTraceFileSlice `protobuf:"bytes,2,rep,name=sources,proto3" json:"sources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Collection of stacks captured.
	Stacks []*StackTrace `protobuf:"bytes,3,rep,name=stacks,proto3" json:"stacks,omitempty"`
}

func (x *EnhancedStackTrace) Reset() {
	*x = EnhancedStackTrace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_sdk_v1_enhanced_stack_trace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnhancedStackTrace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnhancedStackTrace) ProtoMessage() {}

func (x *EnhancedStackTrace) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_sdk_v1_enhanced_stack_trace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnhancedStackTrace.ProtoReflect.Descriptor instead.
func (*EnhancedStackTrace) Descriptor() ([]byte, []int) {
	return file_temporal_api_sdk_v1_enhanced_stack_trace_proto_rawDescGZIP(), []int{0}
}

func (x *EnhancedStackTrace) GetSdk() *StackTraceSDKInfo {
	if x != nil {
		return x.Sdk
	}
	return nil
}

func (x *EnhancedStackTrace) GetSources() map[string]*StackTraceFileSlice {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *EnhancedStackTrace) GetStacks() []*StackTrace {
	if x != nil {
		return x.Stacks
	}
	return nil
}

// Information pertaining to the SDK that the trace has been captured from.
// (-- api-linter: core::0123::resource-annotation=disabled
//
//	aip.dev/not-precedent: Naming SDK version is optional. --)
type StackTraceSDKInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the SDK
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Version string of the SDK
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *StackTraceSDKInfo) Reset() {
	*x = StackTraceSDKInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_sdk_v1_enhanced_stack_trace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackTraceSDKInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackTraceSDKInfo) ProtoMessage() {}

func (x *StackTraceSDKInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_sdk_v1_enhanced_stack_trace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackTraceSDKInfo.ProtoReflect.Descriptor instead.
func (*StackTraceSDKInfo) Descriptor() ([]byte, []int) {
	return file_temporal_api_sdk_v1_enhanced_stack_trace_proto_rawDescGZIP(), []int{1}
}

func (x *StackTraceSDKInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StackTraceSDKInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// "Slice" of a file starting at line_offset -- a line offset and code fragment corresponding to the worker's stack.
type StackTraceFileSlice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Only used (possibly) to trim the file without breaking syntax highlighting. This is not optional, unlike
	// the `line` property of a `StackTraceFileLocation`.
	// (-- api-linter: core::0141::forbidden-types=disabled
	//
	//	aip.dev/not-precedent: These really shouldn't have negative values. --)
	LineOffset uint32 `protobuf:"varint,1,opt,name=line_offset,json=lineOffset,proto3" json:"line_offset,omitempty"`
	// Slice of a file with the respective OS-specific line terminator.
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *StackTraceFileSlice) Reset() {
	*x = StackTraceFileSlice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_sdk_v1_enhanced_stack_trace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackTraceFileSlice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackTraceFileSlice) ProtoMessage() {}

func (x *StackTraceFileSlice) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_sdk_v1_enhanced_stack_trace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackTraceFileSlice.ProtoReflect.Descriptor instead.
func (*StackTraceFileSlice) Descriptor() ([]byte, []int) {
	return file_temporal_api_sdk_v1_enhanced_stack_trace_proto_rawDescGZIP(), []int{2}
}

func (x *StackTraceFileSlice) GetLineOffset() uint32 {
	if x != nil {
		return x.LineOffset
	}
	return 0
}

func (x *StackTraceFileSlice) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// More specific location details of a file: its path, precise line and column numbers if applicable, and function name if available.
// In essence, a pointer to a location in a file
type StackTraceFileLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path to source file (absolute or relative).
	// If the paths are relative, ensure that they are all relative to the same root.
	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	// Optional; If possible, SDK should send this -- this is required for displaying the code location.
	// If not provided, set to -1.
	Line int32 `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
	// Optional; if possible, SDK should send this.
	// If not provided, set to -1.
	Column int32 `protobuf:"varint,3,opt,name=column,proto3" json:"column,omitempty"`
	// Function name this line belongs to, if applicable.
	// Used for falling back to stack trace view.
	FunctionName string `protobuf:"bytes,4,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	// Flag to communicate whether a location should be hidden by default in the stack view.
	InternalCode bool `protobuf:"varint,5,opt,name=internal_code,json=internalCode,proto3" json:"internal_code,omitempty"`
}

func (x *StackTraceFileLocation) Reset() {
	*x = StackTraceFileLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_sdk_v1_enhanced_stack_trace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackTraceFileLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackTraceFileLocation) ProtoMessage() {}

func (x *StackTraceFileLocation) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_sdk_v1_enhanced_stack_trace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackTraceFileLocation.ProtoReflect.Descriptor instead.
func (*StackTraceFileLocation) Descriptor() ([]byte, []int) {
	return file_temporal_api_sdk_v1_enhanced_stack_trace_proto_rawDescGZIP(), []int{3}
}

func (x *StackTraceFileLocation) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *StackTraceFileLocation) GetLine() int32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *StackTraceFileLocation) GetColumn() int32 {
	if x != nil {
		return x.Column
	}
	return 0
}

func (x *StackTraceFileLocation) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *StackTraceFileLocation) GetInternalCode() bool {
	if x != nil {
		return x.InternalCode
	}
	return false
}

// Collection of FileLocation messages from a single stack.
type StackTrace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Collection of `FileLocation`s, each for a stack frame that comprise a stack trace.
	Locations []*StackTraceFileLocation `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
}

func (x *StackTrace) Reset() {
	*x = StackTrace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_sdk_v1_enhanced_stack_trace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackTrace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackTrace) ProtoMessage() {}

func (x *StackTrace) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_sdk_v1_enhanced_stack_trace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackTrace.ProtoReflect.Descriptor instead.
func (*StackTrace) Descriptor() ([]byte, []int) {
	return file_temporal_api_sdk_v1_enhanced_stack_trace_proto_rawDescGZIP(), []int{4}
}

func (x *StackTrace) GetLocations() []*StackTraceFileLocation {
	if x != nil {
		return x.Locations
	}
	return nil
}

var File_temporal_api_sdk_v1_enhanced_stack_trace_proto protoreflect.FileDescriptor

var file_temporal_api_sdk_v1_enhanced_stack_trace_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x64, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x76, 0x31, 0x22, 0xbd, 0x02, 0x0a, 0x12, 0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63,
	0x65, 0x64, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x03,
	0x73, 0x64, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x44, 0x4b, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x03, 0x73, 0x64, 0x6b, 0x12, 0x4e, 0x0a, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e,
	0x68, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x73, 0x1a,
	0x64, 0x0a, 0x0c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x3e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x41, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x53, 0x44, 0x4b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x50, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xab, 0x01, 0x0a, 0x16, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x23,
	0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x57, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x85, 0x01, 0x0a, 0x16, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x42, 0x17, 0x45, 0x6e,
	0x68, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1d, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x64, 0x6b, 0x2f,
	0x76, 0x31, 0x3b, 0x73, 0x64, 0x6b, 0xaa, 0x02, 0x15, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x64, 0x6b, 0x2e, 0x56, 0x31, 0xea, 0x02,
	0x18, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69,
	0x3a, 0x3a, 0x53, 0x64, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_temporal_api_sdk_v1_enhanced_stack_trace_proto_rawDescOnce sync.Once
	file_temporal_api_sdk_v1_enhanced_stack_trace_proto_rawDescData = file_temporal_api_sdk_v1_enhanced_stack_trace_proto_rawDesc
)

func file_temporal_api_sdk_v1_enhanced_stack_trace_proto_rawDescGZIP() []byte {
	file_temporal_api_sdk_v1_enhanced_stack_trace_proto_rawDescOnce.Do(func() {
		file_temporal_api_sdk_v1_enhanced_stack_trace_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_sdk_v1_enhanced_stack_trace_proto_rawDescData)
	})
	return file_temporal_api_sdk_v1_enhanced_stack_trace_proto_rawDescData
}

var file_temporal_api_sdk_v1_enhanced_stack_trace_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_temporal_api_sdk_v1_enhanced_stack_trace_proto_goTypes = []any{
	(*EnhancedStackTrace)(nil),     // 0: temporal.api.sdk.v1.EnhancedStackTrace
	(*StackTraceSDKInfo)(nil),      // 1: temporal.api.sdk.v1.StackTraceSDKInfo
	(*StackTraceFileSlice)(nil),    // 2: temporal.api.sdk.v1.StackTraceFileSlice
	(*StackTraceFileLocation)(nil), // 3: temporal.api.sdk.v1.StackTraceFileLocation
	(*StackTrace)(nil),             // 4: temporal.api.sdk.v1.StackTrace
	nil,                            // 5: temporal.api.sdk.v1.EnhancedStackTrace.SourcesEntry
}
var file_temporal_api_sdk_v1_enhanced_stack_trace_proto_depIdxs = []int32{
	1, // 0: temporal.api.sdk.v1.EnhancedStackTrace.sdk:type_name -> temporal.api.sdk.v1.StackTraceSDKInfo
	5, // 1: temporal.api.sdk.v1.EnhancedStackTrace.sources:type_name -> temporal.api.sdk.v1.EnhancedStackTrace.SourcesEntry
	4, // 2: temporal.api.sdk.v1.EnhancedStackTrace.stacks:type_name -> temporal.api.sdk.v1.StackTrace
	3, // 3: temporal.api.sdk.v1.StackTrace.locations:type_name -> temporal.api.sdk.v1.StackTraceFileLocation
	2, // 4: temporal.api.sdk.v1.EnhancedStackTrace.SourcesEntry.value:type_name -> temporal.api.sdk.v1.StackTraceFileSlice
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_temporal_api_sdk_v1_enhanced_stack_trace_proto_init() }
func file_temporal_api_sdk_v1_enhanced_stack_trace_proto_init() {
	if File_temporal_api_sdk_v1_enhanced_stack_trace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_api_sdk_v1_enhanced_stack_trace_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EnhancedStackTrace); i {
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
		file_temporal_api_sdk_v1_enhanced_stack_trace_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*StackTraceSDKInfo); i {
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
		file_temporal_api_sdk_v1_enhanced_stack_trace_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*StackTraceFileSlice); i {
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
		file_temporal_api_sdk_v1_enhanced_stack_trace_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*StackTraceFileLocation); i {
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
		file_temporal_api_sdk_v1_enhanced_stack_trace_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*StackTrace); i {
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
			RawDescriptor: file_temporal_api_sdk_v1_enhanced_stack_trace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_sdk_v1_enhanced_stack_trace_proto_goTypes,
		DependencyIndexes: file_temporal_api_sdk_v1_enhanced_stack_trace_proto_depIdxs,
		MessageInfos:      file_temporal_api_sdk_v1_enhanced_stack_trace_proto_msgTypes,
	}.Build()
	File_temporal_api_sdk_v1_enhanced_stack_trace_proto = out.File
	file_temporal_api_sdk_v1_enhanced_stack_trace_proto_rawDesc = nil
	file_temporal_api_sdk_v1_enhanced_stack_trace_proto_goTypes = nil
	file_temporal_api_sdk_v1_enhanced_stack_trace_proto_depIdxs = nil
}
