use crate::{dcpu, hardware};
use std::io::{Read, Seek, Write};
use std::{error, fs, io, mem};

#[derive(Debug, Default)]
pub struct Drive {
    file: Option<fs::File>,
    last_error: u16,
    message: u16,
    state: u16,
}

impl Drive {
    pub fn insert(
        &mut self,
        disk: &str,
        write_protected: bool,
    ) -> Result<(), Box<dyn error::Error>> {
        self.file = Some(
            fs::OpenOptions::new()
                .read(true)
                .write(!write_protected)
                .open(disk)?,
        );

        self.state = if write_protected {
            State::ReadyWriteProtected
        } else {
            State::Ready
        } as u16;

        Ok(())
    }

    #[allow(dead_code)]
    pub fn eject(&mut self) -> Result<(), Box<dyn error::Error>> {
        if let Some(file) = &mut self.file {
            file.flush()?;
            self.file = None;
            self.state = State::NoMedia as u16;
        }
        Ok(())
    }
}

impl hardware::Hardware for Drive {
    fn get_id(&self) -> u32 {
        const ID: u32 = 0x4fd524c5;
        ID
    }

    fn get_manufacturer_id(&self) -> u32 {
        const MANUFACTURER_ID: u32 = 0x1eb37e91;
        MANUFACTURER_ID
    }

    fn get_version(&self) -> u16 {
        const VERSION: u16 = 0x000b;
        VERSION
    }

    fn handle_hardware_interrupt(&mut self, dcpu: &mut dcpu::Dcpu) {
        let message = Message::from(dcpu.register_a);
        match message {
            Message::Poll => {
                dcpu.register_b = self.state;
                dcpu.register_c = self.last_error;
                self.last_error = 0;
            }
            Message::SetInterruptMessage => self.message = dcpu.register_x,
            Message::ReadSector => {
                dcpu.register_b = 0;
                match State::from(self.state) {
                    State::NoMedia => {
                        self.last_error = Error::NoMedia as u16;
                        return;
                    }
                    State::Busy => {
                        self.last_error = Error::Busy as u16;
                        return;
                    }
                    _ => dcpu.register_b = 1,
                }

                self.state = State::Busy as u16;
                if self.message > 0 {
                    dcpu.interrupt(self.message);
                }

                let offset = 1024 * (dcpu.register_x as u64 % 1440);
                self.file
                    .as_mut()
                    .unwrap()
                    .seek(io::SeekFrom::Start(offset))
                    .unwrap();

                let mut data = Vec::with_capacity(1024);
                self.file.as_mut().unwrap().read_exact(&mut data).unwrap();

                let data = data
                    .chunks(2)
                    .map(|c| u16::from_be_bytes([c[0], c[1]]))
                    .collect::<Vec<_>>();

                dcpu.memory[dcpu.register_y as usize..(dcpu.register_y + 512) as usize]
                    .copy_from_slice(&data);

                self.state = if self
                    .file
                    .as_ref()
                    .unwrap()
                    .metadata()
                    .unwrap()
                    .permissions()
                    .readonly()
                {
                    State::ReadyWriteProtected
                } else {
                    State::Ready
                } as u16;
                if self.message > 0 {
                    dcpu.interrupt(self.message);
                }
            }
            Message::WriteSector => {
                dcpu.register_b = 0;
                match State::from(self.state) {
                    State::NoMedia => {
                        self.last_error = Error::NoMedia as u16;
                        return;
                    }
                    State::Busy => {
                        self.last_error = Error::Busy as u16;
                        return;
                    }
                    State::ReadyWriteProtected => {
                        self.last_error = Error::Protected as u16;
                        return;
                    }
                    State::Ready => dcpu.register_b = 1,
                }

                self.state = State::Busy as u16;
                if self.message > 0 {
                    dcpu.interrupt(self.message);
                }

                let offset = 1024 * (dcpu.register_x as u64 % 1440);
                self.file
                    .as_mut()
                    .unwrap()
                    .seek(io::SeekFrom::Start(offset))
                    .unwrap();

                let data = dcpu.memory[dcpu.register_y as usize..(dcpu.register_y + 512) as usize]
                    .iter()
                    .flat_map(|c| c.to_be_bytes())
                    .collect::<Vec<_>>();
                self.file.as_mut().unwrap().write_all(&data).unwrap();
                self.file.as_mut().unwrap().flush().unwrap();

                self.state = State::Ready as u16;
                if self.message > 0 {
                    dcpu.interrupt(self.message);
                }
            }
        }
    }
}

#[allow(dead_code)]
#[derive(Debug)]
#[repr(u16)]
enum Message {
    Poll,
    SetInterruptMessage,
    ReadSector,
    WriteSector,
}

impl From<u16> for Message {
    fn from(value: u16) -> Self {
        unsafe { mem::transmute(value & 0b11) }
    }
}

#[allow(dead_code)]
#[derive(Debug)]
#[repr(u16)]
enum State {
    NoMedia,
    Ready,
    ReadyWriteProtected,
    Busy,
}

impl From<u16> for State {
    fn from(value: u16) -> Self {
        unsafe { mem::transmute(value & 0b11) }
    }
}

#[allow(dead_code)]
#[repr(u16)]
enum Error {
    None,
    Busy,
    NoMedia,
    Protected,
    Eject,
    BadSector,
    Broken,
    _Unused07,
}

impl From<u16> for Error {
    fn from(value: u16) -> Self {
        unsafe { mem::transmute(value & 0b111) }
    }
}
