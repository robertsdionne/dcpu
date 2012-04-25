// Copyright 2012 Robert Scott Dionne. All rights reserved.

#include <algorithm>
#include <sstream>
#include <gtest/gtest.h>
#include "dcpu.h"
#include "disassembler.h"

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
      "set 0, 0\n"
      "set 0, 0\n"
      "set 0, 0\n"
      "set 0, 0\n"
      "set 0, 0\n"
      "set 0, 0\n"
      "set 0, 0\n"
      "set 0, 0\n"
      "set 0, 0\n"
      "set 0, 0\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_register_with_register) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set b, 1
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterB, Dcpu::k1),
    // set a, b
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kRegisterB)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterJ, Dcpu::k1),
    // set a, j
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kRegisterJ)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocation, Dcpu::k13),
    0x1000,
    // set b, 0x1000
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterB, Dcpu::kLiteral),
    0x1000,
    // set a, [b]
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLocationInRegisterB)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocation, Dcpu::k13),
    0x1000,
    // set j, 0x1000
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterJ, Dcpu::kLiteral),
    0x1000,
    // set a, [j]
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLocationInRegisterJ)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocation, Dcpu::k13),
    0x100A,
    // set b, 10
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterB, Dcpu::k10),
    // set a, [0x1000+b]
    Dcpu::Instruct(
        Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLocationOffsetByRegisterB),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocation, Dcpu::k13),
    0x100A,
    // set j, 10
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterJ, Dcpu::k10),
    // set a, [0x1000+j]
    Dcpu::Instruct(
        Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLocationOffsetByRegisterJ),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set a, pop
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kPop)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set a, peek
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kPeek)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set push, 0xd\n"
      "set a, peek\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_register_with_push) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k13),
    // set a, push
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kPush)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xd\n"
      "set a, push\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_register_with_stack_pointer) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set a, sp
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kStackPointer)
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
    // reserved 0, 0
    Dcpu::Instruct(Dcpu::kBasicReserved, Dcpu::k0, Dcpu::k0),
    // set a, pc
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kProgramCounter)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set 0, 0\n"
      "set a, pc\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_register_with_overflow) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set ex, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kExtra, Dcpu::k13),
    // set a, ex
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kExtra)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocation, Dcpu::k13),
    0x1000,
    // set a, [0x1000]
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLocation),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k1),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterJ, Dcpu::k1),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0x1000,
    // set [a], 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocationInRegisterA, Dcpu::k13)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterJ, Dcpu::kLiteral),
    0x1000,
    // set [j], 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocationInRegisterJ, Dcpu::k13)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k10),
    // set [0x1000+a], 13
    Dcpu::Instruct(
        Dcpu::kSet, Dcpu::kLocationOffsetByRegisterA, Dcpu::k13),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k10),
    // set [0x1000+a], 13
    Dcpu::Instruct(
        Dcpu::kSet, Dcpu::kLocationOffsetByRegisterA, Dcpu::k13),
    0x1000,
    // set [0x2000+a], [0x1000+a]
    Dcpu::Instruct(Dcpu::kSet,
        Dcpu::kLocationOffsetByRegisterA, Dcpu::kLocationOffsetByRegisterA),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterJ, Dcpu::k10),
    // set [0x1000+j], 13
    Dcpu::Instruct(
        Dcpu::kSet, Dcpu::kLocationOffsetByRegisterJ, Dcpu::k13),
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

TEST(DisassemblerTest, Disassemble_set_pop_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set pop, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPop, Dcpu::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set push, 0xd\n"
      "set pop, 0xe\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_peek_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set peek, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPeek, Dcpu::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set push, 0xd\n"
      "set peek, 0xe\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_push_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set push, 0xd\n", out.str());
}

TEST(DisassemblerTest, Disassemble_set_stack_pointer_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set sp, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kStackPointer, Dcpu::k13)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kProgramCounter, Dcpu::k13)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kExtra, Dcpu::k13)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocation, Dcpu::k13),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLiteral, Dcpu::k13),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k13),
    // add a, 0x0E
    Dcpu::Instruct(Dcpu::kAdd, Dcpu::kRegisterA, Dcpu::k14)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xFFFF,
    // add a, 0xFFFF
    Dcpu::Instruct(Dcpu::kAdd, Dcpu::kRegisterA, Dcpu::kLiteral),
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
    // set a, 0x1F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k31),
    // sub a, 0x10
    Dcpu::Instruct(Dcpu::kSubtract, Dcpu::kRegisterA, Dcpu::k16)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x1f\n"
      "sub a, 0x10\n", out.str());
}

TEST(DisassemblerTest, Disassemble_subtract_register_with_underflow) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x10
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k16),
    // sub a, 0x1F
    Dcpu::Instruct(Dcpu::kSubtract, Dcpu::kRegisterA, Dcpu::k31)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x10\n"
      "sub a, 0x1f\n", out.str());
}

TEST(DisassemblerTest, Disassemble_multiply_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x10
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k16),
    // mul a, 0x1F
    Dcpu::Instruct(Dcpu::kMultiply, Dcpu::kRegisterA, Dcpu::k31)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x10\n"
      "mul a, 0x1f\n", out.str());
}

TEST(DisassemblerTest, Disassemble_multiply_register_with_overflow) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0xFFFF
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xFFFF,
    // mul a, 0xFFFF
    Dcpu::Instruct(Dcpu::kMultiply, Dcpu::kRegisterA, Dcpu::kLiteral),
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
    // set a, 0x1F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k31),
    // div a, 0x10
    Dcpu::Instruct(Dcpu::kDivide, Dcpu::kRegisterA, Dcpu::k16)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x1f\n"
      "div a, 0x10\n", out.str());
}

TEST(DisassemblerTest, Disassemble_divide_register_by_zero) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x1F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k31),
    // div a, 0x00
    Dcpu::Instruct(Dcpu::kDivide, Dcpu::kRegisterA, Dcpu::k0)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x1f\n"
      "div a, 0\n", out.str());
}

TEST(DisassemblerTest, Disassemble_modulo_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x1F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k31),
    // mod a, 0x0B
    Dcpu::Instruct(Dcpu::kModulo, Dcpu::kRegisterA, Dcpu::k11)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x1f\n"
      "mod a, 0xb\n", out.str());
}

TEST(DisassemblerTest, Disassemble_shift_left_register_with_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x1F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k31),
    // shl a, 0x02
    Dcpu::Instruct(Dcpu::kShiftLeft, Dcpu::kRegisterA, Dcpu::k2)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x1f\n"
      "shl a, 0x2\n", out.str());
}

TEST(DisassemblerTest, Disassemble_shift_left_register_with_overflow) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0xFFFF
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xFFFF,
    // shl a, 0x02
    Dcpu::Instruct(Dcpu::kShiftLeft, Dcpu::kRegisterA, Dcpu::k2)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xFFF0,
    // shr a, 0x02
    Dcpu::Instruct(Dcpu::kShiftRight, Dcpu::kRegisterA, Dcpu::k2)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xFFFF,
    // shr a, 0x02
    Dcpu::Instruct(Dcpu::kShiftRight, Dcpu::kRegisterA, Dcpu::k2)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xF0F0,
    // and a, 0x00FF
    Dcpu::Instruct(Dcpu::kBinaryAnd, Dcpu::kRegisterA, Dcpu::kLiteral),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xF0F0,
    // bor a, 0x00FF
    Dcpu::Instruct(Dcpu::kBinaryOr, Dcpu::kRegisterA, Dcpu::kLiteral),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xF0F0,
    // xor a, 0x00FF
    Dcpu::Instruct(Dcpu::kBinaryExclusiveOr, Dcpu::kRegisterA, Dcpu::kLiteral),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k15),
    // ife a, 0x0F
    Dcpu::Instruct(Dcpu::kIfEqual, Dcpu::kRegisterA, Dcpu::k15),
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k15),
    // ife a, 0x00
    Dcpu::Instruct(Dcpu::kIfEqual, Dcpu::kRegisterA, Dcpu::k0),
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k15),
    // ifn a, 0x00
    Dcpu::Instruct(Dcpu::kIfNotEqual, Dcpu::kRegisterA, Dcpu::k0),
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k15),
    // ife a, 0x0F
    Dcpu::Instruct(Dcpu::kIfNotEqual, Dcpu::kRegisterA, Dcpu::k15),
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14)
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
    // set a, 0x1F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k31),
    // ifg a, 0x0F
    Dcpu::Instruct(Dcpu::kIfGreaterThan, Dcpu::kRegisterA, Dcpu::k15),
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x1f\n"
      "ifg a, 0xf\n"
      "set push, 0xd\n"
      "set push, 0xe\n", out.str());
}

TEST(DisassemblerTest,
    Disassemble_if_greater_than_register_with_greater_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x0F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k15),
    // ifg a, 0x1F
    Dcpu::Instruct(Dcpu::kIfGreaterThan, Dcpu::kRegisterA, Dcpu::k31),
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0xf\n"
      "ifg a, 0x1f\n"
      "set push, 0xd\n"
      "set push, 0xe\n", out.str());
}

TEST(DisassemblerTest,
    Disassemble_if_both_register_with_common_bits_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x1F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k31),
    // ifb a, 0x10
    Dcpu::Instruct(Dcpu::kIfBoth, Dcpu::kRegisterA, Dcpu::k16),
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::ostringstream out;
  disassembler.Disassemble(program, program_end, out);
  EXPECT_EQ(
      "set a, 0x1f\n"
      "ifb a, 0x10\n"
      "set push, 0xd\n"
      "set push, 0xe\n", out.str());
}

TEST(DisassemblerTest,
    Disassemble_if_both_register_with_uncommon_bits_low_literal) {
  Disassembler disassembler;
  const Dcpu::Word program[] = {
    // set a, 0x0F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k15),
    // ife a, 0x10
    Dcpu::Instruct(Dcpu::kIfBoth, Dcpu::kRegisterA, Dcpu::k16),
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14)
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
    Dcpu::Instruct(Dcpu::kJumpSubRoutine, Dcpu::k3),
    // set a, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k13),
    // sub pc, 1
    Dcpu::Instruct(Dcpu::kSubtract, Dcpu::kProgramCounter, Dcpu::k1),
    // subroutine: set b, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterB, Dcpu::k14),
    // set pc, pop
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kProgramCounter, Dcpu::kPop)
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
