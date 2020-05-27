// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.12.1
// source: badger.proto

package badgerpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PrefRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Company string `protobuf:"bytes,1,opt,name=company,proto3" json:"company,omitempty"`
}

func (x *PrefRequest) Reset() {
	*x = PrefRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_badger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrefRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrefRequest) ProtoMessage() {}

func (x *PrefRequest) ProtoReflect() protoreflect.Message {
	mi := &file_badger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrefRequest.ProtoReflect.Descriptor instead.
func (*PrefRequest) Descriptor() ([]byte, []int) {
	return file_badger_proto_rawDescGZIP(), []int{0}
}

func (x *PrefRequest) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

type PrefResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefecture string `protobuf:"bytes,1,opt,name=prefecture,proto3" json:"prefecture,omitempty"`
}

func (x *PrefResponse) Reset() {
	*x = PrefResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_badger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrefResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrefResponse) ProtoMessage() {}

func (x *PrefResponse) ProtoReflect() protoreflect.Message {
	mi := &file_badger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrefResponse.ProtoReflect.Descriptor instead.
func (*PrefResponse) Descriptor() ([]byte, []int) {
	return file_badger_proto_rawDescGZIP(), []int{1}
}

func (x *PrefResponse) GetPrefecture() string {
	if x != nil {
		return x.Prefecture
	}
	return ""
}

var File_badger_proto protoreflect.FileDescriptor

var file_badger_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x61, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x62, 0x61, 0x64, 0x67, 0x65, 0x72, 0x70, 0x62, 0x22, 0x27, 0x0a, 0x0b, 0x50, 0x72, 0x65, 0x66,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x22, 0x2e, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x66, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x66, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x32, 0x4a, 0x0a, 0x06, 0x42, 0x61, 0x64, 0x67, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x65, 0x66, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x15, 0x2e, 0x62,
	0x61, 0x64, 0x67, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x62, 0x61, 0x64, 0x67, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x50,
	0x72, 0x65, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x2f, 0x62, 0x61, 0x64, 0x67, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_badger_proto_rawDescOnce sync.Once
	file_badger_proto_rawDescData = file_badger_proto_rawDesc
)

func file_badger_proto_rawDescGZIP() []byte {
	file_badger_proto_rawDescOnce.Do(func() {
		file_badger_proto_rawDescData = protoimpl.X.CompressGZIP(file_badger_proto_rawDescData)
	})
	return file_badger_proto_rawDescData
}

var file_badger_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_badger_proto_goTypes = []interface{}{
	(*PrefRequest)(nil),  // 0: badgerpb.PrefRequest
	(*PrefResponse)(nil), // 1: badgerpb.PrefResponse
}
var file_badger_proto_depIdxs = []int32{
	0, // 0: badgerpb.Badger.GetPrefecture:input_type -> badgerpb.PrefRequest
	1, // 1: badgerpb.Badger.GetPrefecture:output_type -> badgerpb.PrefResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_badger_proto_init() }
func file_badger_proto_init() {
	if File_badger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_badger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrefRequest); i {
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
		file_badger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrefResponse); i {
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
			RawDescriptor: file_badger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_badger_proto_goTypes,
		DependencyIndexes: file_badger_proto_depIdxs,
		MessageInfos:      file_badger_proto_msgTypes,
	}.Build()
	File_badger_proto = out.File
	file_badger_proto_rawDesc = nil
	file_badger_proto_goTypes = nil
	file_badger_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BadgerClient is the client API for Badger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BadgerClient interface {
	GetPrefecture(ctx context.Context, in *PrefRequest, opts ...grpc.CallOption) (*PrefResponse, error)
}

type badgerClient struct {
	cc grpc.ClientConnInterface
}

func NewBadgerClient(cc grpc.ClientConnInterface) BadgerClient {
	return &badgerClient{cc}
}

func (c *badgerClient) GetPrefecture(ctx context.Context, in *PrefRequest, opts ...grpc.CallOption) (*PrefResponse, error) {
	out := new(PrefResponse)
	err := c.cc.Invoke(ctx, "/badgerpb.Badger/GetPrefecture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BadgerServer is the server API for Badger service.
type BadgerServer interface {
	GetPrefecture(context.Context, *PrefRequest) (*PrefResponse, error)
}

// UnimplementedBadgerServer can be embedded to have forward compatible implementations.
type UnimplementedBadgerServer struct {
}

func (*UnimplementedBadgerServer) GetPrefecture(context.Context, *PrefRequest) (*PrefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrefecture not implemented")
}

func RegisterBadgerServer(s *grpc.Server, srv BadgerServer) {
	s.RegisterService(&_Badger_serviceDesc, srv)
}

func _Badger_GetPrefecture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgerServer).GetPrefecture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/badgerpb.Badger/GetPrefecture",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgerServer).GetPrefecture(ctx, req.(*PrefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Badger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "badgerpb.Badger",
	HandlerType: (*BadgerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrefecture",
			Handler:    _Badger_GetPrefecture_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "badger.proto",
}