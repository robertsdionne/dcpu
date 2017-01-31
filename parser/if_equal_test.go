package parser

import (
	"testing"

	. "github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestAssemble_ifEqualRegisterWithEqualSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal15),
		Basic(IfEqual, RegisterA, Literal15),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	}, Assemble(`
		set a, 15
		ife a, 15
		set push, 13
		set push, 14
	`))
}

func TestAssemble_ifEqualRegisterWithUnequalSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal15),
		Basic(IfEqual, RegisterA, Literal0),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	}, Assemble(`
		set a, 15
		ife a, 0
		set push, 13
		set push, 14
	`))
}

func TestAssemble_ifEqualSkipsConditionalsWhenNotEqual(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal15),
		Basic(IfEqual, RegisterA, Literal0),
		Basic(IfEqual, RegisterB, Literal0),
		Basic(Set, Push, Literal12),
		Basic(Set, Push, Literal13),
	}, Assemble(`
		set a, 15
		ife a, 0
		ife b, 0
		set push, 12
		set push, 13
	`))
}

func TestAssemble_ifEqualDoesNotSkipConditionalsWhenEqual(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal15),
		Basic(IfEqual, RegisterA, Literal15),
		Basic(IfEqual, RegisterB, Literal0),
		Basic(Set, Push, Literal12),
		Basic(Set, Push, Literal13),
	}, Assemble(`
		set a, 15
		ife a, 15
		ife b, 0
		set push, 12
		set push, 13
	`))
}
