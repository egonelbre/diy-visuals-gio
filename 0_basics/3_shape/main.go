// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image/color"

	"gioui.org/f32"      // f32 contains float32 points.
	"gioui.org/layout"   // layout is used for layouting widgets.
	"gioui.org/op/clip"  // clip contains operations for clipping painting area.
	"gioui.org/op/paint" // paint contains operations for coloring.

	"github.com/egonelbre/diy-visuals-gio/internal/qapp" // qapp contains convenience funcs for this tutorial
)

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		var triangle clip.Path
		triangle.Begin(gtx.Ops)
		triangle.MoveTo(f32.Pt(30, 30))
		triangle.LineTo(f32.Pt(170, 170))
		triangle.LineTo(f32.Pt(80, 170))
		triangle.Close()

		// set the clip to the outline
		defer clip.Outline{
			Path: triangle.End(),
		}.Op().Push(gtx.Ops).Pop()

		// color the clip area:
		black := color.NRGBA{A: 0xFF}
		paint.ColorOp{Color: black}.Add(gtx.Ops)
		paint.PaintOp{}.Add(gtx.Ops)

		/*
			defer clip.Rect{Max: image.Pt(100, 100)}.Push(gtx.Ops)
			green := color.NRGBA{G: 0xFF, A: 0xFF}
			paint.ColorOp{Color: green}.Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)
		*/

		return layout.Dimensions{}
	})
}
