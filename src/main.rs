use clap::Parser;
use std::{error, fs, io};
use std::io::{IsTerminal, Write};

use dcpu::assembler;
use dcpu::clock;
use dcpu::cursive;
use dcpu::floppy;
use dcpu::hardware;
use dcpu::instructions;
use dcpu::monitor;
use dcpu::printer;
use dcpu::stderr;
use dcpu::stdin;
use dcpu::stdout;

fn main() -> Result<(), Box<dyn error::Error>> {
    match Cli::parse() {
        Cli::Assemble { program } => {
            let program = assembler::assemble_file(&program)?;
            if io::stdout().is_terminal() {
                println!("{:04x?}", program);
            } else {
                let bytes: Vec<u8> = program.iter()
                    .flat_map(|word| word.to_le_bytes())
                    .collect();
                io::stdout().write_all(&bytes)?;
            }
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
    let program = assembler::assemble_file(&program)?;

    let mut stdin = stdin::Device;
    let mut stdout = stdout::Device;
    let mut stderr = stderr::Device;
    let mut clock = clock::Device::default();
    let mut monitor = monitor::Device::default();
    let mut floppy = floppy::Device::default();
    let mut printer = printer::Device::default();
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
    hardware.push(&mut printer);
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
