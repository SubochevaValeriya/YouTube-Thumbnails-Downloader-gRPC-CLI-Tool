// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/youtube_thumbnails.proto

package grpcYoutubeThumbnails

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

// YoutubeThumbnailsServiceClient is the client API for YoutubeThumbnailsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YoutubeThumbnailsServiceClient interface {
	LoadThumbnail(ctx context.Context, in *LoadThumbnailLinkRequest, opts ...grpc.CallOption) (*LoadThumbnailLinkResponse, error)
}

type youtubeThumbnailsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYoutubeThumbnailsServiceClient(cc grpc.ClientConnInterface) YoutubeThumbnailsServiceClient {
	return &youtubeThumbnailsServiceClient{cc}
}

func (c *youtubeThumbnailsServiceClient) LoadThumbnail(ctx context.Context, in *LoadThumbnailLinkRequest, opts ...grpc.CallOption) (*LoadThumbnailLinkResponse, error) {
	out := new(LoadThumbnailLinkResponse)
	err := c.cc.Invoke(ctx, "/proto.YoutubeThumbnailsService/LoadThumbnail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YoutubeThumbnailsServiceServer is the server API for YoutubeThumbnailsService service.
// All implementations must embed UnimplementedYoutubeThumbnailsServiceServer
// for forward compatibility
type YoutubeThumbnailsServiceServer interface {
	LoadThumbnail(context.Context, *LoadThumbnailLinkRequest) (*LoadThumbnailLinkResponse, error)
	mustEmbedUnimplementedYoutubeThumbnailsServiceServer()
}

// UnimplementedYoutubeThumbnailsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedYoutubeThumbnailsServiceServer struct {
}

func (UnimplementedYoutubeThumbnailsServiceServer) LoadThumbnail(context.Context, *LoadThumbnailLinkRequest) (*LoadThumbnailLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadThumbnail not implemented")
}
func (UnimplementedYoutubeThumbnailsServiceServer) mustEmbedUnimplementedYoutubeThumbnailsServiceServer() {
}

// UnsafeYoutubeThumbnailsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YoutubeThumbnailsServiceServer will
// result in compilation errors.
type UnsafeYoutubeThumbnailsServiceServer interface {
	mustEmbedUnimplementedYoutubeThumbnailsServiceServer()
}

func RegisterYoutubeThumbnailsServiceServer(s grpc.ServiceRegistrar, srv YoutubeThumbnailsServiceServer) {
	s.RegisterService(&YoutubeThumbnailsService_ServiceDesc, srv)
}

func _YoutubeThumbnailsService_LoadThumbnail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadThumbnailLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YoutubeThumbnailsServiceServer).LoadThumbnail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.YoutubeThumbnailsService/LoadThumbnail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YoutubeThumbnailsServiceServer).LoadThumbnail(ctx, req.(*LoadThumbnailLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// YoutubeThumbnailsService_ServiceDesc is the grpc.ServiceDesc for YoutubeThumbnailsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YoutubeThumbnailsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.YoutubeThumbnailsService",
	HandlerType: (*YoutubeThumbnailsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoadThumbnail",
			Handler:    _YoutubeThumbnailsService_LoadThumbnail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/youtube_thumbnails.proto",
}
