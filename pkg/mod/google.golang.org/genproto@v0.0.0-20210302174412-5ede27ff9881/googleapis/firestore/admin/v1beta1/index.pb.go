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
// source: google/firestore/admin/v1beta1/index.proto

package admin

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// The mode determines how a field is indexed.
type IndexField_Mode int32

const (
	// The mode is unspecified.
	IndexField_MODE_UNSPECIFIED IndexField_Mode = 0
	// The field's values are indexed so as to support sequencing in
	// ascending order and also query by <, >, <=, >=, and =.
	IndexField_ASCENDING IndexField_Mode = 2
	// The field's values are indexed so as to support sequencing in
	// descending order and also query by <, >, <=, >=, and =.
	IndexField_DESCENDING IndexField_Mode = 3
	// The field's array values are indexed so as to support membership using
	// ARRAY_CONTAINS queries.
	IndexField_ARRAY_CONTAINS IndexField_Mode = 4
)

// Enum value maps for IndexField_Mode.
var (
	IndexField_Mode_name = map[int32]string{
		0: "MODE_UNSPECIFIED",
		2: "ASCENDING",
		3: "DESCENDING",
		4: "ARRAY_CONTAINS",
	}
	IndexField_Mode_value = map[string]int32{
		"MODE_UNSPECIFIED": 0,
		"ASCENDING":        2,
		"DESCENDING":       3,
		"ARRAY_CONTAINS":   4,
	}
)

func (x IndexField_Mode) Enum() *IndexField_Mode {
	p := new(IndexField_Mode)
	*p = x
	return p
}

func (x IndexField_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IndexField_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_google_firestore_admin_v1beta1_index_proto_enumTypes[0].Descriptor()
}

func (IndexField_Mode) Type() protoreflect.EnumType {
	return &file_google_firestore_admin_v1beta1_index_proto_enumTypes[0]
}

func (x IndexField_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IndexField_Mode.Descriptor instead.
func (IndexField_Mode) EnumDescriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1beta1_index_proto_rawDescGZIP(), []int{0, 0}
}

// The state of an index. During index creation, an index will be in the
// `CREATING` state. If the index is created successfully, it will transition
// to the `READY` state. If the index is not able to be created, it will
// transition to the `ERROR` state.
type Index_State int32

const (
	// The state is unspecified.
	Index_STATE_UNSPECIFIED Index_State = 0
	// The index is being created.
	// There is an active long-running operation for the index.
	// The index is updated when writing a document.
	// Some index data may exist.
	Index_CREATING Index_State = 3
	// The index is ready to be used.
	// The index is updated when writing a document.
	// The index is fully populated from all stored documents it applies to.
	Index_READY Index_State = 2
	// The index was being created, but something went wrong.
	// There is no active long-running operation for the index,
	// and the most recently finished long-running operation failed.
	// The index is not updated when writing a document.
	// Some index data may exist.
	Index_ERROR Index_State = 5
)

// Enum value maps for Index_State.
var (
	Index_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		3: "CREATING",
		2: "READY",
		5: "ERROR",
	}
	Index_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          3,
		"READY":             2,
		"ERROR":             5,
	}
)

func (x Index_State) Enum() *Index_State {
	p := new(Index_State)
	*p = x
	return p
}

func (x Index_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Index_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_firestore_admin_v1beta1_index_proto_enumTypes[1].Descriptor()
}

func (Index_State) Type() protoreflect.EnumType {
	return &file_google_firestore_admin_v1beta1_index_proto_enumTypes[1]
}

func (x Index_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Index_State.Descriptor instead.
func (Index_State) EnumDescriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1beta1_index_proto_rawDescGZIP(), []int{1, 0}
}

// A field of an index.
type IndexField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path of the field. Must match the field path specification described
	// by [google.firestore.v1beta1.Document.fields][fields].
	// Special field path `__name__` may be used by itself or at the end of a
	// path. `__type__` may be used only at the end of path.
	FieldPath string `protobuf:"bytes,1,opt,name=field_path,json=fieldPath,proto3" json:"field_path,omitempty"`
	// The field's mode.
	Mode IndexField_Mode `protobuf:"varint,2,opt,name=mode,proto3,enum=google.firestore.admin.v1beta1.IndexField_Mode" json:"mode,omitempty"`
}

func (x *IndexField) Reset() {
	*x = IndexField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firestore_admin_v1beta1_index_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexField) ProtoMessage() {}

func (x *IndexField) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_admin_v1beta1_index_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexField.ProtoReflect.Descriptor instead.
func (*IndexField) Descriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1beta1_index_proto_rawDescGZIP(), []int{0}
}

func (x *IndexField) GetFieldPath() string {
	if x != nil {
		return x.FieldPath
	}
	return ""
}

func (x *IndexField) GetMode() IndexField_Mode {
	if x != nil {
		return x.Mode
	}
	return IndexField_MODE_UNSPECIFIED
}

// An index definition.
type Index struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the index.
	// Output only.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The collection ID to which this index applies. Required.
	CollectionId string `protobuf:"bytes,2,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	// The fields to index.
	Fields []*IndexField `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
	// The state of the index.
	// Output only.
	State Index_State `protobuf:"varint,6,opt,name=state,proto3,enum=google.firestore.admin.v1beta1.Index_State" json:"state,omitempty"`
}

func (x *Index) Reset() {
	*x = Index{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firestore_admin_v1beta1_index_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index) ProtoMessage() {}

func (x *Index) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_admin_v1beta1_index_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Index.ProtoReflect.Descriptor instead.
func (*Index) Descriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1beta1_index_proto_rawDescGZIP(), []int{1}
}

func (x *Index) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Index) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

func (x *Index) GetFields() []*IndexField {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Index) GetState() Index_State {
	if x != nil {
		return x.State
	}
	return Index_STATE_UNSPECIFIED
}

var File_google_firestore_admin_v1beta1_index_proto protoreflect.FileDescriptor

var file_google_firestore_admin_v1beta1_index_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x01, 0x0a, 0x0a, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x68, 0x12, 0x43, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x4f, 0x0a,
	0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41,
	0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x45,
	0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x52,
	0x52, 0x41, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x10, 0x04, 0x22, 0x8b,
	0x02, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x42, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x41, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69,
	0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x42, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x42, 0xa5, 0x01, 0x0a,
	0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x42, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0xa2, 0x02, 0x04, 0x47, 0x43, 0x46, 0x53, 0xaa, 0x02, 0x24,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x46, 0x69, 0x72,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x56, 0x31, 0x42,
	0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_firestore_admin_v1beta1_index_proto_rawDescOnce sync.Once
	file_google_firestore_admin_v1beta1_index_proto_rawDescData = file_google_firestore_admin_v1beta1_index_proto_rawDesc
)

func file_google_firestore_admin_v1beta1_index_proto_rawDescGZIP() []byte {
	file_google_firestore_admin_v1beta1_index_proto_rawDescOnce.Do(func() {
		file_google_firestore_admin_v1beta1_index_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_firestore_admin_v1beta1_index_proto_rawDescData)
	})
	return file_google_firestore_admin_v1beta1_index_proto_rawDescData
}

var file_google_firestore_admin_v1beta1_index_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_firestore_admin_v1beta1_index_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_firestore_admin_v1beta1_index_proto_goTypes = []interface{}{
	(IndexField_Mode)(0), // 0: google.firestore.admin.v1beta1.IndexField.Mode
	(Index_State)(0),     // 1: google.firestore.admin.v1beta1.Index.State
	(*IndexField)(nil),   // 2: google.firestore.admin.v1beta1.IndexField
	(*Index)(nil),        // 3: google.firestore.admin.v1beta1.Index
}
var file_google_firestore_admin_v1beta1_index_proto_depIdxs = []int32{
	0, // 0: google.firestore.admin.v1beta1.IndexField.mode:type_name -> google.firestore.admin.v1beta1.IndexField.Mode
	2, // 1: google.firestore.admin.v1beta1.Index.fields:type_name -> google.firestore.admin.v1beta1.IndexField
	1, // 2: google.firestore.admin.v1beta1.Index.state:type_name -> google.firestore.admin.v1beta1.Index.State
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_firestore_admin_v1beta1_index_proto_init() }
func file_google_firestore_admin_v1beta1_index_proto_init() {
	if File_google_firestore_admin_v1beta1_index_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_firestore_admin_v1beta1_index_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexField); i {
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
		file_google_firestore_admin_v1beta1_index_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Index); i {
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
			RawDescriptor: file_google_firestore_admin_v1beta1_index_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_firestore_admin_v1beta1_index_proto_goTypes,
		DependencyIndexes: file_google_firestore_admin_v1beta1_index_proto_depIdxs,
		EnumInfos:         file_google_firestore_admin_v1beta1_index_proto_enumTypes,
		MessageInfos:      file_google_firestore_admin_v1beta1_index_proto_msgTypes,
	}.Build()
	File_google_firestore_admin_v1beta1_index_proto = out.File
	file_google_firestore_admin_v1beta1_index_proto_rawDesc = nil
	file_google_firestore_admin_v1beta1_index_proto_goTypes = nil
	file_google_firestore_admin_v1beta1_index_proto_depIdxs = nil
}
