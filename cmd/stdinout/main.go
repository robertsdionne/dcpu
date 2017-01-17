package main

import (
	"github.com/robertsdionne/dcpu"
	"github.com/robertsdionne/dcpu/stdin"
	"github.com/robertsdionne/dcpu/stdout"
)

func main() {
	cpu := dcpu.DCPU{}
	in := stdin.Stdin{}
	out := stdout.Device{}

	cpu.Hardware = append(cpu.Hardware, &in, &out)

	cpu.Load(0, []uint16{
		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal0),
		dcpu.Basic(dcpu.Set, dcpu.RegisterX, dcpu.Literal), 0x0400,
		dcpu.Basic(dcpu.Set, dcpu.RegisterY, dcpu.Literal), 0x1000,
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal0),
		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal0),
		dcpu.Basic(dcpu.Set, dcpu.RegisterX, dcpu.RegisterZ),
		dcpu.Basic(dcpu.Set, dcpu.RegisterY, dcpu.Literal), 0x1000,
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal1),
		dcpu.Basic(dcpu.Set, dcpu.ProgramCounter, dcpu.Literal0),
	})

	cpu.Execute()
}
