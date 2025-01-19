use clap::Parser;
use std::{error, fs};

mod assembler;
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
mod sped3;
mod monitor;
mod harold;
mod kulog;

fn main() -> Result<(), Box<dyn error::Error>> {
    match Cli::parse() {
        Cli::Assemble { program } => {
            let program = assembler::assemble(&program)?;
            println!("{:#06x?}", program);
            Ok(())
        }
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
    let data = (0..data.len() - 3)
        .into_iter()
        .map(|i| &data[i..i + 3])
        .collect::<Vec<_>>();

    for instruction in data {
        println!(
            "{:04x?} {} {:?}",
            instruction,
            instructions::Instruction::from(instruction).size(),
            instructions::Instruction::from(instruction)
        );
    }
    Ok(())
}

fn run(program: &str, floppy_disk: Option<String>) -> Result<(), Box<dyn error::Error>> {
    let program = assembler::assemble(&program)?;

    let mut stdin = stdin::Stdin;
    let mut stdout = stdout::Stdout;
    let mut stderr = stderr::Stderr;
    let mut clock = clock::Clock::default();
    let mut monitor = monitor::Device::default();
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
    hardware.push(&mut monitor);
    let mut dcpu = dcpu::Dcpu::default();
    dcpu.load(0, &program);

    dcpu.execute(&mut hardware);
    Ok(())
}

#[derive(clap::Parser)]
enum Cli {
    Assemble {
        #[clap(index = 1)]
        program: String,
    },
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
