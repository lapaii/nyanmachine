package registers

func (r *Registers) SetFlag(value bool, flag Flags) {
	if value {
		r.flags = r.flags | flag // sets the bit to 1
	} else {
		r.flags = r.flags & ^flag // sets the bit to 0
	}
}

func (r *Registers) GetFlag(flag Flags) bool {
	return r.flags&flag == flag
}
