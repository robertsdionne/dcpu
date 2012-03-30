// Copyright 2012 Robert Scott Dionne. All rights reserved.

#include "dcpu.h"

const unsigned int Dcpu::kMemorySize;
const Dcpu::Word Dcpu::kVideoMemoryBegin;
const Dcpu::Word Dcpu::kVideoMemoryEnd;
const Dcpu::Word Dcpu::kOpcodeMask;
const Dcpu::Word Dcpu::kOperandMaskA;
const Dcpu::Word Dcpu::kOperandMaskB;
const Dcpu::Word Dcpu::kOperandShiftA;
const Dcpu::Word Dcpu::kOperandShiftB;

Dcpu::Dcpu()
  : register_a_(0), register_b_(0), register_c_(0), register_x_(0),
    register_y_(0), register_z_(0), register_i_(0), register_j_(0),
    program_counter_(0), stack_pointer_(0), overflow_(0)
{}

Dcpu::Word *Dcpu::address(const Dcpu::Word address_value) {
  return memory_begin() + address_value;
}

const Dcpu::Word *Dcpu::address(const Dcpu::Word address_value) const {
  return memory_begin() + address_value;
}

Dcpu::Word *Dcpu::memory_begin() {
  return &memory_[0];
}

const Dcpu::Word *Dcpu::memory_begin() const {
  return &memory_[0];
}

Dcpu::Word *Dcpu::memory_end() {
  return memory_begin() + kMemorySize;
}

const Dcpu::Word *Dcpu::memory_end() const {
  return memory_begin() + kMemorySize;
}

Dcpu::Word *Dcpu::video_memory_begin() {
  return address(kVideoMemoryBegin);
}

const Dcpu::Word *Dcpu::video_memory_begin() const {
  return address(kVideoMemoryBegin);
}

Dcpu::Word *Dcpu::video_memory_end() {
  return address(kVideoMemoryEnd);
}

const Dcpu::Word *Dcpu::video_memory_end() const {
  return address(kVideoMemoryEnd);
}
