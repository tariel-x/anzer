/*
antlr4 -Dlanguage=Go Anzer.g4
 */
 
grammar Anzer;

system
    : statement+
    ;

statement
    : 'data' dataNameId '=' dataContent
	| funcName '::' dataName '->' dataName
    ;

funcName
    : FUNC_NAME
    ;

FUNC_NAME
    : [a-zA-Z0-9] +
    ;

dataName
    : dataNameId
    | '_'
    ;

dataNameId
    : DATA_NAME_ID
    ;

DATA_NAME_ID
    : [A-Z0-9] +
    ;

dataContent
    : '"' jsonContent '"'
    ;

jsonContent
    : JSON_CONTENT
    ;

JSON_CONTENT
    : '{' [.]+ '}'
    ;

WS: [ \t\r\n]+ -> skip;