package parser

import (
	"testing"

	. "github.com/robertsdionne/dcpu"
	"github.com/stretchr/testify/assert"
)

func TestAssemble_moduloRegisterWithSmallLiteral(t *testing.T) {
	assert.Equal(t, []uint16{
		Basic(Set, RegisterA, Literal30),
		Basic(Modulo, RegisterA, Literal11),
	}, Assemble(`
		set a, 30
		mod a, 11
	`))
}
