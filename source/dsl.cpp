#include <cassert>
#include <google/protobuf/text_format.h>
#include <google/protobuf/io/zero_copy_stream_impl.h>
#include <iostream>

#include "assembler.hpp"
#include "dsl.hpp"
#include "generated/program.pb.h"

namespace dcpu {

  namespace dsl {

    proto::Operand a = MakeRegister(Operand_Register_A);

    proto::Operand b = MakeRegister(Operand_Register_B);

    proto::Operand c = MakeRegister(Operand_Register_C);

    proto::Operand x = MakeRegister(Operand_Register_X);

    proto::Operand y = MakeRegister(Operand_Register_Y);

    proto::Operand z = MakeRegister(Operand_Register_Z);

    proto::Operand i = MakeRegister(Operand_Register_I);

    proto::Operand j = MakeRegister(Operand_Register_J);

    proto::Operand sp = MakeSpecialRegister(Operand_Type_STACK_POINTER);

    proto::Operand pc = MakeSpecialRegister(Operand_Type_PROGRAM_COUNTER);

    proto::Operand ex = MakeSpecialRegister(Operand_Type_EXTRA);

    proto::Operand push = MakeSpecialRegister(Operand_Type_PUSH_POP);

    proto::Operand pop = MakeSpecialRegister(Operand_Type_PUSH_POP);

    proto::Operand peek = MakeSpecialRegister(Operand_Type_PEEK);

    proto::Operand pick(const std::string &label) {
      proto::Operand result;
      result.set_type(Operand_Type_PICK);
      result.set_label(label);
      return result;
    }

    proto::Operand pick(Word literal) {
      proto::Operand result;
      result.set_type(Operand_Type_PICK);
      result.set_value(literal);
      return result;
    }

    proto::Operand MakeRegister(Operand_Register register_) {
      proto::Operand result;
      result.set_type(Operand_Type_REGISTER);
      result.set_register_(register_);
      return result;
    }

    proto::Operand MakeSpecialRegister(Operand_Type type) {
      proto::Operand result;
      result.set_type(type);
      return result;
    }

    proto::Operand operator +(proto::Operand a, proto::Operand b) {
      assert((Operand_Type_LITERAL == a.type() && Operand_Type_REGISTER == b.type())
          || (Operand_Type_REGISTER == a.type() && Operand_Type_LITERAL == b.type()));
      proto::Operand result;
      result.CopyFrom(a);
      if (Operand_Type_LITERAL == a.type() && Operand_Type_REGISTER == b.type()) {
        assert(a.has_value() || a.has_label());
        result.set_type(Operand_Type_LOCATION_OFFSET_BY_REGISTER);
        result.set_register_(b.register_());
      } else if (Operand_Type_REGISTER == a.type() && Operand_Type_LITERAL == b.type()) {
        assert(b.has_value() || b.has_label());
        result.set_type(Operand_Type_LOCATION_OFFSET_BY_REGISTER);
        if (b.has_value()) {
          result.set_value(b.value());
        } else if (b.has_label()) {
          result.set_label(b.label());
        }
      }
      return result;
    }

    proto::Operand operator +(const std::string &label, proto::Operand b) {
      proto::Operand a;
      a.set_type(Operand_Type_LITERAL);
      a.set_label(label);
      return a + b;
    }

    proto::Operand operator +(Word literal, proto::Operand b) {
      proto::Operand a;
      a.set_type(Operand_Type_LITERAL);
      a.set_value(literal);
      return a + b;
    }

    proto::Operand operator +(proto::Operand a, const std::string &label) {
      proto::Operand b;
      b.set_type(Operand_Type_LITERAL);
      b.set_label(label);
      return a + b;
    }

    proto::Operand operator +(proto::Operand a, Word literal) {
      proto::Operand b;
      b.set_type(Operand_Type_LITERAL);
      b.set_value(literal);
      return a + b;
    }

    proto::Operand Dsl::operator [](proto::Operand operand) {
      proto::Operand result;
      result.CopyFrom(operand);
      if (Operand_Type_REGISTER == operand.type()) {
        result.set_type(Operand_Type_LOCATION_IN_REGISTER);
      } else if (Operand_Type_LITERAL == operand.type()) {
        result.set_type(Operand_Type_LOCATION);
      }
      return result;
    }

    proto::Operand Dsl::operator [](const std::string &label) {
      proto::Operand result;
      result.set_type(Operand_Type_LOCATION);
      result.set_label(label);
      return result;
    }

    proto::Operand Dsl::operator [](Word literal) {
      proto::Operand result;
      result.set_type(Operand_Type_LOCATION);
      result.set_value(literal);
      return result;
    }

    Dsl &Dsl::label(const std::string &label) {
      proto::Statement *statement = program.add_statement();
      statement->set_type(Statement_Type_LABEL);
      statement->set_label(label);
      return *this;
    }

    Dsl &Dsl::data(Word value) {
      proto::Statement *statement = program.add_statement();
      statement->set_type(Statement_Type_DATA);
      statement->set_data(value);
      return *this;
    }

    Dsl &Dsl::data(const std::string &string) {
      for (auto character : string) {
        data(character);
      }
      data(0);
      return *this;
    }

    void Dsl::Assemble(Word *const memory_begin) {
      Assembler assembler;
      assembler.Assemble(program, memory_begin);
      program.Clear();
    }

  }  // namespace dsl

}  // namespace dsl
