// Package dcpu implements an emulator for Notch's DCPU 1.7 specification.
package dcpu

type SpecialOpcode uint16

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
