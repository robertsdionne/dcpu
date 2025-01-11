use std::{error, fs};
use clap::Parser;

mod dcpu;
mod hardware;
mod instructions;
#[cfg(test)]
mod dcpu_test;
mod clock;
mod keyboard;
mod cursive;
mod stdout;
mod stdin;

fn main() -> Result<(), Box<dyn error::Error>> {
    match Cli::parse() {
        Cli::Terminal { program } => run(&program),
        Cli::Cursive { program } => cursive::run(&program),
    }
}

fn run(program: &str) -> Result<(), Box<dyn error::Error>> {
    let data = fs::read(program)?;

    let mut stdin = stdin::Stdin;
    let mut stdout = stdout::Stdout;
    let mut clock = clock::Clock::default();
    let mut hardware = vec![&mut stdin as &mut dyn hardware::Hardware];
    hardware.push(&mut stdout);
    hardware.push(&mut clock);
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
    },
}
