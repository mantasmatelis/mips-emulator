package main

import (
  "fmt"
)

func decode_register(inst uint32) (uint8, uint8, uint8) {
  return uint8(inst >> 21 & 0x1F), uint8(inst >> 16 & 0x1F), uint8(inst >> 11 & 0x1F)
}

func decode_immediate(inst uint32) (uint8, uint8, uint16) {
  return uint8(inst >> 21 & 0x1F), uint8(inst >> 16 & 0x1F), uint16(inst & 0xFFFF) 
}

func debug_inst(inst uint32, pc uint16, debug string, verbose bool) {
  if verbose {
    fmt.Printf("%0#8x: %0#8x which is %v\n", pc - 4, inst, debug)
  }
}

func inst_add(inst uint32, mem *[16384]uint32, registers *[34]uint32, pc *uint16, verbose bool) {
  s, t, d := decode_register(inst)
  debug_inst(inst, *pc, fmt.Sprintf("add $%v, $%v, $%v", s, t, d), verbose)
  registers[d] = registers[t] + registers[s]
}

func inst_sub(inst uint32, mem *[16384]uint32, registers *[34]uint32, pc *uint16, verbose bool) {
  s, t, d := decode_register(inst)
  _ = s + t + d
  //TODO: right order?
  registers[d] = registers[t] - registers[s]
}

func inst_mult(inst uint32, mem *[16384]uint32, registers *[34]uint32, pc *uint16, verbose bool) {
  s, t, _ := decode_register(inst)
  result := int64(registers[s]) * int64(registers[t]) 
  registers[32] = uint32(result)
  registers[33] = uint32(result >> 32)
}

func inst_multu(inst uint32, mem *[16384]uint32, registers *[34]uint32, pc *uint16, verbose bool) {
  s, t, _ := decode_register(inst)
  result := uint64(registers[s]) * uint64(registers[t]) 
  registers[32] = uint32(result)
  registers[33] = uint32(result >> 32)
  
}

func inst_div(inst uint32, mem *[16384]uint32, registers *[34]uint32, pc *uint16, verbose bool) {
  s, t, d := decode_register(inst)
  _ = s + t + d
}

func inst_divu(inst uint32, mem *[16384]uint32, registers *[34]uint32, pc *uint16, verbose bool) {
  s, t, d := decode_register(inst)
  _ = s + t + d
}

func inst_mfhi(inst uint32, mem *[16384]uint32, registers *[34]uint32, pc *uint16, verbose bool) {
  _, _, d := decode_register(inst)
  registers[d] = registers[33]
}

func inst_mflo(inst uint32, mem *[16384]uint32, registers *[34]uint32, pc *uint16, verbose bool) {
  _, _, d := decode_register(inst)
  registers[d] = registers[32]
}

func inst_lis(inst uint32, mem *[16384]uint32, registers *[34]uint32, pc *uint16, verbose bool) {
  _, _, d := decode_register(inst)
  registers[d] = mem[*pc]
  *pc += uint16(4)
}

func inst_lw(inst uint32, mem *[16384]uint32, registers *[34]uint32, pc *uint16, verbose bool) {
  s, t, i := decode_immediate(inst)
  registers[t] = mem[uint16(s) + i]
}

func inst_sw(inst uint32, mem *[16384]uint32, registers *[34]uint32, pc *uint16, verbose bool) {
  s, t, i := decode_immediate(inst)
  mem[uint16(s) + i] = registers[t]
}

func inst_slt(inst uint32, mem *[16384]uint32, registers *[34]uint32, pc *uint16, verbose bool) {
  s, t, d := decode_register(inst)
  _ = s + t + d
}

func inst_sltu(inst uint32, mem *[16384]uint32, registers *[34]uint32, pc *uint16, verbose bool) {
  s, t, d := decode_register(inst)
  _ = s + t + d
}

func inst_beq(inst uint32, mem *[16384]uint32, registers *[34]uint32, pc *uint16, verbose bool) {
  s, t, i := decode_immediate(inst)
  if registers[s] == registers[t] {
    *pc += i * 4
  }
}

func inst_bne(inst uint32, mem *[16384]uint32, registers *[34]uint32, pc *uint16, verbose bool) {
  s, t, i := decode_immediate(inst)
  if registers[s] != registers[t] {
    *pc += i * 4
  }
}

func inst_jr(inst uint32, mem *[16384]uint32, registers *[34]uint32, pc *uint16, verbose bool) {
 s, _, _ := decode_register(inst)
 *pc = uint16(registers[s])
}

func inst_jalr(inst uint32, mem *[16384]uint32, registers *[34]uint32, pc *uint16, verbose bool) {
  s, _, _ := decode_register(inst)
  registers[31] = uint32(*pc)
  *pc = uint16(registers[s])
}
