package main

import (
	"flag"
	"fmt"
	"nyantime/decoder"
	"nyantime/runtime"
)

func main() {
	var programPath string
	flag.StringVar(&programPath, "program", "", "the assembled file to run")
	flag.Parse()

	if programPath == "" {
		fmt.Println("No input program specified!")
		return
	}

	contents, err := OpenFile(programPath)

	if err != nil {
		panic(err)
	}

	decoded := decoder.DecodeInstructions(contents)

	err = runtime.Runtime(decoded)

	if err != nil {
		panic(err)
	}
}
