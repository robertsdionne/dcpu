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

TEST(DcpuTest, ExecuteCycle_set_register_with_low_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::k1),
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterB, Dcpu::k2),
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterC, Dcpu::k3),
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterX, Dcpu::k4),
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterY, Dcpu::k5),
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterZ, Dcpu::k6),
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterI, Dcpu::k7),
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterJ, Dcpu::k8),
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kStackPointer, Dcpu::k9),
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kOverflow, Dcpu::k10),
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kProgramCounter, Dcpu::k31)
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycle();
  EXPECT_EQ(1, dcpu.register_a());
  dcpu.ExecuteCycle();
  EXPECT_EQ(2, dcpu.register_b());
  dcpu.ExecuteCycle();
  EXPECT_EQ(3, dcpu.register_c());
  dcpu.ExecuteCycle();
  EXPECT_EQ(4, dcpu.register_x());
  dcpu.ExecuteCycle();
  EXPECT_EQ(5, dcpu.register_y());
  dcpu.ExecuteCycle();
  EXPECT_EQ(6, dcpu.register_z());
  dcpu.ExecuteCycle();
  EXPECT_EQ(7, dcpu.register_i());
  dcpu.ExecuteCycle();
  EXPECT_EQ(8, dcpu.register_j());
  dcpu.ExecuteCycle();
  EXPECT_EQ(9, dcpu.stack_pointer());
  dcpu.ExecuteCycle();
  EXPECT_EQ(10, dcpu.overflow());
  EXPECT_EQ(10, dcpu.program_counter());
  dcpu.ExecuteCycle();
  EXPECT_EQ(31, dcpu.program_counter());
}

TEST(DcpuTest, ExecuteCycle_set_register_with_high_literal) {
  Dcpu dcpu;
  const Dcpu::Word program[] = {
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0x1001,
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterB, Dcpu::kLiteral),
    0x1002,
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterC, Dcpu::kLiteral),
    0x1003,
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterX, Dcpu::kLiteral),
    0x1004,
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterY, Dcpu::kLiteral),
    0x1005,
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterZ, Dcpu::kLiteral),
    0x1006,
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterI, Dcpu::kLiteral),
    0x1007,
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterJ, Dcpu::kLiteral),
    0x1008,
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kStackPointer, Dcpu::kLiteral),
    0x1009,
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kOverflow, Dcpu::kLiteral),
    0x100A,
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kProgramCounter, Dcpu::kLiteral),
    0x100B
  };
  const Dcpu::Word *const program_end =
      program + sizeof(program)/sizeof(Dcpu::Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0x1001, dcpu.register_a());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0x1002, dcpu.register_b());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0x1003, dcpu.register_c());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0x1004, dcpu.register_x());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0x1005, dcpu.register_y());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0x1006, dcpu.register_z());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0x1007, dcpu.register_i());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0x1008, dcpu.register_j());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0x1009, dcpu.stack_pointer());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0x100A, dcpu.overflow());
  EXPECT_EQ(20, dcpu.program_counter());
  dcpu.ExecuteCycle();
  EXPECT_EQ(0x100B, dcpu.program_counter());
}
