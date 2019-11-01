// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/rds/kubernetes/proto/config.proto

package proto

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

type Pods struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pods) Reset()         { *m = Pods{} }
func (m *Pods) String() string { return proto.CompactTextString(m) }
func (*Pods) ProtoMessage()    {}
func (*Pods) Descriptor() ([]byte, []int) {
	return fileDescriptor_b730109f5f77edc1, []int{0}
}

func (m *Pods) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pods.Unmarshal(m, b)
}
func (m *Pods) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pods.Marshal(b, m, deterministic)
}
func (m *Pods) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pods.Merge(m, src)
}
func (m *Pods) XXX_Size() int {
	return xxx_messageInfo_Pods.Size(m)
}
func (m *Pods) XXX_DiscardUnknown() {
	xxx_messageInfo_Pods.DiscardUnknown(m)
}

var xxx_messageInfo_Pods proto.InternalMessageInfo

type Endpoints struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoints) Reset()         { *m = Endpoints{} }
func (m *Endpoints) String() string { return proto.CompactTextString(m) }
func (*Endpoints) ProtoMessage()    {}
func (*Endpoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_b730109f5f77edc1, []int{1}
}

func (m *Endpoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoints.Unmarshal(m, b)
}
func (m *Endpoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoints.Marshal(b, m, deterministic)
}
func (m *Endpoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoints.Merge(m, src)
}
func (m *Endpoints) XXX_Size() int {
	return xxx_messageInfo_Endpoints.Size(m)
}
func (m *Endpoints) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoints.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoints proto.InternalMessageInfo

type Services struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Services) Reset()         { *m = Services{} }
func (m *Services) String() string { return proto.CompactTextString(m) }
func (*Services) ProtoMessage()    {}
func (*Services) Descriptor() ([]byte, []int) {
	return fileDescriptor_b730109f5f77edc1, []int{2}
}

func (m *Services) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Services.Unmarshal(m, b)
}
func (m *Services) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Services.Marshal(b, m, deterministic)
}
func (m *Services) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Services.Merge(m, src)
}
func (m *Services) XXX_Size() int {
	return xxx_messageInfo_Services.Size(m)
}
func (m *Services) XXX_DiscardUnknown() {
	xxx_messageInfo_Services.DiscardUnknown(m)
}

var xxx_messageInfo_Services proto.InternalMessageInfo

// Kubernetes provider config.
type ProviderConfig struct {
	// Namespace to list resources for. If not specified, we default to all
	// namespaces.
	Namespace *string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	// Pods discovery options. This field should be declared for the pods
	// discovery to be enabled.
	Pods *Pods `protobuf:"bytes,2,opt,name=pods" json:"pods,omitempty"`
	// Endpoints discovery options. This field should be declared for the
	// endpoints discovery to be enabled.
	Endpoints *Endpoints `protobuf:"bytes,3,opt,name=endpoints" json:"endpoints,omitempty"`
	// Services discovery options. This field should be declared for the
	// services discovery to be enabled.
	Services *Services `protobuf:"bytes,4,opt,name=services" json:"services,omitempty"`
	// How often resources should be evaluated/expanded.
	ReEvalSec            *int32   `protobuf:"varint,99,opt,name=re_eval_sec,json=reEvalSec,def=60" json:"re_eval_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProviderConfig) Reset()         { *m = ProviderConfig{} }
func (m *ProviderConfig) String() string { return proto.CompactTextString(m) }
func (*ProviderConfig) ProtoMessage()    {}
func (*ProviderConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_b730109f5f77edc1, []int{3}
}

func (m *ProviderConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProviderConfig.Unmarshal(m, b)
}
func (m *ProviderConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProviderConfig.Marshal(b, m, deterministic)
}
func (m *ProviderConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderConfig.Merge(m, src)
}
func (m *ProviderConfig) XXX_Size() int {
	return xxx_messageInfo_ProviderConfig.Size(m)
}
func (m *ProviderConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderConfig proto.InternalMessageInfo

const Default_ProviderConfig_ReEvalSec int32 = 60

func (m *ProviderConfig) GetNamespace() string {
	if m != nil && m.Namespace != nil {
		return *m.Namespace
	}
	return ""
}

func (m *ProviderConfig) GetPods() *Pods {
	if m != nil {
		return m.Pods
	}
	return nil
}

func (m *ProviderConfig) GetEndpoints() *Endpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ProviderConfig) GetServices() *Services {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *ProviderConfig) GetReEvalSec() int32 {
	if m != nil && m.ReEvalSec != nil {
		return *m.ReEvalSec
	}
	return Default_ProviderConfig_ReEvalSec
}

func init() {
	proto.RegisterType((*Pods)(nil), "cloudprober.rds.kubernetes.Pods")
	proto.RegisterType((*Endpoints)(nil), "cloudprober.rds.kubernetes.Endpoints")
	proto.RegisterType((*Services)(nil), "cloudprober.rds.kubernetes.Services")
	proto.RegisterType((*ProviderConfig)(nil), "cloudprober.rds.kubernetes.ProviderConfig")
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/rds/kubernetes/proto/config.proto", fileDescriptor_b730109f5f77edc1)
}

var fileDescriptor_b730109f5f77edc1 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x8c, 0xd2, 0x4c, 0xc0, 0xc3, 0x9e, 0x82, 0x78, 0x08, 0x41, 0x21, 0xa7, 0x5d,
	0x11, 0xf1, 0xe0, 0x45, 0xa1, 0xf4, 0x5e, 0xd2, 0x07, 0x28, 0xc9, 0xee, 0x18, 0x83, 0x69, 0x26,
	0xcc, 0x24, 0x79, 0x0d, 0x5f, 0x59, 0x5c, 0x6c, 0xe3, 0xc5, 0x1e, 0x77, 0xf9, 0xbf, 0xff, 0xfb,
	0x07, 0x5e, 0x9b, 0x76, 0xfc, 0x98, 0x6a, 0x6d, 0xe9, 0x60, 0x1a, 0xa2, 0xa6, 0x43, 0x63, 0x3b,
	0x9a, 0xdc, 0xc0, 0x54, 0x23, 0x1b, 0x76, 0x62, 0x3e, 0xa7, 0x1a, 0xb9, 0xc7, 0x11, 0xc5, 0x0c,
	0x4c, 0x23, 0x19, 0x4b, 0xfd, 0x7b, 0xdb, 0x68, 0xff, 0x50, 0x37, 0x7f, 0xe2, 0x9a, 0x9d, 0xe8,
	0x25, 0x9e, 0x5f, 0x41, 0xb4, 0x25, 0x27, 0x79, 0x02, 0xf1, 0xa6, 0x77, 0x03, 0xb5, 0xfd, 0x28,
	0x39, 0xc0, 0x6a, 0x87, 0x3c, 0xb7, 0x16, 0x25, 0xff, 0x0a, 0xe1, 0x7a, 0xcb, 0x34, 0xb7, 0x0e,
	0x79, 0xed, 0x5b, 0xd5, 0x2d, 0xc4, 0x7d, 0x75, 0x40, 0x19, 0x2a, 0x8b, 0x69, 0x90, 0x05, 0x45,
	0x5c, 0x2e, 0x1f, 0xea, 0x09, 0xa2, 0x81, 0x9c, 0xa4, 0x61, 0x16, 0x14, 0xc9, 0x63, 0xa6, 0xff,
	0x97, 0xeb, 0x1f, 0x73, 0xe9, 0xd3, 0x6a, 0x0d, 0x31, 0x1e, 0xfd, 0xe9, 0x85, 0x47, 0xef, 0xcf,
	0xa1, 0xa7, 0xb1, 0xe5, 0xc2, 0xa9, 0x37, 0x58, 0xc9, 0xef, 0xee, 0x34, 0xf2, 0x1d, 0x77, 0xe7,
	0x3a, 0x8e, 0x37, 0x96, 0x27, 0x4a, 0xe5, 0x90, 0x30, 0xee, 0x71, 0xae, 0xba, 0xbd, 0xa0, 0x4d,
	0x6d, 0x16, 0x14, 0x97, 0x2f, 0xe1, 0xf3, 0x43, 0x19, 0x33, 0x6e, 0xe6, 0xaa, 0xdb, 0xa1, 0xfd,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xc2, 0x18, 0xbe, 0x90, 0x01, 0x00, 0x00,
}
