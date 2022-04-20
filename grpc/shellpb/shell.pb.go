// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: grpc/shellpb/shell.proto

package shellpb

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

type ExecuteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Script string `protobuf:"bytes,1,opt,name=script,proto3" json:"script,omitempty"`
}

func (x *ExecuteRequest) Reset() {
	*x = ExecuteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_shellpb_shell_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteRequest) ProtoMessage() {}

func (x *ExecuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_shellpb_shell_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteRequest.ProtoReflect.Descriptor instead.
func (*ExecuteRequest) Descriptor() ([]byte, []int) {
	return file_grpc_shellpb_shell_proto_rawDescGZIP(), []int{0}
}

func (x *ExecuteRequest) GetScript() string {
	if x != nil {
		return x.Script
	}
	return ""
}

type ExecuteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stdout string `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr string `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	Code   int32  `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *ExecuteResponse) Reset() {
	*x = ExecuteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_shellpb_shell_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteResponse) ProtoMessage() {}

func (x *ExecuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_shellpb_shell_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteResponse.ProtoReflect.Descriptor instead.
func (*ExecuteResponse) Descriptor() ([]byte, []int) {
	return file_grpc_shellpb_shell_proto_rawDescGZIP(), []int{1}
}

func (x *ExecuteResponse) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *ExecuteResponse) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

func (x *ExecuteResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_grpc_shellpb_shell_proto protoreflect.FileDescriptor

var file_grpc_shellpb_shell_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x70, 0x62, 0x2f, 0x73,
	0x68, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x68, 0x65, 0x6c,
	0x6c, 0x22, 0x28, 0x0a, 0x0e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x55, 0x0a, 0x0f, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x4a, 0x0a, 0x0c, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x15, 0x2e,
	0x73, 0x68, 0x65, 0x6c, 0x6c, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_shellpb_shell_proto_rawDescOnce sync.Once
	file_grpc_shellpb_shell_proto_rawDescData = file_grpc_shellpb_shell_proto_rawDesc
)

func file_grpc_shellpb_shell_proto_rawDescGZIP() []byte {
	file_grpc_shellpb_shell_proto_rawDescOnce.Do(func() {
		file_grpc_shellpb_shell_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_shellpb_shell_proto_rawDescData)
	})
	return file_grpc_shellpb_shell_proto_rawDescData
}

var file_grpc_shellpb_shell_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_shellpb_shell_proto_goTypes = []interface{}{
	(*ExecuteRequest)(nil),  // 0: shell.ExecuteRequest
	(*ExecuteResponse)(nil), // 1: shell.ExecuteResponse
}
var file_grpc_shellpb_shell_proto_depIdxs = []int32{
	0, // 0: shell.ShellService.Execute:input_type -> shell.ExecuteRequest
	1, // 1: shell.ShellService.Execute:output_type -> shell.ExecuteResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_shellpb_shell_proto_init() }
func file_grpc_shellpb_shell_proto_init() {
	if File_grpc_shellpb_shell_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_shellpb_shell_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteRequest); i {
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
		file_grpc_shellpb_shell_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteResponse); i {
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
			RawDescriptor: file_grpc_shellpb_shell_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_shellpb_shell_proto_goTypes,
		DependencyIndexes: file_grpc_shellpb_shell_proto_depIdxs,
		MessageInfos:      file_grpc_shellpb_shell_proto_msgTypes,
	}.Build()
	File_grpc_shellpb_shell_proto = out.File
	file_grpc_shellpb_shell_proto_rawDesc = nil
	file_grpc_shellpb_shell_proto_goTypes = nil
	file_grpc_shellpb_shell_proto_depIdxs = nil
}
