package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_setRegisterWithRegister(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterB, Literal1),
		Basic(Set, RegisterA, RegisterB),
	})

	dcpu.ExecuteInstructions(2)

	assert.EqualValues(t, 1, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithLastRegister(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterJ, Literal1),
		Basic(Set, RegisterA, RegisterJ),
	})

	dcpu.ExecuteInstructions(2)

	assert.EqualValues(t, 1, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithLocationInRegister(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, Location, Literal13), 0x1000,
		Basic(Set, RegisterB, Literal), 0x1000,
		Basic(Set, RegisterA, LocationInRegisterB),
	})

	dcpu.ExecuteInstructions(3)

	assert.EqualValues(t, 13, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithLocationInLastRegister(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, Location, Literal13), 0x1000,
		Basic(Set, RegisterJ, Literal), 0x1000,
		Basic(Set, RegisterA, LocationInRegisterJ),
	})

	dcpu.ExecuteInstructions(3)

	assert.EqualValues(t, 13, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithLocationOffsetByRegister(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, Location, Literal13), 0x100a,
		Basic(Set, RegisterB, Literal10),
		Basic(Set, RegisterA, LocationOffsetByRegisterB), 0x1000,
	})

	dcpu.ExecuteInstructions(3)

	assert.EqualValues(t, 13, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithLocationOffsetByLastRegister(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, Location, Literal13), 0x100a,
		Basic(Set, RegisterJ, Literal10),
		Basic(Set, RegisterA, LocationOffsetByRegisterJ), 0x1000,
	})

	dcpu.ExecuteInstructions(3)

	assert.EqualValues(t, 13, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithPop(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, Push, Literal13),
		Basic(Set, RegisterA, Pop),
	})

	dcpu.ExecuteInstructions(1)

	assert.EqualValues(t, 0xffff, dcpu.StackPointer)

	dcpu.ExecuteInstructions(1)

	assert.EqualValues(t, 0, dcpu.StackPointer)
	assert.EqualValues(t, 13, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithPeek(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, Push, Literal13),
		Basic(Set, RegisterA, Peek),
	})

	dcpu.ExecuteInstructions(1)

	assert.EqualValues(t, 0xffff, dcpu.StackPointer)

	dcpu.ExecuteInstructions(1)

	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
	assert.EqualValues(t, 13, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithPick(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
		Basic(Set, RegisterA, Pick), 0x1,
	})

	assert.EqualValues(t, 0, dcpu.StackPointer)

	dcpu.ExecuteInstructions(3)

	assert.EqualValues(t, 0xfffe, dcpu.StackPointer)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer+1])
	assert.EqualValues(t, 13, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithStackPointer(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, Push, Literal13),
		Basic(Set, RegisterA, StackPointer),
	})

	assert.EqualValues(t, 0, dcpu.StackPointer)

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 0xffff, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithProgramCounter(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		0x0,
		Basic(Set, RegisterA, ProgramCounter),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 2, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithExtra(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, Extra, Literal13),
		Basic(Set, RegisterA, Extra),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 13, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithLocation(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, Location, Literal13), 0x1000,
		Basic(Set, RegisterA, Location), 0x1000,
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 13, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithLargeLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal), 0x1001,
	})

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 0x1001, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal1),
	})

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 0x1, dcpu.RegisterA)
}

func TestExecuteInstructions_setLastRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterJ, Literal1),
	})

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 0x1, dcpu.RegisterJ)
}

func TestExecuteInstructions_setLocationInRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal), 0x1000,
		Basic(Set, LocationInRegisterA, Literal13),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 13, dcpu.Memory[0x1000])
}

func TestExecuteInstructions_setLocationInLastRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterJ, Literal), 0x1000,
		Basic(Set, LocationInRegisterJ, Literal13),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 13, dcpu.Memory[0x1000])
}

func TestExecuteInstructions_setLocationOffsetByRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal10),
		Basic(Set, LocationOffsetByRegisterA, Literal13), 0x1000,
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 13, dcpu.Memory[0x100a])
}

func TestExecuteInstructions_setLocationOffsetByLastRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterJ, Literal10),
		Basic(Set, LocationOffsetByRegisterJ, Literal13), 0x1000,
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 13, dcpu.Memory[0x100a])
}

func TestExecuteInstructions_setLocationOffsetByRegisterWithLocationOffsetByRegister(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, RegisterA, Literal10),
		Basic(Set, LocationOffsetByRegisterA, Literal13), 0x1000,
		Basic(Set, LocationOffsetByRegisterA, LocationOffsetByRegisterA), 0x1000, 0x2000,
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 13, dcpu.Memory[0x200a])
}

func TestExecuteInstructions_setPushWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, Push, Literal13),
	})

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer])
}

func TestExecuteInstructions_setPushWithPop(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Pop),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer])
}

func TestExecuteInstructions_setPeekWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, Push, Literal13),
		Basic(Set, Peek, Literal14),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 14, dcpu.Memory[dcpu.StackPointer])
}

func TestExecuteInstructions_setPickWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, Push, Literal12),
		Basic(Set, Push, Literal13),
		Basic(Set, Pick, Literal14), 0x1,
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 14, dcpu.Memory[dcpu.StackPointer+1])
}

func TestExecuteInstructions_setStackPointerWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, StackPointer, Literal13),
	})

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 13, dcpu.StackPointer)
}

func TestExecuteInstructions_setProgramCounterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, ProgramCounter, Literal13),
	})

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 13, dcpu.ProgramCounter)
}

func TestExecuteInstructions_setExtraWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, Extra, Literal13),
	})

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 13, dcpu.Extra)
}

func TestExecuteInstructions_setLocationWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, Location, Literal13), 0x1000,
	})

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 13, dcpu.Memory[0x1000])
}

func TestExecuteInstructions_setLiteralWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load(0, []uint16{
		Basic(Set, Literal, Literal13), 0x1000,
	})

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 0, dcpu.Memory[0x1000])
}
