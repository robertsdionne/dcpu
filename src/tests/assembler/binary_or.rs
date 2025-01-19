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
                BasicOpcode::BinaryOr,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::WithPayload(WithPayload::Literal, 0x00ff)),
            )
            .into(),
            0x00ff,
        ],
        assembler::assemble(
            "set a, 0xf0f0
            bor a, 0x00ff"
        )?,
    );
    Ok(())
}
