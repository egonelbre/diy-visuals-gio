// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"time"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget/material"

	"github.com/egonelbre/diy-visuals-gio/internal/ease"
	"github.com/egonelbre/diy-visuals-gio/internal/g"
	"github.com/egonelbre/diy-visuals-gio/internal/qapp"
)

var Theme = material.NewTheme()
var Start time.Time

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		screenSize := layout.FPt(gtx.Constraints.Max)
		op.InvalidateOp{}.Add(gtx.Ops)

		if Start.IsZero() {
			Start = gtx.Now
		}

		now := float32(gtx.Now.Sub(Start).Seconds())
		animationTime := g.Clamp(g.Mod(now, 3), 0, 1)

		Start := f32.Pt(screenSize.X/2, 0)
		End := f32.Pt(screenSize.X/2, screenSize.Y/2)

		At := g.PtLerp(ease.InOutQuad(animationTime), Start, End)
		defer op.Affine(
			f32.Affine2D{}.Offset(At),
		).Push(gtx.Ops).Pop()

		scale := g.Map(g.Sin(now*1.4), -1, 1, 1, 1.3)
		defer op.Affine(
			f32.Affine2D{}.Scale(f32.Point{},
				f32.Point{
					X: scale + g.Sin(now*1.1)*0.2,
					Y: scale + g.Sin(now*1.2)*0.2,
				}),
		).Push(gtx.Ops).Pop()

		defer op.Affine(
			f32.Affine2D{}.Rotate(f32.Point{},
				g.Sin(now*2.3)*0.1+g.Cos(now*1.7)*0.2,
			),
		).Push(gtx.Ops).Pop()

		const width = 300
		const height = 100
		paint.FillShape(gtx.Ops, g.Black,
			g.FillRect(gtx.Ops, -width/2, -height/2, width, height))

		// using macros
		macro := op.Record(gtx.Ops)
		lbl := material.Label(Theme, gtx.Metric.PxToSp(height/2), "Hello")
		lbl.Color = g.White
		lbl.Alignment = text.Middle
		lblgtx := gtx
		lblgtx.Constraints.Min = image.Point{}
		dims := lbl.Layout(lblgtx)
		call := macro.Stop()
		op.Offset(dims.Size.Div(-2)).Add(gtx.Ops)
		call.Add(gtx.Ops)

		return layout.Dimensions{}
	}, app.Size(1024, 512))
}
