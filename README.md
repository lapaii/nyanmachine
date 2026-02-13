# nyanpreter

simple interpreter for the assembly language provided as part of the [Cambridge International AS & A Level Computer Science syllabus](https://www.cambridgeinternational.org/Images/636089-2024-2025-syllabus.pdf).

the interpreter only makes one general purpose register (the accumulator) available for use. this is because in exams on the course, it is assumed that the accumulator is the only register available for use.

## usage

go build .
./nyanpreter --input \<file\>

## list of instructions

this is a list of instructions currently implemented

| operand | operator | explanation                                                 |
|---------|----------|-------------------------------------------------------------|
| LDM     | #n       | load n into acc                                             |
| ADD     | #n/Bn/&n | add n to the acc                                            |
| OUT     |          | output to stdout the value of the acc as an ascii character |

\# denotes a denary number
B denotes a binary number
& denotes a hex number
