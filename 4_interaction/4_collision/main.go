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
var Last time.Time
var Start time.Time

var Balls []*Ball

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		screenSize := layout.FPt(gtx.Constraints.Max)
		op.InvalidateOp{}.Add(gtx.Ops)
		paint.Fill(gtx.Ops, g.Black)

		for len(Balls) < 30 {
			Balls = append(Balls, NewBall(
				f32.Pt(
					g.Rand()*screenSize.X,
					g.Rand()*screenSize.Y,
				),
			))
		}

		if Last.IsZero() {
			Last = gtx.Now
			Start = gtx.Now
		}

		deltaTime := float32(gtx.Now.Sub(Last).Seconds())
		Last = gtx.Now

		for _, b := range Balls {
			b.Force = f32.Point{}
		}

		for i, a := range Balls {
			for _, b := range Balls[i+1:] {
				delta := a.At.Sub(b.At)
				distance := g.Len(delta)
				if distance > a.Radius+b.Radius {
					continue
				}

				pen := g.Unit(delta).Mul(a.Radius + b.Radius - distance)

				ra := a.Mass / (a.Mass + b.Mass)
				rb := b.Mass / (a.Mass + b.Mass)

				a.At = a.At.Add(pen.Mul(rb))
				b.At = b.At.Sub(pen.Mul(ra))

				unit := g.Unit(delta)
				p := 2 * (g.Dot(a.Velocity, unit) - g.Dot(b.Velocity, unit)) / (a.Mass + b.Mass)

				a.Velocity = a.Velocity.Sub(unit.Mul(p * a.Mass))
				b.Velocity = b.Velocity.Add(unit.Mul(p * b.Mass))
			}
		}

		for _, b := range Balls {
			b.Update(deltaTime, screenSize)
		}

		for _, b := range Balls {
			b.Draw(gtx)
		}

		return layout.Dimensions{}
	})
}

type Ball struct {
	Hue      float32
	Radius   float32
	Mass     float32
	At       f32.Point
	Velocity f32.Point
	Force    f32.Point

	Rotation     float32
	AngularSpeed float32
}

func NewBall(at f32.Point) *Ball {
	radius := g.RandRange(20, 60)
	return &Ball{
		Hue:      g.Rand(),
		At:       at,
		Radius:   radius,
		Mass:     math.Pi * radius * radius,
		Velocity: g.RandUnitPt().Mul(400),

		Rotation:     g.RandRange(0, 2*math.Pi),
		AngularSpeed: g.RandRange(-0.3, 0.3),
	}
}

func (b *Ball) Update(deltaTime float32, screen f32.Point) {
	if b.At.X-b.Radius < 0 {
		b.Velocity.X = g.Abs(b.Velocity.X)
		b.At.X = b.Radius
		b.AngularSpeed += g.Sign(b.Velocity.Y) * 5
	}
	if b.At.X+b.Radius > screen.X {
		b.Velocity.X = -g.Abs(b.Velocity.X)
		b.At.X = screen.X - b.Radius
		b.AngularSpeed -= g.Sign(b.Velocity.Y) * 5
	}
	if b.At.Y-b.Radius < 0 {
		b.Velocity.Y = g.Abs(b.Velocity.Y)
		b.At.Y = b.Radius
		b.AngularSpeed -= g.Sign(b.Velocity.X) * 5
	}
	if b.At.Y+b.Radius > screen.Y {
		b.Velocity.Y = -g.Abs(b.Velocity.Y)
		b.At.Y = screen.Y - b.Radius
		b.AngularSpeed += g.Sign(b.Velocity.X) * 5
	}

	b.AngularSpeed *= 0.99
	b.Rotation += b.AngularSpeed * deltaTime
	b.Velocity = b.Velocity.Add(b.Force.Mul(deltaTime))
	b.At = b.At.Add(b.Velocity.Mul(deltaTime))
}

func (b *Ball) Draw(gtx layout.Context) {
	defer op.Affine(f32.Affine2D{}.
		Rotate(f32.Point{}, b.Rotation).
		Offset(b.At),
	).Push(gtx.Ops).Pop()

	paint.FillShape(gtx.Ops,
		g.HSLA(b.Hue, 0.7, 0.7, 0.5),
		g.FillCircle(gtx.Ops, f32.Point{}, b.Radius),
	)

	paint.FillShape(gtx.Ops, g.White,
		g.FillCircle(gtx.Ops, f32.Point{
			X: b.Radius * 0.4,
			Y: b.Radius * -0.3,
		}, b.Radius*0.2),
	)
	paint.FillShape(gtx.Ops, g.White,
		g.FillCircle(gtx.Ops, f32.Point{
			X: -b.Radius * 0.4,
			Y: b.Radius * -0.3,
		}, b.Radius*0.2),
	)
}
