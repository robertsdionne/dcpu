package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_shiftLeftRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, OperandA(Literal30)),
		Basic(ShiftLeft, RegisterA, OperandA(Literal2)),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0x78, dcpu.RegisterA)
	assert.Zero(t, dcpu.Extra)
}

func TestExecuteInstructions_shiftLeftRegisterWithOverflow(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, OperandA(LiteralNegative1)),
		Basic(ShiftLeft, RegisterA, OperandA(Literal2)),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0xfffc, dcpu.RegisterA)
	assert.EqualValues(t, 0x0003, dcpu.Extra)
}
