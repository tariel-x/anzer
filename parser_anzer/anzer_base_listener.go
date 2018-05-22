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

// EnterSystem is called when production system is entered.
func (s *BaseAnzerListener) EnterSystem(ctx *SystemContext) {}

// ExitSystem is called when production system is exited.
func (s *BaseAnzerListener) ExitSystem(ctx *SystemContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseAnzerListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseAnzerListener) ExitStatement(ctx *StatementContext) {}

// EnterFuncName is called when production funcName is entered.
func (s *BaseAnzerListener) EnterFuncName(ctx *FuncNameContext) {}

// ExitFuncName is called when production funcName is exited.
func (s *BaseAnzerListener) ExitFuncName(ctx *FuncNameContext) {}

// EnterDataName is called when production dataName is entered.
func (s *BaseAnzerListener) EnterDataName(ctx *DataNameContext) {}

// ExitDataName is called when production dataName is exited.
func (s *BaseAnzerListener) ExitDataName(ctx *DataNameContext) {}

// EnterDataNameId is called when production dataNameId is entered.
func (s *BaseAnzerListener) EnterDataNameId(ctx *DataNameIdContext) {}

// ExitDataNameId is called when production dataNameId is exited.
func (s *BaseAnzerListener) ExitDataNameId(ctx *DataNameIdContext) {}

// EnterDataContent is called when production dataContent is entered.
func (s *BaseAnzerListener) EnterDataContent(ctx *DataContentContext) {}

// ExitDataContent is called when production dataContent is exited.
func (s *BaseAnzerListener) ExitDataContent(ctx *DataContentContext) {}

// EnterJsonContent is called when production jsonContent is entered.
func (s *BaseAnzerListener) EnterJsonContent(ctx *JsonContentContext) {}

// ExitJsonContent is called when production jsonContent is exited.
func (s *BaseAnzerListener) ExitJsonContent(ctx *JsonContentContext) {}
