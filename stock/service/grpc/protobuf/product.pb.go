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

type RequestRead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestRead) Reset() {
	*x = RequestRead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestRead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestRead) ProtoMessage() {}

func (x *RequestRead) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestRead.ProtoReflect.Descriptor instead.
func (*RequestRead) Descriptor() ([]byte, []int) {
	return file_protobuf_product_proto_rawDescGZIP(), []int{3}
}

func (x *RequestRead) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReplyRead struct {
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

func (x *ReplyRead) Reset() {
	*x = ReplyRead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyRead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyRead) ProtoMessage() {}

func (x *ReplyRead) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyRead.ProtoReflect.Descriptor instead.
func (*ReplyRead) Descriptor() ([]byte, []int) {
	return file_protobuf_product_proto_rawDescGZIP(), []int{4}
}

func (x *ReplyRead) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ReplyRead) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *ReplyRead) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ReplyRead) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ReplyRead) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *ReplyRead) GetCostPerUnit() int64 {
	if x != nil {
		return x.CostPerUnit
	}
	return 0
}

func (x *ReplyRead) GetEXPDate() int64 {
	if x != nil {
		return x.EXPDate
	}
	return 0
}

func (x *ReplyRead) GetImportDate() int64 {
	if x != nil {
		return x.ImportDate
	}
	return 0
}

func (x *ReplyRead) GetSupplier() string {
	if x != nil {
		return x.Supplier
	}
	return ""
}

func (x *ReplyRead) GetTotalCost() int64 {
	if x != nil {
		return x.TotalCost
	}
	return 0
}

func (x *ReplyRead) GetTotalAmount() int64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *ReplyRead) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ReplyArrRead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply []*ReplyRead `protobuf:"bytes,1,rep,name=Reply,proto3" json:"Reply,omitempty"`
}

func (x *ReplyArrRead) Reset() {
	*x = ReplyArrRead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyArrRead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyArrRead) ProtoMessage() {}

func (x *ReplyArrRead) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyArrRead.ProtoReflect.Descriptor instead.
func (*ReplyArrRead) Descriptor() ([]byte, []int) {
	return file_protobuf_product_proto_rawDescGZIP(), []int{5}
}

func (x *ReplyArrRead) GetReply() []*ReplyRead {
	if x != nil {
		return x.Reply
	}
	return nil
}

type RequestName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemName string `protobuf:"bytes,1,opt,name=itemName,proto3" json:"itemName,omitempty"`
	PerPage  int64  `protobuf:"varint,2,opt,name=perPage,proto3" json:"perPage,omitempty"`
	Page     int64  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *RequestName) Reset() {
	*x = RequestName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestName) ProtoMessage() {}

func (x *RequestName) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestName.ProtoReflect.Descriptor instead.
func (*RequestName) Descriptor() ([]byte, []int) {
	return file_protobuf_product_proto_rawDescGZIP(), []int{6}
}

func (x *RequestName) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *RequestName) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *RequestName) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type RequestCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	PerPage  int64  `protobuf:"varint,2,opt,name=perPage,proto3" json:"perPage,omitempty"`
	Page     int64  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *RequestCategory) Reset() {
	*x = RequestCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_product_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCategory) ProtoMessage() {}

func (x *RequestCategory) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_product_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCategory.ProtoReflect.Descriptor instead.
func (*RequestCategory) Descriptor() ([]byte, []int) {
	return file_protobuf_product_proto_rawDescGZIP(), []int{7}
}

func (x *RequestCategory) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *RequestCategory) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *RequestCategory) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
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
	0x69, 0x65, 0x6e, 0x74, 0x22, 0x1d, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xcf, 0x02, 0x0a, 0x09, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x61,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x73, 0x74, 0x50, 0x65, 0x72,
	0x55, 0x6e, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x43, 0x6f, 0x73, 0x74,
	0x50, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x58, 0x50, 0x44, 0x61,
	0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x45, 0x58, 0x50, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x35, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x41, 0x72,
	0x72, 0x52, 0x65, 0x61, 0x64, 0x12, 0x25, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x52, 0x65, 0x61, 0x64, 0x52, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x57, 0x0a, 0x0b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x5b, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x32, 0x50, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x1a, 0x14,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x32, 0x3c, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x12, 0x2f, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x11,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x61,
	0x64, 0x1a, 0x0f, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65,
	0x61, 0x64, 0x32, 0x47, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x12, 0x36, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x41, 0x72, 0x72, 0x52, 0x65, 0x61, 0x64, 0x32, 0x53, 0x0a, 0x11, 0x52,
	0x65, 0x61, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x12, 0x3e, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x15, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x12, 0x2e, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x41, 0x72, 0x72, 0x52, 0x65, 0x61, 0x64,
	0x42, 0x07, 0x5a, 0x05, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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

var file_protobuf_product_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protobuf_product_proto_goTypes = []interface{}{
	(*ReplyFromStock)(nil),  // 0: cart.ReplyFromStock
	(*CalculateCost)(nil),   // 1: cart.CalculateCost
	(*RequestToStock)(nil),  // 2: cart.RequestToStock
	(*RequestRead)(nil),     // 3: cart.RequestRead
	(*ReplyRead)(nil),       // 4: cart.ReplyRead
	(*ReplyArrRead)(nil),    // 5: cart.ReplyArrRead
	(*RequestName)(nil),     // 6: cart.RequestName
	(*RequestCategory)(nil), // 7: cart.RequestCategory
}
var file_protobuf_product_proto_depIdxs = []int32{
	1, // 0: cart.ReplyFromStock.calculateCost:type_name -> cart.CalculateCost
	4, // 1: cart.ReplyArrRead.Reply:type_name -> cart.ReplyRead
	2, // 2: cart.SendIngredients.SendIngredients:input_type -> cart.RequestToStock
	3, // 3: cart.ReadStock.ReadStock:input_type -> cart.RequestRead
	6, // 4: cart.ReadNameStock.ReadNameStock:input_type -> cart.RequestName
	7, // 5: cart.ReadCategoryStock.ReadCategoryStock:input_type -> cart.RequestCategory
	0, // 6: cart.SendIngredients.SendIngredients:output_type -> cart.ReplyFromStock
	4, // 7: cart.ReadStock.ReadStock:output_type -> cart.ReplyRead
	5, // 8: cart.ReadNameStock.ReadNameStock:output_type -> cart.ReplyArrRead
	5, // 9: cart.ReadCategoryStock.ReadCategoryStock:output_type -> cart.ReplyArrRead
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
		file_protobuf_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestRead); i {
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
		file_protobuf_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyRead); i {
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
		file_protobuf_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyArrRead); i {
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
		file_protobuf_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestName); i {
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
		file_protobuf_product_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCategory); i {
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
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   4,
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

// ReadStockClient is the client API for ReadStock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReadStockClient interface {
	ReadStock(ctx context.Context, in *RequestRead, opts ...grpc.CallOption) (*ReplyRead, error)
}

type readStockClient struct {
	cc grpc.ClientConnInterface
}

func NewReadStockClient(cc grpc.ClientConnInterface) ReadStockClient {
	return &readStockClient{cc}
}

func (c *readStockClient) ReadStock(ctx context.Context, in *RequestRead, opts ...grpc.CallOption) (*ReplyRead, error) {
	out := new(ReplyRead)
	err := c.cc.Invoke(ctx, "/cart.ReadStock/ReadStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReadStockServer is the server API for ReadStock service.
type ReadStockServer interface {
	ReadStock(context.Context, *RequestRead) (*ReplyRead, error)
}

// UnimplementedReadStockServer can be embedded to have forward compatible implementations.
type UnimplementedReadStockServer struct {
}

func (*UnimplementedReadStockServer) ReadStock(context.Context, *RequestRead) (*ReplyRead, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadStock not implemented")
}

func RegisterReadStockServer(s *grpc.Server, srv ReadStockServer) {
	s.RegisterService(&_ReadStock_serviceDesc, srv)
}

func _ReadStock_ReadStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRead)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReadStockServer).ReadStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.ReadStock/ReadStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReadStockServer).ReadStock(ctx, req.(*RequestRead))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReadStock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cart.ReadStock",
	HandlerType: (*ReadStockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadStock",
			Handler:    _ReadStock_ReadStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/product.proto",
}

// ReadNameStockClient is the client API for ReadNameStock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReadNameStockClient interface {
	ReadNameStock(ctx context.Context, in *RequestName, opts ...grpc.CallOption) (*ReplyArrRead, error)
}

type readNameStockClient struct {
	cc grpc.ClientConnInterface
}

func NewReadNameStockClient(cc grpc.ClientConnInterface) ReadNameStockClient {
	return &readNameStockClient{cc}
}

func (c *readNameStockClient) ReadNameStock(ctx context.Context, in *RequestName, opts ...grpc.CallOption) (*ReplyArrRead, error) {
	out := new(ReplyArrRead)
	err := c.cc.Invoke(ctx, "/cart.ReadNameStock/ReadNameStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReadNameStockServer is the server API for ReadNameStock service.
type ReadNameStockServer interface {
	ReadNameStock(context.Context, *RequestName) (*ReplyArrRead, error)
}

// UnimplementedReadNameStockServer can be embedded to have forward compatible implementations.
type UnimplementedReadNameStockServer struct {
}

func (*UnimplementedReadNameStockServer) ReadNameStock(context.Context, *RequestName) (*ReplyArrRead, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadNameStock not implemented")
}

func RegisterReadNameStockServer(s *grpc.Server, srv ReadNameStockServer) {
	s.RegisterService(&_ReadNameStock_serviceDesc, srv)
}

func _ReadNameStock_ReadNameStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReadNameStockServer).ReadNameStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.ReadNameStock/ReadNameStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReadNameStockServer).ReadNameStock(ctx, req.(*RequestName))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReadNameStock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cart.ReadNameStock",
	HandlerType: (*ReadNameStockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadNameStock",
			Handler:    _ReadNameStock_ReadNameStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/product.proto",
}

// ReadCategoryStockClient is the client API for ReadCategoryStock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReadCategoryStockClient interface {
	ReadCategoryStock(ctx context.Context, in *RequestCategory, opts ...grpc.CallOption) (*ReplyArrRead, error)
}

type readCategoryStockClient struct {
	cc grpc.ClientConnInterface
}

func NewReadCategoryStockClient(cc grpc.ClientConnInterface) ReadCategoryStockClient {
	return &readCategoryStockClient{cc}
}

func (c *readCategoryStockClient) ReadCategoryStock(ctx context.Context, in *RequestCategory, opts ...grpc.CallOption) (*ReplyArrRead, error) {
	out := new(ReplyArrRead)
	err := c.cc.Invoke(ctx, "/cart.ReadCategoryStock/ReadCategoryStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReadCategoryStockServer is the server API for ReadCategoryStock service.
type ReadCategoryStockServer interface {
	ReadCategoryStock(context.Context, *RequestCategory) (*ReplyArrRead, error)
}

// UnimplementedReadCategoryStockServer can be embedded to have forward compatible implementations.
type UnimplementedReadCategoryStockServer struct {
}

func (*UnimplementedReadCategoryStockServer) ReadCategoryStock(context.Context, *RequestCategory) (*ReplyArrRead, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCategoryStock not implemented")
}

func RegisterReadCategoryStockServer(s *grpc.Server, srv ReadCategoryStockServer) {
	s.RegisterService(&_ReadCategoryStock_serviceDesc, srv)
}

func _ReadCategoryStock_ReadCategoryStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReadCategoryStockServer).ReadCategoryStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.ReadCategoryStock/ReadCategoryStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReadCategoryStockServer).ReadCategoryStock(ctx, req.(*RequestCategory))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReadCategoryStock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cart.ReadCategoryStock",
	HandlerType: (*ReadCategoryStockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadCategoryStock",
			Handler:    _ReadCategoryStock_ReadCategoryStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/product.proto",
}
