package memory

import (
	"nyantime/registers"
	"nyantime/util"
)

func LDR(r *registers.Registers, operator util.Operator, program *[]util.Instruction) error {
	parsedOperator, err := util.ParseOperator(operator)

	if err != nil {
		return err
	}

	r.SetIndex(parsedOperator)

	return nil
}
