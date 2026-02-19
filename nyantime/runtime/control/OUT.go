package control

import (
	"fmt"
	"nyantime/registers"
	"nyantime/util"
)

func OUT(r *registers.Registers, operator util.Operator, program *[]util.Instruction) error {
	fmt.Println(r.Accumulator)

	return nil
}
