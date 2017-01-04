package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var basicOpcodeCases = []struct{ expected, actual BasicOpcode }{
	{0x00, BasicReserved},
	{0x01, Set},
	{0x02, Add},
	{0x03, Subtract},
	{0x04, Multiply},
	{0x05, MultiplySigned},
	{0x06, Divide},
	{0x07, DivideSigned},
	{0x08, Modulo},
	{0x09, ModuloSigned},
	{0x0a, BinaryAnd},
	{0x0b, BinaryOr},
	{0x0c, BinaryExclusiveOr},
	{0x0d, ShiftRight},
	{0x0e, ArithmeticShiftRight},
	{0x0f, ShiftLeft},
	{0x10, IfBitSet},
	{0x11, IfClear},
	{0x12, IfEqual},
	{0x13, IfNotEqual},
	{0x14, IfGreaterThan},
	{0x15, IfAbove},
	{0x16, IfLessThan},
	{0x17, IfUnder},
	{0x1a, AddWithCarry},
	{0x1b, SubtractWithCarry},
	{0x1e, SetThenIncrement},
	{0x1f, SetThenDecrement},
}

func TestBasicOpcodes(t *testing.T) {
	for _, testCase := range basicOpcodeCases {
		assert.Equal(t, testCase.expected, testCase.actual)
	}
}

var specialOpcodeCases = []struct{ expected, actual SpecialOpcode }{
	{0x00, SpecialReserved},
	{0x01, JumpSubRoutine},
	{0x08, InterruptTrigger},
	{0x09, InterruptAddressGet},
	{0x0a, InterruptAddressSet},
	{0x0b, ReturnFromInterrupt},
	{0x0c, InterruptAddToQueue},
	{0x10, HardwareNumberConnected},
	{0x11, HardwareQuery},
	{0x12, HardwareInterrupt},
}

func TestSpecialOpcodes(t *testing.T) {
	for _, testCase := range specialOpcodeCases {
		assert.Equal(t, testCase.expected, testCase.actual)
	}
}

var operandValueCases = []struct{ expected, actual OperandA }{
	{0x00, OperandA(RegisterA)},
	{0x01, OperandA(RegisterB)},
	{0x02, OperandA(RegisterC)},
	{0x03, OperandA(RegisterX)},
	{0x04, OperandA(RegisterY)},
	{0x05, OperandA(RegisterZ)},
	{0x06, OperandA(RegisterI)},
	{0x07, OperandA(RegisterJ)},
	{0x08, OperandA(LocationInRegisterA)},
	{0x09, OperandA(LocationInRegisterB)},
	{0x0a, OperandA(LocationInRegisterC)},
	{0x0b, OperandA(LocationInRegisterX)},
	{0x0c, OperandA(LocationInRegisterY)},
	{0x0d, OperandA(LocationInRegisterZ)},
	{0x0e, OperandA(LocationInRegisterI)},
	{0x0f, OperandA(LocationInRegisterJ)},
	{0x10, OperandA(LocationOffsetByRegisterA)},
	{0x11, OperandA(LocationOffsetByRegisterB)},
	{0x12, OperandA(LocationOffsetByRegisterC)},
	{0x13, OperandA(LocationOffsetByRegisterX)},
	{0x14, OperandA(LocationOffsetByRegisterY)},
	{0x15, OperandA(LocationOffsetByRegisterZ)},
	{0x16, OperandA(LocationOffsetByRegisterI)},
	{0x17, OperandA(LocationOffsetByRegisterJ)},
	{0x18, OperandA(Push)},
	{0x18, OperandA(Pop)},
	{0x18, OperandA(PushOrPop)},
	{0x19, OperandA(Peek)},
	{0x1a, OperandA(Pick)},
	{0x1b, OperandA(StackPointer)},
	{0x1c, OperandA(ProgramCounter)},
	{0x1d, OperandA(Extra)},
	{0x1e, OperandA(Location)},
	{0x1f, OperandA(Literal)},
	{0x20, LiteralNegative1},
	{0x21, Literal0},
	{0x22, Literal1},
	{0x23, Literal2},
	{0x24, Literal3},
	{0x25, Literal4},
	{0x26, Literal5},
	{0x27, Literal6},
	{0x28, Literal7},
	{0x29, Literal8},
	{0x2a, Literal9},
	{0x2b, Literal10},
	{0x2c, Literal11},
	{0x2d, Literal12},
	{0x2e, Literal13},
	{0x2f, Literal14},
	{0x30, Literal15},
	{0x31, Literal16},
	{0x32, Literal17},
	{0x33, Literal18},
	{0x34, Literal19},
	{0x35, Literal20},
	{0x36, Literal21},
	{0x37, Literal22},
	{0x38, Literal23},
	{0x39, Literal24},
	{0x3a, Literal25},
	{0x3b, Literal26},
	{0x3c, Literal27},
	{0x3d, Literal28},
	{0x3e, Literal29},
	{0x3f, Literal30},
}

func TestOperandValues(t *testing.T) {
	for _, testCase := range operandValueCases {
		assert.Equal(t, testCase.expected, testCase.actual)
	}
}

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
