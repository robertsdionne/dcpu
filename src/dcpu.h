// Copyright 2012 Robert Scott Dionne. All rights reserved.

#ifndef DCPU_DCPU_H_
#define DCPU_DCPU_H_

class Dcpu {
  public:
    typedef unsigned short Word;

  public:
    Dcpu();
    virtual ~Dcpu() {}

    Word *address(const Word address_value);
    const Word *address(const Word address_value) const;
    Word *begin();
    const Word *begin() const;
    Word *end();
    const Word *end() const;

    void ReadVideoMemory(Word *const out) const;
    void WriteVideoMemory(Word *const begin, Word *const end);

  private:
    Word memory_[0x10000];
};

#endif  // DCPU_DCPU_H_
