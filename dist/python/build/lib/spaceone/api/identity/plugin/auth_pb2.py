# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/identity/plugin/auth.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\'spaceone/api/identity/plugin/auth.proto\x12\x1cspaceone.api.identity.plugin\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\"7\n\x0bInitRequest\x12(\n\x07options\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\"w\n\rVerifyRequest\x12(\n\x07options\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\x12,\n\x0bsecret_data\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0e\n\x06schema\x18\x03 \x01(\t\"\x97\x01\n\x0b\x46indRequest\x12(\n\x07options\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\x12,\n\x0bsecret_data\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0f\n\x07user_id\x18\x03 \x01(\t\x12\x0f\n\x07keyword\x18\x04 \x01(\t\x12\x0e\n\x06schema\x18\x05 \x01(\t\"\xa9\x01\n\x0cLoginRequest\x12(\n\x07options\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\x12,\n\x0bsecret_data\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x31\n\x10user_credentials\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0e\n\x06schema\x18\x04 \x01(\t\"\xd4\x01\n\x08UserInfo\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05\x65mail\x18\x03 \x01(\t\x12\x0e\n\x06mobile\x18\x04 \x01(\t\x12\r\n\x05group\x18\x05 \x01(\t\x12;\n\x05state\x18\x06 \x01(\x0e\x32,.spaceone.api.identity.plugin.UserInfo.State\">\n\x05State\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\x12\x10\n\x0cUNIDENTIFIED\x10\x03\"g\n\tUsersInfo\x12\x37\n\x07results\x18\x01 \x03(\x0b\x32&.spaceone.api.identity.plugin.UserInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\x12\x0c\n\x04more\x18\x03 \x01(\x08\":\n\x0e\x41uthVerifyInfo\x12(\n\x07options\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\"7\n\nPluginInfo\x12)\n\x08metadata\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct2\xf3\x02\n\x04\x41uth\x12]\n\x04init\x12).spaceone.api.identity.plugin.InitRequest\x1a(.spaceone.api.identity.plugin.PluginInfo\"\x00\x12O\n\x06verify\x12+.spaceone.api.identity.plugin.VerifyRequest\x1a\x16.google.protobuf.Empty\"\x00\x12\\\n\x04\x66ind\x12).spaceone.api.identity.plugin.FindRequest\x1a\'.spaceone.api.identity.plugin.UsersInfo\"\x00\x12]\n\x05login\x12*.spaceone.api.identity.plugin.LoginRequest\x1a&.spaceone.api.identity.plugin.UserInfo\"\x00\x42\x43ZAgithub.com/cloudforet-io/api/dist/go/spaceone/api/identity/pluginb\x06proto3')



_INITREQUEST = DESCRIPTOR.message_types_by_name['InitRequest']
_VERIFYREQUEST = DESCRIPTOR.message_types_by_name['VerifyRequest']
_FINDREQUEST = DESCRIPTOR.message_types_by_name['FindRequest']
_LOGINREQUEST = DESCRIPTOR.message_types_by_name['LoginRequest']
_USERINFO = DESCRIPTOR.message_types_by_name['UserInfo']
_USERSINFO = DESCRIPTOR.message_types_by_name['UsersInfo']
_AUTHVERIFYINFO = DESCRIPTOR.message_types_by_name['AuthVerifyInfo']
_PLUGININFO = DESCRIPTOR.message_types_by_name['PluginInfo']
_USERINFO_STATE = _USERINFO.enum_types_by_name['State']
InitRequest = _reflection.GeneratedProtocolMessageType('InitRequest', (_message.Message,), {
  'DESCRIPTOR' : _INITREQUEST,
  '__module__' : 'spaceone.api.identity.plugin.auth_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.identity.plugin.InitRequest)
  })
_sym_db.RegisterMessage(InitRequest)

VerifyRequest = _reflection.GeneratedProtocolMessageType('VerifyRequest', (_message.Message,), {
  'DESCRIPTOR' : _VERIFYREQUEST,
  '__module__' : 'spaceone.api.identity.plugin.auth_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.identity.plugin.VerifyRequest)
  })
_sym_db.RegisterMessage(VerifyRequest)

FindRequest = _reflection.GeneratedProtocolMessageType('FindRequest', (_message.Message,), {
  'DESCRIPTOR' : _FINDREQUEST,
  '__module__' : 'spaceone.api.identity.plugin.auth_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.identity.plugin.FindRequest)
  })
_sym_db.RegisterMessage(FindRequest)

LoginRequest = _reflection.GeneratedProtocolMessageType('LoginRequest', (_message.Message,), {
  'DESCRIPTOR' : _LOGINREQUEST,
  '__module__' : 'spaceone.api.identity.plugin.auth_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.identity.plugin.LoginRequest)
  })
_sym_db.RegisterMessage(LoginRequest)

UserInfo = _reflection.GeneratedProtocolMessageType('UserInfo', (_message.Message,), {
  'DESCRIPTOR' : _USERINFO,
  '__module__' : 'spaceone.api.identity.plugin.auth_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.identity.plugin.UserInfo)
  })
_sym_db.RegisterMessage(UserInfo)

UsersInfo = _reflection.GeneratedProtocolMessageType('UsersInfo', (_message.Message,), {
  'DESCRIPTOR' : _USERSINFO,
  '__module__' : 'spaceone.api.identity.plugin.auth_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.identity.plugin.UsersInfo)
  })
_sym_db.RegisterMessage(UsersInfo)

AuthVerifyInfo = _reflection.GeneratedProtocolMessageType('AuthVerifyInfo', (_message.Message,), {
  'DESCRIPTOR' : _AUTHVERIFYINFO,
  '__module__' : 'spaceone.api.identity.plugin.auth_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.identity.plugin.AuthVerifyInfo)
  })
_sym_db.RegisterMessage(AuthVerifyInfo)

PluginInfo = _reflection.GeneratedProtocolMessageType('PluginInfo', (_message.Message,), {
  'DESCRIPTOR' : _PLUGININFO,
  '__module__' : 'spaceone.api.identity.plugin.auth_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.identity.plugin.PluginInfo)
  })
_sym_db.RegisterMessage(PluginInfo)

_AUTH = DESCRIPTOR.services_by_name['Auth']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZAgithub.com/cloudforet-io/api/dist/go/spaceone/api/identity/plugin'
  _INITREQUEST._serialized_start=132
  _INITREQUEST._serialized_end=187
  _VERIFYREQUEST._serialized_start=189
  _VERIFYREQUEST._serialized_end=308
  _FINDREQUEST._serialized_start=311
  _FINDREQUEST._serialized_end=462
  _LOGINREQUEST._serialized_start=465
  _LOGINREQUEST._serialized_end=634
  _USERINFO._serialized_start=637
  _USERINFO._serialized_end=849
  _USERINFO_STATE._serialized_start=787
  _USERINFO_STATE._serialized_end=849
  _USERSINFO._serialized_start=851
  _USERSINFO._serialized_end=954
  _AUTHVERIFYINFO._serialized_start=956
  _AUTHVERIFYINFO._serialized_end=1014
  _PLUGININFO._serialized_start=1016
  _PLUGININFO._serialized_end=1071
  _AUTH._serialized_start=1074
  _AUTH._serialized_end=1445
# @@protoc_insertion_point(module_scope)