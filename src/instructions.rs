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
        let instruction = InstructionValue::from(instruction);
        let basic_opcode = BasicOpcode::from(instruction);
        let special_opcode = SpecialOpcode::from(instruction);

        if basic_opcode != BasicOpcode::Reserved {
            Instruction::Basic(basic_opcode, OperandB::from(instruction), OperandA::from(instruction))
        } else if special_opcode != SpecialOpcode::Reserved {
            Instruction::Special(special_opcode, OperandA::from(instruction))
        } else {
            Instruction::Debug(DebugOpcode::from(instruction))
        }
    }
}

impl Into<u16> for Instruction {
    fn into(self) -> u16 {
        match self {
            Instruction::Basic(basic_opcode, operand_b, operand_a) => {
                let basic_opcode: InstructionValue = basic_opcode.into();
                let operand_a: InstructionValue = operand_a.into();
                let operand_b: InstructionValue = operand_b.into();
                let basic_opcode: u16 = basic_opcode.into();
                let operand_a: u16 = operand_a.into();
                let operand_b: u16 = operand_b.into();
                basic_opcode
                    | operand_b
                    | operand_a
            }
            Instruction::Special(special_opcode, operand_a) => {
                let special_opcode: InstructionValue = special_opcode.into();
                let operand_a: InstructionValue = operand_a.into();
                let special_opcode: u16 = special_opcode.into();
                let operand_a: u16 = operand_a.into();
                special_opcode
                    | operand_a
            }
            Instruction::Debug(debug_opcode) => {
                let debug_opcode: InstructionValue = debug_opcode.into();
                debug_opcode.into()
            },
        }
    }
}

#[derive(Copy, Clone)]
struct InstructionValue(u16);

impl Into<u16> for InstructionValue {
    fn into(self) -> u16 {
        self.0
    }
}

impl From<u16> for InstructionValue {
    fn from(value: u16) -> Self {
        InstructionValue(value)
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

impl From<InstructionValue> for BasicOpcode {
    fn from(value: InstructionValue) -> Self {
        Self::from(value.0 & BASIC_OPCODE_MASK)
    }
}

impl Into<InstructionValue> for BasicOpcode {
    fn into(self) -> InstructionValue {
        InstructionValue::from(self as u16)
    }
}

impl From<u16> for BasicOpcode {
    fn from(value: u16) -> Self {
        unsafe { mem::transmute(value & 0x1f) }
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

impl From<InstructionValue> for SpecialOpcode {
    fn from(value: InstructionValue) -> Self {
        Self::from((value.0 & SPECIAL_OPCODE_MASK) >> SPECIAL_OPCODE_SHIFT)
    }
}

impl Into<InstructionValue> for SpecialOpcode {
    fn into(self) -> InstructionValue {
        InstructionValue::from((self as u16) << SPECIAL_OPCODE_SHIFT)
    }
}

impl From<u16> for SpecialOpcode {
    fn from(value: u16) -> Self {
        unsafe { mem::transmute(value & 0x1f) }
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

impl From<InstructionValue> for DebugOpcode {
    fn from(value: InstructionValue) -> Self {
        Self::from((value.0 & DEBUG_OPCODE_MASK) >> DEBUG_OPCODE_SHIFT)
    }
}

impl Into<InstructionValue> for DebugOpcode {
    fn into(self) -> InstructionValue {
        InstructionValue::from((self as u16) << DEBUG_OPCODE_SHIFT)
    }
}

impl From<u16> for DebugOpcode {
    fn from(value: u16) -> Self {
        unsafe { mem::transmute(value & 0x3f) }
    }
}

#[derive(Copy, Clone, Debug, PartialEq)]
pub enum OperandA {
    LeftValue(OperandB),
    SmallLiteral(i8), // 32 values: -1 to 30
}

impl From<InstructionValue> for OperandA {
    fn from(value: InstructionValue) -> Self {
        OperandA::from(value.0 >> BASIC_VALUE_SHIFT_A)
    }
}

impl Into<InstructionValue> for OperandA {
    fn into(self) -> InstructionValue {
        let value: u16 = self.into();
        InstructionValue(value << BASIC_VALUE_SHIFT_A)
    }
}

impl Into<u16> for OperandA {
    fn into(self) -> u16 {
        match self {
            OperandA::LeftValue(operand_b) => operand_b.into(),
            OperandA::SmallLiteral(literal) => 0x20
                | ((literal + 1) as u16 & BASIC_OPCODE_MASK),
        }
    }
}

impl From<u16> for OperandA {
    fn from(value: u16) -> Self {
        match value & 0x20 {
            0 => OperandA::LeftValue(OperandB::from(value)),
            _ => OperandA::SmallLiteral((value & BASIC_OPCODE_MASK) as i8 - 1)
        }
    }
}

impl From<[u16; 2]> for OperandA {
    fn from(value: [u16; 2]) -> Self {
        match value[0] & 0x20 {
            0 => OperandA::LeftValue(OperandB::from([value[0], 0])),
            _ => OperandA::SmallLiteral(
                (value[0] & BASIC_OPCODE_MASK) as i8 - 1
            ),
        }
    }
}

#[derive(Copy, Clone, Debug, PartialEq)]
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
    LocationOffsetByRegisterA(u16),
    LocationOffsetByRegisterB(u16),
    LocationOffsetByRegisterC(u16),
    LocationOffsetByRegisterX(u16),
    LocationOffsetByRegisterY(u16),
    LocationOffsetByRegisterZ(u16),
    LocationOffsetByRegisterI(u16),
    LocationOffsetByRegisterJ(u16),
    PushOrPop,
    Peek,
    Pick(u16),
    StackPointer,
    ProgramCounter,
    Extra,
    Location(u16),
    Literal(u16),
}

impl From<InstructionValue> for OperandB {
    fn from(value: InstructionValue) -> Self {
        OperandB::from([(value.0 & BASIC_VALUE_MASK_B) >> BASIC_VALUE_SHIFT_B, 0])
    }
}

impl Into<InstructionValue> for OperandB {
    fn into(self) -> InstructionValue {
        let value: u16 = self.into();
        InstructionValue(value << BASIC_VALUE_SHIFT_B)
    }
}

impl From<u16> for OperandB {
    fn from(value: u16) -> Self {
        let operand_value = OperandValue::from(value);
        match operand_value {
            OperandValue::RegisterA => OperandB::RegisterA,
            OperandValue::RegisterB => OperandB::RegisterB,
            OperandValue::RegisterC => OperandB::RegisterC,
            OperandValue::RegisterX => OperandB::RegisterX,
            OperandValue::RegisterY => OperandB::RegisterY,
            OperandValue::RegisterZ => OperandB::RegisterZ,
            OperandValue::RegisterI => OperandB::RegisterI,
            OperandValue::RegisterJ => OperandB::RegisterJ,
            OperandValue::LocationInRegisterA => OperandB::LocationInRegisterA,
            OperandValue::LocationInRegisterB => OperandB::LocationInRegisterB,
            OperandValue::LocationInRegisterC => OperandB::LocationInRegisterC,
            OperandValue::LocationInRegisterX => OperandB::LocationInRegisterX,
            OperandValue::LocationInRegisterY => OperandB::LocationInRegisterY,
            OperandValue::LocationInRegisterZ => OperandB::LocationInRegisterZ,
            OperandValue::LocationInRegisterI => OperandB::LocationInRegisterI,
            OperandValue::LocationInRegisterJ => OperandB::LocationInRegisterJ,
            OperandValue::LocationOffsetByRegisterA => OperandB::LocationOffsetByRegisterA(0),
            OperandValue::LocationOffsetByRegisterB => OperandB::LocationOffsetByRegisterB(0),
            OperandValue::LocationOffsetByRegisterC => OperandB::LocationOffsetByRegisterC(0),
            OperandValue::LocationOffsetByRegisterX => OperandB::LocationOffsetByRegisterX(0),
            OperandValue::LocationOffsetByRegisterY => OperandB::LocationOffsetByRegisterY(0),
            OperandValue::LocationOffsetByRegisterZ => OperandB::LocationOffsetByRegisterZ(0),
            OperandValue::LocationOffsetByRegisterI => OperandB::LocationOffsetByRegisterI(0),
            OperandValue::LocationOffsetByRegisterJ => OperandB::LocationOffsetByRegisterJ(0),
            OperandValue::PushOrPop => OperandB::PushOrPop,
            OperandValue::Peek => OperandB::Peek,
            OperandValue::Pick => OperandB::Pick(0),
            OperandValue::StackPointer => OperandB::StackPointer,
            OperandValue::ProgramCounter => OperandB::ProgramCounter,
            OperandValue::Extra => OperandB::Extra,
            OperandValue::Location => OperandB::Location(0),
            OperandValue::Literal => OperandB::Literal(0),
        }
    }
}

impl Into<u16> for OperandB {
    fn into(self) -> u16 {
        match self {
            OperandB::RegisterA => OperandValue::RegisterA as u16,
            OperandB::RegisterB => OperandValue::RegisterB as u16,
            OperandB::RegisterC => OperandValue::RegisterC as u16,
            OperandB::RegisterX => OperandValue::RegisterX as u16,
            OperandB::RegisterY => OperandValue::RegisterY as u16,
            OperandB::RegisterZ => OperandValue::RegisterZ as u16,
            OperandB::RegisterI => OperandValue::RegisterI as u16,
            OperandB::RegisterJ => OperandValue::RegisterJ as u16,
            OperandB::LocationInRegisterA => OperandValue::LocationInRegisterA as u16,
            OperandB::LocationInRegisterB => OperandValue::LocationInRegisterA as u16,
            OperandB::LocationInRegisterC => OperandValue::LocationInRegisterA as u16,
            OperandB::LocationInRegisterX => OperandValue::LocationInRegisterA as u16,
            OperandB::LocationInRegisterY => OperandValue::LocationInRegisterA as u16,
            OperandB::LocationInRegisterZ => OperandValue::LocationInRegisterA as u16,
            OperandB::LocationInRegisterI => OperandValue::LocationInRegisterA as u16,
            OperandB::LocationInRegisterJ => OperandValue::LocationInRegisterA as u16,
            OperandB::LocationOffsetByRegisterA(_) => OperandValue::LocationOffsetByRegisterA as u16,
            OperandB::LocationOffsetByRegisterB(_) => OperandValue::LocationOffsetByRegisterA as u16,
            OperandB::LocationOffsetByRegisterC(_) => OperandValue::LocationOffsetByRegisterA as u16,
            OperandB::LocationOffsetByRegisterX(_) => OperandValue::LocationOffsetByRegisterA as u16,
            OperandB::LocationOffsetByRegisterY(_) => OperandValue::LocationOffsetByRegisterA as u16,
            OperandB::LocationOffsetByRegisterZ(_) => OperandValue::LocationOffsetByRegisterA as u16,
            OperandB::LocationOffsetByRegisterI(_) => OperandValue::LocationOffsetByRegisterA as u16,
            OperandB::LocationOffsetByRegisterJ(_) => OperandValue::LocationOffsetByRegisterA as u16,
            OperandB::PushOrPop => OperandValue::PushOrPop as u16,
            OperandB::Peek => OperandValue::Peek as u16,
            OperandB::Pick(_) => OperandValue::Pick as u16,
            OperandB::StackPointer => OperandValue::StackPointer as u16,
            OperandB::ProgramCounter => OperandValue::ProgramCounter as u16,
            OperandB::Extra => OperandValue::Extra as u16,
            OperandB::Location(_) => OperandValue::Location as u16,
            OperandB::Literal(_) => OperandValue::Literal as u16,
        }
    }
}

impl From<[u16; 2]> for OperandB {
    fn from(value: [u16; 2]) -> Self {
        let operand_value = OperandValue::from(value[0]);
        match operand_value {
            OperandValue::RegisterA => OperandB::RegisterA,
            OperandValue::RegisterB => OperandB::RegisterB,
            OperandValue::RegisterC => OperandB::RegisterC,
            OperandValue::RegisterX => OperandB::RegisterX,
            OperandValue::RegisterY => OperandB::RegisterY,
            OperandValue::RegisterZ => OperandB::RegisterZ,
            OperandValue::RegisterI => OperandB::RegisterI,
            OperandValue::RegisterJ => OperandB::RegisterJ,
            OperandValue::LocationInRegisterA => OperandB::LocationInRegisterA,
            OperandValue::LocationInRegisterB => OperandB::LocationInRegisterB,
            OperandValue::LocationInRegisterC => OperandB::LocationInRegisterC,
            OperandValue::LocationInRegisterX => OperandB::LocationInRegisterX,
            OperandValue::LocationInRegisterY => OperandB::LocationInRegisterY,
            OperandValue::LocationInRegisterZ => OperandB::LocationInRegisterZ,
            OperandValue::LocationInRegisterI => OperandB::LocationInRegisterI,
            OperandValue::LocationInRegisterJ => OperandB::LocationInRegisterJ,
            OperandValue::LocationOffsetByRegisterA => OperandB::LocationOffsetByRegisterA(value[1]),
            OperandValue::LocationOffsetByRegisterB => OperandB::LocationOffsetByRegisterB(value[1]),
            OperandValue::LocationOffsetByRegisterC => OperandB::LocationOffsetByRegisterC(value[1]),
            OperandValue::LocationOffsetByRegisterX => OperandB::LocationOffsetByRegisterX(value[1]),
            OperandValue::LocationOffsetByRegisterY => OperandB::LocationOffsetByRegisterY(value[1]),
            OperandValue::LocationOffsetByRegisterZ => OperandB::LocationOffsetByRegisterZ(value[1]),
            OperandValue::LocationOffsetByRegisterI => OperandB::LocationOffsetByRegisterI(value[1]),
            OperandValue::LocationOffsetByRegisterJ => OperandB::LocationOffsetByRegisterJ(value[1]),
            OperandValue::PushOrPop => OperandB::PushOrPop,
            OperandValue::Peek => OperandB::Peek,
            OperandValue::Pick => OperandB::Pick(value[1]),
            OperandValue::StackPointer => OperandB::StackPointer,
            OperandValue::ProgramCounter => OperandB::ProgramCounter,
            OperandValue::Extra => OperandB::Extra,
            OperandValue::Location => OperandB::Location(value[1]),
            OperandValue::Literal => OperandB::Literal(value[1]),
        }
    }
}

impl Into<(u16, Option<u16>)> for OperandB {
    fn into(self) -> (u16, Option<u16>) {
        match self {
            OperandB::RegisterA => (OperandValue::RegisterA as u16, None),
            OperandB::RegisterB => (OperandValue::RegisterB as u16, None),
            OperandB::RegisterC => (OperandValue::RegisterC as u16, None),
            OperandB::RegisterX => (OperandValue::RegisterX as u16, None),
            OperandB::RegisterY => (OperandValue::RegisterY as u16, None),
            OperandB::RegisterZ => (OperandValue::RegisterZ as u16, None),
            OperandB::RegisterI => (OperandValue::RegisterI as u16, None),
            OperandB::RegisterJ => (OperandValue::RegisterJ as u16, None),
            OperandB::LocationInRegisterA => (OperandValue::LocationInRegisterA as u16, None),
            OperandB::LocationInRegisterB => (OperandValue::LocationInRegisterA as u16, None),
            OperandB::LocationInRegisterC => (OperandValue::LocationInRegisterA as u16, None),
            OperandB::LocationInRegisterX => (OperandValue::LocationInRegisterA as u16, None),
            OperandB::LocationInRegisterY => (OperandValue::LocationInRegisterA as u16, None),
            OperandB::LocationInRegisterZ => (OperandValue::LocationInRegisterA as u16, None),
            OperandB::LocationInRegisterI => (OperandValue::LocationInRegisterA as u16, None),
            OperandB::LocationInRegisterJ => (OperandValue::LocationInRegisterA as u16, None),
            OperandB::LocationOffsetByRegisterA(location) => (OperandValue::LocationOffsetByRegisterA as u16, Some(location)),
            OperandB::LocationOffsetByRegisterB(location) => (OperandValue::LocationOffsetByRegisterA as u16, Some(location)),
            OperandB::LocationOffsetByRegisterC(location) => (OperandValue::LocationOffsetByRegisterA as u16, Some(location)),
            OperandB::LocationOffsetByRegisterX(location) => (OperandValue::LocationOffsetByRegisterA as u16, Some(location)),
            OperandB::LocationOffsetByRegisterY(location) => (OperandValue::LocationOffsetByRegisterA as u16, Some(location)),
            OperandB::LocationOffsetByRegisterZ(location) => (OperandValue::LocationOffsetByRegisterA as u16, Some(location)),
            OperandB::LocationOffsetByRegisterI(location) => (OperandValue::LocationOffsetByRegisterA as u16, Some(location)),
            OperandB::LocationOffsetByRegisterJ(location) => (OperandValue::LocationOffsetByRegisterA as u16, Some(location)),
            OperandB::PushOrPop => (OperandValue::PushOrPop as u16, None),
            OperandB::Peek => (OperandValue::Peek as u16, None),
            OperandB::Pick(offset) => (OperandValue::Pick as u16, Some(offset)),
            OperandB::StackPointer => (OperandValue::StackPointer as u16, None),
            OperandB::ProgramCounter => (OperandValue::ProgramCounter as u16, None),
            OperandB::Extra => (OperandValue::Extra as u16, None),
            OperandB::Location(location) => (OperandValue::Location as u16, Some(location)),
            OperandB::Literal(literal) => (OperandValue::Literal as u16, Some(literal)),
        }
    }
}

#[allow(dead_code)]
#[repr(u16)]
enum OperandValue {
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

impl From<u16> for OperandValue {
    fn from(value: u16) -> Self {
        unsafe { mem::transmute(value & 0x1f) }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn roundtrip() {
        let instruction = Instruction::Basic(BasicOpcode::Add, OperandB::Extra, OperandA::LeftValue(OperandB::LocationOffsetByRegisterA(0)));
        let code: u16 = instruction.clone().into();
        assert_eq!(instruction, Instruction::from(code));

        let instruction = Instruction::Basic(BasicOpcode::IfEqual, OperandB::Literal(0), OperandA::SmallLiteral(-1));
        let code: u16 = instruction.clone().into();
        assert_eq!(instruction, Instruction::from(code));

        let instruction = Instruction::Special(SpecialOpcode::HardwareNumberConnected, OperandA::LeftValue(OperandB::LocationOffsetByRegisterA(0)));
        let code: u16 = instruction.clone().into();
        assert_eq!(instruction, Instruction::from(code));

        let instruction = Instruction::Special(SpecialOpcode::JumpSubroutine, OperandA::SmallLiteral(30));
        let code: u16 = instruction.clone().into();
        assert_eq!(instruction, Instruction::from(code));

        let instruction = Instruction::Debug(DebugOpcode::DumpState);
        let code: u16 = instruction.clone().into();
        assert_eq!(instruction, Instruction::from(code));
    }
}
