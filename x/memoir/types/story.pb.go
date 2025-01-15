// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: memoir/memoir/story.proto

package types

import (
	fmt "fmt"
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

type Story struct {
	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Author   string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Id       uint64 `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Category string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	Rating   uint64 `protobuf:"varint,6,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (m *Story) Reset()         { *m = Story{} }
func (m *Story) String() string { return proto.CompactTextString(m) }
func (*Story) ProtoMessage()    {}
func (*Story) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfdbb23d078684b6, []int{0}
}
func (m *Story) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Story) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Story.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Story) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Story.Merge(m, src)
}
func (m *Story) XXX_Size() int {
	return m.Size()
}
func (m *Story) XXX_DiscardUnknown() {
	xxx_messageInfo_Story.DiscardUnknown(m)
}

var xxx_messageInfo_Story proto.InternalMessageInfo

func (m *Story) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Story) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Story) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Story) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Story) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Story) GetRating() uint64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func init() {
	proto.RegisterType((*Story)(nil), "memoir.memoir.Story")
}

func init() { proto.RegisterFile("memoir/memoir/story.proto", fileDescriptor_bfdbb23d078684b6) }

var fileDescriptor_bfdbb23d078684b6 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0x4d, 0xcd, 0xcd,
	0xcf, 0x2c, 0xd2, 0x87, 0x52, 0xc5, 0x25, 0xf9, 0x45, 0x95, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0xbc, 0x10, 0x31, 0x3d, 0x08, 0xa5, 0x34, 0x99, 0x91, 0x8b, 0x35, 0x18, 0x24, 0x2d, 0x24,
	0xc2, 0xc5, 0x5a, 0x92, 0x59, 0x92, 0x93, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1,
	0x08, 0x49, 0x70, 0xb1, 0x27, 0xe7, 0xe7, 0x95, 0xa4, 0xe6, 0x95, 0x48, 0x30, 0x81, 0xc5, 0x61,
	0x5c, 0x21, 0x31, 0x2e, 0xb6, 0xc4, 0xd2, 0x92, 0x8c, 0xfc, 0x22, 0x09, 0x66, 0xb0, 0x04, 0x94,
	0x27, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa2, 0xc0, 0xa8, 0xc1, 0x12, 0xc4, 0x94, 0x99,
	0x22, 0x24, 0xc5, 0xc5, 0x91, 0x9c, 0x58, 0x92, 0x9a, 0x9e, 0x5f, 0x54, 0x29, 0xc1, 0x0a, 0x56,
	0x09, 0xe7, 0x83, 0xcc, 0x28, 0x4a, 0x2c, 0xc9, 0xcc, 0x4b, 0x97, 0x60, 0x03, 0xab, 0x87, 0xf2,
	0x9c, 0xf4, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09,
	0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x14, 0xea, 0xa5,
	0x0a, 0x98, 0xdf, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x9e, 0x33, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x75, 0xf9, 0x1c, 0x0a, 0xf9, 0x00, 0x00, 0x00,
}

func (m *Story) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Story) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Story) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Rating != 0 {
		i = encodeVarintStory(dAtA, i, uint64(m.Rating))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Category) > 0 {
		i -= len(m.Category)
		copy(dAtA[i:], m.Category)
		i = encodeVarintStory(dAtA, i, uint64(len(m.Category)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Id != 0 {
		i = encodeVarintStory(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Author) > 0 {
		i -= len(m.Author)
		copy(dAtA[i:], m.Author)
		i = encodeVarintStory(dAtA, i, uint64(len(m.Author)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintStory(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintStory(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStory(dAtA []byte, offset int, v uint64) int {
	offset -= sovStory(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Story) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovStory(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovStory(uint64(l))
	}
	l = len(m.Author)
	if l > 0 {
		n += 1 + l + sovStory(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovStory(uint64(m.Id))
	}
	l = len(m.Category)
	if l > 0 {
		n += 1 + l + sovStory(uint64(l))
	}
	if m.Rating != 0 {
		n += 1 + sovStory(uint64(m.Rating))
	}
	return n
}

func sovStory(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStory(x uint64) (n int) {
	return sovStory(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Story) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStory
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
			return fmt.Errorf("proto: Story: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Story: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStory
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
				return ErrInvalidLengthStory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStory
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
				return ErrInvalidLengthStory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Author", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStory
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
				return ErrInvalidLengthStory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Author = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Category", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStory
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
				return ErrInvalidLengthStory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Category = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rating", wireType)
			}
			m.Rating = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rating |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStory
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
func skipStory(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStory
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
					return 0, ErrIntOverflowStory
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
					return 0, ErrIntOverflowStory
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
				return 0, ErrInvalidLengthStory
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStory
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStory
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStory        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStory          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStory = fmt.Errorf("proto: unexpected end of group")
)
