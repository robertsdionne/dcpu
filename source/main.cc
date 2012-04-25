// Copyright 2012 Robert Scott Dionne. All rights reserved.

#include <algorithm>
#include <ncurses.h>
#include <string>
#include "dcpu.h"

const Dcpu::Word kVideoMemoryBegin = 0x8000;

int main(int argc, char *argv[]) {
  Dcpu dcpu;
  initscr();
  Dcpu::Word program[] = {
    // set a, 0xBEEF
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xBEEF,
    // set [0x1000], a
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocation, Dcpu::kRegisterA),
    0x1000,
    // ifn a, [0x1000]
    Dcpu::Instruct(Dcpu::kIfNotEqual, Dcpu::kRegisterA, Dcpu::kLocation),
    0x1000,
    //   set pc, end
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kProgramCounter, Dcpu::k29),
    // set i, 0
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterI, Dcpu::k0),
    // nextchar: ife i, 0
    Dcpu::Instruct(Dcpu::kIfEqual, Dcpu::kLocationOffsetByRegisterI, Dcpu::k0),
    0x0010,
    //   set pc, end
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kProgramCounter, Dcpu::k29),
    // set [0x8000+i], [data+i]
    Dcpu::Instruct(Dcpu::kSet,
        Dcpu::kLocationOffsetByRegisterI, Dcpu::kLocationOffsetByRegisterI),
    kVideoMemoryBegin,
    0x0010,
    // add i, 1
    Dcpu::Instruct(Dcpu::kAdd, Dcpu::kRegisterI, Dcpu::k1),
    // set pc, nextchar
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kProgramCounter, Dcpu::k8),
    // data: dat "Hello world!", 0
    'H', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd', '!', '\0',
    // end: sub pc, 1
    Dcpu::Instruct(Dcpu::kSubtract, Dcpu::kProgramCounter, Dcpu::k1)
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
    printw("PC: %X SP: %X O: %X\n\n",
        dcpu.program_counter(), dcpu.stack_pointer(), dcpu.overflow());
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
