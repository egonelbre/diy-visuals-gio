// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"image/color"

	"gioui.org/f32"      // f32 contains float32 points.
	"gioui.org/layout"   // op is used for recording different operations.
	"gioui.org/op/clip"  // clip contains operations for clipping painting area.
	"gioui.org/op/paint" // paint contains operations for coloring.

	"github.com/egonelbre/diy-visuals-gio/internal/qapp"
)

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		black := color.NRGBA{A: 0xFF}

		paint.FillShape(gtx.Ops, black,
			clip.Rect{
				Min: image.Pt(50, 50),
				Max: image.Pt(75, 75),
			}.Op(),
		)

		paint.FillShape(gtx.Ops, black,
			clip.Rect{
				Min: image.Pt(50, 100),
				Max: image.Pt(100, 125),
			}.Op(),
		)

		{
			var p clip.Path
			p.Begin(gtx.Ops)
			p.MoveTo(f32.Pt(50, 150))
			p.LineTo(f32.Pt(100, 200))
			p.LineTo(f32.Pt(150, 160))
			p.LineTo(f32.Pt(200, 150))

			paint.FillShape(gtx.Ops, black,
				clip.Stroke{
					Width: 3,
					Path:  p.End(),
				}.Op())
		}

		{
			var p clip.Path
			p.Begin(gtx.Ops)
			p.MoveTo(f32.Pt(50, 250))
			p.LineTo(f32.Pt(100, 270))
			p.LineTo(f32.Pt(150, 220))
			p.LineTo(f32.Pt(200, 260))
			p.Close()

			paint.FillShape(gtx.Ops, black,
				clip.Outline{
					Path: p.End(),
				}.Op())
		}

		return layout.Dimensions{}
	})
}
