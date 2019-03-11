grammar Anzer;

forms : form+ EOF;

form : (typeDeclaration | functionDeclaration) ;

typeDeclaration : 'type' typeName '=' typeDefinition;

typeName: Ident;

typeDefinition:  '{' ( typeField )* '}' ;

typeField : fieldName '::' type* ;

fieldName : Field;

type : Ident;

WS : [ \t\r\n]+ -> skip ;

functionDeclaration: 'func';

// Identifiers

fragment Letter : [A-Za-z_];
fragment LowLetter : [a-z];
fragment DecimalDigit: [0-9];

Ident : Letter (Letter | DecimalDigit)*;

Field : LowLetter (Letter | DecimalDigit)*;