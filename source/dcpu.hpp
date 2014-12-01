#ifndef DCPU_DCPU_HPP_
#define DCPU_DCPU_HPP_

#include <vector>

#include "generated/program.pb.h"

namespace dcpu {

  class Hardware;

  using Word = unsigned short;
  using SignedWord = signed short;

  enum class BasicOpcode {
    kBasicReserved = 0x00,
    kSet = proto::Opcode_Basic_set,
    kAdd = proto::Opcode_Basic_add,
    kSubtract = proto::Opcode_Basic_sub,
    kMultiply = proto::Opcode_Basic_mul,
    kMultiplySigned = proto::Opcode_Basic_mli,
    kDivide = proto::Opcode_Basic_div,
    kDivideSigned = proto::Opcode_Basic_dvi,
    kModulo = proto::Opcode_Basic_mod,
    kModuloSigned = proto::Opcode_Basic_mdi,
    kBinaryAnd = proto::Opcode_Basic_and_,
    kBinaryOr = proto::Opcode_Basic_bor,
    kBinaryExclusiveOr = proto::Opcode_Basic_xor_,
    kShiftRight = proto::Opcode_Basic_shr,
    kArithmeticShiftRight = proto::Opcode_Basic_asr,
    kShiftLeft = proto::Opcode_Basic_shl,
    kIfBitSet = proto::Opcode_Basic_ifb,
    kIfClear = proto::Opcode_Basic_ifc,
    kIfEqual = proto::Opcode_Basic_ife,
    kIfNotEqual = proto::Opcode_Basic_ifn,
    kIfGreaterThan = proto::Opcode_Basic_ifg,
    kIfAbove = proto::Opcode_Basic_ifa,
    kIfLessThan = proto::Opcode_Basic_ifl,
    kIfUnder = proto::Opcode_Basic_ifu,
    kAddWithCarry = proto::Opcode_Basic_adx,
    kSubtractWithCarry = proto::Opcode_Basic_sbx,
    kSetThenIncrement = proto::Opcode_Basic_sti,
    kSetThenDecrement = proto::Opcode_Basic_std
  };

  enum class AdvancedOpcode {
    kAdvancedReserved = 0x00,
    kJumpSubRoutine = proto::Opcode_Advanced_jsr,
    kInterruptTrigger = proto::Opcode_Advanced_int_,
    kInterruptAddressGet = proto::Opcode_Advanced_iag,
    kInterruptAddressSet = proto::Opcode_Advanced_ias,
    kReturnFromInterrupt = proto::Opcode_Advanced_rfi,
    kInterruptAddToQueue = proto::Opcode_Advanced_iaq,
    kHardwareNumberConnected = proto::Opcode_Advanced_hwn,
    kHardwareQuery = proto::Opcode_Advanced_hwq,
    kHardwareInterrupt = proto::Opcode_Advanced_hwi
  };

  enum class Operand {
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
    kPush = 0x18,
    kPop = 0x18,
    kPushPop = 0x18,
    kPeek = 0x19,
    kPick = 0x1A,
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

  /**
   * Memory landmarks.
   */
  static constexpr unsigned int kMemorySize = 0x10000;

  /**
   * Opcodes.
   */
  static constexpr Word kBasicOpcodeMask = 0x001F;
  static constexpr Word kAdvancedOpcodeMask = 0x03E0;
  static constexpr Word kAdvancedOpcodeShift = 0x5;

  /**
   * Operands.
   */
  static constexpr Word kBasicOperandMaskA = 0xFC00;
  static constexpr Word kBasicOperandMaskB = 0x03E0;
  static constexpr Word kBasicOperandShiftA = 0xA;
  static constexpr Word kBasicOperandShiftB = 0x5;
  static constexpr Word kAdvancedOperandMaskA = 0xFC00;
  static constexpr Word kAdvancedOperandShiftA = 0xA;

  Word Noop();

  Word Instruct(const BasicOpcode basic_opcode, const Operand operand_b, const Operand operand_a);

  Word Instruct(const AdvancedOpcode advanced_opcode, const Operand operand_a);

  class Dcpu {
  public:
    Dcpu() = default;

    virtual ~Dcpu() = default;

    Word *address(const Word address_value);

    const Word *address(const Word address_value) const;

    Word *memory_begin();

    const Word *memory_begin() const;

    Word *memory_end();

    const Word *memory_end() const;

    void Connect(Hardware *hardware);

    void Interrupt(const Word message);

    void ExecuteInstruction(const bool skip = false);

    void ExecuteInstructions(const unsigned long int count);

    void Reset();

  private:
    Word *register_address(const Word register_index);
    Word register_value(const Word register_index);

    Word *GetOperandAddressOrLiteral(const Operand operand, const bool assignable, Word &literal);

    void MaybeAssignResult(Word *const slot, const unsigned int result);

  public:
    Word register_a = 0;
    Word register_b = 0;
    Word register_c = 0;
    Word register_x = 0;
    Word register_y = 0;
    Word register_z = 0;
    Word register_i = 0;
    Word register_j = 0;
    Word program_counter = 0;
    Word stack_pointer = 0;
    Word extra = 0;
    Word interrupt_address = 0;
    bool queue_interrupts = false;

  private:
    Word memory_[kMemorySize] = {};
    std::vector<Hardware *> hardware{};
  };

}  // namespace dcpu

#endif  // DCPU_DCPU_HPP_
