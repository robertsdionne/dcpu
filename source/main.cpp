#include <algorithm>
#include <fstream>
#include <iostream>
#include <ncurses.h>
#include <string>

#include "assembler.hpp"
#include "clock.hpp"
#include "dcpu.hpp"

using namespace dcpu;

constexpr Word kVideoMemoryBegin = 0x8000;

int main(int argument_count, char *arguments[]) {
  Dcpu cpu;
  Clock clock;
  cpu.Connect(&clock);
  Assembler assembler;
  initscr();
  Word program[] = {
    Instruct(AdvancedOpcode::kInterruptAddressSet, Operand::kLiteral),
    0x000B,
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k0),
    Instruct(BasicOpcode::kSet, Operand::kRegisterB, Operand::kLiteral),
    0x003C,
    Instruct(AdvancedOpcode::kHardwareInterrupt, Operand::k0),
    Instruct(BasicOpcode::kSet, Operand::kRegisterA, Operand::k2),
    Instruct(BasicOpcode::kSet, Operand::kRegisterB, Operand::k13),
    Instruct(AdvancedOpcode::kHardwareInterrupt, Operand::k0),
    Instruct(BasicOpcode::kSubtract, Operand::kProgramCounter, Operand::k1),
    // state = 0x000A
    0x0000,
    // handler = 0x000B
    Instruct(BasicOpcode::kSet, Operand::kRegisterI, Operand::k0),
    Instruct(BasicOpcode::kIfEqual, Operand::kLocation, Operand::k0),
    0x000A,
    Instruct(BasicOpcode::kSet, Operand::kProgramCounter, Operand::kLiteral),
    0x001B,
    // tock = 0x0010
    Instruct(BasicOpcode::kSet, Operand::kLocation, Operand::k0),
    0x000A,
    Instruct(BasicOpcode::kIfEqual, Operand::kLocationOffsetByRegisterI, Operand::k0),
    0x002B,
    Instruct(AdvancedOpcode::kReturnFromInterrupt, Operand::k0),
    Instruct(BasicOpcode::kSet,
        Operand::kLocationOffsetByRegisterI, Operand::kLocationOffsetByRegisterI),
    0x002B,
    kVideoMemoryBegin,
    Instruct(BasicOpcode::kAdd, Operand::kRegisterI, Operand::k1),
    Instruct(BasicOpcode::kSet, Operand::kProgramCounter, Operand::kLiteral),
    0x0010,
    // tick = 0x001B
    Instruct(BasicOpcode::kSet, Operand::kLocation, Operand::k1),
    0x000A,
    Instruct(BasicOpcode::kIfEqual, Operand::kLocationOffsetByRegisterI, Operand::k0),
    0x0026,
    Instruct(AdvancedOpcode::kReturnFromInterrupt, Operand::k0),
    Instruct(BasicOpcode::kSet,
        Operand::kLocationOffsetByRegisterI, Operand::kLocationOffsetByRegisterI),
    0x0026,
    kVideoMemoryBegin,
    Instruct(BasicOpcode::kAdd, Operand::kRegisterI, Operand::k1),
    Instruct(BasicOpcode::kSet, Operand::kProgramCounter, Operand::kLiteral),
    0x001B,
    // tick_data = 0x0026
    't', 'i', 'c', 'k', 0,
    // tock_data = 0x002B
    't', 'o', 'c', 'k', 0
  };

  std::copy(program,
            program + sizeof(program)/sizeof(Word), cpu.memory_begin());
  bool quit = false;
  while (!quit) {
    move(0, 0);
    printw("Push any key to advance one cycle; push q to quit.\n\n");
    printw("Registers:\n");
    printw("A: %X B: %X C: %X ", cpu.register_a, cpu.register_b, cpu.register_c);
    printw("X: %X Y: %X Z: %X ", cpu.register_x, cpu.register_y, cpu.register_z);
    printw("I: %X J: %X\n", cpu.register_i, cpu.register_j);
    printw("PC: %X SP: %X EX: %X IA: %X\n\n",
        cpu.program_counter, cpu.stack_pointer,
        cpu.extra, cpu.interrupt_address);
    printw("Instruction(s): %X %X %X\n\n",
        *cpu.address(cpu.program_counter),
        *cpu.address(cpu.program_counter + 1),
        *cpu.address(cpu.program_counter + 2));
    printw("Memory:\n");
    printw("0x1000: %X\n", *cpu.address(0x1000));
    printw("Video Memory:\n");
    printw("0x8000: %X %X %X %X %X %X %X %X %X %X %X %X %X %X %X %X\n\n",
        *(cpu.address(kVideoMemoryBegin) + 0x0),
        *(cpu.address(kVideoMemoryBegin) + 0x1),
        *(cpu.address(kVideoMemoryBegin) + 0x2),
        *(cpu.address(kVideoMemoryBegin) + 0x3),
        *(cpu.address(kVideoMemoryBegin) + 0x4),
        *(cpu.address(kVideoMemoryBegin) + 0x5),
        *(cpu.address(kVideoMemoryBegin) + 0x6),
        *(cpu.address(kVideoMemoryBegin) + 0x7),
        *(cpu.address(kVideoMemoryBegin) + 0x8),
        *(cpu.address(kVideoMemoryBegin) + 0x9),
        *(cpu.address(kVideoMemoryBegin) + 0xA),
        *(cpu.address(kVideoMemoryBegin) + 0xB),
        *(cpu.address(kVideoMemoryBegin) + 0xC),
        *(cpu.address(kVideoMemoryBegin) + 0xD),
        *(cpu.address(kVideoMemoryBegin) + 0xE),
        *(cpu.address(kVideoMemoryBegin) + 0xF));
    printw("Display:\n");
    for (int i = 0; i < 160; ++i) {
      char character = *(cpu.address(kVideoMemoryBegin) + i);
      if (character) {
        addch(character);
      } else {
        addch(' ');
      }
    }
    refresh();
    // quit = getch() == 'q';
    cpu.ExecuteInstruction();
    clock.Execute();
  }
  endwin();
  return 0;
}
