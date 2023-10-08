// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: message.proto

package proto

import (
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

// Auth
type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *AuthRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type TokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Freebie   bool   `protobuf:"varint,2,opt,name=freebie,proto3" json:"freebie,omitempty"`
	ExpiredAt int64  `protobuf:"varint,3,opt,name=expiredAt,proto3" json:"expiredAt,omitempty"`
}

func (x *TokenResponse) Reset() {
	*x = TokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenResponse) ProtoMessage() {}

func (x *TokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenResponse.ProtoReflect.Descriptor instead.
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *TokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *TokenResponse) GetFreebie() bool {
	if x != nil {
		return x.Freebie
	}
	return false
}

func (x *TokenResponse) GetExpiredAt() int64 {
	if x != nil {
		return x.ExpiredAt
	}
	return 0
}

// Products Cards
type ProductCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku   uint64 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Image string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *ProductCard) Reset() {
	*x = ProductCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCard) ProtoMessage() {}

func (x *ProductCard) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCard.ProtoReflect.Descriptor instead.
func (*ProductCard) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *ProductCard) GetSku() uint64 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *ProductCard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductCard) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type AddProductsCardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductsCards []*ProductCard `protobuf:"bytes,1,rep,name=productsCards,proto3" json:"productsCards,omitempty"`
}

func (x *AddProductsCardsRequest) Reset() {
	*x = AddProductsCardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProductsCardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProductsCardsRequest) ProtoMessage() {}

func (x *AddProductsCardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProductsCardsRequest.ProtoReflect.Descriptor instead.
func (*AddProductsCardsRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *AddProductsCardsRequest) GetProductsCards() []*ProductCard {
	if x != nil {
		return x.ProductsCards
	}
	return nil
}

type AddProductsCardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Qty int32 `protobuf:"varint,1,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *AddProductsCardsResponse) Reset() {
	*x = AddProductsCardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProductsCardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProductsCardsResponse) ProtoMessage() {}

func (x *AddProductsCardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProductsCardsResponse.ProtoReflect.Descriptor instead.
func (*AddProductsCardsResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *AddProductsCardsResponse) GetQty() int32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

type GetProductsCardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetProductsCardsRequest) Reset() {
	*x = GetProductsCardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductsCardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsCardsRequest) ProtoMessage() {}

func (x *GetProductsCardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsCardsRequest.ProtoReflect.Descriptor instead.
func (*GetProductsCardsRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{6}
}

type GetProductsCardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductsCards []*ProductCard `protobuf:"bytes,1,rep,name=productsCards,proto3" json:"productsCards,omitempty"`
}

func (x *GetProductsCardsResponse) Reset() {
	*x = GetProductsCardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductsCardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsCardsResponse) ProtoMessage() {}

func (x *GetProductsCardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsCardsResponse.ProtoReflect.Descriptor instead.
func (*GetProductsCardsResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{7}
}

func (x *GetProductsCardsResponse) GetProductsCards() []*ProductCard {
	if x != nil {
		return x.ProductsCards
	}
	return nil
}

type GetStocksFromToReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From uint64   `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To   uint64   `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
	Skus []uint64 `protobuf:"varint,3,rep,packed,name=skus,proto3" json:"skus,omitempty"`
}

func (x *GetStocksFromToReq) Reset() {
	*x = GetStocksFromToReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStocksFromToReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStocksFromToReq) ProtoMessage() {}

func (x *GetStocksFromToReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStocksFromToReq.ProtoReflect.Descriptor instead.
func (*GetStocksFromToReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{8}
}

func (x *GetStocksFromToReq) GetFrom() uint64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *GetStocksFromToReq) GetTo() uint64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *GetStocksFromToReq) GetSkus() []uint64 {
	if x != nil {
		return x.Skus
	}
	return nil
}

type Stock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku uint64 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Wh  uint64 `protobuf:"varint,2,opt,name=wh,proto3" json:"wh,omitempty"`
	Qty int32  `protobuf:"varint,3,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *Stock) Reset() {
	*x = Stock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stock) ProtoMessage() {}

func (x *Stock) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stock.ProtoReflect.Descriptor instead.
func (*Stock) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{9}
}

func (x *Stock) GetSku() uint64 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *Stock) GetWh() uint64 {
	if x != nil {
		return x.Wh
	}
	return 0
}

func (x *Stock) GetQty() int32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

type GetStocksFromToResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stocks []*Stock `protobuf:"bytes,1,rep,name=stocks,proto3" json:"stocks,omitempty"`
}

func (x *GetStocksFromToResp) Reset() {
	*x = GetStocksFromToResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStocksFromToResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStocksFromToResp) ProtoMessage() {}

func (x *GetStocksFromToResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStocksFromToResp.ProtoReflect.Descriptor instead.
func (*GetStocksFromToResp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{10}
}

func (x *GetStocksFromToResp) GetStocks() []*Stock {
	if x != nil {
		return x.Stocks
	}
	return nil
}

type DeleteProductCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku uint64 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *DeleteProductCardRequest) Reset() {
	*x = DeleteProductCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProductCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductCardRequest) ProtoMessage() {}

func (x *DeleteProductCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductCardRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductCardRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteProductCardRequest) GetSku() uint64 {
	if x != nil {
		return x.Sku
	}
	return 0
}

type DeleteProductCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteProductCardResponse) Reset() {
	*x = DeleteProductCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProductCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductCardResponse) ProtoMessage() {}

func (x *DeleteProductCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductCardResponse.ProtoReflect.Descriptor instead.
func (*DeleteProductCardResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{12}
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x45,
	0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x5d, 0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x66, 0x72, 0x65, 0x65, 0x62, 0x69, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66,
	0x72, 0x65, 0x65, 0x62, 0x69, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x49, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x61, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22,
	0x52, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0d, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x43, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x22, 0x2c, 0x0a, 0x18, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71, 0x74,
	0x79, 0x22, 0x19, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x53, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x43, 0x61, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x43, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61,
	0x72, 0x64, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x43, 0x61, 0x72, 0x64,
	0x73, 0x22, 0x4c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x46, 0x72,
	0x6f, 0x6d, 0x54, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6b, 0x75, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6b, 0x75, 0x73, 0x22,
	0x3b, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x0e, 0x0a, 0x02, 0x77, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x77, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71, 0x74, 0x79, 0x22, 0x3a, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x52, 0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x2c, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x22, 0x1b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_message_proto_goTypes = []interface{}{
	(*Empty)(nil),                     // 0: main.Empty
	(*AuthRequest)(nil),               // 1: main.AuthRequest
	(*TokenResponse)(nil),             // 2: main.TokenResponse
	(*ProductCard)(nil),               // 3: main.ProductCard
	(*AddProductsCardsRequest)(nil),   // 4: main.AddProductsCardsRequest
	(*AddProductsCardsResponse)(nil),  // 5: main.AddProductsCardsResponse
	(*GetProductsCardsRequest)(nil),   // 6: main.GetProductsCardsRequest
	(*GetProductsCardsResponse)(nil),  // 7: main.GetProductsCardsResponse
	(*GetStocksFromToReq)(nil),        // 8: main.GetStocksFromToReq
	(*Stock)(nil),                     // 9: main.Stock
	(*GetStocksFromToResp)(nil),       // 10: main.GetStocksFromToResp
	(*DeleteProductCardRequest)(nil),  // 11: main.DeleteProductCardRequest
	(*DeleteProductCardResponse)(nil), // 12: main.DeleteProductCardResponse
}
var file_message_proto_depIdxs = []int32{
	3, // 0: main.AddProductsCardsRequest.productsCards:type_name -> main.ProductCard
	3, // 1: main.GetProductsCardsResponse.productsCards:type_name -> main.ProductCard
	9, // 2: main.GetStocksFromToResp.stocks:type_name -> main.Stock
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenResponse); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCard); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddProductsCardsRequest); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddProductsCardsResponse); i {
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
		file_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductsCardsRequest); i {
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
		file_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductsCardsResponse); i {
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
		file_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStocksFromToReq); i {
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
		file_message_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stock); i {
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
		file_message_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStocksFromToResp); i {
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
		file_message_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProductCardRequest); i {
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
		file_message_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProductCardResponse); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
