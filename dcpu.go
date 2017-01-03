// Package dcpu implements an emulator for Notch's DCPU 1.7 specification.
package dcpu

type (
	// Word values are sixteen bits wide.
	Word uint16
	// SignedWord values are sixteen bits wide.
	SignedWord int16
)

// Basic opcodes.
const (
	// BasicReserved indicates to use a special instruction.
	BasicReserved = iota
	// Set sets b to a.
	Set
	// Add sets b to b + a, sets EX to 0x1 upon overflow or 0x0 otherwise.
	Add
	// Subtract sets b to b - a, sets EX to 0xffff upon underflow or 0x0 otherwise.
	Subtract
	// Multiply sets b to b * a, sets EX to ((b * a) >> 16) & 0xffff treating a and b as unsigned.
	Multiply
	// MultiplySigned is like Multiply, but treats a and b as signed.
	MultiplySigned
	// Divide sets b to b / a, sets EX to ((b << 16) / a) & 0xffff. If a is 0, sets b and EX to 0 instead.
	// Treats a and b as unsigned.
	Divide
	// DivideSigned is like Divide, but treats a and b as signed.
	DivideSigned
	// Modulo sets b to b % a. If a is 0, sets b to 0 instead.
	Modulo
	// ModuloSigned is like Modulo, but treats a and b as signed.
	ModuloSigned
	// BinaryAnd sets b to b & a.
	BinaryAnd
	// BinaryOr sets b to b | a.
	BinaryOr
	// BinaryExclusiveOr sets b to b ^ a.
	BinaryExclusiveOr
	// ShiftRight sets b to b >>> a, sets EX to ((b << 16) >> a) & 0xffff.
	ShiftRight
	// ArithmeticShiftRight sets b to b >> a, sets EX to ((b << 16) >>> a) & 0xffff. Treats b as signed.
	ArithmeticShiftRight
	// ShiftLeft sets b to b << a, sets EX to ((b << a) >> 16) & 0xffff.
	ShiftLeft
	// IfBitSet performs the next instruction only if b & a is not 0.
	IfBitSet
	// IfClear performs the next instruction only if b & a is 0.
	IfClear
	// IfEqual performs the next instruction only if b == a.
	IfEqual
	// IfNotEqual performs the next instruction only if b != a.
	IfNotEqual
	// IfGreaterThan performs the next instruction only if b > a.
	IfGreaterThan
	// IfAbove performs the next instruction only if b > a, signed.
	IfAbove
	// IfLessThan performs the next instruction only if b < a.
	IfLessThan
	// IfUnder performs the next instruction only if b < a, signed.
	IfUnder
	_
	_
	// AddWithCarry sets b to b + a + EX, sets EX to 0x1 upon overflow or 0x0 otherwise.
	AddWithCarry
	// SubtractWithCarry sets b to b - a + EX, sets EX to 0xffff upon underflow, 0x1 upon overflow or 0x0 otherwise.
	SubtractWithCarry
	_
	_
	// SetThenIncrement sets b to a, then increments I and J by one.
	SetThenIncrement
	// SetThenDecrement sets b to a, then decrements I and J by one.
	SetThenDecrement
)

// Special opcodes.
const (
	// SpecialReserved is reserved for future expansion.
	SpecialReserved = iota
	// JumpSubRoutine pushes the address of the next instruction onto the stack, then sets PC to a.
	JumpSubRoutine
	_
	_
	_
	_
	_
	_
	// InterruptTrigger triggers a software interrupt with message a.
	InterruptTrigger
	// InterruptAddressGet sets a to IA.
	InterruptAddressGet
	// InterruptAddressSet sets IA to a.
	InterruptAddressSet
	// ReturnFromInterrupt disables interrupt queueing, pops a from the stack, then pops PC from the stack.
	ReturnFromInterrupt
	// InterruptAddToQueue upon nonzero a causes interrupts to add themselves to the queue instead of triggering.
	// Upon zero a causes interrupts to trigger again.
	InterruptAddToQueue
	_
	_
	_
	// HardwareNumberConnected sets a to the number of connected hardware devices.
	HardwareNumberConnected
	// HardwareQuery sets A, B, C, X and Y registers to information about hardware a:
	// A + (B << 16) is a 32-bit hardware id
	// C is the hardware version
	// X + (Y << 16) is a 32-bit manufacturer id
	HardwareQuery
	// HardwareInterrupt sends an interrupt to hardware a.
	HardwareInterrupt
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
	// Literal12 signifies 0xd.
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
	RegisterA        Word
	RegisterB        Word
	RegisterC        Word
	RegisterX        Word
	RegisterY        Word
	RegisterZ        Word
	RegisterI        Word
	RegisterJ        Word
	ProgramCounter   Word
	StackPointer     Word
	Extra            Word
	InterruptAddress Word
	Memory           [MemorySize]Word
}
