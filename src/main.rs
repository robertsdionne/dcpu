use std::{error, fs};
use clap::Parser;

mod clock;
mod cursive;
mod dcpu;
#[cfg(test)]
mod dcpu_test;
mod hardware;
mod instructions;
mod keyboard;
mod stderr;
mod stdin;
mod stdout;
mod floppy;

fn main() -> Result<(), Box<dyn error::Error>> {
    match Cli::parse() {
        Cli::Terminal { program, floppy_disk } => run(&program, floppy_disk),
        Cli::Cursive { program } => cursive::run(&program),
    }
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
}
