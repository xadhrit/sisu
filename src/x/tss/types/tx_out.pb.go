// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tx_out.proto

package types

import (
	fmt "fmt"
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

type TxOut struct {
	Signer        string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	InChain       string `protobuf:"bytes,2,opt,name=inChain,proto3" json:"inChain,omitempty"`
	OutChain      string `protobuf:"bytes,3,opt,name=outChain,proto3" json:"outChain,omitempty"`
	InBlockHeight int64  `protobuf:"varint,4,opt,name=inBlockHeight,proto3" json:"inBlockHeight,omitempty"`
	InHash        string `protobuf:"bytes,5,opt,name=inHash,proto3" json:"inHash,omitempty"`
	OutBytes      []byte `protobuf:"bytes,6,opt,name=outBytes,proto3" json:"outBytes,omitempty"`
}

func (m *TxOut) Reset()         { *m = TxOut{} }
func (m *TxOut) String() string { return proto.CompactTextString(m) }
func (*TxOut) ProtoMessage()    {}
func (*TxOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8b99e4d8e912782, []int{0}
}
func (m *TxOut) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxOut.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxOut.Merge(m, src)
}
func (m *TxOut) XXX_Size() int {
	return m.Size()
}
func (m *TxOut) XXX_DiscardUnknown() {
	xxx_messageInfo_TxOut.DiscardUnknown(m)
}

var xxx_messageInfo_TxOut proto.InternalMessageInfo

func (m *TxOut) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *TxOut) GetInChain() string {
	if m != nil {
		return m.InChain
	}
	return ""
}

func (m *TxOut) GetOutChain() string {
	if m != nil {
		return m.OutChain
	}
	return ""
}

func (m *TxOut) GetInBlockHeight() int64 {
	if m != nil {
		return m.InBlockHeight
	}
	return 0
}

func (m *TxOut) GetInHash() string {
	if m != nil {
		return m.InHash
	}
	return ""
}

func (m *TxOut) GetOutBytes() []byte {
	if m != nil {
		return m.OutBytes
	}
	return nil
}

func init() {
	proto.RegisterType((*TxOut)(nil), "types.TxOut")
}

func init() { proto.RegisterFile("tx_out.proto", fileDescriptor_f8b99e4d8e912782) }

var fileDescriptor_f8b99e4d8e912782 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xa9, 0x88, 0xcf,
	0x2f, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x56,
	0x5a, 0xcf, 0xc8, 0xc5, 0x1a, 0x52, 0xe1, 0x5f, 0x5a, 0x22, 0x24, 0xc6, 0xc5, 0x56, 0x9c, 0x99,
	0x9e, 0x97, 0x5a, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe5, 0x09, 0x49, 0x70, 0xb1,
	0x67, 0xe6, 0x39, 0x67, 0x24, 0x66, 0xe6, 0x49, 0x30, 0x81, 0x25, 0x60, 0x5c, 0x21, 0x29, 0x2e,
	0x8e, 0xfc, 0xd2, 0x12, 0x88, 0x14, 0x33, 0x58, 0x0a, 0xce, 0x17, 0x52, 0xe1, 0xe2, 0xcd, 0xcc,
	0x73, 0xca, 0xc9, 0x4f, 0xce, 0xf6, 0x48, 0xcd, 0x4c, 0xcf, 0x28, 0x91, 0x60, 0x51, 0x60, 0xd4,
	0x60, 0x0e, 0x42, 0x15, 0x04, 0xd9, 0x99, 0x99, 0xe7, 0x91, 0x58, 0x9c, 0x21, 0xc1, 0x0a, 0xb1,
	0x13, 0xc2, 0x83, 0x9a, 0xec, 0x54, 0x59, 0x92, 0x5a, 0x2c, 0xc1, 0xa6, 0xc0, 0xa8, 0xc1, 0x13,
	0x04, 0xe7, 0x3b, 0x49, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72,
	0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x12, 0x1b,
	0xd8, 0x67, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0x06, 0x86, 0x18, 0xe9, 0x00, 0x00,
	0x00,
}

func (m *TxOut) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxOut) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxOut) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OutBytes) > 0 {
		i -= len(m.OutBytes)
		copy(dAtA[i:], m.OutBytes)
		i = encodeVarintTxOut(dAtA, i, uint64(len(m.OutBytes)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.InHash) > 0 {
		i -= len(m.InHash)
		copy(dAtA[i:], m.InHash)
		i = encodeVarintTxOut(dAtA, i, uint64(len(m.InHash)))
		i--
		dAtA[i] = 0x2a
	}
	if m.InBlockHeight != 0 {
		i = encodeVarintTxOut(dAtA, i, uint64(m.InBlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.OutChain) > 0 {
		i -= len(m.OutChain)
		copy(dAtA[i:], m.OutChain)
		i = encodeVarintTxOut(dAtA, i, uint64(len(m.OutChain)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InChain) > 0 {
		i -= len(m.InChain)
		copy(dAtA[i:], m.InChain)
		i = encodeVarintTxOut(dAtA, i, uint64(len(m.InChain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTxOut(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTxOut(dAtA []byte, offset int, v uint64) int {
	offset -= sovTxOut(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TxOut) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTxOut(uint64(l))
	}
	l = len(m.InChain)
	if l > 0 {
		n += 1 + l + sovTxOut(uint64(l))
	}
	l = len(m.OutChain)
	if l > 0 {
		n += 1 + l + sovTxOut(uint64(l))
	}
	if m.InBlockHeight != 0 {
		n += 1 + sovTxOut(uint64(m.InBlockHeight))
	}
	l = len(m.InHash)
	if l > 0 {
		n += 1 + l + sovTxOut(uint64(l))
	}
	l = len(m.OutBytes)
	if l > 0 {
		n += 1 + l + sovTxOut(uint64(l))
	}
	return n
}

func sovTxOut(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTxOut(x uint64) (n int) {
	return sovTxOut(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TxOut) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxOut
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
			return fmt.Errorf("proto: TxOut: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxOut: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
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
				return ErrInvalidLengthTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
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
				return ErrInvalidLengthTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
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
				return ErrInvalidLengthTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InBlockHeight", wireType)
			}
			m.InBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InBlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
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
				return ErrInvalidLengthTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxOut
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutBytes = append(m.OutBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.OutBytes == nil {
				m.OutBytes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTxOut(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTxOut
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
func skipTxOut(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTxOut
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
					return 0, ErrIntOverflowTxOut
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
					return 0, ErrIntOverflowTxOut
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
				return 0, ErrInvalidLengthTxOut
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTxOut
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTxOut
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTxOut        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTxOut          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTxOut = fmt.Errorf("proto: unexpected end of group")
)