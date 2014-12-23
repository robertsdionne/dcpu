#include <algorithm>
#include <fstream>
#include <glfwapplication.hpp>
#include <iostream>
#include <mouse.hpp>
#include <ncurses.h>
#include <string>

#include "assembler.hpp"
#include "clock.hpp"
#include "dcpu.hpp"
#include "dcpurenderer.hpp"
#include "dsl.hpp"
#include "lexer.hpp"
#include "parser.hpp"

using namespace dcpu;
using namespace dcpu::dsl;

constexpr Word kVideoMemoryBegin = 0x8000;

int main(int argument_count, char *arguments[]) {
  auto renderer = DcpuRenderer{};
  auto mouse = rsd::Mouse{};
  auto application = rsd::GlfwApplication{
    argument_count, arguments, 1024, 768, 7, "DCPU", renderer, mouse
  };
  application.Run();

  // auto source = std::string(R"(
  // # comments
  // set a, 0
  // SET b, a
  // set [c], a
  // set [x + 0x0000], 0
  // set [0x0000 + y], 1024
  // set push, 0
  // set a, pop
  // set b, peek
  // set a, [sp]
  // set [sp], a
  // IFN a, peek
  // set peek, NULLPTR
  // set pick 1, nullptr
  // set a, PICK 2
  // set [sp + 0], 0b0011
  // set [0 + sp], a
  // set [0x0000], 0b0000101011110101
  // label:
  // set 1, 0x0000
  // .DATA "asdf", 0
  // .data 1, 2, 3, 4, 5, "asdf", 0
  // .data 0x1000, 0xAAAA, 0x3000
  // set a, label
  // set b, [label]
  // set [label + i], 0
  // set [i + label], 0
  // set sp, 0
  // set pc, 0
  // set ex, 0
  // jsr a
  // int a
  // )");
  //
  // // std::cout << source << std::endl;
  //
  // source = std::string(
  //   "set a, 0xAB00"
  //   "my_label:"
  //   ".data 0x00CD"
  // );
  //
  // auto lexer = Lexer(source);
  // auto parser = Parser(lexer);
  // auto assembler = Assembler();
  //
  // auto program = proto::Program();
  // if (!parser.ParseProgram(&program)) {
  //   std::cout << "error compiling program" << std::endl;
  //   return 1;
  // }
  //
  // return 0;
  //
  //
  // while (true) {
  //   using Type = Token::Type;
  //   auto token = lexer.EatToken();
  //   if (Type::kInvalid == token.type) {
  //     break;
  //   }
  //   if (Type::kWhitespace != token.type) {
  //     std::cout << token << std::endl;
  //   }
  // }
  //
  // return 0;
  // initscr();
  //
  // Dcpu cpu;
  // Clock clock;
  // cpu.Connect(&clock);
  // assembler.Assemble(program, cpu.memory_begin());
  //
  // Dsl d;
  // d.ias("handler")
  //   .set(a, 0)
  //   .set(b, 60)
  //   .hwi(0)
  //   .set(a, 2)
  //   .set(b, 13)
  //   .hwi(0)
  //   .sub(pc, 1)
  //
  //   .label("state")
  //     .data(0)
  //
  //   .label("handler")
  //     .set(i, 0)
  //     .set(j, 0)
  //     .ife(d["state"], 0)
  //       .set(pc, "tick")
  //
  //   .label("tock")
  //     .set(d["state"], 0)
  //     .ife(d["tock_data" + j], 0)
  //       .rfi()
  //     .sti(d[kVideoMemoryBegin + i], d["tock_data" + j])
  //     .set(pc, "tock")
  //
  //   .label("tick")
  //     .set(d["state"], 1)
  //     .ife(d["tick_data" + j], 0)
  //       .rfi()
  //     .sti(d[kVideoMemoryBegin + i], d["tick_data" + j])
  //     .set(pc, "tick")
  //
  //   .label("tick_data")
  //     .data("tick")
  //
  //   .label("tock_data")
  //     .data("tock")
  //
  //   .Assemble(cpu.memory_begin());
  //
  // bool quit = false;
  // while (!quit) {
  //   move(0, 0);
  //   printw("Push any key to advance one cycle; push q to quit.\n\n");
  //   printw("Registers:\n");
  //   printw("A: %X B: %X C: %X ", cpu.register_a, cpu.register_b, cpu.register_c);
  //   printw("X: %X Y: %X Z: %X ", cpu.register_x, cpu.register_y, cpu.register_z);
  //   printw("I: %X J: %X\n", cpu.register_i, cpu.register_j);
  //   printw("PC: %X SP: %X EX: %X IA: %X\n\n",
  //       cpu.program_counter, cpu.stack_pointer,
  //       cpu.extra, cpu.interrupt_address);
  //   printw("Instruction(s): %X %X %X\n\n",
  //       *cpu.address(cpu.program_counter),
  //       *cpu.address(cpu.program_counter + 1),
  //       *cpu.address(cpu.program_counter + 2));
  //   printw("Memory:\n");
  //   printw("0x1000: %X\n", *cpu.address(0x1000));
  //   printw("Video Memory:\n");
  //   printw("0x8000: %X %X %X %X %X %X %X %X %X %X %X %X %X %X %X %X\n\n",
  //       *(cpu.address(kVideoMemoryBegin) + 0x0),
  //       *(cpu.address(kVideoMemoryBegin) + 0x1),
  //       *(cpu.address(kVideoMemoryBegin) + 0x2),
  //       *(cpu.address(kVideoMemoryBegin) + 0x3),
  //       *(cpu.address(kVideoMemoryBegin) + 0x4),
  //       *(cpu.address(kVideoMemoryBegin) + 0x5),
  //       *(cpu.address(kVideoMemoryBegin) + 0x6),
  //       *(cpu.address(kVideoMemoryBegin) + 0x7),
  //       *(cpu.address(kVideoMemoryBegin) + 0x8),
  //       *(cpu.address(kVideoMemoryBegin) + 0x9),
  //       *(cpu.address(kVideoMemoryBegin) + 0xA),
  //       *(cpu.address(kVideoMemoryBegin) + 0xB),
  //       *(cpu.address(kVideoMemoryBegin) + 0xC),
  //       *(cpu.address(kVideoMemoryBegin) + 0xD),
  //       *(cpu.address(kVideoMemoryBegin) + 0xE),
  //       *(cpu.address(kVideoMemoryBegin) + 0xF));
  //   printw("Display:\n");
  //   for (int i = 0; i < 160; ++i) {
  //     char character = *(cpu.address(kVideoMemoryBegin) + i);
  //     if (character) {
  //       addch(character);
  //     } else {
  //       addch(' ');
  //     }
  //   }
  //   refresh();
  //   // quit = getch() == 'q';
  //   cpu.ExecuteInstruction();
  //   clock.Execute();
  // }
  // endwin();
  return 0;
}
