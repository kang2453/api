# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/identity/v2/schema.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n%spaceone/api/identity/v2/schema.proto\x12\x18spaceone.api.identity.v2\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\xa9\x02\n\x13\x43reateSchemaRequest\x12\x11\n\tschema_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x39\n\x0bschema_type\x18\x03 \x01(\x0e\x32$.spaceone.api.identity.v2.SchemaType\x12\'\n\x06schema\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x10\n\x08provider\x18\x05 \x01(\t\x12\x17\n\x0frelated_schemas\x18\x06 \x03(\t\x12(\n\x07options\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"\xdc\x01\n\x13UpdateSchemaRequest\x12\x11\n\tschema_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\'\n\x06schema\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x17\n\x0frelated_schemas\x18\x04 \x03(\t\x12(\n\x07options\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"5\n\rSchemaRequest\x12\x11\n\tschema_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"\xdc\x02\n\nSchemaInfo\x12\x11\n\tschema_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x39\n\x0bschema_type\x18\x03 \x01(\x0e\x32$.spaceone.api.identity.v2.SchemaType\x12\'\n\x06schema\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x10\n\x08provider\x18\x05 \x01(\t\x12\x17\n\x0frelated_schemas\x18\x06 \x03(\t\x12(\n\x07options\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\nis_managed\x18\t \x01(\x08\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"\xef\x01\n\x11SchemaSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x11\n\tschema_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x39\n\x0bschema_type\x18\x04 \x01(\x0e\x32$.spaceone.api.identity.v2.SchemaType\x12\x10\n\x08provider\x18\x05 \x01(\t\x12\x19\n\x11related_schema_id\x18\x06 \x01(\t\x12\x12\n\nis_managed\x18\x07 \x01(\x08\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"Y\n\x0bSchemasInfo\x12\x35\n\x07results\x18\x01 \x03(\x0b\x32$.spaceone.api.identity.v2.SchemaInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"Z\n\x0fSchemaStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery\x12\x11\n\tdomain_id\x18\x15 \x01(\t*m\n\nSchemaType\x12\x14\n\x10SCHEMA_TYPE_NONE\x10\x00\x12\x13\n\x0fSERVICE_ACCOUNT\x10\x01\x12\x13\n\x0fTRUSTED_ACCOUNT\x10\x02\x12\n\n\x06SECRET\x10\x03\x12\x13\n\x0fTRUSTING_SECRET\x10\x04\x32\xf4\x05\n\x06Schema\x12\x84\x01\n\x06\x63reate\x12-.spaceone.api.identity.v2.CreateSchemaRequest\x1a$.spaceone.api.identity.v2.SchemaInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/identity/v2/schema/create:\x01*\x12\x84\x01\n\x06update\x12-.spaceone.api.identity.v2.UpdateSchemaRequest\x1a$.spaceone.api.identity.v2.SchemaInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/identity/v2/schema/update:\x01*\x12p\n\x06\x64\x65lete\x12\'.spaceone.api.identity.v2.SchemaRequest\x1a\x16.google.protobuf.Empty\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/identity/v2/schema/delete:\x01*\x12x\n\x03get\x12\'.spaceone.api.identity.v2.SchemaRequest\x1a$.spaceone.api.identity.v2.SchemaInfo\"\"\x82\xd3\xe4\x93\x02\x1c\"\x17/identity/v2/schema/get:\x01*\x12\x7f\n\x04list\x12+.spaceone.api.identity.v2.SchemaSearchQuery\x1a%.spaceone.api.identity.v2.SchemasInfo\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/identity/v2/schema/list:\x01*\x12o\n\x04stat\x12).spaceone.api.identity.v2.SchemaStatQuery\x1a\x17.google.protobuf.Struct\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/identity/v2/schema/stat:\x01*B?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.identity.v2.schema_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2'
  _SCHEMA.methods_by_name['create']._options = None
  _SCHEMA.methods_by_name['create']._serialized_options = b'\202\323\344\223\002\037\"\032/identity/v2/schema/create:\001*'
  _SCHEMA.methods_by_name['update']._options = None
  _SCHEMA.methods_by_name['update']._serialized_options = b'\202\323\344\223\002\037\"\032/identity/v2/schema/update:\001*'
  _SCHEMA.methods_by_name['delete']._options = None
  _SCHEMA.methods_by_name['delete']._serialized_options = b'\202\323\344\223\002\037\"\032/identity/v2/schema/delete:\001*'
  _SCHEMA.methods_by_name['get']._options = None
  _SCHEMA.methods_by_name['get']._serialized_options = b'\202\323\344\223\002\034\"\027/identity/v2/schema/get:\001*'
  _SCHEMA.methods_by_name['list']._options = None
  _SCHEMA.methods_by_name['list']._serialized_options = b'\202\323\344\223\002\035\"\030/identity/v2/schema/list:\001*'
  _SCHEMA.methods_by_name['stat']._options = None
  _SCHEMA.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\035\"\030/identity/v2/schema/stat:\001*'
  _globals['_SCHEMATYPE']._serialized_start=1544
  _globals['_SCHEMATYPE']._serialized_end=1653
  _globals['_CREATESCHEMAREQUEST']._serialized_start=191
  _globals['_CREATESCHEMAREQUEST']._serialized_end=488
  _globals['_UPDATESCHEMAREQUEST']._serialized_start=491
  _globals['_UPDATESCHEMAREQUEST']._serialized_end=711
  _globals['_SCHEMAREQUEST']._serialized_start=713
  _globals['_SCHEMAREQUEST']._serialized_end=766
  _globals['_SCHEMAINFO']._serialized_start=769
  _globals['_SCHEMAINFO']._serialized_end=1117
  _globals['_SCHEMASEARCHQUERY']._serialized_start=1120
  _globals['_SCHEMASEARCHQUERY']._serialized_end=1359
  _globals['_SCHEMASINFO']._serialized_start=1361
  _globals['_SCHEMASINFO']._serialized_end=1450
  _globals['_SCHEMASTATQUERY']._serialized_start=1452
  _globals['_SCHEMASTATQUERY']._serialized_end=1542
  _globals['_SCHEMA']._serialized_start=1656
  _globals['_SCHEMA']._serialized_end=2412
# @@protoc_insertion_point(module_scope)
