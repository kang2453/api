# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/config/v1/public_config.proto
# Protobuf Python Version: 5.26.1
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n*spaceone/api/config/v1/public_config.proto\x12\x16spaceone.api.config.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\xcc\x02\n\x19\x43reatePublicConfigRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12%\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12W\n\x0eresource_group\x18\x14 \x01(\x0e\x32?.spaceone.api.config.v1.CreatePublicConfigRequest.ResourceGroup\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\x12\x12\n\nproject_id\x18\x16 \x01(\t\"P\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\r\n\tWORKSPACE\x10\x02\x12\x0b\n\x07PROJECT\x10\x03\"\xa1\x01\n\x19UpdatePublicConfigRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12%\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\x12\x12\n\nproject_id\x18\x16 \x01(\t\"M\n\x13PublicConfigRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\x12\x12\n\nproject_id\x18\x16 \x01(\t\"S\n\x17PublicConfigSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x0c\n\x04name\x18\x02 \x01(\t\"\xf5\x02\n\x10PublicConfigInfo\x12\x0c\n\x04name\x18\x01 \x01(\t\x12%\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12N\n\x0eresource_group\x18\x14 \x01(\x0e\x32\x36.spaceone.api.config.v1.PublicConfigInfo.ResourceGroup\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\nproject_id\x18\x17 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"P\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\r\n\tWORKSPACE\x10\x02\x12\x0b\n\x07PROJECT\x10\x03\"c\n\x11PublicConfigsInfo\x12\x39\n\x07results\x18\x01 \x03(\x0b\x32(.spaceone.api.config.v1.PublicConfigInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"M\n\x15PublicConfigStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\xc2\x06\n\x0cPublicConfig\x12\x91\x01\n\x06\x63reate\x12\x31.spaceone.api.config.v1.CreatePublicConfigRequest\x1a(.spaceone.api.config.v1.PublicConfigInfo\"*\x82\xd3\xe4\x93\x02$\"\x1f/config/v1/public-config/create:\x01*\x12\x91\x01\n\x06update\x12\x31.spaceone.api.config.v1.UpdatePublicConfigRequest\x1a(.spaceone.api.config.v1.PublicConfigInfo\"*\x82\xd3\xe4\x93\x02$\"\x1f/config/v1/public-config/update:\x01*\x12y\n\x06\x64\x65lete\x12+.spaceone.api.config.v1.PublicConfigRequest\x1a\x16.google.protobuf.Empty\"*\x82\xd3\xe4\x93\x02$\"\x1f/config/v1/public-config/delete:\x01*\x12\x85\x01\n\x03get\x12+.spaceone.api.config.v1.PublicConfigRequest\x1a(.spaceone.api.config.v1.PublicConfigInfo\"\'\x82\xd3\xe4\x93\x02!\"\x1c/config/v1/public-config/get:\x01*\x12\x8c\x01\n\x04list\x12/.spaceone.api.config.v1.PublicConfigSearchQuery\x1a).spaceone.api.config.v1.PublicConfigsInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/config/v1/public-config/list:\x01*\x12x\n\x04stat\x12-.spaceone.api.config.v1.PublicConfigStatQuery\x1a\x17.google.protobuf.Struct\"(\x82\xd3\xe4\x93\x02\"\"\x1d/config/v1/public-config/stat:\x01*B=Z;github.com/cloudforet-io/api/dist/go/spaceone/api/config/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.config.v1.public_config_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z;github.com/cloudforet-io/api/dist/go/spaceone/api/config/v1'
  _globals['_PUBLICCONFIG'].methods_by_name['create']._loaded_options = None
  _globals['_PUBLICCONFIG'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002$\"\037/config/v1/public-config/create:\001*'
  _globals['_PUBLICCONFIG'].methods_by_name['update']._loaded_options = None
  _globals['_PUBLICCONFIG'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002$\"\037/config/v1/public-config/update:\001*'
  _globals['_PUBLICCONFIG'].methods_by_name['delete']._loaded_options = None
  _globals['_PUBLICCONFIG'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002$\"\037/config/v1/public-config/delete:\001*'
  _globals['_PUBLICCONFIG'].methods_by_name['get']._loaded_options = None
  _globals['_PUBLICCONFIG'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002!\"\034/config/v1/public-config/get:\001*'
  _globals['_PUBLICCONFIG'].methods_by_name['list']._loaded_options = None
  _globals['_PUBLICCONFIG'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002\"\"\035/config/v1/public-config/list:\001*'
  _globals['_PUBLICCONFIG'].methods_by_name['stat']._loaded_options = None
  _globals['_PUBLICCONFIG'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\"\"\035/config/v1/public-config/stat:\001*'
  _globals['_CREATEPUBLICCONFIGREQUEST']._serialized_start=194
  _globals['_CREATEPUBLICCONFIGREQUEST']._serialized_end=526
  _globals['_CREATEPUBLICCONFIGREQUEST_RESOURCEGROUP']._serialized_start=446
  _globals['_CREATEPUBLICCONFIGREQUEST_RESOURCEGROUP']._serialized_end=526
  _globals['_UPDATEPUBLICCONFIGREQUEST']._serialized_start=529
  _globals['_UPDATEPUBLICCONFIGREQUEST']._serialized_end=690
  _globals['_PUBLICCONFIGREQUEST']._serialized_start=692
  _globals['_PUBLICCONFIGREQUEST']._serialized_end=769
  _globals['_PUBLICCONFIGSEARCHQUERY']._serialized_start=771
  _globals['_PUBLICCONFIGSEARCHQUERY']._serialized_end=854
  _globals['_PUBLICCONFIGINFO']._serialized_start=857
  _globals['_PUBLICCONFIGINFO']._serialized_end=1230
  _globals['_PUBLICCONFIGINFO_RESOURCEGROUP']._serialized_start=446
  _globals['_PUBLICCONFIGINFO_RESOURCEGROUP']._serialized_end=526
  _globals['_PUBLICCONFIGSINFO']._serialized_start=1232
  _globals['_PUBLICCONFIGSINFO']._serialized_end=1331
  _globals['_PUBLICCONFIGSTATQUERY']._serialized_start=1333
  _globals['_PUBLICCONFIGSTATQUERY']._serialized_end=1410
  _globals['_PUBLICCONFIG']._serialized_start=1413
  _globals['_PUBLICCONFIG']._serialized_end=2247
# @@protoc_insertion_point(module_scope)
