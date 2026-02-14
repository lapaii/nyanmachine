package control

import (
	"fmt"
	"nyanpreter/instructions"
	"nyanpreter/interpreter/registers"
	"strconv"
)

func JMP(register *registers.Registers, operator instructions.Operator) error {
	conv, err := strconv.Atoi(string(operator))

	if err != nil {
		return fmt.Errorf("error on line %d, invalid jump location!", register.ProgramCounter+1)
	}

	register.SetProgramCounter(conv - 1) // minus one so it jumps to right location

	return nil
}
