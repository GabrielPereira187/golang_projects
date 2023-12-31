// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: proto/cnpj.proto

package gen

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
	CnpjService_GetCnpjInfo_FullMethodName = "/cnpj.CnpjService/GetCnpjInfo"
)

// CnpjServiceClient is the client API for CnpjService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CnpjServiceClient interface {
	GetCnpjInfo(ctx context.Context, in *CnpjRequest, opts ...grpc.CallOption) (*CnpjResponse, error)
}

type cnpjServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCnpjServiceClient(cc grpc.ClientConnInterface) CnpjServiceClient {
	return &cnpjServiceClient{cc}
}

func (c *cnpjServiceClient) GetCnpjInfo(ctx context.Context, in *CnpjRequest, opts ...grpc.CallOption) (*CnpjResponse, error) {
	out := new(CnpjResponse)
	err := c.cc.Invoke(ctx, CnpjService_GetCnpjInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CnpjServiceServer is the server API for CnpjService service.
// All implementations must embed UnimplementedCnpjServiceServer
// for forward compatibility
type CnpjServiceServer interface {
	GetCnpjInfo(context.Context, *CnpjRequest) (*CnpjResponse, error)
	mustEmbedUnimplementedCnpjServiceServer()
}

// UnimplementedCnpjServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCnpjServiceServer struct {
}

func (UnimplementedCnpjServiceServer) GetCnpjInfo(context.Context, *CnpjRequest) (*CnpjResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCnpjInfo not implemented")
}
func (UnimplementedCnpjServiceServer) mustEmbedUnimplementedCnpjServiceServer() {}

// UnsafeCnpjServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CnpjServiceServer will
// result in compilation errors.
type UnsafeCnpjServiceServer interface {
	mustEmbedUnimplementedCnpjServiceServer()
}

func RegisterCnpjServiceServer(s grpc.ServiceRegistrar, srv CnpjServiceServer) {
	s.RegisterService(&CnpjService_ServiceDesc, srv)
}

func _CnpjService_GetCnpjInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CnpjRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CnpjServiceServer).GetCnpjInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CnpjService_GetCnpjInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CnpjServiceServer).GetCnpjInfo(ctx, req.(*CnpjRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CnpjService_ServiceDesc is the grpc.ServiceDesc for CnpjService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CnpjService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cnpj.CnpjService",
	HandlerType: (*CnpjServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCnpjInfo",
			Handler:    _CnpjService_GetCnpjInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cnpj.proto",
}
