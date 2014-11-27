// Copyright 2012 Robert Scott Dionne. All rights reserved.

#ifndef DCPU_DISASSEMBLER_H_
#define DCPU_DISASSEMBLER_H_

#include <iostream>
#include "dcpu.hpp"

class Disassembler {
public:
  Disassembler() {}
  virtual ~Disassembler() {}

  void Disassemble(const Dcpu::Word *const program_begin,
                   const Dcpu::Word *const program_end, std::ostream &out) const;

private:
  char DetermineRegisterName(const Dcpu::Operand operand) const;
  void OutputOperand(const Dcpu::Word *&i, const Dcpu::Operand operand,
                     const bool assignable, std::ostream &out) const;
  void OutputOperands(
                      const Dcpu::Word *&i, const Dcpu::Operand operand_a,
                      const Dcpu::Operand operand_b, std::ostream &out) const;
};

#endif  // DCPU_DISASSEMBLER_H_
