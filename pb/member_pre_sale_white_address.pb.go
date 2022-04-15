// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: member/member_pre_sale_white_address.proto

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

// MemberPreSaleWhiteAddress 预售白名单
type MemberPreSaleWhiteAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	ActivityId uint64 `protobuf:"varint,2,opt,name=activity_id,json=activityId,proto3" json:"activity_id"`
	CardId     uint64 `protobuf:"varint,3,opt,name=card_id,json=cardId,proto3" json:"card_id"`
	Sign       string `protobuf:"bytes,4,opt,name=sign,proto3" json:"sign"`
	Address    string `protobuf:"bytes,5,opt,name=address,proto3" json:"address"`
	Status     uint64 `protobuf:"varint,6,opt,name=status,proto3" json:"status"`
	Checked    bool   `protobuf:"varint,7,opt,name=checked,proto3" json:"checked"`
	CreatedAt  uint64 `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt  uint64 `protobuf:"varint,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *MemberPreSaleWhiteAddress) Reset() {
	*x = MemberPreSaleWhiteAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_pre_sale_white_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPreSaleWhiteAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPreSaleWhiteAddress) ProtoMessage() {}

func (x *MemberPreSaleWhiteAddress) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_pre_sale_white_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPreSaleWhiteAddress.ProtoReflect.Descriptor instead.
func (*MemberPreSaleWhiteAddress) Descriptor() ([]byte, []int) {
	return file_member_member_pre_sale_white_address_proto_rawDescGZIP(), []int{0}
}

func (x *MemberPreSaleWhiteAddress) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberPreSaleWhiteAddress) GetActivityId() uint64 {
	if x != nil {
		return x.ActivityId
	}
	return 0
}

func (x *MemberPreSaleWhiteAddress) GetCardId() uint64 {
	if x != nil {
		return x.CardId
	}
	return 0
}

func (x *MemberPreSaleWhiteAddress) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *MemberPreSaleWhiteAddress) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MemberPreSaleWhiteAddress) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MemberPreSaleWhiteAddress) GetChecked() bool {
	if x != nil {
		return x.Checked
	}
	return false
}

func (x *MemberPreSaleWhiteAddress) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *MemberPreSaleWhiteAddress) GetUpdatedAt() uint64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

// CreateMemberPreSaleWhiteAddressPost 创建预售白名单
type CreateMemberPreSaleWhiteAddressPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId uint64   `protobuf:"varint,1,opt,name=activity_id,json=activityId,proto3" json:"activity_id"`
	CardId     []uint64 `protobuf:"varint,2,rep,packed,name=card_id,json=cardId,proto3" json:"card_id"`
	Addresses  []string `protobuf:"bytes,3,rep,name=addresses,proto3" json:"addresses"`
}

func (x *CreateMemberPreSaleWhiteAddressPost) Reset() {
	*x = CreateMemberPreSaleWhiteAddressPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_pre_sale_white_address_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMemberPreSaleWhiteAddressPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMemberPreSaleWhiteAddressPost) ProtoMessage() {}

func (x *CreateMemberPreSaleWhiteAddressPost) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_pre_sale_white_address_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMemberPreSaleWhiteAddressPost.ProtoReflect.Descriptor instead.
func (*CreateMemberPreSaleWhiteAddressPost) Descriptor() ([]byte, []int) {
	return file_member_member_pre_sale_white_address_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMemberPreSaleWhiteAddressPost) GetActivityId() uint64 {
	if x != nil {
		return x.ActivityId
	}
	return 0
}

func (x *CreateMemberPreSaleWhiteAddressPost) GetCardId() []uint64 {
	if x != nil {
		return x.CardId
	}
	return nil
}

func (x *CreateMemberPreSaleWhiteAddressPost) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

// UpdateMemberPreSaleWhiteAddressPost 更新预售白名单
type UpdateMemberPreSaleWhiteAddressPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	ActivityId uint64   `protobuf:"varint,2,opt,name=activity_id,json=activityId,proto3" json:"activity_id"`
	CardId     []uint64 `protobuf:"varint,3,rep,packed,name=card_id,json=cardId,proto3" json:"card_id"`
	Address    string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address"`
}

func (x *UpdateMemberPreSaleWhiteAddressPost) Reset() {
	*x = UpdateMemberPreSaleWhiteAddressPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_pre_sale_white_address_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMemberPreSaleWhiteAddressPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMemberPreSaleWhiteAddressPost) ProtoMessage() {}

func (x *UpdateMemberPreSaleWhiteAddressPost) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_pre_sale_white_address_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMemberPreSaleWhiteAddressPost.ProtoReflect.Descriptor instead.
func (*UpdateMemberPreSaleWhiteAddressPost) Descriptor() ([]byte, []int) {
	return file_member_member_pre_sale_white_address_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateMemberPreSaleWhiteAddressPost) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateMemberPreSaleWhiteAddressPost) GetActivityId() uint64 {
	if x != nil {
		return x.ActivityId
	}
	return 0
}

func (x *UpdateMemberPreSaleWhiteAddressPost) GetCardId() []uint64 {
	if x != nil {
		return x.CardId
	}
	return nil
}

func (x *UpdateMemberPreSaleWhiteAddressPost) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// MemberPreSaleWhiteAddressListResponse 返回白名单列表
type MemberPreSaleWhiteAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                     `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    *MemberPreSaleWhiteAddress `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *MemberPreSaleWhiteAddressResponse) Reset() {
	*x = MemberPreSaleWhiteAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_pre_sale_white_address_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPreSaleWhiteAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPreSaleWhiteAddressResponse) ProtoMessage() {}

func (x *MemberPreSaleWhiteAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_pre_sale_white_address_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPreSaleWhiteAddressResponse.ProtoReflect.Descriptor instead.
func (*MemberPreSaleWhiteAddressResponse) Descriptor() ([]byte, []int) {
	return file_member_member_pre_sale_white_address_proto_rawDescGZIP(), []int{3}
}

func (x *MemberPreSaleWhiteAddressResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MemberPreSaleWhiteAddressResponse) GetData() *MemberPreSaleWhiteAddress {
	if x != nil {
		return x.Data
	}
	return nil
}

// MemberPreSaleWhiteAddressListResponse 返回白名单列表
type MemberPreSaleWhiteAddressListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                       `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    []*MemberPreSaleWhiteAddress `protobuf:"bytes,2,rep,name=data,proto3" json:"data"`
}

func (x *MemberPreSaleWhiteAddressListResponse) Reset() {
	*x = MemberPreSaleWhiteAddressListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_pre_sale_white_address_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPreSaleWhiteAddressListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPreSaleWhiteAddressListResponse) ProtoMessage() {}

func (x *MemberPreSaleWhiteAddressListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_pre_sale_white_address_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPreSaleWhiteAddressListResponse.ProtoReflect.Descriptor instead.
func (*MemberPreSaleWhiteAddressListResponse) Descriptor() ([]byte, []int) {
	return file_member_member_pre_sale_white_address_proto_rawDescGZIP(), []int{4}
}

func (x *MemberPreSaleWhiteAddressListResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MemberPreSaleWhiteAddressListResponse) GetData() []*MemberPreSaleWhiteAddress {
	if x != nil {
		return x.Data
	}
	return nil
}

// MemberPreSaleWhiteAddressPaginationResponse 返回白名单分页
type MemberPreSaleWhiteAddressPaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                                                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    *MemberPreSaleWhiteAddressPaginationResponse_Addresses `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *MemberPreSaleWhiteAddressPaginationResponse) Reset() {
	*x = MemberPreSaleWhiteAddressPaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_pre_sale_white_address_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPreSaleWhiteAddressPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPreSaleWhiteAddressPaginationResponse) ProtoMessage() {}

func (x *MemberPreSaleWhiteAddressPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_pre_sale_white_address_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPreSaleWhiteAddressPaginationResponse.ProtoReflect.Descriptor instead.
func (*MemberPreSaleWhiteAddressPaginationResponse) Descriptor() ([]byte, []int) {
	return file_member_member_pre_sale_white_address_proto_rawDescGZIP(), []int{5}
}

func (x *MemberPreSaleWhiteAddressPaginationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MemberPreSaleWhiteAddressPaginationResponse) GetData() *MemberPreSaleWhiteAddressPaginationResponse_Addresses {
	if x != nil {
		return x.Data
	}
	return nil
}

type MemberPreSaleWhiteAddressPaginationResponse_Addresses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*MemberPreSaleWhiteAddress `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
	Total int64                        `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
}

func (x *MemberPreSaleWhiteAddressPaginationResponse_Addresses) Reset() {
	*x = MemberPreSaleWhiteAddressPaginationResponse_Addresses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_pre_sale_white_address_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPreSaleWhiteAddressPaginationResponse_Addresses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPreSaleWhiteAddressPaginationResponse_Addresses) ProtoMessage() {}

func (x *MemberPreSaleWhiteAddressPaginationResponse_Addresses) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_pre_sale_white_address_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPreSaleWhiteAddressPaginationResponse_Addresses.ProtoReflect.Descriptor instead.
func (*MemberPreSaleWhiteAddressPaginationResponse_Addresses) Descriptor() ([]byte, []int) {
	return file_member_member_pre_sale_white_address_proto_rawDescGZIP(), []int{5, 0}
}

func (x *MemberPreSaleWhiteAddressPaginationResponse_Addresses) GetList() []*MemberPreSaleWhiteAddress {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *MemberPreSaleWhiteAddressPaginationResponse_Addresses) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_member_member_pre_sale_white_address_proto protoreflect.FileDescriptor

var file_member_member_pre_sale_white_address_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x70, 0x72, 0x65, 0x5f, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x77, 0x68, 0x69, 0x74, 0x65, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x83, 0x02, 0x0a, 0x19, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50,
	0x72, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x67, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7d, 0x0a, 0x23, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x65, 0x53, 0x61, 0x6c,
	0x65, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x06, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x23, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x65, 0x53, 0x61, 0x6c,
	0x65, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x06, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x74, 0x0a, 0x21, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50,
	0x72, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x50, 0x72, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x78, 0x0a, 0x25, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x57, 0x68, 0x69, 0x74,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x35,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x65, 0x53,
	0x61, 0x6c, 0x65, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xf4, 0x01, 0x0a, 0x2b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x50, 0x72, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x51, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x65,
	0x53, 0x61, 0x6c, 0x65, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x58, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12,
	0x35, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x65,
	0x53, 0x61, 0x6c, 0x65, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_member_member_pre_sale_white_address_proto_rawDescOnce sync.Once
	file_member_member_pre_sale_white_address_proto_rawDescData = file_member_member_pre_sale_white_address_proto_rawDesc
)

func file_member_member_pre_sale_white_address_proto_rawDescGZIP() []byte {
	file_member_member_pre_sale_white_address_proto_rawDescOnce.Do(func() {
		file_member_member_pre_sale_white_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_member_member_pre_sale_white_address_proto_rawDescData)
	})
	return file_member_member_pre_sale_white_address_proto_rawDescData
}

var file_member_member_pre_sale_white_address_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_member_member_pre_sale_white_address_proto_goTypes = []interface{}{
	(*MemberPreSaleWhiteAddress)(nil),                             // 0: member.MemberPreSaleWhiteAddress
	(*CreateMemberPreSaleWhiteAddressPost)(nil),                   // 1: member.CreateMemberPreSaleWhiteAddressPost
	(*UpdateMemberPreSaleWhiteAddressPost)(nil),                   // 2: member.UpdateMemberPreSaleWhiteAddressPost
	(*MemberPreSaleWhiteAddressResponse)(nil),                     // 3: member.MemberPreSaleWhiteAddressResponse
	(*MemberPreSaleWhiteAddressListResponse)(nil),                 // 4: member.MemberPreSaleWhiteAddressListResponse
	(*MemberPreSaleWhiteAddressPaginationResponse)(nil),           // 5: member.MemberPreSaleWhiteAddressPaginationResponse
	(*MemberPreSaleWhiteAddressPaginationResponse_Addresses)(nil), // 6: member.MemberPreSaleWhiteAddressPaginationResponse.Addresses
}
var file_member_member_pre_sale_white_address_proto_depIdxs = []int32{
	0, // 0: member.MemberPreSaleWhiteAddressResponse.data:type_name -> member.MemberPreSaleWhiteAddress
	0, // 1: member.MemberPreSaleWhiteAddressListResponse.data:type_name -> member.MemberPreSaleWhiteAddress
	6, // 2: member.MemberPreSaleWhiteAddressPaginationResponse.data:type_name -> member.MemberPreSaleWhiteAddressPaginationResponse.Addresses
	0, // 3: member.MemberPreSaleWhiteAddressPaginationResponse.Addresses.list:type_name -> member.MemberPreSaleWhiteAddress
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_member_member_pre_sale_white_address_proto_init() }
func file_member_member_pre_sale_white_address_proto_init() {
	if File_member_member_pre_sale_white_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_member_member_pre_sale_white_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPreSaleWhiteAddress); i {
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
		file_member_member_pre_sale_white_address_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMemberPreSaleWhiteAddressPost); i {
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
		file_member_member_pre_sale_white_address_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMemberPreSaleWhiteAddressPost); i {
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
		file_member_member_pre_sale_white_address_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPreSaleWhiteAddressResponse); i {
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
		file_member_member_pre_sale_white_address_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPreSaleWhiteAddressListResponse); i {
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
		file_member_member_pre_sale_white_address_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPreSaleWhiteAddressPaginationResponse); i {
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
		file_member_member_pre_sale_white_address_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPreSaleWhiteAddressPaginationResponse_Addresses); i {
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
			RawDescriptor: file_member_member_pre_sale_white_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_member_member_pre_sale_white_address_proto_goTypes,
		DependencyIndexes: file_member_member_pre_sale_white_address_proto_depIdxs,
		MessageInfos:      file_member_member_pre_sale_white_address_proto_msgTypes,
	}.Build()
	File_member_member_pre_sale_white_address_proto = out.File
	file_member_member_pre_sale_white_address_proto_rawDesc = nil
	file_member_member_pre_sale_white_address_proto_goTypes = nil
	file_member_member_pre_sale_white_address_proto_depIdxs = nil
}
