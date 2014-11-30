#ifndef DCPU_DSL_HPP_
#define DCPU_DSL_HPP_

#include <string>

#include "generated/program.pb.h"

namespace dcpu {

  using Word = unsigned short;

  namespace dsl {

    using namespace proto;

    proto::Operand a();

    proto::Operand b();

    proto::Operand c();

    proto::Operand x();

    proto::Operand y();

    proto::Operand z();

    proto::Operand i();

    proto::Operand j();

    proto::Operand sp();

    proto::Operand pc();

    proto::Operand ex();

    proto::Operand label(const std::string &label);

    proto::Operand value(Word value);

    proto::Operand operator +(proto::Operand a, proto::Operand b);

    proto::Operand operator +(const std::string &label, proto::Operand b);

    proto::Operand operator +(Word literal, proto::Operand b);

    proto::Operand operator +(proto::Operand a, const std::string &label);

    proto::Operand operator +(proto::Operand a, Word literal);

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

      template <typename A>
      Dsl &set(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_SET, b, a);
      }

      template <typename A>
      Dsl &add(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_ADD, b, a);
      }

      template <typename A>
      Dsl &sub(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_SUB, b, a);
      }

      template <typename A>
      Dsl &mul(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_MUL, b, a);
      }

      template <typename A>
      Dsl &mli(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_MLI, b, a);
      }

      template <typename A>
      Dsl &div(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_DIV, b, a);
      }

      template <typename A>
      Dsl &dvi(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_DVI, b, a);
      }

      template <typename A>
      Dsl &mod(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_MOD, b, a);
      }

      template <typename A>
      Dsl &mdi(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_MDI, b, a);
      }

      template <typename A>
      Dsl &and_(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_AND, b, a);
      }

      template <typename A>
      Dsl &bor(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_BOR, b, a);
      }

      template <typename A>
      Dsl &xor_(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_XOR, b, a);
      }

      template <typename A>
      Dsl &shr(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_SHR, b, a);
      }

      template <typename A>
      Dsl &asr(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_ASR, b, a);
      }

      template <typename A>
      Dsl &shl(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_SHL, b, a);
      }

      template <typename A>
      Dsl &ifb(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_IFB, b, a);
      }

      template <typename A>
      Dsl &ifc(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_IFC, b, a);
      }

      template <typename A>
      Dsl &ife(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_IFE, b, a);
      }

      template <typename A>
      Dsl &ifn(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_IFN, b, a);
      }

      template <typename A>
      Dsl &ifg(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_IFG, b, a);
      }

      template <typename A>
      Dsl &ifa(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_IFA, b, a);
      }

      template <typename A>
      Dsl &ifl(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_IFL, b, a);
      }

      template <typename A>
      Dsl &ifu(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_IFU, b, a);
      }

      template <typename A>
      Dsl &adx(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_ADX, b, a);
      }

      template <typename A>
      Dsl &sbx(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_SBX, b, a);
      }

      template <typename A>
      Dsl &sti(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_STI, b, a);
      }

      template <typename A>
      Dsl &std(proto::Operand b, A a) {
        return Instruction(Opcode_Basic_STD, b, a);
      }

      template <typename A>
      Dsl &jsr(A a) {
        return Instruction(Opcode_Advanced_JSR, a);
      }

      template <typename A>
      Dsl &int_(A a) {
        return Instruction(Opcode_Advanced_INT, a);
      }

      template <typename A>
      Dsl &iag(A a) {
        return Instruction(Opcode_Advanced_IAG, a);
      }

      template <typename A>
      Dsl &ias(A a) {
        return Instruction(Opcode_Advanced_IAS, a);
      }

      Dsl &rfi() {
        return Instruction(Opcode_Advanced_RFI, 0);
      }

      template <typename A>
      Dsl &iaq(A a) {
        return Instruction(Opcode_Advanced_IAQ, a);
      }

      template <typename A>
      Dsl &hwn(A a) {
        return Instruction(Opcode_Advanced_HWN, a);
      }

      template <typename A>
      Dsl &hwq(A a) {
        return Instruction(Opcode_Advanced_HWQ, a);
      }

      template <typename A>
      Dsl &hwi(A a) {
        return Instruction(Opcode_Advanced_HWI, a);
      }

      void Assemble(const Word *const memory_begin, const Word *const memory_end);

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
