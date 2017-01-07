package dcpu

func (d *DCPU) returnFromInterrupt() {
	d.RegisterA = d.Memory[d.StackPointer]
	d.StackPointer++
	d.ProgramCounter = d.Memory[d.StackPointer]
	d.StackPointer++
	d.QueueInterrupts = false
}
