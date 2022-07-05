// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: mpath/v1/cvr.proto

package v1

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

type Sample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *Metadata    `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Variants *CNVVariants `protobuf:"bytes,2,opt,name=variants,proto3" json:"variants,omitempty"`
}

func (x *Sample) Reset() {
	*x = Sample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpath_v1_cvr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sample) ProtoMessage() {}

func (x *Sample) ProtoReflect() protoreflect.Message {
	mi := &file_mpath_v1_cvr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sample.ProtoReflect.Descriptor instead.
func (*Sample) Descriptor() ([]byte, []int) {
	return file_mpath_v1_cvr_proto_rawDescGZIP(), []int{0}
}

func (x *Sample) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Sample) GetVariants() *CNVVariants {
	if x != nil {
		return x.Variants
	}
	return nil
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DMPSampleId string `protobuf:"bytes,1,opt,name=DMPSampleId,proto3" json:"DMPSampleId,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpath_v1_cvr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_mpath_v1_cvr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_mpath_v1_cvr_proto_rawDescGZIP(), []int{1}
}

func (x *Metadata) GetDMPSampleId() string {
	if x != nil {
		return x.DMPSampleId
	}
	return ""
}

type CNVVariants struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Variant []*CVNVariant `protobuf:"bytes,1,rep,name=Variant,proto3" json:"Variant,omitempty"`
}

func (x *CNVVariants) Reset() {
	*x = CNVVariants{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpath_v1_cvr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CNVVariants) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CNVVariants) ProtoMessage() {}

func (x *CNVVariants) ProtoReflect() protoreflect.Message {
	mi := &file_mpath_v1_cvr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CNVVariants.ProtoReflect.Descriptor instead.
func (*CNVVariants) Descriptor() ([]byte, []int) {
	return file_mpath_v1_cvr_proto_rawDescGZIP(), []int{2}
}

func (x *CNVVariants) GetVariant() []*CVNVariant {
	if x != nil {
		return x.Variant
	}
	return nil
}

type CVNVariant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GeneId string `protobuf:"bytes,1,opt,name=GeneId,proto3" json:"GeneId,omitempty"`
}

func (x *CVNVariant) Reset() {
	*x = CVNVariant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpath_v1_cvr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CVNVariant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CVNVariant) ProtoMessage() {}

func (x *CVNVariant) ProtoReflect() protoreflect.Message {
	mi := &file_mpath_v1_cvr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CVNVariant.ProtoReflect.Descriptor instead.
func (*CVNVariant) Descriptor() ([]byte, []int) {
	return file_mpath_v1_cvr_proto_rawDescGZIP(), []int{3}
}

func (x *CVNVariant) GetGeneId() string {
	if x != nil {
		return x.GeneId
	}
	return ""
}

var File_mpath_v1_cvr_proto protoreflect.FileDescriptor

var file_mpath_v1_cvr_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x70, 0x61, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x76, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x06, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x25,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x4e, 0x56, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x22,
	0x2c, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x4d, 0x50, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x44, 0x4d, 0x50, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x34, 0x0a,
	0x0b, 0x43, 0x4e, 0x56, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x07,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x43, 0x56, 0x4e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x07, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x74, 0x22, 0x24, 0x0a, 0x0a, 0x43, 0x56, 0x4e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x73, 0x6b, 0x63, 0x63, 0x2f, 0x73, 0x6d,
	0x69, 0x6c, 0x65, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x70, 0x61, 0x74,
	0x68, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mpath_v1_cvr_proto_rawDescOnce sync.Once
	file_mpath_v1_cvr_proto_rawDescData = file_mpath_v1_cvr_proto_rawDesc
)

func file_mpath_v1_cvr_proto_rawDescGZIP() []byte {
	file_mpath_v1_cvr_proto_rawDescOnce.Do(func() {
		file_mpath_v1_cvr_proto_rawDescData = protoimpl.X.CompressGZIP(file_mpath_v1_cvr_proto_rawDescData)
	})
	return file_mpath_v1_cvr_proto_rawDescData
}

var file_mpath_v1_cvr_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mpath_v1_cvr_proto_goTypes = []interface{}{
	(*Sample)(nil),      // 0: Sample
	(*Metadata)(nil),    // 1: Metadata
	(*CNVVariants)(nil), // 2: CNVVariants
	(*CVNVariant)(nil),  // 3: CVNVariant
}
var file_mpath_v1_cvr_proto_depIdxs = []int32{
	1, // 0: Sample.metadata:type_name -> Metadata
	2, // 1: Sample.variants:type_name -> CNVVariants
	3, // 2: CNVVariants.Variant:type_name -> CVNVariant
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mpath_v1_cvr_proto_init() }
func file_mpath_v1_cvr_proto_init() {
	if File_mpath_v1_cvr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mpath_v1_cvr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sample); i {
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
		file_mpath_v1_cvr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_mpath_v1_cvr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CNVVariants); i {
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
		file_mpath_v1_cvr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CVNVariant); i {
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
			RawDescriptor: file_mpath_v1_cvr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mpath_v1_cvr_proto_goTypes,
		DependencyIndexes: file_mpath_v1_cvr_proto_depIdxs,
		MessageInfos:      file_mpath_v1_cvr_proto_msgTypes,
	}.Build()
	File_mpath_v1_cvr_proto = out.File
	file_mpath_v1_cvr_proto_rawDesc = nil
	file_mpath_v1_cvr_proto_goTypes = nil
	file_mpath_v1_cvr_proto_depIdxs = nil
}
