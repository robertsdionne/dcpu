use crate::instructions::{
    BasicOpcode, Instruction, OperandA, OperandB, Register, WithPayload, WithRegister,
};
use crate::{assembler, instructions};
use std::error;

#[test]
fn register_with_small_literal() -> Result<(), Box<dyn error::Error>> {
    assert_eq!(
        vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(16),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::DivideSigned,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::WithPayload(WithPayload::Literal, 0xfffe)),
            )
            .into(),
            0xfffe,
        ],
        assembler::assemble(
            "set a, 16
            dvi a, 0xfffe"
        )?,
    );
    Ok(())
}

#[test]
fn by_zero() -> Result<(), Box<dyn error::Error>> {
    assert_eq!(
        instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(30),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::DivideSigned,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(0),
            )
            .into(),
        ]),
        assembler::assemble(
            "set a, 30
            dvi a, 0"
        )?,
    );
    Ok(())
}
