#include <algorithm>

#include "dcpu.hpp"
#include "hardware.hpp"

namespace dcpu {

  Word Noop() {
    return Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kRegisterA);
  }

  Word Instruct(const BasicOpcode basic_opcode,
      const Operand operand_b, const Operand operand_a) {
    return static_cast<int>(basic_opcode) | (static_cast<int>(operand_a)
        << kBasicOperandShiftA) | (static_cast<int>(operand_b) << kBasicOperandShiftB);
  }

  Word Instruct(const AdvancedOpcode advanced_opcode, const Operand operand_a) {
    return (static_cast<int>(advanced_opcode) << kAdvancedOpcodeShift) | (
        static_cast<int>(operand_a) << kAdvancedOperandShiftA);
  }

  Word *Dcpu::address(const Word address_value) {
    return memory_begin() + address_value;
  }

  const Word *Dcpu::address(const Word address_value) const {
    return memory_begin() + address_value;
  }

  Word *Dcpu::memory_begin() {
    return &memory_[0];
  }

  const Word *Dcpu::memory_begin() const {
    return &memory_[0];
  }

  Word *Dcpu::memory_end() {
    return memory_begin() + kMemorySize;
  }

  const Word *Dcpu::memory_end() const {
    return memory_begin() + kMemorySize;
  }

  void Dcpu::Connect(Hardware *hardware) {
    this->hardware.push_back(hardware);
  }

  void Dcpu::Interrupt(const Word message) {
    if (!queue_interrupts && interrupt_address) {
      // TODO(robertsdionne): turn on interrupt queueing
      queue_interrupts = true;
      stack_pointer -= 1;
      *address(stack_pointer) = program_counter;
      stack_pointer -= 1;
      *address(stack_pointer) = register_a;
      program_counter = interrupt_address;
      register_a = message;
    }
  }

  void Dcpu::ExecuteInstruction(const bool skip) {
    const Word stack_pointerbackup = stack_pointer;
    const Word instruction = *address(program_counter);
    program_counter += 1;
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
      const Word operand_a_value = operand_a_address ? *operand_a_address : operand_a_literal;
      if (skip) {
        stack_pointer = stack_pointerbackup;
        return;
      }
      switch (advanced_opcode) {
        case AdvancedOpcode::kAdvancedReserved:
          break;
        case AdvancedOpcode::kJumpSubRoutine: {
          stack_pointer -= 1;
          *address(stack_pointer) = program_counter;
          program_counter = operand_a_value;
          break;
        }
        case AdvancedOpcode::kInterruptTrigger: {
          Interrupt(operand_a_value);
          break;
        }
        case AdvancedOpcode::kInterruptAddressGet: {
          MaybeAssignResult(operand_a_address, interrupt_address);
          break;
        }
        case AdvancedOpcode::kInterruptAddressSet: {
          interrupt_address = operand_a_value;
          break;
        }
        case AdvancedOpcode::kReturnFromInterrupt: {
          // TODO(robertsdionne): disable interrupt queueing
          queue_interrupts = false;
          register_a = *address(stack_pointer);
          stack_pointer += 1;
          program_counter = *address(stack_pointer);
          stack_pointer += 1;
          break;
        }
        case AdvancedOpcode::kHardwareNumberConnected: {
          MaybeAssignResult(operand_a_address, static_cast<Word>(hardware.size()));
          break;
        }
        case AdvancedOpcode::kHardwareQuery: {
          if (operand_a_value < hardware.size()) {
            auto hardware_id = hardware[operand_a_value]->GetId();
            auto hardware_version = hardware[operand_a_value]->GetVersion();
            auto manufacturer_id = hardware[operand_a_value]->GetManufacturerId();
            register_a = hardware_id;
            register_b = hardware_id >> 16;
            register_c = hardware_version;
            register_x = manufacturer_id;
            register_y = manufacturer_id >> 16;
          } else {
            register_a = 0;
            register_b = 0;
            register_c = 0;
            register_x = 0;
            register_y = 0;
          }
          break;
        }
        case AdvancedOpcode::kHardwareInterrupt: {
          if (operand_a_value < hardware.size()) {
            hardware[operand_a_value]->HandleHardwareInterrupt();
          }
          break;
        }
        default: break;
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
      const Word operand_a_value = operand_a_address ? *operand_a_address : operand_a_literal;
      const Word operand_b_value = operand_b_address ? *operand_b_address : operand_b_literal;
      if (skip) {
        stack_pointer = stack_pointerbackup;
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
          extra = result >> 16;
          MaybeAssignResult(operand_b_address, result);
          break;
        case BasicOpcode::kSubtract:
          extra = operand_b_value < operand_a_value;
          MaybeAssignResult(operand_b_address, operand_b_value - operand_a_value);
          break;
        case BasicOpcode::kMultiply:
          result = operand_b_value * operand_a_value;
          extra = result >> 16;
          MaybeAssignResult(operand_b_address, result);
          break;
        case BasicOpcode::kMultiplySigned:
          signed_result = static_cast<SignedWord>(operand_b_value) *
              static_cast<SignedWord>(operand_a_value);
          extra = signed_result >> 16;
          MaybeAssignResult(operand_b_address, static_cast<Word>(signed_result));
          break;
        case BasicOpcode::kDivide:
          if (operand_a_value) {
            extra = 0;
            MaybeAssignResult(operand_b_address, operand_b_value / operand_a_value);
          } else {
            extra = 1;
            MaybeAssignResult(operand_b_address, 0);
          }
          break;
        case BasicOpcode::kDivideSigned:
          if (operand_a_value) {
            extra = 0;
            MaybeAssignResult(operand_b_address, static_cast<Word>(
                static_cast<SignedWord>(operand_b_value) /
                    static_cast<SignedWord>(operand_a_value)));
          } else {
            extra = 1;
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
          extra = operand_b_value << (0x10 - operand_a_value);
          MaybeAssignResult(operand_b_address, result);
          break;
        case BasicOpcode::kArithmeticShiftRight:
          signed_result = static_cast<SignedWord>(operand_b_value) >> operand_a_value;
          extra = operand_b_value << (0x10 - operand_a_value);
          MaybeAssignResult(operand_b_address, static_cast<Word>(signed_result));
          break;
        case BasicOpcode::kShiftLeft:
          result = operand_b_value << operand_a_value;
          extra = result >> 16;
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
        default: break;
      }
    }
  }

  void Dcpu::ExecuteInstructions(const unsigned long int count) {
    for (unsigned long int i = 0; i < count; ++i) {
      ExecuteInstruction();
    }
  }

  void Dcpu::Reset() {
    register_a = register_b = register_c = 0;
    register_x = register_y = register_z = 0;
    register_i = register_j = 0;
    program_counter = 0;
    stack_pointer = 0;
    extra = 0;
    interrupt_address = 0;
    std::fill(memory_begin(), memory_end(), 0);
    hardware.clear();
  }

  Word *Dcpu::register_address(const Word register_index) {
    return &register_a + register_index % static_cast<int>(Operand::kLocationInRegisterA);
  }

  Word Dcpu::register_value(const Word register_index) {
    return *register_address(register_index);
  }

  Word *Dcpu::GetOperandAddressOrLiteral(
      const Operand operand, const bool assignable, Word &literal) {
    if (operand < Operand::kLocationInRegisterA) {
      return register_address(static_cast<int>(operand));
    } else if (Operand::kLocationInRegisterA <= operand
        && operand < Operand::kLocationOffsetByRegisterA) {
      return address(register_value(static_cast<int>(operand)));
    } else if (Operand::kLocationOffsetByRegisterA <= operand && operand < Operand::kPushPop) {
      Word *const result =
          address(*address(program_counter)) + register_value(static_cast<int>(operand));
      program_counter += 1;
      return result;
    } else if (operand == Operand::kPushPop) {
      if (assignable) { // Push
        stack_pointer -= 1;
        return address(stack_pointer);
      } else { // Pop
        Word *const result = address(stack_pointer);
        stack_pointer += 1;
        return result;
      }
    } else if (operand == Operand::kPeek) {
      return address(stack_pointer);
    } else if (operand == Operand::kPick) {
      const Word offset = *address(program_counter);
      program_counter += 1;
      return address(stack_pointer + offset);
    } else if (operand == Operand::kStackPointer) {
      return &stack_pointer;
    } else if (operand == Operand::kProgramCounter) {
      return &program_counter;
    } else if (operand == Operand::kExtra) {
      return &extra;
    } else if (operand == Operand::kLocation) {
      Word *const result = address(*address(program_counter));
      program_counter += 1;
      return result;
    } else if (operand == Operand::kLiteral) {
      Word *const result = address(program_counter);
      program_counter += 1;
      return result;
    } else {
      literal = static_cast<int>(operand) - static_cast<int>(Operand::k0);
      return 0;
    }
  }

  void Dcpu::MaybeAssignResult(Word *const slot, const unsigned int result) {
    if (slot) {
      *slot = static_cast<Word>(result);
    }
  }

}  // namespace dcpu
