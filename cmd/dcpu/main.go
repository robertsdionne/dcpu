package main

import (
	"io/ioutil"
	"log"

	"github.com/robertsdionne/dcpu"
	"github.com/robertsdionne/dcpu/assembler"
	"github.com/robertsdionne/dcpu/clock"
)

func main() {
	source, err := ioutil.ReadFile("programs/dcpu.dasm")
	if err != nil {
		log.Fatalln(err)
	}

	cpu := dcpu.DCPU{}
	clock := clock.Clock{}

	cpu.Hardware = append(cpu.Hardware, &clock)
	cpu.Load(0, assembler.Assemble(string(source)))

	cpu.Execute()
}
