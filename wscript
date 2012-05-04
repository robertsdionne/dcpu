#!/usr/bin/env python
# encoding: utf-8

import sys

APPNAME = 'dcpu'
VERSION = '0.1'

def options(ctx):
  ctx.load('compiler_cxx')

def configure(ctx):
  ctx.load('compiler_cxx')
  ctx.env.LIB_NCURSES = ['ncurses']
  ctx.env.LIB_PTHREAD = ['pthread']
  if sys.platform.startswith('darwin'):
    ctx.env.LIBPATH_PTHREAD = ['/usr/lib']
  else:
    ctx.env.LIBPATH_PTHREAD = ['/usr/lib/x86_64-linux-gnu']

def build(ctx):
  ctx.stlib(
      source = [
        'libraries/gtest-1.6.0/src/gtest-all.cc',
        ],
      target = 'gtest',
      includes = [
        'libraries/gtest-1.6.0/include',
        'libraries/gtest-1.6.0',
        ],
      use = [
        'PTHREAD',
        ])
  ctx.stlib(
      source = [
        'libraries/protobuf-2.4.1/src/google/protobuf/descriptor.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/descriptor_database.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/descriptor.pb.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/dynamic_message.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/extension_set.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/extension_set_heavy.cc',
        ('libraries/protobuf-2.4.1/src/google/protobuf/' +
            'generated_message_reflection.cc'),
        'libraries/protobuf-2.4.1/src/google/protobuf/generated_message_util.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/message.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/message_lite.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/reflection_ops.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/repeated_field.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/service.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/text_format.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/unknown_field_set.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/wire_format.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/wire_format_lite.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/io/coded_stream.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/io/gzip_stream.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/io/printer.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/io/tokenizer.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/io/zero_copy_stream.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/io/zero_copy_stream_impl.cc',
        ('libraries/protobuf-2.4.1/src/google/protobuf/io/' +
            'zero_copy_stream_impl_lite.cc'),
        'libraries/protobuf-2.4.1/src/google/protobuf/stubs/common.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/stubs/once.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/stubs/structurally_valid.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/stubs/strutil.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/stubs/substitute.cc',
        ],
      target = 'protobuf',
      includes = [
        'libraries/protobuf-2.4.1/src/',
        'libraries/protobuf-2.4.1/',
        ],
      use = [
        'PTHREAD',
        ],
      cxxflags = [
          '-g',
        ])
  ctx.program(
      source = [
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/code_generator.cc',
        ('libraries/protobuf-2.4.1/src/google/protobuf/compiler/' +
            'command_line_interface.cc'),
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/importer.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/main.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/parser.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/plugin.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/plugin.pb.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/subprocess.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/zip_writer.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/cpp/cpp_enum.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/cpp/cpp_enum_field.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/cpp/cpp_extension.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/cpp/cpp_field.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/cpp/cpp_file.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/cpp/cpp_generator.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/cpp/cpp_helpers.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/cpp/cpp_message.cc',
        ('libraries/protobuf-2.4.1/src/google/protobuf/compiler/cpp/' +
            'cpp_message_field.cc'),
        ('libraries/protobuf-2.4.1/src/google/protobuf/compiler/cpp/' +
            'cpp_primitive_field.cc'),
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/cpp/cpp_service.cc',
        ('libraries/protobuf-2.4.1/src/google/protobuf/compiler/cpp/' +
            'cpp_string_field.cc'),
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/java/java_enum.cc',
        ('libraries/protobuf-2.4.1/src/google/protobuf/compiler/java/' +
            'java_enum_field.cc'),
        ('libraries/protobuf-2.4.1/src/google/protobuf/compiler/java/' +
            'java_extension.cc'),
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/java/java_field.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/java/java_file.cc',
        ('libraries/protobuf-2.4.1/src/google/protobuf/compiler/java/' +
            'java_generator.cc'),
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/java/java_helpers.cc',
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/java/java_message.cc',
        ('libraries/protobuf-2.4.1/src/google/protobuf/compiler/java/' +
            'java_message_field.cc'),
        ('libraries/protobuf-2.4.1/src/google/protobuf/compiler/java/' +
            'java_primitive_field.cc'),
        'libraries/protobuf-2.4.1/src/google/protobuf/compiler/java/java_service.cc',
        ('libraries/protobuf-2.4.1/src/google/protobuf/compiler/java/' +
            'java_string_field.cc'),
        ('libraries/protobuf-2.4.1/src/google/protobuf/compiler/python/' +
            'python_generator.cc'),
        ],
      target = 'protoc',
      includes = [
        'libraries/protobuf-2.4.1/src/',
        'libraries/protobuf-2.4.1/',
        ],
      use = [
          'protobuf',
          'PTHREAD',
        ],
      cxxflags = [
          '-g',
        ])
  ctx.program(
      source = [
        'libraries/gtest-1.6.0/src/gtest_main.cc',
        ],
      target = 'tests',
      includes = [
        'libraries/gtest-1.6.0/include',
        'source',
        ],
      use = [
        'dcpu_tests',
        'gtest',
        ])
  ctx.recurse('source')
  ctx.add_post_fun(test)

def run(ctx):
  ctx.exec_command('build/dcpu')

def test(ctx):
  ctx.exec_command('build/tests')

def debug(ctx):
  ctx.exec_command('gdb build/dcpu')

def valgrind(ctx):
  ctx.exec_command('valgrind build/dcpu')
