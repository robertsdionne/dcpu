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

#[derive(Copy, Clone, Debug, PartialEq)]
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

impl Into<u16> for Instruction {
    fn into(self) -> u16 {
        match self {
            Instruction::Basic(basic_opcode, operand_b, operand_a) => {
                let operand_a: u16 = operand_a.into();
                basic_opcode as u16 | (operand_b as u16) << BASIC_VALUE_SHIFT_B | operand_a << BASIC_VALUE_SHIFT_A
            }
            Instruction::Special(special_opcode, operand_a) => {
                let operand_a: u16 = operand_a.into();
                (special_opcode as u16) << SPECIAL_OPCODE_SHIFT | operand_a << BASIC_VALUE_SHIFT_A
            }
            Instruction::Debug(debug_opcode) => (debug_opcode as u16) << DEBUG_OPCODE_SHIFT,
        }
    }
}

#[allow(dead_code)]
#[derive(Copy, Clone, Debug, PartialEq)]
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
    _Unused18 = 0x18,
    _Unused19 = 0x19,
    AddWithCarry = 0x1a,
    SubtractWithCarry = 0x1b,
    _Unused1c = 0x1c,
    _Unused1d = 0x1d,
    SetThenIncrement = 0x1e,
    SetThenDecrement = 0x1f,
}

impl From<u16> for BasicOpcode {
    fn from(instruction: u16) -> Self {
        unsafe { mem::transmute(instruction & BASIC_OPCODE_MASK) }
    }
}

#[allow(dead_code)]
#[derive(Copy, Clone, Debug, PartialEq)]
#[repr(u16)]
pub enum SpecialOpcode {
    Reserved = 0x0,
    JumpSubroutine = 0x1,
    _Unused02 = 0x2,
    _Unused03 = 0x3,
    _Unused04 = 0x4,
    _Unused05 = 0x5,
    _Unused06 = 0x6,
    _Unused07 = 0x7,
    InterruptTrigger = 0x8,
    InterruptAddressGet = 0x9,
    InterruptAddressSet = 0xa,
    ReturnFromInterrupt = 0xb,
    InterruptAddToQueue = 0xc,
    _Unused0d = 0xd,
    _Unused0e = 0xe,
    HardwareNumberConnected = 0x10,
    HardwareQuery = 0x11,
    HardwareInterrupt = 0x12,
    _Unused13 = 0x13,
    _Unused14 = 0x14,
    _Unused15 = 0x15,
    _Unused16 = 0x16,
    _Unused17 = 0x17,
    _Unused18 = 0x18,
    _Unused19 = 0x19,
    _Unused1a = 0x1a,
    _Unused1b = 0x1b,
    _Unused1c = 0x1c,
    _Unused1d = 0x1d,
    _Unused1e = 0x1e,
    _Unused1f = 0x1f,
}

impl From<u16> for SpecialOpcode {
    fn from(instruction: u16) -> Self {
        unsafe { mem::transmute((instruction & SPECIAL_OPCODE_MASK) >> SPECIAL_OPCODE_SHIFT) }
    }
}

#[allow(dead_code)]
#[derive(Copy, Clone, Debug, PartialEq)]
#[repr(u16)]
pub enum DebugOpcode {
    Noop = 0x0,
    Alert = 0x1,
    DumpState = 0x2,
    _Unused03 = 0x3,
    _Unused04 = 0x4,
    _Unused05 = 0x5,
    _Unused06 = 0x6,
    _Unused07 = 0x7,
    _Unused08 = 0x8,
    _Unused09 = 0x9,
    _Unused0a = 0xa,
    _Unused0b = 0xb,
    _Unused0c = 0xc,
    _Unused0d = 0xd,
    _Unused0e = 0xe,
    _Unused0f = 0xf,
    _Unused10 = 0x10,
    _Unused11 = 0x11,
    _Unused12 = 0x12,
    _Unused13 = 0x13,
    _Unused14 = 0x14,
    _Unused15 = 0x15,
    _Unused16 = 0x16,
    _Unused17 = 0x17,
    _Unused18 = 0x18,
    _Unused19 = 0x19,
    _Unused1a = 0x1a,
    _Unused1b = 0x1b,
    _Unused1c = 0x1c,
    _Unused1d = 0x1d,
    _Unused1e = 0x1e,
    _Unused1f = 0x1f,
    _Unused20 = 0x20,
    _Unused21 = 0x21,
    _Unused22 = 0x22,
    _Unused23 = 0x23,
    _Unused24 = 0x24,
    _Unused25 = 0x25,
    _Unused26 = 0x26,
    _Unused27 = 0x27,
    _Unused28 = 0x28,
    _Unused29 = 0x29,
    _Unused2a = 0x2a,
    _Unused2b = 0x2b,
    _Unused2c = 0x2c,
    _Unused2d = 0x2d,
    _Unused2e = 0x2e,
    _Unused2f = 0x2f,
    _Unused30 = 0x30,
    _Unused31 = 0x31,
    _Unused32 = 0x32,
    _Unused33 = 0x33,
    _Unused34 = 0x34,
    _Unused35 = 0x35,
    _Unused36 = 0x36,
    _Unused37 = 0x37,
    _Unused38 = 0x38,
    _Unused39 = 0x39,
    _Unused3a = 0x3a,
    _Unused3b = 0x3b,
    _Unused3c = 0x3c,
    _Unused3d = 0x3d,
    _Unused3e = 0x3e,
    _Unused3f = 0x3f,
}

impl From<u16> for DebugOpcode {
    fn from(instruction: u16) -> Self {
        unsafe { mem::transmute((instruction & DEBUG_OPCODE_MASK) >> DEBUG_OPCODE_SHIFT) }
    }
}

#[derive(Copy, Clone, Debug, PartialEq)]
pub enum OperandA {
    LeftValue(OperandB),
    SmallLiteral(i8), // 32 values: -1 to 30
}

impl From<u16> for OperandA {
    fn from(instruction: u16) -> Self {
        match instruction & 0x8000 {
            0 => OperandA::LeftValue(OperandB::from(instruction >> BASIC_VALUE_SHIFT_B)),
            _ => OperandA::SmallLiteral(
                ((instruction & BASIC_SMALL_VALUE_MASK_A) >> BASIC_VALUE_SHIFT_A) as i8 - 1
            ),
        }
    }
}

impl Into<u16> for OperandA {
    fn into(self) -> u16 {
        match self {
            OperandA::LeftValue(operand_b) => operand_b as u16,
            OperandA::SmallLiteral(literal) => 0x20 | ((literal + 1) as u16 & BASIC_OPCODE_MASK),
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

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn roundtrip() {
        let instruction = Instruction::Basic(BasicOpcode::Add, OperandB::Extra, OperandA::LeftValue(OperandB::LocationOffsetByRegisterA));
        let code: u16 = instruction.clone().into();
        assert_eq!(instruction, Instruction::from(code));

        let instruction = Instruction::Special(SpecialOpcode::HardwareNumberConnected, OperandA::LeftValue(OperandB::LocationOffsetByRegisterA));
        let code: u16 = instruction.clone().into();
        assert_eq!(instruction, Instruction::from(code));

        let instruction = Instruction::Debug(DebugOpcode::DumpState);
        let code: u16 = instruction.clone().into();
        assert_eq!(instruction, Instruction::from(code));
    }
}
