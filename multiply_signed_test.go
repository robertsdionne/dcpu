package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_multiplySignedRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterA, OperandA(LiteralNegative1)),
		Basic(MultiplySigned, RegisterA, OperandA(Literal16)),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0xfff0, dcpu.RegisterA)
	assert.EqualValues(t, 0xffff, dcpu.Extra)
}

func TestExecuteInstructions_multiplySignedRegisterWithOverflow(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterA, OperandA(LiteralNegative1)),
		Basic(MultiplySigned, RegisterA, OperandA(LiteralNegative1)),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0x0001, dcpu.RegisterA)
	assert.EqualValues(t, 0x0, dcpu.Extra)
}
