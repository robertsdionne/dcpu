use crate::dcpu::Dcpu;
use crate::instructions;
use crate::instructions::{
    BasicOpcode, DebugOpcode, Instruction, OperandA, OperandB, Register, WithPayload, WithRegister,
};

#[test]
fn register_with_register() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::B),
                OperandA::SmallLiteral(1),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::WithRegister(WithRegister::Register, Register::B)),
            )
            .into(),
        ]),
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0x1, cpu.register_a);
}

#[test]
fn register_with_last_register() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::J),
                OperandA::SmallLiteral(1),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::WithRegister(WithRegister::Register, Register::J)),
            )
            .into(),
        ]),
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0x1, cpu.register_a);
}

#[test]
fn register_with_location_in_register() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithPayload(WithPayload::Location, 0x1000),
                OperandA::SmallLiteral(13),
            )
            .into(),
            0x1000,
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::B),
                OperandA::Operand(OperandB::WithPayload(WithPayload::Literal, 0x1000)),
            )
            .into(),
            0x1000,
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::WithRegister(
                    WithRegister::LocationInRegister,
                    Register::B,
                )),
            )
            .into(),
        ],
    );

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(13, cpu.register_a);
}

#[test]
fn register_with_location_in_last_register() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithPayload(WithPayload::Location, 0x1000),
                OperandA::SmallLiteral(13),
            )
            .into(),
            0x1000,
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::J),
                OperandA::Operand(OperandB::WithPayload(WithPayload::Literal, 0x1000)),
            )
            .into(),
            0x1000,
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::WithRegister(
                    WithRegister::LocationInRegister,
                    Register::J,
                )),
            )
            .into(),
        ],
    );

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(13, cpu.register_a);
}

#[test]
fn register_with_location_offset_by_register() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithPayload(WithPayload::Location, 0x100a),
                OperandA::SmallLiteral(13),
            )
            .into(),
            0x100a,
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::B),
                OperandA::SmallLiteral(10),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::WithPayload(
                    WithPayload::LocationOffsetByRegister(Register::B),
                    0x1000,
                )),
            )
            .into(),
            0x1000,
        ],
    );

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(13, cpu.register_a);
}

#[test]
fn register_with_location_offset_by_last_register() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithPayload(WithPayload::Location, 0x100a),
                OperandA::SmallLiteral(13),
            )
            .into(),
            0x100a,
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::J),
                OperandA::SmallLiteral(10),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::WithPayload(
                    WithPayload::LocationOffsetByRegister(Register::J),
                    0x1000,
                )),
            )
            .into(),
            0x1000,
        ],
    );

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(13, cpu.register_a);
}

#[test]
fn register_with_pop() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::PushOrPop,
                OperandA::SmallLiteral(13),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::PushOrPop),
            )
            .into(),
        ]),
    );

    cpu.execute_instructions(&mut hardware, 1);
    assert_eq!(0xffff, cpu.stack_pointer);

    cpu.execute_instructions(&mut hardware, 1);
    assert_eq!(0, cpu.stack_pointer);
    assert_eq!(13, cpu.register_a);
}

#[test]
fn register_with_peek() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::PushOrPop,
                OperandA::SmallLiteral(13),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::Peek),
            )
            .into(),
        ]),
    );

    cpu.execute_instructions(&mut hardware, 1);
    assert_eq!(0xffff, cpu.stack_pointer);

    cpu.execute_instructions(&mut hardware, 1);
    assert_eq!(0xffff, cpu.stack_pointer);
    assert_eq!(13, cpu.register_a);
}

#[test]
fn register_with_pick() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &vec![
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
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::WithPayload(WithPayload::Pick, 1)),
            )
            .into(),
            1,
        ],
    );

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(0xfffe, cpu.stack_pointer);
    assert_eq!(13, cpu.memory[cpu.stack_pointer as usize + 1]);
    assert_eq!(13, cpu.register_a);
}

#[test]
fn register_with_stack_pointer() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::PushOrPop,
                OperandA::SmallLiteral(13),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::StackPointer),
            )
            .into(),
        ]),
    );

    cpu.execute_instructions(&mut hardware, 1);
    assert_eq!(0xffff, cpu.stack_pointer);

    cpu.execute_instructions(&mut hardware, 1);
    assert_eq!(0xffff, cpu.register_a);
}

#[test]
fn register_with_program_counter() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![
            Instruction::Debug(DebugOpcode::Noop),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::ProgramCounter),
            )
            .into(),
        ]),
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(2, cpu.register_a);
}

#[test]
fn register_with_extra() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::Extra,
                OperandA::SmallLiteral(13),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::Extra),
            )
            .into(),
        ]),
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(13, cpu.register_a);
}

#[test]
fn register_with_location() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithPayload(WithPayload::Location, 0x1000),
                OperandA::SmallLiteral(13),
            )
            .into(),
            0x1000,
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::WithPayload(WithPayload::Location, 0x1000)),
            )
            .into(),
            0x1000,
        ],
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(13, cpu.register_a);
}

#[test]
fn register_with_large_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::WithPayload(WithPayload::Literal, 0x1001)),
            )
            .into(),
            0x1001,
        ],
    );

    cpu.execute_instructions(&mut hardware, 1);
    assert_eq!(0x1001, cpu.register_a);
}

#[test]
fn register_with_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![Instruction::Basic(
            BasicOpcode::Set,
            OperandB::WithRegister(WithRegister::Register, Register::A),
            OperandA::SmallLiteral(1),
        )
        .into()]),
    );

    cpu.execute_instructions(&mut hardware, 1);
    assert_eq!(0x1, cpu.register_a);
}

#[test]
fn last_register_with_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![Instruction::Basic(
            BasicOpcode::Set,
            OperandB::WithRegister(WithRegister::Register, Register::J),
            OperandA::SmallLiteral(1),
        )
        .into()]),
    );

    cpu.execute_instructions(&mut hardware, 1);
    assert_eq!(0x1, cpu.register_j);
}

#[test]
fn location_in_register_with_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::Operand(OperandB::WithPayload(WithPayload::Literal, 0x1000)),
            )
            .into(),
            0x1000,
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::LocationInRegister, Register::A),
                OperandA::SmallLiteral(13),
            )
            .into(),
        ],
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(13, cpu.memory[0x1000]);
}

#[test]
fn location_in_last_register_with_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::J),
                OperandA::Operand(OperandB::WithPayload(WithPayload::Literal, 0x1000)),
            )
            .into(),
            0x1000,
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::LocationInRegister, Register::J),
                OperandA::SmallLiteral(13),
            )
            .into(),
        ],
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(13, cpu.memory[0x1000]);
}

#[test]
fn location_offset_by_register_with_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(10),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithPayload(WithPayload::LocationOffsetByRegister(Register::A), 0x1000),
                OperandA::SmallLiteral(13),
            )
            .into(),
            0x1000,
        ],
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(13, cpu.memory[0x100a]);
}

#[test]
fn location_offset_by_last_register_with_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::J),
                OperandA::SmallLiteral(10),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithPayload(WithPayload::LocationOffsetByRegister(Register::J), 0x1000),
                OperandA::SmallLiteral(13),
            )
            .into(),
            0x1000,
        ],
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(13, cpu.memory[0x100a]);
}

#[test]
fn location_offset_by_register_with_location_offset_by_register() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithRegister(WithRegister::Register, Register::A),
                OperandA::SmallLiteral(10),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithPayload(WithPayload::LocationOffsetByRegister(Register::A), 0x1000),
                OperandA::SmallLiteral(13),
            )
            .into(),
            0x1000,
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithPayload(WithPayload::LocationOffsetByRegister(Register::A), 0x2000),
                OperandA::Operand(OperandB::WithPayload(
                    WithPayload::LocationOffsetByRegister(Register::A),
                    0x1000,
                )),
            )
            .into(),
            0x1000,
            0x2000,
        ],
    );

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(13, cpu.memory[0x200a]);
}

#[test]
fn push_with_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![Instruction::Basic(
            BasicOpcode::Set,
            OperandB::PushOrPop,
            OperandA::SmallLiteral(13),
        )
        .into()]),
    );

    cpu.execute_instructions(&mut hardware, 1);
    assert_eq!(13, cpu.memory[cpu.stack_pointer as usize]);
}

#[test]
fn push_with_pop() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::PushOrPop,
                OperandA::SmallLiteral(13),
            )
            .into(),
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::PushOrPop,
                OperandA::Operand(OperandB::PushOrPop),
            )
            .into(),
        ]),
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(0xffff, cpu.stack_pointer);
    assert_eq!(13, cpu.memory[cpu.stack_pointer as usize]);
}

#[test]
fn peek_with_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::PushOrPop,
                OperandA::SmallLiteral(13),
            )
            .into(),
            Instruction::Basic(BasicOpcode::Set, OperandB::Peek, OperandA::SmallLiteral(14)).into(),
        ]),
    );

    cpu.execute_instructions(&mut hardware, 2);
    assert_eq!(14, cpu.memory[cpu.stack_pointer as usize]);
}

#[test]
fn pick_with_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &vec![
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
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithPayload(WithPayload::Pick, 0x1),
                OperandA::SmallLiteral(14),
            )
            .into(),
            0x1,
        ],
    );

    cpu.execute_instructions(&mut hardware, 3);
    assert_eq!(13, cpu.memory[cpu.stack_pointer as usize]);
    assert_eq!(14, cpu.memory[cpu.stack_pointer as usize + 1]);
}

#[test]
fn stack_pointer_with_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![Instruction::Basic(
            BasicOpcode::Set,
            OperandB::StackPointer,
            OperandA::SmallLiteral(13),
        )
        .into()]),
    );

    cpu.execute_instructions(&mut hardware, 1);
    assert_eq!(13, cpu.stack_pointer);
}

#[test]
fn program_counter_with_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![Instruction::Basic(
            BasicOpcode::Set,
            OperandB::ProgramCounter,
            OperandA::SmallLiteral(13),
        )
        .into()]),
    );

    cpu.execute_instructions(&mut hardware, 1);
    assert_eq!(13, cpu.program_counter);
}

#[test]
fn extra_with_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &instructions::assemble(vec![Instruction::Basic(
            BasicOpcode::Set,
            OperandB::Extra,
            OperandA::SmallLiteral(13),
        )
        .into()]),
    );

    cpu.execute_instructions(&mut hardware, 1);
    assert_eq!(13, cpu.extra);
}

#[test]
fn location_with_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithPayload(WithPayload::Location, 0x1000),
                OperandA::SmallLiteral(13),
            )
            .into(),
            0x1000,
        ],
    );

    cpu.execute_instructions(&mut hardware, 1);
    assert_eq!(13, cpu.memory[0x1000]);
}

#[test]
fn literal_with_small_literal() {
    let mut cpu = Dcpu::default();
    let mut hardware = vec![];
    cpu.load(
        0,
        &vec![
            Instruction::Basic(
                BasicOpcode::Set,
                OperandB::WithPayload(WithPayload::Literal, 0x1000),
                OperandA::SmallLiteral(13),
            )
            .into(),
            0x1000,
        ],
    );

    cpu.execute_instructions(&mut hardware, 1);
    assert_eq!(0, cpu.memory[0x1000]);
}
