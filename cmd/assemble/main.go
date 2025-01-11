package main

import (
    "encoding/binary"
    "log"
    "os"
    "github.com/robertsdionne/dcpu/assembler"
)

func main() {
	program, err := assembler.AssembleFile("/dev/stdin")
	if err != nil {
		log.Fatalln(err)
	}

    for _, instruction := range program {
        err = binary.Write(os.Stdout, binary.LittleEndian, instruction)
        if err != nil {
            log.Fatalln(err)
        }
    }
}
