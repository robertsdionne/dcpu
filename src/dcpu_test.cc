// Copyright 2012 Robert Scott Dionne. All rights reserved.

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
