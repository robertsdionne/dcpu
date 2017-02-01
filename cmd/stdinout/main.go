package main

import (
	"github.com/robertsdionne/dcpu"
	"github.com/robertsdionne/dcpu/parser"
	"github.com/robertsdionne/dcpu/stdin"
	"github.com/robertsdionne/dcpu/stdout"
)

func main() {
	cpu := dcpu.DCPU{}
	in := stdin.Device{}
	out := stdout.Device{}

	cpu.Hardware = append(cpu.Hardware, &in, &out)

	cpu.Load(0, parser.Assemble(`
		:main
			set a, 0
			set x, 0x0400
			set y, 0x1000
			hwi 0
			set a, 0
			set x, z
			set y, 0x1000
			hwi 1
			set pc, main
	`))

	cpu.Execute()
}
