package dcpu

func (d *DCPU) shiftRight(pb *uint16, b, a uint16) {
	result := b >> a
	d.Extra = uint16(b << (0x10 - a))
	d.set(pb, result)
}
