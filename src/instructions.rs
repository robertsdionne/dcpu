use std::mem;

const BASIC_OPCODE_MASK: u16 = 0x001f;
const BASIC_VALUE_MASK_A: u16 = 0xfc00;
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

impl Instruction {
    pub fn size(self) -> usize {
        match self {
            Self::Basic(_, operand_a, operand_b) => 1 + operand_a.size() + operand_b.size(),
            Self::Special(_, operand_a) => 1 + operand_a.size(),
            _ => 1,
        }
    }
}

impl From<&[u16]> for Instruction {
    fn from(value: &[u16]) -> Self {
        let instruction = InstructionValue::from(value[0]);
        let basic_opcode = BasicOpcode::from(instruction);
        let special_opcode = SpecialOpcode::from(instruction);

        if basic_opcode != BasicOpcode::Reserved {
            let operand_a = OperandA::from((instruction, value[1]));
            let operand_b = OperandB::from((instruction, value[1 + operand_a.size()]));
            Instruction::Basic(basic_opcode, operand_b, operand_a)
        } else if special_opcode != SpecialOpcode::Reserved {
            let operand_a = OperandA::from((instruction, value[1]));
            Instruction::Special(special_opcode, operand_a)
        } else {
            Instruction::Debug(DebugOpcode::from(instruction))
        }
    }
}

impl Into<Vec<u16>> for Instruction {
    fn into(self) -> Vec<u16> {
        match self {
            Instruction::Basic(basic_opcode, operand_b, operand_a) => {
                let payload_a = operand_a.payload();
                let payload_b = operand_b.payload();
                let basic_opcode: InstructionValue = basic_opcode.into();
                let operand_a: InstructionValue = operand_a.into();
                let operand_b: InstructionValue = operand_b.into();
                let basic_opcode: u16 = basic_opcode.into();
                let operand_a: u16 = operand_a.into();
                let operand_b: u16 = operand_b.into();
                let mut result = vec![basic_opcode | operand_b | operand_a];
                if let Some(payload) = payload_a {
                    result.push(payload);
                }
                if let Some(payload) = payload_b {
                    result.push(payload);
                }
                result
            }
            Instruction::Special(special_opcode, operand_a) => {
                let payload_a = operand_a.payload();
                let special_opcode: InstructionValue = special_opcode.into();
                let operand_a: InstructionValue = operand_a.into();
                let special_opcode: u16 = special_opcode.into();
                let operand_a: u16 = operand_a.into();
                let mut result = vec![special_opcode | operand_a];
                if let Some(payload) = payload_a {
                    result.push(payload);
                }
                result
            }
            Instruction::Debug(debug_opcode) => {
                let debug_opcode: InstructionValue = debug_opcode.into();
                vec![debug_opcode.into()]
            }
        }
    }
}

impl From<u16> for Instruction {
    fn from(instruction: u16) -> Self {
        let instruction = InstructionValue::from(instruction);
        let basic_opcode = BasicOpcode::from(instruction);
        let special_opcode = SpecialOpcode::from(instruction);

        if basic_opcode != BasicOpcode::Reserved {
            Instruction::Basic(
                basic_opcode,
                OperandB::from(instruction),
                OperandA::from(instruction),
            )
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
                basic_opcode | operand_b | operand_a
            }
            Instruction::Special(special_opcode, operand_a) => {
                let special_opcode: InstructionValue = special_opcode.into();
                let operand_a: InstructionValue = operand_a.into();
                let special_opcode: u16 = special_opcode.into();
                let operand_a: u16 = operand_a.into();
                special_opcode | operand_a
            }
            Instruction::Debug(debug_opcode) => {
                let debug_opcode: InstructionValue = debug_opcode.into();
                debug_opcode.into()
            }
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

impl OperandA {
    fn size(self) -> usize {
        match self {
            OperandA::LeftValue(operand_b) => operand_b.size(),
            _ => 0,
        }
    }

    fn payload(self) -> Option<u16> {
        match self {
            OperandA::LeftValue(operand_b) => operand_b.payload(),
            _ => None,
        }
    }

    fn with_payload(self, payload: u16) -> OperandA {
        match self {
            OperandA::LeftValue(operand_b) => OperandA::LeftValue(operand_b.with_payload(payload)),
            _ => self,
        }
    }
}

impl From<(InstructionValue, u16)> for OperandA {
    fn from(value: (InstructionValue, u16)) -> Self {
        let (instruction, payload) = value;
        Self::from(instruction).with_payload(payload)
    }
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
            OperandA::SmallLiteral(literal) => 0x20 | ((literal + 1) as u16 & BASIC_OPCODE_MASK),
        }
    }
}

impl From<u16> for OperandA {
    fn from(value: u16) -> Self {
        match value & 0x20 {
            0 => OperandA::LeftValue(OperandB::from(value)),
            _ => OperandA::SmallLiteral((value & BASIC_OPCODE_MASK) as i8 - 1),
        }
    }
}

#[derive(Copy, Clone, Debug, PartialEq)]
pub enum OperandB {
    Register(Register),
    LocationInRegister(Register),
    LocationOffsetByRegister(Register, u16),
    PushOrPop,
    Peek,
    Pick(u16),
    StackPointer,
    ProgramCounter,
    Extra,
    Location(u16),
    Literal(u16),
}

impl OperandB {
    fn size(self) -> usize {
        match self {
            OperandB::LocationOffsetByRegister(_, _) => 1,
            OperandB::Pick(_) => 1,
            OperandB::Location(_) => 1,
            OperandB::Literal(_) => 1,
            _ => 0,
        }
    }

    fn payload(self) -> Option<u16> {
        match self {
            OperandB::LocationOffsetByRegister(_, location) | OperandB::Location(location) => {
                Some(location)
            }
            OperandB::Pick(offset) => Some(offset),
            OperandB::Literal(literal) => Some(literal),
            _ => None,
        }
    }

    fn with_payload(self, payload: u16) -> OperandB {
        match self {
            OperandB::LocationOffsetByRegister(register, _) => {
                OperandB::LocationOffsetByRegister(register, payload)
            }
            OperandB::Pick(_) => OperandB::Pick(payload),
            OperandB::Location(_) => OperandB::Location(payload),
            OperandB::Literal(_) => OperandB::Literal(payload),
            _ => self,
        }
    }
}

impl From<(InstructionValue, u16)> for OperandB {
    fn from(value: (InstructionValue, u16)) -> Self {
        let (instruction, payload) = value;
        Self::from(instruction).with_payload(payload)
    }
}

impl From<InstructionValue> for OperandB {
    fn from(value: InstructionValue) -> Self {
        OperandB::from((value.0 & BASIC_VALUE_MASK_B) >> BASIC_VALUE_SHIFT_B)
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
        OperandB::from(operand_value)
    }
}

impl Into<u16> for OperandB {
    fn into(self) -> u16 {
        let operand_value: OperandValue = self.into();
        operand_value as u16
    }
}

impl From<OperandValue> for OperandB {
    fn from(value: OperandValue) -> Self {
        match value {
            OperandValue::RegisterA => OperandB::Register(Register::A),
            OperandValue::RegisterB => OperandB::Register(Register::B),
            OperandValue::RegisterC => OperandB::Register(Register::C),
            OperandValue::RegisterX => OperandB::Register(Register::X),
            OperandValue::RegisterY => OperandB::Register(Register::Y),
            OperandValue::RegisterZ => OperandB::Register(Register::Z),
            OperandValue::RegisterI => OperandB::Register(Register::I),
            OperandValue::RegisterJ => OperandB::Register(Register::J),
            OperandValue::LocationInRegisterA => OperandB::LocationInRegister(Register::A),
            OperandValue::LocationInRegisterB => OperandB::LocationInRegister(Register::B),
            OperandValue::LocationInRegisterC => OperandB::LocationInRegister(Register::C),
            OperandValue::LocationInRegisterX => OperandB::LocationInRegister(Register::X),
            OperandValue::LocationInRegisterY => OperandB::LocationInRegister(Register::Y),
            OperandValue::LocationInRegisterZ => OperandB::LocationInRegister(Register::Z),
            OperandValue::LocationInRegisterI => OperandB::LocationInRegister(Register::I),
            OperandValue::LocationInRegisterJ => OperandB::LocationInRegister(Register::J),
            OperandValue::LocationOffsetByRegisterA => {
                OperandB::LocationOffsetByRegister(Register::A, 0)
            }
            OperandValue::LocationOffsetByRegisterB => {
                OperandB::LocationOffsetByRegister(Register::B, 0)
            }
            OperandValue::LocationOffsetByRegisterC => {
                OperandB::LocationOffsetByRegister(Register::C, 0)
            }
            OperandValue::LocationOffsetByRegisterX => {
                OperandB::LocationOffsetByRegister(Register::X, 0)
            }
            OperandValue::LocationOffsetByRegisterY => {
                OperandB::LocationOffsetByRegister(Register::Y, 0)
            }
            OperandValue::LocationOffsetByRegisterZ => {
                OperandB::LocationOffsetByRegister(Register::Z, 0)
            }
            OperandValue::LocationOffsetByRegisterI => {
                OperandB::LocationOffsetByRegister(Register::I, 0)
            }
            OperandValue::LocationOffsetByRegisterJ => {
                OperandB::LocationOffsetByRegister(Register::J, 0)
            }
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

impl Into<OperandValue> for OperandB {
    fn into(self) -> OperandValue {
        match self {
            OperandB::Register(Register::A) => OperandValue::RegisterA,
            OperandB::Register(Register::B) => OperandValue::RegisterB,
            OperandB::Register(Register::C) => OperandValue::RegisterC,
            OperandB::Register(Register::X) => OperandValue::RegisterX,
            OperandB::Register(Register::Y) => OperandValue::RegisterY,
            OperandB::Register(Register::Z) => OperandValue::RegisterZ,
            OperandB::Register(Register::I) => OperandValue::RegisterI,
            OperandB::Register(Register::J) => OperandValue::RegisterJ,
            OperandB::LocationInRegister(Register::A) => OperandValue::LocationInRegisterA,
            OperandB::LocationInRegister(Register::B) => OperandValue::LocationInRegisterB,
            OperandB::LocationInRegister(Register::C) => OperandValue::LocationInRegisterC,
            OperandB::LocationInRegister(Register::X) => OperandValue::LocationInRegisterX,
            OperandB::LocationInRegister(Register::Y) => OperandValue::LocationInRegisterY,
            OperandB::LocationInRegister(Register::Z) => OperandValue::LocationInRegisterZ,
            OperandB::LocationInRegister(Register::I) => OperandValue::LocationInRegisterI,
            OperandB::LocationInRegister(Register::J) => OperandValue::LocationInRegisterJ,
            OperandB::LocationOffsetByRegister(Register::A, _) => {
                OperandValue::LocationOffsetByRegisterA
            }
            OperandB::LocationOffsetByRegister(Register::B, _) => {
                OperandValue::LocationOffsetByRegisterB
            }
            OperandB::LocationOffsetByRegister(Register::C, _) => {
                OperandValue::LocationOffsetByRegisterC
            }
            OperandB::LocationOffsetByRegister(Register::X, _) => {
                OperandValue::LocationOffsetByRegisterX
            }
            OperandB::LocationOffsetByRegister(Register::Y, _) => {
                OperandValue::LocationOffsetByRegisterY
            }
            OperandB::LocationOffsetByRegister(Register::Z, _) => {
                OperandValue::LocationOffsetByRegisterZ
            }
            OperandB::LocationOffsetByRegister(Register::I, _) => {
                OperandValue::LocationOffsetByRegisterI
            }
            OperandB::LocationOffsetByRegister(Register::J, _) => {
                OperandValue::LocationOffsetByRegisterJ
            }
            OperandB::PushOrPop => OperandValue::PushOrPop,
            OperandB::Peek => OperandValue::Peek,
            OperandB::Pick(_) => OperandValue::Pick,
            OperandB::StackPointer => OperandValue::StackPointer,
            OperandB::ProgramCounter => OperandValue::ProgramCounter,
            OperandB::Extra => OperandValue::Extra,
            OperandB::Location(_) => OperandValue::Location,
            OperandB::Literal(_) => OperandValue::Literal,
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

#[allow(dead_code)]
#[derive(Copy, Clone, Debug, PartialEq)]
#[repr(u16)]
pub enum Register {
    A,
    B,
    C,
    X,
    Y,
    Z,
    I,
    J,
}

impl From<u16> for Register {
    fn from(value: u16) -> Self {
        unsafe { mem::transmute(value & 0x07) }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn roundtrip() {
        let instruction = Instruction::Basic(
            BasicOpcode::Add,
            OperandB::Extra,
            OperandA::LeftValue(OperandB::LocationOffsetByRegister(Register::A, 0)),
        );
        let code: u16 = instruction.clone().into();
        assert_eq!(instruction, Instruction::from(code));

        let instruction = Instruction::Basic(
            BasicOpcode::IfEqual,
            OperandB::Literal(0),
            OperandA::SmallLiteral(-1),
        );
        let code: u16 = instruction.clone().into();
        assert_eq!(instruction, Instruction::from(code));

        let instruction = Instruction::Special(
            SpecialOpcode::HardwareNumberConnected,
            OperandA::LeftValue(OperandB::LocationOffsetByRegister(Register::A, 0)),
        );
        let code: u16 = instruction.clone().into();
        assert_eq!(instruction, Instruction::from(code));

        let instruction =
            Instruction::Special(SpecialOpcode::JumpSubroutine, OperandA::SmallLiteral(30));
        let code: u16 = instruction.clone().into();
        assert_eq!(instruction, Instruction::from(code));

        let instruction = Instruction::Debug(DebugOpcode::DumpState);
        let code: u16 = instruction.clone().into();
        assert_eq!(instruction, Instruction::from(code));
    }
}
