package assembler

import (
	"testing"

	. "github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestAssemble_binaryAndRegisterWithSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal), 0xf0f0,
		Basic(BinaryAnd, RegisterA, Literal), 0x00ff,
	}, Assemble(`
    set a, 0xf0f0
    and a, 0x00ff
    `))

	assert.Equal(t, []uint16{
		Basic(Set, Literal, RegisterA), 0xf0f0,
		Basic(BinaryAnd, Literal, RegisterA), 0x00ff,
		Basic(BinaryAnd, Literal, Literal), 0x00ff, 0xf0f0,
	}, Assemble(`
	set 0xf0f0, a
	and 0x00ff, a
	and 0xf0f0, 0x00ff
	`))
}
