use crate::dcpu::Dcpu;
use crate::hardware;
use std::time;

#[allow(dead_code)]
#[derive(Debug, Default)]
pub struct Device {
    pub border_color: u16,
    font_address: u16,
    palette_address: u16,
    pub video_address: u16,
    boot_png_path: Option<String>,
    start_time: Option<time::Instant>,
}

impl Device {
    pub fn paint(&self, dcpu: &Dcpu, pixels: &mut Vec<Pixel>) {
        pixels.resize_with((WIDTH * HEIGHT) as usize, Pixel::default);

        let font = if self.font_address > 0 {
            &dcpu.memory[self.font_address as usize .. self.font_address as usize + 0x100]
        } else {
            &DEFAULT_FONT[..]
        };

        let palette = if self.palette_address > 0 {
            &dcpu.memory[self.palette_address as usize .. self.palette_address as usize + 0x100]
        } else {
            &DEFAULT_PALETTE[..]
        };

        let time_to_blink = if let Some(start_time) = self.start_time {
            start_time.elapsed().as_secs_f32() % 2.0 < 1.0
        } else {
            false
        };

        for x in 0..WIDTH {
            for y in 0..HEIGHT {
                let i = (x / BORDER_WIDTH) as i16 - 1;
                let j = (y / BORDER_HEIGHT) as i16 - 1;

                let in_border = i < 0 || i == BUFFER_WIDTH as i16 || j < 0 || j == BUFFER_HEIGHT as i16;
                if in_border {
                    Self::set_pixel(x, y, palette[(self.border_color & 0xf) as usize], pixels);
                    continue;
                }

                let offset = BUFFER_WIDTH * j as u16 + i as u16;
                let character = dcpu.memory[self.video_address.wrapping_add(offset) as usize];
                let blink = character & 0x0080 > 0;
                let mut foreground_color = (character & 0xf000) >> 12;
                let mut background_color = (character & 0x0f00) >> 8;
                if blink && time_to_blink {
                    (foreground_color, background_color) = (background_color, foreground_color);
                }

                let foreground = Self::lookup_font_pixel(font, x % BORDER_WIDTH, y % BORDER_HEIGHT, character & 0x7f);
                let value = if foreground {
                    palette[foreground_color as usize]
                } else {
                    palette[background_color as usize]
                };

                Self::set_pixel(x, y, value, pixels);
            }
        }
    }

    fn set_pixel(x: u16, y: u16, value: u16, pixels: &mut [Pixel]) {
        let r = ((value & 0x0f00) >> 8) as u8;
        let r = r | (r << 4);
        let g = ((value & 0x00f0) >> 4) as u8;
        let g = g | (g << 4);
        let b = (value & 0x000f) as u8;
        let b = b | (b << 4);
        pixels[(y * WIDTH + x) as usize] = Pixel {
            r,
            g,
            b,
            a: 0xff,
        };
    }

    fn lookup_font_pixel(font: &[u16], x: u16, y: u16, index: u16) -> bool {
        let lo = font[2 * index as usize];
        let hi = font[2 * index as usize + 1];

        let mut mask = (1 << y) as u16;
        if x < 2 {
            if x == 0 {
                mask = mask << BORDER_HEIGHT;
            }

            lo & mask > 0
        } else {
            if x == 2 {
                mask = mask << BORDER_HEIGHT;
            }

            hi & mask > 0
        }
    }
}

impl hardware::Hardware for Device {
    fn get_id(&self) -> u32 {
        const ID: u32 = 0x7349f615;
        ID
    }

    fn get_manufacturer_id(&self) -> u32 {
        const MANUFACTURER_ID: u32 = 0x1c6c8b36;
        MANUFACTURER_ID
    }

    fn get_version(&self) -> u16 {
        const VERSION: u16 = 0x1802;
        VERSION
    }

    fn handle_hardware_interrupt(&mut self, dcpu: &mut Dcpu) {
        match dcpu.register_a.into() {
            Message::MemoryMapScreen => self.video_address = dcpu.register_b,
            Message::MemoryMapFont => self.font_address = dcpu.register_b,
            Message::MemoryMapPalette => self.palette_address = dcpu.register_b,
            Message::SetBorderColor => self.border_color = dcpu.register_b & 0xf,
            Message::MemoryDumpFont => {
                dcpu.memory[dcpu.register_b as usize..dcpu.register_b as usize + 0x100]
                    .copy_from_slice(&DEFAULT_FONT[..]);
            }
            Message::MemoryDumpPalette => {
                dcpu.memory[dcpu.register_b as usize..dcpu.register_b as usize + 0x10]
                    .copy_from_slice(&DEFAULT_PALETTE[..]);
            }
            _ => {}
        }
    }
}

pub const WIDTH: u16 = AREA_WIDTH + 2 * BORDER_WIDTH;
pub const HEIGHT: u16 = AREA_HEIGHT + 2 * BORDER_HEIGHT;
const AREA_WIDTH: u16 = 128;
const AREA_HEIGHT: u16 = 96;
const BORDER_WIDTH: u16 = 4;
const BORDER_HEIGHT: u16 = 8;
pub const BUFFER_WIDTH: u16 = 32;
pub const BUFFER_HEIGHT: u16 = 12;
const SCALE: u16 = 3;
const TITLE: &'static str = "LEM1802";
const BOOT_PNG: &'static str = "documents/boot.png";

const DEFAULT_FONT: [u16; 8 * 32] = [
    0xb79e, 0x388e, 0x722c, 0x75f4, 0x19bb, 0x7f8f, 0x85f9, 0xb158, 0x242e, 0x2400, 0x082a, 0x0800,
    0x0008, 0x0000, 0x0808, 0x0808, 0x00ff, 0x0000, 0x00f8, 0x0808, 0x08f8, 0x0000, 0x080f, 0x0000,
    0x000f, 0x0808, 0x00ff, 0x0808, 0x08f8, 0x0808, 0x08ff, 0x0000, 0x080f, 0x0808, 0x08ff, 0x0808,
    0x6633, 0x99cc, 0x9933, 0x66cc, 0xfef8, 0xe080, 0x7f1f, 0x0701, 0x0107, 0x1f7f, 0x80e0, 0xf8fe,
    0x5500, 0xaa00, 0x55aa, 0x55aa, 0xffaa, 0xff55, 0x0f0f, 0x0f0f, 0xf0f0, 0xf0f0, 0x0000, 0xffff,
    0xffff, 0x0000, 0xffff, 0xffff, 0x0000, 0x0000, 0x005f, 0x0000, 0x0300, 0x0300, 0x3e14, 0x3e00,
    0x266b, 0x3200, 0x611c, 0x4300, 0x3629, 0x7650, 0x0002, 0x0100, 0x1c22, 0x4100, 0x4122, 0x1c00,
    0x1408, 0x1400, 0x081c, 0x0800, 0x4020, 0x0000, 0x0808, 0x0800, 0x0040, 0x0000, 0x601c, 0x0300,
    0x3e49, 0x3e00, 0x427f, 0x4000, 0x6259, 0x4600, 0x2249, 0x3600, 0x0f08, 0x7f00, 0x2745, 0x3900,
    0x3e49, 0x3200, 0x6119, 0x0700, 0x3649, 0x3600, 0x2649, 0x3e00, 0x0024, 0x0000, 0x4024, 0x0000,
    0x0814, 0x2200, 0x1414, 0x1400, 0x0022, 0x1408, 0x0259, 0x0600, 0x3e59, 0x5e00, 0x7e09, 0x7e00,
    0x7f49, 0x3600, 0x3e41, 0x2200, 0x7f41, 0x3e00, 0x7f49, 0x4100, 0x7f09, 0x0100, 0x3e41, 0x7a00,
    0x7f08, 0x7f00, 0x417f, 0x4100, 0x2040, 0x3f00, 0x7f08, 0x7700, 0x7f40, 0x4000, 0x7f06, 0x7f00,
    0x7f01, 0x7e00, 0x3e41, 0x3e00, 0x7f09, 0x0600, 0x3e61, 0x7e00, 0x7f09, 0x7600, 0x2649, 0x3200,
    0x017f, 0x0100, 0x3f40, 0x7f00, 0x1f60, 0x1f00, 0x7f30, 0x7f00, 0x7708, 0x7700, 0x0778, 0x0700,
    0x7149, 0x4700, 0x007f, 0x4100, 0x031c, 0x6000, 0x0041, 0x7f00, 0x0201, 0x0200, 0x8080, 0x8000,
    0x0001, 0x0200, 0x2454, 0x7800, 0x7f44, 0x3800, 0x3844, 0x2800, 0x3844, 0x7f00, 0x3854, 0x5800,
    0x087e, 0x0900, 0x4854, 0x3c00, 0x7f04, 0x7800, 0x447d, 0x4000, 0x2040, 0x3d00, 0x7f10, 0x6c00,
    0x417f, 0x4000, 0x7c18, 0x7c00, 0x7c04, 0x7800, 0x3844, 0x3800, 0x7c14, 0x0800, 0x0814, 0x7c00,
    0x7c04, 0x0800, 0x4854, 0x2400, 0x043e, 0x4400, 0x3c40, 0x7c00, 0x1c60, 0x1c00, 0x7c30, 0x7c00,
    0x6c10, 0x6c00, 0x4c50, 0x3c00, 0x6454, 0x4c00, 0x0836, 0x4100, 0x0077, 0x0000, 0x4136, 0x0800,
    0x0201, 0x0201, 0x0205, 0x0200,
];

const DEFAULT_PALETTE: [u16; 16] = [
    0x0000, // black
    0x0008, // navy
    0x0080, // green
    0x0088, // teal
    0x0800, // maroon
    0x080f, // purple
    0x0880, // olive
    0x0ccc, // silver
    0x0888, // gray
    0x000f, // blue
    0x00f0, // lime
    0x00ff, // cyan
    0x0f00, // red
    0x0f0f, // fuchsia
    0x0ff0, // yellow
    0x0fff, // white
];

pub const TEST_PATTERN: [u16; 128] = [
    0xf000, 0xf001, 0xf002, 0xf003, 0xf004, 0xf005, 0xf006, 0xf007, 0xf008, 0xf009, 0xf00a, 0xf00b,
    0xf00c, 0xf00d, 0xf00e, 0xf00f, 0xf010, 0xf011, 0xf012, 0xf013, 0xf014, 0xf015, 0xf016, 0xf017,
    0xf018, 0xf019, 0xf01a, 0xf01b, 0xf01c, 0xf01d, 0xf01e, 0xf01f, 0xf020, 0xf021, 0xf022, 0xf023,
    0xf024, 0xf025, 0xf026, 0xf027, 0xf028, 0xf029, 0xf02a, 0xf02b, 0xf02c, 0xf02d, 0xf02e, 0xf02f,
    0xf030, 0xf031, 0xf032, 0xf033, 0xf034, 0xf035, 0xf036, 0xf037, 0xf038, 0xf039, 0xf03a, 0xf03b,
    0xf03c, 0xf03d, 0xf03e, 0xf03f, 0xf040, 0xf041, 0xf042, 0xf043, 0xf044, 0xf045, 0xf046, 0xf047,
    0xf048, 0xf049, 0xf04a, 0xf04b, 0xf04c, 0xf04d, 0xf04e, 0xf04f, 0xf050, 0xf051, 0xf052, 0xf053,
    0xf054, 0xf055, 0xf056, 0xf057, 0xf058, 0xf059, 0xf05a, 0xf05b, 0xf05c, 0xf05d, 0xf05e, 0xf05f,
    0xf060, 0xf061, 0xf062, 0xf063, 0xf064, 0xf065, 0xf066, 0xf067, 0xf068, 0xf069, 0xf06a, 0xf06b,
    0xf06c, 0xf06d, 0xf06e, 0xf06f, 0xf070, 0xf071, 0xf072, 0xf073, 0xf074, 0xf075, 0xf076, 0xf077,
    0xf078, 0xf079, 0xf07a, 0xf07b, 0xf07c, 0xf07d, 0xf07e, 0xf07f,
];

#[repr(u16)]
enum Message {
    MemoryMapScreen,
    MemoryMapFont,
    MemoryMapPalette,
    SetBorderColor,
    MemoryDumpFont,
    MemoryDumpPalette,
    _Unused06,
    _Unused07,
}

impl From<u16> for Message {
    fn from(value: u16) -> Self {
        unsafe { std::mem::transmute(value & 0b111) }
    }
}

#[repr(C)]
#[derive(Clone, Debug, Default)]
pub struct Pixel {
    r: u8,
    g: u8,
    b: u8,
    a: u8,
}
