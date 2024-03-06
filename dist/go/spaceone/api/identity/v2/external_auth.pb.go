// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.6.1
// source: spaceone/api/identity/v2/external_auth.proto

package v2

import (
	v2 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ExternalAuthInfo_State int32

const (
	ExternalAuthInfo_NONE     ExternalAuthInfo_State = 0
	ExternalAuthInfo_ENABLED  ExternalAuthInfo_State = 1
	ExternalAuthInfo_DISABLED ExternalAuthInfo_State = 2
)

// Enum value maps for ExternalAuthInfo_State.
var (
	ExternalAuthInfo_State_name = map[int32]string{
		0: "NONE",
		1: "ENABLED",
		2: "DISABLED",
	}
	ExternalAuthInfo_State_value = map[string]int32{
		"NONE":     0,
		"ENABLED":  1,
		"DISABLED": 2,
	}
)

func (x ExternalAuthInfo_State) Enum() *ExternalAuthInfo_State {
	p := new(ExternalAuthInfo_State)
	*p = x
	return p
}

func (x ExternalAuthInfo_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExternalAuthInfo_State) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_external_auth_proto_enumTypes[0].Descriptor()
}

func (ExternalAuthInfo_State) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_external_auth_proto_enumTypes[0]
}

func (x ExternalAuthInfo_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExternalAuthInfo_State.Descriptor instead.
func (ExternalAuthInfo_State) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_external_auth_proto_rawDescGZIP(), []int{2, 0}
}

type SetExternalAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginInfo *v2.PluginRequest `protobuf:"bytes,2,opt,name=plugin_info,json=pluginInfo,proto3" json:"plugin_info,omitempty"`
}

func (x *SetExternalAuthRequest) Reset() {
	*x = SetExternalAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_external_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetExternalAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetExternalAuthRequest) ProtoMessage() {}

func (x *SetExternalAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_external_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetExternalAuthRequest.ProtoReflect.Descriptor instead.
func (*SetExternalAuthRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_external_auth_proto_rawDescGZIP(), []int{0}
}

func (x *SetExternalAuthRequest) GetPluginInfo() *v2.PluginRequest {
	if x != nil {
		return x.PluginInfo
	}
	return nil
}

type ExternalAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExternalAuthRequest) Reset() {
	*x = ExternalAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_external_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalAuthRequest) ProtoMessage() {}

func (x *ExternalAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_external_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalAuthRequest.ProtoReflect.Descriptor instead.
func (*ExternalAuthRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_external_auth_proto_rawDescGZIP(), []int{1}
}

type ExternalAuthInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainId   string                 `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	State      ExternalAuthInfo_State `protobuf:"varint,2,opt,name=state,proto3,enum=spaceone.api.identity.v2.ExternalAuthInfo_State" json:"state,omitempty"`
	PluginInfo *v2.PluginInfo         `protobuf:"bytes,3,opt,name=plugin_info,json=pluginInfo,proto3" json:"plugin_info,omitempty"`
	UpdatedAt  string                 `protobuf:"bytes,31,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *ExternalAuthInfo) Reset() {
	*x = ExternalAuthInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_external_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalAuthInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalAuthInfo) ProtoMessage() {}

func (x *ExternalAuthInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_external_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalAuthInfo.ProtoReflect.Descriptor instead.
func (*ExternalAuthInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_external_auth_proto_rawDescGZIP(), []int{2}
}

func (x *ExternalAuthInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *ExternalAuthInfo) GetState() ExternalAuthInfo_State {
	if x != nil {
		return x.State
	}
	return ExternalAuthInfo_NONE
}

func (x *ExternalAuthInfo) GetPluginInfo() *v2.PluginInfo {
	if x != nil {
		return x.PluginInfo
	}
	return nil
}

func (x *ExternalAuthInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_spaceone_api_identity_v2_external_auth_proto protoreflect.FileDescriptor

var file_spaceone_api_identity_v2_external_auth_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44,
	0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x87, 0x02, 0x0a, 0x10,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x46, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2c, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e,
	0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42,
	0x4c, 0x45, 0x44, 0x10, 0x02, 0x32, 0xbf, 0x03, 0x0a, 0x0c, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x12, 0x8e, 0x01, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x12, 0x30,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x74, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x29, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2d, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x73, 0x65, 0x74, 0x12, 0x8f, 0x01, 0x0a, 0x05, 0x75, 0x6e, 0x73, 0x65,
	0x74, 0x12, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2d, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x75, 0x6e, 0x73, 0x65, 0x74, 0x12, 0x8b, 0x01, 0x0a, 0x03, 0x67, 0x65,
	0x74, 0x12, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x29, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2d, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x67, 0x65, 0x74, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74,
	0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_identity_v2_external_auth_proto_rawDescOnce sync.Once
	file_spaceone_api_identity_v2_external_auth_proto_rawDescData = file_spaceone_api_identity_v2_external_auth_proto_rawDesc
)

func file_spaceone_api_identity_v2_external_auth_proto_rawDescGZIP() []byte {
	file_spaceone_api_identity_v2_external_auth_proto_rawDescOnce.Do(func() {
		file_spaceone_api_identity_v2_external_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_identity_v2_external_auth_proto_rawDescData)
	})
	return file_spaceone_api_identity_v2_external_auth_proto_rawDescData
}

var file_spaceone_api_identity_v2_external_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spaceone_api_identity_v2_external_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spaceone_api_identity_v2_external_auth_proto_goTypes = []interface{}{
	(ExternalAuthInfo_State)(0),    // 0: spaceone.api.identity.v2.ExternalAuthInfo.State
	(*SetExternalAuthRequest)(nil), // 1: spaceone.api.identity.v2.SetExternalAuthRequest
	(*ExternalAuthRequest)(nil),    // 2: spaceone.api.identity.v2.ExternalAuthRequest
	(*ExternalAuthInfo)(nil),       // 3: spaceone.api.identity.v2.ExternalAuthInfo
	(*v2.PluginRequest)(nil),       // 4: spaceone.api.core.v2.PluginRequest
	(*v2.PluginInfo)(nil),          // 5: spaceone.api.core.v2.PluginInfo
}
var file_spaceone_api_identity_v2_external_auth_proto_depIdxs = []int32{
	4, // 0: spaceone.api.identity.v2.SetExternalAuthRequest.plugin_info:type_name -> spaceone.api.core.v2.PluginRequest
	0, // 1: spaceone.api.identity.v2.ExternalAuthInfo.state:type_name -> spaceone.api.identity.v2.ExternalAuthInfo.State
	5, // 2: spaceone.api.identity.v2.ExternalAuthInfo.plugin_info:type_name -> spaceone.api.core.v2.PluginInfo
	1, // 3: spaceone.api.identity.v2.ExternalAuth.set:input_type -> spaceone.api.identity.v2.SetExternalAuthRequest
	2, // 4: spaceone.api.identity.v2.ExternalAuth.unset:input_type -> spaceone.api.identity.v2.ExternalAuthRequest
	2, // 5: spaceone.api.identity.v2.ExternalAuth.get:input_type -> spaceone.api.identity.v2.ExternalAuthRequest
	3, // 6: spaceone.api.identity.v2.ExternalAuth.set:output_type -> spaceone.api.identity.v2.ExternalAuthInfo
	3, // 7: spaceone.api.identity.v2.ExternalAuth.unset:output_type -> spaceone.api.identity.v2.ExternalAuthInfo
	3, // 8: spaceone.api.identity.v2.ExternalAuth.get:output_type -> spaceone.api.identity.v2.ExternalAuthInfo
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_spaceone_api_identity_v2_external_auth_proto_init() }
func file_spaceone_api_identity_v2_external_auth_proto_init() {
	if File_spaceone_api_identity_v2_external_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_identity_v2_external_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetExternalAuthRequest); i {
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
		file_spaceone_api_identity_v2_external_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalAuthRequest); i {
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
		file_spaceone_api_identity_v2_external_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalAuthInfo); i {
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
			RawDescriptor: file_spaceone_api_identity_v2_external_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_identity_v2_external_auth_proto_goTypes,
		DependencyIndexes: file_spaceone_api_identity_v2_external_auth_proto_depIdxs,
		EnumInfos:         file_spaceone_api_identity_v2_external_auth_proto_enumTypes,
		MessageInfos:      file_spaceone_api_identity_v2_external_auth_proto_msgTypes,
	}.Build()
	File_spaceone_api_identity_v2_external_auth_proto = out.File
	file_spaceone_api_identity_v2_external_auth_proto_rawDesc = nil
	file_spaceone_api_identity_v2_external_auth_proto_goTypes = nil
	file_spaceone_api_identity_v2_external_auth_proto_depIdxs = nil
}
