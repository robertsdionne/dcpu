package dcpu

func (d *DCPU) divideSigned(pb *uint16, b, a uint16) {
	switch int16(a) {
	case 0:
		d.Extra = 1
		d.set(pb, 0)

	default:
		result := int32(int16(b) / int16(a))
		d.Extra = 0
		d.set(pb, uint16(result))
	}
}
