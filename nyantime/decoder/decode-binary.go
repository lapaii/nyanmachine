package decoder

import (
	"encoding/binary"
	"shared"
)

func DecodeBinary(file []byte) ([]shared.Instruction, error) {
	var decodedProgram []shared.Instruction
	index := 0

	for index < len(file) {
		thisInstruction := shared.Instruction{}

		// read next 2 bytes from file
		headerBytes := file[index : index+2]
		index += 2 // increment index by two bytes

		headerNum := binary.BigEndian.Uint16(headerBytes)

		// get opcode
		opcode := shared.Opcode((headerNum >> 4)) // shr 4 gets 5-8th bits
		thisInstruction.Opcode = opcode

		// decode destType
		destType := shared.OperandType((headerNum & 14) >> 2) // anding with 1100 then shr 2 gets the 3rd and 4th bits
		thisInstruction.DestType = destType

		// get destination operand byte or bytes
		if thisInstruction.DestType == shared.Register || thisInstruction.DestType == shared.RegisterPointer {
			// 1 byte to read
			destByte := file[index : index+1][0]
			index += 1

			destVal := uint8(destByte)
			thisInstruction.Dest = shared.Operand(destVal)
		} else {
			// 4 bytes to read
			destBytes := file[index : index+4]
			index += 4

			destVal := binary.BigEndian.Uint32(destBytes)
			thisInstruction.Dest = shared.Operand(destVal)
		}

		// decode sourceType
		sourceType := shared.OperandType(headerNum & 3) // and with 11 to get the 1st and 2nd bits
		thisInstruction.SourceType = sourceType

		// get source operand byte or bytes
		if thisInstruction.SourceType == shared.Register || thisInstruction.SourceType == shared.RegisterPointer {
			// 1 byte to read
			destByte := file[index : index+1][0]
			index += 1

			destVal := uint8(destByte)
			thisInstruction.Source = shared.Operand(destVal)
		} else {
			// 4 bytes to read
			destBytes := file[index : index+4]
			index += 4

			destVal := binary.BigEndian.Uint32(destBytes)
			thisInstruction.Source = shared.Operand(destVal)
		}

		decodedProgram = append(decodedProgram, thisInstruction)
	}

	return decodedProgram, nil
}
