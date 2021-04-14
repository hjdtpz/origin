// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bin/gogopb/gogorpc.proto

package rpc

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

type GoGoPBRpcRequestData struct {
	Seq                  uint64   `protobuf:"varint,1,opt,name=Seq,proto3" json:"Seq,omitempty"`
	RpcMethodId          uint32   `protobuf:"varint,2,opt,name=RpcMethodId,proto3" json:"RpcMethodId,omitempty"`
	ServiceMethod        string   `protobuf:"bytes,3,opt,name=ServiceMethod,proto3" json:"ServiceMethod,omitempty"`
	NoReply              bool     `protobuf:"varint,4,opt,name=NoReply,proto3" json:"NoReply,omitempty"`
	InParam              []byte   `protobuf:"bytes,5,opt,name=InParam,proto3" json:"InParam,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoGoPBRpcRequestData) Reset()         { *m = GoGoPBRpcRequestData{} }
func (m *GoGoPBRpcRequestData) String() string { return proto.CompactTextString(m) }
func (*GoGoPBRpcRequestData) ProtoMessage()    {}
func (*GoGoPBRpcRequestData) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3b900b0f45d7fb5, []int{0}
}
func (m *GoGoPBRpcRequestData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoGoPBRpcRequestData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoGoPBRpcRequestData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoGoPBRpcRequestData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoGoPBRpcRequestData.Merge(m, src)
}
func (m *GoGoPBRpcRequestData) XXX_Size() int {
	return m.Size()
}
func (m *GoGoPBRpcRequestData) XXX_DiscardUnknown() {
	xxx_messageInfo_GoGoPBRpcRequestData.DiscardUnknown(m)
}

var xxx_messageInfo_GoGoPBRpcRequestData proto.InternalMessageInfo

func (m *GoGoPBRpcRequestData) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *GoGoPBRpcRequestData) GetRpcMethodId() uint32 {
	if m != nil {
		return m.RpcMethodId
	}
	return 0
}

func (m *GoGoPBRpcRequestData) GetServiceMethod() string {
	if m != nil {
		return m.ServiceMethod
	}
	return ""
}

func (m *GoGoPBRpcRequestData) GetNoReply() bool {
	if m != nil {
		return m.NoReply
	}
	return false
}

func (m *GoGoPBRpcRequestData) GetInParam() []byte {
	if m != nil {
		return m.InParam
	}
	return nil
}

type GoGoPBRpcResponseData struct {
	Seq                  uint64   `protobuf:"varint,1,opt,name=Seq,proto3" json:"Seq,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	Reply                []byte   `protobuf:"bytes,3,opt,name=Reply,proto3" json:"Reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoGoPBRpcResponseData) Reset()         { *m = GoGoPBRpcResponseData{} }
func (m *GoGoPBRpcResponseData) String() string { return proto.CompactTextString(m) }
func (*GoGoPBRpcResponseData) ProtoMessage()    {}
func (*GoGoPBRpcResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3b900b0f45d7fb5, []int{1}
}
func (m *GoGoPBRpcResponseData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoGoPBRpcResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoGoPBRpcResponseData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoGoPBRpcResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoGoPBRpcResponseData.Merge(m, src)
}
func (m *GoGoPBRpcResponseData) XXX_Size() int {
	return m.Size()
}
func (m *GoGoPBRpcResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_GoGoPBRpcResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_GoGoPBRpcResponseData proto.InternalMessageInfo

func (m *GoGoPBRpcResponseData) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *GoGoPBRpcResponseData) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GoGoPBRpcResponseData) GetReply() []byte {
	if m != nil {
		return m.Reply
	}
	return nil
}

func init() {
	proto.RegisterType((*GoGoPBRpcRequestData)(nil), "rpc.GoGoPBRpcRequestData")
	proto.RegisterType((*GoGoPBRpcResponseData)(nil), "rpc.GoGoPBRpcResponseData")
}

func init() { proto.RegisterFile("bin/gogopb/gogorpc.proto", fileDescriptor_b3b900b0f45d7fb5) }

var fileDescriptor_b3b900b0f45d7fb5 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xca, 0xcc, 0xd3,
	0x4f, 0xcf, 0x4f, 0xcf, 0x2f, 0x48, 0x02, 0x53, 0x45, 0x05, 0xc9, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0xcc, 0x45, 0x05, 0xc9, 0x4a, 0x4b, 0x18, 0xb9, 0x44, 0xdc, 0xf3, 0xdd, 0xf3, 0x03,
	0x9c, 0x82, 0x0a, 0x92, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x5c, 0x12, 0x4b, 0x12, 0x85,
	0x04, 0xb8, 0x98, 0x83, 0x53, 0x0b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0x40, 0x4c, 0x21,
	0x05, 0x2e, 0xee, 0xa0, 0x82, 0x64, 0xdf, 0xd4, 0x92, 0x8c, 0xfc, 0x14, 0xcf, 0x14, 0x09, 0x26,
	0x05, 0x46, 0x0d, 0xde, 0x20, 0x64, 0x21, 0x21, 0x15, 0x2e, 0xde, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc,
	0xe4, 0x54, 0x88, 0x90, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0xaa, 0xa0, 0x90, 0x04, 0x17,
	0xbb, 0x5f, 0x7e, 0x50, 0x6a, 0x41, 0x4e, 0xa5, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x8c,
	0x0b, 0x92, 0xf1, 0xcc, 0x0b, 0x48, 0x2c, 0x4a, 0xcc, 0x95, 0x60, 0x55, 0x60, 0xd4, 0xe0, 0x09,
	0x82, 0x71, 0x95, 0x42, 0xb9, 0x44, 0x91, 0x5c, 0x59, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x8a, 0xc3,
	0x99, 0x22, 0x5c, 0xac, 0xae, 0x45, 0x45, 0xf9, 0x45, 0x60, 0x07, 0x72, 0x06, 0x41, 0x38, 0x20,
	0x51, 0x88, 0x95, 0xcc, 0x60, 0x83, 0x21, 0x1c, 0x27, 0xe1, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c,
	0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x31, 0x8a, 0x55, 0x4f, 0xbf, 0xa8, 0x20, 0x39, 0x89, 0x0d,
	0x1c, 0x3c, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x18, 0x52, 0x1a, 0x3a, 0x01, 0x00,
	0x00,
}

func (m *GoGoPBRpcRequestData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoGoPBRpcRequestData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoGoPBRpcRequestData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.InParam) > 0 {
		i -= len(m.InParam)
		copy(dAtA[i:], m.InParam)
		i = encodeVarintGogorpc(dAtA, i, uint64(len(m.InParam)))
		i--
		dAtA[i] = 0x2a
	}
	if m.NoReply {
		i--
		if m.NoReply {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.ServiceMethod) > 0 {
		i -= len(m.ServiceMethod)
		copy(dAtA[i:], m.ServiceMethod)
		i = encodeVarintGogorpc(dAtA, i, uint64(len(m.ServiceMethod)))
		i--
		dAtA[i] = 0x1a
	}
	if m.RpcMethodId != 0 {
		i = encodeVarintGogorpc(dAtA, i, uint64(m.RpcMethodId))
		i--
		dAtA[i] = 0x10
	}
	if m.Seq != 0 {
		i = encodeVarintGogorpc(dAtA, i, uint64(m.Seq))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GoGoPBRpcResponseData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoGoPBRpcResponseData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoGoPBRpcResponseData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Reply) > 0 {
		i -= len(m.Reply)
		copy(dAtA[i:], m.Reply)
		i = encodeVarintGogorpc(dAtA, i, uint64(len(m.Reply)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintGogorpc(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x12
	}
	if m.Seq != 0 {
		i = encodeVarintGogorpc(dAtA, i, uint64(m.Seq))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGogorpc(dAtA []byte, offset int, v uint64) int {
	offset -= sovGogorpc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GoGoPBRpcRequestData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Seq != 0 {
		n += 1 + sovGogorpc(uint64(m.Seq))
	}
	if m.RpcMethodId != 0 {
		n += 1 + sovGogorpc(uint64(m.RpcMethodId))
	}
	l = len(m.ServiceMethod)
	if l > 0 {
		n += 1 + l + sovGogorpc(uint64(l))
	}
	if m.NoReply {
		n += 2
	}
	l = len(m.InParam)
	if l > 0 {
		n += 1 + l + sovGogorpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GoGoPBRpcResponseData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Seq != 0 {
		n += 1 + sovGogorpc(uint64(m.Seq))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovGogorpc(uint64(l))
	}
	l = len(m.Reply)
	if l > 0 {
		n += 1 + l + sovGogorpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGogorpc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGogorpc(x uint64) (n int) {
	return sovGogorpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GoGoPBRpcRequestData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGogorpc
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
			return fmt.Errorf("proto: GoGoPBRpcRequestData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoGoPBRpcRequestData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seq", wireType)
			}
			m.Seq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogorpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RpcMethodId", wireType)
			}
			m.RpcMethodId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogorpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RpcMethodId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogorpc
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
				return ErrInvalidLengthGogorpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGogorpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceMethod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoReply", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogorpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NoReply = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InParam", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogorpc
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
				return ErrInvalidLengthGogorpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGogorpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InParam = append(m.InParam[:0], dAtA[iNdEx:postIndex]...)
			if m.InParam == nil {
				m.InParam = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGogorpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGogorpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGogorpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GoGoPBRpcResponseData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGogorpc
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
			return fmt.Errorf("proto: GoGoPBRpcResponseData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoGoPBRpcResponseData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seq", wireType)
			}
			m.Seq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogorpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogorpc
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
				return ErrInvalidLengthGogorpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGogorpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reply", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogorpc
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
				return ErrInvalidLengthGogorpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGogorpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reply = append(m.Reply[:0], dAtA[iNdEx:postIndex]...)
			if m.Reply == nil {
				m.Reply = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGogorpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGogorpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGogorpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGogorpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGogorpc
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
					return 0, ErrIntOverflowGogorpc
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
					return 0, ErrIntOverflowGogorpc
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
				return 0, ErrInvalidLengthGogorpc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGogorpc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGogorpc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGogorpc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGogorpc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGogorpc = fmt.Errorf("proto: unexpected end of group")
)