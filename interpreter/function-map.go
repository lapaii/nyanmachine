package interpreter

import (
	"nyanpreter/instructions"
	"nyanpreter/interpreter/control"
	"nyanpreter/interpreter/math"
	"nyanpreter/interpreter/memory"
	"nyanpreter/interpreter/registers"
)

var FunctionMap = map[instructions.Operand]func(*registers.Registers, instructions.Operator) error{
	instructions.LDM: memory.LDM,
	instructions.ADD: math.ADD,
	instructions.OUT: control.OUT,
	instructions.JMP: control.JMP,
}
