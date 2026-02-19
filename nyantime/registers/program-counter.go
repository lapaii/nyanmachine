package registers

func (r *Registers) GetPC() int {
	return r.programCounter
}

func (r *Registers) IncrementPC() {
	r.programCounter = r.programCounter + 1
}
