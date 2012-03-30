// Copyright 2012 Robert Scott Dionne. All rights reserved.

#ifndef DCPU_DCPU_H_
#define DCPU_DCPU_H_

class Dcpu {
  public:
    typedef unsigned short Word;

    static const unsigned int kMemorySize = 0x10000;
    static const Word kVideoMemoryBegin = 0x8000;
    static const Word kVideoMemoryEnd = 0x87D0;

  public:
    Dcpu();
    virtual ~Dcpu() {}

    Word *address(const Word address_value);
    const Word *address(const Word address_value) const;

    Word *memory_begin();
    const Word *memory_begin() const;
    Word *memory_end();
    const Word *memory_end() const;

    Word *video_memory_begin();
    const Word *video_memory_begin() const;
    Word *video_memory_end();
    const Word *video_memory_end() const;

  private:
    Word memory_[kMemorySize];
    Word register_a_;
    Word register_b_;
    Word register_c_;
    Word register_x_;
    Word register_y_;
    Word register_z_;
    Word register_i_;
    Word register_j_;
    Word program_counter_;
    Word stack_pointer_;
    Word overflow_;
};

#endif  // DCPU_DCPU_H_
