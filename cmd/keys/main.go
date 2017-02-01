package main

import (
	"io/ioutil"
	"log"

	"github.com/robertsdionne/dcpu"
	"github.com/robertsdionne/dcpu/hardware"
	"github.com/robertsdionne/dcpu/keyboard"
	"github.com/robertsdionne/dcpu/parser"
)

func main() {
	source, err := ioutil.ReadFile("programs/keys.dasm")
	if err != nil {
		log.Fatalln(err)
	}

	cpu := dcpu.DCPU{}
	keys := keyboard.Device{}

	keys.Init()

	cpu.Hardware = append(cpu.Hardware, &keys)
	cpu.Load(0, parser.Assemble(string(source)))

	cpu.LoadString(0xf000, " ")

	go cpu.Execute()

	loop := hardware.Loop{
		Keyboard: &keys,
	}
	loop.Run()
}
