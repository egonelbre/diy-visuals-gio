// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"image/color"

	"gioui.org/layout"   // layout is used for layouting widgets.
	"gioui.org/op/clip"  // clip contains operations for clipping painting area.
	"gioui.org/op/paint" // paint contains operations for coloring.

	"github.com/egonelbre/diy-visuals-gio/internal/qapp" // qapp contains convenience funcs for this tutorial
)

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		// It's possible to restrict the area where to draw.
		clipping := clip.Rect{Max: image.Pt(100, 100)}.Push(gtx.Ops)
		defer clipping.Pop()
		// defer clip.Rect{Min: image.Pt(40, 50), Max: image.Pt(60, 200)}.Push(gtx.Ops).Pop()

		// color the clip area:
		red := color.NRGBA{R: 0xFF, A: 0xFF}
		paint.ColorOp{Color: red}.Add(gtx.Ops)
		paint.PaintOp{}.Add(gtx.Ops)

		return layout.Dimensions{}
	})
}
