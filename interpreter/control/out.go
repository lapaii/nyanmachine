package control

import (
	"fmt"
	"nyanpreter/instructions"
	"nyanpreter/interpreter/registers"
)

func OUT(register *registers.Registers, operator instructions.Operator) error {
	accumulatorValue := register.GetAccumulator()

	fmt.Printf("%c\n", accumulatorValue)

	return nil
}
