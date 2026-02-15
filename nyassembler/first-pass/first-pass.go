package firstpass

import (
	"nyassembler/util"
	"strings"
)

func FirstPass(input []string) ([]SymbolTable, error) {
	symbolTable := []SymbolTable{}

	for idx, line := range input {
		noComments := strings.TrimSpace(util.CommentsRegex.ReplaceAllString(line, ""))

		if noComments == "" {
			continue
		}

		labelMatches := util.LabelRegex.FindString(noComments)

		if labelMatches == "" {
			continue
		}

		labelMatches = labelMatches[:len(labelMatches)-1] // FindString still returns the trailing ":", this removes it

		symbolTable = append(symbolTable, SymbolTable{
			LabelName: labelMatches,
			LineNum:   idx,
		})
	}

	return symbolTable, nil
}
