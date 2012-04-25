// Copyright 2012 Robert Scott Dionne. All rights reserved.

#include <iostream>
#include "dcpu.h"
#include "disassembler.h"

void Disassembler::Disassemble(const Dcpu::Word *const program_begin,
    const Dcpu::Word *const program_end, std::ostream &out) const {
  for (const Dcpu::Word *i = program_begin; i < program_end; ++i) {
    const Dcpu::Word instruction = *i;
    const Dcpu::Word basic_opcode = instruction & Dcpu::kBasicOpcodeMask;

    if (basic_opcode == Dcpu::kBasicReserved) {
      const Dcpu::Word advanced_opcode = (instruction &
          Dcpu::kAdvancedOpcodeMask) >> Dcpu::kAdvancedOpcodeShift;
      const Dcpu::Operand operand_a = static_cast<Dcpu::Operand>(
          (instruction & Dcpu::kAdvancedOperandMaskA) >>
              Dcpu::kAdvancedOperandShiftA);
      switch (advanced_opcode) {
        case Dcpu::kAdvancedReserved:
          out << "set 0, 0" << std::endl;
          break;
        case Dcpu::kJumpSubRoutine:
          out << "jsr ";
          OutputOperand(i, operand_a, out);
          out << std::endl;
          break;
        default:
          out << "set 0, 0" << std::endl;
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
        case Dcpu::kBasicReserved:
          out << "set 0, 0" << std::endl;
          break;
        case Dcpu::kSet:
          out << "set ";
          OutputOperands(i, operand_a, operand_b, out);
          out << std::endl;
          break;
        case Dcpu::kAdd:
          out << "add ";
          OutputOperands(i, operand_a, operand_b, out);
          out << std::endl;
          break;
        case Dcpu::kSubtract:
          out << "sub ";
          OutputOperands(i, operand_a, operand_b, out);
          out << std::endl;
          break;
        case Dcpu::kMultiply:
          out << "mul ";
          OutputOperands(i, operand_a, operand_b, out);
          out << std::endl;
          break;
        case Dcpu::kDivide:
          out << "div ";
          OutputOperands(i, operand_a, operand_b, out);
          out << std::endl;
          break;
        case Dcpu::kModulo:
          out << "mod ";
          OutputOperands(i, operand_a, operand_b, out);
          out << std::endl;
          break;
        case Dcpu::kShiftLeft:
          out << "shl ";
          OutputOperands(i, operand_a, operand_b, out);
          out << std::endl;
          break;
        case Dcpu::kShiftRight:
          out << "shr ";
          OutputOperands(i, operand_a, operand_b, out);
          out << std::endl;
          break;
        case Dcpu::kBinaryAnd:
          out << "and ";
          OutputOperands(i, operand_a, operand_b, out);
          out << std::endl;
          break;
        case Dcpu::kBinaryOr:
          out << "bor ";
          OutputOperands(i, operand_a, operand_b, out);
          out << std::endl;
          break;
        case Dcpu::kBinaryExclusiveOr:
          out << "xor ";
          OutputOperands(i, operand_a, operand_b, out);
          out << std::endl;
          break;
        case Dcpu::kIfEqual:
          out << "ife ";
          OutputOperands(i, operand_a, operand_b, out);
          out << std::endl;
          break;
        case Dcpu::kIfNotEqual:
          out << "ifn ";
          OutputOperands(i, operand_a, operand_b, out);
          out << std::endl;
          break;
        case Dcpu::kIfGreaterThan:
          out << "ifg ";
          OutputOperands(i, operand_a, operand_b, out);
          out << std::endl;
          break;
        case Dcpu::kIfBoth:
          out << "ifb ";
          OutputOperands(i, operand_a, operand_b, out);
          out << std::endl;
          break;
        default:
          out << "set 0, 0" << std::endl;
          break;
      }
    }
  }
}

char Disassembler::DetermineRegisterName(const Dcpu::Operand operand) const {
  switch (operand % 8) {
    case Dcpu::kRegisterA:
      return 'a';
    case Dcpu::kRegisterB:
      return 'b';
    case Dcpu::kRegisterC:
      return 'c';
    case Dcpu::kRegisterX:
      return 'x';
    case Dcpu::kRegisterY:
      return 'y';
    case Dcpu::kRegisterZ:
      return 'z';
    case Dcpu::kRegisterI:
      return 'i';
    default:
      return 'j';
  }
}

void Disassembler::OutputOperand(const Dcpu::Word *&i,
    const Dcpu::Operand operand, std::ostream &out) const {
  if (operand < Dcpu::kLocationInRegisterA) {
    const char register_name = DetermineRegisterName(operand);
    out << register_name;
  } else if (Dcpu::kLocationInRegisterA <= operand &&
      operand < Dcpu::kLocationOffsetByRegisterA) {
    const char register_name = DetermineRegisterName(operand);
    out << '[' << register_name << ']';
  } else if (Dcpu::kLocationOffsetByRegisterA <= operand &&
      operand < Dcpu::kPop) {
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
      case Dcpu::kPop:
        out << "pop";
        break;
      case Dcpu::kPeek:
        out << "peek";
        break;
      case Dcpu::kPush:
        out << "push";
        break;
      case Dcpu::kStackPointer:
        out << "sp";
        break;
      case Dcpu::kProgramCounter:
        out << "pc";
        break;
      case Dcpu::kOverflow:
        out << "o";
        break;
      case Dcpu::kLocation:
        i += 1;
        out << '[';
        flags = out.flags();
        out << std::showbase << std::hex << *i;
        out.flags(flags);
        out << ']';
        break;
      case Dcpu::kLiteral:
        i += 1;
        flags = out.flags();
        out << std::showbase << std::hex << *i;
        out.flags(flags);
        break;
      default:
        flags = out.flags();
        out << std::showbase << std::hex << operand - Dcpu::k0;
        out.flags(flags);
        break;
    }
  }
}

void Disassembler::OutputOperands(
    const Dcpu::Word *&i, const Dcpu::Operand operand_a,
    const Dcpu::Operand operand_b, std::ostream &out) const {
  OutputOperand(i, operand_a, out);
  out << ", ";
  OutputOperand(i, operand_b, out);
}
