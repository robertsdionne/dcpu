// Package dcpu implements an emulator for Notch's DCPU 1.7 specification.
package dcpu

import "fmt"

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

// Hardware represents peripheral devices that may interact with a DCPU.
type Hardware interface {
	Execute()
	GetID() uint32
	GetManufacturerID() uint32
	GetVersion() uint16
	HandleHardwareInterrupt()
}

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

// Load copies a sequence of instructions (a program) into memory.
func (d *DCPU) Load(instructions []uint16) {
	copy(d.Memory[:], instructions)
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
	return fmt.Sprintln("A:", d.RegisterA, "B:", d.RegisterB, "J:", d.RegisterJ, "PC:", d.ProgramCounter, "SP:", d.StackPointer,
		"[0:10]", d.Memory[0:10], "[0x100a]", d.Memory[0x100a], "[0x200a]", d.Memory[0x200a])
}

// ExecuteInstruction executes a single instruction.
func (d *DCPU) ExecuteInstruction(skip bool) {
	if !skip && !d.QueueInterrupts && len(d.InterruptQueue) > 0 {
		d.maybeTriggerInterrupt(d.InterruptQueue[0])
		d.InterruptQueue = d.InterruptQueue[1:]
	}

	stackPointerBackup := d.StackPointer
	instruction := d.Memory[d.ProgramCounter]
	d.ProgramCounter++
	basicOpcode := BasicOpcode(instruction & BasicOpcodeMask)

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
	} else {
		specialOpcode := SpecialOpcode(instruction & SpecialOpcodeMask)
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
