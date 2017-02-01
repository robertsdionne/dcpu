package main

import (
	"flag"

	"github.com/robertsdionne/dcpu"
	"github.com/robertsdionne/dcpu/floppy"
	"github.com/robertsdionne/dcpu/parser"
	"github.com/robertsdionne/dcpu/stderr"
	"github.com/robertsdionne/dcpu/stdin"
	"github.com/robertsdionne/dcpu/stdout"
)

var (
	floppyDisk = flag.String("floppy-disk", "floppy/0", "The floppy file.")
	sector     = flag.Int("sector", 0, "The floppy sector.")
)

func main() {
	flag.Parse()

	d := dcpu.DCPU{}
	f := floppy.Device{}

	f.Insert(*floppyDisk, false)
	defer f.Eject()

	d.Hardware = append(d.Hardware, &stdin.Device{}, &stdout.Device{}, &stderr.Device{}, &f)

	d.Load(0, parser.Assemble(`
		:main
			set a, 1
			set x, 1024
			set y, 0x1000
			hwi 0
			set a, 3
			set x, i
			set y, 0x1000
			hwi 3
			add i, 1
			dum
			set pc, main
	`))

	d.Execute()
}
