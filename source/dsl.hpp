#ifndef DCPU_DSL_HPP_
#define DCPU_DSL_HPP_

#include <string>

#include "generated/program.pb.h"

namespace dcpu {

  #define BASIC_OPCODE(opcode) template <typename A>\
    Dsl &opcode(proto::Operand b, A a) {\
      return Instruction(Opcode_Basic_##opcode, b, a);\
    }

  #define ADVANCED_OPCODE(opcode) template <typename A>\
    Dsl &opcode(A a) {\
      return Instruction(Opcode_Advanced_##opcode, a);\
    }

  #define ADVANCED_OPCODE_IGNORE_ARGUMENT(opcode)\
    Dsl &opcode() {\
      return Instruction(Opcode_Advanced_##opcode, 0);\
    }

  using Word = unsigned short;

  namespace dsl {

    using namespace proto;

    extern proto::Operand a;

    extern proto::Operand b;

    extern proto::Operand c;

    extern proto::Operand x;

    extern proto::Operand y;

    extern proto::Operand z;

    extern proto::Operand i;

    extern proto::Operand j;

    extern proto::Operand sp;

    extern proto::Operand pc;

    extern proto::Operand ex;

    extern proto::Operand push;

    extern proto::Operand pop;

    extern proto::Operand peek;

    proto::Operand pick(const std::string &label);

    proto::Operand pick(Word n);

    proto::Operand operator +(proto::Operand a, proto::Operand b);

    proto::Operand operator +(const std::string &label, proto::Operand b);

    proto::Operand operator +(Word literal, proto::Operand b);

    proto::Operand operator +(proto::Operand a, const std::string &label);

    proto::Operand operator +(proto::Operand a, Word literal);

    static proto::Operand MakeRegister(Operand_Register register_);

    static proto::Operand MakeSpecialRegister(Operand_Type type);

    class Dsl {
    public:
      Dsl() = default;

      virtual ~Dsl() = default;

      proto::Operand operator [](proto::Operand operand);

      proto::Operand operator [](const std::string &label);

      proto::Operand operator [](Word literal);

      Dsl &label(const std::string &label);

      template <typename... Arguments>
      Dsl &data(Word value, Arguments... arguments) {
        data(value);
        data(arguments...);
        return *this;
      }

      Dsl &data(Word value);

      Dsl &data(const std::string &string);

      BASIC_OPCODE(set);
      BASIC_OPCODE(add);
      BASIC_OPCODE(sub);
      BASIC_OPCODE(mul);
      BASIC_OPCODE(mli);
      BASIC_OPCODE(div);
      BASIC_OPCODE(dvi);
      BASIC_OPCODE(mod);
      BASIC_OPCODE(mdi);
      BASIC_OPCODE(and_);
      BASIC_OPCODE(bor);
      BASIC_OPCODE(xor_);
      BASIC_OPCODE(shr);
      BASIC_OPCODE(asr);
      BASIC_OPCODE(shl);
      BASIC_OPCODE(ifb);
      BASIC_OPCODE(ifc);
      BASIC_OPCODE(ife);
      BASIC_OPCODE(ifn);
      BASIC_OPCODE(ifg);
      BASIC_OPCODE(ifa);
      BASIC_OPCODE(ifl);
      BASIC_OPCODE(ifu);
      BASIC_OPCODE(adx);
      BASIC_OPCODE(sbx);
      BASIC_OPCODE(sti);
      BASIC_OPCODE(std);

      ADVANCED_OPCODE(jsr);
      ADVANCED_OPCODE(int_);
      ADVANCED_OPCODE(iag);
      ADVANCED_OPCODE(ias);
      ADVANCED_OPCODE_IGNORE_ARGUMENT(rfi);
      ADVANCED_OPCODE(iaq);
      ADVANCED_OPCODE(hwn);
      ADVANCED_OPCODE(hwq);
      ADVANCED_OPCODE(hwi);

      void Assemble(Word *const memory_begin);

    private:
      inline Dsl &Instruction(proto::Opcode::Basic basic, proto::Operand b, proto::Operand a) {
        proto::Statement *statement = program.add_statement();
        statement->set_type(Statement_Type_INSTRUCTION);
        proto::Instruction *instruction = statement->mutable_instruction();
        proto::Opcode *opcode = instruction->mutable_opcode();
        opcode->set_type(Opcode_Type_BASIC);
        opcode->set_basic(basic);
        instruction->mutable_operand_b()->CopyFrom(b);
        instruction->mutable_operand_a()->CopyFrom(a);
        return *this;
      }

      inline Dsl &Instruction(
          proto::Opcode::Basic basic, proto::Operand b, const std::string &label) {
        proto::Operand a;
        a.set_type(Operand_Type_LITERAL);
        a.set_label(label);
        return Instruction(basic, b, a);
      }

      inline Dsl &Instruction(proto::Opcode::Basic basic, proto::Operand b, Word literal) {
        proto::Operand a;
        a.set_type(Operand_Type_LITERAL);
        a.set_value(literal);
        return Instruction(basic, b, a);
      }

      inline Dsl &Instruction(proto::Opcode::Advanced advanced, proto::Operand a) {
        proto::Statement *statement = program.add_statement();
        statement->set_type(Statement_Type_INSTRUCTION);
        proto::Instruction *instruction = statement->mutable_instruction();
        proto::Opcode *opcode = instruction->mutable_opcode();
        opcode->set_type(Opcode_Type_ADVANCED);
        opcode->set_advanced(advanced);
        instruction->mutable_operand_a()->CopyFrom(a);
        return *this;
      }

      inline Dsl &Instruction(proto::Opcode::Advanced advanced, const std::string &label) {
        proto::Operand a;
        a.set_type(Operand_Type_LITERAL);
        a.set_label(label);
        return Instruction(advanced, a);
      }

      inline Dsl &Instruction(proto::Opcode::Advanced advanced, Word literal) {
        proto::Operand a;
        a.set_type(Operand_Type_LITERAL);
        a.set_value(literal);
        return Instruction(advanced, a);
      }

    private:
      proto::Program program;
    };

  }  // namespace dsl

}  // namespace dcpu

#endif  // DCPU_DSL_HPP_
