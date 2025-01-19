use crate::dcpu::Dcpu;
use crate::instructions;
use crate::instructions::{
    BasicOpcode, Instruction, OperandA, OperandB, Register, WithPayload, WithRegister,
};

#[test]
fn with_lesser_small_literal() {
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
                BasicOpcode::IfGreaterThan,
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
fn with_greater_small_literal() {
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
                BasicOpcode::IfGreaterThan,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(30),
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
fn with_greater_literal() {
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
                BasicOpcode::IfGreaterThan,
                OperandB::WithPayload(WithPayload::Literal, 0x0a),
                OperandA::Operand(OperandB::WithRegister(WithRegister::Register, Register::A)),
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
fn with_lesser_literal() {
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
                BasicOpcode::IfGreaterThan,
                OperandB::WithPayload(WithPayload::Literal, 0x1e),
                OperandA::Operand(OperandB::WithRegister(WithRegister::Register, Register::A)),
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
