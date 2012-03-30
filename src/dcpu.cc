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

Dcpu::Word &Dcpu::register_a() {
  return register_a_;
}

Dcpu::Word Dcpu::register_a() const {
  return register_a_;
}

Dcpu::Word &Dcpu::register_b() {
  return register_b_;
}

Dcpu::Word Dcpu::register_b() const {
  return register_b_;
}

Dcpu::Word &Dcpu::register_c() {
  return register_c_;
}

Dcpu::Word Dcpu::register_c() const {
  return register_c_;
}

Dcpu::Word &Dcpu::register_x() {
  return register_x_;
}

Dcpu::Word Dcpu::register_x() const {
  return register_x_;
}

Dcpu::Word &Dcpu::register_y() {
  return register_y_;
}

Dcpu::Word Dcpu::register_y() const {
  return register_y_;
}

Dcpu::Word &Dcpu::register_z() {
  return register_z_;
}

Dcpu::Word Dcpu::register_z() const {
  return register_z_;
}

Dcpu::Word &Dcpu::register_i() {
  return register_i_;
}

Dcpu::Word Dcpu::register_i() const {
  return register_i_;
}

Dcpu::Word &Dcpu::register_j() {
  return register_j_;
}

Dcpu::Word Dcpu::register_j() const {
  return register_j_;
}

Dcpu::Word &Dcpu::program_counter() {
  return program_counter_;
}

Dcpu::Word Dcpu::program_counter() const {
  return program_counter_;
}

Dcpu::Word &Dcpu::stack_pointer() {
  return stack_pointer_;
}

Dcpu::Word Dcpu::stack_pointer() const {
  return stack_pointer_;
}

Dcpu::Word &Dcpu::overflow() {
  return overflow_;
}

Dcpu::Word Dcpu::overflow() const {
  return overflow_;
}
