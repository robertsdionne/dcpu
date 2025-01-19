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
                OperandA::SmallLiteral(13),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Add,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(14),
            )
            .into(),
        ]),
        assembler::assemble(
            "set a, 13
            add a, 14"
        )?,
    );
    Ok(())
}

#[test]
fn register_with_overflow() -> Result<(), Box<dyn error::Error>> {
    assert_eq!(
        instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(-1),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Add,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(-1),
            )
            .into(),
        ]),
        assembler::assemble(
            "set a, -1
            add a, -1"
        )?,
    );
    Ok(())
}
