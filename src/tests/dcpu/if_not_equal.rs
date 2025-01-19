use crate::dcpu::Dcpu;
use crate::instructions;
use crate::instructions::{BasicOpcode, Instruction, OperandA, OperandB, Register, WithRegister};

#[test]
fn with_unequal_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(15),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::IfNotEqual,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(0),
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
    );

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(13, cpu.memory[cpu.stack_pointer as usize]);
    assert_eq!(0xffff, cpu.stack_pointer);
}

#[test]
fn with_equal_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(15),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::IfNotEqual,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(15),
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
    );

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(14, cpu.memory[cpu.stack_pointer as usize]);
    assert_eq!(0xffff, cpu.stack_pointer);
}
