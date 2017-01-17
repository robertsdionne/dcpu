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

func main() {
	flag.Parse()
	color := uint16(*foregroundColor<<12 | *backgroundColor<<8)

	cpu := dcpu.DCPU{}
	k := keyboard.Keyboard{}
	m := monitor.Device{}

	cpu.Hardware = append(cpu.Hardware, &k, &m)

	cpu.Load(0, []uint16{
		dcpu.Special(dcpu.InterruptAddressSet, dcpu.Literal), 0x2000,

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

	const (
		printCharacter  = 0x2100
		advanceCursor   = 0x2200
		deleteCharacter = 0x2300
	)

	cpu.Load(0x2000, []uint16{
		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal), keyboard.GetNextKey,
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal0),

		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.RegisterC),

		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal), keyboard.GetKeyState,
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal0),

		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterC, dcpu.Literal0),
		dcpu.Special(dcpu.ReturnFromInterrupt, dcpu.Literal0),

		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterB, dcpu.Literal), 0x0010,
		dcpu.Special(dcpu.JumpSubRoutine, dcpu.Literal), deleteCharacter,

		dcpu.Basic(dcpu.IfNotEqual, dcpu.RegisterB, dcpu.Literal), 0x0010,
		dcpu.Special(dcpu.JumpSubRoutine, dcpu.Literal), printCharacter,

		dcpu.Special(dcpu.ReturnFromInterrupt, dcpu.Literal0),
	})

	cpu.Load(0x1000, monitor.TestPattern)

	cpu.Load(printCharacter, []uint16{
		dcpu.Basic(dcpu.BinaryOr, dcpu.RegisterB, dcpu.Literal), color,
		dcpu.Basic(dcpu.Set, dcpu.LocationOffsetByRegisterI, dcpu.RegisterB), 0x1000,
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
		dcpu.Basic(dcpu.Set, dcpu.RegisterJ, dcpu.LiteralNegative1),
		dcpu.Special(dcpu.JumpSubRoutine, dcpu.Literal), advanceCursor,
		dcpu.Basic(dcpu.Set, dcpu.LocationOffsetByRegisterI, dcpu.Literal0), 0x1000,
		dcpu.Basic(dcpu.Set, dcpu.ProgramCounter, dcpu.Pop),
	})

	go cpu.Execute()
	m.Poll(&cpu)
}
