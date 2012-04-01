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
  EXPECT_EQ(0, dcpu.overflow());
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
  dcpu.overflow() = 12;
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
  EXPECT_EQ(12, dcpu.overflow());
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
  EXPECT_EQ(0, dcpu.overflow());
}

TEST(DcpuTest, ExecuteCycles) {
  Dcpu dcpu;
  // dcpu.program_counter() == 0 by DcpuTest.DefaultConstructor
  dcpu.ExecuteCycles(10);
  EXPECT_EQ(10, dcpu.program_counter());
}

TEST(DcpuTest, ExecuteCycle_set_register_with_register) {
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
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(1, dcpu.register_a());
}

TEST(DcpuTest, ExecuteCycle_set_register_with_location_in_register) {
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
  dcpu.ExecuteCycles(3);
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest, ExecuteCycle_set_register_with_location_offset_by_register) {
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
  dcpu.ExecuteCycles(3);
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest, ExecuteCycle_set_register_with_pop) {
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
  dcpu.ExecuteCycle();
  EXPECT_EQ(0xFFFF, dcpu.stack_pointer());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0, dcpu.stack_pointer());
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest, ExecuteCycle_set_register_with_peek) {
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
  dcpu.ExecuteCycle();
  EXPECT_EQ(0xFFFF, dcpu.stack_pointer());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0xFFFF, dcpu.stack_pointer());
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest, ExecuteCycle_set_register_with_push) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k13),
    // set a, push
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kPush)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0, dcpu.stack_pointer());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0xFFFF, dcpu.stack_pointer());
  EXPECT_EQ(0, dcpu.register_a());
}

TEST(DcpuTest, ExecuteCycle_set_register_with_stack_pointer) {
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
  dcpu.ExecuteCycle();
  EXPECT_EQ(0xFFFF, dcpu.stack_pointer());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0xFFFF, dcpu.register_a());
}

TEST(DcpuTest, ExecuteCycle_set_register_with_program_counter) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // reserved 0, 0
    Dcpu::Instruct(Dcpu::kReserved, Dcpu::k0, Dcpu::k0),
    // set a, pc
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kProgramCounter)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(2, dcpu.register_a());
}

TEST(DcpuTest, ExecuteCycle_set_register_with_overflow) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set o, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kOverflow, Dcpu::k13),
    // set a, o
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kOverflow)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest, ExecuteCycle_set_register_with_location) {
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
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(13, dcpu.register_a());
}

TEST(DcpuTest, ExecuteCycle_set_register_with_high_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x1001
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0x1001
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0x1001, dcpu.register_a());
}

TEST(DcpuTest, ExecuteCycle_set_register_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 1
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k1),
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycle();
  EXPECT_EQ(1, dcpu.register_a());
}

TEST(DcpuTest, ExecuteCycle_set_location_in_register_with_low_literal) {
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
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(13, *dcpu.address(0x1000));
}

TEST(DcpuTest, ExecuteCycle_set_location_offset_by_register_with_low_literal) {
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
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(13, *dcpu.address(0x100A));
}

TEST(DcpuTest, ExecuteCycle_set_pop_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set pop, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPop, Dcpu::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer() - 1));
}

TEST(DcpuTest, ExecuteCycle_set_peek_with_low_literal) {
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
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest, ExecuteCycle_set_push_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycle();
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest, ExecuteCycle_set_stack_pointer_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set sp, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kStackPointer, Dcpu::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycle();
  EXPECT_EQ(13, dcpu.stack_pointer());
}

TEST(DcpuTest, ExecuteCycle_set_program_counter_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set pc, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kProgramCounter, Dcpu::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycle();
  EXPECT_EQ(13, dcpu.program_counter());
}

TEST(DcpuTest, ExecuteCycle_set_overflow_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set o, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kOverflow, Dcpu::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycle();
  EXPECT_EQ(13, dcpu.overflow());
}

TEST(DcpuTest, ExecuteCycle_set_location_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set [0x1000], 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocation, Dcpu::k13),
    0x1000
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycle();
  EXPECT_EQ(13, *dcpu.address(0x1000));
}

TEST(DcpuTest, ExecuteCycle_set_literal_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set 0x1000, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLiteral, Dcpu::k13),
    0x1000
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycle();
  // set 0x1000, 13 should be a noop.
  EXPECT_EQ(0, *dcpu.address(0x1000));
}

TEST(DcpuTest, ExecuteCycle_set_low_literal_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set 10, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::k10, Dcpu::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycle();
  // set 10, 13 should be a noop.
  EXPECT_EQ(0, *dcpu.address(0xA));
}

TEST(DcpuTest, ExecuteCycle_add_register_with_low_literal) {
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
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(0x1B, dcpu.register_a());
  EXPECT_EQ(0, dcpu.overflow());
}

TEST(DcpuTest, ExecuteCycle_add_register_with_overflow) {
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
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(0xFFFE, dcpu.register_a());
  EXPECT_EQ(1, dcpu.overflow());
}

TEST(DcpuTest, ExecuteCycle_subtract_register_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x1F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k31),
    // sub a, 0x10
    Dcpu::Instruct(Dcpu::kSubtract, Dcpu::kRegisterA, Dcpu::k16)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(0xF, dcpu.register_a());
  EXPECT_EQ(0, dcpu.overflow());
}

TEST(DcpuTest, ExecuteCycle_subtract_register_with_underflow) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x10
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k16),
    // sub a, 0x1F
    Dcpu::Instruct(Dcpu::kSubtract, Dcpu::kRegisterA, Dcpu::k31)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(0xFFF1, dcpu.register_a());
  EXPECT_EQ(1, dcpu.overflow());
}

TEST(DcpuTest, ExecuteCycle_multiply_register_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x10
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k16),
    // mul a, 0x1F
    Dcpu::Instruct(Dcpu::kMultiply, Dcpu::kRegisterA, Dcpu::k31)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(0x01F0, dcpu.register_a());
  EXPECT_EQ(0, dcpu.overflow());
}

TEST(DcpuTest, ExecuteCycle_multiply_register_with_overflow) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0xFFFF
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xFFFF,
    // add a, 0xFFFF
    Dcpu::Instruct(Dcpu::kMultiply, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xFFFF
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(0x0001, dcpu.register_a());
  EXPECT_EQ(0xFFFE, dcpu.overflow());
}

TEST(DcpuTest, ExecuteCycle_divide_register_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x1F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k31),
    // div a, 0x10
    Dcpu::Instruct(Dcpu::kDivide, Dcpu::kRegisterA, Dcpu::k16)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(1, dcpu.register_a());
  EXPECT_EQ(0, dcpu.overflow());
}

TEST(DcpuTest, ExecuteCycle_divide_register_by_zero) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x1F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k31),
    // div a, 0x00
    Dcpu::Instruct(Dcpu::kDivide, Dcpu::kRegisterA, Dcpu::k0)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(0, dcpu.register_a());
  EXPECT_EQ(1, dcpu.overflow());
}

TEST(DcpuTest, ExecuteCycle_modulo_register_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x1F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k31),
    // mod a, 0x0B
    Dcpu::Instruct(Dcpu::kModulo, Dcpu::kRegisterA, Dcpu::k11)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(0x9, dcpu.register_a());
}

TEST(DcpuTest, ExecuteCycle_shift_left_register_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x1F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k31),
    // shl a, 0x02
    Dcpu::Instruct(Dcpu::kShiftLeft, Dcpu::kRegisterA, Dcpu::k2)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(0x7C, dcpu.register_a());
  EXPECT_EQ(0, dcpu.overflow());
}

TEST(DcpuTest, ExecuteCycle_shift_left_register_with_overflow) {
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
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(0xFFFC, dcpu.register_a());
  EXPECT_EQ(0x0003, dcpu.overflow());
}

TEST(DcpuTest, ExecuteCycle_shift_right_register_with_low_literal) {
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
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(0x3FFC, dcpu.register_a());
  EXPECT_EQ(0, dcpu.overflow());
}

TEST(DcpuTest, ExecuteCycle_shift_right_register_with_underflow) {
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
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(0x3FFF, dcpu.register_a());
  EXPECT_EQ(0xC000, dcpu.overflow());
}

TEST(DcpuTest, ExecuteCycle_and_register_with_low_literal) {
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
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(0x00F0, dcpu.register_a());
}

TEST(DcpuTest, ExecuteCycle_or_register_with_low_literal) {
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
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(0xF0FF, dcpu.register_a());
}

TEST(DcpuTest, ExecuteCycle_xor_register_with_low_literal) {
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
  dcpu.ExecuteCycles(2);
  EXPECT_EQ(0xF00F, dcpu.register_a());
}

TEST(DcpuTest, ExecuteCycle_if_equal_register_with_equal_low_literal) {
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
  dcpu.ExecuteCycles(3);
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest, ExecuteCycle_if_equal_register_with_unequal_low_literal) {
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
  dcpu.ExecuteCycles(3);
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest, ExecuteCycle_if_not_equal_register_with_unequal_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x0F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k15),
    // ife a, 0x00
    Dcpu::Instruct(Dcpu::kIfNotEqual, Dcpu::kRegisterA, Dcpu::k0),
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycles(3);
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest, ExecuteCycle_if_not_equal_register_with_equal_low_literal) {
  Dcpu dcpu;
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
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycles(3);
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest, ExecuteCycle_if_greater_than_register_with_lesser_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x1F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k31),
    // ife a, 0x0F
    Dcpu::Instruct(Dcpu::kIfGreaterThan, Dcpu::kRegisterA, Dcpu::k15),
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycles(3);
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest, ExecuteCycle_if_greater_than_register_with_greater_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x0F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k15),
    // ife a, 0x1F
    Dcpu::Instruct(Dcpu::kIfGreaterThan, Dcpu::kRegisterA, Dcpu::k31),
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycles(3);
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest, ExecuteCycle_if_both_register_with_common_bits_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x1F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k31),
    // ife a, 0x0F
    Dcpu::Instruct(Dcpu::kIfGreaterThan, Dcpu::kRegisterA, Dcpu::k15),
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycles(3);
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer()));
}

TEST(DcpuTest, ExecuteCycle_if_both_register_with_uncommon_bits_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    // set a, 0x0F
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k15),
    // ife a, 0x10
    Dcpu::Instruct(Dcpu::kIfGreaterThan, Dcpu::kRegisterA, Dcpu::k16),
    // set push, 13
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
    // set push, 14
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k14)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycles(3);
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer()));
}
