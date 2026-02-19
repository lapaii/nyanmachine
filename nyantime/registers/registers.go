package registers

type Registers struct {
	ProgramCounter int
	Accumulator    int
	Index          int
}

func (r *Registers) IncrementPC() {
	r.ProgramCounter = r.ProgramCounter + 1
}

func (r *Registers) SetAccumulator(value int) {
	r.Accumulator = value
}

func (r *Registers) IncrementAccumulator(value int) {
	r.Accumulator = r.Accumulator + value
}
