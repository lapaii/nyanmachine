package registers

type Flags uint32

const (
	ZF Flags = 1 << iota // 1
	SF                   // 2
	OF
)

type Registers struct {
	programCounter int32
	indexRegister  int32
	flags          Flags
	r0             int32
	r1             int32
	r2             int32
	r3             int32
	r4             int32
	r5             int32
	r6             int32
	r7             int32
}

func NewRegisters() Registers {
	return Registers{
		programCounter: 0,
		indexRegister:  0,
		flags:          0,
		r0:             0,
		r1:             0,
		r2:             0,
		r3:             0,
		r4:             0,
		r5:             0,
		r6:             0,
		r7:             0,
	}
}
