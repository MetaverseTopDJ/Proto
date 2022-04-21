// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: member/member_card.proto

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

// MemberCard 会员拥有的卡片
type MemberCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	ActivityId uint64 `protobuf:"varint,2,opt,name=activity_id,json=activityId,proto3" json:"activity_id"`
	CreatorId  uint64 `protobuf:"varint,3,opt,name=creator_id,json=creatorId,proto3" json:"creator_id"`
	OwnerId    uint64 `protobuf:"varint,4,opt,name=owner_id,json=ownerId,proto3" json:"owner_id"`
	CardId     uint64 `protobuf:"varint,5,opt,name=card_id,json=cardId,proto3" json:"card_id"`
	NftId      uint64 `protobuf:"varint,6,opt,name=nft_id,json=nftId,proto3" json:"nft_id"`
	Version    uint64 `protobuf:"varint,7,opt,name=version,proto3" json:"version"`
	Status     uint64 `protobuf:"varint,8,opt,name=status,proto3" json:"status"`
	CreatedAt  string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt  string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *MemberCard) Reset() {
	*x = MemberCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberCard) ProtoMessage() {}

func (x *MemberCard) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberCard.ProtoReflect.Descriptor instead.
func (*MemberCard) Descriptor() ([]byte, []int) {
	return file_member_member_card_proto_rawDescGZIP(), []int{0}
}

func (x *MemberCard) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberCard) GetActivityId() uint64 {
	if x != nil {
		return x.ActivityId
	}
	return 0
}

func (x *MemberCard) GetCreatorId() uint64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *MemberCard) GetOwnerId() uint64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *MemberCard) GetCardId() uint64 {
	if x != nil {
		return x.CardId
	}
	return 0
}

func (x *MemberCard) GetNftId() uint64 {
	if x != nil {
		return x.NftId
	}
	return 0
}

func (x *MemberCard) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *MemberCard) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MemberCard) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *MemberCard) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// MemberCardCreatePost 会员卡 铸造参数
type MemberCardCreatePost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractId  uint64 `protobuf:"varint,1,opt,name=contract_id,json=contractId,proto3" json:"contract_id"`
	TokenId     uint64 `protobuf:"varint,2,opt,name=token_id,json=tokenId,proto3" json:"token_id"`
	BlockNumber uint64 `protobuf:"varint,3,opt,name=block_number,json=blockNumber,proto3" json:"block_number"`
	Hash        string `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash"`
	Address     string `protobuf:"bytes,5,opt,name=address,proto3" json:"address"`
	Timestamp   uint64 `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp"`
}

func (x *MemberCardCreatePost) Reset() {
	*x = MemberCardCreatePost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_card_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberCardCreatePost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberCardCreatePost) ProtoMessage() {}

func (x *MemberCardCreatePost) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_card_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberCardCreatePost.ProtoReflect.Descriptor instead.
func (*MemberCardCreatePost) Descriptor() ([]byte, []int) {
	return file_member_member_card_proto_rawDescGZIP(), []int{1}
}

func (x *MemberCardCreatePost) GetContractId() uint64 {
	if x != nil {
		return x.ContractId
	}
	return 0
}

func (x *MemberCardCreatePost) GetTokenId() uint64 {
	if x != nil {
		return x.TokenId
	}
	return 0
}

func (x *MemberCardCreatePost) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *MemberCardCreatePost) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *MemberCardCreatePost) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MemberCardCreatePost) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// MemberCardTransferPost 会员卡 交易参数
type MemberCardTransferPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractId  uint64 `protobuf:"varint,1,opt,name=contract_id,json=contractId,proto3" json:"contract_id"`
	TokenId     uint64 `protobuf:"varint,2,opt,name=token_id,json=tokenId,proto3" json:"token_id"`
	BlockNumber uint64 `protobuf:"varint,3,opt,name=block_number,json=blockNumber,proto3" json:"block_number"`
	Hash        string `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash"`
	From        string `protobuf:"bytes,5,opt,name=from,proto3" json:"from"`
	To          string `protobuf:"bytes,6,opt,name=to,proto3" json:"to"`
}

func (x *MemberCardTransferPost) Reset() {
	*x = MemberCardTransferPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_card_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberCardTransferPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberCardTransferPost) ProtoMessage() {}

func (x *MemberCardTransferPost) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_card_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberCardTransferPost.ProtoReflect.Descriptor instead.
func (*MemberCardTransferPost) Descriptor() ([]byte, []int) {
	return file_member_member_card_proto_rawDescGZIP(), []int{2}
}

func (x *MemberCardTransferPost) GetContractId() uint64 {
	if x != nil {
		return x.ContractId
	}
	return 0
}

func (x *MemberCardTransferPost) GetTokenId() uint64 {
	if x != nil {
		return x.TokenId
	}
	return 0
}

func (x *MemberCardTransferPost) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *MemberCardTransferPost) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *MemberCardTransferPost) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *MemberCardTransferPost) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

// CreateMemberCardPost 新增会员卡
type CreateMemberCardPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateMemberCardPost) Reset() {
	*x = CreateMemberCardPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_card_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMemberCardPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMemberCardPost) ProtoMessage() {}

func (x *CreateMemberCardPost) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_card_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMemberCardPost.ProtoReflect.Descriptor instead.
func (*CreateMemberCardPost) Descriptor() ([]byte, []int) {
	return file_member_member_card_proto_rawDescGZIP(), []int{3}
}

// UpdateMemberCardPost 更新会员卡
type UpdateMemberCardPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateMemberCardPost) Reset() {
	*x = UpdateMemberCardPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_card_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMemberCardPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMemberCardPost) ProtoMessage() {}

func (x *UpdateMemberCardPost) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_card_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMemberCardPost.ProtoReflect.Descriptor instead.
func (*UpdateMemberCardPost) Descriptor() ([]byte, []int) {
	return file_member_member_card_proto_rawDescGZIP(), []int{4}
}

// MemberCardResponse 会员卡返回
type MemberCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MemberCardResponse) Reset() {
	*x = MemberCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_card_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberCardResponse) ProtoMessage() {}

func (x *MemberCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_card_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberCardResponse.ProtoReflect.Descriptor instead.
func (*MemberCardResponse) Descriptor() ([]byte, []int) {
	return file_member_member_card_proto_rawDescGZIP(), []int{5}
}

var File_member_member_card_proto protoreflect.FileDescriptor

var file_member_member_card_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x97, 0x02, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x61, 0x72,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6e, 0x66, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x66, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc1, 0x01, 0x0a,
	0x14, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0xaf, 0x01, 0x0a, 0x16, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x50, 0x6f,
	0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_member_member_card_proto_rawDescOnce sync.Once
	file_member_member_card_proto_rawDescData = file_member_member_card_proto_rawDesc
)

func file_member_member_card_proto_rawDescGZIP() []byte {
	file_member_member_card_proto_rawDescOnce.Do(func() {
		file_member_member_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_member_member_card_proto_rawDescData)
	})
	return file_member_member_card_proto_rawDescData
}

var file_member_member_card_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_member_member_card_proto_goTypes = []interface{}{
	(*MemberCard)(nil),             // 0: member.MemberCard
	(*MemberCardCreatePost)(nil),   // 1: member.MemberCardCreatePost
	(*MemberCardTransferPost)(nil), // 2: member.MemberCardTransferPost
	(*CreateMemberCardPost)(nil),   // 3: member.CreateMemberCardPost
	(*UpdateMemberCardPost)(nil),   // 4: member.UpdateMemberCardPost
	(*MemberCardResponse)(nil),     // 5: member.MemberCardResponse
}
var file_member_member_card_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_member_member_card_proto_init() }
func file_member_member_card_proto_init() {
	if File_member_member_card_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_member_member_card_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberCard); i {
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
		file_member_member_card_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberCardCreatePost); i {
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
		file_member_member_card_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberCardTransferPost); i {
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
		file_member_member_card_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMemberCardPost); i {
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
		file_member_member_card_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMemberCardPost); i {
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
		file_member_member_card_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberCardResponse); i {
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
			RawDescriptor: file_member_member_card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_member_member_card_proto_goTypes,
		DependencyIndexes: file_member_member_card_proto_depIdxs,
		MessageInfos:      file_member_member_card_proto_msgTypes,
	}.Build()
	File_member_member_card_proto = out.File
	file_member_member_card_proto_rawDesc = nil
	file_member_member_card_proto_goTypes = nil
	file_member_member_card_proto_depIdxs = nil
}
