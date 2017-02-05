package assembler

import (
	"testing"

	. "github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestAssemble_addRegisterWithSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal13),
		Basic(Add, RegisterA, Literal14),
	}, Assemble(`
    set a, 13
    add a, 14
	`))
}

func TestAssemble_addRegisterWithOverflow(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, LiteralNegative1),
		Basic(Add, RegisterA, LiteralNegative1),
	}, Assemble(`
    set a, -1
    add a, -1
	`))
}
