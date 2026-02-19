package math

import (
	"nyantime/registers"
	"nyantime/util"
)

func SUB(r *registers.Registers, operator util.Operator, program *[]util.Instruction) error {
	parsedOperator, err := util.ParseOperator(operator)

	if err != nil {
		return err
	}

	r.DecrementAccumulator(parsedOperator)

	return nil
}
