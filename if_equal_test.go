package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_ifEqualRegisterWithEqualSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal15),
		Basic(IfEqual, RegisterA, Literal15),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}

func TestExecuteInstructions_ifEqualRegisterWithUnequalSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal15),
		Basic(IfEqual, RegisterA, Literal0),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 14, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}

func TestExecuteInstructions_ifEqualSkipsConditionalsWhenNotEqual(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal15),
		Basic(IfEqual, RegisterA, Literal0),
		Basic(IfEqual, RegisterB, Literal0),
		Basic(Set, Push, Literal12),
		Basic(Set, Push, Literal13),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}

func TestExecuteInstructions_ifEqualDoesNotSkipConditionalsWhenEqual(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal15),
		Basic(IfEqual, RegisterA, Literal15),
		Basic(IfEqual, RegisterB, Literal0),
		Basic(Set, Push, Literal12),
		Basic(Set, Push, Literal13),
	})

	dcpu.ExecuteInstructions(4)
	assert.EqualValues(t, 12, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}
