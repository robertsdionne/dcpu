package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_divideRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, OperandA(Literal30)),
		Basic(Divide, RegisterA, OperandA(Literal16)),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0x1, dcpu.RegisterA)
	assert.EqualValues(t, 0x0, dcpu.Extra)
}

func TestExecuteInstructions_divideByZero(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, OperandA(Literal30)),
		Basic(Divide, RegisterA, OperandA(Literal0)),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0x0, dcpu.RegisterA)
	assert.EqualValues(t, 0x1, dcpu.Extra)
}
