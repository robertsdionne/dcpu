use crate::dcpu::Dcpu;
use crate::instructions;
use crate::instructions::{BasicOpcode, Instruction, OperandA, OperandB, Register, WithRegister};

#[test]
fn small_literal_to_register() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![
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
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0x1b, cpu.register_a);
    assert_eq!(0, cpu.extra);
}

#[test]
fn register_with_overflow() {
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
                BasicOpcode::Add,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(-1),
            )
            .into(),
        ]),
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0xfffe, cpu.register_a);
    assert_eq!(1, cpu.extra);
}
