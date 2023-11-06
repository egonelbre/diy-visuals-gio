// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"gioui.org/f32"
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/widget/material"
	
	"github.com/egonelbre/diy-visuals-gio/internal/g"
	"github.com/egonelbre/diy-visuals-gio/internal/qapp"
)

var Theme = material.NewTheme()

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		g.Offset{
			Pos: f32.Pt(100, 100),
		}.Layout(gtx, material.Body1(Theme, "Hello").Layout)

		g.Offset{
			Pos: f32.Pt(100, 200),
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(Theme, 40, "Hello")
			lbl.Color = g.HSL(40, 0.7, 0.7)
			lbl.Font.Weight = font.Black
			return lbl.Layout(gtx)
		})

		return layout.Dimensions{}
	})
}
