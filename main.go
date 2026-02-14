package main

import (
	"flag"
	"nyanpreter/interpreter"
	"nyanpreter/parser"
)

func main() {
	// gets the input file from flags
	var filePath string
	flag.StringVar(&filePath, "input", "input.nyan", "the source file to run")

	flag.Parse()

	fileContents, err := parser.LoadFile(filePath)

	if err != nil {
		panic(err)
	}

	program, err := parser.Parse(fileContents)

	if err != nil {
		panic(err)
	}

	interpreter.Interpreter(program)
}
