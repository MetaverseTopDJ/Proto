// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: member/member_contract_event_log.proto

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

// MemberContractEventLog 会员合约日志
type MemberContractEventLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	ContractId  uint64 `protobuf:"varint,2,opt,name=contract_id,json=contractId,proto3" json:"contract_id"`
	Hash        string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash"`
	BlockNumber uint64 `protobuf:"varint,4,opt,name=block_number,json=blockNumber,proto3" json:"block_number"`
	BlockAt     string `protobuf:"bytes,5,opt,name=block_at,json=blockAt,proto3" json:"block_at"`
	Retry       uint64 `protobuf:"varint,6,opt,name=retry,proto3" json:"retry"`
	Removed     bool   `protobuf:"varint,7,opt,name=removed,proto3" json:"removed"`
	Error       string `protobuf:"bytes,8,opt,name=error,proto3" json:"error"`
	Status      uint64 `protobuf:"varint,9,opt,name=status,proto3" json:"status"`
	CreatedAt   string `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt   string `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *MemberContractEventLog) Reset() {
	*x = MemberContractEventLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_contract_event_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberContractEventLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberContractEventLog) ProtoMessage() {}

func (x *MemberContractEventLog) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_contract_event_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberContractEventLog.ProtoReflect.Descriptor instead.
func (*MemberContractEventLog) Descriptor() ([]byte, []int) {
	return file_member_member_contract_event_log_proto_rawDescGZIP(), []int{0}
}

func (x *MemberContractEventLog) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberContractEventLog) GetContractId() uint64 {
	if x != nil {
		return x.ContractId
	}
	return 0
}

func (x *MemberContractEventLog) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *MemberContractEventLog) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *MemberContractEventLog) GetBlockAt() string {
	if x != nil {
		return x.BlockAt
	}
	return ""
}

func (x *MemberContractEventLog) GetRetry() uint64 {
	if x != nil {
		return x.Retry
	}
	return 0
}

func (x *MemberContractEventLog) GetRemoved() bool {
	if x != nil {
		return x.Removed
	}
	return false
}

func (x *MemberContractEventLog) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *MemberContractEventLog) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MemberContractEventLog) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *MemberContractEventLog) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// MemberContractEventLogResponse 返回
type MemberContractEventLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                  `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    *MemberContractEventLog `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *MemberContractEventLogResponse) Reset() {
	*x = MemberContractEventLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_contract_event_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberContractEventLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberContractEventLogResponse) ProtoMessage() {}

func (x *MemberContractEventLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_contract_event_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberContractEventLogResponse.ProtoReflect.Descriptor instead.
func (*MemberContractEventLogResponse) Descriptor() ([]byte, []int) {
	return file_member_member_contract_event_log_proto_rawDescGZIP(), []int{1}
}

func (x *MemberContractEventLogResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MemberContractEventLogResponse) GetData() *MemberContractEventLog {
	if x != nil {
		return x.Data
	}
	return nil
}

// MemberContractEventLogPaginationResponse 分页
type MemberContractEventLogPaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                                         `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Data    *MemberContractEventLogPaginationResponse_Logs `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *MemberContractEventLogPaginationResponse) Reset() {
	*x = MemberContractEventLogPaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_contract_event_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberContractEventLogPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberContractEventLogPaginationResponse) ProtoMessage() {}

func (x *MemberContractEventLogPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_contract_event_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberContractEventLogPaginationResponse.ProtoReflect.Descriptor instead.
func (*MemberContractEventLogPaginationResponse) Descriptor() ([]byte, []int) {
	return file_member_member_contract_event_log_proto_rawDescGZIP(), []int{2}
}

func (x *MemberContractEventLogPaginationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MemberContractEventLogPaginationResponse) GetData() *MemberContractEventLogPaginationResponse_Logs {
	if x != nil {
		return x.Data
	}
	return nil
}

type MemberContractEventLogPaginationResponse_Logs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*MemberContractEventLog `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
	Total int64                     `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
}

func (x *MemberContractEventLogPaginationResponse_Logs) Reset() {
	*x = MemberContractEventLogPaginationResponse_Logs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_member_contract_event_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberContractEventLogPaginationResponse_Logs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberContractEventLogPaginationResponse_Logs) ProtoMessage() {}

func (x *MemberContractEventLogPaginationResponse_Logs) ProtoReflect() protoreflect.Message {
	mi := &file_member_member_contract_event_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberContractEventLogPaginationResponse_Logs.ProtoReflect.Descriptor instead.
func (*MemberContractEventLogPaginationResponse_Logs) Descriptor() ([]byte, []int) {
	return file_member_member_contract_event_log_proto_rawDescGZIP(), []int{2, 0}
}

func (x *MemberContractEventLogPaginationResponse_Logs) GetList() []*MemberContractEventLog {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *MemberContractEventLogPaginationResponse_Logs) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_member_member_contract_event_log_proto protoreflect.FileDescriptor

var file_member_member_contract_event_log_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0xb7, 0x02, 0x0a, 0x16, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x65, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x72,
	0x65, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x6e, 0x0a, 0x1e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xe1, 0x01, 0x0a, 0x28, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x49, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x50, 0x0a, 0x04,
	0x4c, 0x6f, 0x67, 0x73, 0x12, 0x32, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c,
	0x6f, 0x67, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_member_member_contract_event_log_proto_rawDescOnce sync.Once
	file_member_member_contract_event_log_proto_rawDescData = file_member_member_contract_event_log_proto_rawDesc
)

func file_member_member_contract_event_log_proto_rawDescGZIP() []byte {
	file_member_member_contract_event_log_proto_rawDescOnce.Do(func() {
		file_member_member_contract_event_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_member_member_contract_event_log_proto_rawDescData)
	})
	return file_member_member_contract_event_log_proto_rawDescData
}

var file_member_member_contract_event_log_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_member_member_contract_event_log_proto_goTypes = []interface{}{
	(*MemberContractEventLog)(nil),                        // 0: member.MemberContractEventLog
	(*MemberContractEventLogResponse)(nil),                // 1: member.MemberContractEventLogResponse
	(*MemberContractEventLogPaginationResponse)(nil),      // 2: member.MemberContractEventLogPaginationResponse
	(*MemberContractEventLogPaginationResponse_Logs)(nil), // 3: member.MemberContractEventLogPaginationResponse.Logs
}
var file_member_member_contract_event_log_proto_depIdxs = []int32{
	0, // 0: member.MemberContractEventLogResponse.data:type_name -> member.MemberContractEventLog
	3, // 1: member.MemberContractEventLogPaginationResponse.data:type_name -> member.MemberContractEventLogPaginationResponse.Logs
	0, // 2: member.MemberContractEventLogPaginationResponse.Logs.list:type_name -> member.MemberContractEventLog
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_member_member_contract_event_log_proto_init() }
func file_member_member_contract_event_log_proto_init() {
	if File_member_member_contract_event_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_member_member_contract_event_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberContractEventLog); i {
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
		file_member_member_contract_event_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberContractEventLogResponse); i {
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
		file_member_member_contract_event_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberContractEventLogPaginationResponse); i {
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
		file_member_member_contract_event_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberContractEventLogPaginationResponse_Logs); i {
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
			RawDescriptor: file_member_member_contract_event_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_member_member_contract_event_log_proto_goTypes,
		DependencyIndexes: file_member_member_contract_event_log_proto_depIdxs,
		MessageInfos:      file_member_member_contract_event_log_proto_msgTypes,
	}.Build()
	File_member_member_contract_event_log_proto = out.File
	file_member_member_contract_event_log_proto_rawDesc = nil
	file_member_member_contract_event_log_proto_goTypes = nil
	file_member_member_contract_event_log_proto_depIdxs = nil
}
