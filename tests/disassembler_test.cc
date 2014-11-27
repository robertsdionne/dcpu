// Copyright 2012 Robert Scott Dionne. All rights reserved.

#include <algorithm>
#include <sstream>
#include <gtest/gtest.h>

#include "dcpu.hpp"
#include "disassembler.hpp"

using namespace dcpu;

TEST(DisassemblerTest, Disassemble) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    0,
    0,
    0,
    0,
    0,
    0,
    0,
    0,
    0,
    0
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, a\n"
      "set a, a\n"
      "set a, a\n"
      "set a, a\n"
      "set a, a\n"
      "set a, a\n"
      "set a, a\n"
      "set a, a\n"
      "set a, a\n"
      "set a, a\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_register_with_register) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set b, 1
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterB, Dcpu::Operand::k1),
    // set a, b
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kRegisterB)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set b, 0x1\n"
      "set a, b\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_register_with_last_register) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set j, 1
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterJ, Dcpu::Operand::k1),
    // set a, j
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kRegisterJ)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set j, 0x1\n"
      "set a, j\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_register_with_location_in_register) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set [0x1000], 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kLocation, Dcpu::Operand::k13),
    0x1000,
    // set b, 0x1000
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterB, Dcpu::Operand::kLiteral),
    0x1000,
    // set a, [b]
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLocationInRegisterB)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set [0x1000], 0xd\n"
      "set b, 0x1000\n"
      "set a, [b]\n", out.str());
}

TEST(DisassemblerTest,
    Disassemble_set_register_with_location_in_last_register) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set [0x1000], 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kLocation, Dcpu::Operand::k13),
    0x1000,
    // set j, 0x1000
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterJ, Dcpu::Operand::kLiteral),
    0x1000,
    // set a, [j]
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLocationInRegisterJ)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set [0x1000], 0xd\n"
      "set j, 0x1000\n"
      "set a, [j]\n", out.str());
}

TEST(DisassemblerTest,
    Disassemble_set_register_with_location_offset_by_register) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set [0x100A], 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kLocation, Dcpu::Operand::k13),
    0x100A,
    // set b, 10
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterB, Dcpu::Operand::k10),
    // set a, [0x1000+b]
    Dcpu::Instruct(
        Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLocationOffsetByRegisterB),
    0x1000
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set [0x100a], 0xd\n"
      "set b, 0xa\n"
      "set a, [0x1000+b]\n", out.str());
}

TEST(DisassemblerTest,
    Disassemble_set_register_with_location_offset_by_last_register) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set [0x100A], 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kLocation, Dcpu::Operand::k13),
    0x100A,
    // set j, 10
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterJ, Dcpu::Operand::k10),
    // set a, [0x1000+j]
    Dcpu::Instruct(
        Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLocationOffsetByRegisterJ),
    0x1000
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set [0x100a], 0xd\n"
      "set j, 0xa\n"
      "set a, [0x1000+j]\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_register_with_pop) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set a, pop
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kPop)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set push, 0xd\n"
      "set a, pop\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_register_with_peek) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set a, peek
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kPeek)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set push, 0xd\n"
      "set a, peek\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_register_with_pick) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k14),
    // set a, [sp+1]
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kPick),
    0x1
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set push, 0xd\n"
      "set push, 0xe\n"
      "set a, [sp+0x1]\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_register_with_stack_pointer) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set a, sp
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kStackPointer)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set push, 0xd\n"
      "set a, sp\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_register_with_program_counter) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // noop
    Dcpu::Noop(),
    // set a, pc
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kProgramCounter)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, a\n"
      "set a, pc\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_register_with_overflow) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set ex, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kExtra, Dcpu::Operand::k13),
    // set a, ex
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kExtra)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set ex, 0xd\n"
      "set a, ex\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_register_with_location) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set [0x1000], 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kLocation, Dcpu::Operand::k13),
    0x1000,
    // set a, [0x1000]
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLocation),
    0x1000
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set [0x1000], 0xd\n"
      "set a, [0x1000]\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_register_with_high_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x1001
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0x1001
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x1001\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 1
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k1),
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x1\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_last_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set j, 1
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterJ, Dcpu::Operand::k1),
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set j, 0x1\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_location_in_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x1000
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0x1000,
    // set [a], 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kLocationInRegisterA, Dcpu::Operand::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x1000\n"
      "set [a], 0xd\n", out.str());
}

TEST(DisassemblerTest,
    Disassemble_set_location_in_last_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set j, 0x1000
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterJ, Dcpu::Operand::kLiteral),
    0x1000,
    // set [j], 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kLocationInRegisterJ, Dcpu::Operand::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set j, 0x1000\n"
      "set [j], 0xd\n", out.str());
}

TEST(DisassemblerTest,
    Disassemble_set_location_offset_by_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 10
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k10),
    // set [0x1000+a], 13
    Dcpu::Instruct(
        Dcpu::BasicOpcode::kSet, Dcpu::Operand::kLocationOffsetByRegisterA, Dcpu::Operand::k13),
    0x1000
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xa\n"
      "set [0x1000+a], 0xd\n", out.str());
}

TEST(DisassemblerTest,
    Disassemble_set_location_offset_by_register_with_location_offset_by_register) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 10
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k10),
    // set [0x1000+a], 13
    Dcpu::Instruct(
        Dcpu::BasicOpcode::kSet, Dcpu::Operand::kLocationOffsetByRegisterA, Dcpu::Operand::k13),
    0x1000,
    // set [0x2000+a], [0x1000+a]
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet,
        Dcpu::Operand::kLocationOffsetByRegisterA, Dcpu::Operand::kLocationOffsetByRegisterA),
    0x1000,
    0x2000
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xa\n"
      "set [0x1000+a], 0xd\n"
      "set [0x2000+a], [0x1000+a]\n", out.str());
}


TEST(DisassemblerTest,
    Disassemble_set_location_offset_by_last_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set j, 10
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterJ, Dcpu::Operand::k10),
    // set [0x1000+j], 13
    Dcpu::Instruct(
        Dcpu::BasicOpcode::kSet, Dcpu::Operand::kLocationOffsetByRegisterJ, Dcpu::Operand::k13),
    0x1000
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set j, 0xa\n"
      "set [0x1000+j], 0xd\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_push_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set push, 0xd\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_push_with_pop) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set push, pop
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::kPop)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set push, 0xd\n"
      "set push, pop\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_peek_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set peek, 14
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPeek, Dcpu::Operand::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set push, 0xd\n"
      "set peek, 0xe\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_pick_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set push, 12
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k12),
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set [sp+1], 14
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPick, Dcpu::Operand::k14),
    0x1
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set push, 0xc\n"
      "set push, 0xd\n"
      "set [sp+0x1], 0xe\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_stack_pointer_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set sp, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kStackPointer, Dcpu::Operand::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set sp, 0xd\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_program_counter_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set pc, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kProgramCounter, Dcpu::Operand::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set pc, 0xd\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_overflow_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set ex, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kExtra, Dcpu::Operand::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set ex, 0xd\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_location_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set [0x1000], 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kLocation, Dcpu::Operand::k13),
    0x1000
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set [0x1000], 0xd\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_literal_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set 0x1000, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kLiteral, Dcpu::Operand::k13),
    0x1000
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set 0x1000, 0xd\n", out.str());
}

TEST(DisassemblerTest, Disassemble_add_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x0D
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k13),
    // add a, 0x0E
    Dcpu::Instruct(Dcpu::BasicOpcode::kAdd, Dcpu::Operand::kRegisterA, Dcpu::Operand::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xd\n"
      "add a, 0xe\n", out.str());
}

TEST(DisassemblerTest, Disassemble_add_register_with_overflow) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0xFFFF
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0xFFFF,
    // add a, 0xFFFF
    Dcpu::Instruct(Dcpu::BasicOpcode::kAdd, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0xFFFF
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xffff\n"
      "add a, 0xffff\n", out.str());
}

TEST(DisassemblerTest, Disassemble_subtract_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x1E
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k30),
    // sub a, 0x10
    Dcpu::Instruct(Dcpu::BasicOpcode::kSubtract, Dcpu::Operand::kRegisterA, Dcpu::Operand::k16)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x1e\n"
      "sub a, 0x10\n", out.str());
}

TEST(DisassemblerTest, Disassemble_subtract_register_with_underflow) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x10
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k16),
    // sub a, 0x1E
    Dcpu::Instruct(Dcpu::BasicOpcode::kSubtract, Dcpu::Operand::kRegisterA, Dcpu::Operand::k30)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x10\n"
      "sub a, 0x1e\n", out.str());
}

TEST(DisassemblerTest, Disassemble_multiply_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x10
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k16),
    // mul a, 0x1E
    Dcpu::Instruct(Dcpu::BasicOpcode::kMultiply, Dcpu::Operand::kRegisterA, Dcpu::Operand::k30)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x10\n"
      "mul a, 0x1e\n", out.str());
}

TEST(DisassemblerTest, Disassemble_multiply_register_with_overflow) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0xFFFF
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0xFFFF,
    // mul a, 0xFFFF
    Dcpu::Instruct(Dcpu::BasicOpcode::kMultiply, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0xFFFF
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xffff\n"
      "mul a, 0xffff\n", out.str());
}

TEST(DisassemblerTest, Disassemble_divide_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x1E
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k30),
    // div a, 0x10
    Dcpu::Instruct(Dcpu::BasicOpcode::kDivide, Dcpu::Operand::kRegisterA, Dcpu::Operand::k16)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x1e\n"
      "div a, 0x10\n", out.str());
}

TEST(DisassemblerTest, Disassemble_divide_register_by_zero) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x1E
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k30),
    // div a, 0x00
    Dcpu::Instruct(Dcpu::BasicOpcode::kDivide, Dcpu::Operand::kRegisterA, Dcpu::Operand::k0)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x1e\n"
      "div a, 0\n", out.str());
}

TEST(DisassemblerTest, Disassemble_modulo_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x1E
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k30),
    // mod a, 0x0B
    Dcpu::Instruct(Dcpu::BasicOpcode::kModulo, Dcpu::Operand::kRegisterA, Dcpu::Operand::k11)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x1e\n"
      "mod a, 0xb\n", out.str());
}

TEST(DisassemblerTest, Disassemble_shift_left_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x1E
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k30),
    // shl a, 0x02
    Dcpu::Instruct(Dcpu::BasicOpcode::kShiftLeft, Dcpu::Operand::kRegisterA, Dcpu::Operand::k2)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x1e\n"
      "shl a, 0x2\n", out.str());
}

TEST(DisassemblerTest, Disassemble_shift_left_register_with_overflow) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0xFFFF
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0xFFFF,
    // shl a, 0x02
    Dcpu::Instruct(Dcpu::BasicOpcode::kShiftLeft, Dcpu::Operand::kRegisterA, Dcpu::Operand::k2)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xffff\n"
      "shl a, 0x2\n", out.str());
}

TEST(DisassemblerTest, Disassemble_shift_right_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0xFFF0
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0xFFF0,
    // shr a, 0x02
    Dcpu::Instruct(Dcpu::BasicOpcode::kShiftRight, Dcpu::Operand::kRegisterA, Dcpu::Operand::k2)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xfff0\n"
      "shr a, 0x2\n", out.str());
}

TEST(DisassemblerTest, Disassemble_shift_right_register_with_underflow) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0xFFFF
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0xFFFF,
    // shr a, 0x02
    Dcpu::Instruct(Dcpu::BasicOpcode::kShiftRight, Dcpu::Operand::kRegisterA, Dcpu::Operand::k2)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xffff\n"
      "shr a, 0x2\n", out.str());
}

TEST(DisassemblerTest, Disassemble_and_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0xF0F0
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0xF0F0,
    // and a, 0x00FF
    Dcpu::Instruct(Dcpu::BasicOpcode::kBinaryAnd, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0x00FF
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xf0f0\n"
      "and a, 0xff\n", out.str());
}

TEST(DisassemblerTest, Disassemble_or_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0xF0F0
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0xF0F0,
    // bor a, 0x00FF
    Dcpu::Instruct(Dcpu::BasicOpcode::kBinaryOr, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0x00FF
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xf0f0\n"
      "bor a, 0xff\n", out.str());
}

TEST(DisassemblerTest, Disassemble_xor_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0xF0F0
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0xF0F0,
    // xor a, 0x00FF
    Dcpu::Instruct(Dcpu::BasicOpcode::kBinaryExclusiveOr, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0x00FF
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xf0f0\n"
      "xor a, 0xff\n", out.str());
}

TEST(DisassemblerTest, Disassemble_if_equal_register_with_equal_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x0F
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k15),
    // ife a, 0x0F
    Dcpu::Instruct(Dcpu::BasicOpcode::kIfEqual, Dcpu::Operand::kRegisterA, Dcpu::Operand::k15),
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xf\n"
      "ife a, 0xf\n"
      "set push, 0xd\n"
      "set push, 0xe\n", out.str());
}

TEST(DisassemblerTest, Disassemble_if_equal_register_with_unequal_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x0F
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k15),
    // ife a, 0x00
    Dcpu::Instruct(Dcpu::BasicOpcode::kIfEqual, Dcpu::Operand::kRegisterA, Dcpu::Operand::k0),
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xf\n"
      "ife a, 0\n"
      "set push, 0xd\n"
      "set push, 0xe\n", out.str());
}

TEST(DisassemblerTest,
    Disassemble_if_not_equal_register_with_unequal_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x0F
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k15),
    // ifn a, 0x00
    Dcpu::Instruct(Dcpu::BasicOpcode::kIfNotEqual, Dcpu::Operand::kRegisterA, Dcpu::Operand::k0),
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xf\n"
      "ifn a, 0\n"
      "set push, 0xd\n"
      "set push, 0xe\n", out.str());
}

TEST(DisassemblerTest,
    Disassemble_if_not_equal_register_with_equal_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x0F
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k15),
    // ife a, 0x0F
    Dcpu::Instruct(Dcpu::BasicOpcode::kIfNotEqual, Dcpu::Operand::kRegisterA, Dcpu::Operand::k15),
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xf\n"
      "ifn a, 0xf\n"
      "set push, 0xd\n"
      "set push, 0xe\n", out.str());
}

TEST(DisassemblerTest,
    Disassemble_if_greater_than_register_with_lesser_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x1E
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k30),
    // ifg a, 0x0F
    Dcpu::Instruct(Dcpu::BasicOpcode::kIfGreaterThan, Dcpu::Operand::kRegisterA, Dcpu::Operand::k15),
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x1e\n"
      "ifg a, 0xf\n"
      "set push, 0xd\n"
      "set push, 0xe\n", out.str());
}

TEST(DisassemblerTest,
    Disassemble_if_greater_than_register_with_greater_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x0F
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k15),
    // ifg a, 0x1E
    Dcpu::Instruct(Dcpu::BasicOpcode::kIfGreaterThan, Dcpu::Operand::kRegisterA, Dcpu::Operand::k30),
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xf\n"
      "ifg a, 0x1e\n"
      "set push, 0xd\n"
      "set push, 0xe\n", out.str());
}

TEST(DisassemblerTest,
    Disassemble_if_both_register_with_common_bits_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x1E
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k30),
    // ifb a, 0x10
    Dcpu::Instruct(Dcpu::BasicOpcode::kIfBitSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k16),
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x1e\n"
      "ifb a, 0x10\n"
      "set push, 0xd\n"
      "set push, 0xe\n", out.str());
}

TEST(DisassemblerTest,
    Disassemble_if_both_register_with_uncommon_bits_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x0F
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k15),
    // ife a, 0x10
    Dcpu::Instruct(Dcpu::BasicOpcode::kIfBitSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k16),
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xf\n"
      "ifb a, 0x10\n"
      "set push, 0xd\n"
      "set push, 0xe\n", out.str());
}

TEST(DisassemblerTest, Disassemble_jump_sub_routine) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // jsr subroutine
    Dcpu::Instruct(Dcpu::AdvancedOpcode::kJumpSubRoutine, Dcpu::Operand::k3),
    // set a, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k13),
    // sub pc, 1
    Dcpu::Instruct(Dcpu::BasicOpcode::kSubtract, Dcpu::Operand::kProgramCounter, Dcpu::Operand::k1),
    // subroutine: set b, 14
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterB, Dcpu::Operand::k14),
    // set pc, pop
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kProgramCounter, Dcpu::Operand::kPop)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "jsr 0x3\n"
      "set a, 0xd\n"
      "sub pc, 0x1\n"
      "set b, 0xe\n"
      "set pc, pop\n", out.str());
}
