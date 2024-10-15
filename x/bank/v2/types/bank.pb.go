// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/bank/v2/bank.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters for the bank/v2 module.
type Params struct {
	// DenomCreationFee defines the fee to be charged on the creation of a new
	// denom. The fee is drawn from the MsgCreateDenom's sender account, and
	// transferred to the community pool.
	DenomCreationFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=denom_creation_fee,json=denomCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"denom_creation_fee" yaml:"denom_creation_fee"`
	// DenomCreationGasConsume defines the gas cost for creating a new denom.
	// This is intended as a spam deterrence mechanism.
	//
	// See: https://github.com/CosmWasm/token-factory/issues/11
	DenomCreationGasConsume uint64 `protobuf:"varint,2,opt,name=denom_creation_gas_consume,json=denomCreationGasConsume,proto3" json:"denom_creation_gas_consume,omitempty" yaml:"denom_creation_gas_consume"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e0dfb4485ca624d, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetDenomCreationFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.DenomCreationFee
	}
	return nil
}

func (m *Params) GetDenomCreationGasConsume() uint64 {
	if m != nil {
		return m.DenomCreationGasConsume
	}
	return 0
}

// DenomUnit represents a struct that describes a given
// denomination unit of the basic token.
type DenomUnit struct {
	// denom represents the string name of the given denom unit (e.g uatom).
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// exponent represents power of 10 exponent that one must
	// raise the base_denom to in order to equal the given DenomUnit's denom
	// 1 denom = 10^exponent base_denom
	// (e.g. with a base_denom of uatom, one can create a DenomUnit of 'atom' with
	// exponent = 6, thus: 1 atom = 10^6 uatom).
	Exponent uint32 `protobuf:"varint,2,opt,name=exponent,proto3" json:"exponent,omitempty"`
	// aliases is a list of string aliases for the given denom
	Aliases []string `protobuf:"bytes,3,rep,name=aliases,proto3" json:"aliases,omitempty"`
}

func (m *DenomUnit) Reset()         { *m = DenomUnit{} }
func (m *DenomUnit) String() string { return proto.CompactTextString(m) }
func (*DenomUnit) ProtoMessage()    {}
func (*DenomUnit) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e0dfb4485ca624d, []int{1}
}
func (m *DenomUnit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomUnit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomUnit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomUnit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomUnit.Merge(m, src)
}
func (m *DenomUnit) XXX_Size() int {
	return m.Size()
}
func (m *DenomUnit) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomUnit.DiscardUnknown(m)
}

var xxx_messageInfo_DenomUnit proto.InternalMessageInfo

func (m *DenomUnit) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *DenomUnit) GetExponent() uint32 {
	if m != nil {
		return m.Exponent
	}
	return 0
}

func (m *DenomUnit) GetAliases() []string {
	if m != nil {
		return m.Aliases
	}
	return nil
}

// Metadata represents a struct that describes
// a basic token.
type Metadata struct {
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// denom_units represents the list of DenomUnit's for a given coin
	DenomUnits []*DenomUnit `protobuf:"bytes,2,rep,name=denom_units,json=denomUnits,proto3" json:"denom_units,omitempty"`
	// base represents the base denom (should be the DenomUnit with exponent = 0).
	Base string `protobuf:"bytes,3,opt,name=base,proto3" json:"base,omitempty"`
	// display indicates the suggested denom that should be
	// displayed in clients.
	Display string `protobuf:"bytes,4,opt,name=display,proto3" json:"display,omitempty"`
	// name defines the name of the token (eg: Cosmos Atom)
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// symbol is the token symbol usually shown on exchanges (eg: ATOM). This can
	// be the same as the display.
	Symbol string `protobuf:"bytes,6,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// URI to a document (on or off-chain) that contains additional information. Optional.
	URI string `protobuf:"bytes,7,opt,name=uri,proto3" json:"uri,omitempty"`
	// URIHash is a sha256 hash of a document pointed by URI. It's used to verify that
	// the document didn't change. Optional.
	URIHash string `protobuf:"bytes,8,opt,name=uri_hash,json=uriHash,proto3" json:"uri_hash,omitempty"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e0dfb4485ca624d, []int{2}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return m.Size()
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Metadata) GetDenomUnits() []*DenomUnit {
	if m != nil {
		return m.DenomUnits
	}
	return nil
}

func (m *Metadata) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *Metadata) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

func (m *Metadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metadata) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Metadata) GetURI() string {
	if m != nil {
		return m.URI
	}
	return ""
}

func (m *Metadata) GetURIHash() string {
	if m != nil {
		return m.URIHash
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "cosmos.bank.v2.Params")
	proto.RegisterType((*DenomUnit)(nil), "cosmos.bank.v2.DenomUnit")
	proto.RegisterType((*Metadata)(nil), "cosmos.bank.v2.Metadata")
}

func init() { proto.RegisterFile("cosmos/bank/v2/bank.proto", fileDescriptor_2e0dfb4485ca624d) }

var fileDescriptor_2e0dfb4485ca624d = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0xe3, 0x34, 0x3f, 0x5b, 0xf1, 0xa3, 0xa5, 0x12, 0x9b, 0x08, 0xd9, 0x26, 0x17, 0x82,
	0xa0, 0x36, 0x4d, 0x51, 0x25, 0x7a, 0x4c, 0x10, 0xd0, 0x43, 0x25, 0x64, 0x29, 0x42, 0xe2, 0x12,
	0x6d, 0xec, 0x6d, 0xb2, 0x4a, 0xbc, 0x1b, 0x79, 0xec, 0xa8, 0x79, 0x04, 0x6e, 0x9c, 0x78, 0x08,
	0xce, 0x3c, 0x44, 0x8e, 0xbd, 0x20, 0x21, 0x0e, 0x06, 0x25, 0x6f, 0xd0, 0x27, 0x40, 0x5e, 0x3b,
	0xa6, 0xb4, 0xe5, 0xe4, 0x9d, 0xf9, 0x7e, 0xe6, 0x93, 0x66, 0x8c, 0x9a, 0x9e, 0x84, 0x40, 0x82,
	0x33, 0xa2, 0x62, 0xea, 0x2c, 0xba, 0xea, 0x6b, 0xcf, 0x43, 0x19, 0x49, 0x7c, 0x37, 0x83, 0x6c,
	0xd5, 0x5a, 0x74, 0x5b, 0x7b, 0x63, 0x39, 0x96, 0x0a, 0x72, 0xd2, 0x57, 0xc6, 0x6a, 0xe5, 0x06,
	0xc3, 0x0c, 0xc8, 0x25, 0x19, 0x64, 0x14, 0xde, 0xc0, 0x9c, 0xc5, 0xc1, 0x88, 0x45, 0xf4, 0xc0,
	0xf1, 0x24, 0x17, 0x19, 0xde, 0xfe, 0x54, 0x46, 0xd5, 0xf7, 0x34, 0xa4, 0x01, 0xe0, 0x2f, 0x1a,
	0xc2, 0x3e, 0x13, 0x32, 0x18, 0x7a, 0x21, 0xa3, 0x11, 0x97, 0x62, 0x78, 0xc6, 0x18, 0xd1, 0x2c,
	0xbd, 0xb3, 0xdb, 0x6d, 0xda, 0x45, 0x12, 0x60, 0x76, 0x6e, 0x64, 0xf7, 0x25, 0x17, 0xbd, 0xd3,
	0x55, 0x62, 0x96, 0x2e, 0x13, 0xb3, 0xb9, 0xa4, 0xc1, 0xec, 0xb8, 0x7d, 0xd3, 0xa2, 0xfd, 0xf5,
	0x97, 0xd9, 0x19, 0xf3, 0x68, 0x12, 0x8f, 0x6c, 0x4f, 0x06, 0x79, 0xc0, 0xfc, 0xb3, 0x0f, 0xfe,
	0xd4, 0x89, 0x96, 0x73, 0x06, 0xca, 0x0d, 0xdc, 0xfb, 0xca, 0xa0, 0x9f, 0xeb, 0xdf, 0x30, 0x86,
	0xcf, 0x50, 0xeb, 0x9a, 0xe9, 0x98, 0xc2, 0xd0, 0x93, 0x02, 0xe2, 0x80, 0x91, 0xb2, 0xa5, 0x75,
	0x2a, 0xbd, 0xa7, 0xab, 0xc4, 0xd4, 0x2e, 0x13, 0xf3, 0xf1, 0xad, 0x21, 0xae, 0xf0, 0xdb, 0xee,
	0xc3, 0x7f, 0x06, 0xbc, 0xa5, 0xd0, 0xcf, 0x91, 0x0f, 0xa8, 0xf1, 0x3a, 0x85, 0x06, 0x82, 0x47,
	0x78, 0x0f, 0xed, 0x28, 0x1e, 0xd1, 0x2c, 0xad, 0xd3, 0x70, 0xb3, 0x02, 0xb7, 0x50, 0x9d, 0x9d,
	0xcf, 0xa5, 0x60, 0x22, 0x52, 0x83, 0xef, 0xb8, 0x45, 0x8d, 0x09, 0xaa, 0xd1, 0x19, 0xa7, 0xc0,
	0x80, 0xe8, 0x96, 0xde, 0x69, 0xb8, 0xdb, 0xb2, 0xfd, 0xbd, 0x8c, 0xea, 0xa7, 0x2c, 0xa2, 0x3e,
	0x8d, 0x28, 0xb6, 0xd0, 0xae, 0xcf, 0xc0, 0x0b, 0xf9, 0x3c, 0x1d, 0x9f, 0xdb, 0x5f, 0x6d, 0xe1,
	0xe3, 0x94, 0x91, 0xe6, 0x8f, 0x05, 0x8f, 0x80, 0x94, 0xaf, 0x2f, 0x40, 0x9d, 0x82, 0x5d, 0x44,
	0x75, 0x91, 0xbf, 0x7d, 0x02, 0xc6, 0xa8, 0x92, 0x6e, 0x88, 0xe8, 0xca, 0x56, 0xbd, 0xd3, 0x60,
	0x3e, 0x87, 0xf9, 0x8c, 0x2e, 0x49, 0x45, 0xb5, 0xb7, 0x25, 0x7e, 0x82, 0x2a, 0x82, 0x06, 0x8c,
	0xec, 0xa4, 0xed, 0xde, 0x83, 0x9f, 0xdf, 0xf6, 0xef, 0xfd, 0xdd, 0x8a, 0xf5, 0xc2, 0x7e, 0x79,
	0xe8, 0x2a, 0x02, 0x7e, 0x86, 0xaa, 0xb0, 0x0c, 0x46, 0x72, 0x46, 0xaa, 0xff, 0xa7, 0xe6, 0x14,
	0xfc, 0x1c, 0xe9, 0x71, 0xc8, 0x49, 0x4d, 0x31, 0x5b, 0xeb, 0xc4, 0xd4, 0x07, 0xee, 0xc9, 0x4d,
	0xc1, 0x91, 0x9b, 0xd2, 0xf0, 0x2b, 0x54, 0x8f, 0x43, 0x3e, 0x9c, 0x50, 0x98, 0x90, 0xba, 0x92,
	0x18, 0xeb, 0xc4, 0xac, 0x0d, 0xdc, 0x93, 0x77, 0x14, 0x26, 0xb7, 0xc9, 0x6a, 0x71, 0xc8, 0x53,
	0xac, 0x77, 0xb4, 0x5a, 0x1b, 0xda, 0xc5, 0xda, 0xd0, 0x7e, 0xaf, 0x0d, 0xed, 0xf3, 0xc6, 0x28,
	0x5d, 0x6c, 0x8c, 0xd2, 0x8f, 0x8d, 0x51, 0xfa, 0xf8, 0x28, 0xd3, 0x80, 0x3f, 0xb5, 0xb9, 0x74,
	0xce, 0x8b, 0x5f, 0x4b, 0x1d, 0xda, 0xa8, 0xaa, 0x6e, 0xff, 0xf0, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc2, 0xe8, 0x08, 0x7c, 0x79, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DenomCreationGasConsume != 0 {
		i = encodeVarintBank(dAtA, i, uint64(m.DenomCreationGasConsume))
		i--
		dAtA[i] = 0x10
	}
	if len(m.DenomCreationFee) > 0 {
		for iNdEx := len(m.DenomCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBank(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *DenomUnit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomUnit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomUnit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Aliases) > 0 {
		for iNdEx := len(m.Aliases) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Aliases[iNdEx])
			copy(dAtA[i:], m.Aliases[iNdEx])
			i = encodeVarintBank(dAtA, i, uint64(len(m.Aliases[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Exponent != 0 {
		i = encodeVarintBank(dAtA, i, uint64(m.Exponent))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintBank(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.URIHash) > 0 {
		i -= len(m.URIHash)
		copy(dAtA[i:], m.URIHash)
		i = encodeVarintBank(dAtA, i, uint64(len(m.URIHash)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.URI) > 0 {
		i -= len(m.URI)
		copy(dAtA[i:], m.URI)
		i = encodeVarintBank(dAtA, i, uint64(len(m.URI)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintBank(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintBank(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Display) > 0 {
		i -= len(m.Display)
		copy(dAtA[i:], m.Display)
		i = encodeVarintBank(dAtA, i, uint64(len(m.Display)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Base) > 0 {
		i -= len(m.Base)
		copy(dAtA[i:], m.Base)
		i = encodeVarintBank(dAtA, i, uint64(len(m.Base)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DenomUnits) > 0 {
		for iNdEx := len(m.DenomUnits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomUnits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBank(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintBank(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBank(dAtA []byte, offset int, v uint64) int {
	offset -= sovBank(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DenomCreationFee) > 0 {
		for _, e := range m.DenomCreationFee {
			l = e.Size()
			n += 1 + l + sovBank(uint64(l))
		}
	}
	if m.DenomCreationGasConsume != 0 {
		n += 1 + sovBank(uint64(m.DenomCreationGasConsume))
	}
	return n
}

func (m *DenomUnit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovBank(uint64(l))
	}
	if m.Exponent != 0 {
		n += 1 + sovBank(uint64(m.Exponent))
	}
	if len(m.Aliases) > 0 {
		for _, s := range m.Aliases {
			l = len(s)
			n += 1 + l + sovBank(uint64(l))
		}
	}
	return n
}

func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovBank(uint64(l))
	}
	if len(m.DenomUnits) > 0 {
		for _, e := range m.DenomUnits {
			l = e.Size()
			n += 1 + l + sovBank(uint64(l))
		}
	}
	l = len(m.Base)
	if l > 0 {
		n += 1 + l + sovBank(uint64(l))
	}
	l = len(m.Display)
	if l > 0 {
		n += 1 + l + sovBank(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovBank(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovBank(uint64(l))
	}
	l = len(m.URI)
	if l > 0 {
		n += 1 + l + sovBank(uint64(l))
	}
	l = len(m.URIHash)
	if l > 0 {
		n += 1 + l + sovBank(uint64(l))
	}
	return n
}

func sovBank(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBank(x uint64) (n int) {
	return sovBank(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBank
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomCreationFee = append(m.DenomCreationFee, types.Coin{})
			if err := m.DenomCreationFee[len(m.DenomCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomCreationGasConsume", wireType)
			}
			m.DenomCreationGasConsume = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DenomCreationGasConsume |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBank(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBank
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
func (m *DenomUnit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBank
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
			return fmt.Errorf("proto: DenomUnit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomUnit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exponent", wireType)
			}
			m.Exponent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Exponent |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Aliases", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Aliases = append(m.Aliases, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBank(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBank
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
func (m *Metadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBank
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomUnits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomUnits = append(m.DenomUnits, &DenomUnit{})
			if err := m.DenomUnits[len(m.DenomUnits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Base = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Display", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Display = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URIHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URIHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBank(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBank
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
func skipBank(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBank
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
					return 0, ErrIntOverflowBank
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
					return 0, ErrIntOverflowBank
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
				return 0, ErrInvalidLengthBank
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBank
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBank
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBank        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBank          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBank = fmt.Errorf("proto: unexpected end of group")
)
