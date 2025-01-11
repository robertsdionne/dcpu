use crate::dcpu;

pub trait Hardware {
    fn execute(&mut self, dcpu: &mut dcpu::Dcpu);
    fn get_id(&self) -> u32;
    fn get_manufacturer_id(&self) -> u32;
    fn get_version(&self) -> u16;
    fn handle_hardware_interrupt(&mut self, dcpu: &mut dcpu::Dcpu);
}
