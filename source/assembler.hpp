#ifndef DCPU_ASSEMBLER_HPP_
#define DCPU_ASSEMBLER_HPP_

#include "dcpu.hpp"

namespace dcpu {

  class Data;
  class Instruction;
  class Operand;
  class Program;
  class Statement;

  class Assembler {
  public:
    Assembler() {}
    virtual ~Assembler() {}

    void Assemble(const Program &program,
                  const Dcpu::Word *const memory_begin,
                  const Dcpu::Word *const memory_end) const;
  private:
    Dcpu::Word DetermineStatementSize(const Statement &statement) const;
    Dcpu::Word DetermineDataSize(const Data &data) const;
    Dcpu::Word DetermineInstructionSize(const Instruction &instruction) const;
    Dcpu::Word DetermineOperandSize(const Operand &operand) const;
  };

}  // namespace dcpu

#endif  // DCPU_ASSEMBLER_HPP_
