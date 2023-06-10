// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: complex_map.proto

package proto_gen

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

type ComplexAndHuge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComplexMap map[string]*Huge `protobuf:"bytes,1,rep,name=complex_map,json=complexMap,proto3" json:"complex_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ComplexAndHuge) Reset() {
	*x = ComplexAndHuge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_map_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexAndHuge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexAndHuge) ProtoMessage() {}

func (x *ComplexAndHuge) ProtoReflect() protoreflect.Message {
	mi := &file_complex_map_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexAndHuge.ProtoReflect.Descriptor instead.
func (*ComplexAndHuge) Descriptor() ([]byte, []int) {
	return file_complex_map_proto_rawDescGZIP(), []int{0}
}

func (x *ComplexAndHuge) GetComplexMap() map[string]*Huge {
	if x != nil {
		return x.ComplexMap
	}
	return nil
}

type ComplexAndHugeArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComplexArr []*ComplexAndHuge `protobuf:"bytes,1,rep,name=complex_arr,json=complexArr,proto3" json:"complex_arr,omitempty"`
}

func (x *ComplexAndHugeArray) Reset() {
	*x = ComplexAndHugeArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_map_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexAndHugeArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexAndHugeArray) ProtoMessage() {}

func (x *ComplexAndHugeArray) ProtoReflect() protoreflect.Message {
	mi := &file_complex_map_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexAndHugeArray.ProtoReflect.Descriptor instead.
func (*ComplexAndHugeArray) Descriptor() ([]byte, []int) {
	return file_complex_map_proto_rawDescGZIP(), []int{1}
}

func (x *ComplexAndHugeArray) GetComplexArr() []*ComplexAndHuge {
	if x != nil {
		return x.ComplexArr
	}
	return nil
}

var File_complex_map_proto protoreflect.FileDescriptor

var file_complex_map_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x62, 0x2e, 0x67, 0x65, 0x6e, 0x1a, 0x0a, 0x68, 0x75, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x78, 0x41, 0x6e, 0x64, 0x48, 0x75, 0x67, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x70, 0x62, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78,
	0x41, 0x6e, 0x64, 0x48, 0x75, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78,
	0x4d, 0x61, 0x70, 0x1a, 0x4b, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x67, 0x65, 0x6e,
	0x2e, 0x48, 0x75, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x4e, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x41, 0x6e, 0x64, 0x48, 0x75,
	0x67, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x37, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x5f, 0x61, 0x72, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x62, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x41, 0x6e, 0x64,
	0x48, 0x75, 0x67, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x41, 0x72, 0x72,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52,
	0x53, 0x68, 0x65, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x67, 0x6f, 0x62, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_complex_map_proto_rawDescOnce sync.Once
	file_complex_map_proto_rawDescData = file_complex_map_proto_rawDesc
)

func file_complex_map_proto_rawDescGZIP() []byte {
	file_complex_map_proto_rawDescOnce.Do(func() {
		file_complex_map_proto_rawDescData = protoimpl.X.CompressGZIP(file_complex_map_proto_rawDescData)
	})
	return file_complex_map_proto_rawDescData
}

var file_complex_map_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_complex_map_proto_goTypes = []interface{}{
	(*ComplexAndHuge)(nil),      // 0: pb.gen.ComplexAndHuge
	(*ComplexAndHugeArray)(nil), // 1: pb.gen.ComplexAndHugeArray
	nil,                         // 2: pb.gen.ComplexAndHuge.ComplexMapEntry
	(*Huge)(nil),                // 3: pb.gen.Huge
}
var file_complex_map_proto_depIdxs = []int32{
	2, // 0: pb.gen.ComplexAndHuge.complex_map:type_name -> pb.gen.ComplexAndHuge.ComplexMapEntry
	0, // 1: pb.gen.ComplexAndHugeArray.complex_arr:type_name -> pb.gen.ComplexAndHuge
	3, // 2: pb.gen.ComplexAndHuge.ComplexMapEntry.value:type_name -> pb.gen.Huge
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_complex_map_proto_init() }
func file_complex_map_proto_init() {
	if File_complex_map_proto != nil {
		return
	}
	file_huge_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_complex_map_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexAndHuge); i {
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
		file_complex_map_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexAndHugeArray); i {
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
			RawDescriptor: file_complex_map_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_complex_map_proto_goTypes,
		DependencyIndexes: file_complex_map_proto_depIdxs,
		MessageInfos:      file_complex_map_proto_msgTypes,
	}.Build()
	File_complex_map_proto = out.File
	file_complex_map_proto_rawDesc = nil
	file_complex_map_proto_goTypes = nil
	file_complex_map_proto_depIdxs = nil
}