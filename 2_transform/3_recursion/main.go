// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"math"
	"time"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/widget/material"

	"github.com/egonelbre/diy-visuals-gio/internal/g"
	"github.com/egonelbre/diy-visuals-gio/internal/qapp"
)

var Theme = material.NewTheme()
var Start = time.Now()

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		screenSize := layout.FPt(gtx.Constraints.Max)
		op.InvalidateOp{}.Add(gtx.Ops)

		millis := float32(gtx.Now.Sub(Start).Seconds())

		defer op.Affine(
			f32.Affine2D{}.
				Scale(f32.Point{}, f32.Pt(1, -1)).
				Offset(f32.Pt(
					screenSize.X/2,
					screenSize.Y,
				)),
		).Push(gtx.Ops).Pop()

		//millis = 0.3
		Tree(gtx, 4, millis)

		return layout.Dimensions{}
	})
}

func Tree(gtx layout.Context, level int, time float32) {
	level--
	if level < 0 {
		return
	}

	height := float32(400)

	paint.FillShape(gtx.Ops, g.Black,
		g.FillRect(gtx.Ops, -5, 0, 10, height))

	defer op.Affine(
		f32.Affine2D{}.
			Scale(f32.Point{}, f32.Pt(0.7, 0.7)).
			Offset(f32.Pt(0, height)),
	).Push(gtx.Ops).Pop()

	rotateLeft := g.Sin(time*3 + 2*math.Pi*0.1)
	rotateRight := g.Sin(time*4 - 2*math.Pi*0.1)

	func() {
		defer op.Affine(
			f32.Affine2D{}.Rotate(f32.Pt(0, 0), rotateLeft),
		).Push(gtx.Ops).Pop()

		Tree(gtx, level, time)
	}()

	func() {
		defer op.Affine(
			f32.Affine2D{}.Rotate(f32.Pt(0, 0), rotateRight),
		).Push(gtx.Ops).Pop()

		Tree(gtx, level, time)
	}()
}
