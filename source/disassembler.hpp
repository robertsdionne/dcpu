#ifndef DCPU_DISASSEMBLER_HPP_
#define DCPU_DISASSEMBLER_HPP_

#include <iostream>

#include "dcpu.hpp"

namespace dcpu {

  class Disassembler {
  public:
    Disassembler() {}
    virtual ~Disassembler() {}

    void Disassemble(const Word *const program_begin,
                     const Word *const program_end, std::ostream &out) const;

  private:
    char DetermineRegisterName(const Operand operand) const;
    void OutputOperand(const Word *&i, const Operand operand,
                       const bool assignable, std::ostream &out) const;
    void OutputOperands(
                        const Word *&i, const Operand operand_a,
                        const Operand operand_b, std::ostream &out) const;
  };

}  // namespace dcpu

#endif  // DCPU_DISASSEMBLER_HPP_
