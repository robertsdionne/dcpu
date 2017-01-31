package parser

import (
	"testing"

	. "github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestAssemble_divideSignedRegisterWithSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal16),
		Basic(DivideSigned, RegisterA, Literal), 0xfffe,
	}, Assemble(`
		set a, 16
		dvi a, 0xfffe
	`))
}

func TestAssemble_divideSignedByZero(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal30),
		Basic(DivideSigned, RegisterA, Literal0),
	}, Assemble(`
		set a, 30
		dvi a, 0
	`))
}
