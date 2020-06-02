// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/hash_policy.proto

package envoy_type

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type HashPolicy struct {
	// Types that are valid to be assigned to PolicySpecifier:
	//	*HashPolicy_SourceIp_
	PolicySpecifier      isHashPolicy_PolicySpecifier `protobuf_oneof:"policy_specifier"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *HashPolicy) Reset()         { *m = HashPolicy{} }
func (m *HashPolicy) String() string { return proto.CompactTextString(m) }
func (*HashPolicy) ProtoMessage()    {}
func (*HashPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_b409d49e989a073d, []int{0}
}

func (m *HashPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashPolicy.Unmarshal(m, b)
}
func (m *HashPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashPolicy.Marshal(b, m, deterministic)
}
func (m *HashPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashPolicy.Merge(m, src)
}
func (m *HashPolicy) XXX_Size() int {
	return xxx_messageInfo_HashPolicy.Size(m)
}
func (m *HashPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_HashPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_HashPolicy proto.InternalMessageInfo

type isHashPolicy_PolicySpecifier interface {
	isHashPolicy_PolicySpecifier()
}

type HashPolicy_SourceIp_ struct {
	SourceIp *HashPolicy_SourceIp `protobuf:"bytes,1,opt,name=source_ip,json=sourceIp,proto3,oneof"`
}

func (*HashPolicy_SourceIp_) isHashPolicy_PolicySpecifier() {}

func (m *HashPolicy) GetPolicySpecifier() isHashPolicy_PolicySpecifier {
	if m != nil {
		return m.PolicySpecifier
	}
	return nil
}

func (m *HashPolicy) GetSourceIp() *HashPolicy_SourceIp {
	if x, ok := m.GetPolicySpecifier().(*HashPolicy_SourceIp_); ok {
		return x.SourceIp
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HashPolicy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HashPolicy_SourceIp_)(nil),
	}
}

type HashPolicy_SourceIp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashPolicy_SourceIp) Reset()         { *m = HashPolicy_SourceIp{} }
func (m *HashPolicy_SourceIp) String() string { return proto.CompactTextString(m) }
func (*HashPolicy_SourceIp) ProtoMessage()    {}
func (*HashPolicy_SourceIp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b409d49e989a073d, []int{0, 0}
}

func (m *HashPolicy_SourceIp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashPolicy_SourceIp.Unmarshal(m, b)
}
func (m *HashPolicy_SourceIp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashPolicy_SourceIp.Marshal(b, m, deterministic)
}
func (m *HashPolicy_SourceIp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashPolicy_SourceIp.Merge(m, src)
}
func (m *HashPolicy_SourceIp) XXX_Size() int {
	return xxx_messageInfo_HashPolicy_SourceIp.Size(m)
}
func (m *HashPolicy_SourceIp) XXX_DiscardUnknown() {
	xxx_messageInfo_HashPolicy_SourceIp.DiscardUnknown(m)
}

var xxx_messageInfo_HashPolicy_SourceIp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HashPolicy)(nil), "envoy.type.HashPolicy")
	proto.RegisterType((*HashPolicy_SourceIp)(nil), "envoy.type.HashPolicy.SourceIp")
}

func init() { proto.RegisterFile("envoy/type/hash_policy.proto", fileDescriptor_b409d49e989a073d) }

var fileDescriptor_b409d49e989a073d = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x48, 0x2c, 0xce, 0x88, 0x2f, 0xc8, 0xcf, 0xc9,
	0x4c, 0xae, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02, 0xcb, 0xea, 0x81, 0x64, 0xa5,
	0xc4, 0xcb, 0x12, 0x73, 0x32, 0x53, 0x12, 0x4b, 0x52, 0xf5, 0x61, 0x0c, 0x88, 0x22, 0xa5, 0x42,
	0x2e, 0x2e, 0x8f, 0xc4, 0xe2, 0x8c, 0x00, 0xb0, 0x46, 0x21, 0x3b, 0x2e, 0xce, 0xe2, 0xfc, 0xd2,
	0xa2, 0xe4, 0xd4, 0xf8, 0xcc, 0x02, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x79, 0x3d, 0x84,
	0x31, 0x7a, 0x08, 0xa5, 0x7a, 0xc1, 0x60, 0x75, 0x9e, 0x05, 0x1e, 0x0c, 0x41, 0x1c, 0xc5, 0x50,
	0xb6, 0x14, 0x17, 0x17, 0x07, 0x4c, 0xdc, 0x49, 0x9c, 0x4b, 0x00, 0xe2, 0x9c, 0xf8, 0xe2, 0x82,
	0xd4, 0xe4, 0xcc, 0xb4, 0xcc, 0xd4, 0x22, 0x21, 0xe6, 0x1f, 0x4e, 0x8c, 0x4e, 0xba, 0x5c, 0x12,
	0x99, 0xf9, 0x10, 0x53, 0x0b, 0x8a, 0xf2, 0x2b, 0x2a, 0x91, 0x2c, 0x70, 0xe2, 0x47, 0xd8, 0x10,
	0x00, 0x72, 0x5f, 0x00, 0x63, 0x12, 0x1b, 0xd8, 0xa1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xdf, 0xc6, 0xa6, 0x78, 0xed, 0x00, 0x00, 0x00,
}