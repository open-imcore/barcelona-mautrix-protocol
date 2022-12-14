// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: ipc/v1/v1.proto

package mautrix_nosip_protobuf

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

// BarcelonaClient is the client API for Barcelona service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BarcelonaClient interface {
	StartupSyncHook(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StartupSyncHookResponse, error)
	RequestHistory(ctx context.Context, in *HistoryQuery, opts ...grpc.CallOption) (*MessageList, error)
	RequestChats(ctx context.Context, in *ChatQuery, opts ...grpc.CallOption) (*ChatIDList, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	GetContacts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ContactList, error)
	GetContact(ctx context.Context, in *GUID, opts ...grpc.CallOption) (*Contact, error)
	GetChat(ctx context.Context, in *GUID, opts ...grpc.CallOption) (*ChatInfo, error)
	GetChatAvatar(ctx context.Context, in *GUID, opts ...grpc.CallOption) (*Attachment, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendResult, error)
	SendReadReceipt(ctx context.Context, in *SendReadReceiptRequest, opts ...grpc.CallOption) (*Empty, error)
	SetTyping(ctx context.Context, in *SetTypingRequest, opts ...grpc.CallOption) (*Empty, error)
	ResolveIdentifier(ctx context.Context, in *ResolveIdentifierRequest, opts ...grpc.CallOption) (*ResolveIdentifierResult, error)
	PrepareDM(ctx context.Context, in *PrepareDMRequest, opts ...grpc.CallOption) (*EmptyResult, error)
	LogStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Barcelona_LogStreamClient, error)
	MessageStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Barcelona_MessageStreamClient, error)
	ReceiptStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Barcelona_ReceiptStreamClient, error)
	TypingStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Barcelona_TypingStreamClient, error)
	ChatStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Barcelona_ChatStreamClient, error)
	ContactStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Barcelona_ContactStreamClient, error)
	MessageStatusStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Barcelona_MessageStatusStreamClient, error)
	BridgeStatusStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Barcelona_BridgeStatusStreamClient, error)
}

type barcelonaClient struct {
	cc grpc.ClientConnInterface
}

func NewBarcelonaClient(cc grpc.ClientConnInterface) BarcelonaClient {
	return &barcelonaClient{cc}
}

func (c *barcelonaClient) StartupSyncHook(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StartupSyncHookResponse, error) {
	out := new(StartupSyncHookResponse)
	err := c.cc.Invoke(ctx, "/Barcelona/StartupSyncHook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barcelonaClient) RequestHistory(ctx context.Context, in *HistoryQuery, opts ...grpc.CallOption) (*MessageList, error) {
	out := new(MessageList)
	err := c.cc.Invoke(ctx, "/Barcelona/RequestHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barcelonaClient) RequestChats(ctx context.Context, in *ChatQuery, opts ...grpc.CallOption) (*ChatIDList, error) {
	out := new(ChatIDList)
	err := c.cc.Invoke(ctx, "/Barcelona/RequestChats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barcelonaClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Barcelona/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barcelonaClient) GetContacts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ContactList, error) {
	out := new(ContactList)
	err := c.cc.Invoke(ctx, "/Barcelona/GetContacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barcelonaClient) GetContact(ctx context.Context, in *GUID, opts ...grpc.CallOption) (*Contact, error) {
	out := new(Contact)
	err := c.cc.Invoke(ctx, "/Barcelona/GetContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barcelonaClient) GetChat(ctx context.Context, in *GUID, opts ...grpc.CallOption) (*ChatInfo, error) {
	out := new(ChatInfo)
	err := c.cc.Invoke(ctx, "/Barcelona/GetChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barcelonaClient) GetChatAvatar(ctx context.Context, in *GUID, opts ...grpc.CallOption) (*Attachment, error) {
	out := new(Attachment)
	err := c.cc.Invoke(ctx, "/Barcelona/GetChatAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barcelonaClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendResult, error) {
	out := new(SendResult)
	err := c.cc.Invoke(ctx, "/Barcelona/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barcelonaClient) SendReadReceipt(ctx context.Context, in *SendReadReceiptRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Barcelona/SendReadReceipt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barcelonaClient) SetTyping(ctx context.Context, in *SetTypingRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Barcelona/SetTyping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barcelonaClient) ResolveIdentifier(ctx context.Context, in *ResolveIdentifierRequest, opts ...grpc.CallOption) (*ResolveIdentifierResult, error) {
	out := new(ResolveIdentifierResult)
	err := c.cc.Invoke(ctx, "/Barcelona/ResolveIdentifier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barcelonaClient) PrepareDM(ctx context.Context, in *PrepareDMRequest, opts ...grpc.CallOption) (*EmptyResult, error) {
	out := new(EmptyResult)
	err := c.cc.Invoke(ctx, "/Barcelona/PrepareDM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barcelonaClient) LogStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Barcelona_LogStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Barcelona_ServiceDesc.Streams[0], "/Barcelona/LogStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &barcelonaLogStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Barcelona_LogStreamClient interface {
	Recv() (*LogLine, error)
	grpc.ClientStream
}

type barcelonaLogStreamClient struct {
	grpc.ClientStream
}

func (x *barcelonaLogStreamClient) Recv() (*LogLine, error) {
	m := new(LogLine)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *barcelonaClient) MessageStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Barcelona_MessageStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Barcelona_ServiceDesc.Streams[1], "/Barcelona/MessageStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &barcelonaMessageStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Barcelona_MessageStreamClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type barcelonaMessageStreamClient struct {
	grpc.ClientStream
}

func (x *barcelonaMessageStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *barcelonaClient) ReceiptStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Barcelona_ReceiptStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Barcelona_ServiceDesc.Streams[2], "/Barcelona/ReceiptStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &barcelonaReceiptStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Barcelona_ReceiptStreamClient interface {
	Recv() (*ReadReceipt, error)
	grpc.ClientStream
}

type barcelonaReceiptStreamClient struct {
	grpc.ClientStream
}

func (x *barcelonaReceiptStreamClient) Recv() (*ReadReceipt, error) {
	m := new(ReadReceipt)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *barcelonaClient) TypingStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Barcelona_TypingStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Barcelona_ServiceDesc.Streams[3], "/Barcelona/TypingStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &barcelonaTypingStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Barcelona_TypingStreamClient interface {
	Recv() (*TypingNotification, error)
	grpc.ClientStream
}

type barcelonaTypingStreamClient struct {
	grpc.ClientStream
}

func (x *barcelonaTypingStreamClient) Recv() (*TypingNotification, error) {
	m := new(TypingNotification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *barcelonaClient) ChatStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Barcelona_ChatStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Barcelona_ServiceDesc.Streams[4], "/Barcelona/ChatStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &barcelonaChatStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Barcelona_ChatStreamClient interface {
	Recv() (*ChatInfo, error)
	grpc.ClientStream
}

type barcelonaChatStreamClient struct {
	grpc.ClientStream
}

func (x *barcelonaChatStreamClient) Recv() (*ChatInfo, error) {
	m := new(ChatInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *barcelonaClient) ContactStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Barcelona_ContactStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Barcelona_ServiceDesc.Streams[5], "/Barcelona/ContactStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &barcelonaContactStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Barcelona_ContactStreamClient interface {
	Recv() (*Contact, error)
	grpc.ClientStream
}

type barcelonaContactStreamClient struct {
	grpc.ClientStream
}

func (x *barcelonaContactStreamClient) Recv() (*Contact, error) {
	m := new(Contact)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *barcelonaClient) MessageStatusStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Barcelona_MessageStatusStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Barcelona_ServiceDesc.Streams[6], "/Barcelona/MessageStatusStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &barcelonaMessageStatusStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Barcelona_MessageStatusStreamClient interface {
	Recv() (*SendMessageStatus, error)
	grpc.ClientStream
}

type barcelonaMessageStatusStreamClient struct {
	grpc.ClientStream
}

func (x *barcelonaMessageStatusStreamClient) Recv() (*SendMessageStatus, error) {
	m := new(SendMessageStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *barcelonaClient) BridgeStatusStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Barcelona_BridgeStatusStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Barcelona_ServiceDesc.Streams[7], "/Barcelona/BridgeStatusStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &barcelonaBridgeStatusStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Barcelona_BridgeStatusStreamClient interface {
	Recv() (*BridgeStatus, error)
	grpc.ClientStream
}

type barcelonaBridgeStatusStreamClient struct {
	grpc.ClientStream
}

func (x *barcelonaBridgeStatusStreamClient) Recv() (*BridgeStatus, error) {
	m := new(BridgeStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BarcelonaServer is the server API for Barcelona service.
// All implementations must embed UnimplementedBarcelonaServer
// for forward compatibility
type BarcelonaServer interface {
	StartupSyncHook(context.Context, *Empty) (*StartupSyncHookResponse, error)
	RequestHistory(context.Context, *HistoryQuery) (*MessageList, error)
	RequestChats(context.Context, *ChatQuery) (*ChatIDList, error)
	Ping(context.Context, *Empty) (*Empty, error)
	GetContacts(context.Context, *Empty) (*ContactList, error)
	GetContact(context.Context, *GUID) (*Contact, error)
	GetChat(context.Context, *GUID) (*ChatInfo, error)
	GetChatAvatar(context.Context, *GUID) (*Attachment, error)
	SendMessage(context.Context, *SendMessageRequest) (*SendResult, error)
	SendReadReceipt(context.Context, *SendReadReceiptRequest) (*Empty, error)
	SetTyping(context.Context, *SetTypingRequest) (*Empty, error)
	ResolveIdentifier(context.Context, *ResolveIdentifierRequest) (*ResolveIdentifierResult, error)
	PrepareDM(context.Context, *PrepareDMRequest) (*EmptyResult, error)
	LogStream(*Empty, Barcelona_LogStreamServer) error
	MessageStream(*Empty, Barcelona_MessageStreamServer) error
	ReceiptStream(*Empty, Barcelona_ReceiptStreamServer) error
	TypingStream(*Empty, Barcelona_TypingStreamServer) error
	ChatStream(*Empty, Barcelona_ChatStreamServer) error
	ContactStream(*Empty, Barcelona_ContactStreamServer) error
	MessageStatusStream(*Empty, Barcelona_MessageStatusStreamServer) error
	BridgeStatusStream(*Empty, Barcelona_BridgeStatusStreamServer) error
	mustEmbedUnimplementedBarcelonaServer()
}

// UnimplementedBarcelonaServer must be embedded to have forward compatible implementations.
type UnimplementedBarcelonaServer struct {
}

func (UnimplementedBarcelonaServer) StartupSyncHook(context.Context, *Empty) (*StartupSyncHookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartupSyncHook not implemented")
}
func (UnimplementedBarcelonaServer) RequestHistory(context.Context, *HistoryQuery) (*MessageList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestHistory not implemented")
}
func (UnimplementedBarcelonaServer) RequestChats(context.Context, *ChatQuery) (*ChatIDList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestChats not implemented")
}
func (UnimplementedBarcelonaServer) Ping(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedBarcelonaServer) GetContacts(context.Context, *Empty) (*ContactList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContacts not implemented")
}
func (UnimplementedBarcelonaServer) GetContact(context.Context, *GUID) (*Contact, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContact not implemented")
}
func (UnimplementedBarcelonaServer) GetChat(context.Context, *GUID) (*ChatInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChat not implemented")
}
func (UnimplementedBarcelonaServer) GetChatAvatar(context.Context, *GUID) (*Attachment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatAvatar not implemented")
}
func (UnimplementedBarcelonaServer) SendMessage(context.Context, *SendMessageRequest) (*SendResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedBarcelonaServer) SendReadReceipt(context.Context, *SendReadReceiptRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendReadReceipt not implemented")
}
func (UnimplementedBarcelonaServer) SetTyping(context.Context, *SetTypingRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTyping not implemented")
}
func (UnimplementedBarcelonaServer) ResolveIdentifier(context.Context, *ResolveIdentifierRequest) (*ResolveIdentifierResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveIdentifier not implemented")
}
func (UnimplementedBarcelonaServer) PrepareDM(context.Context, *PrepareDMRequest) (*EmptyResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareDM not implemented")
}
func (UnimplementedBarcelonaServer) LogStream(*Empty, Barcelona_LogStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method LogStream not implemented")
}
func (UnimplementedBarcelonaServer) MessageStream(*Empty, Barcelona_MessageStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method MessageStream not implemented")
}
func (UnimplementedBarcelonaServer) ReceiptStream(*Empty, Barcelona_ReceiptStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiptStream not implemented")
}
func (UnimplementedBarcelonaServer) TypingStream(*Empty, Barcelona_TypingStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method TypingStream not implemented")
}
func (UnimplementedBarcelonaServer) ChatStream(*Empty, Barcelona_ChatStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ChatStream not implemented")
}
func (UnimplementedBarcelonaServer) ContactStream(*Empty, Barcelona_ContactStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ContactStream not implemented")
}
func (UnimplementedBarcelonaServer) MessageStatusStream(*Empty, Barcelona_MessageStatusStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method MessageStatusStream not implemented")
}
func (UnimplementedBarcelonaServer) BridgeStatusStream(*Empty, Barcelona_BridgeStatusStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method BridgeStatusStream not implemented")
}
func (UnimplementedBarcelonaServer) mustEmbedUnimplementedBarcelonaServer() {}

// UnsafeBarcelonaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BarcelonaServer will
// result in compilation errors.
type UnsafeBarcelonaServer interface {
	mustEmbedUnimplementedBarcelonaServer()
}

func RegisterBarcelonaServer(s grpc.ServiceRegistrar, srv BarcelonaServer) {
	s.RegisterService(&Barcelona_ServiceDesc, srv)
}

func _Barcelona_StartupSyncHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarcelonaServer).StartupSyncHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Barcelona/StartupSyncHook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarcelonaServer).StartupSyncHook(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Barcelona_RequestHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistoryQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarcelonaServer).RequestHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Barcelona/RequestHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarcelonaServer).RequestHistory(ctx, req.(*HistoryQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Barcelona_RequestChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarcelonaServer).RequestChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Barcelona/RequestChats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarcelonaServer).RequestChats(ctx, req.(*ChatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Barcelona_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarcelonaServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Barcelona/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarcelonaServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Barcelona_GetContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarcelonaServer).GetContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Barcelona/GetContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarcelonaServer).GetContacts(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Barcelona_GetContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarcelonaServer).GetContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Barcelona/GetContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarcelonaServer).GetContact(ctx, req.(*GUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Barcelona_GetChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarcelonaServer).GetChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Barcelona/GetChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarcelonaServer).GetChat(ctx, req.(*GUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Barcelona_GetChatAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarcelonaServer).GetChatAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Barcelona/GetChatAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarcelonaServer).GetChatAvatar(ctx, req.(*GUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Barcelona_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarcelonaServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Barcelona/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarcelonaServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Barcelona_SendReadReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendReadReceiptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarcelonaServer).SendReadReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Barcelona/SendReadReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarcelonaServer).SendReadReceipt(ctx, req.(*SendReadReceiptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Barcelona_SetTyping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTypingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarcelonaServer).SetTyping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Barcelona/SetTyping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarcelonaServer).SetTyping(ctx, req.(*SetTypingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Barcelona_ResolveIdentifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveIdentifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarcelonaServer).ResolveIdentifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Barcelona/ResolveIdentifier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarcelonaServer).ResolveIdentifier(ctx, req.(*ResolveIdentifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Barcelona_PrepareDM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareDMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarcelonaServer).PrepareDM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Barcelona/PrepareDM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarcelonaServer).PrepareDM(ctx, req.(*PrepareDMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Barcelona_LogStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BarcelonaServer).LogStream(m, &barcelonaLogStreamServer{stream})
}

type Barcelona_LogStreamServer interface {
	Send(*LogLine) error
	grpc.ServerStream
}

type barcelonaLogStreamServer struct {
	grpc.ServerStream
}

func (x *barcelonaLogStreamServer) Send(m *LogLine) error {
	return x.ServerStream.SendMsg(m)
}

func _Barcelona_MessageStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BarcelonaServer).MessageStream(m, &barcelonaMessageStreamServer{stream})
}

type Barcelona_MessageStreamServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type barcelonaMessageStreamServer struct {
	grpc.ServerStream
}

func (x *barcelonaMessageStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Barcelona_ReceiptStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BarcelonaServer).ReceiptStream(m, &barcelonaReceiptStreamServer{stream})
}

type Barcelona_ReceiptStreamServer interface {
	Send(*ReadReceipt) error
	grpc.ServerStream
}

type barcelonaReceiptStreamServer struct {
	grpc.ServerStream
}

func (x *barcelonaReceiptStreamServer) Send(m *ReadReceipt) error {
	return x.ServerStream.SendMsg(m)
}

func _Barcelona_TypingStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BarcelonaServer).TypingStream(m, &barcelonaTypingStreamServer{stream})
}

type Barcelona_TypingStreamServer interface {
	Send(*TypingNotification) error
	grpc.ServerStream
}

type barcelonaTypingStreamServer struct {
	grpc.ServerStream
}

func (x *barcelonaTypingStreamServer) Send(m *TypingNotification) error {
	return x.ServerStream.SendMsg(m)
}

func _Barcelona_ChatStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BarcelonaServer).ChatStream(m, &barcelonaChatStreamServer{stream})
}

type Barcelona_ChatStreamServer interface {
	Send(*ChatInfo) error
	grpc.ServerStream
}

type barcelonaChatStreamServer struct {
	grpc.ServerStream
}

func (x *barcelonaChatStreamServer) Send(m *ChatInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _Barcelona_ContactStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BarcelonaServer).ContactStream(m, &barcelonaContactStreamServer{stream})
}

type Barcelona_ContactStreamServer interface {
	Send(*Contact) error
	grpc.ServerStream
}

type barcelonaContactStreamServer struct {
	grpc.ServerStream
}

func (x *barcelonaContactStreamServer) Send(m *Contact) error {
	return x.ServerStream.SendMsg(m)
}

func _Barcelona_MessageStatusStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BarcelonaServer).MessageStatusStream(m, &barcelonaMessageStatusStreamServer{stream})
}

type Barcelona_MessageStatusStreamServer interface {
	Send(*SendMessageStatus) error
	grpc.ServerStream
}

type barcelonaMessageStatusStreamServer struct {
	grpc.ServerStream
}

func (x *barcelonaMessageStatusStreamServer) Send(m *SendMessageStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _Barcelona_BridgeStatusStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BarcelonaServer).BridgeStatusStream(m, &barcelonaBridgeStatusStreamServer{stream})
}

type Barcelona_BridgeStatusStreamServer interface {
	Send(*BridgeStatus) error
	grpc.ServerStream
}

type barcelonaBridgeStatusStreamServer struct {
	grpc.ServerStream
}

func (x *barcelonaBridgeStatusStreamServer) Send(m *BridgeStatus) error {
	return x.ServerStream.SendMsg(m)
}

// Barcelona_ServiceDesc is the grpc.ServiceDesc for Barcelona service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Barcelona_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Barcelona",
	HandlerType: (*BarcelonaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartupSyncHook",
			Handler:    _Barcelona_StartupSyncHook_Handler,
		},
		{
			MethodName: "RequestHistory",
			Handler:    _Barcelona_RequestHistory_Handler,
		},
		{
			MethodName: "RequestChats",
			Handler:    _Barcelona_RequestChats_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Barcelona_Ping_Handler,
		},
		{
			MethodName: "GetContacts",
			Handler:    _Barcelona_GetContacts_Handler,
		},
		{
			MethodName: "GetContact",
			Handler:    _Barcelona_GetContact_Handler,
		},
		{
			MethodName: "GetChat",
			Handler:    _Barcelona_GetChat_Handler,
		},
		{
			MethodName: "GetChatAvatar",
			Handler:    _Barcelona_GetChatAvatar_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Barcelona_SendMessage_Handler,
		},
		{
			MethodName: "SendReadReceipt",
			Handler:    _Barcelona_SendReadReceipt_Handler,
		},
		{
			MethodName: "SetTyping",
			Handler:    _Barcelona_SetTyping_Handler,
		},
		{
			MethodName: "ResolveIdentifier",
			Handler:    _Barcelona_ResolveIdentifier_Handler,
		},
		{
			MethodName: "PrepareDM",
			Handler:    _Barcelona_PrepareDM_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LogStream",
			Handler:       _Barcelona_LogStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "MessageStream",
			Handler:       _Barcelona_MessageStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReceiptStream",
			Handler:       _Barcelona_ReceiptStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TypingStream",
			Handler:       _Barcelona_TypingStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ChatStream",
			Handler:       _Barcelona_ChatStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ContactStream",
			Handler:       _Barcelona_ContactStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "MessageStatusStream",
			Handler:       _Barcelona_MessageStatusStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BridgeStatusStream",
			Handler:       _Barcelona_BridgeStatusStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ipc/v1/v1.proto",
}

// BridgeClient is the client API for Bridge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BridgeClient interface {
	BarcelonaReady(ctx context.Context, in *BarcelonaStartupInfo, opts ...grpc.CallOption) (*Empty, error)
	TransferRetrievalFailed(ctx context.Context, in *Error, opts ...grpc.CallOption) (*Empty, error)
}

type bridgeClient struct {
	cc grpc.ClientConnInterface
}

func NewBridgeClient(cc grpc.ClientConnInterface) BridgeClient {
	return &bridgeClient{cc}
}

func (c *bridgeClient) BarcelonaReady(ctx context.Context, in *BarcelonaStartupInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Bridge/BarcelonaReady", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeClient) TransferRetrievalFailed(ctx context.Context, in *Error, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Bridge/TransferRetrievalFailed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BridgeServer is the server API for Bridge service.
// All implementations must embed UnimplementedBridgeServer
// for forward compatibility
type BridgeServer interface {
	BarcelonaReady(context.Context, *BarcelonaStartupInfo) (*Empty, error)
	TransferRetrievalFailed(context.Context, *Error) (*Empty, error)
	mustEmbedUnimplementedBridgeServer()
}

// UnimplementedBridgeServer must be embedded to have forward compatible implementations.
type UnimplementedBridgeServer struct {
}

func (UnimplementedBridgeServer) BarcelonaReady(context.Context, *BarcelonaStartupInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BarcelonaReady not implemented")
}
func (UnimplementedBridgeServer) TransferRetrievalFailed(context.Context, *Error) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferRetrievalFailed not implemented")
}
func (UnimplementedBridgeServer) mustEmbedUnimplementedBridgeServer() {}

// UnsafeBridgeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BridgeServer will
// result in compilation errors.
type UnsafeBridgeServer interface {
	mustEmbedUnimplementedBridgeServer()
}

func RegisterBridgeServer(s grpc.ServiceRegistrar, srv BridgeServer) {
	s.RegisterService(&Bridge_ServiceDesc, srv)
}

func _Bridge_BarcelonaReady_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BarcelonaStartupInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServer).BarcelonaReady(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bridge/BarcelonaReady",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServer).BarcelonaReady(ctx, req.(*BarcelonaStartupInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bridge_TransferRetrievalFailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Error)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServer).TransferRetrievalFailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bridge/TransferRetrievalFailed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServer).TransferRetrievalFailed(ctx, req.(*Error))
	}
	return interceptor(ctx, in, info, handler)
}

// Bridge_ServiceDesc is the grpc.ServiceDesc for Bridge service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bridge_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Bridge",
	HandlerType: (*BridgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BarcelonaReady",
			Handler:    _Bridge_BarcelonaReady_Handler,
		},
		{
			MethodName: "TransferRetrievalFailed",
			Handler:    _Bridge_TransferRetrievalFailed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ipc/v1/v1.proto",
}
