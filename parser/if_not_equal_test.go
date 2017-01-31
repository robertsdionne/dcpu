package parser

import (
	"testing"

	. "github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestAssemble_ifNotEqualRegisterWithUnequalSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal15),
		Basic(IfNotEqual, RegisterA, Literal0),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	}, Assemble(`
		set a, 15
		ifn a, 0
		set push, 13
		set push, 14
	`))
}

func TestAssemble_ifNotEqualRegisterWithEqualSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal15),
		Basic(IfNotEqual, RegisterA, Literal15),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	}, Assemble(`
		set a, 15
		ifn a, 15
		set push, 13
		set push, 14
	`))
}
