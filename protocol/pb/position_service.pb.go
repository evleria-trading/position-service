// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: position_service.proto

package pb

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

type OpenPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol    string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	IsBuyType bool   `protobuf:"varint,2,opt,name=is_buy_type,json=isBuyType,proto3" json:"is_buy_type,omitempty"`
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

	PositionId int64 `protobuf:"varint,1,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
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

var File_position_service_proto protoreflect.FileDescriptor

var file_position_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x13, 0x4f, 0x70, 0x65, 0x6e,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x62, 0x75,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x42, 0x75, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x37, 0x0a, 0x14, 0x4f, 0x70, 0x65, 0x6e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x22, 0x37, 0x0a, 0x14, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x15, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x32, 0x92, 0x01, 0x0a, 0x0f, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d,
	0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a,
	0x0d, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15,
	0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_position_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_position_service_proto_goTypes = []interface{}{
	(*OpenPositionRequest)(nil),   // 0: OpenPositionRequest
	(*OpenPositionResponse)(nil),  // 1: OpenPositionResponse
	(*ClosePositionRequest)(nil),  // 2: ClosePositionRequest
	(*ClosePositionResponse)(nil), // 3: ClosePositionResponse
}
var file_position_service_proto_depIdxs = []int32{
	0, // 0: PositionService.OpenPosition:input_type -> OpenPositionRequest
	2, // 1: PositionService.ClosePosition:input_type -> ClosePositionRequest
	1, // 2: PositionService.OpenPosition:output_type -> OpenPositionResponse
	3, // 3: PositionService.ClosePosition:output_type -> ClosePositionResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_position_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
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