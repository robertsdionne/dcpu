use crate::dcpu::Dcpu;
use crate::hardware;

#[derive(Debug, Default)]
pub struct Device {
    last_error: u16,
    region_address: u16,
    state: u16,
    target_rotation: u16,
    vertex_count: u16,
    rotation: u16,
}

impl hardware::Hardware for Device {
    fn get_id(&self) -> u32 {
        const ID: u32 = 0x42babf3c;
        ID
    }

    fn get_manufacturer_id(&self) -> u32 {
        const MANUFACTURER_ID: u32 = 0x1eb37e91;
        MANUFACTURER_ID
    }

    fn get_version(&self) -> u16 {
        const VERSION: u16 = 0x0003;
        VERSION
    }

    fn handle_hardware_interrupt(&mut self, dcpu: &mut Dcpu) {
        match dcpu.register_a.into() {
            Message::PollDevice => {
                dcpu.register_b = self.state;
                dcpu.register_c = self.last_error;
            },
            Message::MapRegion => {
                self.region_address = dcpu.register_x;
                self.vertex_count = dcpu.register_y;
            },
            Message::RotateDevice => self.target_rotation = dcpu.register_x,
            _ => {},
        }
    }
}

#[repr(u16)]
enum Message {
    PollDevice,
    MapRegion,
    RotateDevice,
    _Unused03,
}

impl From<u16> for Message {
    fn from(value: u16) -> Self {
        unsafe { std::mem::transmute(value & 0b11) }
    }
}

#[repr(u16)]
enum State {
    NoData,
    Running,
    Turning,
    _Unused03,
}

impl From<u16> for State {
    fn from(value: u16) -> Self {
        unsafe { std::mem::transmute(value & 0b11) }
    }
}

#[repr(u16)]
enum Error {
    None,
    Broken = 0xffff,
}

impl From<u16> for Error {
    fn from(value: u16) -> Self {
        unsafe { std::mem::transmute(value & 0b1) }
    }
}
