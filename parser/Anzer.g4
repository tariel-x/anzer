/*
This is the grammar for Anzer definition language of services system.
antlr4 -Dlanguage=Go Anzer.g4
 */
 
grammar Anzer;

system: statement+;

statement
    : dataSig
    | funcSig
    | funcDef
    ;

/**
 * Data type
 */

dataSig: jsonDataDef | logicDataDef;

jsonDataDef: 'data' DATA_NAME_ID '=' json;

logicDataDef: 'data' DATA_NAME_ID '=' (dataAnd | dataOr);

dataArg: dataNameId | json;

dataOr: DATA_NAME_ID (DATA_OR DATA_NAME_ID)*;

dataAnd: DATA_NAME_ID (DATA_AND DATA_NAME_ID)*;

DATA_AND: '&';

DATA_OR: '|';

dataNameId: DATA_NAME_ID;

DATA_NAME_ID: [A-Z] [A-Z0-9_] +;

/**
 * Functions
 */

//signature
funcSig: FUNC_NAME_ID '::' dataName '->' dataName funcParams?;

funcParams: json;

//definition
funcDef: FUNC_NAME_ID '=' composeFunc ('.' composeFunc)*;

// composition
composeFunc: FUNC_NAME_ID | productFunc;

productFunc: '<' composeFunc (',' composeFunc)* '>';

dataName: DATA_NAME_ID | '_';

FUNC_NAME_ID: [a-z] [a-zA-Z0-9]*;

FUNC_PARAM_ID: [a-z0-9_]+;
/**
 * JSON Definition
 */

json: STRING
   | NUMBER
   | obj
   | array
   | 'true'
   | 'false'
   | 'null'
   ;

obj: '{' pair (',' pair)* '}'
   | '{' '}'
   ;

pair: STRING ':' json;

array: '[' json (',' json)* ']'
    | '[' ']'
    ;

STRING: '"' (ESC | SAFECODEPOINT)* '"';


fragment ESC: '\\' (["\\/bfnrt] | UNICODE);


fragment UNICODE: 'u' HEX HEX HEX HEX;


fragment HEX: [0-9a-fA-F];


fragment SAFECODEPOINT: ~ ["\\\u0000-\u001F];


NUMBER: '-'? INT ('.' [0-9] +)? EXP?;


fragment INT: '0' | [1-9] [0-9]*;

// no leading zeros

fragment EXP: [Ee] [+\-]? INT;


/**
 * WS
 */

WS: [ \t\n\r] + -> skip;

/**
 * Comments
 */

LineComment: '//' ~[\r\n]* -> skip;