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

	// EnterFuncName is called when entering the funcName production.
	EnterFuncName(c *FuncNameContext)

	// EnterDataName is called when entering the dataName production.
	EnterDataName(c *DataNameContext)

	// EnterDataNameId is called when entering the dataNameId production.
	EnterDataNameId(c *DataNameIdContext)

	// EnterDataContent is called when entering the dataContent production.
	EnterDataContent(c *DataContentContext)

	// EnterJsonContent is called when entering the jsonContent production.
	EnterJsonContent(c *JsonContentContext)

	// ExitSystem is called when exiting the system production.
	ExitSystem(c *SystemContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitFuncName is called when exiting the funcName production.
	ExitFuncName(c *FuncNameContext)

	// ExitDataName is called when exiting the dataName production.
	ExitDataName(c *DataNameContext)

	// ExitDataNameId is called when exiting the dataNameId production.
	ExitDataNameId(c *DataNameIdContext)

	// ExitDataContent is called when exiting the dataContent production.
	ExitDataContent(c *DataContentContext)

	// ExitJsonContent is called when exiting the jsonContent production.
	ExitJsonContent(c *JsonContentContext)
}
