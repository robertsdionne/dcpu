package main

import (
	. "github.com/robertsdionne/dcpu"
	"github.com/robertsdionne/dcpu/hardware"
	"github.com/robertsdionne/dcpu/sped3"
)

func main() {
	d := DCPU{}
	s := sped3.Device{TargetRotation: 90}

	d.Hardware = append(d.Hardware, &s)

	var (
		main                uint16 = 0x0000
		update              uint16 = 0x0007
		clipCoordinateAbove        = update + 43
		clipCoordinateBelow        = clipCoordinateAbove + 10
		displayBoxes               = clipCoordinateBelow + 7
		displayBoundary            = displayBoxes + 21
		x                          = displayBoundary + 6
		y                          = x + 1
		z                          = y + 1
		vx                         = z + 1
		vy                         = vx + 1
		vz                         = vy + 1
		boxes                      = vz + 1
		boxesEnd                   = boxes + 78
		boundary                   = boxesEnd
		boundaryEnd                = boundary + 90
	)

	d.Load(main, []uint16{
		Basic(Set, RegisterA, Literal), sped3.MapRegion,
		Basic(Set, RegisterX, Literal), 0x8000,
		Basic(Set, RegisterY, Literal), 84,
		Special(HardwareInterrupt, Literal0),
	})

	d.Load(update, []uint16{
		Basic(Add, Location, Location), x, vx,
		Basic(Add, Location, Location), y, vy,
		Basic(Add, Location, Location), z, vz,

		Basic(Set, RegisterA, Literal), x,
		Basic(Set, RegisterB, Literal), vx,
		Special(JumpSubRoutine, Literal), clipCoordinateAbove,
		Special(JumpSubRoutine, Literal), clipCoordinateBelow,

		Basic(Set, RegisterA, Literal), y,
		Basic(Set, RegisterB, Literal), vy,
		Special(JumpSubRoutine, Literal), clipCoordinateAbove,
		Special(JumpSubRoutine, Literal), clipCoordinateBelow,

		Basic(Set, RegisterA, Literal), z,
		Basic(Set, RegisterB, Literal), vz,
		Special(JumpSubRoutine, Literal), clipCoordinateAbove,
		Special(JumpSubRoutine, Literal), clipCoordinateBelow,

		Basic(Set, RegisterI, Literal), 0x8000,
		Basic(Set, RegisterJ, Literal), boxes,
		Special(JumpSubRoutine, Literal), displayBoxes,
		Special(JumpSubRoutine, Literal), displayBoundary,

		Basic(Set, ProgramCounter, Literal), update,
	})

	d.Load(clipCoordinateAbove, []uint16{
		Basic(IfEqual, LocationInRegisterA, Literal), 223,
		Basic(Set, ProgramCounter, Pop),
		Basic(IfUnder, LocationInRegisterA, Literal), 223,
		Basic(Set, ProgramCounter, Pop),
		Basic(Set, LocationInRegisterA, Literal), 223,
		Basic(MultiplySigned, LocationInRegisterB, LiteralNegative1),
		Basic(Set, ProgramCounter, Pop),
	})

	d.Load(clipCoordinateBelow, []uint16{
		Basic(IfEqual, LocationInRegisterA, Literal0),
		Basic(Set, ProgramCounter, Pop),
		Basic(IfAbove, LocationInRegisterA, Literal0),
		Basic(Set, ProgramCounter, Pop),
		Basic(Set, LocationInRegisterA, Literal0),
		Basic(MultiplySigned, LocationInRegisterB, LiteralNegative1),
		Basic(Set, ProgramCounter, Pop),
	})

	d.Load(displayBoxes, []uint16{
		Basic(IfEqual, RegisterJ, Literal), boxesEnd,
		Basic(Set, ProgramCounter, Pop),
		Basic(Set, RegisterA, Location), y,
		Basic(ShiftLeft, RegisterA, Literal8),
		Basic(Set, RegisterB, Location), x,
		Basic(BinaryAnd, RegisterB, Literal), 0x00ff,
		Basic(BinaryOr, RegisterA, RegisterB),
		Basic(Add, RegisterA, LocationInRegisterJ),
		Basic(SetThenIncrement, LocationInRegisterI, RegisterA),
		Basic(Set, RegisterA, Location), z,
		Basic(BinaryAnd, RegisterA, Literal), 0x00ff,
		Basic(Add, RegisterA, LocationInRegisterJ),
		Basic(SetThenIncrement, LocationInRegisterI, RegisterA),
		Basic(Set, ProgramCounter, Literal), displayBoxes,
	})

	d.Load(displayBoundary, []uint16{
		Basic(IfEqual, RegisterJ, Literal), boundaryEnd,
		Basic(Set, ProgramCounter, Pop),
		Basic(SetThenIncrement, LocationInRegisterI, LocationInRegisterJ),
		Basic(Set, ProgramCounter, Literal), displayBoundary,
	})

	d.Load(x, []uint16{0})
	d.Load(y, []uint16{0})
	d.Load(z, []uint16{0})
	d.Load(vx, []uint16{1})
	d.Load(vy, []uint16{2})
	d.Load(vz, []uint16{3})

	d.Load(boxes, []uint16{
		0x0000, 0x0700,
		0x2000, 0x0700,
		0x2020, 0x0700,
		0x0020, 0x0700,
		0x0000, 0x0700,

		0x0000, 0x0720,
		0x0000, 0x0700,
		0x2000, 0x0700,
		0x2000, 0x0720,
		0x2020, 0x0720,
		0x2020, 0x0700,
		0x0020, 0x0700,
		0x0020, 0x0720,

		0x0000, 0x0720,
		0x2000, 0x0720,
		0x2020, 0x0720,
		0x0020, 0x0720,
		0x0000, 0x0720,
		0x0000, 0x0020,

		0x3000, 0x0000,
		0x3000, 0x0700,
		0x5000, 0x0700,
		0x5020, 0x0700,
		0x3020, 0x0700,
		0x3000, 0x0700,

		0x3000, 0x0720,
		0x3000, 0x0700,
		0x5000, 0x0700,
		0x5000, 0x0720,
		0x5020, 0x0720,
		0x5020, 0x0700,
		0x3020, 0x0700,
		0x3020, 0x0720,

		0x3000, 0x0720,
		0x5000, 0x0720,
		0x5020, 0x0720,
		0x3020, 0x0720,
		0x3000, 0x0720,
		0x3000, 0x0020,
	})

	d.Load(boundary, []uint16{
		0x0000, 0x0000,
		0x0000, 0x0700,
		0x2000, 0x0700,
		0x2000, 0x0000,
		0xDF00, 0x0000,
		0xDF00, 0x0700,
		0xFF00, 0x0700,
		0xFF20, 0x0700,
		0xFF20, 0x0000,
		0xFFDF, 0x0000,
		0xFFDF, 0x0700,
		0xFFFF, 0x0700,
		0xDFFF, 0x0700,
		0xDFFF, 0x0000,
		0x20FF, 0x0000,
		0x20FF, 0x0700,
		0x00FF, 0x0700,
		0x00DF, 0x0700,
		0x00DF, 0x0000,
		0x0020, 0x0000,
		0x0020, 0x0700,
		0x0000, 0x0700,
		0x0000, 0x0000,

		0x0000, 0x00FF,
		0x0000, 0x07FF,
		0x2000, 0x07FF,
		0x2000, 0x00FF,
		0xDF00, 0x00FF,
		0xDF00, 0x07FF,
		0xFF00, 0x07FF,
		0xFF20, 0x07FF,
		0xFF20, 0x00FF,
		0xFFDF, 0x00FF,
		0xFFDF, 0x07FF,
		0xFFFF, 0x07FF,
		0xDFFF, 0x07FF,
		0xDFFF, 0x00FF,
		0x20FF, 0x00FF,
		0x20FF, 0x07FF,
		0x00FF, 0x07FF,
		0x00DF, 0x07FF,
		0x00DF, 0x00FF,
		0x0020, 0x00FF,
		0x0020, 0x07FF,
		0x0000, 0x07FF,
	})

	go d.Execute()

	loop := hardware.Loop{
		SPED3: &s,
	}
	loop.Run()
}
