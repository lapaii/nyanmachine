package control

import (
	"nyantime/registers"
	"nyantime/util"
	"strconv"
)

func JMP(r *registers.Registers, operator util.Operator, program *[]util.Instruction) error {
	parsedOperator, err := strconv.Atoi(string(operator))

	if err != nil {
		return err
	}

	r.SetPC(parsedOperator - 1) // 1 below so that pc increments to right place

	return nil
}
