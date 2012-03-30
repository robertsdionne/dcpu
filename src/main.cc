// Copyright 2012 Robert Scott Dionne. All rights reserved.

#include <ncurses.h>
#include <string>
#include "dcpu.h"

int main(int argc, char *argv[]) {
  Dcpu dcpu;
  initscr();
  unsigned short screen[1000];
  dcpu.ReadVideoMemory(screen);
  for (int i = 0; i < 1000; ++i) {
    if (screen[i]) {
      addch(screen[i]);
    }
  }
  refresh();
  getch();
  endwin();
  return 0;
}
