grammar Anzer;

forms : form+ EOF;

form : typeDeclaration
    | funcDeclaration
    | localFuncDeclaration
    | invokeCmd
    ;

typeDeclaration : 'type' typeName '=' typeDefinition;

typeName : UpperIdent;

typeDefinition : typeComplexDefinition | typeSimpleDefinition ;

typeComplexDefinition :  '{' typeField * '}' ;

typeSimpleDefinition : typeId + ;

typeField : fieldName '::' typeDefinition ;

fieldName : LowIdent | UpperIdent;

typeId : typeConstructor | typeScalar | typeOther ;

// Tokens

typeConstructor : typeMinLength | typeMaxLength | typeRight | typeLeft | typeList | typeOptional ;

typeMinLength : 'MinLength' ConstructorArg ;
typeMaxLength : 'MaxLength' ConstructorArg ;
typeRight : 'Right' ;
typeLeft : 'Left' ;
typeList : 'List' ;
typeOptional : 'Optional' ;

typeScalar : typeString | typeInteger | typeFloat | typeBool ;

typeString : 'String' ;
typeInteger : 'Integer' ;
typeFloat : 'Float' ;
typeBool : 'Bool' ;

typeOther : UpperIdent ;

// Func

funcDeclaration : funcName? url '[' runtime ']' '::' funcArgument '->' funcResult;

funcName : LowIdent ;

runtime : LowIdent ;

url : URL ;

funcArgument : typeId + ;

funcResult : typeId + ;

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

ConstructorArg : DecimalDigit+ ;

URL : ( Urlpart | '.' )+ '.' Letter+ '/' ( Urlpart | '/' )+ ;

WS : [ \t\r\n]+ -> skip ;