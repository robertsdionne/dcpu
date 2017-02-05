package main

import (
	"log"

	"github.com/robertsdionne/dcpu"
	"github.com/robertsdionne/dcpu/assembler"
	"github.com/robertsdionne/dcpu/hardware"
	"github.com/robertsdionne/dcpu/sped3"
)

func main() {
	program, err := assembler.AssembleFile("programs/hologram.dasm")
	if err != nil {
		log.Fatalln(err)
	}

	d := dcpu.DCPU{}
	s := sped3.Device{TargetRotation: 90}

	d.Hardware = append(d.Hardware, &s)
	d.Load(0, program)

	go d.Execute()

	loop := hardware.Loop{
		SPED3: &s,
	}
	loop.Run()
}
