use std::{error, io};
use std::io::Read;

mod dcpu;
mod hardware;
mod instructions;
#[cfg(test)]
mod dcpu_test;
mod clock;

fn main() -> Result<(), Box<dyn error::Error>> {
    let mut input = io::stdin().lock();
    let mut data = vec![];
    input.read_to_end(&mut data)?;

    let mut clock = clock::Clock::default();
    let mut hardware = vec![&mut clock as &mut dyn hardware::Hardware];
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

    dcpu.execute(&mut hardware);
    Ok(())
}
