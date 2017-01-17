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

type Device struct{}

const (
	ID             = 0x00000001
	ManufacturerID = 0x76543210
	Version        = 0x0000
)

const (
	WriteWordsAsBytes = iota
	WriteBytes
	WriteUTF16String
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
	buffer := &bytes.Buffer{}
	var err error

	switch dcpu.RegisterA {
	case WriteWordsAsBytes:
		for i := uint16(0); i < length; i++ {
			err = buffer.WriteByte(byte(dcpu.Memory[start+i]))
		}

	case WriteBytes:
		for i := uint16(0); i < length; i++ {
			err = binary.Write(buffer, binary.LittleEndian, dcpu.Memory[start+i])
		}

	case WriteUTF16String:
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
