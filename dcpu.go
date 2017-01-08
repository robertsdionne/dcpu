// Package dcpu implements an emulator for Notch's DCPU 1.7 specification.
package dcpu

import (
	"fmt"
	"log"
	"time"
	"unicode/utf16"
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
	// DebugOpcodeMask is used to extract the debug opcode.
	DebugOpcodeMask = BasicValueMaskA
	// DebugOpcodeShift is used to extract the shifted debug opcode.
	DebugOpcodeShift = BasicValueShiftA
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
	InterruptQueue   []uint16
	QueueInterrupts  bool
	InstructionCount uint16
	Memory           [MemorySize]uint16
	Hardware         []Hardware
}

// Basic builds a basic instruction with operands b and a.
func Basic(opcode BasicOpcode, b OperandB, a OperandA) (instruction uint16) {
	instruction = uint16(opcode) | uint16(b<<BasicValueShiftB) | uint16(a<<BasicValueShiftA)
	return
}

// Special builds a special instruction with operand a.
func Special(opcode SpecialOpcode, a OperandA) (instruction uint16) {
	instruction = uint16(opcode<<SpecialOpcodeShift) | uint16(a<<SpecialValueShiftA)
	return
}

// Debug builds a debug instruction.
func Debug(opcode DebugOpcode) (instruction uint16) {
	instruction = uint16(opcode << DebugOpcodeShift)
	return
}

// Load copies a sequence of instructions (a program) into memory.
func (d *DCPU) Load(start uint16, instructions []uint16) {
	copy(d.Memory[start:], instructions)
}

// LoadString copies a string into memory as utf16.
func (d *DCPU) LoadString(start uint16, message string) {
	d.Load(start, []uint16{uint16(len(message))})
	d.Load(start+1, utf16.Encode([]rune(message)))
}

// Execute runs the DCPU machine and all connected hardware.
func (d *DCPU) Execute() {
	for {
		d.ExecuteInstruction(false)
		for i := range d.Hardware {
			d.Hardware[i].Execute()
		}
		time.Sleep(10 * time.Microsecond)
	}
}

// ExecuteInstructions executes multiple instructions.
func (d *DCPU) ExecuteInstructions(count int) {
	for i := 0; i < count; i++ {
		// log.Println(*d)
		d.ExecuteInstruction(false)
	}
	// log.Println(*d)
}

func (d DCPU) String() string {
	state := []interface{}{}
	state = append(state,
		d.RegisterA, d.RegisterB, d.RegisterC,
		d.RegisterX, d.RegisterY, d.RegisterZ,
		d.RegisterI, d.RegisterJ,
		d.ProgramCounter, d.StackPointer, d.Extra,
		d.InterruptAddress, len(d.InterruptQueue), d.QueueInterrupts,
		len(d.Hardware), d.InstructionCount,
	)
	for _, value := range d.Memory[0x0000:0x0020] {
		state = append(state, value)
	}
	for _, value := range d.Memory[0x1000:0x1020] {
		state = append(state, value)
	}
	return fmt.Sprintf(`A: %#04x B: %#04x C: %#04x X: %#04x Y: %#04x Z: %#04x I: %#04x J: %#04x
                    PC %#04x SP %#04x EX %#04x IA %#04x IQ %#04x Q: % 6v HW %#04x IC %#04x

     [0x0000:0x0008]   %#04x    %#04x    %#04x    %#04x    %#04x    %#04x    %#04x    %#04x
     [0x0008:0x0010]   %#04x    %#04x    %#04x    %#04x    %#04x    %#04x    %#04x    %#04x
     [0x0010:0x0018]   %#04x    %#04x    %#04x    %#04x    %#04x    %#04x    %#04x    %#04x
     [0x0018:0x0020]   %#04x    %#04x    %#04x    %#04x    %#04x    %#04x    %#04x    %#04x

     [0x1000:0x1008]   %#04x    %#04x    %#04x    %#04x    %#04x    %#04x    %#04x    %#04x
     [0x1008:0x1010]   %#04x    %#04x    %#04x    %#04x    %#04x    %#04x    %#04x    %#04x
     [0x1010:0x1018]   %#04x    %#04x    %#04x    %#04x    %#04x    %#04x    %#04x    %#04x
     [0x1018:0x1020]   %#04x    %#04x    %#04x    %#04x    %#04x    %#04x    %#04x    %#04x
`, state...)
}

// ExecuteInstruction executes a single instruction.
func (d *DCPU) ExecuteInstruction(skip bool) {
	if !skip && !d.QueueInterrupts && len(d.InterruptQueue) > 0 {
		d.maybeTriggerInterrupt(d.InterruptQueue[0])
		d.InterruptQueue = d.InterruptQueue[1:]
	}

	if !skip {
		d.InstructionCount++
	}

	stackPointerBackup := d.StackPointer
	instruction := d.Memory[d.ProgramCounter]
	d.ProgramCounter++
	basicOpcode := BasicOpcode(instruction & BasicOpcodeMask)
	specialOpcode := SpecialOpcode((instruction & SpecialOpcodeMask) >> SpecialOpcodeShift)

	if basicOpcode != BasicReserved {
		operandA := OperandA((instruction & BasicValueMaskA) >> BasicValueShiftA)
		operandB := OperandB((instruction & BasicValueMaskB) >> BasicValueShiftB)

		pa, a := d.getOperandAddressOrLiteral(operandA, false)
		if pa != nil {
			a = *pa
		}

		pb, b := d.getOperandAddressOrLiteral(OperandA(operandB), true)
		if pb != nil {
			b = *pb
		}

		// log.Println("instruction", instruction)
		// log.Println("basicOpcode", basicOpcode)
		//
		// log.Println("operandB", operandB)
		// log.Println("pb", int(uintptr(unsafe.Pointer(pb)))-int(uintptr(unsafe.Pointer(&d.Memory))))
		// log.Println("b", b)
		//
		// log.Println("operandA", operandA)
		// log.Println("pa", int(uintptr(unsafe.Pointer(pa)))-int(uintptr(unsafe.Pointer(&d.Memory))))
		// log.Println("a", a)
		//
		// log.Println()

		if skip {
			d.StackPointer = stackPointerBackup

			switch basicOpcode {
			case IfBitSet, IfClear, IfEqual, IfNotEqual, IfGreaterThan, IfAbove, IfLessThan, IfUnder:
				d.ExecuteInstruction( /* skip */ true)
			}

			return
		}

		switch basicOpcode {
		case BasicReserved:
			return

		case Set:
			d.set(pb, a)

		case Add:
			d.add(pb, b, a)

		case Subtract:
			d.subtract(pb, b, a)

		case Multiply:
			d.multiply(pb, b, a)

		case MultiplySigned:
			d.multiplySigned(pb, b, a)

		case Divide:
			d.divide(pb, b, a)

		case DivideSigned:
			d.divideSigned(pb, b, a)

		case Modulo:
			d.set(pb, b%a)

		// TODO(robertsdionne): case ModuloSigned:

		case BinaryAnd:
			d.set(pb, b&a)

		case BinaryOr:
			d.set(pb, b|a)

		case BinaryExclusiveOr:
			d.set(pb, b^a)

		case ShiftRight:
			d.shiftRight(pb, b, a)

		// TODO(robertsdionne): case ArithmeticShiftRight:

		case ShiftLeft:
			d.shiftLeft(pb, b, a)

		case IfBitSet:
			d.skipInstructionIf((a & b) == 0)

		case IfClear:
			d.skipInstructionIf((a & b) != 0)

		case IfEqual:
			d.skipInstructionIf(b != a)

		case IfNotEqual:
			d.skipInstructionIf(b == a)

		case IfGreaterThan:
			d.skipInstructionIf(b <= a)

		case IfAbove:
			d.skipInstructionIf(int16(b) <= int16(a))

		case IfLessThan:
			d.skipInstructionIf(b >= a)

		case IfUnder:
			d.skipInstructionIf(int16(b) >= int16(a))

		// TODO(robertsdionne): case AddWithCarry:

		// TODO(robertsdionne): case SubtractWithCarry:

		case SetThenIncrement:
			d.setThenIncrement(pb, a)

		case SetThenDecrement:
			d.setThenDecrement(pb, b)
		}
	} else if specialOpcode != SpecialReserved {
		operandA := OperandA((instruction & SpecialValueMaskA) >> SpecialValueShiftA)
		assignable := specialOpcode == InterruptAddressGet || specialOpcode == HardwareNumberConnected

		pa, a := d.getOperandAddressOrLiteral(operandA, assignable)
		if pa != nil {
			a = *pa
		}

		if skip {
			d.StackPointer = stackPointerBackup
			return
		}

		switch specialOpcode {
		case SpecialReserved:
			return

		case JumpSubRoutine:
			d.jumpSubRoutine(a)

		case InterruptTrigger:
			d.Interrupt(a)

		case InterruptAddressGet:
			d.set(pa, d.InterruptAddress)

		case InterruptAddressSet:
			d.InterruptAddress = a

		case ReturnFromInterrupt:
			d.returnFromInterrupt()

		case InterruptAddToQueue:
			d.QueueInterrupts = a > 0

		case HardwareNumberConnected:
			d.set(pa, uint16(len(d.Hardware)))

		case HardwareQuery:
			d.hardwareQuery(a)

		case HardwareInterrupt:
			if int(a) < len(d.Hardware) {
				d.Hardware[a].HandleHardwareInterrupt()
			}
		}
	} else {
		debugOpcode := DebugOpcode((instruction & DebugOpcodeMask) >> DebugOpcodeShift)

		switch debugOpcode {
		case Alert:
			length := d.Memory[0xf000]
			switch {
			case length > 0:
				alert := string(utf16.Decode(d.Memory[0xf001 : 0xf001+length]))
				log.Println(alert)

			default:
				log.Println("alert")
			}

		case DumpState:
			log.Println(*d)
		}
	}
}

func (d *DCPU) skipInstructionIf(condition bool) {
	if condition {
		d.ExecuteInstruction( /* skip */ true)
	}
}

func (d *DCPU) Interrupt(message uint16) {
	switch {
	case d.QueueInterrupts:
		d.InterruptQueue = append(d.InterruptQueue, message)

	default:
		d.maybeTriggerInterrupt(message)
	}
}

func (d *DCPU) maybeTriggerInterrupt(message uint16) {
	if d.InterruptAddress == 0 {
		return
	}

	d.QueueInterrupts = true
	d.StackPointer--
	d.Memory[d.StackPointer] = d.ProgramCounter
	d.StackPointer--
	d.Memory[d.StackPointer] = d.RegisterA
	d.ProgramCounter = d.InterruptAddress
	d.RegisterA = message
}
