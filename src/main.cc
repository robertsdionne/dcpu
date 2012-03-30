// Copyright 2012 Robert Scott Dionne. All rights reserved.

#include <algorithm>
#include <ncurses.h>
#include <string>
#include "dcpu.h"

int main(int argc, char *argv[]) {
  Dcpu dcpu;
  initscr();
  std::string hello_world("Hello, world!");
  std::copy(hello_world.begin(), hello_world.end(), dcpu.video_memory_begin());
  for (int i = 0; i < 1000; ++i) {
    char character = *(dcpu.video_memory_begin() + i);
    if (character) {
      addch(character);
    }
  }
  refresh();
  getch();
  endwin();
  return 0;
}
