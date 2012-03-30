#!/usr/bin/env python
# encoding: utf-8

APPNAME = 'dcpu'
VERSION = '0.1'

def options(ctx):
  ctx.load('compiler_cxx')

def configure(ctx):
  ctx.load('compiler_cxx')

def build(ctx):
  ctx.recurse('src')
