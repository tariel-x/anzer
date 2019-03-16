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

// EnterTypeComplexDefinition is called when production typeComplexDefinition is entered.
func (s *BaseAnzerListener) EnterTypeComplexDefinition(ctx *TypeComplexDefinitionContext) {}

// ExitTypeComplexDefinition is called when production typeComplexDefinition is exited.
func (s *BaseAnzerListener) ExitTypeComplexDefinition(ctx *TypeComplexDefinitionContext) {}

// EnterTypeSimpleDefinition is called when production typeSimpleDefinition is entered.
func (s *BaseAnzerListener) EnterTypeSimpleDefinition(ctx *TypeSimpleDefinitionContext) {}

// ExitTypeSimpleDefinition is called when production typeSimpleDefinition is exited.
func (s *BaseAnzerListener) ExitTypeSimpleDefinition(ctx *TypeSimpleDefinitionContext) {}

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

// EnterTypeConstructor is called when production typeConstructor is entered.
func (s *BaseAnzerListener) EnterTypeConstructor(ctx *TypeConstructorContext) {}

// ExitTypeConstructor is called when production typeConstructor is exited.
func (s *BaseAnzerListener) ExitTypeConstructor(ctx *TypeConstructorContext) {}

// EnterTypeMinLength is called when production typeMinLength is entered.
func (s *BaseAnzerListener) EnterTypeMinLength(ctx *TypeMinLengthContext) {}

// ExitTypeMinLength is called when production typeMinLength is exited.
func (s *BaseAnzerListener) ExitTypeMinLength(ctx *TypeMinLengthContext) {}

// EnterTypeMaxLength is called when production typeMaxLength is entered.
func (s *BaseAnzerListener) EnterTypeMaxLength(ctx *TypeMaxLengthContext) {}

// ExitTypeMaxLength is called when production typeMaxLength is exited.
func (s *BaseAnzerListener) ExitTypeMaxLength(ctx *TypeMaxLengthContext) {}

// EnterTypeRight is called when production typeRight is entered.
func (s *BaseAnzerListener) EnterTypeRight(ctx *TypeRightContext) {}

// ExitTypeRight is called when production typeRight is exited.
func (s *BaseAnzerListener) ExitTypeRight(ctx *TypeRightContext) {}

// EnterTypeLeft is called when production typeLeft is entered.
func (s *BaseAnzerListener) EnterTypeLeft(ctx *TypeLeftContext) {}

// ExitTypeLeft is called when production typeLeft is exited.
func (s *BaseAnzerListener) ExitTypeLeft(ctx *TypeLeftContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseAnzerListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseAnzerListener) ExitTypeList(ctx *TypeListContext) {}

// EnterTypeOptional is called when production typeOptional is entered.
func (s *BaseAnzerListener) EnterTypeOptional(ctx *TypeOptionalContext) {}

// ExitTypeOptional is called when production typeOptional is exited.
func (s *BaseAnzerListener) ExitTypeOptional(ctx *TypeOptionalContext) {}

// EnterTypeScalar is called when production typeScalar is entered.
func (s *BaseAnzerListener) EnterTypeScalar(ctx *TypeScalarContext) {}

// ExitTypeScalar is called when production typeScalar is exited.
func (s *BaseAnzerListener) ExitTypeScalar(ctx *TypeScalarContext) {}

// EnterTypeString is called when production typeString is entered.
func (s *BaseAnzerListener) EnterTypeString(ctx *TypeStringContext) {}

// ExitTypeString is called when production typeString is exited.
func (s *BaseAnzerListener) ExitTypeString(ctx *TypeStringContext) {}

// EnterTypeInteger is called when production typeInteger is entered.
func (s *BaseAnzerListener) EnterTypeInteger(ctx *TypeIntegerContext) {}

// ExitTypeInteger is called when production typeInteger is exited.
func (s *BaseAnzerListener) ExitTypeInteger(ctx *TypeIntegerContext) {}

// EnterTypeFloat is called when production typeFloat is entered.
func (s *BaseAnzerListener) EnterTypeFloat(ctx *TypeFloatContext) {}

// ExitTypeFloat is called when production typeFloat is exited.
func (s *BaseAnzerListener) ExitTypeFloat(ctx *TypeFloatContext) {}

// EnterTypeBool is called when production typeBool is entered.
func (s *BaseAnzerListener) EnterTypeBool(ctx *TypeBoolContext) {}

// ExitTypeBool is called when production typeBool is exited.
func (s *BaseAnzerListener) ExitTypeBool(ctx *TypeBoolContext) {}

// EnterTypeOther is called when production typeOther is entered.
func (s *BaseAnzerListener) EnterTypeOther(ctx *TypeOtherContext) {}

// ExitTypeOther is called when production typeOther is exited.
func (s *BaseAnzerListener) ExitTypeOther(ctx *TypeOtherContext) {}

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

// EnterFuncArgument is called when production funcArgument is entered.
func (s *BaseAnzerListener) EnterFuncArgument(ctx *FuncArgumentContext) {}

// ExitFuncArgument is called when production funcArgument is exited.
func (s *BaseAnzerListener) ExitFuncArgument(ctx *FuncArgumentContext) {}

// EnterFuncResult is called when production funcResult is entered.
func (s *BaseAnzerListener) EnterFuncResult(ctx *FuncResultContext) {}

// ExitFuncResult is called when production funcResult is exited.
func (s *BaseAnzerListener) ExitFuncResult(ctx *FuncResultContext) {}

// EnterLocalFuncDeclaration is called when production localFuncDeclaration is entered.
func (s *BaseAnzerListener) EnterLocalFuncDeclaration(ctx *LocalFuncDeclarationContext) {}

// ExitLocalFuncDeclaration is called when production localFuncDeclaration is exited.
func (s *BaseAnzerListener) ExitLocalFuncDeclaration(ctx *LocalFuncDeclarationContext) {}

// EnterFuncRef is called when production funcRef is entered.
func (s *BaseAnzerListener) EnterFuncRef(ctx *FuncRefContext) {}

// ExitFuncRef is called when production funcRef is exited.
func (s *BaseAnzerListener) ExitFuncRef(ctx *FuncRefContext) {}

// EnterInvokeCmd is called when production invokeCmd is entered.
func (s *BaseAnzerListener) EnterInvokeCmd(ctx *InvokeCmdContext) {}

// ExitInvokeCmd is called when production invokeCmd is exited.
func (s *BaseAnzerListener) ExitInvokeCmd(ctx *InvokeCmdContext) {}
