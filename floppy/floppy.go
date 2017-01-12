package floppy

import (
	"encoding/binary"
	"log"
	"os"

	"github.com/robertsdionne/dcpu"
)

type Floppy struct {
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

func (f *Floppy) Execute(dcpu *dcpu.DCPU) {}

func (f *Floppy) Insert(disk string, writeProtected bool) {
	var err error
	mode := os.O_RDWR
	if writeProtected {
		mode = os.O_RDONLY
	}

	f.file, err = os.OpenFile(disk, mode, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	f.State = StateReady
	if writeProtected {
		f.State = StateReadyWriteProtected
	}
}

func (f *Floppy) Eject() {
	if f.file != nil {
		f.State = StateNoMedia
		f.file.Close()
		f.file = nil
	}
}

func (f *Floppy) GetID() uint32 {
	return ID
}

func (f *Floppy) GetManufacturerID() uint32 {
	return ManufacturerID
}

func (f *Floppy) GetVersion() uint16 {
	return Version
}

func (f *Floppy) HandleHardwareInterrupt(dcpu *dcpu.DCPU) {
	switch dcpu.RegisterA {
	case Poll:
		dcpu.RegisterB = f.State
		dcpu.RegisterC = f.LastError

	case SetInterruptMessage:
		f.Message = dcpu.RegisterX

	case ReadSector:
		dcpu.RegisterB = 0

		switch f.State {
		case StateNoMedia:
			f.LastError = ErrorNoMedia
			return

		case StateBusy:
			f.LastError = ErrorBusy
			return

		default:
			dcpu.RegisterB = 1
		}

		offset := int64(1024 * (dcpu.RegisterX % 1440))
		_, err := f.file.Seek(offset, os.SEEK_SET)
		if err != nil {
			log.Fatalln(err)
		}

		err = binary.Read(f.file, binary.LittleEndian, dcpu.Memory[dcpu.RegisterY:dcpu.RegisterY+512])
		if err != nil {
			log.Fatalln(err)
		}

		if f.Message > 0 {
			dcpu.Interrupt(f.Message)
		}

	case WriteSector:
		dcpu.RegisterB = 0

		switch f.State {
		case StateNoMedia:
			f.LastError = ErrorNoMedia
			return

		case StateBusy:
			f.LastError = ErrorBusy
			return

		case StateReadyWriteProtected:
			f.LastError = ErrorProtected
			return

		default:
			dcpu.RegisterB = 1
		}

		offset := int64(1024 * (dcpu.RegisterX % 1440))
		_, err := f.file.Seek(offset, os.SEEK_SET)
		if err != nil {
			log.Fatalln(err)
		}

		err = binary.Write(f.file, binary.LittleEndian, dcpu.Memory[dcpu.RegisterY:dcpu.RegisterY+512])
		if err != nil {
			log.Fatalln(err)
		}

		err = f.file.Sync()
		if err != nil {
			log.Fatalln(err)
		}
	}
}
