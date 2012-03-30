#!/usr/bin/env python
# encoding: utf-8

APPNAME = 'dcpu'
VERSION = '0.1'

def options(ctx):
  ctx.load('compiler_cxx')

def configure(ctx):
  ctx.load('compiler_cxx')
  ctx.env.LIB_NCURSES = ['ncurses']

def build(ctx):
  ctx.recurse('src')

def run(ctx):
  ctx.exec_command('build/dcpu')

def debug(ctx):
  ctx.exec_command('gdb build/dcpu')

def valgrind(ctx):
  ctx.exec_command('valgrind build/dcpu')
