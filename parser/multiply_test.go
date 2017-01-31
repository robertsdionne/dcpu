package parser

import (
	"testing"

	. "github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestAssemble_multiplyRegisterWithSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal16),
		Basic(Multiply, RegisterA, Literal30),
	}, Assemble(`
		set a, 16
		mul a, 30
	`))
}

func TestAssemble_multiplyRegisterWithOverflow(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, LiteralNegative1),
		Basic(Multiply, RegisterA, LiteralNegative1),
	}, Assemble(`
		set a, -1
		mul a, -1
	`))
}
