// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: centrifugo/proxy/v1/proxy.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CentrifugoProxy_Connect_FullMethodName                 = "/centrifugal.centrifugo.proxy.CentrifugoProxy/Connect"
	CentrifugoProxy_Refresh_FullMethodName                 = "/centrifugal.centrifugo.proxy.CentrifugoProxy/Refresh"
	CentrifugoProxy_Subscribe_FullMethodName               = "/centrifugal.centrifugo.proxy.CentrifugoProxy/Subscribe"
	CentrifugoProxy_Publish_FullMethodName                 = "/centrifugal.centrifugo.proxy.CentrifugoProxy/Publish"
	CentrifugoProxy_RPC_FullMethodName                     = "/centrifugal.centrifugo.proxy.CentrifugoProxy/RPC"
	CentrifugoProxy_SubRefresh_FullMethodName              = "/centrifugal.centrifugo.proxy.CentrifugoProxy/SubRefresh"
	CentrifugoProxy_SubscribeUnidirectional_FullMethodName = "/centrifugal.centrifugo.proxy.CentrifugoProxy/SubscribeUnidirectional"
	CentrifugoProxy_SubscribeBidirectional_FullMethodName  = "/centrifugal.centrifugo.proxy.CentrifugoProxy/SubscribeBidirectional"
	CentrifugoProxy_NotifyChannelState_FullMethodName      = "/centrifugal.centrifugo.proxy.CentrifugoProxy/NotifyChannelState"
)

// CentrifugoProxyClient is the client API for CentrifugoProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CentrifugoProxyClient interface {
	// Connect to proxy connection authentication and communicate initial state.
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	// Refresh to proxy decision about connection expiration to the app backend.
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error)
	// Subscribe to proxy subscription attempts to channels.
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error)
	// Publish to proxy publication attempts to channels.
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	// RPC to execute custom logic on the backend over request through the real-time connection.
	RPC(ctx context.Context, in *RPCRequest, opts ...grpc.CallOption) (*RPCResponse, error)
	// SubRefresh to proxy decision about subscription expiration to the app backend.
	SubRefresh(ctx context.Context, in *SubRefreshRequest, opts ...grpc.CallOption) (*SubRefreshResponse, error)
	// SubscribeUnidirectional is an EXPERIMENTAL method which allows handling unidirectional
	// subscription streams. Stream starts with SubscribeRequest similar to Subscribe rpc,
	// then expects StreamSubscribeResponse with SubscribeResponse as first message, and
	// StreamSubscribeResponse with Publication afterwards.
	SubscribeUnidirectional(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (CentrifugoProxy_SubscribeUnidirectionalClient, error)
	// SubscribeBidirectional is an EXPERIMENTAL method which allows handling bidirectional
	// subscription streams. Stream receives StreamSubscribeRequest. First StreamSubscribeRequest
	// always contains SubscribeRequest, then StreamSubscribeRequest will contain data published
	// by client. Reverse direction works the same way as in SubscribeUnidirectional.
	SubscribeBidirectional(ctx context.Context, opts ...grpc.CallOption) (CentrifugoProxy_SubscribeBidirectionalClient, error)
	// NotifyChannelState can be used to receive channel events such as channel "occupied" and "vacated".
	// This is a feature in a preview state and is only available in Centrifugo PRO.
	NotifyChannelState(ctx context.Context, in *NotifyChannelStateRequest, opts ...grpc.CallOption) (*NotifyChannelStateResponse, error)
}

type centrifugoProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewCentrifugoProxyClient(cc grpc.ClientConnInterface) CentrifugoProxyClient {
	return &centrifugoProxyClient{cc}
}

func (c *centrifugoProxyClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, CentrifugoProxy_Connect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centrifugoProxyClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	out := new(RefreshResponse)
	err := c.cc.Invoke(ctx, CentrifugoProxy_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centrifugoProxyClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error) {
	out := new(SubscribeResponse)
	err := c.cc.Invoke(ctx, CentrifugoProxy_Subscribe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centrifugoProxyClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, CentrifugoProxy_Publish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centrifugoProxyClient) RPC(ctx context.Context, in *RPCRequest, opts ...grpc.CallOption) (*RPCResponse, error) {
	out := new(RPCResponse)
	err := c.cc.Invoke(ctx, CentrifugoProxy_RPC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centrifugoProxyClient) SubRefresh(ctx context.Context, in *SubRefreshRequest, opts ...grpc.CallOption) (*SubRefreshResponse, error) {
	out := new(SubRefreshResponse)
	err := c.cc.Invoke(ctx, CentrifugoProxy_SubRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centrifugoProxyClient) SubscribeUnidirectional(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (CentrifugoProxy_SubscribeUnidirectionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &CentrifugoProxy_ServiceDesc.Streams[0], CentrifugoProxy_SubscribeUnidirectional_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &centrifugoProxySubscribeUnidirectionalClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CentrifugoProxy_SubscribeUnidirectionalClient interface {
	Recv() (*StreamSubscribeResponse, error)
	grpc.ClientStream
}

type centrifugoProxySubscribeUnidirectionalClient struct {
	grpc.ClientStream
}

func (x *centrifugoProxySubscribeUnidirectionalClient) Recv() (*StreamSubscribeResponse, error) {
	m := new(StreamSubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *centrifugoProxyClient) SubscribeBidirectional(ctx context.Context, opts ...grpc.CallOption) (CentrifugoProxy_SubscribeBidirectionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &CentrifugoProxy_ServiceDesc.Streams[1], CentrifugoProxy_SubscribeBidirectional_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &centrifugoProxySubscribeBidirectionalClient{stream}
	return x, nil
}

type CentrifugoProxy_SubscribeBidirectionalClient interface {
	Send(*StreamSubscribeRequest) error
	Recv() (*StreamSubscribeResponse, error)
	grpc.ClientStream
}

type centrifugoProxySubscribeBidirectionalClient struct {
	grpc.ClientStream
}

func (x *centrifugoProxySubscribeBidirectionalClient) Send(m *StreamSubscribeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *centrifugoProxySubscribeBidirectionalClient) Recv() (*StreamSubscribeResponse, error) {
	m := new(StreamSubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *centrifugoProxyClient) NotifyChannelState(ctx context.Context, in *NotifyChannelStateRequest, opts ...grpc.CallOption) (*NotifyChannelStateResponse, error) {
	out := new(NotifyChannelStateResponse)
	err := c.cc.Invoke(ctx, CentrifugoProxy_NotifyChannelState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CentrifugoProxyServer is the server API for CentrifugoProxy service.
// All implementations should embed UnimplementedCentrifugoProxyServer
// for forward compatibility
type CentrifugoProxyServer interface {
	// Connect to proxy connection authentication and communicate initial state.
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	// Refresh to proxy decision about connection expiration to the app backend.
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)
	// Subscribe to proxy subscription attempts to channels.
	Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error)
	// Publish to proxy publication attempts to channels.
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	// RPC to execute custom logic on the backend over request through the real-time connection.
	RPC(context.Context, *RPCRequest) (*RPCResponse, error)
	// SubRefresh to proxy decision about subscription expiration to the app backend.
	SubRefresh(context.Context, *SubRefreshRequest) (*SubRefreshResponse, error)
	// SubscribeUnidirectional is an EXPERIMENTAL method which allows handling unidirectional
	// subscription streams. Stream starts with SubscribeRequest similar to Subscribe rpc,
	// then expects StreamSubscribeResponse with SubscribeResponse as first message, and
	// StreamSubscribeResponse with Publication afterwards.
	SubscribeUnidirectional(*SubscribeRequest, CentrifugoProxy_SubscribeUnidirectionalServer) error
	// SubscribeBidirectional is an EXPERIMENTAL method which allows handling bidirectional
	// subscription streams. Stream receives StreamSubscribeRequest. First StreamSubscribeRequest
	// always contains SubscribeRequest, then StreamSubscribeRequest will contain data published
	// by client. Reverse direction works the same way as in SubscribeUnidirectional.
	SubscribeBidirectional(CentrifugoProxy_SubscribeBidirectionalServer) error
	// NotifyChannelState can be used to receive channel events such as channel "occupied" and "vacated".
	// This is a feature in a preview state and is only available in Centrifugo PRO.
	NotifyChannelState(context.Context, *NotifyChannelStateRequest) (*NotifyChannelStateResponse, error)
}

// UnimplementedCentrifugoProxyServer should be embedded to have forward compatible implementations.
type UnimplementedCentrifugoProxyServer struct {
}

func (UnimplementedCentrifugoProxyServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedCentrifugoProxyServer) Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedCentrifugoProxyServer) Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedCentrifugoProxyServer) Publish(context.Context, *PublishRequest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedCentrifugoProxyServer) RPC(context.Context, *RPCRequest) (*RPCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RPC not implemented")
}
func (UnimplementedCentrifugoProxyServer) SubRefresh(context.Context, *SubRefreshRequest) (*SubRefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubRefresh not implemented")
}
func (UnimplementedCentrifugoProxyServer) SubscribeUnidirectional(*SubscribeRequest, CentrifugoProxy_SubscribeUnidirectionalServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeUnidirectional not implemented")
}
func (UnimplementedCentrifugoProxyServer) SubscribeBidirectional(CentrifugoProxy_SubscribeBidirectionalServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeBidirectional not implemented")
}
func (UnimplementedCentrifugoProxyServer) NotifyChannelState(context.Context, *NotifyChannelStateRequest) (*NotifyChannelStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyChannelState not implemented")
}

// UnsafeCentrifugoProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CentrifugoProxyServer will
// result in compilation errors.
type UnsafeCentrifugoProxyServer interface {
	mustEmbedUnimplementedCentrifugoProxyServer()
}

func RegisterCentrifugoProxyServer(s grpc.ServiceRegistrar, srv CentrifugoProxyServer) {
	s.RegisterService(&CentrifugoProxy_ServiceDesc, srv)
}

func _CentrifugoProxy_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentrifugoProxyServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CentrifugoProxy_Connect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentrifugoProxyServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentrifugoProxy_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentrifugoProxyServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CentrifugoProxy_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentrifugoProxyServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentrifugoProxy_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentrifugoProxyServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CentrifugoProxy_Subscribe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentrifugoProxyServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentrifugoProxy_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentrifugoProxyServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CentrifugoProxy_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentrifugoProxyServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentrifugoProxy_RPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentrifugoProxyServer).RPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CentrifugoProxy_RPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentrifugoProxyServer).RPC(ctx, req.(*RPCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentrifugoProxy_SubRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubRefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentrifugoProxyServer).SubRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CentrifugoProxy_SubRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentrifugoProxyServer).SubRefresh(ctx, req.(*SubRefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentrifugoProxy_SubscribeUnidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CentrifugoProxyServer).SubscribeUnidirectional(m, &centrifugoProxySubscribeUnidirectionalServer{stream})
}

type CentrifugoProxy_SubscribeUnidirectionalServer interface {
	Send(*StreamSubscribeResponse) error
	grpc.ServerStream
}

type centrifugoProxySubscribeUnidirectionalServer struct {
	grpc.ServerStream
}

func (x *centrifugoProxySubscribeUnidirectionalServer) Send(m *StreamSubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CentrifugoProxy_SubscribeBidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CentrifugoProxyServer).SubscribeBidirectional(&centrifugoProxySubscribeBidirectionalServer{stream})
}

type CentrifugoProxy_SubscribeBidirectionalServer interface {
	Send(*StreamSubscribeResponse) error
	Recv() (*StreamSubscribeRequest, error)
	grpc.ServerStream
}

type centrifugoProxySubscribeBidirectionalServer struct {
	grpc.ServerStream
}

func (x *centrifugoProxySubscribeBidirectionalServer) Send(m *StreamSubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *centrifugoProxySubscribeBidirectionalServer) Recv() (*StreamSubscribeRequest, error) {
	m := new(StreamSubscribeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CentrifugoProxy_NotifyChannelState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyChannelStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentrifugoProxyServer).NotifyChannelState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CentrifugoProxy_NotifyChannelState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentrifugoProxyServer).NotifyChannelState(ctx, req.(*NotifyChannelStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CentrifugoProxy_ServiceDesc is the grpc.ServiceDesc for CentrifugoProxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CentrifugoProxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "centrifugal.centrifugo.proxy.CentrifugoProxy",
	HandlerType: (*CentrifugoProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _CentrifugoProxy_Connect_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _CentrifugoProxy_Refresh_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _CentrifugoProxy_Subscribe_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _CentrifugoProxy_Publish_Handler,
		},
		{
			MethodName: "RPC",
			Handler:    _CentrifugoProxy_RPC_Handler,
		},
		{
			MethodName: "SubRefresh",
			Handler:    _CentrifugoProxy_SubRefresh_Handler,
		},
		{
			MethodName: "NotifyChannelState",
			Handler:    _CentrifugoProxy_NotifyChannelState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeUnidirectional",
			Handler:       _CentrifugoProxy_SubscribeUnidirectional_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeBidirectional",
			Handler:       _CentrifugoProxy_SubscribeBidirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "centrifugo/proxy/v1/proxy.proto",
}
