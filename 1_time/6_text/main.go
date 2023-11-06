// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"math"
	"time"

	"gioui.org/f32"
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/unit"
	"gioui.org/widget/material"

	"github.com/egonelbre/diy-visuals-gio/internal/g"
	"github.com/egonelbre/diy-visuals-gio/internal/qapp"
)

var Theme = material.NewTheme()
var Start = time.Now()

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		op.InvalidateOp{}.Add(gtx.Ops)

		screenSize := layout.FPt(gtx.Constraints.Max)
		millis := float32(gtx.Now.Sub(Start).Seconds()) * 3

		const y0 = 300
		const textSize = unit.Sp(100)
		textHeight := float32(gtx.Metric.Sp(textSize))

		var crop clip.Path
		crop.Begin(gtx.Ops)
		for i := float32(0); i < 3; i++ {
			y := y0 + g.Sin(millis*(i+1.15)+i*math.Phi)*textHeight - 5
			h := float32(10)
			crop.MoveTo(f32.Pt(0, y))
			crop.LineTo(f32.Pt(screenSize.X, y))
			crop.LineTo(f32.Pt(screenSize.X, y+h))
			crop.LineTo(f32.Pt(0, y+h))
			crop.Close()
		}

		defer clip.Outline{
			Path: crop.End(),
		}.Op().Push(gtx.Ops).Pop()

		g.Offset{
			Pos: f32.Pt(300, y0),
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(Theme, textSize, "Hello")
			lbl.Font.Weight = font.Black
			return lbl.Layout(gtx)
		})

		return layout.Dimensions{}
	})
}
