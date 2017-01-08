package clock

import (
	"time"

	"github.com/robertsdionne/dcpu"
)

// Clock implements Generic Clock (compatible).
type Clock struct {
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
func (c *Clock) Execute(dcpu *dcpu.DCPU) {
	if c.Interval > 0 {
		now := time.Now()

		if now.Sub(c.lastTick) > time.Duration(c.Interval)*duration {
			c.lastTick = now
			c.Ticks++
			if dcpu != nil && c.Message > 0 {
				dcpu.Interrupt(c.Message)
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
func (c *Clock) HandleHardwareInterrupt(dcpu *dcpu.DCPU) {
	switch dcpu.RegisterA {
	case setInterval:
		c.Interval = dcpu.RegisterB
		c.Ticks = 0

	case getTicks:
		dcpu.RegisterC = c.Ticks

	case setInterruptMessage:
		c.Message = dcpu.RegisterB
	}
}
