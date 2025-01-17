use std::{error, fs};
use chumsky::{prelude, primitive, text};
use chumsky::prelude::Simple;
use chumsky::Parser;
use chumsky::text::TextParser;
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
    fn parser() -> impl Parser<char, Self, Error = Simple<char>> {
        Statement::parser().padded()
            .repeated().at_least(1)
            .then_ignore(prelude::end())
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
    fn parser() -> impl Parser<char, Self, Error = Simple<char>> {
        Self::label_definition_parser()
            .or(Self::instruction_parser())
            .or(Self::data_section_parser())
    }

    fn label_definition_parser() -> impl Parser<char, Self, Error = Simple<char>> {
        primitive::just(':')
            .ignore_then(text::ident())
            .map(|label| Statement::LabelDefinition(label))
    }

    fn instruction_parser() -> impl Parser<char, Self, Error = Simple<char>> {
        InstructionWithLabels::parser()
            .map(|instruction| Statement::Instruction(instruction))
    }

    fn data_section_parser() -> impl Parser<char, Self, Error = Simple<char>> {
        primitive::just(".DAT")
            .or(primitive::just(".dat"))
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
    fn parser() -> impl Parser<char, Self, Error = Simple<char>> {
        Self::basic_parser()
            .or(Self::special_parser())
    }

    fn basic_parser() -> impl Parser<char, Self, Error = Simple<char>> {
        text::ident().padded()
            .then(text::ident().or(text::digits(10))).padded()
            .then_ignore(primitive::just(',')).padded()
            .then(text::digits(10).or(text::ident()))
            .map(|_| InstructionWithLabels::Debug(instructions::DebugOpcode::Noop))
    }

    fn special_parser() -> impl Parser<char, Self, Error = Simple<char>> {
        text::ident().padded()
            .then(text::digits(10))
            .map(|_| InstructionWithLabels::Debug(instructions::DebugOpcode::Noop))
    }
}

#[derive(Debug)]
enum OperandBWithLabel {
    With(instructions::WithPayload, String),
    Without(instructions::OperandB),
}

#[derive(Debug)]
enum OperandAWithLabel {
    With(instructions::WithPayload, String),
    Without(instructions::OperandA),
}

#[derive(Debug)]
struct Data(Vec<Datum>);

impl Data {
    fn parser() -> impl Parser<char, Self, Error = Simple<char>> {
        Datum::parser().padded()
            .separated_by(primitive::just(','))
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
    fn parser() -> impl Parser<char, Self, Error = Simple<char>> {
        Self::string_parser()
            .or(Self::identifier_parser())
            .or(Self::number_parser())
    }

    fn string_parser() -> impl Parser<char, Self, Error = Simple<char>> {
        primitive::none_of::<char, &str, Simple<char>>("\"")
            .repeated()
            .delimited_by(primitive::just("\""), primitive::just("\""))
            .collect::<String>()
            .map(|string| Datum::String(string))
    }

    fn identifier_parser() -> impl Parser<char, Self, Error = Simple<char>> {
        text::ident()
            .map(|identifier| Datum::String(identifier))
    }

    fn number_parser() -> impl Parser<char, Self, Error = Simple<char>> {
        let decimal = primitive::just('-').or_not()
            .then(text::digits(10))
            .map(|(minus_sign, digits)| {
                let value = digits.parse().unwrap();
                match minus_sign {
                    None => value,
                    _ => -(value as i16) as u16,
                }
            })
            .map(|number| Datum::Number(number));

        let binary = primitive::just("0b")
            .ignore_then(text::digits(2))
            .map(|digits| {
                let digits: String = digits;
                u16::from_str_radix(digits.as_str(), 2).unwrap()
            })
            .map(|number: u16| Datum::Number(number));

        let hexadecimal = primitive::just("0x")
            .ignore_then(text::digits(16))
            .map(|text: String| u16::from_str_radix(text.as_str(), 16).unwrap())
            .map(|number| Datum::Number(number));

        binary
            .or(hexadecimal)
            .or(decimal)
    }
}
