# mips-emulator
A MIPS emulator compatible with the University of Waterloo dialect of MIPS

This emulates MIPS with some features not available in the CS241 provided emulator.

## Features above-and beyond the standard MIPS emulator
1. In `-verbose` mode, each decoded instruction is output along with the registers it's acting upon
2. Complains when you do things that may be a bad idea (e.g. letting your PC go past the loaded program)
3. Registers are passed in as flags
4. Warns on use of uninitialized memory or registers
5. Single executable with no dependencies
6. Fast to startup - it's written in Go, not Java

## Features missing
1. Compatibility can't be guaranteed with the MIPS standard, this is an early project
2. ???

Pull requests are welcomed.

## Usage
	$ ./mips-emulator -f 0:a2p6.mips -r 1:0x10000 -r 2:3 -m "0x10000:1,2,3" -v
	0x00000000: sw $30 (0x0), $2 (0x3) + 0xfffc
	0x00000004: sw $30 (0x0), $3 (0x0) + 0xfff8
	0x00000008: sw $30 (0x0), $31 (0xfde8) + 0xfff4
	0x0000000c: lis $2 (0xc)
	... snip snip snip ...
	0x00000038: lis $2 (0xc)
	0x00000040: add $30 (0xfffffff4 -> 0x0), $30 (0xfffffff4), $2 (0xc)
	0x00000044: lw $30 (0x0), $2 (0xc) + 0xfffc
	0x00000048: lw $30 (0x0), $3 (0xffff000c) + 0xfff8
	0x0000004c: lw $30 (0x0), $31 (0x24) + 0xfff4
	0x00000050: jr $31 (0xfde8)
	
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

