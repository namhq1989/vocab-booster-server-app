// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: exercisepb/hub.proto

package exercisepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ExerciseService_AnswerExercise_FullMethodName                 = "/exercisepb.ExerciseService/AnswerExercise"
	ExerciseService_GetUserExercises_FullMethodName               = "/exercisepb.ExerciseService/GetUserExercises"
	ExerciseService_GetUserReadyForReviewExercises_FullMethodName = "/exercisepb.ExerciseService/GetUserReadyForReviewExercises"
	ExerciseService_GetUserStats_FullMethodName                   = "/exercisepb.ExerciseService/GetUserStats"
)

// ExerciseServiceClient is the client API for ExerciseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExerciseServiceClient interface {
	AnswerExercise(ctx context.Context, in *AnswerExerciseRequest, opts ...grpc.CallOption) (*AnswerExerciseResponse, error)
	GetUserExercises(ctx context.Context, in *GetUserExercisesRequest, opts ...grpc.CallOption) (*GetUserExercisesResponse, error)
	GetUserReadyForReviewExercises(ctx context.Context, in *GetUserReadyForReviewExercisesRequest, opts ...grpc.CallOption) (*GetUserReadyForReviewExercisesResponse, error)
	GetUserStats(ctx context.Context, in *GetUserStatsRequest, opts ...grpc.CallOption) (*GetUserStatsResponse, error)
}

type exerciseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExerciseServiceClient(cc grpc.ClientConnInterface) ExerciseServiceClient {
	return &exerciseServiceClient{cc}
}

func (c *exerciseServiceClient) AnswerExercise(ctx context.Context, in *AnswerExerciseRequest, opts ...grpc.CallOption) (*AnswerExerciseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AnswerExerciseResponse)
	err := c.cc.Invoke(ctx, ExerciseService_AnswerExercise_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exerciseServiceClient) GetUserExercises(ctx context.Context, in *GetUserExercisesRequest, opts ...grpc.CallOption) (*GetUserExercisesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserExercisesResponse)
	err := c.cc.Invoke(ctx, ExerciseService_GetUserExercises_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exerciseServiceClient) GetUserReadyForReviewExercises(ctx context.Context, in *GetUserReadyForReviewExercisesRequest, opts ...grpc.CallOption) (*GetUserReadyForReviewExercisesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserReadyForReviewExercisesResponse)
	err := c.cc.Invoke(ctx, ExerciseService_GetUserReadyForReviewExercises_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exerciseServiceClient) GetUserStats(ctx context.Context, in *GetUserStatsRequest, opts ...grpc.CallOption) (*GetUserStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserStatsResponse)
	err := c.cc.Invoke(ctx, ExerciseService_GetUserStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExerciseServiceServer is the server API for ExerciseService service.
// All implementations should embed UnimplementedExerciseServiceServer
// for forward compatibility
type ExerciseServiceServer interface {
	AnswerExercise(context.Context, *AnswerExerciseRequest) (*AnswerExerciseResponse, error)
	GetUserExercises(context.Context, *GetUserExercisesRequest) (*GetUserExercisesResponse, error)
	GetUserReadyForReviewExercises(context.Context, *GetUserReadyForReviewExercisesRequest) (*GetUserReadyForReviewExercisesResponse, error)
	GetUserStats(context.Context, *GetUserStatsRequest) (*GetUserStatsResponse, error)
}

// UnimplementedExerciseServiceServer should be embedded to have forward compatible implementations.
type UnimplementedExerciseServiceServer struct {
}

func (UnimplementedExerciseServiceServer) AnswerExercise(context.Context, *AnswerExerciseRequest) (*AnswerExerciseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnswerExercise not implemented")
}
func (UnimplementedExerciseServiceServer) GetUserExercises(context.Context, *GetUserExercisesRequest) (*GetUserExercisesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserExercises not implemented")
}
func (UnimplementedExerciseServiceServer) GetUserReadyForReviewExercises(context.Context, *GetUserReadyForReviewExercisesRequest) (*GetUserReadyForReviewExercisesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserReadyForReviewExercises not implemented")
}
func (UnimplementedExerciseServiceServer) GetUserStats(context.Context, *GetUserStatsRequest) (*GetUserStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserStats not implemented")
}

// UnsafeExerciseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExerciseServiceServer will
// result in compilation errors.
type UnsafeExerciseServiceServer interface {
	mustEmbedUnimplementedExerciseServiceServer()
}

func RegisterExerciseServiceServer(s grpc.ServiceRegistrar, srv ExerciseServiceServer) {
	s.RegisterService(&ExerciseService_ServiceDesc, srv)
}

func _ExerciseService_AnswerExercise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnswerExerciseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseServiceServer).AnswerExercise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExerciseService_AnswerExercise_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseServiceServer).AnswerExercise(ctx, req.(*AnswerExerciseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExerciseService_GetUserExercises_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserExercisesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseServiceServer).GetUserExercises(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExerciseService_GetUserExercises_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseServiceServer).GetUserExercises(ctx, req.(*GetUserExercisesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExerciseService_GetUserReadyForReviewExercises_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReadyForReviewExercisesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseServiceServer).GetUserReadyForReviewExercises(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExerciseService_GetUserReadyForReviewExercises_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseServiceServer).GetUserReadyForReviewExercises(ctx, req.(*GetUserReadyForReviewExercisesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExerciseService_GetUserStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseServiceServer).GetUserStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExerciseService_GetUserStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseServiceServer).GetUserStats(ctx, req.(*GetUserStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExerciseService_ServiceDesc is the grpc.ServiceDesc for ExerciseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExerciseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "exercisepb.ExerciseService",
	HandlerType: (*ExerciseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnswerExercise",
			Handler:    _ExerciseService_AnswerExercise_Handler,
		},
		{
			MethodName: "GetUserExercises",
			Handler:    _ExerciseService_GetUserExercises_Handler,
		},
		{
			MethodName: "GetUserReadyForReviewExercises",
			Handler:    _ExerciseService_GetUserReadyForReviewExercises_Handler,
		},
		{
			MethodName: "GetUserStats",
			Handler:    _ExerciseService_GetUserStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exercisepb/hub.proto",
}
