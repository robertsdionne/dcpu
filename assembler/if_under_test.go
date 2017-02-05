package assembler

import (
	"testing"

	. "github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestAssemble_ifUnderRegisterWithLesserSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal30),
		Basic(IfUnder, RegisterA, LiteralNegative1),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	}, Assemble(`
		set a, 30
		ifu a, -1
		set push, 13
		set push, 14
	`))
}

func TestAssemble_ifUnderRegisterWithGreaterSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, LiteralNegative1),
		Basic(IfUnder, RegisterA, Literal30),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	}, Assemble(`
		set a, -1
		ifu a, 30
		set push, 13
		set push, 14
	`))
}
