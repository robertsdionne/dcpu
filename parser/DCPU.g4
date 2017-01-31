grammar DCPU;

program
    : (labelDefinition | instruction | dataSection)+ EOF
    ;

labelDefinition
    : ':' IDENTIFIER
    ;

label : IDENTIFIER ;

instruction
    : binaryOperation
    | unaryOperation
    ;

dataSection
    : ('.dat' | '.DAT' | 'dat' | 'DAT') data
    ;

data
    : datum (',' datum)*
    ;

datum
    : STRING
    | IDENTIFIER
    | NUMBER
    ;

binaryOperation
    : binaryOpcode argumentB ',' argumentA
    ;

binaryOpcode
    : SET
    | ADD
    | SUB
    | MUL
    | MLI
    | DIV
    | DVI
    | MOD
    | MDI
    | AND
    | BOR
    | XOR
    | SHR
    | ASR
    | SHL
    | IFB
    | IFC
    | IFE
    | IFN
    | IFG
    | IFA
    | IFL
    | IFU
    | ADX
    | SBX
    | STI
    | STD
    ;

SET : 'SET' | 'set' ;
ADD : 'ADD' | 'add' ;
SUB : 'SUB' | 'sub' ;
MUL : 'MUL' | 'mul' ;
MLI : 'MLI' | 'mli' ;
DIV : 'DIV' | 'div' ;
DVI : 'DVI' | 'dvi' ;
MOD : 'MOD' | 'mod' ;
MDI : 'MDI' | 'mdi' ;
AND : 'AND' | 'and' ;
BOR : 'BOR' | 'bor' ;
XOR : 'XOR' | 'xor' ;
SHR : 'SHR' | 'shr' ;
ASR : 'ASR' | 'asr' ;
SHL : 'SHL' | 'shl' ;
IFB : 'IFB' | 'ifb' ;
IFC : 'IFC' | 'ifc' ;
IFE : 'IFE' | 'ife' ;
IFN : 'IFN' | 'ifn' ;
IFG : 'IFG' | 'ifg' ;
IFA : 'IFA' | 'ifa' ;
IFL : 'IFL' | 'ifl' ;
IFU : 'IFU' | 'ifu' ;
ADX : 'ADX' | 'adx' ;
SBX : 'SBX' | 'sbx' ;
STI : 'STI' | 'sti' ;
STD : 'STD' | 'std' ;

argumentA
    : register
    | locationInRegister
    | locationOffsetByRegister
    | POP
    | PEEK
    | pick
    | STACK_POINTER
    | PROGRAM_COUNTER
    | EXTRA
    | location
    | label
    | value
    ;

location
    : '[' (label | value) ']'
    ;

POP
    : 'POP'
    | 'pop'
    ;

argumentB
    : register
    | locationInRegister
    | locationOffsetByRegister
    | PUSH
    | PEEK
    | pick
    | STACK_POINTER
    | PROGRAM_COUNTER
    | EXTRA
    | location
    ;

register
    : REGISTER
    ;

PUSH
    : 'PUSH'
    | 'push'
    ;

PEEK
    : 'PEEK'
    | 'peek'
    ;

pick
    : PICK NUMBER
    ;

PICK
    : 'PICK'
    | 'pick'
    ;

STACK_POINTER
    : 'SP'
    | 'sp'
    ;

PROGRAM_COUNTER
    : 'PC'
    | 'pc'
    ;

EXTRA
    : 'EX'
    | 'ex'
    ;

unaryOperation
    : unaryOpcode argumentA
    ;

locationInRegister
    : '[' REGISTER ']'
    ;

locationOffsetByRegister
    : '[' ( REGISTER '+' (label | value) | (label | value) '+' REGISTER ) ']'
    ;

value
    : NUMBER
    ;

NUMBER
    : '0x' [0-9a-fA-F]+
    | '-'? [0-9]+
    ;

REGISTER
    : [abcxyzijABCXYZIJ]
    ;

unaryOpcode
    : JSR
    | INT
    | IAG
    | IAS
    | RFI
    | IAQ
    | HWN
    | HWQ
    | HWI
    ;

JSR : 'JSR' | 'jsr' ;
INT : 'INT' | 'int' ;
IAG : 'IAG' | 'iag' ;
IAS : 'IAS' | 'ias' ;
RFI : 'RFI' | 'rfi' ;
IAQ : 'IAQ' | 'iaq' ;
HWN : 'HWN' | 'hwn' ;
HWQ : 'HWQ' | 'hwq' ;
HWI : 'HWI' | 'hwi' ;

COMMENT
    : ';' ~[\r\n]* -> skip
    ;

IDENTIFIER
    : [._a-zA-Z]+
    ;

STRING
    : '"' (ESCAPE|.)*? '"'
    ;

fragment ESCAPE
    : '\\"' | '\\\\'
    ;

WHITESPACE
    : [ \t\r\n]+ -> skip
    ;
