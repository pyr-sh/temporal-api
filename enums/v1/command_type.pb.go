// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/api/enums/v1/command_type.proto

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

// Whenever this list of command types is changed do change the function shouldBufferEvent in mutableStateBuilder.go to make sure to do the correct event ordering.
type CommandType int32

const (
	COMMAND_TYPE_UNSPECIFIED                                CommandType = 0
	COMMAND_TYPE_SCHEDULE_ACTIVITY_TASK                     CommandType = 1
	COMMAND_TYPE_REQUEST_CANCEL_ACTIVITY_TASK               CommandType = 2
	COMMAND_TYPE_START_TIMER                                CommandType = 3
	COMMAND_TYPE_COMPLETE_WORKFLOW_EXECUTION                CommandType = 4
	COMMAND_TYPE_FAIL_WORKFLOW_EXECUTION                    CommandType = 5
	COMMAND_TYPE_CANCEL_TIMER                               CommandType = 6
	COMMAND_TYPE_CANCEL_WORKFLOW_EXECUTION                  CommandType = 7
	COMMAND_TYPE_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION CommandType = 8
	COMMAND_TYPE_RECORD_MARKER                              CommandType = 9
	COMMAND_TYPE_CONTINUE_AS_NEW_WORKFLOW_EXECUTION         CommandType = 10
	COMMAND_TYPE_START_CHILD_WORKFLOW_EXECUTION             CommandType = 11
	COMMAND_TYPE_SIGNAL_EXTERNAL_WORKFLOW_EXECUTION         CommandType = 12
	COMMAND_TYPE_UPSERT_WORKFLOW_SEARCH_ATTRIBUTES          CommandType = 13
)

var CommandType_name = map[int32]string{
	0:  "Unspecified",
	1:  "ScheduleActivityTask",
	2:  "RequestCancelActivityTask",
	3:  "StartTimer",
	4:  "CompleteWorkflowExecution",
	5:  "FailWorkflowExecution",
	6:  "CancelTimer",
	7:  "CancelWorkflowExecution",
	8:  "RequestCancelExternalWorkflowExecution",
	9:  "RecordMarker",
	10: "ContinueAsNewWorkflowExecution",
	11: "StartChildWorkflowExecution",
	12: "SignalExternalWorkflowExecution",
	13: "UpsertWorkflowSearchAttributes",
}

var CommandType_value = map[string]int32{
	"Unspecified":                            0,
	"ScheduleActivityTask":                   1,
	"RequestCancelActivityTask":              2,
	"StartTimer":                             3,
	"CompleteWorkflowExecution":              4,
	"FailWorkflowExecution":                  5,
	"CancelTimer":                            6,
	"CancelWorkflowExecution":                7,
	"RequestCancelExternalWorkflowExecution": 8,
	"RecordMarker":                           9,
	"ContinueAsNewWorkflowExecution":         10,
	"StartChildWorkflowExecution":            11,
	"SignalExternalWorkflowExecution":        12,
	"UpsertWorkflowSearchAttributes":         13,
}

func (CommandType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e70af5ad8bd10de0, []int{0}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.CommandType", CommandType_name, CommandType_value)
}

func init() {
	proto.RegisterFile("temporal/api/enums/v1/command_type.proto", fileDescriptor_e70af5ad8bd10de0)
}

var fileDescriptor_e70af5ad8bd10de0 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0x3d, 0xfd, 0xfa, 0x15, 0x98, 0x82, 0x64, 0x8d, 0x84, 0x14, 0x10, 0x8c, 0x84, 0x40,
	0x10, 0x0a, 0xd8, 0x0a, 0x5d, 0x20, 0xb9, 0xab, 0xc9, 0xf8, 0x86, 0x8c, 0xe2, 0x7f, 0xcc, 0x8c,
	0x9b, 0x96, 0xcd, 0x28, 0x40, 0x84, 0x2c, 0x91, 0xda, 0x2a, 0xa1, 0x52, 0x77, 0x3c, 0x02, 0x8f,
	0xc0, 0x92, 0x47, 0x61, 0x99, 0x65, 0x97, 0xc4, 0xd9, 0x20, 0x56, 0x79, 0x04, 0x94, 0x04, 0x44,
	0x9d, 0x5a, 0xdd, 0x59, 0xbe, 0xbf, 0x7b, 0xcf, 0xb9, 0xa3, 0x73, 0x71, 0x73, 0x3c, 0x1c, 0x15,
	0xf9, 0xf1, 0xe0, 0x83, 0x3b, 0x28, 0x32, 0x77, 0x78, 0xf4, 0x69, 0xf4, 0xd1, 0x3d, 0x69, 0xb9,
	0x6f, 0xf3, 0xd1, 0x68, 0x70, 0xf4, 0xce, 0x8c, 0x4f, 0x8b, 0xa1, 0x53, 0x1c, 0xe7, 0xe3, 0x9c,
	0xdc, 0xfc, 0x4b, 0x3a, 0x83, 0x22, 0x73, 0x96, 0xa4, 0x73, 0xd2, 0xda, 0x99, 0x6f, 0xe2, 0x6d,
	0xbe, 0xa2, 0xf5, 0x69, 0x31, 0x24, 0x77, 0x70, 0x83, 0xc7, 0x61, 0xc8, 0x22, 0xdf, 0xe8, 0xc3,
	0x04, 0x4c, 0x1a, 0xa9, 0x04, 0xb8, 0xe8, 0x08, 0xf0, 0x6d, 0x8b, 0x3c, 0xc2, 0xf7, 0x2b, 0x55,
	0xc5, 0xbb, 0xe0, 0xa7, 0x01, 0x18, 0xc6, 0xb5, 0xd8, 0x17, 0xfa, 0xd0, 0x68, 0xa6, 0x7a, 0x36,
	0x22, 0xcf, 0xf0, 0xe3, 0x0a, 0x28, 0xe1, 0x55, 0x0a, 0x4a, 0x1b, 0xce, 0x22, 0x0e, 0xc1, 0x1a,
	0xbe, 0x71, 0x41, 0x55, 0x69, 0x26, 0xb5, 0xd1, 0x22, 0x04, 0x69, 0xff, 0x47, 0x9e, 0xe2, 0x66,
	0xa5, 0xca, 0xe3, 0x30, 0x09, 0x40, 0x83, 0xe9, 0xc7, 0xb2, 0xd7, 0x09, 0xe2, 0xbe, 0x81, 0x03,
	0xe0, 0xa9, 0x16, 0x71, 0x64, 0x6f, 0x92, 0x26, 0x7e, 0x50, 0xa1, 0x3b, 0x4c, 0x04, 0x75, 0xe4,
	0xff, 0xe4, 0x2e, 0xbe, 0x55, 0x9d, 0xbb, 0x32, 0xb7, 0x92, 0xdd, 0x22, 0x3b, 0xf8, 0x61, 0x5d,
	0xb9, 0x66, 0xd4, 0x15, 0xb2, 0x87, 0x5f, 0x5c, 0xb6, 0x2f, 0x1c, 0x68, 0x90, 0x11, 0xab, 0x6d,
	0xbe, 0x4a, 0x28, 0xbe, 0xbd, 0xd6, 0xcc, 0x63, 0xe9, 0x9b, 0x90, 0xc9, 0x1e, 0x48, 0xfb, 0x1a,
	0xd9, 0xc5, 0xee, 0xda, 0xfe, 0x91, 0x16, 0x51, 0x0a, 0x86, 0x29, 0x13, 0x41, 0xbf, 0x6e, 0x28,
	0x26, 0x2e, 0x7e, 0x52, 0xf3, 0xa4, 0xbc, 0x2b, 0x02, 0xbf, 0xae, 0x61, 0xfb, 0x82, 0x8a, 0x12,
	0x2f, 0x17, 0x8e, 0x2f, 0xb3, 0x7e, 0x9d, 0x3c, 0xc7, 0x4e, 0x35, 0x2e, 0x89, 0x02, 0xa9, 0xff,
	0xb1, 0x0a, 0x98, 0xe4, 0x5d, 0xc3, 0xb4, 0x96, 0xa2, 0x9d, 0x6a, 0x50, 0xf6, 0x8d, 0xf6, 0x57,
	0x34, 0x99, 0x52, 0xeb, 0x6c, 0x4a, 0xad, 0xf9, 0x94, 0xa2, 0xcf, 0x25, 0x45, 0xdf, 0x4a, 0x8a,
	0xbe, 0x97, 0x14, 0x4d, 0x4a, 0x8a, 0x7e, 0x94, 0x14, 0xfd, 0x2c, 0xa9, 0x35, 0x2f, 0x29, 0xfa,
	0x32, 0xa3, 0xd6, 0x64, 0x46, 0xad, 0xb3, 0x19, 0xb5, 0x70, 0x23, 0xcb, 0x9d, 0xda, 0x0c, 0xb7,
	0xed, 0x73, 0x01, 0x4e, 0x16, 0x61, 0x4f, 0xd0, 0xeb, 0x7b, 0xef, 0xcf, 0xd1, 0x59, 0x5e, 0x39,
	0x8f, 0xbd, 0xe5, 0xc7, 0xaf, 0x8d, 0x86, 0xfe, 0x03, 0x78, 0x1e, 0x2b, 0x32, 0xcf, 0x83, 0xc5,
	0x6f, 0xcf, 0xdb, 0x6f, 0xbd, 0xd9, 0x5a, 0xde, 0xcc, 0xee, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x4c, 0xba, 0x30, 0x76, 0x5f, 0x03, 0x00, 0x00,
}

func (x CommandType) String() string {
	s, ok := CommandType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}