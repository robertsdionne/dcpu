package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_ifLessThanRegisterWithLesserSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal30),
		Basic(IfLessThan, RegisterA, Literal15),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 14, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}

func TestExecuteInstructions_ifLessThanRegisterWithGreaterSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal15),
		Basic(IfLessThan, RegisterA, Literal30),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}
