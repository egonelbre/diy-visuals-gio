// SPDX-License-Identifier: Unlicense OR MIT

package g

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
)

type Offset struct {
	Pos f32.Point
}

func (at Offset) Layout(gtx layout.Context, w layout.Widget) layout.Dimensions {
	var x f32.Affine2D
	x = x.Offset(at.Pos)
	defer op.Affine(x).Push(gtx.Ops).Pop()
	return w(gtx)
}
