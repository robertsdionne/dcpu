use crate::dcpu::Dcpu;
use crate::hardware;
use std::io;
use std::io::Write;

#[derive(Debug)]
pub struct Device;

impl hardware::Hardware for Device {
    fn get_id(&self) -> u32 {
        const ID: u32 = 0x00000001;
        ID
    }

    fn get_manufacturer_id(&self) -> u32 {
        const MANUFACTURER_ID: u32 = 0x76543210;
        MANUFACTURER_ID
    }

    fn get_version(&self) -> u16 {
        const VERSION: u16 = 0x000;
        VERSION
    }

    fn handle_hardware_interrupt(&mut self, dcpu: &mut Dcpu) {
        let mut output = io::stdout().lock();
        let length = dcpu.register_x as usize;
        let start = dcpu.register_y as usize;
        let message = Message::from(dcpu.register_a);
        match message {
            Message::WriteWordsAsBytes => {
                let data = &dcpu.memory[start..start + length];
                let data: Vec<u8> = data.iter().map(|w| *w as u8).collect();
                output.write_all(&data).unwrap();
            }
            Message::WriteBytes => {
                let data = &dcpu.memory[start..start + length];
                let data: Vec<u8> = data.iter().flat_map(|w| w.to_le_bytes()).collect();
                output.write_all(&data).unwrap();
            }
            Message::WriteUTF16 => {
                let data = &dcpu.memory[start..start + length];
                let data: Vec<u8> = String::from_utf16(data).unwrap().into_bytes();
                output.write_all(&data).unwrap();
            }
            _ => todo!(),
        }
        output.flush().unwrap();
    }
}

#[allow(dead_code)]
#[repr(u16)]
enum Message {
    WriteWordsAsBytes,
    WriteBytes,
    WriteUTF16,
    _Unused04,
}

impl From<u16> for Message {
    fn from(value: u16) -> Self {
        unsafe { std::mem::transmute(value & 0x11) }
    }
}
