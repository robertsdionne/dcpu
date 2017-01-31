package parser

import (
	"testing"

	. "github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestAssemble_ifAboveRegisterWithLesserSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal30),
		Basic(IfAbove, RegisterA, LiteralNegative1),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	}, Assemble(`
		set a, 30
		ifa a, -1
		set push, 13
		set push, 14
	`))
}

func TestAssemble_ifAboveRegisterWithGreaterSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, LiteralNegative1),
		Basic(IfAbove, RegisterA, Literal30),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	}, Assemble(`
		set a, -1
		ifa a, 30
		set push, 13
		set push, 14
	`))
}
