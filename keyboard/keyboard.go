package keyboard

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/robertsdionne/dcpu"
)

// Device implements Generic Keyboard (compatible).
type Device struct {
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

func (d *Device) Execute(dcpu *dcpu.DCPU) {
	d.previousState = d.state
	d.state = [0x100]uint16{}

	for key := ebiten.Key0; key <= ebiten.KeyMax; key++ {
		i := getKeyCode(key)
		if i == 0xffff {
			continue
		}

		if ebiten.IsKeyPressed(key) {
			switch key {
			case ebiten.KeyShift:
				// noop
			case ebiten.KeyControl:
				// noop
			case ebiten.KeyRightShift, ebiten.KeyRightControl:
				d.state[i] |= 0x1
			default:
				d.state[i] = 0x1
			}
		}
	}

	for i := uint16(0); i < 0x100; i++ {
		if d.state[i] != d.previousState[i] {
			d.buffer = append(d.buffer, i)
			if d.Message > 0 {
				dcpu.Interrupt(d.Message)
			}
		}
	}
}

func (d *Device) GetID() uint32 {
	return ID
}

func (d *Device) GetManufacturerID() uint32 {
	return 0
}

func (d *Device) GetVersion() uint16 {
	return Version
}

func (d *Device) HandleHardwareInterrupt(dcpu *dcpu.DCPU) {
	switch dcpu.RegisterA {
	case ClearBuffer:
		d.buffer = make([]uint16, 0, 16)

	case GetNextKey:
		dcpu.RegisterC = 0
		if len(d.buffer) > 0 {
			dcpu.RegisterC = d.buffer[0]
			d.buffer = d.buffer[1:]
		}

	case GetKeyState:
		dcpu.RegisterC = d.state[dcpu.RegisterB]

	case SetInterruptMessage:
		d.Message = dcpu.RegisterB
	}
}

func getKeyCode(key ebiten.Key) uint16 {
	switch {
	case ebiten.Key0 <= key && key <= ebiten.Key9:
		return uint16(key-ebiten.Key0) + '0'
	case ebiten.KeyA <= key && key <= ebiten.KeyZ:
		return uint16(key-ebiten.KeyA) + 'A'
	case key == ebiten.KeyBackspace:
		return 0x10
	case key == ebiten.KeyComma:
		return ','
	case key == ebiten.KeyControl || key == ebiten.KeyLeftControl || key == ebiten.KeyRightControl:
		return 0x91
	case key == ebiten.KeyDelete:
		return 0x13
	case key == ebiten.KeyDown:
		return 0x81
	case key == ebiten.KeyEnter:
		return 0x11
	case key == ebiten.KeyEscape:
		return 0x1b
	case key == ebiten.KeyInsert:
		return 0x12
	case key == ebiten.KeyLeft:
		return 0x82
	case key == ebiten.KeyPeriod:
		return '.'
	case key == ebiten.KeyRight:
		return 0x83
	case key == ebiten.KeyShift || key == ebiten.KeyLeftShift || key == ebiten.KeyRightShift:
		return 0x90
	case key == ebiten.KeySpace:
		return ' '
	case key == ebiten.KeyTab:
		return '\t'
	case key == ebiten.KeyUp:
		return 0x80
	case key == ebiten.KeyApostrophe:
		return '\''
	case key == ebiten.KeyMinus:
		return '-'
	case key == ebiten.KeySlash:
		return '/'
	case key == ebiten.KeySemicolon:
		return ';'
	case key == ebiten.KeyEqual:
		return '='
	case key == ebiten.KeyLeftBracket:
		return '{'
	case key == ebiten.KeyBackslash:
		return '\\'
	case key == ebiten.KeyRightBracket:
		return '}'
	case key == ebiten.KeyGraveAccent:
		return '`'
	default:
		return 0xffff
	}
}
