// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"time"

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
		op.InvalidateOp{}.Add(gtx.Ops)

		paint.Fill(gtx.Ops, g.Black)

		screenSize := gtx.Constraints.Max
		screenWidth, screenHeight := float32(screenSize.X), float32(screenSize.Y)

		millis := float32(gtx.Now.Sub(Start).Seconds())

		rx := (g.Sin(millis*2) + 1) * 0.5
		ry := (g.Sin(millis*3) + 1) * 0.5

		x := rx*screenWidth*0.5 + screenWidth*0.25
		y := ry*screenHeight*0.5 + screenHeight*0.25

		paint.FillShape(gtx.Ops,
			g.HSL(millis, 0.5, 0.7),
			g.FillRect(gtx.Ops, x, y, 30, 30),
		)

		paint.FillShape(gtx.Ops, g.White,
			g.FillRect(gtx.Ops, x+15, 50, 2, 30),
		)
		paint.FillShape(gtx.Ops, g.White,
			g.FillRect(gtx.Ops, 50, y+15, 30, 2),
		)

		return layout.Dimensions{}
	})
}
