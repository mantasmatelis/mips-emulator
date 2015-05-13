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
