package stdin

import (
	"encoding/binary"
	"io"
	"log"
	"os"
	"unicode/utf16"

	"github.com/robertsdionne/dcpu"
)

type Stdin struct{}

const (
	ID             = 0x00000000
	ManufacturerID = 0x76543210
	Version        = 0x0000
)

const (
	readBytesAsWords = iota
	readWords
	readUTF8String
)

func (s *Stdin) Execute(dcpu *dcpu.DCPU) {}

func (s *Stdin) GetID() uint32 {
	return ID
}

func (s *Stdin) GetManufacturerID() uint32 {
	return ManufacturerID
}

func (s *Stdin) GetVersion() uint16 {
	return Version
}

func (s *Stdin) HandleHardwareInterrupt(dcpu *dcpu.DCPU) {
	length, start := dcpu.RegisterX, dcpu.RegisterY
	buffer := make([]byte, length)

	bytesRead, err := os.Stdin.Read(buffer)
	if err != nil && err != io.EOF {
		log.Fatalln(err)
	}

	dcpu.RegisterZ = uint16(bytesRead)

	switch dcpu.RegisterA {
	case readBytesAsWords:
		for i, value := range buffer {
			dcpu.Memory[start+uint16(i)] = uint16(value)
		}

	case readWords:
		for i := 0; i < int(length); i += 2 {
			dcpu.Memory[start+uint16(i)] = binary.LittleEndian.Uint16(buffer[i : i+2])
		}

	case readUTF8String:
		copy(dcpu.Memory[start:], utf16.Encode([]rune(string(buffer))))
	}
}
