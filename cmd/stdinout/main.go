package main

import (
	"io/ioutil"
	"log"

	"github.com/robertsdionne/dcpu"
	"github.com/robertsdionne/dcpu/assembler"
	"github.com/robertsdionne/dcpu/stdin"
	"github.com/robertsdionne/dcpu/stdout"
)

func main() {
	source, err := ioutil.ReadFile("programs/stdinout.dasm")
	if err != nil {
		log.Fatalln(err)
	}

	cpu := dcpu.DCPU{}
	in := stdin.Device{}
	out := stdout.Device{}

	cpu.Hardware = append(cpu.Hardware, &in, &out)
	cpu.Load(0, assembler.Assemble(string(source)))

	cpu.Execute()
}
