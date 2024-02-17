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
// source: temporal/api/enums/v1/event_type.proto

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

// Whenever this list of events is changed do change the function shouldBufferEvent in mutableStateBuilder.go to make sure to do the correct event ordering
type EventType int32

const (
	// Place holder and should never appear in a Workflow execution history
	EventType_EVENT_TYPE_UNSPECIFIED EventType = 0
	// Workflow execution has been triggered/started
	// It contains Workflow execution inputs, as well as Workflow timeout configurations
	EventType_EVENT_TYPE_WORKFLOW_EXECUTION_STARTED EventType = 1
	// Workflow execution has successfully completed and contains Workflow execution results
	EventType_EVENT_TYPE_WORKFLOW_EXECUTION_COMPLETED EventType = 2
	// Workflow execution has unsuccessfully completed and contains the Workflow execution error
	EventType_EVENT_TYPE_WORKFLOW_EXECUTION_FAILED EventType = 3
	// Workflow execution has timed out by the Temporal Server
	// Usually due to the Workflow having not been completed within timeout settings
	EventType_EVENT_TYPE_WORKFLOW_EXECUTION_TIMED_OUT EventType = 4
	// Workflow Task has been scheduled and the SDK client should now be able to process any new history events
	EventType_EVENT_TYPE_WORKFLOW_TASK_SCHEDULED EventType = 5
	// Workflow Task has started and the SDK client has picked up the Workflow Task and is processing new history events
	EventType_EVENT_TYPE_WORKFLOW_TASK_STARTED EventType = 6
	// Workflow Task has completed
	// The SDK client picked up the Workflow Task and processed new history events
	// SDK client may or may not ask the Temporal Server to do additional work, such as:
	// EVENT_TYPE_ACTIVITY_TASK_SCHEDULED
	// EVENT_TYPE_TIMER_STARTED
	// EVENT_TYPE_UPSERT_WORKFLOW_SEARCH_ATTRIBUTES
	// EVENT_TYPE_MARKER_RECORDED
	// EVENT_TYPE_START_CHILD_WORKFLOW_EXECUTION_INITIATED
	// EVENT_TYPE_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_INITIATED
	// EVENT_TYPE_SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_INITIATED
	// EVENT_TYPE_WORKFLOW_EXECUTION_COMPLETED
	// EVENT_TYPE_WORKFLOW_EXECUTION_FAILED
	// EVENT_TYPE_WORKFLOW_EXECUTION_CANCELED
	// EVENT_TYPE_WORKFLOW_EXECUTION_CONTINUED_AS_NEW
	EventType_EVENT_TYPE_WORKFLOW_TASK_COMPLETED EventType = 7
	// Workflow Task encountered a timeout
	// Either an SDK client with a local cache was not available at the time, or it took too long for the SDK client to process the task
	EventType_EVENT_TYPE_WORKFLOW_TASK_TIMED_OUT EventType = 8
	// Workflow Task encountered a failure
	// Usually this means that the Workflow was non-deterministic
	// However, the Workflow reset functionality also uses this event
	EventType_EVENT_TYPE_WORKFLOW_TASK_FAILED EventType = 9
	// Activity Task was scheduled
	// The SDK client should pick up this activity task and execute
	// This event type contains activity inputs, as well as activity timeout configurations
	EventType_EVENT_TYPE_ACTIVITY_TASK_SCHEDULED EventType = 10
	// Activity Task has started executing
	// The SDK client has picked up the Activity Task and is processing the Activity invocation
	EventType_EVENT_TYPE_ACTIVITY_TASK_STARTED EventType = 11
	// Activity Task has finished successfully
	// The SDK client has picked up and successfully completed the Activity Task
	// This event type contains Activity execution results
	EventType_EVENT_TYPE_ACTIVITY_TASK_COMPLETED EventType = 12
	// Activity Task has finished unsuccessfully
	// The SDK picked up the Activity Task but unsuccessfully completed it
	// This event type contains Activity execution errors
	EventType_EVENT_TYPE_ACTIVITY_TASK_FAILED EventType = 13
	// Activity has timed out according to the Temporal Server
	// Activity did not complete within the timeout settings
	EventType_EVENT_TYPE_ACTIVITY_TASK_TIMED_OUT EventType = 14
	// A request to cancel the Activity has occurred
	// The SDK client will be able to confirm cancellation of an Activity during an Activity heartbeat
	EventType_EVENT_TYPE_ACTIVITY_TASK_CANCEL_REQUESTED EventType = 15
	// Activity has been cancelled
	EventType_EVENT_TYPE_ACTIVITY_TASK_CANCELED EventType = 16
	// A timer has started
	EventType_EVENT_TYPE_TIMER_STARTED EventType = 17
	// A timer has fired
	EventType_EVENT_TYPE_TIMER_FIRED EventType = 18
	// A time has been cancelled
	EventType_EVENT_TYPE_TIMER_CANCELED EventType = 19
	// A request has been made to cancel the Workflow execution
	EventType_EVENT_TYPE_WORKFLOW_EXECUTION_CANCEL_REQUESTED EventType = 20
	// SDK client has confirmed the cancellation request and the Workflow execution has been cancelled
	EventType_EVENT_TYPE_WORKFLOW_EXECUTION_CANCELED EventType = 21
	// Workflow has requested that the Temporal Server try to cancel another Workflow
	EventType_EVENT_TYPE_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_INITIATED EventType = 22
	// Temporal Server could not cancel the targeted Workflow
	// This is usually because the target Workflow could not be found
	EventType_EVENT_TYPE_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED EventType = 23
	// Temporal Server has successfully requested the cancellation of the target Workflow
	EventType_EVENT_TYPE_EXTERNAL_WORKFLOW_EXECUTION_CANCEL_REQUESTED EventType = 24
	// A marker has been recorded.
	// This event type is transparent to the Temporal Server
	// The Server will only store it and will not try to understand it.
	EventType_EVENT_TYPE_MARKER_RECORDED EventType = 25
	// Workflow has received a Signal event
	// The event type contains the Signal name, as well as a Signal payload
	EventType_EVENT_TYPE_WORKFLOW_EXECUTION_SIGNALED EventType = 26
	// Workflow execution has been forcefully terminated
	// This is usually because the terminate Workflow API was called
	EventType_EVENT_TYPE_WORKFLOW_EXECUTION_TERMINATED EventType = 27
	// Workflow has successfully completed and a new Workflow has been started within the same transaction
	// Contains last Workflow execution results as well as new Workflow execution inputs
	EventType_EVENT_TYPE_WORKFLOW_EXECUTION_CONTINUED_AS_NEW EventType = 28
	// Temporal Server will try to start a child Workflow
	EventType_EVENT_TYPE_START_CHILD_WORKFLOW_EXECUTION_INITIATED EventType = 29
	// Child Workflow execution cannot be started/triggered
	// Usually due to a child Workflow ID collision
	EventType_EVENT_TYPE_START_CHILD_WORKFLOW_EXECUTION_FAILED EventType = 30
	// Child Workflow execution has successfully started/triggered
	EventType_EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_STARTED EventType = 31
	// Child Workflow execution has successfully completed
	EventType_EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_COMPLETED EventType = 32
	// Child Workflow execution has unsuccessfully completed
	EventType_EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_FAILED EventType = 33
	// Child Workflow execution has been cancelled
	EventType_EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_CANCELED EventType = 34
	// Child Workflow execution has timed out by the Temporal Server
	EventType_EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_TIMED_OUT EventType = 35
	// Child Workflow execution has been terminated
	EventType_EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_TERMINATED EventType = 36
	// Temporal Server will try to Signal the targeted Workflow
	// Contains the Signal name, as well as a Signal payload
	EventType_EVENT_TYPE_SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_INITIATED EventType = 37
	// Temporal Server cannot Signal the targeted Workflow
	// Usually because the Workflow could not be found
	EventType_EVENT_TYPE_SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED EventType = 38
	// Temporal Server has successfully Signaled the targeted Workflow
	EventType_EVENT_TYPE_EXTERNAL_WORKFLOW_EXECUTION_SIGNALED EventType = 39
	// Workflow search attributes should be updated and synchronized with the visibility store
	EventType_EVENT_TYPE_UPSERT_WORKFLOW_SEARCH_ATTRIBUTES EventType = 40
	// An update was accepted (i.e. validated)
	EventType_EVENT_TYPE_WORKFLOW_EXECUTION_UPDATE_ACCEPTED EventType = 41
	// An update was rejected (i.e. failed validation)
	EventType_EVENT_TYPE_WORKFLOW_EXECUTION_UPDATE_REJECTED EventType = 42
	// An update completed
	EventType_EVENT_TYPE_WORKFLOW_EXECUTION_UPDATE_COMPLETED EventType = 43
	// Some property or properties of the workflow as a whole have changed by non-workflow code.
	// The distinction of external vs. command-based modification is important so the SDK can
	// maintain determinism when using the command-based approach.
	EventType_EVENT_TYPE_WORKFLOW_PROPERTIES_MODIFIED_EXTERNALLY EventType = 44
	// Some property or properties of an already-scheduled activity have changed by non-workflow code.
	// The distinction of external vs. command-based modification is important so the SDK can
	// maintain determinism when using the command-based approach.
	EventType_EVENT_TYPE_ACTIVITY_PROPERTIES_MODIFIED_EXTERNALLY EventType = 45
	// Workflow properties modified by user workflow code
	EventType_EVENT_TYPE_WORKFLOW_PROPERTIES_MODIFIED EventType = 46
	// An update was requested. Note that not all update requests result in this
	// event. See UpdateRequestedEventOrigin for situations in which this event
	// is created.
	EventType_EVENT_TYPE_WORKFLOW_EXECUTION_UPDATE_REQUESTED EventType = 47
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0:  "EVENT_TYPE_UNSPECIFIED",
		1:  "EVENT_TYPE_WORKFLOW_EXECUTION_STARTED",
		2:  "EVENT_TYPE_WORKFLOW_EXECUTION_COMPLETED",
		3:  "EVENT_TYPE_WORKFLOW_EXECUTION_FAILED",
		4:  "EVENT_TYPE_WORKFLOW_EXECUTION_TIMED_OUT",
		5:  "EVENT_TYPE_WORKFLOW_TASK_SCHEDULED",
		6:  "EVENT_TYPE_WORKFLOW_TASK_STARTED",
		7:  "EVENT_TYPE_WORKFLOW_TASK_COMPLETED",
		8:  "EVENT_TYPE_WORKFLOW_TASK_TIMED_OUT",
		9:  "EVENT_TYPE_WORKFLOW_TASK_FAILED",
		10: "EVENT_TYPE_ACTIVITY_TASK_SCHEDULED",
		11: "EVENT_TYPE_ACTIVITY_TASK_STARTED",
		12: "EVENT_TYPE_ACTIVITY_TASK_COMPLETED",
		13: "EVENT_TYPE_ACTIVITY_TASK_FAILED",
		14: "EVENT_TYPE_ACTIVITY_TASK_TIMED_OUT",
		15: "EVENT_TYPE_ACTIVITY_TASK_CANCEL_REQUESTED",
		16: "EVENT_TYPE_ACTIVITY_TASK_CANCELED",
		17: "EVENT_TYPE_TIMER_STARTED",
		18: "EVENT_TYPE_TIMER_FIRED",
		19: "EVENT_TYPE_TIMER_CANCELED",
		20: "EVENT_TYPE_WORKFLOW_EXECUTION_CANCEL_REQUESTED",
		21: "EVENT_TYPE_WORKFLOW_EXECUTION_CANCELED",
		22: "EVENT_TYPE_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_INITIATED",
		23: "EVENT_TYPE_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED",
		24: "EVENT_TYPE_EXTERNAL_WORKFLOW_EXECUTION_CANCEL_REQUESTED",
		25: "EVENT_TYPE_MARKER_RECORDED",
		26: "EVENT_TYPE_WORKFLOW_EXECUTION_SIGNALED",
		27: "EVENT_TYPE_WORKFLOW_EXECUTION_TERMINATED",
		28: "EVENT_TYPE_WORKFLOW_EXECUTION_CONTINUED_AS_NEW",
		29: "EVENT_TYPE_START_CHILD_WORKFLOW_EXECUTION_INITIATED",
		30: "EVENT_TYPE_START_CHILD_WORKFLOW_EXECUTION_FAILED",
		31: "EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_STARTED",
		32: "EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_COMPLETED",
		33: "EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_FAILED",
		34: "EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_CANCELED",
		35: "EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_TIMED_OUT",
		36: "EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_TERMINATED",
		37: "EVENT_TYPE_SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_INITIATED",
		38: "EVENT_TYPE_SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED",
		39: "EVENT_TYPE_EXTERNAL_WORKFLOW_EXECUTION_SIGNALED",
		40: "EVENT_TYPE_UPSERT_WORKFLOW_SEARCH_ATTRIBUTES",
		41: "EVENT_TYPE_WORKFLOW_EXECUTION_UPDATE_ACCEPTED",
		42: "EVENT_TYPE_WORKFLOW_EXECUTION_UPDATE_REJECTED",
		43: "EVENT_TYPE_WORKFLOW_EXECUTION_UPDATE_COMPLETED",
		44: "EVENT_TYPE_WORKFLOW_PROPERTIES_MODIFIED_EXTERNALLY",
		45: "EVENT_TYPE_ACTIVITY_PROPERTIES_MODIFIED_EXTERNALLY",
		46: "EVENT_TYPE_WORKFLOW_PROPERTIES_MODIFIED",
		47: "EVENT_TYPE_WORKFLOW_EXECUTION_UPDATE_REQUESTED",
	}
	EventType_value = map[string]int32{
		"EVENT_TYPE_UNSPECIFIED":                                          0,
		"EVENT_TYPE_WORKFLOW_EXECUTION_STARTED":                           1,
		"EVENT_TYPE_WORKFLOW_EXECUTION_COMPLETED":                         2,
		"EVENT_TYPE_WORKFLOW_EXECUTION_FAILED":                            3,
		"EVENT_TYPE_WORKFLOW_EXECUTION_TIMED_OUT":                         4,
		"EVENT_TYPE_WORKFLOW_TASK_SCHEDULED":                              5,
		"EVENT_TYPE_WORKFLOW_TASK_STARTED":                                6,
		"EVENT_TYPE_WORKFLOW_TASK_COMPLETED":                              7,
		"EVENT_TYPE_WORKFLOW_TASK_TIMED_OUT":                              8,
		"EVENT_TYPE_WORKFLOW_TASK_FAILED":                                 9,
		"EVENT_TYPE_ACTIVITY_TASK_SCHEDULED":                              10,
		"EVENT_TYPE_ACTIVITY_TASK_STARTED":                                11,
		"EVENT_TYPE_ACTIVITY_TASK_COMPLETED":                              12,
		"EVENT_TYPE_ACTIVITY_TASK_FAILED":                                 13,
		"EVENT_TYPE_ACTIVITY_TASK_TIMED_OUT":                              14,
		"EVENT_TYPE_ACTIVITY_TASK_CANCEL_REQUESTED":                       15,
		"EVENT_TYPE_ACTIVITY_TASK_CANCELED":                               16,
		"EVENT_TYPE_TIMER_STARTED":                                        17,
		"EVENT_TYPE_TIMER_FIRED":                                          18,
		"EVENT_TYPE_TIMER_CANCELED":                                       19,
		"EVENT_TYPE_WORKFLOW_EXECUTION_CANCEL_REQUESTED":                  20,
		"EVENT_TYPE_WORKFLOW_EXECUTION_CANCELED":                          21,
		"EVENT_TYPE_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_INITIATED": 22,
		"EVENT_TYPE_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED":    23,
		"EVENT_TYPE_EXTERNAL_WORKFLOW_EXECUTION_CANCEL_REQUESTED":         24,
		"EVENT_TYPE_MARKER_RECORDED":                                      25,
		"EVENT_TYPE_WORKFLOW_EXECUTION_SIGNALED":                          26,
		"EVENT_TYPE_WORKFLOW_EXECUTION_TERMINATED":                        27,
		"EVENT_TYPE_WORKFLOW_EXECUTION_CONTINUED_AS_NEW":                  28,
		"EVENT_TYPE_START_CHILD_WORKFLOW_EXECUTION_INITIATED":             29,
		"EVENT_TYPE_START_CHILD_WORKFLOW_EXECUTION_FAILED":                30,
		"EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_STARTED":                     31,
		"EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_COMPLETED":                   32,
		"EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_FAILED":                      33,
		"EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_CANCELED":                    34,
		"EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_TIMED_OUT":                   35,
		"EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_TERMINATED":                  36,
		"EVENT_TYPE_SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_INITIATED":         37,
		"EVENT_TYPE_SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED":            38,
		"EVENT_TYPE_EXTERNAL_WORKFLOW_EXECUTION_SIGNALED":                 39,
		"EVENT_TYPE_UPSERT_WORKFLOW_SEARCH_ATTRIBUTES":                    40,
		"EVENT_TYPE_WORKFLOW_EXECUTION_UPDATE_ACCEPTED":                   41,
		"EVENT_TYPE_WORKFLOW_EXECUTION_UPDATE_REJECTED":                   42,
		"EVENT_TYPE_WORKFLOW_EXECUTION_UPDATE_COMPLETED":                  43,
		"EVENT_TYPE_WORKFLOW_PROPERTIES_MODIFIED_EXTERNALLY":              44,
		"EVENT_TYPE_ACTIVITY_PROPERTIES_MODIFIED_EXTERNALLY":              45,
		"EVENT_TYPE_WORKFLOW_PROPERTIES_MODIFIED":                         46,
		"EVENT_TYPE_WORKFLOW_EXECUTION_UPDATE_REQUESTED":                  47,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_event_type_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_event_type_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_event_type_proto_rawDescGZIP(), []int{0}
}

var File_temporal_api_enums_v1_event_type_proto protoreflect.FileDescriptor

var file_temporal_api_enums_v1_event_type_proto_rawDesc = []byte{
	0x0a, 0x26, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2a,
	0xc2, 0x11, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x16, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x29, 0x0a, 0x25, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57,
	0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x2b, 0x0a, 0x27, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43,
	0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x28, 0x0a, 0x24, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x2b, 0x0a, 0x27, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c,
	0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x44, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x04, 0x12, 0x26, 0x0a, 0x22, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f,
	0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0x05,
	0x12, 0x24, 0x0a, 0x20, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57,
	0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53, 0x54, 0x41,
	0x52, 0x54, 0x45, 0x44, 0x10, 0x06, 0x12, 0x26, 0x0a, 0x22, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x54, 0x41,
	0x53, 0x4b, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x07, 0x12, 0x26,
	0x0a, 0x22, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52,
	0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x44,
	0x5f, 0x4f, 0x55, 0x54, 0x10, 0x08, 0x12, 0x23, 0x0a, 0x1f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x54, 0x41,
	0x53, 0x4b, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x09, 0x12, 0x26, 0x0a, 0x22, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49,
	0x54, 0x59, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45,
	0x44, 0x10, 0x0a, 0x12, 0x24, 0x0a, 0x20, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f,
	0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x26, 0x0a, 0x22, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59,
	0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10,
	0x0c, 0x12, 0x23, 0x0a, 0x1f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x0d, 0x12, 0x26, 0x0a, 0x22, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x41,
	0x53, 0x4b, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x44, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x0e, 0x12, 0x2d,
	0x0a, 0x29, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45,
	0x4c, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10, 0x0f, 0x12, 0x25, 0x0a,
	0x21, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c,
	0x45, 0x44, 0x10, 0x10, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44,
	0x10, 0x11, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x54, 0x49, 0x4d, 0x45, 0x52, 0x5f, 0x46, 0x49, 0x52, 0x45, 0x44, 0x10, 0x12, 0x12, 0x1d,
	0x0a, 0x19, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x52, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x13, 0x12, 0x32, 0x0a,
	0x2e, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b,
	0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10,
	0x14, 0x12, 0x2a, 0x0a, 0x26, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x15, 0x12, 0x43, 0x0a,
	0x3f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x52,
	0x4e, 0x41, 0x4c, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45,
	0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x54, 0x45, 0x44,
	0x10, 0x16, 0x12, 0x40, 0x0a, 0x3c, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f,
	0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f,
	0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x17, 0x12, 0x3b, 0x0a, 0x37, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x57, 0x4f, 0x52, 0x4b,
	0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10,
	0x18, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4d, 0x41, 0x52, 0x4b, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x45, 0x44, 0x10,
	0x19, 0x12, 0x2a, 0x0a, 0x26, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x4c, 0x45, 0x44, 0x10, 0x1a, 0x12, 0x2c, 0x0a,
	0x28, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b,
	0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x44, 0x10, 0x1b, 0x12, 0x32, 0x0a, 0x2e, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c,
	0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4e,
	0x54, 0x49, 0x4e, 0x55, 0x45, 0x44, 0x5f, 0x41, 0x53, 0x5f, 0x4e, 0x45, 0x57, 0x10, 0x1c, 0x12,
	0x37, 0x0a, 0x33, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x5f, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c,
	0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x49,
	0x54, 0x49, 0x41, 0x54, 0x45, 0x44, 0x10, 0x1d, 0x12, 0x34, 0x0a, 0x30, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x43, 0x48, 0x49,
	0x4c, 0x44, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43,
	0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x1e, 0x12, 0x2f,
	0x0a, 0x2b, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48, 0x49,
	0x4c, 0x44, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43,
	0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x1f, 0x12,
	0x31, 0x0a, 0x2d, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48,
	0x49, 0x4c, 0x44, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45,
	0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44,
	0x10, 0x20, 0x12, 0x2e, 0x0a, 0x2a, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f,
	0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x21, 0x12, 0x30, 0x0a, 0x2c, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f,
	0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c,
	0x45, 0x44, 0x10, 0x22, 0x12, 0x31, 0x0a, 0x2d, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f,
	0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x49, 0x4d, 0x45,
	0x44, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x23, 0x12, 0x32, 0x0a, 0x2e, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f, 0x57, 0x4f, 0x52, 0x4b,
	0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x44, 0x10, 0x24, 0x12, 0x3b, 0x0a, 0x37, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x4c,
	0x5f, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c,
	0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x49,
	0x54, 0x49, 0x41, 0x54, 0x45, 0x44, 0x10, 0x25, 0x12, 0x38, 0x0a, 0x34, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x58,
	0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f,
	0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x26, 0x12, 0x33, 0x0a, 0x2f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c,
	0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x47,
	0x4e, 0x41, 0x4c, 0x45, 0x44, 0x10, 0x27, 0x12, 0x30, 0x0a, 0x2c, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x53, 0x45, 0x52, 0x54, 0x5f, 0x57, 0x4f, 0x52,
	0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x41, 0x54, 0x54,
	0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53, 0x10, 0x28, 0x12, 0x31, 0x0a, 0x2d, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57,
	0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x29, 0x12, 0x31, 0x0a, 0x2d,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46,
	0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x2a, 0x12,
	0x32, 0x0a, 0x2e, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f,
	0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45,
	0x44, 0x10, 0x2b, 0x12, 0x36, 0x0a, 0x32, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x45,
	0x52, 0x54, 0x49, 0x45, 0x53, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x45,
	0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x4c, 0x59, 0x10, 0x2c, 0x12, 0x36, 0x0a, 0x32, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49,
	0x54, 0x59, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x49, 0x45, 0x53, 0x5f, 0x4d, 0x4f,
	0x44, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x4c,
	0x59, 0x10, 0x2d, 0x12, 0x2b, 0x0a, 0x27, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x45,
	0x52, 0x54, 0x49, 0x45, 0x53, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x2e,
	0x12, 0x32, 0x0a, 0x2e, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57,
	0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x45, 0x44, 0x10, 0x2f, 0x42, 0x86, 0x01, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x42, 0x0e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xaa, 0x02, 0x17, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x56, 0x31,
	0xea, 0x02, 0x1a, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41,
	0x70, 0x69, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_enums_v1_event_type_proto_rawDescOnce sync.Once
	file_temporal_api_enums_v1_event_type_proto_rawDescData = file_temporal_api_enums_v1_event_type_proto_rawDesc
)

func file_temporal_api_enums_v1_event_type_proto_rawDescGZIP() []byte {
	file_temporal_api_enums_v1_event_type_proto_rawDescOnce.Do(func() {
		file_temporal_api_enums_v1_event_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_enums_v1_event_type_proto_rawDescData)
	})
	return file_temporal_api_enums_v1_event_type_proto_rawDescData
}

var file_temporal_api_enums_v1_event_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_temporal_api_enums_v1_event_type_proto_goTypes = []interface{}{
	(EventType)(0), // 0: temporal.api.enums.v1.EventType
}
var file_temporal_api_enums_v1_event_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_enums_v1_event_type_proto_init() }
func file_temporal_api_enums_v1_event_type_proto_init() {
	if File_temporal_api_enums_v1_event_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_enums_v1_event_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_enums_v1_event_type_proto_goTypes,
		DependencyIndexes: file_temporal_api_enums_v1_event_type_proto_depIdxs,
		EnumInfos:         file_temporal_api_enums_v1_event_type_proto_enumTypes,
	}.Build()
	File_temporal_api_enums_v1_event_type_proto = out.File
	file_temporal_api_enums_v1_event_type_proto_rawDesc = nil
	file_temporal_api_enums_v1_event_type_proto_goTypes = nil
	file_temporal_api_enums_v1_event_type_proto_depIdxs = nil
}
