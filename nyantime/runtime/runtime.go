package runtime

import (
	"fmt"
	"nyantime/registers"
	"shared"
)

func Runtime(decodedProgram []shared.Instruction) error {
	registers := registers.NewRegisters()
	shouldContinue := true

	for shouldContinue {
		currentInstruction := decodedProgram[registers.GetPC()]

		fmt.Println(currentInstruction)

		if currentInstruction.Opcode == shared.End {
			shouldContinue = false
			continue
		}

		err := FunctionMap[currentInstruction.Opcode](&registers, currentInstruction, &decodedProgram)

		if err != nil {
			return fmt.Errorf("failed on instruction: %+v\nerror: %s", currentInstruction, err.Error())
		}

		registers.IncrementPC(1)
	}

	return nil
}
