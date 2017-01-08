package dcpu

import "time"

type Clock struct {
	BaseHardware

	lastTick time.Time

	DCPU     *DCPU
	Interval uint16
	Message  uint16
	Ticks    uint16
}

const (
	frequency = 60
	duration  = time.Second / frequency
)

// NewClock builds a new clock.
func NewClock() Clock {
	return Clock{
		BaseHardware: BaseHardware{
			id:             0x12d0b402,
			manufacturerID: 0x00000000,
			version:        0x0001,
		},
	}
}

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
