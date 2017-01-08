package dcpu

import "time"

// Clock implements Generic Clock (compatible).
type Clock struct {
	DCPU     *DCPU
	Interval uint16
	Message  uint16
	Ticks    uint16

	lastTick time.Time
}

const (
	duration  = time.Second / frequency
	frequency = 60
	id        = 0x12d0b402
	version   = 0x0001
)

// Execute runs the clock.
func (c *Clock) Execute() {
	if c.Interval > 0 {
		now := time.Now()

		if now.Sub(c.lastTick) > time.Duration(c.Interval)*duration {
			c.lastTick = now
			c.Ticks++
			if c.DCPU != nil && c.Message > 0 {
				c.DCPU.Interrupt(c.Message)
			}
		}
	}
}

// GetID returns the Clock id.
func (c *Clock) GetID() uint32 {
	return id
}

// GetManufacturerID returns the Clock manufacturer id.
func (c *Clock) GetManufacturerID() uint32 {
	return 0
}

// GetVersion returns the Clock version.
func (c *Clock) GetVersion() uint16 {
	return version
}

// HandleHardwareInterrupt handles messages from the DCPU.
func (c *Clock) HandleHardwareInterrupt() {
	switch {
	case c.DCPU == nil:
		return

	case c.DCPU.RegisterA == 0:
		c.Interval = c.DCPU.RegisterB
		c.Ticks = 0

	case c.DCPU.RegisterA == 1:
		c.DCPU.RegisterC = c.Ticks

	case c.DCPU.RegisterA == 2:
		c.Message = c.DCPU.RegisterB
	}
}
