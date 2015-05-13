package main 

import (
  "io/ioutil"
  "flag"
  "fmt"
)

func main() {
  var filename = flag.String("filename", "program.mips", "the mips program to run")
  var verbose = flag.Bool("verbose", false, "whether or not all decoded instructions are outputed")
  var r1 = flag.Uint64("r1", 0, "the initial value of register 1")
  var r2 = flag.Uint64("r2", 0, "the initial value of register 2")
  var r3 = flag.Uint64("r3", 0, "the initial value of register 3")
  var r4 = flag.Uint64("r4", 0, "the initial value of register 4")
  flag.Parse()

  var mem [16384]uint32 
  var registers [34]uint32 // 32 registers + lo + hi
  var ip uint16

  registers[1] = uint32(*r1)
  registers[2] = uint32(*r2)
  registers[3] = uint32(*r3)
  registers[4] = uint32(*r4)
  registers[31] = 65000
  
  
  // load program at address 0
  fileMem, err := ioutil.ReadFile(*filename)
  if err != nil {
    fmt.Println("could not read file: ", err)
    return
  }

  for i := uint(0); int(i) < len(fileMem); i++ {
    mem[i / 4] += uint32(fileMem[i]) << (8 * (3 - (i % 4)))  
  }

  // fetch, decode, execute cycle
  for {
    if ip % 4 != 0 {
      fmt.Println("instruction pointer was not word aligned, dying.")
      break
    }

    if ip == 65000 {
      fmt.Println("program ended cleanly.")
      break
    }

    if int(ip) > len(fileMem) {
      fmt.Println("instruction pointer went past loaded program. continuing execution but unless you're writing self-modifying code this is almost certainly not what you want.")
    }

    inst := mem[ip / 4]
    ip += 4

    if inst & 0xF0000000 == 0x00000000 { // most
      if inst & 0xFF == 0x20 { // add
        inst_add(inst, &mem, &registers, &ip, *verbose)
      } else if inst & 0xFF == 0x22 { // sub
        inst_sub(inst, &mem, &registers, &ip, *verbose)
      } else if inst & 0xFF == 0x18 { // mult
        inst_mult(inst, &mem, &registers, &ip, *verbose)
      } else if inst & 0xFF == 0x19 { // multu
        inst_multu(inst, &mem, &registers, &ip, *verbose)
      } else if inst & 0xFF == 0x1A { // div
        inst_div(inst, &mem, &registers, &ip, *verbose)
      } else if inst & 0xFF == 0x1B { // divu
        inst_divu(inst, &mem, &registers, &ip, *verbose)
      } else if inst & 0xFF == 0x10 { // mfhi
        inst_mfhi(inst, &mem, &registers, &ip, *verbose)
      } else if inst & 0xFF == 0x12 { // mflo
        inst_mflo(inst, &mem, &registers, &ip, *verbose)
      } else if inst & 0xFF == 0x14 { // lis
        inst_lis(inst, &mem, &registers, &ip, *verbose)
      } else if inst & 0xFF == 0x2A { // slt
        inst_slt(inst, &mem, &registers, &ip, *verbose)
      } else if inst & 0xFF == 0x2B { // sltu
        inst_sltu(inst, &mem, &registers, &ip, *verbose)
      } else if inst & 0xFF == 0x08 { // jr
        inst_jr(inst, &mem, &registers, &ip, *verbose)
      } else if inst & 0xFF == 0x09 { // jalr
        inst_jalr(inst, &mem, &registers, &ip, *verbose)
      } else {
        fmt.Printf("cannot decode instruction %#v, dying.\n", inst)
        break
      }
    } else if inst & 0xFC000000 == 0x14000000 { // branches
      if inst & 0x04000000 == 0x00000000 { //beq
        inst_beq(inst, &mem, &registers, &ip, *verbose)
      } else { //bne
        inst_bne(inst, &mem, &registers, &ip, *verbose)
      } 
    } else if inst & 0xCC000000 == 0x8C000000 { // load/store word
      if inst & 0xFC000000 == 0x8C000000 { // lw
        inst_lw(inst, &mem, &registers, &ip, *verbose)
      } else { // sw
        inst_sw(inst, &mem, &registers, &ip, *verbose)
      }
    } else {
      fmt.Printf("cannot decode instruction %#v, dying.\n", inst)
      break
    }
  }

  for i := 0; i < 32; i++ {
    if registers[i] == 0 || (i == 31 && registers[i] == 65000) {
      continue
    }
    fmt.Printf("register %v has value %#v\n", i, registers[i])
  }
  if ip != 65000 {
    fmt.Printf("ip is %v", ip)
  }
}
