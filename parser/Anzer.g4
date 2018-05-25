/*
antlr4 -Dlanguage=Go Anzer.g4
 */
 
grammar Anzer;

system: statement+;

statement
    : dataSig
    | funcSig
    | funcDef
    | funcParam
    ;

/**
 * Data type
 */

dataSig: 'data' dataNameId '=' dataDefinition;

dataDefinition: json;

dataName: dataNameId | '_';

dataNameId: DATA_NAME_ID;

DATA_NAME_ID: [A-Z] [A-Z0-9_] +;

/**
 * Functions
 */

//signature
funcSig: funcNameId '::' dataName '->' dataName;

//definition
funcDef: funcName '=' funcNameId ('.' funcNameId)*;

//params
funcParam: funcParamConfig | funcParamEnv;

//env
funcParamEnv: funcName '.env[' funcEnvName ']' '=' funcParamValue;

//param
funcParamConfig: funcName '.' funcParamId '=' funcParamValue;

funcName: funcNameId;

funcNameId: FUNC_NAME_ID;

FUNC_NAME_ID: [a-z] [a-zA-Z0-9]*;

funcParamId: FUNC_PARAM_ID;

FUNC_PARAM_ID: [a-z0-9_]+;

funcEnvName: STRING;

funcParamValue: STRING;

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