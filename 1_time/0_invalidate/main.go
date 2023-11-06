// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"

	"github.com/egonelbre/diy-visuals-gio/internal/qapp"
)

var Theme = material.NewTheme()

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		op.InvalidateOp{}.Add(gtx.Ops)
		material.Body1(Theme, gtx.Now.String()).Layout(gtx)
		return layout.Dimensions{}
	})
}
