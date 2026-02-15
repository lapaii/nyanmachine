package main

import (
	"os"
	"strings"
)

func OpenFile(filePath string) ([]string, error) {
	contentBytes, err := os.ReadFile(filePath)

	if err != nil {
		return []string{}, err
	}

	contentString := string(contentBytes)
	lineSplit := strings.Split(contentString, "\n")

	return lineSplit, nil
}
