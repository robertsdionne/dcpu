package dcpu

func (d *DCPU) add(pb *uint16, b, a uint16) {
	result := uint32(b) + uint32(a)
	d.Extra = uint16(result >> 16)
	d.set(pb, uint16(result))
}
