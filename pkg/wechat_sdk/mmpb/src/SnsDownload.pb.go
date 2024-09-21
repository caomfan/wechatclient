// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: SnsDownload.proto

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

type SnsDownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *BaseRequest `protobuf:"bytes,1,req,name=baseRequest" json:"baseRequest,omitempty"`
	StartPos    *int32       `protobuf:"varint,2,req,name=startPos" json:"startPos,omitempty"`
	TotalLen    *int32       `protobuf:"varint,3,opt,name=totalLen" json:"totalLen,omitempty"`
	BufferId    *string      `protobuf:"bytes,4,opt,name=bufferId" json:"bufferId,omitempty"`
	Type        *uint32      `protobuf:"varint,5,req,name=type" json:"type,omitempty"`
	DownBufLen  *uint32      `protobuf:"varint,6,req,name=downBufLen" json:"downBufLen,omitempty"`
}

func (x *SnsDownloadRequest) Reset() {
	*x = SnsDownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SnsDownload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnsDownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnsDownloadRequest) ProtoMessage() {}

func (x *SnsDownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_SnsDownload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnsDownloadRequest.ProtoReflect.Descriptor instead.
func (*SnsDownloadRequest) Descriptor() ([]byte, []int) {
	return file_SnsDownload_proto_rawDescGZIP(), []int{0}
}

func (x *SnsDownloadRequest) GetBaseRequest() *BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *SnsDownloadRequest) GetStartPos() int32 {
	if x != nil && x.StartPos != nil {
		return *x.StartPos
	}
	return 0
}

func (x *SnsDownloadRequest) GetTotalLen() int32 {
	if x != nil && x.TotalLen != nil {
		return *x.TotalLen
	}
	return 0
}

func (x *SnsDownloadRequest) GetBufferId() string {
	if x != nil && x.BufferId != nil {
		return *x.BufferId
	}
	return ""
}

func (x *SnsDownloadRequest) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *SnsDownloadRequest) GetDownBufLen() uint32 {
	if x != nil && x.DownBufLen != nil {
		return *x.DownBufLen
	}
	return 0
}

type SnsDownloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *BaseResponse     `protobuf:"bytes,1,req,name=baseResponse" json:"baseResponse,omitempty"`
	StartPos     *int32            `protobuf:"varint,2,req,name=startPos" json:"startPos,omitempty"`
	TotalLen     *int32            `protobuf:"varint,3,req,name=totalLen" json:"totalLen,omitempty"`
	Buffer       *SKBuiltinBufferT `protobuf:"bytes,4,req,name=buffer" json:"buffer,omitempty"`
	BufferId     *string           `protobuf:"bytes,5,opt,name=bufferId" json:"bufferId,omitempty"`
	Type         *uint32           `protobuf:"varint,6,req,name=type" json:"type,omitempty"`
}

func (x *SnsDownloadResponse) Reset() {
	*x = SnsDownloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SnsDownload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnsDownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnsDownloadResponse) ProtoMessage() {}

func (x *SnsDownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_SnsDownload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnsDownloadResponse.ProtoReflect.Descriptor instead.
func (*SnsDownloadResponse) Descriptor() ([]byte, []int) {
	return file_SnsDownload_proto_rawDescGZIP(), []int{1}
}

func (x *SnsDownloadResponse) GetBaseResponse() *BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

func (x *SnsDownloadResponse) GetStartPos() int32 {
	if x != nil && x.StartPos != nil {
		return *x.StartPos
	}
	return 0
}

func (x *SnsDownloadResponse) GetTotalLen() int32 {
	if x != nil && x.TotalLen != nil {
		return *x.TotalLen
	}
	return 0
}

func (x *SnsDownloadResponse) GetBuffer() *SKBuiltinBufferT {
	if x != nil {
		return x.Buffer
	}
	return nil
}

func (x *SnsDownloadResponse) GetBufferId() string {
	if x != nil && x.BufferId != nil {
		return *x.BufferId
	}
	return ""
}

func (x *SnsDownloadResponse) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

var File_SnsDownload_proto protoreflect.FileDescriptor

var file_SnsDownload_proto_rawDesc = []byte{
	0x0a, 0x11, 0x53, 0x6e, 0x73, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x4d, 0x73, 0x67, 0x42, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x12, 0x53, 0x6e, 0x73, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e,
	0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05,
	0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x4c, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x4c, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x77, 0x6e, 0x42, 0x75,
	0x66, 0x4c, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0a, 0x64, 0x6f, 0x77, 0x6e,
	0x42, 0x75, 0x66, 0x4c, 0x65, 0x6e, 0x22, 0xdc, 0x01, 0x0a, 0x13, 0x53, 0x6e, 0x73, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x73, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x02, 0x28, 0x05, 0x52,
	0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x06, 0x62, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x4b, 0x42, 0x75,
	0x69, 0x6c, 0x74, 0x69, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x74, 0x52, 0x06, 0x62,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x6d, 0x73, 0x67,
}

var (
	file_SnsDownload_proto_rawDescOnce sync.Once
	file_SnsDownload_proto_rawDescData = file_SnsDownload_proto_rawDesc
)

func file_SnsDownload_proto_rawDescGZIP() []byte {
	file_SnsDownload_proto_rawDescOnce.Do(func() {
		file_SnsDownload_proto_rawDescData = protoimpl.X.CompressGZIP(file_SnsDownload_proto_rawDescData)
	})
	return file_SnsDownload_proto_rawDescData
}

var file_SnsDownload_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_SnsDownload_proto_goTypes = []interface{}{
	(*SnsDownloadRequest)(nil),  // 0: SnsDownloadRequest
	(*SnsDownloadResponse)(nil), // 1: SnsDownloadResponse
	(*BaseRequest)(nil),         // 2: BaseRequest
	(*BaseResponse)(nil),        // 3: BaseResponse
	(*SKBuiltinBufferT)(nil),    // 4: SKBuiltinBuffer_t
}
var file_SnsDownload_proto_depIdxs = []int32{
	2, // 0: SnsDownloadRequest.baseRequest:type_name -> BaseRequest
	3, // 1: SnsDownloadResponse.baseResponse:type_name -> BaseResponse
	4, // 2: SnsDownloadResponse.buffer:type_name -> SKBuiltinBuffer_t
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_SnsDownload_proto_init() }
func file_SnsDownload_proto_init() {
	if File_SnsDownload_proto != nil {
		return
	}
	file_MicroMsgBase_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SnsDownload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnsDownloadRequest); i {
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
		file_SnsDownload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnsDownloadResponse); i {
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
			RawDescriptor: file_SnsDownload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SnsDownload_proto_goTypes,
		DependencyIndexes: file_SnsDownload_proto_depIdxs,
		MessageInfos:      file_SnsDownload_proto_msgTypes,
	}.Build()
	File_SnsDownload_proto = out.File
	file_SnsDownload_proto_rawDesc = nil
	file_SnsDownload_proto_goTypes = nil
	file_SnsDownload_proto_depIdxs = nil
}
