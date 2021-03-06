package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_shiftRightRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal), 0xFFF0,
		Basic(ShiftRight, RegisterA, Literal2),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0x3ffc, dcpu.RegisterA)
	assert.Zero(t, dcpu.Extra)
}

func TestExecuteInstructions_shiftRightRegisterWithUnderflow(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, LiteralNegative1),
		Basic(ShiftRight, RegisterA, Literal2),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0x3fff, dcpu.RegisterA)
	assert.EqualValues(t, 0xc000, dcpu.Extra)
}
