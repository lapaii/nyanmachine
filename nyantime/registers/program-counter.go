package registers

func (r *Registers) GetPC() int {
	return r.programCounter
}

func (r *Registers) SetPC(value int) {
	r.programCounter = value
}

func (r *Registers) IncrementPC(value int) {
	r.programCounter = r.programCounter + value
}

func (r *Registers) DecrementPC(value int) {
	r.programCounter = r.programCounter - value
}
