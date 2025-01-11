use std::{error, io};
use std::io::Read;
use clap::Parser;

mod dcpu;
mod hardware;
mod instructions;
#[cfg(test)]
mod dcpu_test;
mod clock;
mod keyboard;
mod cursive;

fn main() -> Result<(), Box<dyn error::Error>> {
    let arguments = Cli::parse();

    match arguments {
        Cli::Terminal => run(),
        Cli::Cursive => cursive::run(),
    }
}

fn run() -> Result<(), Box<dyn error::Error>> {
    let mut input = io::stdin().lock();
    let mut data = vec![];
    input.read_to_end(&mut data)?;

    let mut clock = clock::Clock::default();
    let mut hardware = vec![&mut clock as &mut dyn hardware::Hardware];
    let mut dcpu = dcpu::Dcpu::default();
    dcpu.load_bytes(0, &data);

    dcpu.execute(&mut hardware);
    Ok(())
}

#[derive(clap::Parser)]
enum Cli {
    Cursive,
    Terminal,
}
