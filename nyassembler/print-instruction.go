package main

import (
	"fmt"
	"shared"
)

func reverseMap(m map[string]shared.Opcode) map[shared.Opcode]string {
	n := make(map[shared.Opcode]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

func PrintInstruction(inst shared.Instruction, idx int) {
	reversedMap := reverseMap(shared.InstructionSet)

	fmt.Printf("%d: %s %s\n", idx, reversedMap[inst.Opcode], inst.Operator)
}
