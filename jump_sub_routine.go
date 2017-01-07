package dcpu

func (d *DCPU) jumpSubRoutine(a uint16) {
	d.StackPointer--
	d.Memory[d.StackPointer] = d.ProgramCounter
	d.ProgramCounter = uint16(a)
}
