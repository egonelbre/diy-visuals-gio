// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"math"
	"time"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"

	"github.com/egonelbre/diy-visuals-gio/internal/g"
	"github.com/egonelbre/diy-visuals-gio/internal/qapp"
)

var Theme = material.NewTheme()
var Start = time.Now()

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		op.InvalidateOp{}.Add(gtx.Ops)

		millis := float32(gtx.Now.Sub(Start).Seconds()) * 3

		{ // pie line
			var arc clip.Path
			arc.Begin(gtx.Ops)
			arc.MoveTo(f32.Pt(50, 100))
			arc.ArcTo(f32.Pt(100, 100), f32.Pt(100, 100), g.Mod(millis, 2*math.Pi))
			paint.FillShape(gtx.Ops, g.Black, clip.Stroke{
				Path:  arc.End(),
				Width: 4,
			}.Op())
		}

		{ // closed pie
			var pie clip.Path
			pie.Begin(gtx.Ops)
			pie.MoveTo(f32.Pt(100, 250))
			pie.LineTo(f32.Pt(50, 250))
			pie.ArcTo(f32.Pt(100, 250), f32.Pt(100, 250), g.Mod(millis, 2*math.Pi))
			pie.Close()
			paint.FillShape(gtx.Ops, g.Black, clip.Stroke{
				Path:  pie.End(),
				Width: 4,
			}.Op())
		}

		{ // a morphing round shape
			const N = 250
			var shape clip.Path
			shape.Begin(gtx.Ops)
			center := f32.Pt(550, 250)
			for i := 0; i < N; i++ {
				phi := float32(i) * 2 * math.Pi / N

				radius := 150 +
					g.Sin(millis*2+phi*5)*50 +
					0*g.Cos(millis*4+phi*3)*30 +
					0*g.Cos(millis*8+phi*16)*10

				sn, cs := g.Sincos(phi)

				point := f32.Pt(
					center.X+cs*radius,
					center.Y+sn*radius,
				)
				if i == 0 {
					shape.MoveTo(point)
				} else {
					shape.LineTo(point)
				}
			}
			shape.Close()

			path := shape.End()

			paint.FillShape(gtx.Ops, g.HexRGB1(0xCCC),
				clip.Outline{
					Path: path,
				}.Op())

			paint.FillShape(gtx.Ops, g.HexRGB1(0x888),
				clip.Stroke{
					Width: 8,
					Path:  path,
				}.Op())
		}

		{ // a morphing rectangular shape
			const N = 40
			var top, bottom []f32.Point

			for i := 0; i < N; i++ {
				t := float32(i) / N
				p := 100 * t
				top = append(top, f32.Pt(
					p+100,
					-10*g.Sin(t*math.Pi+millis*3)+350-20,
				))
				bottom = append(bottom, f32.Pt(
					p+100,
					10*g.Sin(t*math.Pi+millis*2)+350+20,
				))
			}

			var shape clip.Path
			shape.Begin(gtx.Ops)
			shape.MoveTo(top[0])
			for _, p := range top[1:] {
				shape.LineTo(p)
			}
			for i := len(bottom) - 1; i >= 0; i-- {
				shape.LineTo(bottom[i])
			}
			shape.Close()

			path := shape.End()

			paint.FillShape(gtx.Ops, g.HexRGB1(0xEEE),
				clip.Outline{
					Path: path,
				}.Op())

			paint.FillShape(gtx.Ops, g.HexRGB1(0x666),
				clip.Stroke{
					Width: 8,
					Path:  path,
				}.Op())
		}

		return layout.Dimensions{}
	})
}
