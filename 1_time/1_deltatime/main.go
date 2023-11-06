// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"fmt"
	"time"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"

	"github.com/egonelbre/diy-visuals-gio/internal/g"
	"github.com/egonelbre/diy-visuals-gio/internal/qapp"
)

var Theme = material.NewTheme()

var LastTime time.Time

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		op.InvalidateOp{}.Add(gtx.Ops)

		if LastTime.IsZero() {
			LastTime = gtx.Now
		}

		DeltaTime := gtx.Now.Sub(LastTime)
		LastTime = gtx.Now

		g.FillText(Theme, gtx, gtx.Now.String(), f32.Pt(100, 100))
		g.FillText(Theme, gtx, DeltaTime.String(), f32.Pt(100, 200))
		g.FillText(Theme, gtx, fmt.Sprint(1/DeltaTime.Seconds())+"fps", f32.Pt(100, 300))

		return layout.Dimensions{}
	})
}
