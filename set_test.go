package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_setRegisterWithRegister(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterB, Literal1),
		Basic(Set, RegisterA, OperandA(RegisterB)),
	})

	dcpu.ExecuteInstructions(2)

	assert.EqualValues(t, 1, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithLastRegister(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterJ, Literal1),
		Basic(Set, RegisterA, OperandA(RegisterJ)),
	})

	dcpu.ExecuteInstructions(2)

	assert.EqualValues(t, 1, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithLocationInRegister(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, Location, Literal13), 0x1000,
		Basic(Set, RegisterB, OperandA(Literal)), 0x1000,
		Basic(Set, RegisterA, OperandA(LocationInRegisterB)),
	})

	dcpu.ExecuteInstructions(3)

	assert.EqualValues(t, 13, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithLocationInLastRegister(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, Location, Literal13), 0x1000,
		Basic(Set, RegisterJ, OperandA(Literal)), 0x1000,
		Basic(Set, RegisterA, OperandA(LocationInRegisterJ)),
	})

	dcpu.ExecuteInstructions(3)

	assert.EqualValues(t, 13, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithLocationOffsetByRegister(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, Location, Literal13), 0x100a,
		Basic(Set, RegisterB, OperandA(Literal10)),
		Basic(Set, RegisterA, OperandA(LocationOffsetByRegisterB)), 0x1000,
	})

	dcpu.ExecuteInstructions(3)

	assert.EqualValues(t, 13, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithLocationOffsetByLastRegister(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, Location, Literal13), 0x100a,
		Basic(Set, RegisterJ, OperandA(Literal10)),
		Basic(Set, RegisterA, OperandA(LocationOffsetByRegisterJ)), 0x1000,
	})

	dcpu.ExecuteInstructions(3)

	assert.EqualValues(t, 13, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithPop(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, Push, OperandA(Literal13)),
		Basic(Set, RegisterA, OperandA(Pop)),
	})

	dcpu.ExecuteInstructions(1)

	assert.EqualValues(t, 0xffff, dcpu.StackPointer)

	dcpu.ExecuteInstructions(1)

	assert.EqualValues(t, 0, dcpu.StackPointer)
	assert.EqualValues(t, 13, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithPeek(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, Push, OperandA(Literal13)),
		Basic(Set, RegisterA, OperandA(Peek)),
	})

	dcpu.ExecuteInstructions(1)

	assert.EqualValues(t, 0xffff, dcpu.StackPointer)

	dcpu.ExecuteInstructions(1)

	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
	assert.EqualValues(t, 13, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithPick(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, Push, OperandA(Literal13)),
		Basic(Set, Push, OperandA(Literal14)),
		Basic(Set, RegisterA, OperandA(Pick)), 0x1,
	})

	assert.EqualValues(t, 0, dcpu.StackPointer)

	dcpu.ExecuteInstructions(3)

	assert.EqualValues(t, 0xfffe, dcpu.StackPointer)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer+1])
	assert.EqualValues(t, 13, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithStackPointer(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, Push, OperandA(Literal13)),
		Basic(Set, RegisterA, OperandA(StackPointer)),
	})

	assert.EqualValues(t, 0, dcpu.StackPointer)

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 0xffff, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithProgramCounter(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		0x0,
		Basic(Set, RegisterA, OperandA(ProgramCounter)),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 2, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithExtra(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, Extra, OperandA(Literal13)),
		Basic(Set, RegisterA, OperandA(Extra)),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 13, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithLocation(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, Location, OperandA(Literal13)), 0x1000,
		Basic(Set, RegisterA, OperandA(Location)), 0x1000,
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 13, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithLargeLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterA, OperandA(Literal)), 0x1001,
	})

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 0x1001, dcpu.RegisterA)
}

func TestExecuteInstructions_setRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterA, OperandA(Literal1)),
	})

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 0x1, dcpu.RegisterA)
}

func TestExecuteInstructions_setLastRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterJ, OperandA(Literal1)),
	})

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 0x1, dcpu.RegisterJ)
}

func TestExecuteInstructions_setLocationInRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterA, OperandA(Literal)), 0x1000,
		Basic(Set, LocationInRegisterA, OperandA(Literal13)),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 13, dcpu.Memory[0x1000])
}

func TestExecuteInstructions_setLocationInLastRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterJ, OperandA(Literal)), 0x1000,
		Basic(Set, LocationInRegisterJ, OperandA(Literal13)),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 13, dcpu.Memory[0x1000])
}

func TestExecuteInstructions_setLocationOffsetByRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterA, OperandA(Literal10)),
		Basic(Set, LocationOffsetByRegisterA, OperandA(Literal13)), 0x1000,
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 13, dcpu.Memory[0x100a])
}

func TestExecuteInstructions_setLocationOffsetByLastRegisterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterJ, OperandA(Literal10)),
		Basic(Set, LocationOffsetByRegisterJ, OperandA(Literal13)), 0x1000,
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 13, dcpu.Memory[0x100a])
}

func TestExecuteInstructions_setLocationOffsetByRegisterWithLocationOffsetByRegister(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, RegisterA, OperandA(Literal10)),
		Basic(Set, LocationOffsetByRegisterA, OperandA(Literal13)), 0x1000,
		Basic(Set, LocationOffsetByRegisterA, OperandA(LocationOffsetByRegisterA)), 0x1000, 0x2000,
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 13, dcpu.Memory[0x200a])
}

func TestExecuteInstructions_setPushWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, Push, OperandA(Literal13)),
	})

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer])
}

func TestExecuteInstructions_setPushWithPop(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, Push, OperandA(Literal13)),
		Basic(Set, Push, OperandA(Pop)),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 0xffff, dcpu.StackPointer)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer])
}

func TestExecuteInstructions_setPeekWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, Push, OperandA(Literal13)),
		Basic(Set, Peek, OperandA(Literal14)),
	})

	dcpu.ExecuteInstructions(2)
	assert.EqualValues(t, 14, dcpu.Memory[dcpu.StackPointer])
}

func TestExecuteInstructions_setPickWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, Push, OperandA(Literal12)),
		Basic(Set, Push, OperandA(Literal13)),
		Basic(Set, Pick, OperandA(Literal14)), 0x1,
	})

	dcpu.ExecuteInstructions(3)
	assert.EqualValues(t, 13, dcpu.Memory[dcpu.StackPointer])
	assert.EqualValues(t, 14, dcpu.Memory[dcpu.StackPointer+1])
}

func TestExecuteInstructions_setStackPointerWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, StackPointer, OperandA(Literal13)),
	})

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 13, dcpu.StackPointer)
}

func TestExecuteInstructions_setProgramCounterWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, ProgramCounter, OperandA(Literal13)),
	})

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 13, dcpu.ProgramCounter)
}

func TestExecuteInstructions_setExtraWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, Extra, OperandA(Literal13)),
	})

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 13, dcpu.Extra)
}

func TestExecuteInstructions_setLocationWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, Location, OperandA(Literal13)), 0x1000,
	})

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 13, dcpu.Memory[0x1000])
}

func TestExecuteInstructions_setLiteralWithSmallLiteral(t *testing.T) {
	dcpu := DCPU{}

	dcpu.Load([]uint16{
		Basic(Set, Literal, OperandA(Literal13)), 0x1000,
	})

	dcpu.ExecuteInstructions(1)
	assert.EqualValues(t, 0, dcpu.Memory[0x1000])
}
