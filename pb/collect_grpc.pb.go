// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// CollectServiceClient is the client API for CollectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectServiceClient interface {
	// 作者
	Author(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*AuthorResponse, error)
	HotAuthors(ctx context.Context, in *EmptyPost, opts ...grpc.CallOption) (*HotAuthorsResponse, error)
	AuthorPagination(ctx context.Context, in *PaginationPost, opts ...grpc.CallOption) (*AuthorPaginationResponse, error)
	CreateAuthor(ctx context.Context, in *CreateAuthorPost, opts ...grpc.CallOption) (*AuthorResponse, error)
	UpdateAuthor(ctx context.Context, in *UpdateAuthorPost, opts ...grpc.CallOption) (*AuthorResponse, error)
	ChangeAuthorStatus(ctx context.Context, in *ChangeStatusPost, opts ...grpc.CallOption) (*AuthorResponse, error)
	// 作品
	Work(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*WorkResponse, error)
	RecommendWorks(ctx context.Context, in *RecommendWorksPost, opts ...grpc.CallOption) (*RecommendWorksResponse, error)
	WorkPagination(ctx context.Context, in *PaginationPost, opts ...grpc.CallOption) (*WorkPaginationResponse, error)
	CreateWork(ctx context.Context, in *CreateWorkPost, opts ...grpc.CallOption) (*WorkResponse, error)
	UpdateWork(ctx context.Context, in *UpdateWorkPost, opts ...grpc.CallOption) (*WorkResponse, error)
	ChangeWorkStatus(ctx context.Context, in *ChangeStatusPost, opts ...grpc.CallOption) (*WorkResponse, error)
	// 盲盒
	BlindBox(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*BlindBoxResponse, error)
	BlindPaginaton(ctx context.Context, in *PaginationPost, opts ...grpc.CallOption) (*BlindBoxPaginationResponse, error)
	CreateBlindBox(ctx context.Context, in *CreateBlindBoxPost, opts ...grpc.CallOption) (*BlindBoxResponse, error)
	UpdateBlindBox(ctx context.Context, in *UpdateBlindBoxPost, opts ...grpc.CallOption) (*BlindBoxResponse, error)
	StartBlindBox(ctx context.Context, in *ChangeStatusPost, opts ...grpc.CallOption) (*BlindBoxResponse, error)
}

type collectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectServiceClient(cc grpc.ClientConnInterface) CollectServiceClient {
	return &collectServiceClient{cc}
}

func (c *collectServiceClient) Author(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*AuthorResponse, error) {
	out := new(AuthorResponse)
	err := c.cc.Invoke(ctx, "/collect.CollectService/Author", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) HotAuthors(ctx context.Context, in *EmptyPost, opts ...grpc.CallOption) (*HotAuthorsResponse, error) {
	out := new(HotAuthorsResponse)
	err := c.cc.Invoke(ctx, "/collect.CollectService/HotAuthors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) AuthorPagination(ctx context.Context, in *PaginationPost, opts ...grpc.CallOption) (*AuthorPaginationResponse, error) {
	out := new(AuthorPaginationResponse)
	err := c.cc.Invoke(ctx, "/collect.CollectService/AuthorPagination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) CreateAuthor(ctx context.Context, in *CreateAuthorPost, opts ...grpc.CallOption) (*AuthorResponse, error) {
	out := new(AuthorResponse)
	err := c.cc.Invoke(ctx, "/collect.CollectService/CreateAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) UpdateAuthor(ctx context.Context, in *UpdateAuthorPost, opts ...grpc.CallOption) (*AuthorResponse, error) {
	out := new(AuthorResponse)
	err := c.cc.Invoke(ctx, "/collect.CollectService/UpdateAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) ChangeAuthorStatus(ctx context.Context, in *ChangeStatusPost, opts ...grpc.CallOption) (*AuthorResponse, error) {
	out := new(AuthorResponse)
	err := c.cc.Invoke(ctx, "/collect.CollectService/ChangeAuthorStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) Work(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*WorkResponse, error) {
	out := new(WorkResponse)
	err := c.cc.Invoke(ctx, "/collect.CollectService/Work", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) RecommendWorks(ctx context.Context, in *RecommendWorksPost, opts ...grpc.CallOption) (*RecommendWorksResponse, error) {
	out := new(RecommendWorksResponse)
	err := c.cc.Invoke(ctx, "/collect.CollectService/RecommendWorks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) WorkPagination(ctx context.Context, in *PaginationPost, opts ...grpc.CallOption) (*WorkPaginationResponse, error) {
	out := new(WorkPaginationResponse)
	err := c.cc.Invoke(ctx, "/collect.CollectService/WorkPagination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) CreateWork(ctx context.Context, in *CreateWorkPost, opts ...grpc.CallOption) (*WorkResponse, error) {
	out := new(WorkResponse)
	err := c.cc.Invoke(ctx, "/collect.CollectService/CreateWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) UpdateWork(ctx context.Context, in *UpdateWorkPost, opts ...grpc.CallOption) (*WorkResponse, error) {
	out := new(WorkResponse)
	err := c.cc.Invoke(ctx, "/collect.CollectService/UpdateWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) ChangeWorkStatus(ctx context.Context, in *ChangeStatusPost, opts ...grpc.CallOption) (*WorkResponse, error) {
	out := new(WorkResponse)
	err := c.cc.Invoke(ctx, "/collect.CollectService/ChangeWorkStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) BlindBox(ctx context.Context, in *InfoPost, opts ...grpc.CallOption) (*BlindBoxResponse, error) {
	out := new(BlindBoxResponse)
	err := c.cc.Invoke(ctx, "/collect.CollectService/BlindBox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) BlindPaginaton(ctx context.Context, in *PaginationPost, opts ...grpc.CallOption) (*BlindBoxPaginationResponse, error) {
	out := new(BlindBoxPaginationResponse)
	err := c.cc.Invoke(ctx, "/collect.CollectService/BlindPaginaton", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) CreateBlindBox(ctx context.Context, in *CreateBlindBoxPost, opts ...grpc.CallOption) (*BlindBoxResponse, error) {
	out := new(BlindBoxResponse)
	err := c.cc.Invoke(ctx, "/collect.CollectService/CreateBlindBox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) UpdateBlindBox(ctx context.Context, in *UpdateBlindBoxPost, opts ...grpc.CallOption) (*BlindBoxResponse, error) {
	out := new(BlindBoxResponse)
	err := c.cc.Invoke(ctx, "/collect.CollectService/UpdateBlindBox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) StartBlindBox(ctx context.Context, in *ChangeStatusPost, opts ...grpc.CallOption) (*BlindBoxResponse, error) {
	out := new(BlindBoxResponse)
	err := c.cc.Invoke(ctx, "/collect.CollectService/StartBlindBox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectServiceServer is the server API for CollectService service.
// All implementations must embed UnimplementedCollectServiceServer
// for forward compatibility
type CollectServiceServer interface {
	// 作者
	Author(context.Context, *InfoPost) (*AuthorResponse, error)
	HotAuthors(context.Context, *EmptyPost) (*HotAuthorsResponse, error)
	AuthorPagination(context.Context, *PaginationPost) (*AuthorPaginationResponse, error)
	CreateAuthor(context.Context, *CreateAuthorPost) (*AuthorResponse, error)
	UpdateAuthor(context.Context, *UpdateAuthorPost) (*AuthorResponse, error)
	ChangeAuthorStatus(context.Context, *ChangeStatusPost) (*AuthorResponse, error)
	// 作品
	Work(context.Context, *InfoPost) (*WorkResponse, error)
	RecommendWorks(context.Context, *RecommendWorksPost) (*RecommendWorksResponse, error)
	WorkPagination(context.Context, *PaginationPost) (*WorkPaginationResponse, error)
	CreateWork(context.Context, *CreateWorkPost) (*WorkResponse, error)
	UpdateWork(context.Context, *UpdateWorkPost) (*WorkResponse, error)
	ChangeWorkStatus(context.Context, *ChangeStatusPost) (*WorkResponse, error)
	// 盲盒
	BlindBox(context.Context, *InfoPost) (*BlindBoxResponse, error)
	BlindPaginaton(context.Context, *PaginationPost) (*BlindBoxPaginationResponse, error)
	CreateBlindBox(context.Context, *CreateBlindBoxPost) (*BlindBoxResponse, error)
	UpdateBlindBox(context.Context, *UpdateBlindBoxPost) (*BlindBoxResponse, error)
	StartBlindBox(context.Context, *ChangeStatusPost) (*BlindBoxResponse, error)
	mustEmbedUnimplementedCollectServiceServer()
}

// UnimplementedCollectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCollectServiceServer struct {
}

func (UnimplementedCollectServiceServer) Author(context.Context, *InfoPost) (*AuthorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Author not implemented")
}
func (UnimplementedCollectServiceServer) HotAuthors(context.Context, *EmptyPost) (*HotAuthorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HotAuthors not implemented")
}
func (UnimplementedCollectServiceServer) AuthorPagination(context.Context, *PaginationPost) (*AuthorPaginationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorPagination not implemented")
}
func (UnimplementedCollectServiceServer) CreateAuthor(context.Context, *CreateAuthorPost) (*AuthorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuthor not implemented")
}
func (UnimplementedCollectServiceServer) UpdateAuthor(context.Context, *UpdateAuthorPost) (*AuthorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthor not implemented")
}
func (UnimplementedCollectServiceServer) ChangeAuthorStatus(context.Context, *ChangeStatusPost) (*AuthorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAuthorStatus not implemented")
}
func (UnimplementedCollectServiceServer) Work(context.Context, *InfoPost) (*WorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Work not implemented")
}
func (UnimplementedCollectServiceServer) RecommendWorks(context.Context, *RecommendWorksPost) (*RecommendWorksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecommendWorks not implemented")
}
func (UnimplementedCollectServiceServer) WorkPagination(context.Context, *PaginationPost) (*WorkPaginationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WorkPagination not implemented")
}
func (UnimplementedCollectServiceServer) CreateWork(context.Context, *CreateWorkPost) (*WorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWork not implemented")
}
func (UnimplementedCollectServiceServer) UpdateWork(context.Context, *UpdateWorkPost) (*WorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWork not implemented")
}
func (UnimplementedCollectServiceServer) ChangeWorkStatus(context.Context, *ChangeStatusPost) (*WorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeWorkStatus not implemented")
}
func (UnimplementedCollectServiceServer) BlindBox(context.Context, *InfoPost) (*BlindBoxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlindBox not implemented")
}
func (UnimplementedCollectServiceServer) BlindPaginaton(context.Context, *PaginationPost) (*BlindBoxPaginationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlindPaginaton not implemented")
}
func (UnimplementedCollectServiceServer) CreateBlindBox(context.Context, *CreateBlindBoxPost) (*BlindBoxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlindBox not implemented")
}
func (UnimplementedCollectServiceServer) UpdateBlindBox(context.Context, *UpdateBlindBoxPost) (*BlindBoxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlindBox not implemented")
}
func (UnimplementedCollectServiceServer) StartBlindBox(context.Context, *ChangeStatusPost) (*BlindBoxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartBlindBox not implemented")
}
func (UnimplementedCollectServiceServer) mustEmbedUnimplementedCollectServiceServer() {}

// UnsafeCollectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectServiceServer will
// result in compilation errors.
type UnsafeCollectServiceServer interface {
	mustEmbedUnimplementedCollectServiceServer()
}

func RegisterCollectServiceServer(s grpc.ServiceRegistrar, srv CollectServiceServer) {
	s.RegisterService(&CollectService_ServiceDesc, srv)
}

func _CollectService_Author_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).Author(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collect.CollectService/Author",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).Author(ctx, req.(*InfoPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_HotAuthors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).HotAuthors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collect.CollectService/HotAuthors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).HotAuthors(ctx, req.(*EmptyPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_AuthorPagination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaginationPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).AuthorPagination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collect.CollectService/AuthorPagination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).AuthorPagination(ctx, req.(*PaginationPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_CreateAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthorPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).CreateAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collect.CollectService/CreateAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).CreateAuthor(ctx, req.(*CreateAuthorPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_UpdateAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthorPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).UpdateAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collect.CollectService/UpdateAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).UpdateAuthor(ctx, req.(*UpdateAuthorPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_ChangeAuthorStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStatusPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).ChangeAuthorStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collect.CollectService/ChangeAuthorStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).ChangeAuthorStatus(ctx, req.(*ChangeStatusPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_Work_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).Work(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collect.CollectService/Work",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).Work(ctx, req.(*InfoPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_RecommendWorks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendWorksPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).RecommendWorks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collect.CollectService/RecommendWorks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).RecommendWorks(ctx, req.(*RecommendWorksPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_WorkPagination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaginationPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).WorkPagination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collect.CollectService/WorkPagination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).WorkPagination(ctx, req.(*PaginationPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_CreateWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).CreateWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collect.CollectService/CreateWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).CreateWork(ctx, req.(*CreateWorkPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_UpdateWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).UpdateWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collect.CollectService/UpdateWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).UpdateWork(ctx, req.(*UpdateWorkPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_ChangeWorkStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStatusPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).ChangeWorkStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collect.CollectService/ChangeWorkStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).ChangeWorkStatus(ctx, req.(*ChangeStatusPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_BlindBox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).BlindBox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collect.CollectService/BlindBox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).BlindBox(ctx, req.(*InfoPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_BlindPaginaton_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaginationPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).BlindPaginaton(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collect.CollectService/BlindPaginaton",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).BlindPaginaton(ctx, req.(*PaginationPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_CreateBlindBox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlindBoxPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).CreateBlindBox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collect.CollectService/CreateBlindBox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).CreateBlindBox(ctx, req.(*CreateBlindBoxPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_UpdateBlindBox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBlindBoxPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).UpdateBlindBox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collect.CollectService/UpdateBlindBox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).UpdateBlindBox(ctx, req.(*UpdateBlindBoxPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_StartBlindBox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStatusPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).StartBlindBox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collect.CollectService/StartBlindBox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).StartBlindBox(ctx, req.(*ChangeStatusPost))
	}
	return interceptor(ctx, in, info, handler)
}

// CollectService_ServiceDesc is the grpc.ServiceDesc for CollectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CollectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "collect.CollectService",
	HandlerType: (*CollectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Author",
			Handler:    _CollectService_Author_Handler,
		},
		{
			MethodName: "HotAuthors",
			Handler:    _CollectService_HotAuthors_Handler,
		},
		{
			MethodName: "AuthorPagination",
			Handler:    _CollectService_AuthorPagination_Handler,
		},
		{
			MethodName: "CreateAuthor",
			Handler:    _CollectService_CreateAuthor_Handler,
		},
		{
			MethodName: "UpdateAuthor",
			Handler:    _CollectService_UpdateAuthor_Handler,
		},
		{
			MethodName: "ChangeAuthorStatus",
			Handler:    _CollectService_ChangeAuthorStatus_Handler,
		},
		{
			MethodName: "Work",
			Handler:    _CollectService_Work_Handler,
		},
		{
			MethodName: "RecommendWorks",
			Handler:    _CollectService_RecommendWorks_Handler,
		},
		{
			MethodName: "WorkPagination",
			Handler:    _CollectService_WorkPagination_Handler,
		},
		{
			MethodName: "CreateWork",
			Handler:    _CollectService_CreateWork_Handler,
		},
		{
			MethodName: "UpdateWork",
			Handler:    _CollectService_UpdateWork_Handler,
		},
		{
			MethodName: "ChangeWorkStatus",
			Handler:    _CollectService_ChangeWorkStatus_Handler,
		},
		{
			MethodName: "BlindBox",
			Handler:    _CollectService_BlindBox_Handler,
		},
		{
			MethodName: "BlindPaginaton",
			Handler:    _CollectService_BlindPaginaton_Handler,
		},
		{
			MethodName: "CreateBlindBox",
			Handler:    _CollectService_CreateBlindBox_Handler,
		},
		{
			MethodName: "UpdateBlindBox",
			Handler:    _CollectService_UpdateBlindBox_Handler,
		},
		{
			MethodName: "StartBlindBox",
			Handler:    _CollectService_StartBlindBox_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collect/collect.proto",
}