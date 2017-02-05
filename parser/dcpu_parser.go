// Generated from DCPU.g4 by ANTLR 4.6.

package parser // DCPU
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 3, 62, 155, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 18,
	4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2, 3, 2, 3,
	2, 6, 2, 48, 10, 2, 13, 2, 14, 2, 49, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 5, 5, 62, 10, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 7, 7, 7, 70, 10, 7, 12, 7, 14, 7, 73, 11, 7, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 96, 10, 11, 3, 12,
	3, 12, 3, 12, 5, 12, 101, 10, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 115, 10, 13, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 134, 10, 18, 3, 18, 3, 18,
	5, 18, 138, 10, 18, 3, 18, 3, 18, 3, 18, 5, 18, 143, 10, 18, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 2, 2,
	23, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 2, 7, 3, 2, 4, 7, 4, 2, 46, 46, 60, 61, 3, 2, 12, 38, 3, 2,
	48, 56, 3, 2, 57, 58, 163, 2, 47, 3, 2, 2, 2, 4, 53, 3, 2, 2, 2, 6, 56,
	3, 2, 2, 2, 8, 61, 3, 2, 2, 2, 10, 63, 3, 2, 2, 2, 12, 66, 3, 2, 2, 2,
	14, 74, 3, 2, 2, 2, 16, 76, 3, 2, 2, 2, 18, 81, 3, 2, 2, 2, 20, 95, 3,
	2, 2, 2, 22, 97, 3, 2, 2, 2, 24, 114, 3, 2, 2, 2, 26, 116, 3, 2, 2, 2,
	28, 118, 3, 2, 2, 2, 30, 121, 3, 2, 2, 2, 32, 124, 3, 2, 2, 2, 34, 128,
	3, 2, 2, 2, 36, 146, 3, 2, 2, 2, 38, 148, 3, 2, 2, 2, 40, 150, 3, 2, 2,
	2, 42, 152, 3, 2, 2, 2, 44, 48, 5, 4, 3, 2, 45, 48, 5, 8, 5, 2, 46, 48,
	5, 10, 6, 2, 47, 44, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 46, 3, 2, 2, 2,
	48, 49, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 51, 3,
	2, 2, 2, 51, 52, 7, 2, 2, 3, 52, 3, 3, 2, 2, 2, 53, 54, 7, 3, 2, 2, 54,
	55, 7, 60, 2, 2, 55, 5, 3, 2, 2, 2, 56, 57, 7, 60, 2, 2, 57, 7, 3, 2, 2,
	2, 58, 62, 5, 16, 9, 2, 59, 62, 5, 30, 16, 2, 60, 62, 5, 40, 21, 2, 61,
	58, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 60, 3, 2, 2, 2, 62, 9, 3, 2, 2,
	2, 63, 64, 9, 2, 2, 2, 64, 65, 5, 12, 7, 2, 65, 11, 3, 2, 2, 2, 66, 71,
	5, 14, 8, 2, 67, 68, 7, 8, 2, 2, 68, 70, 5, 14, 8, 2, 69, 67, 3, 2, 2,
	2, 70, 73, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 13,
	3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 74, 75, 9, 3, 2, 2, 75, 15, 3, 2, 2, 2,
	76, 77, 5, 18, 10, 2, 77, 78, 5, 24, 13, 2, 78, 79, 7, 8, 2, 2, 79, 80,
	5, 20, 11, 2, 80, 17, 3, 2, 2, 2, 81, 82, 9, 4, 2, 2, 82, 19, 3, 2, 2,
	2, 83, 96, 5, 26, 14, 2, 84, 96, 5, 32, 17, 2, 85, 96, 5, 34, 18, 2, 86,
	96, 7, 39, 2, 2, 87, 96, 7, 41, 2, 2, 88, 96, 5, 28, 15, 2, 89, 96, 7,
	43, 2, 2, 90, 96, 7, 44, 2, 2, 91, 96, 7, 45, 2, 2, 92, 96, 5, 22, 12,
	2, 93, 96, 5, 6, 4, 2, 94, 96, 5, 36, 19, 2, 95, 83, 3, 2, 2, 2, 95, 84,
	3, 2, 2, 2, 95, 85, 3, 2, 2, 2, 95, 86, 3, 2, 2, 2, 95, 87, 3, 2, 2, 2,
	95, 88, 3, 2, 2, 2, 95, 89, 3, 2, 2, 2, 95, 90, 3, 2, 2, 2, 95, 91, 3,
	2, 2, 2, 95, 92, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 95, 94, 3, 2, 2, 2, 96,
	21, 3, 2, 2, 2, 97, 100, 7, 9, 2, 2, 98, 101, 5, 6, 4, 2, 99, 101, 5, 36,
	19, 2, 100, 98, 3, 2, 2, 2, 100, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2,
	102, 103, 7, 10, 2, 2, 103, 23, 3, 2, 2, 2, 104, 115, 5, 26, 14, 2, 105,
	115, 5, 32, 17, 2, 106, 115, 5, 34, 18, 2, 107, 115, 7, 40, 2, 2, 108,
	115, 7, 41, 2, 2, 109, 115, 5, 28, 15, 2, 110, 115, 7, 43, 2, 2, 111, 115,
	7, 44, 2, 2, 112, 115, 7, 45, 2, 2, 113, 115, 5, 22, 12, 2, 114, 104, 3,
	2, 2, 2, 114, 105, 3, 2, 2, 2, 114, 106, 3, 2, 2, 2, 114, 107, 3, 2, 2,
	2, 114, 108, 3, 2, 2, 2, 114, 109, 3, 2, 2, 2, 114, 110, 3, 2, 2, 2, 114,
	111, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 114, 113, 3, 2, 2, 2, 115, 25, 3,
	2, 2, 2, 116, 117, 7, 47, 2, 2, 117, 27, 3, 2, 2, 2, 118, 119, 7, 42, 2,
	2, 119, 120, 7, 46, 2, 2, 120, 29, 3, 2, 2, 2, 121, 122, 5, 38, 20, 2,
	122, 123, 5, 20, 11, 2, 123, 31, 3, 2, 2, 2, 124, 125, 7, 9, 2, 2, 125,
	126, 7, 47, 2, 2, 126, 127, 7, 10, 2, 2, 127, 33, 3, 2, 2, 2, 128, 142,
	7, 9, 2, 2, 129, 130, 7, 47, 2, 2, 130, 133, 7, 11, 2, 2, 131, 134, 5,
	6, 4, 2, 132, 134, 5, 36, 19, 2, 133, 131, 3, 2, 2, 2, 133, 132, 3, 2,
	2, 2, 134, 143, 3, 2, 2, 2, 135, 138, 5, 6, 4, 2, 136, 138, 5, 36, 19,
	2, 137, 135, 3, 2, 2, 2, 137, 136, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139,
	140, 7, 11, 2, 2, 140, 141, 7, 47, 2, 2, 141, 143, 3, 2, 2, 2, 142, 129,
	3, 2, 2, 2, 142, 137, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 7, 10,
	2, 2, 145, 35, 3, 2, 2, 2, 146, 147, 7, 46, 2, 2, 147, 37, 3, 2, 2, 2,
	148, 149, 9, 5, 2, 2, 149, 39, 3, 2, 2, 2, 150, 151, 5, 42, 22, 2, 151,
	41, 3, 2, 2, 2, 152, 153, 9, 6, 2, 2, 153, 43, 3, 2, 2, 2, 12, 47, 49,
	61, 71, 95, 100, 114, 133, 137, 142,
}

var deserializer = antlr.NewATNDeserializer(nil)

var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':'", "'.dat'", "'.DAT'", "'dat'", "'DAT'", "','", "'['", "']'", "'+'",
}

var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "SET", "ADD", "SUB", "MUL", "MLI",
	"DIV", "DVI", "MOD", "MDI", "AND", "BOR", "XOR", "SHR", "ASR", "SHL", "IFB",
	"IFC", "IFE", "IFN", "IFG", "IFA", "IFL", "IFU", "ADX", "SBX", "STI", "STD",
	"POP", "PUSH", "PEEK", "PICK", "STACK_POINTER", "PROGRAM_COUNTER", "EXTRA",
	"NUMBER", "REGISTER", "JSR", "INT", "IAG", "IAS", "RFI", "IAQ", "HWN",
	"HWQ", "HWI", "ALT", "DUM", "COMMENT", "IDENTIFIER", "STRING", "WHITESPACE",
}

var ruleNames = []string{
	"program", "labelDefinition", "label", "instruction", "dataSection", "data",
	"datum", "binaryOperation", "binaryOpcode", "argumentA", "location", "argumentB",
	"register", "pick", "unaryOperation", "locationInRegister", "locationOffsetByRegister",
	"value", "unaryOpcode", "debugOperation", "debugOpcode",
}

type DCPUParser struct {
	*antlr.BaseParser
}

func NewDCPUParser(input antlr.TokenStream) *DCPUParser {
	var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	var sharedContextCache = antlr.NewPredictionContextCache()

	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	this := new(DCPUParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, sharedContextCache)
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "DCPU.g4"

	return this
}

// DCPUParser tokens.
const (
	DCPUParserEOF             = antlr.TokenEOF
	DCPUParserT__0            = 1
	DCPUParserT__1            = 2
	DCPUParserT__2            = 3
	DCPUParserT__3            = 4
	DCPUParserT__4            = 5
	DCPUParserT__5            = 6
	DCPUParserT__6            = 7
	DCPUParserT__7            = 8
	DCPUParserT__8            = 9
	DCPUParserSET             = 10
	DCPUParserADD             = 11
	DCPUParserSUB             = 12
	DCPUParserMUL             = 13
	DCPUParserMLI             = 14
	DCPUParserDIV             = 15
	DCPUParserDVI             = 16
	DCPUParserMOD             = 17
	DCPUParserMDI             = 18
	DCPUParserAND             = 19
	DCPUParserBOR             = 20
	DCPUParserXOR             = 21
	DCPUParserSHR             = 22
	DCPUParserASR             = 23
	DCPUParserSHL             = 24
	DCPUParserIFB             = 25
	DCPUParserIFC             = 26
	DCPUParserIFE             = 27
	DCPUParserIFN             = 28
	DCPUParserIFG             = 29
	DCPUParserIFA             = 30
	DCPUParserIFL             = 31
	DCPUParserIFU             = 32
	DCPUParserADX             = 33
	DCPUParserSBX             = 34
	DCPUParserSTI             = 35
	DCPUParserSTD             = 36
	DCPUParserPOP             = 37
	DCPUParserPUSH            = 38
	DCPUParserPEEK            = 39
	DCPUParserPICK            = 40
	DCPUParserSTACK_POINTER   = 41
	DCPUParserPROGRAM_COUNTER = 42
	DCPUParserEXTRA           = 43
	DCPUParserNUMBER          = 44
	DCPUParserREGISTER        = 45
	DCPUParserJSR             = 46
	DCPUParserINT             = 47
	DCPUParserIAG             = 48
	DCPUParserIAS             = 49
	DCPUParserRFI             = 50
	DCPUParserIAQ             = 51
	DCPUParserHWN             = 52
	DCPUParserHWQ             = 53
	DCPUParserHWI             = 54
	DCPUParserALT             = 55
	DCPUParserDUM             = 56
	DCPUParserCOMMENT         = 57
	DCPUParserIDENTIFIER      = 58
	DCPUParserSTRING          = 59
	DCPUParserWHITESPACE      = 60
)

// DCPUParser rules.
const (
	DCPUParserRULE_program                  = 0
	DCPUParserRULE_labelDefinition          = 1
	DCPUParserRULE_label                    = 2
	DCPUParserRULE_instruction              = 3
	DCPUParserRULE_dataSection              = 4
	DCPUParserRULE_data                     = 5
	DCPUParserRULE_datum                    = 6
	DCPUParserRULE_binaryOperation          = 7
	DCPUParserRULE_binaryOpcode             = 8
	DCPUParserRULE_argumentA                = 9
	DCPUParserRULE_location                 = 10
	DCPUParserRULE_argumentB                = 11
	DCPUParserRULE_register                 = 12
	DCPUParserRULE_pick                     = 13
	DCPUParserRULE_unaryOperation           = 14
	DCPUParserRULE_locationInRegister       = 15
	DCPUParserRULE_locationOffsetByRegister = 16
	DCPUParserRULE_value                    = 17
	DCPUParserRULE_unaryOpcode              = 18
	DCPUParserRULE_debugOperation           = 19
	DCPUParserRULE_debugOpcode              = 20
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(DCPUParserEOF, 0)
}

func (s *ProgramContext) AllLabelDefinition() []ILabelDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILabelDefinitionContext)(nil)).Elem())
	var tst = make([]ILabelDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILabelDefinitionContext)
		}
	}

	return tst
}

func (s *ProgramContext) LabelDefinition(i int) ILabelDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILabelDefinitionContext)
}

func (s *ProgramContext) AllInstruction() []IInstructionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstructionContext)(nil)).Elem())
	var tst = make([]IInstructionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstructionContext)
		}
	}

	return tst
}

func (s *ProgramContext) Instruction(i int) IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *ProgramContext) AllDataSection() []IDataSectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDataSectionContext)(nil)).Elem())
	var tst = make([]IDataSectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDataSectionContext)
		}
	}

	return tst
}

func (s *ProgramContext) DataSection(i int) IDataSectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataSectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDataSectionContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DCPUParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DCPUParserT__0)|(1<<DCPUParserT__1)|(1<<DCPUParserT__2)|(1<<DCPUParserT__3)|(1<<DCPUParserT__4)|(1<<DCPUParserSET)|(1<<DCPUParserADD)|(1<<DCPUParserSUB)|(1<<DCPUParserMUL)|(1<<DCPUParserMLI)|(1<<DCPUParserDIV)|(1<<DCPUParserDVI)|(1<<DCPUParserMOD)|(1<<DCPUParserMDI)|(1<<DCPUParserAND)|(1<<DCPUParserBOR)|(1<<DCPUParserXOR)|(1<<DCPUParserSHR)|(1<<DCPUParserASR)|(1<<DCPUParserSHL)|(1<<DCPUParserIFB)|(1<<DCPUParserIFC)|(1<<DCPUParserIFE)|(1<<DCPUParserIFN)|(1<<DCPUParserIFG)|(1<<DCPUParserIFA)|(1<<DCPUParserIFL))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint(_la-32))&((1<<(DCPUParserIFU-32))|(1<<(DCPUParserADX-32))|(1<<(DCPUParserSBX-32))|(1<<(DCPUParserSTI-32))|(1<<(DCPUParserSTD-32))|(1<<(DCPUParserJSR-32))|(1<<(DCPUParserINT-32))|(1<<(DCPUParserIAG-32))|(1<<(DCPUParserIAS-32))|(1<<(DCPUParserRFI-32))|(1<<(DCPUParserIAQ-32))|(1<<(DCPUParserHWN-32))|(1<<(DCPUParserHWQ-32))|(1<<(DCPUParserHWI-32))|(1<<(DCPUParserALT-32))|(1<<(DCPUParserDUM-32)))) != 0) {
		p.SetState(45)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case DCPUParserT__0:
			{
				p.SetState(42)
				p.LabelDefinition()
			}

		case DCPUParserSET, DCPUParserADD, DCPUParserSUB, DCPUParserMUL, DCPUParserMLI, DCPUParserDIV, DCPUParserDVI, DCPUParserMOD, DCPUParserMDI, DCPUParserAND, DCPUParserBOR, DCPUParserXOR, DCPUParserSHR, DCPUParserASR, DCPUParserSHL, DCPUParserIFB, DCPUParserIFC, DCPUParserIFE, DCPUParserIFN, DCPUParserIFG, DCPUParserIFA, DCPUParserIFL, DCPUParserIFU, DCPUParserADX, DCPUParserSBX, DCPUParserSTI, DCPUParserSTD, DCPUParserJSR, DCPUParserINT, DCPUParserIAG, DCPUParserIAS, DCPUParserRFI, DCPUParserIAQ, DCPUParserHWN, DCPUParserHWQ, DCPUParserHWI, DCPUParserALT, DCPUParserDUM:
			{
				p.SetState(43)
				p.Instruction()
			}

		case DCPUParserT__1, DCPUParserT__2, DCPUParserT__3, DCPUParserT__4:
			{
				p.SetState(44)
				p.DataSection()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(49)
		p.Match(DCPUParserEOF)
	}

	return localctx
}

// ILabelDefinitionContext is an interface to support dynamic dispatch.
type ILabelDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelDefinitionContext differentiates from other interfaces.
	IsLabelDefinitionContext()
}

type LabelDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelDefinitionContext() *LabelDefinitionContext {
	var p = new(LabelDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_labelDefinition
	return p
}

func (*LabelDefinitionContext) IsLabelDefinitionContext() {}

func NewLabelDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelDefinitionContext {
	var p = new(LabelDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_labelDefinition

	return p
}

func (s *LabelDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelDefinitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DCPUParserIDENTIFIER, 0)
}

func (s *LabelDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterLabelDefinition(s)
	}
}

func (s *LabelDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitLabelDefinition(s)
	}
}

func (s *LabelDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitLabelDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) LabelDefinition() (localctx ILabelDefinitionContext) {
	localctx = NewLabelDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DCPUParserRULE_labelDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(DCPUParserT__0)
	}
	{
		p.SetState(52)
		p.Match(DCPUParserIDENTIFIER)
	}

	return localctx
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DCPUParserIDENTIFIER, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (s *LabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DCPUParserRULE_label)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(DCPUParserIDENTIFIER)
	}

	return localctx
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) BinaryOperation() IBinaryOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryOperationContext)
}

func (s *InstructionContext) UnaryOperation() IUnaryOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryOperationContext)
}

func (s *InstructionContext) DebugOperation() IDebugOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDebugOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDebugOperationContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (s *InstructionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitInstruction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DCPUParserRULE_instruction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(59)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DCPUParserSET, DCPUParserADD, DCPUParserSUB, DCPUParserMUL, DCPUParserMLI, DCPUParserDIV, DCPUParserDVI, DCPUParserMOD, DCPUParserMDI, DCPUParserAND, DCPUParserBOR, DCPUParserXOR, DCPUParserSHR, DCPUParserASR, DCPUParserSHL, DCPUParserIFB, DCPUParserIFC, DCPUParserIFE, DCPUParserIFN, DCPUParserIFG, DCPUParserIFA, DCPUParserIFL, DCPUParserIFU, DCPUParserADX, DCPUParserSBX, DCPUParserSTI, DCPUParserSTD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.BinaryOperation()
		}

	case DCPUParserJSR, DCPUParserINT, DCPUParserIAG, DCPUParserIAS, DCPUParserRFI, DCPUParserIAQ, DCPUParserHWN, DCPUParserHWQ, DCPUParserHWI:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.UnaryOperation()
		}

	case DCPUParserALT, DCPUParserDUM:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(58)
			p.DebugOperation()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDataSectionContext is an interface to support dynamic dispatch.
type IDataSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataSectionContext differentiates from other interfaces.
	IsDataSectionContext()
}

type DataSectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataSectionContext() *DataSectionContext {
	var p = new(DataSectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_dataSection
	return p
}

func (*DataSectionContext) IsDataSectionContext() {}

func NewDataSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataSectionContext {
	var p = new(DataSectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_dataSection

	return p
}

func (s *DataSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *DataSectionContext) Data() IDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataContext)
}

func (s *DataSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterDataSection(s)
	}
}

func (s *DataSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitDataSection(s)
	}
}

func (s *DataSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitDataSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) DataSection() (localctx IDataSectionContext) {
	localctx = NewDataSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DCPUParserRULE_dataSection)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(61)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DCPUParserT__1)|(1<<DCPUParserT__2)|(1<<DCPUParserT__3)|(1<<DCPUParserT__4))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(62)
		p.Data()
	}

	return localctx
}

// IDataContext is an interface to support dynamic dispatch.
type IDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataContext differentiates from other interfaces.
	IsDataContext()
}

type DataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataContext() *DataContext {
	var p = new(DataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_data
	return p
}

func (*DataContext) IsDataContext() {}

func NewDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataContext {
	var p = new(DataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_data

	return p
}

func (s *DataContext) GetParser() antlr.Parser { return s.parser }

func (s *DataContext) AllDatum() []IDatumContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDatumContext)(nil)).Elem())
	var tst = make([]IDatumContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDatumContext)
		}
	}

	return tst
}

func (s *DataContext) Datum(i int) IDatumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatumContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDatumContext)
}

func (s *DataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterData(s)
	}
}

func (s *DataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitData(s)
	}
}

func (s *DataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitData(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) Data() (localctx IDataContext) {
	localctx = NewDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DCPUParserRULE_data)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Datum()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DCPUParserT__5 {
		{
			p.SetState(65)
			p.Match(DCPUParserT__5)
		}
		{
			p.SetState(66)
			p.Datum()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDatumContext is an interface to support dynamic dispatch.
type IDatumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatumContext differentiates from other interfaces.
	IsDatumContext()
}

type DatumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatumContext() *DatumContext {
	var p = new(DatumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_datum
	return p
}

func (*DatumContext) IsDatumContext() {}

func NewDatumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatumContext {
	var p = new(DatumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_datum

	return p
}

func (s *DatumContext) GetParser() antlr.Parser { return s.parser }

func (s *DatumContext) STRING() antlr.TerminalNode {
	return s.GetToken(DCPUParserSTRING, 0)
}

func (s *DatumContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DCPUParserIDENTIFIER, 0)
}

func (s *DatumContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(DCPUParserNUMBER, 0)
}

func (s *DatumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterDatum(s)
	}
}

func (s *DatumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitDatum(s)
	}
}

func (s *DatumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitDatum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) Datum() (localctx IDatumContext) {
	localctx = NewDatumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DCPUParserRULE_datum)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(72)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-44)&-(0x1f+1)) == 0 && ((1<<uint(_la-44))&((1<<(DCPUParserNUMBER-44))|(1<<(DCPUParserIDENTIFIER-44))|(1<<(DCPUParserSTRING-44)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IBinaryOperationContext is an interface to support dynamic dispatch.
type IBinaryOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryOperationContext differentiates from other interfaces.
	IsBinaryOperationContext()
}

type BinaryOperationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOperationContext() *BinaryOperationContext {
	var p = new(BinaryOperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_binaryOperation
	return p
}

func (*BinaryOperationContext) IsBinaryOperationContext() {}

func NewBinaryOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOperationContext {
	var p = new(BinaryOperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_binaryOperation

	return p
}

func (s *BinaryOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOperationContext) BinaryOpcode() IBinaryOpcodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryOpcodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryOpcodeContext)
}

func (s *BinaryOperationContext) ArgumentB() IArgumentBContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentBContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentBContext)
}

func (s *BinaryOperationContext) ArgumentA() IArgumentAContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentAContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentAContext)
}

func (s *BinaryOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterBinaryOperation(s)
	}
}

func (s *BinaryOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitBinaryOperation(s)
	}
}

func (s *BinaryOperationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitBinaryOperation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) BinaryOperation() (localctx IBinaryOperationContext) {
	localctx = NewBinaryOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DCPUParserRULE_binaryOperation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.BinaryOpcode()
	}
	{
		p.SetState(75)
		p.ArgumentB()
	}
	{
		p.SetState(76)
		p.Match(DCPUParserT__5)
	}
	{
		p.SetState(77)
		p.ArgumentA()
	}

	return localctx
}

// IBinaryOpcodeContext is an interface to support dynamic dispatch.
type IBinaryOpcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryOpcodeContext differentiates from other interfaces.
	IsBinaryOpcodeContext()
}

type BinaryOpcodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOpcodeContext() *BinaryOpcodeContext {
	var p = new(BinaryOpcodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_binaryOpcode
	return p
}

func (*BinaryOpcodeContext) IsBinaryOpcodeContext() {}

func NewBinaryOpcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOpcodeContext {
	var p = new(BinaryOpcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_binaryOpcode

	return p
}

func (s *BinaryOpcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOpcodeContext) SET() antlr.TerminalNode {
	return s.GetToken(DCPUParserSET, 0)
}

func (s *BinaryOpcodeContext) ADD() antlr.TerminalNode {
	return s.GetToken(DCPUParserADD, 0)
}

func (s *BinaryOpcodeContext) SUB() antlr.TerminalNode {
	return s.GetToken(DCPUParserSUB, 0)
}

func (s *BinaryOpcodeContext) MUL() antlr.TerminalNode {
	return s.GetToken(DCPUParserMUL, 0)
}

func (s *BinaryOpcodeContext) MLI() antlr.TerminalNode {
	return s.GetToken(DCPUParserMLI, 0)
}

func (s *BinaryOpcodeContext) DIV() antlr.TerminalNode {
	return s.GetToken(DCPUParserDIV, 0)
}

func (s *BinaryOpcodeContext) DVI() antlr.TerminalNode {
	return s.GetToken(DCPUParserDVI, 0)
}

func (s *BinaryOpcodeContext) MOD() antlr.TerminalNode {
	return s.GetToken(DCPUParserMOD, 0)
}

func (s *BinaryOpcodeContext) MDI() antlr.TerminalNode {
	return s.GetToken(DCPUParserMDI, 0)
}

func (s *BinaryOpcodeContext) AND() antlr.TerminalNode {
	return s.GetToken(DCPUParserAND, 0)
}

func (s *BinaryOpcodeContext) BOR() antlr.TerminalNode {
	return s.GetToken(DCPUParserBOR, 0)
}

func (s *BinaryOpcodeContext) XOR() antlr.TerminalNode {
	return s.GetToken(DCPUParserXOR, 0)
}

func (s *BinaryOpcodeContext) SHR() antlr.TerminalNode {
	return s.GetToken(DCPUParserSHR, 0)
}

func (s *BinaryOpcodeContext) ASR() antlr.TerminalNode {
	return s.GetToken(DCPUParserASR, 0)
}

func (s *BinaryOpcodeContext) SHL() antlr.TerminalNode {
	return s.GetToken(DCPUParserSHL, 0)
}

func (s *BinaryOpcodeContext) IFB() antlr.TerminalNode {
	return s.GetToken(DCPUParserIFB, 0)
}

func (s *BinaryOpcodeContext) IFC() antlr.TerminalNode {
	return s.GetToken(DCPUParserIFC, 0)
}

func (s *BinaryOpcodeContext) IFE() antlr.TerminalNode {
	return s.GetToken(DCPUParserIFE, 0)
}

func (s *BinaryOpcodeContext) IFN() antlr.TerminalNode {
	return s.GetToken(DCPUParserIFN, 0)
}

func (s *BinaryOpcodeContext) IFG() antlr.TerminalNode {
	return s.GetToken(DCPUParserIFG, 0)
}

func (s *BinaryOpcodeContext) IFA() antlr.TerminalNode {
	return s.GetToken(DCPUParserIFA, 0)
}

func (s *BinaryOpcodeContext) IFL() antlr.TerminalNode {
	return s.GetToken(DCPUParserIFL, 0)
}

func (s *BinaryOpcodeContext) IFU() antlr.TerminalNode {
	return s.GetToken(DCPUParserIFU, 0)
}

func (s *BinaryOpcodeContext) ADX() antlr.TerminalNode {
	return s.GetToken(DCPUParserADX, 0)
}

func (s *BinaryOpcodeContext) SBX() antlr.TerminalNode {
	return s.GetToken(DCPUParserSBX, 0)
}

func (s *BinaryOpcodeContext) STI() antlr.TerminalNode {
	return s.GetToken(DCPUParserSTI, 0)
}

func (s *BinaryOpcodeContext) STD() antlr.TerminalNode {
	return s.GetToken(DCPUParserSTD, 0)
}

func (s *BinaryOpcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOpcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOpcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterBinaryOpcode(s)
	}
}

func (s *BinaryOpcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitBinaryOpcode(s)
	}
}

func (s *BinaryOpcodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitBinaryOpcode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) BinaryOpcode() (localctx IBinaryOpcodeContext) {
	localctx = NewBinaryOpcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DCPUParserRULE_binaryOpcode)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-10)&-(0x1f+1)) == 0 && ((1<<uint(_la-10))&((1<<(DCPUParserSET-10))|(1<<(DCPUParserADD-10))|(1<<(DCPUParserSUB-10))|(1<<(DCPUParserMUL-10))|(1<<(DCPUParserMLI-10))|(1<<(DCPUParserDIV-10))|(1<<(DCPUParserDVI-10))|(1<<(DCPUParserMOD-10))|(1<<(DCPUParserMDI-10))|(1<<(DCPUParserAND-10))|(1<<(DCPUParserBOR-10))|(1<<(DCPUParserXOR-10))|(1<<(DCPUParserSHR-10))|(1<<(DCPUParserASR-10))|(1<<(DCPUParserSHL-10))|(1<<(DCPUParserIFB-10))|(1<<(DCPUParserIFC-10))|(1<<(DCPUParserIFE-10))|(1<<(DCPUParserIFN-10))|(1<<(DCPUParserIFG-10))|(1<<(DCPUParserIFA-10))|(1<<(DCPUParserIFL-10))|(1<<(DCPUParserIFU-10))|(1<<(DCPUParserADX-10))|(1<<(DCPUParserSBX-10))|(1<<(DCPUParserSTI-10))|(1<<(DCPUParserSTD-10)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IArgumentAContext is an interface to support dynamic dispatch.
type IArgumentAContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentAContext differentiates from other interfaces.
	IsArgumentAContext()
}

type ArgumentAContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentAContext() *ArgumentAContext {
	var p = new(ArgumentAContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_argumentA
	return p
}

func (*ArgumentAContext) IsArgumentAContext() {}

func NewArgumentAContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentAContext {
	var p = new(ArgumentAContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_argumentA

	return p
}

func (s *ArgumentAContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentAContext) Register() IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *ArgumentAContext) LocationInRegister() ILocationInRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocationInRegisterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocationInRegisterContext)
}

func (s *ArgumentAContext) LocationOffsetByRegister() ILocationOffsetByRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocationOffsetByRegisterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocationOffsetByRegisterContext)
}

func (s *ArgumentAContext) POP() antlr.TerminalNode {
	return s.GetToken(DCPUParserPOP, 0)
}

func (s *ArgumentAContext) PEEK() antlr.TerminalNode {
	return s.GetToken(DCPUParserPEEK, 0)
}

func (s *ArgumentAContext) Pick() IPickContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPickContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPickContext)
}

func (s *ArgumentAContext) STACK_POINTER() antlr.TerminalNode {
	return s.GetToken(DCPUParserSTACK_POINTER, 0)
}

func (s *ArgumentAContext) PROGRAM_COUNTER() antlr.TerminalNode {
	return s.GetToken(DCPUParserPROGRAM_COUNTER, 0)
}

func (s *ArgumentAContext) EXTRA() antlr.TerminalNode {
	return s.GetToken(DCPUParserEXTRA, 0)
}

func (s *ArgumentAContext) Location() ILocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocationContext)
}

func (s *ArgumentAContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *ArgumentAContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArgumentAContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentAContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentAContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterArgumentA(s)
	}
}

func (s *ArgumentAContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitArgumentA(s)
	}
}

func (s *ArgumentAContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitArgumentA(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) ArgumentA() (localctx IArgumentAContext) {
	localctx = NewArgumentAContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DCPUParserRULE_argumentA)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(81)
			p.Register()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.LocationInRegister()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(83)
			p.LocationOffsetByRegister()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(84)
			p.Match(DCPUParserPOP)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(85)
			p.Match(DCPUParserPEEK)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(86)
			p.Pick()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(87)
			p.Match(DCPUParserSTACK_POINTER)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(88)
			p.Match(DCPUParserPROGRAM_COUNTER)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(89)
			p.Match(DCPUParserEXTRA)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(90)
			p.Location()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(91)
			p.Label()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(92)
			p.Value()
		}

	}

	return localctx
}

// ILocationContext is an interface to support dynamic dispatch.
type ILocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocationContext differentiates from other interfaces.
	IsLocationContext()
}

type LocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocationContext() *LocationContext {
	var p = new(LocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_location
	return p
}

func (*LocationContext) IsLocationContext() {}

func NewLocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocationContext {
	var p = new(LocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_location

	return p
}

func (s *LocationContext) GetParser() antlr.Parser { return s.parser }

func (s *LocationContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *LocationContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *LocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterLocation(s)
	}
}

func (s *LocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitLocation(s)
	}
}

func (s *LocationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitLocation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) Location() (localctx ILocationContext) {
	localctx = NewLocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DCPUParserRULE_location)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(DCPUParserT__6)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DCPUParserIDENTIFIER:
		{
			p.SetState(96)
			p.Label()
		}

	case DCPUParserNUMBER:
		{
			p.SetState(97)
			p.Value()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(100)
		p.Match(DCPUParserT__7)
	}

	return localctx
}

// IArgumentBContext is an interface to support dynamic dispatch.
type IArgumentBContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentBContext differentiates from other interfaces.
	IsArgumentBContext()
}

type ArgumentBContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentBContext() *ArgumentBContext {
	var p = new(ArgumentBContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_argumentB
	return p
}

func (*ArgumentBContext) IsArgumentBContext() {}

func NewArgumentBContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentBContext {
	var p = new(ArgumentBContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_argumentB

	return p
}

func (s *ArgumentBContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentBContext) Register() IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *ArgumentBContext) LocationInRegister() ILocationInRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocationInRegisterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocationInRegisterContext)
}

func (s *ArgumentBContext) LocationOffsetByRegister() ILocationOffsetByRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocationOffsetByRegisterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocationOffsetByRegisterContext)
}

func (s *ArgumentBContext) PUSH() antlr.TerminalNode {
	return s.GetToken(DCPUParserPUSH, 0)
}

func (s *ArgumentBContext) PEEK() antlr.TerminalNode {
	return s.GetToken(DCPUParserPEEK, 0)
}

func (s *ArgumentBContext) Pick() IPickContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPickContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPickContext)
}

func (s *ArgumentBContext) STACK_POINTER() antlr.TerminalNode {
	return s.GetToken(DCPUParserSTACK_POINTER, 0)
}

func (s *ArgumentBContext) PROGRAM_COUNTER() antlr.TerminalNode {
	return s.GetToken(DCPUParserPROGRAM_COUNTER, 0)
}

func (s *ArgumentBContext) EXTRA() antlr.TerminalNode {
	return s.GetToken(DCPUParserEXTRA, 0)
}

func (s *ArgumentBContext) Location() ILocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocationContext)
}

func (s *ArgumentBContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentBContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentBContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterArgumentB(s)
	}
}

func (s *ArgumentBContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitArgumentB(s)
	}
}

func (s *ArgumentBContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitArgumentB(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) ArgumentB() (localctx IArgumentBContext) {
	localctx = NewArgumentBContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DCPUParserRULE_argumentB)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.Register()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.LocationInRegister()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(104)
			p.LocationOffsetByRegister()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(105)
			p.Match(DCPUParserPUSH)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(106)
			p.Match(DCPUParserPEEK)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(107)
			p.Pick()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(108)
			p.Match(DCPUParserSTACK_POINTER)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(109)
			p.Match(DCPUParserPROGRAM_COUNTER)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(110)
			p.Match(DCPUParserEXTRA)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(111)
			p.Location()
		}

	}

	return localctx
}

// IRegisterContext is an interface to support dynamic dispatch.
type IRegisterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegisterContext differentiates from other interfaces.
	IsRegisterContext()
}

type RegisterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegisterContext() *RegisterContext {
	var p = new(RegisterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_register
	return p
}

func (*RegisterContext) IsRegisterContext() {}

func NewRegisterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegisterContext {
	var p = new(RegisterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_register

	return p
}

func (s *RegisterContext) GetParser() antlr.Parser { return s.parser }

func (s *RegisterContext) REGISTER() antlr.TerminalNode {
	return s.GetToken(DCPUParserREGISTER, 0)
}

func (s *RegisterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegisterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegisterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterRegister(s)
	}
}

func (s *RegisterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitRegister(s)
	}
}

func (s *RegisterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitRegister(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) Register() (localctx IRegisterContext) {
	localctx = NewRegisterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DCPUParserRULE_register)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(DCPUParserREGISTER)
	}

	return localctx
}

// IPickContext is an interface to support dynamic dispatch.
type IPickContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPickContext differentiates from other interfaces.
	IsPickContext()
}

type PickContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPickContext() *PickContext {
	var p = new(PickContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_pick
	return p
}

func (*PickContext) IsPickContext() {}

func NewPickContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PickContext {
	var p = new(PickContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_pick

	return p
}

func (s *PickContext) GetParser() antlr.Parser { return s.parser }

func (s *PickContext) PICK() antlr.TerminalNode {
	return s.GetToken(DCPUParserPICK, 0)
}

func (s *PickContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(DCPUParserNUMBER, 0)
}

func (s *PickContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PickContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PickContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterPick(s)
	}
}

func (s *PickContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitPick(s)
	}
}

func (s *PickContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitPick(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) Pick() (localctx IPickContext) {
	localctx = NewPickContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DCPUParserRULE_pick)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(DCPUParserPICK)
	}
	{
		p.SetState(117)
		p.Match(DCPUParserNUMBER)
	}

	return localctx
}

// IUnaryOperationContext is an interface to support dynamic dispatch.
type IUnaryOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOperationContext differentiates from other interfaces.
	IsUnaryOperationContext()
}

type UnaryOperationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperationContext() *UnaryOperationContext {
	var p = new(UnaryOperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_unaryOperation
	return p
}

func (*UnaryOperationContext) IsUnaryOperationContext() {}

func NewUnaryOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperationContext {
	var p = new(UnaryOperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_unaryOperation

	return p
}

func (s *UnaryOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOperationContext) UnaryOpcode() IUnaryOpcodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOpcodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryOpcodeContext)
}

func (s *UnaryOperationContext) ArgumentA() IArgumentAContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentAContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentAContext)
}

func (s *UnaryOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterUnaryOperation(s)
	}
}

func (s *UnaryOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitUnaryOperation(s)
	}
}

func (s *UnaryOperationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitUnaryOperation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) UnaryOperation() (localctx IUnaryOperationContext) {
	localctx = NewUnaryOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DCPUParserRULE_unaryOperation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.UnaryOpcode()
	}
	{
		p.SetState(120)
		p.ArgumentA()
	}

	return localctx
}

// ILocationInRegisterContext is an interface to support dynamic dispatch.
type ILocationInRegisterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocationInRegisterContext differentiates from other interfaces.
	IsLocationInRegisterContext()
}

type LocationInRegisterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocationInRegisterContext() *LocationInRegisterContext {
	var p = new(LocationInRegisterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_locationInRegister
	return p
}

func (*LocationInRegisterContext) IsLocationInRegisterContext() {}

func NewLocationInRegisterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocationInRegisterContext {
	var p = new(LocationInRegisterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_locationInRegister

	return p
}

func (s *LocationInRegisterContext) GetParser() antlr.Parser { return s.parser }

func (s *LocationInRegisterContext) REGISTER() antlr.TerminalNode {
	return s.GetToken(DCPUParserREGISTER, 0)
}

func (s *LocationInRegisterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocationInRegisterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocationInRegisterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterLocationInRegister(s)
	}
}

func (s *LocationInRegisterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitLocationInRegister(s)
	}
}

func (s *LocationInRegisterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitLocationInRegister(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) LocationInRegister() (localctx ILocationInRegisterContext) {
	localctx = NewLocationInRegisterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DCPUParserRULE_locationInRegister)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(DCPUParserT__6)
	}
	{
		p.SetState(123)
		p.Match(DCPUParserREGISTER)
	}
	{
		p.SetState(124)
		p.Match(DCPUParserT__7)
	}

	return localctx
}

// ILocationOffsetByRegisterContext is an interface to support dynamic dispatch.
type ILocationOffsetByRegisterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocationOffsetByRegisterContext differentiates from other interfaces.
	IsLocationOffsetByRegisterContext()
}

type LocationOffsetByRegisterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocationOffsetByRegisterContext() *LocationOffsetByRegisterContext {
	var p = new(LocationOffsetByRegisterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_locationOffsetByRegister
	return p
}

func (*LocationOffsetByRegisterContext) IsLocationOffsetByRegisterContext() {}

func NewLocationOffsetByRegisterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocationOffsetByRegisterContext {
	var p = new(LocationOffsetByRegisterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_locationOffsetByRegister

	return p
}

func (s *LocationOffsetByRegisterContext) GetParser() antlr.Parser { return s.parser }

func (s *LocationOffsetByRegisterContext) REGISTER() antlr.TerminalNode {
	return s.GetToken(DCPUParserREGISTER, 0)
}

func (s *LocationOffsetByRegisterContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *LocationOffsetByRegisterContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *LocationOffsetByRegisterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocationOffsetByRegisterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocationOffsetByRegisterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterLocationOffsetByRegister(s)
	}
}

func (s *LocationOffsetByRegisterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitLocationOffsetByRegister(s)
	}
}

func (s *LocationOffsetByRegisterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitLocationOffsetByRegister(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) LocationOffsetByRegister() (localctx ILocationOffsetByRegisterContext) {
	localctx = NewLocationOffsetByRegisterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DCPUParserRULE_locationOffsetByRegister)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(DCPUParserT__6)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DCPUParserREGISTER:
		{
			p.SetState(127)
			p.Match(DCPUParserREGISTER)
		}
		{
			p.SetState(128)
			p.Match(DCPUParserT__8)
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case DCPUParserIDENTIFIER:
			{
				p.SetState(129)
				p.Label()
			}

		case DCPUParserNUMBER:
			{
				p.SetState(130)
				p.Value()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	case DCPUParserNUMBER, DCPUParserIDENTIFIER:
		p.SetState(135)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case DCPUParserIDENTIFIER:
			{
				p.SetState(133)
				p.Label()
			}

		case DCPUParserNUMBER:
			{
				p.SetState(134)
				p.Value()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(137)
			p.Match(DCPUParserT__8)
		}
		{
			p.SetState(138)
			p.Match(DCPUParserREGISTER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(142)
		p.Match(DCPUParserT__7)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(DCPUParserNUMBER, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, DCPUParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(DCPUParserNUMBER)
	}

	return localctx
}

// IUnaryOpcodeContext is an interface to support dynamic dispatch.
type IUnaryOpcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOpcodeContext differentiates from other interfaces.
	IsUnaryOpcodeContext()
}

type UnaryOpcodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOpcodeContext() *UnaryOpcodeContext {
	var p = new(UnaryOpcodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_unaryOpcode
	return p
}

func (*UnaryOpcodeContext) IsUnaryOpcodeContext() {}

func NewUnaryOpcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOpcodeContext {
	var p = new(UnaryOpcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_unaryOpcode

	return p
}

func (s *UnaryOpcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOpcodeContext) JSR() antlr.TerminalNode {
	return s.GetToken(DCPUParserJSR, 0)
}

func (s *UnaryOpcodeContext) INT() antlr.TerminalNode {
	return s.GetToken(DCPUParserINT, 0)
}

func (s *UnaryOpcodeContext) IAG() antlr.TerminalNode {
	return s.GetToken(DCPUParserIAG, 0)
}

func (s *UnaryOpcodeContext) IAS() antlr.TerminalNode {
	return s.GetToken(DCPUParserIAS, 0)
}

func (s *UnaryOpcodeContext) RFI() antlr.TerminalNode {
	return s.GetToken(DCPUParserRFI, 0)
}

func (s *UnaryOpcodeContext) IAQ() antlr.TerminalNode {
	return s.GetToken(DCPUParserIAQ, 0)
}

func (s *UnaryOpcodeContext) HWN() antlr.TerminalNode {
	return s.GetToken(DCPUParserHWN, 0)
}

func (s *UnaryOpcodeContext) HWQ() antlr.TerminalNode {
	return s.GetToken(DCPUParserHWQ, 0)
}

func (s *UnaryOpcodeContext) HWI() antlr.TerminalNode {
	return s.GetToken(DCPUParserHWI, 0)
}

func (s *UnaryOpcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOpcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOpcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterUnaryOpcode(s)
	}
}

func (s *UnaryOpcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitUnaryOpcode(s)
	}
}

func (s *UnaryOpcodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitUnaryOpcode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) UnaryOpcode() (localctx IUnaryOpcodeContext) {
	localctx = NewUnaryOpcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, DCPUParserRULE_unaryOpcode)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(146)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-46)&-(0x1f+1)) == 0 && ((1<<uint(_la-46))&((1<<(DCPUParserJSR-46))|(1<<(DCPUParserINT-46))|(1<<(DCPUParserIAG-46))|(1<<(DCPUParserIAS-46))|(1<<(DCPUParserRFI-46))|(1<<(DCPUParserIAQ-46))|(1<<(DCPUParserHWN-46))|(1<<(DCPUParserHWQ-46))|(1<<(DCPUParserHWI-46)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IDebugOperationContext is an interface to support dynamic dispatch.
type IDebugOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDebugOperationContext differentiates from other interfaces.
	IsDebugOperationContext()
}

type DebugOperationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDebugOperationContext() *DebugOperationContext {
	var p = new(DebugOperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_debugOperation
	return p
}

func (*DebugOperationContext) IsDebugOperationContext() {}

func NewDebugOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DebugOperationContext {
	var p = new(DebugOperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_debugOperation

	return p
}

func (s *DebugOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *DebugOperationContext) DebugOpcode() IDebugOpcodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDebugOpcodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDebugOpcodeContext)
}

func (s *DebugOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DebugOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DebugOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterDebugOperation(s)
	}
}

func (s *DebugOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitDebugOperation(s)
	}
}

func (s *DebugOperationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitDebugOperation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) DebugOperation() (localctx IDebugOperationContext) {
	localctx = NewDebugOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, DCPUParserRULE_debugOperation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.DebugOpcode()
	}

	return localctx
}

// IDebugOpcodeContext is an interface to support dynamic dispatch.
type IDebugOpcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDebugOpcodeContext differentiates from other interfaces.
	IsDebugOpcodeContext()
}

type DebugOpcodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDebugOpcodeContext() *DebugOpcodeContext {
	var p = new(DebugOpcodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCPUParserRULE_debugOpcode
	return p
}

func (*DebugOpcodeContext) IsDebugOpcodeContext() {}

func NewDebugOpcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DebugOpcodeContext {
	var p = new(DebugOpcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_debugOpcode

	return p
}

func (s *DebugOpcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *DebugOpcodeContext) ALT() antlr.TerminalNode {
	return s.GetToken(DCPUParserALT, 0)
}

func (s *DebugOpcodeContext) DUM() antlr.TerminalNode {
	return s.GetToken(DCPUParserDUM, 0)
}

func (s *DebugOpcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DebugOpcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DebugOpcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.EnterDebugOpcode(s)
	}
}

func (s *DebugOpcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCPUListener); ok {
		listenerT.ExitDebugOpcode(s)
	}
}

func (s *DebugOpcodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DCPUVisitor:
		return t.VisitDebugOpcode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DCPUParser) DebugOpcode() (localctx IDebugOpcodeContext) {
	localctx = NewDebugOpcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, DCPUParserRULE_debugOpcode)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(150)
	_la = p.GetTokenStream().LA(1)

	if !(_la == DCPUParserALT || _la == DCPUParserDUM) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
