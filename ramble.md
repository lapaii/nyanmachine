# my rambles

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
- this is where any errors in the syntax of the code will be caught and reported
- ignore label definitions, just translate opcodes and operand definitions into binary
- the memory layout will be created in this step, an instruction will take up 1 unit of this memory, as a struct like:

```go
type Instruction struct {
  Operand int // will be later decoded back into instructions that are useful
  Operator int // Numbr .
}
```

- output binary to file

assembly done!! yayy

running the code:

- basically how i am now, read in file according to the spec i set for the binary file
- run it lol
