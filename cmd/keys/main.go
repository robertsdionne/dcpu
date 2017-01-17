package main

import (
	"github.com/robertsdionne/dcpu"
	"github.com/robertsdionne/dcpu/keyboard"
)

func main() {
	cpu := dcpu.DCPU{}
	keys := keyboard.Device{}

	cpu.Hardware = append(cpu.Hardware, &keys)

	cpu.Load(0, []uint16{
		dcpu.Special(dcpu.InterruptAddressSet, dcpu.Literal), 0x1000,
		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal3),
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.Literal1),
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal0),
		dcpu.Debug(dcpu.DumpState),
		dcpu.Basic(dcpu.Subtract, dcpu.ProgramCounter, dcpu.Literal1),
	})

	cpu.Load(0x1000, []uint16{
		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal1),
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal0),
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.RegisterC),
		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal2),
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal0),
		dcpu.Basic(dcpu.IfEqual, dcpu.RegisterC, dcpu.Literal0),
		dcpu.Special(dcpu.ReturnFromInterrupt, dcpu.Literal0),
		dcpu.Basic(dcpu.Add, dcpu.Location, dcpu.Literal1), 0xf000,
		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Location), 0xf000,
		dcpu.Basic(dcpu.Set, dcpu.LocationOffsetByRegisterA, dcpu.RegisterB), 0xf000,
		dcpu.Debug(dcpu.Alert),
		dcpu.Special(dcpu.ReturnFromInterrupt, dcpu.Literal0),
	})

	cpu.LoadString(0xf000, " ")

	cpu.Execute()
}
