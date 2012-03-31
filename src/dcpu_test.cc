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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterB, Dcpu::k1),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocation, Dcpu::k13),
    0x1000,
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterB, Dcpu::kLiteral),
    0x1000,
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocation, Dcpu::k13),
    0x100A,
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterB, Dcpu::k10),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k13),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
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
    Dcpu::Instruct(Dcpu::kReserved, Dcpu::k0, Dcpu::k0),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kOverflow, Dcpu::k13),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocation, Dcpu::k13),
    0x1000,
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0x1000,
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k10),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kPush, Dcpu::k13),
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
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLiteral, Dcpu::k13),
    0x1000
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0, *dcpu.address(0x1000));
}

TEST(DcpuTest, ExecuteCycle_set_low_literal_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    Dcpu::Instruct(Dcpu::kSet, Dcpu::k10, Dcpu::k13)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0, *dcpu.address(0xA));
}
