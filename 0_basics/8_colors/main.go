// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/paint"

	"github.com/egonelbre/diy-visuals-gio/internal/g"
	"github.com/egonelbre/diy-visuals-gio/internal/qapp"
)

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		screenSize := gtx.Constraints.Max
		width := float32(screenSize.X)

		var alpha float32
		for x := 0; x < screenSize.X; x += 10 {
			paint.FillShape(gtx.Ops,
				g.HSL(alpha, 0.7, 0.5),
				g.FillRect(gtx.Ops, float32(x), 100, 10, 50))
			alpha += 0.03
		}

		var beta float32
		for x := 0; x < screenSize.X; x += 10 {
			paint.FillShape(gtx.Ops,
				g.RGB(beta, 0.2, 0.2),
				g.FillRect(gtx.Ops, float32(x), 200, 10, 50))
			beta += 0.03
		}

		for x := 0; x < screenSize.X; x += 10 {
			xf := float32(x)
			paint.FillShape(gtx.Ops,
				g.RGB(
					g.Sin(xf/float32(screenSize.X)),
					g.Sin(1/3+7*xf/width),
					g.Sin(5*xf/width),
				),
				g.FillRect(gtx.Ops, float32(x), 300, 10, 50))
		}

		func() {
			defer g.FillRect(gtx.Ops, 0, 400, float32(screenSize.X), 50).Push(gtx.Ops).Pop()

			paint.LinearGradientOp{
				Stop1:  f32.Pt(0, 0),
				Color1: g.HSL(0, 0.7, 0.7),
				Stop2:  f32.Pt(width, 0),
				Color2: g.HSL(0.7, 0.7, 0.7),
			}.Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)
		}()

		return layout.Dimensions{}
	})
}
