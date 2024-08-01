// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: task.proto

package task

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

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	// Process
	Start(ctx context.Context, in *StartArgs, opts ...grpc.CallOption) (*StartResp, error)
	Dump(ctx context.Context, in *DumpArgs, opts ...grpc.CallOption) (*DumpResp, error)
	Restore(ctx context.Context, in *RestoreArgs, opts ...grpc.CallOption) (*RestoreResp, error)
	Query(ctx context.Context, in *QueryArgs, opts ...grpc.CallOption) (*QueryResp, error)
	// Containerd
	ContainerdDump(ctx context.Context, in *ContainerdDumpArgs, opts ...grpc.CallOption) (*ContainerdDumpResp, error)
	ContainerdRestore(ctx context.Context, in *ContainerdRestoreArgs, opts ...grpc.CallOption) (*ContainerdRestoreResp, error)
	ContainerdQuery(ctx context.Context, in *ContainerdQueryArgs, opts ...grpc.CallOption) (*ContainerdQueryResp, error)
	ContainerdRootfsDump(ctx context.Context, in *ContainerdRootfsDumpArgs, opts ...grpc.CallOption) (*ContainerdRootfsDumpResp, error)
	ContainerdRootfsRestore(ctx context.Context, in *ContainerdRootfsRestoreArgs, opts ...grpc.CallOption) (*ContainerdRootfsRestoreResp, error)
	// Runc
	RuncDump(ctx context.Context, in *RuncDumpArgs, opts ...grpc.CallOption) (*RuncDumpResp, error)
	RuncRestore(ctx context.Context, in *RuncRestoreArgs, opts ...grpc.CallOption) (*RuncRestoreResp, error)
	RuncQuery(ctx context.Context, in *RuncQueryArgs, opts ...grpc.CallOption) (*RuncQueryResp, error)
	RuncGetPausePid(ctx context.Context, in *RuncGetPausePidArgs, opts ...grpc.CallOption) (*RuncGetPausePidResp, error)
	// CRIO
	CRIORootfsDump(ctx context.Context, in *CRIORootfsDumpArgs, opts ...grpc.CallOption) (*CRIORootfsDumpResp, error)
	CRIOImagePush(ctx context.Context, in *CRIOImagePushArgs, opts ...grpc.CallOption) (*CRIOImagePushResp, error)
	// Streaming
	LogStreaming(ctx context.Context, opts ...grpc.CallOption) (TaskService_LogStreamingClient, error)
	ProcessStateStreaming(ctx context.Context, in *ProcessStateStreamingArgs, opts ...grpc.CallOption) (TaskService_ProcessStateStreamingClient, error)
	// Health
	DetailedHealthCheck(ctx context.Context, in *DetailedHealthCheckRequest, opts ...grpc.CallOption) (*DetailedHealthCheckResponse, error)
	// Config
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) Start(ctx context.Context, in *StartArgs, opts ...grpc.CallOption) (*StartResp, error) {
	out := new(StartResp)
	err := c.cc.Invoke(ctx, "/cedana.services.task.TaskService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) Dump(ctx context.Context, in *DumpArgs, opts ...grpc.CallOption) (*DumpResp, error) {
	out := new(DumpResp)
	err := c.cc.Invoke(ctx, "/cedana.services.task.TaskService/Dump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) Restore(ctx context.Context, in *RestoreArgs, opts ...grpc.CallOption) (*RestoreResp, error) {
	out := new(RestoreResp)
	err := c.cc.Invoke(ctx, "/cedana.services.task.TaskService/Restore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) Query(ctx context.Context, in *QueryArgs, opts ...grpc.CallOption) (*QueryResp, error) {
	out := new(QueryResp)
	err := c.cc.Invoke(ctx, "/cedana.services.task.TaskService/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ContainerdDump(ctx context.Context, in *ContainerdDumpArgs, opts ...grpc.CallOption) (*ContainerdDumpResp, error) {
	out := new(ContainerdDumpResp)
	err := c.cc.Invoke(ctx, "/cedana.services.task.TaskService/ContainerdDump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ContainerdRestore(ctx context.Context, in *ContainerdRestoreArgs, opts ...grpc.CallOption) (*ContainerdRestoreResp, error) {
	out := new(ContainerdRestoreResp)
	err := c.cc.Invoke(ctx, "/cedana.services.task.TaskService/ContainerdRestore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ContainerdQuery(ctx context.Context, in *ContainerdQueryArgs, opts ...grpc.CallOption) (*ContainerdQueryResp, error) {
	out := new(ContainerdQueryResp)
	err := c.cc.Invoke(ctx, "/cedana.services.task.TaskService/ContainerdQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ContainerdRootfsDump(ctx context.Context, in *ContainerdRootfsDumpArgs, opts ...grpc.CallOption) (*ContainerdRootfsDumpResp, error) {
	out := new(ContainerdRootfsDumpResp)
	err := c.cc.Invoke(ctx, "/cedana.services.task.TaskService/ContainerdRootfsDump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ContainerdRootfsRestore(ctx context.Context, in *ContainerdRootfsRestoreArgs, opts ...grpc.CallOption) (*ContainerdRootfsRestoreResp, error) {
	out := new(ContainerdRootfsRestoreResp)
	err := c.cc.Invoke(ctx, "/cedana.services.task.TaskService/ContainerdRootfsRestore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) RuncDump(ctx context.Context, in *RuncDumpArgs, opts ...grpc.CallOption) (*RuncDumpResp, error) {
	out := new(RuncDumpResp)
	err := c.cc.Invoke(ctx, "/cedana.services.task.TaskService/RuncDump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) RuncRestore(ctx context.Context, in *RuncRestoreArgs, opts ...grpc.CallOption) (*RuncRestoreResp, error) {
	out := new(RuncRestoreResp)
	err := c.cc.Invoke(ctx, "/cedana.services.task.TaskService/RuncRestore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) RuncQuery(ctx context.Context, in *RuncQueryArgs, opts ...grpc.CallOption) (*RuncQueryResp, error) {
	out := new(RuncQueryResp)
	err := c.cc.Invoke(ctx, "/cedana.services.task.TaskService/RuncQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) RuncGetPausePid(ctx context.Context, in *RuncGetPausePidArgs, opts ...grpc.CallOption) (*RuncGetPausePidResp, error) {
	out := new(RuncGetPausePidResp)
	err := c.cc.Invoke(ctx, "/cedana.services.task.TaskService/RuncGetPausePid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) CRIORootfsDump(ctx context.Context, in *CRIORootfsDumpArgs, opts ...grpc.CallOption) (*CRIORootfsDumpResp, error) {
	out := new(CRIORootfsDumpResp)
	err := c.cc.Invoke(ctx, "/cedana.services.task.TaskService/CRIORootfsDump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) CRIOImagePush(ctx context.Context, in *CRIOImagePushArgs, opts ...grpc.CallOption) (*CRIOImagePushResp, error) {
	out := new(CRIOImagePushResp)
	err := c.cc.Invoke(ctx, "/cedana.services.task.TaskService/CRIOImagePush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) LogStreaming(ctx context.Context, opts ...grpc.CallOption) (TaskService_LogStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &TaskService_ServiceDesc.Streams[0], "/cedana.services.task.TaskService/LogStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskServiceLogStreamingClient{stream}
	return x, nil
}

type TaskService_LogStreamingClient interface {
	Send(*LogStreamingResp) error
	Recv() (*LogStreamingArgs, error)
	grpc.ClientStream
}

type taskServiceLogStreamingClient struct {
	grpc.ClientStream
}

func (x *taskServiceLogStreamingClient) Send(m *LogStreamingResp) error {
	return x.ClientStream.SendMsg(m)
}

func (x *taskServiceLogStreamingClient) Recv() (*LogStreamingArgs, error) {
	m := new(LogStreamingArgs)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskServiceClient) ProcessStateStreaming(ctx context.Context, in *ProcessStateStreamingArgs, opts ...grpc.CallOption) (TaskService_ProcessStateStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &TaskService_ServiceDesc.Streams[1], "/cedana.services.task.TaskService/ProcessStateStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskServiceProcessStateStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TaskService_ProcessStateStreamingClient interface {
	Recv() (*ProcessState, error)
	grpc.ClientStream
}

type taskServiceProcessStateStreamingClient struct {
	grpc.ClientStream
}

func (x *taskServiceProcessStateStreamingClient) Recv() (*ProcessState, error) {
	m := new(ProcessState)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskServiceClient) DetailedHealthCheck(ctx context.Context, in *DetailedHealthCheckRequest, opts ...grpc.CallOption) (*DetailedHealthCheckResponse, error) {
	out := new(DetailedHealthCheckResponse)
	err := c.cc.Invoke(ctx, "/cedana.services.task.TaskService/DetailedHealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := c.cc.Invoke(ctx, "/cedana.services.task.TaskService/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	// Process
	Start(context.Context, *StartArgs) (*StartResp, error)
	Dump(context.Context, *DumpArgs) (*DumpResp, error)
	Restore(context.Context, *RestoreArgs) (*RestoreResp, error)
	Query(context.Context, *QueryArgs) (*QueryResp, error)
	// Containerd
	ContainerdDump(context.Context, *ContainerdDumpArgs) (*ContainerdDumpResp, error)
	ContainerdRestore(context.Context, *ContainerdRestoreArgs) (*ContainerdRestoreResp, error)
	ContainerdQuery(context.Context, *ContainerdQueryArgs) (*ContainerdQueryResp, error)
	ContainerdRootfsDump(context.Context, *ContainerdRootfsDumpArgs) (*ContainerdRootfsDumpResp, error)
	ContainerdRootfsRestore(context.Context, *ContainerdRootfsRestoreArgs) (*ContainerdRootfsRestoreResp, error)
	// Runc
	RuncDump(context.Context, *RuncDumpArgs) (*RuncDumpResp, error)
	RuncRestore(context.Context, *RuncRestoreArgs) (*RuncRestoreResp, error)
	RuncQuery(context.Context, *RuncQueryArgs) (*RuncQueryResp, error)
	RuncGetPausePid(context.Context, *RuncGetPausePidArgs) (*RuncGetPausePidResp, error)
	// CRIO
	CRIORootfsDump(context.Context, *CRIORootfsDumpArgs) (*CRIORootfsDumpResp, error)
	CRIOImagePush(context.Context, *CRIOImagePushArgs) (*CRIOImagePushResp, error)
	// Streaming
	LogStreaming(TaskService_LogStreamingServer) error
	ProcessStateStreaming(*ProcessStateStreamingArgs, TaskService_ProcessStateStreamingServer) error
	// Health
	DetailedHealthCheck(context.Context, *DetailedHealthCheckRequest) (*DetailedHealthCheckResponse, error)
	// Config
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (UnimplementedTaskServiceServer) Start(context.Context, *StartArgs) (*StartResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedTaskServiceServer) Dump(context.Context, *DumpArgs) (*DumpResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dump not implemented")
}
func (UnimplementedTaskServiceServer) Restore(context.Context, *RestoreArgs) (*RestoreResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedTaskServiceServer) Query(context.Context, *QueryArgs) (*QueryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedTaskServiceServer) ContainerdDump(context.Context, *ContainerdDumpArgs) (*ContainerdDumpResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContainerdDump not implemented")
}
func (UnimplementedTaskServiceServer) ContainerdRestore(context.Context, *ContainerdRestoreArgs) (*ContainerdRestoreResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContainerdRestore not implemented")
}
func (UnimplementedTaskServiceServer) ContainerdQuery(context.Context, *ContainerdQueryArgs) (*ContainerdQueryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContainerdQuery not implemented")
}
func (UnimplementedTaskServiceServer) ContainerdRootfsDump(context.Context, *ContainerdRootfsDumpArgs) (*ContainerdRootfsDumpResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContainerdRootfsDump not implemented")
}
func (UnimplementedTaskServiceServer) ContainerdRootfsRestore(context.Context, *ContainerdRootfsRestoreArgs) (*ContainerdRootfsRestoreResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContainerdRootfsRestore not implemented")
}
func (UnimplementedTaskServiceServer) RuncDump(context.Context, *RuncDumpArgs) (*RuncDumpResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RuncDump not implemented")
}
func (UnimplementedTaskServiceServer) RuncRestore(context.Context, *RuncRestoreArgs) (*RuncRestoreResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RuncRestore not implemented")
}
func (UnimplementedTaskServiceServer) RuncQuery(context.Context, *RuncQueryArgs) (*RuncQueryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RuncQuery not implemented")
}
func (UnimplementedTaskServiceServer) RuncGetPausePid(context.Context, *RuncGetPausePidArgs) (*RuncGetPausePidResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RuncGetPausePid not implemented")
}
func (UnimplementedTaskServiceServer) CRIORootfsDump(context.Context, *CRIORootfsDumpArgs) (*CRIORootfsDumpResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CRIORootfsDump not implemented")
}
func (UnimplementedTaskServiceServer) CRIOImagePush(context.Context, *CRIOImagePushArgs) (*CRIOImagePushResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CRIOImagePush not implemented")
}
func (UnimplementedTaskServiceServer) LogStreaming(TaskService_LogStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method LogStreaming not implemented")
}
func (UnimplementedTaskServiceServer) ProcessStateStreaming(*ProcessStateStreamingArgs, TaskService_ProcessStateStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessStateStreaming not implemented")
}
func (UnimplementedTaskServiceServer) DetailedHealthCheck(context.Context, *DetailedHealthCheckRequest) (*DetailedHealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailedHealthCheck not implemented")
}
func (UnimplementedTaskServiceServer) GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.task.TaskService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Start(ctx, req.(*StartArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_Dump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DumpArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Dump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.task.TaskService/Dump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Dump(ctx, req.(*DumpArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.task.TaskService/Restore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Restore(ctx, req.(*RestoreArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.task.TaskService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Query(ctx, req.(*QueryArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ContainerdDump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerdDumpArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ContainerdDump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.task.TaskService/ContainerdDump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ContainerdDump(ctx, req.(*ContainerdDumpArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ContainerdRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerdRestoreArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ContainerdRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.task.TaskService/ContainerdRestore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ContainerdRestore(ctx, req.(*ContainerdRestoreArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ContainerdQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerdQueryArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ContainerdQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.task.TaskService/ContainerdQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ContainerdQuery(ctx, req.(*ContainerdQueryArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ContainerdRootfsDump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerdRootfsDumpArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ContainerdRootfsDump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.task.TaskService/ContainerdRootfsDump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ContainerdRootfsDump(ctx, req.(*ContainerdRootfsDumpArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ContainerdRootfsRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerdRootfsRestoreArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ContainerdRootfsRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.task.TaskService/ContainerdRootfsRestore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ContainerdRootfsRestore(ctx, req.(*ContainerdRootfsRestoreArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_RuncDump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuncDumpArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).RuncDump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.task.TaskService/RuncDump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).RuncDump(ctx, req.(*RuncDumpArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_RuncRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuncRestoreArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).RuncRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.task.TaskService/RuncRestore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).RuncRestore(ctx, req.(*RuncRestoreArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_RuncQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuncQueryArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).RuncQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.task.TaskService/RuncQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).RuncQuery(ctx, req.(*RuncQueryArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_RuncGetPausePid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuncGetPausePidArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).RuncGetPausePid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.task.TaskService/RuncGetPausePid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).RuncGetPausePid(ctx, req.(*RuncGetPausePidArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_CRIORootfsDump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CRIORootfsDumpArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).CRIORootfsDump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.task.TaskService/CRIORootfsDump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).CRIORootfsDump(ctx, req.(*CRIORootfsDumpArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_CRIOImagePush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CRIOImagePushArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).CRIOImagePush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.task.TaskService/CRIOImagePush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).CRIOImagePush(ctx, req.(*CRIOImagePushArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_LogStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TaskServiceServer).LogStreaming(&taskServiceLogStreamingServer{stream})
}

type TaskService_LogStreamingServer interface {
	Send(*LogStreamingArgs) error
	Recv() (*LogStreamingResp, error)
	grpc.ServerStream
}

type taskServiceLogStreamingServer struct {
	grpc.ServerStream
}

func (x *taskServiceLogStreamingServer) Send(m *LogStreamingArgs) error {
	return x.ServerStream.SendMsg(m)
}

func (x *taskServiceLogStreamingServer) Recv() (*LogStreamingResp, error) {
	m := new(LogStreamingResp)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TaskService_ProcessStateStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProcessStateStreamingArgs)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskServiceServer).ProcessStateStreaming(m, &taskServiceProcessStateStreamingServer{stream})
}

type TaskService_ProcessStateStreamingServer interface {
	Send(*ProcessState) error
	grpc.ServerStream
}

type taskServiceProcessStateStreamingServer struct {
	grpc.ServerStream
}

func (x *taskServiceProcessStateStreamingServer) Send(m *ProcessState) error {
	return x.ServerStream.SendMsg(m)
}

func _TaskService_DetailedHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailedHealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).DetailedHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.task.TaskService/DetailedHealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).DetailedHealthCheck(ctx, req.(*DetailedHealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedana.services.task.TaskService/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cedana.services.task.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _TaskService_Start_Handler,
		},
		{
			MethodName: "Dump",
			Handler:    _TaskService_Dump_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _TaskService_Restore_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _TaskService_Query_Handler,
		},
		{
			MethodName: "ContainerdDump",
			Handler:    _TaskService_ContainerdDump_Handler,
		},
		{
			MethodName: "ContainerdRestore",
			Handler:    _TaskService_ContainerdRestore_Handler,
		},
		{
			MethodName: "ContainerdQuery",
			Handler:    _TaskService_ContainerdQuery_Handler,
		},
		{
			MethodName: "ContainerdRootfsDump",
			Handler:    _TaskService_ContainerdRootfsDump_Handler,
		},
		{
			MethodName: "ContainerdRootfsRestore",
			Handler:    _TaskService_ContainerdRootfsRestore_Handler,
		},
		{
			MethodName: "RuncDump",
			Handler:    _TaskService_RuncDump_Handler,
		},
		{
			MethodName: "RuncRestore",
			Handler:    _TaskService_RuncRestore_Handler,
		},
		{
			MethodName: "RuncQuery",
			Handler:    _TaskService_RuncQuery_Handler,
		},
		{
			MethodName: "RuncGetPausePid",
			Handler:    _TaskService_RuncGetPausePid_Handler,
		},
		{
			MethodName: "CRIORootfsDump",
			Handler:    _TaskService_CRIORootfsDump_Handler,
		},
		{
			MethodName: "CRIOImagePush",
			Handler:    _TaskService_CRIOImagePush_Handler,
		},
		{
			MethodName: "DetailedHealthCheck",
			Handler:    _TaskService_DetailedHealthCheck_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _TaskService_GetConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LogStreaming",
			Handler:       _TaskService_LogStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ProcessStateStreaming",
			Handler:       _TaskService_ProcessStateStreaming_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "task.proto",
}
