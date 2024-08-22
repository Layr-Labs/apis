// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: eigenlayer/proofs_service/v1/api.proto

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
	ProofsService_GetCheckpointProofs_FullMethodName            = "/eigenlayer.apis.api.v1.ProofsService/GetCheckpointProofs"
	ProofsService_GetWithdrawalCredentialsProofs_FullMethodName = "/eigenlayer.apis.api.v1.ProofsService/GetWithdrawalCredentialsProofs"
)

// ProofsServiceClient is the client API for ProofsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProofsServiceClient interface {
	GetCheckpointProofs(ctx context.Context, in *GetCheckpointProofsRequest, opts ...grpc.CallOption) (*GetCheckpointProofsResponse, error)
	GetWithdrawalCredentialsProofs(ctx context.Context, in *GetWithdrawalCredentialsProofsRequest, opts ...grpc.CallOption) (*GetWithdrawalCredentialsProofsResponse, error)
}

type proofsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProofsServiceClient(cc grpc.ClientConnInterface) ProofsServiceClient {
	return &proofsServiceClient{cc}
}

func (c *proofsServiceClient) GetCheckpointProofs(ctx context.Context, in *GetCheckpointProofsRequest, opts ...grpc.CallOption) (*GetCheckpointProofsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCheckpointProofsResponse)
	err := c.cc.Invoke(ctx, ProofsService_GetCheckpointProofs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proofsServiceClient) GetWithdrawalCredentialsProofs(ctx context.Context, in *GetWithdrawalCredentialsProofsRequest, opts ...grpc.CallOption) (*GetWithdrawalCredentialsProofsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWithdrawalCredentialsProofsResponse)
	err := c.cc.Invoke(ctx, ProofsService_GetWithdrawalCredentialsProofs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProofsServiceServer is the server API for ProofsService service.
// All implementations should embed UnimplementedProofsServiceServer
// for forward compatibility.
type ProofsServiceServer interface {
	GetCheckpointProofs(context.Context, *GetCheckpointProofsRequest) (*GetCheckpointProofsResponse, error)
	GetWithdrawalCredentialsProofs(context.Context, *GetWithdrawalCredentialsProofsRequest) (*GetWithdrawalCredentialsProofsResponse, error)
}

// UnimplementedProofsServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProofsServiceServer struct{}

func (UnimplementedProofsServiceServer) GetCheckpointProofs(context.Context, *GetCheckpointProofsRequest) (*GetCheckpointProofsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckpointProofs not implemented")
}
func (UnimplementedProofsServiceServer) GetWithdrawalCredentialsProofs(context.Context, *GetWithdrawalCredentialsProofsRequest) (*GetWithdrawalCredentialsProofsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWithdrawalCredentialsProofs not implemented")
}
func (UnimplementedProofsServiceServer) testEmbeddedByValue() {}

// UnsafeProofsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProofsServiceServer will
// result in compilation errors.
type UnsafeProofsServiceServer interface {
	mustEmbedUnimplementedProofsServiceServer()
}

func RegisterProofsServiceServer(s grpc.ServiceRegistrar, srv ProofsServiceServer) {
	// If the following call pancis, it indicates UnimplementedProofsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProofsService_ServiceDesc, srv)
}

func _ProofsService_GetCheckpointProofs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCheckpointProofsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProofsServiceServer).GetCheckpointProofs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProofsService_GetCheckpointProofs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProofsServiceServer).GetCheckpointProofs(ctx, req.(*GetCheckpointProofsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProofsService_GetWithdrawalCredentialsProofs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWithdrawalCredentialsProofsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProofsServiceServer).GetWithdrawalCredentialsProofs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProofsService_GetWithdrawalCredentialsProofs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProofsServiceServer).GetWithdrawalCredentialsProofs(ctx, req.(*GetWithdrawalCredentialsProofsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProofsService_ServiceDesc is the grpc.ServiceDesc for ProofsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProofsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eigenlayer.apis.api.v1.ProofsService",
	HandlerType: (*ProofsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCheckpointProofs",
			Handler:    _ProofsService_GetCheckpointProofs_Handler,
		},
		{
			MethodName: "GetWithdrawalCredentialsProofs",
			Handler:    _ProofsService_GetWithdrawalCredentialsProofs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eigenlayer/proofs_service/v1/api.proto",
}
