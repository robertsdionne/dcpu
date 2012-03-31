// Copyright 2012 Robert Scott Dionne. All rights reserved.

#include <algorithm>
#include "dcpu.h"

const unsigned int Dcpu::kMemorySize;
const Dcpu::Word Dcpu::kVideoMemoryBegin;
const Dcpu::Word Dcpu::kVideoMemoryEnd;
const Dcpu::Word Dcpu::kOpcodeMask;
const Dcpu::Word Dcpu::kOperandMaskA;
const Dcpu::Word Dcpu::kOperandMaskB;
const Dcpu::Word Dcpu::kOperandShiftA;
const Dcpu::Word Dcpu::kOperandShiftB;

Dcpu::Word Dcpu::Instruct(const Dcpu::Word opcode,
    const Dcpu::Word operand_a, const Dcpu::Word operand_b) {
  return opcode | (operand_a << kOperandShiftA) | (operand_b << kOperandShiftB);
}

Dcpu::Dcpu()
  : register_a_(0), register_b_(0), register_c_(0), register_x_(0),
    register_y_(0), register_z_(0), register_i_(0), register_j_(0),
    program_counter_(0), stack_pointer_(0), overflow_(0)
{
  std::fill(memory_begin(), memory_end(), 0);
}

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

void Dcpu::ExecuteCycle(const bool skip) {
  const Word stack_pointer_backup = stack_pointer_;
  const Word instruction = *address(program_counter_);
  program_counter_ += 1;
  const Word opcode = instruction & kOpcodeMask;
  const Word operand_a = (instruction & kOperandMaskA) >> kOperandShiftA;
  const Word operand_b = (instruction & kOperandMaskB) >> kOperandShiftB;
  Word operand_a_literal = 0;
  Word *const operand_a_address = GetOperandAddressOrLiteral(
      operand_a, operand_a_literal);
  Word operand_b_literal = 0;
  const Word *const operand_b_address = GetOperandAddressOrLiteral(
      operand_b, operand_b_literal);
  const Word operand_a_value = operand_a_address ?
      *operand_a_address : operand_a_literal;
  const Word operand_b_value = operand_b_address ?
      *operand_b_address : operand_b_literal;
  if (skip) {
    stack_pointer_ = stack_pointer_backup;
    return;
  }
  unsigned int result;
  switch (opcode) {
    case kReserved:
      break;
    case kSet:
      MaybeAssignResult(operand_a_address, operand_b_value);
      break;
    case kAdd:
      result = operand_a_value + operand_b_value;
      overflow_ = result >> 16;
      MaybeAssignResult(operand_a_address, result);
      break;
    case kSubtract:
      overflow_ = operand_a_value < operand_b_value;
      MaybeAssignResult(operand_a_address, operand_a_value - operand_b_value);
      break;
    case kMultiply:
      result = operand_a_value * operand_b_value;
      overflow_ = result >> 16;
      MaybeAssignResult(operand_a_address, result);
      break;
    case kDivide:
      if (operand_b_value) {
        overflow_ = 0;
        MaybeAssignResult(operand_a_address, operand_a_value / operand_b_value);
      } else {
        overflow_ = 1;
        MaybeAssignResult(operand_a_address, 0);
      }
      break;
    case kModulo:
      MaybeAssignResult(operand_a_address, operand_a_value % operand_b_value);
      break;
    case kShiftLeft:
      result = operand_a_value << operand_b_value;
      overflow_ = result >> 16;
      MaybeAssignResult(operand_a_address, result);
      break;
    case kShiftRight:
      result = operand_a_value >> operand_b_value;
      overflow_ = operand_a_value & ((1 << operand_b_value) - 1);
      MaybeAssignResult(operand_a_address, result);
      break;
    case kBinaryAnd:
      MaybeAssignResult(operand_a_address, operand_a_value & operand_b_value);
      break;
    case kBinaryOr:
      MaybeAssignResult(operand_a_address, operand_a_value | operand_b_value);
      break;
    case kBinaryExclusiveOr:
      MaybeAssignResult(operand_a_address, operand_a_value ^ operand_b_value);
      break;
    case kIfEqual:
      if (operand_a_value != operand_b_value) {
        ExecuteCycle(/* skip */ true);
      }
      break;
    case kIfNotEqual:
      if (operand_a_value == operand_b_value) {
        ExecuteCycle(/* skip */ true);
      }
      break;
    case kIfGreaterThan:
      if (operand_a_value <= operand_b_value) {
        ExecuteCycle(/* skip */ true);
      }
      break;
    case kIfBoth:
      if (operand_a_value & operand_b_value == 0) {
        ExecuteCycle(/* skip */ true);
      }
      break;
    default:
      break;
  }
}

void Dcpu::ExecuteCycles(const unsigned long int count) {
  for (unsigned long int i = 0; i < count; ++i) {
    ExecuteCycle();
  }
}

void Dcpu::Reset() {
  std::fill(memory_begin(), memory_end(), 0);
  register_a_ = register_b_ = register_c_ = 0;
  register_x_ = register_y_ = register_z_ = 0;
  register_i_ = register_j_ = 0;
  program_counter_ = 0;
  stack_pointer_ = 0;
  overflow_ = 0;
}

Dcpu::Word *Dcpu::register_address(const Word register_index) {
  return &register_a_ + register_index % kLocationInRegisterA;
}

Dcpu::Word Dcpu::register_value(const Word register_index) {
  return *register_address(register_index);
}

Dcpu::Word *Dcpu::GetOperandAddressOrLiteral(
    const Word operand, Word &literal) {
  if (operand < kLocationInRegisterA) {
    return register_address(operand);
  } else if (kLocationInRegisterA <= operand
      && operand < kLocationOffsetByRegisterA) {
    return address(register_value(operand));
  } else if (kLocationOffsetByRegisterA <= operand && operand < kPop) {
    Word *const result =
        address(*address(program_counter_)) + register_value(operand);
    program_counter_ += 1;
    return result;
  } else if (operand == kPop) {
    Word *result = address(stack_pointer_);
    stack_pointer_ += 1;
    return result;
  } else if (operand == kPeek) {
    return address(stack_pointer_);
  } else if (operand == kPush) {
    stack_pointer_ -= 1;
    return address(stack_pointer_);
  } else if (operand == kStackPointer) {
    return &stack_pointer_;
  } else if (operand == kProgramCounter) {
    return &program_counter_;
  } else if (operand == kOverflow) {
    return &overflow_;
  } else if (operand == kLocation) {
    Word *const result = address(*address(program_counter_));
    program_counter_ += 1;
    return result;
  } else if (operand == kLiteral) {
    Word *const result = address(program_counter_);
    program_counter_ += 1;
    return result;
  } else {
    literal = operand - k0;
    return 0;
  }
}

void Dcpu::MaybeAssignResult(Word *const slot, const unsigned int result) {
  if (slot) {
    *slot = static_cast<Word>(result);
  }
}
