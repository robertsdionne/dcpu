package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/robertsdionne/dcpu"
	"github.com/robertsdionne/dcpu/hardware"
	"github.com/robertsdionne/dcpu/keyboard"
	"github.com/robertsdionne/dcpu/monitor"
	"github.com/robertsdionne/dcpu/parser"
)

var (
	borderColor     = flag.Int("border-color", 0, "the border color")
	foregroundColor = flag.Int("foreground-color", 0xf, "the foreground color")
	backgroundColor = flag.Int("background-color", 0x0, "the background color")
)

func main() {
	flag.Parse()

	source, err := ioutil.ReadFile("programs/screen.dasm")
	if err != nil {
		log.Fatalln(err)
	}

	cpu := dcpu.DCPU{}
	k := keyboard.Device{}
	m := monitor.Device{}

	k.Init()

	cpu.Hardware = append(cpu.Hardware, &k, &m)
	cpu.Load(0, parser.Assemble(string(source)))

	go cpu.Execute()

	loop := hardware.Loop{
		Keyboard: &k,
		Monitor:  &m,
	}
	loop.Run()
}
