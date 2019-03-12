// Code generated from Anzer.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Anzer

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAnzerListener is a complete listener for a parse tree produced by AnzerParser.
type BaseAnzerListener struct{}

var _ AnzerListener = &BaseAnzerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAnzerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAnzerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAnzerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAnzerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterForms is called when production forms is entered.
func (s *BaseAnzerListener) EnterForms(ctx *FormsContext) {}

// ExitForms is called when production forms is exited.
func (s *BaseAnzerListener) ExitForms(ctx *FormsContext) {}

// EnterForm is called when production form is entered.
func (s *BaseAnzerListener) EnterForm(ctx *FormContext) {}

// ExitForm is called when production form is exited.
func (s *BaseAnzerListener) ExitForm(ctx *FormContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseAnzerListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseAnzerListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseAnzerListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseAnzerListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterTypeDefinition is called when production typeDefinition is entered.
func (s *BaseAnzerListener) EnterTypeDefinition(ctx *TypeDefinitionContext) {}

// ExitTypeDefinition is called when production typeDefinition is exited.
func (s *BaseAnzerListener) ExitTypeDefinition(ctx *TypeDefinitionContext) {}

// EnterTypeField is called when production typeField is entered.
func (s *BaseAnzerListener) EnterTypeField(ctx *TypeFieldContext) {}

// ExitTypeField is called when production typeField is exited.
func (s *BaseAnzerListener) ExitTypeField(ctx *TypeFieldContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseAnzerListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseAnzerListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterTypeId is called when production typeId is entered.
func (s *BaseAnzerListener) EnterTypeId(ctx *TypeIdContext) {}

// ExitTypeId is called when production typeId is exited.
func (s *BaseAnzerListener) ExitTypeId(ctx *TypeIdContext) {}

// EnterFuncDeclaration is called when production funcDeclaration is entered.
func (s *BaseAnzerListener) EnterFuncDeclaration(ctx *FuncDeclarationContext) {}

// ExitFuncDeclaration is called when production funcDeclaration is exited.
func (s *BaseAnzerListener) ExitFuncDeclaration(ctx *FuncDeclarationContext) {}

// EnterFuncName is called when production funcName is entered.
func (s *BaseAnzerListener) EnterFuncName(ctx *FuncNameContext) {}

// ExitFuncName is called when production funcName is exited.
func (s *BaseAnzerListener) ExitFuncName(ctx *FuncNameContext) {}

// EnterRuntime is called when production runtime is entered.
func (s *BaseAnzerListener) EnterRuntime(ctx *RuntimeContext) {}

// ExitRuntime is called when production runtime is exited.
func (s *BaseAnzerListener) ExitRuntime(ctx *RuntimeContext) {}

// EnterUrl is called when production url is entered.
func (s *BaseAnzerListener) EnterUrl(ctx *UrlContext) {}

// ExitUrl is called when production url is exited.
func (s *BaseAnzerListener) ExitUrl(ctx *UrlContext) {}

// EnterLocalFuncDeclaration is called when production localFuncDeclaration is entered.
func (s *BaseAnzerListener) EnterLocalFuncDeclaration(ctx *LocalFuncDeclarationContext) {}

// ExitLocalFuncDeclaration is called when production localFuncDeclaration is exited.
func (s *BaseAnzerListener) ExitLocalFuncDeclaration(ctx *LocalFuncDeclarationContext) {}

// EnterInvokeCmd is called when production invokeCmd is entered.
func (s *BaseAnzerListener) EnterInvokeCmd(ctx *InvokeCmdContext) {}

// ExitInvokeCmd is called when production invokeCmd is exited.
func (s *BaseAnzerListener) ExitInvokeCmd(ctx *InvokeCmdContext) {}
