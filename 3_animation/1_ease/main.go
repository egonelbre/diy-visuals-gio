// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"math"
	"time"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"

	"github.com/egonelbre/diy-visuals-gio/internal/ease"
	"github.com/egonelbre/diy-visuals-gio/internal/g"
	"github.com/egonelbre/diy-visuals-gio/internal/qapp"
)

var Theme = material.NewTheme()
var Start time.Time

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		screenSize := layout.FPt(gtx.Constraints.Max)
		op.InvalidateOp{}.Add(gtx.Ops)

		height := float32(int(screenSize.Y / float32(len(ease.Funcs))))

		if Start.IsZero() {
			Start = gtx.Now
		}

		now := float32(gtx.Now.Sub(Start).Seconds())
		animationTime := g.Mod(now, 1)

		for i, fn := range ease.Funcs {
			func() {
				defer op.Offset(image.Point{Y: int(height)}).Push(gtx.Ops)

				defer clip.Rect{
					Max: image.Point{
						X: int(screenSize.X),
						Y: int(height),
					},
				}.Push(gtx.Ops).Pop()

				// separator line
				paint.FillShape(gtx.Ops, g.RGBA(0, 0, 0, 0.4),
					g.FillRect(gtx.Ops, 0, 0, screenSize.X, 3))

				easedTime := fn.Ease(animationTime)
				color := g.HSL(float32(i)*math.Phi, 0.5, 0.5)

				left := height / 2
				right := screenSize.X - height/2

				drawSquare := func() {
					halfsize := height / 3
					paint.FillShape(gtx.Ops, color,
						g.FillRect(gtx.Ops, -halfsize, -halfsize, 2*halfsize, 2*halfsize))
				}

				// rotating a square using easing 360deg
				func() {
					rotate := g.Lerp(easedTime, 0, 2*math.Pi)
					defer op.Affine(
						f32.Affine2D{}.
							Rotate(f32.Point{}, rotate).
							Offset(f32.Pt(left, height/2)),
					).Push(gtx.Ops).Pop()
					left += height
					drawSquare()
				}()

				// rotating a square using easing but 90deg
				func() {
					rotate := g.Lerp(easedTime, 0, math.Pi/2)
					defer op.Affine(
						f32.Affine2D{}.
							Rotate(f32.Point{}, rotate).
							Offset(f32.Pt(left, height/2)),
					).Push(gtx.Ops).Pop()
					left += height
					drawSquare()
				}()

				// rotating a square using easing but 90deg and cale
				func() {
					rotate := g.Lerp(easedTime, 0, math.Pi/2)
					scale := g.Lerp(easedTime, 0.8, 1.5)

					defer op.Affine(
						f32.Affine2D{}.
							Rotate(f32.Point{}, rotate).
							Scale(f32.Point{}, f32.Pt(scale, scale)).
							Offset(f32.Pt(left, height/2)),
					).Push(gtx.Ops).Pop()
					left += height
					drawSquare()
				}()

				// draw the label
				g.Offset{
					Pos: f32.Pt(left, 0),
				}.Layout(gtx, material.Body1(Theme, fn.Name).Layout)

				// draw a plot
				func() {
					defer op.Affine(
						f32.Affine2D{}.
							Offset(f32.Pt(right-height/2, 0)),
					).Push(gtx.Ops).Pop()
					right -= height

					gray := float32(0.9)
					if i%2 == 0 {
						gray = 0.8
					}
					paint.FillShape(gtx.Ops,
						g.RGBA(gray, gray, gray, 1),
						g.FillRect(gtx.Ops, 0, 0, height, height),
					)

					var path clip.Path
					path.Begin(gtx.Ops)
					path.Move(f32.Pt(0, 0))
					for x := float32(1); x < height; x++ {
						easedY := fn.Ease(x / height)
						y := easedY * height
						path.LineTo(f32.Pt(x, y))
					}

					paint.FillShape(gtx.Ops, g.Black,
						clip.Stroke{
							Width: 2,
							Path:  path.End(),
						}.Op())

					at := f32.Pt(
						animationTime*height,
						easedTime*height,
					)

					paint.FillShape(gtx.Ops,
						color,
						g.FillCircle(gtx.Ops, at, height/8),
					)
				}()

				// draw the point movement animation
				func() {
					start := f32.Pt(left, height/2)
					end := f32.Pt(right, height/2)
					at := g.PtLerp(easedTime, start, end)

					paint.FillShape(gtx.Ops,
						g.RGBA(0.9, 0.9, 0.9, 1),
						g.FillRect(gtx.Ops, left, height/2-5, right-left, 10),
					)

					paint.FillShape(gtx.Ops,
						color,
						g.FillCircle(gtx.Ops, at, height/3),
					)
				}()

				_ = easedTime
			}()
		}

		return layout.Dimensions{}
	}, app.Size(1024, 512))
}
