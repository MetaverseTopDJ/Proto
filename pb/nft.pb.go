// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: nft/nft.proto

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

var File_nft_nft_proto protoreflect.FileDescriptor

var file_nft_nft_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x66, 0x74, 0x2f, 0x6e, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x6e, 0x66, 0x74, 0x1a, 0x13, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x6e, 0x66, 0x74, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0e, 0x6e, 0x66, 0x74, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x6e, 0x66, 0x74, 0x2f, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x5f, 0x62, 0x6f,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc6, 0x0c, 0x0a, 0x0a, 0x4e, 0x46, 0x54, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x10, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x6f,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0c, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x2e, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6e, 0x66,
	0x74, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x48, 0x6f, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x73, 0x12, 0x11, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x48, 0x6f,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0f, 0x43, 0x61, 0x63, 0x68, 0x65, 0x41, 0x6c, 0x6c, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x11, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x48,
	0x6f, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x15, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x15,
	0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x12,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6e,
	0x66, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x04, 0x57, 0x6f, 0x72, 0x6b, 0x12, 0x10, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x6e, 0x66, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x48, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x12, 0x17, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x6e, 0x66, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0e,
	0x57, 0x6f, 0x72, 0x6b, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x12, 0x13, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x12, 0x13, 0x2e, 0x6e, 0x66,
	0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x50, 0x6f, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x2e, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50,
	0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x08, 0x42, 0x6c, 0x69, 0x6e,
	0x64, 0x42, 0x6f, 0x78, 0x12, 0x10, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x42, 0x6c, 0x69,
	0x6e, 0x64, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4f, 0x0a, 0x12, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x6e, 0x66, 0x74, 0x2e, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42,
	0x6f, 0x78, 0x12, 0x17, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6e, 0x66,
	0x74, 0x2e, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c,
	0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x12, 0x17, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x50, 0x6f, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x12, 0x18, 0x2e, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50,
	0x6f, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x42,
	0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0b,
	0x4e, 0x46, 0x54, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x10, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x12, 0x4e, 0x46, 0x54,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x11, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x6f,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e,
	0x46, 0x54, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4b, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x46, 0x54, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4f, 0x0a, 0x17, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x46, 0x54, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x2e, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x50, 0x6f, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_nft_nft_proto_goTypes = []interface{}{
	(*InfoPost)(nil),                   // 0: global.InfoPost
	(*EmptyPost)(nil),                  // 1: global.EmptyPost
	(*PaginationPost)(nil),             // 2: global.PaginationPost
	(*CreateAuthorPost)(nil),           // 3: nft.CreateAuthorPost
	(*UpdateAuthorPost)(nil),           // 4: nft.UpdateAuthorPost
	(*ChangeStatusPost)(nil),           // 5: global.ChangeStatusPost
	(*RecommendWorksPost)(nil),         // 6: nft.RecommendWorksPost
	(*CreateWorkPost)(nil),             // 7: nft.CreateWorkPost
	(*UpdateWorkPost)(nil),             // 8: nft.UpdateWorkPost
	(*CreateBlindBoxPost)(nil),         // 9: nft.CreateBlindBoxPost
	(*UpdateBlindBoxPost)(nil),         // 10: nft.UpdateBlindBoxPost
	(*CreateContractPost)(nil),         // 11: global.CreateContractPost
	(*UpdateContractPost)(nil),         // 12: global.UpdateContractPost
	(*AuthorResponse)(nil),             // 13: nft.AuthorResponse
	(*SimpleAuthorResponse)(nil),       // 14: nft.SimpleAuthorResponse
	(*HotAuthorsResponse)(nil),         // 15: nft.HotAuthorsResponse
	(*AuthorPaginationResponse)(nil),   // 16: nft.AuthorPaginationResponse
	(*WorkResponse)(nil),               // 17: nft.WorkResponse
	(*RecommendWorksResponse)(nil),     // 18: nft.RecommendWorksResponse
	(*WorkPaginationResponse)(nil),     // 19: nft.WorkPaginationResponse
	(*BlindBoxResponse)(nil),           // 20: nft.BlindBoxResponse
	(*BlindBoxPaginationResponse)(nil), // 21: nft.BlindBoxPaginationResponse
	(*ContractResponse)(nil),           // 22: global.ContractResponse
	(*ContractAddressResponse)(nil),    // 23: global.ContractAddressResponse
}
var file_nft_nft_proto_depIdxs = []int32{
	0,  // 0: nft.NFTService.Author:input_type -> global.InfoPost
	0,  // 1: nft.NFTService.AuthorDetail:input_type -> global.InfoPost
	1,  // 2: nft.NFTService.HotAuthors:input_type -> global.EmptyPost
	1,  // 3: nft.NFTService.CacheAllAuthors:input_type -> global.EmptyPost
	2,  // 4: nft.NFTService.AuthorPagination:input_type -> global.PaginationPost
	3,  // 5: nft.NFTService.CreateAuthor:input_type -> nft.CreateAuthorPost
	4,  // 6: nft.NFTService.UpdateAuthor:input_type -> nft.UpdateAuthorPost
	5,  // 7: nft.NFTService.ChangeAuthorStatus:input_type -> global.ChangeStatusPost
	0,  // 8: nft.NFTService.Work:input_type -> global.InfoPost
	6,  // 9: nft.NFTService.RecommendWorks:input_type -> nft.RecommendWorksPost
	2,  // 10: nft.NFTService.WorkPagination:input_type -> global.PaginationPost
	7,  // 11: nft.NFTService.CreateWork:input_type -> nft.CreateWorkPost
	8,  // 12: nft.NFTService.UpdateWork:input_type -> nft.UpdateWorkPost
	5,  // 13: nft.NFTService.ChangeWorkStatus:input_type -> global.ChangeStatusPost
	0,  // 14: nft.NFTService.BlindBox:input_type -> global.InfoPost
	2,  // 15: nft.NFTService.BlindBoxPagination:input_type -> global.PaginationPost
	9,  // 16: nft.NFTService.CreateBlindBox:input_type -> nft.CreateBlindBoxPost
	10, // 17: nft.NFTService.UpdateBlindBox:input_type -> nft.UpdateBlindBoxPost
	5,  // 18: nft.NFTService.StartBlindBox:input_type -> global.ChangeStatusPost
	0,  // 19: nft.NFTService.NFTContract:input_type -> global.InfoPost
	1,  // 20: nft.NFTService.NFTContractAddress:input_type -> global.EmptyPost
	11, // 21: nft.NFTService.CreateNFTContract:input_type -> global.CreateContractPost
	12, // 22: nft.NFTService.UpdateNFTContract:input_type -> global.UpdateContractPost
	5,  // 23: nft.NFTService.ChangeNFTContractStatus:input_type -> global.ChangeStatusPost
	13, // 24: nft.NFTService.Author:output_type -> nft.AuthorResponse
	14, // 25: nft.NFTService.AuthorDetail:output_type -> nft.SimpleAuthorResponse
	15, // 26: nft.NFTService.HotAuthors:output_type -> nft.HotAuthorsResponse
	15, // 27: nft.NFTService.CacheAllAuthors:output_type -> nft.HotAuthorsResponse
	16, // 28: nft.NFTService.AuthorPagination:output_type -> nft.AuthorPaginationResponse
	13, // 29: nft.NFTService.CreateAuthor:output_type -> nft.AuthorResponse
	13, // 30: nft.NFTService.UpdateAuthor:output_type -> nft.AuthorResponse
	13, // 31: nft.NFTService.ChangeAuthorStatus:output_type -> nft.AuthorResponse
	17, // 32: nft.NFTService.Work:output_type -> nft.WorkResponse
	18, // 33: nft.NFTService.RecommendWorks:output_type -> nft.RecommendWorksResponse
	19, // 34: nft.NFTService.WorkPagination:output_type -> nft.WorkPaginationResponse
	17, // 35: nft.NFTService.CreateWork:output_type -> nft.WorkResponse
	17, // 36: nft.NFTService.UpdateWork:output_type -> nft.WorkResponse
	17, // 37: nft.NFTService.ChangeWorkStatus:output_type -> nft.WorkResponse
	20, // 38: nft.NFTService.BlindBox:output_type -> nft.BlindBoxResponse
	21, // 39: nft.NFTService.BlindBoxPagination:output_type -> nft.BlindBoxPaginationResponse
	20, // 40: nft.NFTService.CreateBlindBox:output_type -> nft.BlindBoxResponse
	20, // 41: nft.NFTService.UpdateBlindBox:output_type -> nft.BlindBoxResponse
	20, // 42: nft.NFTService.StartBlindBox:output_type -> nft.BlindBoxResponse
	22, // 43: nft.NFTService.NFTContract:output_type -> global.ContractResponse
	23, // 44: nft.NFTService.NFTContractAddress:output_type -> global.ContractAddressResponse
	22, // 45: nft.NFTService.CreateNFTContract:output_type -> global.ContractResponse
	22, // 46: nft.NFTService.UpdateNFTContract:output_type -> global.ContractResponse
	22, // 47: nft.NFTService.ChangeNFTContractStatus:output_type -> global.ContractResponse
	24, // [24:48] is the sub-list for method output_type
	0,  // [0:24] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_nft_nft_proto_init() }
func file_nft_nft_proto_init() {
	if File_nft_nft_proto != nil {
		return
	}
	file_global_global_proto_init()
	file_global_contract_proto_init()
	file_nft_author_proto_init()
	file_nft_work_proto_init()
	file_nft_blind_box_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nft_nft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nft_nft_proto_goTypes,
		DependencyIndexes: file_nft_nft_proto_depIdxs,
	}.Build()
	File_nft_nft_proto = out.File
	file_nft_nft_proto_rawDesc = nil
	file_nft_nft_proto_goTypes = nil
	file_nft_nft_proto_depIdxs = nil
}
