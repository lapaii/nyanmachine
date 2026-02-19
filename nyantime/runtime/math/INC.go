package math

import (
	"nyantime/registers"
	"nyantime/util"
)

// operator is IDX or ACC
func INC(r *registers.Registers, operator util.Operator, program *[]util.Instruction) error {
	parsedOperator := string(operator)

	if parsedOperator == "IDX" {
		r.IncrementIndex(1)
	}

	if parsedOperator == "ACC" {
		r.IncrementAccumulator(1)
	}

	return nil
}
