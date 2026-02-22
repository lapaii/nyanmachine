package registers

func (r *Registers) SetIndex(value int32) {
	r.indexRegister = value
}

func (r *Registers) GetIndex() int32 {
	return r.indexRegister
}

func (r *Registers) IncrementIndex(value int32) {
	r.indexRegister = r.indexRegister + value
}

func (r *Registers) DecrementIndex(value int32) {
	r.indexRegister = r.indexRegister + value
}
