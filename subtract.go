package dcpu

func (d *DCPU) subtract(pb *uint16, b, a uint16) {
	d.Extra = 0
	if b < a {
		d.Extra = 1
	}
	d.set(pb, b-a)
}
