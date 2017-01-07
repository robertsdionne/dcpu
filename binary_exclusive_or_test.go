package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_binaryExclusiveOrRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, OperandA(Literal)), 0xf0f0,
		Basic(BinaryExclusiveOr, RegisterA, OperandA(Literal)), 0x00ff,
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0xf00f, dcpu.RegisterA)
}
