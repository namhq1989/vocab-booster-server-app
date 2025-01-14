// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: vocabularypb/hub.proto

package vocabularypb

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
	VocabularyService_SearchVocabulary_FullMethodName               = "/vocabularypb.VocabularyService/SearchVocabulary"
	VocabularyService_BookmarkVocabulary_FullMethodName             = "/vocabularypb.VocabularyService/BookmarkVocabulary"
	VocabularyService_GetUserBookmarkedVocabularies_FullMethodName  = "/vocabularypb.VocabularyService/GetUserBookmarkedVocabularies"
	VocabularyService_GetWordOfTheDay_FullMethodName                = "/vocabularypb.VocabularyService/GetWordOfTheDay"
	VocabularyService_GetCommunitySentences_FullMethodName          = "/vocabularypb.VocabularyService/GetCommunitySentences"
	VocabularyService_GetCommunitySentence_FullMethodName           = "/vocabularypb.VocabularyService/GetCommunitySentence"
	VocabularyService_GetUserCommunitySentenceDrafts_FullMethodName = "/vocabularypb.VocabularyService/GetUserCommunitySentenceDrafts"
)

// VocabularyServiceClient is the client API for VocabularyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VocabularyServiceClient interface {
	SearchVocabulary(ctx context.Context, in *SearchVocabularyRequest, opts ...grpc.CallOption) (*SearchVocabularyResponse, error)
	BookmarkVocabulary(ctx context.Context, in *BookmarkVocabularyRequest, opts ...grpc.CallOption) (*BookmarkVocabularyResponse, error)
	GetUserBookmarkedVocabularies(ctx context.Context, in *GetUserBookmarkedVocabulariesRequest, opts ...grpc.CallOption) (*GetUserBookmarkedVocabulariesResponse, error)
	GetWordOfTheDay(ctx context.Context, in *GetWordOfTheDayRequest, opts ...grpc.CallOption) (*GetWordOfTheDayResponse, error)
	GetCommunitySentences(ctx context.Context, in *GetCommunitySentencesRequest, opts ...grpc.CallOption) (*GetCommunitySentencesResponse, error)
	GetCommunitySentence(ctx context.Context, in *GetCommunitySentenceRequest, opts ...grpc.CallOption) (*GetCommunitySentenceResponse, error)
	GetUserCommunitySentenceDrafts(ctx context.Context, in *GetUserCommunitySentenceDraftsRequest, opts ...grpc.CallOption) (*GetUserCommunitySentenceDraftsResponse, error)
}

type vocabularyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVocabularyServiceClient(cc grpc.ClientConnInterface) VocabularyServiceClient {
	return &vocabularyServiceClient{cc}
}

func (c *vocabularyServiceClient) SearchVocabulary(ctx context.Context, in *SearchVocabularyRequest, opts ...grpc.CallOption) (*SearchVocabularyResponse, error) {
	out := new(SearchVocabularyResponse)
	err := c.cc.Invoke(ctx, VocabularyService_SearchVocabulary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) BookmarkVocabulary(ctx context.Context, in *BookmarkVocabularyRequest, opts ...grpc.CallOption) (*BookmarkVocabularyResponse, error) {
	out := new(BookmarkVocabularyResponse)
	err := c.cc.Invoke(ctx, VocabularyService_BookmarkVocabulary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) GetUserBookmarkedVocabularies(ctx context.Context, in *GetUserBookmarkedVocabulariesRequest, opts ...grpc.CallOption) (*GetUserBookmarkedVocabulariesResponse, error) {
	out := new(GetUserBookmarkedVocabulariesResponse)
	err := c.cc.Invoke(ctx, VocabularyService_GetUserBookmarkedVocabularies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) GetWordOfTheDay(ctx context.Context, in *GetWordOfTheDayRequest, opts ...grpc.CallOption) (*GetWordOfTheDayResponse, error) {
	out := new(GetWordOfTheDayResponse)
	err := c.cc.Invoke(ctx, VocabularyService_GetWordOfTheDay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) GetCommunitySentences(ctx context.Context, in *GetCommunitySentencesRequest, opts ...grpc.CallOption) (*GetCommunitySentencesResponse, error) {
	out := new(GetCommunitySentencesResponse)
	err := c.cc.Invoke(ctx, VocabularyService_GetCommunitySentences_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) GetCommunitySentence(ctx context.Context, in *GetCommunitySentenceRequest, opts ...grpc.CallOption) (*GetCommunitySentenceResponse, error) {
	out := new(GetCommunitySentenceResponse)
	err := c.cc.Invoke(ctx, VocabularyService_GetCommunitySentence_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) GetUserCommunitySentenceDrafts(ctx context.Context, in *GetUserCommunitySentenceDraftsRequest, opts ...grpc.CallOption) (*GetUserCommunitySentenceDraftsResponse, error) {
	out := new(GetUserCommunitySentenceDraftsResponse)
	err := c.cc.Invoke(ctx, VocabularyService_GetUserCommunitySentenceDrafts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VocabularyServiceServer is the server API for VocabularyService service.
// All implementations should embed UnimplementedVocabularyServiceServer
// for forward compatibility
type VocabularyServiceServer interface {
	SearchVocabulary(context.Context, *SearchVocabularyRequest) (*SearchVocabularyResponse, error)
	BookmarkVocabulary(context.Context, *BookmarkVocabularyRequest) (*BookmarkVocabularyResponse, error)
	GetUserBookmarkedVocabularies(context.Context, *GetUserBookmarkedVocabulariesRequest) (*GetUserBookmarkedVocabulariesResponse, error)
	GetWordOfTheDay(context.Context, *GetWordOfTheDayRequest) (*GetWordOfTheDayResponse, error)
	GetCommunitySentences(context.Context, *GetCommunitySentencesRequest) (*GetCommunitySentencesResponse, error)
	GetCommunitySentence(context.Context, *GetCommunitySentenceRequest) (*GetCommunitySentenceResponse, error)
	GetUserCommunitySentenceDrafts(context.Context, *GetUserCommunitySentenceDraftsRequest) (*GetUserCommunitySentenceDraftsResponse, error)
}

// UnimplementedVocabularyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedVocabularyServiceServer struct {
}

func (UnimplementedVocabularyServiceServer) SearchVocabulary(context.Context, *SearchVocabularyRequest) (*SearchVocabularyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchVocabulary not implemented")
}
func (UnimplementedVocabularyServiceServer) BookmarkVocabulary(context.Context, *BookmarkVocabularyRequest) (*BookmarkVocabularyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookmarkVocabulary not implemented")
}
func (UnimplementedVocabularyServiceServer) GetUserBookmarkedVocabularies(context.Context, *GetUserBookmarkedVocabulariesRequest) (*GetUserBookmarkedVocabulariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBookmarkedVocabularies not implemented")
}
func (UnimplementedVocabularyServiceServer) GetWordOfTheDay(context.Context, *GetWordOfTheDayRequest) (*GetWordOfTheDayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWordOfTheDay not implemented")
}
func (UnimplementedVocabularyServiceServer) GetCommunitySentences(context.Context, *GetCommunitySentencesRequest) (*GetCommunitySentencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommunitySentences not implemented")
}
func (UnimplementedVocabularyServiceServer) GetCommunitySentence(context.Context, *GetCommunitySentenceRequest) (*GetCommunitySentenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommunitySentence not implemented")
}
func (UnimplementedVocabularyServiceServer) GetUserCommunitySentenceDrafts(context.Context, *GetUserCommunitySentenceDraftsRequest) (*GetUserCommunitySentenceDraftsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCommunitySentenceDrafts not implemented")
}

// UnsafeVocabularyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VocabularyServiceServer will
// result in compilation errors.
type UnsafeVocabularyServiceServer interface {
	mustEmbedUnimplementedVocabularyServiceServer()
}

func RegisterVocabularyServiceServer(s grpc.ServiceRegistrar, srv VocabularyServiceServer) {
	s.RegisterService(&VocabularyService_ServiceDesc, srv)
}

func _VocabularyService_SearchVocabulary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchVocabularyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).SearchVocabulary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_SearchVocabulary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).SearchVocabulary(ctx, req.(*SearchVocabularyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_BookmarkVocabulary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookmarkVocabularyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).BookmarkVocabulary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_BookmarkVocabulary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).BookmarkVocabulary(ctx, req.(*BookmarkVocabularyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_GetUserBookmarkedVocabularies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserBookmarkedVocabulariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).GetUserBookmarkedVocabularies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_GetUserBookmarkedVocabularies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).GetUserBookmarkedVocabularies(ctx, req.(*GetUserBookmarkedVocabulariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_GetWordOfTheDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWordOfTheDayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).GetWordOfTheDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_GetWordOfTheDay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).GetWordOfTheDay(ctx, req.(*GetWordOfTheDayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_GetCommunitySentences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommunitySentencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).GetCommunitySentences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_GetCommunitySentences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).GetCommunitySentences(ctx, req.(*GetCommunitySentencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_GetCommunitySentence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommunitySentenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).GetCommunitySentence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_GetCommunitySentence_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).GetCommunitySentence(ctx, req.(*GetCommunitySentenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_GetUserCommunitySentenceDrafts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCommunitySentenceDraftsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).GetUserCommunitySentenceDrafts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_GetUserCommunitySentenceDrafts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).GetUserCommunitySentenceDrafts(ctx, req.(*GetUserCommunitySentenceDraftsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VocabularyService_ServiceDesc is the grpc.ServiceDesc for VocabularyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VocabularyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vocabularypb.VocabularyService",
	HandlerType: (*VocabularyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchVocabulary",
			Handler:    _VocabularyService_SearchVocabulary_Handler,
		},
		{
			MethodName: "BookmarkVocabulary",
			Handler:    _VocabularyService_BookmarkVocabulary_Handler,
		},
		{
			MethodName: "GetUserBookmarkedVocabularies",
			Handler:    _VocabularyService_GetUserBookmarkedVocabularies_Handler,
		},
		{
			MethodName: "GetWordOfTheDay",
			Handler:    _VocabularyService_GetWordOfTheDay_Handler,
		},
		{
			MethodName: "GetCommunitySentences",
			Handler:    _VocabularyService_GetCommunitySentences_Handler,
		},
		{
			MethodName: "GetCommunitySentence",
			Handler:    _VocabularyService_GetCommunitySentence_Handler,
		},
		{
			MethodName: "GetUserCommunitySentenceDrafts",
			Handler:    _VocabularyService_GetUserCommunitySentenceDrafts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vocabularypb/hub.proto",
}
