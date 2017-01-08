package dcpu

type SpecialOpcode uint16

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
