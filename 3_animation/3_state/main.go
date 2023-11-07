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

	"github.com/egonelbre/diy-visuals-gio/internal/ease"
	"github.com/egonelbre/diy-visuals-gio/internal/g"
	"github.com/egonelbre/diy-visuals-gio/internal/qapp"
)

var Theme = material.NewTheme()
var Last time.Time
var Critters []*Critter

type Critter struct {
	Start f32.Point
	Food  f32.Point

	Animation float32
	Duration  float32

	Hue float32

	State State
}

type State byte

const (
	Moving = State(iota)
	Eating
)

func NewCritter() *Critter {
	return &Critter{
		Start: g.RandPt(),
		Food:  g.RandPt(),

		Animation: 0,
		Duration:  g.RandRange(1, 2),
		Hue:       g.Rand(),
		State:     Moving,
	}
}

func (c *Critter) Update(deltaTime float32) {
	c.Animation += deltaTime / c.Duration
	if c.Animation > 1 {
		c.Animation = 0

		switch c.State {
		case Moving:
			c.State = Eating
			c.Start = c.Food
			c.Duration = g.RandRange(1, 3)
		case Eating:
			c.State = Moving
			c.Food = g.RandPt()
			c.Duration = g.RandRange(1, 2)
		}
	}
}

func (c *Critter) Draw(gtx layout.Context, screenSize f32.Point) {
	start := f32.Pt(
		c.Start.X*screenSize.X,
		c.Start.Y*screenSize.Y,
	)
	food := f32.Pt(
		c.Food.X*screenSize.X,
		c.Food.Y*screenSize.Y,
	)

	switch c.State {
	case Moving:
		t := ease.InOutQuad(c.Animation)
		at := g.PtLerp(t, start, food)
		paint.FillShape(gtx.Ops,
			g.HSLA(c.Hue, 0.7, 0.3, 0.95),
			g.FillCircle(gtx.Ops, at, 20))
	case Eating:
		t := g.Sin(c.Animation * math.Pi * 8)
		paint.FillShape(gtx.Ops,
			g.HSLA(c.Hue, 0.7, 0.3, 0.95),
			g.FillCircle(gtx.Ops, food, 20+t*5))
	}

	paint.FillShape(gtx.Ops,
		g.HSLA(0, 0.8, 0.6, 0.95),
		g.FillCircle(gtx.Ops, food, 10))
}

func main() {
	for i := 0; i < 4; i++ {
		Critters = append(Critters, NewCritter())
	}

	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		screenSize := layout.FPt(gtx.Constraints.Max)
		op.InvalidateOp{}.Add(gtx.Ops)

		if Last.IsZero() {
			Last = gtx.Now
		}

		deltaTime := float32(gtx.Now.Sub(Last).Seconds())
		Last = gtx.Now

		for _, c := range Critters {
			c.Update(deltaTime)
			c.Draw(gtx, screenSize)
		}

		return layout.Dimensions{}
	})
}
