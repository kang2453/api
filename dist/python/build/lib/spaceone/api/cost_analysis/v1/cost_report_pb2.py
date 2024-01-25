# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/cost_analysis/v1/cost_report.proto
# Protobuf Python Version: 4.25.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.core.v2 import query_pb2 as spaceone_dot_api_dot_core_dot_v2_dot_query__pb2
from spaceone.api.cost_analysis.v1 import job_pb2 as spaceone_dot_api_dot_cost__analysis_dot_v1_dot_job__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n/spaceone/api/cost_analysis/v1/cost_report.proto\x12\x1dspaceone.api.cost_analysis.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a spaceone/api/core/v2/query.proto\x1a\'spaceone/api/cost_analysis/v1/job.proto\"+\n\x11\x43ostReportRequest\x12\x16\n\x0e\x63ost_report_id\x18\x01 \x01(\t\"1\n\x17GetUrlCostReportRequest\x12\x16\n\x0e\x63ost_report_id\x18\x01 \x01(\t\"\xfa\x01\n\x0f\x43ostReportQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x16\n\x0e\x63ost_report_id\x18\x02 \x01(\t\x12\x45\n\x06status\x18\x03 \x01(\x0e\x32\x35.spaceone.api.cost_analysis.v1.CostReportQuery.Status\x12\x12\n\nissue_date\x18\x04 \x01(\t\x12\x16\n\x0eworkspace_name\x18\x05 \x01(\t\"0\n\x06Status\x12\x08\n\x04NONE\x10\x00\x12\x0f\n\x0bIN_PROGRESS\x10\x01\x12\x0b\n\x07SUCCESS\x10\x02\"\xa8\x03\n\x0e\x43ostReportInfo\x12\x16\n\x0e\x63ost_report_id\x18\x01 \x01(\t\x12%\n\x04\x63ost\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x44\n\x06status\x18\x03 \x01(\x0e\x32\x34.spaceone.api.cost_analysis.v1.CostReportInfo.Status\x12\x1a\n\x12\x63ost_report_number\x18\x04 \x01(\t\x12\x10\n\x08\x63urrency\x18\x05 \x01(\t\x12\x12\n\nissue_date\x18\x06 \x01(\t\x12\x13\n\x0breport_year\x18\x07 \x01(\t\x12\x14\n\x0creport_month\x18\x08 \x01(\t\x12\x16\n\x0eworkspace_name\x18\t \x01(\t\x12\x1d\n\x15\x63ost_report_config_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x11\n\tdomain_id\x18\x17 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\"0\n\x06Status\x12\x08\n\x04NONE\x10\x00\x12\x0f\n\x0bIN_PROGRESS\x10\x01\x12\x0b\n\x07SUCCESS\x10\x02\"f\n\x0f\x43ostReportsInfo\x12>\n\x07results\x18\x01 \x03(\x0b\x32-.spaceone.api.cost_analysis.v1.CostReportInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"c\n\x13\x43ostReportStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery\x12\x16\n\x0e\x63ost_report_id\x18\x02 \x01(\t2\xe8\x05\n\nCostReport\x12\x96\x01\n\x04send\x12\x30.spaceone.api.cost_analysis.v1.CostReportRequest\x1a-.spaceone.api.cost_analysis.v1.CostReportInfo\"-\x82\xd3\xe4\x93\x02\'\"\"/cost-analysis/v1/cost-report/send:\x01*\x12\x8c\x01\n\x07get_url\x12\x36.spaceone.api.cost_analysis.v1.GetUrlCostReportRequest\x1a\x17.google.protobuf.Struct\"0\x82\xd3\xe4\x93\x02*\"%/cost-analysis/v1/cost-report/get-url:\x01*\x12\x94\x01\n\x03get\x12\x30.spaceone.api.cost_analysis.v1.CostReportRequest\x1a-.spaceone.api.cost_analysis.v1.CostReportInfo\",\x82\xd3\xe4\x93\x02&\"!/cost-analysis/v1/cost-report/get:\x01*\x12\x95\x01\n\x04list\x12..spaceone.api.cost_analysis.v1.CostReportQuery\x1a..spaceone.api.cost_analysis.v1.CostReportsInfo\"-\x82\xd3\xe4\x93\x02\'\"\"/cost-analysis/v1/cost-report/list:\x01*\x12\x82\x01\n\x04stat\x12\x32.spaceone.api.cost_analysis.v1.CostReportStatQuery\x1a\x17.google.protobuf.Struct\"-\x82\xd3\xe4\x93\x02\'\"\"/cost-analysis/v1/cost-report/stat:\x01*BDZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.cost_analysis.v1.cost_report_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1'
  _globals['_COSTREPORT'].methods_by_name['send']._options = None
  _globals['_COSTREPORT'].methods_by_name['send']._serialized_options = b'\202\323\344\223\002\'\"\"/cost-analysis/v1/cost-report/send:\001*'
  _globals['_COSTREPORT'].methods_by_name['get_url']._options = None
  _globals['_COSTREPORT'].methods_by_name['get_url']._serialized_options = b'\202\323\344\223\002*\"%/cost-analysis/v1/cost-report/get-url:\001*'
  _globals['_COSTREPORT'].methods_by_name['get']._options = None
  _globals['_COSTREPORT'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002&\"!/cost-analysis/v1/cost-report/get:\001*'
  _globals['_COSTREPORT'].methods_by_name['list']._options = None
  _globals['_COSTREPORT'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002\'\"\"/cost-analysis/v1/cost-report/list:\001*'
  _globals['_COSTREPORT'].methods_by_name['stat']._options = None
  _globals['_COSTREPORT'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\'\"\"/cost-analysis/v1/cost-report/stat:\001*'
  _globals['_COSTREPORTREQUEST']._serialized_start=246
  _globals['_COSTREPORTREQUEST']._serialized_end=289
  _globals['_GETURLCOSTREPORTREQUEST']._serialized_start=291
  _globals['_GETURLCOSTREPORTREQUEST']._serialized_end=340
  _globals['_COSTREPORTQUERY']._serialized_start=343
  _globals['_COSTREPORTQUERY']._serialized_end=593
  _globals['_COSTREPORTQUERY_STATUS']._serialized_start=545
  _globals['_COSTREPORTQUERY_STATUS']._serialized_end=593
  _globals['_COSTREPORTINFO']._serialized_start=596
  _globals['_COSTREPORTINFO']._serialized_end=1020
  _globals['_COSTREPORTINFO_STATUS']._serialized_start=545
  _globals['_COSTREPORTINFO_STATUS']._serialized_end=593
  _globals['_COSTREPORTSINFO']._serialized_start=1022
  _globals['_COSTREPORTSINFO']._serialized_end=1124
  _globals['_COSTREPORTSTATQUERY']._serialized_start=1126
  _globals['_COSTREPORTSTATQUERY']._serialized_end=1225
  _globals['_COSTREPORT']._serialized_start=1228
  _globals['_COSTREPORT']._serialized_end=1972
# @@protoc_insertion_point(module_scope)
