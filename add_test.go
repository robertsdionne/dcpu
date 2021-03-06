package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_addRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal13),
		Basic(Add, RegisterA, Literal14),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0x1b, dcpu.RegisterA)
	assert.EqualValues(t, 0x0, dcpu.Extra)
}

func TestExecuteInstructions_addRegisterWithOverflow(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, LiteralNegative1),
		Basic(Add, RegisterA, LiteralNegative1),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0xfffe, dcpu.RegisterA)
	assert.EqualValues(t, 0x1, dcpu.Extra)
}
