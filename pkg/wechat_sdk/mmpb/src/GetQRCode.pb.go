// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: GetQRCode.proto

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

type GetQRCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *BaseRequest      `protobuf:"bytes,1,req,name=baseRequest" json:"baseRequest,omitempty"`
	UserName    *SKBuiltinStringT `protobuf:"bytes,2,req,name=userName" json:"userName,omitempty"`
	Style       *uint32           `protobuf:"varint,3,req,name=style" json:"style,omitempty"`
	OpCode      *uint32           `protobuf:"varint,4,opt,name=opCode" json:"opCode,omitempty"`
}

func (x *GetQRCodeRequest) Reset() {
	*x = GetQRCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetQRCode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQRCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQRCodeRequest) ProtoMessage() {}

func (x *GetQRCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_GetQRCode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQRCodeRequest.ProtoReflect.Descriptor instead.
func (*GetQRCodeRequest) Descriptor() ([]byte, []int) {
	return file_GetQRCode_proto_rawDescGZIP(), []int{0}
}

func (x *GetQRCodeRequest) GetBaseRequest() *BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *GetQRCodeRequest) GetUserName() *SKBuiltinStringT {
	if x != nil {
		return x.UserName
	}
	return nil
}

func (x *GetQRCodeRequest) GetStyle() uint32 {
	if x != nil && x.Style != nil {
		return *x.Style
	}
	return 0
}

func (x *GetQRCodeRequest) GetOpCode() uint32 {
	if x != nil && x.OpCode != nil {
		return *x.OpCode
	}
	return 0
}

type GetQRCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse        *BaseResponse     `protobuf:"bytes,1,req,name=baseResponse" json:"baseResponse,omitempty"`
	Qrcode              *SKBuiltinBufferT `protobuf:"bytes,2,req,name=qrcode" json:"qrcode,omitempty"`
	Style               *uint32           `protobuf:"varint,5,req,name=style" json:"style,omitempty"`
	FooterWording       *string           `protobuf:"bytes,6,opt,name=footerWording" json:"footerWording,omitempty"`
	RevokeQrcodeId      *string           `protobuf:"bytes,7,opt,name=revokeQrcodeId" json:"revokeQrcodeId,omitempty"`
	RevokeQrcodeWording *string           `protobuf:"bytes,8,opt,name=revokeQrcodeWording" json:"revokeQrcodeWording,omitempty"`
}

func (x *GetQRCodeResponse) Reset() {
	*x = GetQRCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetQRCode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQRCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQRCodeResponse) ProtoMessage() {}

func (x *GetQRCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_GetQRCode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQRCodeResponse.ProtoReflect.Descriptor instead.
func (*GetQRCodeResponse) Descriptor() ([]byte, []int) {
	return file_GetQRCode_proto_rawDescGZIP(), []int{1}
}

func (x *GetQRCodeResponse) GetBaseResponse() *BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

func (x *GetQRCodeResponse) GetQrcode() *SKBuiltinBufferT {
	if x != nil {
		return x.Qrcode
	}
	return nil
}

func (x *GetQRCodeResponse) GetStyle() uint32 {
	if x != nil && x.Style != nil {
		return *x.Style
	}
	return 0
}

func (x *GetQRCodeResponse) GetFooterWording() string {
	if x != nil && x.FooterWording != nil {
		return *x.FooterWording
	}
	return ""
}

func (x *GetQRCodeResponse) GetRevokeQrcodeId() string {
	if x != nil && x.RevokeQrcodeId != nil {
		return *x.RevokeQrcodeId
	}
	return ""
}

func (x *GetQRCodeResponse) GetRevokeQrcodeWording() string {
	if x != nil && x.RevokeQrcodeWording != nil {
		return *x.RevokeQrcodeWording
	}
	return ""
}

var File_GetQRCode_proto protoreflect.FileDescriptor

var file_GetQRCode_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x4d, 0x73, 0x67, 0x42, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x51, 0x52, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x62, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53,
	0x4b, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x79, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x79, 0x6c, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x6f, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x88, 0x02, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x06, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x53, 0x4b, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x42, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x5f, 0x74, 0x52, 0x06, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74,
	0x79, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x57, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x6f, 0x6f, 0x74,
	0x65, 0x72, 0x57, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x51, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x51, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x30, 0x0a, 0x13, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x51, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x57, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13,
	0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x51, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x57, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x73,
	0x67,
}

var (
	file_GetQRCode_proto_rawDescOnce sync.Once
	file_GetQRCode_proto_rawDescData = file_GetQRCode_proto_rawDesc
)

func file_GetQRCode_proto_rawDescGZIP() []byte {
	file_GetQRCode_proto_rawDescOnce.Do(func() {
		file_GetQRCode_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetQRCode_proto_rawDescData)
	})
	return file_GetQRCode_proto_rawDescData
}

var file_GetQRCode_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_GetQRCode_proto_goTypes = []interface{}{
	(*GetQRCodeRequest)(nil),  // 0: GetQRCodeRequest
	(*GetQRCodeResponse)(nil), // 1: GetQRCodeResponse
	(*BaseRequest)(nil),       // 2: BaseRequest
	(*SKBuiltinStringT)(nil),  // 3: SKBuiltinString_t
	(*BaseResponse)(nil),      // 4: BaseResponse
	(*SKBuiltinBufferT)(nil),  // 5: SKBuiltinBuffer_t
}
var file_GetQRCode_proto_depIdxs = []int32{
	2, // 0: GetQRCodeRequest.baseRequest:type_name -> BaseRequest
	3, // 1: GetQRCodeRequest.userName:type_name -> SKBuiltinString_t
	4, // 2: GetQRCodeResponse.baseResponse:type_name -> BaseResponse
	5, // 3: GetQRCodeResponse.qrcode:type_name -> SKBuiltinBuffer_t
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_GetQRCode_proto_init() }
func file_GetQRCode_proto_init() {
	if File_GetQRCode_proto != nil {
		return
	}
	file_MicroMsgBase_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetQRCode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQRCodeRequest); i {
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
		file_GetQRCode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQRCodeResponse); i {
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
			RawDescriptor: file_GetQRCode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetQRCode_proto_goTypes,
		DependencyIndexes: file_GetQRCode_proto_depIdxs,
		MessageInfos:      file_GetQRCode_proto_msgTypes,
	}.Build()
	File_GetQRCode_proto = out.File
	file_GetQRCode_proto_rawDesc = nil
	file_GetQRCode_proto_goTypes = nil
	file_GetQRCode_proto_depIdxs = nil
}
