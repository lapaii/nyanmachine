package parser

import (
	"fmt"
	"nyanpreter/instructions"
	"regexp"
	"slices"
	"strings"
)

func Parse(contents []string) ([]instructions.Instruction, error) {
	// go through each line
	// clean any comments, `//` is comment sign
	// split into operand / operator
	// check if operand is legit and convert it into the enum type
	// check if operator exists for the instructions that need it
	// return instruction array

	// regex to match all characters after a // for removing comments
	commentRegex := regexp.MustCompile(`(\/\/.*)`)

	parsedInstructions := []instructions.Instruction{}

	for idx, line := range contents {
		// remove comments, clear any whitespace
		cleansedLine := commentRegex.ReplaceAllString(line, "")
		cleansedLine = strings.TrimSpace(cleansedLine)

		// if line is blank, skip
		if cleansedLine == "" {
			continue
		}

		parts := strings.Split(cleansedLine, " ")

		// parse operand
		parsedOperand, err := ParseOperand(idx, parts[0])

		if err != nil {
			return []instructions.Instruction{}, err
		}

		// have to check here if this line has an
		// operator or not and check if it should have one
		var parsedOperator instructions.Operator = ""
		if len(parts) == 2 {
			// parse operator
			parsedOperator, err = ParseOperator(idx, parsedOperand, parts[1])

			if err != nil {
				return []instructions.Instruction{}, err
			}
		} else if len(parts) == 1 {
			// if the instruction doesnt use an operator
			if !slices.Contains(instructions.NoOperator, parsedOperand) {
				return []instructions.Instruction{}, fmt.Errorf("instruction %d on line %d requires an operator!\n", parsedOperand, idx+1)
			}
		} else {
			return []instructions.Instruction{}, fmt.Errorf("invalid syntax on line %d", idx+1)
		}

		// add parsed instruction to slice
		parsedInstructions = append(parsedInstructions, instructions.Instruction{
			Operand:  parsedOperand,
			Operator: parsedOperator,
		})
	}

	// if last instruction is not an END, add it
	if parsedInstructions[len(parsedInstructions)-1].Operand != instructions.END {
		parsedInstructions = append(parsedInstructions, instructions.Instruction{
			Operand:  instructions.END,
			Operator: "",
		})
	}

	return parsedInstructions, nil
}
