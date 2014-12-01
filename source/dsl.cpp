#include <cassert>
#include <google/protobuf/text_format.h>
#include <google/protobuf/io/zero_copy_stream_impl.h>
#include <iostream>

#include "assembler.hpp"
#include "dsl.hpp"
#include "generated/program.pb.h"

namespace dcpu {

  namespace dsl {

    proto::Operand a() {
      proto::Operand result;
      result.set_type(Operand_Type_REGISTER);
      result.set_register_(Operand_Register_A);
      return result;
    }

    proto::Operand b() {
      proto::Operand result;
      result.set_type(Operand_Type_REGISTER);
      result.set_register_(Operand_Register_B);
      return result;
    }

    proto::Operand c() {
      proto::Operand result;
      result.set_type(Operand_Type_REGISTER);
      result.set_register_(Operand_Register_C);
      return result;
    }

    proto::Operand x() {
      proto::Operand result;
      result.set_type(Operand_Type_REGISTER);
      result.set_register_(Operand_Register_X);
      return result;
    }

    proto::Operand y() {
      proto::Operand result;
      result.set_type(Operand_Type_REGISTER);
      result.set_register_(Operand_Register_Y);
      return result;
    }

    proto::Operand z() {
      proto::Operand result;
      result.set_type(Operand_Type_REGISTER);
      result.set_register_(Operand_Register_Z);
      return result;
    }

    proto::Operand i() {
      proto::Operand result;
      result.set_type(Operand_Type_REGISTER);
      result.set_register_(Operand_Register_I);
      return result;
    }

    proto::Operand j() {
      proto::Operand result;
      result.set_type(Operand_Type_REGISTER);
      result.set_register_(Operand_Register_J);
      return result;
    }

    proto::Operand sp() {
      proto::Operand result;
      result.set_type(Operand_Type_STACK_POINTER);
      return result;
    }

    proto::Operand pc() {
      proto::Operand result;
      result.set_type(Operand_Type_PROGRAM_COUNTER);
      return result;
    }

    proto::Operand ex() {
      proto::Operand result;
      result.set_type(Operand_Type_EXTRA);
      return result;
    }

    proto::Operand push() {
      proto::Operand result;
      result.set_type(Operand_Type_PUSH_POP);
      return result;
    }

    proto::Operand pop() {
      return push();
    }

    proto::Operand peek() {
      proto::Operand result;
      result.set_type(Operand_Type_PEEK);
      return result;
    }

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
      auto out = new google::protobuf::io::OstreamOutputStream(&std::cout);
      google::protobuf::TextFormat::Print(program, out);
      delete out;
      program.Clear();
    }

  }  // namespace dsl

}  // namespace dsl
