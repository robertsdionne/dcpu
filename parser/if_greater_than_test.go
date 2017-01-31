package parser

import (
	"testing"

	. "github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestAssemble_ifGreaterThanRegisterWithLesserSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal30),
		Basic(IfGreaterThan, RegisterA, Literal15),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	}, Assemble(`
		set a, 30
		ifg a, 15
		set push, 13
		set push, 14
	`))
}

func TestAssemble_ifGreaterThanRegisterWithGreaterSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal15),
		Basic(IfGreaterThan, RegisterA, Literal30),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	}, Assemble(`
		set a, 15
		ifg a, 30
		set push, 13
		set push, 14
	`))
}
