/*
antlr4 -Dlanguage=Go Anzer.g4
 */
 
grammar Anzer;

system
    : statement+
    ;

statement
    : dataDef
    | funcDef
    ;

dataDef 
    : 'data' DATA_NAME '=' DATA_CONTENT
    ;

funcDef
    : FUNC_NAME '::' dataName '->' dataName
    ;

FUNC_NAME
    : [a-zA-Z0-9]+
    ;

dataName
    : DATA_NAME
    | '_'
    ;

DATA_NAME
    : ([A-Z0-9]+)
    ;

DATA_CONTENT
    : '"' JSON_STRING '"'
    ;

JSON_STRING
    : '{' [.]+ '}'
    ;

WS
   : [ \t\n\r] + -> skip
   ;
