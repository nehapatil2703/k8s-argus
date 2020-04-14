// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/topic_constant.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Use topics to target or exclude placements in the Google Display Network
// based on the category into which the placement falls (for example,
// "Pets & Animals/Pets/Dogs").
type TopicConstant struct {
	// The resource name of the topic constant.
	// topic constant resource names have the form:
	//
	// `topicConstants/{topic_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the topic.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Resource name of parent of the topic constant.
	TopicConstantParent *wrappers.StringValue `protobuf:"bytes,3,opt,name=topic_constant_parent,json=topicConstantParent,proto3" json:"topic_constant_parent,omitempty"`
	// The category to target or exclude. Each subsequent element in the array
	// describes a more specific sub-category. For example,
	// {"Pets & Animals", "Pets", "Dogs"} represents the
	// "Pets & Animals/Pets/Dogs" category. List of available topic categories at
	// https://developers.google.com/adwords/api/docs/appendix/verticals
	Path                 []*wrappers.StringValue `protobuf:"bytes,4,rep,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *TopicConstant) Reset()         { *m = TopicConstant{} }
func (m *TopicConstant) String() string { return proto.CompactTextString(m) }
func (*TopicConstant) ProtoMessage()    {}
func (*TopicConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_87af0f21f4e047c9, []int{0}
}

func (m *TopicConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicConstant.Unmarshal(m, b)
}
func (m *TopicConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicConstant.Marshal(b, m, deterministic)
}
func (m *TopicConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicConstant.Merge(m, src)
}
func (m *TopicConstant) XXX_Size() int {
	return xxx_messageInfo_TopicConstant.Size(m)
}
func (m *TopicConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicConstant.DiscardUnknown(m)
}

var xxx_messageInfo_TopicConstant proto.InternalMessageInfo

func (m *TopicConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *TopicConstant) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *TopicConstant) GetTopicConstantParent() *wrappers.StringValue {
	if m != nil {
		return m.TopicConstantParent
	}
	return nil
}

func (m *TopicConstant) GetPath() []*wrappers.StringValue {
	if m != nil {
		return m.Path
	}
	return nil
}

func init() {
	proto.RegisterType((*TopicConstant)(nil), "google.ads.googleads.v3.resources.TopicConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/topic_constant.proto", fileDescriptor_87af0f21f4e047c9)
}

var fileDescriptor_87af0f21f4e047c9 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x6b, 0xdb, 0x30,
	0x18, 0xc6, 0xb1, 0x13, 0x06, 0xd3, 0x96, 0x8b, 0xc7, 0xc0, 0xcb, 0xc2, 0x96, 0x6c, 0x64, 0x04,
	0x06, 0xd2, 0x98, 0x47, 0x0e, 0xda, 0xc9, 0xd9, 0x21, 0x64, 0x8c, 0x61, 0xb2, 0xe1, 0x43, 0x31,
	0x04, 0xc5, 0x56, 0x5d, 0x43, 0x2c, 0x09, 0x49, 0x49, 0x0f, 0xa5, 0x5f, 0xa6, 0xc7, 0x5e, 0xfb,
	0x2d, 0xfa, 0x51, 0xf2, 0x25, 0x5a, 0xfc, 0x4f, 0x89, 0x29, 0x34, 0xb7, 0xc7, 0x7e, 0x9f, 0xe7,
	0xd1, 0xcf, 0xaf, 0x05, 0xa6, 0x29, 0xe7, 0xe9, 0x86, 0x22, 0x92, 0x28, 0x54, 0xc9, 0x42, 0xed,
	0x3c, 0x24, 0xa9, 0xe2, 0x5b, 0x19, 0x53, 0x85, 0x34, 0x17, 0x59, 0xbc, 0x8a, 0x39, 0x53, 0x9a,
	0x30, 0x0d, 0x85, 0xe4, 0x9a, 0x3b, 0xa3, 0xca, 0x0c, 0x49, 0xa2, 0xa0, 0xc9, 0xc1, 0x9d, 0x07,
	0x4d, 0xae, 0xff, 0xae, 0xa9, 0x16, 0x99, 0x69, 0xab, 0xd2, 0xfd, 0x0f, 0xf5, 0xa8, 0x7c, 0x5a,
	0x6f, 0xcf, 0xd1, 0xa5, 0x24, 0x42, 0x50, 0xa9, 0xea, 0xf9, 0xe0, 0x28, 0x4a, 0x18, 0xe3, 0x9a,
	0xe8, 0x8c, 0xb3, 0x7a, 0xfa, 0xe9, 0xce, 0x06, 0xbd, 0xff, 0x05, 0xd4, 0xaf, 0x9a, 0xc9, 0xf9,
	0x0c, 0x7a, 0xcd, 0x09, 0x2b, 0x46, 0x72, 0xea, 0x5a, 0x43, 0x6b, 0xf2, 0x72, 0xf9, 0xba, 0x79,
	0xf9, 0x97, 0xe4, 0xd4, 0xf9, 0x0a, 0xec, 0x2c, 0x71, 0xed, 0xa1, 0x35, 0x79, 0xf5, 0xfd, 0x7d,
	0x0d, 0x0d, 0x1b, 0x02, 0xb8, 0x60, 0x7a, 0xfa, 0x23, 0x24, 0x9b, 0x2d, 0x5d, 0xda, 0x59, 0xe2,
	0x04, 0xe0, 0x6d, 0xfb, 0xbb, 0x57, 0x82, 0x48, 0xca, 0xb4, 0xdb, 0x29, 0xf3, 0x83, 0x27, 0xf9,
	0x7f, 0x5a, 0x66, 0x2c, 0xad, 0x0a, 0xde, 0xe8, 0x63, 0xba, 0xa0, 0x0c, 0x3a, 0xdf, 0x40, 0x57,
	0x10, 0x7d, 0xe1, 0x76, 0x87, 0x9d, 0x93, 0x05, 0xa5, 0x13, 0xff, 0xd9, 0xfb, 0x0b, 0xf0, 0xe5,
	0xb0, 0xdd, 0x5a, 0x89, 0x4c, 0xc1, 0x98, 0xe7, 0xa8, 0xbd, 0x82, 0x8f, 0xad, 0x33, 0x15, 0xba,
	0x6a, 0xe3, 0x5f, 0xcf, 0x1e, 0x2c, 0x30, 0x8e, 0x79, 0x0e, 0x4f, 0xfe, 0xb8, 0x99, 0xd3, 0x6a,
	0x0e, 0x0a, 0xc0, 0xc0, 0x3a, 0xfb, 0x5d, 0x07, 0x53, 0xbe, 0x21, 0x2c, 0x85, 0x5c, 0xa6, 0x28,
	0xa5, 0xac, 0xc4, 0x47, 0x07, 0xb2, 0x67, 0x2e, 0xd2, 0x4f, 0xa3, 0x6e, 0xec, 0xce, 0xdc, 0xf7,
	0x6f, 0xed, 0xd1, 0xbc, 0xaa, 0xf4, 0x13, 0x05, 0x2b, 0x59, 0xa8, 0xd0, 0x83, 0xcb, 0xc6, 0x79,
	0xdf, 0x78, 0x22, 0x3f, 0x51, 0x91, 0xf1, 0x44, 0xa1, 0x17, 0x19, 0xcf, 0xde, 0x1e, 0x57, 0x03,
	0x8c, 0xfd, 0x44, 0x61, 0x6c, 0x5c, 0x18, 0x87, 0x1e, 0xc6, 0xc6, 0xb7, 0x7e, 0x51, 0xc2, 0x7a,
	0x8f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xba, 0x66, 0x83, 0xb6, 0xf4, 0x02, 0x00, 0x00,
}
