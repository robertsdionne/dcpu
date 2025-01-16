module Main where

import Data.Bits

main :: IO [()]
main = mapM print [
        Special InterruptAddressSet (LeftValue (WithPayload Literal 11)),
        Basic Set (WithRegister Register RegisterB) (SmallLiteral 1),
        Special HardwareInterrupt (SmallLiteral 0),
        Basic Set (WithRegister Register RegisterA) (SmallLiteral 2),
        Basic Set (WithRegister Register RegisterB) (SmallLiteral 1),
        Special HardwareInterrupt (SmallLiteral 0),
        Basic Set (WithRegister Register RegisterA) (SmallLiteral 3),
        Basic Set (WithRegister Register RegisterB) (SmallLiteral 2),
        Special HardwareInterrupt (SmallLiteral 1),
        Basic Subtract ProgramCounter (SmallLiteral 1),
        Basic IfEqual (WithRegister Register RegisterA) (SmallLiteral 1),
        Special JumpSubroutine (LeftValue (WithPayload Literal 18)),
        Basic IfEqual (WithRegister Register  RegisterA) (SmallLiteral 2),
        Special JumpSubroutine (LeftValue (WithPayload Literal 38)),
        Special ReturnFromInterrupt (SmallLiteral 0),
        Basic Set (WithRegister Register RegisterA) (SmallLiteral 1),
        Special HardwareInterrupt (SmallLiteral 0),
        Basic Modulo (WithRegister Register RegisterC) (SmallLiteral 4),
        Basic Set (WithPayload Location 61440) (LeftValue (WithPayload (LocationOffsetByRegister RegisterC) 34)),
        Basic Set (WithPayload Location 61471) (LeftValue (WithPayload (LocationOffsetByRegister RegisterC) 34)),
        Basic Set (WithPayload Location 61792) (LeftValue (WithPayload (LocationOffsetByRegister RegisterC) 34)),
        Basic Set (WithPayload Location 61823) (LeftValue (WithPayload (LocationOffsetByRegister RegisterC) 34)),
        Basic Set ProgramCounter (LeftValue PushOrPop),
        Basic Set (WithRegister Register RegisterA) (SmallLiteral 1),
        Special HardwareInterrupt (SmallLiteral 1),
        Basic Add (WithPayload Location 61439) (SmallLiteral 1),
        Basic Set (WithRegister Register RegisterA) (LeftValue (WithPayload Location 61439)),
        Basic Modulo (WithRegister Register RegisterA) (LeftValue (WithPayload Literal 384))
    ]

class Sizeable a where
    size :: a -> Int

class From a b where
    from :: a -> b

class Into a b where
    into :: b -> a

newtype InstructionValue = InstructionValue Int

instance From Int InstructionValue where
    from value = InstructionValue value

instance Into Int InstructionValue where
    into (InstructionValue value) = value

data Instruction =
    Basic BasicOpcode OperandB OperandA
  | Special SpecialOpcode OperandA
  | Debug DebugOpcode
    deriving (Show)

instance From Int Instruction where
    from value =
        if basic_opcode /= Reserved then
            let operand_a :: OperandA = from instruction
                operand_b :: OperandB = from instruction
            in Basic basic_opcode operand_b operand_a
        else if special_opcode /= SpecialReserved then
            let operand_a :: OperandA = from instruction
            in Special special_opcode operand_a
        else
            Debug debug_opcode
        where instruction :: InstructionValue = from value
              basic_opcode :: BasicOpcode = from instruction
              special_opcode :: SpecialOpcode = from instruction
              debug_opcode :: DebugOpcode = from instruction

instance Into Int Instruction where
    into (Basic basic_opcode operand_b operand_a) =
        let basic_opcode_value :: InstructionValue = into basic_opcode
            operand_a_value :: InstructionValue = into operand_a
            operand_b_value :: InstructionValue = into operand_b
        in let basic_opcode :: Int = into basic_opcode_value
               operand_a :: Int = into operand_a_value
               operand_b :: Int = into operand_b_value
        in basic_opcode .|. operand_b .|. operand_a
    into (Special special_opcode operand_a) =
        let special_opcode_value :: InstructionValue = into special_opcode
            operand_a_value :: InstructionValue = into operand_a
        in let special_opcode :: Int = into special_opcode_value
               operand_a :: Int = into operand_a_value
        in special_opcode .|. operand_a
    into (Debug debug_opcode) =
        let debug_opcode_value :: InstructionValue = into debug_opcode
        in let debug_opcode :: Int = into debug_opcode_value
        in debug_opcode

instance Sizeable Instruction where
    size (Basic _ operand_b operand_a) = 1 + size operand_b + size operand_a
    size (Special _ operand_a) = 1 + size operand_a
    size _ = 1

data BasicOpcode =
    Reserved
  | Set
  | Add
  | Subtract
  | Multiply
  | MultiplySigned
  | Divide
  | DivideSigned
  | Modulo
  | ModuloSigned
  | BinaryAnd
  | BinaryOr
  | BinaryExclusiveOr
  | ShiftRight
  | ArithmeticShiftRight
  | ShiftLeft
  | IfBitSet
  | IfClear
  | IfEqual
  | IfNotEqual
  | IfGreaterThan
  | IfAbove
  | IfLessThan
  | IfUnder
  | Unused18
  | Unused19
  | AddWithCarry
  | SubtractWithCarry
  | Unused1c
  | Unused1d
  | SetThenIncrement
  | SetThenDecrement
    deriving (Enum, Eq, Show)

instance From InstructionValue BasicOpcode where
    from value = Reserved

instance Into InstructionValue BasicOpcode where
    into basic_opcode = InstructionValue (fromEnum basic_opcode)

data SpecialOpcode =
    SpecialReserved
  | JumpSubroutine
  | Unused02
  | Unused03
  | Unused04
  | Unused05
  | Unused06
  | Unused07
  | InterruptTrigger
  | InterruptAddressGet
  | InterruptAddressSet
  | ReturnFromInterrupt
  | InterruptAddedToQueue
  | Unused0d
  | Unused0e
  | HardwareNumberConnected
  | HardwareQuery
  | HardwareInterrupt
    deriving (Enum, Eq, Show)

instance From InstructionValue SpecialOpcode where
    from value = SpecialReserved

instance Into InstructionValue SpecialOpcode where
    into special_opcode = InstructionValue (fromEnum special_opcode)

data DebugOpcode =
    Noop
  | Alert
  | DumpState
    deriving (Enum, Eq, Show)

instance From InstructionValue DebugOpcode where
    from value = Alert

instance Into InstructionValue DebugOpcode where
    into debug_opcode = InstructionValue (fromEnum debug_opcode)

data OperandA =
    LeftValue OperandB
  | SmallLiteral Int
    deriving (Show)

instance From InstructionValue OperandA where
    from value = SmallLiteral 3

instance Into InstructionValue OperandA where
    into basic_opcode = InstructionValue 2

instance Sizeable OperandA where
    size (LeftValue operand_b) = size operand_b
    size _ = 0

data OperandB =
    WithRegister WithRegister Register
  | WithPayload WithPayload Int
  | PushOrPop
  | Peek
  | StackPointer
  | ProgramCounter
  | Extra
    deriving (Show)

instance From InstructionValue OperandB where
    from value = WithRegister Register RegisterA

instance Into InstructionValue OperandB where
    into basic_opcode = InstructionValue 1

instance Sizeable OperandB where
    size (WithPayload _ _) = 1
    size _ = 0

data WithRegister =
    Register
  | LocationInRegister
    deriving (Show)

data WithPayload =
    LocationOffsetByRegister Register
  | Pick
  | Location
  | Literal
    deriving (Show)

data Register =
    RegisterA
  | RegisterB
  | RegisterC
  | RegisterX
  | RegisterY
  | RegisterZ
  | RegisterI
  | RegisterJ
    deriving (Enum, Show)
