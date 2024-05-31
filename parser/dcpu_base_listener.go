// Code generated from parser/DCPU.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // DCPU
import "github.com/antlr4-go/antlr/v4"

// BaseDCPUListener is a complete listener for a parse tree produced by DCPUParser.
type BaseDCPUListener struct{}

var _ DCPUListener = &BaseDCPUListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDCPUListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDCPUListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDCPUListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDCPUListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseDCPUListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseDCPUListener) ExitProgram(ctx *ProgramContext) {}

// EnterLabelDefinition is called when production labelDefinition is entered.
func (s *BaseDCPUListener) EnterLabelDefinition(ctx *LabelDefinitionContext) {}

// ExitLabelDefinition is called when production labelDefinition is exited.
func (s *BaseDCPUListener) ExitLabelDefinition(ctx *LabelDefinitionContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseDCPUListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseDCPUListener) ExitLabel(ctx *LabelContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseDCPUListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseDCPUListener) ExitInstruction(ctx *InstructionContext) {}

// EnterDataSection is called when production dataSection is entered.
func (s *BaseDCPUListener) EnterDataSection(ctx *DataSectionContext) {}

// ExitDataSection is called when production dataSection is exited.
func (s *BaseDCPUListener) ExitDataSection(ctx *DataSectionContext) {}

// EnterData is called when production data is entered.
func (s *BaseDCPUListener) EnterData(ctx *DataContext) {}

// ExitData is called when production data is exited.
func (s *BaseDCPUListener) ExitData(ctx *DataContext) {}

// EnterDatum is called when production datum is entered.
func (s *BaseDCPUListener) EnterDatum(ctx *DatumContext) {}

// ExitDatum is called when production datum is exited.
func (s *BaseDCPUListener) ExitDatum(ctx *DatumContext) {}

// EnterBinaryOperation is called when production binaryOperation is entered.
func (s *BaseDCPUListener) EnterBinaryOperation(ctx *BinaryOperationContext) {}

// ExitBinaryOperation is called when production binaryOperation is exited.
func (s *BaseDCPUListener) ExitBinaryOperation(ctx *BinaryOperationContext) {}

// EnterBinaryOpcode is called when production binaryOpcode is entered.
func (s *BaseDCPUListener) EnterBinaryOpcode(ctx *BinaryOpcodeContext) {}

// ExitBinaryOpcode is called when production binaryOpcode is exited.
func (s *BaseDCPUListener) ExitBinaryOpcode(ctx *BinaryOpcodeContext) {}

// EnterArgumentA is called when production argumentA is entered.
func (s *BaseDCPUListener) EnterArgumentA(ctx *ArgumentAContext) {}

// ExitArgumentA is called when production argumentA is exited.
func (s *BaseDCPUListener) ExitArgumentA(ctx *ArgumentAContext) {}

// EnterLocation is called when production location is entered.
func (s *BaseDCPUListener) EnterLocation(ctx *LocationContext) {}

// ExitLocation is called when production location is exited.
func (s *BaseDCPUListener) ExitLocation(ctx *LocationContext) {}

// EnterArgumentB is called when production argumentB is entered.
func (s *BaseDCPUListener) EnterArgumentB(ctx *ArgumentBContext) {}

// ExitArgumentB is called when production argumentB is exited.
func (s *BaseDCPUListener) ExitArgumentB(ctx *ArgumentBContext) {}

// EnterRegister is called when production register is entered.
func (s *BaseDCPUListener) EnterRegister(ctx *RegisterContext) {}

// ExitRegister is called when production register is exited.
func (s *BaseDCPUListener) ExitRegister(ctx *RegisterContext) {}

// EnterPick is called when production pick is entered.
func (s *BaseDCPUListener) EnterPick(ctx *PickContext) {}

// ExitPick is called when production pick is exited.
func (s *BaseDCPUListener) ExitPick(ctx *PickContext) {}

// EnterUnaryOperation is called when production unaryOperation is entered.
func (s *BaseDCPUListener) EnterUnaryOperation(ctx *UnaryOperationContext) {}

// ExitUnaryOperation is called when production unaryOperation is exited.
func (s *BaseDCPUListener) ExitUnaryOperation(ctx *UnaryOperationContext) {}

// EnterLocationInRegister is called when production locationInRegister is entered.
func (s *BaseDCPUListener) EnterLocationInRegister(ctx *LocationInRegisterContext) {}

// ExitLocationInRegister is called when production locationInRegister is exited.
func (s *BaseDCPUListener) ExitLocationInRegister(ctx *LocationInRegisterContext) {}

// EnterLocationOffsetByRegister is called when production locationOffsetByRegister is entered.
func (s *BaseDCPUListener) EnterLocationOffsetByRegister(ctx *LocationOffsetByRegisterContext) {}

// ExitLocationOffsetByRegister is called when production locationOffsetByRegister is exited.
func (s *BaseDCPUListener) ExitLocationOffsetByRegister(ctx *LocationOffsetByRegisterContext) {}

// EnterValue is called when production value is entered.
func (s *BaseDCPUListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseDCPUListener) ExitValue(ctx *ValueContext) {}

// EnterUnaryOpcode is called when production unaryOpcode is entered.
func (s *BaseDCPUListener) EnterUnaryOpcode(ctx *UnaryOpcodeContext) {}

// ExitUnaryOpcode is called when production unaryOpcode is exited.
func (s *BaseDCPUListener) ExitUnaryOpcode(ctx *UnaryOpcodeContext) {}

// EnterDebugOperation is called when production debugOperation is entered.
func (s *BaseDCPUListener) EnterDebugOperation(ctx *DebugOperationContext) {}

// ExitDebugOperation is called when production debugOperation is exited.
func (s *BaseDCPUListener) ExitDebugOperation(ctx *DebugOperationContext) {}

// EnterDebugOpcode is called when production debugOpcode is entered.
func (s *BaseDCPUListener) EnterDebugOpcode(ctx *DebugOpcodeContext) {}

// ExitDebugOpcode is called when production debugOpcode is exited.
func (s *BaseDCPUListener) ExitDebugOpcode(ctx *DebugOpcodeContext) {}
