package keyboard

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/robertsdionne/dcpu"
)

// Keyboard implements Generic Keyboard (compatible).
type Keyboard struct {
	Message uint16
	channel chan uint16
	state   [0x100]uint16
}

const (
	ID      = 0x30cf7406
	Version = 0x0001
)

const (
	clearBuffer = iota
	getNextKey
	getKeyState
	setInterruptMessage
)

func (k *Keyboard) Init() {
	k.channel = make(chan uint16, 16)
}

func (k *Keyboard) Execute(dcpu *dcpu.DCPU) {
	if k.Message > 0 && len(k.channel) > 0 {
		dcpu.Interrupt(k.Message)
	}
}

func (k *Keyboard) GetID() uint32 {
	return ID
}

func (k *Keyboard) GetManufacturerID() uint32 {
	return 0
}

func (k *Keyboard) GetVersion() uint16 {
	return Version
}

func (k *Keyboard) HandleHardwareInterrupt(dcpu *dcpu.DCPU) {
	switch dcpu.RegisterA {
	case clearBuffer:
		close(k.channel)
		k.channel = make(chan uint16, 16)

	case getNextKey:
		dcpu.RegisterC = 0
		select {
		case key := <-k.channel:
			dcpu.RegisterC = key
		}

	case getKeyState:
		dcpu.RegisterC = k.state[dcpu.RegisterB]

	case setInterruptMessage:
		k.Message = dcpu.RegisterB
	}
}

func (k *Keyboard) Poll() {
	err := glfw.Init()
	if err != nil {
		log.Fatalln(err)
	}
	defer glfw.Terminate()

	err = gl.Init()
	if err != nil {
		log.Fatalln(err)
	}

	window, err := glfw.CreateWindow(128, 1, "keyboard", nil, nil)
	if err != nil {
		log.Fatalln(err)
	}

	window.SetKeyCallback(func(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		var value uint16
		switch {
		case key == glfw.KeyBackspace:
			value = 0x10
		case key == glfw.KeyEnter:
			value = 0x11
		case key == glfw.KeyInsert:
			value = 0x12
		case key == glfw.KeyDelete:
			value = 0x13
		case key == glfw.KeyUp:
			value = 0x80
		case key == glfw.KeyDown:
			value = 0x81
		case key == glfw.KeyLeft:
			value = 0x82
		case key == glfw.KeyRight:
			value = 0x83
		default:
			value = uint16(key)
		}

		if value > 0 {
			switch action {
			case glfw.Press, glfw.Repeat:
				k.state[value] = 0x1
				k.channel <- value

			case glfw.Release:
				k.state[value] = 0x0
				k.channel <- value
			}
		}
	})

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		glfw.PollEvents()
		gl.Clear(gl.COLOR_BUFFER_BIT)
		window.SwapBuffers()
	}
}
