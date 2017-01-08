// Package dcpu implements an emulator for Notch's DCPU 1.7 specification.
package dcpu

// Hardware represents peripheral devices that may interact with a DCPU.
type Hardware interface {
	Execute()
	GetID() uint32
	GetManufacturerID() uint32
	GetVersion() uint16
	HandleHardwareInterrupt()
}

// BaseHardware provides id, manufacturerID and version.
type BaseHardware struct {
	id             uint32
	manufacturerID uint32
	version        uint16
}

// GetID returns the id of this hardware.
func (b *BaseHardware) GetID() uint32 {
	return b.id
}

// GetManufacturerID returns the manufacturer of this hardware.
func (b *BaseHardware) GetManufacturerID() uint32 {
	return b.manufacturerID
}

// GetVersion returns the version of this hardware.
func (b *BaseHardware) GetVersion() uint16 {
	return b.version
}
