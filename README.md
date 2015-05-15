# mips-emulator
A MIPS emulator compatible with the University of Waterloo dialect of MIPS

This emulates MIPS with some features not available in the CS241 provided emulator.

## Features above-and beyond the standard MIPS emulator
1. In `-verbose` mode, each decoded instruction is output along with the registers it's acting upon
2. Complains when you do things that may be a bad idea (e.g. letting your PC go past the loaded program)
3. Registers are passed in as flags
4. Warns on use of uninitialized memory or registers

## Features missing
1. Compatibility can't be guaranteed with the MIPS standard, this is an early project
2. ???

Pull requests are welcomed.

## Usage
	$ ./mips-emulator -filename="../mips-assembler/a1p3.out" -verbose -r1 5
	0x00000000: add $3 (0x00000000 -> 0x00000005), $0 (0x00000000), $1 (0x00000005)
	0x00000004: add $4 (0x00000000 -> 0x0000000a), $1 (0x00000005), $3 (0x00000005)
	0x00000008: jr $31 (0xfde8)
	
	machine dump: program finished cleanly
	$ 0: 0x00000000, $ 1: 0x00000005, $ 2: 0x00000000, $ 3: 0x00000005, 
	$ 4: 0x0000000a, $ 5: 0x00000000, $ 6: 0x00000000, $ 7: 0x00000000, 
	$ 8: 0x00000000, $ 9: 0x00000000, $10: 0x00000000, $11: 0x00000000, 
	$12: 0x00000000, $13: 0x00000000, $14: 0x00000000, $15: 0x00000000, 
	$16: 0x00000000, $17: 0x00000000, $18: 0x00000000, $19: 0x00000000, 
	$20: 0x00000000, $21: 0x00000000, $22: 0x00000000, $23: 0x00000000, 
	$24: 0x00000000, $25: 0x00000000, $26: 0x00000000, $27: 0x00000000, 
	$28: 0x00000000, $29: 0x00000000, $30: 0x00000000, $31: 0x0000fde8, 

	 pc: 0x0000fde8,  lo: 0x00000000,  hi: 0x00000000

