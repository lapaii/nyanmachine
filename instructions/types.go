package instructions

type Operand int
type Operator string

type Instruction struct {
	Operand  Operand
	Operator Operator
}
