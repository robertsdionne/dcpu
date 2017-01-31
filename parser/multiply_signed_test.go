package parser

import (
	"testing"

	. "github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestAssemble_multiplySignedRegisterWithSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, LiteralNegative1),
		Basic(MultiplySigned, RegisterA, Literal16),
	}, Assemble(`
		set a, -1
		mli a, 16
	`))
}

func TestAssemble_multiplySignedRegisterWithOverflow(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, LiteralNegative1),
		Basic(MultiplySigned, RegisterA, LiteralNegative1),
	}, Assemble(`
		set a, -1
		mli a, -1
	`))
}
