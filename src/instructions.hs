module Main where

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

data Instruction =
    Basic BasicOpcode OperandB OperandA
  | Special SpecialOpcode OperandA
  | Debug DebugOpcode
    deriving (Show)

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
    deriving (Enum, Show)

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
    deriving (Enum, Show)

data DebugOpcode =
    Noop
  | Alert
  | DumpState
    deriving (Enum, Show)

data OperandA =
    LeftValue OperandB
  | SmallLiteral Int
    deriving (Show)

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
