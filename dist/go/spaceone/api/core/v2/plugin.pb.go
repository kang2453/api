// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.6.1
// source: spaceone/api/core/v2/plugin.proto

package v2

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PluginRequest_UpgradeMode int32

const (
	PluginRequest_NONE   PluginRequest_UpgradeMode = 0
	PluginRequest_MANUAL PluginRequest_UpgradeMode = 1
	PluginRequest_AUTO   PluginRequest_UpgradeMode = 2
)

// Enum value maps for PluginRequest_UpgradeMode.
var (
	PluginRequest_UpgradeMode_name = map[int32]string{
		0: "NONE",
		1: "MANUAL",
		2: "AUTO",
	}
	PluginRequest_UpgradeMode_value = map[string]int32{
		"NONE":   0,
		"MANUAL": 1,
		"AUTO":   2,
	}
)

func (x PluginRequest_UpgradeMode) Enum() *PluginRequest_UpgradeMode {
	p := new(PluginRequest_UpgradeMode)
	*p = x
	return p
}

func (x PluginRequest_UpgradeMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PluginRequest_UpgradeMode) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_core_v2_plugin_proto_enumTypes[0].Descriptor()
}

func (PluginRequest_UpgradeMode) Type() protoreflect.EnumType {
	return &file_spaceone_api_core_v2_plugin_proto_enumTypes[0]
}

func (x PluginRequest_UpgradeMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PluginRequest_UpgradeMode.Descriptor instead.
func (PluginRequest_UpgradeMode) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_core_v2_plugin_proto_rawDescGZIP(), []int{0, 0}
}

type PluginInfo_UpgradeMode int32

const (
	PluginInfo_NONE   PluginInfo_UpgradeMode = 0
	PluginInfo_MANUAL PluginInfo_UpgradeMode = 1
	PluginInfo_AUTO   PluginInfo_UpgradeMode = 2
)

// Enum value maps for PluginInfo_UpgradeMode.
var (
	PluginInfo_UpgradeMode_name = map[int32]string{
		0: "NONE",
		1: "MANUAL",
		2: "AUTO",
	}
	PluginInfo_UpgradeMode_value = map[string]int32{
		"NONE":   0,
		"MANUAL": 1,
		"AUTO":   2,
	}
)

func (x PluginInfo_UpgradeMode) Enum() *PluginInfo_UpgradeMode {
	p := new(PluginInfo_UpgradeMode)
	*p = x
	return p
}

func (x PluginInfo_UpgradeMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PluginInfo_UpgradeMode) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_core_v2_plugin_proto_enumTypes[1].Descriptor()
}

func (PluginInfo_UpgradeMode) Type() protoreflect.EnumType {
	return &file_spaceone_api_core_v2_plugin_proto_enumTypes[1]
}

func (x PluginInfo_UpgradeMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PluginInfo_UpgradeMode.Descriptor instead.
func (PluginInfo_UpgradeMode) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_core_v2_plugin_proto_rawDescGZIP(), []int{1, 0}
}

type PluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginId    string                    `protobuf:"bytes,1,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
	Version     string                    `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	UpgradeMode PluginRequest_UpgradeMode `protobuf:"varint,3,opt,name=upgrade_mode,json=upgradeMode,proto3,enum=spaceone.api.core.v2.PluginRequest_UpgradeMode" json:"upgrade_mode,omitempty"`
	Options     *_struct.Struct           `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
	SchemaId    string                    `protobuf:"bytes,5,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	SecretData  *_struct.Struct           `protobuf:"bytes,6,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
}

func (x *PluginRequest) Reset() {
	*x = PluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_core_v2_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginRequest) ProtoMessage() {}

func (x *PluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_core_v2_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginRequest.ProtoReflect.Descriptor instead.
func (*PluginRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_core_v2_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *PluginRequest) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

func (x *PluginRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PluginRequest) GetUpgradeMode() PluginRequest_UpgradeMode {
	if x != nil {
		return x.UpgradeMode
	}
	return PluginRequest_NONE
}

func (x *PluginRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *PluginRequest) GetSchemaId() string {
	if x != nil {
		return x.SchemaId
	}
	return ""
}

func (x *PluginRequest) GetSecretData() *_struct.Struct {
	if x != nil {
		return x.SecretData
	}
	return nil
}

type PluginInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginId    string                 `protobuf:"bytes,1,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
	Version     string                 `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	UpgradeMode PluginInfo_UpgradeMode `protobuf:"varint,3,opt,name=upgrade_mode,json=upgradeMode,proto3,enum=spaceone.api.core.v2.PluginInfo_UpgradeMode" json:"upgrade_mode,omitempty"`
	Options     *_struct.Struct        `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
	Metadata    *_struct.Struct        `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	SchemaId    string                 `protobuf:"bytes,6,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	SecretId    string                 `protobuf:"bytes,7,opt,name=secret_id,json=secretId,proto3" json:"secret_id,omitempty"`
}

func (x *PluginInfo) Reset() {
	*x = PluginInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_core_v2_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginInfo) ProtoMessage() {}

func (x *PluginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_core_v2_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginInfo.ProtoReflect.Descriptor instead.
func (*PluginInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_core_v2_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *PluginInfo) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

func (x *PluginInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PluginInfo) GetUpgradeMode() PluginInfo_UpgradeMode {
	if x != nil {
		return x.UpgradeMode
	}
	return PluginInfo_NONE
}

func (x *PluginInfo) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *PluginInfo) GetMetadata() *_struct.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *PluginInfo) GetSchemaId() string {
	if x != nil {
		return x.SchemaId
	}
	return ""
}

func (x *PluginInfo) GetSecretId() string {
	if x != nil {
		return x.SecretId
	}
	return ""
}

var File_spaceone_api_core_v2_plugin_proto protoreflect.FileDescriptor

var file_spaceone_api_core_v2_plugin_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x02, 0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x52, 0x0a, 0x0c, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x2d,
	0x0a, 0x0b, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x41, 0x4e, 0x55, 0x41,
	0x4c, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x55, 0x54, 0x4f, 0x10, 0x02, 0x22, 0xe5, 0x02,
	0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x0c, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32,
	0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x70, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x0b, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x41,
	0x55, 0x54, 0x4f, 0x10, 0x02, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_core_v2_plugin_proto_rawDescOnce sync.Once
	file_spaceone_api_core_v2_plugin_proto_rawDescData = file_spaceone_api_core_v2_plugin_proto_rawDesc
)

func file_spaceone_api_core_v2_plugin_proto_rawDescGZIP() []byte {
	file_spaceone_api_core_v2_plugin_proto_rawDescOnce.Do(func() {
		file_spaceone_api_core_v2_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_core_v2_plugin_proto_rawDescData)
	})
	return file_spaceone_api_core_v2_plugin_proto_rawDescData
}

var file_spaceone_api_core_v2_plugin_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_spaceone_api_core_v2_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_spaceone_api_core_v2_plugin_proto_goTypes = []interface{}{
	(PluginRequest_UpgradeMode)(0), // 0: spaceone.api.core.v2.PluginRequest.UpgradeMode
	(PluginInfo_UpgradeMode)(0),    // 1: spaceone.api.core.v2.PluginInfo.UpgradeMode
	(*PluginRequest)(nil),          // 2: spaceone.api.core.v2.PluginRequest
	(*PluginInfo)(nil),             // 3: spaceone.api.core.v2.PluginInfo
	(*_struct.Struct)(nil),         // 4: google.protobuf.Struct
}
var file_spaceone_api_core_v2_plugin_proto_depIdxs = []int32{
	0, // 0: spaceone.api.core.v2.PluginRequest.upgrade_mode:type_name -> spaceone.api.core.v2.PluginRequest.UpgradeMode
	4, // 1: spaceone.api.core.v2.PluginRequest.options:type_name -> google.protobuf.Struct
	4, // 2: spaceone.api.core.v2.PluginRequest.secret_data:type_name -> google.protobuf.Struct
	1, // 3: spaceone.api.core.v2.PluginInfo.upgrade_mode:type_name -> spaceone.api.core.v2.PluginInfo.UpgradeMode
	4, // 4: spaceone.api.core.v2.PluginInfo.options:type_name -> google.protobuf.Struct
	4, // 5: spaceone.api.core.v2.PluginInfo.metadata:type_name -> google.protobuf.Struct
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_spaceone_api_core_v2_plugin_proto_init() }
func file_spaceone_api_core_v2_plugin_proto_init() {
	if File_spaceone_api_core_v2_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_core_v2_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginRequest); i {
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
		file_spaceone_api_core_v2_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginInfo); i {
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
			RawDescriptor: file_spaceone_api_core_v2_plugin_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spaceone_api_core_v2_plugin_proto_goTypes,
		DependencyIndexes: file_spaceone_api_core_v2_plugin_proto_depIdxs,
		EnumInfos:         file_spaceone_api_core_v2_plugin_proto_enumTypes,
		MessageInfos:      file_spaceone_api_core_v2_plugin_proto_msgTypes,
	}.Build()
	File_spaceone_api_core_v2_plugin_proto = out.File
	file_spaceone_api_core_v2_plugin_proto_rawDesc = nil
	file_spaceone_api_core_v2_plugin_proto_goTypes = nil
	file_spaceone_api_core_v2_plugin_proto_depIdxs = nil
}
