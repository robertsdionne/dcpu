#ifndef DCPU_PARSER_HPP_
#define DCPU_PARSER_HPP_

#include "generated/program.pb.h"

namespace dcpu {
  class Lexer;

  class Parser {
  public:
    Parser(Lexer &lexer);

    virtual ~Parser() = default;

    bool ParseProgram(proto::Program *program);

  private:
    bool ParseData(proto::Statement *data);
    bool ParseInstruction(proto::Statement *instruction);
    bool ParseInstruction(proto::Instruction *instruction);
    bool ParseLabel(proto::Statement *label);
    bool ParseStatement(proto::Statement *statement);

  private:
    Lexer &lexer;
  };

}  // namespace dcpu

#endif  // DCPU_PARSER_HPP_
