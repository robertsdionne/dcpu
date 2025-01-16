use std::error;
use nom::{branch, bytes, multi, sequence};
use nom::character::complete;
use crate::instructions::{BasicOpcode, DebugOpcode, OperandA, OperandB, SpecialOpcode};

pub fn assemble(_program: &str) -> Result<(), Box<dyn error::Error>> {
    Ok(())
}

struct Program {
    components: Vec<ProgramComponent>,
}

impl Program {
    fn parse(input: &str) -> nom::IResult<&str, Self> {
        let result = multi::many1(ProgramComponent::parse)(input)?;

        Ok((input, Self {
            components: result.1,
        }))
    }
}

enum ProgramComponent {
    LabelDefinition,
    Instruction,
    DataSection,
}

impl ProgramComponent {
    fn parse(input: &str) -> nom::IResult<&str, Self> {
        branch::alt((
            Self::parse_label_definition,
            Self::parse_instruction,
            Self::parse_data_section,
        ))(input)
    }

    fn parse_label_definition(input: &str) -> nom::IResult<&str, ProgramComponent> {
        sequence::preceded(
            complete::char(':'),
            Identifier::parse,
        )(input)?;
        Ok((input, ProgramComponent::LabelDefinition))
    }

    fn parse_instruction(input: &str) -> nom::IResult<&str, ProgramComponent> {
        Instruction::parse(input)?;
        Ok((input, ProgramComponent::Instruction))
    }

    fn parse_data_section(input: &str) -> nom::IResult<&str, ProgramComponent> {
        Ok((input, ProgramComponent::DataSection))
    }
}

struct Identifier;

impl Identifier {
    fn parse(input: &str) -> nom::IResult<&str, Self> {
        Ok((input, Identifier))
    }
}

enum Instruction {
    Basic,
    Special,
    Debug,
}

impl Instruction {
    fn parse(input: &str) -> nom::IResult<&str, Self> {
        branch::alt((
            Self::parse_basic_instruction,
            Self::parse_special_instruction,
            Self::parse_debug_instruction,
        ))(input)
    }

    fn parse_basic_instruction(input: &str) -> nom::IResult<&str, Instruction> {
        sequence::tuple((
            BasicOpcode::parse,
            OperandB::parse,
            bytes::complete::tag(", "),
            OperandA::parse,

        ))(input)?;
        Ok((input, Instruction::Basic))
    }

    fn parse_special_instruction(input: &str) -> nom::IResult<&str, Instruction> {
        Ok((input, Instruction::Special))
    }

    fn parse_debug_instruction(input: &str) -> nom::IResult<&str, Instruction> {
        Ok((input, Instruction::Debug))
    }
}

impl BasicOpcode {
    fn parse(input: &str) -> nom::IResult<&str, BasicOpcode> {
        use bytes::complete::tag_no_case;

        let result: (&str, &str) = branch::alt((
            branch::alt((
                tag_no_case("SET"),
                tag_no_case("ADD"),
                tag_no_case("SUB"),
                tag_no_case("MUL"),
                tag_no_case("DIV"),
                tag_no_case("DVI"),
                tag_no_case("MOD"),
                tag_no_case("MDI"),
                tag_no_case("AND"),
                tag_no_case("BOR"),
                tag_no_case("SHR"),
                tag_no_case("ASR"),
                tag_no_case("SHL"),
                tag_no_case("IFB"),
                tag_no_case("IFC"),
                tag_no_case("IFE"),
            )),
            tag_no_case("IFN"),
            tag_no_case("IFG"),
            tag_no_case("IFA"),
            tag_no_case("IFL"),
            tag_no_case("IFU"),
            tag_no_case("ADX"),
            tag_no_case("SBX"),
            tag_no_case("STI"),
            tag_no_case("STD"),
        ))(input)?;
        Ok((input, BasicOpcode::Reserved))
    }
}

impl SpecialOpcode {
    fn parse(input: &str) -> nom::IResult<&str, SpecialOpcode> {
        use bytes::complete::tag_no_case;

        let result: (&str, &str) = branch::alt((
            tag_no_case("JSR"),
            tag_no_case("INT"),
            tag_no_case("IAG"),
            tag_no_case("IAS"),
            tag_no_case("RFI"),
            tag_no_case("IAQ"),
            tag_no_case("HWN"),
            tag_no_case("HWQ"),
            tag_no_case("HWI"),
        ))(input)?;
        Ok((input, SpecialOpcode::Reserved))
    }
}

impl DebugOpcode {
    fn parse(input: &str) -> nom::IResult<&str, DebugOpcode> {
        use bytes::complete::tag_no_case;
        let result: (&str, &str) = branch::alt((
            tag_no_case("ALT"),
            tag_no_case("DUM"),
        ))(input)?;
        Ok((input, DebugOpcode::Noop))
    }
}

impl OperandB {
    fn parse(input: &str) -> nom::IResult<&str, Self> {
        Ok((input, Self::Extra))
    }
}

impl OperandA {
    fn parse(input: &str) -> nom::IResult<&str, Self> {
        Ok((input, Self::SmallLiteral(0)))
    }
}
