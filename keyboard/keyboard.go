package keyboard

import (
	"github.com/robertsdionne/dcpu"
	"golang.org/x/mobile/event/key"
)

// Device implements Generic Keyboard (compatible).
type Device struct {
	Message uint16
	channel chan uint16
	dcpu    *dcpu.DCPU
	state   map[uint16]bool
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

func (d *Device) Init() {
	d.channel = make(chan uint16, 512)
	d.state = map[uint16]bool{}
}

func (d *Device) Event(event key.Event) {
	code := getKeyCode(event)
	if code == 0 {
		return
	}

	switch event.Direction {
	case key.DirPress:
		d.state[code] = true
	case key.DirNone:
		return
	case key.DirRelease:
		d.state[code] = false
	}

	d.channel <- code
	if d.Message > 0 && d.dcpu != nil {
		d.dcpu.Interrupt(d.Message)
	}
}

func (d *Device) Execute(dcpu *dcpu.DCPU) {
	if d.dcpu == nil {
		d.dcpu = dcpu
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
		d.Init()

	case GetNextKey:
		select {
		case code := <-d.channel:
			dcpu.RegisterC = code
		default:
			dcpu.RegisterC = 0
		}

	case GetKeyState:
		dcpu.RegisterC = 0
		if d.state[dcpu.RegisterB] {
			dcpu.RegisterC = 1
		}

	case SetInterruptMessage:
		d.Message = dcpu.RegisterB
	}
}

func getKeyCode(event key.Event) uint16 {
	switch event.Code {
	case key.CodeDeleteBackspace:
		return 0x10
	case key.CodeReturnEnter, key.CodeKeypadEnter:
		return 0x11
	case key.CodeInsert:
		return 0x12
	case key.CodeDeleteForward:
		return 0x13
	case key.CodeEscape:
		return 0x1b
	case key.CodeUpArrow:
		return 0x80
	case key.CodeDownArrow:
		return 0x81
	case key.CodeLeftArrow:
		return 0x82
	case key.CodeRightArrow:
		return 0x83
	case key.CodeLeftShift, key.CodeRightShift:
		return 0x90
	case key.CodeLeftControl, key.CodeRightControl:
		return 0x91
	case key.CodeF1, key.CodeF2, key.CodeF3, key.CodeF4,
		key.CodeF5, key.CodeF6, key.CodeF7, key.CodeF8,
		key.CodeF9, key.CodeF10, key.CodeF11, key.CodeF12:
		fallthrough
	case key.CodeF13, key.CodeF14, key.CodeF15, key.CodeF16,
		key.CodeF17, key.CodeF18, key.CodeF19, key.CodeF20,
		key.CodeF21, key.CodeF22, key.CodeF23, key.CodeF24:
		fallthrough
	case key.CodePause, key.CodeHome, key.CodeEnd, key.CodePageUp, key.CodePageDown:
		fallthrough
	case key.CodeKeypadNumLock:
		fallthrough
	case key.CodeHelp, key.CodeMute, key.CodeVolumeDown, key.CodeVolumeUp:
		fallthrough
	case key.CodeLeftAlt, key.CodeRightAlt, key.CodeLeftGUI, key.CodeRightGUI, key.CodeCompose:
		fallthrough
	case key.CodeUnknown:
		return 0
	default:
		return uint16(event.Rune)
	}
}
