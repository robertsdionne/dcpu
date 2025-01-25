use crate::{assembler, clock, dcpu, hardware, keyboard, monitor};
use cursive;
use cursive::event::Key;
use cursive::{event, view, views, CursiveExt};
use std::error;

pub fn run(program: &str) -> Result<(), Box<dyn error::Error>> {
    let program = assembler::assemble_file(&program)?;

    let clock = clock::Device::default();
    let keyboard = keyboard::Device::default();
    let monitor = monitor::Device::default();
    let mut dcpu = dcpu::Dcpu::default();
    dcpu.load(0, &program);

    let data = vec![0x20; 0x180];
    dcpu.load(0xf000, &data);

    let mut siv = cursive::Cursive::new();
    siv.add_layer(View {
        view: views::TextView::new("dcpu-16"),
        cpu: dcpu,
        clock,
        keyboard,
        monitor,
    });
    siv.set_autorefresh(true);

    siv.run();

    Ok(())
}

struct View {
    view: views::TextView,
    cpu: dcpu::Dcpu,
    clock: clock::Device,
    keyboard: keyboard::Device,
    monitor: monitor::Device,
}

impl view::ViewWrapper for View {
    cursive::wrap_impl!(self.view: views::TextView);

    fn wrap_on_event(&mut self, event: event::Event) -> event::EventResult {
        match event {
            event::Event::Char(char) => {
                self.keyboard.event(char as u16, &mut self.cpu);
                event::EventResult::Consumed(None)
            }
            event::Event::Key(key_event) => {
                let key = match key_event {
                    Key::Backspace => 0x10,
                    Key::Enter => 0x11,
                    Key::Ins => 0x12,
                    Key::Del => 0x13,
                    Key::Up => 0x80,
                    Key::Down => 0x81,
                    Key::Left => 0x82,
                    Key::Right => 0x83,
                    _ => 0,
                };
                self.keyboard.event(key, &mut self.cpu);
                event::EventResult::Consumed(None)
            }
            event::Event::Refresh => {
                let mut hardware = vec![&mut self.clock as &mut dyn hardware::Hardware];
                hardware.push(&mut self.keyboard);
                hardware.push(&mut self.monitor);
                self.cpu.execute_instructions(&mut hardware, 3333);
                let video = self.monitor.video_address as usize;
                let text: Vec<u8> = self.cpu.memory[video..video + 0x180].iter()
                    .map(|word| *word as u8)
                    .collect();

                self.view.set_content(format!(
                    "{}\n{}\n{}\n{}\n{}\n{}\n{}\n{}\n{}\n{}\n{}\n{}",
                    String::from_utf8_lossy(&text[..0x20]),
                    String::from_utf8_lossy(&text[0x20..0x40]),
                    String::from_utf8_lossy(&text[0x40..0x60]),
                    String::from_utf8_lossy(&text[0x60..0x80]),
                    String::from_utf8_lossy(&text[0x80..0xa0]),
                    String::from_utf8_lossy(&text[0xa0..0xc0]),
                    String::from_utf8_lossy(&text[0xc0..0xe0]),
                    String::from_utf8_lossy(&text[0xe0..0x100]),
                    String::from_utf8_lossy(&text[0x100..0x120]),
                    String::from_utf8_lossy(&text[0x120..0x140]),
                    String::from_utf8_lossy(&text[0x140..0x160]),
                    String::from_utf8_lossy(&text[0x160..0x180]),
                ));
                event::EventResult::Consumed(None)
            }
            _ => event::EventResult::Ignored,
        }
    }
}
