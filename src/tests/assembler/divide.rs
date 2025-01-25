use crate::instructions::{BasicOpcode, Instruction, OperandA, OperandB, Register, WithRegister};
use crate::{assembler, instructions};
use std::error;

#[test]
fn register_with_small_literal() -> Result<(), Box<dyn error::Error>> {
    assert_eq!(
        instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(30),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Divide,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(16),
            )
            .into(),
        ]),
        assembler::assemble(
            "set a, 30
            div a, 16"
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
                BasicOpcode::Divide,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(0),
            )
            .into(),
        ]),
        assembler::assemble(
            "set a, 30
            div a, 0"
        )?,
    );
    Ok(())
}
