// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: console/rule.proto

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

type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuId    uint64 `protobuf:"varint,1,opt,name=menu_id,json=menuId,proto3" json:"menu_id,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Code      string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Urls      string `protobuf:"bytes,5,opt,name=urls,proto3" json:"urls,omitempty"`
	Sort      uint64 `protobuf:"varint,6,opt,name=sort,proto3" json:"sort,omitempty"`
	Status    uint64 `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	Remark    string `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	CreatedAt string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_rule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_console_rule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_console_rule_proto_rawDescGZIP(), []int{0}
}

func (x *Rule) GetMenuId() uint64 {
	if x != nil {
		return x.MenuId
	}
	return 0
}

func (x *Rule) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Rule) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Rule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rule) GetUrls() string {
	if x != nil {
		return x.Urls
	}
	return ""
}

func (x *Rule) GetSort() uint64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Rule) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Rule) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *Rule) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Rule) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateRulePost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuId uint64 `protobuf:"varint,2,opt,name=menu_id,json=menuId,proto3" json:"menu_id,omitempty"`
	Code   string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Name   string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Urls   string `protobuf:"bytes,5,opt,name=urls,proto3" json:"urls,omitempty"`
	Sort   uint64 `protobuf:"varint,6,opt,name=sort,proto3" json:"sort,omitempty"`
	Remark string `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *CreateRulePost) Reset() {
	*x = CreateRulePost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_rule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRulePost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRulePost) ProtoMessage() {}

func (x *CreateRulePost) ProtoReflect() protoreflect.Message {
	mi := &file_console_rule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRulePost.ProtoReflect.Descriptor instead.
func (*CreateRulePost) Descriptor() ([]byte, []int) {
	return file_console_rule_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRulePost) GetMenuId() uint64 {
	if x != nil {
		return x.MenuId
	}
	return 0
}

func (x *CreateRulePost) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateRulePost) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRulePost) GetUrls() string {
	if x != nil {
		return x.Urls
	}
	return ""
}

func (x *CreateRulePost) GetSort() uint64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *CreateRulePost) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type UpdateRulePost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	MenuId uint64 `protobuf:"varint,2,opt,name=menu_id,json=menuId,proto3" json:"menu_id,omitempty"`
	Code   string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Name   string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Urls   string `protobuf:"bytes,5,opt,name=urls,proto3" json:"urls,omitempty"`
	Sort   uint64 `protobuf:"varint,6,opt,name=sort,proto3" json:"sort,omitempty"`
	Remark string `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *UpdateRulePost) Reset() {
	*x = UpdateRulePost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_rule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRulePost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRulePost) ProtoMessage() {}

func (x *UpdateRulePost) ProtoReflect() protoreflect.Message {
	mi := &file_console_rule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRulePost.ProtoReflect.Descriptor instead.
func (*UpdateRulePost) Descriptor() ([]byte, []int) {
	return file_console_rule_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRulePost) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UpdateRulePost) GetMenuId() uint64 {
	if x != nil {
		return x.MenuId
	}
	return 0
}

func (x *UpdateRulePost) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateRulePost) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRulePost) GetUrls() string {
	if x != nil {
		return x.Urls
	}
	return ""
}

func (x *UpdateRulePost) GetSort() uint64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *UpdateRulePost) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

// ??????????????????
type RuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Data    *Rule  `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *RuleResponse) Reset() {
	*x = RuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_rule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleResponse) ProtoMessage() {}

func (x *RuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_console_rule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleResponse.ProtoReflect.Descriptor instead.
func (*RuleResponse) Descriptor() ([]byte, []int) {
	return file_console_rule_proto_rawDescGZIP(), []int{3}
}

func (x *RuleResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RuleResponse) GetData() *Rule {
	if x != nil {
		return x.Data
	}
	return nil
}

// ??????????????????
type RulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string  `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Data    []*Rule `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *RulesResponse) Reset() {
	*x = RulesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_rule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RulesResponse) ProtoMessage() {}

func (x *RulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_console_rule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RulesResponse.ProtoReflect.Descriptor instead.
func (*RulesResponse) Descriptor() ([]byte, []int) {
	return file_console_rule_proto_rawDescGZIP(), []int{4}
}

func (x *RulesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RulesResponse) GetData() []*Rule {
	if x != nil {
		return x.Data
	}
	return nil
}

// ??????????????????
type RulePaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                        `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Data    *RulePaginationResponse_Rules `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *RulePaginationResponse) Reset() {
	*x = RulePaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_rule_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RulePaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RulePaginationResponse) ProtoMessage() {}

func (x *RulePaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_console_rule_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RulePaginationResponse.ProtoReflect.Descriptor instead.
func (*RulePaginationResponse) Descriptor() ([]byte, []int) {
	return file_console_rule_proto_rawDescGZIP(), []int{5}
}

func (x *RulePaginationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RulePaginationResponse) GetData() *RulePaginationResponse_Rules {
	if x != nil {
		return x.Data
	}
	return nil
}

type RulePaginationResponse_Rules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*Rule `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total int64   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *RulePaginationResponse_Rules) Reset() {
	*x = RulePaginationResponse_Rules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_rule_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RulePaginationResponse_Rules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RulePaginationResponse_Rules) ProtoMessage() {}

func (x *RulePaginationResponse_Rules) ProtoReflect() protoreflect.Message {
	mi := &file_console_rule_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RulePaginationResponse_Rules.ProtoReflect.Descriptor instead.
func (*RulePaginationResponse_Rules) Descriptor() ([]byte, []int) {
	return file_console_rule_proto_rawDescGZIP(), []int{5, 0}
}

func (x *RulePaginationResponse_Rules) GetList() []*Rule {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *RulePaginationResponse_Rules) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_console_rule_proto protoreflect.FileDescriptor

var file_console_rule_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x22, 0xef, 0x01,
	0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x91, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x22, 0xa3, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x75,
	0x6c, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x75,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x4b, 0x0a, 0x0c, 0x52, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x4c, 0x0a, 0x0d, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x21, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x22, 0xaf, 0x01, 0x0a, 0x16, 0x52, 0x75, 0x6c, 0x65, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c,
	0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x40, 0x0a, 0x05, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x21, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_console_rule_proto_rawDescOnce sync.Once
	file_console_rule_proto_rawDescData = file_console_rule_proto_rawDesc
)

func file_console_rule_proto_rawDescGZIP() []byte {
	file_console_rule_proto_rawDescOnce.Do(func() {
		file_console_rule_proto_rawDescData = protoimpl.X.CompressGZIP(file_console_rule_proto_rawDescData)
	})
	return file_console_rule_proto_rawDescData
}

var file_console_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_console_rule_proto_goTypes = []interface{}{
	(*Rule)(nil),                         // 0: console.Rule
	(*CreateRulePost)(nil),               // 1: console.CreateRulePost
	(*UpdateRulePost)(nil),               // 2: console.UpdateRulePost
	(*RuleResponse)(nil),                 // 3: console.RuleResponse
	(*RulesResponse)(nil),                // 4: console.RulesResponse
	(*RulePaginationResponse)(nil),       // 5: console.RulePaginationResponse
	(*RulePaginationResponse_Rules)(nil), // 6: console.RulePaginationResponse.Rules
}
var file_console_rule_proto_depIdxs = []int32{
	0, // 0: console.RuleResponse.Data:type_name -> console.Rule
	0, // 1: console.RulesResponse.Data:type_name -> console.Rule
	6, // 2: console.RulePaginationResponse.Data:type_name -> console.RulePaginationResponse.Rules
	0, // 3: console.RulePaginationResponse.Rules.list:type_name -> console.Rule
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_console_rule_proto_init() }
func file_console_rule_proto_init() {
	if File_console_rule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_console_rule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
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
		file_console_rule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRulePost); i {
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
		file_console_rule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRulePost); i {
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
		file_console_rule_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuleResponse); i {
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
		file_console_rule_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RulesResponse); i {
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
		file_console_rule_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RulePaginationResponse); i {
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
		file_console_rule_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RulePaginationResponse_Rules); i {
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
			RawDescriptor: file_console_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_console_rule_proto_goTypes,
		DependencyIndexes: file_console_rule_proto_depIdxs,
		MessageInfos:      file_console_rule_proto_msgTypes,
	}.Build()
	File_console_rule_proto = out.File
	file_console_rule_proto_rawDesc = nil
	file_console_rule_proto_goTypes = nil
	file_console_rule_proto_depIdxs = nil
}
