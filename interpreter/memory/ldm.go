package memory

import (
	"nyanpreter/instructions"
	"nyanpreter/interpreter/registers"
	"nyanpreter/interpreter/util"
)

func LDM(register *registers.Registers, operator instructions.Operator) error {
	result, err := util.ConvertOperator(register, operator)

	if err != nil {
		return err
	}

	register.SetAccumulator(result)

	return nil
}
