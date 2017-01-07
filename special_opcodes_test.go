package dcpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
