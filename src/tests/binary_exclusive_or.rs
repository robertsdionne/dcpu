use crate::dcpu::Dcpu;
use crate::instructions;
use crate::instructions::{
    BasicOpcode, Instruction, OperandA, OperandB, Register, WithPayload, WithRegister,
};

#[test]
fn register_with_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::WithPayload(WithPayload::Literal, 0xf0f0)),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::BinaryExclusiveOr,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::WithPayload(WithPayload::Literal, 0x00ff)),
            )
            .into(),
        ]),
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0xf00f, cpu.register_a);
}
