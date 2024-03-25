// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/perm/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgSetPermissionedRelayer defines msg to set permissioned relyer for
// the specific ibc channel.
type MsgSetPermissionedRelayer struct {
	// authority is the address that controls the module
	// (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty" yaml:"authority"`
	PortId    string `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty" yaml:"port_id"`
	ChannelId string `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty" yaml:"channel_id"`
	Relayer   string `protobuf:"bytes,4,opt,name=relayer,proto3" json:"relayer,omitempty" yaml:"relayer"`
}

func (m *MsgSetPermissionedRelayer) Reset()         { *m = MsgSetPermissionedRelayer{} }
func (m *MsgSetPermissionedRelayer) String() string { return proto.CompactTextString(m) }
func (*MsgSetPermissionedRelayer) ProtoMessage()    {}
func (*MsgSetPermissionedRelayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ea92a5af9136e08, []int{0}
}
func (m *MsgSetPermissionedRelayer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetPermissionedRelayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetPermissionedRelayer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetPermissionedRelayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetPermissionedRelayer.Merge(m, src)
}
func (m *MsgSetPermissionedRelayer) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetPermissionedRelayer) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetPermissionedRelayer.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetPermissionedRelayer proto.InternalMessageInfo

// MsgSetPermissionedRelayerResponse defines the Msg/SetPermissionedRelayer response type.
type MsgSetPermissionedRelayerResponse struct {
}

func (m *MsgSetPermissionedRelayerResponse) Reset()         { *m = MsgSetPermissionedRelayerResponse{} }
func (m *MsgSetPermissionedRelayerResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetPermissionedRelayerResponse) ProtoMessage()    {}
func (*MsgSetPermissionedRelayerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ea92a5af9136e08, []int{1}
}
func (m *MsgSetPermissionedRelayerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetPermissionedRelayerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetPermissionedRelayerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetPermissionedRelayerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetPermissionedRelayerResponse.Merge(m, src)
}
func (m *MsgSetPermissionedRelayerResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetPermissionedRelayerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetPermissionedRelayerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetPermissionedRelayerResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSetPermissionedRelayer)(nil), "ibc.applications.perm.v1.MsgSetPermissionedRelayer")
	proto.RegisterType((*MsgSetPermissionedRelayerResponse)(nil), "ibc.applications.perm.v1.MsgSetPermissionedRelayerResponse")
}

func init() { proto.RegisterFile("ibc/applications/perm/v1/tx.proto", fileDescriptor_9ea92a5af9136e08) }

var fileDescriptor_9ea92a5af9136e08 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x6f, 0xd4, 0x30,
	0x14, 0xc7, 0xe3, 0x16, 0x5a, 0x9d, 0x07, 0x44, 0xa3, 0x02, 0x69, 0x86, 0x84, 0x86, 0x05, 0x1d,
	0xad, 0xad, 0x52, 0x16, 0xca, 0x44, 0x27, 0x8a, 0x54, 0x09, 0xa5, 0x1b, 0x4b, 0xe5, 0x24, 0x96,
	0xcf, 0x52, 0x6c, 0x47, 0xb6, 0x5b, 0x35, 0x1b, 0x62, 0x42, 0x9d, 0x58, 0xd9, 0xfa, 0x11, 0x4e,
	0x82, 0x0f, 0xc1, 0x58, 0x31, 0x31, 0x9d, 0xd0, 0xdd, 0x70, 0xcc, 0xf7, 0x09, 0x50, 0xe2, 0x94,
	0x43, 0x48, 0x59, 0xba, 0x58, 0xef, 0xf9, 0xf7, 0xf7, 0xdf, 0x4f, 0xef, 0x3d, 0xb8, 0xcd, 0xb3,
	0x1c, 0x93, 0xaa, 0x2a, 0x79, 0x4e, 0x2c, 0x57, 0xd2, 0xe0, 0x8a, 0x6a, 0x81, 0xcf, 0xf7, 0xb0,
	0xbd, 0x40, 0x95, 0x56, 0x56, 0xf9, 0x01, 0xcf, 0x72, 0xf4, 0xaf, 0x04, 0x35, 0x12, 0x74, 0xbe,
	0x17, 0x6e, 0x10, 0xc1, 0xa5, 0xc2, 0xed, 0xe9, 0xc4, 0xe1, 0xa3, 0x5c, 0x19, 0xa1, 0x0c, 0x16,
	0x86, 0x35, 0x26, 0xc2, 0xb0, 0x0e, 0x6c, 0x39, 0x70, 0xda, 0x66, 0xd8, 0x25, 0x1d, 0xda, 0x64,
	0x8a, 0x29, 0x77, 0xdf, 0x44, 0xee, 0x36, 0xf9, 0xba, 0x02, 0xb7, 0x8e, 0x0d, 0x3b, 0xa1, 0xf6,
	0x1d, 0xd5, 0x82, 0x1b, 0xc3, 0x95, 0xa4, 0x45, 0x4a, 0x4b, 0x52, 0x53, 0xed, 0xbf, 0x85, 0x03,
	0x72, 0x66, 0x47, 0x4a, 0x73, 0x5b, 0x07, 0xe0, 0x31, 0x78, 0x3a, 0x38, 0xdc, 0x59, 0x4c, 0xe2,
	0xfb, 0x35, 0x11, 0xe5, 0x41, 0xf2, 0x17, 0x25, 0x3f, 0xbe, 0xed, 0x6e, 0x76, 0x9f, 0xbd, 0x2e,
	0x0a, 0x4d, 0x8d, 0x39, 0xb1, 0x9a, 0x4b, 0x96, 0x2e, 0x9f, 0xfb, 0xcf, 0xe0, 0x7a, 0xa5, 0xb4,
	0x3d, 0xe5, 0x45, 0xb0, 0xd2, 0x3a, 0xf9, 0x8b, 0x49, 0x7c, 0xcf, 0x39, 0x75, 0x20, 0x49, 0xd7,
	0x9a, 0xe8, 0xa8, 0xf0, 0x5f, 0x40, 0x98, 0x8f, 0x88, 0x94, 0xb4, 0x6c, 0xf4, 0xab, 0xad, 0xfe,
	0xc1, 0x62, 0x12, 0x6f, 0x38, 0xfd, 0x92, 0x25, 0xe9, 0xa0, 0x4b, 0x8e, 0x0a, 0x7f, 0x07, 0xae,
	0x6b, 0x57, 0x79, 0x70, 0xe7, 0xff, 0x2f, 0x3a, 0x90, 0xa4, 0x37, 0x92, 0x83, 0x97, 0x9f, 0xae,
	0x62, 0xef, 0xf7, 0x55, 0xec, 0x7d, 0x9c, 0x8f, 0x87, 0xcb, 0x42, 0x2f, 0xe7, 0xe3, 0x61, 0xd4,
	0x8e, 0xa7, 0xb7, 0x2f, 0xc9, 0x13, 0xb8, 0xdd, 0x0b, 0x53, 0x6a, 0x2a, 0x25, 0x0d, 0x7d, 0xfe,
	0x05, 0xc0, 0xd5, 0x63, 0xc3, 0xfc, 0x4b, 0x00, 0x1f, 0xf6, 0xf4, 0x77, 0x1f, 0xf5, 0x4d, 0x1d,
	0xf5, 0xfa, 0x87, 0xaf, 0x6e, 0xf1, 0xe8, 0xa6, 0xa8, 0xf0, 0xee, 0x87, 0xf9, 0x78, 0x08, 0x0e,
	0xdf, 0x7c, 0x9f, 0x46, 0xe0, 0x7a, 0x1a, 0x81, 0x5f, 0xd3, 0x08, 0x7c, 0x9e, 0x45, 0xde, 0xf5,
	0x2c, 0xf2, 0x7e, 0xce, 0x22, 0xef, 0x3d, 0x62, 0xdc, 0x8e, 0xce, 0x32, 0x94, 0x2b, 0x81, 0xb9,
	0xe4, 0x96, 0x93, 0xdd, 0x92, 0x64, 0xa6, 0x8b, 0xf1, 0x05, 0x6e, 0x56, 0xb9, 0x6d, 0x8f, 0xad,
	0x2b, 0x6a, 0xb2, 0xb5, 0x76, 0x8f, 0xf6, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xef, 0x15,
	0x46, 0xe3, 0x02, 0x00, 0x00,
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
	// SetPermissionedRelayer defines a rpc handler method for MsgSetPermissionedRelayer.
	SetPermissionedRelayer(ctx context.Context, in *MsgSetPermissionedRelayer, opts ...grpc.CallOption) (*MsgSetPermissionedRelayerResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SetPermissionedRelayer(ctx context.Context, in *MsgSetPermissionedRelayer, opts ...grpc.CallOption) (*MsgSetPermissionedRelayerResponse, error) {
	out := new(MsgSetPermissionedRelayerResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.perm.v1.Msg/SetPermissionedRelayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SetPermissionedRelayer defines a rpc handler method for MsgSetPermissionedRelayer.
	SetPermissionedRelayer(context.Context, *MsgSetPermissionedRelayer) (*MsgSetPermissionedRelayerResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SetPermissionedRelayer(ctx context.Context, req *MsgSetPermissionedRelayer) (*MsgSetPermissionedRelayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPermissionedRelayer not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SetPermissionedRelayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetPermissionedRelayer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetPermissionedRelayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.perm.v1.Msg/SetPermissionedRelayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetPermissionedRelayer(ctx, req.(*MsgSetPermissionedRelayer))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.applications.perm.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetPermissionedRelayer",
			Handler:    _Msg_SetPermissionedRelayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/applications/perm/v1/tx.proto",
}

func (m *MsgSetPermissionedRelayer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetPermissionedRelayer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetPermissionedRelayer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Relayer) > 0 {
		i -= len(m.Relayer)
		copy(dAtA[i:], m.Relayer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Relayer)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetPermissionedRelayerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetPermissionedRelayerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetPermissionedRelayerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgSetPermissionedRelayer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Relayer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSetPermissionedRelayerResponse) Size() (n int) {
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
func (m *MsgSetPermissionedRelayer) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetPermissionedRelayer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetPermissionedRelayer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
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
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
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
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
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
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relayer", wireType)
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
			m.Relayer = string(dAtA[iNdEx:postIndex])
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
func (m *MsgSetPermissionedRelayerResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetPermissionedRelayerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetPermissionedRelayerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
