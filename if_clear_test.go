package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_ifClearRegisterWithCommonBitsSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterA, OperandA(Literal30)),
		Basic(IfClear, RegisterA, OperandA(Literal16)),
		Basic(Set, Push, OperandA(Literal13)),
		Basic(Set, Push, OperandA(Literal14)),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 14, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}

func TestExecuteInstructions_ifClearRegisterWithoutCommonBitsSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterA, OperandA(Literal15)),
		Basic(IfClear, RegisterA, OperandA(Literal16)),
		Basic(Set, Push, OperandA(Literal13)),
		Basic(Set, Push, OperandA(Literal14)),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}