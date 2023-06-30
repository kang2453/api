# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/cost_analysis/v1/data_source.proto
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
from spaceone.api.cost_analysis.v1 import job_pb2 as spaceone_dot_api_dot_cost__analysis_dot_v1_dot_job__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n/spaceone/api/cost_analysis/v1/data_source.proto\x12\x1dspaceone.api.cost_analysis.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\x1a\'spaceone/api/cost_analysis/v1/job.proto\"\xd2\x01\n\x0cSecretFilter\x12L\n\x05state\x18\x01 \x01(\x0e\x32=.spaceone.api.cost_analysis.v1.SecretFilter.SecretFilterState\x12\x0f\n\x07secrets\x18\x02 \x03(\t\x12\x18\n\x10service_accounts\x18\x03 \x03(\t\x12\x0f\n\x07schemas\x18\x04 \x03(\t\"8\n\x11SecretFilterState\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"\xd2\x02\n\nPluginInfo\x12\x11\n\tplugin_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08metadata\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12,\n\x0bsecret_data\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0e\n\x06schema\x18\x06 \x01(\t\x12\x11\n\tsecret_id\x18\x07 \x01(\t\x12K\n\x0cupgrade_mode\x18\x08 \x01(\x0e\x32\x35.spaceone.api.cost_analysis.v1.PluginInfo.UpgradeMode\"-\n\x0bUpgradeMode\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06MANUAL\x10\x01\x12\x08\n\x04\x41UTO\x10\x02\"\xe6\x04\n\x19RegisterDataSourceRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x61\n\x10\x64\x61ta_source_type\x18\x02 \x01(\x0e\x32G.spaceone.api.cost_analysis.v1.RegisterDataSourceRequest.DataSourceType\x12\x10\n\x08provider\x18\x03 \x01(\t\x12)\n\x08template\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12>\n\x0bplugin_info\x18\x05 \x01(\x0b\x32).spaceone.api.cost_analysis.v1.PluginInfo\x12%\n\x04tags\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12X\n\x0bsecret_type\x18\x07 \x01(\x0e\x32\x43.spaceone.api.cost_analysis.v1.RegisterDataSourceRequest.SecretType\x12\x42\n\rsecret_filter\x18\x08 \x01(\x0b\x32+.spaceone.api.cost_analysis.v1.SecretFilter\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"3\n\x0e\x44\x61taSourceType\x12\x08\n\x04NONE\x10\x00\x12\t\n\x05LOCAL\x10\x01\x12\x0c\n\x08\x45XTERNAL\x10\x02\"N\n\nSecretType\x12\x14\n\x10SECRET_TYPE_NONE\x10\x00\x12\n\n\x06MANUAL\x10\x01\x12\x1e\n\x1aUSE_SERVICE_ACCOUNT_SECRET\x10\x02\"\xa4\x01\n\x17UpdateDataSourceRequest\x12\x16\n\x0e\x64\x61ta_source_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12)\n\x08template\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"\x94\x02\n\x1dUpdateDataSourcePluginRequest\x12\x16\n\x0e\x64\x61ta_source_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12^\n\x0cupgrade_mode\x18\x04 \x01(\x0e\x32H.spaceone.api.cost_analysis.v1.UpdateDataSourcePluginRequest.UpgradeMode\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"-\n\x0bUpgradeMode\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06MANUAL\x10\x01\x12\x08\n\x04\x41UTO\x10\x02\"e\n\x1b\x44\x65registerDataSourceRequest\x12\x16\n\x0e\x64\x61ta_source_id\x18\x01 \x01(\t\x12\x1b\n\x13\x63\x61scade_delete_cost\x18\x02 \x01(\x08\x12\x11\n\tdomain_id\x18\x03 \x01(\t\"k\n\x15SyncDataSourceRequest\x12\x16\n\x0e\x64\x61ta_source_id\x18\x01 \x01(\t\x12\r\n\x05start\x18\x02 \x01(\t\x12\x18\n\x10no_preload_cache\x18\x03 \x01(\x08\x12\x11\n\tdomain_id\x18\x0b \x01(\t\">\n\x11\x44\x61taSourceRequest\x12\x16\n\x0e\x64\x61ta_source_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"O\n\x14GetDataSourceRequest\x12\x16\n\x0e\x64\x61ta_source_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x0c\n\x04only\x18\x03 \x03(\t\"\xa5\x02\n\x0f\x44\x61taSourceQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x16\n\x0e\x64\x61ta_source_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\r\n\x05state\x18\x04 \x01(\t\x12W\n\x10\x64\x61ta_source_type\x18\x05 \x01(\x0e\x32=.spaceone.api.cost_analysis.v1.DataSourceQuery.DataSourceType\x12\x10\n\x08provider\x18\x06 \x01(\t\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"3\n\x0e\x44\x61taSourceType\x12\x08\n\x04NONE\x10\x00\x12\t\n\x05LOCAL\x10\x01\x12\x0c\n\x08\x45XTERNAL\x10\x02\"\xc6\x06\n\x0e\x44\x61taSourceInfo\x12\x16\n\x0e\x64\x61ta_source_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x42\n\x05state\x18\x03 \x01(\x0e\x32\x33.spaceone.api.cost_analysis.v1.DataSourceInfo.State\x12V\n\x10\x64\x61ta_source_type\x18\x04 \x01(\x0e\x32<.spaceone.api.cost_analysis.v1.DataSourceInfo.DataSourceType\x12\x10\n\x08provider\x18\x05 \x01(\t\x12>\n\x0bplugin_info\x18\x06 \x01(\x0b\x32).spaceone.api.cost_analysis.v1.PluginInfo\x12)\n\x08template\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x15\n\rcost_tag_keys\x18\t \x03(\t\x12!\n\x19\x63ost_additional_info_keys\x18\n \x03(\t\x12\x11\n\tdomain_id\x18\x0b \x01(\t\x12M\n\x0bsecret_type\x18\x0c \x01(\x0e\x32\x38.spaceone.api.cost_analysis.v1.DataSourceInfo.SecretType\x12\x42\n\rsecret_filter\x18\r \x01(\x0b\x32+.spaceone.api.cost_analysis.v1.SecretFilter\x12\x12\n\ncreated_at\x18\x15 \x01(\t\x12\x1c\n\x14last_synchronized_at\x18\x16 \x01(\t\"2\n\x05State\x12\x0e\n\nSTATE_NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"8\n\x0e\x44\x61taSourceType\x12\r\n\tTYPE_NONE\x10\x00\x12\t\n\x05LOCAL\x10\x01\x12\x0c\n\x08\x45XTERNAL\x10\x02\"N\n\nSecretType\x12\x14\n\x10SECRET_TYPE_NONE\x10\x00\x12\n\n\x06MANUAL\x10\x01\x12\x1e\n\x1aUSE_SERVICE_ACCOUNT_SECRET\x10\x02\"f\n\x0f\x44\x61taSourcesInfo\x12>\n\x07results\x18\x01 \x03(\x0b\x32-.spaceone.api.cost_analysis.v1.DataSourceInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"^\n\x13\x44\x61taSourceStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\x12\x11\n\tdomain_id\x18\x02 \x01(\t2\xc4\r\n\nDataSource\x12\xa6\x01\n\x08register\x12\x38.spaceone.api.cost_analysis.v1.RegisterDataSourceRequest\x1a-.spaceone.api.cost_analysis.v1.DataSourceInfo\"1\x82\xd3\xe4\x93\x02+\"&/cost-analysis/v1/data-source/register:\x01*\x12\xa0\x01\n\x06update\x12\x36.spaceone.api.cost_analysis.v1.UpdateDataSourceRequest\x1a-.spaceone.api.cost_analysis.v1.DataSourceInfo\"/\x82\xd3\xe4\x93\x02)\"$/cost-analysis/v1/data-source/update:\x01*\x12\xb4\x01\n\rupdate_plugin\x12<.spaceone.api.cost_analysis.v1.UpdateDataSourcePluginRequest\x1a-.spaceone.api.cost_analysis.v1.DataSourceInfo\"6\x82\xd3\xe4\x93\x02\x30\"+/cost-analysis/v1/data-source/update-plugin:\x01*\x12\x91\x01\n\rverify_plugin\x12\x30.spaceone.api.cost_analysis.v1.DataSourceRequest\x1a\x16.google.protobuf.Empty\"6\x82\xd3\xe4\x93\x02\x30\"+/cost-analysis/v1/data-source/verify-plugin:\x01*\x12\x9a\x01\n\x06\x65nable\x12\x30.spaceone.api.cost_analysis.v1.DataSourceRequest\x1a-.spaceone.api.cost_analysis.v1.DataSourceInfo\"/\x82\xd3\xe4\x93\x02)\"$/cost-analysis/v1/data-source/enable:\x01*\x12\x9c\x01\n\x07\x64isable\x12\x30.spaceone.api.cost_analysis.v1.DataSourceRequest\x1a-.spaceone.api.cost_analysis.v1.DataSourceInfo\"0\x82\xd3\xe4\x93\x02*\"%/cost-analysis/v1/data-source/disable:\x01*\x12\x95\x01\n\nderegister\x12:.spaceone.api.cost_analysis.v1.DeregisterDataSourceRequest\x1a\x16.google.protobuf.Empty\"3\x82\xd3\xe4\x93\x02-\"(/cost-analysis/v1/data-source/deregister:\x01*\x12\x93\x01\n\x04sync\x12\x34.spaceone.api.cost_analysis.v1.SyncDataSourceRequest\x1a&.spaceone.api.cost_analysis.v1.JobInfo\"-\x82\xd3\xe4\x93\x02\'\"\"/cost-analysis/v1/data-source/sync:\x01*\x12\x97\x01\n\x03get\x12\x33.spaceone.api.cost_analysis.v1.GetDataSourceRequest\x1a-.spaceone.api.cost_analysis.v1.DataSourceInfo\",\x82\xd3\xe4\x93\x02&\"!/cost-analysis/v1/data-source/get:\x01*\x12\x95\x01\n\x04list\x12..spaceone.api.cost_analysis.v1.DataSourceQuery\x1a..spaceone.api.cost_analysis.v1.DataSourcesInfo\"-\x82\xd3\xe4\x93\x02\'\"\"/cost-analysis/v1/data-source/list:\x01*\x12\x82\x01\n\x04stat\x12\x32.spaceone.api.cost_analysis.v1.DataSourceStatQuery\x1a\x17.google.protobuf.Struct\"-\x82\xd3\xe4\x93\x02\'\"\"/cost-analysis/v1/data-source/stat:\x01*BDZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1b\x06proto3')



_SECRETFILTER = DESCRIPTOR.message_types_by_name['SecretFilter']
_PLUGININFO = DESCRIPTOR.message_types_by_name['PluginInfo']
_REGISTERDATASOURCEREQUEST = DESCRIPTOR.message_types_by_name['RegisterDataSourceRequest']
_UPDATEDATASOURCEREQUEST = DESCRIPTOR.message_types_by_name['UpdateDataSourceRequest']
_UPDATEDATASOURCEPLUGINREQUEST = DESCRIPTOR.message_types_by_name['UpdateDataSourcePluginRequest']
_DEREGISTERDATASOURCEREQUEST = DESCRIPTOR.message_types_by_name['DeregisterDataSourceRequest']
_SYNCDATASOURCEREQUEST = DESCRIPTOR.message_types_by_name['SyncDataSourceRequest']
_DATASOURCEREQUEST = DESCRIPTOR.message_types_by_name['DataSourceRequest']
_GETDATASOURCEREQUEST = DESCRIPTOR.message_types_by_name['GetDataSourceRequest']
_DATASOURCEQUERY = DESCRIPTOR.message_types_by_name['DataSourceQuery']
_DATASOURCEINFO = DESCRIPTOR.message_types_by_name['DataSourceInfo']
_DATASOURCESINFO = DESCRIPTOR.message_types_by_name['DataSourcesInfo']
_DATASOURCESTATQUERY = DESCRIPTOR.message_types_by_name['DataSourceStatQuery']
_SECRETFILTER_SECRETFILTERSTATE = _SECRETFILTER.enum_types_by_name['SecretFilterState']
_PLUGININFO_UPGRADEMODE = _PLUGININFO.enum_types_by_name['UpgradeMode']
_REGISTERDATASOURCEREQUEST_DATASOURCETYPE = _REGISTERDATASOURCEREQUEST.enum_types_by_name['DataSourceType']
_REGISTERDATASOURCEREQUEST_SECRETTYPE = _REGISTERDATASOURCEREQUEST.enum_types_by_name['SecretType']
_UPDATEDATASOURCEPLUGINREQUEST_UPGRADEMODE = _UPDATEDATASOURCEPLUGINREQUEST.enum_types_by_name['UpgradeMode']
_DATASOURCEQUERY_DATASOURCETYPE = _DATASOURCEQUERY.enum_types_by_name['DataSourceType']
_DATASOURCEINFO_STATE = _DATASOURCEINFO.enum_types_by_name['State']
_DATASOURCEINFO_DATASOURCETYPE = _DATASOURCEINFO.enum_types_by_name['DataSourceType']
_DATASOURCEINFO_SECRETTYPE = _DATASOURCEINFO.enum_types_by_name['SecretType']
SecretFilter = _reflection.GeneratedProtocolMessageType('SecretFilter', (_message.Message,), {
  'DESCRIPTOR' : _SECRETFILTER,
  '__module__' : 'spaceone.api.cost_analysis.v1.data_source_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.SecretFilter)
  })
_sym_db.RegisterMessage(SecretFilter)

PluginInfo = _reflection.GeneratedProtocolMessageType('PluginInfo', (_message.Message,), {
  'DESCRIPTOR' : _PLUGININFO,
  '__module__' : 'spaceone.api.cost_analysis.v1.data_source_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.PluginInfo)
  })
_sym_db.RegisterMessage(PluginInfo)

RegisterDataSourceRequest = _reflection.GeneratedProtocolMessageType('RegisterDataSourceRequest', (_message.Message,), {
  'DESCRIPTOR' : _REGISTERDATASOURCEREQUEST,
  '__module__' : 'spaceone.api.cost_analysis.v1.data_source_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.RegisterDataSourceRequest)
  })
_sym_db.RegisterMessage(RegisterDataSourceRequest)

UpdateDataSourceRequest = _reflection.GeneratedProtocolMessageType('UpdateDataSourceRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEDATASOURCEREQUEST,
  '__module__' : 'spaceone.api.cost_analysis.v1.data_source_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.UpdateDataSourceRequest)
  })
_sym_db.RegisterMessage(UpdateDataSourceRequest)

UpdateDataSourcePluginRequest = _reflection.GeneratedProtocolMessageType('UpdateDataSourcePluginRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEDATASOURCEPLUGINREQUEST,
  '__module__' : 'spaceone.api.cost_analysis.v1.data_source_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.UpdateDataSourcePluginRequest)
  })
_sym_db.RegisterMessage(UpdateDataSourcePluginRequest)

DeregisterDataSourceRequest = _reflection.GeneratedProtocolMessageType('DeregisterDataSourceRequest', (_message.Message,), {
  'DESCRIPTOR' : _DEREGISTERDATASOURCEREQUEST,
  '__module__' : 'spaceone.api.cost_analysis.v1.data_source_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.DeregisterDataSourceRequest)
  })
_sym_db.RegisterMessage(DeregisterDataSourceRequest)

SyncDataSourceRequest = _reflection.GeneratedProtocolMessageType('SyncDataSourceRequest', (_message.Message,), {
  'DESCRIPTOR' : _SYNCDATASOURCEREQUEST,
  '__module__' : 'spaceone.api.cost_analysis.v1.data_source_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.SyncDataSourceRequest)
  })
_sym_db.RegisterMessage(SyncDataSourceRequest)

DataSourceRequest = _reflection.GeneratedProtocolMessageType('DataSourceRequest', (_message.Message,), {
  'DESCRIPTOR' : _DATASOURCEREQUEST,
  '__module__' : 'spaceone.api.cost_analysis.v1.data_source_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.DataSourceRequest)
  })
_sym_db.RegisterMessage(DataSourceRequest)

GetDataSourceRequest = _reflection.GeneratedProtocolMessageType('GetDataSourceRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETDATASOURCEREQUEST,
  '__module__' : 'spaceone.api.cost_analysis.v1.data_source_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.GetDataSourceRequest)
  })
_sym_db.RegisterMessage(GetDataSourceRequest)

DataSourceQuery = _reflection.GeneratedProtocolMessageType('DataSourceQuery', (_message.Message,), {
  'DESCRIPTOR' : _DATASOURCEQUERY,
  '__module__' : 'spaceone.api.cost_analysis.v1.data_source_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.DataSourceQuery)
  })
_sym_db.RegisterMessage(DataSourceQuery)

DataSourceInfo = _reflection.GeneratedProtocolMessageType('DataSourceInfo', (_message.Message,), {
  'DESCRIPTOR' : _DATASOURCEINFO,
  '__module__' : 'spaceone.api.cost_analysis.v1.data_source_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.DataSourceInfo)
  })
_sym_db.RegisterMessage(DataSourceInfo)

DataSourcesInfo = _reflection.GeneratedProtocolMessageType('DataSourcesInfo', (_message.Message,), {
  'DESCRIPTOR' : _DATASOURCESINFO,
  '__module__' : 'spaceone.api.cost_analysis.v1.data_source_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.DataSourcesInfo)
  })
_sym_db.RegisterMessage(DataSourcesInfo)

DataSourceStatQuery = _reflection.GeneratedProtocolMessageType('DataSourceStatQuery', (_message.Message,), {
  'DESCRIPTOR' : _DATASOURCESTATQUERY,
  '__module__' : 'spaceone.api.cost_analysis.v1.data_source_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.DataSourceStatQuery)
  })
_sym_db.RegisterMessage(DataSourceStatQuery)

_DATASOURCE = DESCRIPTOR.services_by_name['DataSource']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1'
  _DATASOURCE.methods_by_name['register']._options = None
  _DATASOURCE.methods_by_name['register']._serialized_options = b'\202\323\344\223\002+\"&/cost-analysis/v1/data-source/register:\001*'
  _DATASOURCE.methods_by_name['update']._options = None
  _DATASOURCE.methods_by_name['update']._serialized_options = b'\202\323\344\223\002)\"$/cost-analysis/v1/data-source/update:\001*'
  _DATASOURCE.methods_by_name['update_plugin']._options = None
  _DATASOURCE.methods_by_name['update_plugin']._serialized_options = b'\202\323\344\223\0020\"+/cost-analysis/v1/data-source/update-plugin:\001*'
  _DATASOURCE.methods_by_name['verify_plugin']._options = None
  _DATASOURCE.methods_by_name['verify_plugin']._serialized_options = b'\202\323\344\223\0020\"+/cost-analysis/v1/data-source/verify-plugin:\001*'
  _DATASOURCE.methods_by_name['enable']._options = None
  _DATASOURCE.methods_by_name['enable']._serialized_options = b'\202\323\344\223\002)\"$/cost-analysis/v1/data-source/enable:\001*'
  _DATASOURCE.methods_by_name['disable']._options = None
  _DATASOURCE.methods_by_name['disable']._serialized_options = b'\202\323\344\223\002*\"%/cost-analysis/v1/data-source/disable:\001*'
  _DATASOURCE.methods_by_name['deregister']._options = None
  _DATASOURCE.methods_by_name['deregister']._serialized_options = b'\202\323\344\223\002-\"(/cost-analysis/v1/data-source/deregister:\001*'
  _DATASOURCE.methods_by_name['sync']._options = None
  _DATASOURCE.methods_by_name['sync']._serialized_options = b'\202\323\344\223\002\'\"\"/cost-analysis/v1/data-source/sync:\001*'
  _DATASOURCE.methods_by_name['get']._options = None
  _DATASOURCE.methods_by_name['get']._serialized_options = b'\202\323\344\223\002&\"!/cost-analysis/v1/data-source/get:\001*'
  _DATASOURCE.methods_by_name['list']._options = None
  _DATASOURCE.methods_by_name['list']._serialized_options = b'\202\323\344\223\002\'\"\"/cost-analysis/v1/data-source/list:\001*'
  _DATASOURCE.methods_by_name['stat']._options = None
  _DATASOURCE.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\'\"\"/cost-analysis/v1/data-source/stat:\001*'
  _SECRETFILTER._serialized_start=247
  _SECRETFILTER._serialized_end=457
  _SECRETFILTER_SECRETFILTERSTATE._serialized_start=401
  _SECRETFILTER_SECRETFILTERSTATE._serialized_end=457
  _PLUGININFO._serialized_start=460
  _PLUGININFO._serialized_end=798
  _PLUGININFO_UPGRADEMODE._serialized_start=753
  _PLUGININFO_UPGRADEMODE._serialized_end=798
  _REGISTERDATASOURCEREQUEST._serialized_start=801
  _REGISTERDATASOURCEREQUEST._serialized_end=1415
  _REGISTERDATASOURCEREQUEST_DATASOURCETYPE._serialized_start=1284
  _REGISTERDATASOURCEREQUEST_DATASOURCETYPE._serialized_end=1335
  _REGISTERDATASOURCEREQUEST_SECRETTYPE._serialized_start=1337
  _REGISTERDATASOURCEREQUEST_SECRETTYPE._serialized_end=1415
  _UPDATEDATASOURCEREQUEST._serialized_start=1418
  _UPDATEDATASOURCEREQUEST._serialized_end=1582
  _UPDATEDATASOURCEPLUGINREQUEST._serialized_start=1585
  _UPDATEDATASOURCEPLUGINREQUEST._serialized_end=1861
  _UPDATEDATASOURCEPLUGINREQUEST_UPGRADEMODE._serialized_start=753
  _UPDATEDATASOURCEPLUGINREQUEST_UPGRADEMODE._serialized_end=798
  _DEREGISTERDATASOURCEREQUEST._serialized_start=1863
  _DEREGISTERDATASOURCEREQUEST._serialized_end=1964
  _SYNCDATASOURCEREQUEST._serialized_start=1966
  _SYNCDATASOURCEREQUEST._serialized_end=2073
  _DATASOURCEREQUEST._serialized_start=2075
  _DATASOURCEREQUEST._serialized_end=2137
  _GETDATASOURCEREQUEST._serialized_start=2139
  _GETDATASOURCEREQUEST._serialized_end=2218
  _DATASOURCEQUERY._serialized_start=2221
  _DATASOURCEQUERY._serialized_end=2514
  _DATASOURCEQUERY_DATASOURCETYPE._serialized_start=1284
  _DATASOURCEQUERY_DATASOURCETYPE._serialized_end=1335
  _DATASOURCEINFO._serialized_start=2517
  _DATASOURCEINFO._serialized_end=3355
  _DATASOURCEINFO_STATE._serialized_start=3167
  _DATASOURCEINFO_STATE._serialized_end=3217
  _DATASOURCEINFO_DATASOURCETYPE._serialized_start=3219
  _DATASOURCEINFO_DATASOURCETYPE._serialized_end=3275
  _DATASOURCEINFO_SECRETTYPE._serialized_start=1337
  _DATASOURCEINFO_SECRETTYPE._serialized_end=1415
  _DATASOURCESINFO._serialized_start=3357
  _DATASOURCESINFO._serialized_end=3459
  _DATASOURCESTATQUERY._serialized_start=3461
  _DATASOURCESTATQUERY._serialized_end=3555
  _DATASOURCE._serialized_start=3558
  _DATASOURCE._serialized_end=5290
# @@protoc_insertion_point(module_scope)
