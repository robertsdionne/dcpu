// Code generated from parser/DCPU.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // DCPU
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type DCPUParser struct {
	*antlr.BaseParser
}

var DCPUParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dcpuParserInit() {
	staticData := &DCPUParserStaticData
	staticData.LiteralNames = []string{
		"", "':'", "'.dat'", "'.DAT'", "'dat'", "'DAT'", "','", "'['", "']'",
		"'+'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "SET", "ADD", "SUB", "MUL",
		"MLI", "DIV", "DVI", "MOD", "MDI", "AND", "BOR", "XOR", "SHR", "ASR",
		"SHL", "IFB", "IFC", "IFE", "IFN", "IFG", "IFA", "IFL", "IFU", "ADX",
		"SBX", "STI", "STD", "POP", "PUSH", "PEEK", "PICK", "STACK_POINTER",
		"PROGRAM_COUNTER", "EXTRA", "NUMBER", "REGISTER", "JSR", "INT", "IAG",
		"IAS", "RFI", "IAQ", "HWN", "HWQ", "HWI", "ALT", "DUM", "COMMENT", "IDENTIFIER",
		"STRING", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"program", "labelDefinition", "label", "instruction", "dataSection",
		"data", "datum", "binaryOperation", "binaryOpcode", "argumentA", "location",
		"argumentB", "register", "pick", "unaryOperation", "locationInRegister",
		"locationOffsetByRegister", "value", "unaryOpcode", "debugOperation",
		"debugOpcode",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 60, 153, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1,
		0, 1, 0, 1, 0, 4, 0, 46, 8, 0, 11, 0, 12, 0, 47, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3, 60, 8, 3, 1, 4, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 5, 5, 68, 8, 5, 10, 5, 12, 5, 71, 9, 5, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 94, 8, 9, 1, 10, 1, 10, 1, 10,
		3, 10, 99, 8, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 113, 8, 11, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 3, 16, 132, 8, 16, 1, 16, 1, 16, 3, 16, 136, 8, 16,
		1, 16, 1, 16, 1, 16, 3, 16, 141, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 0, 0, 21, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 0, 5, 1,
		0, 2, 5, 2, 0, 44, 44, 58, 59, 1, 0, 10, 36, 1, 0, 46, 54, 1, 0, 55, 56,
		161, 0, 45, 1, 0, 0, 0, 2, 51, 1, 0, 0, 0, 4, 54, 1, 0, 0, 0, 6, 59, 1,
		0, 0, 0, 8, 61, 1, 0, 0, 0, 10, 64, 1, 0, 0, 0, 12, 72, 1, 0, 0, 0, 14,
		74, 1, 0, 0, 0, 16, 79, 1, 0, 0, 0, 18, 93, 1, 0, 0, 0, 20, 95, 1, 0, 0,
		0, 22, 112, 1, 0, 0, 0, 24, 114, 1, 0, 0, 0, 26, 116, 1, 0, 0, 0, 28, 119,
		1, 0, 0, 0, 30, 122, 1, 0, 0, 0, 32, 126, 1, 0, 0, 0, 34, 144, 1, 0, 0,
		0, 36, 146, 1, 0, 0, 0, 38, 148, 1, 0, 0, 0, 40, 150, 1, 0, 0, 0, 42, 46,
		3, 2, 1, 0, 43, 46, 3, 6, 3, 0, 44, 46, 3, 8, 4, 0, 45, 42, 1, 0, 0, 0,
		45, 43, 1, 0, 0, 0, 45, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 45, 1,
		0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 50, 5, 0, 0, 1, 50,
		1, 1, 0, 0, 0, 51, 52, 5, 1, 0, 0, 52, 53, 5, 58, 0, 0, 53, 3, 1, 0, 0,
		0, 54, 55, 5, 58, 0, 0, 55, 5, 1, 0, 0, 0, 56, 60, 3, 14, 7, 0, 57, 60,
		3, 28, 14, 0, 58, 60, 3, 38, 19, 0, 59, 56, 1, 0, 0, 0, 59, 57, 1, 0, 0,
		0, 59, 58, 1, 0, 0, 0, 60, 7, 1, 0, 0, 0, 61, 62, 7, 0, 0, 0, 62, 63, 3,
		10, 5, 0, 63, 9, 1, 0, 0, 0, 64, 69, 3, 12, 6, 0, 65, 66, 5, 6, 0, 0, 66,
		68, 3, 12, 6, 0, 67, 65, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0,
		0, 0, 69, 70, 1, 0, 0, 0, 70, 11, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 73,
		7, 1, 0, 0, 73, 13, 1, 0, 0, 0, 74, 75, 3, 16, 8, 0, 75, 76, 3, 22, 11,
		0, 76, 77, 5, 6, 0, 0, 77, 78, 3, 18, 9, 0, 78, 15, 1, 0, 0, 0, 79, 80,
		7, 2, 0, 0, 80, 17, 1, 0, 0, 0, 81, 94, 3, 24, 12, 0, 82, 94, 3, 30, 15,
		0, 83, 94, 3, 32, 16, 0, 84, 94, 5, 37, 0, 0, 85, 94, 5, 39, 0, 0, 86,
		94, 3, 26, 13, 0, 87, 94, 5, 41, 0, 0, 88, 94, 5, 42, 0, 0, 89, 94, 5,
		43, 0, 0, 90, 94, 3, 20, 10, 0, 91, 94, 3, 4, 2, 0, 92, 94, 3, 34, 17,
		0, 93, 81, 1, 0, 0, 0, 93, 82, 1, 0, 0, 0, 93, 83, 1, 0, 0, 0, 93, 84,
		1, 0, 0, 0, 93, 85, 1, 0, 0, 0, 93, 86, 1, 0, 0, 0, 93, 87, 1, 0, 0, 0,
		93, 88, 1, 0, 0, 0, 93, 89, 1, 0, 0, 0, 93, 90, 1, 0, 0, 0, 93, 91, 1,
		0, 0, 0, 93, 92, 1, 0, 0, 0, 94, 19, 1, 0, 0, 0, 95, 98, 5, 7, 0, 0, 96,
		99, 3, 4, 2, 0, 97, 99, 3, 34, 17, 0, 98, 96, 1, 0, 0, 0, 98, 97, 1, 0,
		0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 5, 8, 0, 0, 101, 21, 1, 0, 0, 0, 102,
		113, 3, 24, 12, 0, 103, 113, 3, 30, 15, 0, 104, 113, 3, 32, 16, 0, 105,
		113, 5, 38, 0, 0, 106, 113, 5, 39, 0, 0, 107, 113, 3, 26, 13, 0, 108, 113,
		5, 41, 0, 0, 109, 113, 5, 42, 0, 0, 110, 113, 5, 43, 0, 0, 111, 113, 3,
		20, 10, 0, 112, 102, 1, 0, 0, 0, 112, 103, 1, 0, 0, 0, 112, 104, 1, 0,
		0, 0, 112, 105, 1, 0, 0, 0, 112, 106, 1, 0, 0, 0, 112, 107, 1, 0, 0, 0,
		112, 108, 1, 0, 0, 0, 112, 109, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112,
		111, 1, 0, 0, 0, 113, 23, 1, 0, 0, 0, 114, 115, 5, 45, 0, 0, 115, 25, 1,
		0, 0, 0, 116, 117, 5, 40, 0, 0, 117, 118, 5, 44, 0, 0, 118, 27, 1, 0, 0,
		0, 119, 120, 3, 36, 18, 0, 120, 121, 3, 18, 9, 0, 121, 29, 1, 0, 0, 0,
		122, 123, 5, 7, 0, 0, 123, 124, 5, 45, 0, 0, 124, 125, 5, 8, 0, 0, 125,
		31, 1, 0, 0, 0, 126, 140, 5, 7, 0, 0, 127, 128, 5, 45, 0, 0, 128, 131,
		5, 9, 0, 0, 129, 132, 3, 4, 2, 0, 130, 132, 3, 34, 17, 0, 131, 129, 1,
		0, 0, 0, 131, 130, 1, 0, 0, 0, 132, 141, 1, 0, 0, 0, 133, 136, 3, 4, 2,
		0, 134, 136, 3, 34, 17, 0, 135, 133, 1, 0, 0, 0, 135, 134, 1, 0, 0, 0,
		136, 137, 1, 0, 0, 0, 137, 138, 5, 9, 0, 0, 138, 139, 5, 45, 0, 0, 139,
		141, 1, 0, 0, 0, 140, 127, 1, 0, 0, 0, 140, 135, 1, 0, 0, 0, 141, 142,
		1, 0, 0, 0, 142, 143, 5, 8, 0, 0, 143, 33, 1, 0, 0, 0, 144, 145, 5, 44,
		0, 0, 145, 35, 1, 0, 0, 0, 146, 147, 7, 3, 0, 0, 147, 37, 1, 0, 0, 0, 148,
		149, 3, 40, 20, 0, 149, 39, 1, 0, 0, 0, 150, 151, 7, 4, 0, 0, 151, 41,
		1, 0, 0, 0, 10, 45, 47, 59, 69, 93, 98, 112, 131, 135, 140,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// DCPUParserInit initializes any static state used to implement DCPUParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDCPUParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DCPUParserInit() {
	staticData := &DCPUParserStaticData
	staticData.once.Do(dcpuParserInit)
}

// NewDCPUParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDCPUParser(input antlr.TokenStream) *DCPUParser {
	DCPUParserInit()
	this := new(DCPUParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &DCPUParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
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

	// Getter signatures
	EOF() antlr.TerminalNode
	AllLabelDefinition() []ILabelDefinitionContext
	LabelDefinition(i int) ILabelDefinitionContext
	AllInstruction() []IInstructionContext
	Instruction(i int) IInstructionContext
	AllDataSection() []IDataSectionContext
	DataSection(i int) IDataSectionContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(DCPUParserEOF, 0)
}

func (s *ProgramContext) AllLabelDefinition() []ILabelDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILabelDefinitionContext); ok {
			len++
		}
	}

	tst := make([]ILabelDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILabelDefinitionContext); ok {
			tst[i] = t.(ILabelDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) LabelDefinition(i int) ILabelDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelDefinitionContext)
}

func (s *ProgramContext) AllInstruction() []IInstructionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstructionContext); ok {
			len++
		}
	}

	tst := make([]IInstructionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstructionContext); ok {
			tst[i] = t.(IInstructionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Instruction(i int) IInstructionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *ProgramContext) AllDataSection() []IDataSectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDataSectionContext); ok {
			len++
		}
	}

	tst := make([]IDataSectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDataSectionContext); ok {
			tst[i] = t.(IDataSectionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) DataSection(i int) IDataSectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataSectionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144044956770630718) != 0) {
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

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
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(49)
		p.Match(DCPUParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabelDefinitionContext is an interface to support dynamic dispatch.
type ILabelDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsLabelDefinitionContext differentiates from other interfaces.
	IsLabelDefinitionContext()
}

type LabelDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelDefinitionContext() *LabelDefinitionContext {
	var p = new(LabelDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_labelDefinition
	return p
}

func InitEmptyLabelDefinitionContext(p *LabelDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_labelDefinition
}

func (*LabelDefinitionContext) IsLabelDefinitionContext() {}

func NewLabelDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelDefinitionContext {
	var p = new(LabelDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(DCPUParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(DCPUParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_label
	return p
}

func InitEmptyLabelContext(p *LabelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_label
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(DCPUParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BinaryOperation() IBinaryOperationContext
	UnaryOperation() IUnaryOperationContext
	DebugOperation() IDebugOperationContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) BinaryOperation() IBinaryOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryOperationContext)
}

func (s *InstructionContext) UnaryOperation() IUnaryOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryOperationContext)
}

func (s *InstructionContext) DebugOperation() IDebugOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDebugOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

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
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDataSectionContext is an interface to support dynamic dispatch.
type IDataSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Data() IDataContext

	// IsDataSectionContext differentiates from other interfaces.
	IsDataSectionContext()
}

type DataSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataSectionContext() *DataSectionContext {
	var p = new(DataSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_dataSection
	return p
}

func InitEmptyDataSectionContext(p *DataSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_dataSection
}

func (*DataSectionContext) IsDataSectionContext() {}

func NewDataSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataSectionContext {
	var p = new(DataSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_dataSection

	return p
}

func (s *DataSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *DataSectionContext) Data() IDataContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&60) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(62)
		p.Data()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDataContext is an interface to support dynamic dispatch.
type IDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDatum() []IDatumContext
	Datum(i int) IDatumContext

	// IsDataContext differentiates from other interfaces.
	IsDataContext()
}

type DataContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataContext() *DataContext {
	var p = new(DataContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_data
	return p
}

func InitEmptyDataContext(p *DataContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_data
}

func (*DataContext) IsDataContext() {}

func NewDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataContext {
	var p = new(DataContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_data

	return p
}

func (s *DataContext) GetParser() antlr.Parser { return s.parser }

func (s *DataContext) AllDatum() []IDatumContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDatumContext); ok {
			len++
		}
	}

	tst := make([]IDatumContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDatumContext); ok {
			tst[i] = t.(IDatumContext)
			i++
		}
	}

	return tst
}

func (s *DataContext) Datum(i int) IDatumContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatumContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Datum()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DCPUParserT__5 {
		{
			p.SetState(65)
			p.Match(DCPUParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Datum()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDatumContext is an interface to support dynamic dispatch.
type IDatumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsDatumContext differentiates from other interfaces.
	IsDatumContext()
}

type DatumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatumContext() *DatumContext {
	var p = new(DatumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_datum
	return p
}

func InitEmptyDatumContext(p *DatumContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_datum
}

func (*DatumContext) IsDatumContext() {}

func NewDatumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatumContext {
	var p = new(DatumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&864708720641179648) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinaryOperationContext is an interface to support dynamic dispatch.
type IBinaryOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BinaryOpcode() IBinaryOpcodeContext
	ArgumentB() IArgumentBContext
	ArgumentA() IArgumentAContext

	// IsBinaryOperationContext differentiates from other interfaces.
	IsBinaryOperationContext()
}

type BinaryOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOperationContext() *BinaryOperationContext {
	var p = new(BinaryOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_binaryOperation
	return p
}

func InitEmptyBinaryOperationContext(p *BinaryOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_binaryOperation
}

func (*BinaryOperationContext) IsBinaryOperationContext() {}

func NewBinaryOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOperationContext {
	var p = new(BinaryOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_binaryOperation

	return p
}

func (s *BinaryOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOperationContext) BinaryOpcode() IBinaryOpcodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryOpcodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryOpcodeContext)
}

func (s *BinaryOperationContext) ArgumentB() IArgumentBContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentBContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentBContext)
}

func (s *BinaryOperationContext) ArgumentA() IArgumentAContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentAContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)
		p.ArgumentA()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinaryOpcodeContext is an interface to support dynamic dispatch.
type IBinaryOpcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SET() antlr.TerminalNode
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode
	MUL() antlr.TerminalNode
	MLI() antlr.TerminalNode
	DIV() antlr.TerminalNode
	DVI() antlr.TerminalNode
	MOD() antlr.TerminalNode
	MDI() antlr.TerminalNode
	AND() antlr.TerminalNode
	BOR() antlr.TerminalNode
	XOR() antlr.TerminalNode
	SHR() antlr.TerminalNode
	ASR() antlr.TerminalNode
	SHL() antlr.TerminalNode
	IFB() antlr.TerminalNode
	IFC() antlr.TerminalNode
	IFE() antlr.TerminalNode
	IFN() antlr.TerminalNode
	IFG() antlr.TerminalNode
	IFA() antlr.TerminalNode
	IFL() antlr.TerminalNode
	IFU() antlr.TerminalNode
	ADX() antlr.TerminalNode
	SBX() antlr.TerminalNode
	STI() antlr.TerminalNode
	STD() antlr.TerminalNode

	// IsBinaryOpcodeContext differentiates from other interfaces.
	IsBinaryOpcodeContext()
}

type BinaryOpcodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOpcodeContext() *BinaryOpcodeContext {
	var p = new(BinaryOpcodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_binaryOpcode
	return p
}

func InitEmptyBinaryOpcodeContext(p *BinaryOpcodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_binaryOpcode
}

func (*BinaryOpcodeContext) IsBinaryOpcodeContext() {}

func NewBinaryOpcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOpcodeContext {
	var p = new(BinaryOpcodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&137438952448) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentAContext is an interface to support dynamic dispatch.
type IArgumentAContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Register() IRegisterContext
	LocationInRegister() ILocationInRegisterContext
	LocationOffsetByRegister() ILocationOffsetByRegisterContext
	POP() antlr.TerminalNode
	PEEK() antlr.TerminalNode
	Pick() IPickContext
	STACK_POINTER() antlr.TerminalNode
	PROGRAM_COUNTER() antlr.TerminalNode
	EXTRA() antlr.TerminalNode
	Location() ILocationContext
	Label() ILabelContext
	Value() IValueContext

	// IsArgumentAContext differentiates from other interfaces.
	IsArgumentAContext()
}

type ArgumentAContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentAContext() *ArgumentAContext {
	var p = new(ArgumentAContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_argumentA
	return p
}

func InitEmptyArgumentAContext(p *ArgumentAContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_argumentA
}

func (*ArgumentAContext) IsArgumentAContext() {}

func NewArgumentAContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentAContext {
	var p = new(ArgumentAContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_argumentA

	return p
}

func (s *ArgumentAContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentAContext) Register() IRegisterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegisterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *ArgumentAContext) LocationInRegister() ILocationInRegisterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocationInRegisterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocationInRegisterContext)
}

func (s *ArgumentAContext) LocationOffsetByRegister() ILocationOffsetByRegisterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocationOffsetByRegisterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPickContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocationContext)
}

func (s *ArgumentAContext) Label() ILabelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *ArgumentAContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
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
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(85)
			p.Match(DCPUParserPEEK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(88)
			p.Match(DCPUParserPROGRAM_COUNTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(89)
			p.Match(DCPUParserEXTRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILocationContext is an interface to support dynamic dispatch.
type ILocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Label() ILabelContext
	Value() IValueContext

	// IsLocationContext differentiates from other interfaces.
	IsLocationContext()
}

type LocationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocationContext() *LocationContext {
	var p = new(LocationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_location
	return p
}

func InitEmptyLocationContext(p *LocationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_location
}

func (*LocationContext) IsLocationContext() {}

func NewLocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocationContext {
	var p = new(LocationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_location

	return p
}

func (s *LocationContext) GetParser() antlr.Parser { return s.parser }

func (s *LocationContext) Label() ILabelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *LocationContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(DCPUParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

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
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(100)
		p.Match(DCPUParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentBContext is an interface to support dynamic dispatch.
type IArgumentBContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Register() IRegisterContext
	LocationInRegister() ILocationInRegisterContext
	LocationOffsetByRegister() ILocationOffsetByRegisterContext
	PUSH() antlr.TerminalNode
	PEEK() antlr.TerminalNode
	Pick() IPickContext
	STACK_POINTER() antlr.TerminalNode
	PROGRAM_COUNTER() antlr.TerminalNode
	EXTRA() antlr.TerminalNode
	Location() ILocationContext

	// IsArgumentBContext differentiates from other interfaces.
	IsArgumentBContext()
}

type ArgumentBContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentBContext() *ArgumentBContext {
	var p = new(ArgumentBContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_argumentB
	return p
}

func InitEmptyArgumentBContext(p *ArgumentBContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_argumentB
}

func (*ArgumentBContext) IsArgumentBContext() {}

func NewArgumentBContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentBContext {
	var p = new(ArgumentBContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_argumentB

	return p
}

func (s *ArgumentBContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentBContext) Register() IRegisterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegisterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *ArgumentBContext) LocationInRegister() ILocationInRegisterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocationInRegisterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocationInRegisterContext)
}

func (s *ArgumentBContext) LocationOffsetByRegister() ILocationOffsetByRegisterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocationOffsetByRegisterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPickContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
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
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(106)
			p.Match(DCPUParserPEEK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(109)
			p.Match(DCPUParserPROGRAM_COUNTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(110)
			p.Match(DCPUParserEXTRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(111)
			p.Location()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRegisterContext is an interface to support dynamic dispatch.
type IRegisterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REGISTER() antlr.TerminalNode

	// IsRegisterContext differentiates from other interfaces.
	IsRegisterContext()
}

type RegisterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegisterContext() *RegisterContext {
	var p = new(RegisterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_register
	return p
}

func InitEmptyRegisterContext(p *RegisterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_register
}

func (*RegisterContext) IsRegisterContext() {}

func NewRegisterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegisterContext {
	var p = new(RegisterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(DCPUParserREGISTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPickContext is an interface to support dynamic dispatch.
type IPickContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PICK() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsPickContext differentiates from other interfaces.
	IsPickContext()
}

type PickContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPickContext() *PickContext {
	var p = new(PickContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_pick
	return p
}

func InitEmptyPickContext(p *PickContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_pick
}

func (*PickContext) IsPickContext() {}

func NewPickContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PickContext {
	var p = new(PickContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(DCPUParserPICK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Match(DCPUParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryOperationContext is an interface to support dynamic dispatch.
type IUnaryOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryOpcode() IUnaryOpcodeContext
	ArgumentA() IArgumentAContext

	// IsUnaryOperationContext differentiates from other interfaces.
	IsUnaryOperationContext()
}

type UnaryOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperationContext() *UnaryOperationContext {
	var p = new(UnaryOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_unaryOperation
	return p
}

func InitEmptyUnaryOperationContext(p *UnaryOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_unaryOperation
}

func (*UnaryOperationContext) IsUnaryOperationContext() {}

func NewUnaryOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperationContext {
	var p = new(UnaryOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_unaryOperation

	return p
}

func (s *UnaryOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOperationContext) UnaryOpcode() IUnaryOpcodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryOpcodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryOpcodeContext)
}

func (s *UnaryOperationContext) ArgumentA() IArgumentAContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentAContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.UnaryOpcode()
	}
	{
		p.SetState(120)
		p.ArgumentA()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILocationInRegisterContext is an interface to support dynamic dispatch.
type ILocationInRegisterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REGISTER() antlr.TerminalNode

	// IsLocationInRegisterContext differentiates from other interfaces.
	IsLocationInRegisterContext()
}

type LocationInRegisterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocationInRegisterContext() *LocationInRegisterContext {
	var p = new(LocationInRegisterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_locationInRegister
	return p
}

func InitEmptyLocationInRegisterContext(p *LocationInRegisterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_locationInRegister
}

func (*LocationInRegisterContext) IsLocationInRegisterContext() {}

func NewLocationInRegisterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocationInRegisterContext {
	var p = new(LocationInRegisterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(DCPUParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(DCPUParserREGISTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(DCPUParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILocationOffsetByRegisterContext is an interface to support dynamic dispatch.
type ILocationOffsetByRegisterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REGISTER() antlr.TerminalNode
	Label() ILabelContext
	Value() IValueContext

	// IsLocationOffsetByRegisterContext differentiates from other interfaces.
	IsLocationOffsetByRegisterContext()
}

type LocationOffsetByRegisterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocationOffsetByRegisterContext() *LocationOffsetByRegisterContext {
	var p = new(LocationOffsetByRegisterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_locationOffsetByRegister
	return p
}

func InitEmptyLocationOffsetByRegisterContext(p *LocationOffsetByRegisterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_locationOffsetByRegister
}

func (*LocationOffsetByRegisterContext) IsLocationOffsetByRegisterContext() {}

func NewLocationOffsetByRegisterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocationOffsetByRegisterContext {
	var p = new(LocationOffsetByRegisterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_locationOffsetByRegister

	return p
}

func (s *LocationOffsetByRegisterContext) GetParser() antlr.Parser { return s.parser }

func (s *LocationOffsetByRegisterContext) REGISTER() antlr.TerminalNode {
	return s.GetToken(DCPUParserREGISTER, 0)
}

func (s *LocationOffsetByRegisterContext) Label() ILabelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *LocationOffsetByRegisterContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(DCPUParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DCPUParserREGISTER:
		{
			p.SetState(127)
			p.Match(DCPUParserREGISTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Match(DCPUParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

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
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	case DCPUParserNUMBER, DCPUParserIDENTIFIER:
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

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
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(137)
			p.Match(DCPUParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(DCPUParserREGISTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(142)
		p.Match(DCPUParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(DCPUParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryOpcodeContext is an interface to support dynamic dispatch.
type IUnaryOpcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JSR() antlr.TerminalNode
	INT() antlr.TerminalNode
	IAG() antlr.TerminalNode
	IAS() antlr.TerminalNode
	RFI() antlr.TerminalNode
	IAQ() antlr.TerminalNode
	HWN() antlr.TerminalNode
	HWQ() antlr.TerminalNode
	HWI() antlr.TerminalNode

	// IsUnaryOpcodeContext differentiates from other interfaces.
	IsUnaryOpcodeContext()
}

type UnaryOpcodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOpcodeContext() *UnaryOpcodeContext {
	var p = new(UnaryOpcodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_unaryOpcode
	return p
}

func InitEmptyUnaryOpcodeContext(p *UnaryOpcodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_unaryOpcode
}

func (*UnaryOpcodeContext) IsUnaryOpcodeContext() {}

func NewUnaryOpcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOpcodeContext {
	var p = new(UnaryOpcodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&35958428274786304) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDebugOperationContext is an interface to support dynamic dispatch.
type IDebugOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DebugOpcode() IDebugOpcodeContext

	// IsDebugOperationContext differentiates from other interfaces.
	IsDebugOperationContext()
}

type DebugOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDebugOperationContext() *DebugOperationContext {
	var p = new(DebugOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_debugOperation
	return p
}

func InitEmptyDebugOperationContext(p *DebugOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_debugOperation
}

func (*DebugOperationContext) IsDebugOperationContext() {}

func NewDebugOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DebugOperationContext {
	var p = new(DebugOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCPUParserRULE_debugOperation

	return p
}

func (s *DebugOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *DebugOperationContext) DebugOpcode() IDebugOpcodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDebugOpcodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.DebugOpcode()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDebugOpcodeContext is an interface to support dynamic dispatch.
type IDebugOpcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ALT() antlr.TerminalNode
	DUM() antlr.TerminalNode

	// IsDebugOpcodeContext differentiates from other interfaces.
	IsDebugOpcodeContext()
}

type DebugOpcodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDebugOpcodeContext() *DebugOpcodeContext {
	var p = new(DebugOpcodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_debugOpcode
	return p
}

func InitEmptyDebugOpcodeContext(p *DebugOpcodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCPUParserRULE_debugOpcode
}

func (*DebugOpcodeContext) IsDebugOpcodeContext() {}

func NewDebugOpcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DebugOpcodeContext {
	var p = new(DebugOpcodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DCPUParserALT || _la == DCPUParserDUM) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
