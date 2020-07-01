// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage_message.proto

package pb

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

type Storage_Driver int32

const (
	Storage_UNKNOWN Storage_Driver = 0
	Storage_HDD     Storage_Driver = 1
	Storage_SSD     Storage_Driver = 2
)

var Storage_Driver_name = map[int32]string{
	0: "UNKNOWN",
	1: "HDD",
	2: "SSD",
}

var Storage_Driver_value = map[string]int32{
	"UNKNOWN": 0,
	"HDD":     1,
	"SSD":     2,
}

func (x Storage_Driver) String() string {
	return proto.EnumName(Storage_Driver_name, int32(x))
}

func (Storage_Driver) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_170f09d838bd8a04, []int{0, 0}
}

type Storage struct {
	Driver               Storage_Driver `protobuf:"varint,1,opt,name=driver,proto3,enum=techschool.pcbook.Storage_Driver" json:"driver,omitempty"`
	Memory               *Memory        `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Storage) Reset()         { *m = Storage{} }
func (m *Storage) String() string { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()    {}
func (*Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_170f09d838bd8a04, []int{0}
}

func (m *Storage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storage.Unmarshal(m, b)
}
func (m *Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storage.Marshal(b, m, deterministic)
}
func (m *Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage.Merge(m, src)
}
func (m *Storage) XXX_Size() int {
	return xxx_messageInfo_Storage.Size(m)
}
func (m *Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_Storage proto.InternalMessageInfo

func (m *Storage) GetDriver() Storage_Driver {
	if m != nil {
		return m.Driver
	}
	return Storage_UNKNOWN
}

func (m *Storage) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func init() {
	proto.RegisterEnum("techschool.pcbook.Storage_Driver", Storage_Driver_name, Storage_Driver_value)
	proto.RegisterType((*Storage)(nil), "techschool.pcbook.Storage")
}

func init() {
	proto.RegisterFile("storage_message.proto", fileDescriptor_170f09d838bd8a04)
}

var fileDescriptor_170f09d838bd8a04 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0x8d, 0xcf, 0x4d, 0x2d, 0x2e, 0x06, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0x82, 0x25, 0xa9, 0xc9, 0x19, 0xc5, 0xc9, 0x19, 0xf9, 0xf9, 0x39, 0x7a, 0x05, 0xc9, 0x49,
	0xf9, 0xf9, 0xd9, 0x52, 0x22, 0xb9, 0xa9, 0xb9, 0xf9, 0x45, 0x95, 0xa8, 0x0a, 0x95, 0x16, 0x30,
	0x72, 0xb1, 0x07, 0x43, 0x8c, 0x10, 0xb2, 0xe4, 0x62, 0x4b, 0x29, 0xca, 0x2c, 0x4b, 0x2d, 0x92,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0x52, 0xd4, 0xc3, 0x30, 0x45, 0x0f, 0xaa, 0x56, 0xcf, 0x05,
	0xac, 0x30, 0x08, 0xaa, 0x41, 0xc8, 0x90, 0x8b, 0x0d, 0x62, 0xbc, 0x04, 0x13, 0x50, 0x2b, 0xb7,
	0x91, 0x24, 0x16, 0xad, 0xbe, 0x60, 0x05, 0x41, 0x50, 0x85, 0x4a, 0xea, 0x5c, 0x6c, 0x10, 0x43,
	0x84, 0xb8, 0xb9, 0xd8, 0x43, 0xfd, 0xbc, 0xfd, 0xfc, 0xc3, 0xfd, 0x04, 0x18, 0x84, 0xd8, 0xb9,
	0x98, 0x3d, 0x5c, 0x5c, 0x04, 0x18, 0x41, 0x8c, 0xe0, 0x60, 0x17, 0x01, 0x26, 0x27, 0x4d, 0x2e,
	0xf9, 0xe4, 0xfc, 0x5c, 0xbd, 0xf4, 0xcc, 0x92, 0x9c, 0xc4, 0x24, 0x2c, 0xe6, 0x16, 0x24, 0x05,
	0x30, 0x46, 0xb1, 0xe8, 0x59, 0x17, 0x24, 0x25, 0xb1, 0x81, 0x3d, 0x65, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x8d, 0x31, 0x5a, 0xbc, 0x16, 0x01, 0x00, 0x00,
}