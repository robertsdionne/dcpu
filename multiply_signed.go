package dcpu

func (d *DCPU) multiplySigned(pb *uint16, b, a uint16) {
	result := int32(int16(b)) * int32(int16(a))
	d.Extra = uint16(result >> 16)
	d.set(pb, uint16(result))
}
