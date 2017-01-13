package main

import (
	"flag"

	"github.com/robertsdionne/dcpu"
	"github.com/robertsdionne/dcpu/clock"
	"github.com/robertsdionne/dcpu/monitor"
	"github.com/robertsdionne/dcpu/stdin"
)

var (
	borderColor     = flag.Int("border-color", 0, "the border color")
	foregroundColor = flag.Int("foreground-color", 0xf, "the foreground color")
	backgroundColor = flag.Int("background-color", 0x0, "the background color")
)

func main() {
	flag.Parse()

	cpu := dcpu.DCPU{}
	c := clock.Clock{}
	in := stdin.Stdin{}
	m := monitor.Monitor{}

	cpu.Hardware = append(cpu.Hardware, &in, &m, &c)

	cpu.Load(0, []uint16{
		dcpu.Special(dcpu.InterruptAddressSet, dcpu.Literal), 0x2000,
		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal), monitor.SetBorderColor,
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), uint16(*borderColor),
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal1),
		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal), monitor.MemoryMapScreen,
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal), 0x1000,
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal1),
		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal0),
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal20),
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal2),
		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal2),
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal1),
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal2),
		// dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal), stdin.ReadWords,
		// dcpu.Basic(dcpu.Set, dcpu.RegisterX, dcpu.Literal), 0x0300,
		// dcpu.Basic(dcpu.Set, dcpu.RegisterY, dcpu.Literal), 0x1000,
		// dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal0),
		dcpu.Basic(dcpu.Subtract, dcpu.ProgramCounter, dcpu.Literal1),
	})

	color := uint16(*foregroundColor<<12 | *backgroundColor<<8)

	cpu.Load(0x1000, []uint16{
		color | 0x48, color | 0x65, color | 0x6c, color | 0x6c, color | 0x6f, color | 0x2e,
	})

	cpu.Load(0x2000, []uint16{
		dcpu.Basic(dcpu.Add, dcpu.Location, dcpu.Literal1), 0x1000,
		dcpu.Basic(dcpu.Add, dcpu.Location, dcpu.Literal1), 0x1001,
		dcpu.Basic(dcpu.Add, dcpu.Location, dcpu.Literal1), 0x1002,
		dcpu.Basic(dcpu.Add, dcpu.Location, dcpu.Literal1), 0x1003,
		dcpu.Basic(dcpu.Add, dcpu.Location, dcpu.Literal1), 0x1004,
		dcpu.Basic(dcpu.Add, dcpu.Location, dcpu.Literal1), 0x1005,
		dcpu.Special(dcpu.ReturnFromInterrupt, dcpu.Literal0),
	})

	go cpu.Execute()
	m.Poll(&cpu)
}
