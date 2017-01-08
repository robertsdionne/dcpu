package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_ifUnderRegisterWithLesserSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal30),
		Basic(IfUnder, RegisterA, LiteralNegative1),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 14, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}

func TestExecuteInstructions_ifUnderRegisterWithGreaterSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, LiteralNegative1),
		Basic(IfUnder, RegisterA, Literal30),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}
