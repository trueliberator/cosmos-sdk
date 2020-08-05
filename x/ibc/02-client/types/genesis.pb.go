// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/client/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisClientState defines an identified ClientState as protobuf Any format.
type GenesisClientState struct {
	ClientID    string     `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" yaml:"client_id"`
	ClientState *types.Any `protobuf:"bytes,2,opt,name=client_state,json=clientState,proto3" json:"client_state,omitempty" yaml:"client_state"`
}

func (m *GenesisClientState) Reset()         { *m = GenesisClientState{} }
func (m *GenesisClientState) String() string { return proto.CompactTextString(m) }
func (*GenesisClientState) ProtoMessage()    {}
func (*GenesisClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eb5d7ff040be5c2, []int{0}
}
func (m *GenesisClientState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisClientState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisClientState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisClientState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisClientState.Merge(m, src)
}
func (m *GenesisClientState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisClientState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisClientState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisClientState proto.InternalMessageInfo

func (m *GenesisClientState) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *GenesisClientState) GetClientState() *types.Any {
	if m != nil {
		return m.ClientState
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisClientState)(nil), "ibc.client.GenesisClientState")
}

func init() { proto.RegisterFile("ibc/client/genesis.proto", fileDescriptor_2eb5d7ff040be5c2) }

var fileDescriptor_2eb5d7ff040be5c2 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0x4c, 0x4a, 0xd6,
	0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0xd1, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0x4c, 0x4a, 0xd6, 0x83, 0xc8, 0x48, 0x89, 0xa4, 0xe7,
	0xa7, 0xe7, 0x83, 0x85, 0xf5, 0x41, 0x2c, 0x88, 0x0a, 0x29, 0xc9, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c,
	0x54, 0x7d, 0x30, 0x2f, 0xa9, 0x34, 0x4d, 0x3f, 0x31, 0xaf, 0x12, 0x22, 0xa5, 0xb4, 0x94, 0x91,
	0x4b, 0xc8, 0x1d, 0x62, 0x9c, 0x33, 0xd8, 0x88, 0xe0, 0x92, 0xc4, 0x92, 0x54, 0x21, 0x5b, 0x2e,
	0x4e, 0x88, 0x89, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x4e, 0x0a, 0x8f, 0xee,
	0xc9, 0x73, 0x40, 0xd4, 0x78, 0xba, 0x7c, 0xba, 0x27, 0x2f, 0x50, 0x99, 0x98, 0x9b, 0x63, 0xa5,
	0x04, 0x57, 0xa6, 0x14, 0xc4, 0x01, 0x61, 0x7b, 0xa6, 0x08, 0x05, 0x70, 0xf1, 0x40, 0xc5, 0x8b,
	0x41, 0xc6, 0x49, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x1b, 0x89, 0xe8, 0x41, 0xdc, 0xa1, 0x07, 0x73,
	0x87, 0x9e, 0x63, 0x5e, 0xa5, 0x93, 0xf8, 0xa7, 0x7b, 0xf2, 0xc2, 0x28, 0x66, 0x81, 0xf5, 0x28,
	0x05, 0x71, 0x27, 0x23, 0x1c, 0xe4, 0xe4, 0x73, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72,
	0x0c, 0x51, 0x46, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc9, 0xf9,
	0xc5, 0xb9, 0xf9, 0xc5, 0x50, 0x4a, 0xb7, 0x38, 0x25, 0x5b, 0xbf, 0x42, 0x1f, 0x14, 0x6e, 0x06,
	0x46, 0xba, 0xd0, 0xa0, 0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xbb, 0xc0, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x12, 0x6b, 0x10, 0x60, 0x55, 0x01, 0x00, 0x00,
}

func (m *GenesisClientState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisClientState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisClientState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClientState != nil {
		{
			size, err := m.ClientState.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisClientState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.ClientState != nil {
		l = m.ClientState.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisClientState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisClientState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisClientState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClientState == nil {
				m.ClientState = &types.Any{}
			}
			if err := m.ClientState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
