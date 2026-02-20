package secondpass

import (
	"fmt"
	"shared"
)

func ParseOpcode(input string, lineNum int) (shared.Opcode, error) {
	var opcode shared.Opcode
	var returnError error = nil

	opcode = shared.InstructionSet[input]

	if opcode == shared.INVALID {
		returnError = fmt.Errorf("Invalid instruction at Line %d: %s", lineNum+1, input)
	}

	return opcode, returnError
}
