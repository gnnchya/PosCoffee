// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: protobuf/product.proto

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

type ReplyFromStock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stock         bool             `protobuf:"varint,1,opt,name=Stock,proto3" json:"Stock,omitempty"`
	CalculateCost []*CalculateCost `protobuf:"bytes,2,rep,name=calculateCost,proto3" json:"calculateCost,omitempty"`
	Err           string           `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *ReplyFromStock) Reset() {
	*x = ReplyFromStock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyFromStock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyFromStock) ProtoMessage() {}

func (x *ReplyFromStock) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyFromStock.ProtoReflect.Descriptor instead.
func (*ReplyFromStock) Descriptor() ([]byte, []int) {
	return file_protobuf_product_proto_rawDescGZIP(), []int{0}
}

func (x *ReplyFromStock) GetStock() bool {
	if x != nil {
		return x.Stock
	}
	return false
}

func (x *ReplyFromStock) GetCalculateCost() []*CalculateCost {
	if x != nil {
		return x.CalculateCost
	}
	return nil
}

func (x *ReplyFromStock) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type CalculateCost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemName    string `protobuf:"bytes,1,opt,name=ItemName,proto3" json:"ItemName,omitempty"`
	CostPerUnit int64  `protobuf:"varint,2,opt,name=CostPerUnit,proto3" json:"CostPerUnit,omitempty"`
}

func (x *CalculateCost) Reset() {
	*x = CalculateCost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateCost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateCost) ProtoMessage() {}

func (x *CalculateCost) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateCost.ProtoReflect.Descriptor instead.
func (*CalculateCost) Descriptor() ([]byte, []int) {
	return file_protobuf_product_proto_rawDescGZIP(), []int{1}
}

func (x *CalculateCost) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *CalculateCost) GetCostPerUnit() int64 {
	if x != nil {
		return x.CostPerUnit
	}
	return 0
}

type RequestToStock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ingredient []string `protobuf:"bytes,1,rep,name=ingredient,proto3" json:"ingredient,omitempty"`
}

func (x *RequestToStock) Reset() {
	*x = RequestToStock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestToStock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestToStock) ProtoMessage() {}

func (x *RequestToStock) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestToStock.ProtoReflect.Descriptor instead.
func (*RequestToStock) Descriptor() ([]byte, []int) {
	return file_protobuf_product_proto_rawDescGZIP(), []int{2}
}

func (x *RequestToStock) GetIngredient() []string {
	if x != nil {
		return x.Ingredient
	}
	return nil
}

var File_protobuf_product_proto protoreflect.FileDescriptor

var file_protobuf_product_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x73,
	0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x39, 0x0a, 0x0d, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x73, 0x74, 0x52, 0x0d, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x65, 0x72, 0x72, 0x22, 0x4d, 0x0a, 0x0d, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x73, 0x74, 0x50, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x43, 0x6f, 0x73, 0x74, 0x50, 0x65, 0x72, 0x55, 0x6e,
	0x69, 0x74, 0x22, 0x30, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64,
	0x69, 0x65, 0x6e, 0x74, 0x32, 0x50, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x1a, 0x14, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x46, 0x72, 0x6f,
	0x6d, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_product_proto_rawDescOnce sync.Once
	file_protobuf_product_proto_rawDescData = file_protobuf_product_proto_rawDesc
)

func file_protobuf_product_proto_rawDescGZIP() []byte {
	file_protobuf_product_proto_rawDescOnce.Do(func() {
		file_protobuf_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_product_proto_rawDescData)
	})
	return file_protobuf_product_proto_rawDescData
}

var file_protobuf_product_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protobuf_product_proto_goTypes = []interface{}{
	(*ReplyFromStock)(nil), // 0: cart.ReplyFromStock
	(*CalculateCost)(nil),  // 1: cart.CalculateCost
	(*RequestToStock)(nil), // 2: cart.RequestToStock
}
var file_protobuf_product_proto_depIdxs = []int32{
	1, // 0: cart.ReplyFromStock.calculateCost:type_name -> cart.CalculateCost
	2, // 1: cart.SendIngredients.SendIngredients:input_type -> cart.RequestToStock
	0, // 2: cart.SendIngredients.SendIngredients:output_type -> cart.ReplyFromStock
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protobuf_product_proto_init() }
func file_protobuf_product_proto_init() {
	if File_protobuf_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyFromStock); i {
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
		file_protobuf_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateCost); i {
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
		file_protobuf_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestToStock); i {
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
			RawDescriptor: file_protobuf_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_product_proto_goTypes,
		DependencyIndexes: file_protobuf_product_proto_depIdxs,
		MessageInfos:      file_protobuf_product_proto_msgTypes,
	}.Build()
	File_protobuf_product_proto = out.File
	file_protobuf_product_proto_rawDesc = nil
	file_protobuf_product_proto_goTypes = nil
	file_protobuf_product_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SendIngredientsClient is the client API for SendIngredients service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SendIngredientsClient interface {
	SendIngredients(ctx context.Context, in *RequestToStock, opts ...grpc.CallOption) (*ReplyFromStock, error)
}

type sendIngredientsClient struct {
	cc grpc.ClientConnInterface
}

func NewSendIngredientsClient(cc grpc.ClientConnInterface) SendIngredientsClient {
	return &sendIngredientsClient{cc}
}

func (c *sendIngredientsClient) SendIngredients(ctx context.Context, in *RequestToStock, opts ...grpc.CallOption) (*ReplyFromStock, error) {
	out := new(ReplyFromStock)
	err := c.cc.Invoke(ctx, "/cart.SendIngredients/SendIngredients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendIngredientsServer is the server API for SendIngredients service.
type SendIngredientsServer interface {
	SendIngredients(context.Context, *RequestToStock) (*ReplyFromStock, error)
}

// UnimplementedSendIngredientsServer can be embedded to have forward compatible implementations.
type UnimplementedSendIngredientsServer struct {
}

func (*UnimplementedSendIngredientsServer) SendIngredients(context.Context, *RequestToStock) (*ReplyFromStock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendIngredients not implemented")
}

func RegisterSendIngredientsServer(s *grpc.Server, srv SendIngredientsServer) {
	s.RegisterService(&_SendIngredients_serviceDesc, srv)
}

func _SendIngredients_SendIngredients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestToStock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendIngredientsServer).SendIngredients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.SendIngredients/SendIngredients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendIngredientsServer).SendIngredients(ctx, req.(*RequestToStock))
	}
	return interceptor(ctx, in, info, handler)
}

var _SendIngredients_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cart.SendIngredients",
	HandlerType: (*SendIngredientsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendIngredients",
			Handler:    _SendIngredients_SendIngredients_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/product.proto",
}