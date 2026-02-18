package firstpass

import (
	"nyassembler/util"
)

func FirstPass(input []string) (SymbolTable, error) {
	symbolTable := SymbolTable{}

	for idx, line := range input {
		labelMatches := util.LabelRegex.FindStringSubmatch(line)

		if len(labelMatches) == 0 {
			continue
		}

		// adding 1 so that i can tell when a label
		// is defined on the first line vs not at all
		symbolTable[labelMatches[1]] = idx + 1
	}

	return symbolTable, nil
}
