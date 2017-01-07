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

func TestInterrupts(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Special(InterruptAddressSet, OperandA(Literal)), 0x1000,
		Basic(Set, RegisterA, OperandA(Literal13)),
		Special(InterruptTrigger, OperandA(Literal14)),
		Basic(Set, RegisterB, OperandA(RegisterA)),
	})

	dcpu.Load(0x1000, []uint16{
		Basic(Set, RegisterA, OperandA(LiteralNegative1)),
		Basic(Set, RegisterB, OperandA(LiteralNegative1)),
		Special(ReturnFromInterrupt, OperandA(Literal0)),
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 0x1000, dcpu.InterruptAddress)
	assert.EqualValues(t, 0x1000, dcpu.ProgramCounter)
	assert.True(t, dcpu.QueueInterrupts)
	assert.Len(t, dcpu.InterruptQueue, 0)

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0xffff, dcpu.RegisterA)
	assert.EqualValues(t, 0xffff, dcpu.RegisterB)
	assert.True(t, dcpu.QueueInterrupts)
	assert.Len(t, dcpu.InterruptQueue, 0)

	dcpu.Interrupt(10)
	assert.Len(t, dcpu.InterruptQueue, 1)

	dcpu.ExecuteInstructions(5)
	assert.EqualValues(t, 13, dcpu.RegisterA)
	assert.EqualValues(t, 13, dcpu.RegisterB)
	assert.EqualValues(t, 5, dcpu.ProgramCounter)
	assert.False(t, dcpu.QueueInterrupts)
	assert.Len(t, dcpu.InterruptQueue, 0)
}
