package parser

import (
	"fmt"
	"nyanpreter/instructions"
	"slices"
)

func ParseOperator(lineNum int, instruction instructions.Operand, part string) (string, error) {
	// check if the operator is of the right type for the instruction
	if slices.Contains(instructions.RegisterOperator, instruction) {
		// instruction needs a register operator
		// is the operator either one of them?
		if part == "ACC" || part == "IDX" {
			// yes it is!!
			return part, nil
		}

		return "", fmt.Errorf(
			"instruction %d on line %d requires a register as the operator, instead got %s",
			instruction, lineNum+1, part,
		)
	}

	return part, nil
}
