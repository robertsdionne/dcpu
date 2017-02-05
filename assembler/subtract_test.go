package assembler

import (
	"testing"

	. "github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestAssemble_subtractRegisterWithSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal30),
		Basic(Subtract, RegisterA, Literal16),
	}, Assemble(`
		set a, 30
		sub a, 16
	`))
}

func TestAssemble_subtractRegisterWithUnderflow(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal16),
		Basic(Subtract, RegisterA, Literal30),
	}, Assemble(`
		set a, 16
		sub a, 30
	`))
}
