package memory

import (
	"nyantime/registers"
	"nyantime/util"
)

func LDD(r *registers.Registers, operator util.Operator, program *[]util.Instruction) error {
	// operator is an address in this instruction so
	// i know its just a line number i can index to

	parsedValue, err := util.ParseOperator(operator, *program)

	if err != nil {
		return err
	}

	r.SetAccumulator(parsedValue)

	return nil
}
