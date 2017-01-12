package main

import (
	"flag"

	"github.com/robertsdionne/dcpu"
	"github.com/robertsdionne/dcpu/floppy"
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
	f := floppy.Floppy{}

	f.Insert(*floppyDisk, false)
	defer f.Eject()

	d.Hardware = append(d.Hardware, &stdin.Stdin{}, &stdout.Stdout{}, &stderr.Stderr{}, &f)

	d.Load(0, []uint16{
		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal), stdin.ReadWords,
		dcpu.Basic(dcpu.Set, dcpu.RegisterX, dcpu.Literal), 1024,
		dcpu.Basic(dcpu.Set, dcpu.RegisterY, dcpu.Literal), 0x1000,
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal0),
		dcpu.Basic(dcpu.Set, dcpu.RegisterA, dcpu.Literal), floppy.WriteSector,
		dcpu.Basic(dcpu.Set, dcpu.RegisterX, dcpu.RegisterI),
		dcpu.Basic(dcpu.Set, dcpu.RegisterY, dcpu.Literal), 0x1000,
		dcpu.Special(dcpu.HardwareInterrupt, dcpu.Literal3),
		dcpu.Basic(dcpu.Add, dcpu.RegisterI, dcpu.Literal1),
		dcpu.Debug(dcpu.DumpState),
		dcpu.Basic(dcpu.Set, dcpu.ProgramCounter, dcpu.Literal0),
	})

	d.Execute()
}
