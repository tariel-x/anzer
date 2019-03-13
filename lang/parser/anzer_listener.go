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

	// EnterTypeField is called when entering the typeField production.
	EnterTypeField(c *TypeFieldContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterTypeId is called when entering the typeId production.
	EnterTypeId(c *TypeIdContext)

	// EnterFuncDeclaration is called when entering the funcDeclaration production.
	EnterFuncDeclaration(c *FuncDeclarationContext)

	// EnterFuncName is called when entering the funcName production.
	EnterFuncName(c *FuncNameContext)

	// EnterRuntime is called when entering the runtime production.
	EnterRuntime(c *RuntimeContext)

	// EnterUrl is called when entering the url production.
	EnterUrl(c *UrlContext)

	// EnterLocalFuncDeclaration is called when entering the localFuncDeclaration production.
	EnterLocalFuncDeclaration(c *LocalFuncDeclarationContext)

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

	// ExitTypeField is called when exiting the typeField production.
	ExitTypeField(c *TypeFieldContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitTypeId is called when exiting the typeId production.
	ExitTypeId(c *TypeIdContext)

	// ExitFuncDeclaration is called when exiting the funcDeclaration production.
	ExitFuncDeclaration(c *FuncDeclarationContext)

	// ExitFuncName is called when exiting the funcName production.
	ExitFuncName(c *FuncNameContext)

	// ExitRuntime is called when exiting the runtime production.
	ExitRuntime(c *RuntimeContext)

	// ExitUrl is called when exiting the url production.
	ExitUrl(c *UrlContext)

	// ExitLocalFuncDeclaration is called when exiting the localFuncDeclaration production.
	ExitLocalFuncDeclaration(c *LocalFuncDeclarationContext)

	// ExitInvokeCmd is called when exiting the invokeCmd production.
	ExitInvokeCmd(c *InvokeCmdContext)
}
