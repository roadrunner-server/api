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
// source: temporal/api/replication/v1/message.proto

package replication

import (
	v1 "go.temporal.io/api/enums/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type ClusterReplicationConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClusterName   string                 `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClusterReplicationConfig) Reset() {
	*x = ClusterReplicationConfig{}
	mi := &file_temporal_api_replication_v1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClusterReplicationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterReplicationConfig) ProtoMessage() {}

func (x *ClusterReplicationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_replication_v1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterReplicationConfig.ProtoReflect.Descriptor instead.
func (*ClusterReplicationConfig) Descriptor() ([]byte, []int) {
	return file_temporal_api_replication_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterReplicationConfig) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

type NamespaceReplicationConfig struct {
	state             protoimpl.MessageState      `protogen:"open.v1"`
	ActiveClusterName string                      `protobuf:"bytes,1,opt,name=active_cluster_name,json=activeClusterName,proto3" json:"active_cluster_name,omitempty"`
	Clusters          []*ClusterReplicationConfig `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	State             v1.ReplicationState         `protobuf:"varint,3,opt,name=state,proto3,enum=temporal.api.enums.v1.ReplicationState" json:"state,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *NamespaceReplicationConfig) Reset() {
	*x = NamespaceReplicationConfig{}
	mi := &file_temporal_api_replication_v1_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamespaceReplicationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceReplicationConfig) ProtoMessage() {}

func (x *NamespaceReplicationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_replication_v1_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceReplicationConfig.ProtoReflect.Descriptor instead.
func (*NamespaceReplicationConfig) Descriptor() ([]byte, []int) {
	return file_temporal_api_replication_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *NamespaceReplicationConfig) GetActiveClusterName() string {
	if x != nil {
		return x.ActiveClusterName
	}
	return ""
}

func (x *NamespaceReplicationConfig) GetClusters() []*ClusterReplicationConfig {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *NamespaceReplicationConfig) GetState() v1.ReplicationState {
	if x != nil {
		return x.State
	}
	return v1.ReplicationState(0)
}

// Represents a historical replication status of a Namespace
type FailoverStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Timestamp when the Cluster switched to the following failover_version
	FailoverTime    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=failover_time,json=failoverTime,proto3" json:"failover_time,omitempty"`
	FailoverVersion int64                  `protobuf:"varint,2,opt,name=failover_version,json=failoverVersion,proto3" json:"failover_version,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *FailoverStatus) Reset() {
	*x = FailoverStatus{}
	mi := &file_temporal_api_replication_v1_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FailoverStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailoverStatus) ProtoMessage() {}

func (x *FailoverStatus) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_replication_v1_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailoverStatus.ProtoReflect.Descriptor instead.
func (*FailoverStatus) Descriptor() ([]byte, []int) {
	return file_temporal_api_replication_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *FailoverStatus) GetFailoverTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FailoverTime
	}
	return nil
}

func (x *FailoverStatus) GetFailoverVersion() int64 {
	if x != nil {
		return x.FailoverVersion
	}
	return 0
}

var File_temporal_api_replication_v1_message_proto protoreflect.FileDescriptor

const file_temporal_api_replication_v1_message_proto_rawDesc = "" +
	"\n" +
	")temporal/api/replication/v1/message.proto\x12\x1btemporal.api.replication.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a%temporal/api/enums/v1/namespace.proto\"=\n" +
	"\x18ClusterReplicationConfig\x12!\n" +
	"\fcluster_name\x18\x01 \x01(\tR\vclusterName\"\xde\x01\n" +
	"\x1aNamespaceReplicationConfig\x12.\n" +
	"\x13active_cluster_name\x18\x01 \x01(\tR\x11activeClusterName\x12Q\n" +
	"\bclusters\x18\x02 \x03(\v25.temporal.api.replication.v1.ClusterReplicationConfigR\bclusters\x12=\n" +
	"\x05state\x18\x03 \x01(\x0e2'.temporal.api.enums.v1.ReplicationStateR\x05state\"|\n" +
	"\x0eFailoverStatus\x12?\n" +
	"\rfailover_time\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\ffailoverTime\x12)\n" +
	"\x10failover_version\x18\x02 \x01(\x03R\x0ffailoverVersionB\xa2\x01\n" +
	"\x1eio.temporal.api.replication.v1B\fMessageProtoP\x01Z-go.temporal.io/api/replication/v1;replication\xaa\x02\x1dTemporalio.Api.Replication.V1\xea\x02 Temporalio::Api::Replication::V1b\x06proto3"

var (
	file_temporal_api_replication_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_replication_v1_message_proto_rawDescData []byte
)

func file_temporal_api_replication_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_replication_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_replication_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_temporal_api_replication_v1_message_proto_rawDesc), len(file_temporal_api_replication_v1_message_proto_rawDesc)))
	})
	return file_temporal_api_replication_v1_message_proto_rawDescData
}

var file_temporal_api_replication_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_temporal_api_replication_v1_message_proto_goTypes = []any{
	(*ClusterReplicationConfig)(nil),   // 0: temporal.api.replication.v1.ClusterReplicationConfig
	(*NamespaceReplicationConfig)(nil), // 1: temporal.api.replication.v1.NamespaceReplicationConfig
	(*FailoverStatus)(nil),             // 2: temporal.api.replication.v1.FailoverStatus
	(v1.ReplicationState)(0),           // 3: temporal.api.enums.v1.ReplicationState
	(*timestamppb.Timestamp)(nil),      // 4: google.protobuf.Timestamp
}
var file_temporal_api_replication_v1_message_proto_depIdxs = []int32{
	0, // 0: temporal.api.replication.v1.NamespaceReplicationConfig.clusters:type_name -> temporal.api.replication.v1.ClusterReplicationConfig
	3, // 1: temporal.api.replication.v1.NamespaceReplicationConfig.state:type_name -> temporal.api.enums.v1.ReplicationState
	4, // 2: temporal.api.replication.v1.FailoverStatus.failover_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_temporal_api_replication_v1_message_proto_init() }
func file_temporal_api_replication_v1_message_proto_init() {
	if File_temporal_api_replication_v1_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_temporal_api_replication_v1_message_proto_rawDesc), len(file_temporal_api_replication_v1_message_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_replication_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_replication_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_api_replication_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_replication_v1_message_proto = out.File
	file_temporal_api_replication_v1_message_proto_goTypes = nil
	file_temporal_api_replication_v1_message_proto_depIdxs = nil
}
