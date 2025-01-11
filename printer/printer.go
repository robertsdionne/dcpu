package printer

import (
	"fmt"
	"github.com/robertsdionne/dcpu"
	"unicode"
	"unicode/utf16"
)

type Device struct {
	mode uint16
}

var _ dcpu.Hardware = &Device{}

const (
	ID             = 0xcff2a11d
	ManufacturerID = 0xf6976d00
	Version        = 0x0001
)

const (
	TextMode = iota
	DataMode
	HexMode
	DataBinMode
	BinMode
)

const (
	SetMode = iota
	GetMode
	CutPage
	PrintSingleLine
	PrintMultipleLines
	FullDump
	BufferStatus
	Reset = 0xffff
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
	switch dcpu.RegisterA {
	case SetMode:
		d.mode = dcpu.RegisterB

	case GetMode:
		dcpu.RegisterA = d.mode

	case CutPage:
		const page = "--------------------------------------------------------------------------------"
		width := 80
		if d.mode == 2 {
			width = 64
		} else if d.mode == 3 || d.mode == 4 {
			width = 72
		}
		fmt.Println(page[:width])
		fmt.Println(page[:width])

	case PrintSingleLine:
		switch d.mode {
		case TextMode:
			data := dcpu.Memory[dcpu.RegisterB : dcpu.RegisterB+80]
			asciiData := make([]uint16, 0, len(data))
			for i := range data {
				asciiData = append(asciiData, data[i]&0x7f)
			}
			runes := utf16.Decode(asciiData)
			printableRunes := make([]rune, 0, len(runes))
			for _, r := range runes {
				if unicode.IsPrint(r) {
					printableRunes = append(printableRunes, r)
				}
			}
			fmt.Println(string(printableRunes))

		case DataMode:
			data := dcpu.Memory[dcpu.RegisterB : dcpu.RegisterB+0x10]
			fmt.Printf(
				"%04x %04x %04x %04x %04x %04x %04x %04x  %04x %04x %04x %04x %04x %04x %04x %04x\n",
				data[0x0], data[0x1], data[0x2], data[0x3], data[0x4], data[0x5], data[0x6], data[0x7],
				data[0x8], data[0x9], data[0xa], data[0xb], data[0xc], data[0xd], data[0xe], data[0xf],
			)

		case HexMode:
			data := dcpu.Memory[dcpu.RegisterB : dcpu.RegisterB+0x08]
			asciiData := make([]uint16, 0, len(data))
			for i := range data {
				asciiData = append(asciiData, data[i]&0x7f)
			}
			runes := utf16.Decode(asciiData)
			printableRunes := make([]rune, 0, len(runes))
			for _, r := range runes {
				if !unicode.IsPrint(r) {
					r = '.'
				}
				printableRunes = append(printableRunes, r)
			}
			fmt.Printf(
				"%04x:  %04x %04x %04x %04x  %04x %04x %04x %04x    %s\n",
				dcpu.RegisterB,
				data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7],
				string(printableRunes),
			)

		case DataBinMode:
			data := dcpu.Memory[dcpu.RegisterB : dcpu.RegisterB+0x04]
			fmt.Printf(
				"%08b %08b %08b %08b %08b %08b %08b %08b\n",
				(data[0]&0xff00)>>8,
				data[0]&0xff,
				(data[1]&0xff00)>>8,
				data[1]&0xff,
				(data[2]&0xff00)>>8,
				data[2]&0xff,
				(data[3]&0xff00)>>8,
				data[3]&0xff,
			)

		case BinMode:
			data := dcpu.Memory[dcpu.RegisterB : dcpu.RegisterB+0x03]
			asciiData := make([]uint16, 0, len(data))
			for i := range data {
				asciiData = append(asciiData, data[i]&0x7f)
			}
			runes := utf16.Decode(asciiData)
			printableRunes := make([]rune, 0, len(runes))
			for _, r := range runes {
				if !unicode.IsPrint(r) {
					r = '.'
				}
				printableRunes = append(printableRunes, r)
			}
			fmt.Printf(
				"%04x:  %08b %08b %08b %08b %08b %08b    %s\n",
				dcpu.RegisterB,
				(data[0]&0xff00)>>8,
				data[0]&0xff,
				(data[1]&0xff00)>>8,
				data[1]&0xff,
				(data[2]&0xff00)>>8,
				data[2]&0xff,
				string(printableRunes),
			)
		}

	case PrintMultipleLines:

	case FullDump:

	case BufferStatus:

	case Reset:

	}
}
