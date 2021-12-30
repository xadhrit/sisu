// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/tss/query.proto

package types

import (
	context "context"
	fmt "fmt"
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

type QueryAllPubKeysRequest struct {
}

func (m *QueryAllPubKeysRequest) Reset()         { *m = QueryAllPubKeysRequest{} }
func (m *QueryAllPubKeysRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllPubKeysRequest) ProtoMessage()    {}
func (*QueryAllPubKeysRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fa0b9365e3b43bc, []int{0}
}
func (m *QueryAllPubKeysRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPubKeysRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPubKeysRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPubKeysRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPubKeysRequest.Merge(m, src)
}
func (m *QueryAllPubKeysRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPubKeysRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPubKeysRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPubKeysRequest proto.InternalMessageInfo

type QueryAllPubKeysResponse struct {
	Pubkeys map[string][]byte `protobuf:"bytes,1,rep,name=pubkeys,proto3" json:"pubkeys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *QueryAllPubKeysResponse) Reset()         { *m = QueryAllPubKeysResponse{} }
func (m *QueryAllPubKeysResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllPubKeysResponse) ProtoMessage()    {}
func (*QueryAllPubKeysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fa0b9365e3b43bc, []int{1}
}
func (m *QueryAllPubKeysResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPubKeysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPubKeysResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPubKeysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPubKeysResponse.Merge(m, src)
}
func (m *QueryAllPubKeysResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPubKeysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPubKeysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPubKeysResponse proto.InternalMessageInfo

func (m *QueryAllPubKeysResponse) GetPubkeys() map[string][]byte {
	if m != nil {
		return m.Pubkeys
	}
	return nil
}

type QueryContractRequest struct {
	Chain string `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	Hash  string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *QueryContractRequest) Reset()         { *m = QueryContractRequest{} }
func (m *QueryContractRequest) String() string { return proto.CompactTextString(m) }
func (*QueryContractRequest) ProtoMessage()    {}
func (*QueryContractRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fa0b9365e3b43bc, []int{2}
}
func (m *QueryContractRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContractRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContractRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractRequest.Merge(m, src)
}
func (m *QueryContractRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryContractRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractRequest proto.InternalMessageInfo

func (m *QueryContractRequest) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *QueryContractRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type QueryContractResponse struct {
	Contract *Contract `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (m *QueryContractResponse) Reset()         { *m = QueryContractResponse{} }
func (m *QueryContractResponse) String() string { return proto.CompactTextString(m) }
func (*QueryContractResponse) ProtoMessage()    {}
func (*QueryContractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fa0b9365e3b43bc, []int{3}
}
func (m *QueryContractResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractResponse.Merge(m, src)
}
func (m *QueryContractResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryContractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractResponse proto.InternalMessageInfo

func (m *QueryContractResponse) GetContract() *Contract {
	if m != nil {
		return m.Contract
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAllPubKeysRequest)(nil), "types.QueryAllPubKeysRequest")
	proto.RegisterType((*QueryAllPubKeysResponse)(nil), "types.QueryAllPubKeysResponse")
	proto.RegisterMapType((map[string][]byte)(nil), "types.QueryAllPubKeysResponse.PubkeysEntry")
	proto.RegisterType((*QueryContractRequest)(nil), "types.QueryContractRequest")
	proto.RegisterType((*QueryContractResponse)(nil), "types.QueryContractResponse")
}

func init() { proto.RegisterFile("proto/tss/query.proto", fileDescriptor_1fa0b9365e3b43bc) }

var fileDescriptor_1fa0b9365e3b43bc = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x4a, 0xc3, 0x40,
	0x14, 0xce, 0xb4, 0x56, 0xed, 0x6b, 0x45, 0x19, 0x5a, 0x1d, 0xa2, 0x0e, 0x25, 0xab, 0x42, 0x21,
	0x85, 0xba, 0x91, 0xae, 0xfc, 0xeb, 0x4a, 0x85, 0x9a, 0x1b, 0x24, 0x61, 0x20, 0xd2, 0x90, 0xa4,
	0x99, 0x89, 0x30, 0xb7, 0xf0, 0x00, 0x2e, 0x3d, 0x8c, 0xcb, 0x2e, 0x5d, 0x4a, 0x72, 0x11, 0xc9,
	0x24, 0x69, 0x6b, 0xb5, 0xee, 0xe6, 0xf1, 0xfd, 0xbc, 0xf7, 0x7d, 0x0c, 0x74, 0xa3, 0x38, 0x14,
	0xe1, 0x50, 0x70, 0x3e, 0x9c, 0x27, 0x2c, 0x96, 0xa6, 0x9a, 0x71, 0x43, 0xc8, 0x88, 0x71, 0x9d,
	0xac, 0x50, 0x37, 0x0c, 0x44, 0x6c, 0xbb, 0xa2, 0x20, 0x18, 0x04, 0x8e, 0x9f, 0x72, 0xfe, 0xb5,
	0xef, 0x4f, 0x13, 0xe7, 0x9e, 0x49, 0x6e, 0xb1, 0x79, 0xc2, 0xb8, 0x30, 0xde, 0x10, 0x9c, 0xfc,
	0x82, 0x78, 0x14, 0x06, 0x9c, 0xe1, 0x09, 0xec, 0x45, 0x89, 0x33, 0x63, 0x92, 0x13, 0xd4, 0xab,
	0xf7, 0x5b, 0xa3, 0x81, 0xa9, 0x16, 0x99, 0x5b, 0x04, 0xe6, 0xb4, 0x60, 0x4f, 0x02, 0x11, 0x4b,
	0xab, 0xd2, 0xea, 0x63, 0x68, 0xaf, 0x03, 0xf8, 0x08, 0xea, 0x33, 0x26, 0x09, 0xea, 0xa1, 0x7e,
	0xd3, 0xca, 0x9f, 0xb8, 0x03, 0x8d, 0x17, 0xdb, 0x4f, 0x18, 0xa9, 0xf5, 0x50, 0xbf, 0x6d, 0x15,
	0xc3, 0xb8, 0x76, 0x89, 0x8c, 0x2b, 0xe8, 0xa8, 0x65, 0xb7, 0x65, 0x9e, 0xf2, 0xec, 0x5c, 0xe1,
	0x7a, 0xf6, 0x73, 0x50, 0xba, 0x14, 0x03, 0xc6, 0xb0, 0xe3, 0xd9, 0xdc, 0x53, 0x36, 0x4d, 0x4b,
	0xbd, 0x8d, 0x3b, 0xe8, 0x6e, 0x38, 0x94, 0xe9, 0x06, 0xb0, 0x5f, 0xb5, 0xa4, 0x5c, 0x5a, 0xa3,
	0xc3, 0x32, 0xde, 0x92, 0xba, 0x24, 0x8c, 0xde, 0x11, 0x34, 0x94, 0x0d, 0x7e, 0x04, 0x58, 0x25,
	0xc7, 0xe7, 0xdb, 0x1a, 0x51, 0x67, 0xea, 0xf4, 0xff, 0xc2, 0x0c, 0x0d, 0x3f, 0xc0, 0xc1, 0x8f,
	0xf3, 0xf0, 0xe9, 0xba, 0x64, 0x23, 0xb6, 0x7e, 0xf6, 0x37, 0x58, 0xb9, 0xdd, 0x90, 0x8f, 0x94,
	0xa2, 0x45, 0x4a, 0xd1, 0x57, 0x4a, 0xd1, 0x6b, 0x46, 0xb5, 0x45, 0x46, 0xb5, 0xcf, 0x8c, 0x6a,
	0xce, 0xae, 0xfa, 0x08, 0x17, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x46, 0x3d, 0x06, 0x9d, 0x42,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	AllPubKeys(ctx context.Context, in *QueryAllPubKeysRequest, opts ...grpc.CallOption) (*QueryAllPubKeysResponse, error)
	QueryContract(ctx context.Context, in *QueryContractRequest, opts ...grpc.CallOption) (*QueryContractResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) AllPubKeys(ctx context.Context, in *QueryAllPubKeysRequest, opts ...grpc.CallOption) (*QueryAllPubKeysResponse, error) {
	out := new(QueryAllPubKeysResponse)
	err := c.cc.Invoke(ctx, "/types.Query/AllPubKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryContract(ctx context.Context, in *QueryContractRequest, opts ...grpc.CallOption) (*QueryContractResponse, error) {
	out := new(QueryContractResponse)
	err := c.cc.Invoke(ctx, "/types.Query/QueryContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	AllPubKeys(context.Context, *QueryAllPubKeysRequest) (*QueryAllPubKeysResponse, error)
	QueryContract(context.Context, *QueryContractRequest) (*QueryContractResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) AllPubKeys(ctx context.Context, req *QueryAllPubKeysRequest) (*QueryAllPubKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllPubKeys not implemented")
}
func (*UnimplementedQueryServer) QueryContract(ctx context.Context, req *QueryContractRequest) (*QueryContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryContract not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_AllPubKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPubKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllPubKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Query/AllPubKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllPubKeys(ctx, req.(*QueryAllPubKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Query/QueryContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryContract(ctx, req.(*QueryContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllPubKeys",
			Handler:    _Query_AllPubKeys_Handler,
		},
		{
			MethodName: "QueryContract",
			Handler:    _Query_QueryContract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tss/query.proto",
}

func (m *QueryAllPubKeysRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPubKeysRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPubKeysRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAllPubKeysResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPubKeysResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPubKeysResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pubkeys) > 0 {
		for k := range m.Pubkeys {
			v := m.Pubkeys[k]
			baseI := i
			if len(v) > 0 {
				i -= len(v)
				copy(dAtA[i:], v)
				i = encodeVarintQuery(dAtA, i, uint64(len(v)))
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintQuery(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintQuery(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryContractRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryContractResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Contract != nil {
		{
			size, err := m.Contract.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryAllPubKeysRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAllPubKeysResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pubkeys) > 0 {
		for k, v := range m.Pubkeys {
			_ = k
			_ = v
			l = 0
			if len(v) > 0 {
				l = 1 + len(v) + sovQuery(uint64(len(v)))
			}
			mapEntrySize := 1 + len(k) + sovQuery(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovQuery(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *QueryContractRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryContractResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Contract != nil {
		l = m.Contract.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAllPubKeysRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllPubKeysRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPubKeysRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllPubKeysResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllPubKeysResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPubKeysResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkeys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pubkeys == nil {
				m.Pubkeys = make(map[string][]byte)
			}
			var mapkey string
			mapvalue := []byte{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuery
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthQuery
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthQuery
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapbyteLen uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuery
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapbyteLen |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intMapbyteLen := int(mapbyteLen)
					if intMapbyteLen < 0 {
						return ErrInvalidLengthQuery
					}
					postbytesIndex := iNdEx + intMapbyteLen
					if postbytesIndex < 0 {
						return ErrInvalidLengthQuery
					}
					if postbytesIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = make([]byte, mapbyteLen)
					copy(mapvalue, dAtA[iNdEx:postbytesIndex])
					iNdEx = postbytesIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipQuery(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthQuery
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Pubkeys[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryContractRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryContractRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
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
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryContractResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryContractResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Contract == nil {
				m.Contract = &Contract{}
			}
			if err := m.Contract.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
