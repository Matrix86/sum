# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/sum.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/sum.proto',
  package='sum',
  syntax='proto3',
  serialized_pb=_b('\n\x0fproto/sum.proto\x12\x03sum\")\n\nNamedValue\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t\"A\n\x06Record\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x02 \x03(\x02\x12\x1d\n\x04meta\x18\x03 \x03(\x0b\x32\x0f.sum.NamedValue\"0\n\x06Oracle\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0c\n\x04\x63ode\x18\x03 \x01(\t\"?\n\x0b\x45valRequest\x12\x11\n\toracle_id\x18\x01 \x01(\t\x12\x1d\n\x04\x61rgs\x18\x02 \x03(\x0b\x32\x0f.sum.NamedValue\"\x13\n\x05Query\x12\n\n\x02id\x18\x01 \x01(\t\"E\n\x08Response\x12\x0f\n\x07success\x18\x01 \x01(\x08\x12\x0b\n\x03msg\x18\x02 \x01(\t\x12\x1b\n\x06record\x18\x03 \x01(\x0b\x32\x0b.sum.Record\"f\n\nServerInfo\x12\x0f\n\x07version\x18\x01 \x01(\t\x12\x0e\n\x06uptime\x18\x02 \x01(\x04\x12\x0b\n\x03pid\x18\x03 \x01(\x04\x12\x0b\n\x03uid\x18\x04 \x01(\x04\x12\x0c\n\x04\x61rgv\x18\x05 \x03(\t\x12\x0f\n\x07records\x18\x06 \x01(\x04\"\x07\n\x05\x45mpty2\xcf\x01\n\nSumService\x12&\n\x06\x43reate\x12\x0b.sum.Record\x1a\r.sum.Response\"\x00\x12&\n\x06Update\x12\x0b.sum.Record\x1a\r.sum.Response\"\x00\x12#\n\x04Read\x12\n.sum.Query\x1a\r.sum.Response\"\x00\x12%\n\x06\x44\x65lete\x12\n.sum.Query\x1a\r.sum.Response\"\x00\x12%\n\x04Info\x12\n.sum.Empty\x1a\x0f.sum.ServerInfo\"\x00\x62\x06proto3')
)




_NAMEDVALUE = _descriptor.Descriptor(
  name='NamedValue',
  full_name='sum.NamedValue',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='sum.NamedValue.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='sum.NamedValue.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=24,
  serialized_end=65,
)


_RECORD = _descriptor.Descriptor(
  name='Record',
  full_name='sum.Record',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='sum.Record.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data', full_name='sum.Record.data', index=1,
      number=2, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='meta', full_name='sum.Record.meta', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=67,
  serialized_end=132,
)


_ORACLE = _descriptor.Descriptor(
  name='Oracle',
  full_name='sum.Oracle',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='sum.Oracle.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='sum.Oracle.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='code', full_name='sum.Oracle.code', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=134,
  serialized_end=182,
)


_EVALREQUEST = _descriptor.Descriptor(
  name='EvalRequest',
  full_name='sum.EvalRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='oracle_id', full_name='sum.EvalRequest.oracle_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='args', full_name='sum.EvalRequest.args', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=184,
  serialized_end=247,
)


_QUERY = _descriptor.Descriptor(
  name='Query',
  full_name='sum.Query',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='sum.Query.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=249,
  serialized_end=268,
)


_RESPONSE = _descriptor.Descriptor(
  name='Response',
  full_name='sum.Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='success', full_name='sum.Response.success', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='msg', full_name='sum.Response.msg', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='record', full_name='sum.Response.record', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=270,
  serialized_end=339,
)


_SERVERINFO = _descriptor.Descriptor(
  name='ServerInfo',
  full_name='sum.ServerInfo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='version', full_name='sum.ServerInfo.version', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='uptime', full_name='sum.ServerInfo.uptime', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='pid', full_name='sum.ServerInfo.pid', index=2,
      number=3, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='uid', full_name='sum.ServerInfo.uid', index=3,
      number=4, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='argv', full_name='sum.ServerInfo.argv', index=4,
      number=5, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='records', full_name='sum.ServerInfo.records', index=5,
      number=6, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=341,
  serialized_end=443,
)


_EMPTY = _descriptor.Descriptor(
  name='Empty',
  full_name='sum.Empty',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=445,
  serialized_end=452,
)

_RECORD.fields_by_name['meta'].message_type = _NAMEDVALUE
_EVALREQUEST.fields_by_name['args'].message_type = _NAMEDVALUE
_RESPONSE.fields_by_name['record'].message_type = _RECORD
DESCRIPTOR.message_types_by_name['NamedValue'] = _NAMEDVALUE
DESCRIPTOR.message_types_by_name['Record'] = _RECORD
DESCRIPTOR.message_types_by_name['Oracle'] = _ORACLE
DESCRIPTOR.message_types_by_name['EvalRequest'] = _EVALREQUEST
DESCRIPTOR.message_types_by_name['Query'] = _QUERY
DESCRIPTOR.message_types_by_name['Response'] = _RESPONSE
DESCRIPTOR.message_types_by_name['ServerInfo'] = _SERVERINFO
DESCRIPTOR.message_types_by_name['Empty'] = _EMPTY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

NamedValue = _reflection.GeneratedProtocolMessageType('NamedValue', (_message.Message,), dict(
  DESCRIPTOR = _NAMEDVALUE,
  __module__ = 'proto.sum_pb2'
  # @@protoc_insertion_point(class_scope:sum.NamedValue)
  ))
_sym_db.RegisterMessage(NamedValue)

Record = _reflection.GeneratedProtocolMessageType('Record', (_message.Message,), dict(
  DESCRIPTOR = _RECORD,
  __module__ = 'proto.sum_pb2'
  # @@protoc_insertion_point(class_scope:sum.Record)
  ))
_sym_db.RegisterMessage(Record)

Oracle = _reflection.GeneratedProtocolMessageType('Oracle', (_message.Message,), dict(
  DESCRIPTOR = _ORACLE,
  __module__ = 'proto.sum_pb2'
  # @@protoc_insertion_point(class_scope:sum.Oracle)
  ))
_sym_db.RegisterMessage(Oracle)

EvalRequest = _reflection.GeneratedProtocolMessageType('EvalRequest', (_message.Message,), dict(
  DESCRIPTOR = _EVALREQUEST,
  __module__ = 'proto.sum_pb2'
  # @@protoc_insertion_point(class_scope:sum.EvalRequest)
  ))
_sym_db.RegisterMessage(EvalRequest)

Query = _reflection.GeneratedProtocolMessageType('Query', (_message.Message,), dict(
  DESCRIPTOR = _QUERY,
  __module__ = 'proto.sum_pb2'
  # @@protoc_insertion_point(class_scope:sum.Query)
  ))
_sym_db.RegisterMessage(Query)

Response = _reflection.GeneratedProtocolMessageType('Response', (_message.Message,), dict(
  DESCRIPTOR = _RESPONSE,
  __module__ = 'proto.sum_pb2'
  # @@protoc_insertion_point(class_scope:sum.Response)
  ))
_sym_db.RegisterMessage(Response)

ServerInfo = _reflection.GeneratedProtocolMessageType('ServerInfo', (_message.Message,), dict(
  DESCRIPTOR = _SERVERINFO,
  __module__ = 'proto.sum_pb2'
  # @@protoc_insertion_point(class_scope:sum.ServerInfo)
  ))
_sym_db.RegisterMessage(ServerInfo)

Empty = _reflection.GeneratedProtocolMessageType('Empty', (_message.Message,), dict(
  DESCRIPTOR = _EMPTY,
  __module__ = 'proto.sum_pb2'
  # @@protoc_insertion_point(class_scope:sum.Empty)
  ))
_sym_db.RegisterMessage(Empty)



_SUMSERVICE = _descriptor.ServiceDescriptor(
  name='SumService',
  full_name='sum.SumService',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=455,
  serialized_end=662,
  methods=[
  _descriptor.MethodDescriptor(
    name='Create',
    full_name='sum.SumService.Create',
    index=0,
    containing_service=None,
    input_type=_RECORD,
    output_type=_RESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Update',
    full_name='sum.SumService.Update',
    index=1,
    containing_service=None,
    input_type=_RECORD,
    output_type=_RESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Read',
    full_name='sum.SumService.Read',
    index=2,
    containing_service=None,
    input_type=_QUERY,
    output_type=_RESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Delete',
    full_name='sum.SumService.Delete',
    index=3,
    containing_service=None,
    input_type=_QUERY,
    output_type=_RESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Info',
    full_name='sum.SumService.Info',
    index=4,
    containing_service=None,
    input_type=_EMPTY,
    output_type=_SERVERINFO,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_SUMSERVICE)

DESCRIPTOR.services_by_name['SumService'] = _SUMSERVICE

# @@protoc_insertion_point(module_scope)