// Generated from /home/nikita/go/src/github.com/tariel-x/anzer/parser/Anzer.g4 by ANTLR 4.7.

package parser // Anzer
import "github.com/antlr/antlr4/runtime/Go/antlr"

// AnzerListener is a complete listener for a parse tree produced by AnzerParser.
type AnzerListener interface {
	antlr.ParseTreeListener

	// EnterSystem is called when entering the system production.
	EnterSystem(c *SystemContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDataSig is called when entering the dataSig production.
	EnterDataSig(c *DataSigContext)

	// EnterJsonDataDef is called when entering the jsonDataDef production.
	EnterJsonDataDef(c *JsonDataDefContext)

	// EnterLogicDataDef is called when entering the logicDataDef production.
	EnterLogicDataDef(c *LogicDataDefContext)

	// EnterDataArg is called when entering the dataArg production.
	EnterDataArg(c *DataArgContext)

	// EnterDataOr is called when entering the dataOr production.
	EnterDataOr(c *DataOrContext)

	// EnterDataAnd is called when entering the dataAnd production.
	EnterDataAnd(c *DataAndContext)

	// EnterDataNameId is called when entering the dataNameId production.
	EnterDataNameId(c *DataNameIdContext)

	// EnterFuncSig is called when entering the funcSig production.
	EnterFuncSig(c *FuncSigContext)

	// EnterFuncParams is called when entering the funcParams production.
	EnterFuncParams(c *FuncParamsContext)

	// EnterFuncDef is called when entering the funcDef production.
	EnterFuncDef(c *FuncDefContext)

	// EnterComposeFunc is called when entering the composeFunc production.
	EnterComposeFunc(c *ComposeFuncContext)

	// EnterProductFunc is called when entering the productFunc production.
	EnterProductFunc(c *ProductFuncContext)

	// EnterDataName is called when entering the dataName production.
	EnterDataName(c *DataNameContext)

	// EnterJson is called when entering the json production.
	EnterJson(c *JsonContext)

	// EnterObj is called when entering the obj production.
	EnterObj(c *ObjContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// ExitSystem is called when exiting the system production.
	ExitSystem(c *SystemContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDataSig is called when exiting the dataSig production.
	ExitDataSig(c *DataSigContext)

	// ExitJsonDataDef is called when exiting the jsonDataDef production.
	ExitJsonDataDef(c *JsonDataDefContext)

	// ExitLogicDataDef is called when exiting the logicDataDef production.
	ExitLogicDataDef(c *LogicDataDefContext)

	// ExitDataArg is called when exiting the dataArg production.
	ExitDataArg(c *DataArgContext)

	// ExitDataOr is called when exiting the dataOr production.
	ExitDataOr(c *DataOrContext)

	// ExitDataAnd is called when exiting the dataAnd production.
	ExitDataAnd(c *DataAndContext)

	// ExitDataNameId is called when exiting the dataNameId production.
	ExitDataNameId(c *DataNameIdContext)

	// ExitFuncSig is called when exiting the funcSig production.
	ExitFuncSig(c *FuncSigContext)

	// ExitFuncParams is called when exiting the funcParams production.
	ExitFuncParams(c *FuncParamsContext)

	// ExitFuncDef is called when exiting the funcDef production.
	ExitFuncDef(c *FuncDefContext)

	// ExitComposeFunc is called when exiting the composeFunc production.
	ExitComposeFunc(c *ComposeFuncContext)

	// ExitProductFunc is called when exiting the productFunc production.
	ExitProductFunc(c *ProductFuncContext)

	// ExitDataName is called when exiting the dataName production.
	ExitDataName(c *DataNameContext)

	// ExitJson is called when exiting the json production.
	ExitJson(c *JsonContext)

	// ExitObj is called when exiting the obj production.
	ExitObj(c *ObjContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)
}
