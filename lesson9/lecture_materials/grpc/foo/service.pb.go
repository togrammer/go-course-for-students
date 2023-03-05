// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: cmd/grpc/foo/service.proto

package foo

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InstrumentType int32

const (
	InstrumentType_Equity InstrumentType = 0
	InstrumentType_Bond   InstrumentType = 1
)

// Enum value maps for InstrumentType.
var (
	InstrumentType_name = map[int32]string{
		0: "Equity",
		1: "Bond",
	}
	InstrumentType_value = map[string]int32{
		"Equity": 0,
		"Bond":   1,
	}
)

func (x InstrumentType) Enum() *InstrumentType {
	p := new(InstrumentType)
	*p = x
	return p
}

func (x InstrumentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InstrumentType) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_grpc_foo_service_proto_enumTypes[0].Descriptor()
}

func (InstrumentType) Type() protoreflect.EnumType {
	return &file_cmd_grpc_foo_service_proto_enumTypes[0]
}

func (x InstrumentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InstrumentType.Descriptor instead.
func (InstrumentType) EnumDescriptor() ([]byte, []int) {
	return file_cmd_grpc_foo_service_proto_rawDescGZIP(), []int{0}
}

type Instrument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Ticker string         `protobuf:"bytes,2,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Type   InstrumentType `protobuf:"varint,3,opt,name=type,proto3,enum=foo.InstrumentType" json:"type,omitempty"`
}

func (x *Instrument) Reset() {
	*x = Instrument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_grpc_foo_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instrument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instrument) ProtoMessage() {}

func (x *Instrument) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_grpc_foo_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instrument.ProtoReflect.Descriptor instead.
func (*Instrument) Descriptor() ([]byte, []int) {
	return file_cmd_grpc_foo_service_proto_rawDescGZIP(), []int{0}
}

func (x *Instrument) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Instrument) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *Instrument) GetType() InstrumentType {
	if x != nil {
		return x.Type
	}
	return InstrumentType_Equity
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstrumentID int64                  `protobuf:"varint,1,opt,name=instrumentID,proto3" json:"instrumentID,omitempty"`
	Value        float64                `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	Ts           *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_grpc_foo_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_grpc_foo_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_cmd_grpc_foo_service_proto_rawDescGZIP(), []int{1}
}

func (x *Price) GetInstrumentID() int64 {
	if x != nil {
		return x.InstrumentID
	}
	return 0
}

func (x *Price) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Price) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

type Prices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prices []*Price `protobuf:"bytes,1,rep,name=prices,proto3" json:"prices,omitempty"`
}

func (x *Prices) Reset() {
	*x = Prices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_grpc_foo_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prices) ProtoMessage() {}

func (x *Prices) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_grpc_foo_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prices.ProtoReflect.Descriptor instead.
func (*Prices) Descriptor() ([]byte, []int) {
	return file_cmd_grpc_foo_service_proto_rawDescGZIP(), []int{2}
}

func (x *Prices) GetPrices() []*Price {
	if x != nil {
		return x.Prices
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload string `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_grpc_foo_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_grpc_foo_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_cmd_grpc_foo_service_proto_rawDescGZIP(), []int{3}
}

func (x *Message) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type BarMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*BarMessage_Id
	//	*BarMessage_Value
	Data isBarMessage_Data `protobuf_oneof:"Data"`
}

func (x *BarMessage) Reset() {
	*x = BarMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_grpc_foo_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarMessage) ProtoMessage() {}

func (x *BarMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_grpc_foo_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarMessage.ProtoReflect.Descriptor instead.
func (*BarMessage) Descriptor() ([]byte, []int) {
	return file_cmd_grpc_foo_service_proto_rawDescGZIP(), []int{4}
}

func (m *BarMessage) GetData() isBarMessage_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *BarMessage) GetId() string {
	if x, ok := x.GetData().(*BarMessage_Id); ok {
		return x.Id
	}
	return ""
}

func (x *BarMessage) GetValue() float64 {
	if x, ok := x.GetData().(*BarMessage_Value); ok {
		return x.Value
	}
	return 0
}

type isBarMessage_Data interface {
	isBarMessage_Data()
}

type BarMessage_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type BarMessage_Value struct {
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3,oneof"`
}

func (*BarMessage_Id) isBarMessage_Data() {}

func (*BarMessage_Value) isBarMessage_Data() {}

var File_cmd_grpc_foo_service_proto protoreflect.FileDescriptor

var file_cmd_grpc_foo_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6d, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x6f, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x66, 0x6f,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5d, 0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x6d,
	0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x22, 0x2c, 0x0a,
	0x06, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x52, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x22, 0x23, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x3e, 0x0a, 0x0a, 0x42, 0x61, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x48,
	0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x2a, 0x26, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x71, 0x75, 0x69, 0x74, 0x79, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x42, 0x6f, 0x6e, 0x64, 0x10, 0x01, 0x32, 0x9b, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x4c, 0x61,
	0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x0f, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x0a, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x0b, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x73, 0x22, 0x00, 0x28, 0x01, 0x12, 0x2f, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x0f, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x49, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x0a, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x22, 0x00, 0x30, 0x01, 0x32, 0x37, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x0c, 0x2e,
	0x66, 0x6f, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0c, 0x2e, 0x66, 0x6f,
	0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x32,
	0x3c, 0x0a, 0x0a, 0x42, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a,
	0x03, 0x42, 0x61, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x66,
	0x6f, 0x6f, 0x2e, 0x42, 0x61, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x07, 0x5a,
	0x05, 0x2e, 0x2f, 0x66, 0x6f, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_grpc_foo_service_proto_rawDescOnce sync.Once
	file_cmd_grpc_foo_service_proto_rawDescData = file_cmd_grpc_foo_service_proto_rawDesc
)

func file_cmd_grpc_foo_service_proto_rawDescGZIP() []byte {
	file_cmd_grpc_foo_service_proto_rawDescOnce.Do(func() {
		file_cmd_grpc_foo_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_grpc_foo_service_proto_rawDescData)
	})
	return file_cmd_grpc_foo_service_proto_rawDescData
}

var file_cmd_grpc_foo_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cmd_grpc_foo_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cmd_grpc_foo_service_proto_goTypes = []interface{}{
	(InstrumentType)(0),           // 0: foo.InstrumentType
	(*Instrument)(nil),            // 1: foo.Instrument
	(*Price)(nil),                 // 2: foo.Price
	(*Prices)(nil),                // 3: foo.Prices
	(*Message)(nil),               // 4: foo.Message
	(*BarMessage)(nil),            // 5: foo.BarMessage
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_cmd_grpc_foo_service_proto_depIdxs = []int32{
	0, // 0: foo.Instrument.type:type_name -> foo.InstrumentType
	6, // 1: foo.Price.ts:type_name -> google.protobuf.Timestamp
	2, // 2: foo.Prices.prices:type_name -> foo.Price
	1, // 3: foo.PricesService.LastPrice:input_type -> foo.Instrument
	1, // 4: foo.PricesService.GetPrices:input_type -> foo.Instrument
	1, // 5: foo.PricesService.PricesStream:input_type -> foo.Instrument
	4, // 6: foo.ChatService.Chat:input_type -> foo.Message
	7, // 7: foo.BarService.Bar:input_type -> google.protobuf.Empty
	2, // 8: foo.PricesService.LastPrice:output_type -> foo.Price
	3, // 9: foo.PricesService.GetPrices:output_type -> foo.Prices
	2, // 10: foo.PricesService.PricesStream:output_type -> foo.Price
	4, // 11: foo.ChatService.Chat:output_type -> foo.Message
	5, // 12: foo.BarService.Bar:output_type -> foo.BarMessage
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cmd_grpc_foo_service_proto_init() }
func file_cmd_grpc_foo_service_proto_init() {
	if File_cmd_grpc_foo_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_grpc_foo_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instrument); i {
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
		file_cmd_grpc_foo_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_cmd_grpc_foo_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prices); i {
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
		file_cmd_grpc_foo_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_cmd_grpc_foo_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarMessage); i {
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
	file_cmd_grpc_foo_service_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*BarMessage_Id)(nil),
		(*BarMessage_Value)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_grpc_foo_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_cmd_grpc_foo_service_proto_goTypes,
		DependencyIndexes: file_cmd_grpc_foo_service_proto_depIdxs,
		EnumInfos:         file_cmd_grpc_foo_service_proto_enumTypes,
		MessageInfos:      file_cmd_grpc_foo_service_proto_msgTypes,
	}.Build()
	File_cmd_grpc_foo_service_proto = out.File
	file_cmd_grpc_foo_service_proto_rawDesc = nil
	file_cmd_grpc_foo_service_proto_goTypes = nil
	file_cmd_grpc_foo_service_proto_depIdxs = nil
}