# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/cost_analysis/v1/cost_query_set.proto
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
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from spaceone.api.core.v1 import query_pb2 as spaceone_dot_api_dot_core_dot_v1_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n2spaceone/api/cost_analysis/v1/cost_query_set.proto\x12\x1dspaceone.api.cost_analysis.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"\xa5\x01\n\x19\x43reateCostQuerySetRequest\x12\x16\n\x0e\x64\x61ta_source_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"\xa8\x01\n\x19UpdateCostQuerySetRequest\x12\x19\n\x11\x63ost_query_set_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"C\n\x13\x43ostQuerySetRequest\x12\x19\n\x11\x63ost_query_set_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"T\n\x16GetCostQuerySetRequest\x12\x19\n\x11\x63ost_query_set_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x0c\n\x04only\x18\x03 \x03(\t\"\xa4\x01\n\x11\x43ostQuerySetQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x16\n\x0e\x64\x61ta_source_id\x18\x02 \x01(\t\x12\x19\n\x11\x63ost_query_set_id\x18\x03 \x01(\t\x12\x0c\n\x04name\x18\x04 \x01(\t\x12\x0f\n\x07user_id\x18\x05 \x01(\t\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"\xd8\x01\n\x10\x43ostQuerySetInfo\x12\x19\n\x11\x63ost_query_set_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0f\n\x07user_id\x18\x0b \x01(\t\x12\x11\n\tdomain_id\x18\x0c \x01(\t\x12\x12\n\ncreated_at\x18\x15 \x01(\t\x12\x12\n\nupdated_at\x18\x16 \x01(\t\"j\n\x11\x43ostQuerySetsInfo\x12@\n\x07results\x18\x01 \x03(\x0b\x32/.spaceone.api.cost_analysis.v1.CostQuerySetInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"x\n\x15\x43ostQuerySetStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\x12\x16\n\x0e\x64\x61ta_source_id\x18\x02 \x01(\t\x12\x11\n\tdomain_id\x18\x03 \x01(\t2\xb9\x07\n\x0c\x43ostQuerySet\x12\xa7\x01\n\x06\x63reate\x12\x38.spaceone.api.cost_analysis.v1.CreateCostQuerySetRequest\x1a/.spaceone.api.cost_analysis.v1.CostQuerySetInfo\"2\x82\xd3\xe4\x93\x02,\"\'/cost-analysis/v1/cost-query-set/create:\x01*\x12\xa7\x01\n\x06update\x12\x38.spaceone.api.cost_analysis.v1.UpdateCostQuerySetRequest\x1a/.spaceone.api.cost_analysis.v1.CostQuerySetInfo\"2\x82\xd3\xe4\x93\x02,\"\'/cost-analysis/v1/cost-query-set/update:\x01*\x12\x88\x01\n\x06\x64\x65lete\x12\x32.spaceone.api.cost_analysis.v1.CostQuerySetRequest\x1a\x16.google.protobuf.Empty\"2\x82\xd3\xe4\x93\x02,\"\'/cost-analysis/v1/cost-query-set/delete:\x01*\x12\x9e\x01\n\x03get\x12\x35.spaceone.api.cost_analysis.v1.GetCostQuerySetRequest\x1a/.spaceone.api.cost_analysis.v1.CostQuerySetInfo\"/\x82\xd3\xe4\x93\x02)\"$/cost-analysis/v1/cost-query-set/get:\x01*\x12\x9d\x01\n\x04list\x12\x30.spaceone.api.cost_analysis.v1.CostQuerySetQuery\x1a\x30.spaceone.api.cost_analysis.v1.CostQuerySetsInfo\"1\x82\xd3\xe4\x93\x02+\"&/cost-analysis/v1/cost-query-sets/list:\x01*\x12\x88\x01\n\x04stat\x12\x34.spaceone.api.cost_analysis.v1.CostQuerySetStatQuery\x1a\x17.google.protobuf.Struct\"1\x82\xd3\xe4\x93\x02+\"&/cost-analysis/v1/cost-query-sets/stat:\x01*BDZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1b\x06proto3')



_CREATECOSTQUERYSETREQUEST = DESCRIPTOR.message_types_by_name['CreateCostQuerySetRequest']
_UPDATECOSTQUERYSETREQUEST = DESCRIPTOR.message_types_by_name['UpdateCostQuerySetRequest']
_COSTQUERYSETREQUEST = DESCRIPTOR.message_types_by_name['CostQuerySetRequest']
_GETCOSTQUERYSETREQUEST = DESCRIPTOR.message_types_by_name['GetCostQuerySetRequest']
_COSTQUERYSETQUERY = DESCRIPTOR.message_types_by_name['CostQuerySetQuery']
_COSTQUERYSETINFO = DESCRIPTOR.message_types_by_name['CostQuerySetInfo']
_COSTQUERYSETSINFO = DESCRIPTOR.message_types_by_name['CostQuerySetsInfo']
_COSTQUERYSETSTATQUERY = DESCRIPTOR.message_types_by_name['CostQuerySetStatQuery']
CreateCostQuerySetRequest = _reflection.GeneratedProtocolMessageType('CreateCostQuerySetRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATECOSTQUERYSETREQUEST,
  '__module__' : 'spaceone.api.cost_analysis.v1.cost_query_set_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.CreateCostQuerySetRequest)
  })
_sym_db.RegisterMessage(CreateCostQuerySetRequest)

UpdateCostQuerySetRequest = _reflection.GeneratedProtocolMessageType('UpdateCostQuerySetRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATECOSTQUERYSETREQUEST,
  '__module__' : 'spaceone.api.cost_analysis.v1.cost_query_set_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.UpdateCostQuerySetRequest)
  })
_sym_db.RegisterMessage(UpdateCostQuerySetRequest)

CostQuerySetRequest = _reflection.GeneratedProtocolMessageType('CostQuerySetRequest', (_message.Message,), {
  'DESCRIPTOR' : _COSTQUERYSETREQUEST,
  '__module__' : 'spaceone.api.cost_analysis.v1.cost_query_set_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.CostQuerySetRequest)
  })
_sym_db.RegisterMessage(CostQuerySetRequest)

GetCostQuerySetRequest = _reflection.GeneratedProtocolMessageType('GetCostQuerySetRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETCOSTQUERYSETREQUEST,
  '__module__' : 'spaceone.api.cost_analysis.v1.cost_query_set_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.GetCostQuerySetRequest)
  })
_sym_db.RegisterMessage(GetCostQuerySetRequest)

CostQuerySetQuery = _reflection.GeneratedProtocolMessageType('CostQuerySetQuery', (_message.Message,), {
  'DESCRIPTOR' : _COSTQUERYSETQUERY,
  '__module__' : 'spaceone.api.cost_analysis.v1.cost_query_set_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.CostQuerySetQuery)
  })
_sym_db.RegisterMessage(CostQuerySetQuery)

CostQuerySetInfo = _reflection.GeneratedProtocolMessageType('CostQuerySetInfo', (_message.Message,), {
  'DESCRIPTOR' : _COSTQUERYSETINFO,
  '__module__' : 'spaceone.api.cost_analysis.v1.cost_query_set_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.CostQuerySetInfo)
  })
_sym_db.RegisterMessage(CostQuerySetInfo)

CostQuerySetsInfo = _reflection.GeneratedProtocolMessageType('CostQuerySetsInfo', (_message.Message,), {
  'DESCRIPTOR' : _COSTQUERYSETSINFO,
  '__module__' : 'spaceone.api.cost_analysis.v1.cost_query_set_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.CostQuerySetsInfo)
  })
_sym_db.RegisterMessage(CostQuerySetsInfo)

CostQuerySetStatQuery = _reflection.GeneratedProtocolMessageType('CostQuerySetStatQuery', (_message.Message,), {
  'DESCRIPTOR' : _COSTQUERYSETSTATQUERY,
  '__module__' : 'spaceone.api.cost_analysis.v1.cost_query_set_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.CostQuerySetStatQuery)
  })
_sym_db.RegisterMessage(CostQuerySetStatQuery)

_COSTQUERYSET = DESCRIPTOR.services_by_name['CostQuerySet']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1'
  _COSTQUERYSET.methods_by_name['create']._options = None
  _COSTQUERYSET.methods_by_name['create']._serialized_options = b'\202\323\344\223\002,\"\'/cost-analysis/v1/cost-query-set/create:\001*'
  _COSTQUERYSET.methods_by_name['update']._options = None
  _COSTQUERYSET.methods_by_name['update']._serialized_options = b'\202\323\344\223\002,\"\'/cost-analysis/v1/cost-query-set/update:\001*'
  _COSTQUERYSET.methods_by_name['delete']._options = None
  _COSTQUERYSET.methods_by_name['delete']._serialized_options = b'\202\323\344\223\002,\"\'/cost-analysis/v1/cost-query-set/delete:\001*'
  _COSTQUERYSET.methods_by_name['get']._options = None
  _COSTQUERYSET.methods_by_name['get']._serialized_options = b'\202\323\344\223\002)\"$/cost-analysis/v1/cost-query-set/get:\001*'
  _COSTQUERYSET.methods_by_name['list']._options = None
  _COSTQUERYSET.methods_by_name['list']._serialized_options = b'\202\323\344\223\002+\"&/cost-analysis/v1/cost-query-sets/list:\001*'
  _COSTQUERYSET.methods_by_name['stat']._options = None
  _COSTQUERYSET.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002+\"&/cost-analysis/v1/cost-query-sets/stat:\001*'
  _CREATECOSTQUERYSETREQUEST._serialized_start=209
  _CREATECOSTQUERYSETREQUEST._serialized_end=374
  _UPDATECOSTQUERYSETREQUEST._serialized_start=377
  _UPDATECOSTQUERYSETREQUEST._serialized_end=545
  _COSTQUERYSETREQUEST._serialized_start=547
  _COSTQUERYSETREQUEST._serialized_end=614
  _GETCOSTQUERYSETREQUEST._serialized_start=616
  _GETCOSTQUERYSETREQUEST._serialized_end=700
  _COSTQUERYSETQUERY._serialized_start=703
  _COSTQUERYSETQUERY._serialized_end=867
  _COSTQUERYSETINFO._serialized_start=870
  _COSTQUERYSETINFO._serialized_end=1086
  _COSTQUERYSETSINFO._serialized_start=1088
  _COSTQUERYSETSINFO._serialized_end=1194
  _COSTQUERYSETSTATQUERY._serialized_start=1196
  _COSTQUERYSETSTATQUERY._serialized_end=1316
  _COSTQUERYSET._serialized_start=1319
  _COSTQUERYSET._serialized_end=2272
# @@protoc_insertion_point(module_scope)
