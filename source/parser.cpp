#include <string>

#include "generated/program.pb.h"
#include "lexer.hpp"
#include "parser.hpp"

namespace dcpu {

  //
  // Goal: Parse the following program:
  //
  // set a, 0x0000
  // my_label:
  //   .data 0x0001
  //

  using namespace proto;
  using Type = Token::Type;

  Parser::Parser(Lexer &lexer) : lexer(lexer) {}

  bool Parser::ParseProgram(Program *program) {
    std::cout << "ParseProgram" << std::endl;
    *program = Program();
    while (Type::kExhausted != lexer.SeeToken().type
        && Type::kInvalid != lexer.SeeToken().type) {
      auto statement = proto::Statement();
      if (ParseStatement(&statement)) {
        program->add_statement()->CopyFrom(statement);
      } else {
        break;
      }
    }
    return program->statement_size() > 0;
  }

  int Parser::ConvertHexadecimalValue(const std::string &value) {
    std::cout << "ConvertHexadecimalValue" << std::endl;
    return stoi(value, nullptr, 16);
  }

  void Parser::MaybeEatWhitespace() {
    std::cout << "MaybeEatWhitespace" << std::endl;
    if (Type::kWhitespace == lexer.SeeToken().type) {
      lexer.EatToken();
    }
  }

  bool Parser::ParseStatement(Statement *statement) {
    std::cout << "ParseStatement" << std::endl;
    return ParseData(statement) || ParseInstruction(statement) || ParseLabel(statement);
  }

  bool Parser::ParseData(Statement *data) {
    std::cout << "ParseData" << std::endl;
    if (Type::kData != lexer.SeeToken().type) {
      return false;
    }
    lexer.EatToken();

    data->set_type(Statement_Type_DATA);

    MaybeEatWhitespace();

    if (Type::kBinary == lexer.SeeToken().type) {
      auto binary = lexer.EatToken();
      data->set_data(1);
      return true;
    } else if (Type::kHexadecimal == lexer.SeeToken().type) {
      auto hexadecimal = lexer.EatToken();
      data->set_data(2);
      return true;
    } else if (Type::kDecimal == lexer.SeeToken().type) {
      auto decimal = lexer.EatToken();
      data->set_data(3);
      return true;
    } else if (Type::kString == lexer.SeeToken().type) {
      auto string = lexer.EatToken();
      data->set_data(4);
      return true;
    } else {
      std::cout << "Expected kBinary, kHexadecimal, kDecimal or kString but got "
          << lexer.SeeToken().type << std::endl;
      return false;
    }
  }

  bool Parser::ParseInstruction(Statement *instruction) {
    std::cout << "ParseInstruction" << std::endl;
    if (Type::kAdvancedOpcode != lexer.SeeToken().type &&
        Type::kBasicOpcode != lexer.SeeToken().type) {
      std::cout << "Expected kAdvancedOpcode or kBasicOpcode but got " << lexer.SeeToken().type
          << std::endl;
      return false;
    }
    instruction->set_type(Statement_Type_INSTRUCTION);
    return ParseInstruction(instruction->mutable_instruction());
  }

  bool Parser::ParseInstruction(Instruction *instruction) {
    std::cout << "ParseInstruction2" << std::endl;
    // std::cout << "ParseInstruction(Instruction) not yet implemented." << std::endl;
    auto token = lexer.EatToken();
    auto opcode = instruction->mutable_opcode();
    if (Type::kAdvancedOpcode == token.type) {
      opcode->set_type(Opcode_Type_ADVANCED);
      auto advanced = Opcode::Advanced();
      Opcode::Advanced_Parse(token.value, &advanced);
      opcode->set_advanced(advanced);
      if (!ParseOperand(instruction->mutable_operand_a())) {
        return false;
      }
      return true;
    } else if (Type::kBasicOpcode == token.type) {
      opcode->set_type(Opcode_Type_BASIC);
      auto basic = Opcode::Basic();
      Opcode::Basic_Parse(token.value, &basic);
      opcode->set_basic(basic);
      if (!ParseOperand(instruction->mutable_operand_a())) {
        return false;
      }
      MaybeEatWhitespace();
      if (Type::kComma != lexer.SeeToken().type) {
        return false;
      }
      lexer.EatToken();
      if (!ParseOperand(instruction->mutable_operand_b())) {
        return false;
      }
      return true;
    } else {
      std::cout << "logic error" << std::endl;
      return false;
    }

    return false;
  }

  bool Parser::ParseOperand(Operand *operand) {
    std::cout << "ParseOperand" << std::endl;
    // std::cout << "ParseOperand(Operand) not yet implemented." << std::endl;
    MaybeEatWhitespace();
    if (Type::kIdentifier == lexer.SeeToken().type) {
      auto token = lexer.EatToken();
      if ("a" == token.value) {
        operand->set_type(Operand_Type_REGISTER);
        operand->set_register_(Operand_Register_A);
        return true;
      }
    } else if (Type::kHexadecimal == lexer.SeeToken().type) {
      auto token = lexer.EatToken();
      operand->set_type(Operand_Type_LITERAL);
      operand->set_value(ConvertHexadecimalValue(token.value));
      return true;
    }
    return false;
  }

  bool Parser::ParseLabel(Statement *label) {
    std::cout << "ParseLabel" << std::endl;
    if (Type::kIdentifier != lexer.SeeToken().type) {
      std::cout << "Expected kIdentifier for label but got " << lexer.SeeToken().type << std::endl;
      return false;
    }
    auto token = lexer.EatToken();
    if (Type::kColon != lexer.SeeToken().type) {
      std::cout << "Expected kColon after label but got " << lexer.SeeToken().type << std::endl;
      return false;
    }
    lexer.EatToken();
    label->set_type(Statement_Type_LABEL);
    label->set_label(token.value);
    return true;
  }

}  // namespace dcpu
