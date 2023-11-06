// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"image/color"

	"gioui.org/f32"      // f32 contains float32 points.
	"gioui.org/layout"   // op is used for recording different operations.
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

		paint.FillShape(gtx.Ops,
			color.NRGBA{R: 0xFF, A: 0xFF},
			clip.Stroke{
				Width: 4,
				Path:  triangle.End(),
			}.Op(),
		)

		paint.FillShape(gtx.Ops,
			color.NRGBA{G: 0xFF, A: 0xFF},
			clip.Stroke{
				Width: 4,
				Path:  clip.Rect{Max: image.Pt(100, 100)}.Path(),
			}.Op(),
		)
		return layout.Dimensions{}
	})
}
