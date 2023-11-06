// SPDX-License-Identifier: Unlicense OR MIT

package g

import (
	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
)

func Rect(ops *op.Ops, x, y, w, h float32) clip.PathSpec {
	var p clip.Path
	p.Begin(ops)
	p.MoveTo(f32.Pt(x, y))
	p.LineTo(f32.Pt(x+w, y))
	p.LineTo(f32.Pt(x+w, y+h))
	p.LineTo(f32.Pt(x, y+h))
	p.Close()
	return p.End()
}

func FillRect(ops *op.Ops, x, y, w, h float32) clip.Op {
	return clip.Outline{
		Path: Rect(ops, x,y,w,h),
	}.Op()
}