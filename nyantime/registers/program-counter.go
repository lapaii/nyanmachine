package registers

func (r *Registers) GetPC() int32 {
	return r.programCounter
}

func (r *Registers) SetPC(value int32) {
	r.programCounter = value
}

func (r *Registers) IncrementPC(value int32) {
	r.programCounter = r.programCounter + value
}

func (r *Registers) DecrementPC(value int32) {
	r.programCounter = r.programCounter - value
}
