package registers

type Registers struct {
	ProgramCounter int
	Accumulator    int
	Index          int
}

func (register *Registers) SetAccumulator(value int) {
	register.Accumulator = value
}

func (register *Registers) AddToAccumulator(value int) {
	register.Accumulator = register.Accumulator + value
}

func (register *Registers) GetAccumulator() int {
	return register.Accumulator
}

func (register *Registers) IncrementPC() {
	register.ProgramCounter = register.ProgramCounter + 1
}
