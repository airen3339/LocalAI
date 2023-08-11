// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pkg/grpc/proto/backend.proto

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

const (
	Backend_Health_FullMethodName             = "/backend.Backend/Health"
	Backend_Predict_FullMethodName            = "/backend.Backend/Predict"
	Backend_LoadModel_FullMethodName          = "/backend.Backend/LoadModel"
	Backend_PredictStream_FullMethodName      = "/backend.Backend/PredictStream"
	Backend_Embedding_FullMethodName          = "/backend.Backend/Embedding"
	Backend_GenerateImage_FullMethodName      = "/backend.Backend/GenerateImage"
	Backend_AudioTranscription_FullMethodName = "/backend.Backend/AudioTranscription"
	Backend_TTS_FullMethodName                = "/backend.Backend/TTS"
	Backend_TokenizeString_FullMethodName     = "/backend.Backend/TokenizeString"
	Backend_State_FullMethodName              = "/backend.Backend/State"
)

// BackendClient is the client API for Backend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackendClient interface {
	Health(ctx context.Context, in *HealthMessage, opts ...grpc.CallOption) (*Reply, error)
	Predict(ctx context.Context, in *PredictOptions, opts ...grpc.CallOption) (*Reply, error)
	LoadModel(ctx context.Context, in *ModelOptions, opts ...grpc.CallOption) (*Result, error)
	PredictStream(ctx context.Context, in *PredictOptions, opts ...grpc.CallOption) (Backend_PredictStreamClient, error)
	Embedding(ctx context.Context, in *PredictOptions, opts ...grpc.CallOption) (*EmbeddingResult, error)
	GenerateImage(ctx context.Context, in *GenerateImageRequest, opts ...grpc.CallOption) (*Result, error)
	AudioTranscription(ctx context.Context, in *TranscriptRequest, opts ...grpc.CallOption) (*TranscriptResult, error)
	TTS(ctx context.Context, in *TTSRequest, opts ...grpc.CallOption) (*Result, error)
	TokenizeString(ctx context.Context, in *PredictOptions, opts ...grpc.CallOption) (*TokenizationResponse, error)
	State(ctx context.Context, in *HealthMessage, opts ...grpc.CallOption) (*StateResponse, error)
}

type backendClient struct {
	cc grpc.ClientConnInterface
}

func NewBackendClient(cc grpc.ClientConnInterface) BackendClient {
	return &backendClient{cc}
}

func (c *backendClient) Health(ctx context.Context, in *HealthMessage, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, Backend_Health_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) Predict(ctx context.Context, in *PredictOptions, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, Backend_Predict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) LoadModel(ctx context.Context, in *ModelOptions, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, Backend_LoadModel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) PredictStream(ctx context.Context, in *PredictOptions, opts ...grpc.CallOption) (Backend_PredictStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Backend_ServiceDesc.Streams[0], Backend_PredictStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &backendPredictStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Backend_PredictStreamClient interface {
	Recv() (*Reply, error)
	grpc.ClientStream
}

type backendPredictStreamClient struct {
	grpc.ClientStream
}

func (x *backendPredictStreamClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *backendClient) Embedding(ctx context.Context, in *PredictOptions, opts ...grpc.CallOption) (*EmbeddingResult, error) {
	out := new(EmbeddingResult)
	err := c.cc.Invoke(ctx, Backend_Embedding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) GenerateImage(ctx context.Context, in *GenerateImageRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, Backend_GenerateImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) AudioTranscription(ctx context.Context, in *TranscriptRequest, opts ...grpc.CallOption) (*TranscriptResult, error) {
	out := new(TranscriptResult)
	err := c.cc.Invoke(ctx, Backend_AudioTranscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) TTS(ctx context.Context, in *TTSRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, Backend_TTS_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) TokenizeString(ctx context.Context, in *PredictOptions, opts ...grpc.CallOption) (*TokenizationResponse, error) {
	out := new(TokenizationResponse)
	err := c.cc.Invoke(ctx, Backend_TokenizeString_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) State(ctx context.Context, in *HealthMessage, opts ...grpc.CallOption) (*StateResponse, error) {
	out := new(StateResponse)
	err := c.cc.Invoke(ctx, Backend_State_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackendServer is the server API for Backend service.
// All implementations must embed UnimplementedBackendServer
// for forward compatibility
type BackendServer interface {
	Health(context.Context, *HealthMessage) (*Reply, error)
	Predict(context.Context, *PredictOptions) (*Reply, error)
	LoadModel(context.Context, *ModelOptions) (*Result, error)
	PredictStream(*PredictOptions, Backend_PredictStreamServer) error
	Embedding(context.Context, *PredictOptions) (*EmbeddingResult, error)
	GenerateImage(context.Context, *GenerateImageRequest) (*Result, error)
	AudioTranscription(context.Context, *TranscriptRequest) (*TranscriptResult, error)
	TTS(context.Context, *TTSRequest) (*Result, error)
	TokenizeString(context.Context, *PredictOptions) (*TokenizationResponse, error)
	State(context.Context, *HealthMessage) (*StateResponse, error)
	mustEmbedUnimplementedBackendServer()
}

// UnimplementedBackendServer must be embedded to have forward compatible implementations.
type UnimplementedBackendServer struct {
}

func (UnimplementedBackendServer) Health(context.Context, *HealthMessage) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedBackendServer) Predict(context.Context, *PredictOptions) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predict not implemented")
}
func (UnimplementedBackendServer) LoadModel(context.Context, *ModelOptions) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadModel not implemented")
}
func (UnimplementedBackendServer) PredictStream(*PredictOptions, Backend_PredictStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method PredictStream not implemented")
}
func (UnimplementedBackendServer) Embedding(context.Context, *PredictOptions) (*EmbeddingResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Embedding not implemented")
}
func (UnimplementedBackendServer) GenerateImage(context.Context, *GenerateImageRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateImage not implemented")
}
func (UnimplementedBackendServer) AudioTranscription(context.Context, *TranscriptRequest) (*TranscriptResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AudioTranscription not implemented")
}
func (UnimplementedBackendServer) TTS(context.Context, *TTSRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TTS not implemented")
}
func (UnimplementedBackendServer) TokenizeString(context.Context, *PredictOptions) (*TokenizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenizeString not implemented")
}
func (UnimplementedBackendServer) State(context.Context, *HealthMessage) (*StateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method State not implemented")
}
func (UnimplementedBackendServer) mustEmbedUnimplementedBackendServer() {}

// UnsafeBackendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackendServer will
// result in compilation errors.
type UnsafeBackendServer interface {
	mustEmbedUnimplementedBackendServer()
}

func RegisterBackendServer(s grpc.ServiceRegistrar, srv BackendServer) {
	s.RegisterService(&Backend_ServiceDesc, srv)
}

func _Backend_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_Health_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).Health(ctx, req.(*HealthMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_Predict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).Predict(ctx, req.(*PredictOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_LoadModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).LoadModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_LoadModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).LoadModel(ctx, req.(*ModelOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_PredictStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PredictOptions)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BackendServer).PredictStream(m, &backendPredictStreamServer{stream})
}

type Backend_PredictStreamServer interface {
	Send(*Reply) error
	grpc.ServerStream
}

type backendPredictStreamServer struct {
	grpc.ServerStream
}

func (x *backendPredictStreamServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

func _Backend_Embedding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).Embedding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_Embedding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).Embedding(ctx, req.(*PredictOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_GenerateImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).GenerateImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_GenerateImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).GenerateImage(ctx, req.(*GenerateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_AudioTranscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranscriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).AudioTranscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_AudioTranscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).AudioTranscription(ctx, req.(*TranscriptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_TTS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TTSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).TTS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_TTS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).TTS(ctx, req.(*TTSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_TokenizeString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).TokenizeString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_TokenizeString_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).TokenizeString(ctx, req.(*PredictOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_State_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).State(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backend_State_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).State(ctx, req.(*HealthMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// Backend_ServiceDesc is the grpc.ServiceDesc for Backend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Backend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "backend.Backend",
	HandlerType: (*BackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _Backend_Health_Handler,
		},
		{
			MethodName: "Predict",
			Handler:    _Backend_Predict_Handler,
		},
		{
			MethodName: "LoadModel",
			Handler:    _Backend_LoadModel_Handler,
		},
		{
			MethodName: "Embedding",
			Handler:    _Backend_Embedding_Handler,
		},
		{
			MethodName: "GenerateImage",
			Handler:    _Backend_GenerateImage_Handler,
		},
		{
			MethodName: "AudioTranscription",
			Handler:    _Backend_AudioTranscription_Handler,
		},
		{
			MethodName: "TTS",
			Handler:    _Backend_TTS_Handler,
		},
		{
			MethodName: "TokenizeString",
			Handler:    _Backend_TokenizeString_Handler,
		},
		{
			MethodName: "State",
			Handler:    _Backend_State_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PredictStream",
			Handler:       _Backend_PredictStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/grpc/proto/backend.proto",
}
