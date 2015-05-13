# mips-emulator
A MIPS emulator compatible with the University of Waterloo dialect of MIPS

This emulates MIPS with some features not available in the CS241 provided emulator.

## Features above-and beyond the standard MIPS emulator
1. In `-v`erbose mode, each decoded instruction is output
2. Complains when you do things that may be a bad idea (e.g. letting your PC go past the loaded program)
3. Registers are passed in as flags

## Features missing
1. Compatibility can't be guaranteed with the MIPS standard, this is an early project
2. Some incorrect instructions will be decoded as successful instructions

Pull requests are welcomed.

## Usage
    $ ./mips-emulator -filename="../mips-assembler/a1p3.out" -r1 1 -verbose
    0x00000000 is add $0, $1, $3
    0x00000004 is add $1, $3, $4
    0x00000008 is jr $31
    program ended cleanly.
    register 1 has value 0x1
    register 3 has value 0x1
    register 4 has value 0x2
