/*
antlr4 -Dlanguage=Go Anzer.g4
 */
 
grammar Anzer;

system: statement+;

statement: data
	| func
    ;

data: 'data' DATA_NAME_ID '=' dataContent;

func: FUNC_NAME '::' dataName '->' dataName;

FUNC_NAME: [a-zA-Z0-9] +;

dataName: DATA_NAME_ID
    | '_'
    ;

DATA_NAME_ID: [A-Z0-9] +;

dataContent: '"' jsonContent '"';

jsonContent: JSON_CONTENT;

JSON_CONTENT: '{' [.]+ '}';

WS: [ \t\n\r]+ -> skip;