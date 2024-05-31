// Code generated from parser/DCPU.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // DCPU
import "github.com/antlr4-go/antlr/v4"

type BaseDCPUVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDCPUVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitLabelDefinition(ctx *LabelDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitLabel(ctx *LabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitInstruction(ctx *InstructionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitDataSection(ctx *DataSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitData(ctx *DataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitDatum(ctx *DatumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitBinaryOperation(ctx *BinaryOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitBinaryOpcode(ctx *BinaryOpcodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitArgumentA(ctx *ArgumentAContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitLocation(ctx *LocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitArgumentB(ctx *ArgumentBContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitRegister(ctx *RegisterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitPick(ctx *PickContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitUnaryOperation(ctx *UnaryOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitLocationInRegister(ctx *LocationInRegisterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitLocationOffsetByRegister(ctx *LocationOffsetByRegisterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitUnaryOpcode(ctx *UnaryOpcodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitDebugOperation(ctx *DebugOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDCPUVisitor) VisitDebugOpcode(ctx *DebugOpcodeContext) interface{} {
	return v.VisitChildren(ctx)
}
