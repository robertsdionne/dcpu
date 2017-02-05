package main

import (
	"log"

	"github.com/robertsdionne/dcpu"
	"github.com/robertsdionne/dcpu/assembler"
	"github.com/robertsdionne/dcpu/hardware"
	"github.com/robertsdionne/dcpu/keyboard"
)

func main() {
	program, err := assembler.AssembleFile("programs/keys.dasm")
	if err != nil {
		log.Fatalln(err)
	}

	cpu := dcpu.DCPU{}
	keys := keyboard.Device{}

	keys.Init()

	cpu.Hardware = append(cpu.Hardware, &keys)
	cpu.Load(0, program)

	cpu.LoadString(0xf000, " ")

	go cpu.Execute()

	loop := hardware.Loop{
		Keyboard: &keys,
	}
	loop.Run()
}
