#ifndef DCPU_ASSEMBLER_HPP_
#define DCPU_ASSEMBLER_HPP_

#include "dcpu.hpp"

namespace dcpu {

  namespace proto {

    class Data;
    class Instruction;
    class Operand;
    class Program;
    class Statement;

  }  // namespace proto

  class Assembler {
  public:
    Assembler() {}
    virtual ~Assembler() {}

    void Assemble(const proto::Program &program,
                  const Word *const memory_begin,
                  const Word *const memory_end) const;
  private:
    Word DetermineStatementSize(const proto::Statement &statement) const;
    Word DetermineDataSize(const proto::Data &data) const;
    Word DetermineInstructionSize(const proto::Instruction &instruction) const;
    Word DetermineOperandSize(const proto::Operand &operand) const;
  };

}  // namespace dcpu

#endif  // DCPU_ASSEMBLER_HPP_
