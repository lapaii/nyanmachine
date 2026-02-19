package registers

type Registers struct {
	programCounter int
	accumulator    int
	indexRegister  int
	compareResult  bool
}

func NewRegisters() Registers {
	return Registers{
		programCounter: 0,
		accumulator:    0,
		indexRegister:  0,
		compareResult:  false,
	}
}
