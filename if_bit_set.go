package dcpu

func (d *DCPU) ifBitSet(pb *uint16, b, a uint16) {
	result := uint32(b) << a
	d.Extra = uint16(result >> 16)
	d.set(pb, uint16(result))
}
