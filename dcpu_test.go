package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var basicOpcodeCases = []struct{ expected, actual Word }{
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

var specialOpcodeCases = []struct{ expected, actual Word }{
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

var operandValueCases = []struct{ expected, actual Word }{
	{0x00, RegisterA},
	{0x01, RegisterB},
	{0x02, RegisterC},
	{0x03, RegisterX},
	{0x04, RegisterY},
	{0x05, RegisterZ},
	{0x06, RegisterI},
	{0x07, RegisterJ},
	{0x08, LocationInRegisterA},
	{0x09, LocationInRegisterB},
	{0x0a, LocationInRegisterC},
	{0x0b, LocationInRegisterX},
	{0x0c, LocationInRegisterY},
	{0x0d, LocationInRegisterZ},
	{0x0e, LocationInRegisterI},
	{0x0f, LocationInRegisterJ},
	{0x10, LocationOffsetByRegisterA},
	{0x11, LocationOffsetByRegisterB},
	{0x12, LocationOffsetByRegisterC},
	{0x13, LocationOffsetByRegisterX},
	{0x14, LocationOffsetByRegisterY},
	{0x15, LocationOffsetByRegisterZ},
	{0x16, LocationOffsetByRegisterI},
	{0x17, LocationOffsetByRegisterJ},
	{0x18, Push},
	{0x18, Pop},
	{0x18, PushOrPop},
	{0x19, Peek},
	{0x1a, Pick},
	{0x1b, StackPointer},
	{0x1c, ProgramCounter},
	{0x1d, Extra},
	{0x1e, Location},
	{0x1f, Literal},
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
