// Generated from DCPU.g4 by ANTLR 4.6.

package parser // DCPU
import "github.com/antlr/antlr4/runtime/Go/antlr"

// DCPUListener is a complete listener for a parse tree produced by DCPUParser.
type DCPUListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterLabelDefinition is called when entering the labelDefinition production.
	EnterLabelDefinition(c *LabelDefinitionContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterDataSection is called when entering the dataSection production.
	EnterDataSection(c *DataSectionContext)

	// EnterData is called when entering the data production.
	EnterData(c *DataContext)

	// EnterDatum is called when entering the datum production.
	EnterDatum(c *DatumContext)

	// EnterBinaryOperation is called when entering the binaryOperation production.
	EnterBinaryOperation(c *BinaryOperationContext)

	// EnterBinaryOpcode is called when entering the binaryOpcode production.
	EnterBinaryOpcode(c *BinaryOpcodeContext)

	// EnterArgumentA is called when entering the argumentA production.
	EnterArgumentA(c *ArgumentAContext)

	// EnterLocation is called when entering the location production.
	EnterLocation(c *LocationContext)

	// EnterArgumentB is called when entering the argumentB production.
	EnterArgumentB(c *ArgumentBContext)

	// EnterRegister is called when entering the register production.
	EnterRegister(c *RegisterContext)

	// EnterPick is called when entering the pick production.
	EnterPick(c *PickContext)

	// EnterUnaryOperation is called when entering the unaryOperation production.
	EnterUnaryOperation(c *UnaryOperationContext)

	// EnterLocationInRegister is called when entering the locationInRegister production.
	EnterLocationInRegister(c *LocationInRegisterContext)

	// EnterLocationOffsetByRegister is called when entering the locationOffsetByRegister production.
	EnterLocationOffsetByRegister(c *LocationOffsetByRegisterContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterUnaryOpcode is called when entering the unaryOpcode production.
	EnterUnaryOpcode(c *UnaryOpcodeContext)

	// EnterDebugOperation is called when entering the debugOperation production.
	EnterDebugOperation(c *DebugOperationContext)

	// EnterDebugOpcode is called when entering the debugOpcode production.
	EnterDebugOpcode(c *DebugOpcodeContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitLabelDefinition is called when exiting the labelDefinition production.
	ExitLabelDefinition(c *LabelDefinitionContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitDataSection is called when exiting the dataSection production.
	ExitDataSection(c *DataSectionContext)

	// ExitData is called when exiting the data production.
	ExitData(c *DataContext)

	// ExitDatum is called when exiting the datum production.
	ExitDatum(c *DatumContext)

	// ExitBinaryOperation is called when exiting the binaryOperation production.
	ExitBinaryOperation(c *BinaryOperationContext)

	// ExitBinaryOpcode is called when exiting the binaryOpcode production.
	ExitBinaryOpcode(c *BinaryOpcodeContext)

	// ExitArgumentA is called when exiting the argumentA production.
	ExitArgumentA(c *ArgumentAContext)

	// ExitLocation is called when exiting the location production.
	ExitLocation(c *LocationContext)

	// ExitArgumentB is called when exiting the argumentB production.
	ExitArgumentB(c *ArgumentBContext)

	// ExitRegister is called when exiting the register production.
	ExitRegister(c *RegisterContext)

	// ExitPick is called when exiting the pick production.
	ExitPick(c *PickContext)

	// ExitUnaryOperation is called when exiting the unaryOperation production.
	ExitUnaryOperation(c *UnaryOperationContext)

	// ExitLocationInRegister is called when exiting the locationInRegister production.
	ExitLocationInRegister(c *LocationInRegisterContext)

	// ExitLocationOffsetByRegister is called when exiting the locationOffsetByRegister production.
	ExitLocationOffsetByRegister(c *LocationOffsetByRegisterContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitUnaryOpcode is called when exiting the unaryOpcode production.
	ExitUnaryOpcode(c *UnaryOpcodeContext)

	// ExitDebugOperation is called when exiting the debugOperation production.
	ExitDebugOperation(c *DebugOperationContext)

	// ExitDebugOpcode is called when exiting the debugOpcode production.
	ExitDebugOpcode(c *DebugOpcodeContext)
}
