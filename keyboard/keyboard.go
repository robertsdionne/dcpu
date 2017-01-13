package keyboard

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/robertsdionne/dcpu"
)

// Keyboard implements Generic Keyboard (compatible).
type Keyboard struct {
	Message       uint16
	buffer        []uint16
	state         [0x100]uint16
	previousState [0x100]uint16
}

const (
	ID      = 0x30cf7406
	Version = 0x0001
)

const (
	ClearBuffer = iota
	GetNextKey
	GetKeyState
	SetInterruptMessage
)

func (k *Keyboard) Execute(dcpu *dcpu.DCPU) {
	k.previousState = k.state
	k.state = [0x100]uint16{}

	for i := uint16(0); i < 0x100; i++ {
		if ebiten.IsKeyPressed(getKeyCode(i)) {
			k.state[i] = 0x1
		}

		if k.state[i] != k.previousState[i] {
			k.buffer = append(k.buffer, i)
			if k.Message > 0 {
				dcpu.Interrupt(k.Message)
			}
		}
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
	case ClearBuffer:
		k.buffer = make([]uint16, 0, 16)

	case GetNextKey:
		dcpu.RegisterC = 0
		if len(k.buffer) > 0 {
			dcpu.RegisterC = k.buffer[0]
			k.buffer = k.buffer[1:]
		}

	case GetKeyState:
		dcpu.RegisterC = k.state[dcpu.RegisterB]

	case SetInterruptMessage:
		k.Message = dcpu.RegisterB
	}
}

func getKeyCode(key uint16) ebiten.Key {
	switch key {
	case '\t':
		return ebiten.KeyTab
	case 0x10:
		return ebiten.KeyBackspace
	case 0x11:
		return ebiten.KeyEnter
	case 0x12:
		return ebiten.KeyInsert
	case 0x13:
		return ebiten.KeyDelete
	case 0x1b:
		return ebiten.KeyEscape
	case 0x20:
		return ebiten.KeySpace
	case ',':
		return ebiten.KeyComma
	case '.':
		return ebiten.KeyPeriod
	case '0':
		return ebiten.Key0
	case '1':
		return ebiten.Key1
	case '2':
		return ebiten.Key2
	case '3':
		return ebiten.Key3
	case '4':
		return ebiten.Key4
	case '5':
		return ebiten.Key5
	case '6':
		return ebiten.Key6
	case '7':
		return ebiten.Key7
	case '8':
		return ebiten.Key8
	case '9':
		return ebiten.Key9
	case 'a':
		return ebiten.KeyA
	case 'b':
		return ebiten.KeyB
	case 'c':
		return ebiten.KeyC
	case 'd':
		return ebiten.KeyD
	case 'e':
		return ebiten.KeyE
	case 'f':
		return ebiten.KeyF
	case 'g':
		return ebiten.KeyG
	case 'h':
		return ebiten.KeyH
	case 'i':
		return ebiten.KeyI
	case 'j':
		return ebiten.KeyJ
	case 'k':
		return ebiten.KeyK
	case 'l':
		return ebiten.KeyL
	case 'm':
		return ebiten.KeyM
	case 'n':
		return ebiten.KeyN
	case 'o':
		return ebiten.KeyO
	case 'p':
		return ebiten.KeyP
	case 'q':
		return ebiten.KeyQ
	case 'r':
		return ebiten.KeyR
	case 's':
		return ebiten.KeyS
	case 't':
		return ebiten.KeyT
	case 'u':
		return ebiten.KeyU
	case 'v':
		return ebiten.KeyV
	case 'w':
		return ebiten.KeyW
	case 'x':
		return ebiten.KeyX
	case 'y':
		return ebiten.KeyY
	case 'z':
		return ebiten.KeyZ
	case 0x80:
		return ebiten.KeyUp
	case 0x81:
		return ebiten.KeyDown
	case 0x82:
		return ebiten.KeyLeft
	case 0x83:
		return ebiten.KeyRight
	case 0x90:
		return ebiten.KeyShift
	case 0x91:
		return ebiten.KeyControl
	}

	return ebiten.Key(ebiten.KeyUp + 1)
}
