package stdout

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"os"
	"unicode/utf16"

	"github.com/robertsdionne/dcpu"
)

type Stdout struct{}

const (
	ID             = 0x00000001
	ManufacturerID = 0x76543210
	Version        = 0x0000
)

const (
	writeWordsAsBytes = iota
	writeBytes
	writeUTF16String
)

func (s *Stdout) Execute(dcpu *dcpu.DCPU) {}

func (s *Stdout) GetID() uint32 {
	return ID
}

func (s *Stdout) GetManufacturerID() uint32 {
	return ManufacturerID
}

func (s *Stdout) GetVersion() uint16 {
	return Version
}

func (s *Stdout) HandleHardwareInterrupt(dcpu *dcpu.DCPU) {
	length, start := dcpu.RegisterX, dcpu.RegisterY
	buffer := &bytes.Buffer{}
	var err error

	switch dcpu.RegisterA {
	case writeWordsAsBytes:
		for i := uint16(0); i < length; i++ {
			err = buffer.WriteByte(byte(dcpu.Memory[start+i]))
		}

	case writeBytes:
		for i := uint16(0); i < length; i++ {
			err = binary.Write(buffer, binary.LittleEndian, dcpu.Memory[start+i])
		}

	case writeUTF16String:
		utf16String := string(utf16.Decode(dcpu.Memory[start : start+length]))
		_, err = buffer.WriteString(utf16String)
	}
	if err != nil {
		log.Fatalln(err)
	}

	bytesWritten, err := io.Copy(os.Stdout, buffer)
	if err != nil {
		log.Fatalln(err)
	}

	dcpu.RegisterZ = uint16(bytesWritten)
}
