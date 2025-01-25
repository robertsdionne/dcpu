use crate::dcpu::Dcpu;
use crate::hardware;
use std::fs;

#[allow(dead_code)]
#[derive(Debug, Default)]
pub struct Device {
    file: Option<fs::File>,
}

impl hardware::Hardware for Device {
    fn get_id(&self) -> u32 {
        const ID: u32 = 0x74fa4cae;
        ID
    }

    fn get_manufacturer_id(&self) -> u32 {
        const MANUFACTURER_ID: u32 = 0x21544948;
        MANUFACTURER_ID
    }

    fn get_version(&self) -> u16 {
        const VERSION: u16 = 0x07c2;
        VERSION
    }

    fn handle_hardware_interrupt(&mut self, _dcpu: &mut Dcpu) {
        todo!()
    }
}

#[allow(dead_code)]
#[repr(u16)]
enum Message {
    QueryMediaPresent,
    QueryMediaParameters,
    QueryDeviceFlags,
    UpdateDeviceFlags,
    QueryInterruptType,
    SetInterruptMessage,
    ReadSectors = 0x10,
    WriteSectors = 0x11,
    QueryMediaQuality = 0xffff,
}

impl From<u16> for Message {
    fn from(value: u16) -> Message {
        unsafe { std::mem::transmute(value) }
    }
}

#[allow(dead_code)]
#[repr(u16)]
enum Flags {
    NonBlocking,
    MediaStatusInterrupt,
}

impl From<u16> for Flags {
    fn from(value: u16) -> Flags {
        unsafe { std::mem::transmute(value & 0b1) }
    }
}

#[allow(dead_code)]
#[repr(u16)]
enum Interrupt {
    None,
    MediaStatus,
    ReadComplete,
    WriteComplete,
}

#[allow(dead_code)]
#[repr(u16)]
enum MediaQuality {
    AuthenticHITMedia = 0x7fff,
    MediaFromOtherCompanies = 0xffff,
}

#[allow(dead_code)]
#[repr(u16)]
enum Error {
    None,
    NoMedia,
    InvalidSector,
    Pending,
}
