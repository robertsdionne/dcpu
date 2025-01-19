use std::mem;
use crate::dcpu::Dcpu;
use crate::hardware;

#[derive(Debug)]
pub struct Device;

impl hardware::Hardware for Device {
    fn get_id(&self) -> u32 {
        const ID: u32 = 0x40e41d9d;
        ID
    }

    fn get_manufacturer_id(&self) -> u32 {
        const MANUFACTURER_ID: u32 = 0x1c6c8b36;
        MANUFACTURER_ID
    }

    fn get_version(&self) -> u16 {
        const VERSION: u16 = 0x005e;
        VERSION
    }

    fn handle_hardware_interrupt(&mut self, dcpu: &mut Dcpu) {
        todo!()
    }
}

#[repr(u16)]
enum Message {
    GetStatus,
    SetUnitToSkip,
    TriggerDevice,
    SetSkipUnit,
}

impl From<u16> for Message {
    fn from(value: u16) -> Message {
        unsafe { mem::transmute(value & 0x3) }
    }
}

#[repr(u16)]
enum Status {
    EvacuateVesselImmediately,
    NotInAVacuum,
    NotEnoughFuel,
    InhomogeneousGravitationalField,
    TooMuchAngularMomentum,
    OneOrMoreCellDoorsAreOpen,
    MechanicalError,
    UnknownError = 0xffff,
}

#[repr(u16)]
enum Unit {
    Milliseconds,
    Minutes,
    Days,
    Years,
}
