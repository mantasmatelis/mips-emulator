package main

type InstructionMask struct {
  mask uint32
  val uint32
}
func (mask InstructionMask) Matches(inst uint32) bool {
  return inst & mask.mask == mask.val
}
 
type Instruction func(*Machine, uint32)
type InstructionSet map[InstructionMask]Instruction

var InstructionMapping InstructionSet = InstructionSet{
        InstructionMask{0xFC0007FF, 0x00000020}: (*Machine).add,
        InstructionMask{0xFC0007FF, 0x00000022}: (*Machine).sub,
        InstructionMask{0xFC00FFFF, 0x00000018}: (*Machine).mult,
        InstructionMask{0xFC00FFFF, 0x00000019}: (*Machine).multu,
        InstructionMask{0xFC00FFFF, 0x0000001A}: (*Machine).div,
        InstructionMask{0xFC00FFFF, 0x0000001B}: (*Machine).divu,
        InstructionMask{0xFFFF07FF, 0x00000010}: (*Machine).mfhi,
        InstructionMask{0xFFFF07FF, 0x00000012}: (*Machine).mflo,
        InstructionMask{0xFFFF07FF, 0x00000014}: (*Machine).lis,
        InstructionMask{0xFC000000, 0x8C000000}: (*Machine).lw,
        InstructionMask{0xFC000000, 0xAC000000}: (*Machine).sw,
        InstructionMask{0xFC0007FF, 0x0000002A}: (*Machine).slt,
        InstructionMask{0xFC0007FF, 0x0000002B}: (*Machine).sltu,
        InstructionMask{0xFC000000, 0x10000000}: (*Machine).beq,
        InstructionMask{0xFC000000, 0x14000000}: (*Machine).bne,
        InstructionMask{0xFC1FFFFF, 0x00000008}: (*Machine).jr,
        InstructionMask{0xFC1FFFFF, 0x00000009}: (*Machine).jalr,
}
