// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: nft/blind_box_content.proto

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

// BlindBoxContent 盲盒内容
type BlindBoxContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlindBoxId  uint64  `protobuf:"varint,1,opt,name=blind_box_id,json=blindBoxId,proto3" json:"blind_box_id"`
	WorkId      uint64  `protobuf:"varint,2,opt,name=work_id,json=workId,proto3" json:"work_id"`
	Rarity      uint64  `protobuf:"varint,3,opt,name=rarity,proto3" json:"rarity"`
	Probability float32 `protobuf:"fixed32,4,opt,name=probability,proto3" json:"probability"`
	Total       uint64  `protobuf:"varint,5,opt,name=total,proto3" json:"total"`
	Remain      uint64  `protobuf:"varint,6,opt,name=remain,proto3" json:"remain"`
	Status      uint64  `protobuf:"varint,7,opt,name=status,proto3" json:"status"`
	CreatedAt   string  `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt   string  `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *BlindBoxContent) Reset() {
	*x = BlindBoxContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_blind_box_content_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlindBoxContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlindBoxContent) ProtoMessage() {}

func (x *BlindBoxContent) ProtoReflect() protoreflect.Message {
	mi := &file_nft_blind_box_content_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlindBoxContent.ProtoReflect.Descriptor instead.
func (*BlindBoxContent) Descriptor() ([]byte, []int) {
	return file_nft_blind_box_content_proto_rawDescGZIP(), []int{0}
}

func (x *BlindBoxContent) GetBlindBoxId() uint64 {
	if x != nil {
		return x.BlindBoxId
	}
	return 0
}

func (x *BlindBoxContent) GetWorkId() uint64 {
	if x != nil {
		return x.WorkId
	}
	return 0
}

func (x *BlindBoxContent) GetRarity() uint64 {
	if x != nil {
		return x.Rarity
	}
	return 0
}

func (x *BlindBoxContent) GetProbability() float32 {
	if x != nil {
		return x.Probability
	}
	return 0
}

func (x *BlindBoxContent) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *BlindBoxContent) GetRemain() uint64 {
	if x != nil {
		return x.Remain
	}
	return 0
}

func (x *BlindBoxContent) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *BlindBoxContent) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *BlindBoxContent) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// BlindBoxContentItem 盲盒内容分页 展示的信息
type BlindBoxContentItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlindBoxId  uint64               `protobuf:"varint,1,opt,name=blind_box_id,json=blindBoxId,proto3" json:"blind_box_id"`
	WorkId      uint64               `protobuf:"varint,2,opt,name=work_id,json=workId,proto3" json:"work_id"`
	Rarity      uint64               `protobuf:"varint,3,opt,name=rarity,proto3" json:"rarity"`
	Probability float32              `protobuf:"fixed32,4,opt,name=probability,proto3" json:"probability"`
	Total       uint64               `protobuf:"varint,5,opt,name=total,proto3" json:"total"`
	Remain      uint64               `protobuf:"varint,6,opt,name=remain,proto3" json:"remain"`
	Status      uint64               `protobuf:"varint,7,opt,name=status,proto3" json:"status"`
	CreatedAt   string               `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt   string               `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Work        *BlindBoxContentWork `protobuf:"bytes,10,opt,name=work,proto3" json:"work"`
}

func (x *BlindBoxContentItem) Reset() {
	*x = BlindBoxContentItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_blind_box_content_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlindBoxContentItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlindBoxContentItem) ProtoMessage() {}

func (x *BlindBoxContentItem) ProtoReflect() protoreflect.Message {
	mi := &file_nft_blind_box_content_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlindBoxContentItem.ProtoReflect.Descriptor instead.
func (*BlindBoxContentItem) Descriptor() ([]byte, []int) {
	return file_nft_blind_box_content_proto_rawDescGZIP(), []int{1}
}

func (x *BlindBoxContentItem) GetBlindBoxId() uint64 {
	if x != nil {
		return x.BlindBoxId
	}
	return 0
}

func (x *BlindBoxContentItem) GetWorkId() uint64 {
	if x != nil {
		return x.WorkId
	}
	return 0
}

func (x *BlindBoxContentItem) GetRarity() uint64 {
	if x != nil {
		return x.Rarity
	}
	return 0
}

func (x *BlindBoxContentItem) GetProbability() float32 {
	if x != nil {
		return x.Probability
	}
	return 0
}

func (x *BlindBoxContentItem) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *BlindBoxContentItem) GetRemain() uint64 {
	if x != nil {
		return x.Remain
	}
	return 0
}

func (x *BlindBoxContentItem) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *BlindBoxContentItem) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *BlindBoxContentItem) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *BlindBoxContentItem) GetWork() *BlindBoxContentWork {
	if x != nil {
		return x.Work
	}
	return nil
}

// BlindBoxContentWork 盲盒内容展示的 作品信息
type BlindBoxContentWork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Type     uint64 `protobuf:"varint,2,opt,name=type,proto3" json:"type"`
	Title    string `protobuf:"bytes,3,opt,name=title,proto3" json:"title"`
	SubTitle string `protobuf:"bytes,4,opt,name=sub_title,json=subTitle,proto3" json:"sub_title"`
}

func (x *BlindBoxContentWork) Reset() {
	*x = BlindBoxContentWork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_blind_box_content_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlindBoxContentWork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlindBoxContentWork) ProtoMessage() {}

func (x *BlindBoxContentWork) ProtoReflect() protoreflect.Message {
	mi := &file_nft_blind_box_content_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlindBoxContentWork.ProtoReflect.Descriptor instead.
func (*BlindBoxContentWork) Descriptor() ([]byte, []int) {
	return file_nft_blind_box_content_proto_rawDescGZIP(), []int{2}
}

func (x *BlindBoxContentWork) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BlindBoxContentWork) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *BlindBoxContentWork) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BlindBoxContentWork) GetSubTitle() string {
	if x != nil {
		return x.SubTitle
	}
	return ""
}

// BlindBoxContentInfoPost 获取盲盒内容
type BlindBoxContentInfoPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlindBoxId uint64 `protobuf:"varint,1,opt,name=blind_box_id,json=blindBoxId,proto3" json:"blind_box_id"`
	WorkId     uint64 `protobuf:"varint,2,opt,name=work_id,json=workId,proto3" json:"work_id"`
}

func (x *BlindBoxContentInfoPost) Reset() {
	*x = BlindBoxContentInfoPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_blind_box_content_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlindBoxContentInfoPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlindBoxContentInfoPost) ProtoMessage() {}

func (x *BlindBoxContentInfoPost) ProtoReflect() protoreflect.Message {
	mi := &file_nft_blind_box_content_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlindBoxContentInfoPost.ProtoReflect.Descriptor instead.
func (*BlindBoxContentInfoPost) Descriptor() ([]byte, []int) {
	return file_nft_blind_box_content_proto_rawDescGZIP(), []int{3}
}

func (x *BlindBoxContentInfoPost) GetBlindBoxId() uint64 {
	if x != nil {
		return x.BlindBoxId
	}
	return 0
}

func (x *BlindBoxContentInfoPost) GetWorkId() uint64 {
	if x != nil {
		return x.WorkId
	}
	return 0
}

// BlindBoxContentPaginationPost 盲盒内容分页
type BlindBoxContentPaginationPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlindBoxId uint64 `protobuf:"varint,1,opt,name=blind_box_id,json=blindBoxId,proto3" json:"blind_box_id"`
	Size       uint64 `protobuf:"varint,2,opt,name=Size,proto3" json:"Size"`
	Page       uint64 `protobuf:"varint,3,opt,name=Page,proto3" json:"Page"`
}

func (x *BlindBoxContentPaginationPost) Reset() {
	*x = BlindBoxContentPaginationPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_blind_box_content_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlindBoxContentPaginationPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlindBoxContentPaginationPost) ProtoMessage() {}

func (x *BlindBoxContentPaginationPost) ProtoReflect() protoreflect.Message {
	mi := &file_nft_blind_box_content_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlindBoxContentPaginationPost.ProtoReflect.Descriptor instead.
func (*BlindBoxContentPaginationPost) Descriptor() ([]byte, []int) {
	return file_nft_blind_box_content_proto_rawDescGZIP(), []int{4}
}

func (x *BlindBoxContentPaginationPost) GetBlindBoxId() uint64 {
	if x != nil {
		return x.BlindBoxId
	}
	return 0
}

func (x *BlindBoxContentPaginationPost) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *BlindBoxContentPaginationPost) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

// CreateBlindBoxContentPost 创建盲盒内容
type CreateBlindBoxContentPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlindBoxId  uint64  `protobuf:"varint,1,opt,name=blind_box_id,json=blindBoxId,proto3" json:"blind_box_id"`
	WorkId      uint64  `protobuf:"varint,2,opt,name=work_id,json=workId,proto3" json:"work_id"`
	Probability float32 `protobuf:"fixed32,3,opt,name=probability,proto3" json:"probability"`
	Total       uint64  `protobuf:"varint,4,opt,name=total,proto3" json:"total"`
	Remain      uint64  `protobuf:"varint,5,opt,name=remain,proto3" json:"remain"`
}

func (x *CreateBlindBoxContentPost) Reset() {
	*x = CreateBlindBoxContentPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_blind_box_content_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlindBoxContentPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlindBoxContentPost) ProtoMessage() {}

func (x *CreateBlindBoxContentPost) ProtoReflect() protoreflect.Message {
	mi := &file_nft_blind_box_content_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlindBoxContentPost.ProtoReflect.Descriptor instead.
func (*CreateBlindBoxContentPost) Descriptor() ([]byte, []int) {
	return file_nft_blind_box_content_proto_rawDescGZIP(), []int{5}
}

func (x *CreateBlindBoxContentPost) GetBlindBoxId() uint64 {
	if x != nil {
		return x.BlindBoxId
	}
	return 0
}

func (x *CreateBlindBoxContentPost) GetWorkId() uint64 {
	if x != nil {
		return x.WorkId
	}
	return 0
}

func (x *CreateBlindBoxContentPost) GetProbability() float32 {
	if x != nil {
		return x.Probability
	}
	return 0
}

func (x *CreateBlindBoxContentPost) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CreateBlindBoxContentPost) GetRemain() uint64 {
	if x != nil {
		return x.Remain
	}
	return 0
}

// UpdateBlindBoxContentPost 更新盲盒内容
type UpdateBlindBoxContentPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlindBoxId  uint64  `protobuf:"varint,1,opt,name=blind_box_id,json=blindBoxId,proto3" json:"blind_box_id"`
	WorkId      uint64  `protobuf:"varint,2,opt,name=work_id,json=workId,proto3" json:"work_id"`
	Probability float32 `protobuf:"fixed32,3,opt,name=probability,proto3" json:"probability"`
	Total       uint64  `protobuf:"varint,4,opt,name=total,proto3" json:"total"`
	Remain      uint64  `protobuf:"varint,5,opt,name=remain,proto3" json:"remain"`
}

func (x *UpdateBlindBoxContentPost) Reset() {
	*x = UpdateBlindBoxContentPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_blind_box_content_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBlindBoxContentPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlindBoxContentPost) ProtoMessage() {}

func (x *UpdateBlindBoxContentPost) ProtoReflect() protoreflect.Message {
	mi := &file_nft_blind_box_content_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlindBoxContentPost.ProtoReflect.Descriptor instead.
func (*UpdateBlindBoxContentPost) Descriptor() ([]byte, []int) {
	return file_nft_blind_box_content_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBlindBoxContentPost) GetBlindBoxId() uint64 {
	if x != nil {
		return x.BlindBoxId
	}
	return 0
}

func (x *UpdateBlindBoxContentPost) GetWorkId() uint64 {
	if x != nil {
		return x.WorkId
	}
	return 0
}

func (x *UpdateBlindBoxContentPost) GetProbability() float32 {
	if x != nil {
		return x.Probability
	}
	return 0
}

func (x *UpdateBlindBoxContentPost) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *UpdateBlindBoxContentPost) GetRemain() uint64 {
	if x != nil {
		return x.Remain
	}
	return 0
}

// BlindBoxContentResponse 返回单个
type BlindBoxContentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string           `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    *BlindBoxContent `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *BlindBoxContentResponse) Reset() {
	*x = BlindBoxContentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_blind_box_content_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlindBoxContentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlindBoxContentResponse) ProtoMessage() {}

func (x *BlindBoxContentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nft_blind_box_content_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlindBoxContentResponse.ProtoReflect.Descriptor instead.
func (*BlindBoxContentResponse) Descriptor() ([]byte, []int) {
	return file_nft_blind_box_content_proto_rawDescGZIP(), []int{7}
}

func (x *BlindBoxContentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BlindBoxContentResponse) GetData() *BlindBoxContent {
	if x != nil {
		return x.Data
	}
	return nil
}

// BlindBoxContentPaginationResponse 返回列表
type BlindBoxContentPaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                                              `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    *BlindBoxContentPaginationResponse_BlindBoxContents `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *BlindBoxContentPaginationResponse) Reset() {
	*x = BlindBoxContentPaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_blind_box_content_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlindBoxContentPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlindBoxContentPaginationResponse) ProtoMessage() {}

func (x *BlindBoxContentPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nft_blind_box_content_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlindBoxContentPaginationResponse.ProtoReflect.Descriptor instead.
func (*BlindBoxContentPaginationResponse) Descriptor() ([]byte, []int) {
	return file_nft_blind_box_content_proto_rawDescGZIP(), []int{8}
}

func (x *BlindBoxContentPaginationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BlindBoxContentPaginationResponse) GetData() *BlindBoxContentPaginationResponse_BlindBoxContents {
	if x != nil {
		return x.Data
	}
	return nil
}

type BlindBoxContentPaginationResponse_BlindBoxContents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*BlindBoxContentItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
	Total int64                  `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
}

func (x *BlindBoxContentPaginationResponse_BlindBoxContents) Reset() {
	*x = BlindBoxContentPaginationResponse_BlindBoxContents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_blind_box_content_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlindBoxContentPaginationResponse_BlindBoxContents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlindBoxContentPaginationResponse_BlindBoxContents) ProtoMessage() {}

func (x *BlindBoxContentPaginationResponse_BlindBoxContents) ProtoReflect() protoreflect.Message {
	mi := &file_nft_blind_box_content_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlindBoxContentPaginationResponse_BlindBoxContents.ProtoReflect.Descriptor instead.
func (*BlindBoxContentPaginationResponse_BlindBoxContents) Descriptor() ([]byte, []int) {
	return file_nft_blind_box_content_proto_rawDescGZIP(), []int{8, 0}
}

func (x *BlindBoxContentPaginationResponse_BlindBoxContents) GetList() []*BlindBoxContentItem {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *BlindBoxContentPaginationResponse_BlindBoxContents) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_nft_blind_box_content_proto protoreflect.FileDescriptor

var file_nft_blind_box_content_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6e, 0x66, 0x74, 0x2f, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x5f, 0x62, 0x6f, 0x78, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e,
	0x66, 0x74, 0x22, 0x8a, 0x02, 0x0a, 0x0f, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x5f,
	0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x6c,
	0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0xbc, 0x02, 0x0a, 0x13, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x6c, 0x69, 0x6e, 0x64,
	0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62,
	0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x2c, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x6c,
	0x0a, 0x13, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x54, 0x0a, 0x17,
	0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x6c, 0x69, 0x6e, 0x64,
	0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62,
	0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b,
	0x49, 0x64, 0x22, 0x69, 0x0a, 0x1d, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x5f, 0x62, 0x6f, 0x78,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x6c, 0x69, 0x6e, 0x64,
	0x42, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x22, 0xa6, 0x01,
	0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x62,
	0x6c, 0x69, 0x6e, 0x64, 0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xa6, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x5f, 0x62, 0x6f,
	0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x6c, 0x69, 0x6e,
	0x64, 0x42, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x22,
	0x5d, 0x0a, 0x17, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f,
	0x78, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xe2,
	0x01, 0x0a, 0x21, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4b,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x6e,
	0x66, 0x74, 0x2e, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x56, 0x0a, 0x10, 0x42,
	0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x2c, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x6e, 0x66, 0x74, 0x2e, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nft_blind_box_content_proto_rawDescOnce sync.Once
	file_nft_blind_box_content_proto_rawDescData = file_nft_blind_box_content_proto_rawDesc
)

func file_nft_blind_box_content_proto_rawDescGZIP() []byte {
	file_nft_blind_box_content_proto_rawDescOnce.Do(func() {
		file_nft_blind_box_content_proto_rawDescData = protoimpl.X.CompressGZIP(file_nft_blind_box_content_proto_rawDescData)
	})
	return file_nft_blind_box_content_proto_rawDescData
}

var file_nft_blind_box_content_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_nft_blind_box_content_proto_goTypes = []interface{}{
	(*BlindBoxContent)(nil),                                    // 0: nft.BlindBoxContent
	(*BlindBoxContentItem)(nil),                                // 1: nft.BlindBoxContentItem
	(*BlindBoxContentWork)(nil),                                // 2: nft.BlindBoxContentWork
	(*BlindBoxContentInfoPost)(nil),                            // 3: nft.BlindBoxContentInfoPost
	(*BlindBoxContentPaginationPost)(nil),                      // 4: nft.BlindBoxContentPaginationPost
	(*CreateBlindBoxContentPost)(nil),                          // 5: nft.CreateBlindBoxContentPost
	(*UpdateBlindBoxContentPost)(nil),                          // 6: nft.UpdateBlindBoxContentPost
	(*BlindBoxContentResponse)(nil),                            // 7: nft.BlindBoxContentResponse
	(*BlindBoxContentPaginationResponse)(nil),                  // 8: nft.BlindBoxContentPaginationResponse
	(*BlindBoxContentPaginationResponse_BlindBoxContents)(nil), // 9: nft.BlindBoxContentPaginationResponse.BlindBoxContents
}
var file_nft_blind_box_content_proto_depIdxs = []int32{
	2, // 0: nft.BlindBoxContentItem.work:type_name -> nft.BlindBoxContentWork
	0, // 1: nft.BlindBoxContentResponse.data:type_name -> nft.BlindBoxContent
	9, // 2: nft.BlindBoxContentPaginationResponse.data:type_name -> nft.BlindBoxContentPaginationResponse.BlindBoxContents
	1, // 3: nft.BlindBoxContentPaginationResponse.BlindBoxContents.list:type_name -> nft.BlindBoxContentItem
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_nft_blind_box_content_proto_init() }
func file_nft_blind_box_content_proto_init() {
	if File_nft_blind_box_content_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nft_blind_box_content_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlindBoxContent); i {
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
		file_nft_blind_box_content_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlindBoxContentItem); i {
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
		file_nft_blind_box_content_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlindBoxContentWork); i {
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
		file_nft_blind_box_content_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlindBoxContentInfoPost); i {
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
		file_nft_blind_box_content_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlindBoxContentPaginationPost); i {
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
		file_nft_blind_box_content_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlindBoxContentPost); i {
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
		file_nft_blind_box_content_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBlindBoxContentPost); i {
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
		file_nft_blind_box_content_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlindBoxContentResponse); i {
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
		file_nft_blind_box_content_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlindBoxContentPaginationResponse); i {
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
		file_nft_blind_box_content_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlindBoxContentPaginationResponse_BlindBoxContents); i {
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
			RawDescriptor: file_nft_blind_box_content_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nft_blind_box_content_proto_goTypes,
		DependencyIndexes: file_nft_blind_box_content_proto_depIdxs,
		MessageInfos:      file_nft_blind_box_content_proto_msgTypes,
	}.Build()
	File_nft_blind_box_content_proto = out.File
	file_nft_blind_box_content_proto_rawDesc = nil
	file_nft_blind_box_content_proto_goTypes = nil
	file_nft_blind_box_content_proto_depIdxs = nil
}
