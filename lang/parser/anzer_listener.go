// Code generated from Anzer.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Anzer

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AnzerListener is a complete listener for a parse tree produced by AnzerParser.
type AnzerListener interface {
	antlr.ParseTreeListener

	// EnterForms is called when entering the forms production.
	EnterForms(c *FormsContext)

	// EnterForm is called when entering the form production.
	EnterForm(c *FormContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterTypeDefinition is called when entering the typeDefinition production.
	EnterTypeDefinition(c *TypeDefinitionContext)

	// EnterTypeComplexDefinition is called when entering the typeComplexDefinition production.
	EnterTypeComplexDefinition(c *TypeComplexDefinitionContext)

	// EnterTypeSimpleDefinition is called when entering the typeSimpleDefinition production.
	EnterTypeSimpleDefinition(c *TypeSimpleDefinitionContext)

	// EnterTypeField is called when entering the typeField production.
	EnterTypeField(c *TypeFieldContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterTypeId is called when entering the typeId production.
	EnterTypeId(c *TypeIdContext)

	// EnterTypeConstructor is called when entering the typeConstructor production.
	EnterTypeConstructor(c *TypeConstructorContext)

	// EnterTypeMinLength is called when entering the typeMinLength production.
	EnterTypeMinLength(c *TypeMinLengthContext)

	// EnterTypeMaxLength is called when entering the typeMaxLength production.
	EnterTypeMaxLength(c *TypeMaxLengthContext)

	// EnterTypeRight is called when entering the typeRight production.
	EnterTypeRight(c *TypeRightContext)

	// EnterTypeLeft is called when entering the typeLeft production.
	EnterTypeLeft(c *TypeLeftContext)

	// EnterTypeList is called when entering the typeList production.
	EnterTypeList(c *TypeListContext)

	// EnterTypeOptional is called when entering the typeOptional production.
	EnterTypeOptional(c *TypeOptionalContext)

	// EnterTypeScalar is called when entering the typeScalar production.
	EnterTypeScalar(c *TypeScalarContext)

	// EnterTypeString is called when entering the typeString production.
	EnterTypeString(c *TypeStringContext)

	// EnterTypeInteger is called when entering the typeInteger production.
	EnterTypeInteger(c *TypeIntegerContext)

	// EnterTypeFloat is called when entering the typeFloat production.
	EnterTypeFloat(c *TypeFloatContext)

	// EnterTypeBool is called when entering the typeBool production.
	EnterTypeBool(c *TypeBoolContext)

	// EnterTypeOther is called when entering the typeOther production.
	EnterTypeOther(c *TypeOtherContext)

	// EnterFuncDeclaration is called when entering the funcDeclaration production.
	EnterFuncDeclaration(c *FuncDeclarationContext)

	// EnterFuncName is called when entering the funcName production.
	EnterFuncName(c *FuncNameContext)

	// EnterRuntime is called when entering the runtime production.
	EnterRuntime(c *RuntimeContext)

	// EnterUrl is called when entering the url production.
	EnterUrl(c *UrlContext)

	// EnterFuncArgument is called when entering the funcArgument production.
	EnterFuncArgument(c *FuncArgumentContext)

	// EnterFuncResult is called when entering the funcResult production.
	EnterFuncResult(c *FuncResultContext)

	// EnterLocalFuncDeclaration is called when entering the localFuncDeclaration production.
	EnterLocalFuncDeclaration(c *LocalFuncDeclarationContext)

	// EnterFuncRef is called when entering the funcRef production.
	EnterFuncRef(c *FuncRefContext)

	// EnterInvokeCmd is called when entering the invokeCmd production.
	EnterInvokeCmd(c *InvokeCmdContext)

	// ExitForms is called when exiting the forms production.
	ExitForms(c *FormsContext)

	// ExitForm is called when exiting the form production.
	ExitForm(c *FormContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitTypeDefinition is called when exiting the typeDefinition production.
	ExitTypeDefinition(c *TypeDefinitionContext)

	// ExitTypeComplexDefinition is called when exiting the typeComplexDefinition production.
	ExitTypeComplexDefinition(c *TypeComplexDefinitionContext)

	// ExitTypeSimpleDefinition is called when exiting the typeSimpleDefinition production.
	ExitTypeSimpleDefinition(c *TypeSimpleDefinitionContext)

	// ExitTypeField is called when exiting the typeField production.
	ExitTypeField(c *TypeFieldContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitTypeId is called when exiting the typeId production.
	ExitTypeId(c *TypeIdContext)

	// ExitTypeConstructor is called when exiting the typeConstructor production.
	ExitTypeConstructor(c *TypeConstructorContext)

	// ExitTypeMinLength is called when exiting the typeMinLength production.
	ExitTypeMinLength(c *TypeMinLengthContext)

	// ExitTypeMaxLength is called when exiting the typeMaxLength production.
	ExitTypeMaxLength(c *TypeMaxLengthContext)

	// ExitTypeRight is called when exiting the typeRight production.
	ExitTypeRight(c *TypeRightContext)

	// ExitTypeLeft is called when exiting the typeLeft production.
	ExitTypeLeft(c *TypeLeftContext)

	// ExitTypeList is called when exiting the typeList production.
	ExitTypeList(c *TypeListContext)

	// ExitTypeOptional is called when exiting the typeOptional production.
	ExitTypeOptional(c *TypeOptionalContext)

	// ExitTypeScalar is called when exiting the typeScalar production.
	ExitTypeScalar(c *TypeScalarContext)

	// ExitTypeString is called when exiting the typeString production.
	ExitTypeString(c *TypeStringContext)

	// ExitTypeInteger is called when exiting the typeInteger production.
	ExitTypeInteger(c *TypeIntegerContext)

	// ExitTypeFloat is called when exiting the typeFloat production.
	ExitTypeFloat(c *TypeFloatContext)

	// ExitTypeBool is called when exiting the typeBool production.
	ExitTypeBool(c *TypeBoolContext)

	// ExitTypeOther is called when exiting the typeOther production.
	ExitTypeOther(c *TypeOtherContext)

	// ExitFuncDeclaration is called when exiting the funcDeclaration production.
	ExitFuncDeclaration(c *FuncDeclarationContext)

	// ExitFuncName is called when exiting the funcName production.
	ExitFuncName(c *FuncNameContext)

	// ExitRuntime is called when exiting the runtime production.
	ExitRuntime(c *RuntimeContext)

	// ExitUrl is called when exiting the url production.
	ExitUrl(c *UrlContext)

	// ExitFuncArgument is called when exiting the funcArgument production.
	ExitFuncArgument(c *FuncArgumentContext)

	// ExitFuncResult is called when exiting the funcResult production.
	ExitFuncResult(c *FuncResultContext)

	// ExitLocalFuncDeclaration is called when exiting the localFuncDeclaration production.
	ExitLocalFuncDeclaration(c *LocalFuncDeclarationContext)

	// ExitFuncRef is called when exiting the funcRef production.
	ExitFuncRef(c *FuncRefContext)

	// ExitInvokeCmd is called when exiting the invokeCmd production.
	ExitInvokeCmd(c *InvokeCmdContext)
}
