// Copyright 2012 Robert Scott Dionne. All rights reserved.

#include <iostream>
#include <map>
#include <string>
#include "assembler.h"
#include "dcpu.h"
#include "program.pb.h"

void Assembler::Assemble(const Program &program,
    const Dcpu::Word *const memory_begin,
    const Dcpu::Word *const memory_end) const {
  std::map<std::string, Dcpu::Word> labels;
  Dcpu::Word label_address = 0;
  for (int i = 0; i < program.statement_size(); ++i) {
    const Statement &statement = program.statement(i);
    if (!statement.has_type()) {
      continue;
    } else if (statement.type() == Statement_Type_LABEL) {
      if (!statement.has_label()) {
        continue;
      } else {
        labels[statement.label()] = label_address;
      }
    } else {
      label_address += DetermineStatementSize(statement);
    }
  }
}

Dcpu::Word Assembler::DetermineStatementSize(const Statement &statement) const {
  if (!statement.has_type()) {
    return 0;
  }
  if (statement.type() == Statement_Type_INSTRUCTION) {
    if (!statement.has_instruction()) {
      return 0;
    }
    return DetermineInstructionSize(statement.instruction());
  } else if (statement.type() == Statement_Type_DATA) {
    return 0;
  } else {
    return 0;
  }
}

Dcpu::Word Assembler::DetermineInstructionSize(
    const Instruction &instruction) const {
  if (!instruction.has_opcode()) {
    return 0;
  }
  const Opcode &opcode = instruction.opcode();
  if (!opcode.has_type()) {
    return 0;
  }
  if (opcode.type() == Opcode_Type_BASIC) {
    if (!instruction.has_operand_a()) {
      return 0;
    }
    if (!instruction.has_operand_b()) {
      return 0;
    }
    return 1 + DetermineOperandSize(instruction.operand_a())
        + DetermineOperandSize(instruction.operand_b());
  } else {
    if (!instruction.has_operand_a()) {
      return 0;
    }
    return 1 + DetermineOperandSize(instruction.operand_a());
  }
}

Dcpu::Word Assembler::DetermineOperandSize(const Operand &operand) const {
  if (!operand.has_type()) {
    return 0;
  }
  switch (operand.type()) {
    case Operand_Type_LOCATION_OFFSET_BY_REGISTER:
    case Operand_Type_PICK:
    case Operand_Type_LOCATION:
    case Operand_Type_LITERAL:
      return 1;
    default:
      return 0;
  }
}
