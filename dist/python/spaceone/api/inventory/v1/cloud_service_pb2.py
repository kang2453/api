# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/inventory/v1/cloud_service.proto
# Protobuf Python Version: 4.25.0
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
from spaceone.api.core.v1 import query_pb2 as spaceone_dot_api_dot_core_dot_v1_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n-spaceone/api/inventory/v1/cloud_service.proto\x12\x19spaceone.api.inventory.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"C\n\x15\x43loudServiceReference\x12\x13\n\x0bresource_id\x18\x01 \x01(\t\x12\x15\n\rexternal_link\x18\x02 \x01(\t\"\x82\x01\n\x0e\x43ollectionInfo\x12\x10\n\x08provider\x18\x01 \x01(\t\x12\x1a\n\x12service_account_id\x18\x02 \x01(\t\x12\x11\n\tsecret_id\x18\x03 \x01(\t\x12\x14\n\x0c\x63ollector_id\x18\x04 \x01(\t\x12\x19\n\x11last_collected_at\x18\x05 \x01(\t\"\xab\x03\n\x14\x43reateServiceRequest\x12\x1a\n\x12\x63loud_service_type\x18\x01 \x01(\t\x12\x10\n\x08provider\x18\x02 \x01(\t\x12\x1b\n\x13\x63loud_service_group\x18\x03 \x01(\t\x12\x0c\n\x04name\x18\x04 \x01(\t\x12\x0f\n\x07\x61\x63\x63ount\x18\x05 \x01(\t\x12\x15\n\rinstance_type\x18\x06 \x01(\t\x12\x15\n\rinstance_size\x18\x07 \x01(\x02\x12\x14\n\x0cip_addresses\x18\n \x03(\t\x12%\n\x04\x64\x61ta\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08metadata\x18\x0c \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x43\n\treference\x18\r \x01(\x0b\x32\x30.spaceone.api.inventory.v1.CloudServiceReference\x12%\n\x04tags\x18\x0e \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x13\n\x0bregion_code\x18\x0f \x01(\t\x12\x12\n\nproject_id\x18\x15 \x01(\t\"\xff\x02\n\x19UpdateCloudServiceRequest\x12\x18\n\x10\x63loud_service_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0f\n\x07\x61\x63\x63ount\x18\x03 \x01(\t\x12\x15\n\rinstance_type\x18\x04 \x01(\t\x12\x15\n\rinstance_size\x18\x05 \x01(\x02\x12\x14\n\x0cip_addresses\x18\x06 \x03(\t\x12%\n\x04\x64\x61ta\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08metadata\x18\x0c \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x43\n\treference\x18\r \x01(\x0b\x32\x30.spaceone.api.inventory.v1.CloudServiceReference\x12%\n\x04tags\x18\x0e \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x13\n\x0bregion_code\x18\x0f \x01(\t\x12\x12\n\nproject_id\x18\x15 \x01(\t\"/\n\x13\x43loudServiceRequest\x12\x18\n\x10\x63loud_service_id\x18\x01 \x01(\t\"\xbc\x02\n\x11\x43loudServiceQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x18\n\x10\x63loud_service_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\r\n\x05state\x18\x04 \x01(\t\x12\x12\n\nip_address\x18\x05 \x01(\t\x12\x0f\n\x07\x61\x63\x63ount\x18\x06 \x01(\t\x12\x15\n\rinstance_type\x18\x07 \x01(\t\x12\x1a\n\x12\x63loud_service_type\x18\x08 \x01(\t\x12\x1b\n\x13\x63loud_service_group\x18\t \x01(\t\x12\x10\n\x08provider\x18\n \x01(\t\x12\x13\n\x0bregion_code\x18\x0b \x01(\t\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\x12\x12\n\nproject_id\x18\x16 \x01(\t\"\xa4\x05\n\x10\x43loudServiceInfo\x12\x18\n\x10\x63loud_service_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05state\x18\x03 \x01(\t\x12\x14\n\x0cip_addresses\x18\x04 \x03(\t\x12\x0f\n\x07\x61\x63\x63ount\x18\x05 \x01(\t\x12\x15\n\rinstance_type\x18\x06 \x01(\t\x12\x15\n\rinstance_size\x18\x07 \x01(\x02\x12\x1a\n\x12\x63loud_service_type\x18\x08 \x01(\t\x12\x1b\n\x13\x63loud_service_group\x18\t \x01(\t\x12\x10\n\x08provider\x18\n \x01(\t\x12%\n\x04\x64\x61ta\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08metadata\x18\x0c \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x43\n\treference\x18\r \x01(\x0b\x32\x30.spaceone.api.inventory.v1.CloudServiceReference\x12%\n\x04tags\x18\x0e \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08tag_keys\x18\x0f \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x42\n\x0f\x63ollection_info\x18\x10 \x03(\x0b\x32).spaceone.api.inventory.v1.CollectionInfo\x12\x13\n\x0bregion_code\x18\x11 \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\nproject_id\x18\x17 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\x12\x12\n\ndeleted_at\x18! \x01(\t\"f\n\x11\x43loudServicesInfo\x12<\n\x07results\x18\x01 \x03(\x0b\x32+.spaceone.api.inventory.v1.CloudServiceInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"\x83\x02\n\x19\x43loudServiceExportRequest\x12\x33\n\x07options\x18\x01 \x03(\x0b\x32\".spaceone.api.core.v1.ExportOption\x12T\n\x0b\x66ile_format\x18\x02 \x01(\x0e\x32?.spaceone.api.inventory.v1.CloudServiceExportRequest.FileFormat\x12\x11\n\tfile_name\x18\x03 \x01(\t\x12\x10\n\x08timezone\x18\x04 \x01(\t\"6\n\nFileFormat\x12\x14\n\x10NONE_FILE_FORMAT\x10\x00\x12\t\n\x05\x45XCEL\x10\x01\x12\x07\n\x03\x43SV\x10\x02\"M\n\x18\x43loudServiceAnalyzeQuery\x12\x31\n\x05query\x18\x01 \x01(\x0b\x32\".spaceone.api.core.v1.AnalyzeQuery\"M\n\x15\x43loudServiceStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery2\xfa\x08\n\x0c\x43loudService\x12\x95\x01\n\x06\x63reate\x12/.spaceone.api.inventory.v1.CreateServiceRequest\x1a+.spaceone.api.inventory.v1.CloudServiceInfo\"-\x82\xd3\xe4\x93\x02\'\"\"/inventory/v1/cloud-service/create:\x01*\x12\x9a\x01\n\x06update\x12\x34.spaceone.api.inventory.v1.UpdateCloudServiceRequest\x1a+.spaceone.api.inventory.v1.CloudServiceInfo\"-\x82\xd3\xe4\x93\x02\'\"\"/inventory/v1/cloud-service/update:\x01*\x12\x7f\n\x06\x64\x65lete\x12..spaceone.api.inventory.v1.CloudServiceRequest\x1a\x16.google.protobuf.Empty\"-\x82\xd3\xe4\x93\x02\'\"\"/inventory/v1/cloud-service/delete:\x01*\x12\x8e\x01\n\x03get\x12..spaceone.api.inventory.v1.CloudServiceRequest\x1a+.spaceone.api.inventory.v1.CloudServiceInfo\"*\x82\xd3\xe4\x93\x02$\"\x1f/inventory/v1/cloud-service/get:\x01*\x12\x8f\x01\n\x04list\x12,.spaceone.api.inventory.v1.CloudServiceQuery\x1a,.spaceone.api.inventory.v1.CloudServicesInfo\"+\x82\xd3\xe4\x93\x02%\" /inventory/v1/cloud-service/list:\x01*\x12\x86\x01\n\x06\x65xport\x12\x34.spaceone.api.inventory.v1.CloudServiceExportRequest\x1a\x17.google.protobuf.Struct\"-\x82\xd3\xe4\x93\x02\'\"\"/inventory/v1/cloud-service/export:\x01*\x12\x87\x01\n\x07\x61nalyze\x12\x33.spaceone.api.inventory.v1.CloudServiceAnalyzeQuery\x1a\x17.google.protobuf.Struct\".\x82\xd3\xe4\x93\x02(\"#/inventory/v1/cloud-service/analyze:\x01*\x12~\n\x04stat\x12\x30.spaceone.api.inventory.v1.CloudServiceStatQuery\x1a\x17.google.protobuf.Struct\"+\x82\xd3\xe4\x93\x02%\" /inventory/v1/cloud-service/stat:\x01*B@Z>github.com/cloudforet-io/api/dist/go/spaceone/api/inventory/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.inventory.v1.cloud_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z>github.com/cloudforet-io/api/dist/go/spaceone/api/inventory/v1'
  _globals['_CLOUDSERVICE'].methods_by_name['create']._options = None
  _globals['_CLOUDSERVICE'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002\'\"\"/inventory/v1/cloud-service/create:\001*'
  _globals['_CLOUDSERVICE'].methods_by_name['update']._options = None
  _globals['_CLOUDSERVICE'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002\'\"\"/inventory/v1/cloud-service/update:\001*'
  _globals['_CLOUDSERVICE'].methods_by_name['delete']._options = None
  _globals['_CLOUDSERVICE'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002\'\"\"/inventory/v1/cloud-service/delete:\001*'
  _globals['_CLOUDSERVICE'].methods_by_name['get']._options = None
  _globals['_CLOUDSERVICE'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002$\"\037/inventory/v1/cloud-service/get:\001*'
  _globals['_CLOUDSERVICE'].methods_by_name['list']._options = None
  _globals['_CLOUDSERVICE'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002%\" /inventory/v1/cloud-service/list:\001*'
  _globals['_CLOUDSERVICE'].methods_by_name['export']._options = None
  _globals['_CLOUDSERVICE'].methods_by_name['export']._serialized_options = b'\202\323\344\223\002\'\"\"/inventory/v1/cloud-service/export:\001*'
  _globals['_CLOUDSERVICE'].methods_by_name['analyze']._options = None
  _globals['_CLOUDSERVICE'].methods_by_name['analyze']._serialized_options = b'\202\323\344\223\002(\"#/inventory/v1/cloud-service/analyze:\001*'
  _globals['_CLOUDSERVICE'].methods_by_name['stat']._options = None
  _globals['_CLOUDSERVICE'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002%\" /inventory/v1/cloud-service/stat:\001*'
  _globals['_CLOUDSERVICEREFERENCE']._serialized_start=199
  _globals['_CLOUDSERVICEREFERENCE']._serialized_end=266
  _globals['_COLLECTIONINFO']._serialized_start=269
  _globals['_COLLECTIONINFO']._serialized_end=399
  _globals['_CREATESERVICEREQUEST']._serialized_start=402
  _globals['_CREATESERVICEREQUEST']._serialized_end=829
  _globals['_UPDATECLOUDSERVICEREQUEST']._serialized_start=832
  _globals['_UPDATECLOUDSERVICEREQUEST']._serialized_end=1215
  _globals['_CLOUDSERVICEREQUEST']._serialized_start=1217
  _globals['_CLOUDSERVICEREQUEST']._serialized_end=1264
  _globals['_CLOUDSERVICEQUERY']._serialized_start=1267
  _globals['_CLOUDSERVICEQUERY']._serialized_end=1583
  _globals['_CLOUDSERVICEINFO']._serialized_start=1586
  _globals['_CLOUDSERVICEINFO']._serialized_end=2262
  _globals['_CLOUDSERVICESINFO']._serialized_start=2264
  _globals['_CLOUDSERVICESINFO']._serialized_end=2366
  _globals['_CLOUDSERVICEEXPORTREQUEST']._serialized_start=2369
  _globals['_CLOUDSERVICEEXPORTREQUEST']._serialized_end=2628
  _globals['_CLOUDSERVICEEXPORTREQUEST_FILEFORMAT']._serialized_start=2574
  _globals['_CLOUDSERVICEEXPORTREQUEST_FILEFORMAT']._serialized_end=2628
  _globals['_CLOUDSERVICEANALYZEQUERY']._serialized_start=2630
  _globals['_CLOUDSERVICEANALYZEQUERY']._serialized_end=2707
  _globals['_CLOUDSERVICESTATQUERY']._serialized_start=2709
  _globals['_CLOUDSERVICESTATQUERY']._serialized_end=2786
  _globals['_CLOUDSERVICE']._serialized_start=2789
  _globals['_CLOUDSERVICE']._serialized_end=3935
# @@protoc_insertion_point(module_scope)
