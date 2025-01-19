use crate::instructions::{BasicOpcode, Instruction, OperandA, OperandB, Register, WithRegister};
use crate::{assembler, instructions};
use std::error;

#[test]
fn register_with_lesser_small_literal() -> Result<(), Box<dyn error::Error>> {
    assert_eq!(
        instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(30),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::IfBitSet,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(16),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::PushOrPop,
                OperandA::SmallLiteral(13),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::PushOrPop,
                OperandA::SmallLiteral(14),
            )
            .into(),
        ]),
        assembler::assemble(
            "
            set a, 30
            ifb a, 16
            set push, 13
            set push, 14
        "
        )?,
    );
    Ok(())
}

#[test]
fn register_with_greater_small_literal() -> Result<(), Box<dyn error::Error>> {
    assert_eq!(
        instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(15),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::IfBitSet,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(16),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::PushOrPop,
                OperandA::SmallLiteral(13),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::PushOrPop,
                OperandA::SmallLiteral(14),
            )
            .into(),
        ]),
        assembler::assemble(
            "
            set a, 15
            ifb a, 16
            set push, 13
            set push, 14
        "
        )?,
    );
    Ok(())
}
