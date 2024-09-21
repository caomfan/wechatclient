// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: BatchGetHeadImg.proto

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

type BatchGetHeadImgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest  *BaseRequest        `protobuf:"bytes,1,req,name=baseRequest" json:"baseRequest,omitempty"`
	Count        *uint32             `protobuf:"varint,2,req,name=count" json:"count,omitempty"`
	UserNameList []*SKBuiltinStringT `protobuf:"bytes,3,rep,name=userNameList" json:"userNameList,omitempty"`
}

func (x *BatchGetHeadImgRequest) Reset() {
	*x = BatchGetHeadImgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BatchGetHeadImg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetHeadImgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetHeadImgRequest) ProtoMessage() {}

func (x *BatchGetHeadImgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_BatchGetHeadImg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetHeadImgRequest.ProtoReflect.Descriptor instead.
func (*BatchGetHeadImgRequest) Descriptor() ([]byte, []int) {
	return file_BatchGetHeadImg_proto_rawDescGZIP(), []int{0}
}

func (x *BatchGetHeadImgRequest) GetBaseRequest() *BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *BatchGetHeadImgRequest) GetCount() uint32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

func (x *BatchGetHeadImgRequest) GetUserNameList() []*SKBuiltinStringT {
	if x != nil {
		return x.UserNameList
	}
	return nil
}

type ImgPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImgBuf   *SKBuiltinBufferT `protobuf:"bytes,1,req,name=imgBuf" json:"imgBuf,omitempty"`
	Username *SKBuiltinStringT `protobuf:"bytes,2,req,name=username" json:"username,omitempty"`
}

func (x *ImgPair) Reset() {
	*x = ImgPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BatchGetHeadImg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImgPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImgPair) ProtoMessage() {}

func (x *ImgPair) ProtoReflect() protoreflect.Message {
	mi := &file_BatchGetHeadImg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImgPair.ProtoReflect.Descriptor instead.
func (*ImgPair) Descriptor() ([]byte, []int) {
	return file_BatchGetHeadImg_proto_rawDescGZIP(), []int{1}
}

func (x *ImgPair) GetImgBuf() *SKBuiltinBufferT {
	if x != nil {
		return x.ImgBuf
	}
	return nil
}

func (x *ImgPair) GetUsername() *SKBuiltinStringT {
	if x != nil {
		return x.Username
	}
	return nil
}

type BatchGetHeadImgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *BaseResponse `protobuf:"bytes,1,req,name=baseResponse" json:"baseResponse,omitempty"`
	Count        *uint32       `protobuf:"varint,2,req,name=count" json:"count,omitempty"`
	ImgPairList  []*ImgPair    `protobuf:"bytes,3,rep,name=imgPairList" json:"imgPairList,omitempty"`
}

func (x *BatchGetHeadImgResponse) Reset() {
	*x = BatchGetHeadImgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BatchGetHeadImg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetHeadImgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetHeadImgResponse) ProtoMessage() {}

func (x *BatchGetHeadImgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_BatchGetHeadImg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetHeadImgResponse.ProtoReflect.Descriptor instead.
func (*BatchGetHeadImgResponse) Descriptor() ([]byte, []int) {
	return file_BatchGetHeadImg_proto_rawDescGZIP(), []int{2}
}

func (x *BatchGetHeadImgResponse) GetBaseResponse() *BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

func (x *BatchGetHeadImgResponse) GetCount() uint32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

func (x *BatchGetHeadImgResponse) GetImgPairList() []*ImgPair {
	if x != nil {
		return x.ImgPairList
	}
	return nil
}

var File_BatchGetHeadImg_proto protoreflect.FileDescriptor

var file_BatchGetHeadImg_proto_rawDesc = []byte{
	0x0a, 0x15, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x49, 0x6d,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x4d, 0x73,
	0x67, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x16,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x49, 0x6d, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x4b, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x65, 0x0a, 0x07, 0x49, 0x6d, 0x67, 0x50, 0x61, 0x69, 0x72, 0x12,
	0x2a, 0x0a, 0x06, 0x69, 0x6d, 0x67, 0x42, 0x75, 0x66, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x53, 0x4b, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x5f, 0x74, 0x52, 0x06, 0x69, 0x6d, 0x67, 0x42, 0x75, 0x66, 0x12, 0x2e, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x53, 0x4b, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x74, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x17,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x49, 0x6d, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x62, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x2a, 0x0a, 0x0b, 0x69, 0x6d, 0x67, 0x50, 0x61, 0x69, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x49, 0x6d, 0x67, 0x50, 0x61, 0x69, 0x72, 0x52,
	0x0b, 0x69, 0x6d, 0x67, 0x50, 0x61, 0x69, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x3b, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x73, 0x67,
}

var (
	file_BatchGetHeadImg_proto_rawDescOnce sync.Once
	file_BatchGetHeadImg_proto_rawDescData = file_BatchGetHeadImg_proto_rawDesc
)

func file_BatchGetHeadImg_proto_rawDescGZIP() []byte {
	file_BatchGetHeadImg_proto_rawDescOnce.Do(func() {
		file_BatchGetHeadImg_proto_rawDescData = protoimpl.X.CompressGZIP(file_BatchGetHeadImg_proto_rawDescData)
	})
	return file_BatchGetHeadImg_proto_rawDescData
}

var file_BatchGetHeadImg_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_BatchGetHeadImg_proto_goTypes = []interface{}{
	(*BatchGetHeadImgRequest)(nil),  // 0: BatchGetHeadImgRequest
	(*ImgPair)(nil),                 // 1: ImgPair
	(*BatchGetHeadImgResponse)(nil), // 2: BatchGetHeadImgResponse
	(*BaseRequest)(nil),             // 3: BaseRequest
	(*SKBuiltinStringT)(nil),        // 4: SKBuiltinString_t
	(*SKBuiltinBufferT)(nil),        // 5: SKBuiltinBuffer_t
	(*BaseResponse)(nil),            // 6: BaseResponse
}
var file_BatchGetHeadImg_proto_depIdxs = []int32{
	3, // 0: BatchGetHeadImgRequest.baseRequest:type_name -> BaseRequest
	4, // 1: BatchGetHeadImgRequest.userNameList:type_name -> SKBuiltinString_t
	5, // 2: ImgPair.imgBuf:type_name -> SKBuiltinBuffer_t
	4, // 3: ImgPair.username:type_name -> SKBuiltinString_t
	6, // 4: BatchGetHeadImgResponse.baseResponse:type_name -> BaseResponse
	1, // 5: BatchGetHeadImgResponse.imgPairList:type_name -> ImgPair
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_BatchGetHeadImg_proto_init() }
func file_BatchGetHeadImg_proto_init() {
	if File_BatchGetHeadImg_proto != nil {
		return
	}
	file_MicroMsgBase_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BatchGetHeadImg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetHeadImgRequest); i {
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
		file_BatchGetHeadImg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImgPair); i {
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
		file_BatchGetHeadImg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetHeadImgResponse); i {
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
			RawDescriptor: file_BatchGetHeadImg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BatchGetHeadImg_proto_goTypes,
		DependencyIndexes: file_BatchGetHeadImg_proto_depIdxs,
		MessageInfos:      file_BatchGetHeadImg_proto_msgTypes,
	}.Build()
	File_BatchGetHeadImg_proto = out.File
	file_BatchGetHeadImg_proto_rawDesc = nil
	file_BatchGetHeadImg_proto_goTypes = nil
	file_BatchGetHeadImg_proto_depIdxs = nil
}
