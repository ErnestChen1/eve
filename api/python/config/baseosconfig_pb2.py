# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: config/baseosconfig.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from config import devcommon_pb2 as config_dot_devcommon__pb2
from config import storage_pb2 as config_dot_storage__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='config/baseosconfig.proto',
  package='org.lfedge.eve.config',
  syntax='proto3',
  serialized_options=b'\n\025org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/config',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x19\x63onfig/baseosconfig.proto\x12\x15org.lfedge.eve.config\x1a\x16\x63onfig/devcommon.proto\x1a\x14\x63onfig/storage.proto\"1\n\tOSKeyTags\x12\x10\n\x08OSVerKey\x18\x01 \x01(\t\x12\x12\n\nOSVerValue\x18\x02 \x01(\t\"F\n\x0cOSVerDetails\x12\x36\n\x0c\x62\x61seOSParams\x18\x0c \x03(\x0b\x32 .org.lfedge.eve.config.OSKeyTags\"\xe0\x01\n\x0c\x42\x61seOSConfig\x12=\n\x0euuidandversion\x18\x01 \x01(\x0b\x32%.org.lfedge.eve.config.UUIDandVersion\x12,\n\x06\x64rives\x18\x03 \x03(\x0b\x32\x1c.org.lfedge.eve.config.Drive\x12\x10\n\x08\x61\x63tivate\x18\x04 \x01(\x08\x12\x15\n\rbaseOSVersion\x18\n \x01(\t\x12:\n\rbaseOSDetails\x18\x0b \x01(\x0b\x32#.org.lfedge.eve.config.OSVerDetailsB=\n\x15org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/configb\x06proto3'
  ,
  dependencies=[config_dot_devcommon__pb2.DESCRIPTOR,config_dot_storage__pb2.DESCRIPTOR,])




_OSKEYTAGS = _descriptor.Descriptor(
  name='OSKeyTags',
  full_name='org.lfedge.eve.config.OSKeyTags',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='OSVerKey', full_name='org.lfedge.eve.config.OSKeyTags.OSVerKey', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='OSVerValue', full_name='org.lfedge.eve.config.OSKeyTags.OSVerValue', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=98,
  serialized_end=147,
)


_OSVERDETAILS = _descriptor.Descriptor(
  name='OSVerDetails',
  full_name='org.lfedge.eve.config.OSVerDetails',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='baseOSParams', full_name='org.lfedge.eve.config.OSVerDetails.baseOSParams', index=0,
      number=12, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=149,
  serialized_end=219,
)


_BASEOSCONFIG = _descriptor.Descriptor(
  name='BaseOSConfig',
  full_name='org.lfedge.eve.config.BaseOSConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='uuidandversion', full_name='org.lfedge.eve.config.BaseOSConfig.uuidandversion', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='drives', full_name='org.lfedge.eve.config.BaseOSConfig.drives', index=1,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='activate', full_name='org.lfedge.eve.config.BaseOSConfig.activate', index=2,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='baseOSVersion', full_name='org.lfedge.eve.config.BaseOSConfig.baseOSVersion', index=3,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='baseOSDetails', full_name='org.lfedge.eve.config.BaseOSConfig.baseOSDetails', index=4,
      number=11, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=222,
  serialized_end=446,
)

_OSVERDETAILS.fields_by_name['baseOSParams'].message_type = _OSKEYTAGS
_BASEOSCONFIG.fields_by_name['uuidandversion'].message_type = config_dot_devcommon__pb2._UUIDANDVERSION
_BASEOSCONFIG.fields_by_name['drives'].message_type = config_dot_storage__pb2._DRIVE
_BASEOSCONFIG.fields_by_name['baseOSDetails'].message_type = _OSVERDETAILS
DESCRIPTOR.message_types_by_name['OSKeyTags'] = _OSKEYTAGS
DESCRIPTOR.message_types_by_name['OSVerDetails'] = _OSVERDETAILS
DESCRIPTOR.message_types_by_name['BaseOSConfig'] = _BASEOSCONFIG
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

OSKeyTags = _reflection.GeneratedProtocolMessageType('OSKeyTags', (_message.Message,), {
  'DESCRIPTOR' : _OSKEYTAGS,
  '__module__' : 'config.baseosconfig_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.config.OSKeyTags)
  })
_sym_db.RegisterMessage(OSKeyTags)

OSVerDetails = _reflection.GeneratedProtocolMessageType('OSVerDetails', (_message.Message,), {
  'DESCRIPTOR' : _OSVERDETAILS,
  '__module__' : 'config.baseosconfig_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.config.OSVerDetails)
  })
_sym_db.RegisterMessage(OSVerDetails)

BaseOSConfig = _reflection.GeneratedProtocolMessageType('BaseOSConfig', (_message.Message,), {
  'DESCRIPTOR' : _BASEOSCONFIG,
  '__module__' : 'config.baseosconfig_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.config.BaseOSConfig)
  })
_sym_db.RegisterMessage(BaseOSConfig)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
