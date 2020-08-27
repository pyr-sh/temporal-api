// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/api/enums/v1/namespace.proto

package enums

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
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

type NamespaceState int32

const (
	NAMESPACE_STATE_UNSPECIFIED NamespaceState = 0
	NAMESPACE_STATE_REGISTERED  NamespaceState = 1
	NAMESPACE_STATE_DEPRECATED  NamespaceState = 2
	NAMESPACE_STATE_DELETED     NamespaceState = 3
)

var NamespaceState_name = map[int32]string{
	0: "Unspecified",
	1: "Registered",
	2: "Deprecated",
	3: "Deleted",
}

var NamespaceState_value = map[string]int32{
	"Unspecified": 0,
	"Registered":  1,
	"Deprecated":  2,
	"Deleted":     3,
}

func (NamespaceState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d8c23ee067d676e7, []int{0}
}

type ArchivalState int32

const (
	ARCHIVAL_STATE_UNSPECIFIED ArchivalState = 0
	ARCHIVAL_STATE_DISABLED    ArchivalState = 1
	ARCHIVAL_STATE_ENABLED     ArchivalState = 2
)

var ArchivalState_name = map[int32]string{
	0: "Unspecified",
	1: "Disabled",
	2: "Enabled",
}

var ArchivalState_value = map[string]int32{
	"Unspecified": 0,
	"Disabled":    1,
	"Enabled":     2,
}

func (ArchivalState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d8c23ee067d676e7, []int{1}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.NamespaceState", NamespaceState_name, NamespaceState_value)
	proto.RegisterEnum("temporal.api.enums.v1.ArchivalState", ArchivalState_name, ArchivalState_value)
}

func init() {
	proto.RegisterFile("temporal/api/enums/v1/namespace.proto", fileDescriptor_d8c23ee067d676e7)
}

var fileDescriptor_d8c23ee067d676e7 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0xef, 0xfa, 0x87, 0xff, 0x70, 0xa0, 0x94, 0x80, 0xb6, 0xb4, 0xf0, 0x8a, 0x83, 0x4b,
	0x87, 0x84, 0xe2, 0x16, 0xa7, 0x6b, 0xf2, 0xaa, 0x81, 0x1a, 0x42, 0x12, 0x3b, 0xb8, 0x94, 0xb3,
	0x04, 0x7b, 0xd0, 0x36, 0x47, 0x1b, 0x3b, 0xfb, 0x09, 0xc4, 0x4f, 0xe0, 0xec, 0x47, 0x71, 0xec,
	0xd8, 0xd1, 0x5e, 0x17, 0x71, 0xea, 0x47, 0x90, 0xa6, 0x16, 0xb5, 0x66, 0x7b, 0x79, 0x7f, 0xbf,
	0xe1, 0xe1, 0x79, 0xd8, 0x49, 0x96, 0x0c, 0x55, 0x3a, 0x16, 0x03, 0x4b, 0x28, 0x69, 0x25, 0xa3,
	0xfb, 0xe1, 0xc4, 0x9a, 0x36, 0xad, 0x91, 0x18, 0x26, 0x13, 0x25, 0x7a, 0x89, 0xa9, 0xc6, 0x69,
	0x96, 0x1a, 0x07, 0x5b, 0xcd, 0x14, 0x4a, 0x9a, 0xb9, 0x66, 0x4e, 0x9b, 0x8d, 0x47, 0xca, 0xf6,
	0xfd, 0xad, 0x1a, 0x65, 0x22, 0x4b, 0x8c, 0x23, 0x56, 0xf7, 0xf9, 0x15, 0x46, 0x01, 0x77, 0xb0,
	0x1b, 0xc5, 0x3c, 0xc6, 0xee, 0xb5, 0x1f, 0x05, 0xe8, 0x78, 0xe7, 0x1e, 0xba, 0x65, 0x62, 0x00,
	0xab, 0xed, 0x0a, 0x21, 0x5e, 0x78, 0x51, 0x8c, 0x21, 0xba, 0x65, 0x5a, 0xc4, 0x5d, 0x0c, 0x42,
	0x74, 0x78, 0x8c, 0x6e, 0xb9, 0x64, 0xd4, 0x59, 0xe5, 0x2f, 0x6f, 0xe3, 0x1a, 0xfe, 0x6b, 0xf4,
	0xd9, 0x1e, 0x1f, 0xf7, 0xfa, 0x72, 0x2a, 0x06, 0x9b, 0x38, 0xc0, 0x6a, 0x3c, 0x74, 0x2e, 0xbd,
	0x0e, 0x6f, 0x17, 0xa6, 0xa9, 0xb3, 0xca, 0x0e, 0x77, 0xbd, 0x88, 0xb7, 0xda, 0x79, 0x94, 0x1a,
	0x3b, 0xdc, 0x81, 0xe8, 0x6f, 0x58, 0xa9, 0xf5, 0x4c, 0x67, 0x0b, 0x20, 0xf3, 0x05, 0x90, 0xd5,
	0x02, 0xe8, 0x83, 0x06, 0xfa, 0xa2, 0x81, 0xbe, 0x6a, 0xa0, 0x33, 0x0d, 0xf4, 0x4d, 0x03, 0x7d,
	0xd7, 0x40, 0x56, 0x1a, 0xe8, 0xd3, 0x12, 0xc8, 0x6c, 0x09, 0x64, 0xbe, 0x04, 0xc2, 0xaa, 0x32,
	0x35, 0x0b, 0xbb, 0x6c, 0x7d, 0x17, 0x19, 0xac, 0x2b, 0x0f, 0xe8, 0xcd, 0xf1, 0xdd, 0x0f, 0x57,
	0xa6, 0xbf, 0x16, 0x3a, 0xcb, 0x8f, 0x8f, 0x52, 0x35, 0xfe, 0x12, 0x6c, 0x9b, 0x2b, 0x69, 0xdb,
	0xb8, 0x7e, 0xdb, 0x76, 0xa7, 0x79, 0xfb, 0x3f, 0x5f, 0xee, 0xf4, 0x33, 0x00, 0x00, 0xff, 0xff,
	0x93, 0xd7, 0xb4, 0x94, 0xe2, 0x01, 0x00, 0x00,
}

func (x NamespaceState) String() string {
	s, ok := NamespaceState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ArchivalState) String() string {
	s, ok := ArchivalState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
