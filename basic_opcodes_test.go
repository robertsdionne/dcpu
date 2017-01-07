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
