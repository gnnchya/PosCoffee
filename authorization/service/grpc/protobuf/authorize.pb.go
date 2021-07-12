// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: protobuf/authorize.proto

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestPermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roles      []string `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	Permission string   `protobuf:"bytes,2,opt,name=permission,proto3" json:"permission,omitempty"`
}

func (x *RequestPermission) Reset() {
	*x = RequestPermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_authorize_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPermission) ProtoMessage() {}

func (x *RequestPermission) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_authorize_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPermission.ProtoReflect.Descriptor instead.
func (*RequestPermission) Descriptor() ([]byte, []int) {
	return file_protobuf_authorize_proto_rawDescGZIP(), []int{0}
}

func (x *RequestPermission) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *RequestPermission) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

type ReplyPermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allow bool `protobuf:"varint,1,opt,name=allow,proto3" json:"allow,omitempty"`
}

func (x *ReplyPermission) Reset() {
	*x = ReplyPermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_authorize_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyPermission) ProtoMessage() {}

func (x *ReplyPermission) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_authorize_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyPermission.ProtoReflect.Descriptor instead.
func (*ReplyPermission) Descriptor() ([]byte, []int) {
	return file_protobuf_authorize_proto_rawDescGZIP(), []int{1}
}

func (x *ReplyPermission) GetAllow() bool {
	if x != nil {
		return x.Allow
	}
	return false
}

var File_protobuf_authorize_proto protoreflect.FileDescriptor

var file_protobuf_authorize_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x22, 0x49, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x27, 0x0a, 0x0f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x32, 0x56, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x12, 0x49, 0x0a, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_authorize_proto_rawDescOnce sync.Once
	file_protobuf_authorize_proto_rawDescData = file_protobuf_authorize_proto_rawDesc
)

func file_protobuf_authorize_proto_rawDescGZIP() []byte {
	file_protobuf_authorize_proto_rawDescOnce.Do(func() {
		file_protobuf_authorize_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_authorize_proto_rawDescData)
	})
	return file_protobuf_authorize_proto_rawDescData
}

var file_protobuf_authorize_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobuf_authorize_proto_goTypes = []interface{}{
	(*RequestPermission)(nil), // 0: protobuf.RequestPermission
	(*ReplyPermission)(nil),   // 1: protobuf.ReplyPermission
}
var file_protobuf_authorize_proto_depIdxs = []int32{
	0, // 0: protobuf.Authorize.checkPermission:input_type -> protobuf.RequestPermission
	1, // 1: protobuf.Authorize.checkPermission:output_type -> protobuf.ReplyPermission
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_authorize_proto_init() }
func file_protobuf_authorize_proto_init() {
	if File_protobuf_authorize_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_authorize_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestPermission); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_authorize_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyPermission); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_authorize_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_authorize_proto_goTypes,
		DependencyIndexes: file_protobuf_authorize_proto_depIdxs,
		MessageInfos:      file_protobuf_authorize_proto_msgTypes,
	}.Build()
	File_protobuf_authorize_proto = out.File
	file_protobuf_authorize_proto_rawDesc = nil
	file_protobuf_authorize_proto_goTypes = nil
	file_protobuf_authorize_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthorizeClient is the client API for Authorize service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizeClient interface {
	CheckPermission(ctx context.Context, in *RequestPermission, opts ...grpc.CallOption) (*ReplyPermission, error)
}

type authorizeClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizeClient(cc grpc.ClientConnInterface) AuthorizeClient {
	return &authorizeClient{cc}
}

func (c *authorizeClient) CheckPermission(ctx context.Context, in *RequestPermission, opts ...grpc.CallOption) (*ReplyPermission, error) {
	out := new(ReplyPermission)
	err := c.cc.Invoke(ctx, "/protobuf.Authorize/checkPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizeServer is the server API for Authorize service.
type AuthorizeServer interface {
	CheckPermission(context.Context, *RequestPermission) (*ReplyPermission, error)
}

// UnimplementedAuthorizeServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorizeServer struct {
}

func (*UnimplementedAuthorizeServer) CheckPermission(context.Context, *RequestPermission) (*ReplyPermission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermission not implemented")
}

func RegisterAuthorizeServer(s *grpc.Server, srv AuthorizeServer) {
	s.RegisterService(&_Authorize_serviceDesc, srv)
}

func _Authorize_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPermission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizeServer).CheckPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Authorize/CheckPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizeServer).CheckPermission(ctx, req.(*RequestPermission))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorize_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Authorize",
	HandlerType: (*AuthorizeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "checkPermission",
			Handler:    _Authorize_CheckPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/authorize.proto",
}
