// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: console/user.proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Username  string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Phone     string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Email     string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Language  string   `protobuf:"bytes,5,opt,name=language,proto3" json:"language,omitempty"`
	Roles     []string `protobuf:"bytes,6,rep,name=roles,proto3" json:"roles,omitempty"`
	Status    uint64   `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	Checked   bool     `protobuf:"varint,8,opt,name=checked,proto3" json:"checked,omitempty"`
	Remark    string   `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark,omitempty"`
	LoginAt   string   `protobuf:"bytes,10,opt,name=login_at,json=loginAt,proto3" json:"login_at,omitempty"`
	CreatedAt string   `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string   `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_console_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_console_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *User) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *User) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *User) GetChecked() bool {
	if x != nil {
		return x.Checked
	}
	return false
}

func (x *User) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *User) GetLoginAt() string {
	if x != nil {
		return x.LoginAt
	}
	return ""
}

func (x *User) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *User) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateUserPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Phone    string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Email    string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Language string   `protobuf:"bytes,6,opt,name=language,proto3" json:"language,omitempty"`
	Remark   string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark,omitempty"`
	Roles    []string `protobuf:"bytes,8,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *CreateUserPost) Reset() {
	*x = CreateUserPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserPost) ProtoMessage() {}

func (x *CreateUserPost) ProtoReflect() protoreflect.Message {
	mi := &file_console_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserPost.ProtoReflect.Descriptor instead.
func (*CreateUserPost) Descriptor() ([]byte, []int) {
	return file_console_user_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserPost) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserPost) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserPost) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateUserPost) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserPost) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *CreateUserPost) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *CreateUserPost) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type UpdateUserPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Username string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Phone    string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Email    string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Language string   `protobuf:"bytes,6,opt,name=language,proto3" json:"language,omitempty"`
	Remark   string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark,omitempty"`
	Roles    []string `protobuf:"bytes,8,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *UpdateUserPost) Reset() {
	*x = UpdateUserPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserPost) ProtoMessage() {}

func (x *UpdateUserPost) ProtoReflect() protoreflect.Message {
	mi := &file_console_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserPost.ProtoReflect.Descriptor instead.
func (*UpdateUserPost) Descriptor() ([]byte, []int) {
	return file_console_user_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateUserPost) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UpdateUserPost) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdateUserPost) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateUserPost) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateUserPost) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateUserPost) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *UpdateUserPost) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *UpdateUserPost) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type ChangeUserPasswordPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Origin   string `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ChangeUserPasswordPost) Reset() {
	*x = ChangeUserPasswordPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeUserPasswordPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeUserPasswordPost) ProtoMessage() {}

func (x *ChangeUserPasswordPost) ProtoReflect() protoreflect.Message {
	mi := &file_console_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeUserPasswordPost.ProtoReflect.Descriptor instead.
func (*ChangeUserPasswordPost) Descriptor() ([]byte, []int) {
	return file_console_user_proto_rawDescGZIP(), []int{3}
}

func (x *ChangeUserPasswordPost) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ChangeUserPasswordPost) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *ChangeUserPasswordPost) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ScopesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Data    string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ScopesResponse) Reset() {
	*x = ScopesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScopesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopesResponse) ProtoMessage() {}

func (x *ScopesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_console_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopesResponse.ProtoReflect.Descriptor instead.
func (*ScopesResponse) Descriptor() ([]byte, []int) {
	return file_console_user_proto_rawDescGZIP(), []int{4}
}

func (x *ScopesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ScopesResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// ??????????????????
type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Data    *User  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_console_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_console_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserResponse) GetData() *User {
	if x != nil {
		return x.Data
	}
	return nil
}

// ??????????????????
type UsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string  `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*User `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *UsersResponse) Reset() {
	*x = UsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersResponse) ProtoMessage() {}

func (x *UsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_console_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersResponse.ProtoReflect.Descriptor instead.
func (*UsersResponse) Descriptor() ([]byte, []int) {
	return file_console_user_proto_rawDescGZIP(), []int{6}
}

func (x *UsersResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UsersResponse) GetData() []*User {
	if x != nil {
		return x.Data
	}
	return nil
}

// ??????????????????
type UserPaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                        `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Data    *UserPaginationResponse_Users `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UserPaginationResponse) Reset() {
	*x = UserPaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPaginationResponse) ProtoMessage() {}

func (x *UserPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_console_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPaginationResponse.ProtoReflect.Descriptor instead.
func (*UserPaginationResponse) Descriptor() ([]byte, []int) {
	return file_console_user_proto_rawDescGZIP(), []int{7}
}

func (x *UserPaginationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserPaginationResponse) GetData() *UserPaginationResponse_Users {
	if x != nil {
		return x.Data
	}
	return nil
}

type UserPaginationResponse_Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*User `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total int64   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *UserPaginationResponse_Users) Reset() {
	*x = UserPaginationResponse_Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPaginationResponse_Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPaginationResponse_Users) ProtoMessage() {}

func (x *UserPaginationResponse_Users) ProtoReflect() protoreflect.Message {
	mi := &file_console_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPaginationResponse_Users.ProtoReflect.Descriptor instead.
func (*UserPaginationResponse_Users) Descriptor() ([]byte, []int) {
	return file_console_user_proto_rawDescGZIP(), []int{7, 0}
}

func (x *UserPaginationResponse_Users) GetList() []*User {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *UserPaginationResponse_Users) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_console_user_proto protoreflect.FileDescriptor

var file_console_user_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x22, 0xb5, 0x02,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x19, 0x0a, 0x08,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xbe, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0xd0, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x5e, 0x0a, 0x16, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3e, 0x0a, 0x0e, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4b, 0x0a, 0x0c, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4c, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xaf, 0x01, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x40, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x21, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_console_user_proto_rawDescOnce sync.Once
	file_console_user_proto_rawDescData = file_console_user_proto_rawDesc
)

func file_console_user_proto_rawDescGZIP() []byte {
	file_console_user_proto_rawDescOnce.Do(func() {
		file_console_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_console_user_proto_rawDescData)
	})
	return file_console_user_proto_rawDescData
}

var file_console_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_console_user_proto_goTypes = []interface{}{
	(*User)(nil),                         // 0: console.User
	(*CreateUserPost)(nil),               // 1: console.CreateUserPost
	(*UpdateUserPost)(nil),               // 2: console.UpdateUserPost
	(*ChangeUserPasswordPost)(nil),       // 3: console.ChangeUserPasswordPost
	(*ScopesResponse)(nil),               // 4: console.ScopesResponse
	(*UserResponse)(nil),                 // 5: console.UserResponse
	(*UsersResponse)(nil),                // 6: console.UsersResponse
	(*UserPaginationResponse)(nil),       // 7: console.UserPaginationResponse
	(*UserPaginationResponse_Users)(nil), // 8: console.UserPaginationResponse.Users
}
var file_console_user_proto_depIdxs = []int32{
	0, // 0: console.UserResponse.data:type_name -> console.User
	0, // 1: console.UsersResponse.data:type_name -> console.User
	8, // 2: console.UserPaginationResponse.data:type_name -> console.UserPaginationResponse.Users
	0, // 3: console.UserPaginationResponse.Users.list:type_name -> console.User
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_console_user_proto_init() }
func file_console_user_proto_init() {
	if File_console_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_console_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_console_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserPost); i {
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
		file_console_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserPost); i {
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
		file_console_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeUserPasswordPost); i {
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
		file_console_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScopesResponse); i {
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
		file_console_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
		file_console_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersResponse); i {
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
		file_console_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPaginationResponse); i {
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
		file_console_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPaginationResponse_Users); i {
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
			RawDescriptor: file_console_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_console_user_proto_goTypes,
		DependencyIndexes: file_console_user_proto_depIdxs,
		MessageInfos:      file_console_user_proto_msgTypes,
	}.Build()
	File_console_user_proto = out.File
	file_console_user_proto_rawDesc = nil
	file_console_user_proto_goTypes = nil
	file_console_user_proto_depIdxs = nil
}
