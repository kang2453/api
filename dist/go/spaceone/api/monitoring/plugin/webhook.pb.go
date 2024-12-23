// A Webhook is a plugin receiving data from external monitoring systems.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.6.1
// source: spaceone/api/monitoring/plugin/webhook.proto

package plugin

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type WebhookInitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Options       *_struct.Struct        `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WebhookInitRequest) Reset() {
	*x = WebhookInitRequest{}
	mi := &file_spaceone_api_monitoring_plugin_webhook_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebhookInitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookInitRequest) ProtoMessage() {}

func (x *WebhookInitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_plugin_webhook_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookInitRequest.ProtoReflect.Descriptor instead.
func (*WebhookInitRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_plugin_webhook_proto_rawDescGZIP(), []int{0}
}

func (x *WebhookInitRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

type WebhookPluginVerifyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Options       *_struct.Struct        `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WebhookPluginVerifyRequest) Reset() {
	*x = WebhookPluginVerifyRequest{}
	mi := &file_spaceone_api_monitoring_plugin_webhook_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebhookPluginVerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookPluginVerifyRequest) ProtoMessage() {}

func (x *WebhookPluginVerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_plugin_webhook_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookPluginVerifyRequest.ProtoReflect.Descriptor instead.
func (*WebhookPluginVerifyRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_plugin_webhook_proto_rawDescGZIP(), []int{1}
}

func (x *WebhookPluginVerifyRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

type WebhookPluginInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *_struct.Struct        `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WebhookPluginInfo) Reset() {
	*x = WebhookPluginInfo{}
	mi := &file_spaceone_api_monitoring_plugin_webhook_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebhookPluginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookPluginInfo) ProtoMessage() {}

func (x *WebhookPluginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_plugin_webhook_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookPluginInfo.ProtoReflect.Descriptor instead.
func (*WebhookPluginInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_plugin_webhook_proto_rawDescGZIP(), []int{2}
}

func (x *WebhookPluginInfo) GetMetadata() *_struct.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_spaceone_api_monitoring_plugin_webhook_proto protoreflect.FileDescriptor

var file_spaceone_api_monitoring_plugin_webhook_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x12, 0x57, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x4f, 0x0a, 0x1a, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x48, 0x0a, 0x11, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0xda, 0x01,
	0x0a, 0x07, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x6f, 0x0a, 0x04, 0x69, 0x6e, 0x69,
	0x74, 0x12, 0x32, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x06, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x12, 0x3a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f,
	0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f,
	0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_monitoring_plugin_webhook_proto_rawDescOnce sync.Once
	file_spaceone_api_monitoring_plugin_webhook_proto_rawDescData = file_spaceone_api_monitoring_plugin_webhook_proto_rawDesc
)

func file_spaceone_api_monitoring_plugin_webhook_proto_rawDescGZIP() []byte {
	file_spaceone_api_monitoring_plugin_webhook_proto_rawDescOnce.Do(func() {
		file_spaceone_api_monitoring_plugin_webhook_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_monitoring_plugin_webhook_proto_rawDescData)
	})
	return file_spaceone_api_monitoring_plugin_webhook_proto_rawDescData
}

var file_spaceone_api_monitoring_plugin_webhook_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spaceone_api_monitoring_plugin_webhook_proto_goTypes = []any{
	(*WebhookInitRequest)(nil),         // 0: spaceone.api.monitoring.plugin.WebhookInitRequest
	(*WebhookPluginVerifyRequest)(nil), // 1: spaceone.api.monitoring.plugin.WebhookPluginVerifyRequest
	(*WebhookPluginInfo)(nil),          // 2: spaceone.api.monitoring.plugin.WebhookPluginInfo
	(*_struct.Struct)(nil),             // 3: google.protobuf.Struct
	(*empty.Empty)(nil),                // 4: google.protobuf.Empty
}
var file_spaceone_api_monitoring_plugin_webhook_proto_depIdxs = []int32{
	3, // 0: spaceone.api.monitoring.plugin.WebhookInitRequest.options:type_name -> google.protobuf.Struct
	3, // 1: spaceone.api.monitoring.plugin.WebhookPluginVerifyRequest.options:type_name -> google.protobuf.Struct
	3, // 2: spaceone.api.monitoring.plugin.WebhookPluginInfo.metadata:type_name -> google.protobuf.Struct
	0, // 3: spaceone.api.monitoring.plugin.Webhook.init:input_type -> spaceone.api.monitoring.plugin.WebhookInitRequest
	1, // 4: spaceone.api.monitoring.plugin.Webhook.verify:input_type -> spaceone.api.monitoring.plugin.WebhookPluginVerifyRequest
	2, // 5: spaceone.api.monitoring.plugin.Webhook.init:output_type -> spaceone.api.monitoring.plugin.WebhookPluginInfo
	4, // 6: spaceone.api.monitoring.plugin.Webhook.verify:output_type -> google.protobuf.Empty
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_spaceone_api_monitoring_plugin_webhook_proto_init() }
func file_spaceone_api_monitoring_plugin_webhook_proto_init() {
	if File_spaceone_api_monitoring_plugin_webhook_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_monitoring_plugin_webhook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_monitoring_plugin_webhook_proto_goTypes,
		DependencyIndexes: file_spaceone_api_monitoring_plugin_webhook_proto_depIdxs,
		MessageInfos:      file_spaceone_api_monitoring_plugin_webhook_proto_msgTypes,
	}.Build()
	File_spaceone_api_monitoring_plugin_webhook_proto = out.File
	file_spaceone_api_monitoring_plugin_webhook_proto_rawDesc = nil
	file_spaceone_api_monitoring_plugin_webhook_proto_goTypes = nil
	file_spaceone_api_monitoring_plugin_webhook_proto_depIdxs = nil
}
