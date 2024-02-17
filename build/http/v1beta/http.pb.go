// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: http/v1beta/http.proto

package v1beta

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoteAddr string                  `protobuf:"bytes,1,opt,name=remote_addr,json=remoteAddr,proto3" json:"remote_addr,omitempty"`
	Protocol   string                  `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Method     string                  `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Uri        string                  `protobuf:"bytes,4,opt,name=uri,proto3" json:"uri,omitempty"`
	Header     map[string]*HeaderValue `protobuf:"bytes,5,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Cookies    map[string]*HeaderValue `protobuf:"bytes,6,rep,name=cookies,proto3" json:"cookies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FormQuery  map[string]*HeaderValue `protobuf:"bytes,7,rep,name=form_query,json=formQuery,proto3" json:"form_query,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Uploads    *Uploads                `protobuf:"bytes,10,opt,name=uploads,proto3" json:"uploads,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_v1beta_http_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_http_v1beta_http_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_http_v1beta_http_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

func (x *Request) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Request) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Request) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Request) GetHeader() map[string]*HeaderValue {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Request) GetCookies() map[string]*HeaderValue {
	if x != nil {
		return x.Cookies
	}
	return nil
}

func (x *Request) GetFormQuery() map[string]*HeaderValue {
	if x != nil {
		return x.FormQuery
	}
	return nil
}

func (x *Request) GetUploads() *Uploads {
	if x != nil {
		return x.Uploads
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int64                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Headers map[string]*HeaderValue `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_v1beta_http_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_http_v1beta_http_proto_msgTypes[1]
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
	return file_http_v1beta_http_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Response) GetHeaders() map[string]*HeaderValue {
	if x != nil {
		return x.Headers
	}
	return nil
}

type Uploads struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*FileUpload `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *Uploads) Reset() {
	*x = Uploads{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_v1beta_http_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Uploads) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uploads) ProtoMessage() {}

func (x *Uploads) ProtoReflect() protoreflect.Message {
	mi := &file_http_v1beta_http_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uploads.ProtoReflect.Descriptor instead.
func (*Uploads) Descriptor() ([]byte, []int) {
	return file_http_v1beta_http_proto_rawDescGZIP(), []int{2}
}

func (x *Uploads) GetList() []*FileUpload {
	if x != nil {
		return x.List
	}
	return nil
}

type FileUpload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Mime         string     `protobuf:"bytes,2,opt,name=mime,proto3" json:"mime,omitempty"`
	Size         int64      `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Error        int64      `protobuf:"varint,4,opt,name=error,proto3" json:"error,omitempty"`
	TempFilename string     `protobuf:"bytes,5,opt,name=temp_filename,json=tempFilename,proto3" json:"temp_filename,omitempty"`
	Header       *anypb.Any `protobuf:"bytes,9,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *FileUpload) Reset() {
	*x = FileUpload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_v1beta_http_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUpload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUpload) ProtoMessage() {}

func (x *FileUpload) ProtoReflect() protoreflect.Message {
	mi := &file_http_v1beta_http_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUpload.ProtoReflect.Descriptor instead.
func (*FileUpload) Descriptor() ([]byte, []int) {
	return file_http_v1beta_http_proto_rawDescGZIP(), []int{3}
}

func (x *FileUpload) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileUpload) GetMime() string {
	if x != nil {
		return x.Mime
	}
	return ""
}

func (x *FileUpload) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileUpload) GetError() int64 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *FileUpload) GetTempFilename() string {
	if x != nil {
		return x.TempFilename
	}
	return ""
}

func (x *FileUpload) GetHeader() *anypb.Any {
	if x != nil {
		return x.Header
	}
	return nil
}

type HeaderValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *HeaderValue) Reset() {
	*x = HeaderValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_v1beta_http_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderValue) ProtoMessage() {}

func (x *HeaderValue) ProtoReflect() protoreflect.Message {
	mi := &file_http_v1beta_http_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderValue.ProtoReflect.Descriptor instead.
func (*HeaderValue) Descriptor() ([]byte, []int) {
	return file_http_v1beta_http_proto_rawDescGZIP(), []int{4}
}

func (x *HeaderValue) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_http_v1beta_http_proto protoreflect.FileDescriptor

var file_http_v1beta_http_proto_rawDesc = []byte{
	0x0a, 0x16, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xde, 0x04, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x69, 0x12, 0x38, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3b, 0x0a,
	0x07, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x0a, 0x66, 0x6f,
	0x72, 0x6d, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x09, 0x66, 0x6f, 0x72, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2e,
	0x0a, 0x07, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x07, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x1a, 0x53,
	0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x54, 0x0a, 0x0c, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x56, 0x0a, 0x0e, 0x46, 0x6f, 0x72,
	0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xb6, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x1a, 0x54, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x36, 0x0a, 0x07, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0xb1, 0x01, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70,
	0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x23, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x55, 0x5a, 0x0b, 0x68,
	0x74, 0x74, 0x70, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0xca, 0x02, 0x1b, 0x52, 0x6f, 0x61,
	0x64, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x5c, 0x48, 0x54, 0x54, 0x50, 0x5c, 0x44, 0x54, 0x4f,
	0x5c, 0x56, 0x31, 0x42, 0x45, 0x54, 0x41, 0x31, 0xe2, 0x02, 0x27, 0x52, 0x6f, 0x61, 0x64, 0x52,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x5c, 0x48, 0x54, 0x54, 0x50, 0x5c, 0x44, 0x54, 0x4f, 0x5c, 0x56,
	0x31, 0x42, 0x45, 0x54, 0x41, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_http_v1beta_http_proto_rawDescOnce sync.Once
	file_http_v1beta_http_proto_rawDescData = file_http_v1beta_http_proto_rawDesc
)

func file_http_v1beta_http_proto_rawDescGZIP() []byte {
	file_http_v1beta_http_proto_rawDescOnce.Do(func() {
		file_http_v1beta_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_http_v1beta_http_proto_rawDescData)
	})
	return file_http_v1beta_http_proto_rawDescData
}

var file_http_v1beta_http_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_http_v1beta_http_proto_goTypes = []interface{}{
	(*Request)(nil),     // 0: http.v1beta.Request
	(*Response)(nil),    // 1: http.v1beta.Response
	(*Uploads)(nil),     // 2: http.v1beta.Uploads
	(*FileUpload)(nil),  // 3: http.v1beta.FileUpload
	(*HeaderValue)(nil), // 4: http.v1beta.HeaderValue
	nil,                 // 5: http.v1beta.Request.HeaderEntry
	nil,                 // 6: http.v1beta.Request.CookiesEntry
	nil,                 // 7: http.v1beta.Request.FormQueryEntry
	nil,                 // 8: http.v1beta.Response.HeadersEntry
	(*anypb.Any)(nil),   // 9: google.protobuf.Any
}
var file_http_v1beta_http_proto_depIdxs = []int32{
	5,  // 0: http.v1beta.Request.header:type_name -> http.v1beta.Request.HeaderEntry
	6,  // 1: http.v1beta.Request.cookies:type_name -> http.v1beta.Request.CookiesEntry
	7,  // 2: http.v1beta.Request.form_query:type_name -> http.v1beta.Request.FormQueryEntry
	2,  // 3: http.v1beta.Request.uploads:type_name -> http.v1beta.Uploads
	8,  // 4: http.v1beta.Response.headers:type_name -> http.v1beta.Response.HeadersEntry
	3,  // 5: http.v1beta.Uploads.list:type_name -> http.v1beta.FileUpload
	9,  // 6: http.v1beta.FileUpload.header:type_name -> google.protobuf.Any
	4,  // 7: http.v1beta.Request.HeaderEntry.value:type_name -> http.v1beta.HeaderValue
	4,  // 8: http.v1beta.Request.CookiesEntry.value:type_name -> http.v1beta.HeaderValue
	4,  // 9: http.v1beta.Request.FormQueryEntry.value:type_name -> http.v1beta.HeaderValue
	4,  // 10: http.v1beta.Response.HeadersEntry.value:type_name -> http.v1beta.HeaderValue
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_http_v1beta_http_proto_init() }
func file_http_v1beta_http_proto_init() {
	if File_http_v1beta_http_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_http_v1beta_http_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_http_v1beta_http_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_http_v1beta_http_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Uploads); i {
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
		file_http_v1beta_http_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUpload); i {
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
		file_http_v1beta_http_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderValue); i {
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
			RawDescriptor: file_http_v1beta_http_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_http_v1beta_http_proto_goTypes,
		DependencyIndexes: file_http_v1beta_http_proto_depIdxs,
		MessageInfos:      file_http_v1beta_http_proto_msgTypes,
	}.Build()
	File_http_v1beta_http_proto = out.File
	file_http_v1beta_http_proto_rawDesc = nil
	file_http_v1beta_http_proto_goTypes = nil
	file_http_v1beta_http_proto_depIdxs = nil
}
