use crate::{assembler, clock, dcpu, hardware, keyboard};
use cursive;
use cursive::event::Key;
use cursive::{event, view, views, CursiveExt};
use std::error;

pub fn run(program: &str) -> Result<(), Box<dyn error::Error>> {
    let program = assembler::assemble(&program)?;

    let clock = clock::Clock::default();
    let keyboard = keyboard::Keyboard::default();
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
    });
    siv.set_autorefresh(true);

    siv.run();

    Ok(())
}

struct View {
    view: views::TextView,
    cpu: dcpu::Dcpu,
    clock: clock::Clock,
    keyboard: keyboard::Keyboard,
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
                hardware.push(&mut self.keyboard as &mut dyn hardware::Hardware);
                self.cpu.execute_instructions(&mut hardware, 3333);
                self.view.set_content(format!(
                    "{}\n{}\n{}\n{}\n{}\n{}\n{}\n{}\n{}\n{}\n{}\n{}",
                    String::from_utf16_lossy(&self.cpu.memory[0xf000..0xf020]),
                    String::from_utf16_lossy(&self.cpu.memory[0xf020..0xf040]),
                    String::from_utf16_lossy(&self.cpu.memory[0xf040..0xf060]),
                    String::from_utf16_lossy(&self.cpu.memory[0xf060..0xf080]),
                    String::from_utf16_lossy(&self.cpu.memory[0xf080..0xf0a0]),
                    String::from_utf16_lossy(&self.cpu.memory[0xf0a0..0xf0c0]),
                    String::from_utf16_lossy(&self.cpu.memory[0xf0c0..0xf0e0]),
                    String::from_utf16_lossy(&self.cpu.memory[0xf0e0..0xf100]),
                    String::from_utf16_lossy(&self.cpu.memory[0xf100..0xf120]),
                    String::from_utf16_lossy(&self.cpu.memory[0xf120..0xf140]),
                    String::from_utf16_lossy(&self.cpu.memory[0xf140..0xf160]),
                    String::from_utf16_lossy(&self.cpu.memory[0xf160..0xf180]),
                ));
                event::EventResult::Consumed(None)
            }
            _ => event::EventResult::Ignored,
        }
    }
}
