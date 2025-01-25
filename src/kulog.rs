use crate::dcpu::Dcpu;
use crate::hardware;

#[derive(Debug)]
pub struct Device;

impl hardware::Hardware for Device {
    fn get_id(&self) -> u32 {
        const ID: u32 = 0x140fd321;
        ID
    }

    fn get_manufacturer_id(&self) -> u32 {
        const MANUFACTURER_ID: u32 = 0x0c72a5cd;
        MANUFACTURER_ID
    }

    fn get_version(&self) -> u16 {
        const VERSION: u16 = 08581;
        VERSION
    }

    fn handle_hardware_interrupt(&mut self, _dcpu: &mut Dcpu) {
        todo!()
    }
}

#[allow(dead_code)]
#[repr(u16)]
enum Message {
    SetEnvelope1,
    SetEnvelope2,
    SetEnvelope3,
    SetEnvelope4,
    SetAttackVolumes,
    SetSquareDuties,
    SetFrequency1,
    SetFrequency2,
    SetFrequency3,
    SetFrequency4,
    SetHighpassCutoff,
    SetLowpassCutoff,
    SetFilterResVoices,
    SetWaveformPlayStop,
    _Unused0e,
    _Unused0f,
    GetEnvelope1,
    GetEnvelope2,
    GetEnvelope3,
    GetEnvelope4,
    GetAttackVolumes,
    GetSquareDuties,
    GetFrequency1,
    GetFrequency2,
    GetFrequency3,
    GetFrequency4,
    GetHighpassCutoff,
    GetLowpassCutoff,
    GetFilterResVoices,
    GetWaveformPlayStop,
    _Unused1e,
    _Unused1f,
}

impl From<u16> for Message {
    fn from(value: u16) -> Message {
        unsafe { std::mem::transmute(value & 0x1f) }
    }
}
