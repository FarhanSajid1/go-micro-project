// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vessel-service/proto/vessel/vessel.proto

package go_micro_srv_vessel

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwnerId              string   `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_986c3e1f103d74ba, []int{0}
}

func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (m *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(m, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_986c3e1f103d74ba, []int{1}
}

func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (m *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(m, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	// take in the vessel
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels,proto3" json:"vessels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_986c3e1f103d74ba, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func init() {
	proto.RegisterType((*Vessel)(nil), "go.micro.srv.vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "go.micro.srv.vessel.Specification")
	proto.RegisterType((*Response)(nil), "go.micro.srv.vessel.Response")
}

func init() {
	proto.RegisterFile("vessel-service/proto/vessel/vessel.proto", fileDescriptor_986c3e1f103d74ba)
}

var fileDescriptor_986c3e1f103d74ba = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0xfd, 0xb6, 0x40, 0x29, 0xf3, 0x05, 0x0f, 0xeb, 0x65, 0x45, 0x49, 0x9a, 0x9e, 0x7a, 0x71,
	0x49, 0x20, 0xfe, 0x00, 0x2f, 0x26, 0x7a, 0x2c, 0x46, 0x8f, 0x64, 0xd9, 0x8e, 0x38, 0x09, 0xed,
	0x36, 0xdd, 0xa6, 0xe0, 0xbf, 0xf1, 0xa7, 0x9a, 0xec, 0x02, 0x06, 0xd3, 0xe8, 0x69, 0x67, 0xdf,
	0xbc, 0x79, 0x79, 0xf3, 0x06, 0xd2, 0x16, 0xad, 0xc5, 0xed, 0xad, 0xc5, 0xba, 0x25, 0x8d, 0xb3,
	0xaa, 0x36, 0x8d, 0x99, 0x79, 0xf0, 0xf0, 0x48, 0x87, 0xf1, 0xcb, 0x8d, 0x91, 0x05, 0xe9, 0xda,
	0x48, 0x5b, 0xb7, 0xd2, 0xb7, 0x92, 0x4f, 0x06, 0xe1, 0x8b, 0x2b, 0xf9, 0x05, 0x04, 0x94, 0x0b,
	0x16, 0xb3, 0x74, 0x94, 0x05, 0x94, 0xf3, 0x09, 0x44, 0x5a, 0x55, 0x4a, 0x53, 0xf3, 0x21, 0x82,
	0x98, 0xa5, 0x83, 0xec, 0xf4, 0xe7, 0x53, 0x80, 0x42, 0xed, 0x57, 0x3b, 0xa4, 0xcd, 0x7b, 0x23,
	0x7a, 0xae, 0x3b, 0x2a, 0xd4, 0xfe, 0xd5, 0x01, 0x9c, 0x43, 0xbf, 0x54, 0x05, 0x8a, 0xbe, 0x13,
	0x73, 0x35, 0xbf, 0x81, 0x91, 0x6a, 0x15, 0x6d, 0xd5, 0x7a, 0x8b, 0x62, 0x10, 0xb3, 0x34, 0xca,
	0xbe, 0x01, 0x7e, 0x05, 0x91, 0xd9, 0x95, 0x58, 0xaf, 0x28, 0x17, 0xa1, 0x9b, 0x1a, 0xba, 0xff,
	0x63, 0x9e, 0x3c, 0xc1, 0x78, 0x59, 0xa1, 0xa6, 0x37, 0xd2, 0xaa, 0x21, 0x53, 0x9e, 0x19, 0x63,
	0xbf, 0x1a, 0x0b, 0x7e, 0x18, 0x4b, 0x5a, 0x88, 0x32, 0xb4, 0x95, 0x29, 0x2d, 0xf2, 0x05, 0x84,
	0x3e, 0x04, 0x27, 0xf2, 0x7f, 0x7e, 0x2d, 0x3b, 0x02, 0x92, 0x3e, 0x9c, 0xec, 0x40, 0xe5, 0x77,
	0x30, 0xf4, 0x95, 0x15, 0x41, 0xdc, 0xfb, 0x6b, 0xea, 0xc8, 0x9d, 0x23, 0x8c, 0x3d, 0xb4, 0xf4,
	0x67, 0xe2, 0xcf, 0x30, 0x7e, 0xa0, 0x32, 0xbf, 0x3f, 0x05, 0x90, 0x74, 0xea, 0x9c, 0x2d, 0x3e,
	0x99, 0x76, 0x72, 0x8e, 0x0b, 0x25, 0xff, 0xd6, 0xa1, 0xbb, 0xf4, 0xe2, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0xc2, 0xc0, 0xb1, 0x04, 0x15, 0x02, 0x00, 0x00,
}
