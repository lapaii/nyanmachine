package util

import (
	"fmt"
	"nyanpreter/instructions"
	"nyanpreter/interpreter/registers"
	"strconv"
)

func ConvertOperator(register *registers.Registers, operator instructions.Operator) (int, error) {
	var numberBase int

	switch operator[0] {
	case '#':
		numberBase = 10
	case 'B':
		numberBase = 2
	case '&':
		numberBase = 16
	default:
		numberBase = -1 // setting to minus one, cant return the function from here?
	}

	if numberBase == -1 {
		return 0, fmt.Errorf("error in line %d, invalid operator syntax!", register.ProgramCounter+1)
	}

	conv, err := strconv.ParseInt(string(operator)[1:], numberBase, 0)

	if err != nil {
		return 0, fmt.Errorf("error on line %d, invalid operator syntax!", register.ProgramCounter+1)
	}

	return int(conv), nil
}
