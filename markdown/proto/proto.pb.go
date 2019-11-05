// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/proto.proto

package proto

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

type RenderMarkdownRequest struct {
	Source               string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenderMarkdownRequest) Reset()         { *m = RenderMarkdownRequest{} }
func (m *RenderMarkdownRequest) String() string { return proto.CompactTextString(m) }
func (*RenderMarkdownRequest) ProtoMessage()    {}
func (*RenderMarkdownRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9897619deba694d0, []int{0}
}

func (m *RenderMarkdownRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenderMarkdownRequest.Unmarshal(m, b)
}
func (m *RenderMarkdownRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenderMarkdownRequest.Marshal(b, m, deterministic)
}
func (m *RenderMarkdownRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenderMarkdownRequest.Merge(m, src)
}
func (m *RenderMarkdownRequest) XXX_Size() int {
	return xxx_messageInfo_RenderMarkdownRequest.Size(m)
}
func (m *RenderMarkdownRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RenderMarkdownRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RenderMarkdownRequest proto.InternalMessageInfo

func (m *RenderMarkdownRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type RenderMarkdownResponse struct {
	Html                 string   `protobuf:"bytes,2,opt,name=html,proto3" json:"html,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenderMarkdownResponse) Reset()         { *m = RenderMarkdownResponse{} }
func (m *RenderMarkdownResponse) String() string { return proto.CompactTextString(m) }
func (*RenderMarkdownResponse) ProtoMessage()    {}
func (*RenderMarkdownResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9897619deba694d0, []int{1}
}

func (m *RenderMarkdownResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenderMarkdownResponse.Unmarshal(m, b)
}
func (m *RenderMarkdownResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenderMarkdownResponse.Marshal(b, m, deterministic)
}
func (m *RenderMarkdownResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenderMarkdownResponse.Merge(m, src)
}
func (m *RenderMarkdownResponse) XXX_Size() int {
	return xxx_messageInfo_RenderMarkdownResponse.Size(m)
}
func (m *RenderMarkdownResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RenderMarkdownResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RenderMarkdownResponse proto.InternalMessageInfo

func (m *RenderMarkdownResponse) GetHtml() string {
	if m != nil {
		return m.Html
	}
	return ""
}

func init() {
	proto.RegisterType((*RenderMarkdownRequest)(nil), "RenderMarkdownRequest")
	proto.RegisterType((*RenderMarkdownResponse)(nil), "RenderMarkdownResponse")
}

func init() { proto.RegisterFile("proto/proto.proto", fileDescriptor_9897619deba694d0) }

var fileDescriptor_9897619deba694d0 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x07, 0x93, 0x7a, 0x60, 0x52, 0x49, 0x9f, 0x4b, 0x34, 0x28, 0x35, 0x2f, 0x25, 0xb5,
	0xc8, 0x37, 0xb1, 0x28, 0x3b, 0x25, 0xbf, 0x3c, 0x2f, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44,
	0x48, 0x8c, 0x8b, 0xad, 0x38, 0xbf, 0xb4, 0x28, 0x39, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33,
	0x08, 0xca, 0x53, 0xd2, 0xe1, 0x12, 0x43, 0xd7, 0x50, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24,
	0xc4, 0xc5, 0x92, 0x51, 0x92, 0x9b, 0x23, 0xc1, 0x04, 0x56, 0x0f, 0x66, 0x1b, 0xf9, 0x73, 0x71,
	0xc0, 0xd4, 0x09, 0x39, 0x73, 0xf1, 0xa1, 0xea, 0x14, 0x12, 0xd3, 0xc3, 0x6a, 0xb7, 0x94, 0xb8,
	0x1e, 0x76, 0x2b, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0xce, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x4f, 0x31, 0x98, 0xd1, 0xcb, 0x00, 0x00, 0x00,
}
