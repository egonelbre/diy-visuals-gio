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
var Start time.Time

var Particles [50]Particle

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		screenSize := layout.FPt(gtx.Constraints.Max)
		op.InvalidateOp{}.Add(gtx.Ops)
		paint.Fill(gtx.Ops, g.Black)

		for i := range Particles {
			p := &Particles[i]
			if !p.Alive {
				Particles[i] = NewParticle(
					f32.Pt(
						g.Rand()*screenSize.X,
						g.Rand()*screenSize.Y,
					),
					f32.Pt(
						g.Rand()-0.5,
						g.Rand()-0.5,
					).Mul(200),
				)
				break
			}
		}

		if Last.IsZero() {
			Last = gtx.Now
			Start = gtx.Now
		}

		now := float32(gtx.Now.Sub(Start).Seconds())
		deltaTime := float32(gtx.Now.Sub(Last).Seconds())
		Last = gtx.Now

		target := screenSize.Mul(0.5)
		target = target.Add(
			f32.Pt(g.Sincos(now * 0.1)).Mul(screenSize.Y / 4),
		)

		paint.FillShape(gtx.Ops,
			g.HSLA(0.4, 0.5, 0.8, 0.9),
			g.FillCircle(gtx.Ops, target, 30))

		for i := range Particles {
			p := &Particles[i]

			// attract to a central point
			delta := target.Sub(p.At).Mul(0.3)
			p.Force = delta
		}

		for i := range Particles {
			a := &Particles[i]
			for k := range Particles[i+1:] {
				b := &Particles[k]

				delta := b.At.Sub(a.At)
				length := g.Len(delta)
				if length < screenSize.Y*0.2 {
					force := delta.Mul(1 / (length + 1))
					a.Force = a.Force.Add(force)
				}
			}
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
}

func (p *Particle) Draw(gtx layout.Context) {
	paint.FillShape(gtx.Ops,
		g.HSLA(p.Hue, 0.7, 0.7, 0.5),
		g.FillCircle(gtx.Ops, p.At, 20),
	)
}
