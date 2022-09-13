// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: channel/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCommitment struct {
	Creator     string      `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	From        string      `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	CoinA       *types.Coin `protobuf:"bytes,3,opt,name=coinA,proto3" json:"coinA,omitempty"`
	ToATimelock string      `protobuf:"bytes,4,opt,name=toATimelock,proto3" json:"toATimelock,omitempty"`
	Blockheight uint64      `protobuf:"varint,5,opt,name=blockheight,proto3" json:"blockheight,omitempty"`
	ToBHashlock string      `protobuf:"bytes,6,opt,name=toBHashlock,proto3" json:"toBHashlock,omitempty"`
	Hashcode    string      `protobuf:"bytes,7,opt,name=hashcode,proto3" json:"hashcode,omitempty"`
	Coinlock    *types.Coin `protobuf:"bytes,8,opt,name=coinlock,proto3" json:"coinlock,omitempty"`
}

func (m *MsgCommitment) Reset()         { *m = MsgCommitment{} }
func (m *MsgCommitment) String() string { return proto.CompactTextString(m) }
func (*MsgCommitment) ProtoMessage()    {}
func (*MsgCommitment) Descriptor() ([]byte, []int) {
	return fileDescriptor_82d382f6faba5dbf, []int{0}
}
func (m *MsgCommitment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCommitment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCommitment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCommitment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCommitment.Merge(m, src)
}
func (m *MsgCommitment) XXX_Size() int {
	return m.Size()
}
func (m *MsgCommitment) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCommitment.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCommitment proto.InternalMessageInfo

func (m *MsgCommitment) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCommitment) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *MsgCommitment) GetCoinA() *types.Coin {
	if m != nil {
		return m.CoinA
	}
	return nil
}

func (m *MsgCommitment) GetToATimelock() string {
	if m != nil {
		return m.ToATimelock
	}
	return ""
}

func (m *MsgCommitment) GetBlockheight() uint64 {
	if m != nil {
		return m.Blockheight
	}
	return 0
}

func (m *MsgCommitment) GetToBHashlock() string {
	if m != nil {
		return m.ToBHashlock
	}
	return ""
}

func (m *MsgCommitment) GetHashcode() string {
	if m != nil {
		return m.Hashcode
	}
	return ""
}

func (m *MsgCommitment) GetCoinlock() *types.Coin {
	if m != nil {
		return m.Coinlock
	}
	return nil
}

type MsgCommitmentResponse struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *MsgCommitmentResponse) Reset()         { *m = MsgCommitmentResponse{} }
func (m *MsgCommitmentResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCommitmentResponse) ProtoMessage()    {}
func (*MsgCommitmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82d382f6faba5dbf, []int{1}
}
func (m *MsgCommitmentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCommitmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCommitmentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCommitmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCommitmentResponse.Merge(m, src)
}
func (m *MsgCommitmentResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCommitmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCommitmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCommitmentResponse proto.InternalMessageInfo

func (m *MsgCommitmentResponse) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type MsgWithdrawTimelock struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	To      string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Index   string `protobuf:"bytes,3,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *MsgWithdrawTimelock) Reset()         { *m = MsgWithdrawTimelock{} }
func (m *MsgWithdrawTimelock) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawTimelock) ProtoMessage()    {}
func (*MsgWithdrawTimelock) Descriptor() ([]byte, []int) {
	return fileDescriptor_82d382f6faba5dbf, []int{2}
}
func (m *MsgWithdrawTimelock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawTimelock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawTimelock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawTimelock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawTimelock.Merge(m, src)
}
func (m *MsgWithdrawTimelock) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawTimelock) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawTimelock.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawTimelock proto.InternalMessageInfo

func (m *MsgWithdrawTimelock) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgWithdrawTimelock) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *MsgWithdrawTimelock) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type MsgWithdrawTimelockResponse struct {
}

func (m *MsgWithdrawTimelockResponse) Reset()         { *m = MsgWithdrawTimelockResponse{} }
func (m *MsgWithdrawTimelockResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawTimelockResponse) ProtoMessage()    {}
func (*MsgWithdrawTimelockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82d382f6faba5dbf, []int{3}
}
func (m *MsgWithdrawTimelockResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawTimelockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawTimelockResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawTimelockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawTimelockResponse.Merge(m, src)
}
func (m *MsgWithdrawTimelockResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawTimelockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawTimelockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawTimelockResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCommitment)(nil), "channel.channel.MsgCommitment")
	proto.RegisterType((*MsgCommitmentResponse)(nil), "channel.channel.MsgCommitmentResponse")
	proto.RegisterType((*MsgWithdrawTimelock)(nil), "channel.channel.MsgWithdrawTimelock")
	proto.RegisterType((*MsgWithdrawTimelockResponse)(nil), "channel.channel.MsgWithdrawTimelockResponse")
}

func init() { proto.RegisterFile("channel/tx.proto", fileDescriptor_82d382f6faba5dbf) }

var fileDescriptor_82d382f6faba5dbf = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcb, 0x6e, 0xda, 0x40,
	0x14, 0x65, 0xcc, 0xb3, 0x17, 0xb5, 0x45, 0xd3, 0x56, 0x75, 0x5d, 0x75, 0x64, 0xa1, 0xaa, 0x62,
	0xd1, 0x8e, 0x05, 0x55, 0x3e, 0x00, 0xd8, 0x64, 0xc3, 0xc6, 0x22, 0x8a, 0x94, 0x9d, 0x6d, 0x06,
	0xdb, 0x0a, 0xf6, 0x20, 0xcf, 0x28, 0x21, 0x7f, 0x91, 0xbf, 0xc8, 0x87, 0x64, 0x93, 0x25, 0xcb,
	0x2c, 0x23, 0xf8, 0x91, 0xc8, 0xcf, 0x18, 0x42, 0x1e, 0x2b, 0xcf, 0x3d, 0xf7, 0xdc, 0x73, 0x1f,
	0x3e, 0xd0, 0x71, 0x3c, 0x2b, 0x0c, 0xd9, 0xc2, 0x90, 0x2b, 0xba, 0x8c, 0xb8, 0xe4, 0xf8, 0x73,
	0x86, 0xd0, 0xec, 0xab, 0x11, 0x87, 0x8b, 0x80, 0x0b, 0xc3, 0xb6, 0x04, 0x33, 0x2e, 0xfa, 0x36,
	0x93, 0x56, 0xdf, 0x70, 0xb8, 0x1f, 0xa6, 0x05, 0xdd, 0x1b, 0x05, 0x3e, 0x4e, 0x84, 0x3b, 0xe6,
	0x41, 0xe0, 0xcb, 0x80, 0x85, 0x12, 0xab, 0xd0, 0x74, 0x22, 0x66, 0x49, 0x1e, 0xa9, 0x48, 0x47,
	0xbd, 0x0f, 0x66, 0x1e, 0x62, 0x0c, 0xb5, 0x79, 0xc4, 0x03, 0x55, 0x49, 0xe0, 0xe4, 0x8d, 0x0d,
	0xa8, 0xc7, 0x6a, 0x43, 0xb5, 0xaa, 0xa3, 0x5e, 0x7b, 0xf0, 0x83, 0xa6, 0xfd, 0x68, 0xdc, 0x8f,
	0x66, 0xfd, 0xe8, 0x98, 0xfb, 0xa1, 0x99, 0xf2, 0xb0, 0x0e, 0x6d, 0xc9, 0x87, 0x53, 0x3f, 0x60,
	0x0b, 0xee, 0x9c, 0xab, 0xb5, 0x44, 0xab, 0x0c, 0xc5, 0x0c, 0x3b, 0x7e, 0x78, 0xcc, 0x77, 0x3d,
	0xa9, 0xd6, 0x75, 0xd4, 0xab, 0x99, 0x65, 0x28, 0xd5, 0x18, 0x1d, 0x5b, 0xc2, 0x4b, 0x34, 0x1a,
	0xb9, 0x46, 0x01, 0x61, 0x0d, 0x5a, 0x9e, 0x25, 0x3c, 0x87, 0xcf, 0x98, 0xda, 0x4c, 0xd2, 0x45,
	0x8c, 0x8f, 0xa0, 0x15, 0x8f, 0x92, 0x94, 0xb6, 0xde, 0x9a, 0xba, 0xa0, 0x76, 0xff, 0xc1, 0xb7,
	0x9d, 0x43, 0x99, 0x4c, 0x2c, 0x79, 0x28, 0x18, 0xfe, 0x0a, 0x75, 0x3f, 0x9c, 0xb1, 0x55, 0x76,
	0xae, 0x34, 0xe8, 0x9e, 0xc0, 0x97, 0x89, 0x70, 0x4f, 0x7d, 0xe9, 0xcd, 0x22, 0xeb, 0xb2, 0x58,
	0xee, 0xe5, 0xeb, 0x7e, 0x02, 0x45, 0xf2, 0xec, 0xb6, 0x8a, 0xe4, 0x4f, 0xb2, 0xd5, 0xb2, 0xec,
	0x2f, 0xf8, 0x79, 0x40, 0x36, 0x9f, 0x65, 0x70, 0x8b, 0xa0, 0x3a, 0x11, 0x2e, 0x9e, 0x02, 0x94,
	0x7e, 0x29, 0xa1, 0x7b, 0xb6, 0xa0, 0x3b, 0x9b, 0x68, 0x7f, 0x5e, 0xcf, 0x17, 0x9b, 0xce, 0xa1,
	0xf3, 0x6c, 0xa1, 0xdf, 0x87, 0x6a, 0xf7, 0x59, 0xda, 0xdf, 0xf7, 0xb0, 0xf2, 0x3e, 0xa3, 0xfe,
	0xdd, 0x86, 0xa0, 0xf5, 0x86, 0xa0, 0x87, 0x0d, 0x41, 0xd7, 0x5b, 0x52, 0x59, 0x6f, 0x49, 0xe5,
	0x7e, 0x4b, 0x2a, 0x67, 0xdf, 0x73, 0xc7, 0xaf, 0x8c, 0xc2, 0xfb, 0x57, 0x4b, 0x26, 0xec, 0x46,
	0x62, 0xe7, 0xff, 0x8f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x73, 0x36, 0xc5, 0xf4, 0x13, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Commitment(ctx context.Context, in *MsgCommitment, opts ...grpc.CallOption) (*MsgCommitmentResponse, error)
	WithdrawTimelock(ctx context.Context, in *MsgWithdrawTimelock, opts ...grpc.CallOption) (*MsgWithdrawTimelockResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Commitment(ctx context.Context, in *MsgCommitment, opts ...grpc.CallOption) (*MsgCommitmentResponse, error) {
	out := new(MsgCommitmentResponse)
	err := c.cc.Invoke(ctx, "/channel.channel.Msg/Commitment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawTimelock(ctx context.Context, in *MsgWithdrawTimelock, opts ...grpc.CallOption) (*MsgWithdrawTimelockResponse, error) {
	out := new(MsgWithdrawTimelockResponse)
	err := c.cc.Invoke(ctx, "/channel.channel.Msg/WithdrawTimelock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Commitment(context.Context, *MsgCommitment) (*MsgCommitmentResponse, error)
	WithdrawTimelock(context.Context, *MsgWithdrawTimelock) (*MsgWithdrawTimelockResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Commitment(ctx context.Context, req *MsgCommitment) (*MsgCommitmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commitment not implemented")
}
func (*UnimplementedMsgServer) WithdrawTimelock(ctx context.Context, req *MsgWithdrawTimelock) (*MsgWithdrawTimelockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawTimelock not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Commitment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCommitment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Commitment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.channel.Msg/Commitment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Commitment(ctx, req.(*MsgCommitment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawTimelock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawTimelock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawTimelock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.channel.Msg/WithdrawTimelock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawTimelock(ctx, req.(*MsgWithdrawTimelock))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "channel.channel.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Commitment",
			Handler:    _Msg_Commitment_Handler,
		},
		{
			MethodName: "WithdrawTimelock",
			Handler:    _Msg_WithdrawTimelock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "channel/tx.proto",
}

func (m *MsgCommitment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCommitment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCommitment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Coinlock != nil {
		{
			size, err := m.Coinlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.Hashcode) > 0 {
		i -= len(m.Hashcode)
		copy(dAtA[i:], m.Hashcode)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Hashcode)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ToBHashlock) > 0 {
		i -= len(m.ToBHashlock)
		copy(dAtA[i:], m.ToBHashlock)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ToBHashlock)))
		i--
		dAtA[i] = 0x32
	}
	if m.Blockheight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Blockheight))
		i--
		dAtA[i] = 0x28
	}
	if len(m.ToATimelock) > 0 {
		i -= len(m.ToATimelock)
		copy(dAtA[i:], m.ToATimelock)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ToATimelock)))
		i--
		dAtA[i] = 0x22
	}
	if m.CoinA != nil {
		{
			size, err := m.CoinA.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCommitmentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCommitmentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCommitmentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawTimelock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawTimelock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawTimelock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintTx(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawTimelockResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawTimelockResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawTimelockResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCommitment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.CoinA != nil {
		l = m.CoinA.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ToATimelock)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Blockheight != 0 {
		n += 1 + sovTx(uint64(m.Blockheight))
	}
	l = len(m.ToBHashlock)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Hashcode)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Coinlock != nil {
		l = m.Coinlock.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCommitmentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgWithdrawTimelock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgWithdrawTimelockResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCommitment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCommitment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCommitment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinA", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CoinA == nil {
				m.CoinA = &types.Coin{}
			}
			if err := m.CoinA.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToATimelock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToATimelock = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blockheight", wireType)
			}
			m.Blockheight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Blockheight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToBHashlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToBHashlock = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hashcode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hashcode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coinlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Coinlock == nil {
				m.Coinlock = &types.Coin{}
			}
			if err := m.Coinlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCommitmentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCommitmentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCommitmentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgWithdrawTimelock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgWithdrawTimelock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawTimelock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgWithdrawTimelockResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgWithdrawTimelockResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawTimelockResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
