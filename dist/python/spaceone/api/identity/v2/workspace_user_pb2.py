# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/identity/v2/workspace_user.proto
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
from spaceone.api.identity.v2 import user_pb2 as spaceone_dot_api_dot_identity_dot_v2_dot_user__pb2
from spaceone.api.identity.v2 import role_binding_pb2 as spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n-spaceone/api/identity/v2/workspace_user.proto\x12\x18spaceone.api.identity.v2\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\x1a#spaceone/api/identity/v2/user.proto\x1a+spaceone/api/identity/v2/role_binding.proto\"\xb0\x02\n\x1a\x43reateWorkspaceUserRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\r\n\x05\x65mail\x18\x04 \x01(\t\x12\x35\n\tauth_type\x18\x05 \x01(\x0e\x32\".spaceone.api.identity.v2.AuthType\x12\x10\n\x08language\x18\x06 \x01(\t\x12\x10\n\x08timezone\x18\x07 \x01(\t\x12%\n\x04tags\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x16\n\x0ereset_password\x18\t \x01(\x08\x12\x0f\n\x07role_id\x18\n \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\"P\n\x14WorkspaceUserRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\"\xa1\x05\n\x11WorkspaceUserInfo\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12@\n\x05state\x18\x03 \x01(\x0e\x32\x31.spaceone.api.identity.v2.WorkspaceUserInfo.State\x12\r\n\x05\x65mail\x18\x04 \x01(\t\x12\x35\n\tauth_type\x18\x05 \x01(\x0e\x32\".spaceone.api.identity.v2.AuthType\x12G\n\trole_type\x18\x06 \x01(\x0e\x32\x34.spaceone.api.identity.v2.WorkspaceUserInfo.RoleType\x12\x10\n\x08language\x18\x07 \x01(\t\x12\x10\n\x08timezone\x18\x08 \x01(\t\x12\x15\n\rapi_key_count\x18\t \x01(\x05\x12%\n\x04tags\x18\n \x01(\x0b\x32\x17.google.protobuf.Struct\x12?\n\x0crole_binding\x18\x0b \x01(\x0b\x32).spaceone.api.identity.v2.RoleBindingInfo\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x18\n\x10last_accessed_at\x18  \x01(\t\"?\n\x05State\x12\x0e\n\nSTATE_NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\x12\x0b\n\x07PENDING\x10\x03\"w\n\x08RoleType\x12\x12\n\x0eROLE_TYPE_NONE\x10\x00\x12\x10\n\x0cSYSTEM_ADMIN\x10\x01\x12\x10\n\x0c\x44OMAIN_ADMIN\x10\x02\x12\x13\n\x0fWORKSPACE_OWNER\x10\x03\x12\x14\n\x10WORKSPACE_MEMBER\x10\x04\x12\x08\n\x04USER\x10\x05\"\xd8\x02\n\x18WorkspaceUserSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x0f\n\x07user_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12G\n\x05state\x18\x04 \x01(\x0e\x32\x38.spaceone.api.identity.v2.WorkspaceUserSearchQuery.State\x12\r\n\x05\x65mail\x18\x05 \x01(\t\x12\x35\n\tauth_type\x18\x06 \x01(\x0e\x32\".spaceone.api.identity.v2.AuthType\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\"9\n\x05State\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\x12\x0b\n\x07PENDING\x10\x03\"g\n\x12WorkspaceUsersInfo\x12<\n\x07results\x18\x01 \x03(\x0b\x32+.spaceone.api.identity.v2.WorkspaceUserInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"w\n\x16WorkspaceUserStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t2\xd5\x04\n\rWorkspaceUser\x12\x9a\x01\n\x06\x63reate\x12\x34.spaceone.api.identity.v2.CreateWorkspaceUserRequest\x1a+.spaceone.api.identity.v2.WorkspaceUserInfo\"-\x82\xd3\xe4\x93\x02\'\"\"/identity/v2/workspace-user/create:\x01*\x12\x8e\x01\n\x03get\x12..spaceone.api.identity.v2.WorkspaceUserRequest\x1a+.spaceone.api.identity.v2.WorkspaceUserInfo\"*\x82\xd3\xe4\x93\x02$\"\x1f/identity/v2/workspace-user/get:\x01*\x12\x95\x01\n\x04list\x12\x32.spaceone.api.identity.v2.WorkspaceUserSearchQuery\x1a,.spaceone.api.identity.v2.WorkspaceUsersInfo\"+\x82\xd3\xe4\x93\x02%\" /identity/v2/workspace-user/list:\x01*\x12~\n\x04stat\x12\x30.spaceone.api.identity.v2.WorkspaceUserStatQuery\x1a\x17.google.protobuf.Struct\"+\x82\xd3\xe4\x93\x02%\" /identity/v1/workspace-user/stat:\x01*B?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.identity.v2.workspace_user_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2'
  _WORKSPACEUSER.methods_by_name['create']._options = None
  _WORKSPACEUSER.methods_by_name['create']._serialized_options = b'\202\323\344\223\002\'\"\"/identity/v2/workspace-user/create:\001*'
  _WORKSPACEUSER.methods_by_name['get']._options = None
  _WORKSPACEUSER.methods_by_name['get']._serialized_options = b'\202\323\344\223\002$\"\037/identity/v2/workspace-user/get:\001*'
  _WORKSPACEUSER.methods_by_name['list']._options = None
  _WORKSPACEUSER.methods_by_name['list']._serialized_options = b'\202\323\344\223\002%\" /identity/v2/workspace-user/list:\001*'
  _WORKSPACEUSER.methods_by_name['stat']._options = None
  _WORKSPACEUSER.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002%\" /identity/v1/workspace-user/stat:\001*'
  _globals['_CREATEWORKSPACEUSERREQUEST']._serialized_start=281
  _globals['_CREATEWORKSPACEUSERREQUEST']._serialized_end=585
  _globals['_WORKSPACEUSERREQUEST']._serialized_start=587
  _globals['_WORKSPACEUSERREQUEST']._serialized_end=667
  _globals['_WORKSPACEUSERINFO']._serialized_start=670
  _globals['_WORKSPACEUSERINFO']._serialized_end=1343
  _globals['_WORKSPACEUSERINFO_STATE']._serialized_start=1159
  _globals['_WORKSPACEUSERINFO_STATE']._serialized_end=1222
  _globals['_WORKSPACEUSERINFO_ROLETYPE']._serialized_start=1224
  _globals['_WORKSPACEUSERINFO_ROLETYPE']._serialized_end=1343
  _globals['_WORKSPACEUSERSEARCHQUERY']._serialized_start=1346
  _globals['_WORKSPACEUSERSEARCHQUERY']._serialized_end=1690
  _globals['_WORKSPACEUSERSEARCHQUERY_STATE']._serialized_start=1633
  _globals['_WORKSPACEUSERSEARCHQUERY_STATE']._serialized_end=1690
  _globals['_WORKSPACEUSERSINFO']._serialized_start=1692
  _globals['_WORKSPACEUSERSINFO']._serialized_end=1795
  _globals['_WORKSPACEUSERSTATQUERY']._serialized_start=1797
  _globals['_WORKSPACEUSERSTATQUERY']._serialized_end=1916
  _globals['_WORKSPACEUSER']._serialized_start=1919
  _globals['_WORKSPACEUSER']._serialized_end=2516
# @@protoc_insertion_point(module_scope)
