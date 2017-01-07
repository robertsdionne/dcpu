package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialState(t *testing.T) {
	dcpu := DCPU{}

	assert.EqualValues(t, 0, dcpu.RegisterA)
	assert.EqualValues(t, 0, dcpu.RegisterB)
	assert.EqualValues(t, 0, dcpu.RegisterC)
	assert.EqualValues(t, 0, dcpu.RegisterX)
	assert.EqualValues(t, 0, dcpu.RegisterY)
	assert.EqualValues(t, 0, dcpu.RegisterZ)
	assert.EqualValues(t, 0, dcpu.RegisterI)
	assert.EqualValues(t, 0, dcpu.RegisterJ)
	assert.EqualValues(t, 0, dcpu.ProgramCounter)
	assert.EqualValues(t, 0, dcpu.StackPointer)
	assert.EqualValues(t, 0, dcpu.Extra)
	assert.EqualValues(t, 0, dcpu.InterruptAddress)
	assert.EqualValues(t, 0, dcpu.Memory[1000])
}
