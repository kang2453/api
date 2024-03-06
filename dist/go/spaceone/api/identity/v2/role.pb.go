// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.6.1
// source: spaceone/api/identity/v2/role.proto

package v2

import (
	v2 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type RoleType int32

const (
	RoleType_NONE             RoleType = 0
	RoleType_DOMAIN_ADMIN     RoleType = 1
	RoleType_WORKSPACE_OWNER  RoleType = 2
	RoleType_WORKSPACE_MEMBER RoleType = 3
)

// Enum value maps for RoleType.
var (
	RoleType_name = map[int32]string{
		0: "NONE",
		1: "DOMAIN_ADMIN",
		2: "WORKSPACE_OWNER",
		3: "WORKSPACE_MEMBER",
	}
	RoleType_value = map[string]int32{
		"NONE":             0,
		"DOMAIN_ADMIN":     1,
		"WORKSPACE_OWNER":  2,
		"WORKSPACE_MEMBER": 3,
	}
)

func (x RoleType) Enum() *RoleType {
	p := new(RoleType)
	*p = x
	return p
}

func (x RoleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoleType) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_role_proto_enumTypes[0].Descriptor()
}

func (RoleType) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_role_proto_enumTypes[0]
}

func (x RoleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoleType.Descriptor instead.
func (RoleType) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_role_proto_rawDescGZIP(), []int{0}
}

type CreateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RoleType RoleType `protobuf:"varint,2,opt,name=role_type,json=roleType,proto3,enum=spaceone.api.identity.v2.RoleType" json:"role_type,omitempty"`
	// +optional
	Permissions []string `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// +optional
	PageAccess []string `protobuf:"bytes,4,rep,name=page_access,json=pageAccess,proto3" json:"page_access,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *CreateRoleRequest) Reset() {
	*x = CreateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRequest) ProtoMessage() {}

func (x *CreateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_role_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRoleRequest) GetRoleType() RoleType {
	if x != nil {
		return x.RoleType
	}
	return RoleType_NONE
}

func (x *CreateRoleRequest) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *CreateRoleRequest) GetPageAccess() []string {
	if x != nil {
		return x.PageAccess
	}
	return nil
}

func (x *CreateRoleRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

type UpdateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Permissions []string `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// +optional
	PageAccess []string `protobuf:"bytes,4,rep,name=page_access,json=pageAccess,proto3" json:"page_access,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *UpdateRoleRequest) Reset() {
	*x = UpdateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleRequest) ProtoMessage() {}

func (x *UpdateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleRequest.ProtoReflect.Descriptor instead.
func (*UpdateRoleRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_role_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRoleRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *UpdateRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRoleRequest) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *UpdateRoleRequest) GetPageAccess() []string {
	if x != nil {
		return x.PageAccess
	}
	return nil
}

func (x *UpdateRoleRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

type RoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
}

func (x *RoleRequest) Reset() {
	*x = RoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleRequest) ProtoMessage() {}

func (x *RoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleRequest.ProtoReflect.Descriptor instead.
func (*RoleRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_role_proto_rawDescGZIP(), []int{2}
}

func (x *RoleRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

type RoleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId      string          `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Name        string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RoleType    RoleType        `protobuf:"varint,3,opt,name=role_type,json=roleType,proto3,enum=spaceone.api.identity.v2.RoleType" json:"role_type,omitempty"`
	Permissions []string        `protobuf:"bytes,4,rep,name=permissions,proto3" json:"permissions,omitempty"`
	PageAccess  []string        `protobuf:"bytes,5,rep,name=page_access,json=pageAccess,proto3" json:"page_access,omitempty"`
	Tags        *_struct.Struct `protobuf:"bytes,6,opt,name=tags,proto3" json:"tags,omitempty"`
	IsManaged   bool            `protobuf:"varint,7,opt,name=is_managed,json=isManaged,proto3" json:"is_managed,omitempty"`
	DomainId    string          `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	CreatedAt   string          `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string          `protobuf:"bytes,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *RoleInfo) Reset() {
	*x = RoleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleInfo) ProtoMessage() {}

func (x *RoleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleInfo.ProtoReflect.Descriptor instead.
func (*RoleInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_role_proto_rawDescGZIP(), []int{3}
}

func (x *RoleInfo) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *RoleInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleInfo) GetRoleType() RoleType {
	if x != nil {
		return x.RoleType
	}
	return RoleType_NONE
}

func (x *RoleInfo) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *RoleInfo) GetPageAccess() []string {
	if x != nil {
		return x.PageAccess
	}
	return nil
}

func (x *RoleInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *RoleInfo) GetIsManaged() bool {
	if x != nil {
		return x.IsManaged
	}
	return false
}

func (x *RoleInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *RoleInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *RoleInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type RoleSearchQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	RoleId string `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	RoleType RoleType `protobuf:"varint,4,opt,name=role_type,json=roleType,proto3,enum=spaceone.api.identity.v2.RoleType" json:"role_type,omitempty"`
}

func (x *RoleSearchQuery) Reset() {
	*x = RoleSearchQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_role_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleSearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleSearchQuery) ProtoMessage() {}

func (x *RoleSearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_role_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleSearchQuery.ProtoReflect.Descriptor instead.
func (*RoleSearchQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_role_proto_rawDescGZIP(), []int{4}
}

func (x *RoleSearchQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *RoleSearchQuery) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *RoleSearchQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleSearchQuery) GetRoleType() RoleType {
	if x != nil {
		return x.RoleType
	}
	return RoleType_NONE
}

type RolesInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results    []*RoleInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount int32       `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *RolesInfo) Reset() {
	*x = RolesInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_role_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RolesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RolesInfo) ProtoMessage() {}

func (x *RolesInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_role_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RolesInfo.ProtoReflect.Descriptor instead.
func (*RolesInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_role_proto_rawDescGZIP(), []int{5}
}

func (x *RolesInfo) GetResults() []*RoleInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *RolesInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type RoleStatQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query *v2.StatisticsQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *RoleStatQuery) Reset() {
	*x = RoleStatQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_role_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleStatQuery) ProtoMessage() {}

func (x *RoleStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_role_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleStatQuery.ProtoReflect.Descriptor instead.
func (*RoleStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_role_proto_rawDescGZIP(), []int{6}
}

func (x *RoleStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_spaceone_api_identity_v2_role_proto protoreflect.FileDescriptor

var file_spaceone_api_identity_v2_role_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x01, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x72, 0x6f, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0xb0, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x26, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64,
	0x22, 0xe2, 0x02, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a,
	0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x72, 0x6f,
	0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2b,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xb2, 0x01, 0x0a, 0x0f, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x72, 0x6f, 0x6c,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x6a, 0x0a, 0x09, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4c, 0x0a, 0x0d, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2a, 0x51, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x4f,
	0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f,
	0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10,
	0x02, 0x12, 0x14, 0x0a, 0x10, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x4d,
	0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x03, 0x32, 0xd0, 0x05, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x7e, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x23, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x7e, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x23, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x6c, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f,
	0x76, 0x32, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x72,
	0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x67,
	0x65, 0x74, 0x12, 0x79, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x23, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2f, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x6b, 0x0a,
	0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a,
	0x01, 0x2a, 0x22, 0x16, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32,
	0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f,
	0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f,
	0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_identity_v2_role_proto_rawDescOnce sync.Once
	file_spaceone_api_identity_v2_role_proto_rawDescData = file_spaceone_api_identity_v2_role_proto_rawDesc
)

func file_spaceone_api_identity_v2_role_proto_rawDescGZIP() []byte {
	file_spaceone_api_identity_v2_role_proto_rawDescOnce.Do(func() {
		file_spaceone_api_identity_v2_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_identity_v2_role_proto_rawDescData)
	})
	return file_spaceone_api_identity_v2_role_proto_rawDescData
}

var file_spaceone_api_identity_v2_role_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spaceone_api_identity_v2_role_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_spaceone_api_identity_v2_role_proto_goTypes = []interface{}{
	(RoleType)(0),              // 0: spaceone.api.identity.v2.RoleType
	(*CreateRoleRequest)(nil),  // 1: spaceone.api.identity.v2.CreateRoleRequest
	(*UpdateRoleRequest)(nil),  // 2: spaceone.api.identity.v2.UpdateRoleRequest
	(*RoleRequest)(nil),        // 3: spaceone.api.identity.v2.RoleRequest
	(*RoleInfo)(nil),           // 4: spaceone.api.identity.v2.RoleInfo
	(*RoleSearchQuery)(nil),    // 5: spaceone.api.identity.v2.RoleSearchQuery
	(*RolesInfo)(nil),          // 6: spaceone.api.identity.v2.RolesInfo
	(*RoleStatQuery)(nil),      // 7: spaceone.api.identity.v2.RoleStatQuery
	(*_struct.Struct)(nil),     // 8: google.protobuf.Struct
	(*v2.Query)(nil),           // 9: spaceone.api.core.v2.Query
	(*v2.StatisticsQuery)(nil), // 10: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),        // 11: google.protobuf.Empty
}
var file_spaceone_api_identity_v2_role_proto_depIdxs = []int32{
	0,  // 0: spaceone.api.identity.v2.CreateRoleRequest.role_type:type_name -> spaceone.api.identity.v2.RoleType
	8,  // 1: spaceone.api.identity.v2.CreateRoleRequest.tags:type_name -> google.protobuf.Struct
	8,  // 2: spaceone.api.identity.v2.UpdateRoleRequest.tags:type_name -> google.protobuf.Struct
	0,  // 3: spaceone.api.identity.v2.RoleInfo.role_type:type_name -> spaceone.api.identity.v2.RoleType
	8,  // 4: spaceone.api.identity.v2.RoleInfo.tags:type_name -> google.protobuf.Struct
	9,  // 5: spaceone.api.identity.v2.RoleSearchQuery.query:type_name -> spaceone.api.core.v2.Query
	0,  // 6: spaceone.api.identity.v2.RoleSearchQuery.role_type:type_name -> spaceone.api.identity.v2.RoleType
	4,  // 7: spaceone.api.identity.v2.RolesInfo.results:type_name -> spaceone.api.identity.v2.RoleInfo
	10, // 8: spaceone.api.identity.v2.RoleStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	1,  // 9: spaceone.api.identity.v2.Role.create:input_type -> spaceone.api.identity.v2.CreateRoleRequest
	2,  // 10: spaceone.api.identity.v2.Role.update:input_type -> spaceone.api.identity.v2.UpdateRoleRequest
	3,  // 11: spaceone.api.identity.v2.Role.delete:input_type -> spaceone.api.identity.v2.RoleRequest
	3,  // 12: spaceone.api.identity.v2.Role.get:input_type -> spaceone.api.identity.v2.RoleRequest
	5,  // 13: spaceone.api.identity.v2.Role.list:input_type -> spaceone.api.identity.v2.RoleSearchQuery
	7,  // 14: spaceone.api.identity.v2.Role.stat:input_type -> spaceone.api.identity.v2.RoleStatQuery
	4,  // 15: spaceone.api.identity.v2.Role.create:output_type -> spaceone.api.identity.v2.RoleInfo
	4,  // 16: spaceone.api.identity.v2.Role.update:output_type -> spaceone.api.identity.v2.RoleInfo
	11, // 17: spaceone.api.identity.v2.Role.delete:output_type -> google.protobuf.Empty
	4,  // 18: spaceone.api.identity.v2.Role.get:output_type -> spaceone.api.identity.v2.RoleInfo
	6,  // 19: spaceone.api.identity.v2.Role.list:output_type -> spaceone.api.identity.v2.RolesInfo
	8,  // 20: spaceone.api.identity.v2.Role.stat:output_type -> google.protobuf.Struct
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_spaceone_api_identity_v2_role_proto_init() }
func file_spaceone_api_identity_v2_role_proto_init() {
	if File_spaceone_api_identity_v2_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_identity_v2_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleRequest); i {
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
		file_spaceone_api_identity_v2_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoleRequest); i {
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
		file_spaceone_api_identity_v2_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleRequest); i {
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
		file_spaceone_api_identity_v2_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleInfo); i {
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
		file_spaceone_api_identity_v2_role_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleSearchQuery); i {
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
		file_spaceone_api_identity_v2_role_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RolesInfo); i {
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
		file_spaceone_api_identity_v2_role_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleStatQuery); i {
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
			RawDescriptor: file_spaceone_api_identity_v2_role_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_identity_v2_role_proto_goTypes,
		DependencyIndexes: file_spaceone_api_identity_v2_role_proto_depIdxs,
		EnumInfos:         file_spaceone_api_identity_v2_role_proto_enumTypes,
		MessageInfos:      file_spaceone_api_identity_v2_role_proto_msgTypes,
	}.Build()
	File_spaceone_api_identity_v2_role_proto = out.File
	file_spaceone_api_identity_v2_role_proto_rawDesc = nil
	file_spaceone_api_identity_v2_role_proto_goTypes = nil
	file_spaceone_api_identity_v2_role_proto_depIdxs = nil
}
