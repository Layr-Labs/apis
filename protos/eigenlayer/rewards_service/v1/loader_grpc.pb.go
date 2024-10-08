// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: eigenlayer/rewards_service/v1/loader.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProofsLoaderService_ReloadProofs_FullMethodName = "/eigenlayer.apis.loader.v1.ProofsLoaderService/ReloadProofs"
)

// ProofsLoaderServiceClient is the client API for ProofsLoaderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProofsLoaderServiceClient interface {
	ReloadProofs(ctx context.Context, in *ReloadProofsRequest, opts ...grpc.CallOption) (*ReloadProofsResponse, error)
}

type proofsLoaderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProofsLoaderServiceClient(cc grpc.ClientConnInterface) ProofsLoaderServiceClient {
	return &proofsLoaderServiceClient{cc}
}

func (c *proofsLoaderServiceClient) ReloadProofs(ctx context.Context, in *ReloadProofsRequest, opts ...grpc.CallOption) (*ReloadProofsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReloadProofsResponse)
	err := c.cc.Invoke(ctx, ProofsLoaderService_ReloadProofs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProofsLoaderServiceServer is the server API for ProofsLoaderService service.
// All implementations should embed UnimplementedProofsLoaderServiceServer
// for forward compatibility.
type ProofsLoaderServiceServer interface {
	ReloadProofs(context.Context, *ReloadProofsRequest) (*ReloadProofsResponse, error)
}

// UnimplementedProofsLoaderServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProofsLoaderServiceServer struct{}

func (UnimplementedProofsLoaderServiceServer) ReloadProofs(context.Context, *ReloadProofsRequest) (*ReloadProofsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReloadProofs not implemented")
}
func (UnimplementedProofsLoaderServiceServer) testEmbeddedByValue() {}

// UnsafeProofsLoaderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProofsLoaderServiceServer will
// result in compilation errors.
type UnsafeProofsLoaderServiceServer interface {
	mustEmbedUnimplementedProofsLoaderServiceServer()
}

func RegisterProofsLoaderServiceServer(s grpc.ServiceRegistrar, srv ProofsLoaderServiceServer) {
	// If the following call pancis, it indicates UnimplementedProofsLoaderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProofsLoaderService_ServiceDesc, srv)
}

func _ProofsLoaderService_ReloadProofs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloadProofsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProofsLoaderServiceServer).ReloadProofs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProofsLoaderService_ReloadProofs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProofsLoaderServiceServer).ReloadProofs(ctx, req.(*ReloadProofsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProofsLoaderService_ServiceDesc is the grpc.ServiceDesc for ProofsLoaderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProofsLoaderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eigenlayer.apis.loader.v1.ProofsLoaderService",
	HandlerType: (*ProofsLoaderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReloadProofs",
			Handler:    _ProofsLoaderService_ReloadProofs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eigenlayer/rewards_service/v1/loader.proto",
}
