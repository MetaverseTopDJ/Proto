// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: nft/signature.proto

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

// Signature 签名结构
type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address   string `protobuf:"bytes,1,opt,name=address,proto3" json:"address"`
	Signature string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_signature_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_nft_signature_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_nft_signature_proto_rawDescGZIP(), []int{0}
}

func (x *Signature) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Signature) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

// BlindBoxSignPost 盲盒白名单签名
type BlindBoxSignPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractId uint64 `protobuf:"varint,1,opt,name=contract_id,json=contractId,proto3" json:"contract_id"`
	Address    string `protobuf:"bytes,2,opt,name=address,proto3" json:"address"`
	Num        uint64 `protobuf:"varint,3,opt,name=num,proto3" json:"num"`
	Price      uint64 `protobuf:"varint,4,opt,name=price,proto3" json:"price"`
	Code       uint64 `protobuf:"varint,5,opt,name=code,proto3" json:"code"`
	Timestamp  uint64 `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp"`
}

func (x *BlindBoxSignPost) Reset() {
	*x = BlindBoxSignPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_signature_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlindBoxSignPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlindBoxSignPost) ProtoMessage() {}

func (x *BlindBoxSignPost) ProtoReflect() protoreflect.Message {
	mi := &file_nft_signature_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlindBoxSignPost.ProtoReflect.Descriptor instead.
func (*BlindBoxSignPost) Descriptor() ([]byte, []int) {
	return file_nft_signature_proto_rawDescGZIP(), []int{1}
}

func (x *BlindBoxSignPost) GetContractId() uint64 {
	if x != nil {
		return x.ContractId
	}
	return 0
}

func (x *BlindBoxSignPost) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BlindBoxSignPost) GetNum() uint64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *BlindBoxSignPost) GetPrice() uint64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *BlindBoxSignPost) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BlindBoxSignPost) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// RemixSignPost 混合 签名
type RemixSignPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractId uint64 `protobuf:"varint,1,opt,name=contract_id,json=contractId,proto3" json:"contract_id"`
	Address    string `protobuf:"bytes,2,opt,name=address,proto3" json:"address"`
	Code       uint64 `protobuf:"varint,5,opt,name=code,proto3" json:"code"`
	Timestamp  uint64 `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp"`
}

func (x *RemixSignPost) Reset() {
	*x = RemixSignPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_signature_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemixSignPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemixSignPost) ProtoMessage() {}

func (x *RemixSignPost) ProtoReflect() protoreflect.Message {
	mi := &file_nft_signature_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemixSignPost.ProtoReflect.Descriptor instead.
func (*RemixSignPost) Descriptor() ([]byte, []int) {
	return file_nft_signature_proto_rawDescGZIP(), []int{2}
}

func (x *RemixSignPost) GetContractId() uint64 {
	if x != nil {
		return x.ContractId
	}
	return 0
}

func (x *RemixSignPost) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RemixSignPost) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RemixSignPost) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type CollectSignPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractId uint64 `protobuf:"varint,1,opt,name=contract_id,json=contractId,proto3" json:"contract_id"`
	Address    string `protobuf:"bytes,2,opt,name=address,proto3" json:"address"`
	Code       uint64 `protobuf:"varint,3,opt,name=code,proto3" json:"code"`
	Timestamp  uint64 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp"`
}

func (x *CollectSignPost) Reset() {
	*x = CollectSignPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_signature_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectSignPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectSignPost) ProtoMessage() {}

func (x *CollectSignPost) ProtoReflect() protoreflect.Message {
	mi := &file_nft_signature_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectSignPost.ProtoReflect.Descriptor instead.
func (*CollectSignPost) Descriptor() ([]byte, []int) {
	return file_nft_signature_proto_rawDescGZIP(), []int{3}
}

func (x *CollectSignPost) GetContractId() uint64 {
	if x != nil {
		return x.ContractId
	}
	return 0
}

func (x *CollectSignPost) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CollectSignPost) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CollectSignPost) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// SignatureResponse 返回 签名
type SignatureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string     `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    *Signature `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *SignatureResponse) Reset() {
	*x = SignatureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_signature_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignatureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureResponse) ProtoMessage() {}

func (x *SignatureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nft_signature_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignatureResponse.ProtoReflect.Descriptor instead.
func (*SignatureResponse) Descriptor() ([]byte, []int) {
	return file_nft_signature_proto_rawDescGZIP(), []int{4}
}

func (x *SignatureResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SignatureResponse) GetData() *Signature {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_nft_signature_proto protoreflect.FileDescriptor

var file_nft_signature_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6e, 0x66, 0x74, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x66, 0x74, 0x22, 0x43, 0x0a, 0x09, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22,
	0xa7, 0x01, 0x0a, 0x10, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x53, 0x69, 0x67, 0x6e,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6e, 0x75,
	0x6d, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x7c, 0x0a, 0x0d, 0x52, 0x65, 0x6d,
	0x69, 0x78, 0x53, 0x69, 0x67, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x7e, 0x0a, 0x0f, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x51, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nft_signature_proto_rawDescOnce sync.Once
	file_nft_signature_proto_rawDescData = file_nft_signature_proto_rawDesc
)

func file_nft_signature_proto_rawDescGZIP() []byte {
	file_nft_signature_proto_rawDescOnce.Do(func() {
		file_nft_signature_proto_rawDescData = protoimpl.X.CompressGZIP(file_nft_signature_proto_rawDescData)
	})
	return file_nft_signature_proto_rawDescData
}

var file_nft_signature_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_nft_signature_proto_goTypes = []interface{}{
	(*Signature)(nil),         // 0: nft.Signature
	(*BlindBoxSignPost)(nil),  // 1: nft.BlindBoxSignPost
	(*RemixSignPost)(nil),     // 2: nft.RemixSignPost
	(*CollectSignPost)(nil),   // 3: nft.CollectSignPost
	(*SignatureResponse)(nil), // 4: nft.SignatureResponse
}
var file_nft_signature_proto_depIdxs = []int32{
	0, // 0: nft.SignatureResponse.data:type_name -> nft.Signature
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nft_signature_proto_init() }
func file_nft_signature_proto_init() {
	if File_nft_signature_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nft_signature_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
		file_nft_signature_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlindBoxSignPost); i {
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
		file_nft_signature_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemixSignPost); i {
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
		file_nft_signature_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectSignPost); i {
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
		file_nft_signature_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignatureResponse); i {
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
			RawDescriptor: file_nft_signature_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nft_signature_proto_goTypes,
		DependencyIndexes: file_nft_signature_proto_depIdxs,
		MessageInfos:      file_nft_signature_proto_msgTypes,
	}.Build()
	File_nft_signature_proto = out.File
	file_nft_signature_proto_rawDesc = nil
	file_nft_signature_proto_goTypes = nil
	file_nft_signature_proto_depIdxs = nil
}
