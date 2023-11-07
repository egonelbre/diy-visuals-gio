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
			defer op.Affine(
				f32.Affine2D{}.Offset(f32.Pt(g.Sincos(millis)).Mul(300)),
			).Push(gtx.Ops).Pop()

			paint.FillShape(gtx.Ops, g.Black,
				g.FillRect(gtx.Ops, -20, -20, 40, 40),
			)
		}()

		return layout.Dimensions{}
	})
}
