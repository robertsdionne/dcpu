package stdin

import (
	"encoding/binary"
	"log"
	"os"
	"unicode/utf16"

	"github.com/robertsdionne/dcpu"
)

type Device struct{}

const (
	ID             = 0x00000000
	ManufacturerID = 0x76543210
	Version        = 0x0000
)

const (
	ReadBytesAsWords = iota
	ReadWords
	ReadUTF8String
)

func (d *Device) Execute(dcpu *dcpu.DCPU) {}

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
	length, start := dcpu.RegisterX, dcpu.RegisterY
	buffer := make([]byte, length)

	bytesRead, err := os.Stdin.Read(buffer)
	if err != nil {
		log.Fatalln(err)
	}

	dcpu.RegisterZ = uint16(bytesRead)

	switch dcpu.RegisterA {
	case ReadBytesAsWords:
		for i, value := range buffer {
			dcpu.Memory[start+uint16(i)] = uint16(value)
		}

	case ReadWords:
		for i := 0; i < int(length); i += 2 {
			dcpu.Memory[start+uint16(i/2)] = binary.LittleEndian.Uint16(buffer[i : i+2])
		}

	case ReadUTF8String:
		copy(dcpu.Memory[start:], utf16.Encode([]rune(string(buffer))))
	}
}
