// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package importpath

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

// ExtensionListerClient is the client API for ExtensionLister service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExtensionListerClient interface {
	Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	ListExtensions(ctx context.Context, in *ListExtensionRequest, opts ...grpc.CallOption) (ExtensionLister_ListExtensionsClient, error)
}

type extensionListerClient struct {
	cc grpc.ClientConnInterface
}

func NewExtensionListerClient(cc grpc.ClientConnInterface) ExtensionListerClient {
	return &extensionListerClient{cc}
}

func (c *extensionListerClient) Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/lister.ExtensionLister/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extensionListerClient) ListExtensions(ctx context.Context, in *ListExtensionRequest, opts ...grpc.CallOption) (ExtensionLister_ListExtensionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExtensionLister_ServiceDesc.Streams[0], "/lister.ExtensionLister/listExtensions", opts...)
	if err != nil {
		return nil, err
	}
	x := &extensionListerListExtensionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExtensionLister_ListExtensionsClient interface {
	Recv() (*Extension, error)
	grpc.ClientStream
}

type extensionListerListExtensionsClient struct {
	grpc.ClientStream
}

func (x *extensionListerListExtensionsClient) Recv() (*Extension, error) {
	m := new(Extension)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExtensionListerServer is the server API for ExtensionLister service.
// All implementations must embed UnimplementedExtensionListerServer
// for forward compatibility
type ExtensionListerServer interface {
	Authenticate(context.Context, *AuthRequest) (*AuthResponse, error)
	ListExtensions(*ListExtensionRequest, ExtensionLister_ListExtensionsServer) error
	mustEmbedUnimplementedExtensionListerServer()
}

// UnimplementedExtensionListerServer must be embedded to have forward compatible implementations.
type UnimplementedExtensionListerServer struct {
}

func (UnimplementedExtensionListerServer) Authenticate(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedExtensionListerServer) ListExtensions(*ListExtensionRequest, ExtensionLister_ListExtensionsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListExtensions not implemented")
}
func (UnimplementedExtensionListerServer) mustEmbedUnimplementedExtensionListerServer() {}

// UnsafeExtensionListerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExtensionListerServer will
// result in compilation errors.
type UnsafeExtensionListerServer interface {
	mustEmbedUnimplementedExtensionListerServer()
}

func RegisterExtensionListerServer(s grpc.ServiceRegistrar, srv ExtensionListerServer) {
	s.RegisterService(&ExtensionLister_ServiceDesc, srv)
}

func _ExtensionLister_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionListerServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lister.ExtensionLister/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionListerServer).Authenticate(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExtensionLister_ListExtensions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListExtensionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExtensionListerServer).ListExtensions(m, &extensionListerListExtensionsServer{stream})
}

type ExtensionLister_ListExtensionsServer interface {
	Send(*Extension) error
	grpc.ServerStream
}

type extensionListerListExtensionsServer struct {
	grpc.ServerStream
}

func (x *extensionListerListExtensionsServer) Send(m *Extension) error {
	return x.ServerStream.SendMsg(m)
}

// ExtensionLister_ServiceDesc is the grpc.ServiceDesc for ExtensionLister service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExtensionLister_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lister.ExtensionLister",
	HandlerType: (*ExtensionListerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _ExtensionLister_Authenticate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "listExtensions",
			Handler:       _ExtensionLister_ListExtensions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/extension_list.proto",
}
