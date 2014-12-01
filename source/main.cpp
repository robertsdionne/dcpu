#include <algorithm>
#include <fstream>
#include <iostream>
#include <ncurses.h>
#include <string>

#include "clock.hpp"
#include "dcpu.hpp"
#include "dsl.hpp"

using namespace dcpu;
using namespace dcpu::dsl;

constexpr Word kVideoMemoryBegin = 0x8000;

int main(int argument_count, char *arguments[]) {
  initscr();

  Dcpu cpu;
  Clock clock;
  cpu.Connect(&clock);

  Dsl d;
  d.ias("handler")
    .set(a(), 0)
    .set(b(), 60)
    .hwi(0)
    .set(a(), 2)
    .set(b(), 13)
    .hwi(0)
    .sub(pc(), 1)

    .label("state")
      .data(0)

    .label("handler")
      .set(i(), 0)
      .set(j(), 0)
      .ife(d["state"], 0)
        .set(pc(), "tick")

    .label("tock")
      .set(d["state"], 0)
      .ife(d["tock_data" + j()], 0)
        .rfi()
      .sti(d[kVideoMemoryBegin + i()], d["tock_data" + j()])
      .set(pc(), "tock")

    .label("tick")
      .set(d["state"], 1)
      .ife(d["tick_data" + j()], 0)
        .rfi()
      .sti(d[kVideoMemoryBegin + i()], d["tick_data" + j()])
      .set(pc(), "tick")

    .label("tick_data")
      .data("tick")

    .label("tock_data")
      .data("tock")

    .Assemble(cpu.memory_begin());

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
