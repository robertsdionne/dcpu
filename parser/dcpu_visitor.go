// Code generated from DCPU.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // DCPU
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by DCPUParser.
type DCPUVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DCPUParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by DCPUParser#labelDefinition.
	VisitLabelDefinition(ctx *LabelDefinitionContext) interface{}

	// Visit a parse tree produced by DCPUParser#label.
	VisitLabel(ctx *LabelContext) interface{}

	// Visit a parse tree produced by DCPUParser#instruction.
	VisitInstruction(ctx *InstructionContext) interface{}

	// Visit a parse tree produced by DCPUParser#dataSection.
	VisitDataSection(ctx *DataSectionContext) interface{}

	// Visit a parse tree produced by DCPUParser#data.
	VisitData(ctx *DataContext) interface{}

	// Visit a parse tree produced by DCPUParser#datum.
	VisitDatum(ctx *DatumContext) interface{}

	// Visit a parse tree produced by DCPUParser#binaryOperation.
	VisitBinaryOperation(ctx *BinaryOperationContext) interface{}

	// Visit a parse tree produced by DCPUParser#binaryOpcode.
	VisitBinaryOpcode(ctx *BinaryOpcodeContext) interface{}

	// Visit a parse tree produced by DCPUParser#argumentA.
	VisitArgumentA(ctx *ArgumentAContext) interface{}

	// Visit a parse tree produced by DCPUParser#location.
	VisitLocation(ctx *LocationContext) interface{}

	// Visit a parse tree produced by DCPUParser#argumentB.
	VisitArgumentB(ctx *ArgumentBContext) interface{}

	// Visit a parse tree produced by DCPUParser#register.
	VisitRegister(ctx *RegisterContext) interface{}

	// Visit a parse tree produced by DCPUParser#pick.
	VisitPick(ctx *PickContext) interface{}

	// Visit a parse tree produced by DCPUParser#unaryOperation.
	VisitUnaryOperation(ctx *UnaryOperationContext) interface{}

	// Visit a parse tree produced by DCPUParser#locationInRegister.
	VisitLocationInRegister(ctx *LocationInRegisterContext) interface{}

	// Visit a parse tree produced by DCPUParser#locationOffsetByRegister.
	VisitLocationOffsetByRegister(ctx *LocationOffsetByRegisterContext) interface{}

	// Visit a parse tree produced by DCPUParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by DCPUParser#unaryOpcode.
	VisitUnaryOpcode(ctx *UnaryOpcodeContext) interface{}

	// Visit a parse tree produced by DCPUParser#debugOperation.
	VisitDebugOperation(ctx *DebugOperationContext) interface{}

	// Visit a parse tree produced by DCPUParser#debugOpcode.
	VisitDebugOpcode(ctx *DebugOpcodeContext) interface{}
}
