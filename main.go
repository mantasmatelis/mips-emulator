package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	var filename = flag.String("filename", "", "the mips program to run")
	var verbose = flag.Bool("verbose", false, "whether or not all decoded instructions are outputed")
	var breakpoints = flag.Bool("breakpoints", false, "whether to interpret an 0xFEFEFEFE instruction as a debug dump") 

	registerFlags := make([]*uint64, 31)
	for i := 1; i < 32; i++ {
		registerFlags[i-1] = flag.Uint64(fmt.Sprintf("r%v", i), 0, fmt.Sprintf("the initial value of register %v", i))
	}

	flag.Parse()

	registers := make([]uint32, 31)
	for i := 0; i < 31; i++ {
		registers[i] = uint32(*registerFlags[i])
	}

	var m Machine

	fileMem, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Println("could not read file: ", err)
		return
	}

	m.LoadProgram(fileMem, registers, *verbose, *breakpoints)
	m.Run()
}
