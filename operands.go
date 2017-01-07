// Package dcpu implements an emulator for Notch's DCPU 1.7 specification.
package dcpu

type (
	OperandA OperandB
	OperandB uint16
)

// Values for operands.
const (
	// RegisterA signifies register A.
	RegisterA OperandB = iota
	// RegisterB signifies register B.
	RegisterB OperandB = iota
	// RegisterC signifies register C.
	RegisterC OperandB = iota
	// RegisterX signifies register X.
	RegisterX OperandB = iota
	// RegisterY signifies register Y.
	RegisterY OperandB = iota
	// RegisterZ signifies register Z.
	RegisterZ OperandB = iota
	// RegisterI signifies register I.
	RegisterI OperandB = iota
	// RegisterJ signifies register J.
	RegisterJ OperandB = iota
	// LocationInRegisterA signifies location [A].
	LocationInRegisterA OperandB = iota
	// LocationInRegisterB signifies location [B].
	LocationInRegisterB OperandB = iota
	// LocationInRegisterC signifies location [C].
	LocationInRegisterC OperandB = iota
	// LocationInRegisterX signifies location [X].
	LocationInRegisterX OperandB = iota
	// LocationInRegisterY signifies location [Y].
	LocationInRegisterY OperandB = iota
	// LocationInRegisterZ signifies location [Z].
	LocationInRegisterZ OperandB = iota
	// LocationInRegisterI signifies location [I].
	LocationInRegisterI OperandB = iota
	// LocationInRegisterI signifies location [J].
	LocationInRegisterJ OperandB = iota
	// LocationOffsetByRegisterA signifies location [A + next word].
	LocationOffsetByRegisterA OperandB = iota
	// LocationOffsetByRegisterB signifies location [B + next word].
	LocationOffsetByRegisterB OperandB = iota
	// LocationOffsetByRegisterC signifies location [C + next word].
	LocationOffsetByRegisterC OperandB = iota
	// LocationOffsetByRegisterX signifies location [X + next word].
	LocationOffsetByRegisterX OperandB = iota
	// LocationOffsetByRegisterY signifies location [Y + next word].
	LocationOffsetByRegisterY OperandB = iota
	// LocationOffsetByRegisterZ signifies location [Z + next word].
	LocationOffsetByRegisterZ OperandB = iota
	// LocationOffsetByRegisterI signifies location [I + next word].
	LocationOffsetByRegisterI OperandB = iota
	// LocationOffsetByRegisterJ signifies location [J + next word].
	LocationOffsetByRegisterJ OperandB = iota
	// Push signifies [--SP] for operand b.
	// Pop signifies [SP++] for operand a.
	// PushOrPop signifies either Push or Pop depending upon context.
	Push, Pop, PushOrPop OperandB = iota, iota, iota
	// Peek signifies [SP].
	Peek OperandB = iota
	// Pick signifies [SP + next word].
	Pick OperandB = iota
	// StackPointer signifies SP.
	StackPointer OperandB = iota
	// ProgramCounter signifies PC.
	ProgramCounter OperandB = iota
	// Extra signifies EX.
	Extra OperandB = iota
	// Location signifies [next word].
	Location OperandB = iota
	// Literal signifies the next word, literally.
	Literal OperandB = iota
	// LiteralNegative1 signifies 0xffff.
	LiteralNegative1 OperandA = iota
	// Literal0 signifies 0x0.
	Literal0 OperandA = iota
	// Literal1 signifies 0x1.
	Literal1 OperandA = iota
	// Literal2 signifies 0x2.
	Literal2 OperandA = iota
	// Literal3 signifies 0x3.
	Literal3 OperandA = iota
	// Literal4 signifies 0x4.
	Literal4 OperandA = iota
	// Literal5 signifies 0x5.
	Literal5 OperandA = iota
	// Literal6 signifies 0x6.
	Literal6 OperandA = iota
	// Literal7 signifies 0x7.
	Literal7 OperandA = iota
	// Literal8 signifies 0x8.
	Literal8 OperandA = iota
	// Literal9 signifies 0x9.
	Literal9 OperandA = iota
	// Literal10 signifies 0xa.
	Literal10 OperandA = iota
	// Literal11 signifies 0xb.
	Literal11 OperandA = iota
	// Literal12 signifies 0xc.
	Literal12 OperandA = iota
	// Literal12 signifies 0xd.
	Literal13 OperandA = iota
	// Literal14 signifies 0xe.
	Literal14 OperandA = iota
	// Literal15 signifies 0xf.
	Literal15 OperandA = iota
	// Literal16 signifies 0x10.
	Literal16 OperandA = iota
	// Literal17 signifies 0x11.
	Literal17 OperandA = iota
	// Literal18 signifies 0x12.
	Literal18 OperandA = iota
	// Literal19 signifies 0x13.
	Literal19 OperandA = iota
	// Literal20 signifies 0x14.
	Literal20 OperandA = iota
	// Literal21 signifies 0x15.
	Literal21 OperandA = iota
	// Literal22 signifies 0x16.
	Literal22 OperandA = iota
	// Literal23 signifies 0x17.
	Literal23 OperandA = iota
	// Literal24 signifies 0x18.
	Literal24 OperandA = iota
	// Literal25 signifies 0x19.
	Literal25 OperandA = iota
	// Literal26 signifies 0x1a.
	Literal26 OperandA = iota
	// Literal27 signifies 0x1b.
	Literal27 OperandA = iota
	// Literal28 signifies 0x1c.
	Literal28 OperandA = iota
	// Literal29 signifies 0x1d.
	Literal29 OperandA = iota
	// Literal30 signifies 0x1e.
	Literal30 OperandA = iota
)
