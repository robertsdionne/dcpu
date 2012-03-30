// Copyright 2012 Robert Scott Dionne. All rights reserved.

#ifndef DCPU_DCPU_H_
#define DCPU_DCPU_H_

class Dcpu {
  public:
    Dcpu();
    virtual ~Dcpu() {}

    unsigned short *address(unsigned short address_value);
    const unsigned short *address(unsigned short address_value) const;
    unsigned short *begin();
    const unsigned short *begin() const;
    unsigned short *end();
    const unsigned short *end() const;

    void ReadVideoMemory(unsigned short *const out) const;
    void WriteVideoMemory(
        unsigned short *const begin, unsigned short *const end) const;

  private:
    unsigned short memory_[0x10000];
};

#endif  // DCPU_DCPU_H_
