package sped3

import (
	"encoding/binary"
	"log"
	"math"

	"github.com/robertsdionne/dcpu"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"
)

type Device struct {
	LastError      uint16
	RegionAddress  uint16
	State          uint16
	TargetRotation uint16
	VertexCount    uint16
	rotation       float32
	dcpu           *dcpu.DCPU

	program         gl.Program
	modelView       gl.Uniform
	projection      gl.Uniform
	color           gl.Attrib
	position        gl.Attrib
	redBeamBuffer   gl.Buffer
	greenBeamBuffer gl.Buffer
	blueBeamBuffer  gl.Buffer
	lineBuffer      gl.Buffer
}

const (
	ID             = 0x42babf3c
	ManufacturerID = 0x1eb37e91
	Version        = 0x0003
)

const (
	PollDevice = iota
	MapRegion
	RotateDevice
)

const (
	StateNoData = iota
	StateRunning
	StateTurning
)

const (
	ErrorNone   = 0x0000
	ErrorBroken = 0xffff
)

const (
	vertexShader = `
#version 100
uniform mat4 projection;
uniform mat4 model_view;

attribute vec4 position;
attribute vec4 color;

varying vec4 fragment_color;

void main() {
	gl_Position = projection * model_view * position;
	fragment_color = color;
}
`

	fragmentShader = `
#version 100
precision mediump float;

varying vec4 fragment_color;

void main() {
	gl_FragColor = fragment_color;
}
`
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
	case PollDevice:
		dcpu.RegisterB = d.State
		dcpu.RegisterC = d.LastError

	case MapRegion:
		d.RegionAddress = dcpu.RegisterX
		d.VertexCount = dcpu.RegisterY

	case RotateDevice:
		d.TargetRotation = dcpu.RegisterX
	}
}

func (d *Device) Start(context gl.Context) {
	context.ClearColor(0, 0, 0, 1)
	context.Enable(gl.BLEND)
	context.DepthMask(false)
	context.Disable(gl.CULL_FACE)
	context.BlendFunc(gl.SRC_ALPHA, gl.ONE)
	context.LineWidth(1)

	var err error

	d.program, err = glutil.CreateProgram(context, vertexShader, fragmentShader)
	if err != nil {
		log.Fatalln(err)
	}

	d.modelView = context.GetUniformLocation(d.program, "model_view")
	d.projection = context.GetUniformLocation(d.program, "projection")

	d.color = context.GetAttribLocation(d.program, "color")
	d.position = context.GetAttribLocation(d.program, "position")

	d.redBeamBuffer = context.CreateBuffer()
	d.greenBeamBuffer = context.CreateBuffer()
	d.blueBeamBuffer = context.CreateBuffer()
	d.lineBuffer = context.CreateBuffer()
}

func (d *Device) Paint(context gl.Context) {
	if d.dcpu == nil {
		return
	}

	d.rotation = 0.99*float32(d.rotation) + 0.01*float32(d.TargetRotation)

	red, green, blue := d.beam()
	redCount := uint16(len(red) / 7)
	greenCount := uint16(len(green) / 7)
	blueCount := uint16(len(blue) / 7)

	lines := d.lines()

	context.BindBuffer(gl.ARRAY_BUFFER, d.redBeamBuffer)
	context.BufferData(gl.ARRAY_BUFFER, f32.Bytes(binary.LittleEndian, red...), gl.STREAM_DRAW)

	context.BindBuffer(gl.ARRAY_BUFFER, d.greenBeamBuffer)
	context.BufferData(gl.ARRAY_BUFFER, f32.Bytes(binary.LittleEndian, green...), gl.STREAM_DRAW)

	context.BindBuffer(gl.ARRAY_BUFFER, d.blueBeamBuffer)
	context.BufferData(gl.ARRAY_BUFFER, f32.Bytes(binary.LittleEndian, blue...), gl.STREAM_DRAW)

	context.BindBuffer(gl.ARRAY_BUFFER, d.lineBuffer)
	context.BufferData(gl.ARRAY_BUFFER, f32.Bytes(binary.LittleEndian, lines...), gl.STREAM_DRAW)

	context.UseProgram(d.program)

	context.UniformMatrix4fv(d.modelView, []float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, -0.8, -1.5, 1,
	})
	context.UniformMatrix4fv(d.projection, getFrustumMatrix(-0.1, 0.1, -0.1, 0.1, 0.1, 1000))

	d.drawBuffer(context, d.redBeamBuffer, gl.TRIANGLES, redCount)
	d.drawBuffer(context, d.greenBeamBuffer, gl.TRIANGLES, greenCount)
	d.drawBuffer(context, d.blueBeamBuffer, gl.TRIANGLES, blueCount)
	d.drawBuffer(context, d.lineBuffer, gl.LINE_STRIP, d.VertexCount)
}

func (d *Device) drawBuffer(context gl.Context, buffer gl.Buffer, mode gl.Enum, count uint16) {
	context.BindBuffer(gl.ARRAY_BUFFER, buffer)

	context.EnableVertexAttribArray(d.position)
	context.EnableVertexAttribArray(d.color)

	context.VertexAttribPointer(d.position, 3, gl.FLOAT, false, 28, 0)
	context.VertexAttribPointer(d.color, 4, gl.FLOAT, false, 28, 12)

	context.DrawArrays(mode, 0, int(count))

	context.DisableVertexAttribArray(d.position)
	context.DisableVertexAttribArray(d.color)
}

func (d *Device) beam() (red, green, blue []float32) {
	alpha := 1.0 / 2 / f32.Sqrt(float32(d.VertexCount))
	offset := float32(-4*math.Pi/3 + math.Pi/2)
	theta0 := 0 + offset
	theta1 := 2*math.Pi/3 + offset
	theta2 := 4*math.Pi/3 + offset
	origin := map[uint16][]float32{
		4: {f32.Cos(theta0) / 2, 0, f32.Sin(theta0) / 2, 0, 0, 0, alpha},
		2: {f32.Cos(theta1) / 2, 0, f32.Sin(theta1) / 2, 0, 0, 0, alpha},
		1: {f32.Cos(theta2) / 2, 0, f32.Sin(theta2) / 2, 0, 0, 0, alpha},
	}

	for i := 0; i < int(d.VertexCount); i++ {
		word0 := d.dcpu.Memory[int(d.RegionAddress)+2*i+0]
		word1 := d.dcpu.Memory[int(d.RegionAddress)+2*i+1]
		word2 := d.dcpu.Memory[int(d.RegionAddress)+2*i+2]
		word3 := d.dcpu.Memory[int(d.RegionAddress)+2*i+3]

		color0 := word1 >> 8
		color1 := word3 >> 8

		v0 := d.buildVertex(word0, word1, alpha)
		v1 := d.buildVertex(word2, word3, alpha)

		v0Black := []float32{v0[0], v0[1], v0[2], 0, 0, 0, v0[6]}
		v1Black := []float32{v1[0], v1[1], v1[2], 0, 0, 0, v1[6]}

		for c := uint16(1); c <= 4; c <<= 1 {
			if color0&c > 0 || color1&c > 0 {
				var r, g, b float32
				if c == 4 {
					r = 1
				}
				if c == 2 {
					g = 1
				}
				if c == 1 {
					b = 1
				}
				v0Prime := []float32{v0[0], v0[1], v0[2], r, g, b, v0[6]}
				v1Prime := []float32{v1[0], v1[1], v1[2], r, g, b, v1[6]}

				if c == 4 {
					if color0&c > 0 {
						red = append(red, v0Prime...)
					} else {
						red = append(red, v0Black...)
					}
					if color1&c > 0 {
						red = append(red, v1Prime...)
					} else {
						red = append(red, v1Black...)
					}
					red = append(red, origin[c]...)
				}

				if c == 2 {
					if color0&c > 0 {
						green = append(green, v0Prime...)
					} else {
						green = append(green, v0Black...)
					}
					if color1&c > 0 {
						green = append(green, v1Prime...)
					} else {
						green = append(green, v1Black...)
					}
					green = append(green, origin[c]...)
				}

				if c == 1 {
					if color0&c > 0 {
						blue = append(blue, v0Prime...)
					} else {
						blue = append(blue, v0Black...)
					}
					if color1&c > 0 {
						blue = append(blue, v1Prime...)
					} else {
						blue = append(blue, v1Black...)
					}
					blue = append(blue, origin[c]...)
				}
			}
		}
	}
	return
}

func (d *Device) lines() (data []float32) {
	for i := 0; i < int(d.VertexCount); i++ {
		data = append(data, d.buildVertex(
			d.dcpu.Memory[int(d.RegionAddress)+2*i], d.dcpu.Memory[int(d.RegionAddress)+2*i+1], 0.8)...)
	}
	return
}

func (d *Device) buildVertex(word0, word1 uint16, alpha float32) (vertex []float32) {
	x := float32(word0&0xff)/255 - 0.5
	z := float32(word0>>8)/255 - 0.5
	y := float32(word1&0xff)/255 + 0.5
	color := word1 >> 8
	r := float32(0x4 & color)
	g := float32(0x2 & color)
	b := float32(0x1 & color)
	theta := math.Pi / 180 * d.rotation
	sin, cos := f32.Sin(theta), f32.Cos(theta)
	vertex = []float32{
		x*cos + z*sin, y, x*sin - z*cos, r, g, b, alpha,
	}
	return
}

func getFrustumMatrix(left, right, bottom, top, near, far float32) (matrix []float32) {
	a := (right + left) / (right - left)
	b := (top + bottom) / (top - bottom)
	c := -(far + near) / (far - near)
	d := -(2 * far * near) / (far - near)
	matrix = []float32{
		2 * near / (right - left), 0, 0, 0,
		0, 2 * near / (top - bottom), 0, 0,
		a, b, c, -1,
		0, 0, d, 0,
	}
	return
}

func (d *Device) Stop(context gl.Context) {
	context.DeleteProgram(d.program)

	context.DeleteBuffer(d.redBeamBuffer)
	context.DeleteBuffer(d.greenBeamBuffer)
	context.DeleteBuffer(d.blueBeamBuffer)
	context.DeleteBuffer(d.lineBuffer)
}
