use crate::dcpu::Dcpu;
use crate::instructions::{BasicOpcode, Instruction, OperandA, OperandB, Register, WithRegister};

#[test]
fn subtract_small_literal_from_register() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(30),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Subtract,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(16),
            )
            .into(),
        ],
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0xe, cpu.register_a);
    assert_eq!(0, cpu.extra);
}

#[test]
fn subtract_register_with_underflow() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(16),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Subtract,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(30),
            )
            .into(),
        ],
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0xfff2, cpu.register_a);
    assert_eq!(0xffff, cpu.extra);
}
