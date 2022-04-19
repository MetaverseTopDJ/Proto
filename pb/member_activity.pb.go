// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: member/member_activity.proto

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

// MemberActivity 会员活动
type MemberActivity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
	VipStart    string `protobuf:"bytes,4,opt,name=vip_start,json=vipStart,proto3" json:"vip_start"`
	VipEnd      string `protobuf:"bytes,5,opt,name=vip_end,json=vipEnd,proto3" json:"vip_end"`
	VipLimit    uint64 `protobuf:"varint,6,opt,name=vip_limit,json=vipLimit,proto3" json:"vip_limit"`
	PreStart    string `protobuf:"bytes,7,opt,name=pre_start,json=preStart,proto3" json:"pre_start"`
	PreEnd      string `protobuf:"bytes,8,opt,name=pre_end,json=preEnd,proto3" json:"pre_end"`
	PreLimit    uint64 `protobuf:"varint,9,opt,name=pre_limit,json=preLimit,proto3" json:"pre_limit"`
	SaleStart   string `protobuf:"bytes,10,opt,name=sale_start,json=saleStart,proto3" json:"sale_start"`
	SaleEnd     string `protobuf:"bytes,11,opt,name=sale_end,json=saleEnd,proto3" json:"sale_end"`
	SaleLimit   uint64 `protobuf:"varint,12,opt,name=sale_limit,json=saleLimit,proto3" json:"sale_limit"`
	Status      uint64 `protobuf:"varint,13,opt,name=status,proto3" json:"status"`
	Checked     bool   `protobuf:"varint,14,opt,name=checked,proto3" json:"checked"`
	CreatedAt   string `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt   string `protobuf:"bytes,16,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *MemberActivity) Reset() {
	*x = MemberActivity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_activity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberActivity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberActivity) ProtoMessage() {}

func (x *MemberActivity) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_activity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberActivity.ProtoReflect.Descriptor instead.
func (*MemberActivity) Descriptor() ([]byte, []int) {
	return file_member_member_activity_proto_rawDescGZIP(), []int{0}
}

func (x *MemberActivity) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberActivity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MemberActivity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MemberActivity) GetVipStart() string {
	if x != nil {
		return x.VipStart
	}
	return ""
}

func (x *MemberActivity) GetVipEnd() string {
	if x != nil {
		return x.VipEnd
	}
	return ""
}

func (x *MemberActivity) GetVipLimit() uint64 {
	if x != nil {
		return x.VipLimit
	}
	return 0
}

func (x *MemberActivity) GetPreStart() string {
	if x != nil {
		return x.PreStart
	}
	return ""
}

func (x *MemberActivity) GetPreEnd() string {
	if x != nil {
		return x.PreEnd
	}
	return ""
}

func (x *MemberActivity) GetPreLimit() uint64 {
	if x != nil {
		return x.PreLimit
	}
	return 0
}

func (x *MemberActivity) GetSaleStart() string {
	if x != nil {
		return x.SaleStart
	}
	return ""
}

func (x *MemberActivity) GetSaleEnd() string {
	if x != nil {
		return x.SaleEnd
	}
	return ""
}

func (x *MemberActivity) GetSaleLimit() uint64 {
	if x != nil {
		return x.SaleLimit
	}
	return 0
}

func (x *MemberActivity) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MemberActivity) GetChecked() bool {
	if x != nil {
		return x.Checked
	}
	return false
}

func (x *MemberActivity) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *MemberActivity) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// CreateMemberActivityPost 创建会员活动参数
type CreateMemberActivityPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
	VipStart    string `protobuf:"bytes,4,opt,name=vip_start,json=vipStart,proto3" json:"vip_start"`
	VipEnd      string `protobuf:"bytes,5,opt,name=vip_end,json=vipEnd,proto3" json:"vip_end"`
	VipLimit    uint64 `protobuf:"varint,6,opt,name=vip_limit,json=vipLimit,proto3" json:"vip_limit"`
	PreStart    string `protobuf:"bytes,7,opt,name=pre_start,json=preStart,proto3" json:"pre_start"`
	PreEnd      string `protobuf:"bytes,8,opt,name=pre_end,json=preEnd,proto3" json:"pre_end"`
	PreLimit    uint64 `protobuf:"varint,9,opt,name=pre_limit,json=preLimit,proto3" json:"pre_limit"`
	SaleStart   string `protobuf:"bytes,10,opt,name=sale_start,json=saleStart,proto3" json:"sale_start"`
	SaleEnd     string `protobuf:"bytes,11,opt,name=sale_end,json=saleEnd,proto3" json:"sale_end"`
	SaleLimit   uint64 `protobuf:"varint,12,opt,name=sale_limit,json=saleLimit,proto3" json:"sale_limit"`
}

func (x *CreateMemberActivityPost) Reset() {
	*x = CreateMemberActivityPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_activity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMemberActivityPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMemberActivityPost) ProtoMessage() {}

func (x *CreateMemberActivityPost) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_activity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMemberActivityPost.ProtoReflect.Descriptor instead.
func (*CreateMemberActivityPost) Descriptor() ([]byte, []int) {
	return file_member_member_activity_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMemberActivityPost) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateMemberActivityPost) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateMemberActivityPost) GetVipStart() string {
	if x != nil {
		return x.VipStart
	}
	return ""
}

func (x *CreateMemberActivityPost) GetVipEnd() string {
	if x != nil {
		return x.VipEnd
	}
	return ""
}

func (x *CreateMemberActivityPost) GetVipLimit() uint64 {
	if x != nil {
		return x.VipLimit
	}
	return 0
}

func (x *CreateMemberActivityPost) GetPreStart() string {
	if x != nil {
		return x.PreStart
	}
	return ""
}

func (x *CreateMemberActivityPost) GetPreEnd() string {
	if x != nil {
		return x.PreEnd
	}
	return ""
}

func (x *CreateMemberActivityPost) GetPreLimit() uint64 {
	if x != nil {
		return x.PreLimit
	}
	return 0
}

func (x *CreateMemberActivityPost) GetSaleStart() string {
	if x != nil {
		return x.SaleStart
	}
	return ""
}

func (x *CreateMemberActivityPost) GetSaleEnd() string {
	if x != nil {
		return x.SaleEnd
	}
	return ""
}

func (x *CreateMemberActivityPost) GetSaleLimit() uint64 {
	if x != nil {
		return x.SaleLimit
	}
	return 0
}

// UpdateMemberActivityPost 更逊会员活动参数
type UpdateMemberActivityPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
	VipStart    string `protobuf:"bytes,4,opt,name=vip_start,json=vipStart,proto3" json:"vip_start"`
	VipEnd      string `protobuf:"bytes,5,opt,name=vip_end,json=vipEnd,proto3" json:"vip_end"`
	VipLimit    uint64 `protobuf:"varint,6,opt,name=vip_limit,json=vipLimit,proto3" json:"vip_limit"`
	PreStart    string `protobuf:"bytes,7,opt,name=pre_start,json=preStart,proto3" json:"pre_start"`
	PreEnd      string `protobuf:"bytes,8,opt,name=pre_end,json=preEnd,proto3" json:"pre_end"`
	PreLimit    uint64 `protobuf:"varint,9,opt,name=pre_limit,json=preLimit,proto3" json:"pre_limit"`
	SaleStart   string `protobuf:"bytes,10,opt,name=sale_start,json=saleStart,proto3" json:"sale_start"`
	SaleEnd     string `protobuf:"bytes,11,opt,name=sale_end,json=saleEnd,proto3" json:"sale_end"`
	SaleLimit   uint64 `protobuf:"varint,12,opt,name=sale_limit,json=saleLimit,proto3" json:"sale_limit"`
}

func (x *UpdateMemberActivityPost) Reset() {
	*x = UpdateMemberActivityPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_activity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMemberActivityPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMemberActivityPost) ProtoMessage() {}

func (x *UpdateMemberActivityPost) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_activity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMemberActivityPost.ProtoReflect.Descriptor instead.
func (*UpdateMemberActivityPost) Descriptor() ([]byte, []int) {
	return file_member_member_activity_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateMemberActivityPost) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateMemberActivityPost) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateMemberActivityPost) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateMemberActivityPost) GetVipStart() string {
	if x != nil {
		return x.VipStart
	}
	return ""
}

func (x *UpdateMemberActivityPost) GetVipEnd() string {
	if x != nil {
		return x.VipEnd
	}
	return ""
}

func (x *UpdateMemberActivityPost) GetVipLimit() uint64 {
	if x != nil {
		return x.VipLimit
	}
	return 0
}

func (x *UpdateMemberActivityPost) GetPreStart() string {
	if x != nil {
		return x.PreStart
	}
	return ""
}

func (x *UpdateMemberActivityPost) GetPreEnd() string {
	if x != nil {
		return x.PreEnd
	}
	return ""
}

func (x *UpdateMemberActivityPost) GetPreLimit() uint64 {
	if x != nil {
		return x.PreLimit
	}
	return 0
}

func (x *UpdateMemberActivityPost) GetSaleStart() string {
	if x != nil {
		return x.SaleStart
	}
	return ""
}

func (x *UpdateMemberActivityPost) GetSaleEnd() string {
	if x != nil {
		return x.SaleEnd
	}
	return ""
}

func (x *UpdateMemberActivityPost) GetSaleLimit() uint64 {
	if x != nil {
		return x.SaleLimit
	}
	return 0
}

// MemberActivityResponse 返回 会员活动
type MemberActivityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string          `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    *MemberActivity `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *MemberActivityResponse) Reset() {
	*x = MemberActivityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_activity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberActivityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberActivityResponse) ProtoMessage() {}

func (x *MemberActivityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_activity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberActivityResponse.ProtoReflect.Descriptor instead.
func (*MemberActivityResponse) Descriptor() ([]byte, []int) {
	return file_member_member_activity_proto_rawDescGZIP(), []int{3}
}

func (x *MemberActivityResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MemberActivityResponse) GetData() *MemberActivity {
	if x != nil {
		return x.Data
	}
	return nil
}

// MemberActivitiesResponse 返回 会员活动 列表
type MemberActivitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string            `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    []*MemberActivity `protobuf:"bytes,2,rep,name=data,proto3" json:"data"`
}

func (x *MemberActivitiesResponse) Reset() {
	*x = MemberActivitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_activity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberActivitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberActivitiesResponse) ProtoMessage() {}

func (x *MemberActivitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_activity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberActivitiesResponse.ProtoReflect.Descriptor instead.
func (*MemberActivitiesResponse) Descriptor() ([]byte, []int) {
	return file_member_member_activity_proto_rawDescGZIP(), []int{4}
}

func (x *MemberActivitiesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MemberActivitiesResponse) GetData() []*MemberActivity {
	if x != nil {
		return x.Data
	}
	return nil
}

// MemberActivityPaginationResponse 返回 会员活动分页
type MemberActivityPaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                                             `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    *MemberActivityPaginationResponse_MemberActivities `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *MemberActivityPaginationResponse) Reset() {
	*x = MemberActivityPaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_activity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberActivityPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberActivityPaginationResponse) ProtoMessage() {}

func (x *MemberActivityPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_activity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberActivityPaginationResponse.ProtoReflect.Descriptor instead.
func (*MemberActivityPaginationResponse) Descriptor() ([]byte, []int) {
	return file_member_member_activity_proto_rawDescGZIP(), []int{5}
}

func (x *MemberActivityPaginationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MemberActivityPaginationResponse) GetData() *MemberActivityPaginationResponse_MemberActivities {
	if x != nil {
		return x.Data
	}
	return nil
}

type MemberActivityPaginationResponse_MemberActivities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*MemberActivity `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
	Total int64             `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
}

func (x *MemberActivityPaginationResponse_MemberActivities) Reset() {
	*x = MemberActivityPaginationResponse_MemberActivities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_activity_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberActivityPaginationResponse_MemberActivities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberActivityPaginationResponse_MemberActivities) ProtoMessage() {}

func (x *MemberActivityPaginationResponse_MemberActivities) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_activity_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberActivityPaginationResponse_MemberActivities.ProtoReflect.Descriptor instead.
func (*MemberActivityPaginationResponse_MemberActivities) Descriptor() ([]byte, []int) {
	return file_member_member_activity_proto_rawDescGZIP(), []int{5, 0}
}

func (x *MemberActivityPaginationResponse_MemberActivities) GetList() []*MemberActivity {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *MemberActivityPaginationResponse_MemberActivities) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_member_member_activity_proto protoreflect.FileDescriptor

var file_member_member_activity_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xc5, 0x03, 0x0a, 0x0e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x76, 0x69, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x70, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x76, 0x69, 0x70, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76,
	0x69, 0x70, 0x45, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x70, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x76, 0x69, 0x70, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x72, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x72, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x61, 0x6c, 0x65, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x61, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x61, 0x6c, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xcf,
	0x02, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x70, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x76, 0x69, 0x70, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x76, 0x69, 0x70, 0x45, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x70, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x76, 0x69, 0x70, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x65, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72,
	0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70,
	0x72, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x61, 0x6c,
	0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x65,
	0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x61, 0x6c, 0x65, 0x45, 0x6e,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x61, 0x6c, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0xdf, 0x02, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x70, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x76, 0x69, 0x70, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x76, 0x69, 0x70, 0x45, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x70,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x76, 0x69,
	0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x65, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x72, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x70, 0x72, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6c,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x61, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x61, 0x6c, 0x65,
	0x5f, 0x65, 0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x61, 0x6c, 0x65,
	0x45, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x61, 0x6c, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x5e, 0x0a, 0x16, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x60, 0x0a, 0x18, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xe1, 0x01, 0x0a, 0x20, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x4d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x39, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x54, 0x0a, 0x10, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_member_member_activity_proto_rawDescOnce sync.Once
	file_member_member_activity_proto_rawDescData = file_member_member_activity_proto_rawDesc
)

func file_member_member_activity_proto_rawDescGZIP() []byte {
	file_member_member_activity_proto_rawDescOnce.Do(func() {
		file_member_member_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_member_member_activity_proto_rawDescData)
	})
	return file_member_member_activity_proto_rawDescData
}

var file_member_member_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_member_member_activity_proto_goTypes = []interface{}{
	(*MemberActivity)(nil),                                    // 0: member.MemberActivity
	(*CreateMemberActivityPost)(nil),                          // 1: member.CreateMemberActivityPost
	(*UpdateMemberActivityPost)(nil),                          // 2: member.UpdateMemberActivityPost
	(*MemberActivityResponse)(nil),                            // 3: member.MemberActivityResponse
	(*MemberActivitiesResponse)(nil),                          // 4: member.MemberActivitiesResponse
	(*MemberActivityPaginationResponse)(nil),                  // 5: member.MemberActivityPaginationResponse
	(*MemberActivityPaginationResponse_MemberActivities)(nil), // 6: member.MemberActivityPaginationResponse.MemberActivities
}
var file_member_member_activity_proto_depIdxs = []int32{
	0, // 0: member.MemberActivityResponse.data:type_name -> member.MemberActivity
	0, // 1: member.MemberActivitiesResponse.data:type_name -> member.MemberActivity
	6, // 2: member.MemberActivityPaginationResponse.data:type_name -> member.MemberActivityPaginationResponse.MemberActivities
	0, // 3: member.MemberActivityPaginationResponse.MemberActivities.list:type_name -> member.MemberActivity
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_member_member_activity_proto_init() }
func file_member_member_activity_proto_init() {
	if File_member_member_activity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_member_member_activity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberActivity); i {
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
		file_member_member_activity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMemberActivityPost); i {
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
		file_member_member_activity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMemberActivityPost); i {
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
		file_member_member_activity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberActivityResponse); i {
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
		file_member_member_activity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberActivitiesResponse); i {
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
		file_member_member_activity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberActivityPaginationResponse); i {
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
		file_member_member_activity_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberActivityPaginationResponse_MemberActivities); i {
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
			RawDescriptor: file_member_member_activity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_member_member_activity_proto_goTypes,
		DependencyIndexes: file_member_member_activity_proto_depIdxs,
		MessageInfos:      file_member_member_activity_proto_msgTypes,
	}.Build()
	File_member_member_activity_proto = out.File
	file_member_member_activity_proto_rawDesc = nil
	file_member_member_activity_proto_goTypes = nil
	file_member_member_activity_proto_depIdxs = nil
}
