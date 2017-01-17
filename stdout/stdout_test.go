package stdout

import (
	"testing"

	"github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestStdout(t *testing.T) {
	d := dcpu.DCPU{}
	stdout := Device{}

	d.Hardware = append(d.Hardware, &stdout)

	d.Load(0, []uint16{
		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal2),
		dcpu.Basic(dcpu.Set, dcpu.RegisterX, dcpu.Location), 0x0fff,
		dcpu.Basic(dcpu.Set, dcpu.RegisterY, dcpu.Literal), 0x1000,
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal0),
	})

	d.LoadString(0x0fff, "Hello there\n")

	d.ExecuteInstructions(4)
	assert.EqualValues(t, 0x000c, d.RegisterZ)
}
