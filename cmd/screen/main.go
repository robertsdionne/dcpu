package main

import (
	"flag"

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

	cpu := dcpu.DCPU{}
	k := keyboard.Device{}
	m := monitor.Device{}

	k.Init()

	cpu.Hardware = append(cpu.Hardware, &k, &m)

	cpu.Load(0, parser.Assemble(`
			ias interruptHandler

		:main
			set i, 0x0180
			jsr clearScreen

			set a, 3
			set b, [0x1000]	; borderColor
			hwi 1

			set a, 0	; monitor.SetBorderColor
			set b, screenBuffer
			hwi 1

			set a, 3	; keyboard.SetInterruptMessage
			set b, 1
			hwi 0

			sub pc, main

		:interruptHandler
			set a, 1	; keyboard.GetNextKey
			hwi 0

			set b, c

			set a, 2	; keyboard.GetKeyState
			hwi 0

			ife b, 0x90
				set [shiftFlag], c
			ife b, 0x90
				rfi 0

			ife c, 0
				rfi 0

			ife b, 0x11
				jsr newline
			ife b, 0x11
				rfi 0

			ife b, 0x10
				jsr deleteCharacter

			ifn b, 0x10
				jsr printCharacter

			rfi 0

		:printCharacter
			bor b, [color]
			set [screenBuffer + i], b
			add i, 1
			set [screenBuffer + i], [cursor]
			sub i, 1
			set j, 1
			jsr advanceCursor
			set pc, pop

		:advanceCursor
			add i, j
			ifg i, 0x17f
				set i, 0
			ifu i, 0
				set i, 0
			set pc, pop

		:deleteCharacter
			set [screenBuffer + i], [clearColor]
			set j, -1
			jsr advanceCursor
			ife [screenBuffer + i], [clearColor]
				ifa i, 0
					set pc, deleteCharacter
			set [screenBuffer + i], [cursor]
			set pc, pop

		:newline
			set [screenBuffer + i], [clearColor]
			set push, i
			mod peek, 0x20
			set j, 0x20
			sub j, pop
			jsr advanceCursor
			set [screenBuffer + i], [cursor]
			set pc, pop

		:clearScreen
			set [screenBuffer + i], [clearColor]
			sub i, 1
			ifn i, 0
				set pc, clearScreen
			set pc, pop

		:cursor dat 0xf09f
		:color dat 0xf000
		:clearColor dat 0x0000
		:shiftFlag dat 0
		:screenBuffer dat 0xf09f
	`))

	go cpu.Execute()

	loop := hardware.Loop{
		Keyboard: &k,
		Monitor:  &m,
	}
	loop.Run()
}
