package parser

import (
	"testing"

	. "github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestAssemble_divideRegisterWithSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal30),
		Basic(Divide, RegisterA, Literal16),
	}, Assemble(`
		set a, 30
		div a, 16
	`))
}

func TestAssemble_divideByZero(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal30),
		Basic(Divide, RegisterA, Literal0),
	}, Assemble(`
		set a, 30
		div a, 0
	`))
}
