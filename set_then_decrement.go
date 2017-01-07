package dcpu

func (d *DCPU) setThenDecrement(pb *uint16, a uint16) {
	d.set(pb, a)
	d.RegisterI--
	d.RegisterJ--
}
