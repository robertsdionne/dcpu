package main

import (
	"flag"
	"io/ioutil"
	"log"

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

	source, err := ioutil.ReadFile("programs/floppydisk.dasm")
	if err != nil {
		log.Fatalln(err)
	}

	d := dcpu.DCPU{}
	f := floppy.Device{}

	f.Insert(*floppyDisk, false)
	defer f.Eject()

	d.Hardware = append(d.Hardware, &stdin.Device{}, &stdout.Device{}, &stderr.Device{}, &f)
	d.Load(0, parser.Assemble(string(source)))

	d.Execute()
}
