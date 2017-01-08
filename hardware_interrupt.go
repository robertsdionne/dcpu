package dcpu

func (d *DCPU) hardwareInterrupt(a uint16) {
	if int(a) < len(d.Hardware) {
		d.Hardware[a].HandleHardwareInterrupt(d)
	}
}
