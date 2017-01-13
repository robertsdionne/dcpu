package monitor

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/robertsdionne/dcpu"
)

type Monitor struct {
	BorderColor    uint16
	FontAddress    uint16
	PaletteAddress uint16
	VideoAddress   uint16
}

const (
	areaWidth    = 128
	areaHeight   = 96
	borderWidth  = 4
	borderHeight = 8
	bufferWidth  = 32
	bufferHeight = 12
	width        = areaWidth + 2*borderWidth
	height       = areaHeight + 2*borderHeight
	scale        = 3
	title        = "LEM1802"
)

const (
	ID             = 0x7349f615
	ManufacturerID = 0x1c6c8b36
	Version        = 0x1802
)

const (
	MemoryMapScreen = iota
	MemoryMapFont
	MemoryMapPalette
	SetBorderColor
	MemoryDumpFont
	MemoryDumpPalette
)

var (
	defaultFont = []uint16{
		0xb79e, 0x388e, 0x722c, 0x75f4, 0x19bb, 0x7f8f, 0x85f9, 0xb158,
		0x242e, 0x2400, 0x082a, 0x0800, 0x0008, 0x0000, 0x0808, 0x0808,
		0x00ff, 0x0000, 0x00f8, 0x0808, 0x08f8, 0x0000, 0x080f, 0x0000,
		0x000f, 0x0808, 0x00ff, 0x0808, 0x08f8, 0x0808, 0x08ff, 0x0000,
		0x080f, 0x0808, 0x08ff, 0x0808, 0x6633, 0x99cc, 0x9933, 0x66cc,
		0xfef8, 0xe080, 0x7f1f, 0x0701, 0x0107, 0x1f7f, 0x80e0, 0xf8fe,
		0x5500, 0xaa00, 0x55aa, 0x55aa, 0xffaa, 0xff55, 0x0f0f, 0x0f0f,
		0xf0f0, 0xf0f0, 0x0000, 0xffff, 0xffff, 0x0000, 0xffff, 0xffff,
		0x0000, 0x0000, 0x005f, 0x0000, 0x0300, 0x0300, 0x3e14, 0x3e00,
		0x266b, 0x3200, 0x611c, 0x4300, 0x3629, 0x7650, 0x0002, 0x0100,
		0x1c22, 0x4100, 0x4122, 0x1c00, 0x1408, 0x1400, 0x081c, 0x0800,
		0x4020, 0x0000, 0x0808, 0x0800, 0x0040, 0x0000, 0x601c, 0x0300,
		0x3e49, 0x3e00, 0x427f, 0x4000, 0x6259, 0x4600, 0x2249, 0x3600,
		0x0f08, 0x7f00, 0x2745, 0x3900, 0x3e49, 0x3200, 0x6119, 0x0700,
		0x3649, 0x3600, 0x2649, 0x3e00, 0x0024, 0x0000, 0x4024, 0x0000,
		0x0814, 0x2241, 0x1414, 0x1400, 0x4122, 0x1408, 0x0259, 0x0600,
		0x3e59, 0x5e00, 0x7e09, 0x7e00, 0x7f49, 0x3600, 0x3e41, 0x2200,
		0x7f41, 0x3e00, 0x7f49, 0x4100, 0x7f09, 0x0100, 0x3e41, 0x7a00,
		0x7f08, 0x7f00, 0x417f, 0x4100, 0x2040, 0x3f00, 0x7f08, 0x7700,
		0x7f40, 0x4000, 0x7f06, 0x7f00, 0x7f01, 0x7e00, 0x3e41, 0x3e00,
		0x7f09, 0x0600, 0x3e41, 0xbe00, 0x7f09, 0x7600, 0x2649, 0x3200,
		0x017f, 0x0100, 0x3f40, 0x3f00, 0x1f60, 0x1f00, 0x7f30, 0x7f00,
		0x7708, 0x7700, 0x0778, 0x0700, 0x7149, 0x4700, 0x007f, 0x4100,
		0x031c, 0x6000, 0x0041, 0x7f00, 0x0201, 0x0200, 0x8080, 0x8000,
		0x0001, 0x0200, 0x2454, 0x7800, 0x7f44, 0x3800, 0x3844, 0x2800,
		0x3844, 0x7f00, 0x3854, 0x5800, 0x087e, 0x0900, 0x4854, 0x3c00,
		0x7f04, 0x7800, 0x447d, 0x4000, 0x2040, 0x3d00, 0x7f10, 0x6c00,
		0x417f, 0x4000, 0x7c18, 0x7c00, 0x7c04, 0x7800, 0x3844, 0x3800,
		0x7c14, 0x0800, 0x0814, 0x7c00, 0x7c04, 0x0800, 0x4854, 0x2400,
		0x043e, 0x4400, 0x3c40, 0x7c00, 0x1c60, 0x1c00, 0x7c30, 0x7c00,
		0x6c10, 0x6c00, 0x4c50, 0x3c00, 0x6454, 0x4c00, 0x0836, 0x4100,
		0x0077, 0x0000, 0x4136, 0x0800, 0x0201, 0x0201, 0x0205, 0x0200,
	}

	defaultPalette = []uint16{
		0x0000, // black
		0x0008, // navy
		0x0080, // green
		0x0088, // teal
		0x0800, // maroon
		0x080f, // purple
		0x0880, // olive
		0x0ccc, // silver

		0x0888, // gray
		0x000f, // blue
		0x00f0, // lime
		0x00ff, // cyan
		0x0f00, // red
		0x0f0f, // fuchsia
		0x0ff0, // yellow
		0x0fff, // white
	}
)

func (m *Monitor) Execute(dcpu *dcpu.DCPU) {}

func (m *Monitor) GetID() uint32 {
	return ID
}

func (m *Monitor) GetManufacturerID() uint32 {
	return ManufacturerID
}

func (m *Monitor) GetVersion() uint16 {
	return Version
}

func (m *Monitor) HandleHardwareInterrupt(dcpu *dcpu.DCPU) {
	switch dcpu.RegisterA {
	case MemoryMapScreen:
		m.VideoAddress = dcpu.RegisterB

	case MemoryMapFont:
		m.FontAddress = dcpu.RegisterB

	case MemoryMapPalette:
		m.PaletteAddress = dcpu.RegisterB

	case SetBorderColor:
		m.BorderColor = dcpu.RegisterB & 0xf

	case MemoryDumpFont:
		copy(dcpu.Memory[dcpu.RegisterB:dcpu.RegisterB+0x100], defaultFont)

	case MemoryDumpPalette:
		copy(dcpu.Memory[dcpu.RegisterB:dcpu.RegisterB+0x10], defaultPalette)
	}
}

func (m *Monitor) Poll(dcpu *dcpu.DCPU) {
	log.Fatalln(ebiten.Run(m.update(dcpu), width, height, scale, title))
}

func (m *Monitor) update(dcpu *dcpu.DCPU) func(*ebiten.Image) error {
	return func(screen *ebiten.Image) error {
		img := image.NewRGBA(image.Rect(0, 0, width, height))

		var font, palette []uint16 = defaultFont, defaultPalette

		if m.FontAddress > 0 {
			font = dcpu.Memory[m.FontAddress : m.FontAddress+0x100]
		}

		if m.PaletteAddress > 0 {
			palette = dcpu.Memory[m.PaletteAddress : m.PaletteAddress+0x10]
		}

		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {

				i := x/borderWidth - 1
				j := y/borderHeight - 1

				inBorder := i < 0 || i == bufferWidth || j < 0 || j == bufferHeight
				if inBorder {
					img.Set(x, y, colorFromUint16(palette[m.BorderColor&0xf]))
					continue
				}

				offset := uint16(bufferWidth*j + i)
				character := dcpu.Memory[m.VideoAddress+offset]
				foregroundColor := character & 0xf000 >> 12
				backgroundColor := character & 0x0f00 >> 8

				foreground := lookupPixel(font, x%borderWidth, y%borderHeight, character&0x7f)
				if foreground {
					img.Set(x, y, colorFromUint16(palette[foregroundColor]))
					continue
				}

				img.Set(x, y, colorFromUint16(palette[backgroundColor]))
			}
		}

		return screen.ReplacePixels(img.Pix)
	}
}

func colorFromUint16(value uint16) color.RGBA {
	r := uint8(value & 0x0f00 >> 8)
	g := uint8(value & 0x00f0 >> 4)
	b := uint8(value & 0x000f)
	return color.RGBA{
		R: r | r<<4,
		G: g | g<<4,
		B: b | b<<4,
		A: 0xff,
	}
}

func lookupPixel(font []uint16, x, y int, characterIndex uint16) bool {
	lo := font[2*characterIndex]
	hi := font[2*characterIndex+1]

	if x < 3 {
		mask := uint16(0x1 << uint(y))
		if x == 0 {
			mask = mask << borderHeight
		}

		return lo&mask > 0
	}

	mask := uint16(0x1 << uint(y))
	if x == 3 {
		mask = mask << borderHeight
	}

	return hi&mask > 0
}
