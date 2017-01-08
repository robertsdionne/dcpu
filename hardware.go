package dcpu

// Hardware represents peripheral devices that may interact with a DCPU.
type Hardware interface {
	Execute()
	GetID() uint32
	GetManufacturerID() uint32
	GetVersion() uint16
	HandleHardwareInterrupt()
}
