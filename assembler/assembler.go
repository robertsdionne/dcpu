package assembler

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode/utf16"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/robertsdionne/dcpu"
	"github.com/robertsdionne/dcpu/parser"
)

//go:generate java -jar $HOME/Downloads/antlr-4.6-complete.jar -Dlanguage=Go -package parser DCPU.g4 -visitor

func AssembleFile(filename string) (program []uint16, err error) {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	program = Assemble(string(source))
	return
}

func Assemble(source string) (program []uint16) {
	input := antlr.NewInputStream(source)
	lexer := parser.NewDCPULexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewDCPUParser(stream)
	p.BuildParseTrees = true
	visitor := assembler{
		BaseDCPUVisitor: &parser.BaseDCPUVisitor{},
	}
	program = visitor.Visit(p.Program()).([]uint16)
	return
}

type assembler struct {
	*parser.BaseDCPUVisitor
}

func (a *assembler) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(a)
}

func (a *assembler) VisitChildren(node antlr.RuleNode) interface{} {
	var results []interface{}
	for _, child := range node.GetChildren() {
		results = append(results, child.(antlr.ParseTree).Accept(a))
	}
	return results
}

func (a *assembler) VisitProgram(ctx *parser.ProgramContext) interface{} {
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

func (a *assembler) VisitInstruction(ctx *parser.InstructionContext) interface{} {
	switch {
	case ctx.BinaryOperation() != nil:
		return a.Visit(ctx.BinaryOperation())

	case ctx.UnaryOperation() != nil:
		return a.Visit(ctx.UnaryOperation())

	case ctx.DebugOperation() != nil:
		return a.Visit(ctx.DebugOperation())
	}
	return nil
}

func (a *assembler) VisitLabelDefinition(ctx *parser.LabelDefinitionContext) interface{} {
	return ctx.IDENTIFIER().GetText()
}

func (a *assembler) VisitDataSection(ctx *parser.DataSectionContext) interface{} {
	return a.Visit(ctx.Data())
}

func (a *assembler) VisitData(ctx *parser.DataContext) interface{} {
	var result []interface{}
	for _, child := range ctx.GetChildren() {
		switch child := child.(type) {
		case *parser.DatumContext:
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

func (a *assembler) VisitDatum(ctx *parser.DatumContext) interface{} {
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

func (a *assembler) VisitBinaryOperation(ctx *parser.BinaryOperationContext) interface{} {
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

func (a *assembler) VisitBinaryOpcode(ctx *parser.BinaryOpcodeContext) interface{} {
	return binaryOpcodeValue(ctx)
}

func binaryOpcodeValue(ctx *parser.BinaryOpcodeContext) uint16 {
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

func (a *assembler) VisitUnaryOperation(ctx *parser.UnaryOperationContext) interface{} {
	opcode := a.Visit(ctx.UnaryOpcode()).(uint16)
	argument := a.Visit(ctx.ArgumentA()).([]interface{})
	result := []interface{}{opcode<<dcpu.SpecialOpcodeShift | argument[0].(uint16)<<dcpu.SpecialValueShiftA}
	result = append(result, argument[1:]...)
	return result
}

func (a *assembler) VisitUnaryOpcode(ctx *parser.UnaryOpcodeContext) interface{} {
	return unaryOpcodeValue(ctx)
}

func unaryOpcodeValue(ctx *parser.UnaryOpcodeContext) uint16 {
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

func (a *assembler) VisitDebugOperation(ctx *parser.DebugOperationContext) interface{} {
	opcode := a.Visit(ctx.DebugOpcode()).(uint16)
	return []interface{}{opcode << dcpu.DebugOpcodeShift}
}

func (a *assembler) VisitDebugOpcode(ctx *parser.DebugOpcodeContext) interface{} {
	return debugOpcodeValue(ctx)
}

func debugOpcodeValue(ctx *parser.DebugOpcodeContext) uint16 {
	switch {
	case ctx.ALT() != nil:
		return dcpu.Alert

	case ctx.DUM() != nil:
		return dcpu.DumpState

	default:
		return 0
	}
}

func (a *assembler) VisitArgumentB(ctx *parser.ArgumentBContext) interface{} {
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

func (a *assembler) VisitRegister(ctx *parser.RegisterContext) interface{} {
	return registerValue(ctx.REGISTER())
}

func (a *assembler) VisitLocationInRegister(ctx *parser.LocationInRegisterContext) interface{} {
	return registerValue(ctx.REGISTER()) + dcpu.LocationInRegisterA
}

func (a *assembler) VisitLocationOffsetByRegister(ctx *parser.LocationOffsetByRegisterContext) interface{} {
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

func (a *assembler) VisitArgumentA(ctx *parser.ArgumentAContext) interface{} {
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

func (a *assembler) VisitLocation(ctx *parser.LocationContext) interface{} {
	switch {
	case ctx.Label() != nil:
		return a.Visit(ctx.Label())

	case ctx.Value() != nil:
		return a.Visit(ctx.Value())
	}
	return nil
}

func (a *assembler) VisitLabel(ctx *parser.LabelContext) interface{} {
	return ctx.IDENTIFIER().GetText()
}

func (a *assembler) VisitValue(ctx *parser.ValueContext) interface{} {
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
