package main

import (
	"flag"
	"log"

	"github.com/robertsdionne/dcpu"
	"github.com/robertsdionne/dcpu/assembler"
	"github.com/robertsdionne/dcpu/clock"
	"github.com/robertsdionne/dcpu/floppy"
	"github.com/robertsdionne/dcpu/hardware"
	"github.com/robertsdionne/dcpu/keyboard"
	"github.com/robertsdionne/dcpu/monitor"
	"github.com/robertsdionne/dcpu/printer"
	"github.com/robertsdionne/dcpu/sped3"
	"github.com/robertsdionne/dcpu/stderr"
	"github.com/robertsdionne/dcpu/stdin"
	"github.com/robertsdionne/dcpu/stdout"
)

var (
	floppyDisk = flag.String("floppy-disk", "", "The floppy file.")
	program    = flag.String("program", "/dev/stdin", "Which program to run.")
	useSPED3   = flag.Bool("sped3", false, "Use SPED3 instead of LEM1802")
)

func main() {
	flag.Parse()

	program, err := assembler.AssembleFile(*program)
	if err != nil {
		log.Fatalln(err)
	}

	cpu := dcpu.DCPU{}
	k := keyboard.Device{}
	m := monitor.Device{
		VideoAddress: 0x8000,
	}
	s := sped3.Device{TargetRotation: 90}
	c := clock.Clock{}
	i := stdin.Device{}
	o := stdout.Device{}
	e := stderr.Device{}
	f := floppy.Device{}
	p := printer.Device{}

	if *useSPED3 {
		cpu.Hardware = append(cpu.Hardware, &s)
	} else {
		cpu.Hardware = append(cpu.Hardware, &m)
	}
	cpu.Hardware = append(cpu.Hardware, &k, &c, &i, &o, &e, &f, &p)

	k.Init()
	if *floppyDisk != "" {
		f.Insert(*floppyDisk, false)
	}
	cpu.Load(0, program)

	go cpu.Execute()

	loop := hardware.Loop{
		Keyboard: &k,
	}
	if *useSPED3 {
		loop.SPED3 = &s
	} else {
		loop.Monitor = &m
	}
	loop.Run()
}
