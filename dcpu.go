// Package dcpu implements an emulator for Notch's DCPU 1.7 specification.
package dcpu

type (
	BasicOpcode   uint16
	SpecialOpcode uint16
	OperandAValue OperandBValue
	OperandBValue uint16
)

// Basic opcodes.
const (
	// BasicReserved indicates to use a special instruction.
	BasicReserved BasicOpcode = iota
	// Set sets b to a.
	Set BasicOpcode = iota
	// Add sets b to b + a, sets EX to 0x1 upon overflow or 0x0 otherwise.
	Add BasicOpcode = iota
	// Subtract sets b to b - a, sets EX to 0xffff upon underflow or 0x0 otherwise.
	Subtract BasicOpcode = iota
	// Multiply sets b to b * a, sets EX to ((b * a) >> 16) & 0xffff treating a and b as unsigned.
	Multiply BasicOpcode = iota
	// MultiplySigned is like Multiply, but treats a and b as signed.
	MultiplySigned BasicOpcode = iota
	// Divide sets b to b / a, sets EX to ((b << 16) / a) & 0xffff. If a is 0, sets b and EX to 0 instead.
	// Treats a and b as unsigned.
	Divide BasicOpcode = iota
	// DivideSigned is like Divide, but treats a and b as signed.
	DivideSigned BasicOpcode = iota
	// Modulo sets b to b % a. If a is 0, sets b to 0 instead.
	Modulo BasicOpcode = iota
	// ModuloSigned is like Modulo, but treats a and b as signed.
	ModuloSigned BasicOpcode = iota
	// BinaryAnd sets b to b & a.
	BinaryAnd BasicOpcode = iota
	// BinaryOr sets b to b | a.
	BinaryOr BasicOpcode = iota
	// BinaryExclusiveOr sets b to b ^ a.
	BinaryExclusiveOr BasicOpcode = iota
	// ShiftRight sets b to b >>> a, sets EX to ((b << 16) >> a) & 0xffff.
	ShiftRight BasicOpcode = iota
	// ArithmeticShiftRight sets b to b >> a, sets EX to ((b << 16) >>> a) & 0xffff. Treats b as signed.
	ArithmeticShiftRight BasicOpcode = iota
	// ShiftLeft sets b to b << a, sets EX to ((b << a) >> 16) & 0xffff.
	ShiftLeft BasicOpcode = iota
	// IfBitSet performs the next instruction only if b & a is not 0.
	IfBitSet BasicOpcode = iota
	// IfClear performs the next instruction only if b & a is 0.
	IfClear BasicOpcode = iota
	// IfEqual performs the next instruction only if b == a.
	IfEqual BasicOpcode = iota
	// IfNotEqual performs the next instruction only if b != a.
	IfNotEqual BasicOpcode = iota
	// IfGreaterThan performs the next instruction only if b > a.
	IfGreaterThan BasicOpcode = iota
	// IfAbove performs the next instruction only if b > a, signed.
	IfAbove BasicOpcode = iota
	// IfLessThan performs the next instruction only if b < a.
	IfLessThan BasicOpcode = iota
	// IfUnder performs the next instruction only if b < a, signed.
	IfUnder BasicOpcode = iota
	_
	_
	// AddWithCarry sets b to b + a + EX, sets EX to 0x1 upon overflow or 0x0 otherwise.
	AddWithCarry BasicOpcode = iota
	// SubtractWithCarry sets b to b - a + EX, sets EX to 0xffff upon underflow, 0x1 upon overflow or 0x0 otherwise.
	SubtractWithCarry BasicOpcode = iota
	_
	_
	// SetThenIncrement sets b to a, then increments I and J by one.
	SetThenIncrement BasicOpcode = iota
	// SetThenDecrement sets b to a, then decrements I and J by one.
	SetThenDecrement BasicOpcode = iota
)

// Special opcodes.
const (
	// SpecialReserved is reserved for future expansion.
	SpecialReserved SpecialOpcode = iota
	// JumpSubRoutine pushes the address of the next instruction onto the stack, then sets PC to a.
	JumpSubRoutine SpecialOpcode = iota
	_
	_
	_
	_
	_
	_
	// InterruptTrigger triggers a software interrupt with message a.
	InterruptTrigger SpecialOpcode = iota
	// InterruptAddressGet sets a to IA.
	InterruptAddressGet SpecialOpcode = iota
	// InterruptAddressSet sets IA to a.
	InterruptAddressSet SpecialOpcode = iota
	// ReturnFromInterrupt disables interrupt queueing, pops a from the stack, then pops PC from the stack.
	ReturnFromInterrupt SpecialOpcode = iota
	// InterruptAddToQueue upon nonzero a causes interrupts to add themselves to the queue instead of triggering.
	// Upon zero a causes interrupts to trigger again.
	InterruptAddToQueue SpecialOpcode = iota
	_
	_
	_
	// HardwareNumberConnected sets a to the number of connected hardware devices.
	HardwareNumberConnected SpecialOpcode = iota
	// HardwareQuery sets A, B, C, X and Y registers to information about hardware a:
	// A + (B << 16) is a 32-bit hardware id
	// C is the hardware version
	// X + (Y << 16) is a 32-bit manufacturer id
	HardwareQuery SpecialOpcode = iota
	// HardwareInterrupt sends an interrupt to hardware a.
	HardwareInterrupt SpecialOpcode = iota
)

// Values for operands.
const (
	// RegisterA signifies register A.
	RegisterA OperandBValue = iota
	// RegisterB signifies register B.
	RegisterB OperandBValue = iota
	// RegisterC signifies register C.
	RegisterC OperandBValue = iota
	// RegisterX signifies register X.
	RegisterX OperandBValue = iota
	// RegisterY signifies register Y.
	RegisterY OperandBValue = iota
	// RegisterZ signifies register Z.
	RegisterZ OperandBValue = iota
	// RegisterI signifies register I.
	RegisterI OperandBValue = iota
	// RegisterJ signifies register J.
	RegisterJ OperandBValue = iota
	// LocationInRegisterA signifies location [A].
	LocationInRegisterA OperandBValue = iota
	// LocationInRegisterB signifies location [B].
	LocationInRegisterB OperandBValue = iota
	// LocationInRegisterC signifies location [C].
	LocationInRegisterC OperandBValue = iota
	// LocationInRegisterX signifies location [X].
	LocationInRegisterX OperandBValue = iota
	// LocationInRegisterY signifies location [Y].
	LocationInRegisterY OperandBValue = iota
	// LocationInRegisterZ signifies location [Z].
	LocationInRegisterZ OperandBValue = iota
	// LocationInRegisterI signifies location [I].
	LocationInRegisterI OperandBValue = iota
	// LocationInRegisterI signifies location [J].
	LocationInRegisterJ OperandBValue = iota
	// LocationOffsetByRegisterA signifies location [A + next word].
	LocationOffsetByRegisterA OperandBValue = iota
	// LocationOffsetByRegisterB signifies location [B + next word].
	LocationOffsetByRegisterB OperandBValue = iota
	// LocationOffsetByRegisterC signifies location [C + next word].
	LocationOffsetByRegisterC OperandBValue = iota
	// LocationOffsetByRegisterX signifies location [X + next word].
	LocationOffsetByRegisterX OperandBValue = iota
	// LocationOffsetByRegisterY signifies location [Y + next word].
	LocationOffsetByRegisterY OperandBValue = iota
	// LocationOffsetByRegisterZ signifies location [Z + next word].
	LocationOffsetByRegisterZ OperandBValue = iota
	// LocationOffsetByRegisterI signifies location [I + next word].
	LocationOffsetByRegisterI OperandBValue = iota
	// LocationOffsetByRegisterJ signifies location [J + next word].
	LocationOffsetByRegisterJ OperandBValue = iota
	// Push signifies [--SP] for operand b.
	// Pop signifies [SP++] for operand a.
	// PushOrPop signifies either Push or Pop depending upon context.
	Push, Pop, PushOrPop OperandBValue = iota, iota, iota
	// Peek signifies [SP].
	Peek OperandBValue = iota
	// Pick signifies [SP + next word].
	Pick OperandBValue = iota
	// StackPointer signifies SP.
	StackPointer OperandBValue = iota
	// ProgramCounter signifies PC.
	ProgramCounter OperandBValue = iota
	// Extra signifies EX.
	Extra OperandBValue = iota
	// Location signifies [next word].
	Location OperandBValue = iota
	// Literal signifies the next word, literally.
	Literal OperandBValue = iota
	// LiteralNegative1 signifies 0xffff.
	LiteralNegative1 OperandAValue = iota
	// Literal0 signifies 0x0.
	Literal0 OperandAValue = iota
	// Literal1 signifies 0x1.
	Literal1 OperandAValue = iota
	// Literal2 signifies 0x2.
	Literal2 OperandAValue = iota
	// Literal3 signifies 0x3.
	Literal3 OperandAValue = iota
	// Literal4 signifies 0x4.
	Literal4 OperandAValue = iota
	// Literal5 signifies 0x5.
	Literal5 OperandAValue = iota
	// Literal6 signifies 0x6.
	Literal6 OperandAValue = iota
	// Literal7 signifies 0x7.
	Literal7 OperandAValue = iota
	// Literal8 signifies 0x8.
	Literal8 OperandAValue = iota
	// Literal9 signifies 0x9.
	Literal9 OperandAValue = iota
	// Literal10 signifies 0xa.
	Literal10 OperandAValue = iota
	// Literal11 signifies 0xb.
	Literal11 OperandAValue = iota
	// Literal12 signifies 0xc.
	Literal12 OperandAValue = iota
	// Literal12 signifies 0xd.
	Literal13 OperandAValue = iota
	// Literal14 signifies 0xe.
	Literal14 OperandAValue = iota
	// Literal15 signifies 0xf.
	Literal15 OperandAValue = iota
	// Literal16 signifies 0x10.
	Literal16 OperandAValue = iota
	// Literal17 signifies 0x11.
	Literal17 OperandAValue = iota
	// Literal18 signifies 0x12.
	Literal18 OperandAValue = iota
	// Literal19 signifies 0x13.
	Literal19 OperandAValue = iota
	// Literal20 signifies 0x14.
	Literal20 OperandAValue = iota
	// Literal21 signifies 0x15.
	Literal21 OperandAValue = iota
	// Literal22 signifies 0x16.
	Literal22 OperandAValue = iota
	// Literal23 signifies 0x17.
	Literal23 OperandAValue = iota
	// Literal24 signifies 0x18.
	Literal24 OperandAValue = iota
	// Literal25 signifies 0x19.
	Literal25 OperandAValue = iota
	// Literal26 signifies 0x1a.
	Literal26 OperandAValue = iota
	// Literal27 signifies 0x1b.
	Literal27 OperandAValue = iota
	// Literal28 signifies 0x1c.
	Literal28 OperandAValue = iota
	// Literal29 signifies 0x1d.
	Literal29 OperandAValue = iota
	// Literal30 signifies 0x1e.
	Literal30 OperandAValue = iota
)

const (
	// MemorySize is the total size of the DCPU memory array.
	MemorySize = 0x10000
	// BasicOpcodeMask is used to extract the basic opcode.
	BasicOpcodeMask = 0x001f
	// BasicValueMaskA is used to extract the value of operand a.
	BasicValueMaskA = 0xfc00
	// BasicValueMaskB is used to extract the value of operand b.
	BasicValueMaskB = 0x03e0
	// BasicValueShiftA is used to extract the shifted value of operand a.
	BasicValueShiftA = 0xa
	// BasicValueShiftB is used to extract the shifted value of operand b.
	BasicValueShiftB = 0x5
	// SpecialOpcodeMask is used to extract the special opcode.
	SpecialOpcodeMask = BasicValueMaskB
	// SpecialOpcodeShift is used to extract the shifted special opcode.
	SpecialOpcodeShift = BasicValueShiftB
	// SpecialValueMaskA is used to extract the value of operand a.
	SpecialValueMaskA = BasicValueMaskA
	// SpecialValueShiftA is used to extract the shifted value of operand a.
	SpecialValueShiftA = BasicValueShiftA
)

// DCPU represents the state of a DCPU machine.
type DCPU struct {
	RegisterA        uint16
	RegisterB        uint16
	RegisterC        uint16
	RegisterX        uint16
	RegisterY        uint16
	RegisterZ        uint16
	RegisterI        uint16
	RegisterJ        uint16
	ProgramCounter   uint16
	StackPointer     uint16
	Extra            uint16
	InterruptAddress uint16
	Memory           [MemorySize]uint16
}

// ExecuteInstruction executes a single instruction.
func (d *DCPU) ExecuteInstruction(skip bool) {
	stackPointerBackup := d.StackPointer
	instruction := d.Memory[d.ProgramCounter]
	d.ProgramCounter++
	basicOpcode := BasicOpcode(instruction & BasicOpcodeMask)

	if basicOpcode != BasicReserved {
		operandA := OperandAValue((instruction & BasicValueMaskA) >> BasicValueShiftA)
		operandB := OperandBValue((instruction & BasicValueMaskB) >> BasicValueShiftB)

		operandAAddress, operandAValue := d.getOperandAddressOrLiteral(operandA, false)
		if operandAAddress != nil {
			operandAValue = uint32(*operandAAddress)
		}

		operandBAddress, operandBValue := d.getOperandAddressOrLiteral(OperandAValue(operandB), true)
		if operandBAddress != nil {
			operandBValue = uint32(*operandBAddress)
		}

		if skip {
			d.StackPointer = stackPointerBackup
			return
		}

		switch basicOpcode {
		case BasicReserved:
			return

		case Set:
			d.maybeAssignResult(operandBAddress, operandAValue)

		case Add:
			result := operandBValue + operandAValue
			d.Extra = uint16(result >> 16)
			d.maybeAssignResult(operandBAddress, result)

		case Subtract:
			d.Extra = 0
			if operandBValue < operandAValue {
				d.Extra = 1
			}
			d.maybeAssignResult(operandBAddress, operandBValue-operandAValue)

		case Multiply:
			result := operandBValue * operandAValue
			d.Extra = uint16(result >> 16)
			d.maybeAssignResult(operandBAddress, result)

		case MultiplySigned:
			result := int32(operandBValue) * int32(operandAValue)
			d.Extra = uint16(result >> 16)
			d.maybeAssignResult(operandBAddress, uint32(result))

			// TODO(robertsdionne): Finish basic opcode cases.
		}
	} else {
		// TODO(robertsdionne): Finish special opcode cases.
	}
}

func (d *DCPU) getOperandAddressOrLiteral(
	operandTypeA OperandAValue, assignable bool) (address *uint16, literal uint32) {

	operandTypeB := OperandBValue(operandTypeA)
	push := operandTypeB == PushOrPop && assignable
	pop := operandTypeB == PushOrPop && !assignable

	switch {
	case operandTypeB < LocationInRegisterA:
		address = d.registerAddress(operandTypeB)

	case operandTypeB < LocationOffsetByRegisterA:
		address = &d.Memory[d.registerValue(operandTypeB)]

	case operandTypeB < PushOrPop:
		address = &d.Memory[d.ProgramCounter+d.registerValue(operandTypeB)]
		d.ProgramCounter++

	case push:
		d.StackPointer--
		address = &d.Memory[d.StackPointer]

	case pop:
		address = &d.Memory[d.StackPointer]
		d.StackPointer++

	case operandTypeB == Peek:
		address = &d.Memory[d.StackPointer]

	case operandTypeB == Pick:
		address = &d.Memory[d.StackPointer+d.Memory[d.ProgramCounter]]
		d.ProgramCounter++

	case operandTypeB == StackPointer:
		address = &d.StackPointer

	case operandTypeB == ProgramCounter:
		address = &d.ProgramCounter

	case operandTypeB == Extra:
		address = &d.Extra

	case operandTypeB == Location:
		address = &d.Memory[d.Memory[d.ProgramCounter]]
		d.ProgramCounter++

	case operandTypeB == Literal:
		address = &d.Memory[d.ProgramCounter]
		d.ProgramCounter++

	default:
		literal = uint32(operandTypeA - Literal0)
	}
	return
}

func (d *DCPU) registerAddress(registerIndex OperandBValue) (address *uint16) {
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

func (d *DCPU) registerValue(registerIndex OperandBValue) (value uint16) {
	switch registerIndex {
	case RegisterA:
		value = d.RegisterA
	case RegisterB:
		value = d.RegisterB
	case RegisterC:
		value = d.RegisterC
	case RegisterX:
		value = d.RegisterX
	case RegisterY:
		value = d.RegisterY
	case RegisterZ:
		value = d.RegisterZ
	case RegisterI:
		value = d.RegisterI
	case RegisterJ:
		value = d.RegisterJ
	}
	return
}

func (d *DCPU) maybeAssignResult(address *uint16, value uint32) {
	if address != nil {
		*address = uint16(value)
	}
}
