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

	// EnterFuncDef is called when entering the funcDef production.
	EnterFuncDef(c *FuncDefContext)

	// EnterDataName is called when entering the dataName production.
	EnterDataName(c *DataNameContext)

	// ExitSystem is called when exiting the system production.
	ExitSystem(c *SystemContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDataDef is called when exiting the dataDef production.
	ExitDataDef(c *DataDefContext)

	// ExitFuncDef is called when exiting the funcDef production.
	ExitFuncDef(c *FuncDefContext)

	// ExitDataName is called when exiting the dataName production.
	ExitDataName(c *DataNameContext)
}
