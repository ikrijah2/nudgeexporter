// Copyright 2020 Google LLC
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/cloud/automl/v1beta1/column_spec.proto

package automl

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

// A representation of a column in a relational table. When listing them, column specs are returned in the same order in which they were
// given on import .
// Used by:
//   *   Tables
type ColumnSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the column specs.
	// Form:
	//
	// `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/tableSpecs/{table_spec_id}/columnSpecs/{column_spec_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The data type of elements stored in the column.
	DataType *DataType `protobuf:"bytes,2,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	// Output only. The name of the column to show in the interface. The name can
	// be up to 100 characters long and can consist only of ASCII Latin letters
	// A-Z and a-z, ASCII digits 0-9, underscores(_), and forward slashes(/), and
	// must start with a letter or a digit.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. Stats of the series of values in the column.
	// This field may be stale, see the ancestor's
	// Dataset.tables_dataset_metadata.stats_update_time field
	// for the timestamp at which these stats were last updated.
	DataStats *DataStats `protobuf:"bytes,4,opt,name=data_stats,json=dataStats,proto3" json:"data_stats,omitempty"`
	// Deprecated.
	TopCorrelatedColumns []*ColumnSpec_CorrelatedColumn `protobuf:"bytes,5,rep,name=top_correlated_columns,json=topCorrelatedColumns,proto3" json:"top_correlated_columns,omitempty"`
	// Used to perform consistent read-modify-write updates. If not set, a blind
	// "overwrite" update happens.
	Etag string `protobuf:"bytes,6,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (x *ColumnSpec) Reset() {
	*x = ColumnSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_column_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColumnSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnSpec) ProtoMessage() {}

func (x *ColumnSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_column_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnSpec.ProtoReflect.Descriptor instead.
func (*ColumnSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_column_spec_proto_rawDescGZIP(), []int{0}
}

func (x *ColumnSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ColumnSpec) GetDataType() *DataType {
	if x != nil {
		return x.DataType
	}
	return nil
}

func (x *ColumnSpec) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ColumnSpec) GetDataStats() *DataStats {
	if x != nil {
		return x.DataStats
	}
	return nil
}

func (x *ColumnSpec) GetTopCorrelatedColumns() []*ColumnSpec_CorrelatedColumn {
	if x != nil {
		return x.TopCorrelatedColumns
	}
	return nil
}

func (x *ColumnSpec) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

// Identifies the table's column, and its correlation with the column this
// ColumnSpec describes.
type ColumnSpec_CorrelatedColumn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The column_spec_id of the correlated column, which belongs to the same
	// table as the in-context column.
	ColumnSpecId string `protobuf:"bytes,1,opt,name=column_spec_id,json=columnSpecId,proto3" json:"column_spec_id,omitempty"`
	// Correlation between this and the in-context column.
	CorrelationStats *CorrelationStats `protobuf:"bytes,2,opt,name=correlation_stats,json=correlationStats,proto3" json:"correlation_stats,omitempty"`
}

func (x *ColumnSpec_CorrelatedColumn) Reset() {
	*x = ColumnSpec_CorrelatedColumn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_column_spec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColumnSpec_CorrelatedColumn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnSpec_CorrelatedColumn) ProtoMessage() {}

func (x *ColumnSpec_CorrelatedColumn) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_column_spec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnSpec_CorrelatedColumn.ProtoReflect.Descriptor instead.
func (*ColumnSpec_CorrelatedColumn) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_column_spec_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ColumnSpec_CorrelatedColumn) GetColumnSpecId() string {
	if x != nil {
		return x.ColumnSpecId
	}
	return ""
}

func (x *ColumnSpec_CorrelatedColumn) GetCorrelationStats() *CorrelationStats {
	if x != nil {
		return x.CorrelationStats
	}
	return nil
}

var File_google_cloud_automl_v1beta1_column_spec_proto protoreflect.FileDescriptor

var file_google_cloud_automl_v1beta1_column_spec_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x80, 0x05, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x6e, 0x0a, 0x16, 0x74, 0x6f, 0x70, 0x5f, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x43, 0x6f, 0x72, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x14, 0x74, 0x6f,
	0x70, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x1a, 0x94, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x72, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x49,
	0x64, 0x12, 0x5a, 0x0a, 0x11, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x72, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x72,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x3a, 0x94, 0x01,
	0xea, 0x41, 0x90, 0x01, 0x0a, 0x20, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x6c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x7d, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x7b, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x7d, 0x2f, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x53, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x7d, 0x42, 0xa5, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0xca, 0x02, 0x1b,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x75, 0x74,
	0x6f, 0x4d, 0x6c, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x1e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x75, 0x74,
	0x6f, 0x4d, 0x4c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_automl_v1beta1_column_spec_proto_rawDescOnce sync.Once
	file_google_cloud_automl_v1beta1_column_spec_proto_rawDescData = file_google_cloud_automl_v1beta1_column_spec_proto_rawDesc
)

func file_google_cloud_automl_v1beta1_column_spec_proto_rawDescGZIP() []byte {
	file_google_cloud_automl_v1beta1_column_spec_proto_rawDescOnce.Do(func() {
		file_google_cloud_automl_v1beta1_column_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_automl_v1beta1_column_spec_proto_rawDescData)
	})
	return file_google_cloud_automl_v1beta1_column_spec_proto_rawDescData
}

var file_google_cloud_automl_v1beta1_column_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_automl_v1beta1_column_spec_proto_goTypes = []interface{}{
	(*ColumnSpec)(nil),                  // 0: google.cloud.automl.v1beta1.ColumnSpec
	(*ColumnSpec_CorrelatedColumn)(nil), // 1: google.cloud.automl.v1beta1.ColumnSpec.CorrelatedColumn
	(*DataType)(nil),                    // 2: google.cloud.automl.v1beta1.DataType
	(*DataStats)(nil),                   // 3: google.cloud.automl.v1beta1.DataStats
	(*CorrelationStats)(nil),            // 4: google.cloud.automl.v1beta1.CorrelationStats
}
var file_google_cloud_automl_v1beta1_column_spec_proto_depIdxs = []int32{
	2, // 0: google.cloud.automl.v1beta1.ColumnSpec.data_type:type_name -> google.cloud.automl.v1beta1.DataType
	3, // 1: google.cloud.automl.v1beta1.ColumnSpec.data_stats:type_name -> google.cloud.automl.v1beta1.DataStats
	1, // 2: google.cloud.automl.v1beta1.ColumnSpec.top_correlated_columns:type_name -> google.cloud.automl.v1beta1.ColumnSpec.CorrelatedColumn
	4, // 3: google.cloud.automl.v1beta1.ColumnSpec.CorrelatedColumn.correlation_stats:type_name -> google.cloud.automl.v1beta1.CorrelationStats
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_automl_v1beta1_column_spec_proto_init() }
func file_google_cloud_automl_v1beta1_column_spec_proto_init() {
	if File_google_cloud_automl_v1beta1_column_spec_proto != nil {
		return
	}
	file_google_cloud_automl_v1beta1_data_stats_proto_init()
	file_google_cloud_automl_v1beta1_data_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_automl_v1beta1_column_spec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColumnSpec); i {
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
		file_google_cloud_automl_v1beta1_column_spec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColumnSpec_CorrelatedColumn); i {
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
			RawDescriptor: file_google_cloud_automl_v1beta1_column_spec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_automl_v1beta1_column_spec_proto_goTypes,
		DependencyIndexes: file_google_cloud_automl_v1beta1_column_spec_proto_depIdxs,
		MessageInfos:      file_google_cloud_automl_v1beta1_column_spec_proto_msgTypes,
	}.Build()
	File_google_cloud_automl_v1beta1_column_spec_proto = out.File
	file_google_cloud_automl_v1beta1_column_spec_proto_rawDesc = nil
	file_google_cloud_automl_v1beta1_column_spec_proto_goTypes = nil
	file_google_cloud_automl_v1beta1_column_spec_proto_depIdxs = nil
}
