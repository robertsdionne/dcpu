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
