// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: proto/fibonacci.proto

package grpcdelivery

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Range struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From int32 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To   int32 `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *Range) Reset() {
	*x = Range{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fibonacci_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Range) ProtoMessage() {}

func (x *Range) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fibonacci_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Range.ProtoReflect.Descriptor instead.
func (*Range) Descriptor() ([]byte, []int) {
	return file_proto_fibonacci_proto_rawDescGZIP(), []int{0}
}

func (x *Range) GetFrom() int32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *Range) GetTo() int32 {
	if x != nil {
		return x.To
	}
	return 0
}

type FibonacciNum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FibonacciNum) Reset() {
	*x = FibonacciNum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fibonacci_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FibonacciNum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciNum) ProtoMessage() {}

func (x *FibonacciNum) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fibonacci_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciNum.ProtoReflect.Descriptor instead.
func (*FibonacciNum) Descriptor() ([]byte, []int) {
	return file_proto_fibonacci_proto_rawDescGZIP(), []int{1}
}

func (x *FibonacciNum) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_proto_fibonacci_proto protoreflect.FileDescriptor

var file_proto_fibonacci_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x05, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x74,
	0x6f, 0x22, 0x24, 0x0a, 0x0c, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x4e, 0x75,
	0x6d, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x45, 0x0a, 0x09, 0x46, 0x69, 0x62, 0x6f, 0x6e,
	0x61, 0x63, 0x63, 0x69, 0x12, 0x38, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x1a,
	0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x46, 0x69,
	0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x4e, 0x75, 0x6d, 0x22, 0x00, 0x30, 0x01, 0x42, 0x1f,
	0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x68,
	0x72, 0x64, 0x6f, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_fibonacci_proto_rawDescOnce sync.Once
	file_proto_fibonacci_proto_rawDescData = file_proto_fibonacci_proto_rawDesc
)

func file_proto_fibonacci_proto_rawDescGZIP() []byte {
	file_proto_fibonacci_proto_rawDescOnce.Do(func() {
		file_proto_fibonacci_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_fibonacci_proto_rawDescData)
	})
	return file_proto_fibonacci_proto_rawDescData
}

var file_proto_fibonacci_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_fibonacci_proto_goTypes = []interface{}{
	(*Range)(nil),        // 0: grpchandler.Range
	(*FibonacciNum)(nil), // 1: grpchandler.FibonacciNum
}
var file_proto_fibonacci_proto_depIdxs = []int32{
	0, // 0: grpchandler.Fibonacci.Get:input_type -> grpchandler.Range
	1, // 1: grpchandler.Fibonacci.Get:output_type -> grpchandler.FibonacciNum
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_fibonacci_proto_init() }
func file_proto_fibonacci_proto_init() {
	if File_proto_fibonacci_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_fibonacci_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Range); i {
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
		file_proto_fibonacci_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FibonacciNum); i {
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
			RawDescriptor: file_proto_fibonacci_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_fibonacci_proto_goTypes,
		DependencyIndexes: file_proto_fibonacci_proto_depIdxs,
		MessageInfos:      file_proto_fibonacci_proto_msgTypes,
	}.Build()
	File_proto_fibonacci_proto = out.File
	file_proto_fibonacci_proto_rawDesc = nil
	file_proto_fibonacci_proto_goTypes = nil
	file_proto_fibonacci_proto_depIdxs = nil
}
