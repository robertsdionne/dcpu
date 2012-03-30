// Copyright 2012 Robert Scott Dionne. All rights reserved.

#include <algorithm>
#include "dcpu.h"

Dcpu::Dcpu() {}

unsigned short *Dcpu::address(unsigned short address_value) {
  return &memory_[address_value];
}

const unsigned short *Dcpu::address(unsigned short address_value) const {
  return &memory_[address_value];
}

unsigned short *Dcpu::begin() {
  return &memory_[0];
}

const unsigned short *Dcpu::begin() const {
  return &memory_[0];
}

unsigned short *Dcpu::end() {
  return begin() + 0x10000;
}

const unsigned short *Dcpu::end() const {
  return begin() + 0x10000;
}

void Dcpu::ReadVideoMemory(unsigned short *const out) const {
  const unsigned short *const begin = &memory_[0x8000];
  const unsigned short *const end = &memory_[0x8000+1000];
  std::copy(begin, end, out);
}

void Dcpu::WriteVideoMemory(
    unsigned short *const begin, unsigned short *const end) const {
}
