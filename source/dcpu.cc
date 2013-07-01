// Copyright 2012 Robert Scott Dionne. All rights reserved.

#include <algorithm>
#include "dcpu.h"

const unsigned int Dcpu::kMemorySize;
const Dcpu::Word Dcpu::kBasicOpcodeMask;
const Dcpu::Word Dcpu::kAdvancedOpcodeMask;
const Dcpu::Word Dcpu::kBasicOperandMaskA;
const Dcpu::Word Dcpu::kBasicOperandMaskB;
const Dcpu::Word Dcpu::kBasicOperandShiftA;
const Dcpu::Word Dcpu::kBasicOperandShiftB;
const Dcpu::Word Dcpu::kAdvancedOperandMaskA;
const Dcpu::Word Dcpu::kAdvancedOperandShiftA;

Dcpu::Word Dcpu::Noop() {
  return Dcpu::Instruct(BasicOpcode::kSet, kRegisterA, kRegisterA);
}

Dcpu::Word Dcpu::Instruct(const Dcpu::BasicOpcode basic_opcode,
    const Dcpu::Operand operand_b, const Dcpu::Operand operand_a) {
  return static_cast<int>(basic_opcode) | (
      operand_a << kBasicOperandShiftA) | (operand_b << kBasicOperandShiftB);
}

Dcpu::Word Dcpu::Instruct(
    const Dcpu::AdvancedOpcode advanced_opcode, const Dcpu::Operand operand_a) {
  return (static_cast<int>(advanced_opcode) << kAdvancedOpcodeShift) | (
      operand_a << kAdvancedOperandShiftA);
}

Dcpu::Dcpu()
  : register_a_(0), register_b_(0), register_c_(0), register_x_(0),
    register_y_(0), register_z_(0), register_i_(0), register_j_(0),
    program_counter_(0), stack_pointer_(0), extra_(0), interrupt_address_(0)
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

Dcpu::Word &Dcpu::extra() {
  return extra_;
}

Dcpu::Word Dcpu::extra() const {
  return extra_;
}

Dcpu::Word &Dcpu::interrupt_address() {
  return interrupt_address_;
}

Dcpu::Word Dcpu::interrupt_address() const {
  return interrupt_address_;
}

void Dcpu::Interrupt(const Word message) {
  if (interrupt_address_) {
    stack_pointer_ -= 1;
    *address(stack_pointer_) = program_counter_;
    stack_pointer_ -= 1;
    *address(stack_pointer_) = register_a_;
    program_counter_ = interrupt_address_;
    register_a_ = message;
  }
}

void Dcpu::ExecuteInstruction(const bool skip) {
  const Word stack_pointer_backup = stack_pointer_;
  const Word instruction = *address(program_counter_);
  program_counter_ += 1;
  const BasicOpcode basic_opcode = static_cast<BasicOpcode>(instruction & kBasicOpcodeMask);

  if (basic_opcode == BasicOpcode::kBasicReserved) {
    const AdvancedOpcode advanced_opcode = static_cast<AdvancedOpcode>((
        instruction & kAdvancedOpcodeMask) >> kAdvancedOpcodeShift);
    const Operand operand_a = static_cast<Operand>(
        (instruction & kAdvancedOperandMaskA) >> kAdvancedOperandShiftA);
    const bool assignable = advanced_opcode == AdvancedOpcode::kInterruptAddressGet ||
        advanced_opcode == AdvancedOpcode::kHardwareNumberConnected;
    Word operand_a_literal = 0;
    Word *const operand_a_address = GetOperandAddressOrLiteral(
        operand_a, assignable, operand_a_literal);
    const Word operand_a_value = operand_a_address ?
      *operand_a_address : operand_a_literal;
    if (skip) {
      stack_pointer_ = stack_pointer_backup;
      return;
    }
    switch (advanced_opcode) {
      case AdvancedOpcode::kAdvancedReserved:
        break;
      case AdvancedOpcode::kJumpSubRoutine:
        stack_pointer_ -= 1;
        *address(stack_pointer_) = program_counter_;
        program_counter_ = operand_a_value;
        break;
      case AdvancedOpcode::kInterruptTrigger:
        Interrupt(operand_a_value);
        break;
      case AdvancedOpcode::kInterruptAddressGet:
        MaybeAssignResult(operand_a_address, interrupt_address_);
        break;
      case AdvancedOpcode::kInterruptAddressSet:
        interrupt_address_ = operand_a_value;
        break;
      case AdvancedOpcode::kHardwareNumberConnected:
        MaybeAssignResult(operand_a_address, 0);
        break;
      case AdvancedOpcode::kHardwareQuery:
        register_a_ = 0;
        register_b_ = 0;
        register_c_ = 0;
        register_x_ = 0;
        register_y_ = 0;
        break;
      case AdvancedOpcode::kHardwareInterrupt:
        break;
      default:
        break;
    }
  } else {
    const Operand operand_a = static_cast<Operand>(
        (instruction & kBasicOperandMaskA) >> kBasicOperandShiftA);
    const Operand operand_b = static_cast<Operand>(
        (instruction & kBasicOperandMaskB) >> kBasicOperandShiftB);
    Word operand_a_literal = 0;
    const Word *const operand_a_address = GetOperandAddressOrLiteral(
        operand_a, /* assignable */ false, operand_a_literal);
    Word operand_b_literal = 0;
    Word *const operand_b_address = GetOperandAddressOrLiteral(
        operand_b, /* assignable */ true, operand_b_literal);
    const Word operand_a_value = operand_a_address ?
        *operand_a_address : operand_a_literal;
    const Word operand_b_value = operand_b_address ?
        *operand_b_address : operand_b_literal;
    if (skip) {
      stack_pointer_ = stack_pointer_backup;
      return;
    }
    unsigned int result;
    signed int signed_result;
    switch (basic_opcode) {
      case BasicOpcode::kBasicReserved:
        break;
      case BasicOpcode::kSet:
        MaybeAssignResult(operand_b_address, operand_a_value);
        break;
      case BasicOpcode::kAdd:
        result = operand_b_value + operand_a_value;
        extra_ = result >> 16;
        MaybeAssignResult(operand_b_address, result);
        break;
      case BasicOpcode::kSubtract:
        extra_ = operand_b_value < operand_a_value;
        MaybeAssignResult(operand_b_address, operand_b_value - operand_a_value);
        break;
      case BasicOpcode::kMultiply:
        result = operand_b_value * operand_a_value;
        extra_ = result >> 16;
        MaybeAssignResult(operand_b_address, result);
        break;
      case BasicOpcode::kMultiplySigned:
        signed_result = static_cast<SignedWord>(operand_b_value) *
            static_cast<SignedWord>(operand_a_value);
        extra_ = signed_result >> 16;
        MaybeAssignResult(operand_b_address, static_cast<Word>(signed_result));
        break;
      case BasicOpcode::kDivide:
        if (operand_a_value) {
          extra_ = 0;
          MaybeAssignResult(
              operand_b_address, operand_b_value / operand_a_value);
        } else {
          extra_ = 1;
          MaybeAssignResult(operand_b_address, 0);
        }
        break;
      case BasicOpcode::kDivideSigned:
        if (operand_a_value) {
          extra_ = 0;
          MaybeAssignResult(
              operand_b_address, static_cast<Word>(
                  static_cast<SignedWord>(operand_b_value) /
                      static_cast<SignedWord>(operand_a_value)));
        } else {
          extra_ = 1;
          MaybeAssignResult(operand_b_address, 0);
        }
        break;
      case BasicOpcode::kModulo:
        MaybeAssignResult(operand_b_address, operand_b_value % operand_a_value);
        break;
      case BasicOpcode::kBinaryAnd:
        MaybeAssignResult(operand_b_address, operand_b_value & operand_a_value);
        break;
      case BasicOpcode::kBinaryOr:
        MaybeAssignResult(operand_b_address, operand_b_value | operand_a_value);
        break;
      case BasicOpcode::kBinaryExclusiveOr:
        MaybeAssignResult(operand_b_address, operand_b_value ^ operand_a_value);
        break;
      case BasicOpcode::kShiftRight:
        result = operand_b_value >> operand_a_value;
        extra_ = operand_b_value << (0x10 - operand_a_value);
        MaybeAssignResult(operand_b_address, result);
        break;
      case BasicOpcode::kArithmeticShiftRight:
        signed_result = static_cast<SignedWord>(operand_b_value)
            >> operand_a_value;
        extra_ = operand_b_value << (0x10 - operand_a_value);
        MaybeAssignResult(operand_b_address, static_cast<Word>(signed_result));
        break;
      case BasicOpcode::kShiftLeft:
        result = operand_b_value << operand_a_value;
        extra_ = result >> 16;
        MaybeAssignResult(operand_b_address, result);
        break;
      case BasicOpcode::kIfBitSet:
        if ((operand_b_value & operand_a_value) == 0) {
          ExecuteInstruction(/* skip */ true);
        }
        break;
      case BasicOpcode::kIfClear:
        if ((operand_b_value & operand_a_value) != 0) {
          ExecuteInstruction(/* skip */ true);
        }
        break;
      case BasicOpcode::kIfEqual:
        if (operand_b_value != operand_a_value) {
          ExecuteInstruction(/* skip */ true);
        }
        break;
      case BasicOpcode::kIfNotEqual:
        if (operand_b_value == operand_a_value) {
          ExecuteInstruction(/* skip */ true);
        }
        break;
      case BasicOpcode::kIfGreaterThan:
        if (operand_b_value <= operand_a_value) {
          ExecuteInstruction(/* skip */ true);
        }
        break;
      case BasicOpcode::kIfAbove:
        if (static_cast<SignedWord>(operand_b_value)
            <= static_cast<SignedWord>(operand_a_value)) {
          ExecuteInstruction(/* skip */ true);
        }
        break;
      case BasicOpcode::kIfLessThan:
        if (operand_b_value >= operand_a_value) {
          ExecuteInstruction(/* skip */ true);
        }
        break;
      case BasicOpcode::kIfUnder:
        if (static_cast<SignedWord>(operand_b_value)
            >= static_cast<SignedWord>(operand_a_value)) {
          ExecuteInstruction(/* skip */ true);
        }
        break;
      default:
        break;
    }
  }
}

void Dcpu::ExecuteInstructions(const unsigned long int count) {
  for (unsigned long int i = 0; i < count; ++i) {
    ExecuteInstruction();
  }
}

void Dcpu::Reset() {
  std::fill(memory_begin(), memory_end(), 0);
  register_a_ = register_b_ = register_c_ = 0;
  register_x_ = register_y_ = register_z_ = 0;
  register_i_ = register_j_ = 0;
  program_counter_ = 0;
  stack_pointer_ = 0;
  extra_ = 0;
  interrupt_address_ = 0;
}

Dcpu::Word *Dcpu::register_address(const Word register_index) {
  return &register_a_ + register_index % kLocationInRegisterA;
}

Dcpu::Word Dcpu::register_value(const Word register_index) {
  return *register_address(register_index);
}

Dcpu::Word *Dcpu::GetOperandAddressOrLiteral(
    const Operand operand, const bool assignable, Word &literal) {
  if (operand < kLocationInRegisterA) {
    return register_address(operand);
  } else if (kLocationInRegisterA <= operand
      && operand < kLocationOffsetByRegisterA) {
    return address(register_value(operand));
  } else if (kLocationOffsetByRegisterA <= operand && operand < kPushPop) {
    Word *const result =
        address(*address(program_counter_)) + register_value(operand);
    program_counter_ += 1;
    return result;
  } else if (operand == kPushPop) {
    if (assignable) { // Push
      stack_pointer_ -= 1;
      return address(stack_pointer_);
    } else { // Pop
      Word *const result = address(stack_pointer_);
      stack_pointer_ += 1;
      return result;
    }
  } else if (operand == kPeek) {
    return address(stack_pointer_);
  } else if (operand == kPick) {
    const Word offset = *address(program_counter_);
    program_counter_ += 1;
    return address(stack_pointer_ + offset);
  } else if (operand == kStackPointer) {
    return &stack_pointer_;
  } else if (operand == kProgramCounter) {
    return &program_counter_;
  } else if (operand == kExtra) {
    return &extra_;
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
