#include <iostream>
#include <map>
#include <string>

#include "assembler.hpp"
#include "dcpu.hpp"
#include "generated/program.pb.h"

namespace dcpu {

  void Assembler::Assemble(const proto::Program &program,
      const Word *const memory_begin,
      const Word *const memory_end) const {
    std::map<std::string, Word> labels;
    Word label_address = 0;
    for (int i = 0; i < program.statement_size(); ++i) {
      const proto::Statement &statement = program.statement(i);
      if (!statement.has_type()) {
        std::cout << "Encountered statement without type!" << std::endl;
        continue;
      } else if (statement.type() == proto::Statement_Type_LABEL) {
        if (!statement.has_label()) {
          continue;
        } else if (labels.find(statement.label()) != labels.end()) {
          std::cout << "Encountered duplicate label \""
              << statement.label() << "\"!" << std::endl;
        } else {
          labels[statement.label()] = label_address;
        }
      } else {
        label_address += DetermineStatementSize(statement);
      }
    }
    std::map<std::string, Word>::const_iterator i;
    for (i = labels.begin(); i != labels.end(); ++i) {
      std::cout << i->first << ": " << i->second << std::endl;
    }
  }

  Word Assembler::DetermineStatementSize(const proto::Statement &statement) const {
    if (!statement.has_type()) {
      std::cout << "Encountered statement without type!" << std::endl;
      return 0;
    }
    if (statement.type() == proto::Statement_Type_INSTRUCTION) {
      if (!statement.has_instruction()) {
        std::cout << "Encountered statement without instruction!" << std::endl;
        return 0;
      }
      return DetermineInstructionSize(statement.instruction());
    } else if (statement.type() == proto::Statement_Type_DATA) {
      return DetermineDataSize(statement.data());
    } else {
      return 0;
    }
  }

  Word Assembler::DetermineDataSize(const proto::Data &data) const {
    if (!data.has_type()) {
      std::cout << "Encountered data without type!" << std::endl;
      return 0;
    }
    if (data.type() == proto::Data_Type_STRING) {
      return data.string().size();
    } else {
      return data.bytes().size();
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
    if (opcode.type() == proto::Opcode_Type_BASIC) {
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
      case proto::Operand_Type_LOCATION:
      case proto::Operand_Type_LITERAL:
        return 1;
      default:
        return 0;
    }
  }

}  // namespace dcpu
