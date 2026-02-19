package memory

import (
	"fmt"
	"nyantime/registers"
	"nyantime/util"
	"strconv"
)

func STO(r *registers.Registers, operator util.Operator, program *[]util.Instruction) error {
	// sto just sets the value of the operator on the instruction
	// that this instructions operator points to the value of the acc

	parsedOperator, err := strconv.Atoi(string(operator))

	if err != nil {
		return err
	}

	util.ModifyInstruction(program, parsedOperator, util.Instruction{
		Operand:  (*program)[parsedOperator].Operand,
		Operator: util.Operator(fmt.Sprintf("#%d", r.Accumulator)),
	})

	return nil
}
