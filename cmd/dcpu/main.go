package main

import (
	"github.com/robertsdionne/dcpu"
	"github.com/robertsdionne/dcpu/clock"
	"github.com/robertsdionne/dcpu/parser"
)

func main() {
	cpu := dcpu.DCPU{}
	clock := clock.Clock{}

	cpu.Hardware = append(cpu.Hardware, &clock)

	cpu.Load(0, parser.Assemble(`
		:main
			ias handleInterrupt
			set a, 0
			set b, 120
			hwi 0
			set a, 2
			set b, 1
			hwi 0
			dum
			sub pc, 1

		:handleInterrupt
			dum
			set a, 1
			hwi 0
			rfi 0
	`))

	cpu.Execute()
}
