use crate::dcpu::Dcpu;
use crate::instructions;
use crate::instructions::{BasicOpcode, Instruction, OperandA, OperandB, Register, WithRegister};

#[test]
fn if_clear_with_common_bits() {
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
                BasicOpcode::IfClear,
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
    );

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(14, cpu.memory[cpu.stack_pointer as usize]);
    assert_eq!(0xffff, cpu.stack_pointer);
}

#[test]
fn if_clear_without_common_bits() {
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
                BasicOpcode::IfClear,
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
    );

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(13, cpu.memory[cpu.stack_pointer as usize]);
    assert_eq!(0xffff, cpu.stack_pointer);
}
