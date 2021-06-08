// Copyright 2019 Google LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/devtools/resultstore/v2/file.proto

package resultstore

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// If known, the hash function used to compute this digest.
type File_HashType int32

const (
	// Unknown
	File_HASH_TYPE_UNSPECIFIED File_HashType = 0
	// MD5
	File_MD5 File_HashType = 1
	// SHA-1
	File_SHA1 File_HashType = 2
	// SHA-256
	File_SHA256 File_HashType = 3
)

// Enum value maps for File_HashType.
var (
	File_HashType_name = map[int32]string{
		0: "HASH_TYPE_UNSPECIFIED",
		1: "MD5",
		2: "SHA1",
		3: "SHA256",
	}
	File_HashType_value = map[string]int32{
		"HASH_TYPE_UNSPECIFIED": 0,
		"MD5":                   1,
		"SHA1":                  2,
		"SHA256":                3,
	}
)

func (x File_HashType) Enum() *File_HashType {
	p := new(File_HashType)
	*p = x
	return p
}

func (x File_HashType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (File_HashType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_resultstore_v2_file_proto_enumTypes[0].Descriptor()
}

func (File_HashType) Type() protoreflect.EnumType {
	return &file_google_devtools_resultstore_v2_file_proto_enumTypes[0]
}

func (x File_HashType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use File_HashType.Descriptor instead.
func (File_HashType) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_file_proto_rawDescGZIP(), []int{0, 0}
}

// The metadata for a file or an archive file entry.
type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identifier of the file or archive entry.
	// User-provided, must be unique for the repeated field it is in. When an
	// Append RPC is called with a Files field populated, if a File already exists
	// with this ID, that File will be overwritten with the new File proto.
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// The URI of a file.
	// This could also be the URI of an entire archive.
	// Most log data doesn't need to be stored forever, so a ttl is suggested.
	// Note that if you ever move or delete the file at this URI, the link from
	// the server will be broken.
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	// (Optional) The length of the file in bytes.  Allows the filesize to be
	// shown in the UI.  Omit if file is still being written or length is
	// not known.  This could also be the length of an entire archive.
	Length *wrapperspb.Int64Value `protobuf:"bytes,3,opt,name=length,proto3" json:"length,omitempty"`
	// (Optional) The content-type (aka MIME-type) of the file.  This is sent to
	// the web browser so it knows how to handle the file. (e.g. text/plain,
	// image/jpeg, text/html, etc). For zip archives, use "application/zip".
	ContentType string `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// (Optional) If the above path, length, and content_type are referring to an
	// archive, and you wish to refer to a particular entry within that archive,
	// put the particular archive entry data here.
	ArchiveEntry *ArchiveEntry `protobuf:"bytes,5,opt,name=archive_entry,json=archiveEntry,proto3" json:"archive_entry,omitempty"`
	// (Optional) A url to a content display app/site for this file or archive
	// entry.
	ContentViewer string `protobuf:"bytes,6,opt,name=content_viewer,json=contentViewer,proto3" json:"content_viewer,omitempty"`
	// (Optional) Whether to hide this file or archive entry in the UI.  Defaults
	// to false. A checkbox lets users see hidden files, but they're hidden by
	// default.
	Hidden bool `protobuf:"varint,7,opt,name=hidden,proto3" json:"hidden,omitempty"`
	// (Optional) A short description of what this file or archive entry
	// contains. This description should help someone viewing the list of these
	// files to understand the purpose of this file and what they would want to
	// view it for.
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	// (Optional) digest of this file in hexadecimal-like string if known.
	Digest string `protobuf:"bytes,9,opt,name=digest,proto3" json:"digest,omitempty"`
	// (Optional) The algorithm corresponding to the digest if known.
	HashType File_HashType `protobuf:"varint,10,opt,name=hash_type,json=hashType,proto3,enum=google.devtools.resultstore.v2.File_HashType" json:"hash_type,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_file_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *File) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *File) GetLength() *wrapperspb.Int64Value {
	if x != nil {
		return x.Length
	}
	return nil
}

func (x *File) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *File) GetArchiveEntry() *ArchiveEntry {
	if x != nil {
		return x.ArchiveEntry
	}
	return nil
}

func (x *File) GetContentViewer() string {
	if x != nil {
		return x.ContentViewer
	}
	return ""
}

func (x *File) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

func (x *File) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *File) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *File) GetHashType() File_HashType {
	if x != nil {
		return x.HashType
	}
	return File_HASH_TYPE_UNSPECIFIED
}

// Information specific to an entry in an archive.
type ArchiveEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The relative path of the entry within the archive.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// (Optional) The uncompressed length of the archive entry in bytes.  Allows
	// the entry size to be shown in the UI.  Omit if the length is not known.
	Length *wrapperspb.Int64Value `protobuf:"bytes,2,opt,name=length,proto3" json:"length,omitempty"`
	// (Optional) The content-type (aka MIME-type) of the archive entry. (e.g.
	// text/plain, image/jpeg, text/html, etc). This is sent to the web browser
	// so it knows how to handle the entry.
	ContentType string `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
}

func (x *ArchiveEntry) Reset() {
	*x = ArchiveEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArchiveEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArchiveEntry) ProtoMessage() {}

func (x *ArchiveEntry) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArchiveEntry.ProtoReflect.Descriptor instead.
func (*ArchiveEntry) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_file_proto_rawDescGZIP(), []int{1}
}

func (x *ArchiveEntry) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ArchiveEntry) GetLength() *wrapperspb.Int64Value {
	if x != nil {
		return x.Length
	}
	return nil
}

func (x *ArchiveEntry) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

var File_google_devtools_resultstore_v2_file_proto protoreflect.FileDescriptor

var file_google_devtools_resultstore_v2_file_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x03, 0x0a, 0x04,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x33, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x51, 0x0a, 0x0d, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x76,
	0x69, 0x65, 0x77, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69,
	0x64, 0x64, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64,
	0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x09,
	0x68, 0x61, 0x73, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08,
	0x68, 0x61, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x22, 0x44, 0x0a, 0x08, 0x48, 0x61, 0x73, 0x68,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x48, 0x41, 0x53, 0x48, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x4d, 0x44, 0x35, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x48, 0x41, 0x31,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x03, 0x22, 0x7a,
	0x0a, 0x0c, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x33, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x71, 0x0a, 0x22, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32,
	0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x32, 0x3b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_resultstore_v2_file_proto_rawDescOnce sync.Once
	file_google_devtools_resultstore_v2_file_proto_rawDescData = file_google_devtools_resultstore_v2_file_proto_rawDesc
)

func file_google_devtools_resultstore_v2_file_proto_rawDescGZIP() []byte {
	file_google_devtools_resultstore_v2_file_proto_rawDescOnce.Do(func() {
		file_google_devtools_resultstore_v2_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_resultstore_v2_file_proto_rawDescData)
	})
	return file_google_devtools_resultstore_v2_file_proto_rawDescData
}

var file_google_devtools_resultstore_v2_file_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_devtools_resultstore_v2_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_devtools_resultstore_v2_file_proto_goTypes = []interface{}{
	(File_HashType)(0),            // 0: google.devtools.resultstore.v2.File.HashType
	(*File)(nil),                  // 1: google.devtools.resultstore.v2.File
	(*ArchiveEntry)(nil),          // 2: google.devtools.resultstore.v2.ArchiveEntry
	(*wrapperspb.Int64Value)(nil), // 3: google.protobuf.Int64Value
}
var file_google_devtools_resultstore_v2_file_proto_depIdxs = []int32{
	3, // 0: google.devtools.resultstore.v2.File.length:type_name -> google.protobuf.Int64Value
	2, // 1: google.devtools.resultstore.v2.File.archive_entry:type_name -> google.devtools.resultstore.v2.ArchiveEntry
	0, // 2: google.devtools.resultstore.v2.File.hash_type:type_name -> google.devtools.resultstore.v2.File.HashType
	3, // 3: google.devtools.resultstore.v2.ArchiveEntry.length:type_name -> google.protobuf.Int64Value
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_devtools_resultstore_v2_file_proto_init() }
func file_google_devtools_resultstore_v2_file_proto_init() {
	if File_google_devtools_resultstore_v2_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_resultstore_v2_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_devtools_resultstore_v2_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArchiveEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_devtools_resultstore_v2_file_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_resultstore_v2_file_proto_goTypes,
		DependencyIndexes: file_google_devtools_resultstore_v2_file_proto_depIdxs,
		EnumInfos:         file_google_devtools_resultstore_v2_file_proto_enumTypes,
		MessageInfos:      file_google_devtools_resultstore_v2_file_proto_msgTypes,
	}.Build()
	File_google_devtools_resultstore_v2_file_proto = out.File
	file_google_devtools_resultstore_v2_file_proto_rawDesc = nil
	file_google_devtools_resultstore_v2_file_proto_goTypes = nil
	file_google_devtools_resultstore_v2_file_proto_depIdxs = nil
}
