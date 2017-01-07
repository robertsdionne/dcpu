package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_ifGreaterThanRegisterWithLesserSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterA, OperandA(Literal30)),
		Basic(IfGreaterThan, RegisterA, OperandA(Literal15)),
		Basic(Set, Push, OperandA(Literal13)),
		Basic(Set, Push, OperandA(Literal14)),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer])
}

func TestExecuteInstructions_ifGreaterThanRegisterWithGreaterSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterA, OperandA(Literal15)),
		Basic(IfGreaterThan, RegisterA, OperandA(Literal30)),
		Basic(Set, Push, OperandA(Literal13)),
		Basic(Set, Push, OperandA(Literal14)),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 14, dcpu.Memory[dcpu.StackPointer])
}
