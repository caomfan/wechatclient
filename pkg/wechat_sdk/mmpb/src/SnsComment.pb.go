// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: SnsComment.proto

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

type SnsCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *BaseRequest    `protobuf:"bytes,1,req,name=baseRequest" json:"baseRequest,omitempty"`
	Action      *SnsActionGroup `protobuf:"bytes,2,req,name=action" json:"action,omitempty"`
	ClientId    *string         `protobuf:"bytes,3,opt,name=clientId" json:"clientId,omitempty"`
}

func (x *SnsCommentRequest) Reset() {
	*x = SnsCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SnsComment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnsCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnsCommentRequest) ProtoMessage() {}

func (x *SnsCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_SnsComment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnsCommentRequest.ProtoReflect.Descriptor instead.
func (*SnsCommentRequest) Descriptor() ([]byte, []int) {
	return file_SnsComment_proto_rawDescGZIP(), []int{0}
}

func (x *SnsCommentRequest) GetBaseRequest() *BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *SnsCommentRequest) GetAction() *SnsActionGroup {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *SnsCommentRequest) GetClientId() string {
	if x != nil && x.ClientId != nil {
		return *x.ClientId
	}
	return ""
}

type SnsCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *BaseResponse `protobuf:"bytes,1,req,name=baseResponse" json:"baseResponse,omitempty"`
	SnsObject    *SnsObject    `protobuf:"bytes,2,req,name=snsObject" json:"snsObject,omitempty"`
}

func (x *SnsCommentResponse) Reset() {
	*x = SnsCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SnsComment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnsCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnsCommentResponse) ProtoMessage() {}

func (x *SnsCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_SnsComment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnsCommentResponse.ProtoReflect.Descriptor instead.
func (*SnsCommentResponse) Descriptor() ([]byte, []int) {
	return file_SnsComment_proto_rawDescGZIP(), []int{1}
}

func (x *SnsCommentResponse) GetBaseResponse() *BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

func (x *SnsCommentResponse) GetSnsObject() *SnsObject {
	if x != nil {
		return x.SnsObject
	}
	return nil
}

var File_SnsComment_proto protoreflect.FileDescriptor

var file_SnsComment_proto_rawDesc = []byte{
	0x0a, 0x10, 0x53, 0x6e, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x4d, 0x73, 0x67, 0x42, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x6e, 0x73,
	0x42, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x11, 0x53,
	0x6e, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2e, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x27, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x53, 0x6e, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x71, 0x0a, 0x12, 0x53, 0x6e, 0x73, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0c, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x09, 0x73, 0x6e, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x6e, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x09, 0x73,
	0x6e, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x6d, 0x73, 0x67,
}

var (
	file_SnsComment_proto_rawDescOnce sync.Once
	file_SnsComment_proto_rawDescData = file_SnsComment_proto_rawDesc
)

func file_SnsComment_proto_rawDescGZIP() []byte {
	file_SnsComment_proto_rawDescOnce.Do(func() {
		file_SnsComment_proto_rawDescData = protoimpl.X.CompressGZIP(file_SnsComment_proto_rawDescData)
	})
	return file_SnsComment_proto_rawDescData
}

var file_SnsComment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_SnsComment_proto_goTypes = []interface{}{
	(*SnsCommentRequest)(nil),  // 0: SnsCommentRequest
	(*SnsCommentResponse)(nil), // 1: SnsCommentResponse
	(*BaseRequest)(nil),        // 2: BaseRequest
	(*SnsActionGroup)(nil),     // 3: SnsActionGroup
	(*BaseResponse)(nil),       // 4: BaseResponse
	(*SnsObject)(nil),          // 5: SnsObject
}
var file_SnsComment_proto_depIdxs = []int32{
	2, // 0: SnsCommentRequest.baseRequest:type_name -> BaseRequest
	3, // 1: SnsCommentRequest.action:type_name -> SnsActionGroup
	4, // 2: SnsCommentResponse.baseResponse:type_name -> BaseResponse
	5, // 3: SnsCommentResponse.snsObject:type_name -> SnsObject
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_SnsComment_proto_init() }
func file_SnsComment_proto_init() {
	if File_SnsComment_proto != nil {
		return
	}
	file_MicroMsgBase_proto_init()
	file_MicroSnsBase_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SnsComment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnsCommentRequest); i {
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
		file_SnsComment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnsCommentResponse); i {
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
			RawDescriptor: file_SnsComment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SnsComment_proto_goTypes,
		DependencyIndexes: file_SnsComment_proto_depIdxs,
		MessageInfos:      file_SnsComment_proto_msgTypes,
	}.Build()
	File_SnsComment_proto = out.File
	file_SnsComment_proto_rawDesc = nil
	file_SnsComment_proto_goTypes = nil
	file_SnsComment_proto_depIdxs = nil
}
