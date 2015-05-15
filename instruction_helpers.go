package main

import (
	"fmt"
	"os"
)


func parseStd(inst uint32) (uint8, uint8, uint8) {
        return uint8(inst >> 21 & 0x1F), uint8(inst >> 16 & 0x1F), uint8(inst >> 11 & 0x1F)
}

func parseSt(inst uint32) (uint8, uint8) {
        s, t, _ := parseStd(inst)
        return s, t
}

func parseD(inst uint32) uint8 {
        _, _, d := parseStd(inst)
        return d
}

func parseS(inst uint32) uint8 {
        s, _, _ := parseStd(inst)
        return s
}

func parseSti(inst uint32) (uint8, uint8, uint16) {
        return uint8(inst >> 21 & 0x1F), uint8(inst >> 16 & 0x1F), uint16(inst & 0xFFFF)
}

func (m *Machine) logStd(inst string, s, t, d uint8, result uint32) {
        m.logString(fmt.Sprintf("%v $%v (%0#x -> %0#x), $%v (%0#x), $%v (%0#x)", inst, d, m.registers[d], result, s, m.registers[s], t, m.registers[t]))
}

func (m *Machine) logSt(inst string, s, t uint8, lo, hi uint32) {
        m.logString(fmt.Sprintf("%v $%v (%0#x), $%v (%0#x). lo: %0#x, hi: %0#x", inst, s, m.registers[s], t, m.registers[t], lo, hi))
}

func (m *Machine) logD(inst string, d uint8, result uint32) {
        m.logString(fmt.Sprintf("%v $%v (%0#8x)", inst, d, result))
}

func (m *Machine) logStiWord(inst string, s, t uint8, i uint16, mem uint32) {

}

func (m *Machine) logStiBranch(inst string, s, t uint8, i uint16) {

}

func (m *Machine) logS(inst string, s uint8, addr uint16) {
        m.logString(fmt.Sprintf("%v $%v (%0#4x)", inst, s, addr))
}

func (m *Machine) logString(inst string) {
        if m.verbose {
                fmt.Fprintf(os.Stdout, "%0#8x: %v\n", m.pc - 4, inst)
        }
}

