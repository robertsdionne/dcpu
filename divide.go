package dcpu

func (d *DCPU) divide(pb *uint16, b, a uint16) {
	switch a {
	case 0:
		d.Extra = 1
		d.set(pb, 0)

	default:
		result := b / a
		d.Extra = 0
		d.set(pb, result)
	}
}
