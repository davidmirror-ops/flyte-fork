// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/metrics.proto

package core

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

// Span represents a duration trace of Flyte execution. The id field denotes a Flyte execution entity or an operation
// which uniquely identifies the Span. The spans attribute allows this Span to be further broken down into more
// precise definitions.
type Span struct {
	// start_time defines the instance this span began.
	StartTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// end_time defines the instance this span completed.
	EndTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Types that are valid to be assigned to Id:
	//	*Span_WorkflowId
	//	*Span_NodeId
	//	*Span_TaskId
	//	*Span_OperationId
	Id isSpan_Id `protobuf_oneof:"id"`
	// spans defines a collection of Spans that breakdown this execution.
	Spans                []*Span  `protobuf:"bytes,7,rep,name=spans,proto3" json:"spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Span) Reset()         { *m = Span{} }
func (m *Span) String() string { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()    {}
func (*Span) Descriptor() ([]byte, []int) {
	return fileDescriptor_756935f796ae3119, []int{0}
}

func (m *Span) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span.Unmarshal(m, b)
}
func (m *Span) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span.Marshal(b, m, deterministic)
}
func (m *Span) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span.Merge(m, src)
}
func (m *Span) XXX_Size() int {
	return xxx_messageInfo_Span.Size(m)
}
func (m *Span) XXX_DiscardUnknown() {
	xxx_messageInfo_Span.DiscardUnknown(m)
}

var xxx_messageInfo_Span proto.InternalMessageInfo

func (m *Span) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Span) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type isSpan_Id interface {
	isSpan_Id()
}

type Span_WorkflowId struct {
	WorkflowId *WorkflowExecutionIdentifier `protobuf:"bytes,3,opt,name=workflow_id,json=workflowId,proto3,oneof"`
}

type Span_NodeId struct {
	NodeId *NodeExecutionIdentifier `protobuf:"bytes,4,opt,name=node_id,json=nodeId,proto3,oneof"`
}

type Span_TaskId struct {
	TaskId *TaskExecutionIdentifier `protobuf:"bytes,5,opt,name=task_id,json=taskId,proto3,oneof"`
}

type Span_OperationId struct {
	OperationId string `protobuf:"bytes,6,opt,name=operation_id,json=operationId,proto3,oneof"`
}

func (*Span_WorkflowId) isSpan_Id() {}

func (*Span_NodeId) isSpan_Id() {}

func (*Span_TaskId) isSpan_Id() {}

func (*Span_OperationId) isSpan_Id() {}

func (m *Span) GetId() isSpan_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Span) GetWorkflowId() *WorkflowExecutionIdentifier {
	if x, ok := m.GetId().(*Span_WorkflowId); ok {
		return x.WorkflowId
	}
	return nil
}

func (m *Span) GetNodeId() *NodeExecutionIdentifier {
	if x, ok := m.GetId().(*Span_NodeId); ok {
		return x.NodeId
	}
	return nil
}

func (m *Span) GetTaskId() *TaskExecutionIdentifier {
	if x, ok := m.GetId().(*Span_TaskId); ok {
		return x.TaskId
	}
	return nil
}

func (m *Span) GetOperationId() string {
	if x, ok := m.GetId().(*Span_OperationId); ok {
		return x.OperationId
	}
	return ""
}

func (m *Span) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Span) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Span_WorkflowId)(nil),
		(*Span_NodeId)(nil),
		(*Span_TaskId)(nil),
		(*Span_OperationId)(nil),
	}
}

func init() {
	proto.RegisterType((*Span)(nil), "flyteidl.core.Span")
}

func init() { proto.RegisterFile("flyteidl/core/metrics.proto", fileDescriptor_756935f796ae3119) }

var fileDescriptor_756935f796ae3119 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0x40, 0xfb, 0xdd, 0xd7, 0xc9, 0x7b, 0x9b, 0xbc, 0x4d, 0xe9, 0x83, 0x67, 0x51, 0x90, 0x2a,
	0x38, 0x03, 0xf5, 0x03, 0x5c, 0x5a, 0x10, 0x9a, 0x85, 0x2e, 0x62, 0x41, 0x70, 0x53, 0x92, 0xcc,
	0x4d, 0x1c, 0x9a, 0xcc, 0x0d, 0x33, 0x53, 0xaa, 0xbf, 0xc7, 0x3f, 0x2a, 0x33, 0x31, 0x85, 0x96,
	0x8a, 0xcb, 0xc9, 0x9c, 0x73, 0xc2, 0xcd, 0x0d, 0xf9, 0x97, 0xe6, 0xef, 0x06, 0x04, 0xcf, 0x59,
	0x82, 0x0a, 0x58, 0x01, 0x46, 0x89, 0x44, 0xd3, 0x52, 0xa1, 0x41, 0xff, 0x4f, 0x7d, 0x49, 0xed,
	0xe5, 0xe8, 0xff, 0x2e, 0x2b, 0x38, 0x48, 0x23, 0x52, 0x01, 0xaa, 0xc2, 0x47, 0x47, 0x19, 0x62,
	0x96, 0x03, 0x73, 0xa7, 0x78, 0x9d, 0x32, 0x23, 0x0a, 0xd0, 0x26, 0x2a, 0xca, 0x0a, 0x38, 0xfe,
	0x68, 0x93, 0xce, 0x53, 0x19, 0x49, 0xff, 0x96, 0x10, 0x6d, 0x22, 0x65, 0x96, 0x96, 0x18, 0x36,
	0xc7, 0xcd, 0x89, 0x37, 0x1d, 0xd1, 0x4a, 0xa7, 0xb5, 0x4e, 0x17, 0xb5, 0x1e, 0x0e, 0x1c, 0x6d,
	0xcf, 0xfe, 0x35, 0xf9, 0x05, 0x92, 0x57, 0x62, 0xeb, 0x47, 0xb1, 0x0f, 0x92, 0x3b, 0xed, 0x81,
	0x78, 0x1b, 0x54, 0xab, 0x34, 0xc7, 0xcd, 0x52, 0xf0, 0x61, 0xdb, 0x99, 0xe7, 0x74, 0x67, 0x40,
	0xfa, 0xfc, 0x45, 0xdc, 0xbf, 0x41, 0xb2, 0x36, 0x02, 0x65, 0xb0, 0x1d, 0x71, 0xde, 0x08, 0x49,
	0x1d, 0x08, 0xb8, 0x7f, 0x47, 0xfa, 0x12, 0x39, 0xd8, 0x54, 0xc7, 0xa5, 0x4e, 0xf7, 0x52, 0x8f,
	0xc8, 0xe1, 0x70, 0xa6, 0x67, 0xc5, 0x2a, 0x61, 0x22, 0xbd, 0xb2, 0x89, 0xee, 0xc1, 0xc4, 0x22,
	0xd2, 0xab, 0x6f, 0x12, 0x56, 0x0c, 0xb8, 0x7f, 0x42, 0x7e, 0x63, 0x09, 0x2a, 0xb2, 0x80, 0xed,
	0xf4, 0xc6, 0xcd, 0xc9, 0x60, 0xde, 0x08, 0xbd, 0xed, 0xd3, 0x80, 0xfb, 0x67, 0xa4, 0xab, 0xcb,
	0x48, 0xea, 0x61, 0x7f, 0xdc, 0x9e, 0x78, 0xd3, 0xbf, 0x7b, 0x6f, 0xb1, 0xfb, 0x08, 0x2b, 0x62,
	0xd6, 0x21, 0x2d, 0xc1, 0x67, 0x37, 0x2f, 0x57, 0x99, 0x30, 0xaf, 0xeb, 0x98, 0x26, 0x58, 0x30,
	0x47, 0xa3, 0xca, 0xd8, 0x76, 0xf9, 0x19, 0x48, 0x56, 0xc6, 0x17, 0x19, 0xb2, 0x9d, 0xff, 0x21,
	0xee, 0xb9, 0xef, 0x7f, 0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x94, 0xbf, 0x5a, 0x53, 0x02,
	0x00, 0x00,
}