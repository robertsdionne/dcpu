use crate::dcpu::Dcpu;
use crate::instructions;
use crate::instructions::{
    BasicOpcode, Instruction, OperandA, OperandB, Register, WithPayload, WithRegister,
};

#[test]
fn divide_signed_register_with_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(16),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::DivideSigned,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::LeftValue(OperandB::WithPayload(WithPayload::Literal, 0xfffe)),
            )
            .into(),
        ]),
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0xfff8, cpu.register_a);
    assert_eq!(0, cpu.extra);
}

#[test]
fn divide_signed_by_zero() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![
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
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0x0, cpu.register_a);
    assert_eq!(0x1, cpu.extra);
}
