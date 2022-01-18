// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/tss/keygen.proto

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

type KeygenResult_Result int32

const (
	KeygenResult_SUCCESS      KeygenResult_Result = 0
	KeygenResult_FAILURE      KeygenResult_Result = 1
	KeygenResult_NOT_SELECTED KeygenResult_Result = 2
)

var KeygenResult_Result_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILURE",
	2: "NOT_SELECTED",
}

var KeygenResult_Result_value = map[string]int32{
	"SUCCESS":      0,
	"FAILURE":      1,
	"NOT_SELECTED": 2,
}

func (x KeygenResult_Result) String() string {
	return proto.EnumName(KeygenResult_Result_name, int32(x))
}

func (KeygenResult_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8de4706b10e43954, []int{3, 0}
}

type KeygenWithSigner struct {
	Signer string  `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	Data   *Keygen `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *KeygenWithSigner) Reset()         { *m = KeygenWithSigner{} }
func (m *KeygenWithSigner) String() string { return proto.CompactTextString(m) }
func (*KeygenWithSigner) ProtoMessage()    {}
func (*KeygenWithSigner) Descriptor() ([]byte, []int) {
	return fileDescriptor_8de4706b10e43954, []int{0}
}
func (m *KeygenWithSigner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeygenWithSigner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeygenWithSigner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeygenWithSigner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeygenWithSigner.Merge(m, src)
}
func (m *KeygenWithSigner) XXX_Size() int {
	return m.Size()
}
func (m *KeygenWithSigner) XXX_DiscardUnknown() {
	xxx_messageInfo_KeygenWithSigner.DiscardUnknown(m)
}

var xxx_messageInfo_KeygenWithSigner proto.InternalMessageInfo

func (m *KeygenWithSigner) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *KeygenWithSigner) GetData() *Keygen {
	if m != nil {
		return m.Data
	}
	return nil
}

type Keygen struct {
	KeyType     string `protobuf:"bytes,1,opt,name=keyType,proto3" json:"keyType,omitempty"`
	Index       int32  `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	PubKeyBytes []byte `protobuf:"bytes,3,opt,name=pubKeyBytes,proto3" json:"pubKeyBytes,omitempty"`
	Address     string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	StartBlock  int64  `protobuf:"varint,5,opt,name=startBlock,proto3" json:"startBlock,omitempty"`
}

func (m *Keygen) Reset()         { *m = Keygen{} }
func (m *Keygen) String() string { return proto.CompactTextString(m) }
func (*Keygen) ProtoMessage()    {}
func (*Keygen) Descriptor() ([]byte, []int) {
	return fileDescriptor_8de4706b10e43954, []int{1}
}
func (m *Keygen) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Keygen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Keygen.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Keygen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keygen.Merge(m, src)
}
func (m *Keygen) XXX_Size() int {
	return m.Size()
}
func (m *Keygen) XXX_DiscardUnknown() {
	xxx_messageInfo_Keygen.DiscardUnknown(m)
}

var xxx_messageInfo_Keygen proto.InternalMessageInfo

func (m *Keygen) GetKeyType() string {
	if m != nil {
		return m.KeyType
	}
	return ""
}

func (m *Keygen) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Keygen) GetPubKeyBytes() []byte {
	if m != nil {
		return m.PubKeyBytes
	}
	return nil
}

func (m *Keygen) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Keygen) GetStartBlock() int64 {
	if m != nil {
		return m.StartBlock
	}
	return 0
}

type KeygenResultWithSigner struct {
	Signer string        `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	Keygen *Keygen       `protobuf:"bytes,2,opt,name=keygen,proto3" json:"keygen,omitempty"`
	Data   *KeygenResult `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *KeygenResultWithSigner) Reset()         { *m = KeygenResultWithSigner{} }
func (m *KeygenResultWithSigner) String() string { return proto.CompactTextString(m) }
func (*KeygenResultWithSigner) ProtoMessage()    {}
func (*KeygenResultWithSigner) Descriptor() ([]byte, []int) {
	return fileDescriptor_8de4706b10e43954, []int{2}
}
func (m *KeygenResultWithSigner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeygenResultWithSigner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeygenResultWithSigner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeygenResultWithSigner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeygenResultWithSigner.Merge(m, src)
}
func (m *KeygenResultWithSigner) XXX_Size() int {
	return m.Size()
}
func (m *KeygenResultWithSigner) XXX_DiscardUnknown() {
	xxx_messageInfo_KeygenResultWithSigner.DiscardUnknown(m)
}

var xxx_messageInfo_KeygenResultWithSigner proto.InternalMessageInfo

func (m *KeygenResultWithSigner) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *KeygenResultWithSigner) GetKeygen() *Keygen {
	if m != nil {
		return m.Keygen
	}
	return nil
}

func (m *KeygenResultWithSigner) GetData() *KeygenResult {
	if m != nil {
		return m.Data
	}
	return nil
}

type KeygenResult struct {
	From   string              `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Result KeygenResult_Result `protobuf:"varint,2,opt,name=result,proto3,enum=types.KeygenResult_Result" json:"result,omitempty"`
}

func (m *KeygenResult) Reset()         { *m = KeygenResult{} }
func (m *KeygenResult) String() string { return proto.CompactTextString(m) }
func (*KeygenResult) ProtoMessage()    {}
func (*KeygenResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_8de4706b10e43954, []int{3}
}
func (m *KeygenResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeygenResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeygenResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeygenResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeygenResult.Merge(m, src)
}
func (m *KeygenResult) XXX_Size() int {
	return m.Size()
}
func (m *KeygenResult) XXX_DiscardUnknown() {
	xxx_messageInfo_KeygenResult.DiscardUnknown(m)
}

var xxx_messageInfo_KeygenResult proto.InternalMessageInfo

func (m *KeygenResult) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *KeygenResult) GetResult() KeygenResult_Result {
	if m != nil {
		return m.Result
	}
	return KeygenResult_SUCCESS
}

func init() {
	proto.RegisterEnum("types.KeygenResult_Result", KeygenResult_Result_name, KeygenResult_Result_value)
	proto.RegisterType((*KeygenWithSigner)(nil), "types.KeygenWithSigner")
	proto.RegisterType((*Keygen)(nil), "types.Keygen")
	proto.RegisterType((*KeygenResultWithSigner)(nil), "types.KeygenResultWithSigner")
	proto.RegisterType((*KeygenResult)(nil), "types.KeygenResult")
}

func init() { proto.RegisterFile("proto/tss/keygen.proto", fileDescriptor_8de4706b10e43954) }

var fileDescriptor_8de4706b10e43954 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xb3, 0x6d, 0x93, 0xe2, 0xb4, 0x4a, 0x58, 0xa5, 0x2c, 0x1e, 0x96, 0x18, 0x10, 0x73,
	0x4a, 0xa1, 0xfa, 0x02, 0xb6, 0x46, 0x90, 0x56, 0x85, 0xa4, 0xc5, 0xa3, 0xa4, 0x66, 0xad, 0xa5,
	0xb5, 0x09, 0xd9, 0x2d, 0xb8, 0x37, 0x1f, 0xc0, 0x83, 0x07, 0x1f, 0xca, 0x63, 0x8f, 0x1e, 0xa5,
	0x7d, 0x11, 0xe9, 0x66, 0x85, 0x0a, 0x3d, 0x78, 0x9b, 0xff, 0x9f, 0xc9, 0x37, 0x99, 0x9f, 0x85,
	0x46, 0x96, 0xa7, 0x22, 0x6d, 0x0a, 0xce, 0x9b, 0x13, 0x26, 0x47, 0x6c, 0xe6, 0x2b, 0x03, 0x9b,
	0x42, 0x66, 0x8c, 0xbb, 0xd7, 0x60, 0x77, 0x95, 0x7d, 0x37, 0x16, 0x4f, 0xd1, 0x78, 0x34, 0x63,
	0x39, 0x6e, 0x80, 0xc5, 0x55, 0x45, 0x90, 0x83, 0xbc, 0x9d, 0x50, 0x2b, 0x7c, 0x04, 0x95, 0x24,
	0x16, 0x31, 0x29, 0x39, 0xc8, 0xab, 0xb5, 0x76, 0x7d, 0x45, 0xf0, 0x8b, 0xcf, 0x43, 0xd5, 0x72,
	0x3f, 0x10, 0x58, 0x85, 0x81, 0x09, 0x54, 0x27, 0x4c, 0xf6, 0x65, 0xc6, 0x34, 0xe6, 0x57, 0xe2,
	0x03, 0x30, 0xc7, 0xb3, 0x84, 0xbd, 0x28, 0x90, 0x19, 0x16, 0x02, 0x3b, 0x50, 0xcb, 0xe6, 0xc3,
	0x2e, 0x93, 0x6d, 0x29, 0x18, 0x27, 0x65, 0x07, 0x79, 0xf5, 0x70, 0xd3, 0x5a, 0x13, 0xe3, 0x24,
	0xc9, 0x19, 0xe7, 0xa4, 0x52, 0x10, 0xb5, 0xc4, 0x14, 0x80, 0x8b, 0x38, 0x17, 0xed, 0x69, 0xfa,
	0x30, 0x21, 0xa6, 0x83, 0xbc, 0x72, 0xb8, 0xe1, 0xb8, 0xaf, 0x08, 0x1a, 0xfa, 0x3f, 0x19, 0x9f,
	0x4f, 0xc5, 0x3f, 0x8e, 0x3d, 0x06, 0xab, 0xc8, 0x6b, 0xfb, 0xb9, 0xba, 0x89, 0x4f, 0x74, 0x26,
	0x65, 0x35, 0xb4, 0xff, 0x77, 0x48, 0xed, 0xd2, 0xc9, 0xbc, 0x21, 0xa8, 0x6f, 0xda, 0x18, 0x43,
	0xe5, 0x31, 0x4f, 0x9f, 0xf5, 0x5a, 0x55, 0xe3, 0x16, 0x58, 0xb9, 0xea, 0xaa, 0xa5, 0x7b, 0xad,
	0xc3, 0x2d, 0x3c, 0x5f, 0x63, 0xf5, 0xa4, 0x7b, 0x06, 0x96, 0x26, 0xd6, 0xa0, 0x1a, 0x0d, 0x3a,
	0x9d, 0x20, 0x8a, 0x6c, 0x63, 0x2d, 0x2e, 0xcf, 0xaf, 0x7a, 0x83, 0x30, 0xb0, 0x11, 0xb6, 0xa1,
	0x7e, 0x73, 0xdb, 0xbf, 0x8f, 0x82, 0x5e, 0xd0, 0xe9, 0x07, 0x17, 0x76, 0xa9, 0x4d, 0x3e, 0x97,
	0x14, 0x2d, 0x96, 0x14, 0x7d, 0x2f, 0x29, 0x7a, 0x5f, 0x51, 0x63, 0xb1, 0xa2, 0xc6, 0xd7, 0x8a,
	0x1a, 0x43, 0x4b, 0xbd, 0x8f, 0xd3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x33, 0x18, 0x3a,
	0x39, 0x02, 0x00, 0x00,
}

func (m *KeygenWithSigner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeygenWithSigner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeygenWithSigner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintKeygen(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintKeygen(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Keygen) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Keygen) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Keygen) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StartBlock != 0 {
		i = encodeVarintKeygen(dAtA, i, uint64(m.StartBlock))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintKeygen(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PubKeyBytes) > 0 {
		i -= len(m.PubKeyBytes)
		copy(dAtA[i:], m.PubKeyBytes)
		i = encodeVarintKeygen(dAtA, i, uint64(len(m.PubKeyBytes)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Index != 0 {
		i = encodeVarintKeygen(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x10
	}
	if len(m.KeyType) > 0 {
		i -= len(m.KeyType)
		copy(dAtA[i:], m.KeyType)
		i = encodeVarintKeygen(dAtA, i, uint64(len(m.KeyType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *KeygenResultWithSigner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeygenResultWithSigner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeygenResultWithSigner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintKeygen(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Keygen != nil {
		{
			size, err := m.Keygen.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintKeygen(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintKeygen(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *KeygenResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeygenResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeygenResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Result != 0 {
		i = encodeVarintKeygen(dAtA, i, uint64(m.Result))
		i--
		dAtA[i] = 0x10
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintKeygen(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintKeygen(dAtA []byte, offset int, v uint64) int {
	offset -= sovKeygen(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KeygenWithSigner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovKeygen(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovKeygen(uint64(l))
	}
	return n
}

func (m *Keygen) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.KeyType)
	if l > 0 {
		n += 1 + l + sovKeygen(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovKeygen(uint64(m.Index))
	}
	l = len(m.PubKeyBytes)
	if l > 0 {
		n += 1 + l + sovKeygen(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovKeygen(uint64(l))
	}
	if m.StartBlock != 0 {
		n += 1 + sovKeygen(uint64(m.StartBlock))
	}
	return n
}

func (m *KeygenResultWithSigner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovKeygen(uint64(l))
	}
	if m.Keygen != nil {
		l = m.Keygen.Size()
		n += 1 + l + sovKeygen(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovKeygen(uint64(l))
	}
	return n
}

func (m *KeygenResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovKeygen(uint64(l))
	}
	if m.Result != 0 {
		n += 1 + sovKeygen(uint64(m.Result))
	}
	return n
}

func sovKeygen(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKeygen(x uint64) (n int) {
	return sovKeygen(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeygenWithSigner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeygen
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
			return fmt.Errorf("proto: KeygenWithSigner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeygenWithSigner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeygen
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
				return ErrInvalidLengthKeygen
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeygen
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
					return ErrIntOverflowKeygen
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
				return ErrInvalidLengthKeygen
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeygen
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &Keygen{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeygen(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKeygen
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
func (m *Keygen) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeygen
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
			return fmt.Errorf("proto: Keygen: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Keygen: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeygen
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
				return ErrInvalidLengthKeygen
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeygen
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeygen
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKeyBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeygen
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
				return ErrInvalidLengthKeygen
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKeygen
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKeyBytes = append(m.PubKeyBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKeyBytes == nil {
				m.PubKeyBytes = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeygen
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
				return ErrInvalidLengthKeygen
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeygen
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartBlock", wireType)
			}
			m.StartBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeygen
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipKeygen(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKeygen
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
func (m *KeygenResultWithSigner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeygen
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
			return fmt.Errorf("proto: KeygenResultWithSigner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeygenResultWithSigner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeygen
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
				return ErrInvalidLengthKeygen
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeygen
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keygen", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeygen
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
				return ErrInvalidLengthKeygen
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeygen
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Keygen == nil {
				m.Keygen = &Keygen{}
			}
			if err := m.Keygen.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeygen
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
				return ErrInvalidLengthKeygen
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeygen
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &KeygenResult{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeygen(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKeygen
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
func (m *KeygenResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeygen
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
			return fmt.Errorf("proto: KeygenResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeygenResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeygen
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
				return ErrInvalidLengthKeygen
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeygen
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeygen
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= KeygenResult_Result(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipKeygen(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKeygen
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
func skipKeygen(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKeygen
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
					return 0, ErrIntOverflowKeygen
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
					return 0, ErrIntOverflowKeygen
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
				return 0, ErrInvalidLengthKeygen
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKeygen
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKeygen
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKeygen        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKeygen          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKeygen = fmt.Errorf("proto: unexpected end of group")
)
