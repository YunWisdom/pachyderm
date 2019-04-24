// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/pachyderm/pachyderm/src/server/pkg/storage/chunk/chunk.proto

/*
Package chunk is a generated protocol buffer package.

It is generated from these files:
	github.com/pachyderm/pachyderm/src/server/pkg/storage/chunk/chunk.proto

It has these top-level messages:
	Chunk
	DataRef
*/
package chunk

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Chunk struct {
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
}

func (m *Chunk) Reset()                    { *m = Chunk{} }
func (m *Chunk) String() string            { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()               {}
func (*Chunk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Chunk) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// DataRef is a reference to data within a chunk.
type DataRef struct {
	// The chunk the referenced data is located in.
	Chunk *Chunk `protobuf:"bytes,1,opt,name=chunk" json:"chunk,omitempty"`
	// The hash of the data being referenced.
	// This field is empty when it is equal to the chunk hash (the ref is the whole chunk).
	Hash string `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	// The offset and size used for accessing the data within the chunk.
	Offset int64 `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
	Size   int64 `protobuf:"varint,4,opt,name=size" json:"size,omitempty"`
}

func (m *DataRef) Reset()                    { *m = DataRef{} }
func (m *DataRef) String() string            { return proto.CompactTextString(m) }
func (*DataRef) ProtoMessage()               {}
func (*DataRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DataRef) GetChunk() *Chunk {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func (m *DataRef) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *DataRef) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DataRef) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func init() {
	proto.RegisterType((*Chunk)(nil), "chunk.Chunk")
	proto.RegisterType((*DataRef)(nil), "chunk.DataRef")
}

func init() {
	proto.RegisterFile("github.com/pachyderm/pachyderm/src/server/pkg/storage/chunk/chunk.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x48, 0x4c, 0xce, 0xa8, 0x4c, 0x49, 0x2d, 0x42,
	0x66, 0x15, 0x17, 0x25, 0xeb, 0x17, 0xa7, 0x16, 0x95, 0xa5, 0x16, 0xe9, 0x17, 0x64, 0xa7, 0xeb,
	0x17, 0x97, 0xe4, 0x17, 0x25, 0xa6, 0xa7, 0xea, 0x27, 0x67, 0x94, 0xe6, 0x65, 0x43, 0x48, 0xbd,
	0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x56, 0x30, 0x47, 0x49, 0x9a, 0x8b, 0xd5, 0x19, 0xc4, 0x10,
	0x12, 0xe2, 0x62, 0xc9, 0x48, 0x2c, 0xce, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3,
	0x95, 0x72, 0xb9, 0xd8, 0x5d, 0x12, 0x4b, 0x12, 0x83, 0x52, 0xd3, 0x84, 0x94, 0xb8, 0x20, 0x1a,
	0xc0, 0xf2, 0xdc, 0x46, 0x3c, 0x7a, 0x10, 0xb3, 0xc0, 0x7a, 0x83, 0x20, 0x52, 0x70, 0x23, 0x98,
	0x10, 0x46, 0x08, 0x89, 0x71, 0xb1, 0xe5, 0xa7, 0xa5, 0x15, 0xa7, 0x96, 0x48, 0x30, 0x2b, 0x30,
	0x6a, 0x30, 0x07, 0x41, 0x79, 0x20, 0xb5, 0xc5, 0x99, 0x55, 0xa9, 0x12, 0x2c, 0x60, 0x51, 0x30,
	0xdb, 0xc9, 0x36, 0xca, 0x9a, 0x02, 0xdf, 0x25, 0xb1, 0x81, 0x3d, 0x66, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xb4, 0x01, 0xa1, 0x7c, 0x23, 0x01, 0x00, 0x00,
}
