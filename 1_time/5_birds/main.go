// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"time"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"

	"github.com/egonelbre/diy-visuals-gio/internal/g"
	"github.com/egonelbre/diy-visuals-gio/internal/qapp"
)

var Theme = material.NewTheme()
var LastTime time.Time
var Birds []*Bird

type Bird struct {
	Pos   f32.Point
	Time  float32
	Size  float32
	Speed float32
}

func NewBird(at f32.Point, size float32) *Bird {
	return &Bird{
		Pos:   at,
		Size:  size,
		Speed: size * 2,
	}
}

func (b *Bird) Update(deltaTime float32) {
	b.Time += deltaTime * b.Speed
}

func (b *Bird) Draw(gtx layout.Context) {
	screenSize := layout.FPt(gtx.Constraints.Max)

	pos := f32.Pt(
		b.Pos.X*screenSize.X,
		(b.Pos.Y+g.Sin(b.Time*0.1)*0.5)*screenSize.Y,
	)

	size := screenSize.X * 0.1 * b.Size / 10

	var wings clip.Path
	wings.Begin(gtx.Ops)
	wings.MoveTo(pos.Add(f32.Pt(-size*2.5, g.Sin(b.Time-2)*size)))
	wings.LineTo(pos.Add(f32.Pt(-size*2, g.Sin(b.Time-1)*size*0.75)))
	wings.LineTo(pos.Add(f32.Pt(-size, g.Sin(b.Time)*size*0.75)))
	wings.LineTo(pos)
	wings.LineTo(pos.Add(f32.Pt(size, g.Sin(b.Time)*size*0.75)))
	wings.LineTo(pos.Add(f32.Pt(size*2, g.Sin(b.Time-1)*size*0.75)))
	wings.LineTo(pos.Add(f32.Pt(size*2.5, g.Sin(b.Time-2)*size)))

	paint.FillShape(gtx.Ops, g.Black, clip.Stroke{
		Width: 10,
		Path:  wings.End(),
	}.Op())
}

func main() {
	const N = 10
	for i := 1; i < N; i++ {
		Birds = append(Birds, NewBird(f32.Pt(float32(i)/N, 0.5), float32(i)))
	}

	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		defer op.InvalidateOp{}.Add(gtx.Ops)

		if LastTime.IsZero() {
			LastTime = gtx.Now
		}
		deltaTime := gtx.Now.Sub(LastTime)
		LastTime = gtx.Now

		seconds := float32(deltaTime.Seconds())

		for _, bird := range Birds {
			bird.Update(seconds)
			bird.Draw(gtx)
		}

		return layout.Dimensions{}
	})
}
