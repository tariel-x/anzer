grammar Anzer;

forms : form+ EOF;

form : typeDeclaration
    | funcDeclaration
    | localFuncDeclaration
    | invokeCmd
    | packageDeclaration
    ;

packageDeclaration: 'package' packageName;

packageName: LowIdent | UpperIdent;

typeDeclaration : 'type' typeName '=' typeDefinition;

typeName : UpperIdent;

typeDefinition : typeComplexDefinition | typeSimpleDefinition ;

typeComplexDefinition : typeConstructor* '{' typeField * '}' ;

typeSimpleDefinition : typeId + ;

typeField : fieldName '::' typeDefinition ;

fieldName : LowIdent | UpperIdent;

typeId : typeConstructor | typeScalar | typeOther ;

// Tokens

typeConstructor : typeMinLength | typeMaxLength | typeRight | typeLeft | typeList | typeOptional | typeEither ;

typeMinLength : 'MinLength' ConstructorArg ;
typeMaxLength : 'MaxLength' ConstructorArg ;
typeEither : 'Either';
typeRight : 'Right' ;
typeLeft : 'Left' ;
typeList : 'List' | '[]' ;
typeOptional : 'Optional' | '*' ;

typeScalar : typeString | typeInteger | typeFloat | typeBool ;

typeString : 'String' ;
typeInteger : 'Integer' ;
typeFloat : 'Float' ;
typeBool : 'Bool' | 'Boolean' ;

typeOther : UpperIdent ;

// Func

funcDeclaration : funcName? url '[' runtime ']' '::' funcArgument '->' funcResult;

funcName : LowIdent ;

runtime : LowIdent ;

url : URL ;

funcArgument : typeId + ;

funcResult : typeId + ;

// Local func

localFuncDeclaration : funcName '=' (funcRef | funcBind | '.')+ ;

funcRef : LowIdent ;

funcBind : '>>=' funcApplied;

funcApplied : LowIdent;

// Invoke

invokeCmd : 'invoke' '(' (invokeFuncName ',')+ ')' ;

invokeFuncName : LowIdent ;

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

LINE_COMMENT : '//' ~[\r\n]* -> skip ;

WS : [ \t\r\n]+ -> skip ;