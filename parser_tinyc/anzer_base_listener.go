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

// EnterProgram is called when production program is entered.
func (s *BaseAnzerListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseAnzerListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseAnzerListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseAnzerListener) ExitStatement(ctx *StatementContext) {}

// EnterParen_expr is called when production paren_expr is entered.
func (s *BaseAnzerListener) EnterParen_expr(ctx *Paren_exprContext) {}

// ExitParen_expr is called when production paren_expr is exited.
func (s *BaseAnzerListener) ExitParen_expr(ctx *Paren_exprContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseAnzerListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseAnzerListener) ExitExpr(ctx *ExprContext) {}

// EnterTest is called when production test is entered.
func (s *BaseAnzerListener) EnterTest(ctx *TestContext) {}

// ExitTest is called when production test is exited.
func (s *BaseAnzerListener) ExitTest(ctx *TestContext) {}

// EnterSum is called when production sum is entered.
func (s *BaseAnzerListener) EnterSum(ctx *SumContext) {}

// ExitSum is called when production sum is exited.
func (s *BaseAnzerListener) ExitSum(ctx *SumContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseAnzerListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseAnzerListener) ExitTerm(ctx *TermContext) {}

// EnterId is called when production id is entered.
func (s *BaseAnzerListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseAnzerListener) ExitId(ctx *IdContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaseAnzerListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaseAnzerListener) ExitInteger(ctx *IntegerContext) {}
