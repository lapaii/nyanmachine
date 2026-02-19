package registers

func (r *Registers) GetPC() int {
	return r.programCounter
}

func (r *Registers) SetPC(value int) {
	r.programCounter = value
}

func (r *Registers) IncrementPC() {
	r.programCounter = r.programCounter + 1
}
