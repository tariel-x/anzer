// Generated from /home/nikita/go/src/github.com/tariel-x/anzer/parser/Anzer.g4 by ANTLR 4.7.

package parser // Anzer
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by AnzerParser.
type AnzerVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by AnzerParser#system.
	VisitSystem(ctx *SystemContext) interface{}

	// Visit a parse tree produced by AnzerParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by AnzerParser#dataSig.
	VisitDataSig(ctx *DataSigContext) interface{}

	// Visit a parse tree produced by AnzerParser#funcSig.
	VisitFuncSig(ctx *FuncSigContext) interface{}

	// Visit a parse tree produced by AnzerParser#funcDef.
	VisitFuncDef(ctx *FuncDefContext) interface{}

	// Visit a parse tree produced by AnzerParser#definingFuncName.
	VisitDefiningFuncName(ctx *DefiningFuncNameContext) interface{}

	// Visit a parse tree produced by AnzerParser#dataName.
	VisitDataName(ctx *DataNameContext) interface{}

	// Visit a parse tree produced by AnzerParser#dataNameId.
	VisitDataNameId(ctx *DataNameIdContext) interface{}

	// Visit a parse tree produced by AnzerParser#funcNameId.
	VisitFuncNameId(ctx *FuncNameIdContext) interface{}

	// Visit a parse tree produced by AnzerParser#json.
	VisitJson(ctx *JsonContext) interface{}

	// Visit a parse tree produced by AnzerParser#obj.
	VisitObj(ctx *ObjContext) interface{}

	// Visit a parse tree produced by AnzerParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by AnzerParser#array.
	VisitArray(ctx *ArrayContext) interface{}
}
