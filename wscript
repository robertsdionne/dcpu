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
