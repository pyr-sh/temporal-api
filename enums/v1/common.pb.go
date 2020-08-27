// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/api/enums/v1/common.proto

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

type EncodingType int32

const (
	ENCODING_TYPE_UNSPECIFIED EncodingType = 0
	ENCODING_TYPE_PROTO3      EncodingType = 1
	ENCODING_TYPE_JSON        EncodingType = 2
)

var EncodingType_name = map[int32]string{
	0: "Unspecified",
	1: "Proto3",
	2: "Json",
}

var EncodingType_value = map[string]int32{
	"Unspecified": 0,
	"Proto3":      1,
	"Json":        2,
}

func (EncodingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_768283dde248a0f8, []int{0}
}

type IndexedValueType int32

const (
	INDEXED_VALUE_TYPE_UNSPECIFIED IndexedValueType = 0
	INDEXED_VALUE_TYPE_STRING      IndexedValueType = 1
	INDEXED_VALUE_TYPE_KEYWORD     IndexedValueType = 2
	INDEXED_VALUE_TYPE_INT         IndexedValueType = 3
	INDEXED_VALUE_TYPE_DOUBLE      IndexedValueType = 4
	INDEXED_VALUE_TYPE_BOOL        IndexedValueType = 5
	INDEXED_VALUE_TYPE_DATETIME    IndexedValueType = 6
)

var IndexedValueType_name = map[int32]string{
	0: "Unspecified",
	1: "String",
	2: "Keyword",
	3: "Int",
	4: "Double",
	5: "Bool",
	6: "Datetime",
}

var IndexedValueType_value = map[string]int32{
	"Unspecified": 0,
	"String":      1,
	"Keyword":     2,
	"Int":         3,
	"Double":      4,
	"Bool":        5,
	"Datetime":    6,
}

func (IndexedValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_768283dde248a0f8, []int{1}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.EncodingType", EncodingType_name, EncodingType_value)
	proto.RegisterEnum("temporal.api.enums.v1.IndexedValueType", IndexedValueType_name, IndexedValueType_value)
}

func init() {
	proto.RegisterFile("temporal/api/enums/v1/common.proto", fileDescriptor_768283dde248a0f8)
}

var fileDescriptor_768283dde248a0f8 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xbf, 0x6e, 0xe2, 0x30,
	0x1c, 0x80, 0x63, 0xee, 0x8e, 0xc1, 0x77, 0x83, 0x65, 0xdd, 0x71, 0xfc, 0x11, 0xae, 0xca, 0xc8,
	0x90, 0x08, 0xb1, 0xa5, 0x13, 0x21, 0x2e, 0x72, 0x4b, 0xe3, 0x08, 0x0c, 0x2d, 0x5d, 0xa2, 0x14,
	0x22, 0x14, 0x89, 0xc4, 0x11, 0x05, 0xd4, 0x6e, 0x7d, 0x84, 0xee, 0x7d, 0x81, 0x3e, 0x4a, 0x47,
	0x46, 0xc6, 0x12, 0x96, 0xaa, 0x43, 0xc5, 0x23, 0x54, 0x44, 0xad, 0x54, 0xa4, 0x74, 0xb3, 0xfc,
	0x7d, 0xf6, 0x6f, 0xf8, 0x7e, 0xb0, 0x32, 0xf3, 0x82, 0x48, 0x4e, 0xdd, 0x89, 0xe6, 0x46, 0xbe,
	0xe6, 0x85, 0xf3, 0xe0, 0x5a, 0x5b, 0xd4, 0xb4, 0xa1, 0x0c, 0x02, 0x19, 0xaa, 0xd1, 0x54, 0xce,
	0x24, 0xfe, 0xf7, 0xe9, 0xa8, 0x6e, 0xe4, 0xab, 0x89, 0xa3, 0x2e, 0x6a, 0x55, 0x07, 0xfe, 0xa1,
	0xe1, 0x50, 0x8e, 0xfc, 0x70, 0x2c, 0x6e, 0x23, 0x0f, 0x97, 0x61, 0x81, 0x5a, 0x4d, 0x6e, 0x32,
	0xab, 0xe5, 0x88, 0x81, 0x4d, 0x9d, 0x9e, 0xd5, 0xb5, 0x69, 0x93, 0x1d, 0x33, 0x6a, 0x22, 0x05,
	0xe7, 0xe1, 0xdf, 0x7d, 0x6c, 0x77, 0xb8, 0xe0, 0x75, 0x04, 0x70, 0x0e, 0xe2, 0x7d, 0x72, 0xd2,
	0xe5, 0x16, 0xca, 0x54, 0xdf, 0x00, 0x44, 0x2c, 0x1c, 0x79, 0x37, 0xde, 0xa8, 0xef, 0x4e, 0xe6,
	0x5e, 0x32, 0xa5, 0x02, 0x09, 0xb3, 0x4c, 0x7a, 0x41, 0x4d, 0xa7, 0xdf, 0x68, 0xf7, 0x68, 0xda,
	0xa8, 0x32, 0x2c, 0xa4, 0x38, 0x5d, 0xd1, 0x61, 0x56, 0x0b, 0x01, 0x4c, 0x60, 0x31, 0x05, 0x9f,
	0xd2, 0xc1, 0x39, 0xef, 0x98, 0x28, 0x83, 0x8b, 0x30, 0x97, 0xc2, 0x99, 0x25, 0xd0, 0x8f, 0x6f,
	0xbe, 0x36, 0x79, 0xcf, 0x68, 0x53, 0xf4, 0x13, 0x97, 0xe0, 0xff, 0x14, 0x6c, 0x70, 0xde, 0x46,
	0xbf, 0xf0, 0x01, 0x2c, 0xa5, 0xbd, 0x6d, 0x08, 0x2a, 0xd8, 0x19, 0x45, 0x59, 0xe3, 0x01, 0x2c,
	0xd7, 0x44, 0x59, 0xad, 0x89, 0xb2, 0x5d, 0x13, 0x70, 0x17, 0x13, 0xf0, 0x18, 0x13, 0xf0, 0x14,
	0x13, 0xb0, 0x8c, 0x09, 0x78, 0x8e, 0x09, 0x78, 0x89, 0x89, 0xb2, 0x8d, 0x09, 0xb8, 0xdf, 0x10,
	0x65, 0xb9, 0x21, 0xca, 0x6a, 0x43, 0x14, 0x98, 0xf7, 0xa5, 0x9a, 0x9a, 0xc8, 0xf8, 0xdd, 0x4c,
	0x3a, 0xda, 0xbb, 0x8c, 0x36, 0xb8, 0x3c, 0x1c, 0x7f, 0x11, 0x7d, 0xb9, 0x97, 0xfc, 0x28, 0x39,
	0xbc, 0x66, 0xf2, 0xe2, 0x43, 0xd0, 0xf5, 0x46, 0xe4, 0xeb, 0x3a, 0xdd, 0x5d, 0xeb, 0x7a, 0xbf,
	0x76, 0x95, 0x4d, 0xb6, 0xa1, 0xfe, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x94, 0x9b, 0x54, 0x33,
	0x02, 0x00, 0x00,
}

func (x EncodingType) String() string {
	s, ok := EncodingType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x IndexedValueType) String() string {
	s, ok := IndexedValueType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}