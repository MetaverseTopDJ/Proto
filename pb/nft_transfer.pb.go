// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: nft/nft_transfer.proto

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

// NFTTransfer NFT 交易信息
type NFTTransfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64            `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	BlockNumber uint64            `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number"`
	BlockTime   string            `protobuf:"bytes,3,opt,name=block_time,json=blockTime,proto3" json:"block_time"`
	ContractId  uint64            `protobuf:"varint,4,opt,name=contract_id,json=contractId,proto3" json:"contract_id"`
	TokenId     uint64            `protobuf:"varint,5,opt,name=token_id,json=tokenId,proto3" json:"token_id"`
	Hash        string            `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash"`
	FormId      uint64            `protobuf:"varint,7,opt,name=form_id,json=formId,proto3" json:"form_id"`
	Form        string            `protobuf:"bytes,8,opt,name=form,proto3" json:"form"`
	ToId        uint64            `protobuf:"varint,9,opt,name=to_id,json=toId,proto3" json:"to_id"`
	To          string            `protobuf:"bytes,10,opt,name=to,proto3" json:"to"`
	Contract    *RelationContract `protobuf:"bytes,11,opt,name=contract,proto3" json:"contract"`
}

func (x *NFTTransfer) Reset() {
	*x = NFTTransfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_nft_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NFTTransfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NFTTransfer) ProtoMessage() {}

func (x *NFTTransfer) ProtoReflect() protoreflect.Message {
	mi := &file_nft_nft_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NFTTransfer.ProtoReflect.Descriptor instead.
func (*NFTTransfer) Descriptor() ([]byte, []int) {
	return file_nft_nft_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *NFTTransfer) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NFTTransfer) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *NFTTransfer) GetBlockTime() string {
	if x != nil {
		return x.BlockTime
	}
	return ""
}

func (x *NFTTransfer) GetContractId() uint64 {
	if x != nil {
		return x.ContractId
	}
	return 0
}

func (x *NFTTransfer) GetTokenId() uint64 {
	if x != nil {
		return x.TokenId
	}
	return 0
}

func (x *NFTTransfer) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *NFTTransfer) GetFormId() uint64 {
	if x != nil {
		return x.FormId
	}
	return 0
}

func (x *NFTTransfer) GetForm() string {
	if x != nil {
		return x.Form
	}
	return ""
}

func (x *NFTTransfer) GetToId() uint64 {
	if x != nil {
		return x.ToId
	}
	return 0
}

func (x *NFTTransfer) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *NFTTransfer) GetContract() *RelationContract {
	if x != nil {
		return x.Contract
	}
	return nil
}

// NftTransferPaginationPost NFT 交易日志 分页请求
type NftTransferPaginationPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractId uint64 `protobuf:"varint,1,opt,name=contract_id,json=contractId,proto3" json:"contract_id"`
	From       string `protobuf:"bytes,2,opt,name=from,proto3" json:"from"`
	To         string `protobuf:"bytes,3,opt,name=to,proto3" json:"to"`
	Page       uint64 `protobuf:"varint,4,opt,name=page,proto3" json:"page"`
	Size       uint64 `protobuf:"varint,5,opt,name=size,proto3" json:"size"`
}

func (x *NftTransferPaginationPost) Reset() {
	*x = NftTransferPaginationPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_nft_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftTransferPaginationPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftTransferPaginationPost) ProtoMessage() {}

func (x *NftTransferPaginationPost) ProtoReflect() protoreflect.Message {
	mi := &file_nft_nft_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftTransferPaginationPost.ProtoReflect.Descriptor instead.
func (*NftTransferPaginationPost) Descriptor() ([]byte, []int) {
	return file_nft_nft_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *NftTransferPaginationPost) GetContractId() uint64 {
	if x != nil {
		return x.ContractId
	}
	return 0
}

func (x *NftTransferPaginationPost) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *NftTransferPaginationPost) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *NftTransferPaginationPost) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *NftTransferPaginationPost) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

// NFTTransferPaginationResponse 交易日志 分页
type NFTTransferPaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                              `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    *NFTTransferPaginationResponse_NFTs `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *NFTTransferPaginationResponse) Reset() {
	*x = NFTTransferPaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_nft_transfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NFTTransferPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NFTTransferPaginationResponse) ProtoMessage() {}

func (x *NFTTransferPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nft_nft_transfer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NFTTransferPaginationResponse.ProtoReflect.Descriptor instead.
func (*NFTTransferPaginationResponse) Descriptor() ([]byte, []int) {
	return file_nft_nft_transfer_proto_rawDescGZIP(), []int{2}
}

func (x *NFTTransferPaginationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *NFTTransferPaginationResponse) GetData() *NFTTransferPaginationResponse_NFTs {
	if x != nil {
		return x.Data
	}
	return nil
}

type NFTTransferPaginationResponse_NFTs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*NFTTransfer `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
	Total int64          `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
}

func (x *NFTTransferPaginationResponse_NFTs) Reset() {
	*x = NFTTransferPaginationResponse_NFTs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_nft_transfer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NFTTransferPaginationResponse_NFTs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NFTTransferPaginationResponse_NFTs) ProtoMessage() {}

func (x *NFTTransferPaginationResponse_NFTs) ProtoReflect() protoreflect.Message {
	mi := &file_nft_nft_transfer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NFTTransferPaginationResponse_NFTs.ProtoReflect.Descriptor instead.
func (*NFTTransferPaginationResponse_NFTs) Descriptor() ([]byte, []int) {
	return file_nft_nft_transfer_proto_rawDescGZIP(), []int{2, 0}
}

func (x *NFTTransferPaginationResponse_NFTs) GetList() []*NFTTransfer {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *NFTTransferPaginationResponse_NFTs) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_nft_nft_transfer_proto protoreflect.FileDescriptor

var file_nft_nft_transfer_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6e, 0x66, 0x74, 0x2f, 0x6e, 0x66, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x66, 0x74, 0x1a, 0x15, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x02, 0x0a, 0x0b, 0x4e, 0x46, 0x54, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x6f, 0x72, 0x6d, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x34, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x22, 0x88,
	0x01, 0x0a, 0x19, 0x4e, 0x66, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0xba, 0x01, 0x0a, 0x1d, 0x4e, 0x46,
	0x54, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x4e, 0x46, 0x54, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4e, 0x46, 0x54, 0x73, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x42, 0x0a, 0x04, 0x4e, 0x46, 0x54, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x4e,
	0x46, 0x54, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nft_nft_transfer_proto_rawDescOnce sync.Once
	file_nft_nft_transfer_proto_rawDescData = file_nft_nft_transfer_proto_rawDesc
)

func file_nft_nft_transfer_proto_rawDescGZIP() []byte {
	file_nft_nft_transfer_proto_rawDescOnce.Do(func() {
		file_nft_nft_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_nft_nft_transfer_proto_rawDescData)
	})
	return file_nft_nft_transfer_proto_rawDescData
}

var file_nft_nft_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nft_nft_transfer_proto_goTypes = []interface{}{
	(*NFTTransfer)(nil),                        // 0: nft.NFTTransfer
	(*NftTransferPaginationPost)(nil),          // 1: nft.NftTransferPaginationPost
	(*NFTTransferPaginationResponse)(nil),      // 2: nft.NFTTransferPaginationResponse
	(*NFTTransferPaginationResponse_NFTs)(nil), // 3: nft.NFTTransferPaginationResponse.NFTs
	(*RelationContract)(nil),                   // 4: global.RelationContract
}
var file_nft_nft_transfer_proto_depIdxs = []int32{
	4, // 0: nft.NFTTransfer.contract:type_name -> global.RelationContract
	3, // 1: nft.NFTTransferPaginationResponse.data:type_name -> nft.NFTTransferPaginationResponse.NFTs
	0, // 2: nft.NFTTransferPaginationResponse.NFTs.list:type_name -> nft.NFTTransfer
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_nft_nft_transfer_proto_init() }
func file_nft_nft_transfer_proto_init() {
	if File_nft_nft_transfer_proto != nil {
		return
	}
	file_global_contract_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_nft_nft_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NFTTransfer); i {
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
		file_nft_nft_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftTransferPaginationPost); i {
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
		file_nft_nft_transfer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NFTTransferPaginationResponse); i {
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
		file_nft_nft_transfer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NFTTransferPaginationResponse_NFTs); i {
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
			RawDescriptor: file_nft_nft_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nft_nft_transfer_proto_goTypes,
		DependencyIndexes: file_nft_nft_transfer_proto_depIdxs,
		MessageInfos:      file_nft_nft_transfer_proto_msgTypes,
	}.Build()
	File_nft_nft_transfer_proto = out.File
	file_nft_nft_transfer_proto_rawDesc = nil
	file_nft_nft_transfer_proto_goTypes = nil
	file_nft_nft_transfer_proto_depIdxs = nil
}
