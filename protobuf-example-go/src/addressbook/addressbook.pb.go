// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressbook/addressbook.proto

package tutorial

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7f7566f922065dad, []int{0, 0}
}

// [START messages]
type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated          *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f7566f922065dad, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=tutorial.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f7566f922065dad, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// Our address book file is just one of these.
type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f7566f922065dad, []int{1}
}

func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterEnum("tutorial.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Person)(nil), "tutorial.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "tutorial.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "tutorial.AddressBook")
}

func init() { proto.RegisterFile("addressbook/addressbook.proto", fileDescriptor_7f7566f922065dad) }

var fileDescriptor_7f7566f922065dad = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x51, 0x6b, 0xab, 0x30,
	0x18, 0x86, 0x8f, 0xd6, 0x4a, 0xfb, 0x79, 0x28, 0x9e, 0x50, 0x0e, 0x22, 0xe7, 0x30, 0x29, 0xbb,
	0x10, 0x06, 0x29, 0x74, 0x83, 0x5d, 0xed, 0x62, 0x42, 0xd9, 0xc6, 0xd6, 0x55, 0xa4, 0x65, 0x97,
	0x23, 0xce, 0xac, 0x93, 0xaa, 0x09, 0x26, 0xc2, 0xfa, 0x97, 0xf6, 0x17, 0xf6, 0xe7, 0x86, 0x46,
	0x8b, 0x8c, 0xdd, 0xbd, 0xc9, 0xf7, 0xf8, 0xf9, 0xe6, 0x81, 0xff, 0x24, 0x49, 0x4a, 0x2a, 0x44,
	0xcc, 0xd8, 0x7e, 0xde, 0xcb, 0x98, 0x97, 0x4c, 0x32, 0x34, 0x92, 0x95, 0x64, 0x65, 0x4a, 0x32,
	0xf7, 0x64, 0xc7, 0xd8, 0x2e, 0xa3, 0xf3, 0xe6, 0x3e, 0xae, 0x5e, 0xe7, 0x32, 0xcd, 0xa9, 0x90,
	0x24, 0xe7, 0x0a, 0x9d, 0x7d, 0xea, 0x60, 0x86, 0xb4, 0x14, 0xac, 0x40, 0x08, 0x8c, 0x82, 0xe4,
	0xd4, 0xd1, 0x3c, 0xcd, 0x1f, 0x47, 0x4d, 0x46, 0x13, 0xd0, 0xd3, 0xc4, 0xd1, 0x3d, 0xcd, 0x1f,
	0x46, 0x7a, 0x9a, 0xa0, 0x29, 0x0c, 0x69, 0x4e, 0xd2, 0xcc, 0x19, 0x34, 0x90, 0x3a, 0xa0, 0x0b,
	0x30, 0xf9, 0x1b, 0x2b, 0xa8, 0x70, 0x0c, 0x6f, 0xe0, 0x5b, 0x8b, 0x7f, 0xb8, 0x2b, 0x80, 0xd5,
	0x6e, 0x1c, 0xd6, 0xe3, 0xc7, 0x2a, 0x8f, 0x69, 0x19, 0xb5, 0x2c, 0xba, 0x82, 0xdf, 0x19, 0x11,
	0xf2, 0xb9, 0xe2, 0x09, 0x91, 0x34, 0x71, 0x86, 0x9e, 0xe6, 0x5b, 0x0b, 0x17, 0xab, 0xca, 0xb8,
	0xab, 0x8c, 0x37, 0x5d, 0xe5, 0xc8, 0xaa, 0xf9, 0xad, 0xc2, 0xdd, 0x2d, 0x58, 0xbd, 0xad, 0xe8,
	0x2f, 0x98, 0x45, 0x93, 0xda, 0xfe, 0xed, 0x09, 0x61, 0x30, 0xe4, 0x81, 0xd3, 0xe6, 0x0d, 0x93,
	0x85, 0xfb, 0x73, 0xb3, 0xcd, 0x81, 0xd3, 0xa8, 0xe1, 0x66, 0x67, 0x30, 0x3e, 0x5e, 0x21, 0x00,
	0x73, 0xb5, 0x0e, 0xee, 0x1e, 0x96, 0xf6, 0x2f, 0x34, 0x02, 0xe3, 0x76, 0xbd, 0x5a, 0xda, 0x5a,
	0x9d, 0x9e, 0xd6, 0xd1, 0xbd, 0xad, 0xcf, 0x2e, 0xc1, 0xba, 0x56, 0xf6, 0x03, 0xc6, 0xf6, 0xc8,
	0x07, 0x93, 0x53, 0xc6, 0xb3, 0xda, 0x61, 0xed, 0xc1, 0xfe, 0xfe, 0xb7, 0xa8, 0x9d, 0x07, 0x21,
	0x4c, 0x5f, 0x58, 0x8e, 0xe9, 0x3b, 0xc9, 0x79, 0x46, 0x8f, 0x58, 0xf0, 0xa7, 0xb7, 0x2e, 0xac,
	0x05, 0x88, 0x0f, 0xfd, 0xf4, 0x46, 0x09, 0x09, 0x3b, 0x21, 0x4b, 0xf5, 0x95, 0xc0, 0x3d, 0x38,
	0x36, 0x1b, 0x5f, 0xe7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x5a, 0x35, 0x89, 0x1b, 0x02,
	0x00, 0x00,
}