// Copyright 2012 Robert Scott Dionne. All rights reserved.

#include <algorithm>
#include <ncurses.h>
#include <string>
#include "dcpu.h"

int main(int argc, char *argv[]) {
  Dcpu dcpu;
  initscr();
  Dcpu::Word program[] = {
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterA, Dcpu::kLiteral),
    0xBEEF,
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kLocation, Dcpu::kRegisterA),
    0x1000,
    Dcpu::Instruct(Dcpu::kIfNotEqual, Dcpu::kRegisterA, Dcpu::kLocation),
    0x1000,
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kProgramCounter, Dcpu::k29),
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kRegisterI, Dcpu::k0),
    Dcpu::Instruct(Dcpu::kIfEqual, Dcpu::kLocationOffsetByRegisterI, Dcpu::k0),
    0x0010,
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kProgramCounter, Dcpu::k29),
    Dcpu::Instruct(Dcpu::kSet,
        Dcpu::kLocationOffsetByRegisterI, Dcpu::kLocationOffsetByRegisterI),
    Dcpu::kVideoMemoryBegin,
    0x0010,
    Dcpu::Instruct(Dcpu::kAdd, Dcpu::kRegisterI, Dcpu::k1),
    Dcpu::Instruct(Dcpu::kSet, Dcpu::kProgramCounter, Dcpu::k8),
    'H', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd', '!', '\0',
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
        *(dcpu.video_memory_begin() + 0x0),
        *(dcpu.video_memory_begin() + 0x1),
        *(dcpu.video_memory_begin() + 0x2),
        *(dcpu.video_memory_begin() + 0x3),
        *(dcpu.video_memory_begin() + 0x4),
        *(dcpu.video_memory_begin() + 0x5),
        *(dcpu.video_memory_begin() + 0x6),
        *(dcpu.video_memory_begin() + 0x7),
        *(dcpu.video_memory_begin() + 0x8),
        *(dcpu.video_memory_begin() + 0x9),
        *(dcpu.video_memory_begin() + 0xA),
        *(dcpu.video_memory_begin() + 0xB),
        *(dcpu.video_memory_begin() + 0xC),
        *(dcpu.video_memory_begin() + 0xD),
        *(dcpu.video_memory_begin() + 0xE),
        *(dcpu.video_memory_begin() + 0xF));
    printw("Display:\n");
    for (int i = 0; i < 160; ++i) {
      char character = *(dcpu.video_memory_begin() + i);
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
