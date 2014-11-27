#include <iostream>
#include <sstream>

#include "dcpu.hpp"
#include "disassembler.hpp"

namespace dcpu {

  void Disassembler::Disassemble(const Word *const program_begin,
      const Word *const program_end, std::ostream &out) const {
    for (const Word *i = program_begin; i < program_end; ++i) {
      const Word instruction = *i;
      const BasicOpcode basic_opcode = static_cast<BasicOpcode>(
          instruction & kBasicOpcodeMask);

      if (basic_opcode == BasicOpcode::kBasicReserved) {
        const AdvancedOpcode advanced_opcode = static_cast<AdvancedOpcode>((
            instruction & kAdvancedOpcodeMask) >> kAdvancedOpcodeShift);
        const Operand operand_a = static_cast<Operand>(
            (instruction & kAdvancedOperandMaskA) >>
                kAdvancedOperandShiftA);
        switch (advanced_opcode) {
          case AdvancedOpcode::kAdvancedReserved:
            out << "set a, a" << std::endl;
            break;
          case AdvancedOpcode::kJumpSubRoutine:
            out << "jsr ";
            OutputOperand(i, operand_a, /* assignable */ false, out);
            out << std::endl;
            break;
          default:
            out << "set a, a" << std::endl;
            break;
        }
      } else {
        const Operand operand_a = static_cast<Operand>(
            (instruction & kBasicOperandMaskA) >>
                kBasicOperandShiftA);
        const Operand operand_b = static_cast<Operand>(
            (instruction & kBasicOperandMaskB) >>
                kBasicOperandShiftB);
        switch (basic_opcode) {
          case BasicOpcode::kBasicReserved:
            out << "set a, a" << std::endl;
            break;
          case BasicOpcode::kSet:
            out << "set ";
            OutputOperands(i, operand_b, operand_a, out);
            out << std::endl;
            break;
          case BasicOpcode::kAdd:
            out << "add ";
            OutputOperands(i, operand_b, operand_a, out);
            out << std::endl;
            break;
          case BasicOpcode::kSubtract:
            out << "sub ";
            OutputOperands(i, operand_b, operand_a, out);
            out << std::endl;
            break;
          case BasicOpcode::kMultiply:
            out << "mul ";
            OutputOperands(i, operand_b, operand_a, out);
            out << std::endl;
            break;
          case BasicOpcode::kDivide:
            out << "div ";
            OutputOperands(i, operand_b, operand_a, out);
            out << std::endl;
            break;
          case BasicOpcode::kModulo:
            out << "mod ";
            OutputOperands(i, operand_b, operand_a, out);
            out << std::endl;
            break;
          case BasicOpcode::kShiftLeft:
            out << "shl ";
            OutputOperands(i, operand_b, operand_a, out);
            out << std::endl;
            break;
          case BasicOpcode::kShiftRight:
            out << "shr ";
            OutputOperands(i, operand_b, operand_a, out);
            out << std::endl;
            break;
          case BasicOpcode::kBinaryAnd:
            out << "and ";
            OutputOperands(i, operand_b, operand_a, out);
            out << std::endl;
            break;
          case BasicOpcode::kBinaryOr:
            out << "bor ";
            OutputOperands(i, operand_b, operand_a, out);
            out << std::endl;
            break;
          case BasicOpcode::kBinaryExclusiveOr:
            out << "xor ";
            OutputOperands(i, operand_b, operand_a, out);
            out << std::endl;
            break;
          case BasicOpcode::kIfEqual:
            out << "ife ";
            OutputOperands(i, operand_b, operand_a, out);
            out << std::endl;
            break;
          case BasicOpcode::kIfNotEqual:
            out << "ifn ";
            OutputOperands(i, operand_b, operand_a, out);
            out << std::endl;
            break;
          case BasicOpcode::kIfGreaterThan:
            out << "ifg ";
            OutputOperands(i, operand_b, operand_a, out);
            out << std::endl;
            break;
          case BasicOpcode::kIfBitSet:
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

  char Disassembler::DetermineRegisterName(const Operand operand) const {
    switch (static_cast<Operand>(static_cast<int>(operand) % 8)) {
      case Operand::kRegisterA:
        return 'a';
      case Operand::kRegisterB:
        return 'b';
      case Operand::kRegisterC:
        return 'c';
      case Operand::kRegisterX:
        return 'x';
      case Operand::kRegisterY:
        return 'y';
      case Operand::kRegisterZ:
        return 'z';
      case Operand::kRegisterI:
        return 'i';
      default:
        return 'j';
    }
  }

  void Disassembler::OutputOperand(
      const Word *&i, const Operand operand,
      const bool assignable, std::ostream &out) const {
    if (operand < Operand::kLocationInRegisterA) {
      const char register_name = DetermineRegisterName(operand);
      out << register_name;
    } else if (Operand::kLocationInRegisterA <= operand &&
        operand < Operand::kLocationOffsetByRegisterA) {
      const char register_name = DetermineRegisterName(operand);
      out << '[' << register_name << ']';
    } else if (Operand::kLocationOffsetByRegisterA <= operand &&
        operand < Operand::kPop) {
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
        case Operand::kPushPop:
          if (assignable) {
            out << "push";
          } else {
            out << "pop";
          }
          break;
        case Operand::kPeek:
          out << "peek";
          break;
        case Operand::kPick:
          i += 1;
          out << "[sp+";
          flags = out.flags();
          out << std::showbase << std::hex << *i;
          out.flags(flags);
          out << ']';
          break;
        case Operand::kStackPointer:
          out << "sp";
          break;
        case Operand::kProgramCounter:
          out << "pc";
          break;
        case Operand::kExtra:
          out << "ex";
          break;
        case Operand::kLocation:
          i += 1;
          out << '[';
          flags = out.flags();
          out << std::showbase << std::hex << *i;
          out.flags(flags);
          out << ']';
          break;
        case Operand::kLiteral:
          i += 1;
          flags = out.flags();
          out << std::showbase << std::hex << *i;
          out.flags(flags);
          break;
        default:
          flags = out.flags();
          out << std::showbase << std::hex
              << static_cast<int>(operand) - static_cast<int>(Operand::k0);
          out.flags(flags);
          break;
      }
    }
  }

  void Disassembler::OutputOperands(
      const Word *&i, const Operand operand_b,
      const Operand operand_a, std::ostream &out) const {
    std::ostringstream string_out;
    OutputOperand(i, operand_a, /* assignable */ false, string_out);
    OutputOperand(i, operand_b, /* assignable */ true, out);
    out << ", " << string_out.str();
  }

}
