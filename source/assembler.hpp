#ifndef DCPU_ASSEMBLER_HPP_
#define DCPU_ASSEMBLER_HPP_

#include <map>
#include <string>

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
    Assembler() = default;

    virtual ~Assembler() = default;

    void Assemble(const proto::Program &program, Word *const memory_begin) const;

  private:
    using Binary = std::vector<Word>;
    using Labels = std::map<std::string, Word>;

    Word DetermineStatementSize(const proto::Statement &statement) const;

    Word DetermineInstructionSize(const proto::Instruction &instruction) const;

    Word DetermineOperandSize(const proto::Operand &operand) const;

    Binary EncodeInstruction(const Labels &labels, const proto::Instruction &instruction) const;

    Operand EncodeOperand(const proto::Operand &operand) const;

    bool RequiresAdditionalWord(const proto::Operand &operand) const;

    bool IsSmallValue(const proto::Operand &operand) const;

    Word EncodeLiteral(const Labels &labels, const proto::Operand &operand) const;
  };

}  // namespace dcpu

#endif  // DCPU_ASSEMBLER_HPP_
