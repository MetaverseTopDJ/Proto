// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: collect/blind_box_white_address.proto

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

// BlindBoxWhiteAddress 盲盒白名单
type BlindBoxWhiteAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlindBoxId uint64                         `protobuf:"varint,1,opt,name=blind_box_id,json=blindBoxId,proto3" json:"blind_box_id"`
	MemberId   uint64                         `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	Address    string                         `protobuf:"bytes,3,opt,name=address,proto3" json:"address"`
	Signature  string                         `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature"`
	Status     uint64                         `protobuf:"varint,5,opt,name=status,proto3" json:"status"`
	Checked    bool                           `protobuf:"varint,6,opt,name=checked,proto3" json:"checked"`
	CreatedAt  string                         `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt  string                         `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	BlindBox   *BlindBoxWhiteAddress_BlindBox `protobuf:"bytes,9,opt,name=blind_box,json=blindBox,proto3" json:"blind_box"`
}

func (x *BlindBoxWhiteAddress) Reset() {
	*x = BlindBoxWhiteAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_blind_box_white_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlindBoxWhiteAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlindBoxWhiteAddress) ProtoMessage() {}

func (x *BlindBoxWhiteAddress) ProtoReflect() protoreflect.Message {
	mi := &file_collect_blind_box_white_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlindBoxWhiteAddress.ProtoReflect.Descriptor instead.
func (*BlindBoxWhiteAddress) Descriptor() ([]byte, []int) {
	return file_collect_blind_box_white_address_proto_rawDescGZIP(), []int{0}
}

func (x *BlindBoxWhiteAddress) GetBlindBoxId() uint64 {
	if x != nil {
		return x.BlindBoxId
	}
	return 0
}

func (x *BlindBoxWhiteAddress) GetMemberId() uint64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *BlindBoxWhiteAddress) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BlindBoxWhiteAddress) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *BlindBoxWhiteAddress) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *BlindBoxWhiteAddress) GetChecked() bool {
	if x != nil {
		return x.Checked
	}
	return false
}

func (x *BlindBoxWhiteAddress) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *BlindBoxWhiteAddress) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *BlindBoxWhiteAddress) GetBlindBox() *BlindBoxWhiteAddress_BlindBox {
	if x != nil {
		return x.BlindBox
	}
	return nil
}

// BlindBoxWhiteAddressPaginationPost 盲盒白名单分页参数
type BlindBoxWhiteAddressPaginationPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlindBoxId uint64 `protobuf:"varint,1,opt,name=blind_box_id,json=blindBoxId,proto3" json:"blind_box_id"`
	Address    string `protobuf:"bytes,2,opt,name=address,proto3" json:"address"`
	Page       uint64 `protobuf:"varint,3,opt,name=page,proto3" json:"page"`
	Size       uint64 `protobuf:"varint,4,opt,name=size,proto3" json:"size"`
}

func (x *BlindBoxWhiteAddressPaginationPost) Reset() {
	*x = BlindBoxWhiteAddressPaginationPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_blind_box_white_address_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlindBoxWhiteAddressPaginationPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlindBoxWhiteAddressPaginationPost) ProtoMessage() {}

func (x *BlindBoxWhiteAddressPaginationPost) ProtoReflect() protoreflect.Message {
	mi := &file_collect_blind_box_white_address_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlindBoxWhiteAddressPaginationPost.ProtoReflect.Descriptor instead.
func (*BlindBoxWhiteAddressPaginationPost) Descriptor() ([]byte, []int) {
	return file_collect_blind_box_white_address_proto_rawDescGZIP(), []int{1}
}

func (x *BlindBoxWhiteAddressPaginationPost) GetBlindBoxId() uint64 {
	if x != nil {
		return x.BlindBoxId
	}
	return 0
}

func (x *BlindBoxWhiteAddressPaginationPost) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BlindBoxWhiteAddressPaginationPost) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *BlindBoxWhiteAddressPaginationPost) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

// CreateBlindBoxWhiteAddressPost 新增白名单
type CreateBlindBoxWhiteAddressPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlindBoxId uint64 `protobuf:"varint,1,opt,name=blind_box_id,json=blindBoxId,proto3" json:"blind_box_id"`
	Addresses  string `protobuf:"bytes,2,opt,name=addresses,proto3" json:"addresses"`
}

func (x *CreateBlindBoxWhiteAddressPost) Reset() {
	*x = CreateBlindBoxWhiteAddressPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_blind_box_white_address_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlindBoxWhiteAddressPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlindBoxWhiteAddressPost) ProtoMessage() {}

func (x *CreateBlindBoxWhiteAddressPost) ProtoReflect() protoreflect.Message {
	mi := &file_collect_blind_box_white_address_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlindBoxWhiteAddressPost.ProtoReflect.Descriptor instead.
func (*CreateBlindBoxWhiteAddressPost) Descriptor() ([]byte, []int) {
	return file_collect_blind_box_white_address_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBlindBoxWhiteAddressPost) GetBlindBoxId() uint64 {
	if x != nil {
		return x.BlindBoxId
	}
	return 0
}

func (x *CreateBlindBoxWhiteAddressPost) GetAddresses() string {
	if x != nil {
		return x.Addresses
	}
	return ""
}

// RemoveBlindBoxWhiteAddressPost 移除白名单
type RemoveBlindBoxWhiteAddressPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlindBoxId uint64 `protobuf:"varint,1,opt,name=blind_box_id,json=blindBoxId,proto3" json:"blind_box_id"`
	MemberId   uint64 `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id"`
}

func (x *RemoveBlindBoxWhiteAddressPost) Reset() {
	*x = RemoveBlindBoxWhiteAddressPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_blind_box_white_address_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveBlindBoxWhiteAddressPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveBlindBoxWhiteAddressPost) ProtoMessage() {}

func (x *RemoveBlindBoxWhiteAddressPost) ProtoReflect() protoreflect.Message {
	mi := &file_collect_blind_box_white_address_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveBlindBoxWhiteAddressPost.ProtoReflect.Descriptor instead.
func (*RemoveBlindBoxWhiteAddressPost) Descriptor() ([]byte, []int) {
	return file_collect_blind_box_white_address_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveBlindBoxWhiteAddressPost) GetBlindBoxId() uint64 {
	if x != nil {
		return x.BlindBoxId
	}
	return 0
}

func (x *RemoveBlindBoxWhiteAddressPost) GetMemberId() uint64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

// BlindBoxWhiteAddressPaginationResponse 盲盒白名单分页
type BlindBoxWhiteAddressPaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                                                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    *BlindBoxWhiteAddressPaginationResponse_WhiteAddresses `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *BlindBoxWhiteAddressPaginationResponse) Reset() {
	*x = BlindBoxWhiteAddressPaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_blind_box_white_address_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlindBoxWhiteAddressPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlindBoxWhiteAddressPaginationResponse) ProtoMessage() {}

func (x *BlindBoxWhiteAddressPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_collect_blind_box_white_address_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlindBoxWhiteAddressPaginationResponse.ProtoReflect.Descriptor instead.
func (*BlindBoxWhiteAddressPaginationResponse) Descriptor() ([]byte, []int) {
	return file_collect_blind_box_white_address_proto_rawDescGZIP(), []int{4}
}

func (x *BlindBoxWhiteAddressPaginationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BlindBoxWhiteAddressPaginationResponse) GetData() *BlindBoxWhiteAddressPaginationResponse_WhiteAddresses {
	if x != nil {
		return x.Data
	}
	return nil
}

type BlindBoxWhiteAddress_BlindBox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title"`
}

func (x *BlindBoxWhiteAddress_BlindBox) Reset() {
	*x = BlindBoxWhiteAddress_BlindBox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_blind_box_white_address_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlindBoxWhiteAddress_BlindBox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlindBoxWhiteAddress_BlindBox) ProtoMessage() {}

func (x *BlindBoxWhiteAddress_BlindBox) ProtoReflect() protoreflect.Message {
	mi := &file_collect_blind_box_white_address_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlindBoxWhiteAddress_BlindBox.ProtoReflect.Descriptor instead.
func (*BlindBoxWhiteAddress_BlindBox) Descriptor() ([]byte, []int) {
	return file_collect_blind_box_white_address_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BlindBoxWhiteAddress_BlindBox) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BlindBoxWhiteAddress_BlindBox) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type BlindBoxWhiteAddressPaginationResponse_WhiteAddresses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*BlindBoxWhiteAddress `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
	Total int64                   `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
}

func (x *BlindBoxWhiteAddressPaginationResponse_WhiteAddresses) Reset() {
	*x = BlindBoxWhiteAddressPaginationResponse_WhiteAddresses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_blind_box_white_address_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlindBoxWhiteAddressPaginationResponse_WhiteAddresses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlindBoxWhiteAddressPaginationResponse_WhiteAddresses) ProtoMessage() {}

func (x *BlindBoxWhiteAddressPaginationResponse_WhiteAddresses) ProtoReflect() protoreflect.Message {
	mi := &file_collect_blind_box_white_address_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlindBoxWhiteAddressPaginationResponse_WhiteAddresses.ProtoReflect.Descriptor instead.
func (*BlindBoxWhiteAddressPaginationResponse_WhiteAddresses) Descriptor() ([]byte, []int) {
	return file_collect_blind_box_white_address_proto_rawDescGZIP(), []int{4, 0}
}

func (x *BlindBoxWhiteAddressPaginationResponse_WhiteAddresses) GetList() []*BlindBoxWhiteAddress {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *BlindBoxWhiteAddressPaginationResponse_WhiteAddresses) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_collect_blind_box_white_address_proto protoreflect.FileDescriptor

var file_collect_blind_box_white_address_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2f, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x5f,
	0x62, 0x6f, 0x78, 0x5f, 0x77, 0x68, 0x69, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x22, 0xf4, 0x02, 0x0a, 0x14, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x57, 0x68, 0x69,
	0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x6c, 0x69,
	0x6e, 0x64, 0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x43, 0x0a, 0x09, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x5f, 0x62, 0x6f, 0x78, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x42, 0x6c,
	0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x52, 0x08, 0x62, 0x6c, 0x69,
	0x6e, 0x64, 0x42, 0x6f, 0x78, 0x1a, 0x30, 0x0a, 0x08, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f,
	0x78, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x22, 0x42, 0x6c, 0x69, 0x6e,
	0x64, 0x42, 0x6f, 0x78, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0c, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x22, 0x60, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x69, 0x6e,
	0x64, 0x42, 0x6f, 0x78, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x5f, 0x62, 0x6f,
	0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x6c, 0x69, 0x6e,
	0x64, 0x42, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x22, 0x5f, 0x0a, 0x1e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x6c,
	0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x5f,
	0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x6c,
	0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x64, 0x22, 0xf1, 0x01, 0x0a, 0x26, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42,
	0x6f, 0x78, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x52, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x2e, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x57, 0x68, 0x69, 0x74, 0x65,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x59,
	0x0a, 0x0e, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x12, 0x31, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f,
	0x78, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collect_blind_box_white_address_proto_rawDescOnce sync.Once
	file_collect_blind_box_white_address_proto_rawDescData = file_collect_blind_box_white_address_proto_rawDesc
)

func file_collect_blind_box_white_address_proto_rawDescGZIP() []byte {
	file_collect_blind_box_white_address_proto_rawDescOnce.Do(func() {
		file_collect_blind_box_white_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_collect_blind_box_white_address_proto_rawDescData)
	})
	return file_collect_blind_box_white_address_proto_rawDescData
}

var file_collect_blind_box_white_address_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_collect_blind_box_white_address_proto_goTypes = []interface{}{
	(*BlindBoxWhiteAddress)(nil),                                  // 0: collect.BlindBoxWhiteAddress
	(*BlindBoxWhiteAddressPaginationPost)(nil),                    // 1: collect.BlindBoxWhiteAddressPaginationPost
	(*CreateBlindBoxWhiteAddressPost)(nil),                        // 2: collect.CreateBlindBoxWhiteAddressPost
	(*RemoveBlindBoxWhiteAddressPost)(nil),                        // 3: collect.RemoveBlindBoxWhiteAddressPost
	(*BlindBoxWhiteAddressPaginationResponse)(nil),                // 4: collect.BlindBoxWhiteAddressPaginationResponse
	(*BlindBoxWhiteAddress_BlindBox)(nil),                         // 5: collect.BlindBoxWhiteAddress.BlindBox
	(*BlindBoxWhiteAddressPaginationResponse_WhiteAddresses)(nil), // 6: collect.BlindBoxWhiteAddressPaginationResponse.WhiteAddresses
}
var file_collect_blind_box_white_address_proto_depIdxs = []int32{
	5, // 0: collect.BlindBoxWhiteAddress.blind_box:type_name -> collect.BlindBoxWhiteAddress.BlindBox
	6, // 1: collect.BlindBoxWhiteAddressPaginationResponse.data:type_name -> collect.BlindBoxWhiteAddressPaginationResponse.WhiteAddresses
	0, // 2: collect.BlindBoxWhiteAddressPaginationResponse.WhiteAddresses.list:type_name -> collect.BlindBoxWhiteAddress
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_collect_blind_box_white_address_proto_init() }
func file_collect_blind_box_white_address_proto_init() {
	if File_collect_blind_box_white_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collect_blind_box_white_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlindBoxWhiteAddress); i {
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
		file_collect_blind_box_white_address_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlindBoxWhiteAddressPaginationPost); i {
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
		file_collect_blind_box_white_address_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlindBoxWhiteAddressPost); i {
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
		file_collect_blind_box_white_address_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveBlindBoxWhiteAddressPost); i {
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
		file_collect_blind_box_white_address_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlindBoxWhiteAddressPaginationResponse); i {
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
		file_collect_blind_box_white_address_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlindBoxWhiteAddress_BlindBox); i {
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
		file_collect_blind_box_white_address_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlindBoxWhiteAddressPaginationResponse_WhiteAddresses); i {
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
			RawDescriptor: file_collect_blind_box_white_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_collect_blind_box_white_address_proto_goTypes,
		DependencyIndexes: file_collect_blind_box_white_address_proto_depIdxs,
		MessageInfos:      file_collect_blind_box_white_address_proto_msgTypes,
	}.Build()
	File_collect_blind_box_white_address_proto = out.File
	file_collect_blind_box_white_address_proto_rawDesc = nil
	file_collect_blind_box_white_address_proto_goTypes = nil
	file_collect_blind_box_white_address_proto_depIdxs = nil
}
