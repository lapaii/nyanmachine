package runtime

import (
	"nyantime/registers"
	"nyantime/runtime/memory"
	"shared"
)

var FunctionMap = map[shared.Opcode]func(*registers.Registers, shared.Instruction, *[]shared.Instruction) error{
	shared.Mov: memory.Mov,
}
