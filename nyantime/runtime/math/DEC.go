package math

import (
	"nyantime/registers"
	"nyantime/util"
)

// operator is IDX or ACC
func DEC(r *registers.Registers, operator util.Operator, program *[]util.Instruction) error {
	parsedOperator := string(operator)

	if parsedOperator == "IDX" {
		r.DecrementIndex(1)
	}

	if parsedOperator == "ACC" {
		r.DecrementAccumulator(1)
	}

	return nil
}
