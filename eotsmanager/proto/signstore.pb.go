// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: signstore.proto

package proto

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

// SigningRecord represents a record of a signing operation.
// it is keyed by (chain_id || public_key || height)
type SigningRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// msg is the message that the signature is signed over
	Msg []byte `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	// eots_sig is the eots signature
	EotsSig []byte `protobuf:"bytes,2,opt,name=eots_sig,json=eotsSig,proto3" json:"eots_sig,omitempty"`
	// timestamp is the timestamp of the signing operation, in Unix seconds.
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SigningRecord) Reset() {
	*x = SigningRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signstore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningRecord) ProtoMessage() {}

func (x *SigningRecord) ProtoReflect() protoreflect.Message {
	mi := &file_signstore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigningRecord.ProtoReflect.Descriptor instead.
func (*SigningRecord) Descriptor() ([]byte, []int) {
	return file_signstore_proto_rawDescGZIP(), []int{0}
}

func (x *SigningRecord) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *SigningRecord) GetEotsSig() []byte {
	if x != nil {
		return x.EotsSig
	}
	return nil
}

func (x *SigningRecord) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_signstore_proto protoreflect.FileDescriptor

var file_signstore_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6f, 0x74, 0x73, 0x5f, 0x73, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x65,
	0x6f, 0x74, 0x73, 0x53, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x62, 0x79, 0x6c, 0x6f, 0x6e, 0x6c, 0x61, 0x62, 0x73, 0x2d, 0x69,
	0x6f, 0x2f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2f, 0x65, 0x6f, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_signstore_proto_rawDescOnce sync.Once
	file_signstore_proto_rawDescData = file_signstore_proto_rawDesc
)

func file_signstore_proto_rawDescGZIP() []byte {
	file_signstore_proto_rawDescOnce.Do(func() {
		file_signstore_proto_rawDescData = protoimpl.X.CompressGZIP(file_signstore_proto_rawDescData)
	})
	return file_signstore_proto_rawDescData
}

var file_signstore_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_signstore_proto_goTypes = []interface{}{
	(*SigningRecord)(nil), // 0: proto.SigningRecord
}
var file_signstore_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_signstore_proto_init() }
func file_signstore_proto_init() {
	if File_signstore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_signstore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigningRecord); i {
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
			RawDescriptor: file_signstore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_signstore_proto_goTypes,
		DependencyIndexes: file_signstore_proto_depIdxs,
		MessageInfos:      file_signstore_proto_msgTypes,
	}.Build()
	File_signstore_proto = out.File
	file_signstore_proto_rawDesc = nil
	file_signstore_proto_goTypes = nil
	file_signstore_proto_depIdxs = nil
}