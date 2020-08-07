// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v4.0.0
// source: parkir.proto

package entity

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type ServerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ServerReply) Reset() {
	*x = ServerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkir_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerReply) ProtoMessage() {}

func (x *ServerReply) ProtoReflect() protoreflect.Message {
	mi := &file_parkir_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerReply.ProtoReflect.Descriptor instead.
func (*ServerReply) Descriptor() ([]byte, []int) {
	return file_parkir_proto_rawDescGZIP(), []int{0}
}

func (x *ServerReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ParkID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ParkID) Reset() {
	*x = ParkID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkir_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParkID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParkID) ProtoMessage() {}

func (x *ParkID) ProtoReflect() protoreflect.Message {
	mi := &file_parkir_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParkID.ProtoReflect.Descriptor instead.
func (*ParkID) Descriptor() ([]byte, []int) {
	return file_parkir_proto_rawDescGZIP(), []int{1}
}

func (x *ParkID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Parked struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Vehicle string `protobuf:"bytes,2,opt,name=vehicle,proto3" json:"vehicle,omitempty"`
	Plate   string `protobuf:"bytes,3,opt,name=plate,proto3" json:"plate,omitempty"`
}

func (x *Parked) Reset() {
	*x = Parked{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkir_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Parked) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parked) ProtoMessage() {}

func (x *Parked) ProtoReflect() protoreflect.Message {
	mi := &file_parkir_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parked.ProtoReflect.Descriptor instead.
func (*Parked) Descriptor() ([]byte, []int) {
	return file_parkir_proto_rawDescGZIP(), []int{2}
}

func (x *Parked) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Parked) GetVehicle() string {
	if x != nil {
		return x.Vehicle
	}
	return ""
}

func (x *Parked) GetPlate() string {
	if x != nil {
		return x.Plate
	}
	return ""
}

type ParkBill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vehicle string `protobuf:"bytes,1,opt,name=vehicle,proto3" json:"vehicle,omitempty"`
	Plate   string `protobuf:"bytes,2,opt,name=plate,proto3" json:"plate,omitempty"`
	Entry   string `protobuf:"bytes,3,opt,name=entry,proto3" json:"entry,omitempty"`
	Exit    string `protobuf:"bytes,4,opt,name=exit,proto3" json:"exit,omitempty"`
	Bill    int32  `protobuf:"varint,5,opt,name=Bill,proto3" json:"Bill,omitempty"`
}

func (x *ParkBill) Reset() {
	*x = ParkBill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkir_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParkBill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParkBill) ProtoMessage() {}

func (x *ParkBill) ProtoReflect() protoreflect.Message {
	mi := &file_parkir_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParkBill.ProtoReflect.Descriptor instead.
func (*ParkBill) Descriptor() ([]byte, []int) {
	return file_parkir_proto_rawDescGZIP(), []int{3}
}

func (x *ParkBill) GetVehicle() string {
	if x != nil {
		return x.Vehicle
	}
	return ""
}

func (x *ParkBill) GetPlate() string {
	if x != nil {
		return x.Plate
	}
	return ""
}

func (x *ParkBill) GetEntry() string {
	if x != nil {
		return x.Entry
	}
	return ""
}

func (x *ParkBill) GetExit() string {
	if x != nil {
		return x.Exit
	}
	return ""
}

func (x *ParkBill) GetBill() int32 {
	if x != nil {
		return x.Bill
	}
	return 0
}

var File_parkir_proto protoreflect.FileDescriptor

var file_parkir_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x74, 0x61, 0x73, 0x6b, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x27, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x18, 0x0a, 0x06, 0x50, 0x61,
	0x72, 0x6b, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x48, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x78,
	0x0a, 0x08, 0x50, 0x61, 0x72, 0x6b, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x65, 0x78, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x69, 0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x42, 0x69, 0x6c, 0x6c, 0x32, 0xad, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x12, 0x3e, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x11, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x6b, 0x49,
	0x44, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x50, 0x61, 0x72, 0x6b, 0x49, 0x44, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x72, 0x6b, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x0c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x50, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x1a, 0x0e, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x50, 0x61,
	0x72, 0x6b, 0x42, 0x69, 0x6c, 0x6c, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_parkir_proto_rawDescOnce sync.Once
	file_parkir_proto_rawDescData = file_parkir_proto_rawDesc
)

func file_parkir_proto_rawDescGZIP() []byte {
	file_parkir_proto_rawDescOnce.Do(func() {
		file_parkir_proto_rawDescData = protoimpl.X.CompressGZIP(file_parkir_proto_rawDescData)
	})
	return file_parkir_proto_rawDescData
}

var file_parkir_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_parkir_proto_goTypes = []interface{}{
	(*ServerReply)(nil), // 0: task.ServerReply
	(*ParkID)(nil),      // 1: task.ParkID
	(*Parked)(nil),      // 2: task.Parked
	(*ParkBill)(nil),    // 3: task.ParkBill
	(*empty.Empty)(nil), // 4: google.protobuf.Empty
}
var file_parkir_proto_depIdxs = []int32{
	4, // 0: task.Connect.ConnectToServer:input_type -> google.protobuf.Empty
	4, // 1: task.Connect.GetParkID:input_type -> google.protobuf.Empty
	2, // 2: task.Connect.GetParkBill:input_type -> task.Parked
	0, // 3: task.Connect.ConnectToServer:output_type -> task.ServerReply
	1, // 4: task.Connect.GetParkID:output_type -> task.ParkID
	3, // 5: task.Connect.GetParkBill:output_type -> task.ParkBill
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_parkir_proto_init() }
func file_parkir_proto_init() {
	if File_parkir_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_parkir_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerReply); i {
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
		file_parkir_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParkID); i {
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
		file_parkir_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Parked); i {
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
		file_parkir_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParkBill); i {
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
			RawDescriptor: file_parkir_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_parkir_proto_goTypes,
		DependencyIndexes: file_parkir_proto_depIdxs,
		MessageInfos:      file_parkir_proto_msgTypes,
	}.Build()
	File_parkir_proto = out.File
	file_parkir_proto_rawDesc = nil
	file_parkir_proto_goTypes = nil
	file_parkir_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConnectClient is the client API for Connect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnectClient interface {
	ConnectToServer(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ServerReply, error)
	GetParkID(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ParkID, error)
	GetParkBill(ctx context.Context, in *Parked, opts ...grpc.CallOption) (*ParkBill, error)
}

type connectClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectClient(cc grpc.ClientConnInterface) ConnectClient {
	return &connectClient{cc}
}

func (c *connectClient) ConnectToServer(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ServerReply, error) {
	out := new(ServerReply)
	err := c.cc.Invoke(ctx, "/task.Connect/ConnectToServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) GetParkID(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ParkID, error) {
	out := new(ParkID)
	err := c.cc.Invoke(ctx, "/task.Connect/GetParkID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) GetParkBill(ctx context.Context, in *Parked, opts ...grpc.CallOption) (*ParkBill, error) {
	out := new(ParkBill)
	err := c.cc.Invoke(ctx, "/task.Connect/GetParkBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectServer is the server API for Connect service.
type ConnectServer interface {
	ConnectToServer(context.Context, *empty.Empty) (*ServerReply, error)
	GetParkID(context.Context, *empty.Empty) (*ParkID, error)
	GetParkBill(context.Context, *Parked) (*ParkBill, error)
}

// UnimplementedConnectServer can be embedded to have forward compatible implementations.
type UnimplementedConnectServer struct {
}

func (*UnimplementedConnectServer) ConnectToServer(context.Context, *empty.Empty) (*ServerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectToServer not implemented")
}
func (*UnimplementedConnectServer) GetParkID(context.Context, *empty.Empty) (*ParkID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParkID not implemented")
}
func (*UnimplementedConnectServer) GetParkBill(context.Context, *Parked) (*ParkBill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParkBill not implemented")
}

func RegisterConnectServer(s *grpc.Server, srv ConnectServer) {
	s.RegisterService(&_Connect_serviceDesc, srv)
}

func _Connect_ConnectToServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).ConnectToServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Connect/ConnectToServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).ConnectToServer(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_GetParkID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).GetParkID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Connect/GetParkID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).GetParkID(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_GetParkBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Parked)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).GetParkBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Connect/GetParkBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).GetParkBill(ctx, req.(*Parked))
	}
	return interceptor(ctx, in, info, handler)
}

var _Connect_serviceDesc = grpc.ServiceDesc{
	ServiceName: "task.Connect",
	HandlerType: (*ConnectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConnectToServer",
			Handler:    _Connect_ConnectToServer_Handler,
		},
		{
			MethodName: "GetParkID",
			Handler:    _Connect_GetParkID_Handler,
		},
		{
			MethodName: "GetParkBill",
			Handler:    _Connect_GetParkBill_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parkir.proto",
}