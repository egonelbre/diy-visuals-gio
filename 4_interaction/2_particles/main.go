// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
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

var Particles [300]Particle

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		screenSize := layout.FPt(gtx.Constraints.Max)
		op.InvalidateOp{}.Add(gtx.Ops)
		paint.Fill(gtx.Ops, g.Black)

		for i := range Particles {
			p := &Particles[i]
			if !p.Alive {
				Particles[i] = NewParticle(
					f32.Pt(screenSize.X/2, 0),
					f32.Pt(
						g.RandRange(-200, 200),
						g.RandRange(0, 500),
					),
				)
				break
			}
		}

		if Last.IsZero() {
			Last = gtx.Now
		}

		deltaTime := float32(gtx.Now.Sub(Last).Seconds())
		Last = gtx.Now

		wind := g.Sin(float32(gtx.Now.Second())) * 10
		globalForce := f32.Pt(
			wind,
			-100,
		)

		for i := range Particles {
			p := &Particles[i]
			p.Force = globalForce
		}

		for i := range Particles {
			p := &Particles[i]
			p.Update(deltaTime)
		}

		for i := range Particles {
			p := &Particles[i]
			p.Draw(gtx)
		}

		return layout.Dimensions{}
	})
}

type Particle struct {
	Hue   float32
	At    f32.Point
	Speed f32.Point
	Force f32.Point
	Alive bool
}

func NewParticle(at, speed f32.Point) Particle {
	return Particle{
		Hue:   g.Rand(),
		At:    at,
		Speed: speed,
		Alive: true,
	}
}

func (p *Particle) Update(deltaTime float32) {
	p.Speed = p.Speed.Add(p.Force.Mul(deltaTime))
	p.At = p.At.Add(p.Speed.Mul(deltaTime))
	if p.At.Y < 0 {
		p.Alive = false
	}
}

func (p *Particle) Draw(gtx layout.Context) {
	paint.FillShape(gtx.Ops,
		g.HSLA(p.Hue, 0.7, 0.7, 0.5),
		g.FillCircle(gtx.Ops, p.At, 20),
	)
}
