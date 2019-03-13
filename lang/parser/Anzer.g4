grammar Anzer;

forms : form+ EOF;

form : typeDeclaration
    | funcDeclaration
    | localFuncDeclaration
    | invokeCmd
    ;

typeDeclaration : 'type' typeName '=' typeDefinition;

typeName : UpperIdent;

typeDefinition :  '{' typeField * '}' ;

typeField : fieldName '::' typeId * ;

fieldName : LowIdent | UpperIdent;

typeId : LowIdent | UpperIdent;

// Func

funcDeclaration : funcName? url '[' runtime ']' '::' typeName+ '->' typeName+;

funcName : LowIdent ;

runtime : LowIdent ;

url : URL ;

// Local func

localFuncDeclaration : funcName '=' (funcName | '.')+ ;

// Invoke

invokeCmd : 'invoke' '(' funcName+ ')' ;

// Identifiers

fragment Letter : [A-Za-z_] ;
fragment LowLetter : [a-z] ;
fragment UpperLetter : [A-Z] ;
fragment DecimalDigit: [0-9] ;
fragment Urlpart : ('-' | Letter | DecimalDigit ) ;

LowIdent : LowLetter (Letter | DecimalDigit)* ;

UpperIdent : UpperLetter (Letter | DecimalDigit)* ;

URL : ( Urlpart | '.' )+ '.' Letter+ '/' ( Urlpart | '/' )+ ;

WS : [ \t\r\n]+ -> skip ;