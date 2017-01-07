package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_ifEqualRegisterWithEqualSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, OperandA(Literal15)),
		Basic(IfEqual, RegisterA, OperandA(Literal15)),
		Basic(Set, Push, OperandA(Literal13)),
		Basic(Set, Push, OperandA(Literal14)),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}

func TestExecuteInstructions_ifEqualRegisterWithUnequalSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, OperandA(Literal15)),
		Basic(IfEqual, RegisterA, OperandA(Literal0)),
		Basic(Set, Push, OperandA(Literal13)),
		Basic(Set, Push, OperandA(Literal14)),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 14, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}

func TestExecuteInstructions_ifEqualSkipsConditionalsWhenNotEqual(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, OperandA(Literal15)),
		Basic(IfEqual, RegisterA, OperandA(Literal0)),
		Basic(IfEqual, RegisterB, OperandA(Literal0)),
		Basic(Set, Push, OperandA(Literal12)),
		Basic(Set, Push, OperandA(Literal13)),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}

func TestExecuteInstructions_ifEqualDoesNotSkipConditionalsWhenEqual(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, OperandA(Literal15)),
		Basic(IfEqual, RegisterA, OperandA(Literal15)),
		Basic(IfEqual, RegisterB, OperandA(Literal0)),
		Basic(Set, Push, OperandA(Literal12)),
		Basic(Set, Push, OperandA(Literal13)),
	})

	dcpu.ExecuteInstructions(4)
	assert.EqualValues(t, 12, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}
