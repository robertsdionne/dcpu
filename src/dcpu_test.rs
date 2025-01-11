use crate::dcpu::*;
use crate::instructions::{BasicOpcode, Instruction, OperandA, OperandB};

#[test]
fn if_above_with_lesser_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(0, &vec![
        Instruction::Basic(BasicOpcode::Set, OperandB::RegisterA, OperandA::SmallLiteral(30)).into(),
        Instruction::Basic(BasicOpcode::IfAbove, OperandB::RegisterA, OperandA::SmallLiteral(-1)).into(),
        Instruction::Basic(BasicOpcode::Set, OperandB::PushOrPop, OperandA::SmallLiteral(13)).into(),
        Instruction::Basic(BasicOpcode::Set, OperandB::PushOrPop, OperandA::SmallLiteral(14)).into(),
    ]);

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(13, cpu.memory[cpu.stack_pointer as usize]);
    assert_eq!(0xffff, cpu.stack_pointer);
}

#[test]
fn if_above_with_greater_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(0, &vec![
        Instruction::Basic(BasicOpcode::Set, OperandB::RegisterA, OperandA::SmallLiteral(-1)).into(),
        Instruction::Basic(BasicOpcode::IfAbove, OperandB::RegisterA, OperandA::SmallLiteral(30)).into(),
        Instruction::Basic(BasicOpcode::Set, OperandB::PushOrPop, OperandA::SmallLiteral(13)).into(),
        Instruction::Basic(BasicOpcode::Set, OperandB::PushOrPop, OperandA::SmallLiteral(14)).into(),
    ]);

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(14, cpu.memory[cpu.stack_pointer as usize]);
    assert_eq!(0xffff, cpu.stack_pointer);
}

#[test]
fn if_clear_with_common_bits() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(0, &vec![
        Instruction::Basic(BasicOpcode::Set, OperandB::RegisterA, OperandA::SmallLiteral(30)).into(),
        Instruction::Basic(BasicOpcode::IfClear, OperandB::RegisterA, OperandA::SmallLiteral(16)).into(),
        Instruction::Basic(BasicOpcode::Set, OperandB::PushOrPop, OperandA::SmallLiteral(13)).into(),
        Instruction::Basic(BasicOpcode::Set, OperandB::PushOrPop, OperandA::SmallLiteral(14)).into(),
    ]);

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(14, cpu.memory[cpu.stack_pointer as usize]);
    assert_eq!(0xffff, cpu.stack_pointer);
}

#[test]
fn if_clear_without_common_bits() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(0, &vec![
        Instruction::Basic(BasicOpcode::Set, OperandB::RegisterA, OperandA::SmallLiteral(15)).into(),
        Instruction::Basic(BasicOpcode::IfClear, OperandB::RegisterA, OperandA::SmallLiteral(16)).into(),
        Instruction::Basic(BasicOpcode::Set, OperandB::PushOrPop, OperandA::SmallLiteral(13)).into(),
        Instruction::Basic(BasicOpcode::Set, OperandB::PushOrPop, OperandA::SmallLiteral(14)).into(),
    ]);

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(13, cpu.memory[cpu.stack_pointer as usize]);
    assert_eq!(0xffff, cpu.stack_pointer);
}

#[test]
fn if_equal_register_with_equal_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(0, &vec![
        Instruction::Basic(BasicOpcode::Set, OperandB::RegisterA, OperandA::SmallLiteral(15)).into(),
        Instruction::Basic(BasicOpcode::IfEqual, OperandB::RegisterA, OperandA::SmallLiteral(15)).into(),
        Instruction::Basic(BasicOpcode::Set, OperandB::PushOrPop, OperandA::SmallLiteral(13)).into(),
        Instruction::Basic(BasicOpcode::Set, OperandB::PushOrPop, OperandA::SmallLiteral(14)).into(),
    ]);

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(13, cpu.memory[cpu.stack_pointer as usize]);
    assert_eq!(0xffff, cpu.stack_pointer);
}

#[test]
fn if_equal_register_with_unequal_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(0, &vec![
        Instruction::Basic(BasicOpcode::Set, OperandB::RegisterA, OperandA::SmallLiteral(15)).into(),
        Instruction::Basic(BasicOpcode::IfEqual, OperandB::RegisterA, OperandA::SmallLiteral(10)).into(),
        Instruction::Basic(BasicOpcode::Set, OperandB::PushOrPop, OperandA::SmallLiteral(13)).into(),
        Instruction::Basic(BasicOpcode::Set, OperandB::PushOrPop, OperandA::SmallLiteral(14)).into(),
    ]);

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(14, cpu.memory[cpu.stack_pointer as usize]);
    assert_eq!(0xffff, cpu.stack_pointer);
}

#[test]
fn if_equal_skips_conditionals_when_not_equal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(0, &vec![
        Instruction::Basic(BasicOpcode::Set, OperandB::RegisterA, OperandA::SmallLiteral(15)).into(),
        Instruction::Basic(BasicOpcode::IfEqual, OperandB::RegisterA, OperandA::SmallLiteral(0)).into(),
        Instruction::Basic(BasicOpcode::IfEqual, OperandB::RegisterB, OperandA::SmallLiteral(0)).into(),
        Instruction::Basic(BasicOpcode::Set, OperandB::PushOrPop, OperandA::SmallLiteral(12)).into(),
        Instruction::Basic(BasicOpcode::Set, OperandB::PushOrPop, OperandA::SmallLiteral(13)).into(),
    ]);

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(13, cpu.memory[cpu.stack_pointer as usize]);
    assert_eq!(0xffff, cpu.stack_pointer);
}

#[test]
fn if_equal_does_not_skip_conditionals_when_equal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(0, &vec![
        Instruction::Basic(BasicOpcode::Set, OperandB::RegisterA, OperandA::SmallLiteral(15)).into(),
        Instruction::Basic(BasicOpcode::IfEqual, OperandB::RegisterA, OperandA::SmallLiteral(15)).into(),
        Instruction::Basic(BasicOpcode::IfEqual, OperandB::RegisterB, OperandA::SmallLiteral(0)).into(),
        Instruction::Basic(BasicOpcode::Set, OperandB::PushOrPop, OperandA::SmallLiteral(12)).into(),
        Instruction::Basic(BasicOpcode::Set, OperandB::PushOrPop, OperandA::SmallLiteral(13)).into(),
    ]);

    cpu.execute_instructions(&mut hardware, 4);
    assert_eq!(12, cpu.memory[cpu.stack_pointer as usize]);
    assert_eq!(0xffff, cpu.stack_pointer);
}
