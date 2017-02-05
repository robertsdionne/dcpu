package assembler

import (
	"testing"

	. "github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestAssemble_ifBitSetRegisterWithCommonBitsSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal30),
		Basic(IfBitSet, RegisterA, Literal16),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	}, Assemble(`
		set a, 30
		ifb a, 16
		set push, 13
		set push, 14
	`))
}

func TestAssemble_ifBitSetRegisterWithoutCommonBitsSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal15),
		Basic(IfBitSet, RegisterA, Literal16),
		Basic(Set, Push, Literal13),
		Basic(Set, Push, Literal14),
	}, Assemble(`
		set a, 15
		ifb a, 16
		set push, 13
		set push, 14
	`))
}
