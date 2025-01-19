use crate::dcpu::Dcpu;
use crate::hardware;
use std::{collections, mem};

#[derive(Debug, Default)]
pub struct Device {
    message: u16,
    buffer: Vec<u16>,
    state: collections::HashSet<u16>,
    previous_state: collections::HashSet<u16>,
}

impl Device {
    pub fn event(&mut self, key: u16, dcpu: &mut Dcpu) {
        self.buffer.push(key);
        self.state.insert(key);
        if self.message > 0 {
            dcpu.interrupt(self.message);
        }
    }
}

impl hardware::Hardware for Device {
    fn execute(&mut self, _dcpu: &mut Dcpu) {
        self.previous_state = self.state.clone();
        self.state.clear();
    }

    fn get_id(&self) -> u32 {
        const ID: u32 = 0x30fc7406;
        ID
    }

    fn get_manufacturer_id(&self) -> u32 {
        0
    }

    fn get_version(&self) -> u16 {
        const VERSION: u16 = 0x0001;
        VERSION
    }

    fn handle_hardware_interrupt(&mut self, dcpu: &mut Dcpu) {
        let message = Message::from(dcpu.register_a);
        match message {
            Message::ClearBuffer => self.buffer.clear(),
            Message::GetNextKey => {
                if self.buffer.len() > 0 {
                    dcpu.register_c = self.buffer[0];
                    self.buffer.remove(0);
                } else {
                    dcpu.register_c = 0;
                }
            }
            // TODO: GetKeyState seems broken.
            Message::GetKeyState => dcpu.register_c = self.state.contains(&dcpu.register_b) as u16,
            Message::SetInterruptMessage => self.message = dcpu.register_b,
        }
    }
}

#[allow(dead_code)]
#[repr(u16)]
enum Message {
    ClearBuffer,
    GetNextKey,
    GetKeyState,
    SetInterruptMessage,
}

impl From<u16> for Message {
    fn from(value: u16) -> Self {
        unsafe { mem::transmute(value & 0b11) }
    }
}
