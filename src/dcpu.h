// Copyright 2012 Robert Scott Dionne. All rights reserved.

#ifndef DCPU_DCPU_H_
#define DCPU_DCPU_H_

class Dcpu {
  public:
    Dcpu();
    virtual ~Dcpu() {}

  private:
    unsigned short memory_[0x10000];
};

#endif  // DCPU_DCPU_H_
