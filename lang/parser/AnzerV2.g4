grammar AnzerV2;

scheme : term+ EOF;
term : type | func;
// types
type : 'type' UpperIdent '=' typeDefinition;
typeDefinition : record | typeId + ;
record : constructor* '{' field * '}' ;
field : ident '::' typeDefinition ;
typeId : constructor | scalar | other ;
scalar : 'String' | 'Integer' | 'Float' | 'Bool' | 'Boolean' ;
other : UpperIdent ;
constructor : 'MinLength' CArg | 'MaxLength' CArg | 'Right' | 'Left' | 'List' | '[]' | 'Optional' | '*' | 'Either' ;
// funcs
func : composition | reference;
reference : funcName? URL '[' runtime ']' '::' typeId+ '->' typeId+;
composition : LowIdent '=' (funcName | funcBind | '.')+ ;
funcName : LowIdent ;
funcBind : '>>=' funcName;

// Identifiers

fragment Letter : [A-Za-z_] ;
fragment LowLetter : [a-z] ;
fragment UpperLetter : [A-Z] ;
fragment Digit: [0-9] ;
fragment Urlpart : ('-' | '_' | '.' | Letter | Digit ) ;

LowIdent : LowLetter (Letter | Digit)* ;

UpperIdent : UpperLetter (Letter | Digit)* ;

ident : LowIdent | UpperIdent;

CArg : Digit+ ;

URL : ( Urlpart | '.' )+ '.' Letter+ '/' ( Urlpart | '/' )+ ;
runtime : LowIdent;

LINE_COMMENT : '//' ~[\r\n]* -> skip ;

WS : [ \t\r\n]+ -> skip ;