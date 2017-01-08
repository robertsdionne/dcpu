package clock

import (
	"time"

	"github.com/robertsdionne/dcpu"
)

// Clock implements Generic Clock (compatible).
type Clock struct {
	DCPU     *dcpu.DCPU
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

const (
	setInterval = iota
	getTicks
	setInterruptMessage
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
	if c.DCPU == nil {
		return
	}

	switch c.DCPU.RegisterA {
	case setInterval:
		c.Interval = c.DCPU.RegisterB
		c.Ticks = 0

	case getTicks:
		c.DCPU.RegisterC = c.Ticks

	case setInterruptMessage:
		c.Message = c.DCPU.RegisterB
	}
}
