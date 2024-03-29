// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: record.proto

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

//状态
type State int32

const (
	// 未指定
	State_Unspecified State = 0
	//完成
	State_Finish State = 1
	//待审核
	State_Review State = 2
	//取消
	State_Cancel State = 3
	//执行失败
	State_Fail State = 4
	//申诉
	State_Appeal State = 5
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "Unspecified",
		1: "Finish",
		2: "Review",
		3: "Cancel",
		4: "Fail",
		5: "Appeal",
	}
	State_value = map[string]int32{
		"Unspecified": 0,
		"Finish":      1,
		"Review":      2,
		"Cancel":      3,
		"Fail":        4,
		"Appeal":      5,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_record_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_record_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_record_proto_rawDescGZIP(), []int{0}
}

//查询流水
type RecordListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//token
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	//页数
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	//每页条数
	From int32 `protobuf:"varint,3,opt,name=from,proto3" json:"from,omitempty"`
	//状态
	State *State `protobuf:"varint,4,opt,name=state,proto3,enum=interact.State,oneof" json:"state,omitempty"`
	//代理商
	Agent *string `protobuf:"bytes,5,opt,name=agent,proto3,oneof" json:"agent,omitempty"`
}

func (x *RecordListRequest) Reset() {
	*x = RecordListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordListRequest) ProtoMessage() {}

func (x *RecordListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordListRequest.ProtoReflect.Descriptor instead.
func (*RecordListRequest) Descriptor() ([]byte, []int) {
	return file_record_proto_rawDescGZIP(), []int{0}
}

func (x *RecordListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RecordListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *RecordListRequest) GetFrom() int32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *RecordListRequest) GetState() State {
	if x != nil && x.State != nil {
		return *x.State
	}
	return State_Unspecified
}

func (x *RecordListRequest) GetAgent() string {
	if x != nil && x.Agent != nil {
		return *x.Agent
	}
	return ""
}

//返回流水信息
type RecordListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//数据总条数
	TotalCount int32 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	//订单数据
	List []*RecordInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *RecordListResponse) Reset() {
	*x = RecordListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordListResponse) ProtoMessage() {}

func (x *RecordListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordListResponse.ProtoReflect.Descriptor instead.
func (*RecordListResponse) Descriptor() ([]byte, []int) {
	return file_record_proto_rawDescGZIP(), []int{1}
}

func (x *RecordListResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *RecordListResponse) GetList() []*RecordInfo {
	if x != nil {
		return x.List
	}
	return nil
}

//取消订单
type CancelRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//token
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	//标识
	Mark int32 `protobuf:"varint,2,opt,name=mark,proto3" json:"mark,omitempty"`
}

func (x *CancelRecordRequest) Reset() {
	*x = CancelRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_record_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelRecordRequest) ProtoMessage() {}

func (x *CancelRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_record_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelRecordRequest.ProtoReflect.Descriptor instead.
func (*CancelRecordRequest) Descriptor() ([]byte, []int) {
	return file_record_proto_rawDescGZIP(), []int{2}
}

func (x *CancelRecordRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CancelRecordRequest) GetMark() int32 {
	if x != nil {
		return x.Mark
	}
	return 0
}

//充值完成
type FinishRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//token
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	//标识
	Mark int32 `protobuf:"varint,2,opt,name=mark,proto3" json:"mark,omitempty"`
}

func (x *FinishRecordRequest) Reset() {
	*x = FinishRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_record_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishRecordRequest) ProtoMessage() {}

func (x *FinishRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_record_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishRecordRequest.ProtoReflect.Descriptor instead.
func (*FinishRecordRequest) Descriptor() ([]byte, []int) {
	return file_record_proto_rawDescGZIP(), []int{3}
}

func (x *FinishRecordRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *FinishRecordRequest) GetMark() int32 {
	if x != nil {
		return x.Mark
	}
	return 0
}

//申请售后
type ApplyAfterSalesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//token
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	//标识
	Mark int32 `protobuf:"varint,2,opt,name=mark,proto3" json:"mark,omitempty"`
	//售后内容
	Msg string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ApplyAfterSalesRequest) Reset() {
	*x = ApplyAfterSalesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_record_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyAfterSalesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyAfterSalesRequest) ProtoMessage() {}

func (x *ApplyAfterSalesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_record_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyAfterSalesRequest.ProtoReflect.Descriptor instead.
func (*ApplyAfterSalesRequest) Descriptor() ([]byte, []int) {
	return file_record_proto_rawDescGZIP(), []int{4}
}

func (x *ApplyAfterSalesRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ApplyAfterSalesRequest) GetMark() int32 {
	if x != nil {
		return x.Mark
	}
	return 0
}

func (x *ApplyAfterSalesRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

//填写备注
type MemoRepuest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//token
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	//标识
	Mark int32 `protobuf:"varint,2,opt,name=mark,proto3" json:"mark,omitempty"`
	//备注
	Memo string `protobuf:"bytes,3,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *MemoRepuest) Reset() {
	*x = MemoRepuest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_record_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoRepuest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoRepuest) ProtoMessage() {}

func (x *MemoRepuest) ProtoReflect() protoreflect.Message {
	mi := &file_record_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoRepuest.ProtoReflect.Descriptor instead.
func (*MemoRepuest) Descriptor() ([]byte, []int) {
	return file_record_proto_rawDescGZIP(), []int{5}
}

func (x *MemoRepuest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *MemoRepuest) GetMark() int32 {
	if x != nil {
		return x.Mark
	}
	return 0
}

func (x *MemoRepuest) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

//流水信息
type RecordInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//代理商
	Agent string `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty"`
	//充值帐户
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	//充值会员种类
	Vip string `protobuf:"bytes,3,opt,name=vip,proto3" json:"vip,omitempty"`
	//价格
	Price float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	//状态
	State State `protobuf:"varint,5,opt,name=state,proto3,enum=interact.State" json:"state,omitempty"`
	//执行次数
	ExecCount int32 `protobuf:"varint,6,opt,name=exec_count,json=execCount,proto3" json:"exec_count,omitempty"`
	//申诉内容
	Msg string `protobuf:"bytes,7,opt,name=msg,proto3" json:"msg,omitempty"`
	//备注信息
	Memo string `protobuf:"bytes,8,opt,name=memo,proto3" json:"memo,omitempty"`
	//标识
	Mark int32 `protobuf:"varint,9,opt,name=mark,proto3" json:"mark,omitempty"`
}

func (x *RecordInfo) Reset() {
	*x = RecordInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_record_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordInfo) ProtoMessage() {}

func (x *RecordInfo) ProtoReflect() protoreflect.Message {
	mi := &file_record_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordInfo.ProtoReflect.Descriptor instead.
func (*RecordInfo) Descriptor() ([]byte, []int) {
	return file_record_proto_rawDescGZIP(), []int{6}
}

func (x *RecordInfo) GetAgent() string {
	if x != nil {
		return x.Agent
	}
	return ""
}

func (x *RecordInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *RecordInfo) GetVip() string {
	if x != nil {
		return x.Vip
	}
	return ""
}

func (x *RecordInfo) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *RecordInfo) GetState() State {
	if x != nil {
		return x.State
	}
	return State_Unspecified
}

func (x *RecordInfo) GetExecCount() int32 {
	if x != nil {
		return x.ExecCount
	}
	return 0
}

func (x *RecordInfo) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *RecordInfo) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *RecordInfo) GetMark() int32 {
	if x != nil {
		return x.Mark
	}
	return 0
}

//执行返回
type RecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//执行状态
	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	//执行错误时错误信息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *RecordResponse) Reset() {
	*x = RecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_record_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordResponse) ProtoMessage() {}

func (x *RecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_record_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordResponse.ProtoReflect.Descriptor instead.
func (*RecordResponse) Descriptor() ([]byte, []int) {
	return file_record_proto_rawDescGZIP(), []int{7}
}

func (x *RecordResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *RecordResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_record_proto protoreflect.FileDescriptor

var file_record_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x22, 0xac, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2a, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x5f, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x13, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x3f, 0x0a, 0x13, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x54, 0x0a, 0x16, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x41, 0x66, 0x74, 0x65, 0x72, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61,
	0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x22, 0x4b, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x22, 0xe4, 0x01,
	0x0a, 0x0a, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x76, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x69, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x78, 0x65, 0x63, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x65, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x6d, 0x61, 0x72, 0x6b, 0x22, 0x3a, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x2a, 0x52, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x10, 0x03, 0x12, 0x08,
	0x0a, 0x04, 0x46, 0x61, 0x69, 0x6c, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x70, 0x70, 0x65,
	0x61, 0x6c, 0x10, 0x05, 0x32, 0xe3, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x43, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x1d,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x12, 0x1d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f,
	0x0a, 0x0f, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x41, 0x66, 0x74, 0x65, 0x72, 0x53, 0x61, 0x6c, 0x65,
	0x73, 0x12, 0x20, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x41, 0x66, 0x74, 0x65, 0x72, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x39, 0x0a, 0x04, 0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x15, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x70, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f,
	0x70, 0x6b, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_record_proto_rawDescOnce sync.Once
	file_record_proto_rawDescData = file_record_proto_rawDesc
)

func file_record_proto_rawDescGZIP() []byte {
	file_record_proto_rawDescOnce.Do(func() {
		file_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_record_proto_rawDescData)
	})
	return file_record_proto_rawDescData
}

var file_record_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_record_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_record_proto_goTypes = []interface{}{
	(State)(0),                     // 0: interact.State
	(*RecordListRequest)(nil),      // 1: interact.RecordListRequest
	(*RecordListResponse)(nil),     // 2: interact.RecordListResponse
	(*CancelRecordRequest)(nil),    // 3: interact.CancelRecordRequest
	(*FinishRecordRequest)(nil),    // 4: interact.FinishRecordRequest
	(*ApplyAfterSalesRequest)(nil), // 5: interact.ApplyAfterSalesRequest
	(*MemoRepuest)(nil),            // 6: interact.MemoRepuest
	(*RecordInfo)(nil),             // 7: interact.RecordInfo
	(*RecordResponse)(nil),         // 8: interact.RecordResponse
}
var file_record_proto_depIdxs = []int32{
	0, // 0: interact.RecordListRequest.state:type_name -> interact.State
	7, // 1: interact.RecordListResponse.list:type_name -> interact.RecordInfo
	0, // 2: interact.RecordInfo.state:type_name -> interact.State
	1, // 3: interact.Record.List:input_type -> interact.RecordListRequest
	3, // 4: interact.Record.Cancel:input_type -> interact.CancelRecordRequest
	4, // 5: interact.Record.Finish:input_type -> interact.FinishRecordRequest
	5, // 6: interact.Record.ApplyAfterSales:input_type -> interact.ApplyAfterSalesRequest
	6, // 7: interact.Record.Memo:input_type -> interact.MemoRepuest
	2, // 8: interact.Record.List:output_type -> interact.RecordListResponse
	8, // 9: interact.Record.Cancel:output_type -> interact.RecordResponse
	8, // 10: interact.Record.Finish:output_type -> interact.RecordResponse
	8, // 11: interact.Record.ApplyAfterSales:output_type -> interact.RecordResponse
	8, // 12: interact.Record.Memo:output_type -> interact.RecordResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_record_proto_init() }
func file_record_proto_init() {
	if File_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordListRequest); i {
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
		file_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordListResponse); i {
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
		file_record_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelRecordRequest); i {
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
		file_record_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishRecordRequest); i {
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
		file_record_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyAfterSalesRequest); i {
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
		file_record_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoRepuest); i {
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
		file_record_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordInfo); i {
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
		file_record_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordResponse); i {
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
	file_record_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_record_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_record_proto_goTypes,
		DependencyIndexes: file_record_proto_depIdxs,
		EnumInfos:         file_record_proto_enumTypes,
		MessageInfos:      file_record_proto_msgTypes,
	}.Build()
	File_record_proto = out.File
	file_record_proto_rawDesc = nil
	file_record_proto_goTypes = nil
	file_record_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RecordClient is the client API for Record service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecordClient interface {
	//查询流水
	List(ctx context.Context, in *RecordListRequest, opts ...grpc.CallOption) (*RecordListResponse, error)
	//取消充值
	Cancel(ctx context.Context, in *CancelRecordRequest, opts ...grpc.CallOption) (*RecordResponse, error)
	//充值完成
	Finish(ctx context.Context, in *FinishRecordRequest, opts ...grpc.CallOption) (*RecordResponse, error)
	//申请售后
	ApplyAfterSales(ctx context.Context, in *ApplyAfterSalesRequest, opts ...grpc.CallOption) (*RecordResponse, error)
	//填写备注
	Memo(ctx context.Context, in *MemoRepuest, opts ...grpc.CallOption) (*RecordResponse, error)
}

type recordClient struct {
	cc grpc.ClientConnInterface
}

func NewRecordClient(cc grpc.ClientConnInterface) RecordClient {
	return &recordClient{cc}
}

func (c *recordClient) List(ctx context.Context, in *RecordListRequest, opts ...grpc.CallOption) (*RecordListResponse, error) {
	out := new(RecordListResponse)
	err := c.cc.Invoke(ctx, "/interact.Record/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordClient) Cancel(ctx context.Context, in *CancelRecordRequest, opts ...grpc.CallOption) (*RecordResponse, error) {
	out := new(RecordResponse)
	err := c.cc.Invoke(ctx, "/interact.Record/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordClient) Finish(ctx context.Context, in *FinishRecordRequest, opts ...grpc.CallOption) (*RecordResponse, error) {
	out := new(RecordResponse)
	err := c.cc.Invoke(ctx, "/interact.Record/Finish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordClient) ApplyAfterSales(ctx context.Context, in *ApplyAfterSalesRequest, opts ...grpc.CallOption) (*RecordResponse, error) {
	out := new(RecordResponse)
	err := c.cc.Invoke(ctx, "/interact.Record/ApplyAfterSales", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordClient) Memo(ctx context.Context, in *MemoRepuest, opts ...grpc.CallOption) (*RecordResponse, error) {
	out := new(RecordResponse)
	err := c.cc.Invoke(ctx, "/interact.Record/Memo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecordServer is the server API for Record service.
type RecordServer interface {
	//查询流水
	List(context.Context, *RecordListRequest) (*RecordListResponse, error)
	//取消充值
	Cancel(context.Context, *CancelRecordRequest) (*RecordResponse, error)
	//充值完成
	Finish(context.Context, *FinishRecordRequest) (*RecordResponse, error)
	//申请售后
	ApplyAfterSales(context.Context, *ApplyAfterSalesRequest) (*RecordResponse, error)
	//填写备注
	Memo(context.Context, *MemoRepuest) (*RecordResponse, error)
}

// UnimplementedRecordServer can be embedded to have forward compatible implementations.
type UnimplementedRecordServer struct {
}

func (*UnimplementedRecordServer) List(context.Context, *RecordListRequest) (*RecordListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedRecordServer) Cancel(context.Context, *CancelRecordRequest) (*RecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (*UnimplementedRecordServer) Finish(context.Context, *FinishRecordRequest) (*RecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Finish not implemented")
}
func (*UnimplementedRecordServer) ApplyAfterSales(context.Context, *ApplyAfterSalesRequest) (*RecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyAfterSales not implemented")
}
func (*UnimplementedRecordServer) Memo(context.Context, *MemoRepuest) (*RecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Memo not implemented")
}

func RegisterRecordServer(s *grpc.Server, srv RecordServer) {
	s.RegisterService(&_Record_serviceDesc, srv)
}

func _Record_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.Record/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).List(ctx, req.(*RecordListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Record_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.Record/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).Cancel(ctx, req.(*CancelRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Record_Finish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).Finish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.Record/Finish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).Finish(ctx, req.(*FinishRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Record_ApplyAfterSales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyAfterSalesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).ApplyAfterSales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.Record/ApplyAfterSales",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).ApplyAfterSales(ctx, req.(*ApplyAfterSalesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Record_Memo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemoRepuest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).Memo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.Record/Memo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).Memo(ctx, req.(*MemoRepuest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Record_serviceDesc = grpc.ServiceDesc{
	ServiceName: "interact.Record",
	HandlerType: (*RecordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Record_List_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Record_Cancel_Handler,
		},
		{
			MethodName: "Finish",
			Handler:    _Record_Finish_Handler,
		},
		{
			MethodName: "ApplyAfterSales",
			Handler:    _Record_ApplyAfterSales_Handler,
		},
		{
			MethodName: "Memo",
			Handler:    _Record_Memo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "record.proto",
}
