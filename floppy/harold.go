package floppy

import (
	"github.com/robertsdionne/dcpu"
	"os"
)

type Harold struct {
	file *os.File
}

var _ dcpu.Hardware = &Harold{}

const (
	HaroldID             = 0x74fa4cae
	HaroldManufacturerID = 0x21544948
	HaroldVersion        = 0x07c2
)

// Interrupts.
const (
	HaroldQueryMediaPresent = iota
	HaroldQueryMediaParameters
	HaroldQueryDeviceFlags
	HaroldUpdateDeviceFlags
	HaroldQueryInterruptType
	HaroldSetInterruptMessage
	HaroldReadSectors       = 0x10
	HaroldWriteSectors      = 0x11
	HaroldQueryMediaQuality = 0xffff
)

// Device Flags.
const (
	HaroldNonBlocking = 1 << iota
	HaroldMediaStatusInterrupt
)

// Interrupt Type.
const (
	HaroldNone = iota
	HaroldMediaStatus
	HaroldReadComplete
	HaroldWriteComplete
)

// Media Quality.
const (
	HaroldAuthenticHITMedia       = 0x7fff
	HaroldMediaFromOtherCompanies = 0xffff
)

// Errors.
const (
	HaroldErrorNone = iota
	HaroldErrorNoMedia
	HaroldErrorInvalidSector
	HaroldErrorPending
)

func (h *Harold) Execute(dcpu *dcpu.DCPU) {}

func (h *Harold) Insert(disk string) {}

func (h *Harold) Eject() {}

func (h *Harold) GetID() uint32 {
	return 0
}

func (h *Harold) GetManufacturerID() uint32 {
	return 0
}

func (h *Harold) GetVersion() uint16 {
	return 0
}

func (h *Harold) HandleHardwareInterrupt(dcpu *dcpu.DCPU) {
}
