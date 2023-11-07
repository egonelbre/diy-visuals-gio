// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"time"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/widget/material"

	"github.com/egonelbre/diy-visuals-gio/internal/g"
	"github.com/egonelbre/diy-visuals-gio/internal/qapp"
)

var Theme = material.NewTheme()
var Start = time.Now()

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		screenSize := layout.FPt(gtx.Constraints.Max)
		op.InvalidateOp{}.Add(gtx.Ops)

		millis := float32(gtx.Now.Sub(Start).Seconds()) * 3

		func() {
			// center the image
			defer op.Affine(
				f32.Affine2D{}.Offset(screenSize.Div(2)),
			).Push(gtx.Ops).Pop()

			// offset by sin & cos
			sn, cs := g.Sincos(millis)
			sn *= 0.5
			cs *= 0.5
			sn += 1
			cs += 1
			defer op.Affine(
				f32.Affine2D{}.Scale(f32.Point{}, f32.Pt(sn, cs)),
			).Push(gtx.Ops).Pop()

			size := screenSize.X / 5

			paint.FillShape(gtx.Ops, g.Black,
				g.FillRect(gtx.Ops, 0, 0, size, size),
			)
		}()

		return layout.Dimensions{}
	})
}