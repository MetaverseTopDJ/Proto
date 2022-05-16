// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: console/menu.proto

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

type TopMenu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MenuId uint64 `protobuf:"varint,2,opt,name=menu_id,json=menuId,proto3" json:"menu_id,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	EnName string `protobuf:"bytes,4,opt,name=en_name,json=enName,proto3" json:"en_name,omitempty"`
	Status uint64 `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *TopMenu) Reset() {
	*x = TopMenu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_menu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopMenu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopMenu) ProtoMessage() {}

func (x *TopMenu) ProtoReflect() protoreflect.Message {
	mi := &file_console_menu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopMenu.ProtoReflect.Descriptor instead.
func (*TopMenu) Descriptor() ([]byte, []int) {
	return file_console_menu_proto_rawDescGZIP(), []int{0}
}

func (x *TopMenu) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TopMenu) GetMenuId() uint64 {
	if x != nil {
		return x.MenuId
	}
	return 0
}

func (x *TopMenu) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TopMenu) GetEnName() string {
	if x != nil {
		return x.EnName
	}
	return ""
}

func (x *TopMenu) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type MenuRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuId uint64 `protobuf:"varint,1,opt,name=menu_id,json=menuId,proto3" json:"menu_id,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Code   string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Name   string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Sort   uint64 `protobuf:"varint,5,opt,name=sort,proto3" json:"sort,omitempty"`
	Value  string `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	Label  string `protobuf:"bytes,7,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *MenuRule) Reset() {
	*x = MenuRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_menu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuRule) ProtoMessage() {}

func (x *MenuRule) ProtoReflect() protoreflect.Message {
	mi := &file_console_menu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuRule.ProtoReflect.Descriptor instead.
func (*MenuRule) Descriptor() ([]byte, []int) {
	return file_console_menu_proto_rawDescGZIP(), []int{1}
}

func (x *MenuRule) GetMenuId() uint64 {
	if x != nil {
		return x.MenuId
	}
	return 0
}

func (x *MenuRule) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MenuRule) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MenuRule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MenuRule) GetSort() uint64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *MenuRule) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *MenuRule) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type Menu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MenuId    uint64      `protobuf:"varint,2,opt,name=menu_id,json=menuId,proto3" json:"menu_id,omitempty"`
	Sort      uint64      `protobuf:"varint,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Name      string      `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	EnName    string      `protobuf:"bytes,5,opt,name=en_name,json=enName,proto3" json:"en_name,omitempty"`
	Status    uint64      `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Checked   bool        `protobuf:"varint,7,opt,name=checked,proto3" json:"checked,omitempty"`
	Remark    string      `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	CreatedAt string      `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string      `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	TopMenu   *TopMenu    `protobuf:"bytes,11,opt,name=top_menu,json=topMenu,proto3" json:"top_menu,omitempty"`
	Rules     []*MenuRule `protobuf:"bytes,12,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *Menu) Reset() {
	*x = Menu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_menu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Menu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Menu) ProtoMessage() {}

func (x *Menu) ProtoReflect() protoreflect.Message {
	mi := &file_console_menu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Menu.ProtoReflect.Descriptor instead.
func (*Menu) Descriptor() ([]byte, []int) {
	return file_console_menu_proto_rawDescGZIP(), []int{2}
}

func (x *Menu) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Menu) GetMenuId() uint64 {
	if x != nil {
		return x.MenuId
	}
	return 0
}

func (x *Menu) GetSort() uint64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Menu) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Menu) GetEnName() string {
	if x != nil {
		return x.EnName
	}
	return ""
}

func (x *Menu) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Menu) GetChecked() bool {
	if x != nil {
		return x.Checked
	}
	return false
}

func (x *Menu) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *Menu) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Menu) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Menu) GetTopMenu() *TopMenu {
	if x != nil {
		return x.TopMenu
	}
	return nil
}

func (x *Menu) GetRules() []*MenuRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type CreateMenuPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuId uint64 `protobuf:"varint,2,opt,name=menu_id,json=menuId,proto3" json:"menu_id,omitempty"`
	Sort   uint64 `protobuf:"varint,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Name   string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	EnName string `protobuf:"bytes,5,opt,name=en_name,json=enName,proto3" json:"en_name,omitempty"`
	Remark string `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *CreateMenuPost) Reset() {
	*x = CreateMenuPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_menu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMenuPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMenuPost) ProtoMessage() {}

func (x *CreateMenuPost) ProtoReflect() protoreflect.Message {
	mi := &file_console_menu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMenuPost.ProtoReflect.Descriptor instead.
func (*CreateMenuPost) Descriptor() ([]byte, []int) {
	return file_console_menu_proto_rawDescGZIP(), []int{3}
}

func (x *CreateMenuPost) GetMenuId() uint64 {
	if x != nil {
		return x.MenuId
	}
	return 0
}

func (x *CreateMenuPost) GetSort() uint64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *CreateMenuPost) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateMenuPost) GetEnName() string {
	if x != nil {
		return x.EnName
	}
	return ""
}

func (x *CreateMenuPost) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type UpdateMenuPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MenuId uint64 `protobuf:"varint,2,opt,name=menu_id,json=menuId,proto3" json:"menu_id,omitempty"`
	Sort   uint64 `protobuf:"varint,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Name   string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	EnName string `protobuf:"bytes,5,opt,name=en_name,json=enName,proto3" json:"en_name,omitempty"`
	Remark string `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *UpdateMenuPost) Reset() {
	*x = UpdateMenuPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_menu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMenuPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMenuPost) ProtoMessage() {}

func (x *UpdateMenuPost) ProtoReflect() protoreflect.Message {
	mi := &file_console_menu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMenuPost.ProtoReflect.Descriptor instead.
func (*UpdateMenuPost) Descriptor() ([]byte, []int) {
	return file_console_menu_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateMenuPost) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateMenuPost) GetMenuId() uint64 {
	if x != nil {
		return x.MenuId
	}
	return 0
}

func (x *UpdateMenuPost) GetSort() uint64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *UpdateMenuPost) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateMenuPost) GetEnName() string {
	if x != nil {
		return x.EnName
	}
	return ""
}

func (x *UpdateMenuPost) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

// 返回单条数据
type MenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Data    *Menu  `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *MenuResponse) Reset() {
	*x = MenuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_menu_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuResponse) ProtoMessage() {}

func (x *MenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_console_menu_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuResponse.ProtoReflect.Descriptor instead.
func (*MenuResponse) Descriptor() ([]byte, []int) {
	return file_console_menu_proto_rawDescGZIP(), []int{5}
}

func (x *MenuResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MenuResponse) GetData() *Menu {
	if x != nil {
		return x.Data
	}
	return nil
}

// 返回列表数据
type MenusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string  `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Data    []*Menu `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *MenusResponse) Reset() {
	*x = MenusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_menu_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenusResponse) ProtoMessage() {}

func (x *MenusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_console_menu_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenusResponse.ProtoReflect.Descriptor instead.
func (*MenusResponse) Descriptor() ([]byte, []int) {
	return file_console_menu_proto_rawDescGZIP(), []int{6}
}

func (x *MenusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MenusResponse) GetData() []*Menu {
	if x != nil {
		return x.Data
	}
	return nil
}

// 菜单 包含权限
type MenuRulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string  `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Data    []*Menu `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *MenuRulesResponse) Reset() {
	*x = MenuRulesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_menu_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuRulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuRulesResponse) ProtoMessage() {}

func (x *MenuRulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_console_menu_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuRulesResponse.ProtoReflect.Descriptor instead.
func (*MenuRulesResponse) Descriptor() ([]byte, []int) {
	return file_console_menu_proto_rawDescGZIP(), []int{7}
}

func (x *MenuRulesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MenuRulesResponse) GetData() []*Menu {
	if x != nil {
		return x.Data
	}
	return nil
}

// 返回分页数据
type MenuPaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                        `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Data    *MenuPaginationResponse_Menus `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *MenuPaginationResponse) Reset() {
	*x = MenuPaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_menu_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuPaginationResponse) ProtoMessage() {}

func (x *MenuPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_console_menu_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuPaginationResponse.ProtoReflect.Descriptor instead.
func (*MenuPaginationResponse) Descriptor() ([]byte, []int) {
	return file_console_menu_proto_rawDescGZIP(), []int{8}
}

func (x *MenuPaginationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MenuPaginationResponse) GetData() *MenuPaginationResponse_Menus {
	if x != nil {
		return x.Data
	}
	return nil
}

type MenuPaginationResponse_Menus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*Menu `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total int64   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *MenuPaginationResponse_Menus) Reset() {
	*x = MenuPaginationResponse_Menus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_menu_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuPaginationResponse_Menus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuPaginationResponse_Menus) ProtoMessage() {}

func (x *MenuPaginationResponse_Menus) ProtoReflect() protoreflect.Message {
	mi := &file_console_menu_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuPaginationResponse_Menus.ProtoReflect.Descriptor instead.
func (*MenuPaginationResponse_Menus) Descriptor() ([]byte, []int) {
	return file_console_menu_proto_rawDescGZIP(), []int{8, 0}
}

func (x *MenuPaginationResponse_Menus) GetList() []*Menu {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *MenuPaginationResponse_Menus) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_console_menu_proto protoreflect.FileDescriptor

var file_console_menu_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x22, 0x77, 0x0a,
	0x07, 0x54, 0x6f, 0x70, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x75,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x52,
	0x75, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0xce, 0x02, 0x0a, 0x04, 0x4d, 0x65, 0x6e, 0x75, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x65, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x2b, 0x0a, 0x08, 0x74, 0x6f, 0x70, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x54, 0x6f,
	0x70, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x27,
	0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x65,
	0x6e, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6d, 0x65, 0x6e,
	0x75, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x65,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x92, 0x01, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x65, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x22, 0x4b, 0x0a, 0x0c, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x6f, 0x6c, 0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x4c,
	0x0a, 0x0d, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c,
	0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x50, 0x0a, 0x11,
	0x4d, 0x65, 0x6e, 0x75, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x6f, 0x6c, 0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0xaf,
	0x01, 0x0a, 0x16, 0x4d, 0x65, 0x6e, 0x75, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x75,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x40,
	0x0a, 0x05, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e,
	0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_console_menu_proto_rawDescOnce sync.Once
	file_console_menu_proto_rawDescData = file_console_menu_proto_rawDesc
)

func file_console_menu_proto_rawDescGZIP() []byte {
	file_console_menu_proto_rawDescOnce.Do(func() {
		file_console_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_console_menu_proto_rawDescData)
	})
	return file_console_menu_proto_rawDescData
}

var file_console_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_console_menu_proto_goTypes = []interface{}{
	(*TopMenu)(nil),                      // 0: console.TopMenu
	(*MenuRule)(nil),                     // 1: console.MenuRule
	(*Menu)(nil),                         // 2: console.Menu
	(*CreateMenuPost)(nil),               // 3: console.CreateMenuPost
	(*UpdateMenuPost)(nil),               // 4: console.UpdateMenuPost
	(*MenuResponse)(nil),                 // 5: console.MenuResponse
	(*MenusResponse)(nil),                // 6: console.MenusResponse
	(*MenuRulesResponse)(nil),            // 7: console.MenuRulesResponse
	(*MenuPaginationResponse)(nil),       // 8: console.MenuPaginationResponse
	(*MenuPaginationResponse_Menus)(nil), // 9: console.MenuPaginationResponse.Menus
}
var file_console_menu_proto_depIdxs = []int32{
	0, // 0: console.Menu.top_menu:type_name -> console.TopMenu
	1, // 1: console.Menu.rules:type_name -> console.MenuRule
	2, // 2: console.MenuResponse.Data:type_name -> console.Menu
	2, // 3: console.MenusResponse.Data:type_name -> console.Menu
	2, // 4: console.MenuRulesResponse.Data:type_name -> console.Menu
	9, // 5: console.MenuPaginationResponse.Data:type_name -> console.MenuPaginationResponse.Menus
	2, // 6: console.MenuPaginationResponse.Menus.list:type_name -> console.Menu
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_console_menu_proto_init() }
func file_console_menu_proto_init() {
	if File_console_menu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_console_menu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopMenu); i {
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
		file_console_menu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuRule); i {
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
		file_console_menu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Menu); i {
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
		file_console_menu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMenuPost); i {
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
		file_console_menu_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMenuPost); i {
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
		file_console_menu_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuResponse); i {
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
		file_console_menu_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenusResponse); i {
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
		file_console_menu_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuRulesResponse); i {
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
		file_console_menu_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuPaginationResponse); i {
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
		file_console_menu_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuPaginationResponse_Menus); i {
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
			RawDescriptor: file_console_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_console_menu_proto_goTypes,
		DependencyIndexes: file_console_menu_proto_depIdxs,
		MessageInfos:      file_console_menu_proto_msgTypes,
	}.Build()
	File_console_menu_proto = out.File
	file_console_menu_proto_rawDesc = nil
	file_console_menu_proto_goTypes = nil
	file_console_menu_proto_depIdxs = nil
}
