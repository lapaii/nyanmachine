package registers

func (r *Registers) SetIndex(value int) {
	r.indexRegister = value
}

func (r *Registers) GetIndex() int {
	return r.indexRegister
}

func (r *Registers) IncrementIndex(value int) {
	r.indexRegister = r.indexRegister + value
}

func (r *Registers) DecrementIndex(value int) {
	r.indexRegister = r.indexRegister + value
}
