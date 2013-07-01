// Copyright 2012 Robert Scott Dionne. All rights reserved.

#include <algorithm>
#include <gtest/gtest.h>

#include "dcpu.h"

TEST(DcpuTest, DefaultConstructor) {
  Dcpu dcpu;
  EXPECT_EQ(0, *dcpu.address(0x1000));
  EXPECT_EQ(0, dcpu.register_a());
  EXPECT_EQ(0, dcpu.register_b());
  EXPECT_EQ(0, dcpu.register_c());
  EXPECT_EQ(0, dcpu.register_x());
  EXPECT_EQ(0, dcpu.register_y());
  EXPECT_EQ(0, dcpu.register_z());
  EXPECT_EQ(0, dcpu.register_i());
  EXPECT_EQ(0, dcpu.register_j());
  EXPECT_EQ(0, dcpu.program_counter());
  EXPECT_EQ(0, dcpu.stack_pointer());
  EXPECT_EQ(0, dcpu.extra());
  EXPECT_EQ(0, dcpu.interrupt_address());
}

TEST(DcpuTest, Reset) {
  Dcpu dcpu;
  *dcpu.address(0x1000) = 1;
  dcpu.register_a() = 2;
  dcpu.register_b() = 3;
  dcpu.register_c() = 4;
  dcpu.register_x() = 5;
  dcpu.register_y() = 6;
  dcpu.register_z() = 7;
  dcpu.register_i() = 8;
  dcpu.register_j() = 9;
  dcpu.program_counter() = 10;
  dcpu.stack_pointer() = 11;
  dcpu.extra() = 12;
  dcpu.interrupt_address() = 13;
  EXPECT_EQ(1, *dcpu.address(0x1000));
  EXPECT_EQ(2, dcpu.register_a());
  EXPECT_EQ(3, dcpu.register_b());
  EXPECT_EQ(4, dcpu.register_c());
  EXPECT_EQ(5, dcpu.register_x());
  EXPECT_EQ(6, dcpu.register_y());
  EXPECT_EQ(7, dcpu.register_z());
  EXPECT_EQ(8, dcpu.register_i());
  EXPECT_EQ(9, dcpu.register_j());
  EXPECT_EQ(10, dcpu.program_counter());
  EXPECT_EQ(11, dcpu.stack_pointer());
  EXPECT_EQ(12, dcpu.extra());
  EXPECT_EQ(13, dcpu.interrupt_address());
  dcpu.Reset();
  EXPECT_EQ(0, *dcpu.address(0x1000));
  EXPECT_EQ(0, dcpu.register_a());
  EXPECT_EQ(0, dcpu.register_b());
  EXPECT_EQ(0, dcpu.register_c());
  EXPECT_EQ(0, dcpu.register_x());
  EXPECT_EQ(0, dcpu.register_y());
  EXPECT_EQ(0, dcpu.register_z());
  EXPECT_EQ(0, dcpu.register_i());
  EXPECT_EQ(0, dcpu.register_j());
  EXPECT_EQ(0, dcpu.program_counter());
  EXPECT_EQ(0, dcpu.stack_pointer());
  EXPECT_EQ(0, dcpu.extra());
  EXPECT_EQ(0, dcpu.interrupt_address());
}

TEST(DcpuTest, ExecuteInstructions) {
  Dcpu dcpu;
  // dcpu.program_counter() == 0 by DcpuTest.DefaultConstructor
  dcpu.ExecuteInstructions(10);
  EXPECT_EQ(10, dcpu.program_counter());
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_register) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set b, 1
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterB, Dcpu::Operand::k1),
    // set a, b
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kRegisterB)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  // dcpu.register_a() == 0 by DcpuTest.DefaultConstructor
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(1, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_last_register) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set j, 1
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterJ, Dcpu::Operand::k1),
    // set a, j
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kRegisterJ)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  // dcpu.register_a() == 0 by DcpuTest.DefaultConstructor
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(1, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_location_in_register) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_location_in_last_register) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest,
    ExecuteInstruction_set_register_with_location_offset_by_register) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest,
    ExecuteInstruction_set_register_with_location_offset_by_last_register) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_pop) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set a, pop
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kPop)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, dcpu.stack_pointer());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(0, dcpu.stack_pointer());
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_peek) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set a, peek
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kPeek)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, dcpu.stack_pointer());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, dcpu.stack_pointer());
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_pick) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  EXPECT_EQ(0, dcpu.stack_pointer());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(0xFFFE, dcpu.stack_pointer());
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer() + 1));
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_stack_pointer) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set a, sp
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kStackPointer)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, dcpu.stack_pointer());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_program_counter) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // noop
    Dcpu::Noop(),
    // set a, pc
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kProgramCounter)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(2, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_overflow) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set ex, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kExtra, Dcpu::Operand::k13),
    // set a, ex
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kExtra)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_location) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_high_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x1001
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0x1001
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(0x1001, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 1
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k1),
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(1, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_set_last_register_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set j, 1
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterJ, Dcpu::Operand::k1),
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  // dcpu.register_a() == 0 by DcpuTest.DefaultConstructor
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(1, dcpu.register_j());
}

TEST(DcpuTest, ExecuteInstruction_set_location_in_register_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x1000
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0x1000,
    // set [a], 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kLocationInRegisterA, Dcpu::Operand::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(13, *dcpu.address(0x1000));
}

TEST(DcpuTest,
    ExecuteInstruction_set_location_in_last_register_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set j, 0x1000
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterJ, Dcpu::Operand::kLiteral),
    0x1000,
    // set [j], 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kLocationInRegisterJ, Dcpu::Operand::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(13, *dcpu.address(0x1000));
}

TEST(DcpuTest,
    ExecuteInstruction_set_location_offset_by_register_with_low_literal) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(13, *dcpu.address(0x100A));
}

TEST(DcpuTest,
    ExecuteInstruction_set_location_offset_by_register_with_location_offset_by_register) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 10
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k10),
    // set [0x1000+a], 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kLocationOffsetByRegisterA, Dcpu::Operand::k13),
    0x1000,
    // set [0x2000+a], [0x1000+a]
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet,
        Dcpu::Operand::kLocationOffsetByRegisterA, Dcpu::Operand::kLocationOffsetByRegisterA),
    0x1000,
    0x2000
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *dcpu.address(0x200A));
}

TEST(DcpuTest,
    ExecuteInstruction_set_location_offset_by_last_register_with_low_literal) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(13, *dcpu.address(0x100A));
}

TEST(DcpuTest, ExecuteInstruction_set_push_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest, ExecuteInstruction_set_push_with_pop) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set push, pop
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::kPop)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0xFFFF, dcpu.stack_pointer());
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest, ExecuteInstruction_set_peek_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set peek, 14
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPeek, Dcpu::Operand::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest, ExecuteInstruction_set_pick_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k12),
    // set push, 14
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set [SP+0x1], 14
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPick, Dcpu::Operand::k14),
    0x1
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer()));
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer() + 1));
}

TEST(DcpuTest, ExecuteInstruction_set_stack_pointer_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set sp, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kStackPointer, Dcpu::Operand::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(13, dcpu.stack_pointer());
}

TEST(DcpuTest, ExecuteInstruction_set_program_counter_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set pc, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kProgramCounter, Dcpu::Operand::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(13, dcpu.program_counter());
}

TEST(DcpuTest, ExecuteInstruction_set_overflow_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set ex, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kExtra, Dcpu::Operand::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(13, dcpu.extra());
}

TEST(DcpuTest, ExecuteInstruction_set_location_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set [0x1000], 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kLocation, Dcpu::Operand::k13),
    0x1000
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(13, *dcpu.address(0x1000));
}

TEST(DcpuTest, ExecuteInstruction_set_literal_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set 0x1000, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kLiteral, Dcpu::Operand::k13),
    0x1000
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  // set 0x1000, 13 should be a noop.
  EXPECT_EQ(0, *dcpu.address(0x1000));
}

TEST(DcpuTest, ExecuteInstruction_add_register_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x0D
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k13),
    // add a, 0x0E
    Dcpu::Instruct(Dcpu::BasicOpcode::kAdd, Dcpu::Operand::kRegisterA, Dcpu::Operand::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0x1B, dcpu.register_a());
  EXPECT_EQ(0, dcpu.extra());
}

TEST(DcpuTest, ExecuteInstruction_add_register_with_overflow) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0xFFFE, dcpu.register_a());
  EXPECT_EQ(1, dcpu.extra());
}

TEST(DcpuTest, ExecuteInstruction_subtract_register_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x1E
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k30),
    // sub a, 0x10
    Dcpu::Instruct(Dcpu::BasicOpcode::kSubtract, Dcpu::Operand::kRegisterA, Dcpu::Operand::k16)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0xE, dcpu.register_a());
  EXPECT_EQ(0, dcpu.extra());
}

TEST(DcpuTest, ExecuteInstruction_subtract_register_with_underflow) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x10
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k16),
    // sub a, 0x1E
    Dcpu::Instruct(Dcpu::BasicOpcode::kSubtract, Dcpu::Operand::kRegisterA, Dcpu::Operand::k30)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0xFFF2, dcpu.register_a());
  EXPECT_EQ(1, dcpu.extra());
}

TEST(DcpuTest, ExecuteInstruction_multiply_register_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x10
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k16),
    // mul a, 0x1E
    Dcpu::Instruct(Dcpu::BasicOpcode::kMultiply, Dcpu::Operand::kRegisterA, Dcpu::Operand::k30)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0x01E0, dcpu.register_a());
  EXPECT_EQ(0, dcpu.extra());
}

TEST(DcpuTest, ExecuteInstruction_multiply_register_with_overflow) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0x0001, dcpu.register_a());
  EXPECT_EQ(0xFFFE, dcpu.extra());
}

TEST(DcpuTest, ExecuteInstruction_divide_register_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x1E
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k30),
    // div a, 0x10
    Dcpu::Instruct(Dcpu::BasicOpcode::kDivide, Dcpu::Operand::kRegisterA, Dcpu::Operand::k16)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(1, dcpu.register_a());
  EXPECT_EQ(0, dcpu.extra());
}

TEST(DcpuTest, ExecuteInstruction_divide_register_by_zero) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x1E
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k30),
    // div a, 0x00
    Dcpu::Instruct(Dcpu::BasicOpcode::kDivide, Dcpu::Operand::kRegisterA, Dcpu::Operand::k0)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0, dcpu.register_a());
  EXPECT_EQ(1, dcpu.extra());
}

TEST(DcpuTest, ExecuteInstruction_modulo_register_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x1E
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k30),
    // mod a, 0x0B
    Dcpu::Instruct(Dcpu::BasicOpcode::kModulo, Dcpu::Operand::kRegisterA, Dcpu::Operand::k11)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0x8, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_shift_left_register_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x1E
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k30),
    // shl a, 0x02
    Dcpu::Instruct(Dcpu::BasicOpcode::kShiftLeft, Dcpu::Operand::kRegisterA, Dcpu::Operand::k2)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0x78, dcpu.register_a());
  EXPECT_EQ(0, dcpu.extra());
}

TEST(DcpuTest, ExecuteInstruction_shift_left_register_with_overflow) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0xFFFF
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0xFFFF,
    // shl a, 0x02
    Dcpu::Instruct(Dcpu::BasicOpcode::kShiftLeft, Dcpu::Operand::kRegisterA, Dcpu::Operand::k2)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0xFFFC, dcpu.register_a());
  EXPECT_EQ(0x0003, dcpu.extra());
}

TEST(DcpuTest, ExecuteInstruction_shift_right_register_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0xFFF0
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0xFFF0,
    // shr a, 0x02
    Dcpu::Instruct(Dcpu::BasicOpcode::kShiftRight, Dcpu::Operand::kRegisterA, Dcpu::Operand::k2)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0x3FFC, dcpu.register_a());
  EXPECT_EQ(0, dcpu.extra());
}

TEST(DcpuTest, ExecuteInstruction_shift_right_register_with_underflow) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0xFFFF
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::kLiteral),
    0xFFFF,
    // shr a, 0x02
    Dcpu::Instruct(Dcpu::BasicOpcode::kShiftRight, Dcpu::Operand::kRegisterA, Dcpu::Operand::k2)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0x3FFF, dcpu.register_a());
  EXPECT_EQ(0xC000, dcpu.extra());
}

TEST(DcpuTest, ExecuteInstruction_and_register_with_low_literal) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0x00F0, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_or_register_with_low_literal) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0xF0FF, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_xor_register_with_low_literal) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0xF00F, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_if_equal_register_with_equal_low_literal) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest, ExecuteInstruction_if_equal_register_with_unequal_low_literal) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest,
    ExecuteInstruction_if_not_equal_register_with_unequal_low_literal) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest,
    ExecuteInstruction_if_not_equal_register_with_equal_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x0F
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k15),
    // ifn a, 0x0F
    Dcpu::Instruct(Dcpu::BasicOpcode::kIfNotEqual, Dcpu::Operand::kRegisterA, Dcpu::Operand::k15),
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest,
    ExecuteInstruction_if_greater_than_register_with_lesser_low_literal) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest,
    ExecuteInstruction_if_greater_than_register_with_greater_low_literal) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest,
    ExecuteInstruction_if_both_register_with_common_bits_low_literal) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest,
    ExecuteInstruction_if_both_register_with_uncommon_bits_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x0F
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k15),
    // ifb a, 0x10
    Dcpu::Instruct(Dcpu::BasicOpcode::kIfBitSet, Dcpu::Operand::kRegisterA, Dcpu::Operand::k16),
    // set push, 13
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::Operand::kPush, Dcpu::Operand::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest, ExecuteInstruction_jump_sub_routine) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, dcpu.stack_pointer());
  dcpu.ExecuteInstructions(4);
  EXPECT_EQ(13, dcpu.register_a());
  EXPECT_EQ(14, dcpu.register_b());
  EXPECT_EQ(2, dcpu.program_counter());
  EXPECT_EQ(0, dcpu.stack_pointer());
}
