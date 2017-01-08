package dcpu

type DebugOpcode uint16

// Debug opcodes.
const (
	// Noop performs no action.
	Noop      DebugOpcode = iota
	Alert     DebugOpcode = iota
	DumpState DebugOpcode = iota
)
