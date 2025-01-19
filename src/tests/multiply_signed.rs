use crate::dcpu::Dcpu;
use crate::instructions;
use crate::instructions::{BasicOpcode, Instruction, OperandA, OperandB, Register, WithRegister};

#[test]
fn multiply_signed_small_literal_with_register() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(-1),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::MultiplySigned,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(16),
            )
            .into(),
        ]),
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0xfff0, cpu.register_a);
    assert_eq!(0xffff, cpu.extra);
}

#[test]
fn multiply_signed_register_with_overflow() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(-1),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::MultiplySigned,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(-1),
            )
            .into(),
        ]),
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0x1, cpu.register_a);
    assert_eq!(0x0, cpu.extra);
}
