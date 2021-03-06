package monitor

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/robertsdionne/dcpu"
)

type Device struct {
	BorderColor    uint16
	FontAddress    uint16
	PaletteAddress uint16
	VideoAddress   uint16
	BootPNGPath    string
	dcpu           *dcpu.DCPU
	startTime      time.Time
}

const (
	Width  = areaWidth + 2*borderWidth
	Height = areaHeight + 2*borderHeight

	areaWidth    = 128
	areaHeight   = 96
	borderWidth  = 4
	borderHeight = 8
	bufferWidth  = 32
	bufferHeight = 12
	scale        = 3
	title        = "LEM1802"
	bootPNG      = "documents/boot.png"
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
		0x0814, 0x2200, 0x1414, 0x1400, 0x0022, 0x1408, 0x0259, 0x0600,
		0x3e59, 0x5e00, 0x7e09, 0x7e00, 0x7f49, 0x3600, 0x3e41, 0x2200,
		0x7f41, 0x3e00, 0x7f49, 0x4100, 0x7f09, 0x0100, 0x3e41, 0x7a00,
		0x7f08, 0x7f00, 0x417f, 0x4100, 0x2040, 0x3f00, 0x7f08, 0x7700,
		0x7f40, 0x4000, 0x7f06, 0x7f00, 0x7f01, 0x7e00, 0x3e41, 0x3e00,
		0x7f09, 0x0600, 0x3e61, 0x7e00, 0x7f09, 0x7600, 0x2649, 0x3200,
		0x017f, 0x0100, 0x3f40, 0x7f00, 0x1f60, 0x1f00, 0x7f30, 0x7f00,
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

	TestPattern = []uint16{
		0xf000, 0xf001, 0xf002, 0xf003, 0xf004, 0xf005, 0xf006, 0xf007,
		0xf008, 0xf009, 0xf00a, 0xf00b, 0xf00c, 0xf00d, 0xf00e, 0xf00f,
		0xf010, 0xf011, 0xf012, 0xf013, 0xf014, 0xf015, 0xf016, 0xf017,
		0xf018, 0xf019, 0xf01a, 0xf01b, 0xf01c, 0xf01d, 0xf01e, 0xf01f,
		0xf020, 0xf021, 0xf022, 0xf023, 0xf024, 0xf025, 0xf026, 0xf027,
		0xf028, 0xf029, 0xf02a, 0xf02b, 0xf02c, 0xf02d, 0xf02e, 0xf02f,
		0xf030, 0xf031, 0xf032, 0xf033, 0xf034, 0xf035, 0xf036, 0xf037,
		0xf038, 0xf039, 0xf03a, 0xf03b, 0xf03c, 0xf03d, 0xf03e, 0xf03f,
		0xf040, 0xf041, 0xf042, 0xf043, 0xf044, 0xf045, 0xf046, 0xf047,
		0xf048, 0xf049, 0xf04a, 0xf04b, 0xf04c, 0xf04d, 0xf04e, 0xf04f,
		0xf050, 0xf051, 0xf052, 0xf053, 0xf054, 0xf055, 0xf056, 0xf057,
		0xf058, 0xf059, 0xf05a, 0xf05b, 0xf05c, 0xf05d, 0xf05e, 0xf05f,
		0xf060, 0xf061, 0xf062, 0xf063, 0xf064, 0xf065, 0xf066, 0xf067,
		0xf068, 0xf069, 0xf06a, 0xf06b, 0xf06c, 0xf06d, 0xf06e, 0xf06f,
		0xf070, 0xf071, 0xf072, 0xf073, 0xf074, 0xf075, 0xf076, 0xf077,
		0xf078, 0xf079, 0xf07a, 0xf07b, 0xf07c, 0xf07d, 0xf07e, 0xf07f,
	}

	boot image.Image
)

func (d *Device) Execute(dcpu *dcpu.DCPU) {
	if d.dcpu == nil {
		d.dcpu = dcpu
	}
}

func (d *Device) GetID() uint32 {
	return ID
}

func (d *Device) GetManufacturerID() uint32 {
	return ManufacturerID
}

func (d *Device) GetVersion() uint16 {
	return Version
}

func (d *Device) HandleHardwareInterrupt(dcpu *dcpu.DCPU) {
	switch dcpu.RegisterA {
	case MemoryMapScreen:
		d.VideoAddress = dcpu.RegisterB

	case MemoryMapFont:
		d.FontAddress = dcpu.RegisterB

	case MemoryMapPalette:
		d.PaletteAddress = dcpu.RegisterB

	case SetBorderColor:
		d.BorderColor = dcpu.RegisterB & 0xf

	case MemoryDumpFont:
		copy(dcpu.Memory[dcpu.RegisterB:dcpu.RegisterB+0x100], defaultFont)

	case MemoryDumpPalette:
		copy(dcpu.Memory[dcpu.RegisterB:dcpu.RegisterB+0x10], defaultPalette)
	}
}

type imagePart struct {
	src image.Rectangle
	dst image.Rectangle
}

func (i *imagePart) Len() int {
	return 1
}

func (i *imagePart) Src(int) (x0, y0, x1, y1 int) {
	return i.src.Min.X, i.src.Min.Y, i.src.Max.X, i.src.Max.Y
}

func (i *imagePart) Dst(int) (x0, y0, x1, y1 int) {
	return i.dst.Min.X, i.dst.Min.Y, i.dst.Max.X, i.dst.Max.Y
}

func (d *Device) Dimensions() (int, int) {
	d.startTime = time.Now()
	return Width, Height
}

func (d *Device) Paint(img *image.RGBA) {
	if d.dcpu == nil {
		return
	}

	var font, palette []uint16 = defaultFont, defaultPalette

	if d.FontAddress > 0 {
		font = d.dcpu.Memory[d.FontAddress : d.FontAddress+0x100]
	}

	if d.PaletteAddress > 0 {
		palette = d.dcpu.Memory[d.PaletteAddress : d.PaletteAddress+0x10]
	}

	timeToBlink := time.Since(d.startTime)%(2*time.Second) < 1*time.Second

	for x := 0; x < Width; x++ {
		for y := 0; y < Height; y++ {

			i := x/borderWidth - 1
			j := y/borderHeight - 1

			inBorder := i < 0 || i == bufferWidth || j < 0 || j == bufferHeight
			if inBorder {
				img.Set(x, y, colorFromUint16(palette[d.BorderColor&0xf]))
				continue
			}

			offset := uint16(bufferWidth*j + i)
			character := d.dcpu.Memory[d.VideoAddress+offset]
			blink := character&0x0080 > 0
			foregroundColor := character & 0xf000 >> 12
			backgroundColor := character & 0x0f00 >> 8
			if blink && timeToBlink {
				foregroundColor, backgroundColor = backgroundColor, foregroundColor
			}

			foreground := lookupPixel(font, x%borderWidth, y%borderHeight, character&0x7f)
			if foreground {
				img.Set(x, y, colorFromUint16(palette[foregroundColor]))
				continue
			}

			img.Set(x, y, colorFromUint16(palette[backgroundColor]))
		}
	}

	if time.Since(d.startTime) < 5*time.Second {
		if boot == nil {
			if d.BootPNGPath == "" {
				d.BootPNGPath = bootPNG
			}

			file, err := os.Open(d.BootPNGPath)
			if err != nil {
				log.Fatalln(err)
			}
			defer file.Close()

			boot, err = png.Decode(file)
			if err != nil {
				log.Fatalln(err)
			}
		}

		pt := image.Pt(borderWidth, borderHeight)
		draw.Draw(img, image.Rectangle{pt, pt.Add(boot.Bounds().Size())}, boot, image.ZP, draw.Src)
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

	if x < 2 {
		mask := uint16(0x1 << uint(y))
		if x == 0 {
			mask = mask << borderHeight
		}

		return lo&mask > 0
	}

	mask := uint16(0x1 << uint(y))
	if x == 2 {
		mask = mask << borderHeight
	}

	return hi&mask > 0
}
