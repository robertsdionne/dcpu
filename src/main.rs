use std::{error, io};
use std::io::Read;
use crate::dcpu::Dcpu;
use crate::hardware::Hardware;

mod dcpu;
mod hardware;
mod instructions;

fn main() -> Result<(), Box<dyn error::Error>> {
    let mut input = io::stdin().lock();
    let mut data = vec![];
    input.read_to_end(&mut data)?;
    let mut monitor = Monitor;
    let mut hardware = vec![&mut monitor as &mut dyn Hardware];
    let mut dcpu = dcpu::Dcpu::default();
    dcpu.load_bytes(0, &data);
    let data2 = vec![
        10, 0,
        'a' as u8, 0,
        'b' as u8, 0,
        'c' as u8, 0,
        'd' as u8, 0,
        'e' as u8, 0,
        'f' as u8, 0,
        'g' as u8, 0,
        'h' as u8, 0,
        'i' as u8, 0,
        'j' as u8, 0,
    ];
    dcpu.load_bytes(0xf000, &data2);
    dcpu.execute_instructions(&mut hardware, 6);
    Ok(())
}

struct Monitor;

impl Hardware for Monitor {
    fn execute(&mut self, _dcpu: &mut Dcpu) {
        println!("hardware execute");
    }

    fn get_id(&self) -> u32 {
        0x1234
    }

    fn get_manufacturer_id(&self) -> u32 {
        0x2345
    }

    fn get_version(&self) -> u16 {
        0x17
    }

    fn handle_hardware_interrupt(&mut self, _dcpu: &mut Dcpu) {
        println!("hardware interrupt from hardware");
    }
}
