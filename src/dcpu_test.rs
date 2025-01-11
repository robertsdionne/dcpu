use crate::dcpu::*;
use crate::instructions::{BasicOpcode, Instruction, OperandA, OperandB};

#[test]
fn add_small_literal_to_register() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(0, &vec![
        Instruction::Basic(BasicOpcode::Set, OperandB::RegisterA, OperandA::SmallLiteral(13)).into(),
        Instruction::Basic(BasicOpcode::Add, OperandB::RegisterA, OperandA::SmallLiteral(14)).into(),
    ]);

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0x1b, cpu.register_a);
    assert_eq!(0, cpu.extra);
}

#[test]
fn add_register_with_overflow() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(0, &vec![
        Instruction::Basic(BasicOpcode::Set, OperandB::RegisterA, OperandA::SmallLiteral(-1)).into(),
        Instruction::Basic(BasicOpcode::Add, OperandB::RegisterA, OperandA::SmallLiteral(-1)).into(),
    ]);

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0xfffe, cpu.register_a);
    assert_eq!(1, cpu.extra);
}

#[test]
fn subtract_small_literal_from_register() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(0, &vec![
        Instruction::Basic(BasicOpcode::Set, OperandB::RegisterA, OperandA::SmallLiteral(30)).into(),
        Instruction::Basic(BasicOpcode::Subtract, OperandB::RegisterA, OperandA::SmallLiteral(16)).into(),
    ]);

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0xe, cpu.register_a);
    assert_eq!(0, cpu.extra);
}

#[test]
fn subtract_register_with_underflow() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(0, &vec![
        Instruction::Basic(BasicOpcode::Set, OperandB::RegisterA, OperandA::SmallLiteral(16)).into(),
        Instruction::Basic(BasicOpcode::Subtract, OperandB::RegisterA, OperandA::SmallLiteral(30)).into(),
    ]);

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0xfff2, cpu.register_a);
    assert_eq!(0xffff, cpu.extra);
}

#[test]
fn multiply_small_literal_with_register() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(0, &vec![
        Instruction::Basic(BasicOpcode::Set, OperandB::RegisterA, OperandA::SmallLiteral(16)).into(),
        Instruction::Basic(BasicOpcode::Multiply, OperandB::RegisterA, OperandA::SmallLiteral(30)).into(),
    ]);

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0x1e0, cpu.register_a);
    assert_eq!(0, cpu.extra);
}

#[test]
fn multiply_multiply_register_with_overflow() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(0, &vec![
        Instruction::Basic(BasicOpcode::Set, OperandB::RegisterA, OperandA::SmallLiteral(-1)).into(),
        Instruction::Basic(BasicOpcode::Multiply, OperandB::RegisterA, OperandA::SmallLiteral(-1)).into(),
    ]);

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0x1, cpu.register_a);
    assert_eq!(0xfffe, cpu.extra);
}

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
