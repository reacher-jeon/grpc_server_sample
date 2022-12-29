// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: sample/v1/sample_service.proto

package samplev1

import (
	v1 "github.com/reacher-jeon/grpc_server_sample/pb/type/v1"
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

type GetInfoInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *v1.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Info   *GetInfo   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *GetInfoInfoRequest) Reset() {
	*x = GetInfoInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_v1_sample_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoInfoRequest) ProtoMessage() {}

func (x *GetInfoInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sample_v1_sample_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoInfoRequest) Descriptor() ([]byte, []int) {
	return file_sample_v1_sample_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetInfoInfoRequest) GetHeader() *v1.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetInfoInfoRequest) GetInfo() *GetInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *v1.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Error   *v1.Error  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Message string     `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetInfoResponse) Reset() {
	*x = GetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_v1_sample_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse) ProtoMessage() {}

func (x *GetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sample_v1_sample_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return file_sample_v1_sample_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetInfoResponse) GetHeader() *v1.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetInfoResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GetInfoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_sample_v1_sample_service_proto protoreflect.FileDescriptor

var file_sample_v1_sample_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x65, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x26,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x7a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x57, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d,
	0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x7f, 0x0a, 0x0d, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x15, 0x70, 0x62, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa,
	0x02, 0x09, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0a, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sample_v1_sample_service_proto_rawDescOnce sync.Once
	file_sample_v1_sample_service_proto_rawDescData = file_sample_v1_sample_service_proto_rawDesc
)

func file_sample_v1_sample_service_proto_rawDescGZIP() []byte {
	file_sample_v1_sample_service_proto_rawDescOnce.Do(func() {
		file_sample_v1_sample_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_sample_v1_sample_service_proto_rawDescData)
	})
	return file_sample_v1_sample_service_proto_rawDescData
}

var file_sample_v1_sample_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sample_v1_sample_service_proto_goTypes = []interface{}{
	(*GetInfoInfoRequest)(nil), // 0: sample.v1.GetInfoInfoRequest
	(*GetInfoResponse)(nil),    // 1: sample.v1.GetInfoResponse
	(*v1.Header)(nil),          // 2: type.v1.Header
	(*GetInfo)(nil),            // 3: sample.v1.GetInfo
	(*v1.Error)(nil),           // 4: type.v1.Error
}
var file_sample_v1_sample_service_proto_depIdxs = []int32{
	2, // 0: sample.v1.GetInfoInfoRequest.header:type_name -> type.v1.Header
	3, // 1: sample.v1.GetInfoInfoRequest.info:type_name -> sample.v1.GetInfo
	2, // 2: sample.v1.GetInfoResponse.header:type_name -> type.v1.Header
	4, // 3: sample.v1.GetInfoResponse.error:type_name -> type.v1.Error
	0, // 4: sample.v1.SampleService.GetInfo:input_type -> sample.v1.GetInfoInfoRequest
	1, // 5: sample.v1.SampleService.GetInfo:output_type -> sample.v1.GetInfoResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_sample_v1_sample_service_proto_init() }
func file_sample_v1_sample_service_proto_init() {
	if File_sample_v1_sample_service_proto != nil {
		return
	}
	file_sample_v1_sample_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sample_v1_sample_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoInfoRequest); i {
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
		file_sample_v1_sample_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoResponse); i {
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
			RawDescriptor: file_sample_v1_sample_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sample_v1_sample_service_proto_goTypes,
		DependencyIndexes: file_sample_v1_sample_service_proto_depIdxs,
		MessageInfos:      file_sample_v1_sample_service_proto_msgTypes,
	}.Build()
	File_sample_v1_sample_service_proto = out.File
	file_sample_v1_sample_service_proto_rawDesc = nil
	file_sample_v1_sample_service_proto_goTypes = nil
	file_sample_v1_sample_service_proto_depIdxs = nil
}
