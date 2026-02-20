package shared

type Opcode uint8
type Operator string

type Instruction struct {
	Opcode   Opcode
	Operator Operator
}

func ModifyInstruction(instructions *[]Instruction, index int, newInstruction Instruction) {
	(*instructions)[index] = newInstruction
}

func (inst *Instruction) ModifyOperator(newOperator string) {
	inst.Operator = Operator(newOperator)
}

const (
	INVALID Opcode = iota

	// memory instructions
	LDM
	LDD
	LDI
	LDX
	LDR
	MOV
	STO

	DAT

	// maths
	ADD
	SUB
	INC
	DEC

	// control flow
	JMP
	CMP
	CMX
	CMI
	JE
	JNE
	JGE
	JLE
	IN
	OUT
	OUTD
	END

	// bitwise operations
	AND
	XOR
	OR
	LSL
	LSR
)

var InstructionSet = map[string]Opcode{
	"LDM":  LDM,
	"LDD":  LDD,
	"LDI":  LDI,
	"LDX":  LDX,
	"LDR":  LDR,
	"MOV":  MOV,
	"STO":  STO,
	"DAT":  DAT,
	"ADD":  ADD,
	"SUB":  SUB,
	"INC":  INC,
	"DEC":  DEC,
	"JMP":  JMP,
	"CMP":  CMP,
	"CMX":  CMX,
	"CMI":  CMI,
	"JE":   JE,
	"JNE":  JNE,
	"JGE":  JGE,
	"JLE":  JLE,
	"IN":   IN,
	"OUT":  OUT,
	"OUTD": OUTD,
	"END":  END,
	"AND":  AND,
	"XOR":  XOR,
	"OR":   OR,
	"LSL":  LSL,
	"LSR":  LSR,
}

// list of instructions that dont use an operator
var NoOperator = []Opcode{IN, OUT, OUTD, END}

// list of instructions which the operator needs to be a register (ACC, IDX or PC)
var RegisterOperator = []Opcode{MOV, INC, DEC}

// list of instructions that require a user defined number as the operator (#/B/&)
var NumberOperator = []Opcode{LDM, LDR, LSL, LSR}

// list of instructions that require addresses
var AddressOperator = []Opcode{LDD, LDI, LDX, STO, JMP, CMI, JE, JNE}

// all other instructions can be either defined number or address
