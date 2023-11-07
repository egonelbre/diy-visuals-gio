// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"math"
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
var Start time.Time
var LastTime time.Time
var Mouse f32.Point
var Followers []*Ball

func main() {
	target := &Mouse
	for i := 0; i < 10; i++ {
		ball := &Ball{
			Target:   target,
			Position: f32.Pt(100, 100),
			Index:    i,
		}
		target = &ball.Position
		Followers = append(Followers, ball)
	}

	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		screenSize := layout.FPt(gtx.Constraints.Max)
		op.InvalidateOp{}.Add(gtx.Ops)

		log := g.Log{
			Theme: Theme,
			Size:  14,
		}
		defer log.Layout(gtx)

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

		if LastTime.IsZero() {
			Start = gtx.Now
			LastTime = gtx.Now
		}
		deltaMillis := float32(gtx.Now.Sub(LastTime).Seconds() * 1000)
		log.Printf("dt: %v", deltaMillis)
		secs := float32(gtx.Now.Sub(Start).Seconds())
		log.Printf(" t: %v", secs)
		log.Printf(".t: %v", g.Mod(secs, 1))

		LastTime = gtx.Now

		log.Printf("m: %v", Mouse)
		paint.FillShape(gtx.Ops,
			g.HSLA(0, 0.5, 0.5, 1.0),
			g.FillCircle(gtx.Ops, Mouse, ballRadius),
		)

		start := screenSize.Div(2)
		paint.FillShape(gtx.Ops,
			g.Black,
			g.FillCircle(gtx.Ops, start, ballRadius),
		)
		middle := g.PtLerp(g.Mod(secs, 1), start, Mouse)
		paint.FillShape(gtx.Ops,
			g.Blue,
			g.FillCircle(gtx.Ops, middle, ballRadius),
		)

		for i, ball := range Followers {
			log.Printf("%d: %v", i, ball.Position)
			ball.Update(deltaMillis)
			ball.Layout(gtx)
		}

		return layout.Dimensions{}
	})
}

type Ball struct {
	Target   *f32.Point
	Position f32.Point
	Time     float32
	Index    int
}

const ballRadius = 20

func (ball *Ball) Update(deltaTime float32) {
	ball.Time += deltaTime

	if g.Len(ball.Position.Sub(*ball.Target)) < 2*ballRadius {
		return
	}

	ball.Position = g.PtLerp(0.1, ball.Position, *ball.Target)
}

func (ball *Ball) Layout(gtx layout.Context) {
	hue := float32(ball.Index) * math.Phi
	paint.FillShape(gtx.Ops,
		g.HSLA(hue, 0.7, 0.7, 0.9),
		g.FillCircle(gtx.Ops, ball.Position, ballRadius),
	)
}
