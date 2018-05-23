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

	// EnterDataDef is called when entering the dataDef production.
	EnterDataDef(c *DataDefContext)

	// EnterDataName is called when entering the dataName production.
	EnterDataName(c *DataNameContext)

	// EnterDataNameId is called when entering the dataNameId production.
	EnterDataNameId(c *DataNameIdContext)

	// EnterFuncNameId is called when entering the funcNameId production.
	EnterFuncNameId(c *FuncNameIdContext)

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

	// ExitDataDef is called when exiting the dataDef production.
	ExitDataDef(c *DataDefContext)

	// ExitDataName is called when exiting the dataName production.
	ExitDataName(c *DataNameContext)

	// ExitDataNameId is called when exiting the dataNameId production.
	ExitDataNameId(c *DataNameIdContext)

	// ExitFuncNameId is called when exiting the funcNameId production.
	ExitFuncNameId(c *FuncNameIdContext)

	// ExitJson is called when exiting the json production.
	ExitJson(c *JsonContext)

	// ExitObj is called when exiting the obj production.
	ExitObj(c *ObjContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)
}
