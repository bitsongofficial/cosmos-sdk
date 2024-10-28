// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accounts/defaults/authz/authz.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	any "github.com/cosmos/gogoproto/types/any"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Grant defines a authorization grant for grantee
type Grant struct {
	Authorization *any.Any `protobuf:"bytes,1,opt,name=authorization,proto3" json:"authorization,omitempty"`
	// time when the grant will expired
	Expiration *time.Time `protobuf:"bytes,2,opt,name=expiration,proto3,stdtime" json:"expiration,omitempty"`
}

func (m *Grant) Reset()         { *m = Grant{} }
func (m *Grant) String() string { return proto.CompactTextString(m) }
func (*Grant) ProtoMessage()    {}
func (*Grant) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae337a09c25c97cc, []int{0}
}
func (m *Grant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Grant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Grant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Grant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Grant.Merge(m, src)
}
func (m *Grant) XXX_Size() int {
	return m.Size()
}
func (m *Grant) XXX_DiscardUnknown() {
	xxx_messageInfo_Grant.DiscardUnknown(m)
}

var xxx_messageInfo_Grant proto.InternalMessageInfo

func (m *Grant) GetAuthorization() *any.Any {
	if m != nil {
		return m.Authorization
	}
	return nil
}

func (m *Grant) GetExpiration() *time.Time {
	if m != nil {
		return m.Expiration
	}
	return nil
}

type GenericAuthoriztion struct {
	MsgTypeUrl string `protobuf:"bytes,1,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty"`
}

func (m *GenericAuthoriztion) Reset()         { *m = GenericAuthoriztion{} }
func (m *GenericAuthoriztion) String() string { return proto.CompactTextString(m) }
func (*GenericAuthoriztion) ProtoMessage()    {}
func (*GenericAuthoriztion) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae337a09c25c97cc, []int{1}
}
func (m *GenericAuthoriztion) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenericAuthoriztion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenericAuthoriztion.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenericAuthoriztion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericAuthoriztion.Merge(m, src)
}
func (m *GenericAuthoriztion) XXX_Size() int {
	return m.Size()
}
func (m *GenericAuthoriztion) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericAuthoriztion.DiscardUnknown(m)
}

var xxx_messageInfo_GenericAuthoriztion proto.InternalMessageInfo

func (m *GenericAuthoriztion) GetMsgTypeUrl() string {
	if m != nil {
		return m.MsgTypeUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*Grant)(nil), "cosmos.accounts.defaults.authz.Grant")
	proto.RegisterType((*GenericAuthoriztion)(nil), "cosmos.accounts.defaults.authz.GenericAuthoriztion")
}

func init() {
	proto.RegisterFile("cosmos/accounts/defaults/authz/authz.proto", fileDescriptor_ae337a09c25c97cc)
}

var fileDescriptor_ae337a09c25c97cc = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x4b, 0x3b, 0x31,
	0x18, 0xc6, 0x9b, 0x3f, 0xfc, 0x45, 0xa3, 0x2e, 0xb5, 0x83, 0x76, 0x48, 0x4b, 0x27, 0x11, 0x4d,
	0x44, 0x37, 0xb7, 0x16, 0xb1, 0x7b, 0xa9, 0x8b, 0x4b, 0x49, 0xaf, 0xe9, 0x19, 0xbc, 0xcb, 0x7b,
	0x24, 0x39, 0xe8, 0xf5, 0x53, 0xf4, 0xc3, 0xb8, 0xf9, 0x05, 0x8a, 0x53, 0x47, 0x27, 0x95, 0xf6,
	0x8b, 0xc8, 0x25, 0x0d, 0x54, 0x45, 0xc1, 0x25, 0x24, 0xbc, 0xcf, 0xef, 0x79, 0xde, 0x07, 0x82,
	0x4f, 0x22, 0x30, 0x29, 0x18, 0xc6, 0xa3, 0x08, 0x72, 0x65, 0x0d, 0x1b, 0x89, 0x31, 0xcf, 0x13,
	0x6b, 0x18, 0xcf, 0xed, 0xfd, 0xd4, 0x9f, 0x34, 0xd3, 0x60, 0xa1, 0x4a, 0xbc, 0x96, 0x06, 0x2d,
	0x0d, 0x5a, 0xea, 0x54, 0xf5, 0x5a, 0x0c, 0x31, 0x38, 0x29, 0x2b, 0x6f, 0x9e, 0xaa, 0x1f, 0x79,
	0x6a, 0xe0, 0x07, 0x6b, 0x0b, 0x3f, 0x6a, 0xc4, 0x00, 0x71, 0x22, 0x98, 0x7b, 0x0d, 0xf3, 0x31,
	0xb3, 0x32, 0x15, 0xc6, 0xf2, 0x34, 0x0b, 0xec, 0x57, 0x01, 0x57, 0x85, 0x1f, 0xb5, 0x9e, 0x10,
	0xfe, 0xdf, 0xd5, 0x5c, 0xd9, 0x6a, 0x82, 0xf7, 0xcb, 0x7c, 0xd0, 0x72, 0xca, 0xad, 0x04, 0x75,
	0x88, 0x9a, 0xe8, 0x78, 0xf7, 0xa2, 0x46, 0x3d, 0x4c, 0x03, 0x4c, 0xdb, 0xaa, 0xe8, 0x9c, 0x3f,
	0x3f, 0x9e, 0x9d, 0xfe, 0xde, 0x83, 0xb6, 0x37, 0xdd, 0x7a, 0x9f, 0xcd, 0xab, 0xd7, 0x18, 0x8b,
	0x49, 0x26, 0xb5, 0x8f, 0xfa, 0xe7, 0xa2, 0xea, 0xdf, 0xa2, 0xfa, 0xa1, 0x48, 0x67, 0x7b, 0xfe,
	0xda, 0x40, 0xb3, 0xb7, 0x06, 0xea, 0x6d, 0x70, 0x2d, 0x89, 0x0f, 0xba, 0x42, 0x09, 0x2d, 0xa3,
	0x10, 0xe6, 0xcc, 0x9b, 0x78, 0x2f, 0x35, 0xf1, 0xc0, 0x16, 0x99, 0x18, 0xe4, 0x3a, 0x71, 0x4d,
	0x76, 0x7a, 0x38, 0x35, 0x71, 0xbf, 0xc8, 0xc4, 0xad, 0x4e, 0xae, 0xfe, 0xbc, 0x7f, 0xe7, 0x66,
	0xbe, 0x24, 0x68, 0xb1, 0x24, 0xe8, 0x7d, 0x49, 0xd0, 0x6c, 0x45, 0x2a, 0x8b, 0x15, 0xa9, 0xbc,
	0xac, 0x48, 0xe5, 0x6e, 0xed, 0x63, 0x46, 0x0f, 0x54, 0x02, 0x9b, 0xfc, 0xf8, 0x07, 0xca, 0x75,
	0xcc, 0x70, 0xcb, 0x95, 0xbb, 0xfc, 0x08, 0x00, 0x00, 0xff, 0xff, 0x98, 0x21, 0x29, 0xf8, 0x32,
	0x02, 0x00, 0x00,
}

func (m *Grant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Grant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Grant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Expiration != nil {
		n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.Expiration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Expiration):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintAuthz(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x12
	}
	if m.Authorization != nil {
		{
			size, err := m.Authorization.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuthz(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenericAuthoriztion) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenericAuthoriztion) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenericAuthoriztion) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MsgTypeUrl) > 0 {
		i -= len(m.MsgTypeUrl)
		copy(dAtA[i:], m.MsgTypeUrl)
		i = encodeVarintAuthz(dAtA, i, uint64(len(m.MsgTypeUrl)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuthz(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthz(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Grant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Authorization != nil {
		l = m.Authorization.Size()
		n += 1 + l + sovAuthz(uint64(l))
	}
	if m.Expiration != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Expiration)
		n += 1 + l + sovAuthz(uint64(l))
	}
	return n
}

func (m *GenericAuthoriztion) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MsgTypeUrl)
	if l > 0 {
		n += 1 + l + sovAuthz(uint64(l))
	}
	return n
}

func sovAuthz(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthz(x uint64) (n int) {
	return sovAuthz(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Grant) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: Grant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Grant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authorization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Authorization == nil {
				m.Authorization = &any.Any{}
			}
			if err := m.Authorization.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expiration == nil {
				m.Expiration = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.Expiration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func (m *GenericAuthoriztion) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: GenericAuthoriztion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenericAuthoriztion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func skipAuthz(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
				return 0, ErrInvalidLengthAuthz
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthz
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthz
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthz        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthz          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthz = fmt.Errorf("proto: unexpected end of group")
)