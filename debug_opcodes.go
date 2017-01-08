package dcpu

type DebugOpcode uint16

// Debug opcodes.
const (
	// Noop performs no action.
	Noop = iota
	// Alert prints "alert" or the utf-16 string at 0xf000.
	Alert
	// DumpState prints a representation of the DCPU state.
	DumpState
)
