// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: code.proto

package code

import (
	context "context"
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

type SendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Biz   string `protobuf:"bytes,2,opt,name=biz,proto3" json:"biz,omitempty"`
}

func (x *SendReq) Reset() {
	*x = SendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReq) ProtoMessage() {}

func (x *SendReq) ProtoReflect() protoreflect.Message {
	mi := &file_code_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReq.ProtoReflect.Descriptor instead.
func (*SendReq) Descriptor() ([]byte, []int) {
	return file_code_proto_rawDescGZIP(), []int{0}
}

func (x *SendReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SendReq) GetBiz() string {
	if x != nil {
		return x.Biz
	}
	return ""
}

type SendResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendResp) Reset() {
	*x = SendResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResp) ProtoMessage() {}

func (x *SendResp) ProtoReflect() protoreflect.Message {
	mi := &file_code_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResp.ProtoReflect.Descriptor instead.
func (*SendResp) Descriptor() ([]byte, []int) {
	return file_code_proto_rawDescGZIP(), []int{1}
}

type VerifyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Biz   string `protobuf:"bytes,2,opt,name=biz,proto3" json:"biz,omitempty"`
	Code  string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *VerifyReq) Reset() {
	*x = VerifyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyReq) ProtoMessage() {}

func (x *VerifyReq) ProtoReflect() protoreflect.Message {
	mi := &file_code_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyReq.ProtoReflect.Descriptor instead.
func (*VerifyReq) Descriptor() ([]byte, []int) {
	return file_code_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *VerifyReq) GetBiz() string {
	if x != nil {
		return x.Biz
	}
	return ""
}

func (x *VerifyReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type VerifyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VerifyResp) Reset() {
	*x = VerifyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyResp) ProtoMessage() {}

func (x *VerifyResp) ProtoReflect() protoreflect.Message {
	mi := &file_code_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyResp.ProtoReflect.Descriptor instead.
func (*VerifyResp) Descriptor() ([]byte, []int) {
	return file_code_proto_rawDescGZIP(), []int{3}
}

var File_code_proto protoreflect.FileDescriptor

var file_code_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x31, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x7a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x62, 0x69, 0x7a, 0x22, 0x0a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x47, 0x0a, 0x09, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x7a, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x62, 0x69, 0x7a, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x32, 0x65, 0x0a, 0x0b, 0x43, 0x6f, 0x64, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12,
	0x0d, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0e,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x2d, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42,
	0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x58,
	0x4a, 0x30, 0x30, 0x30, 0x30, 0x2f, 0x67, 0x6f, 0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f,
	0x63, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_code_proto_rawDescOnce sync.Once
	file_code_proto_rawDescData = file_code_proto_rawDesc
)

func file_code_proto_rawDescGZIP() []byte {
	file_code_proto_rawDescOnce.Do(func() {
		file_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_code_proto_rawDescData)
	})
	return file_code_proto_rawDescData
}

var file_code_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_code_proto_goTypes = []interface{}{
	(*SendReq)(nil),    // 0: code.SendReq
	(*SendResp)(nil),   // 1: code.SendResp
	(*VerifyReq)(nil),  // 2: code.VerifyReq
	(*VerifyResp)(nil), // 3: code.VerifyResp
}
var file_code_proto_depIdxs = []int32{
	0, // 0: code.CodeService.Send:input_type -> code.SendReq
	2, // 1: code.CodeService.Verify:input_type -> code.VerifyReq
	1, // 2: code.CodeService.Send:output_type -> code.SendResp
	3, // 3: code.CodeService.Verify:output_type -> code.VerifyResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_code_proto_init() }
func file_code_proto_init() {
	if File_code_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_code_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendReq); i {
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
		file_code_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResp); i {
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
		file_code_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyReq); i {
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
		file_code_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyResp); i {
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
			RawDescriptor: file_code_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_code_proto_goTypes,
		DependencyIndexes: file_code_proto_depIdxs,
		MessageInfos:      file_code_proto_msgTypes,
	}.Build()
	File_code_proto = out.File
	file_code_proto_rawDesc = nil
	file_code_proto_goTypes = nil
	file_code_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.1. DO NOT EDIT.

type CodeService interface {
	Send(ctx context.Context, req *SendReq) (res *SendResp, err error)
	Verify(ctx context.Context, req *VerifyReq) (res *VerifyResp, err error)
}
