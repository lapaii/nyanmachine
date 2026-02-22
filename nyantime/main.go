package main

import (
	"flag"
	"fmt"
	"nyantime/decoder"
	"nyantime/runtime"
	"nyantime/util"
)

func main() {
	var programPath string
	flag.StringVar(&programPath, "program", "../test-programs/add-3-nums.nyobj", "the assembled file to run")
	flag.Parse()

	if programPath == "" {
		fmt.Println("No input program specified!")
		return
	}

	contents, err := util.OpenFile(programPath)

	if err != nil {
		panic(err)
	}

	decoded, err := decoder.DecodeBinary(contents)

	if err != nil {
		panic(err)
	}

	err = runtime.Runtime(decoded)

	if err != nil {
		panic(err)
	}
}
