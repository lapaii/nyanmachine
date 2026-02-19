package control

import (
	"fmt"
	"nyantime/registers"
	"shared"
)

func OUTD(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	fmt.Printf("%d", r.GetAccumulator())

	return nil
}
