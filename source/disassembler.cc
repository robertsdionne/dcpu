// Copyright 2012 Robert Scott Dionne. All rights reserved.

#include <iostream>
#include <sstream>
#include "dcpu.h"
#include "disassembler.h"

void Disassembler::Disassemble(const Dcpu::Word *const program_begin,
    const Dcpu::Word *const program_end, std::ostream &out) const {
  for (const Dcpu::Word *i = program_begin; i < program_end; ++i) {
    const Dcpu::Word instruction = *i;
    const Dcpu::BasicOpcode basic_opcode = static_cast<Dcpu::BasicOpcode>(instruction & Dcpu::kBasicOpcodeMask);

    if (basic_opcode == Dcpu::BasicOpcode::kBasicReserved) {
      const Dcpu::AdvancedOpcode advanced_opcode = static_cast<Dcpu::AdvancedOpcode>((instruction &
          Dcpu::kAdvancedOpcodeMask) >> Dcpu::kAdvancedOpcodeShift);
      const Dcpu::Operand operand_a = static_cast<Dcpu::Operand>(
          (instruction & Dcpu::kAdvancedOperandMaskA) >>
              Dcpu::kAdvancedOperandShiftA);
      switch (advanced_opcode) {
        case Dcpu::AdvancedOpcode::kAdvancedReserved:
          out << "set a, a" << std::endl;
          break;
        case Dcpu::AdvancedOpcode::kJumpSubRoutine:
          out << "jsr ";
          OutputOperand(i, operand_a, /* assignable */ false, out);
          out << std::endl;
          break;
        default:
          out << "set a, a" << std::endl;
          break;
      }
    } else {
      const Dcpu::Operand operand_a = static_cast<Dcpu::Operand>(
          (instruction & Dcpu::kBasicOperandMaskA) >>
              Dcpu::kBasicOperandShiftA);
      const Dcpu::Operand operand_b = static_cast<Dcpu::Operand>(
          (instruction & Dcpu::kBasicOperandMaskB) >>
              Dcpu::kBasicOperandShiftB);
      switch (basic_opcode) {
        case Dcpu::BasicOpcode::kBasicReserved:
          out << "set a, a" << std::endl;
          break;
        case Dcpu::BasicOpcode::kSet:
          out << "set ";
          OutputOperands(i, operand_b, operand_a, out);
          out << std::endl;
          break;
        case Dcpu::BasicOpcode::kAdd:
          out << "add ";
          OutputOperands(i, operand_b, operand_a, out);
          out << std::endl;
          break;
        case Dcpu::BasicOpcode::kSubtract:
          out << "sub ";
          OutputOperands(i, operand_b, operand_a, out);
          out << std::endl;
          break;
        case Dcpu::BasicOpcode::kMultiply:
          out << "mul ";
          OutputOperands(i, operand_b, operand_a, out);
          out << std::endl;
          break;
        case Dcpu::BasicOpcode::kDivide:
          out << "div ";
          OutputOperands(i, operand_b, operand_a, out);
          out << std::endl;
          break;
        case Dcpu::BasicOpcode::kModulo:
          out << "mod ";
          OutputOperands(i, operand_b, operand_a, out);
          out << std::endl;
          break;
        case Dcpu::BasicOpcode::kShiftLeft:
          out << "shl ";
          OutputOperands(i, operand_b, operand_a, out);
          out << std::endl;
          break;
        case Dcpu::BasicOpcode::kShiftRight:
          out << "shr ";
          OutputOperands(i, operand_b, operand_a, out);
          out << std::endl;
          break;
        case Dcpu::BasicOpcode::kBinaryAnd:
          out << "and ";
          OutputOperands(i, operand_b, operand_a, out);
          out << std::endl;
          break;
        case Dcpu::BasicOpcode::kBinaryOr:
          out << "bor ";
          OutputOperands(i, operand_b, operand_a, out);
          out << std::endl;
          break;
        case Dcpu::BasicOpcode::kBinaryExclusiveOr:
          out << "xor ";
          OutputOperands(i, operand_b, operand_a, out);
          out << std::endl;
          break;
        case Dcpu::BasicOpcode::kIfEqual:
          out << "ife ";
          OutputOperands(i, operand_b, operand_a, out);
          out << std::endl;
          break;
        case Dcpu::BasicOpcode::kIfNotEqual:
          out << "ifn ";
          OutputOperands(i, operand_b, operand_a, out);
          out << std::endl;
          break;
        case Dcpu::BasicOpcode::kIfGreaterThan:
          out << "ifg ";
          OutputOperands(i, operand_b, operand_a, out);
          out << std::endl;
          break;
        case Dcpu::BasicOpcode::kIfBitSet:
          out << "ifb ";
          OutputOperands(i, operand_b, operand_a, out);
          out << std::endl;
          break;
        default:
          out << "set a, a" << std::endl;
          break;
      }
    }
  }
}

char Disassembler::DetermineRegisterName(const Dcpu::Operand operand) const {
  switch (static_cast<Dcpu::Operand>(static_cast<int>(operand) % 8)) {
    case Dcpu::Operand::kRegisterA:
      return 'a';
    case Dcpu::Operand::kRegisterB:
      return 'b';
    case Dcpu::Operand::kRegisterC:
      return 'c';
    case Dcpu::Operand::kRegisterX:
      return 'x';
    case Dcpu::Operand::kRegisterY:
      return 'y';
    case Dcpu::Operand::kRegisterZ:
      return 'z';
    case Dcpu::Operand::kRegisterI:
      return 'i';
    default:
      return 'j';
  }
}

void Disassembler::OutputOperand(
    const Dcpu::Word *&i, const Dcpu::Operand operand,
    const bool assignable, std::ostream &out) const {
  if (operand < Dcpu::Operand::kLocationInRegisterA) {
    const char register_name = DetermineRegisterName(operand);
    out << register_name;
  } else if (Dcpu::Operand::kLocationInRegisterA <= operand &&
      operand < Dcpu::Operand::kLocationOffsetByRegisterA) {
    const char register_name = DetermineRegisterName(operand);
    out << '[' << register_name << ']';
  } else if (Dcpu::Operand::kLocationOffsetByRegisterA <= operand &&
      operand < Dcpu::Operand::kPop) {
    const char register_name = DetermineRegisterName(operand);
    i += 1;
    out << '[';
    std::ios_base::fmtflags flags = out.flags();
    out << std::showbase << std::hex << *i;
    out.flags(flags);
    out << '+' << register_name << ']';
  } else {
    std::ios_base::fmtflags flags;
    switch (operand) {
      case Dcpu::Operand::kPushPop:
        if (assignable) {
          out << "push";
        } else {
          out << "pop";
        }
        break;
      case Dcpu::Operand::kPeek:
        out << "peek";
        break;
      case Dcpu::Operand::kPick:
        i += 1;
        out << "[sp+";
        flags = out.flags();
        out << std::showbase << std::hex << *i;
        out.flags(flags);
        out << ']';
        break;
      case Dcpu::Operand::kStackPointer:
        out << "sp";
        break;
      case Dcpu::Operand::kProgramCounter:
        out << "pc";
        break;
      case Dcpu::Operand::kExtra:
        out << "ex";
        break;
      case Dcpu::Operand::kLocation:
        i += 1;
        out << '[';
        flags = out.flags();
        out << std::showbase << std::hex << *i;
        out.flags(flags);
        out << ']';
        break;
      case Dcpu::Operand::kLiteral:
        i += 1;
        flags = out.flags();
        out << std::showbase << std::hex << *i;
        out.flags(flags);
        break;
      default:
        flags = out.flags();
        out << std::showbase << std::hex << static_cast<int>(operand) - static_cast<int>(Dcpu::Operand::k0);
        out.flags(flags);
        break;
    }
  }
}

void Disassembler::OutputOperands(
    const Dcpu::Word *&i, const Dcpu::Operand operand_b,
    const Dcpu::Operand operand_a, std::ostream &out) const {
  std::ostringstream string_out;
  OutputOperand(i, operand_a, /* assignable */ false, string_out);
  OutputOperand(i, operand_b, /* assignable */ true, out);
  out << ", " << string_out.str();
}
