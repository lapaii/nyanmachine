package control

import (
	"nyantime/registers"
	"nyantime/util"
)

// CMP operator is an address
func CMP(r *registers.Registers, operator util.Operator, program *[]util.Instruction) error {
	parsedOperator, err := util.ParseOperator(operator) // just line number

	valueToCompare, err := util.ParseOperator((*program)[parsedOperator].Operator)

	if err != nil {
		return err
	}

	if r.GetAccumulator() == valueToCompare {
		r.SetCompareResult(true)
	} else {
		r.SetCompareResult(false)
	}

	return nil
}
