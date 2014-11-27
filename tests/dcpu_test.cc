#include <algorithm>
#include <gtest/gtest.h>

#include "dcpu.hpp"

using namespace dcpu;

TEST(DcpuTest, DefaultConstructor) {
  Dcpu dcpu;
  EXPECT_EQ(0, *dcpu.address(0x1000));
  EXPECT_EQ(0, dcpu.register_a);
  EXPECT_EQ(0, dcpu.register_b);
  EXPECT_EQ(0, dcpu.register_c);
  EXPECT_EQ(0, dcpu.register_x);
  EXPECT_EQ(0, dcpu.register_y);
  EXPECT_EQ(0, dcpu.register_z);
  EXPECT_EQ(0, dcpu.register_i);
  EXPECT_EQ(0, dcpu.register_j);
  EXPECT_EQ(0, dcpu.program_counter);
  EXPECT_EQ(0, dcpu.stack_pointer);
  EXPECT_EQ(0, dcpu.extra);
  EXPECT_EQ(0, dcpu.interrupt_address);
}

TEST(DcpuTest, Reset) {
  Dcpu dcpu;
  *dcpu.address(0x1000) = 1;
  dcpu.register_a = 2;
  dcpu.register_b = 3;
  dcpu.register_c = 4;
  dcpu.register_x = 5;
  dcpu.register_y = 6;
  dcpu.register_z = 7;
  dcpu.register_i = 8;
  dcpu.register_j = 9;
  dcpu.program_counter = 10;
  dcpu.stack_pointer = 11;
  dcpu.extra = 12;
  dcpu.interrupt_address = 13;
  EXPECT_EQ(1, *dcpu.address(0x1000));
  EXPECT_EQ(2, dcpu.register_a);
  EXPECT_EQ(3, dcpu.register_b);
  EXPECT_EQ(4, dcpu.register_c);
  EXPECT_EQ(5, dcpu.register_x);
  EXPECT_EQ(6, dcpu.register_y);
  EXPECT_EQ(7, dcpu.register_z);
  EXPECT_EQ(8, dcpu.register_i);
  EXPECT_EQ(9, dcpu.register_j);
  EXPECT_EQ(10, dcpu.program_counter);
  EXPECT_EQ(11, dcpu.stack_pointer);
  EXPECT_EQ(12, dcpu.extra);
  EXPECT_EQ(13, dcpu.interrupt_address);
  dcpu.Reset();
  EXPECT_EQ(0, *dcpu.address(0x1000));
  EXPECT_EQ(0, dcpu.register_a);
  EXPECT_EQ(0, dcpu.register_b);
  EXPECT_EQ(0, dcpu.register_c);
  EXPECT_EQ(0, dcpu.register_x);
  EXPECT_EQ(0, dcpu.register_y);
  EXPECT_EQ(0, dcpu.register_z);
  EXPECT_EQ(0, dcpu.register_i);
  EXPECT_EQ(0, dcpu.register_j);
  EXPECT_EQ(0, dcpu.program_counter);
  EXPECT_EQ(0, dcpu.stack_pointer);
  EXPECT_EQ(0, dcpu.extra);
  EXPECT_EQ(0, dcpu.interrupt_address);
}

TEST(DcpuTest, ExecuteInstructions) {
  Dcpu dcpu;
  // dcpu.program_counter == 0 by DcpuTest.DefaultConstructor
  dcpu.ExecuteInstructions(10);
  EXPECT_EQ(10, dcpu.program_counter);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_register) {
  Dcpu dcpu;
  const Word program[] = {
    // set b, 1
    Instruct(BasicOpcode::kSet, Operand::kRegisterB, Operand::k1),
    // set a, b
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kRegisterB)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  // dcpu.register_a == 0 by DcpuTest.DefaultConstructor
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(1, dcpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_last_register) {
  Dcpu dcpu;
  const Word program[] = {
    // set j, 1
    Instruct(BasicOpcode::kSet, Operand::kRegisterJ, Operand::k1),
    // set a, j
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kRegisterJ)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  // dcpu.register_a == 0 by DcpuTest.DefaultConstructor
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(1, dcpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_location_in_register) {
  Dcpu dcpu;
  const Word program[] = {
    // set [0x1000], 13
    Instruct(BasicOpcode::kSet, Operand::kLocation, Operand::k13), 0x1000,
    // set b, 0x1000
    Instruct(BasicOpcode::kSet, Operand::kRegisterB, Operand::kLiteral), 0x1000,
    // set a, [b]
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kLocationInRegisterB)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, dcpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_location_in_last_register) {
  Dcpu dcpu;
  const Word program[] = {
    // set [0x1000], 13
    Instruct(BasicOpcode::kSet, Operand::kLocation, Operand::k13), 0x1000,
    // set j, 0x1000
    Instruct(BasicOpcode::kSet, Operand::kRegisterJ, Operand::kLiteral), 0x1000,
    // set a, [j]
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kLocationInRegisterJ)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, dcpu.register_a);
}

TEST(DcpuTest,
    ExecuteInstruction_set_register_with_location_offset_by_register) {
  Dcpu dcpu;
  const Word program[] = {
    // set [0x100A], 13
    Instruct(BasicOpcode::kSet, Operand::kLocation, Operand::k13), 0x100A,
    // set b, 10
    Instruct(BasicOpcode::kSet, Operand::kRegisterB, Operand::k10),
    // set a, [0x1000+b]
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kLocationOffsetByRegisterB), 0x1000
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, dcpu.register_a);
}

TEST(DcpuTest,
    ExecuteInstruction_set_register_with_location_offset_by_last_register) {
  Dcpu dcpu;
  const Word program[] = {
    // set [0x100A], 13
    Instruct(BasicOpcode::kSet, Operand::kLocation, Operand::k13), 0x100A,
    // set j, 10
    Instruct(BasicOpcode::kSet, Operand::kRegisterJ, Operand::k10),
    // set a, [0x1000+j]
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kLocationOffsetByRegisterJ), 0x1000
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, dcpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_pop) {
  Dcpu dcpu;
  const Word program[] = {
    // set push, 13
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k13),
    // set a, pop
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kPop)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, dcpu.stack_pointer);
  dcpu.ExecuteInstruction();
  EXPECT_EQ(0, dcpu.stack_pointer);
  EXPECT_EQ(13, dcpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_peek) {
  Dcpu dcpu;
  const Word program[] = {
    // set push, 13
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k13),
    // set a, peek
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kPeek)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, dcpu.stack_pointer);
  dcpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, dcpu.stack_pointer);
  EXPECT_EQ(13, dcpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_pick) {
  Dcpu dcpu;
  const Word program[] = {
    // set push, 13
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k13),
    // set push, 14
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k14),
    // set a, [sp+1]
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kPick), 0x1
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  EXPECT_EQ(0, dcpu.stack_pointer);
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(0xFFFE, dcpu.stack_pointer);
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer + 1));
  EXPECT_EQ(13, dcpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_stack_pointer) {
  Dcpu dcpu;
  const Word program[] = {
    // set push, 13
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k13),
    // set a, sp
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kStackPointer)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, dcpu.stack_pointer);
  dcpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, dcpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_program_counter) {
  Dcpu dcpu;
  const Word program[] = {
    // noop
    Noop(),
    // set a, pc
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kProgramCounter)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(2, dcpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_overflow) {
  Dcpu dcpu;
  const Word program[] = {
    // set ex, 13
    Instruct(BasicOpcode::kSet, Operand::kExtra, Operand::k13),
    // set a, ex
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kExtra)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(13, dcpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_location) {
  Dcpu dcpu;
  const Word program[] = {
    // set [0x1000], 13
    Instruct(BasicOpcode::kSet, Operand::kLocation, Operand::k13), 0x1000,
    // set a, [0x1000]
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kLocation), 0x1000
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(13, dcpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_high_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0x1001
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kLiteral), 0x1001
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(0x1001, dcpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 1
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k1),
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(1, dcpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_last_register_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set j, 1
    Instruct(BasicOpcode::kSet, Operand::kRegisterJ, Operand::k1),
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  // dcpu.register_a == 0 by DcpuTest.DefaultConstructor
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(1, dcpu.register_j);
}

TEST(DcpuTest, ExecuteInstruction_set_location_in_register_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0x1000
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kLiteral), 0x1000,
    // set [a], 13
    Instruct(BasicOpcode::kSet, Operand::kLocationInRegisterA, Operand::k13)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(13, *dcpu.address(0x1000));
}

TEST(DcpuTest,
    ExecuteInstruction_set_location_in_last_register_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set j, 0x1000
    Instruct(BasicOpcode::kSet, Operand::kRegisterJ, Operand::kLiteral), 0x1000,
    // set [j], 13
    Instruct(BasicOpcode::kSet, Operand::kLocationInRegisterJ, Operand::k13)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(13, *dcpu.address(0x1000));
}

TEST(DcpuTest,
    ExecuteInstruction_set_location_offset_by_register_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 10
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k10),
    // set [0x1000+a], 13
    Instruct(BasicOpcode::kSet, Operand::kLocationOffsetByRegisterA, Operand::k13), 0x1000
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(13, *dcpu.address(0x100A));
}

TEST(DcpuTest,
    ExecuteInstruction_set_location_offset_by_register_with_location_offset_by_register) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 10
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k10),
    // set [0x1000+a], 13
    Instruct(BasicOpcode::kSet, Operand::kLocationOffsetByRegisterA, Operand::k13), 0x1000,
    // set [0x2000+a], [0x1000+a]
    Instruct(BasicOpcode::kSet,
        Operand::kLocationOffsetByRegisterA, Operand::kLocationOffsetByRegisterA), 0x1000, 0x2000
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *dcpu.address(0x200A));
}

TEST(DcpuTest,
    ExecuteInstruction_set_location_offset_by_last_register_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set j, 10
    Instruct(BasicOpcode::kSet, Operand::kRegisterJ, Operand::k10),
    // set [0x1000+j], 13
    Instruct(BasicOpcode::kSet, Operand::kLocationOffsetByRegisterJ, Operand::k13), 0x1000
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(13, *dcpu.address(0x100A));
}

TEST(DcpuTest, ExecuteInstruction_set_push_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set push, 13
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k13)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_set_push_with_pop) {
  Dcpu dcpu;
  const Word program[] = {
    // set push, 13
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k13),
    // set push, pop
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::kPop)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0xFFFF, dcpu.stack_pointer);
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_set_peek_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set push, 13
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k13),
    // set peek, 14
    Instruct(BasicOpcode::kSet, Operand::kPeek, Operand::k14)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_set_pick_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set push, 13
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k12),
    // set push, 14
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k13),
    // set [SP+0x1], 14
    Instruct(BasicOpcode::kSet, Operand::kPick, Operand::k14), 0x1
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer));
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer + 1));
}

TEST(DcpuTest, ExecuteInstruction_set_stack_pointer_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set sp, 13
    Instruct(BasicOpcode::kSet, Operand::kStackPointer, Operand::k13)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(13, dcpu.stack_pointer);
}

TEST(DcpuTest, ExecuteInstruction_set_program_counter_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set pc, 13
    Instruct(BasicOpcode::kSet, Operand::kProgramCounter, Operand::k13)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(13, dcpu.program_counter);
}

TEST(DcpuTest, ExecuteInstruction_set_overflow_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set ex, 13
    Instruct(BasicOpcode::kSet, Operand::kExtra, Operand::k13)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(13, dcpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_set_location_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set [0x1000], 13
    Instruct(BasicOpcode::kSet, Operand::kLocation, Operand::k13), 0x1000
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(13, *dcpu.address(0x1000));
}

TEST(DcpuTest, ExecuteInstruction_set_literal_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set 0x1000, 13
    Instruct(BasicOpcode::kSet, Operand::kLiteral, Operand::k13), 0x1000
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  // set 0x1000, 13 should be a noop.
  EXPECT_EQ(0, *dcpu.address(0x1000));
}

TEST(DcpuTest, ExecuteInstruction_add_register_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0x0D
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k13),
    // add a, 0x0E
    Instruct(BasicOpcode::kAdd, Operand::kRegisterA, Operand::k14)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0x1B, dcpu.register_a);
  EXPECT_EQ(0, dcpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_add_register_with_overflow) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0xFFFF
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kLiteral), 0xFFFF,
    // add a, 0xFFFF
    Instruct(BasicOpcode::kAdd, Operand::kRegisterA, Operand::kLiteral), 0xFFFF
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0xFFFE, dcpu.register_a);
  EXPECT_EQ(1, dcpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_subtract_register_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0x1E
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k30),
    // sub a, 0x10
    Instruct(BasicOpcode::kSubtract, Operand::kRegisterA, Operand::k16)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0xE, dcpu.register_a);
  EXPECT_EQ(0, dcpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_subtract_register_with_underflow) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0x10
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k16),
    // sub a, 0x1E
    Instruct(BasicOpcode::kSubtract, Operand::kRegisterA, Operand::k30)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0xFFF2, dcpu.register_a);
  EXPECT_EQ(1, dcpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_multiply_register_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0x10
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k16),
    // mul a, 0x1E
    Instruct(BasicOpcode::kMultiply, Operand::kRegisterA, Operand::k30)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0x01E0, dcpu.register_a);
  EXPECT_EQ(0, dcpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_multiply_register_with_overflow) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0xFFFF
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kLiteral), 0xFFFF,
    // mul a, 0xFFFF
    Instruct(BasicOpcode::kMultiply, Operand::kRegisterA, Operand::kLiteral), 0xFFFF
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0x0001, dcpu.register_a);
  EXPECT_EQ(0xFFFE, dcpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_divide_register_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0x1E
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k30),
    // div a, 0x10
    Instruct(BasicOpcode::kDivide, Operand::kRegisterA, Operand::k16)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(1, dcpu.register_a);
  EXPECT_EQ(0, dcpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_divide_register_by_zero) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0x1E
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k30),
    // div a, 0x00
    Instruct(BasicOpcode::kDivide, Operand::kRegisterA, Operand::k0)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0, dcpu.register_a);
  EXPECT_EQ(1, dcpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_modulo_register_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0x1E
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k30),
    // mod a, 0x0B
    Instruct(BasicOpcode::kModulo, Operand::kRegisterA, Operand::k11)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0x8, dcpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_shift_left_register_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0x1E
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k30),
    // shl a, 0x02
    Instruct(BasicOpcode::kShiftLeft, Operand::kRegisterA, Operand::k2)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0x78, dcpu.register_a);
  EXPECT_EQ(0, dcpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_shift_left_register_with_overflow) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0xFFFF
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kLiteral), 0xFFFF,
    // shl a, 0x02
    Instruct(BasicOpcode::kShiftLeft, Operand::kRegisterA, Operand::k2)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0xFFFC, dcpu.register_a);
  EXPECT_EQ(0x0003, dcpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_shift_right_register_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0xFFF0
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kLiteral), 0xFFF0,
    // shr a, 0x02
    Instruct(BasicOpcode::kShiftRight, Operand::kRegisterA, Operand::k2)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0x3FFC, dcpu.register_a);
  EXPECT_EQ(0, dcpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_shift_right_register_with_underflow) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0xFFFF
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kLiteral), 0xFFFF,
    // shr a, 0x02
    Instruct(BasicOpcode::kShiftRight, Operand::kRegisterA, Operand::k2)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0x3FFF, dcpu.register_a);
  EXPECT_EQ(0xC000, dcpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_and_register_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0xF0F0
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kLiteral), 0xF0F0,
    // and a, 0x00FF
    Instruct(BasicOpcode::kBinaryAnd, Operand::kRegisterA, Operand::kLiteral), 0x00FF
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0x00F0, dcpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_or_register_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0xF0F0
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kLiteral), 0xF0F0,
    // bor a, 0x00FF
    Instruct(BasicOpcode::kBinaryOr, Operand::kRegisterA, Operand::kLiteral), 0x00FF
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0xF0FF, dcpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_xor_register_with_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0xF0F0
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::kLiteral), 0xF0F0,
    // xor a, 0x00FF
    Instruct(BasicOpcode::kBinaryExclusiveOr, Operand::kRegisterA, Operand::kLiteral), 0x00FF
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(2);
  EXPECT_EQ(0xF00F, dcpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_if_equal_register_with_equal_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0x0F
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k15),
    // ife a, 0x0F
    Instruct(BasicOpcode::kIfEqual, Operand::kRegisterA, Operand::k15),
    // set push, 13
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k13),
    // set push, 14
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k14)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_if_equal_register_with_unequal_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0x0F
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k15),
    // ife a, 0x00
    Instruct(BasicOpcode::kIfEqual, Operand::kRegisterA, Operand::k0),
    // set push, 13
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k13),
    // set push, 14
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k14)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_if_not_equal_register_with_unequal_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0x0F
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k15),
    // ifn a, 0x00
    Instruct(BasicOpcode::kIfNotEqual, Operand::kRegisterA, Operand::k0),
    // set push, 13
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k13),
    // set push, 14
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k14)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_if_not_equal_register_with_equal_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0x0F
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k15),
    // ifn a, 0x0F
    Instruct(BasicOpcode::kIfNotEqual, Operand::kRegisterA, Operand::k15),
    // set push, 13
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k13),
    // set push, 14
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k14)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_if_greater_than_register_with_lesser_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0x1E
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k30),
    // ifg a, 0x0F
    Instruct(BasicOpcode::kIfGreaterThan, Operand::kRegisterA, Operand::k15),
    // set push, 13
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k13),
    // set push, 14
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k14)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_if_greater_than_register_with_greater_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0x0F
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k15),
    // ifg a, 0x1E
    Instruct(BasicOpcode::kIfGreaterThan, Operand::kRegisterA, Operand::k30),
    // set push, 13
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k13),
    // set push, 14
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k14)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_if_both_register_with_common_bits_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0x1E
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k30),
    // ifb a, 0x10
    Instruct(BasicOpcode::kIfBitSet, Operand::kRegisterA, Operand::k16),
    // set push, 13
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k13),
    // set push, 14
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k14)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *dcpu.address(dcpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_if_both_register_with_uncommon_bits_low_literal) {
  Dcpu dcpu;
  const Word program[] = {
    // set a, 0x0F
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k15),
    // ifb a, 0x10
    Instruct(BasicOpcode::kIfBitSet, Operand::kRegisterA, Operand::k16),
    // set push, 13
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k13),
    // set push, 14
    Instruct(BasicOpcode::kSet, Operand::kPush, Operand::k14)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstructions(3);
  EXPECT_EQ(14, *dcpu.address(dcpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_jump_sub_routine) {
  Dcpu dcpu;
  const Word program[] = {
    // jsr subroutine
    Instruct(AdvancedOpcode::kJumpSubRoutine, Operand::k3),
    // set a, 13
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k13),
    // sub pc, 1
    Instruct(BasicOpcode::kSubtract, Operand::kProgramCounter, Operand::k1),
    // subroutine: set b, 14
    Instruct(BasicOpcode::kSet, Operand::kRegisterB, Operand::k14),
    // set pc, pop
    Instruct(BasicOpcode::kSet, Operand::kProgramCounter, Operand::kPop)
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, dcpu.memory_begin());
  dcpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, dcpu.stack_pointer);
  dcpu.ExecuteInstructions(4);
  EXPECT_EQ(13, dcpu.register_a);
  EXPECT_EQ(14, dcpu.register_b);
  EXPECT_EQ(2, dcpu.program_counter);
  EXPECT_EQ(0, dcpu.stack_pointer);
}
