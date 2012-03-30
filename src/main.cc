// Copyright 2012 Robert Scott Dionne. All rights reserved.

#include <algorithm>
#include <ncurses.h>
#include <string>
#include "dcpu.h"

int main(int argc, char *argv[]) {
  Dcpu dcpu;
  initscr();
  std::string hello_world("Hello, world!");
  std::copy(hello_world.begin(), hello_world.end(), dcpu.address(0x8000));
  for (int i = 0; i < 1000; ++i) {
    char character = *dcpu.address(0x8000 + i);
    if (character) {
      addch(character);
    }
  }
  refresh();
  getch();
  endwin();
  return 0;
}
