// Copyright 2012 Robert Scott Dionne. All rights reserved.

#ifndef DCPU_DCPU_H_
#define DCPU_DCPU_H_

class Dcpu {
  public:
    typedef unsigned short Word;

    /**
     * Memory landmarks.
     */
    static const unsigned int kMemorySize = 0x10000;

    /**
     * Opcodes.
     */
    static const Word kBasicOpcodeMask = 0x001F;
    static const Word kAdvancedOpcodeMask = 0x03E0;
    static const Word kAdvancedOpcodeShift = 0x5;

    enum BasicOpcode {
      kBasicReserved = 0x0,
      kSet = 0x1,
      kAdd = 0x2,
      kSubtract = 0x3,
      kMultiply = 0x4,
      kDivide = 0x5,
      kModulo = 0x6,
      kShiftLeft = 0x7,
      kShiftRight = 0x8,
      kBinaryAnd = 0x9,
      kBinaryOr = 0xA,
      kBinaryExclusiveOr = 0xB,
      kIfEqual = 0xC,
      kIfNotEqual = 0xD,
      kIfGreaterThan = 0xE,
      kIfBoth = 0xF
    };

    enum AdvancedOpcode {
      kAdvancedReserved = 0x00,
      kJumpSubRoutine = 0x01
    };

    /**
     * Operands.
     */
    static const Word kBasicOperandMaskA = 0xFC00;
    static const Word kBasicOperandMaskB = 0x03E0;
    static const Word kBasicOperandShiftA = 0xA;
    static const Word kBasicOperandShiftB = 0x5;
    static const Word kAdvancedOperandMaskA = 0xFC00;
    static const Word kAdvancedOperandShiftA = 0xA;

    enum Operand {
      kRegisterA = 0x00,
      kRegisterB = 0x01,
      kRegisterC = 0x02,
      kRegisterX = 0x03,
      kRegisterY = 0x04,
      kRegisterZ = 0x05,
      kRegisterI = 0x06,
      kRegisterJ = 0x07,
      kLocationInRegisterA = 0x08,
      kLocationInRegisterB = 0x09,
      kLocationInRegisterC = 0x0A,
      kLocationInRegisterX = 0x0B,
      kLocationInRegisterY = 0x0C,
      kLocationInRegisterZ = 0x0D,
      kLocationInRegisterI = 0x0E,
      kLocationInRegisterJ = 0x0F,
      kLocationOffsetByRegisterA = 0x10,
      kLocationOffsetByRegisterB = 0x11,
      kLocationOffsetByRegisterC = 0x12,
      kLocationOffsetByRegsiterX = 0x13,
      kLocationOffsetByRegisterY = 0x14,
      kLocationOffsetByRegisterZ = 0x15,
      kLocationOffsetByRegisterI = 0x16,
      kLocationOffsetByRegisterJ = 0x17,
      kPop = 0x18,
      kPeek = 0x19,
      kPush = 0x1A,
      kStackPointer = 0x1B,
      kProgramCounter = 0x1C,
      kExtra = 0x1D,
      kLocation = 0x1E,
      kLiteral = 0x1F,
      kNegative1 = 0x20,
      k0 = 0x21,
      k1 = 0x22,
      k2 = 0x23,
      k3 = 0x24,
      k4 = 0x25,
      k5 = 0x26,
      k6 = 0x27,
      k7 = 0x28,
      k8 = 0x29,
      k9 = 0x2A,
      k10 = 0x2B,
      k11 = 0x2C,
      k12 = 0x2D,
      k13 = 0x2E,
      k14 = 0x2F,
      k15 = 0x30,
      k16 = 0x31,
      k17 = 0x32,
      k18 = 0x33,
      k19 = 0x34,
      k20 = 0x35,
      k21 = 0x36,
      k22 = 0x37,
      k23 = 0x38,
      k24 = 0x39,
      k25 = 0x3A,
      k26 = 0x3B,
      k27 = 0x3C,
      k28 = 0x3D,
      k29 = 0x3E,
      k30 = 0x3F
    };

  public:
    static Word Noop();
    static Word Instruct(const BasicOpcode basic_opcode,
        const Operand operand_a, const Operand operand_b);
    static Word Instruct(
        const AdvancedOpcode advanced_opcode, const Operand operand_a);

    Dcpu();
    virtual ~Dcpu() {}

    Word *address(const Word address_value);
    const Word *address(const Word address_value) const;

    Word *memory_begin();
    const Word *memory_begin() const;
    Word *memory_end();
    const Word *memory_end() const;

    Word &register_a();
    Word register_a() const;

    Word &register_b();
    Word register_b() const;

    Word &register_c();
    Word register_c() const;

    Word &register_x();
    Word register_x() const;

    Word &register_y();
    Word register_y() const;

    Word &register_z();
    Word register_z() const;

    Word &register_i();
    Word register_i() const;

    Word &register_j();
    Word register_j() const;

    Word &program_counter();
    Word program_counter() const;

    Word &stack_pointer();
    Word stack_pointer() const;

    Word &extra();
    Word extra() const;

    Word &interrupt_address();
    Word interrupt_address() const;

    void ExecuteInstruction(const bool skip = false);
    void ExecuteInstructions(const unsigned long int count);

    void Reset();

  private:
    Word *register_address(const Word register_index);
    Word register_value(const Word register_index);

    Word *GetOperandAddressOrLiteral(const Operand operand, Word &literal);
    void MaybeAssignResult(Word *const slot, const unsigned int result);

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
    Word extra_;
    Word interrupt_address_;
};

#endif  // DCPU_DCPU_H_
