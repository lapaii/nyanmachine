# nyasm spec

this is the doc for drafting up the new assembly-like language that this project will be based on moving forward.

it is mainly based of the existing language from the cambridge international 9618 cs course.

most instructions will have two operands, unlike just the one in the previous language

## syntax

the syntax for nyasm is a mix of my favourite parts of gas syntax and intel syntax

convention will be writing mnemonics in lowercase but assembler will also support uppercase

comments are done in with a single ; anywhere in a line to comment out and ignore the rest of that line

parameter order: destination before source (mnemonic, destination, source)
parameter size: all registers and stuff is 32-bit

registers are to be prefixed with a %
immediate values prefixed with $, can define decimal numbers ($95) or hex numbers ($0x5F)

## (WIP) instruction table

| opcode | operator    | explanation                                          |
| ------ | ----------- | ---------------------------------------------------- |
| mov    | src, target | move from value to r1                                |
| add    | src, target | add value to r1                                      |
| sub    | src, target | sub value to r1                                      |
| mul    | src, target | multiply value to r1                                 |
| div    | src, target | integer divide value to r1                           |
| mod    | src, target | modulus value to r1                                  |
| and    | src, target | bitwise AND value to r1                              |
| not    | src         | bitwise NOT r1                                       |
| xor    | src, target | bitwise XOR value to r1                              |
| or     | src, target | bitwise OR value to r1                               |
| shr    | src, target | logical shift r1 right by value bits                 |
| shl    | src, target | logical shift r1 left by value bits                  |
| jmp    | src         | jump to address                                      |
| cmp    | src, target | compare value with r1, setting flags based on result |
| je     | src         | jump to address if equal flag set                    |
| jne    | src         | jump to address if equal flag NOT set                |
| jz     | src         | jump to address if zero flag set                     |
| jnz    | src         | jump to address if zero flag not set                 |
| jge    | src         | jump to address if equal or greater than             |
| jg     | src         | jump to address only if greater than                 |
| jle    | src         | jump to address if equal or less than                |
| jl     | src         | jump to address only if less than                    |
| in     | src         | save ascii value of 1 char entered into stdin to r1  |
| out    | src         | output ascii value from r1 to stdout                 |
| outd   | src         | output value from r1 to stdout as a decimal number   |
| call   | src         | jump to address, while saving return address         |
| ret    |             | jump back to latest return address                   |
| dn     | number      | declare number (not _real_ instruction)              |
| ds     | string      | declare string                                       |
