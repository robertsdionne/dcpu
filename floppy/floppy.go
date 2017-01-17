package floppy

import (
	"encoding/binary"
	"log"
	"os"

	"github.com/robertsdionne/dcpu"
)

type Device struct {
	file      *os.File
	LastError uint16
	Message   uint16
	State     uint16
}

const (
	ID             = 0x4fd524c5
	ManufacturerID = 0x1eb37e91
	Version        = 0x000b
)

// Interrupts.
const (
	Poll = iota
	SetInterruptMessage
	ReadSector
	WriteSector
)

// States.
const (
	StateNoMedia = iota
	StateReady
	StateReadyWriteProtected
	StateBusy
)

// Errors.
const (
	ErrorNone = iota
	ErrorBusy
	ErrorNoMedia
	ErrorProtected
	ErrorEject
	ErrorBadSector
	ErrorBroken
)

func (d *Device) Execute(dcpu *dcpu.DCPU) {}

func (d *Device) Insert(disk string, writeProtected bool) {
	var err error
	mode := os.O_RDWR
	if writeProtected {
		mode = os.O_RDONLY
	}

	d.file, err = os.OpenFile(disk, mode, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	d.State = StateReady
	if writeProtected {
		d.State = StateReadyWriteProtected
	}
}

func (d *Device) Eject() {
	if d.file != nil {
		d.State = StateNoMedia
		d.file.Close()
		d.file = nil
	}
}

func (d *Device) GetID() uint32 {
	return ID
}

func (d *Device) GetManufacturerID() uint32 {
	return ManufacturerID
}

func (d *Device) GetVersion() uint16 {
	return Version
}

func (d *Device) HandleHardwareInterrupt(dcpu *dcpu.DCPU) {
	switch dcpu.RegisterA {
	case Poll:
		dcpu.RegisterB = d.State
		dcpu.RegisterC = d.LastError

	case SetInterruptMessage:
		d.Message = dcpu.RegisterX

	case ReadSector:
		dcpu.RegisterB = 0

		switch d.State {
		case StateNoMedia:
			d.LastError = ErrorNoMedia
			return

		case StateBusy:
			d.LastError = ErrorBusy
			return

		default:
			dcpu.RegisterB = 1
		}

		offset := int64(1024 * (dcpu.RegisterX % 1440))
		_, err := d.file.Seek(offset, os.SEEK_SET)
		if err != nil {
			log.Fatalln(err)
		}

		err = binary.Read(d.file, binary.LittleEndian, dcpu.Memory[dcpu.RegisterY:dcpu.RegisterY+512])
		if err != nil {
			log.Fatalln(err)
		}

		if d.Message > 0 {
			dcpu.Interrupt(d.Message)
		}

	case WriteSector:
		dcpu.RegisterB = 0

		switch d.State {
		case StateNoMedia:
			d.LastError = ErrorNoMedia
			return

		case StateBusy:
			d.LastError = ErrorBusy
			return

		case StateReadyWriteProtected:
			d.LastError = ErrorProtected
			return

		default:
			dcpu.RegisterB = 1
		}

		offset := int64(1024 * (dcpu.RegisterX % 1440))
		_, err := d.file.Seek(offset, os.SEEK_SET)
		if err != nil {
			log.Fatalln(err)
		}

		err = binary.Write(d.file, binary.LittleEndian, dcpu.Memory[dcpu.RegisterY:dcpu.RegisterY+512])
		if err != nil {
			log.Fatalln(err)
		}

		err = d.file.Sync()
		if err != nil {
			log.Fatalln(err)
		}
	}
}
