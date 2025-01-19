use crate::dcpu::Dcpu;
use crate::hardware;

#[derive(Debug, Default)]
pub struct Device {
    mode: u16,
}

impl hardware::Hardware for Device {
    fn get_id(&self) -> u32 {
        const ID: u32 = 0xcff2a11d;
        ID
    }

    fn get_manufacturer_id(&self) -> u32 {
        const MANUFACTURER_ID: u32 = 0xf6976d00;
        MANUFACTURER_ID
    }

    fn get_version(&self) -> u16 {
        const VERSION: u16 = 0x0001;
        VERSION
    }

    fn handle_hardware_interrupt(&mut self, dcpu: &mut Dcpu) {
        match dcpu.register_a.into() {
            Message::SetMode => self.mode = dcpu.register_b,
            Message::GetMode => dcpu.register_a = self.mode,
            Message::CutPage => {
                const PAGE: &str = "--------------------------------------------------------------------------------";
                let width = match self.mode.into() {
                    Mode::Hex => 64,
                    Mode::DataBin | Mode::Bin => 72,
                    _ => 80,
                };
                println!("{}", &PAGE[0..width]);
                println!("{}", &PAGE[0..width]);
            }
            Message::PrintSingleLine => match self.mode.into() {
                Mode::Text => {
                    let data =
                        &dcpu.memory[dcpu.register_b as usize..dcpu.register_b as usize + 80];
                    let text = String::from_utf16_lossy(data);
                    let text: String = text
                        .chars()
                        .map(|c| if c.is_control() { '.' } else { c })
                        .collect();
                    println!("{}", text);
                }
                Mode::Data => {
                    let data =
                        &dcpu.memory[dcpu.register_b as usize..dcpu.register_b as usize + 0x10];
                    println!(
                            "{:04x?} {:04x?} {:04x?} {:04x?} {:04x?} {:04x?} {:04x?} {:04x?}  {:04x?} {:04x?} {:04x?} {:04x?} {:04x?} {:04x?} {:04x?} {:04x?}",
                            data[0x0], data[0x1], data[0x2], data[0x3], data[0x4], data[0x5], data[0x6], data[0x7],
                            data[0x8], data[0x9], data[0xa], data[0xb], data[0xc], data[0xd], data[0xe], data[0xf],
                        );
                }
                Mode::Hex => {
                    let data =
                        &dcpu.memory[dcpu.register_b as usize..dcpu.register_b as usize + 0x8];
                    let text = String::from_utf16_lossy(data);
                    let text: String = text
                        .chars()
                        .map(|c| if c.is_control() { '.' } else { c })
                        .collect();
                    println!(
                            "{:04x?}:  {:04x?} {:04x?} {:04x?} {:04x?}  {:04x?} {:04x?} {:04x?} {:04x?}: {}",
                            dcpu.register_b,
                            data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7],
                             text,
                        );
                }
                Mode::DataBin => {
                    let data =
                        &dcpu.memory[dcpu.register_b as usize..dcpu.register_b as usize + 0x4];
                    println!(
                        "{:08b} {:08b} {:08b} {:08b} {:08b} {:08b} {:08b} {:08b}",
                        (data[0] & 0xff00) >> 8,
                        data[0] & 0xff,
                        (data[1] & 0xff00) >> 8,
                        data[1] & 0xff,
                        (data[2] & 0xff00) >> 8,
                        data[2] & 0xff,
                        (data[3] & 0xff00) >> 8,
                        data[3] & 0xff,
                    );
                }
                Mode::Bin => {
                    let data =
                        &dcpu.memory[dcpu.register_b as usize..dcpu.register_b as usize + 0x3];
                    let text = String::from_utf16_lossy(data);
                    let text: String = text
                        .chars()
                        .map(|c| if c.is_control() { '.' } else { c })
                        .collect();
                    println!(
                        "{:04x?}:  {:08b} {:08b} {:08b} {:08b} {:08b} {:08b}: {}",
                        dcpu.register_b,
                        (data[0] & 0xff00) >> 8,
                        data[0] & 0xff,
                        (data[1] & 0xff00) >> 8,
                        data[1] & 0xff,
                        (data[2] & 0xff00) >> 8,
                        data[2] & 0xff,
                        text,
                    );
                }
            },
            _ => {}
        }
    }
}

#[repr(u16)]
enum Message {
    SetMode,
    GetMode,
    CutPage,
    PrintSingleLine,
    PrintMultipleLines,
    FullDump,
    BufferStatus,
    Reset = 0xffff,
}

impl From<u16> for Message {
    fn from(value: u16) -> Message {
        unsafe { std::mem::transmute(value) }
    }
}

#[repr(u16)]
enum Mode {
    Text,
    Data,
    Hex,
    DataBin,
    Bin,
}

impl From<u16> for Mode {
    fn from(value: u16) -> Mode {
        unsafe { std::mem::transmute(value) }
    }
}
