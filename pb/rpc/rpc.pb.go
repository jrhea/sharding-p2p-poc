// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package proto_rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import message "github.com/ethresearch/sharding-p2p-poc/pb/message"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Response_Status int32

const (
	Response_SUCCESS Response_Status = 0
	Response_FAILURE Response_Status = 1
)

var Response_Status_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILURE",
}
var Response_Status_value = map[string]int32{
	"SUCCESS": 0,
	"FAILURE": 1,
}

func (x Response_Status) String() string {
	return proto.EnumName(Response_Status_name, int32(x))
}
func (Response_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rpc_8f40528141eba4bc, []int{7, 0}
}

// Request
type RPCAddPeerRequest struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int64    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Seed                 int64    `protobuf:"varint,3,opt,name=seed,proto3" json:"seed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCAddPeerRequest) Reset()         { *m = RPCAddPeerRequest{} }
func (m *RPCAddPeerRequest) String() string { return proto.CompactTextString(m) }
func (*RPCAddPeerRequest) ProtoMessage()    {}
func (*RPCAddPeerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_8f40528141eba4bc, []int{0}
}
func (m *RPCAddPeerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCAddPeerRequest.Unmarshal(m, b)
}
func (m *RPCAddPeerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCAddPeerRequest.Marshal(b, m, deterministic)
}
func (dst *RPCAddPeerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCAddPeerRequest.Merge(dst, src)
}
func (m *RPCAddPeerRequest) XXX_Size() int {
	return xxx_messageInfo_RPCAddPeerRequest.Size(m)
}
func (m *RPCAddPeerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCAddPeerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RPCAddPeerRequest proto.InternalMessageInfo

func (m *RPCAddPeerRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *RPCAddPeerRequest) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *RPCAddPeerRequest) GetSeed() int64 {
	if m != nil {
		return m.Seed
	}
	return 0
}

type RPCSubscribeShardRequest struct {
	ShardIDs             []int64  `protobuf:"varint,1,rep,packed,name=shardIDs,proto3" json:"shardIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCSubscribeShardRequest) Reset()         { *m = RPCSubscribeShardRequest{} }
func (m *RPCSubscribeShardRequest) String() string { return proto.CompactTextString(m) }
func (*RPCSubscribeShardRequest) ProtoMessage()    {}
func (*RPCSubscribeShardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_8f40528141eba4bc, []int{1}
}
func (m *RPCSubscribeShardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCSubscribeShardRequest.Unmarshal(m, b)
}
func (m *RPCSubscribeShardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCSubscribeShardRequest.Marshal(b, m, deterministic)
}
func (dst *RPCSubscribeShardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCSubscribeShardRequest.Merge(dst, src)
}
func (m *RPCSubscribeShardRequest) XXX_Size() int {
	return xxx_messageInfo_RPCSubscribeShardRequest.Size(m)
}
func (m *RPCSubscribeShardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCSubscribeShardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RPCSubscribeShardRequest proto.InternalMessageInfo

func (m *RPCSubscribeShardRequest) GetShardIDs() []int64 {
	if m != nil {
		return m.ShardIDs
	}
	return nil
}

type RPCUnsubscribeShardRequest struct {
	ShardIDs             []int64  `protobuf:"varint,1,rep,packed,name=shardIDs,proto3" json:"shardIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCUnsubscribeShardRequest) Reset()         { *m = RPCUnsubscribeShardRequest{} }
func (m *RPCUnsubscribeShardRequest) String() string { return proto.CompactTextString(m) }
func (*RPCUnsubscribeShardRequest) ProtoMessage()    {}
func (*RPCUnsubscribeShardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_8f40528141eba4bc, []int{2}
}
func (m *RPCUnsubscribeShardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCUnsubscribeShardRequest.Unmarshal(m, b)
}
func (m *RPCUnsubscribeShardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCUnsubscribeShardRequest.Marshal(b, m, deterministic)
}
func (dst *RPCUnsubscribeShardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCUnsubscribeShardRequest.Merge(dst, src)
}
func (m *RPCUnsubscribeShardRequest) XXX_Size() int {
	return xxx_messageInfo_RPCUnsubscribeShardRequest.Size(m)
}
func (m *RPCUnsubscribeShardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCUnsubscribeShardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RPCUnsubscribeShardRequest proto.InternalMessageInfo

func (m *RPCUnsubscribeShardRequest) GetShardIDs() []int64 {
	if m != nil {
		return m.ShardIDs
	}
	return nil
}

type RPCGetSubscribedShardRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCGetSubscribedShardRequest) Reset()         { *m = RPCGetSubscribedShardRequest{} }
func (m *RPCGetSubscribedShardRequest) String() string { return proto.CompactTextString(m) }
func (*RPCGetSubscribedShardRequest) ProtoMessage()    {}
func (*RPCGetSubscribedShardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_8f40528141eba4bc, []int{3}
}
func (m *RPCGetSubscribedShardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCGetSubscribedShardRequest.Unmarshal(m, b)
}
func (m *RPCGetSubscribedShardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCGetSubscribedShardRequest.Marshal(b, m, deterministic)
}
func (dst *RPCGetSubscribedShardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCGetSubscribedShardRequest.Merge(dst, src)
}
func (m *RPCGetSubscribedShardRequest) XXX_Size() int {
	return xxx_messageInfo_RPCGetSubscribedShardRequest.Size(m)
}
func (m *RPCGetSubscribedShardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCGetSubscribedShardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RPCGetSubscribedShardRequest proto.InternalMessageInfo

type RPCBroadcastCollationRequest struct {
	ShardID              int64    `protobuf:"varint,1,opt,name=shardID,proto3" json:"shardID,omitempty"`
	Number               int64    `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Size                 int64    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Period               int64    `protobuf:"varint,4,opt,name=period,proto3" json:"period,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCBroadcastCollationRequest) Reset()         { *m = RPCBroadcastCollationRequest{} }
func (m *RPCBroadcastCollationRequest) String() string { return proto.CompactTextString(m) }
func (*RPCBroadcastCollationRequest) ProtoMessage()    {}
func (*RPCBroadcastCollationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_8f40528141eba4bc, []int{4}
}
func (m *RPCBroadcastCollationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCBroadcastCollationRequest.Unmarshal(m, b)
}
func (m *RPCBroadcastCollationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCBroadcastCollationRequest.Marshal(b, m, deterministic)
}
func (dst *RPCBroadcastCollationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCBroadcastCollationRequest.Merge(dst, src)
}
func (m *RPCBroadcastCollationRequest) XXX_Size() int {
	return xxx_messageInfo_RPCBroadcastCollationRequest.Size(m)
}
func (m *RPCBroadcastCollationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCBroadcastCollationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RPCBroadcastCollationRequest proto.InternalMessageInfo

func (m *RPCBroadcastCollationRequest) GetShardID() int64 {
	if m != nil {
		return m.ShardID
	}
	return 0
}

func (m *RPCBroadcastCollationRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *RPCBroadcastCollationRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *RPCBroadcastCollationRequest) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

type RPCSendCollationRequest struct {
	Collation            *message.Collation `protobuf:"bytes,1,opt,name=collation,proto3" json:"collation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RPCSendCollationRequest) Reset()         { *m = RPCSendCollationRequest{} }
func (m *RPCSendCollationRequest) String() string { return proto.CompactTextString(m) }
func (*RPCSendCollationRequest) ProtoMessage()    {}
func (*RPCSendCollationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_8f40528141eba4bc, []int{5}
}
func (m *RPCSendCollationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCSendCollationRequest.Unmarshal(m, b)
}
func (m *RPCSendCollationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCSendCollationRequest.Marshal(b, m, deterministic)
}
func (dst *RPCSendCollationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCSendCollationRequest.Merge(dst, src)
}
func (m *RPCSendCollationRequest) XXX_Size() int {
	return xxx_messageInfo_RPCSendCollationRequest.Size(m)
}
func (m *RPCSendCollationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCSendCollationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RPCSendCollationRequest proto.InternalMessageInfo

func (m *RPCSendCollationRequest) GetCollation() *message.Collation {
	if m != nil {
		return m.Collation
	}
	return nil
}

type RPCStopServerRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCStopServerRequest) Reset()         { *m = RPCStopServerRequest{} }
func (m *RPCStopServerRequest) String() string { return proto.CompactTextString(m) }
func (*RPCStopServerRequest) ProtoMessage()    {}
func (*RPCStopServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_8f40528141eba4bc, []int{6}
}
func (m *RPCStopServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCStopServerRequest.Unmarshal(m, b)
}
func (m *RPCStopServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCStopServerRequest.Marshal(b, m, deterministic)
}
func (dst *RPCStopServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCStopServerRequest.Merge(dst, src)
}
func (m *RPCStopServerRequest) XXX_Size() int {
	return xxx_messageInfo_RPCStopServerRequest.Size(m)
}
func (m *RPCStopServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCStopServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RPCStopServerRequest proto.InternalMessageInfo

// Response
type Response struct {
	Status               Response_Status `protobuf:"varint,1,opt,name=status,proto3,enum=proto.rpc.Response_Status" json:"status,omitempty"`
	Message              string          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_8f40528141eba4bc, []int{7}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() Response_Status {
	if m != nil {
		return m.Status
	}
	return Response_SUCCESS
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RPCPlainResponse struct {
	Response             *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RPCPlainResponse) Reset()         { *m = RPCPlainResponse{} }
func (m *RPCPlainResponse) String() string { return proto.CompactTextString(m) }
func (*RPCPlainResponse) ProtoMessage()    {}
func (*RPCPlainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_8f40528141eba4bc, []int{8}
}
func (m *RPCPlainResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCPlainResponse.Unmarshal(m, b)
}
func (m *RPCPlainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCPlainResponse.Marshal(b, m, deterministic)
}
func (dst *RPCPlainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCPlainResponse.Merge(dst, src)
}
func (m *RPCPlainResponse) XXX_Size() int {
	return xxx_messageInfo_RPCPlainResponse.Size(m)
}
func (m *RPCPlainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCPlainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RPCPlainResponse proto.InternalMessageInfo

func (m *RPCPlainResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type RPCGetSubscribedShardResponse struct {
	Response             *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	ShardIDs             []int64   `protobuf:"varint,2,rep,packed,name=shardIDs,proto3" json:"shardIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RPCGetSubscribedShardResponse) Reset()         { *m = RPCGetSubscribedShardResponse{} }
func (m *RPCGetSubscribedShardResponse) String() string { return proto.CompactTextString(m) }
func (*RPCGetSubscribedShardResponse) ProtoMessage()    {}
func (*RPCGetSubscribedShardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_8f40528141eba4bc, []int{9}
}
func (m *RPCGetSubscribedShardResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCGetSubscribedShardResponse.Unmarshal(m, b)
}
func (m *RPCGetSubscribedShardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCGetSubscribedShardResponse.Marshal(b, m, deterministic)
}
func (dst *RPCGetSubscribedShardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCGetSubscribedShardResponse.Merge(dst, src)
}
func (m *RPCGetSubscribedShardResponse) XXX_Size() int {
	return xxx_messageInfo_RPCGetSubscribedShardResponse.Size(m)
}
func (m *RPCGetSubscribedShardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCGetSubscribedShardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RPCGetSubscribedShardResponse proto.InternalMessageInfo

func (m *RPCGetSubscribedShardResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *RPCGetSubscribedShardResponse) GetShardIDs() []int64 {
	if m != nil {
		return m.ShardIDs
	}
	return nil
}

type SendRequest struct {
	PeerID               string   `protobuf:"bytes,1,opt,name=peerID,proto3" json:"peerID,omitempty"`
	Topic                string   `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	MsgType              int64    `protobuf:"varint,3,opt,name=msgType,proto3" json:"msgType,omitempty"`
	Data                 []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_8f40528141eba4bc, []int{10}
}
func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (dst *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(dst, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

func (m *SendRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *SendRequest) GetMsgType() int64 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *SendRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SendResponse struct {
	Response             *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Data                 []byte    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_8f40528141eba4bc, []int{11}
}
func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (dst *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(dst, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

func (m *SendResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *SendResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*RPCAddPeerRequest)(nil), "proto.rpc.RPCAddPeerRequest")
	proto.RegisterType((*RPCSubscribeShardRequest)(nil), "proto.rpc.RPCSubscribeShardRequest")
	proto.RegisterType((*RPCUnsubscribeShardRequest)(nil), "proto.rpc.RPCUnsubscribeShardRequest")
	proto.RegisterType((*RPCGetSubscribedShardRequest)(nil), "proto.rpc.RPCGetSubscribedShardRequest")
	proto.RegisterType((*RPCBroadcastCollationRequest)(nil), "proto.rpc.RPCBroadcastCollationRequest")
	proto.RegisterType((*RPCSendCollationRequest)(nil), "proto.rpc.RPCSendCollationRequest")
	proto.RegisterType((*RPCStopServerRequest)(nil), "proto.rpc.RPCStopServerRequest")
	proto.RegisterType((*Response)(nil), "proto.rpc.Response")
	proto.RegisterType((*RPCPlainResponse)(nil), "proto.rpc.RPCPlainResponse")
	proto.RegisterType((*RPCGetSubscribedShardResponse)(nil), "proto.rpc.RPCGetSubscribedShardResponse")
	proto.RegisterType((*SendRequest)(nil), "proto.rpc.SendRequest")
	proto.RegisterType((*SendResponse)(nil), "proto.rpc.SendResponse")
	proto.RegisterEnum("proto.rpc.Response_Status", Response_Status_name, Response_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PocClient is the client API for Poc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PocClient interface {
	AddPeer(ctx context.Context, in *RPCAddPeerRequest, opts ...grpc.CallOption) (*RPCPlainResponse, error)
	SubscribeShard(ctx context.Context, in *RPCSubscribeShardRequest, opts ...grpc.CallOption) (*RPCPlainResponse, error)
	UnsubscribeShard(ctx context.Context, in *RPCUnsubscribeShardRequest, opts ...grpc.CallOption) (*RPCPlainResponse, error)
	GetSubscribedShard(ctx context.Context, in *RPCGetSubscribedShardRequest, opts ...grpc.CallOption) (*RPCGetSubscribedShardResponse, error)
	BroadcastCollation(ctx context.Context, in *RPCBroadcastCollationRequest, opts ...grpc.CallOption) (*RPCPlainResponse, error)
	SendCollation(ctx context.Context, in *RPCSendCollationRequest, opts ...grpc.CallOption) (*RPCPlainResponse, error)
	StopServer(ctx context.Context, in *RPCStopServerRequest, opts ...grpc.CallOption) (*RPCPlainResponse, error)
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
}

type pocClient struct {
	cc *grpc.ClientConn
}

func NewPocClient(cc *grpc.ClientConn) PocClient {
	return &pocClient{cc}
}

func (c *pocClient) AddPeer(ctx context.Context, in *RPCAddPeerRequest, opts ...grpc.CallOption) (*RPCPlainResponse, error) {
	out := new(RPCPlainResponse)
	err := c.cc.Invoke(ctx, "/proto.rpc.Poc/AddPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pocClient) SubscribeShard(ctx context.Context, in *RPCSubscribeShardRequest, opts ...grpc.CallOption) (*RPCPlainResponse, error) {
	out := new(RPCPlainResponse)
	err := c.cc.Invoke(ctx, "/proto.rpc.Poc/SubscribeShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pocClient) UnsubscribeShard(ctx context.Context, in *RPCUnsubscribeShardRequest, opts ...grpc.CallOption) (*RPCPlainResponse, error) {
	out := new(RPCPlainResponse)
	err := c.cc.Invoke(ctx, "/proto.rpc.Poc/UnsubscribeShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pocClient) GetSubscribedShard(ctx context.Context, in *RPCGetSubscribedShardRequest, opts ...grpc.CallOption) (*RPCGetSubscribedShardResponse, error) {
	out := new(RPCGetSubscribedShardResponse)
	err := c.cc.Invoke(ctx, "/proto.rpc.Poc/GetSubscribedShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pocClient) BroadcastCollation(ctx context.Context, in *RPCBroadcastCollationRequest, opts ...grpc.CallOption) (*RPCPlainResponse, error) {
	out := new(RPCPlainResponse)
	err := c.cc.Invoke(ctx, "/proto.rpc.Poc/BroadcastCollation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pocClient) SendCollation(ctx context.Context, in *RPCSendCollationRequest, opts ...grpc.CallOption) (*RPCPlainResponse, error) {
	out := new(RPCPlainResponse)
	err := c.cc.Invoke(ctx, "/proto.rpc.Poc/SendCollation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pocClient) StopServer(ctx context.Context, in *RPCStopServerRequest, opts ...grpc.CallOption) (*RPCPlainResponse, error) {
	out := new(RPCPlainResponse)
	err := c.cc.Invoke(ctx, "/proto.rpc.Poc/StopServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pocClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/proto.rpc.Poc/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PocServer is the server API for Poc service.
type PocServer interface {
	AddPeer(context.Context, *RPCAddPeerRequest) (*RPCPlainResponse, error)
	SubscribeShard(context.Context, *RPCSubscribeShardRequest) (*RPCPlainResponse, error)
	UnsubscribeShard(context.Context, *RPCUnsubscribeShardRequest) (*RPCPlainResponse, error)
	GetSubscribedShard(context.Context, *RPCGetSubscribedShardRequest) (*RPCGetSubscribedShardResponse, error)
	BroadcastCollation(context.Context, *RPCBroadcastCollationRequest) (*RPCPlainResponse, error)
	SendCollation(context.Context, *RPCSendCollationRequest) (*RPCPlainResponse, error)
	StopServer(context.Context, *RPCStopServerRequest) (*RPCPlainResponse, error)
	Send(context.Context, *SendRequest) (*SendResponse, error)
}

func RegisterPocServer(s *grpc.Server, srv PocServer) {
	s.RegisterService(&_Poc_serviceDesc, srv)
}

func _Poc_AddPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCAddPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).AddPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.rpc.Poc/AddPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).AddPeer(ctx, req.(*RPCAddPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poc_SubscribeShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCSubscribeShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).SubscribeShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.rpc.Poc/SubscribeShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).SubscribeShard(ctx, req.(*RPCSubscribeShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poc_UnsubscribeShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCUnsubscribeShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).UnsubscribeShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.rpc.Poc/UnsubscribeShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).UnsubscribeShard(ctx, req.(*RPCUnsubscribeShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poc_GetSubscribedShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCGetSubscribedShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).GetSubscribedShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.rpc.Poc/GetSubscribedShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).GetSubscribedShard(ctx, req.(*RPCGetSubscribedShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poc_BroadcastCollation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCBroadcastCollationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).BroadcastCollation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.rpc.Poc/BroadcastCollation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).BroadcastCollation(ctx, req.(*RPCBroadcastCollationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poc_SendCollation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCSendCollationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).SendCollation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.rpc.Poc/SendCollation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).SendCollation(ctx, req.(*RPCSendCollationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poc_StopServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCStopServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).StopServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.rpc.Poc/StopServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).StopServer(ctx, req.(*RPCStopServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poc_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.rpc.Poc/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Poc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.rpc.Poc",
	HandlerType: (*PocServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPeer",
			Handler:    _Poc_AddPeer_Handler,
		},
		{
			MethodName: "SubscribeShard",
			Handler:    _Poc_SubscribeShard_Handler,
		},
		{
			MethodName: "UnsubscribeShard",
			Handler:    _Poc_UnsubscribeShard_Handler,
		},
		{
			MethodName: "GetSubscribedShard",
			Handler:    _Poc_GetSubscribedShard_Handler,
		},
		{
			MethodName: "BroadcastCollation",
			Handler:    _Poc_BroadcastCollation_Handler,
		},
		{
			MethodName: "SendCollation",
			Handler:    _Poc_SendCollation_Handler,
		},
		{
			MethodName: "StopServer",
			Handler:    _Poc_StopServer_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Poc_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_rpc_8f40528141eba4bc) }

var fileDescriptor_rpc_8f40528141eba4bc = []byte{
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x4f, 0xdb, 0x3c,
	0x14, 0xa6, 0x2d, 0x14, 0x7a, 0xe0, 0x45, 0x7d, 0x3d, 0x04, 0x51, 0xc6, 0x36, 0xe4, 0x69, 0x5a,
	0x6f, 0x48, 0xa5, 0x4e, 0x42, 0xdb, 0xdd, 0x58, 0x06, 0x13, 0x1a, 0x17, 0x99, 0x03, 0xd2, 0x2e,
	0x76, 0x93, 0x26, 0x56, 0x6b, 0xa9, 0xc4, 0x9e, 0xed, 0x4e, 0x1a, 0xe2, 0xdf, 0xec, 0x8f, 0x4e,
	0x71, 0xe2, 0xb4, 0x49, 0xa9, 0xa2, 0x71, 0x65, 0x1f, 0xfb, 0x39, 0xcf, 0xf9, 0x3e, 0xd0, 0x93,
	0x22, 0xf6, 0x84, 0xe4, 0x9a, 0xa3, 0x9e, 0x39, 0x3c, 0x29, 0x62, 0xf7, 0xe3, 0x84, 0xe9, 0xe9,
	0x7c, 0xec, 0xc5, 0xfc, 0x6e, 0x48, 0xf5, 0x54, 0x52, 0x45, 0x23, 0x19, 0x4f, 0x87, 0x6a, 0x1a,
	0xc9, 0x84, 0xa5, 0x93, 0x53, 0x31, 0x12, 0xa7, 0x82, 0xc7, 0x43, 0x31, 0x1e, 0xde, 0x51, 0xa5,
	0xa2, 0x09, 0xb5, 0x67, 0x4e, 0x86, 0xbf, 0xc2, 0xff, 0x24, 0xf0, 0xcf, 0x93, 0x24, 0xa0, 0x54,
	0x12, 0xfa, 0x73, 0x4e, 0x95, 0x46, 0xfb, 0xd0, 0x66, 0xc2, 0x69, 0x9d, 0xb4, 0x06, 0x3d, 0xd2,
	0x66, 0x02, 0x21, 0xd8, 0x14, 0x5c, 0x6a, 0xa7, 0x7d, 0xd2, 0x1a, 0x74, 0x88, 0xb9, 0x67, 0x6f,
	0x8a, 0xd2, 0xc4, 0xe9, 0xe4, 0x6f, 0xd9, 0x1d, 0x9f, 0x81, 0x43, 0x02, 0x3f, 0x9c, 0x8f, 0x55,
	0x2c, 0xd9, 0x98, 0x86, 0x99, 0x17, 0x96, 0xd3, 0x85, 0x1d, 0xe3, 0xd5, 0xd5, 0x67, 0xe5, 0xb4,
	0x4e, 0x3a, 0x83, 0x0e, 0x29, 0x65, 0xfc, 0x1e, 0x5c, 0x12, 0xf8, 0xb7, 0xa9, 0xfa, 0x67, 0xcd,
	0x97, 0x70, 0x4c, 0x02, 0xff, 0x0b, 0xd5, 0xa5, 0xd1, 0x64, 0x59, 0x17, 0x3f, 0x98, 0xff, 0x4f,
	0x92, 0x47, 0x49, 0x1c, 0x29, 0xed, 0xf3, 0xd9, 0x2c, 0xd2, 0x8c, 0xa7, 0x96, 0xdb, 0x81, 0xed,
	0x82, 0xcb, 0x84, 0xdb, 0x21, 0x56, 0x44, 0x87, 0xd0, 0x4d, 0xe7, 0x77, 0x63, 0x2a, 0x8b, 0xa8,
	0x0b, 0xc9, 0xc4, 0xcd, 0xee, 0x69, 0x19, 0x37, 0xbb, 0xa7, 0x19, 0x56, 0x50, 0xc9, 0x78, 0xe2,
	0x6c, 0xe6, 0xd8, 0x5c, 0xc2, 0xdf, 0xe0, 0x28, 0xcb, 0x07, 0x4d, 0x93, 0x15, 0xc3, 0x67, 0xd0,
	0x8b, 0xed, 0x9b, 0x31, 0xbd, 0x3b, 0x72, 0xf2, 0x92, 0x78, 0xb6, 0x40, 0x0b, 0x9d, 0x05, 0x14,
	0x1f, 0xc2, 0x41, 0x46, 0xa9, 0xb9, 0x08, 0xa9, 0xfc, 0x55, 0x96, 0x0c, 0x3f, 0xc0, 0x0e, 0xa1,
	0x4a, 0xf0, 0x54, 0x51, 0x34, 0x82, 0xae, 0xd2, 0x91, 0x9e, 0x2b, 0x43, 0xbc, 0x3f, 0x72, 0xbd,
	0xb2, 0x63, 0x3c, 0x0b, 0xf2, 0x42, 0x83, 0x20, 0x05, 0x32, 0x4b, 0x44, 0x61, 0xd7, 0xc4, 0xdb,
	0x23, 0x56, 0xc4, 0x18, 0xba, 0x39, 0x16, 0xed, 0xc2, 0x76, 0x78, 0xeb, 0xfb, 0x17, 0x61, 0xd8,
	0xdf, 0xc8, 0x84, 0xcb, 0xf3, 0xab, 0xeb, 0x5b, 0x72, 0xd1, 0x6f, 0x61, 0x1f, 0xfa, 0x24, 0xf0,
	0x83, 0x59, 0xc4, 0xd2, 0xd2, 0x8b, 0x21, 0xec, 0xc8, 0xe2, 0x5e, 0x04, 0xf8, 0xec, 0x11, 0x3f,
	0x48, 0x09, 0xc2, 0x33, 0x78, 0xb1, 0xa6, 0x96, 0x4f, 0x64, 0xac, 0x74, 0x4e, 0xbb, 0xd6, 0x39,
	0x0c, 0x76, 0xb3, 0xc2, 0xd8, 0x7a, 0x98, 0x12, 0x52, 0x59, 0xf4, 0x41, 0x8f, 0x14, 0x12, 0x3a,
	0x80, 0x2d, 0xcd, 0x05, 0x8b, 0x8b, 0xac, 0xe4, 0x82, 0xc9, 0x96, 0x9a, 0xdc, 0xfc, 0x16, 0xb6,
	0x0f, 0xac, 0x98, 0xb5, 0x47, 0x12, 0xe9, 0xc8, 0x34, 0xc2, 0x1e, 0x31, 0x77, 0x1c, 0xc2, 0x5e,
	0x6e, 0xea, 0xa9, 0x71, 0x58, 0xd2, 0xf6, 0x82, 0x74, 0xf4, 0x67, 0x0b, 0x3a, 0x01, 0x8f, 0xd1,
	0x25, 0x6c, 0x17, 0xd3, 0x8b, 0x8e, 0x97, 0x59, 0xea, 0x43, 0xed, 0x3e, 0xaf, 0xfe, 0x56, 0x8a,
	0x85, 0x37, 0xd0, 0x0d, 0xec, 0x57, 0x07, 0x17, 0xbd, 0xae, 0x2a, 0x3c, 0x3a, 0xd6, 0x4d, 0xac,
	0xdf, 0xa1, 0x5f, 0x1f, 0x6b, 0xf4, 0xa6, 0xaa, 0xb2, 0x66, 0xec, 0x9b, 0x98, 0x19, 0xa0, 0xd5,
	0x56, 0x41, 0x6f, 0xab, 0x4a, 0x6b, 0x17, 0x83, 0x3b, 0x68, 0x06, 0x96, 0xa6, 0x7e, 0x00, 0x5a,
	0xdd, 0x20, 0x75, 0x53, 0x6b, 0x77, 0x4c, 0x53, 0x20, 0x04, 0xfe, 0xab, 0x6c, 0x08, 0x84, 0x6b,
	0x79, 0x7f, 0x64, 0x7d, 0x34, 0x71, 0x5e, 0x03, 0x2c, 0x56, 0x04, 0x7a, 0x55, 0x23, 0xac, 0x2f,
	0x8f, 0x26, 0xb6, 0x0f, 0xb0, 0x99, 0x39, 0x81, 0x0e, 0x97, 0x60, 0x4b, 0xb3, 0xe3, 0x1e, 0xad,
	0xbc, 0x5b, 0xd5, 0x71, 0xd7, 0xfc, 0xbc, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x64, 0x44,
	0x49, 0xbf, 0x06, 0x00, 0x00,
}
