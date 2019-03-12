grammar Anzer;

forms : form+ EOF;

form : typeDeclaration
    | functionDeclaration
    ;

typeDeclaration : 'type' typeName '=' typeDefinition;

typeName: UpperIdent;

typeDefinition:  '{' typeField * '}' ;

typeField : fieldName '::' type * ;

fieldName : LowIdent | UpperIdent;

type : LowIdent | UpperIdent;

functionDeclaration: 'func';

// Identifiers

fragment Letter : [A-Za-z_] ;
fragment LowLetter : [a-z] ;
fragment UpperLetter : [A-Z] ;
fragment DecimalDigit: [0-9] ;

LowIdent : LowLetter (Letter | DecimalDigit)* ;

UpperIdent : UpperLetter (Letter | DecimalDigit)* ;

WS : [ \t\r\n]+ -> skip ;