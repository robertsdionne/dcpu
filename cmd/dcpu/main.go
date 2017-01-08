package main

import "github.com/robertsdionne/dcpu"

func main() {
	cpu := dcpu.DCPU{}
	clock := dcpu.Clock{}

	cpu.Hardware = append(cpu.Hardware, &clock)
	clock.DCPU = &cpu

	cpu.Load(0, []uint16{
		dcpu.Special(dcpu.InterruptAddressSet, dcpu.OperandA(dcpu.Literal)), 0x1000,
		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.OperandA(dcpu.Literal0)),
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.OperandA(dcpu.Literal)), 120,
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.OperandA(dcpu.Literal0)),
		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.OperandA(dcpu.Literal2)),
		dcpu.Basic(dcpu.Set, dcpu.RegisterB, dcpu.OperandA(dcpu.Literal1)),
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.OperandA(dcpu.Literal0)),
		dcpu.Debug(dcpu.DumpState),
		dcpu.Basic(dcpu.Subtract, dcpu.ProgramCounter, dcpu.OperandA(dcpu.Literal1)),
	})

	cpu.Load(0x1000, []uint16{
		dcpu.Debug(dcpu.DumpState),
		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.OperandA(dcpu.Literal1)),
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.OperandA(dcpu.Literal0)),
		dcpu.Special(dcpu.ReturnFromInterrupt, dcpu.OperandA(dcpu.Literal0)),
	})

	cpu.Execute()
}
