use crate::{assembler, clock, dcpu, hardware, keyboard, monitor};
use cursive;
use cursive::event::Key;
use cursive::{event, theme, view, views, CursiveExt};
use std::error;
use cursive::utils::span;

pub fn run(program: &str) -> Result<(), Box<dyn error::Error>> {
    let program = assembler::assemble_file(&program)?;

    let clock = clock::Device::default();
    let keyboard = keyboard::Device::default();
    let mut monitor = monitor::Device::default();
    monitor.border_color = 1;
    let mut dcpu = dcpu::Dcpu::default();
    dcpu.memory[0xf000..0xf080].copy_from_slice(&monitor::TEST_PATTERN);
    for i in 0..16 {
        for j in 0..16 {
            let offset = 0xf000 + (j * monitor::BUFFER_WIDTH + i) as usize;
            dcpu.memory[offset] = (dcpu.memory[offset] & 0x00ff) | (j << 12) | (i << 8);
        }
    }
    dcpu.load(0, &program);

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

                let styled_text = span::SpannedString::concatenate(
                    self.cpu.memory[video..video + 0x180].chunks(monitor::BUFFER_WIDTH as usize)
                        .map(|words| self.word_to_styled(false, words))
                );
                let styled_text = span::SpannedString::concatenate(vec![
                    self.word_to_styled(true, &[0x20; monitor::BUFFER_WIDTH as usize]),
                    styled_text,
                    self.word_to_styled(true, &[0x20; monitor::BUFFER_WIDTH as usize]),
                ]);

                self.view.set_content(styled_text);
                event::EventResult::Consumed(None)
            }
            _ => event::EventResult::Ignored,
        }
    }
}

impl View {
    fn word_to_styled(&self, is_border: bool, words: &[u16]) -> span::SpannedString<theme::Style> {
        let palette = monitor::DEFAULT_PALETTE;
        let border_color = monitor::Pixel::from(palette[(self.monitor.border_color & 0xf) as usize]);
        let border = span::SpannedString::styled(" ", theme::Style {
            color: theme::ColorStyle {
                back: border_color.clone().into(),
                ..Default::default()
            },
            ..Default::default()
        });
        let newline = span::SpannedString::styled("\n", theme::Style::default());
        let content = span::SpannedString::concatenate(
            words.iter()
                .map(|word| {
                    let bytes = [*word as u8];
                    let text = if (*word as u8 as char).is_control() {
                        ".".into()
                    } else {
                        String::from_utf8_lossy(&bytes)
                    };
                    let foreground_color = (word & 0xf000) >> 12;
                    let foreground_color = monitor::Pixel::from(palette[foreground_color as usize]);
                    let background_color = if is_border {
                        self.monitor.border_color & 0xf
                    } else {
                        (word & 0x0f00) >> 8
                    };
                    let background_color = monitor::Pixel::from(palette[background_color as usize]);
                    let blink = word & 0x80 > 0;
                    span::SpannedString::styled(text, theme::Style {
                        color: theme::ColorStyle {
                            front: foreground_color.into(),
                            back: background_color.into(),
                        },
                        effects: if blink {
                            theme::Effects::only(theme::Effect::Blink)
                        } else {
                            theme::Effects::default()
                        },
                    })
                }));
        span::SpannedString::concatenate(vec![border.clone(), content, border, newline])
    }
}

impl Into<theme::ColorType> for monitor::Pixel {
    fn into(self) -> theme::ColorType {
        theme::ColorType::Color(theme::Color::Rgb(self.r, self.g, self.b))
    }
}
