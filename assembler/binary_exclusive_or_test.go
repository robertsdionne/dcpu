package assembler

import (
	"testing"

	. "github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestAssemble_binaryExclusiveOrRegisterWithSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal), 0xf0f0,
		Basic(BinaryExclusiveOr, RegisterA, Literal), 0x00ff,
	}, Assemble(`
		set a, 0xf0f0
		xor a, 0x00ff
	`))
}
