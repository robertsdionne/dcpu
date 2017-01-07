package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_multiplyRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterA, OperandA(Literal16)),
		Basic(Multiply, RegisterA, OperandA(Literal30)),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0x01e0, dcpu.RegisterA)
	assert.EqualValues(t, 0x0, dcpu.Extra)
}

func TestExecuteInstructions_multiplyRegisterWithOverflow(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterA, OperandA(LiteralNegative1)),
		Basic(Multiply, RegisterA, OperandA(LiteralNegative1)),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0x1, dcpu.RegisterA)
	assert.EqualValues(t, 0xfffe, dcpu.Extra)
}
