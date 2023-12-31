// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: api/requestEth/v1/transaction.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type UsdtBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *UsdtBalanceRequest) Reset() {
	*x = UsdtBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_requestEth_v1_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsdtBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsdtBalanceRequest) ProtoMessage() {}

func (x *UsdtBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_requestEth_v1_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsdtBalanceRequest.ProtoReflect.Descriptor instead.
func (*UsdtBalanceRequest) Descriptor() ([]byte, []int) {
	return file_api_requestEth_v1_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *UsdtBalanceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type UsdtBalanceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance string `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	Err     string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *UsdtBalanceReply) Reset() {
	*x = UsdtBalanceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_requestEth_v1_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsdtBalanceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsdtBalanceReply) ProtoMessage() {}

func (x *UsdtBalanceReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_requestEth_v1_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsdtBalanceReply.ProtoReflect.Descriptor instead.
func (*UsdtBalanceReply) Descriptor() ([]byte, []int) {
	return file_api_requestEth_v1_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *UsdtBalanceReply) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

func (x *UsdtBalanceReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_api_requestEth_v1_transaction_proto protoreflect.FileDescriptor

var file_api_requestEth_v1_transaction_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x74, 0x68,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x45, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x12, 0x55, 0x73, 0x64, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x3e, 0x0a, 0x10, 0x55, 0x73, 0x64, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x32, 0x83, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x74, 0x0a, 0x0b, 0x55, 0x73, 0x64, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x45, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x64, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x64, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x75, 0x73, 0x64, 0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x54, 0x0a, 0x1c,
	0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50,
	0x01, 0x5a, 0x1f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x74, 0x68, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_requestEth_v1_transaction_proto_rawDescOnce sync.Once
	file_api_requestEth_v1_transaction_proto_rawDescData = file_api_requestEth_v1_transaction_proto_rawDesc
)

func file_api_requestEth_v1_transaction_proto_rawDescGZIP() []byte {
	file_api_requestEth_v1_transaction_proto_rawDescOnce.Do(func() {
		file_api_requestEth_v1_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_requestEth_v1_transaction_proto_rawDescData)
	})
	return file_api_requestEth_v1_transaction_proto_rawDescData
}

var file_api_requestEth_v1_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_requestEth_v1_transaction_proto_goTypes = []interface{}{
	(*UsdtBalanceRequest)(nil), // 0: api.requestEth.v1.UsdtBalanceRequest
	(*UsdtBalanceReply)(nil),   // 1: api.requestEth.v1.UsdtBalanceReply
}
var file_api_requestEth_v1_transaction_proto_depIdxs = []int32{
	0, // 0: api.requestEth.v1.Transaction.UsdtBalance:input_type -> api.requestEth.v1.UsdtBalanceRequest
	1, // 1: api.requestEth.v1.Transaction.UsdtBalance:output_type -> api.requestEth.v1.UsdtBalanceReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_requestEth_v1_transaction_proto_init() }
func file_api_requestEth_v1_transaction_proto_init() {
	if File_api_requestEth_v1_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_requestEth_v1_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsdtBalanceRequest); i {
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
		file_api_requestEth_v1_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsdtBalanceReply); i {
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
			RawDescriptor: file_api_requestEth_v1_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_requestEth_v1_transaction_proto_goTypes,
		DependencyIndexes: file_api_requestEth_v1_transaction_proto_depIdxs,
		MessageInfos:      file_api_requestEth_v1_transaction_proto_msgTypes,
	}.Build()
	File_api_requestEth_v1_transaction_proto = out.File
	file_api_requestEth_v1_transaction_proto_rawDesc = nil
	file_api_requestEth_v1_transaction_proto_goTypes = nil
	file_api_requestEth_v1_transaction_proto_depIdxs = nil
}
