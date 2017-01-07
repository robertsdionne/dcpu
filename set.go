package dcpu

func (d *DCPU) set(pb *uint16, a uint16) {
	if pb != nil {
		*pb = a
	}
}
