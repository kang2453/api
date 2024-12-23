// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.6.1
// source: spaceone/api/alert_manager/plugin/event.proto

package plugin

import (
	v1 "github.com/cloudforet-io/api/dist/go/spaceone/api/alert-manager/v1"
	_ "github.com/golang/protobuf/ptypes/empty"
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

type EventType int32

const (
	EventType_EVENT_TYPE_NONE EventType = 0
	EventType_RECOVERY        EventType = 1
	EventType_ALERT           EventType = 2
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "EVENT_TYPE_NONE",
		1: "RECOVERY",
		2: "ALERT",
	}
	EventType_value = map[string]int32{
		"EVENT_TYPE_NONE": 0,
		"RECOVERY":        1,
		"ALERT":           2,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_alert_manager_plugin_event_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_spaceone_api_alert_manager_plugin_event_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_plugin_event_proto_rawDescGZIP(), []int{0}
}

type ParseRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Options       *_struct.Struct        `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	Data          *_struct.Struct        `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	DomainId      string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ParseRequest) Reset() {
	*x = ParseRequest{}
	mi := &file_spaceone_api_alert_manager_plugin_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseRequest) ProtoMessage() {}

func (x *ParseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_alert_manager_plugin_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseRequest.ProtoReflect.Descriptor instead.
func (*ParseRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_plugin_event_proto_rawDescGZIP(), []int{0}
}

func (x *ParseRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *ParseRequest) GetData() *_struct.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ParseRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type EventResource struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// +optional
	ResourceType string `protobuf:"bytes,2,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// +optional
	Name          string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventResource) Reset() {
	*x = EventResource{}
	mi := &file_spaceone_api_alert_manager_plugin_event_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResource) ProtoMessage() {}

func (x *EventResource) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_alert_manager_plugin_event_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResource.ProtoReflect.Descriptor instead.
func (*EventResource) Descriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_plugin_event_proto_rawDescGZIP(), []int{1}
}

func (x *EventResource) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *EventResource) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *EventResource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type EventInfo struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	EventKey       string                 `protobuf:"bytes,1,opt,name=event_key,json=eventKey,proto3" json:"event_key,omitempty"`
	EventType      EventType              `protobuf:"varint,2,opt,name=event_type,json=eventType,proto3,enum=spaceone.api.alert_manager.plugin.EventType" json:"event_type,omitempty"`
	Title          string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description    string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Severity       v1.EventSeverity       `protobuf:"varint,5,opt,name=severity,proto3,enum=spaceone.api.alert_manager.v1.EventSeverity" json:"severity,omitempty"`
	Rule           string                 `protobuf:"bytes,6,opt,name=rule,proto3" json:"rule,omitempty"`
	ImageUrl       string                 `protobuf:"bytes,7,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Resources      []string               `protobuf:"bytes,8,rep,name=resources,proto3" json:"resources,omitempty"`
	Provider       string                 `protobuf:"bytes,9,opt,name=provider,proto3" json:"provider,omitempty"`
	Account        string                 `protobuf:"bytes,10,opt,name=account,proto3" json:"account,omitempty"`
	AdditionalInfo *_struct.Struct        `protobuf:"bytes,11,opt,name=additional_info,json=additionalInfo,proto3" json:"additional_info,omitempty"`
	OccurredAt     string                 `protobuf:"bytes,31,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *EventInfo) Reset() {
	*x = EventInfo{}
	mi := &file_spaceone_api_alert_manager_plugin_event_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventInfo) ProtoMessage() {}

func (x *EventInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_alert_manager_plugin_event_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventInfo.ProtoReflect.Descriptor instead.
func (*EventInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_plugin_event_proto_rawDescGZIP(), []int{2}
}

func (x *EventInfo) GetEventKey() string {
	if x != nil {
		return x.EventKey
	}
	return ""
}

func (x *EventInfo) GetEventType() EventType {
	if x != nil {
		return x.EventType
	}
	return EventType_EVENT_TYPE_NONE
}

func (x *EventInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EventInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *EventInfo) GetSeverity() v1.EventSeverity {
	if x != nil {
		return x.Severity
	}
	return v1.EventSeverity(0)
}

func (x *EventInfo) GetRule() string {
	if x != nil {
		return x.Rule
	}
	return ""
}

func (x *EventInfo) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *EventInfo) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *EventInfo) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *EventInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *EventInfo) GetAdditionalInfo() *_struct.Struct {
	if x != nil {
		return x.AdditionalInfo
	}
	return nil
}

func (x *EventInfo) GetOccurredAt() string {
	if x != nil {
		return x.OccurredAt
	}
	return ""
}

type EventsInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*EventInfo           `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventsInfo) Reset() {
	*x = EventsInfo{}
	mi := &file_spaceone_api_alert_manager_plugin_event_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsInfo) ProtoMessage() {}

func (x *EventsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_alert_manager_plugin_event_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsInfo.ProtoReflect.Descriptor instead.
func (*EventsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_plugin_event_proto_rawDescGZIP(), []int{3}
}

func (x *EventsInfo) GetResults() []*EventInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_spaceone_api_alert_manager_plugin_event_proto protoreflect.FileDescriptor

var file_spaceone_api_alert_manager_plugin_event_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x0c, 0x50, 0x61, 0x72,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x69, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0xdf, 0x03, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1b, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x4b, 0x0a, 0x0a,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x48, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x75, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x40, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x0e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x54, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x46, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2a, 0x39, 0x0a, 0x09, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52,
	0x45, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x45,
	0x52, 0x54, 0x10, 0x02, 0x32, 0x72, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x69, 0x0a,
	0x05, 0x70, 0x61, 0x72, 0x73, 0x65, 0x12, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65,
	0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_alert_manager_plugin_event_proto_rawDescOnce sync.Once
	file_spaceone_api_alert_manager_plugin_event_proto_rawDescData = file_spaceone_api_alert_manager_plugin_event_proto_rawDesc
)

func file_spaceone_api_alert_manager_plugin_event_proto_rawDescGZIP() []byte {
	file_spaceone_api_alert_manager_plugin_event_proto_rawDescOnce.Do(func() {
		file_spaceone_api_alert_manager_plugin_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_alert_manager_plugin_event_proto_rawDescData)
	})
	return file_spaceone_api_alert_manager_plugin_event_proto_rawDescData
}

var file_spaceone_api_alert_manager_plugin_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spaceone_api_alert_manager_plugin_event_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spaceone_api_alert_manager_plugin_event_proto_goTypes = []any{
	(EventType)(0),         // 0: spaceone.api.alert_manager.plugin.EventType
	(*ParseRequest)(nil),   // 1: spaceone.api.alert_manager.plugin.ParseRequest
	(*EventResource)(nil),  // 2: spaceone.api.alert_manager.plugin.EventResource
	(*EventInfo)(nil),      // 3: spaceone.api.alert_manager.plugin.EventInfo
	(*EventsInfo)(nil),     // 4: spaceone.api.alert_manager.plugin.EventsInfo
	(*_struct.Struct)(nil), // 5: google.protobuf.Struct
	(v1.EventSeverity)(0),  // 6: spaceone.api.alert_manager.v1.EventSeverity
}
var file_spaceone_api_alert_manager_plugin_event_proto_depIdxs = []int32{
	5, // 0: spaceone.api.alert_manager.plugin.ParseRequest.options:type_name -> google.protobuf.Struct
	5, // 1: spaceone.api.alert_manager.plugin.ParseRequest.data:type_name -> google.protobuf.Struct
	0, // 2: spaceone.api.alert_manager.plugin.EventInfo.event_type:type_name -> spaceone.api.alert_manager.plugin.EventType
	6, // 3: spaceone.api.alert_manager.plugin.EventInfo.severity:type_name -> spaceone.api.alert_manager.v1.EventSeverity
	5, // 4: spaceone.api.alert_manager.plugin.EventInfo.additional_info:type_name -> google.protobuf.Struct
	3, // 5: spaceone.api.alert_manager.plugin.EventsInfo.results:type_name -> spaceone.api.alert_manager.plugin.EventInfo
	1, // 6: spaceone.api.alert_manager.plugin.Event.parse:input_type -> spaceone.api.alert_manager.plugin.ParseRequest
	4, // 7: spaceone.api.alert_manager.plugin.Event.parse:output_type -> spaceone.api.alert_manager.plugin.EventsInfo
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_spaceone_api_alert_manager_plugin_event_proto_init() }
func file_spaceone_api_alert_manager_plugin_event_proto_init() {
	if File_spaceone_api_alert_manager_plugin_event_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_alert_manager_plugin_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_alert_manager_plugin_event_proto_goTypes,
		DependencyIndexes: file_spaceone_api_alert_manager_plugin_event_proto_depIdxs,
		EnumInfos:         file_spaceone_api_alert_manager_plugin_event_proto_enumTypes,
		MessageInfos:      file_spaceone_api_alert_manager_plugin_event_proto_msgTypes,
	}.Build()
	File_spaceone_api_alert_manager_plugin_event_proto = out.File
	file_spaceone_api_alert_manager_plugin_event_proto_rawDesc = nil
	file_spaceone_api_alert_manager_plugin_event_proto_goTypes = nil
	file_spaceone_api_alert_manager_plugin_event_proto_depIdxs = nil
}
