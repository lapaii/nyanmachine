package main

import "os"

func WriteFile(outputPath string, data []byte) error {
	outputData := []byte("nya:3") // magic byte thing
	outputData = append(outputData, data...)

	err := os.WriteFile(outputPath, outputData, 0644)

	return err
}
