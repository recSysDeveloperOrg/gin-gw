// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package movie

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

// MovieServiceClient is the client API for MovieService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieServiceClient interface {
	RecommendMovies(ctx context.Context, in *RecommendReq, opts ...grpc.CallOption) (*RecommendResp, error)
	GetMovieDetail(ctx context.Context, in *MovieDetailReq, opts ...grpc.CallOption) (*MovieDetailResp, error)
	SearchMovies(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchResp, error)
	RecommendFeedback(ctx context.Context, in *FeedbackReq, opts ...grpc.CallOption) (*FeedbackResp, error)
	ModifyMovieRating(ctx context.Context, in *ModifyMovieRatingReq, opts ...grpc.CallOption) (*ModifyMovieRatingResp, error)
}

type movieServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieServiceClient(cc grpc.ClientConnInterface) MovieServiceClient {
	return &movieServiceClient{cc}
}

func (c *movieServiceClient) RecommendMovies(ctx context.Context, in *RecommendReq, opts ...grpc.CallOption) (*RecommendResp, error) {
	out := new(RecommendResp)
	err := c.cc.Invoke(ctx, "/MovieService/RecommendMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) GetMovieDetail(ctx context.Context, in *MovieDetailReq, opts ...grpc.CallOption) (*MovieDetailResp, error) {
	out := new(MovieDetailResp)
	err := c.cc.Invoke(ctx, "/MovieService/GetMovieDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) SearchMovies(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchResp, error) {
	out := new(SearchResp)
	err := c.cc.Invoke(ctx, "/MovieService/SearchMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) RecommendFeedback(ctx context.Context, in *FeedbackReq, opts ...grpc.CallOption) (*FeedbackResp, error) {
	out := new(FeedbackResp)
	err := c.cc.Invoke(ctx, "/MovieService/RecommendFeedback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) ModifyMovieRating(ctx context.Context, in *ModifyMovieRatingReq, opts ...grpc.CallOption) (*ModifyMovieRatingResp, error) {
	out := new(ModifyMovieRatingResp)
	err := c.cc.Invoke(ctx, "/MovieService/ModifyMovieRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieServiceServer is the server API for MovieService service.
// All implementations must embed UnimplementedMovieServiceServer
// for forward compatibility
type MovieServiceServer interface {
	RecommendMovies(context.Context, *RecommendReq) (*RecommendResp, error)
	GetMovieDetail(context.Context, *MovieDetailReq) (*MovieDetailResp, error)
	SearchMovies(context.Context, *SearchReq) (*SearchResp, error)
	RecommendFeedback(context.Context, *FeedbackReq) (*FeedbackResp, error)
	ModifyMovieRating(context.Context, *ModifyMovieRatingReq) (*ModifyMovieRatingResp, error)
	mustEmbedUnimplementedMovieServiceServer()
}

// UnimplementedMovieServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMovieServiceServer struct {
}

func (UnimplementedMovieServiceServer) RecommendMovies(context.Context, *RecommendReq) (*RecommendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecommendMovies not implemented")
}
func (UnimplementedMovieServiceServer) GetMovieDetail(context.Context, *MovieDetailReq) (*MovieDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieDetail not implemented")
}
func (UnimplementedMovieServiceServer) SearchMovies(context.Context, *SearchReq) (*SearchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMovies not implemented")
}
func (UnimplementedMovieServiceServer) RecommendFeedback(context.Context, *FeedbackReq) (*FeedbackResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecommendFeedback not implemented")
}
func (UnimplementedMovieServiceServer) ModifyMovieRating(context.Context, *ModifyMovieRatingReq) (*ModifyMovieRatingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyMovieRating not implemented")
}
func (UnimplementedMovieServiceServer) mustEmbedUnimplementedMovieServiceServer() {}

// UnsafeMovieServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieServiceServer will
// result in compilation errors.
type UnsafeMovieServiceServer interface {
	mustEmbedUnimplementedMovieServiceServer()
}

func RegisterMovieServiceServer(s grpc.ServiceRegistrar, srv MovieServiceServer) {
	s.RegisterService(&MovieService_ServiceDesc, srv)
}

func _MovieService_RecommendMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).RecommendMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieService/RecommendMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).RecommendMovies(ctx, req.(*RecommendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_GetMovieDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).GetMovieDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieService/GetMovieDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).GetMovieDetail(ctx, req.(*MovieDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_SearchMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).SearchMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieService/SearchMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).SearchMovies(ctx, req.(*SearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_RecommendFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedbackReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).RecommendFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieService/RecommendFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).RecommendFeedback(ctx, req.(*FeedbackReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_ModifyMovieRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyMovieRatingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).ModifyMovieRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieService/ModifyMovieRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).ModifyMovieRating(ctx, req.(*ModifyMovieRatingReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MovieService_ServiceDesc is the grpc.ServiceDesc for MovieService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MovieService",
	HandlerType: (*MovieServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecommendMovies",
			Handler:    _MovieService_RecommendMovies_Handler,
		},
		{
			MethodName: "GetMovieDetail",
			Handler:    _MovieService_GetMovieDetail_Handler,
		},
		{
			MethodName: "SearchMovies",
			Handler:    _MovieService_SearchMovies_Handler,
		},
		{
			MethodName: "RecommendFeedback",
			Handler:    _MovieService_RecommendFeedback_Handler,
		},
		{
			MethodName: "ModifyMovieRating",
			Handler:    _MovieService_ModifyMovieRating_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idl/movie.proto",
}