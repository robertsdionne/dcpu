use crate::{hardware, instructions};
use std::fmt;

const MEMORY_SIZE: usize = 0x10000;

pub struct Dcpu {
    pub register_a: u16,
    pub register_b: u16,
    pub register_c: u16,
    pub register_x: u16,
    pub register_y: u16,
    pub register_z: u16,
    pub register_i: u16,
    pub register_j: u16,
    pub program_counter: u16,
    pub stack_pointer: u16,
    pub extra: u16,
    pub instruction_count: u16,
    pub interrupt_address: u16,
    pub queue_interrupts: bool,
    pub interrupt_queue: Vec<u16>,
    pub memory: Vec<u16>,
}

impl Dcpu {
    #[allow(dead_code)]
    pub fn load_bytes(&mut self, i: usize, program: &Vec<u8>) {
        let program = program
            .chunks(2)
            .map(|c| match c {
                [a, b] => u16::from_le_bytes([*a, *b]),
                [b] => *b as u16,
                _ => unreachable!(),
            })
            .collect::<Vec<_>>();
        self.load(i, &program);
    }

    pub fn load(&mut self, i: usize, program: &Vec<u16>) {
        self.memory[i..i + program.len()].copy_from_slice(program);
    }

    #[allow(dead_code)]
    pub fn execute(&mut self, hardware: &mut [&mut dyn hardware::Hardware]) {
        loop {
            self.execute_instruction(hardware, false);
            for hardware in &mut *hardware {
                hardware.execute(self)
            }
        }
    }

    #[allow(dead_code)]
    pub fn execute_instructions(
        &mut self,
        hardware: &mut [&mut dyn hardware::Hardware],
        count: usize,
    ) {
        for _ in 0..count {
            self.execute_instruction(hardware, false);
            for hardware in &mut *hardware {
                hardware.execute(self);
            }
        }
    }

    fn execute_instruction(&mut self, hardware: &mut [&mut dyn hardware::Hardware], skip: bool) {
        use instructions::Instruction;

        if !skip && !self.queue_interrupts && self.interrupt_queue.len() > 0 {
            self.maybe_trigger_interrupt(self.interrupt_queue[0]);
            self.interrupt_queue.remove(0);
        }

        if !skip {
            self.instruction_count = self.instruction_count.wrapping_add(1)
        }

        let carry = self.extra;
        let stack_pointer = self.stack_pointer;
        let instruction_words = (0..3)
            .into_iter()
            .map(|i| self.program_counter.wrapping_add(i))
            .map(|i| self.memory[i as usize])
            .collect::<Vec<_>>();
        let instruction = Instruction::from(&instruction_words[..]);
        self.program_counter = self.program_counter.wrapping_add(1);

        if !skip {
            // println!("{:?}", instruction);
        }

        match instruction {
            Instruction::Basic(basic_opcode, operand_b, operand_a) => {
                use instructions::{BasicOpcode, OperandA};

                let a = self.get_operand(operand_a, false);
                let (_, a) = a.dereference();

                let b = self.get_operand(OperandA::Operand(operand_b), true);
                let (pb, b) = b.dereference();

                if skip {
                    self.stack_pointer = stack_pointer;
                    match basic_opcode {
                        BasicOpcode::IfBitSet
                        | BasicOpcode::IfClear
                        | BasicOpcode::IfEqual
                        | BasicOpcode::IfNotEqual
                        | BasicOpcode::IfGreaterThan
                        | BasicOpcode::IfAbove
                        | BasicOpcode::IfLessThan
                        | BasicOpcode::IfUnder => self.execute_instruction(hardware, true),
                        _ => {}
                    }

                    return;
                }

                match basic_opcode {
                    BasicOpcode::Reserved => unreachable!("BasicOpcode::Reserved"),
                    BasicOpcode::Set => {
                        if let Some(pb) = pb {
                            Self::set(pb, a);
                        }
                    }
                    BasicOpcode::Add => {
                        if let Some(pb) = pb {
                            self.extra = Self::add(pb, b, a);
                        }
                    }
                    BasicOpcode::Subtract => {
                        if let Some(pb) = pb {
                            self.extra = Self::subtract(pb, b, a);
                        }
                    }
                    BasicOpcode::Multiply => {
                        if let Some(pb) = pb {
                            self.extra = Self::multiply(pb, b, a);
                        }
                    }
                    BasicOpcode::MultiplySigned => {
                        if let Some(pb) = pb {
                            self.extra = Self::multiply_signed(pb, b, a);
                        }
                    }
                    BasicOpcode::Divide => {
                        if let Some(pb) = pb {
                            self.extra = Self::divide(pb, b, a);
                        }
                    }
                    BasicOpcode::DivideSigned => {
                        if let Some(pb) = pb {
                            self.extra = Self::divide_signed(pb, b, a);
                        }
                    }
                    BasicOpcode::Modulo => {
                        if let Some(pb) = pb {
                            Self::modulo(pb, b, a);
                        }
                    }
                    BasicOpcode::ModuloSigned => {
                        if let Some(pb) = pb {
                            Self::modulo_signed(pb, b, a);
                        }
                    }
                    BasicOpcode::BinaryAnd => {
                        if let Some(pb) = pb {
                            *pb = b & a;
                        }
                    }
                    BasicOpcode::BinaryOr => {
                        if let Some(pb) = pb {
                            *pb = b | a;
                        }
                    }
                    BasicOpcode::BinaryExclusiveOr => {
                        if let Some(pb) = pb {
                            *pb = b ^ a;
                        }
                    }
                    BasicOpcode::ShiftRight => {
                        if let Some(pb) = pb {
                            self.extra = Self::shift_right(pb, b, a);
                        }
                    }
                    BasicOpcode::ArithmeticShiftRight => {
                        if let Some(pb) = pb {
                            self.extra = Self::arithmetic_shift_right(pb, b, a);
                        }
                    }
                    BasicOpcode::ShiftLeft => {
                        if let Some(pb) = pb {
                            self.extra = Self::shift_left(pb, b, a);
                        }
                    }
                    BasicOpcode::IfBitSet => self.do_next_instruction_if(hardware, (b & a) != 0),
                    BasicOpcode::IfClear => self.do_next_instruction_if(hardware, (b & a) == 0),
                    BasicOpcode::IfEqual => self.do_next_instruction_if(hardware, b == a),
                    BasicOpcode::IfNotEqual => self.do_next_instruction_if(hardware, b != a),
                    BasicOpcode::IfGreaterThan => self.do_next_instruction_if(hardware, b > a),
                    BasicOpcode::IfAbove => {
                        self.do_next_instruction_if(hardware, (b as i16) > (a as i16))
                    }
                    BasicOpcode::IfLessThan => self.do_next_instruction_if(hardware, b < a),
                    BasicOpcode::IfUnder => {
                        self.do_next_instruction_if(hardware, (b as i16) < (a as i16))
                    }
                    BasicOpcode::SetThenIncrement => {
                        if let Some(pb) = pb {
                            *pb = a;
                            self.register_i = self.register_i.wrapping_add(1);
                            self.register_j = self.register_j.wrapping_add(1);
                        }
                    }
                    BasicOpcode::SetThenDecrement => {
                        if let Some(pb) = pb {
                            *pb = a;
                            self.register_i = self.register_i.wrapping_sub(1);
                            self.register_j = self.register_j.wrapping_sub(1);
                        }
                    }
                    BasicOpcode::AddWithCarry => {
                        if let Some(pb) = pb {
                            self.extra = Self::add_with_carry(pb, b, a, carry);
                        }
                    }
                    BasicOpcode::SubtractWithCarry => {
                        if let Some(pb) = pb {
                            self.extra = Self::subtract_with_carry(pb, b, a, carry);
                        }
                    }
                    _ => todo!(),
                }
            }
            Instruction::Special(special_opcode, operand_a) => {
                use instructions::SpecialOpcode;

                let assignable = special_opcode == SpecialOpcode::InterruptAddressGet
                    || special_opcode == SpecialOpcode::HardwareNumberConnected;

                let interrupt_address = self.interrupt_address;

                let a = self.get_operand(operand_a, assignable);
                let (pa, a) = a.dereference();

                if skip {
                    self.stack_pointer = stack_pointer;
                    return;
                }

                match special_opcode {
                    SpecialOpcode::Reserved => unreachable!("SpecialOpcode::Reserved"),
                    SpecialOpcode::JumpSubroutine => self.jump_sub_routine(a),
                    SpecialOpcode::InterruptTrigger => self.interrupt(a),
                    SpecialOpcode::InterruptAddressGet => {
                        if let Some(pa) = pa {
                            *pa = interrupt_address;
                        }
                    }
                    SpecialOpcode::InterruptAddressSet => {
                        self.interrupt_address = a;
                    }
                    SpecialOpcode::ReturnFromInterrupt => self.return_from_interrupt(),
                    SpecialOpcode::InterruptAddToQueue => self.queue_interrupts = a > 0,
                    SpecialOpcode::HardwareNumberConnected => {
                        if let Some(pa) = pa {
                            *pa = hardware.len() as u16;
                        }
                    }
                    SpecialOpcode::HardwareQuery => self.hardware_query(hardware, a),
                    SpecialOpcode::HardwareInterrupt => {
                        if let Some(hardware) = hardware.get_mut(a as usize) {
                            hardware.handle_hardware_interrupt(self);
                        }
                    }
                    _ => todo!(),
                }
            }
            Instruction::Debug(debug_opcode) => {
                use instructions::DebugOpcode;

                if skip {
                    self.stack_pointer = stack_pointer;
                    return;
                }

                match debug_opcode {
                    DebugOpcode::Noop => {}
                    DebugOpcode::Alert => {
                        let length = self.memory[0xf000];
                        if length > 0 {
                            let _alert = String::from_utf16_lossy(
                                &self.memory[0xf001..0xf001 + length as usize],
                            );
                            // println!("alert: {}", alert);
                        } else {
                            // println!("alert");
                        }
                    }
                    DebugOpcode::DumpState => {} //println!("{:04x?} {:04x?}", self, hardware),
                    _ => todo!(),
                }
            }
        }
    }

    fn set(pb: &mut u16, a: u16) {
        *pb = a;
    }

    fn add(pb: &mut u16, b: u16, a: u16) -> u16 {
        let result = b as u32 + a as u32;
        let extra = (result >> 16) as u16;
        *pb = result as u16;
        extra
    }

    fn subtract(pb: &mut u16, b: u16, a: u16) -> u16 {
        let result = b as i32 - a as i32;
        let extra = (result >> 16) as u16;
        *pb = result as u16;
        extra
    }

    fn multiply(pb: &mut u16, b: u16, a: u16) -> u16 {
        let result = b as u32 * a as u32;
        let extra = (result >> 16) as u16;
        *pb = result as u16;
        extra
    }

    fn multiply_signed(pb: &mut u16, b: u16, a: u16) -> u16 {
        let result = b as i16 as i32 * a as i16 as i32;
        let extra = (result >> 16) as u16;
        *pb = result as u16;
        extra
    }

    fn divide(pb: &mut u16, b: u16, a: u16) -> u16 {
        if a == 0 {
            *pb = 0;
            1
        } else {
            *pb = b / a;
            0
        }
    }

    fn divide_signed(pb: &mut u16, b: u16, a: u16) -> u16 {
        if a == 0 {
            *pb = 0;
            1
        } else {
            *pb = (b as i16 / a as i16) as i32 as u16;
            0
        }
    }

    fn modulo(pb: &mut u16, b: u16, a: u16) {
        *pb = b % a;
    }

    fn modulo_signed(pb: &mut u16, b: u16, a: u16) {
        *pb = (b as i16 % a as i16) as u16;
    }

    fn shift_right(pb: &mut u16, b: u16, a: u16) -> u16 {
        *pb = b >> a;
        b << (0x10 - a)
    }

    fn arithmetic_shift_right(pb: &mut u16, b: u16, a: u16) -> u16 {
        *pb = (b as i16 >> a) as u16;
        ((b as i16) << (0x10 - a)) as u16
    }

    fn shift_left(pb: &mut u16, b: u16, a: u16) -> u16 {
        let result = (b as u32) << a;
        *pb = result as u16;
        (result >> 16) as u16
    }

    fn do_next_instruction_if(
        &mut self,
        hardware: &mut [&mut dyn hardware::Hardware],
        condition: bool,
    ) {
        if !condition {
            self.execute_instruction(hardware, true);
        }
    }

    fn add_with_carry(pb: &mut u16, a: u16, b: u16, carry: u16) -> u16 {
        let result = b as u32 + a as u32 + carry as u32;
        let extra = (result >> 16) as u16;
        *pb = result as u16;
        extra
    }

    fn subtract_with_carry(pb: &mut u16, a: u16, b: u16, carry: u16) -> u16 {
        let result = b as i32 - a as i32 + carry as i32;
        let extra = (result >> 16) as u16;
        *pb = result as u16;
        extra
    }

    fn jump_sub_routine(&mut self, a: u16) {
        self.stack_pointer = self.stack_pointer.wrapping_sub(1);
        self.memory[self.stack_pointer as usize] = self.program_counter;
        self.program_counter = a;
    }

    pub fn interrupt(&mut self, a: u16) {
        if self.queue_interrupts {
            self.interrupt_queue.push(a);
        } else {
            self.maybe_trigger_interrupt(a)
        }
    }

    fn return_from_interrupt(&mut self) {
        self.register_a = self.memory[self.stack_pointer as usize];
        self.stack_pointer = self.stack_pointer.wrapping_add(1);
        self.program_counter = self.memory[self.stack_pointer as usize];
        self.stack_pointer = self.stack_pointer.wrapping_add(1);
        self.queue_interrupts = false;
    }

    fn hardware_query(&mut self, hardware: &[&mut dyn hardware::Hardware], a: u16) {
        match hardware.get(a as usize) {
            Some(hardware) => {
                let hardware_id = hardware.get_id();
                let version = hardware.get_version();
                let manufacturer_id = hardware.get_manufacturer_id();

                self.register_a = hardware_id as u16;
                self.register_b = (hardware_id >> 16) as u16;
                self.register_c = version;
                self.register_x = manufacturer_id as u16;
                self.register_y = (manufacturer_id >> 16) as u16;
            }
            None => {
                self.register_a = 0;
                self.register_b = 0;
                self.register_c = 0;
                self.register_x = 0;
                self.register_y = 0;
            }
        }
    }

    fn get_operand(&mut self, operand_a: instructions::OperandA, assignable: bool) -> Operand {
        use instructions::{OperandA, OperandB, WithPayload, WithRegister};

        match operand_a {
            OperandA::Operand(operand_b) => match operand_b {
                OperandB::WithRegister(with_register, register) => {
                    Operand::Address(match with_register {
                        WithRegister::Register => self.register_address(register),
                        WithRegister::LocationInRegister => {
                            self.location_in_register_address(register)
                        }
                    })
                }
                OperandB::WithPayload(with_payload, payload) => match with_payload {
                    WithPayload::LocationOffsetByRegister(register) => Operand::Address(
                        self.location_offset_by_register_address(payload, register),
                    ),
                    WithPayload::Pick => {
                        Operand::Address(self.address_derived_from_program_counter(
                            self.stack_pointer.wrapping_add(payload),
                        ))
                    }
                    WithPayload::Location => {
                        Operand::Address(self.address_derived_from_program_counter(payload))
                    }
                    WithPayload::Literal => {
                        self.program_counter = self.program_counter.wrapping_add(1);
                        Operand::Literal(payload)
                    }
                },
                OperandB::PushOrPop => {
                    if assignable {
                        self.stack_pointer = self.stack_pointer.wrapping_sub(1);
                        Operand::Address(&mut self.memory[self.stack_pointer as usize])
                    } else {
                        let value = Operand::Address(&mut self.memory[self.stack_pointer as usize]);
                        self.stack_pointer = self.stack_pointer.wrapping_add(1);
                        value
                    }
                }
                OperandB::Peek => Operand::Address(&mut self.memory[self.stack_pointer as usize]),
                OperandB::StackPointer => Operand::Address(&mut self.stack_pointer),
                OperandB::ProgramCounter => Operand::Address(&mut self.program_counter),
                OperandB::Extra => Operand::Address(&mut self.extra),
            },
            OperandA::SmallLiteral(literal) => Operand::Literal(literal as u16),
        }
    }

    fn register_address(&mut self, register: instructions::Register) -> &mut u16 {
        use instructions::Register;
        match register {
            Register::A => &mut self.register_a,
            Register::B => &mut self.register_b,
            Register::C => &mut self.register_c,
            Register::X => &mut self.register_x,
            Register::Y => &mut self.register_y,
            Register::Z => &mut self.register_z,
            Register::I => &mut self.register_i,
            Register::J => &mut self.register_j,
        }
    }

    fn location_in_register_address(&mut self, register: instructions::Register) -> &mut u16 {
        use instructions::Register;
        match register {
            Register::A => &mut self.memory[self.register_a as usize],
            Register::B => &mut self.memory[self.register_b as usize],
            Register::C => &mut self.memory[self.register_c as usize],
            Register::X => &mut self.memory[self.register_x as usize],
            Register::Y => &mut self.memory[self.register_y as usize],
            Register::Z => &mut self.memory[self.register_z as usize],
            Register::I => &mut self.memory[self.register_i as usize],
            Register::J => &mut self.memory[self.register_j as usize],
        }
    }

    fn location_offset_by_register_address(
        &mut self,
        location: u16,
        register: instructions::Register,
    ) -> &mut u16 {
        use instructions::Register;
        match register {
            Register::A => self.address_derived_from_program_counter(location + self.register_a),
            Register::B => self.address_derived_from_program_counter(location + self.register_b),
            Register::C => self.address_derived_from_program_counter(location + self.register_c),
            Register::X => self.address_derived_from_program_counter(location + self.register_x),
            Register::Y => self.address_derived_from_program_counter(location + self.register_y),
            Register::Z => self.address_derived_from_program_counter(location + self.register_z),
            Register::I => self.address_derived_from_program_counter(location + self.register_i),
            Register::J => self.address_derived_from_program_counter(location + self.register_j),
        }
    }

    fn address_derived_from_program_counter(&mut self, offset: u16) -> &mut u16 {
        let address = &mut self.memory[offset as usize];
        self.program_counter = self.program_counter.wrapping_add(1);
        address
    }

    fn maybe_trigger_interrupt(&mut self, interrupt: u16) {
        if self.interrupt_address == 0 {
            return;
        }

        self.queue_interrupts = true;
        self.stack_pointer = self.stack_pointer.wrapping_sub(1);
        self.memory[self.stack_pointer as usize] = self.program_counter;
        self.stack_pointer = self.stack_pointer.wrapping_sub(1);
        self.memory[self.stack_pointer as usize] = self.register_a;
        self.program_counter = self.interrupt_address;
        self.register_a = interrupt;
    }
}

enum Operand<'a> {
    Address(&'a mut u16),
    Literal(u16),
}

impl<'a> Operand<'a> {
    fn dereference(self) -> (Option<&'a mut u16>, u16) {
        match self {
            Operand::Address(address) => {
                let value = *address;
                (Some(address), value)
            }
            Operand::Literal(literal) => (None, literal),
        }
    }
}

impl Default for Dcpu {
    fn default() -> Dcpu {
        Dcpu {
            register_a: 0,
            register_b: 0,
            register_c: 0,
            register_x: 0,
            register_y: 0,
            register_z: 0,
            register_i: 0,
            register_j: 0,
            program_counter: 0,
            stack_pointer: 0,
            extra: 0,
            instruction_count: 0,
            interrupt_address: 0,
            queue_interrupts: false,
            interrupt_queue: vec![],
            memory: vec![0; MEMORY_SIZE],
        }
    }
}

impl fmt::Debug for Dcpu {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        f.debug_struct("Dcpu")
            .field("register_a", &self.register_a)
            .field("register_b", &self.register_b)
            .field("register_c", &self.register_c)
            .field("register_x", &self.register_x)
            .field("register_y", &self.register_y)
            .field("register_z", &self.register_z)
            .field("register_i", &self.register_i)
            .field("register_j", &self.register_j)
            .field("program_counter", &self.program_counter)
            .field("stack_pointer", &self.stack_pointer)
            .field("extra", &self.extra)
            .field("interrupt_address", &self.interrupt_address)
            .field("interrupt_queue", &self.interrupt_queue.len())
            .field("queue_interrupts", &self.queue_interrupts)
            .field("instruction_count", &self.instruction_count)
            .field("memory", &&self.memory[0..64])
            .finish()
    }
}
