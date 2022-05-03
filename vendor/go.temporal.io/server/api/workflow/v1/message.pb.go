// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/server/api/workflow/v1/message.proto

package workflow

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	proto "github.com/gogo/protobuf/proto"
	v1 "go.temporal.io/api/common/v1"
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

type ParentExecutionInfo struct {
	NamespaceId string                `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	Namespace   string                `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Execution   *v1.WorkflowExecution `protobuf:"bytes,3,opt,name=execution,proto3" json:"execution,omitempty"`
	InitiatedId int64                 `protobuf:"varint,4,opt,name=initiated_id,json=initiatedId,proto3" json:"initiated_id,omitempty"`
}

func (m *ParentExecutionInfo) Reset()      { *m = ParentExecutionInfo{} }
func (*ParentExecutionInfo) ProtoMessage() {}
func (*ParentExecutionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f1ca48d03c9ded, []int{0}
}
func (m *ParentExecutionInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParentExecutionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParentExecutionInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParentExecutionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParentExecutionInfo.Merge(m, src)
}
func (m *ParentExecutionInfo) XXX_Size() int {
	return m.Size()
}
func (m *ParentExecutionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ParentExecutionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ParentExecutionInfo proto.InternalMessageInfo

func (m *ParentExecutionInfo) GetNamespaceId() string {
	if m != nil {
		return m.NamespaceId
	}
	return ""
}

func (m *ParentExecutionInfo) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ParentExecutionInfo) GetExecution() *v1.WorkflowExecution {
	if m != nil {
		return m.Execution
	}
	return nil
}

func (m *ParentExecutionInfo) GetInitiatedId() int64 {
	if m != nil {
		return m.InitiatedId
	}
	return 0
}

func init() {
	proto.RegisterType((*ParentExecutionInfo)(nil), "temporal.server.api.workflow.v1.ParentExecutionInfo")
}

func init() {
	proto.RegisterFile("temporal/server/api/workflow/v1/message.proto", fileDescriptor_c4f1ca48d03c9ded)
}

var fileDescriptor_c4f1ca48d03c9ded = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x18, 0x85, 0xfd, 0x53, 0x84, 0xd4, 0x94, 0x29, 0x2c, 0x15, 0x42, 0x3f, 0x05, 0x31, 0x94, 0x01,
	0x47, 0x85, 0x91, 0x0d, 0x09, 0xa1, 0x6c, 0x28, 0x0b, 0x12, 0x0b, 0x32, 0x89, 0x5b, 0x59, 0x34,
	0x76, 0xe4, 0x98, 0x94, 0x91, 0x23, 0x70, 0x0c, 0xce, 0xc1, 0xc4, 0x98, 0xb1, 0x23, 0x71, 0x16,
	0xc6, 0x1e, 0x01, 0xa5, 0x25, 0xce, 0x80, 0xd8, 0xec, 0xe7, 0xe7, 0x4f, 0x9f, 0x9e, 0x77, 0x66,
	0x78, 0x9a, 0x29, 0xcd, 0xe6, 0x41, 0xce, 0x75, 0xc1, 0x75, 0xc0, 0x32, 0x11, 0x2c, 0x94, 0x7e,
	0x9a, 0xce, 0xd5, 0x22, 0x28, 0x26, 0x41, 0xca, 0xf3, 0x9c, 0xcd, 0x38, 0xcd, 0xb4, 0x32, 0xca,
	0x3f, 0x6c, 0xeb, 0x74, 0x53, 0xa7, 0x2c, 0x13, 0xb4, 0xad, 0xd3, 0x62, 0xb2, 0x7f, 0xe2, 0x78,
	0x0d, 0x28, 0x56, 0x69, 0xaa, 0xe4, 0x1f, 0xcc, 0xf1, 0x07, 0x78, 0x7b, 0xb7, 0x4c, 0x73, 0x69,
	0xae, 0x5f, 0x78, 0xfc, 0x6c, 0x84, 0x92, 0xa1, 0x9c, 0x2a, 0xff, 0xc8, 0xdb, 0x95, 0x2c, 0xe5,
	0x79, 0xc6, 0x62, 0xfe, 0x20, 0x92, 0x21, 0x8c, 0x60, 0xdc, 0x8f, 0x06, 0x2e, 0x0b, 0x13, 0xff,
	0xc0, 0xeb, 0xbb, 0xeb, 0x70, 0x6b, 0xfd, 0xde, 0x05, 0xfe, 0x8d, 0xd7, 0xe7, 0x2d, 0x71, 0xd8,
	0x1b, 0xc1, 0x78, 0x70, 0x7e, 0x4a, 0x9d, 0x73, 0x23, 0xbb, 0x51, 0xa2, 0xc5, 0x84, 0xde, 0xfd,
	0x6a, 0x3b, 0x85, 0xa8, 0xfb, 0xdb, 0x98, 0x08, 0x29, 0x8c, 0x60, 0x86, 0x27, 0x8d, 0xc9, 0xf6,
	0x08, 0xc6, 0xbd, 0x68, 0xe0, 0xb2, 0x30, 0xb9, 0x4a, 0xca, 0x0a, 0xc9, 0xb2, 0x42, 0xb2, 0xaa,
	0x10, 0x5e, 0x2d, 0xc2, 0xbb, 0x45, 0xf8, 0xb4, 0x08, 0xa5, 0x45, 0xf8, 0xb2, 0x08, 0xdf, 0x16,
	0xc9, 0xca, 0x22, 0xbc, 0xd5, 0x48, 0xca, 0x1a, 0xc9, 0xb2, 0x46, 0x72, 0x4f, 0x67, 0xaa, 0x13,
	0x12, 0xea, 0x9f, 0xd9, 0x2f, 0xdb, 0xf3, 0xe3, 0xce, 0x7a, 0xb1, 0x8b, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd4, 0x7b, 0x8c, 0xc8, 0xa9, 0x01, 0x00, 0x00,
}

func (this *ParentExecutionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ParentExecutionInfo)
	if !ok {
		that2, ok := that.(ParentExecutionInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.NamespaceId != that1.NamespaceId {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if !this.Execution.Equal(that1.Execution) {
		return false
	}
	if this.InitiatedId != that1.InitiatedId {
		return false
	}
	return true
}
func (this *ParentExecutionInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&workflow.ParentExecutionInfo{")
	s = append(s, "NamespaceId: "+fmt.Sprintf("%#v", this.NamespaceId)+",\n")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	if this.Execution != nil {
		s = append(s, "Execution: "+fmt.Sprintf("%#v", this.Execution)+",\n")
	}
	s = append(s, "InitiatedId: "+fmt.Sprintf("%#v", this.InitiatedId)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMessage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ParentExecutionInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParentExecutionInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParentExecutionInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InitiatedId != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.InitiatedId))
		i--
		dAtA[i] = 0x20
	}
	if m.Execution != nil {
		{
			size, err := m.Execution.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.NamespaceId) > 0 {
		i -= len(m.NamespaceId)
		copy(dAtA[i:], m.NamespaceId)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.NamespaceId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ParentExecutionInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NamespaceId)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.Execution != nil {
		l = m.Execution.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.InitiatedId != 0 {
		n += 1 + sovMessage(uint64(m.InitiatedId))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ParentExecutionInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ParentExecutionInfo{`,
		`NamespaceId:` + fmt.Sprintf("%v", this.NamespaceId) + `,`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Execution:` + strings.Replace(fmt.Sprintf("%v", this.Execution), "WorkflowExecution", "v1.WorkflowExecution", 1) + `,`,
		`InitiatedId:` + fmt.Sprintf("%v", this.InitiatedId) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMessage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ParentExecutionInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: ParentExecutionInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParentExecutionInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NamespaceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NamespaceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Execution", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Execution == nil {
				m.Execution = &v1.WorkflowExecution{}
			}
			if err := m.Execution.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitiatedId", wireType)
			}
			m.InitiatedId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitiatedId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)
