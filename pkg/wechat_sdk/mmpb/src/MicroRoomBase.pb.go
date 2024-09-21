// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: MicroRoomBase.proto

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

type MemberReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberName *SKBuiltinStringT `protobuf:"bytes,1,req,name=memberName" json:"memberName,omitempty"`
}

func (x *MemberReq) Reset() {
	*x = MemberReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MicroRoomBase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberReq) ProtoMessage() {}

func (x *MemberReq) ProtoReflect() protoreflect.Message {
	mi := &file_MicroRoomBase_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberReq.ProtoReflect.Descriptor instead.
func (*MemberReq) Descriptor() ([]byte, []int) {
	return file_MicroRoomBase_proto_rawDescGZIP(), []int{0}
}

func (x *MemberReq) GetMemberName() *SKBuiltinStringT {
	if x != nil {
		return x.MemberName
	}
	return nil
}

type MemberResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberName      *SKBuiltinStringT `protobuf:"bytes,1,req,name=memberName" json:"memberName,omitempty"`
	MemberStatus    *uint32           `protobuf:"varint,2,req,name=memberStatus" json:"memberStatus,omitempty"`
	NickName        *SKBuiltinStringT `protobuf:"bytes,3,req,name=nickName" json:"nickName,omitempty"`
	PyInitial       *SKBuiltinStringT `protobuf:"bytes,4,req,name=pyInitial" json:"pyInitial,omitempty"`
	QuanPin         *SKBuiltinStringT `protobuf:"bytes,5,req,name=quanPin" json:"quanPin,omitempty"`
	Sex             *int32            `protobuf:"varint,6,req,name=sex" json:"sex,omitempty"`
	Remark          *SKBuiltinStringT `protobuf:"bytes,9,req,name=remark" json:"remark,omitempty"`
	RemarkPYInitial *SKBuiltinStringT `protobuf:"bytes,10,req,name=remarkPYInitial" json:"remarkPYInitial,omitempty"`
	RemarkQuanPin   *SKBuiltinStringT `protobuf:"bytes,11,req,name=remarkQuanPin" json:"remarkQuanPin,omitempty"`
	ContactType     *uint32           `protobuf:"varint,12,req,name=contactType" json:"contactType,omitempty"`
	Province        *string           `protobuf:"bytes,13,opt,name=province" json:"province,omitempty"`
	City            *string           `protobuf:"bytes,14,opt,name=city" json:"city,omitempty"`
	Signature       *string           `protobuf:"bytes,15,opt,name=signature" json:"signature,omitempty"`
	PersonalCard    *uint32           `protobuf:"varint,16,opt,name=personalCard" json:"personalCard,omitempty"`
	VerifyFlag      *uint32           `protobuf:"varint,17,opt,name=verifyFlag" json:"verifyFlag,omitempty"`
	VerifyInfo      *string           `protobuf:"bytes,18,opt,name=verifyInfo" json:"verifyInfo,omitempty"`
	Country         *string           `protobuf:"bytes,19,opt,name=country" json:"country,omitempty"`
}

func (x *MemberResp) Reset() {
	*x = MemberResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MicroRoomBase_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberResp) ProtoMessage() {}

func (x *MemberResp) ProtoReflect() protoreflect.Message {
	mi := &file_MicroRoomBase_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberResp.ProtoReflect.Descriptor instead.
func (*MemberResp) Descriptor() ([]byte, []int) {
	return file_MicroRoomBase_proto_rawDescGZIP(), []int{1}
}

func (x *MemberResp) GetMemberName() *SKBuiltinStringT {
	if x != nil {
		return x.MemberName
	}
	return nil
}

func (x *MemberResp) GetMemberStatus() uint32 {
	if x != nil && x.MemberStatus != nil {
		return *x.MemberStatus
	}
	return 0
}

func (x *MemberResp) GetNickName() *SKBuiltinStringT {
	if x != nil {
		return x.NickName
	}
	return nil
}

func (x *MemberResp) GetPyInitial() *SKBuiltinStringT {
	if x != nil {
		return x.PyInitial
	}
	return nil
}

func (x *MemberResp) GetQuanPin() *SKBuiltinStringT {
	if x != nil {
		return x.QuanPin
	}
	return nil
}

func (x *MemberResp) GetSex() int32 {
	if x != nil && x.Sex != nil {
		return *x.Sex
	}
	return 0
}

func (x *MemberResp) GetRemark() *SKBuiltinStringT {
	if x != nil {
		return x.Remark
	}
	return nil
}

func (x *MemberResp) GetRemarkPYInitial() *SKBuiltinStringT {
	if x != nil {
		return x.RemarkPYInitial
	}
	return nil
}

func (x *MemberResp) GetRemarkQuanPin() *SKBuiltinStringT {
	if x != nil {
		return x.RemarkQuanPin
	}
	return nil
}

func (x *MemberResp) GetContactType() uint32 {
	if x != nil && x.ContactType != nil {
		return *x.ContactType
	}
	return 0
}

func (x *MemberResp) GetProvince() string {
	if x != nil && x.Province != nil {
		return *x.Province
	}
	return ""
}

func (x *MemberResp) GetCity() string {
	if x != nil && x.City != nil {
		return *x.City
	}
	return ""
}

func (x *MemberResp) GetSignature() string {
	if x != nil && x.Signature != nil {
		return *x.Signature
	}
	return ""
}

func (x *MemberResp) GetPersonalCard() uint32 {
	if x != nil && x.PersonalCard != nil {
		return *x.PersonalCard
	}
	return 0
}

func (x *MemberResp) GetVerifyFlag() uint32 {
	if x != nil && x.VerifyFlag != nil {
		return *x.VerifyFlag
	}
	return 0
}

func (x *MemberResp) GetVerifyInfo() string {
	if x != nil && x.VerifyInfo != nil {
		return *x.VerifyInfo
	}
	return ""
}

func (x *MemberResp) GetCountry() string {
	if x != nil && x.Country != nil {
		return *x.Country
	}
	return ""
}

type DelMemberReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberName *SKBuiltinStringT `protobuf:"bytes,1,req,name=memberName" json:"memberName,omitempty"`
}

func (x *DelMemberReq) Reset() {
	*x = DelMemberReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MicroRoomBase_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelMemberReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelMemberReq) ProtoMessage() {}

func (x *DelMemberReq) ProtoReflect() protoreflect.Message {
	mi := &file_MicroRoomBase_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelMemberReq.ProtoReflect.Descriptor instead.
func (*DelMemberReq) Descriptor() ([]byte, []int) {
	return file_MicroRoomBase_proto_rawDescGZIP(), []int{2}
}

func (x *DelMemberReq) GetMemberName() *SKBuiltinStringT {
	if x != nil {
		return x.MemberName
	}
	return nil
}

type DelMemberResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberName *SKBuiltinStringT `protobuf:"bytes,1,req,name=memberName" json:"memberName,omitempty"`
}

func (x *DelMemberResp) Reset() {
	*x = DelMemberResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MicroRoomBase_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelMemberResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelMemberResp) ProtoMessage() {}

func (x *DelMemberResp) ProtoReflect() protoreflect.Message {
	mi := &file_MicroRoomBase_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelMemberResp.ProtoReflect.Descriptor instead.
func (*DelMemberResp) Descriptor() ([]byte, []int) {
	return file_MicroRoomBase_proto_rawDescGZIP(), []int{3}
}

func (x *DelMemberResp) GetMemberName() *SKBuiltinStringT {
	if x != nil {
		return x.MemberName
	}
	return nil
}

var File_MicroRoomBase_proto protoreflect.FileDescriptor

var file_MicroRoomBase_proto_rawDesc = []byte{
	0x0a, 0x13, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x4d, 0x73, 0x67, 0x42,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x09, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x32, 0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x4b, 0x42,
	0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x52, 0x0a,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x98, 0x05, 0x0a, 0x0a, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x0a, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x53, 0x4b, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x74, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x2e, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x4b, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x30, 0x0a, 0x09, 0x70, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x04,
	0x20, 0x02, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x4b, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x52, 0x09, 0x70, 0x79, 0x49, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x07, 0x71, 0x75, 0x61, 0x6e, 0x50, 0x69, 0x6e, 0x18, 0x05,
	0x20, 0x02, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x4b, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x52, 0x07, 0x71, 0x75, 0x61, 0x6e, 0x50, 0x69,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x06, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03,
	0x73, 0x65, 0x78, 0x12, 0x2a, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x09, 0x20,
	0x02, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x4b, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x3c, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x50, 0x59, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x18, 0x0a, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x4b, 0x42, 0x75, 0x69,
	0x6c, 0x74, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x52, 0x0f, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x50, 0x59, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x38, 0x0a,
	0x0d, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x51, 0x75, 0x61, 0x6e, 0x50, 0x69, 0x6e, 0x18, 0x0b,
	0x20, 0x02, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x4b, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x51, 0x75, 0x61, 0x6e, 0x50, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x43, 0x61, 0x72, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x43, 0x61, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x42, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x32, 0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x4b, 0x42, 0x75,
	0x69, 0x6c, 0x74, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x52, 0x0a, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x43, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x0a, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x53, 0x4b, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x74, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x3b, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x73, 0x67,
}

var (
	file_MicroRoomBase_proto_rawDescOnce sync.Once
	file_MicroRoomBase_proto_rawDescData = file_MicroRoomBase_proto_rawDesc
)

func file_MicroRoomBase_proto_rawDescGZIP() []byte {
	file_MicroRoomBase_proto_rawDescOnce.Do(func() {
		file_MicroRoomBase_proto_rawDescData = protoimpl.X.CompressGZIP(file_MicroRoomBase_proto_rawDescData)
	})
	return file_MicroRoomBase_proto_rawDescData
}

var file_MicroRoomBase_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_MicroRoomBase_proto_goTypes = []interface{}{
	(*MemberReq)(nil),        // 0: MemberReq
	(*MemberResp)(nil),       // 1: MemberResp
	(*DelMemberReq)(nil),     // 2: DelMemberReq
	(*DelMemberResp)(nil),    // 3: DelMemberResp
	(*SKBuiltinStringT)(nil), // 4: SKBuiltinString_t
}
var file_MicroRoomBase_proto_depIdxs = []int32{
	4,  // 0: MemberReq.memberName:type_name -> SKBuiltinString_t
	4,  // 1: MemberResp.memberName:type_name -> SKBuiltinString_t
	4,  // 2: MemberResp.nickName:type_name -> SKBuiltinString_t
	4,  // 3: MemberResp.pyInitial:type_name -> SKBuiltinString_t
	4,  // 4: MemberResp.quanPin:type_name -> SKBuiltinString_t
	4,  // 5: MemberResp.remark:type_name -> SKBuiltinString_t
	4,  // 6: MemberResp.remarkPYInitial:type_name -> SKBuiltinString_t
	4,  // 7: MemberResp.remarkQuanPin:type_name -> SKBuiltinString_t
	4,  // 8: DelMemberReq.memberName:type_name -> SKBuiltinString_t
	4,  // 9: DelMemberResp.memberName:type_name -> SKBuiltinString_t
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_MicroRoomBase_proto_init() }
func file_MicroRoomBase_proto_init() {
	if File_MicroRoomBase_proto != nil {
		return
	}
	file_MicroMsgBase_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MicroRoomBase_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberReq); i {
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
		file_MicroRoomBase_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberResp); i {
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
		file_MicroRoomBase_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelMemberReq); i {
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
		file_MicroRoomBase_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelMemberResp); i {
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
			RawDescriptor: file_MicroRoomBase_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MicroRoomBase_proto_goTypes,
		DependencyIndexes: file_MicroRoomBase_proto_depIdxs,
		MessageInfos:      file_MicroRoomBase_proto_msgTypes,
	}.Build()
	File_MicroRoomBase_proto = out.File
	file_MicroRoomBase_proto_rawDesc = nil
	file_MicroRoomBase_proto_goTypes = nil
	file_MicroRoomBase_proto_depIdxs = nil
}
