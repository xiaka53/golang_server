// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: recommender.proto

package pkg

import (
	context "context"
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

//推荐人列表请求
type RecommenderListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//token
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	//每页条数
	Amount int32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	//页数
	From int32 `protobuf:"varint,3,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *RecommenderListRequest) Reset() {
	*x = RecommenderListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommender_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommenderListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommenderListRequest) ProtoMessage() {}

func (x *RecommenderListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recommender_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommenderListRequest.ProtoReflect.Descriptor instead.
func (*RecommenderListRequest) Descriptor() ([]byte, []int) {
	return file_recommender_proto_rawDescGZIP(), []int{0}
}

func (x *RecommenderListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RecommenderListRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *RecommenderListRequest) GetFrom() int32 {
	if x != nil {
		return x.From
	}
	return 0
}

//推荐人列表返回
type RecommenderListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//数据数量
	TotalCount int32 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	//推荐人数据
	List []*RecommenderInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *RecommenderListResponse) Reset() {
	*x = RecommenderListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommender_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommenderListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommenderListResponse) ProtoMessage() {}

func (x *RecommenderListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recommender_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommenderListResponse.ProtoReflect.Descriptor instead.
func (*RecommenderListResponse) Descriptor() ([]byte, []int) {
	return file_recommender_proto_rawDescGZIP(), []int{1}
}

func (x *RecommenderListResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *RecommenderListResponse) GetList() []*RecommenderInfo {
	if x != nil {
		return x.List
	}
	return nil
}

//用户信息
type RecommenderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//手机号
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	//URL
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	//账号
	Amount float32 `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
	//备注
	Memo string `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *RecommenderInfo) Reset() {
	*x = RecommenderInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommender_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommenderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommenderInfo) ProtoMessage() {}

func (x *RecommenderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_recommender_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommenderInfo.ProtoReflect.Descriptor instead.
func (*RecommenderInfo) Descriptor() ([]byte, []int) {
	return file_recommender_proto_rawDescGZIP(), []int{2}
}

func (x *RecommenderInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RecommenderInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RecommenderInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *RecommenderInfo) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *RecommenderInfo) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

var File_recommender_proto protoreflect.FileDescriptor

var file_recommender_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x22, 0x5a, 0x0a,
	0x16, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x22, 0x69, 0x0a, 0x17, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x79, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x32,
	0x5c, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x4d,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a,
	0x05, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recommender_proto_rawDescOnce sync.Once
	file_recommender_proto_rawDescData = file_recommender_proto_rawDesc
)

func file_recommender_proto_rawDescGZIP() []byte {
	file_recommender_proto_rawDescOnce.Do(func() {
		file_recommender_proto_rawDescData = protoimpl.X.CompressGZIP(file_recommender_proto_rawDescData)
	})
	return file_recommender_proto_rawDescData
}

var file_recommender_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_recommender_proto_goTypes = []interface{}{
	(*RecommenderListRequest)(nil),  // 0: interact.RecommenderListRequest
	(*RecommenderListResponse)(nil), // 1: interact.RecommenderListResponse
	(*RecommenderInfo)(nil),         // 2: interact.RecommenderInfo
}
var file_recommender_proto_depIdxs = []int32{
	2, // 0: interact.RecommenderListResponse.list:type_name -> interact.RecommenderInfo
	0, // 1: interact.Recommender.List:input_type -> interact.RecommenderListRequest
	1, // 2: interact.Recommender.List:output_type -> interact.RecommenderListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_recommender_proto_init() }
func file_recommender_proto_init() {
	if File_recommender_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recommender_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommenderListRequest); i {
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
		file_recommender_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommenderListResponse); i {
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
		file_recommender_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommenderInfo); i {
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
			RawDescriptor: file_recommender_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recommender_proto_goTypes,
		DependencyIndexes: file_recommender_proto_depIdxs,
		MessageInfos:      file_recommender_proto_msgTypes,
	}.Build()
	File_recommender_proto = out.File
	file_recommender_proto_rawDesc = nil
	file_recommender_proto_goTypes = nil
	file_recommender_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RecommenderClient is the client API for Recommender service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecommenderClient interface {
	//推荐人列表
	List(ctx context.Context, in *RecommenderListRequest, opts ...grpc.CallOption) (*RecommenderListResponse, error)
}

type recommenderClient struct {
	cc grpc.ClientConnInterface
}

func NewRecommenderClient(cc grpc.ClientConnInterface) RecommenderClient {
	return &recommenderClient{cc}
}

func (c *recommenderClient) List(ctx context.Context, in *RecommenderListRequest, opts ...grpc.CallOption) (*RecommenderListResponse, error) {
	out := new(RecommenderListResponse)
	err := c.cc.Invoke(ctx, "/interact.Recommender/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommenderServer is the server API for Recommender service.
type RecommenderServer interface {
	//推荐人列表
	List(context.Context, *RecommenderListRequest) (*RecommenderListResponse, error)
}

// UnimplementedRecommenderServer can be embedded to have forward compatible implementations.
type UnimplementedRecommenderServer struct {
}

func (*UnimplementedRecommenderServer) List(context.Context, *RecommenderListRequest) (*RecommenderListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterRecommenderServer(s *grpc.Server, srv RecommenderServer) {
	s.RegisterService(&_Recommender_serviceDesc, srv)
}

func _Recommender_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommenderListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommenderServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.Recommender/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommenderServer).List(ctx, req.(*RecommenderListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Recommender_serviceDesc = grpc.ServiceDesc{
	ServiceName: "interact.Recommender",
	HandlerType: (*RecommenderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Recommender_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recommender.proto",
}
