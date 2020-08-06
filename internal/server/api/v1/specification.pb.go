// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: api/proto/v1/specification.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SpecResourcesOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *SpecResourcesOptionRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Limit   *SpecResourcesOptionLimit   `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *SpecResourcesOptions) Reset() {
	*x = SpecResourcesOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_specification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecResourcesOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecResourcesOptions) ProtoMessage() {}

func (x *SpecResourcesOptions) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_specification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecResourcesOptions.ProtoReflect.Descriptor instead.
func (*SpecResourcesOptions) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_specification_proto_rawDescGZIP(), []int{0}
}

func (x *SpecResourcesOptions) GetRequest() *SpecResourcesOptionRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *SpecResourcesOptions) GetLimit() *SpecResourcesOptionLimit {
	if x != nil {
		return x.Limit
	}
	return nil
}

type SpecResourcesOptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ram     string `protobuf:"bytes,1,opt,name=ram,proto3" json:"ram,omitempty"`
	Cpu     string `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Storage string `protobuf:"bytes,3,opt,name=storage,proto3" json:"storage,omitempty"`
}

func (x *SpecResourcesOptionRequest) Reset() {
	*x = SpecResourcesOptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_specification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecResourcesOptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecResourcesOptionRequest) ProtoMessage() {}

func (x *SpecResourcesOptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_specification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecResourcesOptionRequest.ProtoReflect.Descriptor instead.
func (*SpecResourcesOptionRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_specification_proto_rawDescGZIP(), []int{1}
}

func (x *SpecResourcesOptionRequest) GetRam() string {
	if x != nil {
		return x.Ram
	}
	return ""
}

func (x *SpecResourcesOptionRequest) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *SpecResourcesOptionRequest) GetStorage() string {
	if x != nil {
		return x.Storage
	}
	return ""
}

type SpecResourcesOptionLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ram     string `protobuf:"bytes,1,opt,name=ram,proto3" json:"ram,omitempty"`
	Cpu     string `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Storage string `protobuf:"bytes,3,opt,name=storage,proto3" json:"storage,omitempty"`
}

func (x *SpecResourcesOptionLimit) Reset() {
	*x = SpecResourcesOptionLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_specification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecResourcesOptionLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecResourcesOptionLimit) ProtoMessage() {}

func (x *SpecResourcesOptionLimit) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_specification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecResourcesOptionLimit.ProtoReflect.Descriptor instead.
func (*SpecResourcesOptionLimit) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_specification_proto_rawDescGZIP(), []int{2}
}

func (x *SpecResourcesOptionLimit) GetRam() string {
	if x != nil {
		return x.Ram
	}
	return ""
}

func (x *SpecResourcesOptionLimit) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *SpecResourcesOptionLimit) GetStorage() string {
	if x != nil {
		return x.Storage
	}
	return ""
}

var File_api_proto_v1_specification_proto protoreflect.FileDescriptor

var file_api_proto_v1_specification_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0x84, 0x01, 0x0a, 0x14, 0x53, 0x70, 0x65, 0x63, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x38, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70,
	0x65, 0x63, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5a, 0x0a,
	0x1a, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x72,
	0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x61, 0x6d, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x70, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0x58, 0x0a, 0x18, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x72, 0x61, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_v1_specification_proto_rawDescOnce sync.Once
	file_api_proto_v1_specification_proto_rawDescData = file_api_proto_v1_specification_proto_rawDesc
)

func file_api_proto_v1_specification_proto_rawDescGZIP() []byte {
	file_api_proto_v1_specification_proto_rawDescOnce.Do(func() {
		file_api_proto_v1_specification_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_v1_specification_proto_rawDescData)
	})
	return file_api_proto_v1_specification_proto_rawDescData
}

var file_api_proto_v1_specification_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_proto_v1_specification_proto_goTypes = []interface{}{
	(*SpecResourcesOptions)(nil),       // 0: v1.SpecResourcesOptions
	(*SpecResourcesOptionRequest)(nil), // 1: v1.SpecResourcesOptionRequest
	(*SpecResourcesOptionLimit)(nil),   // 2: v1.SpecResourcesOptionLimit
}
var file_api_proto_v1_specification_proto_depIdxs = []int32{
	1, // 0: v1.SpecResourcesOptions.request:type_name -> v1.SpecResourcesOptionRequest
	2, // 1: v1.SpecResourcesOptions.limit:type_name -> v1.SpecResourcesOptionLimit
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_v1_specification_proto_init() }
func file_api_proto_v1_specification_proto_init() {
	if File_api_proto_v1_specification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_v1_specification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecResourcesOptions); i {
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
		file_api_proto_v1_specification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecResourcesOptionRequest); i {
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
		file_api_proto_v1_specification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecResourcesOptionLimit); i {
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
			RawDescriptor: file_api_proto_v1_specification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_v1_specification_proto_goTypes,
		DependencyIndexes: file_api_proto_v1_specification_proto_depIdxs,
		MessageInfos:      file_api_proto_v1_specification_proto_msgTypes,
	}.Build()
	File_api_proto_v1_specification_proto = out.File
	file_api_proto_v1_specification_proto_rawDesc = nil
	file_api_proto_v1_specification_proto_goTypes = nil
	file_api_proto_v1_specification_proto_depIdxs = nil
}