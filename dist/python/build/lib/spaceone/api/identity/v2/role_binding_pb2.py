# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/identity/v2/role_binding.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from spaceone.api.core.v2 import query_pb2 as spaceone_dot_api_dot_core_dot_v2_dot_query__pb2
from spaceone.api.identity.v2 import project_pb2 as spaceone_dot_api_dot_identity_dot_v2_dot_project__pb2
from spaceone.api.identity.v2 import project_group_pb2 as spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2
from spaceone.api.identity.v2 import role_pb2 as spaceone_dot_api_dot_identity_dot_v2_dot_role__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n+spaceone/api/identity/v2/role_binding.proto\x12\x18spaceone.api.identity.v2\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\x1a&spaceone/api/identity/v2/project.proto\x1a,spaceone/api/identity/v2/project_group.proto\x1a#spaceone/api/identity/v2/role.proto\"\xfb\x01\n\x18\x43reateRoleBindingRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x0f\n\x07role_id\x18\x02 \x01(\t\x12\\\n\x10permission_group\x18\x03 \x01(\x0e\x32\x42.spaceone.api.identity.v2.CreateRoleBindingRequest.PermissionGroup\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\"6\n\x0fPermissionGroup\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\r\n\tWORKSPACE\x10\x02\"m\n\x18UpdateRoleBindingRequest\x12\x17\n\x0frole_binding_id\x18\x01 \x01(\t\x12\x0f\n\x07role_id\x18\x02 \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\"V\n\x12RoleBindingRequest\x12\x17\n\x0frole_binding_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\"\xde\x03\n\x0fRoleBindingInfo\x12\x17\n\x0frole_binding_id\x18\x01 \x01(\t\x12\x45\n\trole_type\x18\x02 \x01(\x0e\x32\x32.spaceone.api.identity.v2.RoleBindingInfo.RoleType\x12S\n\x10permission_group\x18\x03 \x01(\x0e\x32\x39.spaceone.api.identity.v2.RoleBindingInfo.PermissionGroup\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x0f\n\x07role_id\x18\x17 \x01(\t\x12\x0f\n\x07user_id\x18\x18 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\"<\n\x0fPermissionGroup\x12\x0e\n\nGROUP_NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\r\n\tWORKSPACE\x10\x02\"y\n\x08RoleType\x12\x12\n\x0eROLE_TYPE_NONE\x10\x00\x12\n\n\x06SYSTEM\x10\x01\x12\x10\n\x0cSYSTEM_ADMIN\x10\x02\x12\x10\n\x0c\x44OMAIN_ADMIN\x10\x03\x12\x13\n\x0fWORKSPACE_OWNER\x10\x04\x12\x14\n\x10WORKSPACE_MEMBER\x10\x05\"\x8b\x04\n\x16RoleBindingSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x17\n\x0frole_binding_id\x18\x02 \x01(\t\x12L\n\trole_type\x18\x03 \x01(\x0e\x32\x39.spaceone.api.identity.v2.RoleBindingSearchQuery.RoleType\x12Z\n\x10permission_group\x18\x04 \x01(\x0e\x32@.spaceone.api.identity.v2.RoleBindingSearchQuery.PermissionGroup\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x0f\n\x07role_id\x18\x17 \x01(\t\x12\x0f\n\x07user_id\x18\x18 \x01(\t\"<\n\x0fPermissionGroup\x12\x0e\n\nGROUP_NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\r\n\tWORKSPACE\x10\x02\"y\n\x08RoleType\x12\x12\n\x0eROLE_TYPE_NONE\x10\x00\x12\n\n\x06SYSTEM\x10\x01\x12\x10\n\x0cSYSTEM_ADMIN\x10\x02\x12\x10\n\x0c\x44OMAIN_ADMIN\x10\x03\x12\x13\n\x0fWORKSPACE_OWNER\x10\x04\x12\x14\n\x10WORKSPACE_MEMBER\x10\x05\"c\n\x10RoleBindingsInfo\x12:\n\x07results\x18\x01 \x03(\x0b\x32).spaceone.api.identity.v2.RoleBindingInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"u\n\x14RoleBindingStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t2\xdb\x06\n\x0bRoleBinding\x12\x94\x01\n\x06\x63reate\x12\x32.spaceone.api.identity.v2.CreateRoleBindingRequest\x1a).spaceone.api.identity.v2.RoleBindingInfo\"+\x82\xd3\xe4\x93\x02%\" /identity/v1/role-binding/create:\x01*\x12\x9e\x01\n\x0bupdate_role\x12\x32.spaceone.api.identity.v2.UpdateRoleBindingRequest\x1a).spaceone.api.identity.v2.RoleBindingInfo\"0\x82\xd3\xe4\x93\x02*\"%/identity/v1/role-binding/update-role:\x01*\x12{\n\x06\x64\x65lete\x12,.spaceone.api.identity.v2.RoleBindingRequest\x1a\x16.google.protobuf.Empty\"+\x82\xd3\xe4\x93\x02%\" /identity/v1/role-binding/delete:\x01*\x12\x88\x01\n\x03get\x12,.spaceone.api.identity.v2.RoleBindingRequest\x1a).spaceone.api.identity.v2.RoleBindingInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/identity/v1/role-binding/get:\x01*\x12\x8f\x01\n\x04list\x12\x30.spaceone.api.identity.v2.RoleBindingSearchQuery\x1a*.spaceone.api.identity.v2.RoleBindingsInfo\")\x82\xd3\xe4\x93\x02#\"\x1e/identity/v1/role-binding/list:\x01*\x12z\n\x04stat\x12..spaceone.api.identity.v2.RoleBindingStatQuery\x1a\x17.google.protobuf.Struct\")\x82\xd3\xe4\x93\x02#\"\x1e/identity/v1/role-binding/stat:\x01*B?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.identity.v2.role_binding_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2'
  _ROLEBINDING.methods_by_name['create']._options = None
  _ROLEBINDING.methods_by_name['create']._serialized_options = b'\202\323\344\223\002%\" /identity/v1/role-binding/create:\001*'
  _ROLEBINDING.methods_by_name['update_role']._options = None
  _ROLEBINDING.methods_by_name['update_role']._serialized_options = b'\202\323\344\223\002*\"%/identity/v1/role-binding/update-role:\001*'
  _ROLEBINDING.methods_by_name['delete']._options = None
  _ROLEBINDING.methods_by_name['delete']._serialized_options = b'\202\323\344\223\002%\" /identity/v1/role-binding/delete:\001*'
  _ROLEBINDING.methods_by_name['get']._options = None
  _ROLEBINDING.methods_by_name['get']._serialized_options = b'\202\323\344\223\002\"\"\035/identity/v1/role-binding/get:\001*'
  _ROLEBINDING.methods_by_name['list']._options = None
  _ROLEBINDING.methods_by_name['list']._serialized_options = b'\202\323\344\223\002#\"\036/identity/v1/role-binding/list:\001*'
  _ROLEBINDING.methods_by_name['stat']._options = None
  _ROLEBINDING.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002#\"\036/identity/v1/role-binding/stat:\001*'
  _globals['_CREATEROLEBINDINGREQUEST']._serialized_start=320
  _globals['_CREATEROLEBINDINGREQUEST']._serialized_end=571
  _globals['_CREATEROLEBINDINGREQUEST_PERMISSIONGROUP']._serialized_start=517
  _globals['_CREATEROLEBINDINGREQUEST_PERMISSIONGROUP']._serialized_end=571
  _globals['_UPDATEROLEBINDINGREQUEST']._serialized_start=573
  _globals['_UPDATEROLEBINDINGREQUEST']._serialized_end=682
  _globals['_ROLEBINDINGREQUEST']._serialized_start=684
  _globals['_ROLEBINDINGREQUEST']._serialized_end=770
  _globals['_ROLEBINDINGINFO']._serialized_start=773
  _globals['_ROLEBINDINGINFO']._serialized_end=1251
  _globals['_ROLEBINDINGINFO_PERMISSIONGROUP']._serialized_start=1068
  _globals['_ROLEBINDINGINFO_PERMISSIONGROUP']._serialized_end=1128
  _globals['_ROLEBINDINGINFO_ROLETYPE']._serialized_start=1130
  _globals['_ROLEBINDINGINFO_ROLETYPE']._serialized_end=1251
  _globals['_ROLEBINDINGSEARCHQUERY']._serialized_start=1254
  _globals['_ROLEBINDINGSEARCHQUERY']._serialized_end=1777
  _globals['_ROLEBINDINGSEARCHQUERY_PERMISSIONGROUP']._serialized_start=1068
  _globals['_ROLEBINDINGSEARCHQUERY_PERMISSIONGROUP']._serialized_end=1128
  _globals['_ROLEBINDINGSEARCHQUERY_ROLETYPE']._serialized_start=1130
  _globals['_ROLEBINDINGSEARCHQUERY_ROLETYPE']._serialized_end=1251
  _globals['_ROLEBINDINGSINFO']._serialized_start=1779
  _globals['_ROLEBINDINGSINFO']._serialized_end=1878
  _globals['_ROLEBINDINGSTATQUERY']._serialized_start=1880
  _globals['_ROLEBINDINGSTATQUERY']._serialized_end=1997
  _globals['_ROLEBINDING']._serialized_start=2000
  _globals['_ROLEBINDING']._serialized_end=2859
# @@protoc_insertion_point(module_scope)
