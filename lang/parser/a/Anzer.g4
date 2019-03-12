grammar Anzer;

forms : form+ EOF;

form : typeDeclaration
    | funcDeclaration
    ;

typeDeclaration : 'type' typeName '=' typeDefinition;

typeName: UpperIdent;

typeDefinition:  '{' typeField * '}' ;

typeField : fieldName '::' type * ;

fieldName : LowIdent | UpperIdent;

type : LowIdent | UpperIdent;

// Func

funcDeclaration: funcName? url '[' runtime ']' '::' typeName '->' typeName;

funcName : LowIdent ;

runtime : LowIdent ;

url : URL ;

// Identifiers

fragment Letter : [A-Za-z_] ;
fragment LowLetter : [a-z] ;
fragment UpperLetter : [A-Z] ;
fragment DecimalDigit: [0-9] ;

LowIdent : LowLetter (Letter | DecimalDigit)* ;

UpperIdent : UpperLetter (Letter | DecimalDigit)* ;

URL : ('a' .. 'z')+ '.' ('a' .. 'z')+ '/' ( ('0' .. '9') | ('a' .. 'z') | '-' | '_' | '/' )+ ;

WS : [ \t\r\n]+ -> skip ;