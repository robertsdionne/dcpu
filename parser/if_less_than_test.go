package parser

import (
	"testing"

	. "github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestAssemble_ifLessThanRegisterWithLesserSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal30),
		Basic(IfLessThan, RegisterA, Literal15),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	}, Assemble(`
		set a, 30
		ifl a, 15
		set push, 13
		set push, 14
	`))
}

func TestAssemble_ifLessThanRegisterWithGreaterSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal15),
		Basic(IfLessThan, RegisterA, Literal30),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	}, Assemble(`
		set a, 15
		ifl a, 30
		set push, 13
		set push, 14
	`))
}
