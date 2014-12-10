#include "generated/program.pb.h"
#include "lexer.hpp"
#include "parser.hpp"

namespace dcpu {

  using namespace proto;

  Parser::Parser(Lexer &lexer) : lexer(lexer) {}


  //
  // Goal: Parse the following program:
  //
  // set a, 0x0000
  // my_label:
  //   .data 0x0001
  //

  bool Parser::ParseProgram(Program *program) {
    *program = Program();
    while (true) {
      auto statement = proto::Statement();
      if (ParseStatement(&statement)) {
        program->add_statement()->CopyFrom(statement);
      } else {
        break;
      }
    }
    return program->statement_size() > 0;
  }

  bool Parser::ParseStatement(Statement *statement) {
    return ParseData(statement) || ParseInstruction(statement) || ParseLabel(statement);
  }

  bool Parser::ParseData(Statement *data) {
    if (Token::Type::kData != lexer.SeeToken().type) {
      return false;
    }
    lexer.EatToken();

    data->set_type(Statement_Type_DATA);

    if (Token::Type::kWhitespace == lexer.SeeToken().type) {
      lexer.EatToken();
    }

    if (Token::Type::kBinary == lexer.SeeToken().type) {
      auto binary = lexer.EatToken();
      data->set_data(1);
      return true;
    } else if (Token::Type::kHexadecimal == lexer.SeeToken().type) {
      auto hexadecimal = lexer.EatToken();
      data->set_data(2);
      return true;
    } else if (Token::Type::kDecimal == lexer.SeeToken().type) {
      auto decimal = lexer.EatToken();
      data->set_data(3);
      return true;
    } else if (Token::Type::kString == lexer.SeeToken().type) {
      auto string = lexer.EatToken();
      data->set_data(4);
      return true;
    } else {
      return false;
    }
  }

  bool Parser::ParseInstruction(Statement *instruction) {
    if (Token::Type::kAdvancedOpcode != lexer.SeeToken().type ||
        Token::Type::kBasicOpcode != lexer.SeeToken().type) {
      return false;
    }
    instruction->set_type(Statement_Type_INSTRUCTION);
    return ParseInstruction(instruction->mutable_instruction());
  }

  bool Parser::ParseInstruction(Instruction *instruction) {
    return false;
  }

  bool Parser::ParseLabel(Statement *label) {
    return false;
  }

}  // namespace dcpu
