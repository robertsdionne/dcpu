use std::mem;

const BASIC_OPCODE_MASK: u16 = 0x001f;
const BASIC_VALUE_MASK_A: u16 = 0xfc00;
const BASIC_SMALL_VALUE_MASK_A: u16 = 0x7c00;
const BASIC_VALUE_MASK_B: u16 = 0x03e0;
const BASIC_VALUE_SHIFT_A: u16 = 0x000a;
const BASIC_VALUE_SHIFT_B: u16 = 0x0005;
const SPECIAL_OPCODE_MASK: u16 = BASIC_VALUE_MASK_B;
const SPECIAL_OPCODE_SHIFT: u16 = BASIC_VALUE_SHIFT_B;
const DEBUG_OPCODE_MASK: u16 = BASIC_VALUE_MASK_A;
const DEBUG_OPCODE_SHIFT: u16 = BASIC_VALUE_SHIFT_A;

#[derive(Debug)]
pub enum Instruction {
    Basic(BasicOpcode, OperandB, OperandA),
    Special(SpecialOpcode, OperandA),
    Debug(DebugOpcode),
}

impl From<u16> for Instruction {
    fn from(instruction: u16) -> Self {
        let basic_opcode = BasicOpcode::from(instruction);
        let special_opcode = SpecialOpcode::from(instruction);
        let debug_opcode = DebugOpcode::from(instruction);

        if basic_opcode != BasicOpcode::Reserved {
            let operand_a = OperandA::from(instruction);
            let operand_b = OperandB::from(instruction);
            Instruction::Basic(basic_opcode, operand_b, operand_a)
        } else if special_opcode != SpecialOpcode::Reserved {
            let operand_a = OperandA::from(instruction);
            Instruction::Special(special_opcode, operand_a)
        } else {
            Instruction::Debug(debug_opcode)
        }
    }
}

#[allow(dead_code)]
#[derive(Debug, PartialEq)]
#[repr(u16)]
pub enum BasicOpcode {
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

impl From<u16> for BasicOpcode {
    fn from(instruction: u16) -> Self {
        unsafe { mem::transmute(instruction & BASIC_OPCODE_MASK) }
    }
}

#[allow(dead_code)]
#[derive(Debug, PartialEq)]
#[repr(u16)]
pub enum SpecialOpcode {
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

impl From<u16> for SpecialOpcode {
    fn from(instruction: u16) -> Self {
        unsafe { mem::transmute((instruction & SPECIAL_OPCODE_MASK) >> SPECIAL_OPCODE_SHIFT) }
    }
}

#[allow(dead_code)]
#[derive(Debug, PartialEq)]
#[repr(u16)]
pub enum DebugOpcode {
    Noop = 0x0,
    Alert = 0x1,
    DumpState = 0x2,
}

impl From<u16> for DebugOpcode {
    fn from(instruction: u16) -> Self {
        unsafe { mem::transmute((instruction & DEBUG_OPCODE_MASK) >> DEBUG_OPCODE_SHIFT) }
    }
}

#[derive(Debug, PartialEq)]
pub enum OperandA {
    LeftValue(OperandB),
    SmallLiteral(i8), // 32 values: -1 to 30
}

impl From<u16> for OperandA {
    fn from(instruction: u16) -> Self {
        match instruction & 0x8000 {
            0 => OperandA::LeftValue(OperandB::from(instruction >> 5)),
            _ => OperandA::SmallLiteral(
                ((instruction & BASIC_SMALL_VALUE_MASK_A) >> BASIC_VALUE_SHIFT_A) as i8 - 1
            ),
        }
    }
}

#[allow(dead_code)]
#[derive(Copy, Clone, Debug, PartialEq)]
#[repr(u16)]
pub enum OperandB {
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

impl From<u16> for OperandB {
    fn from(instruction: u16) -> Self {
        unsafe { mem::transmute((instruction & BASIC_VALUE_MASK_B) >> BASIC_VALUE_SHIFT_B) }
    }
}
