package main

import (
	"flag"

	"github.com/robertsdionne/dcpu"
	"github.com/robertsdionne/dcpu/keyboard"
	"github.com/robertsdionne/dcpu/monitor"
)

var (
	borderColor     = flag.Int("border-color", 0, "the border color")
	foregroundColor = flag.Int("foreground-color", 0xf, "the foreground color")
	backgroundColor = flag.Int("background-color", 0x0, "the background color")
)

const (
	printCharacter  = 0x2100
	advanceCursor   = 0x2200
	deleteCharacter = 0x2300
	newline         = 0x2400
	shiftCharacter  = 0x2500
	clearScreen     = 0x2600
	cursor          = 0x009f
	shiftFlag       = 0x0f00
)

func main() {
	flag.Parse()
	clearColor := uint16(*backgroundColor<<12 | *backgroundColor<<8)
	color := uint16(*foregroundColor<<12 | *backgroundColor<<8)

	cpu := dcpu.DCPU{}
	k := keyboard.Device{}
	m := monitor.Device{}

	cpu.Hardware = append(cpu.Hardware, &k, &m)

	cpu.Load(0, []uint16{
		dcpu.Special(dcpu.InterruptAddressSet, dcpu.Literal), 0x2000,

		dcpu.Basic(dcpu.Set, dcpu.RegisterI, dcpu.Literal), 0x0180,
		dcpu.Special(dcpu.JumpSubRoutine, dcpu.Literal), clearScreen,

		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal), monitor.SetBorderColor,
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), uint16(*borderColor),
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal1),

		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal), monitor.MemoryMapScreen,
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 0x1000,
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal1),

		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal), keyboard.SetInterruptMessage,
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal1),
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal0),

		dcpu.Basic(dcpu.Subtract, dcpu.ProgramCounter, dcpu.Literal1),
	})

	cpu.Load(0x2000, []uint16{
		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal), keyboard.GetNextKey,
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal0),

		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.RegisterC),

		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal), keyboard.GetKeyState,
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal0),

		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 0x0090,
		dcpu.Basic(dcpu.Set, dcpu.Location, dcpu.RegisterC), shiftFlag,
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 0x0090,
		dcpu.Special(dcpu.ReturnFromInterrupt, dcpu.Literal0),

		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterC, dcpu.Literal0),
		dcpu.Special(dcpu.ReturnFromInterrupt, dcpu.Literal0),

		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 0x0011,
		dcpu.Special(dcpu.JumpSubRoutine, dcpu.Literal), newline,
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 0x0011,
		dcpu.Special(dcpu.ReturnFromInterrupt, dcpu.Literal0),

		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 0x0010,
		dcpu.Special(dcpu.JumpSubRoutine, dcpu.Literal), deleteCharacter,

		dcpu.Basic(dcpu.IfNotEqual, dcpu.RegisterB, dcpu.Literal), 0x0010,
		dcpu.Special(dcpu.JumpSubRoutine, dcpu.Literal), printCharacter,

		dcpu.Special(dcpu.ReturnFromInterrupt, dcpu.Literal0),
	})

	cpu.Load(0x1000, []uint16{color | cursor})

	cpu.Load(printCharacter, []uint16{
		dcpu.Basic(dcpu.IfGreaterThan, dcpu.Location, dcpu.Literal0), shiftFlag,
		dcpu.Special(dcpu.JumpSubRoutine, dcpu.Literal), shiftCharacter,
		dcpu.Basic(dcpu.BinaryOr, dcpu.RegisterB, dcpu.Literal), color,
		dcpu.Basic(dcpu.Set, dcpu.LocationOffsetByRegisterI, dcpu.RegisterB), 0x1000,
		dcpu.Basic(dcpu.Set, dcpu.LocationOffsetByRegisterI, dcpu.Literal), 0x1001, color | cursor,
		dcpu.Basic(dcpu.Set, dcpu.RegisterJ, dcpu.Literal1),
		dcpu.Special(dcpu.JumpSubRoutine, dcpu.Literal), advanceCursor,
		dcpu.Basic(dcpu.Set, dcpu.ProgramCounter, dcpu.Pop),
	})

	cpu.Load(advanceCursor, []uint16{
		dcpu.Basic(dcpu.Add, dcpu.RegisterI, dcpu.RegisterJ),
		dcpu.Basic(dcpu.IfGreaterThan, dcpu.RegisterI, dcpu.Literal), 0x17f,
		dcpu.Basic(dcpu.Set, dcpu.RegisterI, dcpu.Literal0),
		dcpu.Basic(dcpu.IfUnder, dcpu.RegisterI, dcpu.Literal0),
		dcpu.Basic(dcpu.Set, dcpu.RegisterI, dcpu.Literal0),
		dcpu.Basic(dcpu.Set, dcpu.ProgramCounter, dcpu.Pop),
	})

	cpu.Load(deleteCharacter, []uint16{
		dcpu.Basic(dcpu.Set, dcpu.LocationOffsetByRegisterI, dcpu.Literal), 0x1000, clearColor,
		dcpu.Basic(dcpu.Set, dcpu.RegisterJ, dcpu.LiteralNegative1),
		dcpu.Special(dcpu.JumpSubRoutine, dcpu.Literal), advanceCursor,
		dcpu.Basic(dcpu.IfEqual, dcpu.LocationOffsetByRegisterI, dcpu.Literal), 0x1000, clearColor,
		dcpu.Basic(dcpu.IfAbove, dcpu.RegisterI, dcpu.Literal0),
		dcpu.Basic(dcpu.Set, dcpu.ProgramCounter, dcpu.Literal), deleteCharacter,
		dcpu.Basic(dcpu.Set, dcpu.LocationOffsetByRegisterI, dcpu.Literal), 0x1000, color | cursor,
		dcpu.Basic(dcpu.Set, dcpu.ProgramCounter, dcpu.Pop),
	})

	cpu.Load(newline, []uint16{
		dcpu.Basic(dcpu.Set, dcpu.LocationOffsetByRegisterI, dcpu.Literal), 0x1000, clearColor,
		dcpu.Basic(dcpu.Set, dcpu.Push, dcpu.RegisterI),
		dcpu.Basic(dcpu.Modulo, dcpu.Peek, dcpu.Literal), 0x0020,
		dcpu.Basic(dcpu.Set, dcpu.RegisterJ, dcpu.Literal), 0x0020,
		dcpu.Basic(dcpu.Subtract, dcpu.RegisterJ, dcpu.Pop),
		dcpu.Special(dcpu.JumpSubRoutine, dcpu.Literal), advanceCursor,
		dcpu.Basic(dcpu.Set, dcpu.LocationOffsetByRegisterI, dcpu.Literal), 0x1000, color | cursor,
		dcpu.Basic(dcpu.Set, dcpu.ProgramCounter, dcpu.Pop),
	})

	cpu.Load(shiftCharacter, []uint16{
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), '`',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '~',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), '1',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '!',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), '2',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '@',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), '3',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '#',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), '4',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '$',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), '5',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '%',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), '6',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '^',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), '7',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '&',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), '8',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '*',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), '9',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '(',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), '0',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), ')',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), '-',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '_',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), '=',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '+',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'q',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'Q',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'w',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'W',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'e',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'E',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'r',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'R',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 't',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'T',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'y',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'Y',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'u',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'U',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'i',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'I',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'o',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'O',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'p',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'P',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), '[',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '{',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), ']',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '}',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), '\\',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '|',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'a',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'A',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 's',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'S',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'd',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'D',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'f',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'F',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'g',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'G',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'h',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'H',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'j',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'J',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'k',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'K',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'l',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'L',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), ';',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), ':',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), '\'',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '"',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'z',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'Z',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'x',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'X',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'c',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'C',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'v',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'V',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'b',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'B',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'n',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'N',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 'm',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 'M',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), ',',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '<',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), '.',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '>',
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), '/',
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), '?',
		dcpu.Basic(dcpu.Set, dcpu.ProgramCounter, dcpu.Pop),
	})

	cpu.Load(clearScreen, []uint16{
		dcpu.Basic(dcpu.Set, dcpu.LocationOffsetByRegisterI, dcpu.Literal), 0x1000, clearColor,
		dcpu.Basic(dcpu.Subtract, dcpu.RegisterI, dcpu.Literal1),
		dcpu.Basic(dcpu.IfNotEqual, dcpu.RegisterI, dcpu.Literal0),
		dcpu.Basic(dcpu.Set, dcpu.ProgramCounter, dcpu.Literal), clearScreen,
		dcpu.Basic(dcpu.Set, dcpu.ProgramCounter, dcpu.Pop),
	})

	go cpu.Execute()
	m.Poll(&cpu)
}
