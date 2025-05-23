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
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: temporal/api/update/v1/message.proto

package update

import (
	v11 "go.temporal.io/api/common/v1"
	v1 "go.temporal.io/api/enums/v1"
	v12 "go.temporal.io/api/failure/v1"
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

// Specifies client's intent to wait for Update results.
type WaitPolicy struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Indicates the Update lifecycle stage that the Update must reach before
	// API call is returned.
	// NOTE: This field works together with API call timeout which is limited by
	// server timeout (maximum wait time). If server timeout is expired before
	// user specified timeout, API call returns even if specified stage is not reached.
	LifecycleStage v1.UpdateWorkflowExecutionLifecycleStage `protobuf:"varint,1,opt,name=lifecycle_stage,json=lifecycleStage,proto3,enum=temporal.api.enums.v1.UpdateWorkflowExecutionLifecycleStage" json:"lifecycle_stage,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *WaitPolicy) Reset() {
	*x = WaitPolicy{}
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WaitPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaitPolicy) ProtoMessage() {}

func (x *WaitPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaitPolicy.ProtoReflect.Descriptor instead.
func (*WaitPolicy) Descriptor() ([]byte, []int) {
	return file_temporal_api_update_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *WaitPolicy) GetLifecycleStage() v1.UpdateWorkflowExecutionLifecycleStage {
	if x != nil {
		return x.LifecycleStage
	}
	return v1.UpdateWorkflowExecutionLifecycleStage(0)
}

// The data needed by a client to refer to a previously invoked Workflow Update.
type UpdateRef struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	WorkflowExecution *v11.WorkflowExecution `protobuf:"bytes,1,opt,name=workflow_execution,json=workflowExecution,proto3" json:"workflow_execution,omitempty"`
	UpdateId          string                 `protobuf:"bytes,2,opt,name=update_id,json=updateId,proto3" json:"update_id,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *UpdateRef) Reset() {
	*x = UpdateRef{}
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRef) ProtoMessage() {}

func (x *UpdateRef) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRef.ProtoReflect.Descriptor instead.
func (*UpdateRef) Descriptor() ([]byte, []int) {
	return file_temporal_api_update_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRef) GetWorkflowExecution() *v11.WorkflowExecution {
	if x != nil {
		return x.WorkflowExecution
	}
	return nil
}

func (x *UpdateRef) GetUpdateId() string {
	if x != nil {
		return x.UpdateId
	}
	return ""
}

// The outcome of a Workflow Update: success or failure.
type Outcome struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Value:
	//
	//	*Outcome_Success
	//	*Outcome_Failure
	Value         isOutcome_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Outcome) Reset() {
	*x = Outcome{}
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Outcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outcome) ProtoMessage() {}

func (x *Outcome) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outcome.ProtoReflect.Descriptor instead.
func (*Outcome) Descriptor() ([]byte, []int) {
	return file_temporal_api_update_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *Outcome) GetValue() isOutcome_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Outcome) GetSuccess() *v11.Payloads {
	if x != nil {
		if x, ok := x.Value.(*Outcome_Success); ok {
			return x.Success
		}
	}
	return nil
}

func (x *Outcome) GetFailure() *v12.Failure {
	if x != nil {
		if x, ok := x.Value.(*Outcome_Failure); ok {
			return x.Failure
		}
	}
	return nil
}

type isOutcome_Value interface {
	isOutcome_Value()
}

type Outcome_Success struct {
	Success *v11.Payloads `protobuf:"bytes,1,opt,name=success,proto3,oneof"`
}

type Outcome_Failure struct {
	Failure *v12.Failure `protobuf:"bytes,2,opt,name=failure,proto3,oneof"`
}

func (*Outcome_Success) isOutcome_Value() {}

func (*Outcome_Failure) isOutcome_Value() {}

// Metadata about a Workflow Update.
type Meta struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// An ID with workflow-scoped uniqueness for this Update.
	UpdateId string `protobuf:"bytes,1,opt,name=update_id,json=updateId,proto3" json:"update_id,omitempty"`
	// A string identifying the agent that requested this Update.
	Identity      string `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Meta) Reset() {
	*x = Meta{}
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_temporal_api_update_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *Meta) GetUpdateId() string {
	if x != nil {
		return x.UpdateId
	}
	return ""
}

func (x *Meta) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

type Input struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Headers that are passed with the Update from the requesting entity.
	// These can include things like auth or tracing tokens.
	Header *v11.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The name of the Update handler to invoke on the target Workflow.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The arguments to pass to the named Update handler.
	Args          *v11.Payloads `protobuf:"bytes,3,opt,name=args,proto3" json:"args,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Input) Reset() {
	*x = Input{}
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_temporal_api_update_v1_message_proto_rawDescGZIP(), []int{4}
}

func (x *Input) GetHeader() *v11.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Input) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Input) GetArgs() *v11.Payloads {
	if x != nil {
		return x.Args
	}
	return nil
}

// The client request that triggers a Workflow Update.
type Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Meta          *Meta                  `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Input         *Input                 `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[5]
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
	return file_temporal_api_update_v1_message_proto_rawDescGZIP(), []int{5}
}

func (x *Request) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Request) GetInput() *Input {
	if x != nil {
		return x.Input
	}
	return nil
}

// An Update protocol message indicating that a Workflow Update has been rejected.
type Rejection struct {
	state                            protoimpl.MessageState `protogen:"open.v1"`
	RejectedRequestMessageId         string                 `protobuf:"bytes,1,opt,name=rejected_request_message_id,json=rejectedRequestMessageId,proto3" json:"rejected_request_message_id,omitempty"`
	RejectedRequestSequencingEventId int64                  `protobuf:"varint,2,opt,name=rejected_request_sequencing_event_id,json=rejectedRequestSequencingEventId,proto3" json:"rejected_request_sequencing_event_id,omitempty"`
	RejectedRequest                  *Request               `protobuf:"bytes,3,opt,name=rejected_request,json=rejectedRequest,proto3" json:"rejected_request,omitempty"`
	Failure                          *v12.Failure           `protobuf:"bytes,4,opt,name=failure,proto3" json:"failure,omitempty"`
	unknownFields                    protoimpl.UnknownFields
	sizeCache                        protoimpl.SizeCache
}

func (x *Rejection) Reset() {
	*x = Rejection{}
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Rejection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rejection) ProtoMessage() {}

func (x *Rejection) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rejection.ProtoReflect.Descriptor instead.
func (*Rejection) Descriptor() ([]byte, []int) {
	return file_temporal_api_update_v1_message_proto_rawDescGZIP(), []int{6}
}

func (x *Rejection) GetRejectedRequestMessageId() string {
	if x != nil {
		return x.RejectedRequestMessageId
	}
	return ""
}

func (x *Rejection) GetRejectedRequestSequencingEventId() int64 {
	if x != nil {
		return x.RejectedRequestSequencingEventId
	}
	return 0
}

func (x *Rejection) GetRejectedRequest() *Request {
	if x != nil {
		return x.RejectedRequest
	}
	return nil
}

func (x *Rejection) GetFailure() *v12.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

// An Update protocol message indicating that a Workflow Update has
// been accepted (i.e. passed the worker-side validation phase).
type Acceptance struct {
	state                            protoimpl.MessageState `protogen:"open.v1"`
	AcceptedRequestMessageId         string                 `protobuf:"bytes,1,opt,name=accepted_request_message_id,json=acceptedRequestMessageId,proto3" json:"accepted_request_message_id,omitempty"`
	AcceptedRequestSequencingEventId int64                  `protobuf:"varint,2,opt,name=accepted_request_sequencing_event_id,json=acceptedRequestSequencingEventId,proto3" json:"accepted_request_sequencing_event_id,omitempty"`
	AcceptedRequest                  *Request               `protobuf:"bytes,3,opt,name=accepted_request,json=acceptedRequest,proto3" json:"accepted_request,omitempty"`
	unknownFields                    protoimpl.UnknownFields
	sizeCache                        protoimpl.SizeCache
}

func (x *Acceptance) Reset() {
	*x = Acceptance{}
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Acceptance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Acceptance) ProtoMessage() {}

func (x *Acceptance) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Acceptance.ProtoReflect.Descriptor instead.
func (*Acceptance) Descriptor() ([]byte, []int) {
	return file_temporal_api_update_v1_message_proto_rawDescGZIP(), []int{7}
}

func (x *Acceptance) GetAcceptedRequestMessageId() string {
	if x != nil {
		return x.AcceptedRequestMessageId
	}
	return ""
}

func (x *Acceptance) GetAcceptedRequestSequencingEventId() int64 {
	if x != nil {
		return x.AcceptedRequestSequencingEventId
	}
	return 0
}

func (x *Acceptance) GetAcceptedRequest() *Request {
	if x != nil {
		return x.AcceptedRequest
	}
	return nil
}

// An Update protocol message indicating that a Workflow Update has
// completed with the contained outcome.
type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Meta          *Meta                  `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Outcome       *Outcome               `protobuf:"bytes,2,opt,name=outcome,proto3" json:"outcome,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[8]
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
	return file_temporal_api_update_v1_message_proto_rawDescGZIP(), []int{8}
}

func (x *Response) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Response) GetOutcome() *Outcome {
	if x != nil {
		return x.Outcome
	}
	return nil
}

var File_temporal_api_update_v1_message_proto protoreflect.FileDescriptor

const file_temporal_api_update_v1_message_proto_rawDesc = "" +
	"\n" +
	"$temporal/api/update/v1/message.proto\x12\x16temporal.api.update.v1\x1a$temporal/api/common/v1/message.proto\x1a\"temporal/api/enums/v1/update.proto\x1a%temporal/api/failure/v1/message.proto\"s\n" +
	"\n" +
	"WaitPolicy\x12e\n" +
	"\x0flifecycle_stage\x18\x01 \x01(\x0e2<.temporal.api.enums.v1.UpdateWorkflowExecutionLifecycleStageR\x0elifecycleStage\"\x82\x01\n" +
	"\tUpdateRef\x12X\n" +
	"\x12workflow_execution\x18\x01 \x01(\v2).temporal.api.common.v1.WorkflowExecutionR\x11workflowExecution\x12\x1b\n" +
	"\tupdate_id\x18\x02 \x01(\tR\bupdateId\"\x8e\x01\n" +
	"\aOutcome\x12<\n" +
	"\asuccess\x18\x01 \x01(\v2 .temporal.api.common.v1.PayloadsH\x00R\asuccess\x12<\n" +
	"\afailure\x18\x02 \x01(\v2 .temporal.api.failure.v1.FailureH\x00R\afailureB\a\n" +
	"\x05value\"?\n" +
	"\x04Meta\x12\x1b\n" +
	"\tupdate_id\x18\x01 \x01(\tR\bupdateId\x12\x1a\n" +
	"\bidentity\x18\x02 \x01(\tR\bidentity\"\x89\x01\n" +
	"\x05Input\x126\n" +
	"\x06header\x18\x01 \x01(\v2\x1e.temporal.api.common.v1.HeaderR\x06header\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x124\n" +
	"\x04args\x18\x03 \x01(\v2 .temporal.api.common.v1.PayloadsR\x04args\"p\n" +
	"\aRequest\x120\n" +
	"\x04meta\x18\x01 \x01(\v2\x1c.temporal.api.update.v1.MetaR\x04meta\x123\n" +
	"\x05input\x18\x02 \x01(\v2\x1d.temporal.api.update.v1.InputR\x05input\"\xa2\x02\n" +
	"\tRejection\x12=\n" +
	"\x1brejected_request_message_id\x18\x01 \x01(\tR\x18rejectedRequestMessageId\x12N\n" +
	"$rejected_request_sequencing_event_id\x18\x02 \x01(\x03R rejectedRequestSequencingEventId\x12J\n" +
	"\x10rejected_request\x18\x03 \x01(\v2\x1f.temporal.api.update.v1.RequestR\x0frejectedRequest\x12:\n" +
	"\afailure\x18\x04 \x01(\v2 .temporal.api.failure.v1.FailureR\afailure\"\xe7\x01\n" +
	"\n" +
	"Acceptance\x12=\n" +
	"\x1baccepted_request_message_id\x18\x01 \x01(\tR\x18acceptedRequestMessageId\x12N\n" +
	"$accepted_request_sequencing_event_id\x18\x02 \x01(\x03R acceptedRequestSequencingEventId\x12J\n" +
	"\x10accepted_request\x18\x03 \x01(\v2\x1f.temporal.api.update.v1.RequestR\x0facceptedRequest\"w\n" +
	"\bResponse\x120\n" +
	"\x04meta\x18\x01 \x01(\v2\x1c.temporal.api.update.v1.MetaR\x04meta\x129\n" +
	"\aoutcome\x18\x02 \x01(\v2\x1f.temporal.api.update.v1.OutcomeR\aoutcomeB\x89\x01\n" +
	"\x19io.temporal.api.update.v1B\fMessageProtoP\x01Z#go.temporal.io/api/update/v1;update\xaa\x02\x18Temporalio.Api.Update.V1\xea\x02\x1bTemporalio::Api::Update::V1b\x06proto3"

var (
	file_temporal_api_update_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_update_v1_message_proto_rawDescData []byte
)

func file_temporal_api_update_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_update_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_update_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_temporal_api_update_v1_message_proto_rawDesc), len(file_temporal_api_update_v1_message_proto_rawDesc)))
	})
	return file_temporal_api_update_v1_message_proto_rawDescData
}

var file_temporal_api_update_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_temporal_api_update_v1_message_proto_goTypes = []any{
	(*WaitPolicy)(nil), // 0: temporal.api.update.v1.WaitPolicy
	(*UpdateRef)(nil),  // 1: temporal.api.update.v1.UpdateRef
	(*Outcome)(nil),    // 2: temporal.api.update.v1.Outcome
	(*Meta)(nil),       // 3: temporal.api.update.v1.Meta
	(*Input)(nil),      // 4: temporal.api.update.v1.Input
	(*Request)(nil),    // 5: temporal.api.update.v1.Request
	(*Rejection)(nil),  // 6: temporal.api.update.v1.Rejection
	(*Acceptance)(nil), // 7: temporal.api.update.v1.Acceptance
	(*Response)(nil),   // 8: temporal.api.update.v1.Response
	(v1.UpdateWorkflowExecutionLifecycleStage)(0), // 9: temporal.api.enums.v1.UpdateWorkflowExecutionLifecycleStage
	(*v11.WorkflowExecution)(nil),                 // 10: temporal.api.common.v1.WorkflowExecution
	(*v11.Payloads)(nil),                          // 11: temporal.api.common.v1.Payloads
	(*v12.Failure)(nil),                           // 12: temporal.api.failure.v1.Failure
	(*v11.Header)(nil),                            // 13: temporal.api.common.v1.Header
}
var file_temporal_api_update_v1_message_proto_depIdxs = []int32{
	9,  // 0: temporal.api.update.v1.WaitPolicy.lifecycle_stage:type_name -> temporal.api.enums.v1.UpdateWorkflowExecutionLifecycleStage
	10, // 1: temporal.api.update.v1.UpdateRef.workflow_execution:type_name -> temporal.api.common.v1.WorkflowExecution
	11, // 2: temporal.api.update.v1.Outcome.success:type_name -> temporal.api.common.v1.Payloads
	12, // 3: temporal.api.update.v1.Outcome.failure:type_name -> temporal.api.failure.v1.Failure
	13, // 4: temporal.api.update.v1.Input.header:type_name -> temporal.api.common.v1.Header
	11, // 5: temporal.api.update.v1.Input.args:type_name -> temporal.api.common.v1.Payloads
	3,  // 6: temporal.api.update.v1.Request.meta:type_name -> temporal.api.update.v1.Meta
	4,  // 7: temporal.api.update.v1.Request.input:type_name -> temporal.api.update.v1.Input
	5,  // 8: temporal.api.update.v1.Rejection.rejected_request:type_name -> temporal.api.update.v1.Request
	12, // 9: temporal.api.update.v1.Rejection.failure:type_name -> temporal.api.failure.v1.Failure
	5,  // 10: temporal.api.update.v1.Acceptance.accepted_request:type_name -> temporal.api.update.v1.Request
	3,  // 11: temporal.api.update.v1.Response.meta:type_name -> temporal.api.update.v1.Meta
	2,  // 12: temporal.api.update.v1.Response.outcome:type_name -> temporal.api.update.v1.Outcome
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_temporal_api_update_v1_message_proto_init() }
func file_temporal_api_update_v1_message_proto_init() {
	if File_temporal_api_update_v1_message_proto != nil {
		return
	}
	file_temporal_api_update_v1_message_proto_msgTypes[2].OneofWrappers = []any{
		(*Outcome_Success)(nil),
		(*Outcome_Failure)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_temporal_api_update_v1_message_proto_rawDesc), len(file_temporal_api_update_v1_message_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_update_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_update_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_api_update_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_update_v1_message_proto = out.File
	file_temporal_api_update_v1_message_proto_goTypes = nil
	file_temporal_api_update_v1_message_proto_depIdxs = nil
}
