// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/dashboard/v1/layouts.proto

package dashboard

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// A basic layout divides the available space into vertical columns of equal
// width and arranges a list of widgets using a row-first strategy.
type GridLayout struct {
	// The number of columns into which the view's width is divided. If omitted
	// or set to zero, a system default will be used while rendering.
	Columns int64 `protobuf:"varint,1,opt,name=columns,proto3" json:"columns,omitempty"`
	// The informational elements that are arranged into the columns row-first.
	Widgets              []*Widget `protobuf:"bytes,2,rep,name=widgets,proto3" json:"widgets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GridLayout) Reset()         { *m = GridLayout{} }
func (m *GridLayout) String() string { return proto.CompactTextString(m) }
func (*GridLayout) ProtoMessage()    {}
func (*GridLayout) Descriptor() ([]byte, []int) {
	return fileDescriptor_02d4908152646d47, []int{0}
}

func (m *GridLayout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GridLayout.Unmarshal(m, b)
}
func (m *GridLayout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GridLayout.Marshal(b, m, deterministic)
}
func (m *GridLayout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GridLayout.Merge(m, src)
}
func (m *GridLayout) XXX_Size() int {
	return xxx_messageInfo_GridLayout.Size(m)
}
func (m *GridLayout) XXX_DiscardUnknown() {
	xxx_messageInfo_GridLayout.DiscardUnknown(m)
}

var xxx_messageInfo_GridLayout proto.InternalMessageInfo

func (m *GridLayout) GetColumns() int64 {
	if m != nil {
		return m.Columns
	}
	return 0
}

func (m *GridLayout) GetWidgets() []*Widget {
	if m != nil {
		return m.Widgets
	}
	return nil
}

// A simplified layout that divides the available space into rows
// and arranges a set of widgets horizontally in each row.
type RowLayout struct {
	// The rows of content to display.
	Rows                 []*RowLayout_Row `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RowLayout) Reset()         { *m = RowLayout{} }
func (m *RowLayout) String() string { return proto.CompactTextString(m) }
func (*RowLayout) ProtoMessage()    {}
func (*RowLayout) Descriptor() ([]byte, []int) {
	return fileDescriptor_02d4908152646d47, []int{1}
}

func (m *RowLayout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RowLayout.Unmarshal(m, b)
}
func (m *RowLayout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RowLayout.Marshal(b, m, deterministic)
}
func (m *RowLayout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RowLayout.Merge(m, src)
}
func (m *RowLayout) XXX_Size() int {
	return xxx_messageInfo_RowLayout.Size(m)
}
func (m *RowLayout) XXX_DiscardUnknown() {
	xxx_messageInfo_RowLayout.DiscardUnknown(m)
}

var xxx_messageInfo_RowLayout proto.InternalMessageInfo

func (m *RowLayout) GetRows() []*RowLayout_Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

// Defines the layout properties and content for a row.
type RowLayout_Row struct {
	// The relative weight of this row. The row weight is used to adjust the
	// height of rows on the screen (relative to peers). Greater the weight,
	// greater the height of the row on the screen. If omitted, a value
	// of 1 is used while rendering.
	Weight int64 `protobuf:"varint,1,opt,name=weight,proto3" json:"weight,omitempty"`
	// The display widgets arranged horizontally in this row.
	Widgets              []*Widget `protobuf:"bytes,2,rep,name=widgets,proto3" json:"widgets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RowLayout_Row) Reset()         { *m = RowLayout_Row{} }
func (m *RowLayout_Row) String() string { return proto.CompactTextString(m) }
func (*RowLayout_Row) ProtoMessage()    {}
func (*RowLayout_Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_02d4908152646d47, []int{1, 0}
}

func (m *RowLayout_Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RowLayout_Row.Unmarshal(m, b)
}
func (m *RowLayout_Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RowLayout_Row.Marshal(b, m, deterministic)
}
func (m *RowLayout_Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RowLayout_Row.Merge(m, src)
}
func (m *RowLayout_Row) XXX_Size() int {
	return xxx_messageInfo_RowLayout_Row.Size(m)
}
func (m *RowLayout_Row) XXX_DiscardUnknown() {
	xxx_messageInfo_RowLayout_Row.DiscardUnknown(m)
}

var xxx_messageInfo_RowLayout_Row proto.InternalMessageInfo

func (m *RowLayout_Row) GetWeight() int64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *RowLayout_Row) GetWidgets() []*Widget {
	if m != nil {
		return m.Widgets
	}
	return nil
}

// A simplified layout that divides the available space into vertical columns
// and arranges a set of widgets vertically in each column.
type ColumnLayout struct {
	// The columns of content to display.
	Columns              []*ColumnLayout_Column `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ColumnLayout) Reset()         { *m = ColumnLayout{} }
func (m *ColumnLayout) String() string { return proto.CompactTextString(m) }
func (*ColumnLayout) ProtoMessage()    {}
func (*ColumnLayout) Descriptor() ([]byte, []int) {
	return fileDescriptor_02d4908152646d47, []int{2}
}

func (m *ColumnLayout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ColumnLayout.Unmarshal(m, b)
}
func (m *ColumnLayout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ColumnLayout.Marshal(b, m, deterministic)
}
func (m *ColumnLayout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColumnLayout.Merge(m, src)
}
func (m *ColumnLayout) XXX_Size() int {
	return xxx_messageInfo_ColumnLayout.Size(m)
}
func (m *ColumnLayout) XXX_DiscardUnknown() {
	xxx_messageInfo_ColumnLayout.DiscardUnknown(m)
}

var xxx_messageInfo_ColumnLayout proto.InternalMessageInfo

func (m *ColumnLayout) GetColumns() []*ColumnLayout_Column {
	if m != nil {
		return m.Columns
	}
	return nil
}

// Defines the layout properties and content for a column.
type ColumnLayout_Column struct {
	// The relative weight of this column. The column weight is used to adjust
	// the width of columns on the screen (relative to peers).
	// Greater the weight, greater the width of the column on the screen.
	// If omitted, a value of 1 is used while rendering.
	Weight int64 `protobuf:"varint,1,opt,name=weight,proto3" json:"weight,omitempty"`
	// The display widgets arranged vertically in this column.
	Widgets              []*Widget `protobuf:"bytes,2,rep,name=widgets,proto3" json:"widgets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ColumnLayout_Column) Reset()         { *m = ColumnLayout_Column{} }
func (m *ColumnLayout_Column) String() string { return proto.CompactTextString(m) }
func (*ColumnLayout_Column) ProtoMessage()    {}
func (*ColumnLayout_Column) Descriptor() ([]byte, []int) {
	return fileDescriptor_02d4908152646d47, []int{2, 0}
}

func (m *ColumnLayout_Column) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ColumnLayout_Column.Unmarshal(m, b)
}
func (m *ColumnLayout_Column) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ColumnLayout_Column.Marshal(b, m, deterministic)
}
func (m *ColumnLayout_Column) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColumnLayout_Column.Merge(m, src)
}
func (m *ColumnLayout_Column) XXX_Size() int {
	return xxx_messageInfo_ColumnLayout_Column.Size(m)
}
func (m *ColumnLayout_Column) XXX_DiscardUnknown() {
	xxx_messageInfo_ColumnLayout_Column.DiscardUnknown(m)
}

var xxx_messageInfo_ColumnLayout_Column proto.InternalMessageInfo

func (m *ColumnLayout_Column) GetWeight() int64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *ColumnLayout_Column) GetWidgets() []*Widget {
	if m != nil {
		return m.Widgets
	}
	return nil
}

func init() {
	proto.RegisterType((*GridLayout)(nil), "google.monitoring.dashboard.v1.GridLayout")
	proto.RegisterType((*RowLayout)(nil), "google.monitoring.dashboard.v1.RowLayout")
	proto.RegisterType((*RowLayout_Row)(nil), "google.monitoring.dashboard.v1.RowLayout.Row")
	proto.RegisterType((*ColumnLayout)(nil), "google.monitoring.dashboard.v1.ColumnLayout")
	proto.RegisterType((*ColumnLayout_Column)(nil), "google.monitoring.dashboard.v1.ColumnLayout.Column")
}

func init() {
	proto.RegisterFile("google/monitoring/dashboard/v1/layouts.proto", fileDescriptor_02d4908152646d47)
}

var fileDescriptor_02d4908152646d47 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc6, 0x89, 0x91, 0x14, 0x5f, 0x3b, 0x65, 0x90, 0x90, 0x41, 0x4a, 0x06, 0x29, 0xa8, 0x77,
	0xd4, 0x8e, 0x2e, 0x5a, 0x87, 0x2e, 0x0a, 0x25, 0x8b, 0xe0, 0x22, 0x49, 0x13, 0x2e, 0x07, 0x49,
	0xde, 0x72, 0x97, 0xf6, 0x70, 0xf0, 0x33, 0xf9, 0x19, 0xfc, 0x66, 0x72, 0x7f, 0x1a, 0xbb, 0x68,
	0x16, 0xdd, 0xf2, 0xe4, 0x9e, 0xdf, 0xf3, 0x3e, 0x77, 0xbc, 0x70, 0xc5, 0x10, 0x59, 0x5d, 0xd2,
	0x06, 0x5b, 0xde, 0xa1, 0xe0, 0x2d, 0xa3, 0x45, 0x26, 0xab, 0x1c, 0x33, 0x51, 0xd0, 0xdd, 0x8c,
	0xd6, 0xd9, 0x1b, 0x6e, 0x3b, 0x49, 0x36, 0x02, 0x3b, 0x0c, 0xcf, 0xad, 0x9b, 0x7c, 0xbb, 0x49,
	0xef, 0x26, 0xbb, 0x59, 0x7c, 0x39, 0x90, 0xa6, 0x78, 0xc1, 0xca, 0xce, 0x86, 0x25, 0x15, 0xc0,
	0x52, 0xf0, 0xe2, 0xd1, 0x4c, 0x08, 0x23, 0x18, 0xad, 0xb1, 0xde, 0x36, 0xad, 0x8c, 0xbc, 0x89,
	0x37, 0xf5, 0xd3, 0xbd, 0x0c, 0xef, 0x60, 0x64, 0x39, 0x19, 0x1d, 0x4d, 0xfc, 0xe9, 0xe9, 0xcd,
	0x05, 0xf9, 0xbd, 0x06, 0x79, 0x36, 0xf6, 0x74, 0x8f, 0x25, 0x1f, 0x1e, 0x9c, 0xa4, 0xa8, 0xdc,
	0xa4, 0x7b, 0x38, 0x16, 0xa8, 0xf4, 0x18, 0x1d, 0x76, 0x3d, 0x14, 0xd6, 0x83, 0xfa, 0x2b, 0x35,
	0x68, 0xfc, 0x0a, 0x7e, 0x8a, 0x2a, 0x3c, 0x83, 0x40, 0x95, 0x9c, 0x55, 0x9d, 0xab, 0xec, 0xd4,
	0x1f, 0x34, 0xfe, 0xf4, 0x60, 0xfc, 0x60, 0xee, 0xef, 0x4a, 0x3f, 0x1d, 0x3e, 0x8f, 0x8e, 0x9c,
	0x0f, 0x45, 0x1e, 0xe2, 0x4e, 0xf4, 0x6f, 0x1a, 0xe7, 0x10, 0xd8, 0x5f, 0xff, 0x77, 0x87, 0xc5,
	0x3b, 0x24, 0x6b, 0x6c, 0x06, 0xa8, 0xc5, 0xd8, 0x36, 0x94, 0x2b, 0xbd, 0x13, 0x2b, 0xef, 0x65,
	0xe9, 0xfc, 0x0c, 0xeb, 0xac, 0x65, 0x04, 0x05, 0xa3, 0xac, 0x6c, 0xcd, 0xc6, 0x50, 0x7b, 0x94,
	0x6d, 0xb8, 0xfc, 0x69, 0xc3, 0x6e, 0x7b, 0x91, 0x07, 0x86, 0x99, 0x7f, 0x05, 0x00, 0x00, 0xff,
	0xff, 0xaf, 0xb5, 0x7a, 0xef, 0xe2, 0x02, 0x00, 0x00,
}
