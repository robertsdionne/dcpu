package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_ifNotEqualRegisterWithUnequalSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal15),
		Basic(IfNotEqual, RegisterA, Literal0),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}

func TestExecuteInstructions_ifNotEqualRegisterWithEqualSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal15),
		Basic(IfNotEqual, RegisterA, Literal15),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 14, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}
