use crate::assembler;
use crate::instructions::{
    BasicOpcode, Instruction, OperandA, OperandB, Register, WithPayload, WithRegister,
};
use std::error;

#[test]
fn register_with_small_literal() -> Result<(), Box<dyn error::Error>> {
    assert_eq!(
        vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::WithPayload(WithPayload::Literal, 0xf0f0)),
            )
            .into(),
            0xf0f0,
            Instruction::Basic(
                BasicOpcode::BinaryAnd,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::WithPayload(WithPayload::Literal, 0x00ff)),
            )
            .into(),
            0x00ff,
        ],
        assembler::assemble(
            "
            set a, 0xf0f0
            and a, 0x00ff
        "
        )?,
    );

    assert_eq!(
        vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithPayload(WithPayload::Literal, 0xf0f0),
                OperandA::Operand(OperandB::WithRegister(WithRegister::Register, Register::A)),
            )
            .into(),
            0xf0f0,
            Instruction::Basic(
                BasicOpcode::BinaryAnd,
                OperandB::WithPayload(WithPayload::Literal, 0x00ff),
                OperandA::Operand(OperandB::WithRegister(WithRegister::Register, Register::A)),
            )
            .into(),
            0x00ff,
            Instruction::Basic(
                BasicOpcode::BinaryAnd,
                OperandB::WithPayload(WithPayload::Literal, 0xf0f0),
                OperandA::Operand(OperandB::WithPayload(WithPayload::Literal, 0x00ff)),
            )
            .into(),
            0x00ff,
            0xf0f0,
        ],
        assembler::assemble(
            "
            set 0xf0f0, a
            and 0x00ff, a
            and 0xf0f0, 0x00ff
        "
        )?,
    );
    Ok(())
}
