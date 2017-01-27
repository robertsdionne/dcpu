package hardware

import (
	"github.com/robertsdionne/dcpu/keyboard"
	"github.com/robertsdionne/dcpu/monitor"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/geom"
	"golang.org/x/mobile/gl"
)

type Loop struct {
	Keyboard *keyboard.Device
	Monitor  *monitor.Device
	images   *glutil.Images
	image    *glutil.Image
}

func (l *Loop) Run() {
	app.Main(func(a app.App) {
		var context gl.Context
		var sz size.Event
		for event := range a.Events() {
			switch event := a.Filter(event).(type) {
			case lifecycle.Event:
				switch event.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					context, _ = event.DrawContext.(gl.Context)
					l.onStart(context)
					a.Send(paint.Event{})
				case lifecycle.CrossOff:
					l.onStop()
					context = nil
				}
			case size.Event:
				sz = event
			case paint.Event:
				if context == nil || event.External {
					continue
				}
				l.onPaint(context, sz)
				a.Publish()
				a.Send(paint.Event{})
			case key.Event:
				if l.Keyboard != nil {
					l.Keyboard.Event(event)
				}
			}
		}
	})
}

func (l *Loop) onStart(context gl.Context) {
	l.images = glutil.NewImages(context)
	l.image = l.images.NewImage(l.Monitor.Dimensions())
}

func (l *Loop) onPaint(context gl.Context, sz size.Event) {
	context.Clear(gl.COLOR_BUFFER_BIT)
	if l.Monitor != nil {
		l.Monitor.Paint(l.image.RGBA)
		l.image.Upload()
		context.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
		context.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
		l.image.Draw(sz, geom.Point{}, geom.Point{X: sz.WidthPt}, geom.Point{Y: sz.HeightPt}, l.image.RGBA.Bounds())
	}
}

func (l *Loop) onStop() {
	l.image.Release()
	l.images.Release()
}
