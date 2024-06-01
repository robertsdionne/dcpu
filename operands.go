package dcpu

type (
	OperandA OperandB
	OperandB uint16
)

// Values for operands.
const (
	// RegisterA signifies register A.
	RegisterA = iota
	// RegisterB signifies register B.
	RegisterB
	// RegisterC signifies register C.
	RegisterC
	// RegisterX signifies register X.
	RegisterX
	// RegisterY signifies register Y.
	RegisterY
	// RegisterZ signifies register Z.
	RegisterZ
	// RegisterI signifies register I.
	RegisterI
	// RegisterJ signifies register J.
	RegisterJ
	// LocationInRegisterA signifies location [A].
	LocationInRegisterA
	// LocationInRegisterB signifies location [B].
	LocationInRegisterB
	// LocationInRegisterC signifies location [C].
	LocationInRegisterC
	// LocationInRegisterX signifies location [X].
	LocationInRegisterX
	// LocationInRegisterY signifies location [Y].
	LocationInRegisterY
	// LocationInRegisterZ signifies location [Z].
	LocationInRegisterZ
	// LocationInRegisterI signifies location [I].
	LocationInRegisterI
	// LocationInRegisterI signifies location [J].
	LocationInRegisterJ
	// LocationOffsetByRegisterA signifies location [A + next word].
	LocationOffsetByRegisterA
	// LocationOffsetByRegisterB signifies location [B + next word].
	LocationOffsetByRegisterB
	// LocationOffsetByRegisterC signifies location [C + next word].
	LocationOffsetByRegisterC
	// LocationOffsetByRegisterX signifies location [X + next word].
	LocationOffsetByRegisterX
	// LocationOffsetByRegisterY signifies location [Y + next word].
	LocationOffsetByRegisterY
	// LocationOffsetByRegisterZ signifies location [Z + next word].
	LocationOffsetByRegisterZ
	// LocationOffsetByRegisterI signifies location [I + next word].
	LocationOffsetByRegisterI
	// LocationOffsetByRegisterJ signifies location [J + next word].
	LocationOffsetByRegisterJ
	// Push signifies [--SP] for operand b.
	// Pop signifies [SP++] for operand a.
	// PushOrPop signifies either Push or Pop depending upon context.
	Push, Pop, PushOrPop = iota, iota, iota
	// Peek signifies [SP].
	Peek = iota
	// Pick signifies [SP + next word].
	Pick
	// StackPointer signifies SP.
	StackPointer
	// ProgramCounter signifies PC.
	ProgramCounter
	// Extra signifies EX.
	Extra
	// Location signifies [next word].
	Location
	// Literal signifies the next word, literally.
	Literal
	// LiteralNegative1 signifies 0xffff.
	LiteralNegative1
	// Literal0 signifies 0x0.
	Literal0
	// Literal1 signifies 0x1.
	Literal1
	// Literal2 signifies 0x2.
	Literal2
	// Literal3 signifies 0x3.
	Literal3
	// Literal4 signifies 0x4.
	Literal4
	// Literal5 signifies 0x5.
	Literal5
	// Literal6 signifies 0x6.
	Literal6
	// Literal7 signifies 0x7.
	Literal7
	// Literal8 signifies 0x8.
	Literal8
	// Literal9 signifies 0x9.
	Literal9
	// Literal10 signifies 0xa.
	Literal10
	// Literal11 signifies 0xb.
	Literal11
	// Literal12 signifies 0xc.
	Literal12
	// Literal13 signifies 0xd.
	Literal13
	// Literal14 signifies 0xe.
	Literal14
	// Literal15 signifies 0xf.
	Literal15
	// Literal16 signifies 0x10.
	Literal16
	// Literal17 signifies 0x11.
	Literal17
	// Literal18 signifies 0x12.
	Literal18
	// Literal19 signifies 0x13.
	Literal19
	// Literal20 signifies 0x14.
	Literal20
	// Literal21 signifies 0x15.
	Literal21
	// Literal22 signifies 0x16.
	Literal22
	// Literal23 signifies 0x17.
	Literal23
	// Literal24 signifies 0x18.
	Literal24
	// Literal25 signifies 0x19.
	Literal25
	// Literal26 signifies 0x1a.
	Literal26
	// Literal27 signifies 0x1b.
	Literal27
	// Literal28 signifies 0x1c.
	Literal28
	// Literal29 signifies 0x1d.
	Literal29
	// Literal30 signifies 0x1e.
	Literal30
)

func (d *DCPU) getOperandAddressOrLiteral(operand uint16, assignable bool) (address *uint16, literal uint16) {
	push := operand == PushOrPop && assignable
	pop := operand == PushOrPop && !assignable

	switch {
	case operand <= RegisterJ:
		address = d.registerAddress(operand)

	case operand <= LocationInRegisterJ:
		address = &d.Memory[d.registerValue(operand)]

	case operand <= LocationOffsetByRegisterJ:
		address = &d.Memory[d.Memory[d.ProgramCounter]+d.registerValue(operand)]
		d.ProgramCounter++

	case push:
		d.StackPointer--
		address = &d.Memory[d.StackPointer]

	case pop:
		address = &d.Memory[d.StackPointer]
		d.StackPointer++

	case operand == Peek:
		address = &d.Memory[d.StackPointer]

	case operand == Pick:
		address = &d.Memory[d.StackPointer+d.Memory[d.ProgramCounter]]
		d.ProgramCounter++

	case operand == StackPointer:
		address = &d.StackPointer

	case operand == ProgramCounter:
		address = &d.ProgramCounter

	case operand == Extra:
		address = &d.Extra

	case operand == Location:
		address = &d.Memory[d.Memory[d.ProgramCounter]]
		d.ProgramCounter++

	case operand == Literal:
		address = &d.Memory[d.ProgramCounter]
		d.ProgramCounter++

	default:
		literal = operand - Literal0
	}
	return
}

func (d *DCPU) registerAddress(registerIndex uint16) (address *uint16) {
	switch registerIndex {
	case RegisterA:
		address = &d.RegisterA
	case RegisterB:
		address = &d.RegisterB
	case RegisterC:
		address = &d.RegisterC
	case RegisterX:
		address = &d.RegisterX
	case RegisterY:
		address = &d.RegisterY
	case RegisterZ:
		address = &d.RegisterZ
	case RegisterI:
		address = &d.RegisterI
	case RegisterJ:
		address = &d.RegisterJ
	}
	return
}

func (d *DCPU) registerValue(registerIndex uint16) (value uint16) {
	switch registerIndex {
	case RegisterA, LocationInRegisterA, LocationOffsetByRegisterA:
		value = d.RegisterA
	case RegisterB, LocationInRegisterB, LocationOffsetByRegisterB:
		value = d.RegisterB
	case RegisterC, LocationInRegisterC, LocationOffsetByRegisterC:
		value = d.RegisterC
	case RegisterX, LocationInRegisterX, LocationOffsetByRegisterX:
		value = d.RegisterX
	case RegisterY, LocationInRegisterY, LocationOffsetByRegisterY:
		value = d.RegisterY
	case RegisterZ, LocationInRegisterZ, LocationOffsetByRegisterZ:
		value = d.RegisterZ
	case RegisterI, LocationInRegisterI, LocationOffsetByRegisterI:
		value = d.RegisterI
	case RegisterJ, LocationInRegisterJ, LocationOffsetByRegisterJ:
		value = d.RegisterJ
	}
	return
}
