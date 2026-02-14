package math

import (
	"nyanpreter/instructions"
	"nyanpreter/interpreter/registers"
	"nyanpreter/interpreter/util"
)

func ADD(register *registers.Registers, operator instructions.Operator) error {
	result, err := util.ConvertOperator(register, operator)

	if err != nil {
		return err
	}

	register.AddToAccumulator(result)

	return nil
}
