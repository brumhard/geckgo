// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package geckgov1

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

// GeckgoServiceClient is the client API for GeckgoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeckgoServiceClient interface {
	AddList(ctx context.Context, in *AddListRequest, opts ...grpc.CallOption) (*AddListResponse, error)
	GetLists(ctx context.Context, in *GetListsRequest, opts ...grpc.CallOption) (*GetListsResponse, error)
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	UpdateList(ctx context.Context, in *UpdateListRequest, opts ...grpc.CallOption) (*UpdateListResponse, error)
	DeleteList(ctx context.Context, in *DeleteListRequest, opts ...grpc.CallOption) (*DeleteListResponse, error)
	AddDay(ctx context.Context, in *AddDayRequest, opts ...grpc.CallOption) (*AddDayResponse, error)
	GetDays(ctx context.Context, in *GetDaysRequest, opts ...grpc.CallOption) (*GetDaysResponse, error)
	GetDay(ctx context.Context, in *GetDayRequest, opts ...grpc.CallOption) (*GetDayResponse, error)
	UpdateDay(ctx context.Context, in *UpdateDayRequest, opts ...grpc.CallOption) (*UpdateDayResponse, error)
	DeleteDay(ctx context.Context, in *DeleteDayRequest, opts ...grpc.CallOption) (*DeleteDayResponse, error)
	StartDay(ctx context.Context, in *StartDayRequest, opts ...grpc.CallOption) (*StartDayResponse, error)
	StartBreak(ctx context.Context, in *StartBreakRequest, opts ...grpc.CallOption) (*StartBreakResponse, error)
	EndBreak(ctx context.Context, in *EndBreakRequest, opts ...grpc.CallOption) (*EndBreakResponse, error)
	EndDay(ctx context.Context, in *EndDayRequest, opts ...grpc.CallOption) (*EndDayResponse, error)
}

type geckgoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGeckgoServiceClient(cc grpc.ClientConnInterface) GeckgoServiceClient {
	return &geckgoServiceClient{cc}
}

func (c *geckgoServiceClient) AddList(ctx context.Context, in *AddListRequest, opts ...grpc.CallOption) (*AddListResponse, error) {
	out := new(AddListResponse)
	err := c.cc.Invoke(ctx, "/geckgo.v1.GeckgoService/AddList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geckgoServiceClient) GetLists(ctx context.Context, in *GetListsRequest, opts ...grpc.CallOption) (*GetListsResponse, error) {
	out := new(GetListsResponse)
	err := c.cc.Invoke(ctx, "/geckgo.v1.GeckgoService/GetLists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geckgoServiceClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/geckgo.v1.GeckgoService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geckgoServiceClient) UpdateList(ctx context.Context, in *UpdateListRequest, opts ...grpc.CallOption) (*UpdateListResponse, error) {
	out := new(UpdateListResponse)
	err := c.cc.Invoke(ctx, "/geckgo.v1.GeckgoService/UpdateList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geckgoServiceClient) DeleteList(ctx context.Context, in *DeleteListRequest, opts ...grpc.CallOption) (*DeleteListResponse, error) {
	out := new(DeleteListResponse)
	err := c.cc.Invoke(ctx, "/geckgo.v1.GeckgoService/DeleteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geckgoServiceClient) AddDay(ctx context.Context, in *AddDayRequest, opts ...grpc.CallOption) (*AddDayResponse, error) {
	out := new(AddDayResponse)
	err := c.cc.Invoke(ctx, "/geckgo.v1.GeckgoService/AddDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geckgoServiceClient) GetDays(ctx context.Context, in *GetDaysRequest, opts ...grpc.CallOption) (*GetDaysResponse, error) {
	out := new(GetDaysResponse)
	err := c.cc.Invoke(ctx, "/geckgo.v1.GeckgoService/GetDays", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geckgoServiceClient) GetDay(ctx context.Context, in *GetDayRequest, opts ...grpc.CallOption) (*GetDayResponse, error) {
	out := new(GetDayResponse)
	err := c.cc.Invoke(ctx, "/geckgo.v1.GeckgoService/GetDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geckgoServiceClient) UpdateDay(ctx context.Context, in *UpdateDayRequest, opts ...grpc.CallOption) (*UpdateDayResponse, error) {
	out := new(UpdateDayResponse)
	err := c.cc.Invoke(ctx, "/geckgo.v1.GeckgoService/UpdateDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geckgoServiceClient) DeleteDay(ctx context.Context, in *DeleteDayRequest, opts ...grpc.CallOption) (*DeleteDayResponse, error) {
	out := new(DeleteDayResponse)
	err := c.cc.Invoke(ctx, "/geckgo.v1.GeckgoService/DeleteDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geckgoServiceClient) StartDay(ctx context.Context, in *StartDayRequest, opts ...grpc.CallOption) (*StartDayResponse, error) {
	out := new(StartDayResponse)
	err := c.cc.Invoke(ctx, "/geckgo.v1.GeckgoService/StartDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geckgoServiceClient) StartBreak(ctx context.Context, in *StartBreakRequest, opts ...grpc.CallOption) (*StartBreakResponse, error) {
	out := new(StartBreakResponse)
	err := c.cc.Invoke(ctx, "/geckgo.v1.GeckgoService/StartBreak", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geckgoServiceClient) EndBreak(ctx context.Context, in *EndBreakRequest, opts ...grpc.CallOption) (*EndBreakResponse, error) {
	out := new(EndBreakResponse)
	err := c.cc.Invoke(ctx, "/geckgo.v1.GeckgoService/EndBreak", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geckgoServiceClient) EndDay(ctx context.Context, in *EndDayRequest, opts ...grpc.CallOption) (*EndDayResponse, error) {
	out := new(EndDayResponse)
	err := c.cc.Invoke(ctx, "/geckgo.v1.GeckgoService/EndDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeckgoServiceServer is the server API for GeckgoService service.
// All implementations must embed UnimplementedGeckgoServiceServer
// for forward compatibility
type GeckgoServiceServer interface {
	AddList(context.Context, *AddListRequest) (*AddListResponse, error)
	GetLists(context.Context, *GetListsRequest) (*GetListsResponse, error)
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	UpdateList(context.Context, *UpdateListRequest) (*UpdateListResponse, error)
	DeleteList(context.Context, *DeleteListRequest) (*DeleteListResponse, error)
	AddDay(context.Context, *AddDayRequest) (*AddDayResponse, error)
	GetDays(context.Context, *GetDaysRequest) (*GetDaysResponse, error)
	GetDay(context.Context, *GetDayRequest) (*GetDayResponse, error)
	UpdateDay(context.Context, *UpdateDayRequest) (*UpdateDayResponse, error)
	DeleteDay(context.Context, *DeleteDayRequest) (*DeleteDayResponse, error)
	StartDay(context.Context, *StartDayRequest) (*StartDayResponse, error)
	StartBreak(context.Context, *StartBreakRequest) (*StartBreakResponse, error)
	EndBreak(context.Context, *EndBreakRequest) (*EndBreakResponse, error)
	EndDay(context.Context, *EndDayRequest) (*EndDayResponse, error)
	mustEmbedUnimplementedGeckgoServiceServer()
}

// UnimplementedGeckgoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGeckgoServiceServer struct {
}

func (UnimplementedGeckgoServiceServer) AddList(context.Context, *AddListRequest) (*AddListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddList not implemented")
}
func (UnimplementedGeckgoServiceServer) GetLists(context.Context, *GetListsRequest) (*GetListsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLists not implemented")
}
func (UnimplementedGeckgoServiceServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedGeckgoServiceServer) UpdateList(context.Context, *UpdateListRequest) (*UpdateListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateList not implemented")
}
func (UnimplementedGeckgoServiceServer) DeleteList(context.Context, *DeleteListRequest) (*DeleteListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteList not implemented")
}
func (UnimplementedGeckgoServiceServer) AddDay(context.Context, *AddDayRequest) (*AddDayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDay not implemented")
}
func (UnimplementedGeckgoServiceServer) GetDays(context.Context, *GetDaysRequest) (*GetDaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDays not implemented")
}
func (UnimplementedGeckgoServiceServer) GetDay(context.Context, *GetDayRequest) (*GetDayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDay not implemented")
}
func (UnimplementedGeckgoServiceServer) UpdateDay(context.Context, *UpdateDayRequest) (*UpdateDayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDay not implemented")
}
func (UnimplementedGeckgoServiceServer) DeleteDay(context.Context, *DeleteDayRequest) (*DeleteDayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDay not implemented")
}
func (UnimplementedGeckgoServiceServer) StartDay(context.Context, *StartDayRequest) (*StartDayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartDay not implemented")
}
func (UnimplementedGeckgoServiceServer) StartBreak(context.Context, *StartBreakRequest) (*StartBreakResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartBreak not implemented")
}
func (UnimplementedGeckgoServiceServer) EndBreak(context.Context, *EndBreakRequest) (*EndBreakResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndBreak not implemented")
}
func (UnimplementedGeckgoServiceServer) EndDay(context.Context, *EndDayRequest) (*EndDayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndDay not implemented")
}
func (UnimplementedGeckgoServiceServer) mustEmbedUnimplementedGeckgoServiceServer() {}

// UnsafeGeckgoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeckgoServiceServer will
// result in compilation errors.
type UnsafeGeckgoServiceServer interface {
	mustEmbedUnimplementedGeckgoServiceServer()
}

func RegisterGeckgoServiceServer(s grpc.ServiceRegistrar, srv GeckgoServiceServer) {
	s.RegisterService(&GeckgoService_ServiceDesc, srv)
}

func _GeckgoService_AddList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeckgoServiceServer).AddList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geckgo.v1.GeckgoService/AddList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeckgoServiceServer).AddList(ctx, req.(*AddListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeckgoService_GetLists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeckgoServiceServer).GetLists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geckgo.v1.GeckgoService/GetLists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeckgoServiceServer).GetLists(ctx, req.(*GetListsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeckgoService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeckgoServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geckgo.v1.GeckgoService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeckgoServiceServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeckgoService_UpdateList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeckgoServiceServer).UpdateList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geckgo.v1.GeckgoService/UpdateList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeckgoServiceServer).UpdateList(ctx, req.(*UpdateListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeckgoService_DeleteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeckgoServiceServer).DeleteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geckgo.v1.GeckgoService/DeleteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeckgoServiceServer).DeleteList(ctx, req.(*DeleteListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeckgoService_AddDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeckgoServiceServer).AddDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geckgo.v1.GeckgoService/AddDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeckgoServiceServer).AddDay(ctx, req.(*AddDayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeckgoService_GetDays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeckgoServiceServer).GetDays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geckgo.v1.GeckgoService/GetDays",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeckgoServiceServer).GetDays(ctx, req.(*GetDaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeckgoService_GetDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeckgoServiceServer).GetDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geckgo.v1.GeckgoService/GetDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeckgoServiceServer).GetDay(ctx, req.(*GetDayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeckgoService_UpdateDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeckgoServiceServer).UpdateDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geckgo.v1.GeckgoService/UpdateDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeckgoServiceServer).UpdateDay(ctx, req.(*UpdateDayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeckgoService_DeleteDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeckgoServiceServer).DeleteDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geckgo.v1.GeckgoService/DeleteDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeckgoServiceServer).DeleteDay(ctx, req.(*DeleteDayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeckgoService_StartDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartDayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeckgoServiceServer).StartDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geckgo.v1.GeckgoService/StartDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeckgoServiceServer).StartDay(ctx, req.(*StartDayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeckgoService_StartBreak_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartBreakRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeckgoServiceServer).StartBreak(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geckgo.v1.GeckgoService/StartBreak",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeckgoServiceServer).StartBreak(ctx, req.(*StartBreakRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeckgoService_EndBreak_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndBreakRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeckgoServiceServer).EndBreak(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geckgo.v1.GeckgoService/EndBreak",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeckgoServiceServer).EndBreak(ctx, req.(*EndBreakRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeckgoService_EndDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndDayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeckgoServiceServer).EndDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geckgo.v1.GeckgoService/EndDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeckgoServiceServer).EndDay(ctx, req.(*EndDayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GeckgoService_ServiceDesc is the grpc.ServiceDesc for GeckgoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GeckgoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "geckgo.v1.GeckgoService",
	HandlerType: (*GeckgoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddList",
			Handler:    _GeckgoService_AddList_Handler,
		},
		{
			MethodName: "GetLists",
			Handler:    _GeckgoService_GetLists_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _GeckgoService_GetList_Handler,
		},
		{
			MethodName: "UpdateList",
			Handler:    _GeckgoService_UpdateList_Handler,
		},
		{
			MethodName: "DeleteList",
			Handler:    _GeckgoService_DeleteList_Handler,
		},
		{
			MethodName: "AddDay",
			Handler:    _GeckgoService_AddDay_Handler,
		},
		{
			MethodName: "GetDays",
			Handler:    _GeckgoService_GetDays_Handler,
		},
		{
			MethodName: "GetDay",
			Handler:    _GeckgoService_GetDay_Handler,
		},
		{
			MethodName: "UpdateDay",
			Handler:    _GeckgoService_UpdateDay_Handler,
		},
		{
			MethodName: "DeleteDay",
			Handler:    _GeckgoService_DeleteDay_Handler,
		},
		{
			MethodName: "StartDay",
			Handler:    _GeckgoService_StartDay_Handler,
		},
		{
			MethodName: "StartBreak",
			Handler:    _GeckgoService_StartBreak_Handler,
		},
		{
			MethodName: "EndBreak",
			Handler:    _GeckgoService_EndBreak_Handler,
		},
		{
			MethodName: "EndDay",
			Handler:    _GeckgoService_EndDay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "geckgo/v1/geckgo.proto",
}