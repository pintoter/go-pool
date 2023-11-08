// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: api/proto/math.proto

package pb

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

// GeneratorClient is the client API for Generator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeneratorClient interface {
	GetGenerateData(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (Generator_GetGenerateDataClient, error)
}

type generatorClient struct {
	cc grpc.ClientConnInterface
}

func NewGeneratorClient(cc grpc.ClientConnInterface) GeneratorClient {
	return &generatorClient{cc}
}

func (c *generatorClient) GetGenerateData(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (Generator_GetGenerateDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &Generator_ServiceDesc.Streams[0], "/pkg.Generator/GetGenerateData", opts...)
	if err != nil {
		return nil, err
	}
	x := &generatorGetGenerateDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Generator_GetGenerateDataClient interface {
	Recv() (*DataResponse, error)
	grpc.ClientStream
}

type generatorGetGenerateDataClient struct {
	grpc.ClientStream
}

func (x *generatorGetGenerateDataClient) Recv() (*DataResponse, error) {
	m := new(DataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GeneratorServer is the server API for Generator service.
// All implementations must embed UnimplementedGeneratorServer
// for forward compatibility
type GeneratorServer interface {
	GetGenerateData(*DataRequest, Generator_GetGenerateDataServer) error
	mustEmbedUnimplementedGeneratorServer()
}

// UnimplementedGeneratorServer must be embedded to have forward compatible implementations.
type UnimplementedGeneratorServer struct {
}

func (UnimplementedGeneratorServer) GetGenerateData(*DataRequest, Generator_GetGenerateDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetGenerateData not implemented")
}
func (UnimplementedGeneratorServer) mustEmbedUnimplementedGeneratorServer() {}

// UnsafeGeneratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeneratorServer will
// result in compilation errors.
type UnsafeGeneratorServer interface {
	mustEmbedUnimplementedGeneratorServer()
}

func RegisterGeneratorServer(s grpc.ServiceRegistrar, srv GeneratorServer) {
	s.RegisterService(&Generator_ServiceDesc, srv)
}

func _Generator_GetGenerateData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GeneratorServer).GetGenerateData(m, &generatorGetGenerateDataServer{stream})
}

type Generator_GetGenerateDataServer interface {
	Send(*DataResponse) error
	grpc.ServerStream
}

type generatorGetGenerateDataServer struct {
	grpc.ServerStream
}

func (x *generatorGetGenerateDataServer) Send(m *DataResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Generator_ServiceDesc is the grpc.ServiceDesc for Generator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Generator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pkg.Generator",
	HandlerType: (*GeneratorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetGenerateData",
			Handler:       _Generator_GetGenerateData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/proto/math.proto",
}
