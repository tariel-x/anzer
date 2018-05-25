// Code generated from Anzer.g4 by ANTLR 4.7.1. DO NOT EDIT.

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

	// EnterDataDefinition is called when entering the dataDefinition production.
	EnterDataDefinition(c *DataDefinitionContext)

	// EnterFuncSig is called when entering the funcSig production.
	EnterFuncSig(c *FuncSigContext)

	// EnterFuncDef is called when entering the funcDef production.
	EnterFuncDef(c *FuncDefContext)

	// EnterFuncParam is called when entering the funcParam production.
	EnterFuncParam(c *FuncParamContext)

	// EnterFuncParamEnv is called when entering the funcParamEnv production.
	EnterFuncParamEnv(c *FuncParamEnvContext)

	// EnterFuncParamConfig is called when entering the funcParamConfig production.
	EnterFuncParamConfig(c *FuncParamConfigContext)

	// EnterComposeFunc is called when entering the composeFunc production.
	EnterComposeFunc(c *ComposeFuncContext)

	// EnterProductFunc is called when entering the productFunc production.
	EnterProductFunc(c *ProductFuncContext)

	// EnterDataName is called when entering the dataName production.
	EnterDataName(c *DataNameContext)

	// EnterFuncEnvName is called when entering the funcEnvName production.
	EnterFuncEnvName(c *FuncEnvNameContext)

	// EnterFuncParamValue is called when entering the funcParamValue production.
	EnterFuncParamValue(c *FuncParamValueContext)

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

	// ExitDataDefinition is called when exiting the dataDefinition production.
	ExitDataDefinition(c *DataDefinitionContext)

	// ExitFuncSig is called when exiting the funcSig production.
	ExitFuncSig(c *FuncSigContext)

	// ExitFuncDef is called when exiting the funcDef production.
	ExitFuncDef(c *FuncDefContext)

	// ExitFuncParam is called when exiting the funcParam production.
	ExitFuncParam(c *FuncParamContext)

	// ExitFuncParamEnv is called when exiting the funcParamEnv production.
	ExitFuncParamEnv(c *FuncParamEnvContext)

	// ExitFuncParamConfig is called when exiting the funcParamConfig production.
	ExitFuncParamConfig(c *FuncParamConfigContext)

	// ExitComposeFunc is called when exiting the composeFunc production.
	ExitComposeFunc(c *ComposeFuncContext)

	// ExitProductFunc is called when exiting the productFunc production.
	ExitProductFunc(c *ProductFuncContext)

	// ExitDataName is called when exiting the dataName production.
	ExitDataName(c *DataNameContext)

	// ExitFuncEnvName is called when exiting the funcEnvName production.
	ExitFuncEnvName(c *FuncEnvNameContext)

	// ExitFuncParamValue is called when exiting the funcParamValue production.
	ExitFuncParamValue(c *FuncParamValueContext)

	// ExitJson is called when exiting the json production.
	ExitJson(c *JsonContext)

	// ExitObj is called when exiting the obj production.
	ExitObj(c *ObjContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)
}
