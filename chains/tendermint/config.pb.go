// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: relayer/tendermint/config.proto

package tendermint

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
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

type ChainConfig struct {
	Key            string  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ChainId        string  `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	RpcAddr        string  `protobuf:"bytes,3,opt,name=rpc_addr,json=rpcAddr,proto3" json:"rpc_addr,omitempty"`
	AccountPrefix  string  `protobuf:"bytes,4,opt,name=account_prefix,json=accountPrefix,proto3" json:"account_prefix,omitempty"`
	GasAdjustment  float32 `protobuf:"fixed32,5,opt,name=gas_adjustment,json=gasAdjustment,proto3" json:"gas_adjustment,omitempty"`
	GasPrices      string  `protobuf:"bytes,6,opt,name=gas_prices,json=gasPrices,proto3" json:"gas_prices,omitempty"`
	TrustingPeriod string  `protobuf:"bytes,7,opt,name=trusting_period,json=trustingPeriod,proto3" json:"trusting_period,omitempty"`
}

func (m *ChainConfig) Reset()         { *m = ChainConfig{} }
func (m *ChainConfig) String() string { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()    {}
func (*ChainConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_db89c5c80bed2ca6, []int{0}
}
func (m *ChainConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainConfig.Merge(m, src)
}
func (m *ChainConfig) XXX_Size() int {
	return m.Size()
}
func (m *ChainConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ChainConfig proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ChainConfig)(nil), "relayer.tendermint.config.ChainConfig")
}

func init() { proto.RegisterFile("relayer/tendermint/config.proto", fileDescriptor_db89c5c80bed2ca6) }

var fileDescriptor_db89c5c80bed2ca6 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4e, 0x2a, 0x31,
	0x18, 0x85, 0xa7, 0x70, 0x2f, 0x48, 0x0d, 0x68, 0x26, 0x2e, 0x8a, 0x89, 0x95, 0x98, 0x18, 0x59,
	0xcd, 0xc4, 0xf8, 0x04, 0xc8, 0xca, 0x95, 0x84, 0xa5, 0x9b, 0x49, 0x69, 0x4b, 0xa9, 0x42, 0x3b,
	0x69, 0x3b, 0x89, 0xbc, 0x85, 0x8f, 0xc5, 0x92, 0xa5, 0x4b, 0x85, 0xa5, 0x2f, 0x61, 0xfa, 0x03,
	0xea, 0xee, 0xf4, 0x3b, 0x5f, 0x9b, 0xf4, 0xe0, 0x4b, 0x27, 0xe7, 0x6c, 0x29, 0x5d, 0x1e, 0xa4,
	0x11, 0xd2, 0x2d, 0xb4, 0x09, 0x39, 0xb7, 0x66, 0xaa, 0x55, 0x56, 0x3a, 0x1b, 0x6c, 0xda, 0xdd,
	0x0b, 0xd9, 0xaf, 0x90, 0xed, 0x84, 0xf3, 0x33, 0x65, 0x95, 0x05, 0x2b, 0x8f, 0x69, 0x77, 0xe1,
	0xea, 0x0b, 0xe1, 0xe3, 0xe1, 0x8c, 0x69, 0x33, 0x04, 0x2b, 0x3d, 0xc5, 0xf5, 0x17, 0xb9, 0x24,
	0xa8, 0x87, 0xfa, 0xad, 0x71, 0x8c, 0x69, 0x17, 0x1f, 0xf1, 0x28, 0x14, 0x5a, 0x90, 0x1a, 0xe0,
	0x26, 0x9c, 0x1f, 0x44, 0xac, 0x5c, 0xc9, 0x0b, 0x26, 0x84, 0x23, 0xf5, 0x5d, 0xe5, 0x4a, 0x3e,
	0x10, 0xc2, 0xa5, 0xd7, 0xb8, 0xc3, 0x38, 0xb7, 0x95, 0x09, 0x45, 0xe9, 0xe4, 0x54, 0xbf, 0x92,
	0x7f, 0x20, 0xb4, 0xf7, 0x74, 0x04, 0x30, 0x6a, 0x8a, 0xf9, 0x82, 0x89, 0xe7, 0xca, 0x87, 0x85,
	0x34, 0x81, 0xfc, 0xef, 0xa1, 0x7e, 0x6d, 0xdc, 0x56, 0xcc, 0x0f, 0x7e, 0x60, 0x7a, 0x81, 0x71,
	0xd4, 0x4a, 0xa7, 0xb9, 0xf4, 0xa4, 0x01, 0x2f, 0xb5, 0x14, 0xf3, 0x23, 0x00, 0xe9, 0x0d, 0x3e,
	0x09, 0xae, 0xf2, 0x41, 0x1b, 0x55, 0x94, 0xd2, 0x69, 0x2b, 0x48, 0x13, 0x9c, 0xce, 0x01, 0x8f,
	0x80, 0xde, 0x3f, 0xae, 0x3e, 0x69, 0xb2, 0xda, 0x50, 0xb4, 0xde, 0x50, 0xf4, 0xb1, 0xa1, 0xe8,
	0x6d, 0x4b, 0x93, 0xf5, 0x96, 0x26, 0xef, 0x5b, 0x9a, 0x3c, 0xdd, 0x2a, 0x1d, 0x66, 0xd5, 0x24,
	0xe3, 0x76, 0x91, 0x0b, 0x16, 0x18, 0x7c, 0x73, 0xce, 0x26, 0xf9, 0x61, 0x75, 0x00, 0xfe, 0xcf,
	0xf8, 0x93, 0x06, 0xac, 0x78, 0xf7, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x29, 0x48, 0x6e, 0x99,
	0x01, 0x00, 0x00,
}

func (m *ChainConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TrustingPeriod) > 0 {
		i -= len(m.TrustingPeriod)
		copy(dAtA[i:], m.TrustingPeriod)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.TrustingPeriod)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.GasPrices) > 0 {
		i -= len(m.GasPrices)
		copy(dAtA[i:], m.GasPrices)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.GasPrices)))
		i--
		dAtA[i] = 0x32
	}
	if m.GasAdjustment != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.GasAdjustment))))
		i--
		dAtA[i] = 0x2d
	}
	if len(m.AccountPrefix) > 0 {
		i -= len(m.AccountPrefix)
		copy(dAtA[i:], m.AccountPrefix)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.AccountPrefix)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RpcAddr) > 0 {
		i -= len(m.RpcAddr)
		copy(dAtA[i:], m.RpcAddr)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.RpcAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChainConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.RpcAddr)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.AccountPrefix)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.GasAdjustment != 0 {
		n += 5
	}
	l = len(m.GasPrices)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.TrustingPeriod)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: ChainConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RpcAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RpcAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasAdjustment", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.GasAdjustment = float32(math.Float32frombits(v))
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrices", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasPrices = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustingPeriod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrustingPeriod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfig = fmt.Errorf("proto: unexpected end of group")
)