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
}
