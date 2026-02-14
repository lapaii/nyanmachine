package interpreter

import (
	"nyanpreter/instructions"
	"nyanpreter/interpreter/registers"
)

func Interpreter(program []instructions.Instruction) {
	// init all registers (PC, ACC, IDX)
	// go through each instruction of the program
	// run the appropriate function for each operand

	registers := registers.Registers{ProgramCounter: 0, Accumulator: 0, Index: 0}
	shouldContinue := true

	for shouldContinue {
		currentInstruction := program[registers.ProgramCounter]

		// if END instruction, stop running
		if currentInstruction.Operand == instructions.END {
			shouldContinue = false
			continue
		}

		err := FunctionMap[currentInstruction.Operand](&registers, currentInstruction.Operator)

		if err != nil {
			panic(err)
		}

		registers.IncrementPC()
	}
}
