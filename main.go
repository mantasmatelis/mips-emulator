package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"os"
	"strings"
)

var opts struct {
	Verbose     []bool            `short:"v" long:"verbose" description:"Show verbose debug information"`
	Breakpoints bool              `short:"b" long:"breakpoints" description:"Whether to interpret 0xFEFEFEFE instructions as instructions to dump information and continue."`
	Files       map[string]string `short:"f" long:"file" description:"Filenames to load into memory"`
	Memory      map[string]string `short:"m" long:"memory" description:"Direct memory to load"`
	Registers   map[uint8]string  `short:"r" long:"register" description:"Registers to set"`
}

func main() {
	_, err := flags.Parse(&opts)

	if err != nil {
		os.Exit(1)
	}

	registers := make([]uint32, 32)
	for k, v := range opts.Registers {
		_, err := fmt.Sscan(v, &registers[k])
		if err != nil {
			fmt.Printf("register %v with value %v cannot be converted to uint32\n", k, v) 
			os.Exit(1)
		}
	}

	memory := make(map[uint32][]uint32)

	for k, v := range opts.Files {
		var addr uint32
		_, err := fmt.Sscan(k, &addr)
		if err != nil {
			fmt.Println("argument file has invalid address: ", err)
		}
		rawData, err := ioutil.ReadFile(v)
		if err != nil {
			fmt.Println("could not read file: ", err)
			os.Exit(1)
		}
		data := make([]uint32, (len(rawData)+3)/4)
		for k, v := range rawData {
			data[k/4] += uint32(v) << uint(8*(3-(k%4)))
		}
		memory[addr] = data
	}

	for k, v := range opts.Memory {
		var addr uint32
		_, err := fmt.Sscan(k, &addr)
		if err != nil {
			fmt.Println("argument memory has invalid address: ", err)
		}
		strData := strings.Split(v, ",")
		data := make([]uint32, len(strData))
		for k, v := range strData {
			_, err := fmt.Sscan(v, &data[k])
			if err != nil {
				fmt.Println("argument memory is not comma-delimited integers: ", err)
				os.Exit(1)
			}
		}
		memory[addr] = data
	}

	var m Machine
	m.LoadProgram(memory, registers, len(opts.Verbose) > 0, opts.Breakpoints)
	m.Run()
}
