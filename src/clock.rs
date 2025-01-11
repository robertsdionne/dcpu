use std::{mem, time};
use crate::dcpu;
use crate::hardware;

#[derive(Debug, Default)]
pub struct Clock {
    interval: u16,
    message: u16,
    ticks: u16,
    last_tick: Option<time::Instant>,
}

impl Clock {
    fn tick(&mut self, dcpu: &mut dcpu::Dcpu) {
        self.last_tick = Some(time::Instant::now());
        self.ticks = self.ticks.wrapping_add(1);
        if self.message > 0 {
            dcpu.interrupt(self.message);
        }
    }
}

impl hardware::Hardware for Clock {
    fn execute(&mut self, dcpu: &mut dcpu::Dcpu) {
        if self.interval > 0 {
            if self.last_tick.is_none()
                || self.last_tick.unwrap().elapsed() > self.interval as u32 * DURATION
            {
                self.tick(dcpu);
            }
        }
    }

    fn get_id(&self) -> u32 {
        ID
    }

    fn get_manufacturer_id(&self) -> u32 {
        0
    }

    fn get_version(&self) -> u16 {
        VERSION
    }

    fn handle_hardware_interrupt(&mut self, dcpu: &mut dcpu::Dcpu) {
        let message = Message::from(dcpu.register_a);
        match message {
            Message::SetInterval => {
                if dcpu.register_b == 0 {
                    self.ticks = 0;
                }
                self.interval = dcpu.register_b;
            }
            Message::GetTicks => dcpu.register_c = self.ticks,
            Message::SetInterruptMessage => self.message = dcpu.register_b,
            _ => todo!(),
        }
    }
}

const DURATION: time::Duration = time::Duration::from_nanos(16666666);
const ID: u32 = 0x12d0b402;
const VERSION: u16 = 0x0001;

#[allow(dead_code)]
#[repr(u16)]
pub enum Message {
    SetInterval,
    GetTicks,
    SetInterruptMessage,
    _Unused03,
}

impl From<u16> for Message {
    fn from(value: u16) -> Self {
        unsafe { mem::transmute(value & 0b11) }
    }
}
