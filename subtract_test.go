package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_subtractRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal30),
		Basic(Subtract, RegisterA, Literal16),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0xe, dcpu.RegisterA)
	assert.EqualValues(t, 0x0, dcpu.Extra)
}

func TestExecuteInstructions_subtractRegisterWithUnderflow(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal16),
		Basic(Subtract, RegisterA, Literal30),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0xfff2, dcpu.RegisterA)
	assert.EqualValues(t, 0x1, dcpu.Extra)
}
