package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_ifGreaterThanRegisterWithLesserSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal30),
		Basic(IfGreaterThan, RegisterA, Literal15),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}

func TestExecuteInstructions_ifGreaterThanRegisterWithGreaterSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal15),
		Basic(IfGreaterThan, RegisterA, Literal30),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 14, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}

func TestExecuteInstructions_ifGreaterThanLiteralWithLesserRegister(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal30),
		Basic(IfGreaterThan, Literal, RegisterA), 0x0a,
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 14, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}

func TestExecuteInstructions_ifGreaterThanLiteralWithGreaterRegister(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal15),
		Basic(IfGreaterThan, Literal, RegisterA), 0x1e,
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
}
