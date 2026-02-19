package memory

import (
	"nyantime/registers"
	"nyantime/util"
)

// register operator!!
func MOV(r *registers.Registers, operator util.Operator, program *[]util.Instruction) error {
	parsedOperator := string(operator)

	if parsedOperator == "IDX" {
		accValue := r.GetAccumulator()
		r.SetIndex(accValue)
	}

	// no point doing anything if its ACC, as the value of ACC is already there???

	return nil
}
