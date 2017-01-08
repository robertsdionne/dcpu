package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_multiplyRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal16),
		Basic(Multiply, RegisterA, Literal30),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0x01e0, dcpu.RegisterA)
	assert.EqualValues(t, 0x0, dcpu.Extra)
}

func TestExecuteInstructions_multiplyRegisterWithOverflow(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, LiteralNegative1),
		Basic(Multiply, RegisterA, LiteralNegative1),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0x1, dcpu.RegisterA)
	assert.EqualValues(t, 0xfffe, dcpu.Extra)
}
