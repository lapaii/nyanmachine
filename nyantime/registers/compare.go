package registers

func (r *Registers) GetCompareResult() bool {
	return r.compareResult
}

func (r *Registers) SetCompareResult(value bool) {
	r.compareResult = value
}
