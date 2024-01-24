# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.cost_analysis.v1 import cost_report_pb2 as spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2


class CostReportStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.send = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.CostReport/send',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportInfo.FromString,
                )
        self.get_url = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.CostReport/get_url',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.GetUrlCostReportRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                )
        self.get = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.CostReport/get',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportInfo.FromString,
                )
        self.list = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.CostReport/list',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportsInfo.FromString,
                )
        self.stat = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.CostReport/stat',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                )


class CostReportServicer(object):
    """Missing associated documentation comment in .proto file."""

    def send(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get_url(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def list(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def stat(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_CostReportServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'send': grpc.unary_unary_rpc_method_handler(
                    servicer.send,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportInfo.SerializeToString,
            ),
            'get_url': grpc.unary_unary_rpc_method_handler(
                    servicer.get_url,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.GetUrlCostReportRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportsInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.cost_analysis.v1.CostReport', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class CostReport(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def send(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.cost_analysis.v1.CostReport/send',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportRequest.SerializeToString,
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def get_url(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.cost_analysis.v1.CostReport/get_url',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.GetUrlCostReportRequest.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.cost_analysis.v1.CostReport/get',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportRequest.SerializeToString,
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def list(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.cost_analysis.v1.CostReport/list',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportQuery.SerializeToString,
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportsInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def stat(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.cost_analysis.v1.CostReport/stat',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__report__pb2.CostReportStatQuery.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
