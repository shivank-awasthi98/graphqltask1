// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sa/main.proto

package Sa

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
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

type User struct {
	Id                   uint64               `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Username             string               `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Email                string               `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string               `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d0d27f403456fd9, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "Sa.User")
}

func init() { proto.RegisterFile("sa/main.proto", fileDescriptor_0d0d27f403456fd9) }

var fileDescriptor_0d0d27f403456fd9 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8e, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0xe5, 0x90, 0x46, 0xad, 0x11, 0x0c, 0x11, 0x83, 0xc9, 0x42, 0xc4, 0x94, 0xa5, 0xb6,
	0x04, 0x13, 0xdd, 0x3a, 0x30, 0x74, 0x43, 0x01, 0x66, 0xe4, 0xc4, 0xae, 0xb1, 0x14, 0xfb, 0x59,
	0xb6, 0x23, 0xf8, 0x35, 0xf8, 0x1c, 0xbe, 0x04, 0x25, 0x6e, 0x58, 0x19, 0xcf, 0x7b, 0xe7, 0x48,
	0x17, 0x5f, 0x04, 0xce, 0x0c, 0xd7, 0x96, 0x3a, 0x0f, 0x11, 0xca, 0xec, 0x99, 0x57, 0x3b, 0xa5,
	0xe3, 0xfb, 0xd8, 0xd1, 0x1e, 0x0c, 0xd3, 0xf6, 0x08, 0xdd, 0x00, 0x9f, 0xe0, 0xa4, 0x65, 0xb3,
	0xd2, 0x6f, 0x95, 0xb4, 0x5b, 0x05, 0xde, 0x30, 0x70, 0x51, 0x83, 0x0d, 0x6c, 0x82, 0xd4, 0x57,
	0x37, 0x0a, 0x40, 0x0d, 0x32, 0xa9, 0xdd, 0x78, 0x64, 0x51, 0x1b, 0x19, 0x22, 0x37, 0x2e, 0x09,
	0xb7, 0x3f, 0x08, 0xe7, 0xaf, 0x41, 0xfa, 0xf2, 0x12, 0x67, 0x07, 0x41, 0x50, 0x8d, 0x9a, 0xbc,
	0xcd, 0x0e, 0xa2, 0xac, 0xf0, 0x7a, 0xba, 0x5b, 0x6e, 0x24, 0xc9, 0x6a, 0xd4, 0x6c, 0xda, 0x3f,
	0x2e, 0xaf, 0xf0, 0xea, 0xd1, 0x70, 0x3d, 0x90, 0xb3, 0xf9, 0x91, 0x60, 0x2a, 0x9e, 0x78, 0x08,
	0x1f, 0xe0, 0x05, 0xc9, 0x53, 0xb1, 0x70, 0xf9, 0x80, 0x71, 0xef, 0x25, 0x8f, 0x52, 0xbc, 0xf1,
	0x48, 0x56, 0x35, 0x6a, 0xce, 0xef, 0x2a, 0x9a, 0xc6, 0xd1, 0x65, 0x1c, 0x7d, 0x59, 0xc6, 0xb5,
	0x9b, 0x93, 0xbd, 0x8f, 0x53, 0x3a, 0x3a, 0xb1, 0xa4, 0xc5, 0xff, 0xe9, 0xc9, 0xde, 0xc7, 0x5d,
	0xf1, 0xfd, 0x75, 0x9d, 0xad, 0x51, 0x57, 0xcc, 0xda, 0xfd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x67, 0x16, 0xab, 0x52, 0x5d, 0x01, 0x00, 0x00,
}