#include <algorithm>
#include <gtest/gtest.h>

#include "dcpu.hpp"
#include "dsl.hpp"

using namespace dcpu;
using namespace dcpu::dsl;

TEST(DcpuTest, DefaultConstructor) {
  Dcpu cpu;
  EXPECT_EQ(0, *cpu.address(0x1000));
  EXPECT_EQ(0, cpu.register_a);
  EXPECT_EQ(0, cpu.register_b);
  EXPECT_EQ(0, cpu.register_c);
  EXPECT_EQ(0, cpu.register_x);
  EXPECT_EQ(0, cpu.register_y);
  EXPECT_EQ(0, cpu.register_z);
  EXPECT_EQ(0, cpu.register_i);
  EXPECT_EQ(0, cpu.register_j);
  EXPECT_EQ(0, cpu.program_counter);
  EXPECT_EQ(0, cpu.stack_pointer);
  EXPECT_EQ(0, cpu.extra);
  EXPECT_EQ(0, cpu.interrupt_address);
}

TEST(DcpuTest, Reset) {
  Dcpu cpu;
  *cpu.address(0x1000) = 1;
  cpu.register_a = 2;
  cpu.register_b = 3;
  cpu.register_c = 4;
  cpu.register_x = 5;
  cpu.register_y = 6;
  cpu.register_z = 7;
  cpu.register_i = 8;
  cpu.register_j = 9;
  cpu.program_counter = 10;
  cpu.stack_pointer = 11;
  cpu.extra = 12;
  cpu.interrupt_address = 13;
  EXPECT_EQ(1, *cpu.address(0x1000));
  EXPECT_EQ(2, cpu.register_a);
  EXPECT_EQ(3, cpu.register_b);
  EXPECT_EQ(4, cpu.register_c);
  EXPECT_EQ(5, cpu.register_x);
  EXPECT_EQ(6, cpu.register_y);
  EXPECT_EQ(7, cpu.register_z);
  EXPECT_EQ(8, cpu.register_i);
  EXPECT_EQ(9, cpu.register_j);
  EXPECT_EQ(10, cpu.program_counter);
  EXPECT_EQ(11, cpu.stack_pointer);
  EXPECT_EQ(12, cpu.extra);
  EXPECT_EQ(13, cpu.interrupt_address);
  cpu.Reset();
  EXPECT_EQ(0, *cpu.address(0x1000));
  EXPECT_EQ(0, cpu.register_a);
  EXPECT_EQ(0, cpu.register_b);
  EXPECT_EQ(0, cpu.register_c);
  EXPECT_EQ(0, cpu.register_x);
  EXPECT_EQ(0, cpu.register_y);
  EXPECT_EQ(0, cpu.register_z);
  EXPECT_EQ(0, cpu.register_i);
  EXPECT_EQ(0, cpu.register_j);
  EXPECT_EQ(0, cpu.program_counter);
  EXPECT_EQ(0, cpu.stack_pointer);
  EXPECT_EQ(0, cpu.extra);
  EXPECT_EQ(0, cpu.interrupt_address);
}

TEST(DcpuTest, ExecuteInstructions) {
  Dcpu cpu;
  // cpu.program_counter == 0 by DcpuTest.DefaultConstructor
  cpu.ExecuteInstructions(10);
  EXPECT_EQ(10, cpu.program_counter);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_register) {
  Dcpu cpu;
  const Word program[] = {
    // set b, 1
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterB, dcpu::Operand::k1),
    // set a, b
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kRegisterB)
  };

  Dsl d;
  d.set(b(), 1)
    .set(a(), b())
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  // cpu.register_a == 0 by DcpuTest.DefaultConstructor
  cpu.ExecuteInstructions(2);
  EXPECT_EQ(1, cpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_last_register) {
  Dcpu cpu;
  const Word program[] = {
    // set j, 1
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterJ, dcpu::Operand::k1),
    // set a, j
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kRegisterJ)
  };

  Dsl d;
  d.set(j(), 1)
    .set(a(), j())
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  // cpu.register_a == 0 by DcpuTest.DefaultConstructor
  cpu.ExecuteInstructions(2);
  EXPECT_EQ(1, cpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_location_in_register) {
  Dcpu cpu;
  const Word program[] = {
    // set [0x1000], 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kLocation, dcpu::Operand::k13), 0x1000,
    // set b, 0x1000
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterB, dcpu::Operand::kLiteral), 0x1000,
    // set a, [b]
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kLocationInRegisterB)
  };

  Dsl d;
  d.set(d[0x1000], 13)
    .set(b(), 0x1000)
    .set(a(), d[b()])
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(3);
  EXPECT_EQ(13, cpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_location_in_last_register) {
  Dcpu cpu;
  const Word program[] = {
    // set [0x1000], 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kLocation, dcpu::Operand::k13), 0x1000,
    // set j, 0x1000
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterJ, dcpu::Operand::kLiteral), 0x1000,
    // set a, [j]
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kLocationInRegisterJ)
  };

  Dsl d;
  d.set(d[0x1000], 13)
    .set(j(), 0x1000)
    .set(a(), d[j()])
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(3);
  EXPECT_EQ(13, cpu.register_a);
}

TEST(DcpuTest,
    ExecuteInstruction_set_register_with_location_offset_by_register) {
  Dcpu cpu;
  const Word program[] = {
    // set [0x100A], 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kLocation, dcpu::Operand::k13), 0x100A,
    // set b, 10
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterB, dcpu::Operand::k10),
    // set a, [0x1000+b]
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kLocationOffsetByRegisterB), 0x1000
  };

  Dsl d;
  d.set(d[0x100A], 13)
    .set(b(), 0xA)
    .set(a(), d[0x1000+ b()])
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(3);
  EXPECT_EQ(13, cpu.register_a);
}

TEST(DcpuTest,
    ExecuteInstruction_set_register_with_location_offset_by_last_register) {
  Dcpu cpu;
  const Word program[] = {
    // set [0x100A], 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kLocation, dcpu::Operand::k13), 0x100A,
    // set j, 10
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterJ, dcpu::Operand::k10),
    // set a, [0x1000+j]
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kLocationOffsetByRegisterJ), 0x1000
  };

  Dsl d;
  d.set(d[0x100A], 13)
    .set(j(), 0xA)
    .set(a(), d[0x1000 + j()])
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(3);
  EXPECT_EQ(13, cpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_pop) {
  Dcpu cpu;
  const Word program[] = {
    // set push, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k13),
    // set a, pop
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kPop)
  };

  Dsl d;
  d.set(push(), 13)
    .set(a(), pop())
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, cpu.stack_pointer);
  cpu.ExecuteInstruction();
  EXPECT_EQ(0, cpu.stack_pointer);
  EXPECT_EQ(13, cpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_peek) {
  Dcpu cpu;
  const Word program[] = {
    // set push, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k13),
    // set a, peek
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kPeek)
  };

  Dsl d;
  d.set(push(), 13)
    .set(a(), peek())
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, cpu.stack_pointer);
  cpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, cpu.stack_pointer);
  EXPECT_EQ(13, cpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_pick) {
  Dcpu cpu;
  const Word program[] = {
    // set push, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k13),
    // set push, 14
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k14),
    // set a, [sp+1]
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kPick), 0x1
  };

  Dsl d;
  d.set(push(), 13)
    .set(push(), 14)
    .set(a(), pick(0x1))
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  EXPECT_EQ(0, cpu.stack_pointer);
  cpu.ExecuteInstructions(3);
  EXPECT_EQ(0xFFFE, cpu.stack_pointer);
  EXPECT_EQ(13, *cpu.address(cpu.stack_pointer + 1));
  EXPECT_EQ(13, cpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_stack_pointer) {
  Dcpu cpu;
  const Word program[] = {
    // set push, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k13),
    // set a, sp
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kStackPointer)
  };

  Dsl d;
  d.set(push(), 13)
    .set(a(), sp())
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, cpu.stack_pointer);
  cpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, cpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_program_counter) {
  Dcpu cpu;
  const Word program[] = {
    // noop
    Noop(),
    // set a, pc
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kProgramCounter)
  };

  Dsl d;
  d.set(a(), a())
    .set(a(), pc())
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(2, cpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_overflow) {
  Dcpu cpu;
  const Word program[] = {
    // set ex, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kExtra, dcpu::Operand::k13),
    // set a, ex
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kExtra)
  };

  Dsl d;
  d.set(ex(), 13)
    .set(a(), ex())
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(13, cpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_location) {
  Dcpu cpu;
  const Word program[] = {
    // set [0x1000], 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kLocation, dcpu::Operand::k13), 0x1000,
    // set a, [0x1000]
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kLocation), 0x1000
  };

  Dsl d;
  d.set(d[0x1000], 13)
    .set(a(), d[0x1000])
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(13, cpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_high_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0x1001
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kLiteral), 0x1001
  };

  Dsl d;
  d.set(a(), 0x1001)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstruction();
  EXPECT_EQ(0x1001, cpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_register_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 1
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k1),
  };

  Dsl d;
  d.set(a(), 1)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstruction();
  EXPECT_EQ(1, cpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_set_last_register_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set j, 1
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterJ, dcpu::Operand::k1),
  };

  Dsl d;
  d.set(j(), 1)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  // cpu.register_a == 0 by DcpuTest.DefaultConstructor
  cpu.ExecuteInstructions(2);
  EXPECT_EQ(1, cpu.register_j);
}

TEST(DcpuTest, ExecuteInstruction_set_location_in_register_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0x1000
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kLiteral), 0x1000,
    // set [a], 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kLocationInRegisterA, dcpu::Operand::k13)
  };

  Dsl d;
  d.set(a(), 0x1000)
    .set(d[a()], 13)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(13, *cpu.address(0x1000));
}

TEST(DcpuTest,
    ExecuteInstruction_set_location_in_last_register_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set j, 0x1000
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterJ, dcpu::Operand::kLiteral), 0x1000,
    // set [j], 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kLocationInRegisterJ, dcpu::Operand::k13)
  };

  Dsl d;
  d.set(j(), 0x1000)
    .set(d[j()], 13)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(13, *cpu.address(0x1000));
}

TEST(DcpuTest,
    ExecuteInstruction_set_location_offset_by_register_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 10
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k10),
    // set [0x1000+a], 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kLocationOffsetByRegisterA, dcpu::Operand::k13), 0x1000
  };

  Dsl d;
  d.set(a(), 0xA)
    .set(d[0x1000 + a()], 13)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(13, *cpu.address(0x100A));
}

TEST(DcpuTest,
    ExecuteInstruction_set_location_offset_by_register_with_location_offset_by_register) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 10
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k10),
    // set [0x1000+a], 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kLocationOffsetByRegisterA, dcpu::Operand::k13), 0x1000,
    // set [0x2000+a], [0x1000+a]
    Instruct(BasicOpcode::kSet,
        dcpu::Operand::kLocationOffsetByRegisterA, dcpu::Operand::kLocationOffsetByRegisterA), 0x1000, 0x2000
  };

  Dsl d;
  d.set(a(), 0xA)
    .set(d[0x1000 + a()], 13)
    .set(d[0x2000 + a()], d[0x1000 + a()])
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *cpu.address(0x200A));
}

TEST(DcpuTest,
    ExecuteInstruction_set_location_offset_by_last_register_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set j, 10
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterJ, dcpu::Operand::k10),
    // set [0x1000+j], 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kLocationOffsetByRegisterJ, dcpu::Operand::k13), 0x1000
  };

  Dsl d;
  d.set(j(), 0xA)
    .set(d[0x1000 + j()], 13)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(13, *cpu.address(0x100A));
}

TEST(DcpuTest, ExecuteInstruction_set_push_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set push, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k13)
  };

  Dsl d;
  d.set(push(), 13)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstruction();
  EXPECT_EQ(13, *cpu.address(cpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_set_push_with_pop) {
  Dcpu cpu;
  const Word program[] = {
    // set push, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k13),
    // set push, pop
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::kPop)
  };

  Dsl d;
  d.set(push(), 13)
    .set(push(), pop())
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(0xFFFF, cpu.stack_pointer);
  EXPECT_EQ(13, *cpu.address(cpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_set_peek_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set push, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k13),
    // set peek, 14
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPeek, dcpu::Operand::k14)
  };

  Dsl d;
  d.set(push(), 13)
    .set(peek(), 14)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(14, *cpu.address(cpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_set_pick_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set push, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k12),
    // set push, 14
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k13),
    // set [SP+0x1], 14
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPick, dcpu::Operand::k14), 0x1
  };

  Dsl d;
  d.set(push(), 12)
    .set(push(), 13)
    .set(pick(0x1), 14)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *cpu.address(cpu.stack_pointer));
  EXPECT_EQ(14, *cpu.address(cpu.stack_pointer + 1));
}

TEST(DcpuTest, ExecuteInstruction_set_stack_pointer_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set sp, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kStackPointer, dcpu::Operand::k13)
  };

  Dsl d;
  d.set(sp(), 13)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstruction();
  EXPECT_EQ(13, cpu.stack_pointer);
}

TEST(DcpuTest, ExecuteInstruction_set_program_counter_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set pc, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kProgramCounter, dcpu::Operand::k13)
  };

  Dsl d;
  d.set(pc(), 13)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstruction();
  EXPECT_EQ(13, cpu.program_counter);
}

TEST(DcpuTest, ExecuteInstruction_set_overflow_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set ex, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kExtra, dcpu::Operand::k13)
  };

  Dsl d;
  d.set(ex(), 13)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstruction();
  EXPECT_EQ(13, cpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_set_location_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set [0x1000], 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kLocation, dcpu::Operand::k13), 0x1000
  };

  Dsl d;
  d.set(d[0x1000], 13)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstruction();
  EXPECT_EQ(13, *cpu.address(0x1000));
}

TEST(DcpuTest, ExecuteInstruction_set_literal_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set 0x1000, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kLiteral, dcpu::Operand::k13), 0x1000
  };
  const Word *const program_end = program + sizeof(program) / sizeof(Word);
  std::copy(program, program_end, cpu.memory_begin());
  cpu.ExecuteInstruction();
  // set 0x1000, 13 should be a noop.
  EXPECT_EQ(0, *cpu.address(0x1000));
}

TEST(DcpuTest, ExecuteInstruction_add_register_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0x0D
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k13),
    // add a, 0x0E
    Instruct(BasicOpcode::kAdd, dcpu::Operand::kRegisterA, dcpu::Operand::k14)
  };

  Dsl d;
  d.set(a(), 13)
    .add(a(), 14)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(0x1B, cpu.register_a);
  EXPECT_EQ(0, cpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_add_register_with_overflow) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0xFFFF
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kNegative1),
    // add a, 0xFFFF
    Instruct(BasicOpcode::kAdd, dcpu::Operand::kRegisterA, dcpu::Operand::kNegative1)
  };

  Dsl d;
  d.set(a(), 0xFFFF)
    .add(a(), 0xFFFF)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(0xFFFE, cpu.register_a);
  EXPECT_EQ(1, cpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_subtract_register_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0x1E
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k30),
    // sub a, 0x10
    Instruct(BasicOpcode::kSubtract, dcpu::Operand::kRegisterA, dcpu::Operand::k16)
  };

  Dsl d;
  d.set(a(), 30)
    .sub(a(), 16)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(0xE, cpu.register_a);
  EXPECT_EQ(0, cpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_subtract_register_with_underflow) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0x10
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k16),
    // sub a, 0x1E
    Instruct(BasicOpcode::kSubtract, dcpu::Operand::kRegisterA, dcpu::Operand::k30)
  };

  Dsl d;
  d.set(a(), 16)
    .sub(a(), 30)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(0xFFF2, cpu.register_a);
  EXPECT_EQ(1, cpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_multiply_register_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0x10
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k16),
    // mul a, 0x1E
    Instruct(BasicOpcode::kMultiply, dcpu::Operand::kRegisterA, dcpu::Operand::k30)
  };

  Dsl d;
  d.set(a(), 16)
    .mul(a(), 30)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(0x01E0, cpu.register_a);
  EXPECT_EQ(0, cpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_multiply_register_with_overflow) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0xFFFF
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kNegative1),
    // mul a, 0xFFFF
    Instruct(BasicOpcode::kMultiply, dcpu::Operand::kRegisterA, dcpu::Operand::kNegative1)
  };

  Dsl d;
  d.set(a(), 0xFFFF)
    .mul(a(), 0xFFFF)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(0x0001, cpu.register_a);
  EXPECT_EQ(0xFFFE, cpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_divide_register_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0x1E
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k30),
    // div a, 0x10
    Instruct(BasicOpcode::kDivide, dcpu::Operand::kRegisterA, dcpu::Operand::k16)
  };

  Dsl d;
  d.set(a(), 30)
    .div(a(), 16)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(1, cpu.register_a);
  EXPECT_EQ(0, cpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_divide_register_by_zero) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0x1E
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k30),
    // div a, 0x00
    Instruct(BasicOpcode::kDivide, dcpu::Operand::kRegisterA, dcpu::Operand::k0)
  };

  Dsl d;
  d.set(a(), 30)
    .div(a(), 0)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(0, cpu.register_a);
  EXPECT_EQ(1, cpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_modulo_register_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0x1E
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k30),
    // mod a, 0x0B
    Instruct(BasicOpcode::kModulo, dcpu::Operand::kRegisterA, dcpu::Operand::k11)
  };

  Dsl d;
  d.set(a(), 30)
    .mod(a(), 11)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(0x8, cpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_shift_left_register_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0x1E
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k30),
    // shl a, 0x02
    Instruct(BasicOpcode::kShiftLeft, dcpu::Operand::kRegisterA, dcpu::Operand::k2)
  };

  Dsl d;
  d.set(a(), 30)
    .shl(a(), 2)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(0x78, cpu.register_a);
  EXPECT_EQ(0, cpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_shift_left_register_with_overflow) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0xFFFF
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kNegative1),
    // shl a, 0x02
    Instruct(BasicOpcode::kShiftLeft, dcpu::Operand::kRegisterA, dcpu::Operand::k2)
  };

  Dsl d;
  d.set(a(), 0xFFFF)
    .shl(a(), 2)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(0xFFFC, cpu.register_a);
  EXPECT_EQ(0x0003, cpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_shift_right_register_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0xFFF0
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kLiteral), 0xFFF0,
    // shr a, 0x02
    Instruct(BasicOpcode::kShiftRight, dcpu::Operand::kRegisterA, dcpu::Operand::k2)
  };

  Dsl d;
  d.set(a(), 0xFFF0)
    .shr(a(), 2)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(0x3FFC, cpu.register_a);
  EXPECT_EQ(0, cpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_shift_right_register_with_underflow) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0xFFFF
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kNegative1),
    // shr a, 0x02
    Instruct(BasicOpcode::kShiftRight, dcpu::Operand::kRegisterA, dcpu::Operand::k2)
  };

  Dsl d;
  d.set(a(), 0xFFFF)
    .shr(a(), 2)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(0x3FFF, cpu.register_a);
  EXPECT_EQ(0xC000, cpu.extra);
}

TEST(DcpuTest, ExecuteInstruction_and_register_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0xF0F0
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kLiteral), 0xF0F0,
    // and a, 0x00FF
    Instruct(BasicOpcode::kBinaryAnd, dcpu::Operand::kRegisterA, dcpu::Operand::kLiteral), 0x00FF
  };

  Dsl d;
  d.set(a(), 0xF0F0)
    .and_(a(), 0x00FF)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(0x00F0, cpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_or_register_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0xF0F0
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kLiteral), 0xF0F0,
    // bor a, 0x00FF
    Instruct(BasicOpcode::kBinaryOr, dcpu::Operand::kRegisterA, dcpu::Operand::kLiteral), 0x00FF
  };

  Dsl d;
  d.set(a(), 0xF0F0)
    .bor(a(), 0x00FF)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(0xF0FF, cpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_xor_register_with_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0xF0F0
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::kLiteral), 0xF0F0,
    // xor a, 0x00FF
    Instruct(BasicOpcode::kBinaryExclusiveOr, dcpu::Operand::kRegisterA, dcpu::Operand::kLiteral), 0x00FF
  };

  Dsl d;
  d.set(a(), 0xF0F0)
    .xor_(a(), 0x00FF)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(2);
  EXPECT_EQ(0xF00F, cpu.register_a);
}

TEST(DcpuTest, ExecuteInstruction_if_equal_register_with_equal_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0x0F
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k15),
    // ife a, 0x0F
    Instruct(BasicOpcode::kIfEqual, dcpu::Operand::kRegisterA, dcpu::Operand::k15),
    // set push, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k13),
    // set push, 14
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k14)
  };

  Dsl d;
  d.set(a(), 15)
    .ife(a(), 15)
      .set(push(), 13)
    .set(push(), 14)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *cpu.address(cpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_if_equal_register_with_unequal_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0x0F
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k15),
    // ife a, 0x00
    Instruct(BasicOpcode::kIfEqual, dcpu::Operand::kRegisterA, dcpu::Operand::k0),
    // set push, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k13),
    // set push, 14
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k14)
  };

  Dsl d;
  d.set(a(), 15)
    .ife(a(), 0)
      .set(push(), 13)
    .set(push(), 14)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(3);
  EXPECT_EQ(14, *cpu.address(cpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_if_not_equal_register_with_unequal_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0x0F
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k15),
    // ifn a, 0x00
    Instruct(BasicOpcode::kIfNotEqual, dcpu::Operand::kRegisterA, dcpu::Operand::k0),
    // set push, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k13),
    // set push, 14
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k14)
  };

  Dsl d;
  d.set(a(), 15)
    .ifn(a(), 0)
      .set(push(), 13)
    .set(push(), 14)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *cpu.address(cpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_if_not_equal_register_with_equal_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0x0F
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k15),
    // ifn a, 0x0F
    Instruct(BasicOpcode::kIfNotEqual, dcpu::Operand::kRegisterA, dcpu::Operand::k15),
    // set push, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k13),
    // set push, 14
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k14)
  };

  Dsl d;
  d.set(a(), 15)
    .ifn(a(), 15)
      .set(push(), 13)
    .set(push(), 14)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(3);
  EXPECT_EQ(14, *cpu.address(cpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_if_greater_than_register_with_lesser_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0x1E
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k30),
    // ifg a, 0x0F
    Instruct(BasicOpcode::kIfGreaterThan, dcpu::Operand::kRegisterA, dcpu::Operand::k15),
    // set push, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k13),
    // set push, 14
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k14)
  };

  Dsl d;
  d.set(a(), 30)
    .ifg(a(), 15)
      .set(push(), 13)
    .set(push(), 14)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *cpu.address(cpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_if_greater_than_register_with_greater_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0x0F
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k15),
    // ifg a, 0x1E
    Instruct(BasicOpcode::kIfGreaterThan, dcpu::Operand::kRegisterA, dcpu::Operand::k30),
    // set push, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k13),
    // set push, 14
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k14)
  };

  Dsl d;
  d.set(a(), 15)
    .ifg(a(), 30)
      .set(push(), 13)
    .set(push(), 14)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(3);
  EXPECT_EQ(14, *cpu.address(cpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_if_both_register_with_common_bits_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0x1E
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k30),
    // ifb a, 0x10
    Instruct(BasicOpcode::kIfBitSet, dcpu::Operand::kRegisterA, dcpu::Operand::k16),
    // set push, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k13),
    // set push, 14
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k14)
  };

  Dsl d;
  d.set(a(), 0x1E)
    .ifb(a(), 0x10)
      .set(push(), 13)
    .set(push(), 14)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(3);
  EXPECT_EQ(13, *cpu.address(cpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_if_both_register_with_uncommon_bits_low_literal) {
  Dcpu cpu;
  const Word program[] = {
    // set a, 0x0F
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k15),
    // ifb a, 0x10
    Instruct(BasicOpcode::kIfBitSet, dcpu::Operand::kRegisterA, dcpu::Operand::k16),
    // set push, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k13),
    // set push, 14
    Instruct(BasicOpcode::kSet, dcpu::Operand::kPush, dcpu::Operand::k14)
  };

  Dsl d;
  d.set(a(), 0x0F)
    .ifb(a(), 0x10)
      .set(push(), 13)
    .set(push(), 14)
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstructions(3);
  EXPECT_EQ(14, *cpu.address(cpu.stack_pointer));
}

TEST(DcpuTest, ExecuteInstruction_jump_sub_routine) {
  Dcpu cpu;
  const Word program[] = {
    // jsr subroutine
    Instruct(AdvancedOpcode::kJumpSubRoutine, dcpu::Operand::kLiteral), 0x0004,
    // set a, 13
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterA, dcpu::Operand::k13),
    // sub pc, 1
    Instruct(BasicOpcode::kSubtract, dcpu::Operand::kProgramCounter, dcpu::Operand::k1),
    // subroutine: set b, 14
    Instruct(BasicOpcode::kSet, dcpu::Operand::kRegisterB, dcpu::Operand::k14),
    // set pc, pop
    Instruct(BasicOpcode::kSet, dcpu::Operand::kProgramCounter, dcpu::Operand::kPop)
  };

  Dsl d;
  d.jsr("subroutine")
    .set(a(), 13)
    .sub(pc(), 1)
    .label("subroutine")
      .set(b(), 14)
      .set(pc(), pop())
    .Assemble(cpu.memory_begin());

  for (auto i = 0; i < sizeof(program) / sizeof(Word); ++i) {
    EXPECT_EQ(program[i], *cpu.address(i));
  }

  cpu.ExecuteInstruction();
  EXPECT_EQ(0xFFFF, cpu.stack_pointer);
  cpu.ExecuteInstructions(4);
  EXPECT_EQ(13, cpu.register_a);
  EXPECT_EQ(14, cpu.register_b);
  EXPECT_EQ(3, cpu.program_counter);
  EXPECT_EQ(0, cpu.stack_pointer);
}
