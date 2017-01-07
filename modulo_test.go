package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_moduloRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, OperandA(Literal30)),
		Basic(Modulo, RegisterA, OperandA(Literal11)),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0x8, dcpu.RegisterA)
}
