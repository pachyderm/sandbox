// Code generated by protoc-gen-go.
// source: server/pfs/fuse/fuse.proto
// DO NOT EDIT!

/*
Package fuse is a generated protocol buffer package.

It is generated from these files:
	server/pfs/fuse/fuse.proto

It has these top-level messages:
	CommitMount
	Filesystem
	Node
	Attr
	Dirent
	Root
	DirectoryAttr
	DirectoryLookup
	DirectoryReadDirAll
	DirectoryCreate
	DirectoryMkdir
	FileAttr
	FileRead
	FileOpen
	FileWrite
*/
package fuse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import pfs "github.com/pachyderm/pachyderm/src/server/pfs"
import google_protobuf2 "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type CommitMount struct {
	Commit     *pfs.Commit `protobuf:"bytes,1,opt,name=commit" json:"commit,omitempty"`
	FromCommit *pfs.Commit `protobuf:"bytes,2,opt,name=from_commit,json=fromCommit" json:"from_commit,omitempty"`
	Alias      string      `protobuf:"bytes,3,opt,name=alias" json:"alias,omitempty"`
	Shard      *pfs.Shard  `protobuf:"bytes,4,opt,name=shard" json:"shard,omitempty"`
}

func (m *CommitMount) Reset()                    { *m = CommitMount{} }
func (m *CommitMount) String() string            { return proto.CompactTextString(m) }
func (*CommitMount) ProtoMessage()               {}
func (*CommitMount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CommitMount) GetCommit() *pfs.Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

func (m *CommitMount) GetFromCommit() *pfs.Commit {
	if m != nil {
		return m.FromCommit
	}
	return nil
}

func (m *CommitMount) GetShard() *pfs.Shard {
	if m != nil {
		return m.Shard
	}
	return nil
}

type Filesystem struct {
	Shard        *pfs.Shard     `protobuf:"bytes,1,opt,name=shard" json:"shard,omitempty"`
	CommitMounts []*CommitMount `protobuf:"bytes,2,rep,name=commit_mounts,json=commitMounts" json:"commit_mounts,omitempty"`
}

func (m *Filesystem) Reset()                    { *m = Filesystem{} }
func (m *Filesystem) String() string            { return proto.CompactTextString(m) }
func (*Filesystem) ProtoMessage()               {}
func (*Filesystem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Filesystem) GetShard() *pfs.Shard {
	if m != nil {
		return m.Shard
	}
	return nil
}

func (m *Filesystem) GetCommitMounts() []*CommitMount {
	if m != nil {
		return m.CommitMounts
	}
	return nil
}

type Node struct {
	File      *pfs.File                   `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	RepoAlias string                      `protobuf:"bytes,2,opt,name=repo_alias,json=repoAlias" json:"repo_alias,omitempty"`
	Write     bool                        `protobuf:"varint,3,opt,name=write" json:"write,omitempty"`
	Shard     *pfs.Shard                  `protobuf:"bytes,4,opt,name=shard" json:"shard,omitempty"`
	Modified  *google_protobuf2.Timestamp `protobuf:"bytes,5,opt,name=modified" json:"modified,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Node) GetFile() *pfs.File {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *Node) GetShard() *pfs.Shard {
	if m != nil {
		return m.Shard
	}
	return nil
}

func (m *Node) GetModified() *google_protobuf2.Timestamp {
	if m != nil {
		return m.Modified
	}
	return nil
}

type Attr struct {
	Mode uint32 `protobuf:"varint,1,opt,name=Mode,json=mode" json:"Mode,omitempty"`
}

func (m *Attr) Reset()                    { *m = Attr{} }
func (m *Attr) String() string            { return proto.CompactTextString(m) }
func (*Attr) ProtoMessage()               {}
func (*Attr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Dirent struct {
	Inode uint64 `protobuf:"varint,1,opt,name=inode" json:"inode,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Dirent) Reset()                    { *m = Dirent{} }
func (m *Dirent) String() string            { return proto.CompactTextString(m) }
func (*Dirent) ProtoMessage()               {}
func (*Dirent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Root struct {
	Filesystem *Filesystem `protobuf:"bytes,1,opt,name=filesystem" json:"filesystem,omitempty"`
	Result     *Node       `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error      string      `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *Root) Reset()                    { *m = Root{} }
func (m *Root) String() string            { return proto.CompactTextString(m) }
func (*Root) ProtoMessage()               {}
func (*Root) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Root) GetFilesystem() *Filesystem {
	if m != nil {
		return m.Filesystem
	}
	return nil
}

func (m *Root) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryAttr struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    *Attr  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error     string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *DirectoryAttr) Reset()                    { *m = DirectoryAttr{} }
func (m *DirectoryAttr) String() string            { return proto.CompactTextString(m) }
func (*DirectoryAttr) ProtoMessage()               {}
func (*DirectoryAttr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DirectoryAttr) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryAttr) GetResult() *Attr {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryLookup struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Result    *Node  `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
	Err       string `protobuf:"bytes,4,opt,name=err" json:"err,omitempty"`
}

func (m *DirectoryLookup) Reset()                    { *m = DirectoryLookup{} }
func (m *DirectoryLookup) String() string            { return proto.CompactTextString(m) }
func (*DirectoryLookup) ProtoMessage()               {}
func (*DirectoryLookup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DirectoryLookup) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryLookup) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryReadDirAll struct {
	Directory *Node     `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    []*Dirent `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
	Error     string    `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *DirectoryReadDirAll) Reset()                    { *m = DirectoryReadDirAll{} }
func (m *DirectoryReadDirAll) String() string            { return proto.CompactTextString(m) }
func (*DirectoryReadDirAll) ProtoMessage()               {}
func (*DirectoryReadDirAll) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DirectoryReadDirAll) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryReadDirAll) GetResult() []*Dirent {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryCreate struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    *Node  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error     string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *DirectoryCreate) Reset()                    { *m = DirectoryCreate{} }
func (m *DirectoryCreate) String() string            { return proto.CompactTextString(m) }
func (*DirectoryCreate) ProtoMessage()               {}
func (*DirectoryCreate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DirectoryCreate) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryCreate) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryMkdir struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    *Node  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error     string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *DirectoryMkdir) Reset()                    { *m = DirectoryMkdir{} }
func (m *DirectoryMkdir) String() string            { return proto.CompactTextString(m) }
func (*DirectoryMkdir) ProtoMessage()               {}
func (*DirectoryMkdir) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DirectoryMkdir) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryMkdir) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

type FileAttr struct {
	File   *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Result *Attr  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *FileAttr) Reset()                    { *m = FileAttr{} }
func (m *FileAttr) String() string            { return proto.CompactTextString(m) }
func (*FileAttr) ProtoMessage()               {}
func (*FileAttr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *FileAttr) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *FileAttr) GetResult() *Attr {
	if m != nil {
		return m.Result
	}
	return nil
}

type FileRead struct {
	File  *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *FileRead) Reset()                    { *m = FileRead{} }
func (m *FileRead) String() string            { return proto.CompactTextString(m) }
func (*FileRead) ProtoMessage()               {}
func (*FileRead) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *FileRead) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

type FileOpen struct {
	File  *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *FileOpen) Reset()                    { *m = FileOpen{} }
func (m *FileOpen) String() string            { return proto.CompactTextString(m) }
func (*FileOpen) ProtoMessage()               {}
func (*FileOpen) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *FileOpen) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

type FileWrite struct {
	File  *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *FileWrite) Reset()                    { *m = FileWrite{} }
func (m *FileWrite) String() string            { return proto.CompactTextString(m) }
func (*FileWrite) ProtoMessage()               {}
func (*FileWrite) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *FileWrite) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

func init() {
	proto.RegisterType((*CommitMount)(nil), "fuse.CommitMount")
	proto.RegisterType((*Filesystem)(nil), "fuse.Filesystem")
	proto.RegisterType((*Node)(nil), "fuse.Node")
	proto.RegisterType((*Attr)(nil), "fuse.Attr")
	proto.RegisterType((*Dirent)(nil), "fuse.Dirent")
	proto.RegisterType((*Root)(nil), "fuse.Root")
	proto.RegisterType((*DirectoryAttr)(nil), "fuse.DirectoryAttr")
	proto.RegisterType((*DirectoryLookup)(nil), "fuse.DirectoryLookup")
	proto.RegisterType((*DirectoryReadDirAll)(nil), "fuse.DirectoryReadDirAll")
	proto.RegisterType((*DirectoryCreate)(nil), "fuse.DirectoryCreate")
	proto.RegisterType((*DirectoryMkdir)(nil), "fuse.DirectoryMkdir")
	proto.RegisterType((*FileAttr)(nil), "fuse.FileAttr")
	proto.RegisterType((*FileRead)(nil), "fuse.FileRead")
	proto.RegisterType((*FileOpen)(nil), "fuse.FileOpen")
	proto.RegisterType((*FileWrite)(nil), "fuse.FileWrite")
}

var fileDescriptor0 = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x13, 0x27, 0x8a, 0x27, 0x0d, 0x94, 0xa5, 0x87, 0xc8, 0x12, 0x50, 0x19, 0x0e, 0x3d,
	0x20, 0x07, 0x15, 0x89, 0x33, 0x51, 0x11, 0x27, 0x02, 0xd2, 0x82, 0xc4, 0x31, 0x72, 0xe3, 0x75,
	0xb1, 0x6a, 0x7b, 0xad, 0xdd, 0x35, 0xa8, 0xe2, 0xcc, 0x7f, 0xe0, 0x87, 0xf0, 0x03, 0xd9, 0x9d,
	0xf5, 0x17, 0xa2, 0x51, 0x48, 0x90, 0x7a, 0x48, 0xb4, 0xb3, 0xf3, 0x3c, 0xef, 0xcd, 0x9b, 0xb1,
	0xc1, 0x97, 0x4c, 0x7c, 0x65, 0x62, 0x51, 0x26, 0x72, 0x91, 0x54, 0x92, 0xe1, 0x5f, 0x58, 0x0a,
	0xae, 0x38, 0x71, 0xcd, 0xd9, 0x3f, 0xe9, 0x21, 0xf4, 0xcf, 0xe6, 0xfc, 0x27, 0x57, 0x9c, 0x5f,
	0x65, 0x6c, 0x81, 0xd1, 0x65, 0x95, 0x2c, 0x54, 0x9a, 0x33, 0xa9, 0xa2, 0xbc, 0xb4, 0x80, 0xe0,
	0xa7, 0x03, 0xd3, 0x0b, 0x9e, 0xe7, 0xa9, 0x5a, 0xf1, 0xaa, 0x50, 0xe4, 0x29, 0x8c, 0x37, 0x18,
	0xce, 0x9d, 0x53, 0xe7, 0x6c, 0x7a, 0x3e, 0x0d, 0x4d, 0x31, 0x8b, 0xa0, 0x75, 0x8a, 0x3c, 0x87,
	0x69, 0x22, 0x78, 0xbe, 0xae, 0x91, 0x83, 0xbf, 0x91, 0x60, 0xf2, 0xf6, 0x4c, 0x4e, 0x60, 0x14,
	0x65, 0x69, 0x24, 0xe7, 0x43, 0x8d, 0xf3, 0xa8, 0x0d, 0xc8, 0x29, 0x8c, 0xe4, 0x97, 0x48, 0xc4,
	0x73, 0x17, 0x9f, 0x06, 0x7c, 0xfa, 0xa3, 0xb9, 0xa1, 0x36, 0x11, 0x24, 0x00, 0x6f, 0xd3, 0x8c,
	0xc9, 0x1b, 0xa9, 0x58, 0xde, 0xe1, 0x9d, 0x2d, 0x78, 0xf2, 0x0a, 0x66, 0x56, 0xd0, 0x3a, 0x37,
	0xad, 0x48, 0xad, 0x6b, 0xa8, 0x91, 0x0f, 0x42, 0xf4, 0xaa, 0xd7, 0x24, 0x3d, 0xda, 0x74, 0x81,
	0x0c, 0x7e, 0x39, 0xe0, 0xbe, 0xe7, 0x31, 0x23, 0x8f, 0xc0, 0x4d, 0x34, 0x61, 0xcd, 0xe0, 0x21,
	0x83, 0x51, 0x40, 0xf1, 0x5a, 0xa7, 0x41, 0xb0, 0x92, 0xaf, 0x6d, 0x33, 0x03, 0x6c, 0xc6, 0x33,
	0x37, 0x4b, 0x6c, 0x48, 0xb7, 0xf9, 0x4d, 0xa4, 0x8a, 0x61, 0x9b, 0x13, 0x6a, 0x83, 0xdd, 0x6d,
	0x6a, 0xd9, 0x93, 0x9c, 0xc7, 0x69, 0x92, 0xb2, 0x78, 0x3e, 0x42, 0x90, 0x1f, 0xda, 0xa9, 0x85,
	0xcd, 0xd4, 0xc2, 0x4f, 0xcd, 0xd4, 0x68, 0x8b, 0x0d, 0x7c, 0x70, 0x97, 0x4a, 0x09, 0x42, 0xc0,
	0x5d, 0x69, 0xf5, 0xa8, 0x7a, 0x46, 0x5d, 0x9d, 0x67, 0xc1, 0x39, 0x8c, 0xdf, 0xa4, 0x82, 0x15,
	0x68, 0x7e, 0x5a, 0x34, 0x69, 0x97, 0xda, 0xc0, 0x3c, 0x53, 0x44, 0x39, 0xab, 0x9b, 0xc0, 0x73,
	0x20, 0xc0, 0xa5, 0x9c, 0x2b, 0xf2, 0x02, 0x20, 0x69, 0x6d, 0xaf, 0xbd, 0x38, 0xb6, 0x1e, 0x76,
	0xe3, 0xa0, 0x3d, 0x0c, 0x09, 0x60, 0x2c, 0x98, 0xac, 0xb2, 0x66, 0x13, 0xc0, 0xa2, 0x8d, 0xa7,
	0xb4, 0xce, 0x18, 0x1d, 0x4c, 0x08, 0x2e, 0x9a, 0x25, 0xc0, 0x20, 0x90, 0x30, 0x33, 0x3a, 0x37,
	0x8a, 0x8b, 0x1b, 0x6c, 0xe6, 0x0c, 0xbc, 0xb8, 0xb9, 0x68, 0x27, 0xdd, 0x55, 0xeb, 0x92, 0xdb,
	0x48, 0x4d, 0x95, 0x1d, 0xa4, 0x3f, 0x1c, 0xb8, 0xdf, 0xb2, 0xbe, 0xe3, 0xfc, 0xba, 0x2a, 0xf7,
	0xe0, 0xbd, 0xc5, 0xba, 0x9e, 0x96, 0xe1, 0x56, 0x03, 0x8e, 0x61, 0xa8, 0xe9, 0x71, 0x0d, 0x3c,
	0x6a, 0x8e, 0xc1, 0x77, 0x78, 0xd8, 0xca, 0xa0, 0x2c, 0x8a, 0x75, 0xb0, 0xcc, 0xb2, 0x3d, 0xa4,
	0x3c, 0xeb, 0x59, 0x60, 0x36, 0xfd, 0xc8, 0xc2, 0xec, 0xe4, 0x77, 0x98, 0x50, 0xf5, 0x3c, 0xb8,
	0x10, 0x2c, 0xd2, 0xab, 0xfa, 0xdf, 0xde, 0xff, 0xc3, 0xc0, 0x15, 0xdc, 0x6b, 0x69, 0x57, 0xd7,
	0xba, 0xe2, 0x9d, 0xb0, 0xc6, 0x30, 0x31, 0xab, 0x8b, 0x1b, 0xf6, 0xf8, 0x8f, 0x97, 0xbc, 0x5f,
	0xc3, 0xbe, 0xe5, 0x87, 0xef, 0xd5, 0x6b, 0xcb, 0x62, 0x46, 0xb9, 0x93, 0xa5, 0xad, 0x30, 0xb8,
	0xa5, 0xc2, 0x87, 0x92, 0x15, 0x07, 0x56, 0x58, 0x82, 0x67, 0x2a, 0x7c, 0xc6, 0x6f, 0xcf, 0x41,
	0x25, 0x2e, 0xc7, 0xf8, 0xd5, 0x79, 0xf9, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x9d, 0x1f, 0x58,
	0x73, 0x06, 0x00, 0x00,
}
