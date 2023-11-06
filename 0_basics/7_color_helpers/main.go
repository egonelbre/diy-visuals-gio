// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"gioui.org/layout"   // layout is used for layouting widgets.
	"gioui.org/op/paint" // paint contains operations for coloring.

	"github.com/egonelbre/diy-visuals-gio/internal/g"
	"github.com/egonelbre/diy-visuals-gio/internal/qapp"
)

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		paint.FillShape(gtx.Ops,
			g.HexRGB1(0xf00),
			g.FillRect(gtx.Ops, 100, 100, 50, 50),
		)

		paint.FillShape(gtx.Ops,
			g.HexRGB(0x00ff00),
			g.FillRect(gtx.Ops, 200, 100, 50, 50),
		)

		paint.FillShape(gtx.Ops,
			g.HexRGBA(0x00ff0080),
			g.FillRect(gtx.Ops, 300, 100, 50, 50),
		)

		paint.FillShape(gtx.Ops,
			g.RGB(0, 0.3, 0.8),
			g.FillRect(gtx.Ops, 100, 200, 50, 50),
		)

		paint.FillShape(gtx.Ops,
			g.RGBA(0, 0.3, 0.8, 0.8),
			g.FillRect(gtx.Ops, 200, 200, 50, 50),
		)

		paint.FillShape(gtx.Ops,
			g.HSL(0.5, 0.6, 0.6),
			g.FillRect(gtx.Ops, 100, 300, 50, 50),
		)

		paint.FillShape(gtx.Ops,
			g.HSLA(0.5, 0.7, 0.7, 0.5),
			g.FillRect(gtx.Ops, 200, 300, 50, 50),
		)

		return layout.Dimensions{}
	})
}
