#include <iomanip>
#include <iostream>
#include <map>
#include <string>

#include "assembler.hpp"
#include "dcpu.hpp"
#include "generated/program.pb.h"

namespace dcpu {

  void Assembler::Assemble(const proto::Program &program, Word *const memory_begin) const {
    Labels labels;
    Word label_address = 0;
    for (int i = 0; i < program.statement_size(); ++i) {
      auto statement = program.statement(i);
      if (!statement.has_type()) {
        std::cout << "Encountered statement without type!" << std::endl;
      } else if (proto::Statement_Type_LABEL == statement.type()) {
        if (!statement.has_label()) {
          std::cout << "Encountered label without a label!" << std::endl;
        } else if (labels.cend() != labels.find(statement.label())) {
          std::cout << "Encountered duplicate label \"" << statement.label() << "\"!" << std::endl;
        } else {
          labels[statement.label()] = label_address;
        }
      } else {
        label_address += DetermineStatementSize(statement);
      }
    }
    for (auto i = labels.begin(); i != labels.end(); ++i) {
      std::cout << i->first << ": 0x" << std::hex << i->second << std::dec << std::endl;
    }
    std::vector<Word> binary;
    for (int i = 0; i < program.statement_size(); ++i) {
      auto statement = program.statement(i);
      if (proto::Statement_Type_INSTRUCTION == statement.type()) {
        auto instruction = EncodeInstruction(labels, statement.instruction());
        binary.insert(binary.cend(), instruction.begin(), instruction.end());
      } else if (proto::Statement_Type_DATA == statement.type()) {
        binary.push_back(statement.data());
      }
    }
    std::copy(binary.begin(), binary.end(), memory_begin);
  }

  Word Assembler::DetermineStatementSize(const proto::Statement &statement) const {
    if (!statement.has_type()) {
      std::cout << "Encountered statement without type!" << std::endl;
      return 0;
    }
    if (proto::Statement_Type_INSTRUCTION == statement.type()) {
      if (!statement.has_instruction()) {
        std::cout << "Encountered statement without instruction!" << std::endl;
        return 0;
      }
      return DetermineInstructionSize(statement.instruction());
    } else if (proto::Statement_Type_DATA == statement.type()) {
      return 1;
    } else {
      return 0;
    }
  }

  Word Assembler::DetermineInstructionSize(const proto::Instruction &instruction) const {
    if (!instruction.has_opcode()) {
      std::cout << "Encountered instruction without opcode!" << std::endl;
      return 0;
    }
    const proto::Opcode &opcode = instruction.opcode();
    if (!opcode.has_type()) {
      std::cout << "Encountered opcode without type!" << std::endl;
      return 0;
    }
    if (proto::Opcode_Type_BASIC == opcode.type()) {
      if (!instruction.has_operand_a()) {
        std::cout << "Encountered basic opcode without operand A!" << std::endl;
        return 0;
      }
      if (!instruction.has_operand_b()) {
        std::cout << "Encountered basic opcode without operand B!" << std::endl;
        return 0;
      }
      return 1 + DetermineOperandSize(instruction.operand_a())
          + DetermineOperandSize(instruction.operand_b());
    } else {
      if (!instruction.has_operand_a()) {
        std::cout << "Encountered advanced opcode without operand A!"
            << std::endl;
        return 0;
      }
      return 1 + DetermineOperandSize(instruction.operand_a());
    }
  }

  Word Assembler::DetermineOperandSize(const proto::Operand &operand) const {
    if (!operand.has_type()) {
      std::cout << "Encountered operand without type!" << std::endl;
      return 0;
    }
    switch (operand.type()) {
      case proto::Operand_Type_LOCATION_OFFSET_BY_REGISTER:
      case proto::Operand_Type_PICK:
      case proto::Operand_Type_LOCATION: {
        return 1;
      }
      case proto::Operand_Type_LITERAL: {
        if (operand.has_value() && IsSmallValue(operand)) {
          return 0;
        } else {
          return 1;
        }
      }
      default:{
        return 0;
      }
    }
  }

  Assembler::Binary Assembler::EncodeInstruction(
      const Labels &labels, const proto::Instruction &instruction) const {
    auto result = Binary();
    const auto opcode = instruction.opcode();
    const auto operand_a = instruction.operand_a();
    const auto a = EncodeOperand(operand_a);
    if (proto::Opcode_Type_BASIC == opcode.type()) {
      const auto operand_b = instruction.operand_b();
      const auto b = EncodeOperand(operand_b);
      result.push_back(dcpu::Instruct(static_cast<BasicOpcode>(opcode.basic()), b, a));
      if (RequiresAdditionalWord(operand_a)) {
        result.push_back(EncodeLiteral(labels, operand_a));
      }
      if (RequiresAdditionalWord(operand_b)) {
        result.push_back(EncodeLiteral(labels, operand_b));
      }
    } else if (proto::Opcode_Type_ADVANCED == opcode.type()) {
      result.push_back(dcpu::Instruct(static_cast<AdvancedOpcode>(opcode.advanced()), a));
      if (RequiresAdditionalWord(operand_a)) {
        result.push_back(EncodeLiteral(labels, operand_a));
      }
    }
    return result;
  }

  Operand Assembler::EncodeOperand(const proto::Operand &operand) const {
    switch (operand.type()) {
      case proto::Operand_Type_REGISTER:
      case proto::Operand_Type_LOCATION_IN_REGISTER:
      case proto::Operand_Type_LOCATION_OFFSET_BY_REGISTER: {
        return static_cast<Operand>(operand.type() + operand.register_());
      }
      case proto::Operand_Type_LITERAL: {
        if (operand.has_value() && IsSmallValue(operand)) {
          return static_cast<Operand>(operand.value() + static_cast<Word>(Operand::k0));
        } else {
          return static_cast<Operand>(operand.type());
        }
      }
      default: return static_cast<Operand>(operand.type());
    }
  }

  bool Assembler::RequiresAdditionalWord(const proto::Operand &operand) const {
    return operand.has_label() || (operand.has_value() && !IsSmallValue(operand));
  }

  bool Assembler::IsSmallValue(const proto::Operand &operand) const {
    return proto::Operand_Type_LITERAL == operand.type() && operand.has_value()
        && (0xffff == operand.value() || operand.value() <= 30);
  }

  Word Assembler::EncodeLiteral(const Labels &labels, const proto::Operand &operand) const {
    if (operand.has_label()) {
      if (labels.cend() == labels.find(operand.label())) {
        std::cout << "Encountered unknown label \"" << operand.label() << "\"!" << std::endl;
        return 0;
      } else {
        return labels.at(operand.label());
      }
    } else if (operand.has_value()) {
      return operand.value();
    } else {
      return 0;
    }
  }

}  // namespace dcpu
