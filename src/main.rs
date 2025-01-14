use clap::Parser;
use std::{error, fs};

mod clock;
mod cursive;
mod dcpu;
#[cfg(test)]
mod dcpu_test;
mod floppy;
mod hardware;
mod instructions;
mod keyboard;
mod stderr;
mod stdin;
mod stdout;

fn main() -> Result<(), Box<dyn error::Error>> {
    match Cli::parse() {
        Cli::Print { program } => print(&program),
        Cli::Terminal {
            program,
            floppy_disk,
        } => run(&program, floppy_disk),
        Cli::Cursive { program } => cursive::run(&program),
    }
}

fn print(program: &str) -> Result<(), Box<dyn error::Error>> {
    let data = fs::read(program)?;

    let data = data
        .chunks(2)
        .map(|c| match c {
            [a, b] => u16::from_le_bytes([*a, *b]),
            [b] => *b as u16,
            _ => unreachable!(),
        })
        .collect::<Vec<_>>();

    for instruction in data {
        println!(
            "{:#06x?} {} {:?}",
            instruction,
            instructions::Instruction::from(instruction).size(),
            instructions::Instruction::from(instruction)
        );
    }
    Ok(())
}

fn run(program: &str, floppy_disk: Option<String>) -> Result<(), Box<dyn error::Error>> {
    let data = fs::read(program)?;

    let mut stdin = stdin::Stdin;
    let mut stdout = stdout::Stdout;
    let mut stderr = stderr::Stderr;
    let mut clock = clock::Clock::default();
    let mut floppy = floppy::Drive::default();
    if let Some(floppy_disk) = floppy_disk {
        floppy.insert(&floppy_disk, false)?;
    }
    let mut hardware = vec![];
    hardware.push(&mut stdin as &mut dyn hardware::Hardware);
    hardware.push(&mut stdout);
    hardware.push(&mut stderr);
    hardware.push(&mut clock);
    hardware.push(&mut floppy);
    let mut dcpu = dcpu::Dcpu::default();
    dcpu.load_bytes(0, &data);

    dcpu.execute(&mut hardware);
    Ok(())
}

#[derive(clap::Parser)]
enum Cli {
    Cursive {
        #[clap(index = 1)]
        program: String,
    },
    Terminal {
        #[clap(index = 1)]
        program: String,
        #[clap(long)]
        floppy_disk: Option<String>,
    },
    Print {
        #[clap(index = 1)]
        program: String,
    },
}
