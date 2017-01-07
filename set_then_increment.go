package dcpu

func (d *DCPU) setThenIncrement(pb *uint16, a uint16) {
	d.set(pb, a)
	d.RegisterI++
	d.RegisterJ++
}
