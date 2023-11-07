// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"time"

	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/widget/material"

	"github.com/egonelbre/diy-visuals-gio/internal/g"
	"github.com/egonelbre/diy-visuals-gio/internal/qapp"
)

var Theme = material.NewTheme()
var Start = time.Now()
var Mouse f32.Point

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		screenSize := layout.FPt(gtx.Constraints.Max)
		op.InvalidateOp{}.Add(gtx.Ops)

		pointer.InputOp{
			Tag:   &Mouse,
			Types: pointer.Press | pointer.Drag | pointer.Release | pointer.Move,
		}.Add(gtx.Ops)
		for _, ev := range gtx.Events(&Mouse) {
			switch ev := ev.(type) {
			case pointer.Event:
				Mouse = ev.Position
			}
		}

		millis := float32(gtx.Now.Sub(Start).Seconds())
		_ = millis

		defer op.Affine(
			f32.Affine2D{}.
				Scale(f32.Point{}, f32.Pt(1, -1)).
				Offset(f32.Pt(
					screenSize.X/2,
					screenSize.Y,
				)),
		).Push(gtx.Ops).Pop()

		relativeMouse := Mouse
		relativeMouse.X /= screenSize.X
		relativeMouse.Y /= screenSize.Y
		relativeMouse = relativeMouse.Sub(f32.Pt(0.5, 0.5)).Mul(2)

		Tree(gtx, 7, relativeMouse.X, relativeMouse.Y)

		return layout.Dimensions{}
	})
}

func Tree(gtx layout.Context, level int, rotateLeft, rotateRight float32) {
	level--
	if level < 0 {
		return
	}

	height := float32(400)

	paint.FillShape(gtx.Ops, g.Black,
		g.FillRect(gtx.Ops, -10, 0, 20, height))

	defer op.Affine(
		f32.Affine2D{}.
			Scale(f32.Point{}, f32.Pt(0.7, 0.7)).
			Offset(f32.Pt(0, height)),
	).Push(gtx.Ops).Pop()

	func() {
		defer op.Affine(
			f32.Affine2D{}.Rotate(f32.Pt(0, 0), -rotateLeft),
		).Push(gtx.Ops).Pop()

		Tree(gtx, level, rotateLeft, rotateRight)
	}()

	func() {
		defer op.Affine(
			f32.Affine2D{}.Rotate(f32.Pt(0, 0), rotateRight),
		).Push(gtx.Ops).Pop()

		Tree(gtx, level, rotateLeft, rotateRight)
	}()
}
