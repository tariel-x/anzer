// Generated from /home/nikita/go/src/github.com/tariel-x/anzer/parser/Anzer.g4 by ANTLR 4.7.

package parser // Anzer
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseAnzerVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseAnzerVisitor) VisitSystem(ctx *SystemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAnzerVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAnzerVisitor) VisitDataDef(ctx *DataDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAnzerVisitor) VisitDataName(ctx *DataNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAnzerVisitor) VisitDataNameId(ctx *DataNameIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAnzerVisitor) VisitFuncNameId(ctx *FuncNameIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAnzerVisitor) VisitJson(ctx *JsonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAnzerVisitor) VisitObj(ctx *ObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAnzerVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAnzerVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}
