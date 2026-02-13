package parser

import (
	"fmt"
	"nyanpreter/instructions"
)

func ParseOperand(lineNum int, part string) (instructions.Operand, error) {
	var operand instructions.Operand
	var returnError error = nil

	operand = instructions.InstructionSet[part]

	if operand == instructions.INVALID {
		returnError = fmt.Errorf("Invalid instruction at Line %d: %s", lineNum+1, part)
	}

	return operand, returnError
}
