// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: report/report.proto

package report

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

type ReportReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report []*ReportReplyStruct `protobuf:"bytes,1,rep,name=Report,proto3" json:"Report,omitempty"`
}

func (x *ReportReply) Reset() {
	*x = ReportReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportReply) ProtoMessage() {}

func (x *ReportReply) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportReply.ProtoReflect.Descriptor instead.
func (*ReportReply) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{0}
}

func (x *ReportReply) GetReport() []*ReportReplyStruct {
	if x != nil {
		return x.Report
	}
	return nil
}

type ReportReplyStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ItemName    string `protobuf:"bytes,2,opt,name=ItemName,proto3" json:"ItemName,omitempty"`
	Category    string `protobuf:"bytes,3,opt,name=Category,proto3" json:"Category,omitempty"`
	Amount      int64  `protobuf:"varint,4,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Unit        string `protobuf:"bytes,5,opt,name=Unit,proto3" json:"Unit,omitempty"`
	CostPerUnit int64  `protobuf:"varint,6,opt,name=CostPerUnit,proto3" json:"CostPerUnit,omitempty"`
	EXPDate     int64  `protobuf:"varint,7,opt,name=EXPDate,proto3" json:"EXPDate,omitempty"`
	ImportDate  int64  `protobuf:"varint,8,opt,name=ImportDate,proto3" json:"ImportDate,omitempty"`
	Supplier    string `protobuf:"bytes,9,opt,name=Supplier,proto3" json:"Supplier,omitempty"`
	TotalCost   int64  `protobuf:"varint,10,opt,name=TotalCost,proto3" json:"TotalCost,omitempty"`
	TotalAmount int64  `protobuf:"varint,11,opt,name=TotalAmount,proto3" json:"TotalAmount,omitempty"`
	Status      string `protobuf:"bytes,12,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *ReportReplyStruct) Reset() {
	*x = ReportReplyStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportReplyStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportReplyStruct) ProtoMessage() {}

func (x *ReportReplyStruct) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportReplyStruct.ProtoReflect.Descriptor instead.
func (*ReportReplyStruct) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{1}
}

func (x *ReportReplyStruct) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ReportReplyStruct) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *ReportReplyStruct) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ReportReplyStruct) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ReportReplyStruct) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *ReportReplyStruct) GetCostPerUnit() int64 {
	if x != nil {
		return x.CostPerUnit
	}
	return 0
}

func (x *ReportReplyStruct) GetEXPDate() int64 {
	if x != nil {
		return x.EXPDate
	}
	return 0
}

func (x *ReportReplyStruct) GetImportDate() int64 {
	if x != nil {
		return x.ImportDate
	}
	return 0
}

func (x *ReportReplyStruct) GetSupplier() string {
	if x != nil {
		return x.Supplier
	}
	return ""
}

func (x *ReportReplyStruct) GetTotalCost() int64 {
	if x != nil {
		return x.TotalCost
	}
	return 0
}

func (x *ReportReplyStruct) GetTotalAmount() int64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *ReportReplyStruct) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request string `protobuf:"bytes,1,opt,name=Request,proto3" json:"Request,omitempty"`
	Field   string `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
	Order   string `protobuf:"bytes,3,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *ReportRequest) Reset() {
	*x = ReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRequest) ProtoMessage() {}

func (x *ReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRequest.ProtoReflect.Descriptor instead.
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{2}
}

func (x *ReportRequest) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *ReportRequest) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *ReportRequest) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

var File_report_report_proto protoreflect.FileDescriptor

var file_report_report_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x22, 0xd7, 0x02, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f,
	0x73, 0x74, 0x50, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x43, 0x6f, 0x73, 0x74, 0x50, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x45, 0x58, 0x50, 0x44, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x45,
	0x58, 0x50, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x55, 0x0a, 0x0d, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x32, 0x46, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x6f, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x31, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x6f, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_report_report_proto_rawDescOnce sync.Once
	file_report_report_proto_rawDescData = file_report_report_proto_rawDesc
)

func file_report_report_proto_rawDescGZIP() []byte {
	file_report_report_proto_rawDescOnce.Do(func() {
		file_report_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_report_report_proto_rawDescData)
	})
	return file_report_report_proto_rawDescData
}

var file_report_report_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_report_report_proto_goTypes = []interface{}{
	(*ReportReply)(nil),       // 0: ReportReply
	(*ReportReplyStruct)(nil), // 1: ReportReplyStruct
	(*ReportRequest)(nil),     // 2: ReportRequest
}
var file_report_report_proto_depIdxs = []int32{
	1, // 0: ReportReply.Report:type_name -> ReportReplyStruct
	2, // 1: SendReportToStock.SendReportToStock:input_type -> ReportRequest
	0, // 2: SendReportToStock.SendReportToStock:output_type -> ReportReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_report_report_proto_init() }
func file_report_report_proto_init() {
	if File_report_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_report_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportReply); i {
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
		file_report_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportReplyStruct); i {
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
		file_report_report_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportRequest); i {
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
			RawDescriptor: file_report_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_report_report_proto_goTypes,
		DependencyIndexes: file_report_report_proto_depIdxs,
		MessageInfos:      file_report_report_proto_msgTypes,
	}.Build()
	File_report_report_proto = out.File
	file_report_report_proto_rawDesc = nil
	file_report_report_proto_goTypes = nil
	file_report_report_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SendReportToStockClient is the client API for SendReportToStock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SendReportToStockClient interface {
	SendReportToStock(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportReply, error)
}

type sendReportToStockClient struct {
	cc grpc.ClientConnInterface
}

func NewSendReportToStockClient(cc grpc.ClientConnInterface) SendReportToStockClient {
	return &sendReportToStockClient{cc}
}

func (c *sendReportToStockClient) SendReportToStock(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportReply, error) {
	out := new(ReportReply)
	err := c.cc.Invoke(ctx, "/SendReportToStock/SendReportToStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendReportToStockServer is the server API for SendReportToStock service.
type SendReportToStockServer interface {
	SendReportToStock(context.Context, *ReportRequest) (*ReportReply, error)
}

// UnimplementedSendReportToStockServer can be embedded to have forward compatible implementations.
type UnimplementedSendReportToStockServer struct {
}

func (*UnimplementedSendReportToStockServer) SendReportToStock(context.Context, *ReportRequest) (*ReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendReportToStock not implemented")
}

func RegisterSendReportToStockServer(s *grpc.Server, srv SendReportToStockServer) {
	s.RegisterService(&_SendReportToStock_serviceDesc, srv)
}

func _SendReportToStock_SendReportToStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendReportToStockServer).SendReportToStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SendReportToStock/SendReportToStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendReportToStockServer).SendReportToStock(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SendReportToStock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SendReportToStock",
	HandlerType: (*SendReportToStockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendReportToStock",
			Handler:    _SendReportToStock_SendReportToStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "report/report.proto",
}
