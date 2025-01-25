#[allow(dead_code)]
pub mod assembler;
#[allow(dead_code)]
pub mod clock;
#[allow(dead_code)]
pub mod cursive;
#[allow(dead_code)]
mod dcpu;
#[allow(dead_code)]
pub mod floppy;
#[allow(dead_code)]
pub mod hardware;
#[allow(dead_code)]
pub mod harold;
#[allow(dead_code)]
pub mod instructions;
#[allow(dead_code)]
pub mod keyboard;
#[allow(dead_code)]
pub mod kulog;
#[allow(dead_code)]
pub mod monitor;
#[allow(dead_code)]
pub mod printer;
#[allow(dead_code)]
pub mod sleep_chamber;
#[allow(dead_code)]
pub mod sped3;
#[allow(dead_code)]
pub mod stderr;
#[allow(dead_code)]
pub mod stdin;
#[allow(dead_code)]
pub mod stdout;
#[cfg(test)]
mod tests;

pub use crate::dcpu::*;
