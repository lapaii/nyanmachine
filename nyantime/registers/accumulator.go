package registers

func (r *Registers) GetAccumulator() int {
	return r.accumulator
}

func (r *Registers) SetAccumulator(value int) {
	r.accumulator = value
}

func (r *Registers) IncrementAccumulator(value int) {
	r.accumulator = r.accumulator + value
}
