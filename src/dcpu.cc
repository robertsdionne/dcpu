// Copyright 2012 Robert Scott Dionne. All rights reserved.

#include <algorithm>
#include "dcpu.h"

Dcpu::Dcpu() {}

Dcpu::Word *Dcpu::address(const Dcpu::Word address_value) {
  return begin() + address_value;
}

const Dcpu::Word *Dcpu::address(const Dcpu::Word address_value) const {
  return begin() + address_value;
}

Dcpu::Word *Dcpu::begin() {
  return &memory_[0];
}

const Dcpu::Word *Dcpu::begin() const {
  return &memory_[0];
}

Dcpu::Word *Dcpu::end() {
  return begin() + 0x10000;
}

const Dcpu::Word *Dcpu::end() const {
  return begin() + 0x10000;
}

void Dcpu::ReadVideoMemory(Dcpu::Word *const out) const {
  std::copy(address(0x8000), address(0x8000+1000), out);
}

void Dcpu::WriteVideoMemory(Dcpu::Word *const begin, Dcpu::Word *const end) {
  std::copy(begin, end, address(0x8000));
}
