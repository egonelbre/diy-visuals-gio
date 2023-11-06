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
var Start = time.Now()

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		op.InvalidateOp{}.Add(gtx.Ops)

		millis := float32(gtx.Now.Sub(Start).Seconds()) * 3

		center := layout.FPt(gtx.Constraints.Max).Div(2)

		var left clip.Path
		left.Begin(gtx.Ops)
		left.MoveTo(center)
		left.LineTo(center.Add(f32.Pt(-150, g.Sin(millis)*100)))
		left.LineTo(center.Add(f32.Pt(-250, g.Sin(millis-1)*100)))
		left.LineTo(center.Add(f32.Pt(-300, g.Sin(millis-2)*150)))

		paint.FillShape(gtx.Ops, g.Black, clip.Stroke{
			Width: 8,
			Path:  left.End(),
		}.Op())
		return layout.Dimensions{}

		var right clip.Path
		right.Begin(gtx.Ops)
		right.MoveTo(center)
		right.LineTo(center.Add(f32.Pt(150, g.Sin(millis)*100)))
		right.LineTo(center.Add(f32.Pt(250, g.Sin(millis-1)*100)))
		right.LineTo(center.Add(f32.Pt(300, g.Sin(millis-2)*150)))

		paint.FillShape(gtx.Ops, g.Black, clip.Stroke{
			Width: 8,
			Path:  right.End(),
		}.Op())

		return layout.Dimensions{}
	})
}
