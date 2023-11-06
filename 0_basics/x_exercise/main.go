// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	
	"github.com/egonelbre/diy-visuals-gio/internal/g"
	"github.com/egonelbre/diy-visuals-gio/internal/qapp"
)

var Theme = material.NewTheme()

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		screenSize := gtx.Constraints.Max
		screenWidth, screenHeight := float32(screenSize.X), float32(screenSize.Y)
		_, _ = screenWidth, screenHeight

		/*

			Exercise, draw a house.

			You can draw whatever you like, but here are
			a few ideas on what you can add to make the
			scene nicer.

			- some windows
			- a street sign or house number
			- a chimney with smoke
			- a fence
			- some grass
			- bush or a tree
			- a cat or a dog
			- a sun and clouds

		*/

		// Here's a fancier roof you can start with.

		const roofTopX = 300
		const roofTopY = 100
		const roofHeight = 100
		const roofWidth = 100
		const roofCurveX = 50

		var roof clip.Path
		roof.Begin(gtx.Ops)
		const roofCurveY = 90
		roof.MoveTo(
			f32.Pt(roofTopX-roofWidth, roofTopY+roofHeight))
		roof.QuadTo(
			f32.Pt(roofTopX-roofCurveX, roofTopY+roofCurveY),
			f32.Pt(roofTopX, roofTopY))
		roof.QuadTo(
			f32.Pt(roofTopX+roofCurveX, roofTopY+roofCurveY),
			f32.Pt(roofTopX+roofWidth, roofTopY+roofHeight))
		roof.Close()

		paint.FillShape(gtx.Ops, g.HSL(0, 0.6, 0.4), clip.Outline{Path: roof.End()}.Op())

		return layout.Dimensions{}
	})
}
