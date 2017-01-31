package parser

import (
	"testing"

	. "github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestAssemble_shiftRightRegisterWithSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal), 0xFFF0,
		Basic(ShiftRight, RegisterA, Literal2),
	}, Assemble(`
		set a, 0xfff0
		shr a, 2
	`))
}

func TestAssemble_shiftRightRegisterWithUnderflow(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, LiteralNegative1),
		Basic(ShiftRight, RegisterA, Literal2),
	}, Assemble(`
		set a, -1
		shr a, 2
	`))
}
