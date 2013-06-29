grammar Dcpu;

program
    : (label | instruction | data_section)+ EOF
    ;

label
    : ':' IDENTIFIER
    ;

instruction
    : binary_operation
    | unary_operation
    ;

data_section
    : '.dat' data
    ;

data
    : datum (',' datum)*
    ;

datum
    : STRING
    | IDENTIFIER
    | NUMBER
    ;

binary_operation
    : binary_opcode argument_b ',' argument_a
    ;

binary_opcode
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

argument_a
    : REGISTER
    | location_in_register
    | location_offset_by_register
    | POP
    | PEEK
    | pick
    | STACK_POINTER
    | PROGRAM_COUNTER
    | EXTRA
    | location
    | value
    ;

location
    : '[' value ']'
    ;

POP
    : 'POP'
    | 'pop'
    ;

argument_b
    : REGISTER
    | location_in_register
    | location_offset_by_register
    | PUSH
    | PEEK
    | pick
    | STACK_POINTER
    | PROGRAM_COUNTER
    | EXTRA
    | location
    | value
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

unary_operation
    : unary_opcode argument_a
    ;

location_in_register
    : '[' REGISTER ']'
    ;

location_offset_by_register
    : '[' ( REGISTER '+' value | value '+' REGISTER ) ']'
    ;

value
    : IDENTIFIER
    | NUMBER
    ;

NUMBER
    : '0x' [0-9a-fA-F]+
    | [0-9]+
    ;

REGISTER
    : [abcxyzijABCXYZIJ]
    ;

unary_opcode
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
    : [_a-zA-Z]+
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
