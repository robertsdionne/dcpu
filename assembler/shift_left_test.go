package assembler

import (
	"testing"

	. "github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestAssemble_shiftLeftRegisterWithSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal30),
		Basic(ShiftLeft, RegisterA, Literal2),
	}, Assemble(`
		set a, 30
		shl a, 2
	`))
}

func TestAssemble_shiftLeftRegisterWithOverflow(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, LiteralNegative1),
		Basic(ShiftLeft, RegisterA, Literal2),
	}, Assemble(`
		set a, -1
		shl a, 2
	`))
}
