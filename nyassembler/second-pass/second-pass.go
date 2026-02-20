package secondpass

import (
	"fmt"
	firstpass "nyassembler/first-pass"
	"regexp"
	"shared"
	"slices"
	"strconv"
	"strings"
)

// function is rather "messy", should get cleaned up some time
func SecondPass(contents []string, symbolTable firstpass.SymbolTable) ([]shared.Instruction, error) {
	var outputProgram []shared.Instruction

	numberOperator := regexp.MustCompile(`(#[0-9]+)|(B[0-1]+)|(\^[0-9A-F]+)`)
	rawNumberOperator := regexp.MustCompile(`([0-9]+)`)

	for idx, line := range contents {
		// stuff this step does
		// check if opcode is valid, if not: throw error and stop
		// check if operator is a label, if so: replace with value from symbol table

		// going to keep storing operator as a string, this is so that i can
		// encode numbers as #n/Bn/&n and addresses as just `n`

		// if the operator isn't compatible with the opcode, throw an error and stop

		// output

		lineParts := strings.Fields(line)

		// operator will always be 1 more than opcode
		// setting to 0 so that i can tell when something goes wrong
		var opcodeIndex = -1

		switch len(lineParts) {
		// label definition lines
		case 3:
			opcodeIndex = 1
		// opcode operator or just opcode lines
		case 2, 1:
			opcodeIndex = 0
		}

		// something has gone wrong, throw error!!
		if opcodeIndex == -1 {
			return []shared.Instruction{}, fmt.Errorf("something went wrong on line %d!! %s", idx, line)
		}

		opcode, err := ParseOpcode(lineParts[opcodeIndex], idx)

		if err != nil {
			return []shared.Instruction{}, err
		}

		// opcode DOESNT get an operator
		if slices.Contains(shared.NoOperator, opcode) {
			// no operator in the line
			if opcodeIndex+1 > len(lineParts)-1 {
				outputProgram = append(outputProgram, shared.Instruction{
					Opcode:   opcode,
					Operator: "",
				})

				continue
			}

			return []shared.Instruction{}, fmt.Errorf("operator with opcode that doesn't support an operator! on line %d %s", idx, line)
		}

		// dealt with opcodes without operators, now just throw an error if there isn't one
		if opcodeIndex+1 > len(lineParts)-1 {
			return []shared.Instruction{}, fmt.Errorf("no operator found on line %d! %s", idx, line)
		}

		operator := lineParts[opcodeIndex+1]

		// opcode requires a register for the operator (ACC, IDX or PC)
		if slices.Contains(shared.RegisterOperator, opcode) {
			// is it a register?
			if operator != "ACC" && operator != "IDX" && operator != "PC" {
				return []shared.Instruction{}, fmt.Errorf("operator on line %d isn't a register! %s", idx, line)
			}

			outputProgram = append(outputProgram, shared.Instruction{
				Opcode:   opcode,
				Operator: shared.Operator(operator),
			})

			continue
		}

		// now we are dealing with only instructions
		// that use numbers / addresses (or labels to represent either of those)

		// instruction can only take defined numbers
		if slices.Contains(shared.NumberOperator, opcode) {
			// this operator is a number
			if numberOperator.MatchString(operator) {
				outputProgram = append(outputProgram, shared.Instruction{
					Opcode:   opcode,
					Operator: shared.Operator(operator),
				})

				continue
			}

			return []shared.Instruction{}, fmt.Errorf("operator on line %d isn't a defined number! %s", idx, line)
		}

		// instruction can only take addresses
		if slices.Contains(shared.AddressOperator, opcode) {
			// operator is a defined number, not allowed
			if numberOperator.MatchString(operator) {
				return []shared.Instruction{}, fmt.Errorf("operator on line %d isn't an address! %s", idx, line)
			}

			// operator is a "raw" address, i.e. LDD 5 <- load operator at address 5
			if rawNumberOperator.MatchString(operator) {
				// just add to output
				outputProgram = append(outputProgram, shared.Instruction{
					Opcode:   opcode,
					Operator: shared.Operator(operator),
				})

				continue
			}

			// operator must be a label now
			symbolTableLine := symbolTable[operator]

			// operator isnt in symbol table
			if symbolTableLine == 0 {
				return []shared.Instruction{}, fmt.Errorf("%s is used as a symbol but isn't defined at all! line %d", operator, idx)
			}

			// label is properly defined and used
			outputProgram = append(outputProgram, shared.Instruction{
				Opcode:   opcode,
				Operator: shared.Operator(strconv.Itoa(symbolTableLine - 1)),
			})

			continue
		}

		// instruction can take either user defined number or an address
		if numberOperator.MatchString(operator) {
			outputProgram = append(outputProgram, shared.Instruction{
				Opcode:   opcode,
				Operator: shared.Operator(operator),
			})

			continue
		}

		// not a defined number, at this point it must be an address
		// operator is a "raw" address
		if rawNumberOperator.MatchString(operator) {
			// just add to output
			outputProgram = append(outputProgram, shared.Instruction{
				Opcode:   opcode,
				Operator: shared.Operator(operator),
			})

			continue
		}

		// must be a label now
		symbolTableLine := symbolTable[operator]

		// operator isnt in symbol table
		if symbolTableLine == 0 {
			return []shared.Instruction{}, fmt.Errorf("%s is used as a symbol but isn't defined at all! line %d", operator, idx)
		}

		// label is properly defined and used
		outputProgram = append(outputProgram, shared.Instruction{
			Opcode:   opcode,
			Operator: shared.Operator(strconv.Itoa(symbolTableLine - 1)),
		})

		continue
	}

	return outputProgram, nil
}
