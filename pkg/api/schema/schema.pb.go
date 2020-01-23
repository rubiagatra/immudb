// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.proto

package schema

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Key struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0}
}

func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type KeyValue struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}
func (*KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{1}
}

func (m *KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValue.Unmarshal(m, b)
}
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValue.Marshal(b, m, deterministic)
}
func (m *KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValue.Merge(m, src)
}
func (m *KeyValue) XXX_Size() int {
	return xxx_messageInfo_KeyValue.Size(m)
}
func (m *KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValue proto.InternalMessageInfo

func (m *KeyValue) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Index struct {
	Index                uint64   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Index) Reset()         { *m = Index{} }
func (m *Index) String() string { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()    {}
func (*Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{2}
}

func (m *Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index.Unmarshal(m, b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index.Marshal(b, m, deterministic)
}
func (m *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(m, src)
}
func (m *Index) XXX_Size() int {
	return xxx_messageInfo_Index.Size(m)
}
func (m *Index) XXX_DiscardUnknown() {
	xxx_messageInfo_Index.DiscardUnknown(m)
}

var xxx_messageInfo_Index proto.InternalMessageInfo

func (m *Index) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type Item struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Index                uint64   `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{3}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Item) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Item) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type KVList struct {
	KVs                  []*KeyValue `protobuf:"bytes,1,rep,name=KVs,proto3" json:"KVs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *KVList) Reset()         { *m = KVList{} }
func (m *KVList) String() string { return proto.CompactTextString(m) }
func (*KVList) ProtoMessage()    {}
func (*KVList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{4}
}

func (m *KVList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVList.Unmarshal(m, b)
}
func (m *KVList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVList.Marshal(b, m, deterministic)
}
func (m *KVList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVList.Merge(m, src)
}
func (m *KVList) XXX_Size() int {
	return xxx_messageInfo_KVList.Size(m)
}
func (m *KVList) XXX_DiscardUnknown() {
	xxx_messageInfo_KVList.DiscardUnknown(m)
}

var xxx_messageInfo_KVList proto.InternalMessageInfo

func (m *KVList) GetKVs() []*KeyValue {
	if m != nil {
		return m.KVs
	}
	return nil
}

type KeyList struct {
	Keys                 []*Key   `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyList) Reset()         { *m = KeyList{} }
func (m *KeyList) String() string { return proto.CompactTextString(m) }
func (*KeyList) ProtoMessage()    {}
func (*KeyList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{5}
}

func (m *KeyList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyList.Unmarshal(m, b)
}
func (m *KeyList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyList.Marshal(b, m, deterministic)
}
func (m *KeyList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyList.Merge(m, src)
}
func (m *KeyList) XXX_Size() int {
	return xxx_messageInfo_KeyList.Size(m)
}
func (m *KeyList) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyList.DiscardUnknown(m)
}

var xxx_messageInfo_KeyList proto.InternalMessageInfo

func (m *KeyList) GetKeys() []*Key {
	if m != nil {
		return m.Keys
	}
	return nil
}

type ItemList struct {
	Items                []*Item  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemList) Reset()         { *m = ItemList{} }
func (m *ItemList) String() string { return proto.CompactTextString(m) }
func (*ItemList) ProtoMessage()    {}
func (*ItemList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{6}
}

func (m *ItemList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemList.Unmarshal(m, b)
}
func (m *ItemList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemList.Marshal(b, m, deterministic)
}
func (m *ItemList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemList.Merge(m, src)
}
func (m *ItemList) XXX_Size() int {
	return xxx_messageInfo_ItemList.Size(m)
}
func (m *ItemList) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemList.DiscardUnknown(m)
}

var xxx_messageInfo_ItemList proto.InternalMessageInfo

func (m *ItemList) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type ScanOptions struct {
	Prefix               []byte   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Offset               []byte   `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint64   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Reverse              bool     `protobuf:"varint,4,opt,name=reverse,proto3" json:"reverse,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanOptions) Reset()         { *m = ScanOptions{} }
func (m *ScanOptions) String() string { return proto.CompactTextString(m) }
func (*ScanOptions) ProtoMessage()    {}
func (*ScanOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{7}
}

func (m *ScanOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanOptions.Unmarshal(m, b)
}
func (m *ScanOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanOptions.Marshal(b, m, deterministic)
}
func (m *ScanOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanOptions.Merge(m, src)
}
func (m *ScanOptions) XXX_Size() int {
	return xxx_messageInfo_ScanOptions.Size(m)
}
func (m *ScanOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ScanOptions proto.InternalMessageInfo

func (m *ScanOptions) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *ScanOptions) GetOffset() []byte {
	if m != nil {
		return m.Offset
	}
	return nil
}

func (m *ScanOptions) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ScanOptions) GetReverse() bool {
	if m != nil {
		return m.Reverse
	}
	return false
}

type InclusionProof struct {
	At                   uint64   `protobuf:"varint,1,opt,name=at,proto3" json:"at,omitempty"`
	Index                uint64   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Root                 []byte   `protobuf:"bytes,3,opt,name=root,proto3" json:"root,omitempty"`
	Leaf                 []byte   `protobuf:"bytes,4,opt,name=leaf,proto3" json:"leaf,omitempty"`
	Path                 [][]byte `protobuf:"bytes,5,rep,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InclusionProof) Reset()         { *m = InclusionProof{} }
func (m *InclusionProof) String() string { return proto.CompactTextString(m) }
func (*InclusionProof) ProtoMessage()    {}
func (*InclusionProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{8}
}

func (m *InclusionProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InclusionProof.Unmarshal(m, b)
}
func (m *InclusionProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InclusionProof.Marshal(b, m, deterministic)
}
func (m *InclusionProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InclusionProof.Merge(m, src)
}
func (m *InclusionProof) XXX_Size() int {
	return xxx_messageInfo_InclusionProof.Size(m)
}
func (m *InclusionProof) XXX_DiscardUnknown() {
	xxx_messageInfo_InclusionProof.DiscardUnknown(m)
}

var xxx_messageInfo_InclusionProof proto.InternalMessageInfo

func (m *InclusionProof) GetAt() uint64 {
	if m != nil {
		return m.At
	}
	return 0
}

func (m *InclusionProof) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *InclusionProof) GetRoot() []byte {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *InclusionProof) GetLeaf() []byte {
	if m != nil {
		return m.Leaf
	}
	return nil
}

func (m *InclusionProof) GetPath() [][]byte {
	if m != nil {
		return m.Path
	}
	return nil
}

type ConsistencyProof struct {
	First                uint64   `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
	Second               uint64   `protobuf:"varint,2,opt,name=second,proto3" json:"second,omitempty"`
	FirstRoot            []byte   `protobuf:"bytes,3,opt,name=firstRoot,proto3" json:"firstRoot,omitempty"`
	SecondRoot           []byte   `protobuf:"bytes,4,opt,name=secondRoot,proto3" json:"secondRoot,omitempty"`
	Path                 [][]byte `protobuf:"bytes,5,rep,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsistencyProof) Reset()         { *m = ConsistencyProof{} }
func (m *ConsistencyProof) String() string { return proto.CompactTextString(m) }
func (*ConsistencyProof) ProtoMessage()    {}
func (*ConsistencyProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{9}
}

func (m *ConsistencyProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsistencyProof.Unmarshal(m, b)
}
func (m *ConsistencyProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsistencyProof.Marshal(b, m, deterministic)
}
func (m *ConsistencyProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsistencyProof.Merge(m, src)
}
func (m *ConsistencyProof) XXX_Size() int {
	return xxx_messageInfo_ConsistencyProof.Size(m)
}
func (m *ConsistencyProof) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsistencyProof.DiscardUnknown(m)
}

var xxx_messageInfo_ConsistencyProof proto.InternalMessageInfo

func (m *ConsistencyProof) GetFirst() uint64 {
	if m != nil {
		return m.First
	}
	return 0
}

func (m *ConsistencyProof) GetSecond() uint64 {
	if m != nil {
		return m.Second
	}
	return 0
}

func (m *ConsistencyProof) GetFirstRoot() []byte {
	if m != nil {
		return m.FirstRoot
	}
	return nil
}

func (m *ConsistencyProof) GetSecondRoot() []byte {
	if m != nil {
		return m.SecondRoot
	}
	return nil
}

func (m *ConsistencyProof) GetPath() [][]byte {
	if m != nil {
		return m.Path
	}
	return nil
}

type HealthResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{10}
}

func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResponse.Unmarshal(m, b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
}
func (m *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(m, src)
}
func (m *HealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthResponse.Size(m)
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*Key)(nil), "immudb.schema.Key")
	proto.RegisterType((*KeyValue)(nil), "immudb.schema.KeyValue")
	proto.RegisterType((*Index)(nil), "immudb.schema.Index")
	proto.RegisterType((*Item)(nil), "immudb.schema.Item")
	proto.RegisterType((*KVList)(nil), "immudb.schema.KVList")
	proto.RegisterType((*KeyList)(nil), "immudb.schema.KeyList")
	proto.RegisterType((*ItemList)(nil), "immudb.schema.ItemList")
	proto.RegisterType((*ScanOptions)(nil), "immudb.schema.ScanOptions")
	proto.RegisterType((*InclusionProof)(nil), "immudb.schema.InclusionProof")
	proto.RegisterType((*ConsistencyProof)(nil), "immudb.schema.ConsistencyProof")
	proto.RegisterType((*HealthResponse)(nil), "immudb.schema.HealthResponse")
}

func init() { proto.RegisterFile("schema.proto", fileDescriptor_1c5fb4d8cc22d66a) }

var fileDescriptor_1c5fb4d8cc22d66a = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x56, 0xe2, 0x7c, 0x75, 0x92, 0xb7, 0xaa, 0xf6, 0x2d, 0x6d, 0x14, 0x28, 0x44, 0x3e, 0xa0,
	0x20, 0x84, 0x2d, 0xda, 0x02, 0x07, 0x0e, 0xa0, 0x52, 0xd4, 0x46, 0x41, 0x02, 0x39, 0x52, 0x0f,
	0xdc, 0x36, 0xce, 0x38, 0x59, 0x35, 0xf6, 0x5a, 0xde, 0x75, 0x55, 0xff, 0x0d, 0x7e, 0x01, 0x3f,
	0x15, 0xed, 0xae, 0x4d, 0x3e, 0xec, 0x48, 0xdc, 0x66, 0x9e, 0x99, 0x67, 0xbe, 0xf6, 0xd1, 0x42,
	0x4f, 0xf8, 0x4b, 0x0c, 0xa9, 0x13, 0x27, 0x5c, 0x72, 0xf2, 0x1f, 0x0b, 0xc3, 0x74, 0x3e, 0x73,
	0x0c, 0x38, 0x78, 0xba, 0xe0, 0x7c, 0xb1, 0x42, 0x57, 0x07, 0x67, 0x69, 0xe0, 0x62, 0x18, 0xcb,
	0xcc, 0xe4, 0xda, 0xa7, 0x60, 0x4d, 0x30, 0x23, 0x47, 0x60, 0xdd, 0x63, 0xd6, 0xaf, 0x0d, 0x6b,
	0xa3, 0x9e, 0xa7, 0x4c, 0xfb, 0x1c, 0x3a, 0x13, 0xcc, 0xee, 0xe8, 0x2a, 0xc5, 0x72, 0x94, 0x1c,
	0x43, 0xf3, 0x41, 0x85, 0xfa, 0x75, 0x8d, 0x19, 0xc7, 0x3e, 0x83, 0xe6, 0x38, 0x9a, 0xe3, 0xa3,
	0x0a, 0x33, 0x65, 0x68, 0x4a, 0xc3, 0x33, 0x8e, 0x7d, 0x0d, 0x8d, 0xb1, 0xc4, 0xf0, 0x5f, 0xcb,
	0xad, 0xab, 0x58, 0x9b, 0x55, 0x2e, 0xa0, 0x35, 0xb9, 0xfb, 0xc6, 0x84, 0x24, 0xaf, 0xc0, 0x9a,
	0xdc, 0x89, 0x7e, 0x6d, 0x68, 0x8d, 0xba, 0xe7, 0xa7, 0xce, 0xd6, 0xd6, 0x4e, 0x31, 0xbc, 0xa7,
	0x72, 0xec, 0xb7, 0xd0, 0x9e, 0x60, 0xa6, 0x59, 0x2f, 0xa1, 0x71, 0x8f, 0x59, 0x41, 0x23, 0x65,
	0x9a, 0xa7, 0xe3, 0xf6, 0x3b, 0xe8, 0xa8, 0x69, 0xf3, 0x4e, 0x4d, 0x26, 0x31, 0x2c, 0x48, 0xff,
	0xef, 0x90, 0x54, 0x9e, 0x67, 0x32, 0xec, 0x10, 0xba, 0x53, 0x9f, 0x46, 0xdf, 0x63, 0xc9, 0x78,
	0x24, 0xc8, 0x09, 0xb4, 0xe2, 0x04, 0x03, 0xf6, 0x98, 0xaf, 0x9b, 0x7b, 0x0a, 0xe7, 0x41, 0x20,
	0x50, 0xe6, 0x2b, 0xe7, 0x9e, 0xda, 0x79, 0xc5, 0x42, 0x26, 0x8b, 0x9d, 0xb5, 0x43, 0xfa, 0xd0,
	0x4e, 0xf0, 0x01, 0x13, 0x81, 0xfd, 0xc6, 0xb0, 0x36, 0xea, 0x78, 0x85, 0x6b, 0x27, 0x70, 0x38,
	0x8e, 0xfc, 0x55, 0x2a, 0x18, 0x8f, 0x7e, 0x24, 0x9c, 0x07, 0xe4, 0x10, 0xea, 0x54, 0xe6, 0x87,
	0xaf, 0x53, 0xb9, 0xbe, 0x62, 0x7d, 0xe3, 0x8a, 0x84, 0x40, 0x23, 0xe1, 0xdc, 0xb4, 0xe9, 0x79,
	0xda, 0x56, 0xd8, 0x0a, 0x69, 0xa0, 0x5b, 0xf4, 0x3c, 0x6d, 0x2b, 0x2c, 0xa6, 0x72, 0xd9, 0x6f,
	0x0e, 0x2d, 0x85, 0x29, 0xdb, 0xfe, 0x55, 0x83, 0xa3, 0x2f, 0x3c, 0x12, 0x4c, 0x48, 0x8c, 0xfc,
	0xcc, 0xb4, 0x3d, 0x86, 0x66, 0xc0, 0x12, 0x51, 0x74, 0x36, 0x8e, 0x5a, 0x53, 0xa0, 0xcf, 0xa3,
	0x79, 0xde, 0x3d, 0xf7, 0xc8, 0x33, 0x38, 0xd0, 0x09, 0xde, 0x7a, 0x86, 0x35, 0x40, 0x9e, 0x03,
	0x98, 0x3c, 0x1d, 0x36, 0xe3, 0x6c, 0x20, 0x95, 0x43, 0x8d, 0xe0, 0xf0, 0x16, 0xe9, 0x4a, 0x2e,
	0x3d, 0x14, 0x31, 0x8f, 0x04, 0xea, 0xde, 0x92, 0xca, 0x54, 0xe8, 0x91, 0x3a, 0x5e, 0xee, 0x9d,
	0xff, 0x6e, 0x40, 0x77, 0x1c, 0x86, 0xe9, 0x14, 0x93, 0x07, 0xe6, 0x23, 0xb9, 0x04, 0x6b, 0x8a,
	0x92, 0xec, 0x13, 0xd0, 0xe0, 0x78, 0xf7, 0xb5, 0xf5, 0x01, 0x3f, 0x40, 0x67, 0x8a, 0xf2, 0x8a,
	0x4a, 0x7f, 0x49, 0x9e, 0xec, 0x52, 0xb5, 0x3e, 0xf7, 0x10, 0x1d, 0xb0, 0x6e, 0x50, 0x92, 0x0a,
	0xe1, 0x0d, 0xaa, 0x74, 0x45, 0x3e, 0x42, 0xe7, 0xa6, 0x68, 0x74, 0x52, 0x26, 0xe9, 0x4e, 0xa7,
	0x15, 0x44, 0x2d, 0xdc, 0xcf, 0x70, 0xf0, 0x57, 0x1e, 0xa4, 0x72, 0x9e, 0xc1, 0x59, 0x09, 0xdd,
	0x92, 0xd3, 0x35, 0x74, 0x37, 0xde, 0x7a, 0x4f, 0x8d, 0x17, 0x3b, 0x68, 0x49, 0x1d, 0x97, 0xd0,
	0xbe, 0xca, 0xf2, 0xbf, 0xa1, 0xb2, 0x42, 0xe5, 0xea, 0xef, 0xa1, 0x7d, 0xcb, 0x84, 0xe4, 0x49,
	0x56, 0x79, 0xae, 0xbd, 0x5b, 0x7f, 0x82, 0x96, 0xd1, 0x02, 0x39, 0x71, 0xcc, 0xe7, 0xe7, 0x14,
	0x9f, 0x9f, 0xf3, 0x55, 0x7d, 0x7e, 0xa5, 0xa5, 0xb7, 0xa5, 0x73, 0xf5, 0xe6, 0xe7, 0xeb, 0x05,
	0x93, 0xcb, 0x74, 0xe6, 0xf8, 0x3c, 0x74, 0x7d, 0x3e, 0xc7, 0x88, 0x4b, 0x9a, 0x64, 0xae, 0x61,
	0xb9, 0xf1, 0xfd, 0xc2, 0xa5, 0x31, 0x73, 0x0d, 0x7b, 0xd6, 0xd2, 0xd5, 0x2f, 0xfe, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x90, 0xf4, 0x2b, 0x82, 0x87, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ImmuServiceClient is the client API for ImmuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImmuServiceClient interface {
	Set(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*Index, error)
	SetBatch(ctx context.Context, in *KVList, opts ...grpc.CallOption) (*Index, error)
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Item, error)
	GetBatch(ctx context.Context, in *KeyList, opts ...grpc.CallOption) (*ItemList, error)
	Inclusion(ctx context.Context, in *Index, opts ...grpc.CallOption) (*InclusionProof, error)
	Consistency(ctx context.Context, in *Index, opts ...grpc.CallOption) (*ConsistencyProof, error)
	ByIndex(ctx context.Context, in *Index, opts ...grpc.CallOption) (*Item, error)
	History(ctx context.Context, in *Key, opts ...grpc.CallOption) (*ItemList, error)
	Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthResponse, error)
}

type immuServiceClient struct {
	cc *grpc.ClientConn
}

func NewImmuServiceClient(cc *grpc.ClientConn) ImmuServiceClient {
	return &immuServiceClient{cc}
}

func (c *immuServiceClient) Set(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*Index, error) {
	out := new(Index)
	err := c.cc.Invoke(ctx, "/immudb.schema.ImmuService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceClient) SetBatch(ctx context.Context, in *KVList, opts ...grpc.CallOption) (*Index, error) {
	out := new(Index)
	err := c.cc.Invoke(ctx, "/immudb.schema.ImmuService/SetBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/immudb.schema.ImmuService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceClient) GetBatch(ctx context.Context, in *KeyList, opts ...grpc.CallOption) (*ItemList, error) {
	out := new(ItemList)
	err := c.cc.Invoke(ctx, "/immudb.schema.ImmuService/GetBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceClient) Inclusion(ctx context.Context, in *Index, opts ...grpc.CallOption) (*InclusionProof, error) {
	out := new(InclusionProof)
	err := c.cc.Invoke(ctx, "/immudb.schema.ImmuService/Inclusion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceClient) Consistency(ctx context.Context, in *Index, opts ...grpc.CallOption) (*ConsistencyProof, error) {
	out := new(ConsistencyProof)
	err := c.cc.Invoke(ctx, "/immudb.schema.ImmuService/Consistency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceClient) ByIndex(ctx context.Context, in *Index, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/immudb.schema.ImmuService/ByIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceClient) History(ctx context.Context, in *Key, opts ...grpc.CallOption) (*ItemList, error) {
	out := new(ItemList)
	err := c.cc.Invoke(ctx, "/immudb.schema.ImmuService/History", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceClient) Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/immudb.schema.ImmuService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImmuServiceServer is the server API for ImmuService service.
type ImmuServiceServer interface {
	Set(context.Context, *KeyValue) (*Index, error)
	SetBatch(context.Context, *KVList) (*Index, error)
	Get(context.Context, *Key) (*Item, error)
	GetBatch(context.Context, *KeyList) (*ItemList, error)
	Inclusion(context.Context, *Index) (*InclusionProof, error)
	Consistency(context.Context, *Index) (*ConsistencyProof, error)
	ByIndex(context.Context, *Index) (*Item, error)
	History(context.Context, *Key) (*ItemList, error)
	Health(context.Context, *empty.Empty) (*HealthResponse, error)
}

// UnimplementedImmuServiceServer can be embedded to have forward compatible implementations.
type UnimplementedImmuServiceServer struct {
}

func (*UnimplementedImmuServiceServer) Set(ctx context.Context, req *KeyValue) (*Index, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedImmuServiceServer) SetBatch(ctx context.Context, req *KVList) (*Index, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBatch not implemented")
}
func (*UnimplementedImmuServiceServer) Get(ctx context.Context, req *Key) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedImmuServiceServer) GetBatch(ctx context.Context, req *KeyList) (*ItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBatch not implemented")
}
func (*UnimplementedImmuServiceServer) Inclusion(ctx context.Context, req *Index) (*InclusionProof, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inclusion not implemented")
}
func (*UnimplementedImmuServiceServer) Consistency(ctx context.Context, req *Index) (*ConsistencyProof, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Consistency not implemented")
}
func (*UnimplementedImmuServiceServer) ByIndex(ctx context.Context, req *Index) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByIndex not implemented")
}
func (*UnimplementedImmuServiceServer) History(ctx context.Context, req *Key) (*ItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method History not implemented")
}
func (*UnimplementedImmuServiceServer) Health(ctx context.Context, req *empty.Empty) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}

func RegisterImmuServiceServer(s *grpc.Server, srv ImmuServiceServer) {
	s.RegisterService(&_ImmuService_serviceDesc, srv)
}

func _ImmuService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schema.ImmuService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceServer).Set(ctx, req.(*KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuService_SetBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KVList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceServer).SetBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schema.ImmuService/SetBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceServer).SetBatch(ctx, req.(*KVList))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schema.ImmuService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceServer).Get(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuService_GetBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceServer).GetBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schema.ImmuService/GetBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceServer).GetBatch(ctx, req.(*KeyList))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuService_Inclusion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Index)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceServer).Inclusion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schema.ImmuService/Inclusion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceServer).Inclusion(ctx, req.(*Index))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuService_Consistency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Index)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceServer).Consistency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schema.ImmuService/Consistency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceServer).Consistency(ctx, req.(*Index))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuService_ByIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Index)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceServer).ByIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schema.ImmuService/ByIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceServer).ByIndex(ctx, req.(*Index))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuService_History_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceServer).History(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schema.ImmuService/History",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceServer).History(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schema.ImmuService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceServer).Health(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImmuService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "immudb.schema.ImmuService",
	HandlerType: (*ImmuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _ImmuService_Set_Handler,
		},
		{
			MethodName: "SetBatch",
			Handler:    _ImmuService_SetBatch_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ImmuService_Get_Handler,
		},
		{
			MethodName: "GetBatch",
			Handler:    _ImmuService_GetBatch_Handler,
		},
		{
			MethodName: "Inclusion",
			Handler:    _ImmuService_Inclusion_Handler,
		},
		{
			MethodName: "Consistency",
			Handler:    _ImmuService_Consistency_Handler,
		},
		{
			MethodName: "ByIndex",
			Handler:    _ImmuService_ByIndex_Handler,
		},
		{
			MethodName: "History",
			Handler:    _ImmuService_History_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _ImmuService_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schema.proto",
}
