use std::{error, fs};
use chumsky::prelude::*;
use crate::instructions;
use crate::instructions::OperandB::WithPayload;

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
            .or(just("DAT"))
            .or(just("dat"))
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
            .then(OperandBWithLabel::parser(true)).padded()
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
    ///     | pushOrPop
    ///     | PEEK
    ///     | pick
    ///     | STACK_POINTER
    ///     | PROGRAM_COUNTER
    ///     | EXTRA
    ///     | location
    ///     | literal
    ///     ;
    fn parser(operand_b: bool) -> impl Parser<char, Self, Error = Simple<char>> {
        use instructions::{OperandB, Register, WithRegister, WithPayload};

        let register = Register::parser()
            .map(|register|
                OperandBWithLabel::Without(OperandB::WithRegister(WithRegister::Register, register)));

        // locationInRegister
        //     : '[' REGISTER ']'
        //     ;
        let location_in_register = Register::parser()
            .delimited_by(just('['), just(']'))
            .map(|register|
                OperandBWithLabel::Without(OperandB::WithRegister(WithRegister::LocationInRegister, register)));

        // registerOffsetByLiteral
        //     : REGISTER '+' (label | value)
        //     ;
        let register_offset_by_literal = Register::parser().padded()
            .then_ignore(just('+').padded())
            .then(Datum::identifier_parser().or(Datum::number_parser()).padded());

        // literalOffsetByRegister
        //     : (label | value) '+' REGISTER
        //     ;
        let literal_offset_by_register = Datum::identifier_parser()
            .or(Datum::number_parser()).padded()
            .then_ignore(just('+').padded())
            .then(Register::parser()).padded()
            .map(|(datum, register)| (register, datum));

        // locationOffsetByRegister
        //     : '[' (registerOffsetByLiteral | literalOffsetByRegister) ']'
        //     ;
        let location_offset_by_register = register_offset_by_literal.or(literal_offset_by_register)
            .delimited_by(just('['), just(']'))
            .map(|(register, datum)| match datum {
                Datum::Identifier(label) =>
                    OperandBWithLabel::With(WithPayload::LocationOffsetByRegister(register), label),
                Datum::Number(value) =>
                    OperandBWithLabel::Without(OperandB::WithPayload(WithPayload::LocationOffsetByRegister(register), value)),
                _ => unreachable!(),
            });

        // pushOrPop
        //     : PUSH
        //     | POP
        //     ;
        let push_or_pop = if operand_b {
            just("push").or(just("PUSH"))
        } else {
            just("pop").or(just("POP"))
        }.ignored()
            .map(|_| OperandBWithLabel::Without(OperandB::PushOrPop));

        let peek = just("peek").or(just("PEEK")).ignored()
            .map(|_| OperandBWithLabel::Without(OperandB::Peek));

        // pick
        //     : PICK (label | value)
        //     ;
        let pick = just("pick").or(just("PICK")).padded().ignored()
            .then(Datum::identifier_parser().or(Datum::number_parser()))
            .map(|(_, datum)| match datum {
                Datum::Identifier(label) => OperandBWithLabel::With(WithPayload::Pick, label),
                Datum::Number(value) => OperandBWithLabel::Without(OperandB::WithPayload(WithPayload::Pick, value)),
                _ => unreachable!(),
            });

        let stack_pointer = just("sp").or(just("SP")).ignored()
            .map(|_| OperandBWithLabel::Without(OperandB::StackPointer));

        let program_counter = just("pc").or(just("PC")).ignored()
            .map(|_| OperandBWithLabel::Without(OperandB::ProgramCounter));

        let extra = just("ex").or(just("EX")).ignored()
            .map(|_| OperandBWithLabel::Without(OperandB::Extra));

        // location
        //     : '[' (label | value) ']'
        //     ;
        let location = Datum::identifier_parser().or(Datum::number_parser()).padded()
            .delimited_by(just('['), just(']'))
            .map(|datum| match datum {
                Datum::Identifier(label) => OperandBWithLabel::With(WithPayload::Location, label),
                Datum::Number(value) => OperandBWithLabel::Without(OperandB::WithPayload(WithPayload::Location, value)),
                _ => unreachable!(),
            });

        // literal
        //     : (label | value)
        //     ;
        let literal = Datum::identifier_parser().or(Datum::number_parser()).padded()
            .map(|datum| match datum {
                Datum::Identifier(label) => OperandBWithLabel::With(WithPayload::Literal, label),
                Datum::Number(value) => OperandBWithLabel::Without(OperandB::WithPayload(WithPayload::Literal, value)),
                _ => unreachable!(),
            });

        register
            .or(location_in_register)
            .or(location_offset_by_register)
            .or(push_or_pop)
            .or(peek)
            .or(pick)
            .or(stack_pointer)
            .or(program_counter)
            .or(extra)
            .or(location)
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
    ///     : operandB
    ///     | smallLiteral
    ///     ;
    fn parser() -> impl Parser<char, Self, Error = Simple<char>> {
        use instructions::{OperandA, OperandB};

        OperandBWithLabel::parser(false)
            .map(|operand_b| match operand_b {
                OperandBWithLabel::With(with_payload, label) =>
                    OperandAWithLabel::With(with_payload, label),
                OperandBWithLabel::Without(operand_b @ OperandB::WithPayload(_, payload)) => {
                    if payload == 0xffff || payload <= 30 {
                        OperandAWithLabel::Without(OperandA::SmallLiteral(payload as i8))
                    } else {
                        OperandAWithLabel::Without(OperandA::LeftValue(operand_b))
                    }
                }
                OperandBWithLabel::Without(operand_b) =>
                    OperandAWithLabel::Without(OperandA::LeftValue(operand_b)),
            })
    }
}

impl instructions::Register {
    /// REGSITER
    ///     : [abcxyzijABCXYZIJ]
    ///     ;
    fn parser() -> impl Parser<char, Self, Error = Simple<char>> {
        use instructions::Register;

        one_of("abcxyzijABCXYZIJ").padded()
            .then_ignore(one_of(",+[]").rewind())
            .map(|register| match register {
                'a' | 'A' => Register::A,
                'b' | 'B' => Register::B,
                'c' | 'C' => Register::C,
                'x' | 'X' => Register::X,
                'y' | 'Y' => Register::Y,
                'z' | 'Z' => Register::Z,
                'i' | 'I' => Register::I,
                'j' | 'J' => Register::J,
                _ => unreachable!(),
            })
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
            .map(|identifier| Datum::Identifier(identifier))
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
