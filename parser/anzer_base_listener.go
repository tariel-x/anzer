// Generated from /home/nikita/go/src/github.com/tariel-x/anzer/parser/Anzer.g4 by ANTLR 4.7.

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

// EnterDataSig is called when production dataSig is entered.
func (s *BaseAnzerListener) EnterDataSig(ctx *DataSigContext) {}

// ExitDataSig is called when production dataSig is exited.
func (s *BaseAnzerListener) ExitDataSig(ctx *DataSigContext) {}

// EnterJsonDataDef is called when production jsonDataDef is entered.
func (s *BaseAnzerListener) EnterJsonDataDef(ctx *JsonDataDefContext) {}

// ExitJsonDataDef is called when production jsonDataDef is exited.
func (s *BaseAnzerListener) ExitJsonDataDef(ctx *JsonDataDefContext) {}

// EnterLogicDataDef is called when production logicDataDef is entered.
func (s *BaseAnzerListener) EnterLogicDataDef(ctx *LogicDataDefContext) {}

// ExitLogicDataDef is called when production logicDataDef is exited.
func (s *BaseAnzerListener) ExitLogicDataDef(ctx *LogicDataDefContext) {}

// EnterDataArg is called when production dataArg is entered.
func (s *BaseAnzerListener) EnterDataArg(ctx *DataArgContext) {}

// ExitDataArg is called when production dataArg is exited.
func (s *BaseAnzerListener) ExitDataArg(ctx *DataArgContext) {}

// EnterDataOr is called when production dataOr is entered.
func (s *BaseAnzerListener) EnterDataOr(ctx *DataOrContext) {}

// ExitDataOr is called when production dataOr is exited.
func (s *BaseAnzerListener) ExitDataOr(ctx *DataOrContext) {}

// EnterDataAnd is called when production dataAnd is entered.
func (s *BaseAnzerListener) EnterDataAnd(ctx *DataAndContext) {}

// ExitDataAnd is called when production dataAnd is exited.
func (s *BaseAnzerListener) ExitDataAnd(ctx *DataAndContext) {}

// EnterDataNameId is called when production dataNameId is entered.
func (s *BaseAnzerListener) EnterDataNameId(ctx *DataNameIdContext) {}

// ExitDataNameId is called when production dataNameId is exited.
func (s *BaseAnzerListener) ExitDataNameId(ctx *DataNameIdContext) {}

// EnterFuncSig is called when production funcSig is entered.
func (s *BaseAnzerListener) EnterFuncSig(ctx *FuncSigContext) {}

// ExitFuncSig is called when production funcSig is exited.
func (s *BaseAnzerListener) ExitFuncSig(ctx *FuncSigContext) {}

// EnterFuncParams is called when production funcParams is entered.
func (s *BaseAnzerListener) EnterFuncParams(ctx *FuncParamsContext) {}

// ExitFuncParams is called when production funcParams is exited.
func (s *BaseAnzerListener) ExitFuncParams(ctx *FuncParamsContext) {}

// EnterFuncDef is called when production funcDef is entered.
func (s *BaseAnzerListener) EnterFuncDef(ctx *FuncDefContext) {}

// ExitFuncDef is called when production funcDef is exited.
func (s *BaseAnzerListener) ExitFuncDef(ctx *FuncDefContext) {}

// EnterComposeFunc is called when production composeFunc is entered.
func (s *BaseAnzerListener) EnterComposeFunc(ctx *ComposeFuncContext) {}

// ExitComposeFunc is called when production composeFunc is exited.
func (s *BaseAnzerListener) ExitComposeFunc(ctx *ComposeFuncContext) {}

// EnterProductFunc is called when production productFunc is entered.
func (s *BaseAnzerListener) EnterProductFunc(ctx *ProductFuncContext) {}

// ExitProductFunc is called when production productFunc is exited.
func (s *BaseAnzerListener) ExitProductFunc(ctx *ProductFuncContext) {}

// EnterDataName is called when production dataName is entered.
func (s *BaseAnzerListener) EnterDataName(ctx *DataNameContext) {}

// ExitDataName is called when production dataName is exited.
func (s *BaseAnzerListener) ExitDataName(ctx *DataNameContext) {}

// EnterJson is called when production json is entered.
func (s *BaseAnzerListener) EnterJson(ctx *JsonContext) {}

// ExitJson is called when production json is exited.
func (s *BaseAnzerListener) ExitJson(ctx *JsonContext) {}

// EnterObj is called when production obj is entered.
func (s *BaseAnzerListener) EnterObj(ctx *ObjContext) {}

// ExitObj is called when production obj is exited.
func (s *BaseAnzerListener) ExitObj(ctx *ObjContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseAnzerListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseAnzerListener) ExitPair(ctx *PairContext) {}

// EnterArray is called when production array is entered.
func (s *BaseAnzerListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseAnzerListener) ExitArray(ctx *ArrayContext) {}
