// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nodestats.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AuditCheckRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuditCheckRequest) Reset()         { *m = AuditCheckRequest{} }
func (m *AuditCheckRequest) String() string { return proto.CompactTextString(m) }
func (*AuditCheckRequest) ProtoMessage()    {}
func (*AuditCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{0}
}
func (m *AuditCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditCheckRequest.Unmarshal(m, b)
}
func (m *AuditCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditCheckRequest.Marshal(b, m, deterministic)
}
func (m *AuditCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditCheckRequest.Merge(m, src)
}
func (m *AuditCheckRequest) XXX_Size() int {
	return xxx_messageInfo_AuditCheckRequest.Size(m)
}
func (m *AuditCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuditCheckRequest proto.InternalMessageInfo

type AuditCheckResponse struct {
	TotalCount           int64    `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	SuccessCount         int64    `protobuf:"varint,2,opt,name=success_count,json=successCount,proto3" json:"success_count,omitempty"`
	ReputationAlpha      float64  `protobuf:"fixed64,3,opt,name=reputation_alpha,json=reputationAlpha,proto3" json:"reputation_alpha,omitempty"`
	ReputationBeta       float64  `protobuf:"fixed64,4,opt,name=reputation_beta,json=reputationBeta,proto3" json:"reputation_beta,omitempty"`
	ReputationScore      float64  `protobuf:"fixed64,5,opt,name=reputation_score,json=reputationScore,proto3" json:"reputation_score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuditCheckResponse) Reset()         { *m = AuditCheckResponse{} }
func (m *AuditCheckResponse) String() string { return proto.CompactTextString(m) }
func (*AuditCheckResponse) ProtoMessage()    {}
func (*AuditCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{1}
}
func (m *AuditCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditCheckResponse.Unmarshal(m, b)
}
func (m *AuditCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditCheckResponse.Marshal(b, m, deterministic)
}
func (m *AuditCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditCheckResponse.Merge(m, src)
}
func (m *AuditCheckResponse) XXX_Size() int {
	return xxx_messageInfo_AuditCheckResponse.Size(m)
}
func (m *AuditCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuditCheckResponse proto.InternalMessageInfo

func (m *AuditCheckResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *AuditCheckResponse) GetSuccessCount() int64 {
	if m != nil {
		return m.SuccessCount
	}
	return 0
}

func (m *AuditCheckResponse) GetReputationAlpha() float64 {
	if m != nil {
		return m.ReputationAlpha
	}
	return 0
}

func (m *AuditCheckResponse) GetReputationBeta() float64 {
	if m != nil {
		return m.ReputationBeta
	}
	return 0
}

func (m *AuditCheckResponse) GetReputationScore() float64 {
	if m != nil {
		return m.ReputationScore
	}
	return 0
}

type UptimeCheckRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UptimeCheckRequest) Reset()         { *m = UptimeCheckRequest{} }
func (m *UptimeCheckRequest) String() string { return proto.CompactTextString(m) }
func (*UptimeCheckRequest) ProtoMessage()    {}
func (*UptimeCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{2}
}
func (m *UptimeCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UptimeCheckRequest.Unmarshal(m, b)
}
func (m *UptimeCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UptimeCheckRequest.Marshal(b, m, deterministic)
}
func (m *UptimeCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UptimeCheckRequest.Merge(m, src)
}
func (m *UptimeCheckRequest) XXX_Size() int {
	return xxx_messageInfo_UptimeCheckRequest.Size(m)
}
func (m *UptimeCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UptimeCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UptimeCheckRequest proto.InternalMessageInfo

type UptimeCheckResponse struct {
	TotalCount           int64    `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	SuccessCount         int64    `protobuf:"varint,2,opt,name=success_count,json=successCount,proto3" json:"success_count,omitempty"`
	ReputationAlpha      float64  `protobuf:"fixed64,3,opt,name=reputation_alpha,json=reputationAlpha,proto3" json:"reputation_alpha,omitempty"`
	ReputationBeta       float64  `protobuf:"fixed64,4,opt,name=reputation_beta,json=reputationBeta,proto3" json:"reputation_beta,omitempty"`
	ReputationScore      float64  `protobuf:"fixed64,5,opt,name=reputation_score,json=reputationScore,proto3" json:"reputation_score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UptimeCheckResponse) Reset()         { *m = UptimeCheckResponse{} }
func (m *UptimeCheckResponse) String() string { return proto.CompactTextString(m) }
func (*UptimeCheckResponse) ProtoMessage()    {}
func (*UptimeCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{3}
}
func (m *UptimeCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UptimeCheckResponse.Unmarshal(m, b)
}
func (m *UptimeCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UptimeCheckResponse.Marshal(b, m, deterministic)
}
func (m *UptimeCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UptimeCheckResponse.Merge(m, src)
}
func (m *UptimeCheckResponse) XXX_Size() int {
	return xxx_messageInfo_UptimeCheckResponse.Size(m)
}
func (m *UptimeCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UptimeCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UptimeCheckResponse proto.InternalMessageInfo

func (m *UptimeCheckResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *UptimeCheckResponse) GetSuccessCount() int64 {
	if m != nil {
		return m.SuccessCount
	}
	return 0
}

func (m *UptimeCheckResponse) GetReputationAlpha() float64 {
	if m != nil {
		return m.ReputationAlpha
	}
	return 0
}

func (m *UptimeCheckResponse) GetReputationBeta() float64 {
	if m != nil {
		return m.ReputationBeta
	}
	return 0
}

func (m *UptimeCheckResponse) GetReputationScore() float64 {
	if m != nil {
		return m.ReputationScore
	}
	return 0
}

type DailyStorageUsageRequest struct {
	From                 time.Time `protobuf:"bytes,1,opt,name=from,proto3,stdtime" json:"from"`
	To                   time.Time `protobuf:"bytes,2,opt,name=to,proto3,stdtime" json:"to"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DailyStorageUsageRequest) Reset()         { *m = DailyStorageUsageRequest{} }
func (m *DailyStorageUsageRequest) String() string { return proto.CompactTextString(m) }
func (*DailyStorageUsageRequest) ProtoMessage()    {}
func (*DailyStorageUsageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{4}
}
func (m *DailyStorageUsageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DailyStorageUsageRequest.Unmarshal(m, b)
}
func (m *DailyStorageUsageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DailyStorageUsageRequest.Marshal(b, m, deterministic)
}
func (m *DailyStorageUsageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DailyStorageUsageRequest.Merge(m, src)
}
func (m *DailyStorageUsageRequest) XXX_Size() int {
	return xxx_messageInfo_DailyStorageUsageRequest.Size(m)
}
func (m *DailyStorageUsageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DailyStorageUsageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DailyStorageUsageRequest proto.InternalMessageInfo

func (m *DailyStorageUsageRequest) GetFrom() time.Time {
	if m != nil {
		return m.From
	}
	return time.Time{}
}

func (m *DailyStorageUsageRequest) GetTo() time.Time {
	if m != nil {
		return m.To
	}
	return time.Time{}
}

type DailyStorageUsageResponse struct {
	NodeId               NodeID                                    `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	DailyStorageUsage    []*DailyStorageUsageResponse_StorageUsage `protobuf:"bytes,2,rep,name=daily_storage_usage,json=dailyStorageUsage,proto3" json:"daily_storage_usage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *DailyStorageUsageResponse) Reset()         { *m = DailyStorageUsageResponse{} }
func (m *DailyStorageUsageResponse) String() string { return proto.CompactTextString(m) }
func (*DailyStorageUsageResponse) ProtoMessage()    {}
func (*DailyStorageUsageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{5}
}
func (m *DailyStorageUsageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DailyStorageUsageResponse.Unmarshal(m, b)
}
func (m *DailyStorageUsageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DailyStorageUsageResponse.Marshal(b, m, deterministic)
}
func (m *DailyStorageUsageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DailyStorageUsageResponse.Merge(m, src)
}
func (m *DailyStorageUsageResponse) XXX_Size() int {
	return xxx_messageInfo_DailyStorageUsageResponse.Size(m)
}
func (m *DailyStorageUsageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DailyStorageUsageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DailyStorageUsageResponse proto.InternalMessageInfo

func (m *DailyStorageUsageResponse) GetDailyStorageUsage() []*DailyStorageUsageResponse_StorageUsage {
	if m != nil {
		return m.DailyStorageUsage
	}
	return nil
}

type DailyStorageUsageResponse_StorageUsage struct {
	AtRestTotal          float64   `protobuf:"fixed64,2,opt,name=at_rest_total,json=atRestTotal,proto3" json:"at_rest_total,omitempty"`
	TimeStamp            time.Time `protobuf:"bytes,3,opt,name=time_stamp,json=timeStamp,proto3,stdtime" json:"time_stamp"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DailyStorageUsageResponse_StorageUsage) Reset() {
	*m = DailyStorageUsageResponse_StorageUsage{}
}
func (m *DailyStorageUsageResponse_StorageUsage) String() string { return proto.CompactTextString(m) }
func (*DailyStorageUsageResponse_StorageUsage) ProtoMessage()    {}
func (*DailyStorageUsageResponse_StorageUsage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{5, 0}
}
func (m *DailyStorageUsageResponse_StorageUsage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DailyStorageUsageResponse_StorageUsage.Unmarshal(m, b)
}
func (m *DailyStorageUsageResponse_StorageUsage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DailyStorageUsageResponse_StorageUsage.Marshal(b, m, deterministic)
}
func (m *DailyStorageUsageResponse_StorageUsage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DailyStorageUsageResponse_StorageUsage.Merge(m, src)
}
func (m *DailyStorageUsageResponse_StorageUsage) XXX_Size() int {
	return xxx_messageInfo_DailyStorageUsageResponse_StorageUsage.Size(m)
}
func (m *DailyStorageUsageResponse_StorageUsage) XXX_DiscardUnknown() {
	xxx_messageInfo_DailyStorageUsageResponse_StorageUsage.DiscardUnknown(m)
}

var xxx_messageInfo_DailyStorageUsageResponse_StorageUsage proto.InternalMessageInfo

func (m *DailyStorageUsageResponse_StorageUsage) GetAtRestTotal() float64 {
	if m != nil {
		return m.AtRestTotal
	}
	return 0
}

func (m *DailyStorageUsageResponse_StorageUsage) GetTimeStamp() time.Time {
	if m != nil {
		return m.TimeStamp
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*AuditCheckRequest)(nil), "nodestats.AuditCheckRequest")
	proto.RegisterType((*AuditCheckResponse)(nil), "nodestats.AuditCheckResponse")
	proto.RegisterType((*UptimeCheckRequest)(nil), "nodestats.UptimeCheckRequest")
	proto.RegisterType((*UptimeCheckResponse)(nil), "nodestats.UptimeCheckResponse")
	proto.RegisterType((*DailyStorageUsageRequest)(nil), "nodestats.DailyStorageUsageRequest")
	proto.RegisterType((*DailyStorageUsageResponse)(nil), "nodestats.DailyStorageUsageResponse")
	proto.RegisterType((*DailyStorageUsageResponse_StorageUsage)(nil), "nodestats.DailyStorageUsageResponse.StorageUsage")
}

func init() { proto.RegisterFile("nodestats.proto", fileDescriptor_e0b184ee117142aa) }

var fileDescriptor_e0b184ee117142aa = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x39, 0x37, 0x04, 0xf2, 0x9c, 0xb6, 0xe4, 0xc2, 0x60, 0x2c, 0xc0, 0x91, 0x8b, 0xd4,
	0xb0, 0xb8, 0x22, 0x30, 0xb0, 0x36, 0xe9, 0x12, 0x09, 0x31, 0x38, 0xed, 0xc2, 0xc0, 0xe9, 0x62,
	0x5f, 0x5d, 0x8b, 0x24, 0x67, 0x7c, 0xcf, 0x42, 0xec, 0x4c, 0x4c, 0x7c, 0x04, 0x3e, 0x0e, 0x5f,
	0x81, 0x0e, 0xe5, 0xab, 0xa0, 0x3b, 0xbb, 0xc4, 0x49, 0x08, 0xca, 0xcc, 0xe8, 0xdf, 0xfd, 0xdf,
	0xf3, 0xfd, 0xdf, 0xfd, 0x1f, 0x1c, 0x2e, 0x64, 0x2c, 0x14, 0x72, 0x54, 0x41, 0x96, 0x4b, 0x94,
	0xb4, 0xf5, 0x07, 0xb8, 0x90, 0xc8, 0x44, 0x96, 0xd8, 0xf5, 0x12, 0x29, 0x93, 0x99, 0x38, 0x31,
	0x5f, 0xd3, 0xe2, 0xf2, 0x04, 0xd3, 0xb9, 0x96, 0xcd, 0xb3, 0x52, 0xe0, 0x77, 0xa1, 0x73, 0x5a,
	0xc4, 0x29, 0x8e, 0xae, 0x44, 0xf4, 0x21, 0x14, 0x1f, 0x0b, 0xa1, 0xd0, 0xff, 0x49, 0x80, 0xd6,
	0xa9, 0xca, 0xe4, 0x42, 0x09, 0xea, 0x81, 0x8d, 0x12, 0xf9, 0x8c, 0x45, 0xb2, 0x58, 0xa0, 0x43,
	0x7a, 0xa4, 0xbf, 0x17, 0x82, 0x41, 0x23, 0x4d, 0xe8, 0x11, 0xec, 0xab, 0x22, 0x8a, 0x84, 0x52,
	0x95, 0xc4, 0x32, 0x92, 0x76, 0x05, 0x4b, 0xd1, 0x73, 0x78, 0x90, 0x8b, 0xac, 0x40, 0x8e, 0xa9,
	0x5c, 0x30, 0x3e, 0xcb, 0xae, 0xb8, 0xb3, 0xd7, 0x23, 0x7d, 0x12, 0x1e, 0x2e, 0xf9, 0xa9, 0xc6,
	0xf4, 0x18, 0x6a, 0x88, 0x4d, 0x05, 0x72, 0xa7, 0x61, 0x94, 0x07, 0x4b, 0x3c, 0x14, 0xc8, 0xd7,
	0x7a, 0xaa, 0x48, 0xe6, 0xc2, 0xb9, 0xbb, 0xde, 0x73, 0xa2, 0xb1, 0xff, 0x10, 0xe8, 0x45, 0xa6,
	0xa7, 0xb0, 0xe2, 0xf8, 0x9a, 0x40, 0x77, 0x05, 0xff, 0x4f, 0x96, 0xbf, 0x12, 0x70, 0xce, 0x78,
	0x3a, 0xfb, 0x3c, 0x41, 0x99, 0xf3, 0x44, 0x5c, 0x28, 0x9e, 0x88, 0xca, 0x39, 0x7d, 0x0d, 0x8d,
	0xcb, 0x5c, 0xce, 0x8d, 0x35, 0x7b, 0xe0, 0x06, 0x65, 0x60, 0x82, 0xdb, 0xc0, 0x04, 0xe7, 0xb7,
	0x81, 0x19, 0xde, 0xff, 0x71, 0xe3, 0xdd, 0xf9, 0xf6, 0xcb, 0x23, 0xa1, 0xa9, 0xa0, 0xaf, 0xc0,
	0x42, 0x69, 0xfc, 0xee, 0x5a, 0x67, 0xa1, 0xf4, 0xbf, 0x5b, 0xf0, 0xe8, 0x2f, 0x97, 0xa9, 0xe6,
	0x7d, 0x0c, 0xf7, 0x74, 0x90, 0x59, 0x1a, 0x9b, 0x0b, 0xb5, 0x87, 0x07, 0xba, 0xf8, 0xfa, 0xc6,
	0x6b, 0xbe, 0x95, 0xb1, 0x18, 0x9f, 0x85, 0x4d, 0x7d, 0x3c, 0x8e, 0x29, 0x87, 0x6e, 0xac, 0xbb,
	0x30, 0x55, 0xb6, 0x61, 0x85, 0xee, 0xe3, 0x58, 0xbd, 0xbd, 0xbe, 0x3d, 0x78, 0x11, 0x2c, 0xd7,
	0x63, 0xeb, 0xbf, 0x82, 0x15, 0xd8, 0x89, 0xd7, 0x75, 0xee, 0x27, 0x68, 0xd7, 0xbf, 0xa9, 0x0f,
	0xfb, 0x1c, 0x59, 0x2e, 0x14, 0x32, 0x13, 0x00, 0x63, 0x9d, 0x84, 0x36, 0xc7, 0x50, 0x28, 0x3c,
	0xd7, 0x88, 0x8e, 0x00, 0x74, 0x88, 0x98, 0x71, 0x6e, 0xde, 0x78, 0xd7, 0xd9, 0xb4, 0x74, 0xdd,
	0x44, 0xc3, 0xc1, 0x17, 0x0b, 0x5a, 0xda, 0xee, 0x44, 0x1b, 0xa0, 0x63, 0x80, 0xe5, 0x2e, 0xd2,
	0xc7, 0x35, 0x6b, 0x1b, 0x8b, 0xeb, 0x3e, 0xd9, 0x72, 0x5a, 0x4d, 0xf7, 0x0d, 0xd8, 0xb5, 0x90,
	0xd3, 0xba, 0x7a, 0x73, 0x27, 0xdc, 0xa7, 0xdb, 0x8e, 0xab, 0x6e, 0xef, 0xa1, 0xb3, 0x31, 0x5c,
	0x7a, 0xf4, 0xef, 0xd1, 0x97, 0x9d, 0x9f, 0xed, 0xf2, 0x3e, 0xc3, 0xc6, 0x3b, 0x2b, 0x9b, 0x4e,
	0x9b, 0x66, 0x6a, 0x2f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x2d, 0xea, 0x1a, 0xf2, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeStatsClient is the client API for NodeStats service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeStatsClient interface {
	AuditCheck(ctx context.Context, in *AuditCheckRequest, opts ...grpc.CallOption) (*AuditCheckResponse, error)
	UptimeCheck(ctx context.Context, in *UptimeCheckRequest, opts ...grpc.CallOption) (*UptimeCheckResponse, error)
	DailyStorageUsage(ctx context.Context, in *DailyStorageUsageRequest, opts ...grpc.CallOption) (*DailyStorageUsageResponse, error)
}

type nodeStatsClient struct {
	cc *grpc.ClientConn
}

func NewNodeStatsClient(cc *grpc.ClientConn) NodeStatsClient {
	return &nodeStatsClient{cc}
}

func (c *nodeStatsClient) AuditCheck(ctx context.Context, in *AuditCheckRequest, opts ...grpc.CallOption) (*AuditCheckResponse, error) {
	out := new(AuditCheckResponse)
	err := c.cc.Invoke(ctx, "/nodestats.NodeStats/AuditCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeStatsClient) UptimeCheck(ctx context.Context, in *UptimeCheckRequest, opts ...grpc.CallOption) (*UptimeCheckResponse, error) {
	out := new(UptimeCheckResponse)
	err := c.cc.Invoke(ctx, "/nodestats.NodeStats/UptimeCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeStatsClient) DailyStorageUsage(ctx context.Context, in *DailyStorageUsageRequest, opts ...grpc.CallOption) (*DailyStorageUsageResponse, error) {
	out := new(DailyStorageUsageResponse)
	err := c.cc.Invoke(ctx, "/nodestats.NodeStats/DailyStorageUsage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeStatsServer is the server API for NodeStats service.
type NodeStatsServer interface {
	AuditCheck(context.Context, *AuditCheckRequest) (*AuditCheckResponse, error)
	UptimeCheck(context.Context, *UptimeCheckRequest) (*UptimeCheckResponse, error)
	DailyStorageUsage(context.Context, *DailyStorageUsageRequest) (*DailyStorageUsageResponse, error)
}

func RegisterNodeStatsServer(s *grpc.Server, srv NodeStatsServer) {
	s.RegisterService(&_NodeStats_serviceDesc, srv)
}

func _NodeStats_AuditCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeStatsServer).AuditCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodestats.NodeStats/AuditCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeStatsServer).AuditCheck(ctx, req.(*AuditCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeStats_UptimeCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UptimeCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeStatsServer).UptimeCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodestats.NodeStats/UptimeCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeStatsServer).UptimeCheck(ctx, req.(*UptimeCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeStats_DailyStorageUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DailyStorageUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeStatsServer).DailyStorageUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodestats.NodeStats/DailyStorageUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeStatsServer).DailyStorageUsage(ctx, req.(*DailyStorageUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeStats_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nodestats.NodeStats",
	HandlerType: (*NodeStatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuditCheck",
			Handler:    _NodeStats_AuditCheck_Handler,
		},
		{
			MethodName: "UptimeCheck",
			Handler:    _NodeStats_UptimeCheck_Handler,
		},
		{
			MethodName: "DailyStorageUsage",
			Handler:    _NodeStats_DailyStorageUsage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodestats.proto",
}