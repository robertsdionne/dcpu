package dcpu

func (d *DCPU) interruptAddToQueue(a uint16) {
	switch {
	case a == 0:
		d.QueueInterrupts = false
	default:
		d.QueueInterrupts = true
	}
}
