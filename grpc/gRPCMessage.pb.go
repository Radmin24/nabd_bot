// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.28.2
// source: gRPCMessage.proto

package grpcmessage

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Сообщение для кнопок
type Button struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Caption       string                 `protobuf:"bytes,1,opt,name=caption,proto3" json:"caption,omitempty"`
	Data          string                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Order         int32                  `protobuf:"varint,3,opt,name=order,proto3" json:"order,omitempty"`
	Row           int32                  `protobuf:"varint,4,opt,name=row,proto3" json:"row,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Button) Reset() {
	*x = Button{}
	mi := &file_gRPCMessage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Button) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Button) ProtoMessage() {}

func (x *Button) ProtoReflect() protoreflect.Message {
	mi := &file_gRPCMessage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Button.ProtoReflect.Descriptor instead.
func (*Button) Descriptor() ([]byte, []int) {
	return file_gRPCMessage_proto_rawDescGZIP(), []int{0}
}

func (x *Button) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

func (x *Button) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Button) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *Button) GetRow() int32 {
	if x != nil {
		return x.Row
	}
	return 0
}

// Сообщение для клавиатуры
type Keyboard struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Buttons       []*Button              `protobuf:"bytes,1,rep,name=buttons,proto3" json:"buttons,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Keyboard) Reset() {
	*x = Keyboard{}
	mi := &file_gRPCMessage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Keyboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keyboard) ProtoMessage() {}

func (x *Keyboard) ProtoReflect() protoreflect.Message {
	mi := &file_gRPCMessage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keyboard.ProtoReflect.Descriptor instead.
func (*Keyboard) Descriptor() ([]byte, []int) {
	return file_gRPCMessage_proto_rawDescGZIP(), []int{1}
}

func (x *Keyboard) GetButtons() []*Button {
	if x != nil {
		return x.Buttons
	}
	return nil
}

func (x *Keyboard) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

// Основное сообщение
type GRPCMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Mes           string                 `protobuf:"bytes,1,opt,name=mes,proto3" json:"mes,omitempty"`
	Delay         int32                  `protobuf:"varint,2,opt,name=delay,proto3" json:"delay,omitempty"`
	Keyboard      *Keyboard              `protobuf:"bytes,3,opt,name=keyboard,proto3" json:"keyboard,omitempty"`
	IsKb          bool                   `protobuf:"varint,4,opt,name=isKb,proto3" json:"isKb,omitempty"`
	Image         string                 `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	ChatId        int64                  `protobuf:"varint,6,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GRPCMessage) Reset() {
	*x = GRPCMessage{}
	mi := &file_gRPCMessage_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GRPCMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPCMessage) ProtoMessage() {}

func (x *GRPCMessage) ProtoReflect() protoreflect.Message {
	mi := &file_gRPCMessage_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPCMessage.ProtoReflect.Descriptor instead.
func (*GRPCMessage) Descriptor() ([]byte, []int) {
	return file_gRPCMessage_proto_rawDescGZIP(), []int{2}
}

func (x *GRPCMessage) GetMes() string {
	if x != nil {
		return x.Mes
	}
	return ""
}

func (x *GRPCMessage) GetDelay() int32 {
	if x != nil {
		return x.Delay
	}
	return 0
}

func (x *GRPCMessage) GetKeyboard() *Keyboard {
	if x != nil {
		return x.Keyboard
	}
	return nil
}

func (x *GRPCMessage) GetIsKb() bool {
	if x != nil {
		return x.IsKb
	}
	return false
}

func (x *GRPCMessage) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *GRPCMessage) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

// Ответ
type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_gRPCMessage_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_gRPCMessage_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_gRPCMessage_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_gRPCMessage_proto protoreflect.FileDescriptor

var file_gRPCMessage_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x67, 0x52, 0x50, 0x43, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x5e, 0x0a, 0x06, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x6f, 0x77,
	0x22, 0x4d, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x2d, 0x0a, 0x07,
	0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x42, 0x75, 0x74, 0x74,
	0x6f, 0x6e, 0x52, 0x07, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0xab, 0x01, 0x0a, 0x0b, 0x67, 0x52, 0x50, 0x43, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x31, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x52, 0x08, 0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73,
	0x4b, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4b, 0x62, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x3c, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x4d, 0x0a, 0x0b, 0x67,
	0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_gRPCMessage_proto_rawDescOnce sync.Once
	file_gRPCMessage_proto_rawDescData []byte
)

func file_gRPCMessage_proto_rawDescGZIP() []byte {
	file_gRPCMessage_proto_rawDescOnce.Do(func() {
		file_gRPCMessage_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_gRPCMessage_proto_rawDesc), len(file_gRPCMessage_proto_rawDesc)))
	})
	return file_gRPCMessage_proto_rawDescData
}

var file_gRPCMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gRPCMessage_proto_goTypes = []any{
	(*Button)(nil),      // 0: grpcmessage.Button
	(*Keyboard)(nil),    // 1: grpcmessage.Keyboard
	(*GRPCMessage)(nil), // 2: grpcmessage.gRPCMessage
	(*Response)(nil),    // 3: grpcmessage.Response
}
var file_gRPCMessage_proto_depIdxs = []int32{
	0, // 0: grpcmessage.Keyboard.buttons:type_name -> grpcmessage.Button
	1, // 1: grpcmessage.gRPCMessage.keyboard:type_name -> grpcmessage.Keyboard
	2, // 2: grpcmessage.gRPCService.SendMessage:input_type -> grpcmessage.gRPCMessage
	3, // 3: grpcmessage.gRPCService.SendMessage:output_type -> grpcmessage.Response
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gRPCMessage_proto_init() }
func file_gRPCMessage_proto_init() {
	if File_gRPCMessage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_gRPCMessage_proto_rawDesc), len(file_gRPCMessage_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gRPCMessage_proto_goTypes,
		DependencyIndexes: file_gRPCMessage_proto_depIdxs,
		MessageInfos:      file_gRPCMessage_proto_msgTypes,
	}.Build()
	File_gRPCMessage_proto = out.File
	file_gRPCMessage_proto_goTypes = nil
	file_gRPCMessage_proto_depIdxs = nil
}
