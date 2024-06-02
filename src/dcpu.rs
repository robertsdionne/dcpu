use crate::hardware;

const MEMORY_SIZE: usize = 0x10000;
const BASIC_OPCODE_MASK: u16 = 0x001f;
const BASIC_VALUE_MASK_A: u16 = 0xfc00;
const BASIC_VALUE_MASK_B: u16 = 0x03e0;
const BASIC_VALUE_SHIFT_A: u16 = 0x000a;
const BASIC_VALUE_SHIFT_B: u16 = 0x0005;
const SPECIAL_OPCODE_MASK: u16 = BASIC_VALUE_MASK_B;
const SPECIAL_OPCODE_SHIFT: u16 = BASIC_VALUE_SHIFT_B;
const SPECIAL_VALUE_MASK_A: u16 = BASIC_VALUE_MASK_A;
const SPECIAL_VALUE_SHIFT_A: u16 = BASIC_VALUE_SHIFT_A;
const DEBUG_OPCODE_MASK: u16 = BASIC_VALUE_MASK_A;
const DEBUG_OPCODE_SHIFT: u16 = BASIC_VALUE_SHIFT_A;

struct Dcpu {
    register_a: u16,
    register_b: u16,
    register_c: u16,
    register_x: u16,
    register_y: u16,
    register_z: u16,
    register_i: u16,
    register_j: u16,
    program_counter: u16,
    stack_pointer: u16,
    extra: u16,
    instruction_count: u16,
    interrupt_address: u16,
    queue_interrupts: bool,
    interrupt_queue: Vec<u16>,
    memory: [u16; MEMORY_SIZE],
    hardware: Vec<Box<dyn hardware::Hardware>>,
}
