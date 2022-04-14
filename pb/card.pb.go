// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: member/card.proto

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

// Card 会员卡
type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Type       uint64 `protobuf:"varint,3,opt,name=type,proto3" json:"type"`
	Icon       string `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon"`
	Cover      string `protobuf:"bytes,5,opt,name=cover,proto3" json:"cover"`
	Desription string `protobuf:"bytes,6,opt,name=desription,proto3" json:"desription"`
	Num        uint64 `protobuf:"varint,7,opt,name=num,proto3" json:"num"`
	RemainNum  uint64 `protobuf:"varint,8,opt,name=remain_num,json=remainNum,proto3" json:"remain_num"`
	Price      uint64 `protobuf:"varint,9,opt,name=price,proto3" json:"price"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_member_card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_member_card_proto_rawDescGZIP(), []int{0}
}

func (x *Card) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Card) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Card) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Card) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Card) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *Card) GetDesription() string {
	if x != nil {
		return x.Desription
	}
	return ""
}

func (x *Card) GetNum() uint64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *Card) GetRemainNum() uint64 {
	if x != nil {
		return x.RemainNum
	}
	return 0
}

func (x *Card) GetPrice() uint64 {
	if x != nil {
		return x.Price
	}
	return 0
}

// CardResponse 返回会员卡
type CardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CardResponse) Reset() {
	*x = CardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_card_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardResponse) ProtoMessage() {}

func (x *CardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_member_card_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardResponse.ProtoReflect.Descriptor instead.
func (*CardResponse) Descriptor() ([]byte, []int) {
	return file_member_card_proto_rawDescGZIP(), []int{1}
}

var File_member_card_proto protoreflect.FileDescriptor

var file_member_card_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xcf, 0x01, 0x0a, 0x04,
	0x43, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x73, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x72, 0x65,
	0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x0e, 0x0a,
	0x0c, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_member_card_proto_rawDescOnce sync.Once
	file_member_card_proto_rawDescData = file_member_card_proto_rawDesc
)

func file_member_card_proto_rawDescGZIP() []byte {
	file_member_card_proto_rawDescOnce.Do(func() {
		file_member_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_member_card_proto_rawDescData)
	})
	return file_member_card_proto_rawDescData
}

var file_member_card_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_member_card_proto_goTypes = []interface{}{
	(*Card)(nil),         // 0: member.Card
	(*CardResponse)(nil), // 1: member.CardResponse
}
var file_member_card_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_member_card_proto_init() }
func file_member_card_proto_init() {
	if File_member_card_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_member_card_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
		file_member_card_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardResponse); i {
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
			RawDescriptor: file_member_card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_member_card_proto_goTypes,
		DependencyIndexes: file_member_card_proto_depIdxs,
		MessageInfos:      file_member_card_proto_msgTypes,
	}.Build()
	File_member_card_proto = out.File
	file_member_card_proto_rawDesc = nil
	file_member_card_proto_goTypes = nil
	file_member_card_proto_depIdxs = nil
}
