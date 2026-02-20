package serialiser

import (
	"shared"
)

func Serialise(input []shared.Instruction) []byte {
	// this is whats going to output a buffer with the binary format

	// 1 byte for opcode
	// as many bytes as needed for operator, followed by a null byte

	buf := []byte{}

	for _, instruction := range input {
		// write byte for opcode
		buf = append(buf, byte(instruction.Opcode))

		// write bytes for operator
		buf = append(buf, []byte(instruction.Operator)...)

		buf = append(buf, 0)
	}

	return buf
}
