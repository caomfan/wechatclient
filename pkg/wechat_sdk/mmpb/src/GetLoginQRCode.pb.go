// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: GetLoginQRCode.proto

package micromsg

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

type GetLoginQRCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest      *BaseRequest      `protobuf:"bytes,1,opt,name=baseRequest" json:"baseRequest,omitempty"`
	RandomEncryKey   *SKBuiltinBufferT `protobuf:"bytes,2,opt,name=randomEncryKey" json:"randomEncryKey,omitempty"`
	Opcode           *uint32           `protobuf:"varint,3,opt,name=opcode" json:"opcode,omitempty"`
	DeviceName       []byte            `protobuf:"bytes,4,opt,name=deviceName" json:"deviceName,omitempty"`
	UserName         *string           `protobuf:"bytes,5,opt,name=userName" json:"userName,omitempty"`
	ExtDevLoginType  *uint32           `protobuf:"varint,6,opt,name=extDevLoginType" json:"extDevLoginType,omitempty"`
	HardwareExtra    *string           `protobuf:"bytes,7,opt,name=hardwareExtra" json:"hardwareExtra,omitempty"`
	SoftType         *string           `protobuf:"bytes,8,opt,name=softType" json:"softType,omitempty"`
	MsgContextPubKey *SKBuiltinBufferT `protobuf:"bytes,9,opt,name=msgContextPubKey" json:"msgContextPubKey,omitempty"`
}

func (x *GetLoginQRCodeRequest) Reset() {
	*x = GetLoginQRCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetLoginQRCode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoginQRCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginQRCodeRequest) ProtoMessage() {}

func (x *GetLoginQRCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_GetLoginQRCode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginQRCodeRequest.ProtoReflect.Descriptor instead.
func (*GetLoginQRCodeRequest) Descriptor() ([]byte, []int) {
	return file_GetLoginQRCode_proto_rawDescGZIP(), []int{0}
}

func (x *GetLoginQRCodeRequest) GetBaseRequest() *BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *GetLoginQRCodeRequest) GetRandomEncryKey() *SKBuiltinBufferT {
	if x != nil {
		return x.RandomEncryKey
	}
	return nil
}

func (x *GetLoginQRCodeRequest) GetOpcode() uint32 {
	if x != nil && x.Opcode != nil {
		return *x.Opcode
	}
	return 0
}

func (x *GetLoginQRCodeRequest) GetDeviceName() []byte {
	if x != nil {
		return x.DeviceName
	}
	return nil
}

func (x *GetLoginQRCodeRequest) GetUserName() string {
	if x != nil && x.UserName != nil {
		return *x.UserName
	}
	return ""
}

func (x *GetLoginQRCodeRequest) GetExtDevLoginType() uint32 {
	if x != nil && x.ExtDevLoginType != nil {
		return *x.ExtDevLoginType
	}
	return 0
}

func (x *GetLoginQRCodeRequest) GetHardwareExtra() string {
	if x != nil && x.HardwareExtra != nil {
		return *x.HardwareExtra
	}
	return ""
}

func (x *GetLoginQRCodeRequest) GetSoftType() string {
	if x != nil && x.SoftType != nil {
		return *x.SoftType
	}
	return ""
}

func (x *GetLoginQRCodeRequest) GetMsgContextPubKey() *SKBuiltinBufferT {
	if x != nil {
		return x.MsgContextPubKey
	}
	return nil
}

type GetLoginQRCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse              *BaseResponse     `protobuf:"bytes,1,opt,name=baseResponse" json:"baseResponse,omitempty"`
	Qrcode                    *SKBuiltinBufferT `protobuf:"bytes,2,opt,name=qrcode" json:"qrcode,omitempty"`
	Uuid                      *string           `protobuf:"bytes,3,opt,name=uuid" json:"uuid,omitempty"`
	CheckTime                 *uint32           `protobuf:"varint,4,opt,name=checkTime" json:"checkTime,omitempty"`
	NotifyKey                 *SKBuiltinBufferT `protobuf:"bytes,5,opt,name=notifyKey" json:"notifyKey,omitempty"`
	ExpiredTime               *uint32           `protobuf:"varint,6,opt,name=expiredTime" json:"expiredTime,omitempty"`
	BlueToothBroadCastUuid    *string           `protobuf:"bytes,7,opt,name=blueToothBroadCastUuid" json:"blueToothBroadCastUuid,omitempty"`
	BlueToothBroadCastContent *SKBuiltinBufferT `protobuf:"bytes,8,opt,name=blueToothBroadCastContent" json:"blueToothBroadCastContent,omitempty"`
	FileTransferAssistant     *string           `protobuf:"bytes,9,opt,name=fileTransferAssistant" json:"fileTransferAssistant,omitempty"`
	QrScanUrl                 *string           `protobuf:"bytes,10,opt,name=qrScanUrl" json:"qrScanUrl,omitempty"`
}

func (x *GetLoginQRCodeResponse) Reset() {
	*x = GetLoginQRCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetLoginQRCode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoginQRCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginQRCodeResponse) ProtoMessage() {}

func (x *GetLoginQRCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_GetLoginQRCode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginQRCodeResponse.ProtoReflect.Descriptor instead.
func (*GetLoginQRCodeResponse) Descriptor() ([]byte, []int) {
	return file_GetLoginQRCode_proto_rawDescGZIP(), []int{1}
}

func (x *GetLoginQRCodeResponse) GetBaseResponse() *BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

func (x *GetLoginQRCodeResponse) GetQrcode() *SKBuiltinBufferT {
	if x != nil {
		return x.Qrcode
	}
	return nil
}

func (x *GetLoginQRCodeResponse) GetUuid() string {
	if x != nil && x.Uuid != nil {
		return *x.Uuid
	}
	return ""
}

func (x *GetLoginQRCodeResponse) GetCheckTime() uint32 {
	if x != nil && x.CheckTime != nil {
		return *x.CheckTime
	}
	return 0
}

func (x *GetLoginQRCodeResponse) GetNotifyKey() *SKBuiltinBufferT {
	if x != nil {
		return x.NotifyKey
	}
	return nil
}

func (x *GetLoginQRCodeResponse) GetExpiredTime() uint32 {
	if x != nil && x.ExpiredTime != nil {
		return *x.ExpiredTime
	}
	return 0
}

func (x *GetLoginQRCodeResponse) GetBlueToothBroadCastUuid() string {
	if x != nil && x.BlueToothBroadCastUuid != nil {
		return *x.BlueToothBroadCastUuid
	}
	return ""
}

func (x *GetLoginQRCodeResponse) GetBlueToothBroadCastContent() *SKBuiltinBufferT {
	if x != nil {
		return x.BlueToothBroadCastContent
	}
	return nil
}

func (x *GetLoginQRCodeResponse) GetFileTransferAssistant() string {
	if x != nil && x.FileTransferAssistant != nil {
		return *x.FileTransferAssistant
	}
	return ""
}

func (x *GetLoginQRCodeResponse) GetQrScanUrl() string {
	if x != nil && x.QrScanUrl != nil {
		return *x.QrScanUrl
	}
	return ""
}

var File_GetLoginQRCode_proto protoreflect.FileDescriptor

var file_GetLoginQRCode_proto_rawDesc = []byte{
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x4d, 0x73, 0x67,
	0x42, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x03, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0e, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53,
	0x4b, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x74,
	0x52, 0x0e, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x4b, 0x65, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x6f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x44, 0x65, 0x76, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x65,
	0x78, 0x74, 0x44, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24,
	0x0a, 0x0d, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x45, 0x78, 0x74, 0x72, 0x61, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x66, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x66, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x3e, 0x0a, 0x10, 0x6d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x50, 0x75,
	0x62, 0x4b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x4b, 0x42,
	0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x74, 0x52, 0x10,
	0x6d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79,
	0x22, 0xdb, 0x03, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x51, 0x52, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0c, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x06, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x53, 0x4b, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x5f, 0x74, 0x52, 0x06, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x09,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x53, 0x4b, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x5f, 0x74, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x20,
	0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x36, 0x0a, 0x16, 0x62, 0x6c, 0x75, 0x65, 0x54, 0x6f, 0x6f, 0x74, 0x68, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x55, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x16, 0x62, 0x6c, 0x75, 0x65, 0x54, 0x6f, 0x6f, 0x74, 0x68, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x43, 0x61, 0x73, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x50, 0x0a, 0x19, 0x62, 0x6c, 0x75, 0x65,
	0x54, 0x6f, 0x6f, 0x74, 0x68, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x4b,
	0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x74, 0x52,
	0x19, 0x62, 0x6c, 0x75, 0x65, 0x54, 0x6f, 0x6f, 0x74, 0x68, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43,
	0x61, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x15, 0x66, 0x69,
	0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x66, 0x69, 0x6c, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x71, 0x72, 0x53, 0x63, 0x61, 0x6e, 0x55, 0x72, 0x6c, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x72, 0x53, 0x63, 0x61, 0x6e, 0x55, 0x72, 0x6c, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x3b, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x73, 0x67,
}

var (
	file_GetLoginQRCode_proto_rawDescOnce sync.Once
	file_GetLoginQRCode_proto_rawDescData = file_GetLoginQRCode_proto_rawDesc
)

func file_GetLoginQRCode_proto_rawDescGZIP() []byte {
	file_GetLoginQRCode_proto_rawDescOnce.Do(func() {
		file_GetLoginQRCode_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetLoginQRCode_proto_rawDescData)
	})
	return file_GetLoginQRCode_proto_rawDescData
}

var file_GetLoginQRCode_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_GetLoginQRCode_proto_goTypes = []interface{}{
	(*GetLoginQRCodeRequest)(nil),  // 0: GetLoginQRCodeRequest
	(*GetLoginQRCodeResponse)(nil), // 1: GetLoginQRCodeResponse
	(*BaseRequest)(nil),            // 2: BaseRequest
	(*SKBuiltinBufferT)(nil),       // 3: SKBuiltinBuffer_t
	(*BaseResponse)(nil),           // 4: BaseResponse
}
var file_GetLoginQRCode_proto_depIdxs = []int32{
	2, // 0: GetLoginQRCodeRequest.baseRequest:type_name -> BaseRequest
	3, // 1: GetLoginQRCodeRequest.randomEncryKey:type_name -> SKBuiltinBuffer_t
	3, // 2: GetLoginQRCodeRequest.msgContextPubKey:type_name -> SKBuiltinBuffer_t
	4, // 3: GetLoginQRCodeResponse.baseResponse:type_name -> BaseResponse
	3, // 4: GetLoginQRCodeResponse.qrcode:type_name -> SKBuiltinBuffer_t
	3, // 5: GetLoginQRCodeResponse.notifyKey:type_name -> SKBuiltinBuffer_t
	3, // 6: GetLoginQRCodeResponse.blueToothBroadCastContent:type_name -> SKBuiltinBuffer_t
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_GetLoginQRCode_proto_init() }
func file_GetLoginQRCode_proto_init() {
	if File_GetLoginQRCode_proto != nil {
		return
	}
	file_MicroMsgBase_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetLoginQRCode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoginQRCodeRequest); i {
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
		file_GetLoginQRCode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoginQRCodeResponse); i {
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
			RawDescriptor: file_GetLoginQRCode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetLoginQRCode_proto_goTypes,
		DependencyIndexes: file_GetLoginQRCode_proto_depIdxs,
		MessageInfos:      file_GetLoginQRCode_proto_msgTypes,
	}.Build()
	File_GetLoginQRCode_proto = out.File
	file_GetLoginQRCode_proto_rawDesc = nil
	file_GetLoginQRCode_proto_goTypes = nil
	file_GetLoginQRCode_proto_depIdxs = nil
}
