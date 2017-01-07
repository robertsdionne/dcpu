package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_binaryOrRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, OperandA(Literal)), 0xf0f0,
		Basic(BinaryOr, RegisterA, OperandA(Literal)), 0x00ff,
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0xf0ff, dcpu.RegisterA)
}
