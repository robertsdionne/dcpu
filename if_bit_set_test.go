package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_ifBitSetRegisterWithCommonBitsSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterA, OperandA(Literal30)),
		Basic(IfBitSet, RegisterA, OperandA(Literal16)),
		Basic(Set, Push, OperandA(Literal13)),
		Basic(Set, Push, OperandA(Literal14)),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer])
}

func TestExecuteInstructions_ifBitSetRegisterWithoutCommonBitsSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterA, OperandA(Literal15)),
		Basic(IfBitSet, RegisterA, OperandA(Literal16)),
		Basic(Set, Push, OperandA(Literal13)),
		Basic(Set, Push, OperandA(Literal14)),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 14, dcpu.Memory[dcpu.StackPointer])
}
