// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// FilmServiceClient is the client API for FilmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilmServiceClient interface {
	ListFilms(ctx context.Context, in *ListFilmReq, opts ...grpc.CallOption) (FilmService_ListFilmsClient, error)
}

type filmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFilmServiceClient(cc grpc.ClientConnInterface) FilmServiceClient {
	return &filmServiceClient{cc}
}

func (c *filmServiceClient) ListFilms(ctx context.Context, in *ListFilmReq, opts ...grpc.CallOption) (FilmService_ListFilmsClient, error) {
	stream, err := c.cc.NewStream(ctx, &FilmService_ServiceDesc.Streams[0], "/pb.FilmService/ListFilms", opts...)
	if err != nil {
		return nil, err
	}
	x := &filmServiceListFilmsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FilmService_ListFilmsClient interface {
	Recv() (*ListFilmRes, error)
	grpc.ClientStream
}

type filmServiceListFilmsClient struct {
	grpc.ClientStream
}

func (x *filmServiceListFilmsClient) Recv() (*ListFilmRes, error) {
	m := new(ListFilmRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FilmServiceServer is the server API for FilmService service.
// All implementations must embed UnimplementedFilmServiceServer
// for forward compatibility
type FilmServiceServer interface {
	ListFilms(*ListFilmReq, FilmService_ListFilmsServer) error
	mustEmbedUnimplementedFilmServiceServer()
}

// UnimplementedFilmServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFilmServiceServer struct {
}

func (UnimplementedFilmServiceServer) ListFilms(*ListFilmReq, FilmService_ListFilmsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListFilms not implemented")
}
func (UnimplementedFilmServiceServer) mustEmbedUnimplementedFilmServiceServer() {}

// UnsafeFilmServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilmServiceServer will
// result in compilation errors.
type UnsafeFilmServiceServer interface {
	mustEmbedUnimplementedFilmServiceServer()
}

func RegisterFilmServiceServer(s grpc.ServiceRegistrar, srv FilmServiceServer) {
	s.RegisterService(&FilmService_ServiceDesc, srv)
}

func _FilmService_ListFilms_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListFilmReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FilmServiceServer).ListFilms(m, &filmServiceListFilmsServer{stream})
}

type FilmService_ListFilmsServer interface {
	Send(*ListFilmRes) error
	grpc.ServerStream
}

type filmServiceListFilmsServer struct {
	grpc.ServerStream
}

func (x *filmServiceListFilmsServer) Send(m *ListFilmRes) error {
	return x.ServerStream.SendMsg(m)
}

// FilmService_ServiceDesc is the grpc.ServiceDesc for FilmService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FilmService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.FilmService",
	HandlerType: (*FilmServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFilms",
			Handler:       _FilmService_ListFilms_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "films.proto",
}
