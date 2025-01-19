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
                OperandA::Operand(OperandB::WithPayload(WithPayload::Literal, 0xfff0)),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::ArithmeticShiftRight,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(2),
            )
            .into(),
        ]),
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0xfffc, cpu.register_a);
    assert_eq!(0, cpu.extra);
}

#[test]
fn register_with_underflow() {
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
                BasicOpcode::ArithmeticShiftRight,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(2),
            )
            .into(),
        ]),
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0xffff, cpu.register_a);
    assert_eq!(0xc000, cpu.extra);
}
