package control

import (
	"nyantime/registers"
	"nyantime/util"
)

func CMP(r *registers.Registers, operator util.Operator, program *[]util.Instruction) error {
	parsedOperator, err := util.ParseOperator(operator, *program)

	if err != nil {
		return err
	}

	if r.GetAccumulator() == parsedOperator {
		r.SetCompareResult(true)
	} else {
		r.SetCompareResult(false)
	}

	return nil
}
