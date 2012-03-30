// Copyright 2012 Robert Scott Dionne. All rights reserved.

#include <ncurses.h>
#include "dcpu.h"

int main(int argc, char *argv[]) {
  initscr();
  printw("Hello, world!");
  refresh();
  getch();
  endwin();
  return 0;
}
