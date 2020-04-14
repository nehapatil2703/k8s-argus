// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/customer_label_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for [CustomerLabelService.GetCustomerLabel][google.ads.googleads.v2.services.CustomerLabelService.GetCustomerLabel].
type GetCustomerLabelRequest struct {
	// The resource name of the customer-label relationship to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerLabelRequest) Reset()         { *m = GetCustomerLabelRequest{} }
func (m *GetCustomerLabelRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerLabelRequest) ProtoMessage()    {}
func (*GetCustomerLabelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65bc127b03981ed2, []int{0}
}

func (m *GetCustomerLabelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerLabelRequest.Unmarshal(m, b)
}
func (m *GetCustomerLabelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerLabelRequest.Marshal(b, m, deterministic)
}
func (m *GetCustomerLabelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerLabelRequest.Merge(m, src)
}
func (m *GetCustomerLabelRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerLabelRequest.Size(m)
}
func (m *GetCustomerLabelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerLabelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerLabelRequest proto.InternalMessageInfo

func (m *GetCustomerLabelRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CustomerLabelService.MutateCustomerLabels][google.ads.googleads.v2.services.CustomerLabelService.MutateCustomerLabels].
type MutateCustomerLabelsRequest struct {
	// ID of the customer whose customer-label relationships are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on customer-label relationships.
	Operations []*CustomerLabelOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCustomerLabelsRequest) Reset()         { *m = MutateCustomerLabelsRequest{} }
func (m *MutateCustomerLabelsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerLabelsRequest) ProtoMessage()    {}
func (*MutateCustomerLabelsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65bc127b03981ed2, []int{1}
}

func (m *MutateCustomerLabelsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerLabelsRequest.Unmarshal(m, b)
}
func (m *MutateCustomerLabelsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerLabelsRequest.Marshal(b, m, deterministic)
}
func (m *MutateCustomerLabelsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerLabelsRequest.Merge(m, src)
}
func (m *MutateCustomerLabelsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerLabelsRequest.Size(m)
}
func (m *MutateCustomerLabelsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerLabelsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerLabelsRequest proto.InternalMessageInfo

func (m *MutateCustomerLabelsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCustomerLabelsRequest) GetOperations() []*CustomerLabelOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCustomerLabelsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCustomerLabelsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on a customer-label relationship.
type CustomerLabelOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CustomerLabelOperation_Create
	//	*CustomerLabelOperation_Remove
	Operation            isCustomerLabelOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *CustomerLabelOperation) Reset()         { *m = CustomerLabelOperation{} }
func (m *CustomerLabelOperation) String() string { return proto.CompactTextString(m) }
func (*CustomerLabelOperation) ProtoMessage()    {}
func (*CustomerLabelOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_65bc127b03981ed2, []int{2}
}

func (m *CustomerLabelOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerLabelOperation.Unmarshal(m, b)
}
func (m *CustomerLabelOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerLabelOperation.Marshal(b, m, deterministic)
}
func (m *CustomerLabelOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerLabelOperation.Merge(m, src)
}
func (m *CustomerLabelOperation) XXX_Size() int {
	return xxx_messageInfo_CustomerLabelOperation.Size(m)
}
func (m *CustomerLabelOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerLabelOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerLabelOperation proto.InternalMessageInfo

type isCustomerLabelOperation_Operation interface {
	isCustomerLabelOperation_Operation()
}

type CustomerLabelOperation_Create struct {
	Create *resources.CustomerLabel `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CustomerLabelOperation_Remove struct {
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*CustomerLabelOperation_Create) isCustomerLabelOperation_Operation() {}

func (*CustomerLabelOperation_Remove) isCustomerLabelOperation_Operation() {}

func (m *CustomerLabelOperation) GetOperation() isCustomerLabelOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CustomerLabelOperation) GetCreate() *resources.CustomerLabel {
	if x, ok := m.GetOperation().(*CustomerLabelOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CustomerLabelOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CustomerLabelOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CustomerLabelOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CustomerLabelOperation_Create)(nil),
		(*CustomerLabelOperation_Remove)(nil),
	}
}

// Response message for a customer labels mutate.
type MutateCustomerLabelsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCustomerLabelResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *MutateCustomerLabelsResponse) Reset()         { *m = MutateCustomerLabelsResponse{} }
func (m *MutateCustomerLabelsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerLabelsResponse) ProtoMessage()    {}
func (*MutateCustomerLabelsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65bc127b03981ed2, []int{3}
}

func (m *MutateCustomerLabelsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerLabelsResponse.Unmarshal(m, b)
}
func (m *MutateCustomerLabelsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerLabelsResponse.Marshal(b, m, deterministic)
}
func (m *MutateCustomerLabelsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerLabelsResponse.Merge(m, src)
}
func (m *MutateCustomerLabelsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerLabelsResponse.Size(m)
}
func (m *MutateCustomerLabelsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerLabelsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerLabelsResponse proto.InternalMessageInfo

func (m *MutateCustomerLabelsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCustomerLabelsResponse) GetResults() []*MutateCustomerLabelResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for a customer label mutate.
type MutateCustomerLabelResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCustomerLabelResult) Reset()         { *m = MutateCustomerLabelResult{} }
func (m *MutateCustomerLabelResult) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerLabelResult) ProtoMessage()    {}
func (*MutateCustomerLabelResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_65bc127b03981ed2, []int{4}
}

func (m *MutateCustomerLabelResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerLabelResult.Unmarshal(m, b)
}
func (m *MutateCustomerLabelResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerLabelResult.Marshal(b, m, deterministic)
}
func (m *MutateCustomerLabelResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerLabelResult.Merge(m, src)
}
func (m *MutateCustomerLabelResult) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerLabelResult.Size(m)
}
func (m *MutateCustomerLabelResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerLabelResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerLabelResult proto.InternalMessageInfo

func (m *MutateCustomerLabelResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCustomerLabelRequest)(nil), "google.ads.googleads.v2.services.GetCustomerLabelRequest")
	proto.RegisterType((*MutateCustomerLabelsRequest)(nil), "google.ads.googleads.v2.services.MutateCustomerLabelsRequest")
	proto.RegisterType((*CustomerLabelOperation)(nil), "google.ads.googleads.v2.services.CustomerLabelOperation")
	proto.RegisterType((*MutateCustomerLabelsResponse)(nil), "google.ads.googleads.v2.services.MutateCustomerLabelsResponse")
	proto.RegisterType((*MutateCustomerLabelResult)(nil), "google.ads.googleads.v2.services.MutateCustomerLabelResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/customer_label_service.proto", fileDescriptor_65bc127b03981ed2)
}

var fileDescriptor_65bc127b03981ed2 = []byte{
	// 668 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdf, 0x6a, 0xd4, 0x4e,
	0x14, 0xfe, 0x25, 0xfb, 0xa3, 0xda, 0xd9, 0xfa, 0x87, 0xb1, 0xda, 0x74, 0x5b, 0x70, 0x89, 0x05,
	0xcb, 0x5e, 0x4c, 0x4a, 0x0a, 0x52, 0x53, 0xb6, 0xb8, 0x15, 0xdb, 0x2a, 0x6a, 0x4b, 0x8a, 0x45,
	0x64, 0x61, 0x99, 0x26, 0xe3, 0x12, 0x48, 0x32, 0x71, 0x66, 0xb2, 0x50, 0x4a, 0x41, 0xbc, 0xf5,
	0xd2, 0x37, 0xf0, 0xd2, 0x37, 0xf0, 0xc2, 0x6b, 0xa1, 0xb7, 0xe2, 0x1b, 0x88, 0x17, 0x3e, 0x85,
	0x24, 0x93, 0x49, 0x9b, 0x75, 0x97, 0xd5, 0xde, 0x9d, 0x9c, 0x39, 0xe7, 0xfb, 0xce, 0x77, 0xfe,
	0x04, 0xb4, 0xfb, 0x94, 0xf6, 0x43, 0x62, 0x61, 0x9f, 0x5b, 0xd2, 0xcc, 0xac, 0x81, 0x6d, 0x71,
	0xc2, 0x06, 0x81, 0x47, 0xb8, 0xe5, 0xa5, 0x5c, 0xd0, 0x88, 0xb0, 0x5e, 0x88, 0x0f, 0x49, 0xd8,
	0x2b, 0xfc, 0x28, 0x61, 0x54, 0x50, 0xd8, 0x94, 0x39, 0x08, 0xfb, 0x1c, 0x95, 0xe9, 0x68, 0x60,
	0x23, 0x95, 0xde, 0xb8, 0x37, 0x8e, 0x80, 0x11, 0x4e, 0x53, 0xf6, 0x27, 0x83, 0x44, 0x6e, 0x2c,
	0xaa, 0xbc, 0x24, 0xb0, 0x70, 0x1c, 0x53, 0x81, 0x45, 0x40, 0x63, 0x5e, 0xbc, 0xce, 0x15, 0xaf,
	0x2c, 0xf1, 0x2c, 0x2e, 0xb0, 0x48, 0x87, 0x1f, 0xb2, 0x34, 0x2f, 0x0c, 0x48, 0x2c, 0xe4, 0x83,
	0xb9, 0x01, 0xe6, 0xb6, 0x89, 0x78, 0x58, 0x50, 0x3d, 0xcd, 0x98, 0x5c, 0xf2, 0x26, 0x25, 0x5c,
	0xc0, 0x3b, 0xe0, 0x8a, 0x2a, 0xa6, 0x17, 0xe3, 0x88, 0x18, 0x5a, 0x53, 0x5b, 0x9e, 0x76, 0x67,
	0x94, 0xf3, 0x39, 0x8e, 0x88, 0xf9, 0x53, 0x03, 0x0b, 0xcf, 0x52, 0x81, 0x05, 0xa9, 0x60, 0x70,
	0x05, 0x72, 0x1b, 0xd4, 0x4b, 0x1d, 0x81, 0x5f, 0x40, 0x00, 0xe5, 0x7a, 0xec, 0xc3, 0x97, 0x00,
	0xd0, 0x84, 0x30, 0x29, 0xc3, 0xd0, 0x9b, 0xb5, 0xe5, 0xba, 0xbd, 0x86, 0x26, 0xf5, 0x0f, 0x55,
	0xd8, 0x76, 0x15, 0x80, 0x7b, 0x0e, 0x0b, 0xde, 0x05, 0xd7, 0x12, 0xcc, 0x44, 0x80, 0xc3, 0xde,
	0x6b, 0x1c, 0x84, 0x29, 0x23, 0x46, 0xad, 0xa9, 0x2d, 0x5f, 0x76, 0xaf, 0x16, 0xee, 0x2d, 0xe9,
	0xcd, 0x84, 0x0e, 0x70, 0x18, 0xf8, 0x58, 0x90, 0x1e, 0x8d, 0xc3, 0x23, 0xe3, 0xff, 0x3c, 0x6c,
	0x46, 0x39, 0x77, 0xe3, 0xf0, 0xc8, 0x7c, 0xaf, 0x81, 0x5b, 0xa3, 0x49, 0xe1, 0x13, 0x30, 0xe5,
	0x31, 0x82, 0x85, 0xec, 0x50, 0xdd, 0x5e, 0x19, 0x5b, 0x7e, 0x39, 0xdc, 0x6a, 0xfd, 0x3b, 0xff,
	0xb9, 0x05, 0x02, 0x34, 0xc0, 0x14, 0x23, 0x11, 0x1d, 0x10, 0x43, 0xcf, 0x5a, 0x95, 0xbd, 0xc8,
	0xef, 0xcd, 0x3a, 0x98, 0x2e, 0xc5, 0x99, 0x5f, 0x34, 0xb0, 0x38, 0xba, 0xed, 0x3c, 0xa1, 0x31,
	0x27, 0x70, 0x0b, 0xdc, 0x1c, 0x12, 0xdf, 0x23, 0x8c, 0x51, 0x96, 0xb7, 0xa0, 0x6e, 0x43, 0x55,
	0x22, 0x4b, 0x3c, 0xb4, 0x9f, 0x6f, 0x8a, 0x7b, 0xa3, 0xda, 0x96, 0x47, 0x59, 0x38, 0x7c, 0x01,
	0x2e, 0x31, 0xc2, 0xd3, 0x50, 0xa8, 0xd9, 0xac, 0x4f, 0x9e, 0xcd, 0x88, 0xc2, 0xdc, 0x1c, 0xc3,
	0x55, 0x58, 0xe6, 0x03, 0x30, 0x3f, 0x36, 0xea, 0xaf, 0x16, 0xcf, 0xfe, 0x5a, 0x03, 0xb3, 0x95,
	0xe4, 0x7d, 0x49, 0x0f, 0x3f, 0x6b, 0xe0, 0xfa, 0xf0, 0x4a, 0xc3, 0xfb, 0x93, 0xab, 0x1e, 0x73,
	0x06, 0x8d, 0x7f, 0x9e, 0xa6, 0xb9, 0xf6, 0xee, 0xdb, 0x8f, 0x0f, 0xba, 0x0d, 0x57, 0xb2, 0x7b,
	0x3e, 0xae, 0x48, 0x69, 0xab, 0xcd, 0xe7, 0x56, 0xab, 0x3c, 0x70, 0x39, 0x3a, 0xab, 0x75, 0x02,
	0xbf, 0x6b, 0x60, 0x76, 0xd4, 0x58, 0x61, 0xfb, 0x42, 0x5d, 0x57, 0x57, 0xd8, 0xd8, 0xb8, 0x68,
	0xba, 0xdc, 0x26, 0x73, 0x23, 0x57, 0xb4, 0x66, 0xae, 0x66, 0x8a, 0xce, 0x24, 0x1c, 0x9f, 0x3b,
	0xed, 0x76, 0xeb, 0x64, 0x48, 0x90, 0x13, 0xe5, 0x90, 0x8e, 0xd6, 0x6a, 0x2c, 0x9c, 0x76, 0x8c,
	0x33, 0xda, 0xc2, 0x4a, 0x02, 0x8e, 0x3c, 0x1a, 0x6d, 0xbe, 0xd5, 0xc1, 0x92, 0x47, 0xa3, 0x89,
	0x25, 0x6e, 0xce, 0x8f, 0x9a, 0xf7, 0x5e, 0xf6, 0x1b, 0xdb, 0xd3, 0x5e, 0xed, 0x14, 0xe9, 0x7d,
	0x1a, 0xe2, 0xb8, 0x8f, 0x28, 0xeb, 0x5b, 0x7d, 0x12, 0xe7, 0x3f, 0x39, 0xeb, 0x8c, 0x70, 0xfc,
	0x0f, 0x7d, 0x5d, 0x19, 0x1f, 0xf5, 0xda, 0x76, 0xa7, 0xf3, 0x49, 0x6f, 0x6e, 0x4b, 0xc0, 0x8e,
	0xcf, 0x91, 0x34, 0x33, 0xeb, 0xc0, 0x46, 0x05, 0x31, 0x3f, 0x55, 0x21, 0xdd, 0x8e, 0xcf, 0xbb,
	0x65, 0x48, 0xf7, 0xc0, 0xee, 0xaa, 0x90, 0x5f, 0xfa, 0x92, 0xf4, 0x3b, 0x4e, 0xc7, 0xe7, 0x8e,
	0x53, 0x06, 0x39, 0xce, 0x81, 0xed, 0x38, 0x2a, 0xec, 0x70, 0x2a, 0xaf, 0x73, 0xf5, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x32, 0x77, 0x45, 0x5f, 0x77, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerLabelServiceClient is the client API for CustomerLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerLabelServiceClient interface {
	// Returns the requested customer-label relationship in full detail.
	GetCustomerLabel(ctx context.Context, in *GetCustomerLabelRequest, opts ...grpc.CallOption) (*resources.CustomerLabel, error)
	// Creates and removes customer-label relationships.
	// Operation statuses are returned.
	MutateCustomerLabels(ctx context.Context, in *MutateCustomerLabelsRequest, opts ...grpc.CallOption) (*MutateCustomerLabelsResponse, error)
}

type customerLabelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerLabelServiceClient(cc grpc.ClientConnInterface) CustomerLabelServiceClient {
	return &customerLabelServiceClient{cc}
}

func (c *customerLabelServiceClient) GetCustomerLabel(ctx context.Context, in *GetCustomerLabelRequest, opts ...grpc.CallOption) (*resources.CustomerLabel, error) {
	out := new(resources.CustomerLabel)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CustomerLabelService/GetCustomerLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerLabelServiceClient) MutateCustomerLabels(ctx context.Context, in *MutateCustomerLabelsRequest, opts ...grpc.CallOption) (*MutateCustomerLabelsResponse, error) {
	out := new(MutateCustomerLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CustomerLabelService/MutateCustomerLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerLabelServiceServer is the server API for CustomerLabelService service.
type CustomerLabelServiceServer interface {
	// Returns the requested customer-label relationship in full detail.
	GetCustomerLabel(context.Context, *GetCustomerLabelRequest) (*resources.CustomerLabel, error)
	// Creates and removes customer-label relationships.
	// Operation statuses are returned.
	MutateCustomerLabels(context.Context, *MutateCustomerLabelsRequest) (*MutateCustomerLabelsResponse, error)
}

// UnimplementedCustomerLabelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerLabelServiceServer struct {
}

func (*UnimplementedCustomerLabelServiceServer) GetCustomerLabel(ctx context.Context, req *GetCustomerLabelRequest) (*resources.CustomerLabel, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCustomerLabel not implemented")
}
func (*UnimplementedCustomerLabelServiceServer) MutateCustomerLabels(ctx context.Context, req *MutateCustomerLabelsRequest) (*MutateCustomerLabelsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCustomerLabels not implemented")
}

func RegisterCustomerLabelServiceServer(s *grpc.Server, srv CustomerLabelServiceServer) {
	s.RegisterService(&_CustomerLabelService_serviceDesc, srv)
}

func _CustomerLabelService_GetCustomerLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerLabelServiceServer).GetCustomerLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CustomerLabelService/GetCustomerLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerLabelServiceServer).GetCustomerLabel(ctx, req.(*GetCustomerLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerLabelService_MutateCustomerLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerLabelServiceServer).MutateCustomerLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CustomerLabelService/MutateCustomerLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerLabelServiceServer).MutateCustomerLabels(ctx, req.(*MutateCustomerLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerLabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.CustomerLabelService",
	HandlerType: (*CustomerLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerLabel",
			Handler:    _CustomerLabelService_GetCustomerLabel_Handler,
		},
		{
			MethodName: "MutateCustomerLabels",
			Handler:    _CustomerLabelService_MutateCustomerLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/customer_label_service.proto",
}
