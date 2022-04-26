// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: global/contract.proto

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

// Contract 智能合约
type Contract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Type          uint64 `protobuf:"varint,2,opt,name=type,proto3" json:"type"`
	Address       string `protobuf:"bytes,3,opt,name=address,proto3" json:"address"`
	Noise         string `protobuf:"bytes,4,opt,name=noise,proto3" json:"noise"`
	EncodeKey     string `protobuf:"bytes,5,opt,name=encode_key,json=encodeKey,proto3" json:"encode_key"`
	Node          string `protobuf:"bytes,6,opt,name=node,proto3" json:"node"`
	Interval      uint64 `protobuf:"varint,7,opt,name=interval,proto3" json:"interval"`
	Current       uint64 `protobuf:"varint,8,opt,name=current,proto3" json:"current"`
	Next          uint64 `protobuf:"varint,9,opt,name=next,proto3" json:"next"`
	Exec          uint64 `protobuf:"varint,10,opt,name=exec,proto3" json:"exec"`
	Check         uint64 `protobuf:"varint,11,opt,name=check,proto3" json:"check"`
	BlockInterval uint64 `protobuf:"varint,12,opt,name=block_interval,json=blockInterval,proto3" json:"block_interval"`
	Status        uint64 `protobuf:"varint,13,opt,name=status,proto3" json:"status"`
	Checked       bool   `protobuf:"varint,14,opt,name=checked,proto3" json:"checked"`
	CreatedAt     string `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt     string `protobuf:"bytes,16,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *Contract) Reset() {
	*x = Contract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contract) ProtoMessage() {}

func (x *Contract) ProtoReflect() protoreflect.Message {
	mi := &file_global_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contract.ProtoReflect.Descriptor instead.
func (*Contract) Descriptor() ([]byte, []int) {
	return file_global_contract_proto_rawDescGZIP(), []int{0}
}

func (x *Contract) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Contract) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Contract) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Contract) GetNoise() string {
	if x != nil {
		return x.Noise
	}
	return ""
}

func (x *Contract) GetEncodeKey() string {
	if x != nil {
		return x.EncodeKey
	}
	return ""
}

func (x *Contract) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *Contract) GetInterval() uint64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *Contract) GetCurrent() uint64 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *Contract) GetNext() uint64 {
	if x != nil {
		return x.Next
	}
	return 0
}

func (x *Contract) GetExec() uint64 {
	if x != nil {
		return x.Exec
	}
	return 0
}

func (x *Contract) GetCheck() uint64 {
	if x != nil {
		return x.Check
	}
	return 0
}

func (x *Contract) GetBlockInterval() uint64 {
	if x != nil {
		return x.BlockInterval
	}
	return 0
}

func (x *Contract) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Contract) GetChecked() bool {
	if x != nil {
		return x.Checked
	}
	return false
}

func (x *Contract) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Contract) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// CreateContractPost 创建合约
type CreateContractPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          uint64 `protobuf:"varint,2,opt,name=type,proto3" json:"type"`
	Address       string `protobuf:"bytes,3,opt,name=address,proto3" json:"address"`
	Noise         string `protobuf:"bytes,4,opt,name=noise,proto3" json:"noise"`
	EncodeKey     string `protobuf:"bytes,5,opt,name=encode_key,json=encodeKey,proto3" json:"encode_key"`
	Node          string `protobuf:"bytes,6,opt,name=node,proto3" json:"node"`
	Interval      uint64 `protobuf:"varint,7,opt,name=interval,proto3" json:"interval"`
	Current       uint64 `protobuf:"varint,8,opt,name=current,proto3" json:"current"`
	Next          uint64 `protobuf:"varint,9,opt,name=next,proto3" json:"next"`
	Check         uint64 `protobuf:"varint,10,opt,name=check,proto3" json:"check"`
	BlockInterval uint64 `protobuf:"varint,11,opt,name=block_interval,json=blockInterval,proto3" json:"block_interval"`
	Status        uint64 `protobuf:"varint,12,opt,name=status,proto3" json:"status"`
}

func (x *CreateContractPost) Reset() {
	*x = CreateContractPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContractPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContractPost) ProtoMessage() {}

func (x *CreateContractPost) ProtoReflect() protoreflect.Message {
	mi := &file_global_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContractPost.ProtoReflect.Descriptor instead.
func (*CreateContractPost) Descriptor() ([]byte, []int) {
	return file_global_contract_proto_rawDescGZIP(), []int{1}
}

func (x *CreateContractPost) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *CreateContractPost) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateContractPost) GetNoise() string {
	if x != nil {
		return x.Noise
	}
	return ""
}

func (x *CreateContractPost) GetEncodeKey() string {
	if x != nil {
		return x.EncodeKey
	}
	return ""
}

func (x *CreateContractPost) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *CreateContractPost) GetInterval() uint64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *CreateContractPost) GetCurrent() uint64 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *CreateContractPost) GetNext() uint64 {
	if x != nil {
		return x.Next
	}
	return 0
}

func (x *CreateContractPost) GetCheck() uint64 {
	if x != nil {
		return x.Check
	}
	return 0
}

func (x *CreateContractPost) GetBlockInterval() uint64 {
	if x != nil {
		return x.BlockInterval
	}
	return 0
}

func (x *CreateContractPost) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

// UpdateContractPost 更新合约
type UpdateContractPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Type          uint64 `protobuf:"varint,2,opt,name=type,proto3" json:"type"`
	Address       string `protobuf:"bytes,3,opt,name=address,proto3" json:"address"`
	Noise         string `protobuf:"bytes,4,opt,name=noise,proto3" json:"noise"`
	EncodeKey     string `protobuf:"bytes,5,opt,name=encode_key,json=encodeKey,proto3" json:"encode_key"`
	Node          string `protobuf:"bytes,6,opt,name=node,proto3" json:"node"`
	Interval      uint64 `protobuf:"varint,7,opt,name=interval,proto3" json:"interval"`
	Current       uint64 `protobuf:"varint,8,opt,name=current,proto3" json:"current"`
	Next          uint64 `protobuf:"varint,9,opt,name=next,proto3" json:"next"`
	Check         uint64 `protobuf:"varint,10,opt,name=check,proto3" json:"check"`
	BlockInterval uint64 `protobuf:"varint,11,opt,name=block_interval,json=blockInterval,proto3" json:"block_interval"`
	Status        uint64 `protobuf:"varint,12,opt,name=status,proto3" json:"status"`
}

func (x *UpdateContractPost) Reset() {
	*x = UpdateContractPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateContractPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateContractPost) ProtoMessage() {}

func (x *UpdateContractPost) ProtoReflect() protoreflect.Message {
	mi := &file_global_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateContractPost.ProtoReflect.Descriptor instead.
func (*UpdateContractPost) Descriptor() ([]byte, []int) {
	return file_global_contract_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateContractPost) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateContractPost) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *UpdateContractPost) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateContractPost) GetNoise() string {
	if x != nil {
		return x.Noise
	}
	return ""
}

func (x *UpdateContractPost) GetEncodeKey() string {
	if x != nil {
		return x.EncodeKey
	}
	return ""
}

func (x *UpdateContractPost) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *UpdateContractPost) GetInterval() uint64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *UpdateContractPost) GetCurrent() uint64 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *UpdateContractPost) GetNext() uint64 {
	if x != nil {
		return x.Next
	}
	return 0
}

func (x *UpdateContractPost) GetCheck() uint64 {
	if x != nil {
		return x.Check
	}
	return 0
}

func (x *UpdateContractPost) GetBlockInterval() uint64 {
	if x != nil {
		return x.BlockInterval
	}
	return 0
}

func (x *UpdateContractPost) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

// ContractResponse 返回合约
type ContractResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string    `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    *Contract `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *ContractResponse) Reset() {
	*x = ContractResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractResponse) ProtoMessage() {}

func (x *ContractResponse) ProtoReflect() protoreflect.Message {
	mi := &file_global_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractResponse.ProtoReflect.Descriptor instead.
func (*ContractResponse) Descriptor() ([]byte, []int) {
	return file_global_contract_proto_rawDescGZIP(), []int{3}
}

func (x *ContractResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ContractResponse) GetData() *Contract {
	if x != nil {
		return x.Data
	}
	return nil
}

// ContractAddressResponse 仅返回 合约地址
type ContractAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    string `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *ContractAddressResponse) Reset() {
	*x = ContractAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_contract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractAddressResponse) ProtoMessage() {}

func (x *ContractAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_global_contract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractAddressResponse.ProtoReflect.Descriptor instead.
func (*ContractAddressResponse) Descriptor() ([]byte, []int) {
	return file_global_contract_proto_rawDescGZIP(), []int{4}
}

func (x *ContractAddressResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ContractAddressResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_global_contract_proto protoreflect.FileDescriptor

var file_global_contract_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x22,
	0x9c, 0x03, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f,
	0x69, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x69, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x65, 0x78, 0x65, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x65, 0x78, 0x65,
	0x63, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xaa,
	0x02, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x69, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x69, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x25, 0x0a, 0x0e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xba, 0x02, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x69, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6e, 0x6f, 0x69, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x6e,
	0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x52, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x47, 0x0a, 0x17,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x13, 0x42, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x50,
	0x01, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_global_contract_proto_rawDescOnce sync.Once
	file_global_contract_proto_rawDescData = file_global_contract_proto_rawDesc
)

func file_global_contract_proto_rawDescGZIP() []byte {
	file_global_contract_proto_rawDescOnce.Do(func() {
		file_global_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_global_contract_proto_rawDescData)
	})
	return file_global_contract_proto_rawDescData
}

var file_global_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_global_contract_proto_goTypes = []interface{}{
	(*Contract)(nil),                // 0: global.Contract
	(*CreateContractPost)(nil),      // 1: global.CreateContractPost
	(*UpdateContractPost)(nil),      // 2: global.UpdateContractPost
	(*ContractResponse)(nil),        // 3: global.ContractResponse
	(*ContractAddressResponse)(nil), // 4: global.ContractAddressResponse
}
var file_global_contract_proto_depIdxs = []int32{
	0, // 0: global.ContractResponse.data:type_name -> global.Contract
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_global_contract_proto_init() }
func file_global_contract_proto_init() {
	if File_global_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_global_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contract); i {
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
		file_global_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContractPost); i {
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
		file_global_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateContractPost); i {
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
		file_global_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractResponse); i {
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
		file_global_contract_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractAddressResponse); i {
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
			RawDescriptor: file_global_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_global_contract_proto_goTypes,
		DependencyIndexes: file_global_contract_proto_depIdxs,
		MessageInfos:      file_global_contract_proto_msgTypes,
	}.Build()
	File_global_contract_proto = out.File
	file_global_contract_proto_rawDesc = nil
	file_global_contract_proto_goTypes = nil
	file_global_contract_proto_depIdxs = nil
}
