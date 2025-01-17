use std::{error, fs};
use chumsky::prelude::*;
use crate::instructions;

pub fn assemble(program: &str) -> Result<(), Box<dyn error::Error>> {
    let source = fs::read_to_string(program)?;
    let program = Program::parser().parse_recovery_verbose(source);

    println!("{:#?}", program);
    Ok(())
}

#[derive(Debug)]
struct Program(Vec<Statement>);

impl Program {
    /// program
    ///     : statement+ EOF
    ///     ;
    fn parser() -> impl Parser<char, Self, Error = Simple<char>> {
        Statement::parser().padded()
            .repeated().at_least(1)
            .then_ignore(end())
            .map(|statements| Program(statements))
    }
}

#[derive(Debug)]
enum Statement {
    LabelDefinition(String),
    Instruction(InstructionWithLabels),
    DataSection(Data),
}

impl Statement {
    /// statement
    ///     : labelDefinition
    ///     | instruction
    ///     | dataSection
    ///     ;
    fn parser() -> impl Parser<char, Self, Error = Simple<char>> {
        Self::label_definition_parser()
            .or(Self::instruction_parser())
            .or(Self::data_section_parser())
    }

    /// labelDefinition
    ///     : ':' IDENTIFIER
    ///     ;
    fn label_definition_parser() -> impl Parser<char, Self, Error = Simple<char>> {
        just(':')
            .ignore_then(text::ident())
            .map(|label| Statement::LabelDefinition(label))
    }

    fn instruction_parser() -> impl Parser<char, Self, Error = Simple<char>> {
        InstructionWithLabels::parser()
            .map(|instruction| Statement::Instruction(instruction))
    }

    /// dataSection
    ///     : ('.dat' | '.DAT') data
    ///     ;
    fn data_section_parser() -> impl Parser<char, Self, Error = Simple<char>> {
        just(".DAT")
            .or(just(".dat"))
            .padded()
            .ignore_then(Data::parser())
            .map(|data| Statement::DataSection(data))
    }
}

#[derive(Debug)]
enum InstructionWithLabels {
    Basic(instructions::BasicOpcode, OperandBWithLabel, OperandAWithLabel),
    Special(instructions::SpecialOpcode, OperandAWithLabel),
    Debug(instructions::DebugOpcode),
}

impl InstructionWithLabels {
    /// instruction
    ///     : basic
    ///     : special
    ///     : debug
    ///     ;
    fn parser() -> impl Parser<char, Self, Error = Simple<char>> {
        Self::basic_parser()
            .or(Self::special_parser())
            .or(Self::debug_parser())
    }

    /// basic
    ///     : basicOpcode operandB ',' operandA
    ///     ;
    fn basic_parser() -> impl Parser<char, Self, Error = Simple<char>> {
        instructions::BasicOpcode::parser().padded()
            .then(OperandBWithLabel::parser()).padded()
            .then_ignore(just(',')).padded()
            .then(OperandAWithLabel::parser())
            .map(|((basic_opcode, operand_b), operand_a)|
                InstructionWithLabels::Basic(basic_opcode, operand_b, operand_a))
    }

    /// special
    ///     : specialOpcode operandA
    ///     ;
    fn special_parser() -> impl Parser<char, Self, Error = Simple<char>> {
        instructions::SpecialOpcode::parser().padded()
            .then(OperandAWithLabel::parser())
            .map(|(special_opcode, operand_a)|
                InstructionWithLabels::Special(special_opcode, operand_a))
    }

    /// debug
    ///     : debugOpcode
    ///     ;
    fn debug_parser() -> impl Parser<char, Self, Error = Simple<char>> {
        instructions::DebugOpcode::parser().padded()
            .map(|debug_opcode| InstructionWithLabels::Debug(debug_opcode))
    }
}

impl instructions::BasicOpcode {
    /// basicOpcode
    ///     : SET
    ///     | ADD
    ///     | SUB
    ///     | MUL
    ///     | MLI
    ///     | DIV
    ///     | DVI
    ///     | AND
    ///     | BOR
    ///     | XOR
    ///     | SHR
    ///     | ASR
    ///     | SHL
    ///     | IFB
    ///     | IFC
    ///     | IFE
    ///     | IFN
    ///     | IFG
    ///     | IFA
    ///     | IFL
    ///     | IFU
    ///     | ADX
    ///     | SBX
    ///     | STI
    ///     | STD
    ///     ;
    fn parser() -> impl Parser<char, Self, Error = Simple<char>> {
        use instructions::BasicOpcode;

        choice([
            just("SET").or(just("set")).to(BasicOpcode::Set),
            just("ADD").or(just("add")).to(BasicOpcode::Add),
            just("SUB").or(just("sub")).to(BasicOpcode::Subtract),
            just("MUL").or(just("mul")).to(BasicOpcode::Multiply),
            just("MLI").or(just("mli")).to(BasicOpcode::MultiplySigned),
            just("DIV").or(just("div")).to(BasicOpcode::Divide),
            just("DVI").or(just("dvi")).to(BasicOpcode::DivideSigned),
            just("MOD").or(just("mod")).to(BasicOpcode::Modulo),
            just("MDI").or(just("mdi")).to(BasicOpcode::ModuloSigned),
            just("AND").or(just("and")).to(BasicOpcode::BinaryAnd),
            just("BOR").or(just("bor")).to(BasicOpcode::BinaryOr),
            just("XOR").or(just("xor")).to(BasicOpcode::BinaryExclusiveOr),
            just("SHR").or(just("shr")).to(BasicOpcode::ShiftRight),
            just("ASR").or(just("asr")).to(BasicOpcode::ArithmeticShiftRight),
            just("SHL").or(just("shl")).to(BasicOpcode::ShiftLeft),
            just("IFB").or(just("ifb")).to(BasicOpcode::IfBitSet),
            just("IFC").or(just("ifc")).to(BasicOpcode::IfClear),
            just("IFE").or(just("ife")).to(BasicOpcode::IfEqual),
            just("IFN").or(just("ifn")).to(BasicOpcode::IfNotEqual),
            just("IFG").or(just("ifg")).to(BasicOpcode::IfGreaterThan),
            just("IFA").or(just("ifa")).to(BasicOpcode::IfAbove),
            just("IFL").or(just("ifl")).to(BasicOpcode::IfLessThan),
            just("IFU").or(just("ifu")).to(BasicOpcode::IfUnder),
            just("ADX").or(just("adx")).to(BasicOpcode::AddWithCarry),
            just("SBX").or(just("sbx")).to(BasicOpcode::SubtractWithCarry),
            just("STI").or(just("sti")).to(BasicOpcode::SetThenIncrement),
            just("STD").or(just("std")).to(BasicOpcode::SetThenDecrement),
        ])
    }
}

impl instructions::SpecialOpcode {
    /// specialOpcode
    ///     : JSR
    ///     | INT
    ///     | IAG
    ///     | IAS
    ///     | RFI
    ///     | IAQ
    ///     | HWN
    ///     | HWQ
    ///     | HWI
    ///     ;
    fn parser() -> impl Parser<char, Self, Error = Simple<char>> {
        use instructions::SpecialOpcode;

        choice([
            just("JSR").or(just("jsr")).to(SpecialOpcode::JumpSubroutine),
            just("INT").or(just("int")).to(SpecialOpcode::InterruptTrigger),
            just("IAG").or(just("iag")).to(SpecialOpcode::InterruptAddressGet),
            just("IAS").or(just("ias")).to(SpecialOpcode::InterruptAddressSet),
            just("RFI").or(just("rfi")).to(SpecialOpcode::ReturnFromInterrupt),
            just("IAQ").or(just("iaq")).to(SpecialOpcode::InterruptAddToQueue),
            just("HWN").or(just("hwn")).to(SpecialOpcode::HardwareNumberConnected),
            just("HWQ").or(just("hwq")).to(SpecialOpcode::HardwareQuery),
            just("HWI").or(just("hwi")).to(SpecialOpcode::HardwareInterrupt),
        ])
    }
}

impl instructions::DebugOpcode {
    /// debugOpcode
    ///     : ALT
    ///     : DUM
    ///     ;
    fn parser() -> impl Parser<char, Self, Error = Simple<char>> {
        use instructions::DebugOpcode;

        choice([
            just("ALT").or(just("alt")).to(DebugOpcode::Alert),
            just("DUM").or(just("dum")).to(DebugOpcode::DumpState),
        ])
    }
}

#[derive(Debug)]
enum OperandBWithLabel {
    With(instructions::WithPayload, String),
    Without(instructions::OperandB),
}

impl OperandBWithLabel {
    /// operandB
    ///     : register
    ///     | locationInRegister
    ///     | locationOffsetByRegister
    ///     | PUSH
    ///     | PEEK
    ///     | pick
    ///     | STACK_POINTER
    ///     | PROGRAM_COUNTER
    ///     | EXTRA
    ///     | location
    ///     | literal
    ///     ;
    fn parser() -> impl Parser<char, Self, Error = Simple<char>> {
        use instructions::{OperandB, Register, WithRegister, WithPayload};

        let register = Register::parser()
            .map(|register|
                OperandBWithLabel::Without(OperandB::WithRegister(WithRegister::Register, register)));

        let literal = text::ident().or(text::digits(10))
            .map(|_| OperandBWithLabel::Without(OperandB::WithPayload(WithPayload::Literal, 0)));

        register
            .or(literal)
    }
}

#[derive(Debug)]
enum OperandAWithLabel {
    With(instructions::WithPayload, String),
    Without(instructions::OperandA),
}

impl OperandAWithLabel {
    /// operandA
    ///     : register
    ///     | locationInRegister
    ///     | locationOffsetByRegister
    ///     | POP
    ///     | PEEK
    ///     | pick
    ///     | STACK_POINTER
    ///     | PROGRAM_COUNTER
    ///     | EXTRA
    ///     | location
    ///     | literal
    ///     | smallLiteral
    ///     ;
    fn parser() -> impl Parser<char, Self, Error = Simple<char>> {
        use instructions::{OperandA, OperandB, Register, WithRegister};

        let register = Register::parser()
            .map(|register|
                OperandAWithLabel::Without(OperandA::LeftValue(OperandB::WithRegister(WithRegister::Register, register))));

        let literal = text::ident().or(text::digits(10))
            .map(|_| OperandAWithLabel::Without(OperandA::SmallLiteral(0)));

        register
            .or(literal)
    }
}

impl instructions::Register {
    fn parser() -> impl Parser<char, Self, Error = Simple<char>> {
        use instructions::Register;

        choice([
            just('A').or(just('a')).to(Register::A),
            just('B').or(just('b')).to(Register::B),
            just('C').or(just('c')).to(Register::C),
            just('X').or(just('x')).to(Register::X),
            just('Y').or(just('y')).to(Register::Y),
            just('Z').or(just('z')).to(Register::Z),
            just('I').or(just('i')).to(Register::I),
            just('J').or(just('j')).to(Register::J),
        ]).padded()
            .then_ignore(just(',').rewind())
    }
}

#[derive(Debug)]
struct Data(Vec<Datum>);

impl Data {
    /// data
    ///     : datum (',' datum)*
    ///     ;
    fn parser() -> impl Parser<char, Self, Error = Simple<char>> {
        Datum::parser().padded()
            .separated_by(just(','))
            .map(|data| Data(data))
    }
}

#[derive(Debug)]
enum Datum {
    String(String),
    Identifier(String),
    Number(u16),
}

impl Datum {
    /// datum
    ///     : STRING
    ///     | IDENTIFIER
    ///     | NUMBER
    ///     ;
    fn parser() -> impl Parser<char, Self, Error = Simple<char>> {
        Self::string_parser()
            .or(Self::identifier_parser())
            .or(Self::number_parser())
    }

    /// STRING
    ///     : '"' (ESCAPE|.)*? '"'
    ///     ;
    fn string_parser() -> impl Parser<char, Self, Error = Simple<char>> {
        none_of::<char, &str, Simple<char>>("\"")
            .repeated()
            .delimited_by(just("\""), just("\""))
            .collect::<String>()
            .map(|string| Datum::String(string))
    }

    /// IDENTIFIER
    ///     : [._a-zA-Z]+[._a-zA-Z0-9]*
    ///     ;
    fn identifier_parser() -> impl Parser<char, Self, Error = Simple<char>> {
        text::ident()
            .map(|identifier| Datum::String(identifier))
    }

    /// NUMBER
    ///     : '0x' [0-9a-fA-F]+
    ///     | '0b' [0-1]+
    ///     | '-'? [0-9]+
    ///     ;
    fn number_parser() -> impl Parser<char, Self, Error = Simple<char>> {
        let decimal = just('-').or_not()
            .then(text::digits(10))
            .map(|(minus_sign, digits)| {
                let value = digits.parse().unwrap();
                match minus_sign {
                    None => value,
                    _ => -(value as i16) as u16,
                }
            })
            .map(|number| Datum::Number(number));

        let binary = just("0b")
            .ignore_then(text::digits(2))
            .map(|digits| {
                let digits: String = digits;
                u16::from_str_radix(digits.as_str(), 2).unwrap()
            })
            .map(|number: u16| Datum::Number(number));

        let hexadecimal = just("0x")
            .ignore_then(text::digits(16))
            .map(|text: String| u16::from_str_radix(text.as_str(), 16).unwrap())
            .map(|number| Datum::Number(number));

        binary
            .or(hexadecimal)
            .or(decimal)
    }
}
