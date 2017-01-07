package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_binaryAndRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, OperandA(Literal)), 0xf0f0,
		Basic(BinaryAnd, RegisterA, OperandA(Literal)), 0x00ff,
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0x00f0, dcpu.RegisterA)
}
