// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sisu/pause_contract.proto

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

type PauseContractMsg struct {
	Signer string         `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	Data   *PauseContract `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *PauseContractMsg) Reset()         { *m = PauseContractMsg{} }
func (m *PauseContractMsg) String() string { return proto.CompactTextString(m) }
func (*PauseContractMsg) ProtoMessage()    {}
func (*PauseContractMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ae7eec68344454, []int{0}
}
func (m *PauseContractMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PauseContractMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PauseContractMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PauseContractMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PauseContractMsg.Merge(m, src)
}
func (m *PauseContractMsg) XXX_Size() int {
	return m.Size()
}
func (m *PauseContractMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_PauseContractMsg.DiscardUnknown(m)
}

var xxx_messageInfo_PauseContractMsg proto.InternalMessageInfo

func (m *PauseContractMsg) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *PauseContractMsg) GetData() *PauseContract {
	if m != nil {
		return m.Data
	}
	return nil
}

type PauseContract struct {
	Chain string `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	Hash  string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Index int32  `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *PauseContract) Reset()         { *m = PauseContract{} }
func (m *PauseContract) String() string { return proto.CompactTextString(m) }
func (*PauseContract) ProtoMessage()    {}
func (*PauseContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ae7eec68344454, []int{1}
}
func (m *PauseContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PauseContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PauseContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PauseContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PauseContract.Merge(m, src)
}
func (m *PauseContract) XXX_Size() int {
	return m.Size()
}
func (m *PauseContract) XXX_DiscardUnknown() {
	xxx_messageInfo_PauseContract.DiscardUnknown(m)
}

var xxx_messageInfo_PauseContract proto.InternalMessageInfo

func (m *PauseContract) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *PauseContract) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *PauseContract) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func init() {
	proto.RegisterType((*PauseContractMsg)(nil), "types.PauseContractMsg")
	proto.RegisterType((*PauseContract)(nil), "types.PauseContract")
}

func init() { proto.RegisterFile("sisu/pause_contract.proto", fileDescriptor_04ae7eec68344454) }

var fileDescriptor_04ae7eec68344454 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0xce, 0x2c, 0x2e,
	0xd5, 0x2f, 0x48, 0x2c, 0x2d, 0x4e, 0x8d, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x56, 0x0a, 0xe1, 0x12,
	0x08, 0x00, 0x49, 0x3b, 0x43, 0x65, 0x7d, 0x8b, 0xd3, 0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0x33, 0xd3,
	0xf3, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x21, 0x0d, 0x2e, 0x96,
	0x94, 0xc4, 0x92, 0x44, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x11, 0x3d, 0xb0, 0x09, 0x7a,
	0x28, 0xda, 0x83, 0xc0, 0x2a, 0x94, 0xfc, 0xb9, 0x78, 0x51, 0x84, 0x85, 0x44, 0xb8, 0x58, 0x93,
	0x33, 0x12, 0x33, 0xf3, 0xa0, 0x26, 0x42, 0x38, 0x42, 0x42, 0x5c, 0x2c, 0x19, 0x89, 0xc5, 0x19,
	0x60, 0x03, 0x39, 0x83, 0xc0, 0x6c, 0x90, 0xca, 0xcc, 0xbc, 0x94, 0xd4, 0x0a, 0x09, 0x66, 0x05,
	0x46, 0x0d, 0xd6, 0x20, 0x08, 0xc7, 0xc9, 0xf9, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18,
	0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5,
	0x18, 0xa2, 0x34, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x41, 0xbe,
	0xd5, 0xcd, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0x06, 0x73, 0xf4, 0x2b, 0x20, 0x14, 0xd8, 0xa5,
	0x49, 0x6c, 0x60, 0x9f, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x04, 0xd1, 0x6f, 0x5e, 0x16,
	0x01, 0x00, 0x00,
}

func (m *PauseContractMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PauseContractMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PauseContractMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPauseContract(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintPauseContract(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PauseContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PauseContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PauseContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		i = encodeVarintPauseContract(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintPauseContract(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintPauseContract(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPauseContract(dAtA []byte, offset int, v uint64) int {
	offset -= sovPauseContract(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PauseContractMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovPauseContract(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovPauseContract(uint64(l))
	}
	return n
}

func (m *PauseContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovPauseContract(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovPauseContract(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovPauseContract(uint64(m.Index))
	}
	return n
}

func sovPauseContract(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPauseContract(x uint64) (n int) {
	return sovPauseContract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PauseContractMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPauseContract
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
			return fmt.Errorf("proto: PauseContractMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PauseContractMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPauseContract
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
				return ErrInvalidLengthPauseContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPauseContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPauseContract
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
				return ErrInvalidLengthPauseContract
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPauseContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &PauseContract{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPauseContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPauseContract
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
func (m *PauseContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPauseContract
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
			return fmt.Errorf("proto: PauseContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PauseContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPauseContract
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
				return ErrInvalidLengthPauseContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPauseContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPauseContract
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
				return ErrInvalidLengthPauseContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPauseContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPauseContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPauseContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPauseContract
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
func skipPauseContract(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPauseContract
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
					return 0, ErrIntOverflowPauseContract
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
					return 0, ErrIntOverflowPauseContract
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
				return 0, ErrInvalidLengthPauseContract
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPauseContract
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPauseContract
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPauseContract        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPauseContract          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPauseContract = fmt.Errorf("proto: unexpected end of group")
)
