use crate::dcpu::Dcpu;
use crate::hardware;
use std::io;
use std::io::Read;

#[derive(Debug)]
pub struct Stdin;

impl hardware::Hardware for Stdin {
    fn get_id(&self) -> u32 {
        ID
    }

    fn get_manufacturer_id(&self) -> u32 {
        MANUFACTURER_ID
    }

    fn get_version(&self) -> u16 {
        VERSION
    }

    fn handle_hardware_interrupt(&mut self, dcpu: &mut Dcpu) {
        let mut input = io::stdin().lock();
        let length = dcpu.register_x;
        let start = dcpu.register_y;
        let mut buffer = vec![0; length as usize];

        let n = input.read(&mut buffer).unwrap() as u16;
        dcpu.register_z = n;

        buffer.truncate(n as usize);

        let message = Message::from(dcpu.register_a);
        match message {
            Message::ReadWordsAsBytes => {
                for i in 0..n {
                    dcpu.memory[(start + i) as usize] = buffer[i as usize] as u16;
                }
            }
            Message::ReadWords => {
                for i in (0..n).step_by(2) {
                    dcpu.memory[(start + i / 2) as usize] =
                        u16::from_be_bytes([buffer[i as usize], buffer[(i + 1) as usize]]);
                }
            }
            Message::ReadUTF8 => {
                let data = String::from_utf8(buffer).unwrap();
                let data = data.encode_utf16().collect::<Vec<u16>>();
                dcpu.memory[start as usize..(start + data.len() as u16) as usize]
                    .copy_from_slice(&data);
            }
            _ => todo!(),
        }
    }
}

const ID: u32 = 0x00000000;
const MANUFACTURER_ID: u32 = 0x76543210;
const VERSION: u16 = 0x000;

#[allow(dead_code)]
#[repr(u16)]
enum Message {
    ReadWordsAsBytes,
    ReadWords,
    ReadUTF8,
    _Unused04,
}

impl From<u16> for Message {
    fn from(value: u16) -> Self {
        unsafe { std::mem::transmute(value & 0x11) }
    }
}
