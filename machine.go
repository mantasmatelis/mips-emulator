package main

import (
	"fmt"
	"os"
)

type Machine struct {
        mem       [16384]uint32
        registers [32]uint32
        lo        uint32
        hi        uint32
        pc        uint16

        verbose         bool
        breakpoints     bool
        returnAddress   uint16

        memSet       [16384]bool
        registersSet [32]bool
        loHiSet      bool
}

func (m *Machine) GetMem(loc uint16) uint32 {
        if !m.memSet[loc/4] {
		m.DebugDump(fmt.Sprintf("read from memory location %0#8x before writing to it", loc))
        }
        return m.mem[loc/4]
}


func (m *Machine) SetMem(loc uint16, val uint32) {
        m.mem[loc/4] = val
        m.memSet[loc/4] = true
}

func (m *Machine) GetReg(register uint8) uint32 {
	if !m.registersSet[register] {
		m.DebugDump(fmt.Sprintf("read from register $%2v before writing to it", register))
	}
        return m.registers[register]
}

func (m *Machine) SetReg(register uint8, val uint32) {
        m.registers[register] = val
        m.registersSet[register] = true
}

func (m *Machine) SetLoHi(lo uint32, hi uint32) {
        m.lo = lo
        m.hi = hi
        m.loHiSet = true
}

func (m *Machine) GetLo() uint32 {
        return m.lo
}

func (m *Machine) GetHi() uint32 {
        return m.hi
}

func (m *Machine) LoadProgram(program []byte, registers []uint32, verbose bool, breakpoints bool) {
        for i := uint(0); int(i) < len(program); i++ {
                m.mem[i/4] += uint32(program[i]) << (8 * (3 - (i % 4)))
		m.memSet[i/4] = true
        }

        for i := 0; i < len(registers); i++ {
                m.registers[i+1] = registers[i]
		m.registersSet[i+1] = true
        }
	m.registersSet[0] = true

        m.returnAddress = 65000
	m.registers[31] = uint32(m.returnAddress)
        m.verbose = verbose
        m.breakpoints = breakpoints
}

func (m *Machine) DebugDump(msg string) {
	fmt.Fprintf(os.Stderr, "\nmachine dump: %v\n", msg) 
	for i, val := range m.registers {
		fmt.Fprintf(os.Stderr, "$%2v: %0#8x, ", i, val)	
		if i % 4 == 3 {
			fmt.Fprintf(os.Stderr, "\n")
		}
	}
	fmt.Fprintf(os.Stderr, "\n pc:     %0#4x,  lo: %0#8x,  hi: %0#8x\n", m.pc, m.lo, m.hi)
}

func (m *Machine) Run() {
        instructionLoop:
        for {
                if m.pc%4 != 0 {
			m.DebugDump("program counter is not word aligned, dying")
                        break
                }
                if m.pc == m.returnAddress {
			m.DebugDump("program finished cleanly")
                        break
                }

                inst := m.mem[m.pc/4]
                m.pc += 4
	
		if inst == 0xFEFEFEFE && m.breakpoints {
			m.DebugDump("breakpoint (0xFEFEFEFE) hit")
			continue
		}

                for k, v := range InstructionMapping {
                        if k.Matches(inst) {
                                v(m, inst)
                                continue instructionLoop
                        }
                }
		m.DebugDump(fmt.Sprintf("invalid instruction encountered (%0#8x), dying", inst))
                break
        }
}

