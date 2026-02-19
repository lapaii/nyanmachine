package runtime

import (
	"nyantime/registers"
	"nyantime/runtime/control"
	"nyantime/runtime/math"
	"nyantime/runtime/memory"
	"nyantime/util"
)

var FunctionMap = map[util.Operand]func(*registers.Registers, util.Operator, *[]util.Instruction) error{
	util.LDM: memory.LDM,
	util.LDD: memory.LDD,
	// util.LDI
	util.LDX: memory.LDX,
	util.LDR: memory.LDR,
	util.MOV: memory.MOV,
	util.STO: memory.STO,

	util.ADD: math.ADD,
	util.SUB: math.SUB,
	util.INC: math.INC,
	util.DEC: math.DEC,

	util.JMP: control.JMP,
	util.CMP: control.CMP,
	// util.CMI
	util.JPE: control.JPE,
	util.JPN: control.JPN,
	// util.IN
	util.OUT: control.OUT,

	// util.AND
	// util.XOR
	// util.OR
	// util.LSL
	// util.LSR
}
