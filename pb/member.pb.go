// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: member/member.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_member_member_proto protoreflect.FileDescriptor

var file_member_member_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x13, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2f,
	0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x65, 0x5f, 0x73, 0x61, 0x6c, 0x65,
	0x5f, 0x77, 0x68, 0x69, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8a, 0x09, 0x0a, 0x0d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x12, 0x10, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x6f,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x07,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x73, 0x12, 0x11, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x10, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x18, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x18, 0x2e, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50,
	0x6f, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a,
	0x12, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12,
	0x10, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x6f, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x10, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5e, 0x0a, 0x18, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x20, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x12, 0x20, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x15, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x12, 0x10, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x6f,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x10, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x62, 0x0a, 0x1d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x65, 0x53,
	0x61, 0x6c, 0x65, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x57, 0x68, 0x69, 0x74,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x57, 0x68, 0x69,
	0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2b, 0x2e, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50,
	0x72, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x57, 0x68, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_member_member_proto_goTypes = []interface{}{
	(*InfoPost)(nil),                              // 0: global.InfoPost
	(*EmptyPost)(nil),                             // 1: global.EmptyPost
	(*PaginationPost)(nil),                        // 2: global.PaginationPost
	(*CreateAvatarPost)(nil),                      // 3: member.CreateAvatarPost
	(*UpdateAvatarPost)(nil),                      // 4: member.UpdateAvatarPost
	(*ChangeStatusPost)(nil),                      // 5: global.ChangeStatusPost
	(*CreateMemberActivityPost)(nil),              // 6: member.CreateMemberActivityPost
	(*UpdateMemberActivityPost)(nil),              // 7: member.UpdateMemberActivityPost
	(*CreateMemberPreSaleWhiteAddressPost)(nil),   // 8: member.CreateMemberPreSaleWhiteAddressPost
	(*AvatarResponse)(nil),                        // 9: member.AvatarResponse
	(*AvatarsResponse)(nil),                       // 10: member.AvatarsResponse
	(*AvatarPaginationResponse)(nil),              // 11: member.AvatarPaginationResponse
	(*CardResponse)(nil),                          // 12: member.CardResponse
	(*MemberActivityResponse)(nil),                // 13: member.MemberActivityResponse
	(*MemberActivityPaginationResponse)(nil),      // 14: member.MemberActivityPaginationResponse
	(*MessageResponse)(nil),                       // 15: global.MessageResponse
	(*MemberPreSaleWhiteAddressListResponse)(nil), // 16: member.MemberPreSaleWhiteAddressListResponse
}
var file_member_member_proto_depIdxs = []int32{
	0,  // 0: member.MemberService.Avatar:input_type -> global.InfoPost
	1,  // 1: member.MemberService.Avatars:input_type -> global.EmptyPost
	2,  // 2: member.MemberService.AvatarPagination:input_type -> global.PaginationPost
	3,  // 3: member.MemberService.CreateAvatar:input_type -> member.CreateAvatarPost
	4,  // 4: member.MemberService.UpdateAvatar:input_type -> member.UpdateAvatarPost
	5,  // 5: member.MemberService.ChangeAvatarStatus:input_type -> global.ChangeStatusPost
	0,  // 6: member.MemberService.Card:input_type -> global.InfoPost
	0,  // 7: member.MemberService.MemberActivity:input_type -> global.InfoPost
	2,  // 8: member.MemberService.MemberActivityPagination:input_type -> global.PaginationPost
	6,  // 9: member.MemberService.CreateMemberActivity:input_type -> member.CreateMemberActivityPost
	7,  // 10: member.MemberService.UpdateMemberActivity:input_type -> member.UpdateMemberActivityPost
	0,  // 11: member.MemberService.RefreshMemberActivity:input_type -> global.InfoPost
	0,  // 12: member.MemberService.DeleteMemberActivity:input_type -> global.InfoPost
	0,  // 13: member.MemberService.MemberPreSaleWhiteAddressList:input_type -> global.InfoPost
	8,  // 14: member.MemberService.CreateMemberPreSaleWhiteAddress:input_type -> member.CreateMemberPreSaleWhiteAddressPost
	9,  // 15: member.MemberService.Avatar:output_type -> member.AvatarResponse
	10, // 16: member.MemberService.Avatars:output_type -> member.AvatarsResponse
	11, // 17: member.MemberService.AvatarPagination:output_type -> member.AvatarPaginationResponse
	9,  // 18: member.MemberService.CreateAvatar:output_type -> member.AvatarResponse
	9,  // 19: member.MemberService.UpdateAvatar:output_type -> member.AvatarResponse
	9,  // 20: member.MemberService.ChangeAvatarStatus:output_type -> member.AvatarResponse
	12, // 21: member.MemberService.Card:output_type -> member.CardResponse
	13, // 22: member.MemberService.MemberActivity:output_type -> member.MemberActivityResponse
	14, // 23: member.MemberService.MemberActivityPagination:output_type -> member.MemberActivityPaginationResponse
	13, // 24: member.MemberService.CreateMemberActivity:output_type -> member.MemberActivityResponse
	13, // 25: member.MemberService.UpdateMemberActivity:output_type -> member.MemberActivityResponse
	15, // 26: member.MemberService.RefreshMemberActivity:output_type -> global.MessageResponse
	15, // 27: member.MemberService.DeleteMemberActivity:output_type -> global.MessageResponse
	16, // 28: member.MemberService.MemberPreSaleWhiteAddressList:output_type -> member.MemberPreSaleWhiteAddressListResponse
	15, // 29: member.MemberService.CreateMemberPreSaleWhiteAddress:output_type -> global.MessageResponse
	15, // [15:30] is the sub-list for method output_type
	0,  // [0:15] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_member_member_proto_init() }
func file_member_member_proto_init() {
	if File_member_member_proto != nil {
		return
	}
	file_global_global_proto_init()
	file_member_avatar_proto_init()
	file_member_card_proto_init()
	file_member_member_activity_proto_init()
	file_member_member_pre_sale_white_address_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_member_member_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_member_member_proto_goTypes,
		DependencyIndexes: file_member_member_proto_depIdxs,
	}.Build()
	File_member_member_proto = out.File
	file_member_member_proto_rawDesc = nil
	file_member_member_proto_goTypes = nil
	file_member_member_proto_depIdxs = nil
}
