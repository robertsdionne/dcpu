package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_divideSignedRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal16),
		Basic(DivideSigned, RegisterA, Literal), 0xfffe,
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0xfff8, dcpu.RegisterA)
	assert.EqualValues(t, 0x0, dcpu.Extra)
}

func TestExecuteInstructions_divideSignedByZero(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal30),
		Basic(DivideSigned, RegisterA, Literal0),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0x0, dcpu.RegisterA)
	assert.EqualValues(t, 0x1, dcpu.Extra)
}
