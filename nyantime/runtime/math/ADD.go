package math

import (
	"nyantime/registers"
	"nyantime/util"
)

func ADD(r *registers.Registers, operator util.Operator, program *[]util.Instruction) error {
	parsedOperator, err := util.ParseOperator(operator, *program)

	if err != nil {
		return err
	}

	r.IncrementAccumulator(parsedOperator)

	return nil
}
