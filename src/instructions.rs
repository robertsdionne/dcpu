
enum Instruction {
    Basic(BasicOpcode, OperandB, OperandA),
    Special(SpecialOpcode, OperandA),
    Debug(DebugOpcode),
}

enum BasicOpcode {
    Reserved = 0x0,
    Set = 0x1,
    Add = 0x2,
    Subtract = 0x3,
    Multiply = 0x4,
    MultiplySigned = 0x5,
    Divide = 0x6,
    DivideSigned = 0x7,
    Modulo = 0x8,
    ModuloSigned = 0x9,
    BinaryAnd = 0xa,
    BinaryOr = 0xb,
    BinaryExclusiveOr = 0xc,
    ShiftRight = 0xd,
    ArithmeticShiftRight = 0xe,
    ShiftLeft = 0xf,
    IfBitSet = 0x10,
    IfClear = 0x11,
    IfEqual = 0x12,
    IfNotEqual = 0x13,
    IfGreaterThan = 0x14,
    IfAbove = 0x15,
    IfLessThan = 0x16,
    IfUnder = 0x17,
    AddWithCarry = 0x1a,
    SubtractWithCarry = 0x1b,
    SetThenIncrement = 0x1e,
    SetThenDecrement = 0x1f,
}

enum SpecialOpcode {
    Reserved = 0x0,
    JumpSubroutine = 0x1,
    InterruptTrigger = 0x8,
    InterruptAddressGet = 0x9,
    InterruptAddressSet = 0xa,
    ReturnFromInterrupt = 0xb,
    InterruptAddToQueue = 0xc,
    HardwareNumberConnected = 0x10,
    HardwareQuery = 0x11,
    HardwareInterrupt = 0x12,
}

enum DebugOpcode {
    Noop = 0x0,
    Alert = 0x1,
    DumpState = 0x2,
}

enum OperandA {
    LeftValue(OperandB),
    SmallLiteral(u8), // 32 values: -1 to 30
}

enum OperandB {
    RegisterA,
    RegisterB,
    RegisterC,
    RegisterX,
    RegisterY,
    RegisterZ,
    RegisterI,
    RegisterJ,
    LocationInRegisterA,
    LocationInRegisterB,
    LocationInRegisterC,
    LocationInRegisterX,
    LocationInRegisterY,
    LocationInRegisterZ,
    LocationInRegisterI,
    LocationInRegisterJ,
    LocationOffsetByRegisterA,
    LocationOffsetByRegisterB,
    LocationOffsetByRegisterC,
    LocationOffsetByRegisterX,
    LocationOffsetByRegisterY,
    LocationOffsetByRegisterZ,
    LocationOffsetByRegisterI,
    LocationOffsetByRegisterJ,
    PushOrPop,
    Peek,
    Pick,
    StackPointer,
    ProgramCounter,
    Extra,
    Location,
    Literal,
}
