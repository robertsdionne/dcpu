use crate::dcpu::Dcpu;
use crate::instructions;
use crate::instructions::{BasicOpcode, Instruction, OperandA, OperandB, Register, WithRegister};

#[test]
fn register_with_equal_small_literal() {
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
                BasicOpcode::IfEqual,
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
    assert_eq!(13, cpu.memory[cpu.stack_pointer as usize]);
    assert_eq!(0xffff, cpu.stack_pointer);
}

#[test]
fn register_with_unequal_small_literal() {
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
                BasicOpcode::IfEqual,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(10),
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
fn skips_conditionals_when_not_equal() {
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
                BasicOpcode::IfEqual,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(0),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::IfEqual,
                OperandB::WithRegister(WithRegister::Register, Register::B),
                OperandA::SmallLiteral(0),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::PushOrPop,
                OperandA::SmallLiteral(12),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::PushOrPop,
                OperandA::SmallLiteral(13),
            )
            .into(),
        ]),
    );

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(13, cpu.memory[cpu.stack_pointer as usize]);
    assert_eq!(0xffff, cpu.stack_pointer);
}

#[test]
fn does_not_skip_conditionals_when_equal() {
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
                BasicOpcode::IfEqual,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(15),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::IfEqual,
                OperandB::WithRegister(WithRegister::Register, Register::B),
                OperandA::SmallLiteral(0),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::PushOrPop,
                OperandA::SmallLiteral(12),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::PushOrPop,
                OperandA::SmallLiteral(13),
            )
            .into(),
        ]),
    );

    cpu.execute_instructions(&mut hardware, 4);
    assert_eq!(12, cpu.memory[cpu.stack_pointer as usize]);
    assert_eq!(0xffff, cpu.stack_pointer);
}
