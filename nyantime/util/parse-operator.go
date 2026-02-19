package util

import (
	"strconv"
)

func ParseOperator(operator Operator, program []Instruction) (int, error) {
	numberBase := -1

	switch operator[0] {
	case '#':
		numberBase = 10
	case 'B':
		numberBase = 2
	case '&':
		numberBase = 16
	}

	// operator is a label
	if numberBase == -1 {
		parsedOperator, err := strconv.Atoi(string(operator))

		if err != nil {
			return 0, err
		}

		return ParseOperator(program[parsedOperator].Operator, program)
	}

	conv, err := strconv.ParseInt(string(operator)[1:], numberBase, 0)

	if err != nil {
		return 0, err
	}

	return int(conv), nil
}
