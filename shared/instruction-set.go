package shared

type Opcode uint8
type Operand int32

type OperandType uint8

const (
	Register         OperandType = iota // 00
	RegisterPointer                     // 01
	Immediate                           // 10
	ImmediatePointer                    // 11
)

type Instruction struct {
	SourceType OperandType
	DestType   OperandType
	Opcode     Opcode
	Source     Operand
	Dest       Operand
}

func ModifyInstruction(instructions *[]Instruction, index int, newInstruction Instruction) {
	(*instructions)[index] = newInstruction
}

func (inst *Instruction) ModifyOperator(newOperator string) {
	// inst.Source = Operator(newOperator)
}

const (
	INVALID Opcode = iota

	Mov

	Add
	Sub
	Mul
	Div
	Mod

	And
	Not
	Xor
	Or
	Shr
	Shl

	Jmp
	Cmp
	Je
	Jne
	Jz
	Jnz
	Jge
	Jg
	Jle
	Jl

	In
	Out
	Outd
	End
	Call
	Ret

	Dn
)

type OpcodeInfo struct {
	Opcode     Opcode
	NumOperand uint8
}

var InstructionSet = map[string]OpcodeInfo{
	"mov": {
		Opcode:     Mov,
		NumOperand: 2,
	},
	"add": {
		Opcode:     Add,
		NumOperand: 2,
	},
	"sub": {
		Opcode:     Sub,
		NumOperand: 2,
	},
	"mul": {
		Opcode:     Mul,
		NumOperand: 2,
	},
	"div": {
		Opcode:     Div,
		NumOperand: 2,
	},
	"mod": {
		Opcode:     Mod,
		NumOperand: 2,
	},
	"and": {
		Opcode:     And,
		NumOperand: 2,
	},
	"not": {
		Opcode:     Not,
		NumOperand: 1,
	},
	"xor": {
		Opcode:     Xor,
		NumOperand: 2,
	},
	"or": {
		Opcode:     Or,
		NumOperand: 2,
	},
	"shr": {
		Opcode:     Shr,
		NumOperand: 2,
	},
	"shl": {
		Opcode:     Shl,
		NumOperand: 2,
	},
	"jmp": {
		Opcode:     Jmp,
		NumOperand: 1,
	},
	"cmp": {
		Opcode:     Cmp,
		NumOperand: 2,
	},
	"je": {
		Opcode:     Je,
		NumOperand: 1,
	},
	"jne": {
		Opcode:     Jne,
		NumOperand: 1,
	},
	"jz": {
		Opcode:     Jz,
		NumOperand: 1,
	},
	"jnz": {
		Opcode:     Jnz,
		NumOperand: 1,
	},
	"jge": {
		Opcode:     Jge,
		NumOperand: 1,
	},
	"jg": {
		Opcode:     Jg,
		NumOperand: 1,
	},
	"jle": {
		Opcode:     Jle,
		NumOperand: 1,
	},
	"jl": {
		Opcode:     Jl,
		NumOperand: 1,
	},
	"in": {
		Opcode:     In,
		NumOperand: 1,
	},
	"out": {
		Opcode:     Out,
		NumOperand: 1,
	},
	"outd": {
		Opcode:     Outd,
		NumOperand: 1,
	},
	"end": {
		Opcode:     End,
		NumOperand: 0,
	},
	"call": {
		Opcode:     Call,
		NumOperand: 1,
	},
	"ret": {
		Opcode:     Ret,
		NumOperand: 0,
	},
	"dn": {
		Opcode:     Dn,
		NumOperand: 1,
	},
}

// list of instructions that dont use an operand
var NoOperand = []Opcode{End, Ret}

// list of instructions that take 1 operand
var OneOperand = []Opcode{Not, Jmp, Je, Jne, Jz, Jnz, Jge, Jg, Jle, Jl, In, Out, Outd, Call, Dn}

// list of instructions that take 2 operands
var TwoOperand = []Opcode{Mov, Add, Sub, Mul, Div, Mod, And, Xor, Or, Shr, Shl, Cmp}

type RegisterType uint8

const (
	r0 RegisterType = iota + 1
	r1
	r2
	r3
	r4
	r5
	r6
	r7

	idx

	pc
)

var RegisterMap = map[string]RegisterType{
	"%r0": r0,
	"%r1": r1,
	"%r2": r2,
	"%r3": r3,
	"%r4": r4,
	"%r5": r5,
	"%r6": r6,
	"%r7": r7,

	"%idx": idx,

	"%pc": pc,
}
