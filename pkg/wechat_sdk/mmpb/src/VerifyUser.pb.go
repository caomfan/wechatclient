// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: VerifyUser.proto

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

type VerifyUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value                 *string           `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	VerifyUserTicket      *string           `protobuf:"bytes,2,opt,name=verifyUserTicket" json:"verifyUserTicket,omitempty"`
	AntispamTicket        *string           `protobuf:"bytes,3,opt,name=antispamTicket" json:"antispamTicket,omitempty"`
	FriendFlag            *uint32           `protobuf:"varint,4,opt,name=friendFlag" json:"friendFlag,omitempty"`
	ChatRoomUserName      *string           `protobuf:"bytes,5,opt,name=chatRoomUserName" json:"chatRoomUserName,omitempty"`
	SourceUserName        *string           `protobuf:"bytes,6,opt,name=sourceUserName" json:"sourceUserName,omitempty"`
	SourceNickName        *string           `protobuf:"bytes,7,opt,name=sourceNickName" json:"sourceNickName,omitempty"`
	ScanQrcodeFromScene   *uint32           `protobuf:"varint,8,opt,name=scanQrcodeFromScene" json:"scanQrcodeFromScene,omitempty"`
	ReportInfo            *string           `protobuf:"bytes,9,opt,name=reportInfo" json:"reportInfo,omitempty"`
	ShareCardForwardLevel *uint32           `protobuf:"varint,10,opt,name=shareCardForwardLevel" json:"shareCardForwardLevel,omitempty"`
	ShareCardForwardInfo  *SKBuiltinBufferT `protobuf:"bytes,11,opt,name=shareCardForwardInfo" json:"shareCardForwardInfo,omitempty"`
	OuterUrl              *string           `protobuf:"bytes,12,opt,name=outerUrl" json:"outerUrl,omitempty"`
	SubScene              *uint32           `protobuf:"varint,13,opt,name=subScene" json:"subScene,omitempty"`
	BizReportInfo         *SKBuiltinBufferT `protobuf:"bytes,14,opt,name=bizReportInfo" json:"bizReportInfo,omitempty"`
}

func (x *VerifyUser) Reset() {
	*x = VerifyUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_VerifyUser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyUser) ProtoMessage() {}

func (x *VerifyUser) ProtoReflect() protoreflect.Message {
	mi := &file_VerifyUser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyUser.ProtoReflect.Descriptor instead.
func (*VerifyUser) Descriptor() ([]byte, []int) {
	return file_VerifyUser_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyUser) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

func (x *VerifyUser) GetVerifyUserTicket() string {
	if x != nil && x.VerifyUserTicket != nil {
		return *x.VerifyUserTicket
	}
	return ""
}

func (x *VerifyUser) GetAntispamTicket() string {
	if x != nil && x.AntispamTicket != nil {
		return *x.AntispamTicket
	}
	return ""
}

func (x *VerifyUser) GetFriendFlag() uint32 {
	if x != nil && x.FriendFlag != nil {
		return *x.FriendFlag
	}
	return 0
}

func (x *VerifyUser) GetChatRoomUserName() string {
	if x != nil && x.ChatRoomUserName != nil {
		return *x.ChatRoomUserName
	}
	return ""
}

func (x *VerifyUser) GetSourceUserName() string {
	if x != nil && x.SourceUserName != nil {
		return *x.SourceUserName
	}
	return ""
}

func (x *VerifyUser) GetSourceNickName() string {
	if x != nil && x.SourceNickName != nil {
		return *x.SourceNickName
	}
	return ""
}

func (x *VerifyUser) GetScanQrcodeFromScene() uint32 {
	if x != nil && x.ScanQrcodeFromScene != nil {
		return *x.ScanQrcodeFromScene
	}
	return 0
}

func (x *VerifyUser) GetReportInfo() string {
	if x != nil && x.ReportInfo != nil {
		return *x.ReportInfo
	}
	return ""
}

func (x *VerifyUser) GetShareCardForwardLevel() uint32 {
	if x != nil && x.ShareCardForwardLevel != nil {
		return *x.ShareCardForwardLevel
	}
	return 0
}

func (x *VerifyUser) GetShareCardForwardInfo() *SKBuiltinBufferT {
	if x != nil {
		return x.ShareCardForwardInfo
	}
	return nil
}

func (x *VerifyUser) GetOuterUrl() string {
	if x != nil && x.OuterUrl != nil {
		return *x.OuterUrl
	}
	return ""
}

func (x *VerifyUser) GetSubScene() uint32 {
	if x != nil && x.SubScene != nil {
		return *x.SubScene
	}
	return 0
}

func (x *VerifyUser) GetBizReportInfo() *SKBuiltinBufferT {
	if x != nil {
		return x.BizReportInfo
	}
	return nil
}

type VerifyUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (x *VerifyUserInfo) Reset() {
	*x = VerifyUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_VerifyUser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyUserInfo) ProtoMessage() {}

func (x *VerifyUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_VerifyUser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyUserInfo.ProtoReflect.Descriptor instead.
func (*VerifyUserInfo) Descriptor() ([]byte, []int) {
	return file_VerifyUser_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyUserInfo) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *VerifyUserInfo) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

type VerifyUserSpamInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block         *int32  `protobuf:"varint,1,opt,name=block" json:"block,omitempty"`
	WordingTitle  *string `protobuf:"bytes,2,opt,name=wordingTitle" json:"wordingTitle,omitempty"`
	WordingDetail *string `protobuf:"bytes,3,opt,name=wordingDetail" json:"wordingDetail,omitempty"`
	WordingUrl    *string `protobuf:"bytes,4,opt,name=wordingUrl" json:"wordingUrl,omitempty"`
}

func (x *VerifyUserSpamInfo) Reset() {
	*x = VerifyUserSpamInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_VerifyUser_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyUserSpamInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyUserSpamInfo) ProtoMessage() {}

func (x *VerifyUserSpamInfo) ProtoReflect() protoreflect.Message {
	mi := &file_VerifyUser_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyUserSpamInfo.ProtoReflect.Descriptor instead.
func (*VerifyUserSpamInfo) Descriptor() ([]byte, []int) {
	return file_VerifyUser_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyUserSpamInfo) GetBlock() int32 {
	if x != nil && x.Block != nil {
		return *x.Block
	}
	return 0
}

func (x *VerifyUserSpamInfo) GetWordingTitle() string {
	if x != nil && x.WordingTitle != nil {
		return *x.WordingTitle
	}
	return ""
}

func (x *VerifyUserSpamInfo) GetWordingDetail() string {
	if x != nil && x.WordingDetail != nil {
		return *x.WordingDetail
	}
	return ""
}

func (x *VerifyUserSpamInfo) GetWordingUrl() string {
	if x != nil && x.WordingUrl != nil {
		return *x.WordingUrl
	}
	return ""
}

type VerifyUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest         *BaseRequest      `protobuf:"bytes,1,req,name=baseRequest" json:"baseRequest,omitempty"`
	Opcode              *uint32           `protobuf:"varint,2,req,name=opcode" json:"opcode,omitempty"`
	VerifyUserListSize  *uint32           `protobuf:"varint,3,req,name=verifyUserListSize" json:"verifyUserListSize,omitempty"`
	VerifyUserList      []*VerifyUser     `protobuf:"bytes,4,rep,name=verifyUserList" json:"verifyUserList,omitempty"`
	VerifyContent       *string           `protobuf:"bytes,5,opt,name=verifyContent" json:"verifyContent,omitempty"`
	SceneListCount      *uint32           `protobuf:"varint,6,opt,name=sceneListCount" json:"sceneListCount,omitempty"`
	SceneList           []uint32          `protobuf:"varint,7,rep,packed,name=sceneList" json:"sceneList,omitempty"`
	VerifyInfoListCount *uint32           `protobuf:"varint,8,opt,name=verifyInfoListCount" json:"verifyInfoListCount,omitempty"`
	VerifyInfoList      []*VerifyUserInfo `protobuf:"bytes,9,rep,name=verifyInfoList" json:"verifyInfoList,omitempty"`
	ClientCheckData     *SKBuiltinBufferT `protobuf:"bytes,10,opt,name=clientCheckData" json:"clientCheckData,omitempty"`
	ExtSpamInfo         *SKBuiltinBufferT `protobuf:"bytes,11,opt,name=extSpamInfo" json:"extSpamInfo,omitempty"`
	NeedConfirm         *uint32           `protobuf:"varint,12,opt,name=needConfirm" json:"needConfirm,omitempty"`
	Ctx                 *string           `protobuf:"bytes,13,opt,name=ctx" json:"ctx,omitempty"`
}

func (x *VerifyUserRequest) Reset() {
	*x = VerifyUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_VerifyUser_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyUserRequest) ProtoMessage() {}

func (x *VerifyUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_VerifyUser_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyUserRequest.ProtoReflect.Descriptor instead.
func (*VerifyUserRequest) Descriptor() ([]byte, []int) {
	return file_VerifyUser_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyUserRequest) GetBaseRequest() *BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *VerifyUserRequest) GetOpcode() uint32 {
	if x != nil && x.Opcode != nil {
		return *x.Opcode
	}
	return 0
}

func (x *VerifyUserRequest) GetVerifyUserListSize() uint32 {
	if x != nil && x.VerifyUserListSize != nil {
		return *x.VerifyUserListSize
	}
	return 0
}

func (x *VerifyUserRequest) GetVerifyUserList() []*VerifyUser {
	if x != nil {
		return x.VerifyUserList
	}
	return nil
}

func (x *VerifyUserRequest) GetVerifyContent() string {
	if x != nil && x.VerifyContent != nil {
		return *x.VerifyContent
	}
	return ""
}

func (x *VerifyUserRequest) GetSceneListCount() uint32 {
	if x != nil && x.SceneListCount != nil {
		return *x.SceneListCount
	}
	return 0
}

func (x *VerifyUserRequest) GetSceneList() []uint32 {
	if x != nil {
		return x.SceneList
	}
	return nil
}

func (x *VerifyUserRequest) GetVerifyInfoListCount() uint32 {
	if x != nil && x.VerifyInfoListCount != nil {
		return *x.VerifyInfoListCount
	}
	return 0
}

func (x *VerifyUserRequest) GetVerifyInfoList() []*VerifyUserInfo {
	if x != nil {
		return x.VerifyInfoList
	}
	return nil
}

func (x *VerifyUserRequest) GetClientCheckData() *SKBuiltinBufferT {
	if x != nil {
		return x.ClientCheckData
	}
	return nil
}

func (x *VerifyUserRequest) GetExtSpamInfo() *SKBuiltinBufferT {
	if x != nil {
		return x.ExtSpamInfo
	}
	return nil
}

func (x *VerifyUserRequest) GetNeedConfirm() uint32 {
	if x != nil && x.NeedConfirm != nil {
		return *x.NeedConfirm
	}
	return 0
}

func (x *VerifyUserRequest) GetCtx() string {
	if x != nil && x.Ctx != nil {
		return *x.Ctx
	}
	return ""
}

type VerifyUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse       *BaseResponse       `protobuf:"bytes,1,req,name=baseResponse" json:"baseResponse,omitempty"`
	UserName           *string             `protobuf:"bytes,2,opt,name=userName" json:"userName,omitempty"`
	EncryptUserName    *string             `protobuf:"bytes,3,opt,name=encryptUserName" json:"encryptUserName,omitempty"`
	VerifyUserSpamInfo *VerifyUserSpamInfo `protobuf:"bytes,4,opt,name=verifyUserSpamInfo" json:"verifyUserSpamInfo,omitempty"`
	Ctx                *string             `protobuf:"bytes,5,opt,name=ctx" json:"ctx,omitempty"`
}

func (x *VerifyUserResponse) Reset() {
	*x = VerifyUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_VerifyUser_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyUserResponse) ProtoMessage() {}

func (x *VerifyUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_VerifyUser_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyUserResponse.ProtoReflect.Descriptor instead.
func (*VerifyUserResponse) Descriptor() ([]byte, []int) {
	return file_VerifyUser_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyUserResponse) GetBaseResponse() *BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

func (x *VerifyUserResponse) GetUserName() string {
	if x != nil && x.UserName != nil {
		return *x.UserName
	}
	return ""
}

func (x *VerifyUserResponse) GetEncryptUserName() string {
	if x != nil && x.EncryptUserName != nil {
		return *x.EncryptUserName
	}
	return ""
}

func (x *VerifyUserResponse) GetVerifyUserSpamInfo() *VerifyUserSpamInfo {
	if x != nil {
		return x.VerifyUserSpamInfo
	}
	return nil
}

func (x *VerifyUserResponse) GetCtx() string {
	if x != nil && x.Ctx != nil {
		return *x.Ctx
	}
	return ""
}

var File_VerifyUser_proto protoreflect.FileDescriptor

var file_VerifyUser_proto_rawDesc = []byte{
	0x0a, 0x10, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x4d, 0x73, 0x67, 0x42, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x04, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x6e, 0x74, 0x69, 0x73,
	0x70, 0x61, 0x6d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x70, 0x61, 0x6d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x12,
	0x2a, 0x0a, 0x10, 0x63, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x68, 0x61, 0x74, 0x52,
	0x6f, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x69, 0x63,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x73,
	0x63, 0x61, 0x6e, 0x51, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x73, 0x63, 0x61, 0x6e, 0x51, 0x72,
	0x63, 0x6f, 0x64, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a,
	0x15, 0x73, 0x68, 0x61, 0x72, 0x65, 0x43, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x43, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x46, 0x0a, 0x14, 0x73, 0x68, 0x61, 0x72, 0x65, 0x43, 0x61, 0x72, 0x64,
	0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x53, 0x4b, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x42, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x5f, 0x74, 0x52, 0x14, 0x73, 0x68, 0x61, 0x72, 0x65, 0x43, 0x61, 0x72, 0x64,
	0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x75, 0x62, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x62, 0x69, 0x7a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x4b, 0x42,
	0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x74, 0x52, 0x0d,
	0x62, 0x69, 0x7a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x38, 0x0a,
	0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70, 0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x22, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x77, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1e,
	0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x22, 0xc3,
	0x04, 0x0a, 0x11, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x12,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x33, 0x0a, 0x0e,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x0e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x0d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x63, 0x65, 0x6e, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0e, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x20, 0x0a, 0x09, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0d, 0x42, 0x02, 0x10, 0x01, 0x52, 0x09, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x30, 0x0a, 0x13, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x0e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x4b, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69,
	0x6e, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x74, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x0b, 0x65, 0x78,
	0x74, 0x53, 0x70, 0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x53, 0x4b, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x5f, 0x74, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x53, 0x70, 0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6e, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x74, 0x78, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x63, 0x74, 0x78, 0x22, 0xe4, 0x01, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0c, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x70, 0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70, 0x61,
	0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x70, 0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x74, 0x78,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x74, 0x78, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x3b, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x73, 0x67,
}

var (
	file_VerifyUser_proto_rawDescOnce sync.Once
	file_VerifyUser_proto_rawDescData = file_VerifyUser_proto_rawDesc
)

func file_VerifyUser_proto_rawDescGZIP() []byte {
	file_VerifyUser_proto_rawDescOnce.Do(func() {
		file_VerifyUser_proto_rawDescData = protoimpl.X.CompressGZIP(file_VerifyUser_proto_rawDescData)
	})
	return file_VerifyUser_proto_rawDescData
}

var file_VerifyUser_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_VerifyUser_proto_goTypes = []interface{}{
	(*VerifyUser)(nil),         // 0: VerifyUser
	(*VerifyUserInfo)(nil),     // 1: VerifyUserInfo
	(*VerifyUserSpamInfo)(nil), // 2: VerifyUserSpamInfo
	(*VerifyUserRequest)(nil),  // 3: VerifyUserRequest
	(*VerifyUserResponse)(nil), // 4: VerifyUserResponse
	(*SKBuiltinBufferT)(nil),   // 5: SKBuiltinBuffer_t
	(*BaseRequest)(nil),        // 6: BaseRequest
	(*BaseResponse)(nil),       // 7: BaseResponse
}
var file_VerifyUser_proto_depIdxs = []int32{
	5, // 0: VerifyUser.shareCardForwardInfo:type_name -> SKBuiltinBuffer_t
	5, // 1: VerifyUser.bizReportInfo:type_name -> SKBuiltinBuffer_t
	6, // 2: VerifyUserRequest.baseRequest:type_name -> BaseRequest
	0, // 3: VerifyUserRequest.verifyUserList:type_name -> VerifyUser
	1, // 4: VerifyUserRequest.verifyInfoList:type_name -> VerifyUserInfo
	5, // 5: VerifyUserRequest.clientCheckData:type_name -> SKBuiltinBuffer_t
	5, // 6: VerifyUserRequest.extSpamInfo:type_name -> SKBuiltinBuffer_t
	7, // 7: VerifyUserResponse.baseResponse:type_name -> BaseResponse
	2, // 8: VerifyUserResponse.verifyUserSpamInfo:type_name -> VerifyUserSpamInfo
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_VerifyUser_proto_init() }
func file_VerifyUser_proto_init() {
	if File_VerifyUser_proto != nil {
		return
	}
	file_MicroMsgBase_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_VerifyUser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyUser); i {
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
		file_VerifyUser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyUserInfo); i {
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
		file_VerifyUser_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyUserSpamInfo); i {
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
		file_VerifyUser_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyUserRequest); i {
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
		file_VerifyUser_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyUserResponse); i {
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
			RawDescriptor: file_VerifyUser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_VerifyUser_proto_goTypes,
		DependencyIndexes: file_VerifyUser_proto_depIdxs,
		MessageInfos:      file_VerifyUser_proto_msgTypes,
	}.Build()
	File_VerifyUser_proto = out.File
	file_VerifyUser_proto_rawDesc = nil
	file_VerifyUser_proto_goTypes = nil
	file_VerifyUser_proto_depIdxs = nil
}
