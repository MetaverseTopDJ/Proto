// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: collect/author.proto

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

// Author 作者数据结构
type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	AvatarId        uint64 `protobuf:"varint,2,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id"`
	FirstName       string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name"`
	LastName        string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name"`
	Title           string `protobuf:"bytes,5,opt,name=title,proto3" json:"title"`
	BackgroundOut   string `protobuf:"bytes,6,opt,name=background_out,json=backgroundOut,proto3" json:"background_out"`
	BackgroundInner string `protobuf:"bytes,7,opt,name=background_inner,json=backgroundInner,proto3" json:"background_inner"`
	Hot             bool   `protobuf:"varint,8,opt,name=hot,proto3" json:"hot"`
	Facebook        string `protobuf:"bytes,9,opt,name=facebook,proto3" json:"facebook"`
	Twitter         string `protobuf:"bytes,10,opt,name=twitter,proto3" json:"twitter"`
	Follower        uint64 `protobuf:"varint,11,opt,name=follower,proto3" json:"follower"`
	Status          uint64 `protobuf:"varint,12,opt,name=status,proto3" json:"status"`
	Checked         bool   `protobuf:"varint,13,opt,name=checked,proto3" json:"checked"`
	CreatedAt       string `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt       string `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_author_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_collect_author_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_collect_author_proto_rawDescGZIP(), []int{0}
}

func (x *Author) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Author) GetAvatarId() uint64 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *Author) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Author) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Author) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Author) GetBackgroundOut() string {
	if x != nil {
		return x.BackgroundOut
	}
	return ""
}

func (x *Author) GetBackgroundInner() string {
	if x != nil {
		return x.BackgroundInner
	}
	return ""
}

func (x *Author) GetHot() bool {
	if x != nil {
		return x.Hot
	}
	return false
}

func (x *Author) GetFacebook() string {
	if x != nil {
		return x.Facebook
	}
	return ""
}

func (x *Author) GetTwitter() string {
	if x != nil {
		return x.Twitter
	}
	return ""
}

func (x *Author) GetFollower() uint64 {
	if x != nil {
		return x.Follower
	}
	return 0
}

func (x *Author) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Author) GetChecked() bool {
	if x != nil {
		return x.Checked
	}
	return false
}

func (x *Author) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Author) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// SimpleAuthor 返回前端的作者数据
type SimpleAuthor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	AvatarUrl       string `protobuf:"bytes,2,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url"`
	FirstName       string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name"`
	LastName        string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name"`
	Title           string `protobuf:"bytes,5,opt,name=title,proto3" json:"title"`
	BackgroundOut   string `protobuf:"bytes,6,opt,name=background_out,json=backgroundOut,proto3" json:"background_out"`
	BackgroundInner string `protobuf:"bytes,7,opt,name=background_inner,json=backgroundInner,proto3" json:"background_inner"`
	Hot             bool   `protobuf:"varint,8,opt,name=hot,proto3" json:"hot"`
	Facebook        string `protobuf:"bytes,9,opt,name=facebook,proto3" json:"facebook"`
	Twitter         string `protobuf:"bytes,10,opt,name=twitter,proto3" json:"twitter"`
	Follower        uint64 `protobuf:"varint,11,opt,name=follower,proto3" json:"follower"`
	Status          uint64 `protobuf:"varint,12,opt,name=status,proto3" json:"status"`
}

func (x *SimpleAuthor) Reset() {
	*x = SimpleAuthor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_author_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleAuthor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleAuthor) ProtoMessage() {}

func (x *SimpleAuthor) ProtoReflect() protoreflect.Message {
	mi := &file_collect_author_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleAuthor.ProtoReflect.Descriptor instead.
func (*SimpleAuthor) Descriptor() ([]byte, []int) {
	return file_collect_author_proto_rawDescGZIP(), []int{1}
}

func (x *SimpleAuthor) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SimpleAuthor) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *SimpleAuthor) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *SimpleAuthor) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *SimpleAuthor) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SimpleAuthor) GetBackgroundOut() string {
	if x != nil {
		return x.BackgroundOut
	}
	return ""
}

func (x *SimpleAuthor) GetBackgroundInner() string {
	if x != nil {
		return x.BackgroundInner
	}
	return ""
}

func (x *SimpleAuthor) GetHot() bool {
	if x != nil {
		return x.Hot
	}
	return false
}

func (x *SimpleAuthor) GetFacebook() string {
	if x != nil {
		return x.Facebook
	}
	return ""
}

func (x *SimpleAuthor) GetTwitter() string {
	if x != nil {
		return x.Twitter
	}
	return ""
}

func (x *SimpleAuthor) GetFollower() uint64 {
	if x != nil {
		return x.Follower
	}
	return 0
}

func (x *SimpleAuthor) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

// CreateAuthorPost 创建作者
type CreateAuthorPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarId        uint64 `protobuf:"varint,2,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id"`
	FirstName       string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name"`
	LastName        string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name"`
	Title           string `protobuf:"bytes,5,opt,name=title,proto3" json:"title"`
	BackgroundOut   string `protobuf:"bytes,6,opt,name=background_out,json=backgroundOut,proto3" json:"background_out"`
	BackgroundInner string `protobuf:"bytes,7,opt,name=background_inner,json=backgroundInner,proto3" json:"background_inner"`
	Hot             bool   `protobuf:"varint,8,opt,name=hot,proto3" json:"hot"`
	Facebook        string `protobuf:"bytes,9,opt,name=facebook,proto3" json:"facebook"`
	Twitter         string `protobuf:"bytes,10,opt,name=twitter,proto3" json:"twitter"`
	Follower        uint64 `protobuf:"varint,11,opt,name=follower,proto3" json:"follower"`
}

func (x *CreateAuthorPost) Reset() {
	*x = CreateAuthorPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_author_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthorPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthorPost) ProtoMessage() {}

func (x *CreateAuthorPost) ProtoReflect() protoreflect.Message {
	mi := &file_collect_author_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthorPost.ProtoReflect.Descriptor instead.
func (*CreateAuthorPost) Descriptor() ([]byte, []int) {
	return file_collect_author_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAuthorPost) GetAvatarId() uint64 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *CreateAuthorPost) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateAuthorPost) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateAuthorPost) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateAuthorPost) GetBackgroundOut() string {
	if x != nil {
		return x.BackgroundOut
	}
	return ""
}

func (x *CreateAuthorPost) GetBackgroundInner() string {
	if x != nil {
		return x.BackgroundInner
	}
	return ""
}

func (x *CreateAuthorPost) GetHot() bool {
	if x != nil {
		return x.Hot
	}
	return false
}

func (x *CreateAuthorPost) GetFacebook() string {
	if x != nil {
		return x.Facebook
	}
	return ""
}

func (x *CreateAuthorPost) GetTwitter() string {
	if x != nil {
		return x.Twitter
	}
	return ""
}

func (x *CreateAuthorPost) GetFollower() uint64 {
	if x != nil {
		return x.Follower
	}
	return 0
}

// UpdateAuthorPost 更新作者
type UpdateAuthorPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	AvatarId        uint64 `protobuf:"varint,2,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id"`
	FirstName       string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name"`
	LastName        string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name"`
	Title           string `protobuf:"bytes,5,opt,name=title,proto3" json:"title"`
	BackgroundOut   string `protobuf:"bytes,6,opt,name=background_out,json=backgroundOut,proto3" json:"background_out"`
	BackgroundInner string `protobuf:"bytes,7,opt,name=background_inner,json=backgroundInner,proto3" json:"background_inner"`
	Hot             bool   `protobuf:"varint,8,opt,name=hot,proto3" json:"hot"`
	Facebook        string `protobuf:"bytes,9,opt,name=facebook,proto3" json:"facebook"`
	Twitter         string `protobuf:"bytes,10,opt,name=twitter,proto3" json:"twitter"`
	Follower        uint64 `protobuf:"varint,11,opt,name=follower,proto3" json:"follower"`
}

func (x *UpdateAuthorPost) Reset() {
	*x = UpdateAuthorPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_author_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAuthorPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuthorPost) ProtoMessage() {}

func (x *UpdateAuthorPost) ProtoReflect() protoreflect.Message {
	mi := &file_collect_author_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuthorPost.ProtoReflect.Descriptor instead.
func (*UpdateAuthorPost) Descriptor() ([]byte, []int) {
	return file_collect_author_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAuthorPost) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateAuthorPost) GetAvatarId() uint64 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *UpdateAuthorPost) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdateAuthorPost) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdateAuthorPost) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateAuthorPost) GetBackgroundOut() string {
	if x != nil {
		return x.BackgroundOut
	}
	return ""
}

func (x *UpdateAuthorPost) GetBackgroundInner() string {
	if x != nil {
		return x.BackgroundInner
	}
	return ""
}

func (x *UpdateAuthorPost) GetHot() bool {
	if x != nil {
		return x.Hot
	}
	return false
}

func (x *UpdateAuthorPost) GetFacebook() string {
	if x != nil {
		return x.Facebook
	}
	return ""
}

func (x *UpdateAuthorPost) GetTwitter() string {
	if x != nil {
		return x.Twitter
	}
	return ""
}

func (x *UpdateAuthorPost) GetFollower() uint64 {
	if x != nil {
		return x.Follower
	}
	return 0
}

// AuthorResponse 返回
type AuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string  `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    *Author `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *AuthorResponse) Reset() {
	*x = AuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_author_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorResponse) ProtoMessage() {}

func (x *AuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_collect_author_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorResponse.ProtoReflect.Descriptor instead.
func (*AuthorResponse) Descriptor() ([]byte, []int) {
	return file_collect_author_proto_rawDescGZIP(), []int{4}
}

func (x *AuthorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AuthorResponse) GetData() *Author {
	if x != nil {
		return x.Data
	}
	return nil
}

type SimpleAuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string        `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    *SimpleAuthor `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *SimpleAuthorResponse) Reset() {
	*x = SimpleAuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_author_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleAuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleAuthorResponse) ProtoMessage() {}

func (x *SimpleAuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_collect_author_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleAuthorResponse.ProtoReflect.Descriptor instead.
func (*SimpleAuthorResponse) Descriptor() ([]byte, []int) {
	return file_collect_author_proto_rawDescGZIP(), []int{5}
}

func (x *SimpleAuthorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SimpleAuthorResponse) GetData() *SimpleAuthor {
	if x != nil {
		return x.Data
	}
	return nil
}

// HotAuthorsResponse 返回 热门艺术家
type HotAuthorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string          `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    []*SimpleAuthor `protobuf:"bytes,2,rep,name=data,proto3" json:"data"`
}

func (x *HotAuthorsResponse) Reset() {
	*x = HotAuthorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_author_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HotAuthorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotAuthorsResponse) ProtoMessage() {}

func (x *HotAuthorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_collect_author_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotAuthorsResponse.ProtoReflect.Descriptor instead.
func (*HotAuthorsResponse) Descriptor() ([]byte, []int) {
	return file_collect_author_proto_rawDescGZIP(), []int{6}
}

func (x *HotAuthorsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HotAuthorsResponse) GetData() []*SimpleAuthor {
	if x != nil {
		return x.Data
	}
	return nil
}

// AuthorPaginationResponse 返回 分页
type AuthorPaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                                  `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    *AuthorPaginationResponse_SimpleAuthors `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *AuthorPaginationResponse) Reset() {
	*x = AuthorPaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_author_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorPaginationResponse) ProtoMessage() {}

func (x *AuthorPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_collect_author_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorPaginationResponse.ProtoReflect.Descriptor instead.
func (*AuthorPaginationResponse) Descriptor() ([]byte, []int) {
	return file_collect_author_proto_rawDescGZIP(), []int{7}
}

func (x *AuthorPaginationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AuthorPaginationResponse) GetData() *AuthorPaginationResponse_SimpleAuthors {
	if x != nil {
		return x.Data
	}
	return nil
}

type AuthorPaginationResponse_SimpleAuthors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*SimpleAuthor `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
	Total int64           `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
}

func (x *AuthorPaginationResponse_SimpleAuthors) Reset() {
	*x = AuthorPaginationResponse_SimpleAuthors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_author_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorPaginationResponse_SimpleAuthors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorPaginationResponse_SimpleAuthors) ProtoMessage() {}

func (x *AuthorPaginationResponse_SimpleAuthors) ProtoReflect() protoreflect.Message {
	mi := &file_collect_author_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorPaginationResponse_SimpleAuthors.ProtoReflect.Descriptor instead.
func (*AuthorPaginationResponse_SimpleAuthors) Descriptor() ([]byte, []int) {
	return file_collect_author_proto_rawDescGZIP(), []int{7, 0}
}

func (x *AuthorPaginationResponse_SimpleAuthors) GetList() []*SimpleAuthor {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *AuthorPaginationResponse_SimpleAuthors) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_collect_author_proto protoreflect.FileDescriptor

var file_collect_author_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x22,
	0xad, 0x03, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x61, 0x63,
	0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4f, 0x75, 0x74,
	0x12, 0x29, 0x0a, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69,
	0x6e, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x63, 0x6b,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x68,
	0x6f, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x68, 0x6f, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x77, 0x69,
	0x74, 0x74, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x77, 0x69, 0x74,
	0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0xdd, 0x02, 0x0a, 0x0c, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12,
	0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f,
	0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x62, 0x61, 0x63, 0x6b,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x6f, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x03, 0x68, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f,
	0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f,
	0x6b, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0xb7, 0x02, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x61, 0x63,
	0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x62, 0x61,
	0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x6f, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x03, 0x68, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x65, 0x62,
	0x6f, 0x6f, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x63, 0x65, 0x62,
	0x6f, 0x6f, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x22, 0xc7, 0x02, 0x0a, 0x10, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x4f, 0x75, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x68, 0x6f, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x68, 0x6f,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x22, 0x4f, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x5b, 0x0a, 0x14, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x59, 0x0a, 0x12, 0x48, 0x6f, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xcb, 0x01, 0x0a,
	0x18, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x50, 0x0a, 0x0d, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collect_author_proto_rawDescOnce sync.Once
	file_collect_author_proto_rawDescData = file_collect_author_proto_rawDesc
)

func file_collect_author_proto_rawDescGZIP() []byte {
	file_collect_author_proto_rawDescOnce.Do(func() {
		file_collect_author_proto_rawDescData = protoimpl.X.CompressGZIP(file_collect_author_proto_rawDescData)
	})
	return file_collect_author_proto_rawDescData
}

var file_collect_author_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_collect_author_proto_goTypes = []interface{}{
	(*Author)(nil),                                 // 0: collect.Author
	(*SimpleAuthor)(nil),                           // 1: collect.SimpleAuthor
	(*CreateAuthorPost)(nil),                       // 2: collect.CreateAuthorPost
	(*UpdateAuthorPost)(nil),                       // 3: collect.UpdateAuthorPost
	(*AuthorResponse)(nil),                         // 4: collect.AuthorResponse
	(*SimpleAuthorResponse)(nil),                   // 5: collect.SimpleAuthorResponse
	(*HotAuthorsResponse)(nil),                     // 6: collect.HotAuthorsResponse
	(*AuthorPaginationResponse)(nil),               // 7: collect.AuthorPaginationResponse
	(*AuthorPaginationResponse_SimpleAuthors)(nil), // 8: collect.AuthorPaginationResponse.SimpleAuthors
}
var file_collect_author_proto_depIdxs = []int32{
	0, // 0: collect.AuthorResponse.data:type_name -> collect.Author
	1, // 1: collect.SimpleAuthorResponse.data:type_name -> collect.SimpleAuthor
	1, // 2: collect.HotAuthorsResponse.data:type_name -> collect.SimpleAuthor
	8, // 3: collect.AuthorPaginationResponse.data:type_name -> collect.AuthorPaginationResponse.SimpleAuthors
	1, // 4: collect.AuthorPaginationResponse.SimpleAuthors.list:type_name -> collect.SimpleAuthor
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_collect_author_proto_init() }
func file_collect_author_proto_init() {
	if File_collect_author_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collect_author_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_collect_author_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleAuthor); i {
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
		file_collect_author_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuthorPost); i {
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
		file_collect_author_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAuthorPost); i {
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
		file_collect_author_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorResponse); i {
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
		file_collect_author_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleAuthorResponse); i {
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
		file_collect_author_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HotAuthorsResponse); i {
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
		file_collect_author_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorPaginationResponse); i {
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
		file_collect_author_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorPaginationResponse_SimpleAuthors); i {
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
			RawDescriptor: file_collect_author_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_collect_author_proto_goTypes,
		DependencyIndexes: file_collect_author_proto_depIdxs,
		MessageInfos:      file_collect_author_proto_msgTypes,
	}.Build()
	File_collect_author_proto = out.File
	file_collect_author_proto_rawDesc = nil
	file_collect_author_proto_goTypes = nil
	file_collect_author_proto_depIdxs = nil
}
