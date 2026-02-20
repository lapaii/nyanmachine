# my rambles

## new based swag rambles

with the new spec bringing support for multi operand instructions, i need to rethink my assembler and binary format a little bit

definition below for the binary format

instructions are encoded in 2 bytes like so:

bits 0-1: source type
bits 2-3: target type
bits 4-8: opcode
bits 9-15: currently unused

| 15 - 9 | 8 - 4  | 3 - 2       | 1 - 0       |
| ------ | ------ | ----------- | ----------- |
| unused | opcode | target type | source type |

operand types:

00 -> register (8 bits)
01 -> register pointer (8 bits) (when running the instruction, the value of register is treated as an address to do the something to)
10 -> immediate (32 bit)
11 -> immediate pointer (32 bit) (can be just user defined number or a number thats been replaced because it was a label)

## OLD RAMBLES

basically going to cmd+a delete everything

instead of being treated like an interpreter, im going to actually assemble the source code into object programs.

Why am i doing this?

1. because i want to learn more about assemblers and stuff because i think they are Cool.
2. my code is a mess and an interpreter isn't _really_ the most _ideal_ way to do something like this.
3. i want to learn more about assemblers

how it will be implemented next

there will be an additional DAT instruction (DAT #n/Bn/&n), for use with labels for defining data.

example:

```nyan
LDA letterA
OUT // will output "A"

letterA: DAT #65
```

will be a two-pass assembler

on the first pass i:

- go through each line of code
- is it a label definition?
  - yes:
    - store label name and line in a table (address symbol table)
  - no:
    - is it END instruction?
      - yes:
        - go to second pass
      - no:
        - go to next line

on the second pass:

- go through each line
- is instruction valid?
  - yes:
    - is operator valid? (some instructions dont take one or require an address etc)
      - yes:
        - is the operator a label?
          - yes:
            - using the symbol table, replace with label line and save to output
          - no:
            - save to output
      - no:
        - stop and report error to user
  - no:
    - stop and report error to user
- output the assembled program to the binary format

```go
type Instruction struct {
  Opcode int // will be later decoded back into instructions that are useful
  Operator int // Numbr .
}
```

- output binary to file

assembly done!! yayy

running the code:

- basically how i am now, read in file according to the spec i set for the binary file
- run it lol
