// Package dcpu implements an emulator for Notch's DCPU 1.7 specification.
package dcpu

type BasicOpcode uint16

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
