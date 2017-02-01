package main

import (
	"github.com/robertsdionne/dcpu"
	"github.com/robertsdionne/dcpu/hardware"
	"github.com/robertsdionne/dcpu/keyboard"
	"github.com/robertsdionne/dcpu/parser"
)

func main() {
	cpu := dcpu.DCPU{}
	keys := keyboard.Device{}

	keys.Init()

	cpu.Hardware = append(cpu.Hardware, &keys)

	cpu.Load(0, parser.Assemble(`
		:main
			ias handleInterrupt
			set a, 3
			set b, 1
			hwi 0
			dum
			sub pc, 1

		:handleInterrupt
			set a, 1
			hwi 0
			set b, c
			set a, 2
			hwi 0
			ife c, 0
				rfi 0
			add [0xf000], 1
			set a, [0xf000]
			set [0xf000+a], b
			alt
			rfi 0
	`))

	cpu.LoadString(0xf000, " ")

	go cpu.Execute()

	loop := hardware.Loop{
		Keyboard: &keys,
	}
	loop.Run()
}
