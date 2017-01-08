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
	duration     = time.Second / frequency
	frequency    = 60
	clockID      = 0x12d0b402
	clockVersion = 0x0001
)

const (
	clockSetInterval = iota
	clockGetTicks
	clockSetInterruptMessage
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
	return clockID
}

// GetManufacturerID returns the Clock manufacturer id.
func (c *Clock) GetManufacturerID() uint32 {
	return 0
}

// GetVersion returns the Clock version.
func (c *Clock) GetVersion() uint16 {
	return clockVersion
}

// HandleHardwareInterrupt handles messages from the DCPU.
func (c *Clock) HandleHardwareInterrupt() {
	if c.DCPU == nil {
		return
	}

	switch c.DCPU.RegisterA {
	case clockSetInterval:
		c.Interval = c.DCPU.RegisterB
		c.Ticks = 0

	case clockGetTicks:
		c.DCPU.RegisterC = c.Ticks

	case clockSetInterruptMessage:
		c.Message = c.DCPU.RegisterB
	}
}
