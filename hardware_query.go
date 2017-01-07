package dcpu

func (d *DCPU) hardwareQuery(a uint16) {
	switch {
	case int(a) < len(d.Hardware):
		hardwareID := d.Hardware[a].GetID()
		d.RegisterA = uint16(hardwareID)
		d.RegisterB = uint16(hardwareID >> 16)

		d.RegisterC = d.Hardware[a].GetVersion()

		manufacturerID := d.Hardware[a].GetManufacturerID()
		d.RegisterX = uint16(manufacturerID)
		d.RegisterY = uint16(manufacturerID >> 16)

	default:
		d.RegisterA = 0
		d.RegisterB = 0
		d.RegisterC = 0
		d.RegisterX = 0
		d.RegisterY = 0
	}
}
