// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: position_service.proto

package pb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type OpenPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol    string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	IsBuyType bool   `protobuf:"varint,2,opt,name=is_buy_type,json=isBuyType,proto3" json:"is_buy_type,omitempty"`
	PriceId   string `protobuf:"bytes,3,opt,name=price_id,json=priceId,proto3" json:"price_id,omitempty"`
	UserId    int64  `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *OpenPositionRequest) Reset() {
	*x = OpenPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenPositionRequest) ProtoMessage() {}

func (x *OpenPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_position_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenPositionRequest.ProtoReflect.Descriptor instead.
func (*OpenPositionRequest) Descriptor() ([]byte, []int) {
	return file_position_service_proto_rawDescGZIP(), []int{0}
}

func (x *OpenPositionRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *OpenPositionRequest) GetIsBuyType() bool {
	if x != nil {
		return x.IsBuyType
	}
	return false
}

func (x *OpenPositionRequest) GetPriceId() string {
	if x != nil {
		return x.PriceId
	}
	return ""
}

func (x *OpenPositionRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type OpenPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionId int64 `protobuf:"varint,1,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
}

func (x *OpenPositionResponse) Reset() {
	*x = OpenPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenPositionResponse) ProtoMessage() {}

func (x *OpenPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_position_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenPositionResponse.ProtoReflect.Descriptor instead.
func (*OpenPositionResponse) Descriptor() ([]byte, []int) {
	return file_position_service_proto_rawDescGZIP(), []int{1}
}

func (x *OpenPositionResponse) GetPositionId() int64 {
	if x != nil {
		return x.PositionId
	}
	return 0
}

type ClosePositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionId int64  `protobuf:"varint,1,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
	PriceId    string `protobuf:"bytes,2,opt,name=price_id,json=priceId,proto3" json:"price_id,omitempty"`
}

func (x *ClosePositionRequest) Reset() {
	*x = ClosePositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClosePositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClosePositionRequest) ProtoMessage() {}

func (x *ClosePositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_position_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClosePositionRequest.ProtoReflect.Descriptor instead.
func (*ClosePositionRequest) Descriptor() ([]byte, []int) {
	return file_position_service_proto_rawDescGZIP(), []int{2}
}

func (x *ClosePositionRequest) GetPositionId() int64 {
	if x != nil {
		return x.PositionId
	}
	return 0
}

func (x *ClosePositionRequest) GetPriceId() string {
	if x != nil {
		return x.PriceId
	}
	return ""
}

type ClosePositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profit float64 `protobuf:"fixed64,1,opt,name=profit,proto3" json:"profit,omitempty"`
}

func (x *ClosePositionResponse) Reset() {
	*x = ClosePositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClosePositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClosePositionResponse) ProtoMessage() {}

func (x *ClosePositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_position_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClosePositionResponse.ProtoReflect.Descriptor instead.
func (*ClosePositionResponse) Descriptor() ([]byte, []int) {
	return file_position_service_proto_rawDescGZIP(), []int{3}
}

func (x *ClosePositionResponse) GetProfit() float64 {
	if x != nil {
		return x.Profit
	}
	return 0
}

type GetOpenPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionId int64 `protobuf:"varint,1,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
}

func (x *GetOpenPositionRequest) Reset() {
	*x = GetOpenPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOpenPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOpenPositionRequest) ProtoMessage() {}

func (x *GetOpenPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_position_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOpenPositionRequest.ProtoReflect.Descriptor instead.
func (*GetOpenPositionRequest) Descriptor() ([]byte, []int) {
	return file_position_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetOpenPositionRequest) GetPositionId() int64 {
	if x != nil {
		return x.PositionId
	}
	return 0
}

type GetOpenPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddPrice   float64               `protobuf:"fixed64,1,opt,name=add_price,json=addPrice,proto3" json:"add_price,omitempty"`
	Symbol     string                `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	IsBuyType  bool                  `protobuf:"varint,3,opt,name=is_buy_type,json=isBuyType,proto3" json:"is_buy_type,omitempty"`
	StopLoss   *wrappers.DoubleValue `protobuf:"bytes,4,opt,name=stop_loss,json=stopLoss,proto3" json:"stop_loss,omitempty"`
	TakeProfit *wrappers.DoubleValue `protobuf:"bytes,5,opt,name=take_profit,json=takeProfit,proto3" json:"take_profit,omitempty"`
}

func (x *GetOpenPositionResponse) Reset() {
	*x = GetOpenPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOpenPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOpenPositionResponse) ProtoMessage() {}

func (x *GetOpenPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_position_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOpenPositionResponse.ProtoReflect.Descriptor instead.
func (*GetOpenPositionResponse) Descriptor() ([]byte, []int) {
	return file_position_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetOpenPositionResponse) GetAddPrice() float64 {
	if x != nil {
		return x.AddPrice
	}
	return 0
}

func (x *GetOpenPositionResponse) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *GetOpenPositionResponse) GetIsBuyType() bool {
	if x != nil {
		return x.IsBuyType
	}
	return false
}

func (x *GetOpenPositionResponse) GetStopLoss() *wrappers.DoubleValue {
	if x != nil {
		return x.StopLoss
	}
	return nil
}

func (x *GetOpenPositionResponse) GetTakeProfit() *wrappers.DoubleValue {
	if x != nil {
		return x.TakeProfit
	}
	return nil
}

type SetStopLossRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionId int64   `protobuf:"varint,1,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
	StopLoss   float64 `protobuf:"fixed64,2,opt,name=stop_loss,json=stopLoss,proto3" json:"stop_loss,omitempty"`
}

func (x *SetStopLossRequest) Reset() {
	*x = SetStopLossRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetStopLossRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStopLossRequest) ProtoMessage() {}

func (x *SetStopLossRequest) ProtoReflect() protoreflect.Message {
	mi := &file_position_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStopLossRequest.ProtoReflect.Descriptor instead.
func (*SetStopLossRequest) Descriptor() ([]byte, []int) {
	return file_position_service_proto_rawDescGZIP(), []int{6}
}

func (x *SetStopLossRequest) GetPositionId() int64 {
	if x != nil {
		return x.PositionId
	}
	return 0
}

func (x *SetStopLossRequest) GetStopLoss() float64 {
	if x != nil {
		return x.StopLoss
	}
	return 0
}

type SetTakeProfitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionId int64   `protobuf:"varint,1,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
	TakeProfit float64 `protobuf:"fixed64,2,opt,name=take_profit,json=takeProfit,proto3" json:"take_profit,omitempty"`
}

func (x *SetTakeProfitRequest) Reset() {
	*x = SetTakeProfitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTakeProfitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTakeProfitRequest) ProtoMessage() {}

func (x *SetTakeProfitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_position_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTakeProfitRequest.ProtoReflect.Descriptor instead.
func (*SetTakeProfitRequest) Descriptor() ([]byte, []int) {
	return file_position_service_proto_rawDescGZIP(), []int{7}
}

func (x *SetTakeProfitRequest) GetPositionId() int64 {
	if x != nil {
		return x.PositionId
	}
	return 0
}

func (x *SetTakeProfitRequest) GetTakeProfit() float64 {
	if x != nil {
		return x.TakeProfit
	}
	return 0
}

var File_position_service_proto protoreflect.FileDescriptor

var file_position_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x13, 0x4f, 0x70,
	0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x73, 0x5f,
	0x62, 0x75, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x42, 0x75, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x37, 0x0a,
	0x14, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x14, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x15, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x22, 0x39, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xe8, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4f, 0x70,
	0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61, 0x64, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x62, 0x75,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x42, 0x75, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x70, 0x5f,
	0x6c, 0x6f, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x70, 0x4c, 0x6f,
	0x73, 0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x74, 0x61, 0x6b, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x74, 0x22, 0x52, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x70,
	0x5f, 0x6c, 0x6f, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x73, 0x74, 0x6f,
	0x70, 0x4c, 0x6f, 0x73, 0x73, 0x22, 0x58, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x54, 0x61, 0x6b, 0x65,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x61, 0x6b, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x32,
	0xf2, 0x02, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x6e,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x0b, 0x53, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x73, 0x73, 0x12, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x54, 0x61, 0x6b, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74,
	0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x6b, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_position_service_proto_rawDescOnce sync.Once
	file_position_service_proto_rawDescData = file_position_service_proto_rawDesc
)

func file_position_service_proto_rawDescGZIP() []byte {
	file_position_service_proto_rawDescOnce.Do(func() {
		file_position_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_position_service_proto_rawDescData)
	})
	return file_position_service_proto_rawDescData
}

var file_position_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_position_service_proto_goTypes = []interface{}{
	(*OpenPositionRequest)(nil),     // 0: pb.OpenPositionRequest
	(*OpenPositionResponse)(nil),    // 1: pb.OpenPositionResponse
	(*ClosePositionRequest)(nil),    // 2: pb.ClosePositionRequest
	(*ClosePositionResponse)(nil),   // 3: pb.ClosePositionResponse
	(*GetOpenPositionRequest)(nil),  // 4: pb.GetOpenPositionRequest
	(*GetOpenPositionResponse)(nil), // 5: pb.GetOpenPositionResponse
	(*SetStopLossRequest)(nil),      // 6: pb.SetStopLossRequest
	(*SetTakeProfitRequest)(nil),    // 7: pb.SetTakeProfitRequest
	(*wrappers.DoubleValue)(nil),    // 8: google.protobuf.DoubleValue
	(*empty.Empty)(nil),             // 9: google.protobuf.Empty
}
var file_position_service_proto_depIdxs = []int32{
	8, // 0: pb.GetOpenPositionResponse.stop_loss:type_name -> google.protobuf.DoubleValue
	8, // 1: pb.GetOpenPositionResponse.take_profit:type_name -> google.protobuf.DoubleValue
	0, // 2: pb.PositionService.OpenPosition:input_type -> pb.OpenPositionRequest
	2, // 3: pb.PositionService.ClosePosition:input_type -> pb.ClosePositionRequest
	4, // 4: pb.PositionService.GetOpenPosition:input_type -> pb.GetOpenPositionRequest
	6, // 5: pb.PositionService.SetStopLoss:input_type -> pb.SetStopLossRequest
	7, // 6: pb.PositionService.SetTakeProfit:input_type -> pb.SetTakeProfitRequest
	1, // 7: pb.PositionService.OpenPosition:output_type -> pb.OpenPositionResponse
	3, // 8: pb.PositionService.ClosePosition:output_type -> pb.ClosePositionResponse
	5, // 9: pb.PositionService.GetOpenPosition:output_type -> pb.GetOpenPositionResponse
	9, // 10: pb.PositionService.SetStopLoss:output_type -> google.protobuf.Empty
	9, // 11: pb.PositionService.SetTakeProfit:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_position_service_proto_init() }
func file_position_service_proto_init() {
	if File_position_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_position_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenPositionRequest); i {
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
		file_position_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenPositionResponse); i {
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
		file_position_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClosePositionRequest); i {
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
		file_position_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClosePositionResponse); i {
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
		file_position_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOpenPositionRequest); i {
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
		file_position_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOpenPositionResponse); i {
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
		file_position_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetStopLossRequest); i {
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
		file_position_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTakeProfitRequest); i {
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
			RawDescriptor: file_position_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_position_service_proto_goTypes,
		DependencyIndexes: file_position_service_proto_depIdxs,
		MessageInfos:      file_position_service_proto_msgTypes,
	}.Build()
	File_position_service_proto = out.File
	file_position_service_proto_rawDesc = nil
	file_position_service_proto_goTypes = nil
	file_position_service_proto_depIdxs = nil
}
