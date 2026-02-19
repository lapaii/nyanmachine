package control

import (
	"nyantime/registers"
	"nyantime/util"
	"shared"
)

// CMP operator is an address
func CMP(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	parsedOperator, err := util.ParseOperator(operator) // just line number

	valueToCompare, err := util.ParseOperator((*program)[parsedOperator].Operator)

	if err != nil {
		return err
	}

	r.SetEqualFlag(r.GetAccumulator() == valueToCompare)
	r.SetGreaterThanFlag(r.GetAccumulator() > valueToCompare)

	return nil
}
