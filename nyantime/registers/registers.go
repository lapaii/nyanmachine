package registers

type Flags uint8

const (
	Equal       Flags = 1 << iota // 1
	GreaterThan                   // 2
)

type Registers struct {
	programCounter int
	accumulator    int
	indexRegister  int
	flags          Flags
}

func NewRegisters() Registers {
	return Registers{
		programCounter: 0,
		accumulator:    0,
		indexRegister:  0,
		flags:          0,
	}
}
