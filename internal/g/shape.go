// SPDX-License-Identifier: Unlicense OR MIT

package g

import (
	"math"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/widget/material"
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
		Path: Rect(ops, x, y, w, h),
	}.Op()
}

func FillText(th *material.Theme, gtx layout.Context, s string, at f32.Point) {
	Offset{
		Pos: at,
	}.Layout(gtx, material.Body1(th, s).Layout)
}

func Circle(ops *op.Ops, at f32.Point, r float32) clip.PathSpec {
	var p clip.Path
	p.Begin(ops)
	p.MoveTo(at.Add(f32.Pt(r, 0)))
	p.ArcTo(at, at, 2*math.Pi)
	p.Close()
	return p.End()
}

func FillCircle(ops *op.Ops, at f32.Point, r float32) clip.Op {
	return clip.Outline{
		Path: Circle(ops, at, r),
	}.Op()
}
