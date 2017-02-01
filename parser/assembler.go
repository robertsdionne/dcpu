package parser

import (
	"log"
	"strconv"
	"strings"
	"unicode/utf16"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/robertsdionne/dcpu"
)

//go:generate java -jar $HOME/Downloads/antlr-4.6-complete.jar -Dlanguage=Go -package parser DCPU.g4 -visitor

func Assemble(source string) (program []uint16) {
	input := antlr.NewInputStream(source)
	lexer := NewDCPULexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewDCPUParser(stream)
	parser.BuildParseTrees = true
	visitor := Assembler{
		BaseDCPUVisitor: &BaseDCPUVisitor{},
	}
	program = visitor.Visit(parser.Program()).([]uint16)
	return
}

type Assembler struct {
	*BaseDCPUVisitor
}

func (a *Assembler) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(a)
}

func (a *Assembler) VisitChildren(node antlr.RuleNode) interface{} {
	var results []interface{}
	for _, child := range node.GetChildren() {
		results = append(results, child.(antlr.ParseTree).Accept(a))
	}
	return results
}

func (a *Assembler) VisitProgram(ctx *ProgramContext) interface{} {
	program := []interface{}{}
	labelAddresses := map[string]uint16{}

	results := a.VisitChildren(ctx)
	for _, result := range results.([]interface{}) {
		switch result := result.(type) {
		case []interface{}:
			for _, word := range result {
				program = append(program, word)
			}

		case string:
			labelAddresses[result] = uint16(len(program))
		}
	}

	assembledProgram := []uint16{}

	for _, word := range program {
		switch word := word.(type) {
		case uint16:
			assembledProgram = append(assembledProgram, word)
		case string:
			assembledProgram = append(assembledProgram, labelAddresses[word])
		default:
			log.Fatalln("Found bad type", word)
		}
	}

	return assembledProgram
}

func (a *Assembler) VisitInstruction(ctx *InstructionContext) interface{} {
	switch {
	case ctx.BinaryOperation() != nil:
		return a.Visit(ctx.BinaryOperation())

	case ctx.UnaryOperation() != nil:
		return a.Visit(ctx.UnaryOperation())
	}
	return nil
}

func (a *Assembler) VisitLabelDefinition(ctx *LabelDefinitionContext) interface{} {
	return ctx.IDENTIFIER().GetText()
}

func (a *Assembler) VisitDataSection(ctx *DataSectionContext) interface{} {
	return a.Visit(ctx.Data())
}

func (a *Assembler) VisitData(ctx *DataContext) interface{} {
	var result []interface{}
	for _, child := range ctx.GetChildren() {
		switch child := child.(type) {
		case *DatumContext:
			datum := a.Visit(child)
			switch datum := datum.(type) {
			case []uint16:
				for _, value := range datum {
					result = append(result, value)
				}

			case string, uint16:
				result = append(result, datum)
			}
		}
	}
	return result
}

func (a *Assembler) VisitDatum(ctx *DatumContext) interface{} {
	switch {
	case ctx.STRING() != nil:
		text := ctx.STRING().GetText()
		encoded := utf16.Encode([]rune(text[1 : len(text)-2]))
		result := []uint16{uint16(len(encoded))}
		result = append(result, encoded...)
		return result

	case ctx.IDENTIFIER() != nil:
		return ctx.IDENTIFIER().GetText()

	case ctx.NUMBER() != nil:
		return parseValue(ctx.NUMBER().GetText())
	}

	return nil
}

func (a *Assembler) VisitBinaryOperation(ctx *BinaryOperationContext) interface{} {
	opcode := a.Visit(ctx.BinaryOpcode()).(uint16)
	argumentB := a.Visit(ctx.ArgumentB()).([]interface{})
	argumentA := a.Visit(ctx.ArgumentA()).([]interface{})
	result := []interface{}{
		opcode | argumentB[0].(uint16)<<dcpu.BasicValueShiftB | argumentA[0].(uint16)<<dcpu.BasicValueShiftA,
	}
	result = append(result, argumentB[1:]...)
	result = append(result, argumentA[1:]...)
	return result
}

func (a *Assembler) VisitBinaryOpcode(ctx *BinaryOpcodeContext) interface{} {
	return binaryOpcodeValue(ctx)
}

func binaryOpcodeValue(ctx *BinaryOpcodeContext) uint16 {
	switch {
	case ctx.SET() != nil:
		return dcpu.Set

	case ctx.ADD() != nil:
		return dcpu.Add

	case ctx.SUB() != nil:
		return dcpu.Subtract

	case ctx.MUL() != nil:
		return dcpu.Multiply

	case ctx.MLI() != nil:
		return dcpu.MultiplySigned

	case ctx.DIV() != nil:
		return dcpu.Divide

	case ctx.DVI() != nil:
		return dcpu.DivideSigned

	case ctx.MOD() != nil:
		return dcpu.Modulo

	case ctx.MDI() != nil:
		return dcpu.ModuloSigned

	case ctx.AND() != nil:
		return dcpu.BinaryAnd

	case ctx.BOR() != nil:
		return dcpu.BinaryOr

	case ctx.XOR() != nil:
		return dcpu.BinaryExclusiveOr

	case ctx.SHR() != nil:
		return dcpu.ShiftRight

	case ctx.ASR() != nil:
		return dcpu.ArithmeticShiftRight

	case ctx.SHL() != nil:
		return dcpu.ShiftLeft

	case ctx.IFB() != nil:
		return dcpu.IfBitSet

	case ctx.IFC() != nil:
		return dcpu.IfClear

	case ctx.IFE() != nil:
		return dcpu.IfEqual

	case ctx.IFN() != nil:
		return dcpu.IfNotEqual

	case ctx.IFG() != nil:
		return dcpu.IfGreaterThan

	case ctx.IFA() != nil:
		return dcpu.IfAbove

	case ctx.IFL() != nil:
		return dcpu.IfLessThan

	case ctx.IFU() != nil:
		return dcpu.IfUnder

	case ctx.ADX() != nil:
		return dcpu.AddWithCarry

	case ctx.SBX() != nil:
		return dcpu.SubtractWithCarry

	case ctx.STI() != nil:
		return dcpu.SetThenIncrement

	case ctx.STD() != nil:
		return dcpu.SetThenDecrement

	default:
		return 0
	}
}

func (a *Assembler) VisitUnaryOperation(ctx *UnaryOperationContext) interface{} {
	opcode := a.Visit(ctx.UnaryOpcode()).(uint16)
	argument := a.Visit(ctx.ArgumentA()).([]interface{})
	result := []interface{}{opcode<<dcpu.SpecialOpcodeShift | argument[0].(uint16)<<dcpu.SpecialValueShiftA}
	result = append(result, argument[1:]...)
	return result
}

func (a *Assembler) VisitUnaryOpcode(ctx *UnaryOpcodeContext) interface{} {
	return unaryOpcodeValue(ctx)
}

func unaryOpcodeValue(ctx *UnaryOpcodeContext) uint16 {
	switch {
	case ctx.JSR() != nil:
		return dcpu.JumpSubRoutine

	case ctx.INT() != nil:
		return dcpu.InterruptTrigger

	case ctx.IAG() != nil:
		return dcpu.InterruptAddressGet

	case ctx.IAS() != nil:
		return dcpu.InterruptAddressSet

	case ctx.RFI() != nil:
		return dcpu.ReturnFromInterrupt

	case ctx.IAQ() != nil:
		return dcpu.InterruptAddToQueue

	case ctx.HWN() != nil:
		return dcpu.HardwareNumberConnected

	case ctx.HWQ() != nil:
		return dcpu.HardwareQuery

	case ctx.HWI() != nil:
		return dcpu.HardwareInterrupt

	default:
		return 0
	}
}

func (a *Assembler) VisitArgumentB(ctx *ArgumentBContext) interface{} {
	switch {
	case ctx.Register() != nil:
		return []interface{}{a.Visit(ctx.Register())}

	case ctx.LocationInRegister() != nil:
		return []interface{}{a.Visit(ctx.LocationInRegister())}

	case ctx.LocationOffsetByRegister() != nil:
		return a.Visit(ctx.LocationOffsetByRegister())

	case ctx.PUSH() != nil:
		return []interface{}{uint16(dcpu.Push)}

	case ctx.PEEK() != nil:
		return []interface{}{uint16(dcpu.Peek)}

	case ctx.Pick() != nil:
		return a.Visit(ctx.Pick())

	case ctx.STACK_POINTER() != nil:
		return []interface{}{uint16(dcpu.StackPointer)}

	case ctx.PROGRAM_COUNTER() != nil:
		return []interface{}{uint16(dcpu.ProgramCounter)}

	case ctx.EXTRA() != nil:
		return []interface{}{uint16(dcpu.Extra)}

	case ctx.Location() != nil:
		return []interface{}{uint16(dcpu.Location), a.Visit(ctx.Location())}
	}
	return []interface{}{uint16(0)}
}

func (a *Assembler) VisitRegister(ctx *RegisterContext) interface{} {
	return registerValue(ctx.REGISTER())
}

func (a *Assembler) VisitLocationInRegister(ctx *LocationInRegisterContext) interface{} {
	return registerValue(ctx.REGISTER()) + dcpu.LocationInRegisterA
}

func (a *Assembler) VisitLocationOffsetByRegister(ctx *LocationOffsetByRegisterContext) interface{} {
	var location interface{}
	switch {
	case ctx.Label() != nil:
		location = a.Visit(ctx.Label())
	case ctx.Value() != nil:
		location = a.Visit(ctx.Value())
	}
	return []interface{}{registerValue(ctx.REGISTER()) + dcpu.LocationOffsetByRegisterA, location}
}

func registerValue(register antlr.TerminalNode) uint16 {
	switch strings.ToUpper(register.GetText()) {
	case "A":
		return dcpu.RegisterA

	case "B":
		return dcpu.RegisterB

	case "C":
		return dcpu.RegisterC

	case "X":
		return dcpu.RegisterX

	case "Y":
		return dcpu.RegisterY

	case "Z":
		return dcpu.RegisterZ

	case "I":
		return dcpu.RegisterI

	case "J":
		return dcpu.RegisterJ
	}

	return 0
}

func (a *Assembler) VisitArgumentA(ctx *ArgumentAContext) interface{} {
	switch {
	case ctx.Register() != nil:
		return []interface{}{a.Visit(ctx.Register())}

	case ctx.LocationInRegister() != nil:
		return []interface{}{a.Visit(ctx.LocationInRegister())}

	case ctx.LocationOffsetByRegister() != nil:
		return a.Visit(ctx.LocationOffsetByRegister())

	case ctx.POP() != nil:
		return []interface{}{uint16(dcpu.Pop)}

	case ctx.PEEK() != nil:
		return []interface{}{uint16(dcpu.Peek)}

	case ctx.Pick() != nil:
		return a.Visit(ctx.Pick())

	case ctx.STACK_POINTER() != nil:
		return []interface{}{uint16(dcpu.StackPointer)}

	case ctx.PROGRAM_COUNTER() != nil:
		return []interface{}{uint16(dcpu.ProgramCounter)}

	case ctx.EXTRA() != nil:
		return []interface{}{uint16(dcpu.Extra)}

	case ctx.Location() != nil:
		return []interface{}{uint16(dcpu.Location), a.Visit(ctx.Location())}

	case ctx.Label() != nil:
		return []interface{}{uint16(dcpu.Literal), a.Visit(ctx.Label())}

	case ctx.Value() != nil:
		value := a.Visit(ctx.Value()).(uint16)
		switch {
		case value == 0xffff:
			return []interface{}{uint16(dcpu.LiteralNegative1)}
		case 0 <= value && value <= 30:
			return []interface{}{uint16(dcpu.Literal0 + value)}
		default:
			return []interface{}{uint16(dcpu.Literal), value}
		}
	}
	return []interface{}{uint16(0)}
}

func (a *Assembler) VisitLocation(ctx *LocationContext) interface{} {
	switch {
	case ctx.Label() != nil:
		return a.Visit(ctx.Label())

	case ctx.Value() != nil:
		return a.Visit(ctx.Value())
	}
	return nil
}

func (a *Assembler) VisitLabel(ctx *LabelContext) interface{} {
	return ctx.IDENTIFIER().GetText()
}

func (a *Assembler) VisitValue(ctx *ValueContext) interface{} {
	return parseValue(ctx.NUMBER().GetText())
}

func parseValue(text string) uint16 {
	switch {
	case strings.HasPrefix(text, "0x"):
		value, err := strconv.ParseUint(text[2:], 16, 16)
		if err != nil {
			log.Fatalln(err)
		}
		return uint16(value)

	case strings.HasPrefix(text, "0b"):
		value, err := strconv.ParseUint(text[2:], 2, 16)
		if err != nil {
			log.Fatalln(err)
		}
		return uint16(value)

	case strings.HasPrefix(text, "0") && text != "0":
		value, err := strconv.ParseUint(text[1:], 8, 16)
		if err != nil {
			log.Fatalln(err)
		}
		return uint16(value)

	default:
		value, err := strconv.Atoi(text)
		if err != nil {
			log.Fatalln(err)
		}
		return uint16(value)
	}
}
