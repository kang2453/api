# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/identity/v2/domain.proto
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
from spaceone.api.core.v2 import handler_pb2 as spaceone_dot_api_dot_core_dot_v2_dot_handler__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n%spaceone/api/identity/v2/domain.proto\x12\x18spaceone.api.identity.v2\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\x1a\"spaceone/api/core/v2/handler.proto\"\x92\x01\n\x05\x41\x64min\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x10\n\x08password\x18\x03 \x01(\t\x12\r\n\x05\x65mail\x18\x04 \x01(\t\x12\x10\n\x08language\x18\x05 \x01(\t\x12\x10\n\x08timezone\x18\x06 \x01(\t\x12%\n\x04tags\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\"z\n\x13\x43reateDomainRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12.\n\x05\x61\x64min\x18\x02 \x01(\x0b\x32\x1f.spaceone.api.identity.v2.Admin\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"]\n\x13UpdateDomainRequest\x12\x11\n\tdomain_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"\"\n\rDomainRequest\x12\x11\n\tdomain_id\x18\x01 \x01(\t\"$\n\x14GetDomainAuthRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"\xd0\x01\n\x11\x44omainSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12@\n\x05state\x18\x04 \x01(\x0e\x32\x31.spaceone.api.identity.v2.DomainSearchQuery.State\",\n\x05State\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"\xef\x01\n\x0e\x44omainAuthInfo\x12\x11\n\tdomain_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12W\n\x13\x65xternal_auth_state\x18\x03 \x01(\x0e\x32:.spaceone.api.identity.v2.DomainAuthInfo.ExternalAuthState\x12)\n\x08metadata\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\"8\n\x11\x45xternalAuthState\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"\xe5\x01\n\nDomainInfo\x12\x11\n\tdomain_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x39\n\x05state\x18\x03 \x01(\x0e\x32*.spaceone.api.identity.v2.DomainInfo.State\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\ndeleted_at\x18  \x01(\t\",\n\x05State\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"Y\n\x0b\x44omainsInfo\x12\x35\n\x07results\x18\x01 \x03(\x0b\x32$.spaceone.api.identity.v2.DomainInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"G\n\x0f\x44omainStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\xdd\t\n\x06\x44omain\x12\x84\x01\n\x06\x63reate\x12-.spaceone.api.identity.v2.CreateDomainRequest\x1a$.spaceone.api.identity.v2.DomainInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/identity/v2/domain/create:\x01*\x12\x84\x01\n\x06update\x12-.spaceone.api.identity.v2.UpdateDomainRequest\x1a$.spaceone.api.identity.v2.DomainInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/identity/v2/domain/update:\x01*\x12p\n\x06\x64\x65lete\x12\'.spaceone.api.identity.v2.DomainRequest\x1a\x16.google.protobuf.Empty\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/identity/v2/domain/delete:\x01*\x12~\n\x06\x65nable\x12\'.spaceone.api.identity.v2.DomainRequest\x1a$.spaceone.api.identity.v2.DomainInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/identity/v2/domain/enable:\x01*\x12\x80\x01\n\x07\x64isable\x12\'.spaceone.api.identity.v2.DomainRequest\x1a$.spaceone.api.identity.v2.DomainInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/identity/v2/domain/disable:\x01*\x12x\n\x03get\x12\'.spaceone.api.identity.v2.DomainRequest\x1a$.spaceone.api.identity.v2.DomainInfo\"\"\x82\xd3\xe4\x93\x02\x1c\"\x17/identity/v2/domain/get:\x01*\x12\x97\x01\n\rget_auth_info\x12..spaceone.api.identity.v2.GetDomainAuthRequest\x1a(.spaceone.api.identity.v2.DomainAuthInfo\",\x82\xd3\xe4\x93\x02&\"!/identity/v2/domain/get-auth-info:\x01*\x12m\n\x0eget_public_key\x12+.spaceone.api.core.v2.AuthenticationRequest\x1a,.spaceone.api.core.v2.AuthenticationResponse\"\x00\x12\x7f\n\x04list\x12+.spaceone.api.identity.v2.DomainSearchQuery\x1a%.spaceone.api.identity.v2.DomainsInfo\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/identity/v2/domain/list:\x01*\x12L\n\x04stat\x12).spaceone.api.identity.v2.DomainStatQuery\x1a\x17.google.protobuf.Struct\"\x00\x42?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.identity.v2.domain_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2'
  _DOMAIN.methods_by_name['create']._options = None
  _DOMAIN.methods_by_name['create']._serialized_options = b'\202\323\344\223\002\037\"\032/identity/v2/domain/create:\001*'
  _DOMAIN.methods_by_name['update']._options = None
  _DOMAIN.methods_by_name['update']._serialized_options = b'\202\323\344\223\002\037\"\032/identity/v2/domain/update:\001*'
  _DOMAIN.methods_by_name['delete']._options = None
  _DOMAIN.methods_by_name['delete']._serialized_options = b'\202\323\344\223\002\037\"\032/identity/v2/domain/delete:\001*'
  _DOMAIN.methods_by_name['enable']._options = None
  _DOMAIN.methods_by_name['enable']._serialized_options = b'\202\323\344\223\002\037\"\032/identity/v2/domain/enable:\001*'
  _DOMAIN.methods_by_name['disable']._options = None
  _DOMAIN.methods_by_name['disable']._serialized_options = b'\202\323\344\223\002 \"\033/identity/v2/domain/disable:\001*'
  _DOMAIN.methods_by_name['get']._options = None
  _DOMAIN.methods_by_name['get']._serialized_options = b'\202\323\344\223\002\034\"\027/identity/v2/domain/get:\001*'
  _DOMAIN.methods_by_name['get_auth_info']._options = None
  _DOMAIN.methods_by_name['get_auth_info']._serialized_options = b'\202\323\344\223\002&\"!/identity/v2/domain/get-auth-info:\001*'
  _DOMAIN.methods_by_name['list']._options = None
  _DOMAIN.methods_by_name['list']._serialized_options = b'\202\323\344\223\002\035\"\030/identity/v2/domain/list:\001*'
  _globals['_ADMIN']._serialized_start=227
  _globals['_ADMIN']._serialized_end=373
  _globals['_CREATEDOMAINREQUEST']._serialized_start=375
  _globals['_CREATEDOMAINREQUEST']._serialized_end=497
  _globals['_UPDATEDOMAINREQUEST']._serialized_start=499
  _globals['_UPDATEDOMAINREQUEST']._serialized_end=592
  _globals['_DOMAINREQUEST']._serialized_start=594
  _globals['_DOMAINREQUEST']._serialized_end=628
  _globals['_GETDOMAINAUTHREQUEST']._serialized_start=630
  _globals['_GETDOMAINAUTHREQUEST']._serialized_end=666
  _globals['_DOMAINSEARCHQUERY']._serialized_start=669
  _globals['_DOMAINSEARCHQUERY']._serialized_end=877
  _globals['_DOMAINSEARCHQUERY_STATE']._serialized_start=833
  _globals['_DOMAINSEARCHQUERY_STATE']._serialized_end=877
  _globals['_DOMAINAUTHINFO']._serialized_start=880
  _globals['_DOMAINAUTHINFO']._serialized_end=1119
  _globals['_DOMAINAUTHINFO_EXTERNALAUTHSTATE']._serialized_start=1063
  _globals['_DOMAINAUTHINFO_EXTERNALAUTHSTATE']._serialized_end=1119
  _globals['_DOMAININFO']._serialized_start=1122
  _globals['_DOMAININFO']._serialized_end=1351
  _globals['_DOMAININFO_STATE']._serialized_start=833
  _globals['_DOMAININFO_STATE']._serialized_end=877
  _globals['_DOMAINSINFO']._serialized_start=1353
  _globals['_DOMAINSINFO']._serialized_end=1442
  _globals['_DOMAINSTATQUERY']._serialized_start=1444
  _globals['_DOMAINSTATQUERY']._serialized_end=1515
  _globals['_DOMAIN']._serialized_start=1518
  _globals['_DOMAIN']._serialized_end=2763
# @@protoc_insertion_point(module_scope)
