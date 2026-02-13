# nyanpreter

the nyanpreter is a simple interpreter for the assembly language provided as part of the [Cambridge International AS & A Level Computer Science syllabus](https://www.cambridgeinternational.org/Images/636089-2024-2025-syllabus.pdf).

the nyanpreter only provides one general-purpose register (the accumulator). this is because in exam scenarios, the accumulator is the only available register.

## usage

```bash
go build .
./nyanpreter --input input.nyan
```

## instruction set

Currently implemented instructions:

| operand | operator | explanation                                                  |
|---------|----------|--------------------------------------------------------------|
| LDM     | #n       | Load n into the accumulator                                  |
| ADD     | #n/Bn/&n | Add n to the accumulator                                     |
| OUT     |          | Output the accumulator value to stdout as an ASCII character |

- `#` denotes a denary (decimal) number
- `B` denotes a binary number
- `&` denotes a hexadecimal number

## future plans

- [ ] implement all instructions as specified in the syllabus
- [ ] implement support for labels, for use with labeling data and instructions
- [ ] possibly allow user to "build" program to an intermediate pseudo-machine code file that can then be run by the interpreter
