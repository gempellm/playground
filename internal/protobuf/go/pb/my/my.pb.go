// Code generated by protoc-gen-go. DO NOT EDIT.
// source: my/my.proto

package gempellm_proto_demo

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

type ParamValue struct {
	// Types that are valid to be assigned to ValueOneof:
	//	*ParamValue_Double
	//	*ParamValue_Int
	//	*ParamValue_Bool
	//	*ParamValue_String_
	ValueOneof           isParamValue_ValueOneof `protobuf_oneof:"value_oneof"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ParamValue) Reset()         { *m = ParamValue{} }
func (m *ParamValue) String() string { return proto.CompactTextString(m) }
func (*ParamValue) ProtoMessage()    {}
func (*ParamValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_c42ba998dd4d6253, []int{0}
}

func (m *ParamValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParamValue.Unmarshal(m, b)
}
func (m *ParamValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParamValue.Marshal(b, m, deterministic)
}
func (m *ParamValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamValue.Merge(m, src)
}
func (m *ParamValue) XXX_Size() int {
	return xxx_messageInfo_ParamValue.Size(m)
}
func (m *ParamValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamValue.DiscardUnknown(m)
}

var xxx_messageInfo_ParamValue proto.InternalMessageInfo

type isParamValue_ValueOneof interface {
	isParamValue_ValueOneof()
}

type ParamValue_Double struct {
	Double float64 `protobuf:"fixed64,1,opt,name=double,proto3,oneof"`
}

type ParamValue_Int struct {
	Int int64 `protobuf:"varint,2,opt,name=int,proto3,oneof"`
}

type ParamValue_Bool struct {
	Bool bool `protobuf:"varint,3,opt,name=bool,proto3,oneof"`
}

type ParamValue_String_ struct {
	String_ string `protobuf:"bytes,4,opt,name=string,proto3,oneof"`
}

func (*ParamValue_Double) isParamValue_ValueOneof() {}

func (*ParamValue_Int) isParamValue_ValueOneof() {}

func (*ParamValue_Bool) isParamValue_ValueOneof() {}

func (*ParamValue_String_) isParamValue_ValueOneof() {}

func (m *ParamValue) GetValueOneof() isParamValue_ValueOneof {
	if m != nil {
		return m.ValueOneof
	}
	return nil
}

func (m *ParamValue) GetDouble() float64 {
	if x, ok := m.GetValueOneof().(*ParamValue_Double); ok {
		return x.Double
	}
	return 0
}

func (m *ParamValue) GetInt() int64 {
	if x, ok := m.GetValueOneof().(*ParamValue_Int); ok {
		return x.Int
	}
	return 0
}

func (m *ParamValue) GetBool() bool {
	if x, ok := m.GetValueOneof().(*ParamValue_Bool); ok {
		return x.Bool
	}
	return false
}

func (m *ParamValue) GetString_() string {
	if x, ok := m.GetValueOneof().(*ParamValue_String_); ok {
		return x.String_
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ParamValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ParamValue_Double)(nil),
		(*ParamValue_Int)(nil),
		(*ParamValue_Bool)(nil),
		(*ParamValue_String_)(nil),
	}
}

type SellParams struct {
	Result               []*SellParams_Item `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SellParams) Reset()         { *m = SellParams{} }
func (m *SellParams) String() string { return proto.CompactTextString(m) }
func (*SellParams) ProtoMessage()    {}
func (*SellParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c42ba998dd4d6253, []int{1}
}

func (m *SellParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellParams.Unmarshal(m, b)
}
func (m *SellParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellParams.Marshal(b, m, deterministic)
}
func (m *SellParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellParams.Merge(m, src)
}
func (m *SellParams) XXX_Size() int {
	return xxx_messageInfo_SellParams.Size(m)
}
func (m *SellParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SellParams.DiscardUnknown(m)
}

var xxx_messageInfo_SellParams proto.InternalMessageInfo

func (m *SellParams) GetResult() []*SellParams_Item {
	if m != nil {
		return m.Result
	}
	return nil
}

type SellParams_Item struct {
	SellerId             int64                  `protobuf:"varint,1,opt,name=seller_id,json=sellerId,proto3" json:"seller_id,omitempty"`
	Rating               float64                `protobuf:"fixed64,2,opt,name=rating,proto3" json:"rating,omitempty"`
	Params               map[string]*ParamValue `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SellParams_Item) Reset()         { *m = SellParams_Item{} }
func (m *SellParams_Item) String() string { return proto.CompactTextString(m) }
func (*SellParams_Item) ProtoMessage()    {}
func (*SellParams_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_c42ba998dd4d6253, []int{1, 0}
}

func (m *SellParams_Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellParams_Item.Unmarshal(m, b)
}
func (m *SellParams_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellParams_Item.Marshal(b, m, deterministic)
}
func (m *SellParams_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellParams_Item.Merge(m, src)
}
func (m *SellParams_Item) XXX_Size() int {
	return xxx_messageInfo_SellParams_Item.Size(m)
}
func (m *SellParams_Item) XXX_DiscardUnknown() {
	xxx_messageInfo_SellParams_Item.DiscardUnknown(m)
}

var xxx_messageInfo_SellParams_Item proto.InternalMessageInfo

func (m *SellParams_Item) GetSellerId() int64 {
	if m != nil {
		return m.SellerId
	}
	return 0
}

func (m *SellParams_Item) GetRating() float64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *SellParams_Item) GetParams() map[string]*ParamValue {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*ParamValue)(nil), "gempellm.proto.demo.ParamValue")
	proto.RegisterType((*SellParams)(nil), "gempellm.proto.demo.SellParams")
	proto.RegisterType((*SellParams_Item)(nil), "gempellm.proto.demo.SellParams.Item")
	proto.RegisterMapType((map[string]*ParamValue)(nil), "gempellm.proto.demo.SellParams.Item.ParamsEntry")
}

func init() {
	proto.RegisterFile("my/my.proto", fileDescriptor_c42ba998dd4d6253)
}

var fileDescriptor_c42ba998dd4d6253 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x37, 0x9b, 0xb5, 0xec, 0x4e, 0x11, 0x24, 0x8a, 0x84, 0xf5, 0x60, 0x59, 0x3c, 0xf4,
	0x14, 0x65, 0x45, 0x10, 0xf1, 0x24, 0x08, 0xdd, 0x9b, 0x44, 0xf0, 0xe0, 0x65, 0x69, 0xe9, 0xb8,
	0x14, 0x93, 0xa6, 0xa4, 0xa9, 0xd0, 0x97, 0xf2, 0x9d, 0x7c, 0x13, 0x49, 0x5a, 0x58, 0x0f, 0x7b,
	0xf0, 0x36, 0xdf, 0x30, 0xf3, 0xff, 0xff, 0x24, 0x10, 0xeb, 0xfe, 0x5a, 0xf7, 0xa2, 0xb1, 0xc6,
	0x19, 0x76, 0xba, 0x43, 0xdd, 0xa0, 0x52, 0x7a, 0x60, 0x51, 0xa2, 0x36, 0xab, 0x1e, 0xe0, 0x25,
	0xb7, 0xb9, 0x7e, 0xcb, 0x55, 0x87, 0x8c, 0x43, 0x54, 0x9a, 0xae, 0x50, 0xc8, 0x49, 0x42, 0x52,
	0x92, 0x4d, 0xe4, 0xc8, 0x8c, 0x01, 0xad, 0x6a, 0xc7, 0xa7, 0x09, 0x49, 0x69, 0x36, 0x91, 0x1e,
	0xd8, 0x19, 0xcc, 0x0a, 0x63, 0x14, 0xa7, 0x09, 0x49, 0xe7, 0xd9, 0x44, 0x06, 0xf2, 0x1a, 0xad,
	0xb3, 0x55, 0xbd, 0xe3, 0xb3, 0x84, 0xa4, 0x0b, 0xaf, 0x31, 0xf0, 0xd3, 0x31, 0xc4, 0x5f, 0xde,
	0x66, 0x6b, 0x6a, 0x34, 0x1f, 0xab, 0xef, 0x29, 0xc0, 0x2b, 0x2a, 0x15, 0xfc, 0x5b, 0xf6, 0x08,
	0x91, 0xc5, 0xb6, 0x53, 0x8e, 0x93, 0x84, 0xa6, 0xf1, 0xfa, 0x4a, 0x1c, 0xc8, 0x2b, 0xf6, 0x0b,
	0x62, 0xe3, 0x50, 0xcb, 0x71, 0x67, 0xf9, 0x43, 0x60, 0xe6, 0x1b, 0xec, 0x02, 0x16, 0x2d, 0x2a,
	0x85, 0x76, 0x5b, 0x95, 0xe1, 0x0a, 0x2a, 0xe7, 0x43, 0x63, 0x53, 0xb2, 0x73, 0x88, 0x6c, 0xee,
	0x7c, 0x36, 0x7f, 0x08, 0x91, 0x23, 0xb1, 0x0c, 0xa2, 0x26, 0x88, 0x72, 0x1a, 0xbc, 0x6f, 0xfe,
	0xe3, 0x2d, 0x86, 0xfa, 0xb9, 0x76, 0xb6, 0x97, 0xe3, 0xfe, 0xf2, 0x1d, 0xe2, 0x3f, 0x6d, 0x76,
	0x02, 0xf4, 0x13, 0xfb, 0x90, 0x63, 0x21, 0x7d, 0xc9, 0xee, 0xe0, 0x28, 0x3c, 0x42, 0x48, 0x10,
	0xaf, 0x2f, 0x0f, 0x3a, 0xed, 0xbf, 0x44, 0x0e, 0xd3, 0x0f, 0xd3, 0x7b, 0x52, 0x44, 0x61, 0xe2,
	0xf6, 0x37, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xff, 0x46, 0x7f, 0xd6, 0x01, 0x00, 0x00,
}
