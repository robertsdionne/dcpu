package dcpu

type BasicOpcode uint16

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
