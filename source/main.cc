// Copyright 2012 Robert Scott Dionne. All rights reserved.

#include <algorithm>
#include <fstream>
#include <iostream>
#include <ncurses.h>
#include <string>

#include "assembler.h"
#include "dcpu.h"

const Dcpu::Word kVideoMemoryBegin = 0x8000;

int main(int argument_count, char *arguments[]) {
  Dcpu dcpu;
  Assembler assembler;
  initscr();
  Dcpu::Word program[] = {
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::kRegisterA, Dcpu::kLiteral), 0xBEEF,               // set a, 0xBEEF
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::kLocation, Dcpu::kRegisterA), 0x1000,              // set [0x1000], a
    Dcpu::Instruct(Dcpu::BasicOpcode::kIfNotEqual, Dcpu::kRegisterA, Dcpu::kLocation), 0x1000,       // ifn a, [0x1000]
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::kProgramCounter, Dcpu::k29),                       //     set pc, end
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::kRegisterI, Dcpu::k0),                             // set i, 0
    Dcpu::Instruct(Dcpu::BasicOpcode::kIfEqual, Dcpu::kLocationOffsetByRegisterI, Dcpu::k0), 0x0010, // nextchar: ife [data+i], 0
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::kProgramCounter, Dcpu::k29),                       //     set pc, end
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet,
                   Dcpu::kLocationOffsetByRegisterI,
                   Dcpu::kLocationOffsetByRegisterI), 0x0010, kVideoMemoryBegin,        // set [0x8000+i], [data+i]
    Dcpu::Instruct(Dcpu::BasicOpcode::kAdd, Dcpu::kRegisterI, Dcpu::k1),                             // add i, 1
    Dcpu::Instruct(Dcpu::BasicOpcode::kSet, Dcpu::kProgramCounter, Dcpu::k8),                        // set pc, nextchar
    'H', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd', '!', '\0',                   // data: dat "Hello world!", 0
    Dcpu::Instruct(Dcpu::BasicOpcode::kSubtract, Dcpu::kProgramCounter, Dcpu::k1)                    // end: sub pc, 1
  };
  std::copy(program,
            program + sizeof(program)/sizeof(Dcpu::Word), dcpu.memory_begin());
  bool quit = false;
  while (!quit) {
    move(0, 0);
    printw("Push any key to advance one cycle; push q to quit.\n\n");
    printw("Registers:\n");
    printw("A: %X B: %X C: %X ",
           dcpu.register_a(), dcpu.register_b(), dcpu.register_c());
    printw("X: %X Y: %X Z: %X ",
           dcpu.register_x(), dcpu.register_y(), dcpu.register_z());
    printw("I: %X J: %X\n", dcpu.register_i(), dcpu.register_j());
    printw("PC: %X SP: %X EX: %X IA: %X\n\n",
           dcpu.program_counter(), dcpu.stack_pointer(),
           dcpu.extra(), dcpu.interrupt_address());
    printw("Instruction(s): %X %X %X\n\n",
           *dcpu.address(dcpu.program_counter()),
           *dcpu.address(dcpu.program_counter() + 1),
           *dcpu.address(dcpu.program_counter() + 2));
    printw("Memory:\n");
    printw("0x1000: %X\n", *dcpu.address(0x1000));
    printw("Video Memory:\n");
    printw("0x8000: %X %X %X %X %X %X %X %X %X %X %X %X %X %X %X %X\n\n",
           *(dcpu.address(kVideoMemoryBegin) + 0x0),
           *(dcpu.address(kVideoMemoryBegin) + 0x1),
           *(dcpu.address(kVideoMemoryBegin) + 0x2),
           *(dcpu.address(kVideoMemoryBegin) + 0x3),
           *(dcpu.address(kVideoMemoryBegin) + 0x4),
           *(dcpu.address(kVideoMemoryBegin) + 0x5),
           *(dcpu.address(kVideoMemoryBegin) + 0x6),
           *(dcpu.address(kVideoMemoryBegin) + 0x7),
           *(dcpu.address(kVideoMemoryBegin) + 0x8),
           *(dcpu.address(kVideoMemoryBegin) + 0x9),
           *(dcpu.address(kVideoMemoryBegin) + 0xA),
           *(dcpu.address(kVideoMemoryBegin) + 0xB),
           *(dcpu.address(kVideoMemoryBegin) + 0xC),
           *(dcpu.address(kVideoMemoryBegin) + 0xD),
           *(dcpu.address(kVideoMemoryBegin) + 0xE),
           *(dcpu.address(kVideoMemoryBegin) + 0xF));
    printw("Display:\n");
    for (int i = 0; i < 160; ++i) {
      char character = *(dcpu.address(kVideoMemoryBegin) + i);
      if (character) {
        addch(character);
      } else {
        addch(' ');
      }
    }
    refresh();
    quit = getch() == 'q';
    dcpu.ExecuteInstruction();
  }
  endwin();
  return 0;
}
