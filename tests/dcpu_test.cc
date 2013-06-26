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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterB, Dcpu::k1),
    // set a, b
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kRegisterB)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterJ, Dcpu::k1),
    // set a, j
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kRegisterJ)
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_location_in_last_register) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest,
    ExecuteInstruction_set_register_with_location_offset_by_register) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest,
    ExecuteInstruction_set_register_with_location_offset_by_last_register) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_pop) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set a, pop
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kPop)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set a, peek
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kPeek)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14),
    // set a, [sp+1]
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kPick),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set a, sp
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kStackPointer)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kProgramCounter)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kExtra, Dcpu::k13),
    // set a, ex
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kExtra)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocation, Dcpu::k13),
    0x1000,
    // set a, [0x1000]
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLocation),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k1),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterJ, Dcpu::k1),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0x1000,
    // set [a], 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocationInRegisterA, Dcpu::k13)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterJ, Dcpu::kLiteral),
    0x1000,
    // set [j], 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocationInRegisterJ, Dcpu::k13)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k10),
    // set [0x1000+a], 13
    Dcpu::Instruct(
        Dcpu::kSet, Dcpu::kLocationOffsetByRegisterA, Dcpu::k13),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k10),
    // set [0x1000+a], 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocationOffsetByRegisterA, Dcpu::k13),
    0x1000,
    // set [0x2000+a], [0x1000+a]
    Dcpu::Instruct(Dcpu::kSet,
        Dcpu::kLocationOffsetByRegisterA, Dcpu::kLocationOffsetByRegisterA),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterJ, Dcpu::k10),
    // set [0x1000+j], 13
    Dcpu::Instruct(
        Dcpu::kSet, Dcpu::kLocationOffsetByRegisterJ, Dcpu::k13),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, pop
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::kPop)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set peek, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPeek, Dcpu::k14)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k12),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set [SP+0x1], 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPick, Dcpu::k14),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kStackPointer, Dcpu::k13)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kProgramCounter, Dcpu::k13)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kExtra, Dcpu::k13)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocation, Dcpu::k13),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLiteral, Dcpu::k13),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k13),
    // add a, 0x0E
    Dcpu::Instruct(Dcpu::kAdd, Dcpu::kRegisterA, Dcpu::k14)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xFFFF,
    // add a, 0xFFFF
    Dcpu::Instruct(Dcpu::kAdd, Dcpu::kRegisterA, Dcpu::kLiteral),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k30),
    // sub a, 0x10
    Dcpu::Instruct(Dcpu::kSubtract, Dcpu::kRegisterA, Dcpu::k16)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k16),
    // sub a, 0x1E
    Dcpu::Instruct(Dcpu::kSubtract, Dcpu::kRegisterA, Dcpu::k30)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k16),
    // mul a, 0x1E
    Dcpu::Instruct(Dcpu::kMultiply, Dcpu::kRegisterA, Dcpu::k30)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xFFFF,
    // mul a, 0xFFFF
    Dcpu::Instruct(Dcpu::kMultiply, Dcpu::kRegisterA, Dcpu::kLiteral),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k30),
    // div a, 0x10
    Dcpu::Instruct(Dcpu::kDivide, Dcpu::kRegisterA, Dcpu::k16)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k30),
    // div a, 0x00
    Dcpu::Instruct(Dcpu::kDivide, Dcpu::kRegisterA, Dcpu::k0)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k30),
    // mod a, 0x0B
    Dcpu::Instruct(Dcpu::kModulo, Dcpu::kRegisterA, Dcpu::k11)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k30),
    // shl a, 0x02
    Dcpu::Instruct(Dcpu::kShiftLeft, Dcpu::kRegisterA, Dcpu::k2)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xFFFF,
    // shl a, 0x02
    Dcpu::Instruct(Dcpu::kShiftLeft, Dcpu::kRegisterA, Dcpu::k2)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xFFF0,
    // shr a, 0x02
    Dcpu::Instruct(Dcpu::kShiftRight, Dcpu::kRegisterA, Dcpu::k2)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xFFFF,
    // shr a, 0x02
    Dcpu::Instruct(Dcpu::kShiftRight, Dcpu::kRegisterA, Dcpu::k2)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xF0F0,
    // and a, 0x00FF
    Dcpu::Instruct(Dcpu::kBinaryAnd, Dcpu::kRegisterA, Dcpu::kLiteral),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xF0F0,
    // bor a, 0x00FF
    Dcpu::Instruct(Dcpu::kBinaryOr, Dcpu::kRegisterA, Dcpu::kLiteral),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xF0F0,
    // xor a, 0x00FF
    Dcpu::Instruct(Dcpu::kBinaryExclusiveOr, Dcpu::kRegisterA, Dcpu::kLiteral),
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest, ExecuteInstruction_if_equal_register_with_unequal_low_literal) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest,
    ExecuteInstruction_if_not_equal_register_with_unequal_low_literal) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest,
    ExecuteInstruction_if_not_equal_register_with_equal_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x0F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k15),
    // ifn a, 0x0F
    Dcpu::Instruct(Dcpu::kIfNotEqual, Dcpu::kRegisterA, Dcpu::k15),
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k30),
    // ifg a, 0x0F
    Dcpu::Instruct(Dcpu::kIfGreaterThan, Dcpu::kRegisterA, Dcpu::k15),
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k15),
    // ifg a, 0x1E
    Dcpu::Instruct(Dcpu::kIfGreaterThan, Dcpu::kRegisterA, Dcpu::k30),
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k30),
    // ifb a, 0x10
    Dcpu::Instruct(Dcpu::kIfBitSet, Dcpu::kRegisterA, Dcpu::k16),
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14)
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k15),
    // ifb a, 0x10
    Dcpu::Instruct(Dcpu::kIfBitSet, Dcpu::kRegisterA, Dcpu::k16),
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14)
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, dcpu.stack_pointer());
  dcpu.ExecuteInstructions(4);
  EXPECT_EQ(13, dcpu.register_a());
  EXPECT_EQ(14, dcpu.register_b());
  EXPECT_EQ(2, dcpu.program_counter());
  EXPECT_EQ(0, dcpu.stack_pointer());
}
