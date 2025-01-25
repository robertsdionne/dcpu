use crate::dcpu::Dcpu;
use crate::hardware;
use std::f32::consts;

#[derive(Debug, Default)]
pub struct Device {
    last_error: u16,
    region_address: u16,
    state: u16,
    target_rotation: u16,
    vertex_count: u16,
    rotation: f32,
}

impl Device {
    pub fn update(&mut self) {
        self.rotation = 0.99 * self.rotation + 0.01 * (self.target_rotation % 360) as f32;
    }

    pub fn paint(
        &self,
        dcpu: &Dcpu,
    ) -> (Vec<[f32; 7]>, Vec<[f32; 7]>, Vec<[f32; 7]>, Vec<[f32; 7]>) {
        let (red, green, blue) = self.beam(dcpu);
        let lines = self.lines(dcpu);

        (red, green, blue, lines)
    }

    fn beam(&self, dcpu: &Dcpu) -> (Vec<[f32; 7]>, Vec<[f32; 7]>, Vec<[f32; 7]>) {
        let alpha = 1.0f32 / 2.0 / (self.vertex_count as f32).sqrt();
        let offset = -4.0 * consts::PI / 3.0 + consts::PI / 2.0;
        let theta0 = offset;
        let theta1 = 2.0 * consts::PI / 3.0 + offset;
        let theta2 = 4.0 * consts::PI / 3.0 + offset;
        let origin = [
            [
                theta0.cos() / 2.0,
                0.0,
                theta0.sin() / 2.0,
                0.0,
                0.0,
                0.0,
                alpha,
            ],
            [
                theta1.cos() / 2.0,
                0.0,
                theta1.sin() / 2.0,
                0.0,
                0.0,
                0.0,
                alpha,
            ],
            [
                theta2.cos() / 2.0,
                0.0,
                theta2.sin() / 2.0,
                0.0,
                0.0,
                0.0,
                alpha,
            ],
        ];

        let mut red = vec![];
        let mut green = vec![];
        let mut blue = vec![];

        for i in 0..self.vertex_count {
            let offset = (self.region_address + 2 * i) as usize;
            let words = &dcpu.memory[offset..offset + i as usize];

            let color0 = words[1] >> 8;
            let color1 = words[3] >> 8;

            let v0 = self.build_vertex(words[0], words[1], alpha);
            let v1 = self.build_vertex(words[2], words[3], alpha);

            let v0_black = [v0[0], v0[1], v0[2], 0.0, 0.0, 0.0, v0[6]];
            let v1_black = [v1[0], v1[1], v1[2], 0.0, 0.0, 0.0, v1[6]];

            for s in 0..2 {
                let c = 1u16 << s;
                let r = (c == 4) as u16 as f32;
                let g = (c == 2) as u16 as f32;
                let b = (c == 1) as u16 as f32;
                let v0_prime = [v0[0], v0[1], v0[2], r, g, b, v0[6]];
                let v1_prime = [v1[0], v1[1], v1[2], r, g, b, v1[6]];

                match c {
                    4 => {
                        red.push(if color0 & c > 0 { v0_prime } else { v0_black });
                        red.push(if color1 & c > 0 { v1_prime } else { v1_black });
                        red.push(origin[s]);
                    }
                    2 => {
                        green.push(if color0 & c > 0 { v0_prime } else { v0_black });
                        green.push(if color1 & c > 0 { v1_prime } else { v1_black });
                        green.push(origin[s]);
                    }
                    1 => {
                        blue.push(if color0 & c > 0 { v0_prime } else { v0_black });
                        blue.push(if color1 & c > 0 { v1_prime } else { v1_black });
                        blue.push(origin[s]);
                    }
                    _ => unreachable!(),
                }
            }
        }

        (red, green, blue)
    }

    fn build_vertex(&self, word0: u16, word1: u16, alpha: f32) -> [f32; 7] {
        let x = (word0 & 0xff) as f32 / 255.0 - 0.5;
        let y = (word0 >> 8) as f32 / 255.0 - 0.5;
        let z = (word1 & 0xff) as f32 / 255.0 - 0.5;
        let color = word1 >> 8;
        let r = ((0x4 & color) >> 3) as f32;
        let g = ((0x2 & color) >> 1) as f32;
        let b = (0x1 & color) as f32;
        let theta = consts::PI / 180.0 + self.rotation;
        [
            x * theta.cos() + z * theta.sin(),
            y,
            x * theta.sin() - z * theta.cos(),
            r,
            g,
            b,
            alpha,
        ]
    }

    fn lines(&self, dcpu: &Dcpu) -> Vec<[f32; 7]> {
        let mut lines = vec![];
        for i in 0..self.vertex_count {
            let offset = (self.region_address + 2 * i) as usize;
            lines.push(self.build_vertex(dcpu.memory[offset], dcpu.memory[offset + 1], 0.8));
        }
        lines
    }
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
            }
            Message::MapRegion => {
                self.region_address = dcpu.register_x;
                self.vertex_count = dcpu.register_y;
            }
            Message::RotateDevice => self.target_rotation = dcpu.register_x,
            _ => {}
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
